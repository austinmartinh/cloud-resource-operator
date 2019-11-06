package e2e

import (
	"testing"
	"time"

	"github.com/integr8ly/cloud-resource-operator/pkg/apis"
	"github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
)

const (
	postgresName = "example-postgres"
	redisName    = "example-redis"
)

var (
	retryInterval = time.Second * 20
	timeout       = time.Minute * 5
)

func TestCRO(t *testing.T) {
	// adding redis scheme to framework
	redisList := &v1alpha1.Redis{}
	if err := framework.AddToFrameworkScheme(apis.AddToScheme, redisList); err != nil {
		t.Fatalf("failed to add Redis custom resource scheme to framework: %v", err)
	}

	// adding postgres scheme to framework
	postgresList := &v1alpha1.Postgres{}
	if err := framework.AddToFrameworkScheme(apis.AddToScheme, postgresList); err != nil {
		t.Fatalf("failed to add Postgres custom resource scheme to framework: %v", err)
	}

	// run subtests
	t.Run("cro-test", func(t *testing.T) {
		t.Run("Cluster", BasicTestCluster)
	})

}

func BasicTestCluster(t *testing.T) {
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()
	err := ctx.InitializeClusterResources(getCleanupOptions(t))
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}
	t.Log("initialized cluster resources")

	// get global framework variables
	f := framework.Global

	// run postgres test
	if err = OpenshiftPostgresBasicTest(t, f, *ctx); err != nil {
		t.Fatal(err)
	}

	// run redis test
	if err = OpenshiftRedisBasicTest(t, f, *ctx); err != nil {
		t.Fatal(err)
	}

	// run postgres permission test
	if err = OpenshiftVerifyPostgresTest(t, f, *ctx); err != nil {
		t.Fatal(err)
	}

	// run postgres deployment recover test
	if err = OpenshiftVerifyPostgresDeploymentRecovery(t, f, *ctx); err != nil {
		t.Fatal(err)
	}

	// run postgres service recover test
	if err = OpenshiftVerifyPostgresServiceRecovery(t, f, *ctx); err != nil {
		t.Fatal(err)
	}

	// run postgres pvc recover test
	if err = OpenshiftVerifyPostgresPVCRecovery(t, f, *ctx); err != nil {
		t.Fatal(err)
	}

	// run postgres deployment update recover test
	if err = OpenshiftVerifyPostgresDeploymentUpdate(t, f, *ctx); err != nil {
		t.Fatal(err)
	}
}

// returns cleanup options
func getCleanupOptions(t *testing.T) *framework.CleanupOptions {
	return &framework.CleanupOptions{
		TestContext:   framework.NewTestCtx(t),
		Timeout:       timeout,
		RetryInterval: retryInterval,
	}
}
