package resources

import (
	"context"
	"errors"
	"github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1/types"
	"reflect"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/integr8ly/cloud-resource-operator/pkg/apis"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func buildTestScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	err := apis.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}
	return scheme, nil
}

func TestReconcileBlobStorage(t *testing.T) {
	scheme, err := buildTestScheme()
	if err != nil {
		t.Fatal("failed to build scheme", err)
	}

	type args struct {
		ctx            context.Context
		client         client.Client
		deploymentType string
		tier           string
		name           string
		ns             string
		secretName     string
		secretNs       string
		modifyFunc     modifyResourceFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *v1alpha1.BlobStorage
		wantErr bool
	}{
		{
			name: "test successful creation",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc:     nil,
			},
			want: &v1alpha1.BlobStorage{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"productName": "test",
					},
				},
				Spec: v1alpha1.BlobStorageSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					cr.SetLabels(map[string]string{
						"cro": "test",
					})
					return nil
				},
			},
			want: &v1alpha1.BlobStorage{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"cro": "test",
					},
				},
				Spec: v1alpha1.BlobStorageSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function error",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "workshop",
				tier:           "development",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					return errors.New("error executing function")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReconcileBlobStorage(tt.args.ctx, tt.args.client, tt.args.deploymentType, tt.args.tier, tt.args.name, tt.args.ns, tt.args.secretName, tt.args.secretNs, tt.args.modifyFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconcileBlobStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcileBlobStorage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReconcileSMTPCredentialSet(t *testing.T) {
	scheme, err := buildTestScheme()
	if err != nil {
		t.Fatal("failed to build scheme", err)
	}

	type args struct {
		ctx            context.Context
		client         client.Client
		deploymentType string
		tier           string
		name           string
		ns             string
		secretName     string
		secretNs       string
		modifyFunc     modifyResourceFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *v1alpha1.SMTPCredentialSet
		wantErr bool
	}{
		{
			name: "test successful creation",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc:     nil,
			},
			want: &v1alpha1.SMTPCredentialSet{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"productName": "test",
					},
				},
				Spec: v1alpha1.SMTPCredentialSetSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					cr.SetLabels(map[string]string{
						"cro": "test",
					})
					return nil
				},
			},
			want: &v1alpha1.SMTPCredentialSet{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"cro": "test",
					},
				},
				Spec: v1alpha1.SMTPCredentialSetSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function error",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "workshop",
				tier:           "development",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					return errors.New("error executing function")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReconcileSMTPCredentialSet(tt.args.ctx, tt.args.client, tt.args.deploymentType, tt.args.tier, tt.args.name, tt.args.ns, tt.args.secretName, tt.args.secretNs, tt.args.modifyFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconcileSMTPCredentialSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcileSMTPCredentialSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReconcilePostgres(t *testing.T) {
	scheme, err := buildTestScheme()
	if err != nil {
		t.Fatal("failed to build scheme", err)
	}

	type args struct {
		ctx            context.Context
		client         client.Client
		deploymentType string
		tier           string
		name           string
		ns             string
		secretName     string
		secretNs       string
		modifyFunc     modifyResourceFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *v1alpha1.Postgres
		wantErr bool
	}{
		{
			name: "test successful creation",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc:     nil,
			},
			want: &v1alpha1.Postgres{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"productName": "test",
					},
				},
				Spec: v1alpha1.PostgresSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					cr.SetLabels(map[string]string{
						"cro": "test",
					})
					return nil
				},
			},
			want: &v1alpha1.Postgres{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"cro": "test",
					},
				},
				Spec: v1alpha1.PostgresSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function error",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "workshop",
				tier:           "development",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					return errors.New("error executing function")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReconcilePostgres(tt.args.ctx, tt.args.client, tt.args.deploymentType, tt.args.tier, tt.args.name, tt.args.ns, tt.args.secretName, tt.args.secretNs, tt.args.modifyFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconcilePostgres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcilePostgres() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReconcileRedis(t *testing.T) {
	scheme, err := buildTestScheme()
	if err != nil {
		t.Fatal("failed to build scheme", err)
	}
	type args struct {
		ctx            context.Context
		client         client.Client
		deploymentType string
		tier           string
		name           string
		ns             string
		secretName     string
		secretNs       string
		modifyFunc     modifyResourceFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *v1alpha1.Redis
		wantErr bool
	}{
		{
			name: "test successful creation",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc:     nil,
			},
			want: &v1alpha1.Redis{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"productName": "test",
					},
				},
				Spec: v1alpha1.RedisSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "managed",
				tier:           "production",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					cr.SetLabels(map[string]string{
						"cro": "test",
					})
					return nil
				},
			},
			want: &v1alpha1.Redis{
				ObjectMeta: v1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
					Labels: map[string]string{
						"cro": "test",
					},
				},
				Spec: v1alpha1.RedisSpec{
					Type: "managed",
					Tier: "production",
					SecretRef: &types.SecretRef{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test modification function error",
			args: args{
				ctx:            context.TODO(),
				client:         fake.NewFakeClientWithScheme(scheme),
				deploymentType: "workshop",
				tier:           "development",
				name:           "test",
				ns:             "test",
				secretName:     "test",
				secretNs:       "test",
				modifyFunc: func(cr v1.Object) error {
					return errors.New("error executing function")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReconcileRedis(tt.args.ctx, tt.args.client, tt.args.deploymentType, tt.args.tier, tt.args.name, tt.args.ns, tt.args.secretName, tt.args.secretNs, tt.args.modifyFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReconcileRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcileRedis() got = %v, want %v", got, tt.want)
			}
		})
	}
}
