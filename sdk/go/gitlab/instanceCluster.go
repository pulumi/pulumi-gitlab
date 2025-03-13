// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `InstanceCluster` resource allows to manage the lifecycle of an instance cluster.
//
// > This is deprecated GitLab feature since 14.5
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/instance_clusters/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewInstanceCluster(ctx, "bar", &gitlab.InstanceClusterArgs{
//				Name:                        pulumi.String("bar-cluster"),
//				Domain:                      pulumi.String("example.com"),
//				Enabled:                     pulumi.Bool(true),
//				KubernetesApiUrl:            pulumi.String("https://124.124.124"),
//				KubernetesToken:             pulumi.String("some-token"),
//				KubernetesCaCert:            pulumi.String("some-cert"),
//				KubernetesNamespace:         pulumi.String("namespace"),
//				KubernetesAuthorizationType: pulumi.String("rbac"),
//				EnvironmentScope:            pulumi.String("*"),
//				ManagementProjectId:         pulumi.String("123456"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_instance_cluster`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_instance_cluster.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// GitLab instance clusters can be imported using a `clusterid`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/instanceCluster:InstanceCluster bar 123
// ```
type InstanceCluster struct {
	pulumi.CustomResourceState

	// Cluster type.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// Create time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The base domain of the cluster.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope pulumi.StringPtrOutput `pulumi:"environmentScope"`
	// The URL to access the Kubernetes API.
	KubernetesApiUrl pulumi.StringOutput `pulumi:"kubernetesApiUrl"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType pulumi.StringPtrOutput `pulumi:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert pulumi.StringPtrOutput `pulumi:"kubernetesCaCert"`
	// The unique namespace related to the instance.
	KubernetesNamespace pulumi.StringPtrOutput `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes. This attribute cannot be read.
	KubernetesToken pulumi.StringOutput `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrOutput `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrOutput `pulumi:"managementProjectId"`
	// The name of cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Platform type.
	PlatformType pulumi.StringOutput `pulumi:"platformType"`
	// Provider type.
	ProviderType pulumi.StringOutput `pulumi:"providerType"`
}

// NewInstanceCluster registers a new resource with the given unique name, arguments, and options.
func NewInstanceCluster(ctx *pulumi.Context,
	name string, args *InstanceClusterArgs, opts ...pulumi.ResourceOption) (*InstanceCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KubernetesApiUrl == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesApiUrl'")
	}
	if args.KubernetesToken == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesToken'")
	}
	if args.KubernetesToken != nil {
		args.KubernetesToken = pulumi.ToSecret(args.KubernetesToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"kubernetesToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceCluster
	err := ctx.RegisterResource("gitlab:index/instanceCluster:InstanceCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceCluster gets an existing InstanceCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceClusterState, opts ...pulumi.ResourceOption) (*InstanceCluster, error) {
	var resource InstanceCluster
	err := ctx.ReadResource("gitlab:index/instanceCluster:InstanceCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceCluster resources.
type instanceClusterState struct {
	// Cluster type.
	ClusterType *string `pulumi:"clusterType"`
	// Create time.
	CreatedAt *string `pulumi:"createdAt"`
	// The base domain of the cluster.
	Domain *string `pulumi:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled *bool `pulumi:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The URL to access the Kubernetes API.
	KubernetesApiUrl *string `pulumi:"kubernetesApiUrl"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType *string `pulumi:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// The unique namespace related to the instance.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes. This attribute cannot be read.
	KubernetesToken *string `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed *bool `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId *string `pulumi:"managementProjectId"`
	// The name of cluster.
	Name *string `pulumi:"name"`
	// Platform type.
	PlatformType *string `pulumi:"platformType"`
	// Provider type.
	ProviderType *string `pulumi:"providerType"`
}

type InstanceClusterState struct {
	// Cluster type.
	ClusterType pulumi.StringPtrInput
	// Create time.
	CreatedAt pulumi.StringPtrInput
	// The base domain of the cluster.
	Domain pulumi.StringPtrInput
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled pulumi.BoolPtrInput
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope pulumi.StringPtrInput
	// The URL to access the Kubernetes API.
	KubernetesApiUrl pulumi.StringPtrInput
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType pulumi.StringPtrInput
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert pulumi.StringPtrInput
	// The unique namespace related to the instance.
	KubernetesNamespace pulumi.StringPtrInput
	// The token to authenticate against Kubernetes. This attribute cannot be read.
	KubernetesToken pulumi.StringPtrInput
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrInput
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrInput
	// The name of cluster.
	Name pulumi.StringPtrInput
	// Platform type.
	PlatformType pulumi.StringPtrInput
	// Provider type.
	ProviderType pulumi.StringPtrInput
}

func (InstanceClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceClusterState)(nil)).Elem()
}

type instanceClusterArgs struct {
	// The base domain of the cluster.
	Domain *string `pulumi:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled *bool `pulumi:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The URL to access the Kubernetes API.
	KubernetesApiUrl string `pulumi:"kubernetesApiUrl"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType *string `pulumi:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// The unique namespace related to the instance.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes. This attribute cannot be read.
	KubernetesToken string `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed *bool `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId *string `pulumi:"managementProjectId"`
	// The name of cluster.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a InstanceCluster resource.
type InstanceClusterArgs struct {
	// The base domain of the cluster.
	Domain pulumi.StringPtrInput
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled pulumi.BoolPtrInput
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope pulumi.StringPtrInput
	// The URL to access the Kubernetes API.
	KubernetesApiUrl pulumi.StringInput
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType pulumi.StringPtrInput
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert pulumi.StringPtrInput
	// The unique namespace related to the instance.
	KubernetesNamespace pulumi.StringPtrInput
	// The token to authenticate against Kubernetes. This attribute cannot be read.
	KubernetesToken pulumi.StringInput
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrInput
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrInput
	// The name of cluster.
	Name pulumi.StringPtrInput
}

func (InstanceClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceClusterArgs)(nil)).Elem()
}

type InstanceClusterInput interface {
	pulumi.Input

	ToInstanceClusterOutput() InstanceClusterOutput
	ToInstanceClusterOutputWithContext(ctx context.Context) InstanceClusterOutput
}

func (*InstanceCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceCluster)(nil)).Elem()
}

func (i *InstanceCluster) ToInstanceClusterOutput() InstanceClusterOutput {
	return i.ToInstanceClusterOutputWithContext(context.Background())
}

func (i *InstanceCluster) ToInstanceClusterOutputWithContext(ctx context.Context) InstanceClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceClusterOutput)
}

// InstanceClusterArrayInput is an input type that accepts InstanceClusterArray and InstanceClusterArrayOutput values.
// You can construct a concrete instance of `InstanceClusterArrayInput` via:
//
//	InstanceClusterArray{ InstanceClusterArgs{...} }
type InstanceClusterArrayInput interface {
	pulumi.Input

	ToInstanceClusterArrayOutput() InstanceClusterArrayOutput
	ToInstanceClusterArrayOutputWithContext(context.Context) InstanceClusterArrayOutput
}

type InstanceClusterArray []InstanceClusterInput

func (InstanceClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceCluster)(nil)).Elem()
}

func (i InstanceClusterArray) ToInstanceClusterArrayOutput() InstanceClusterArrayOutput {
	return i.ToInstanceClusterArrayOutputWithContext(context.Background())
}

func (i InstanceClusterArray) ToInstanceClusterArrayOutputWithContext(ctx context.Context) InstanceClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceClusterArrayOutput)
}

// InstanceClusterMapInput is an input type that accepts InstanceClusterMap and InstanceClusterMapOutput values.
// You can construct a concrete instance of `InstanceClusterMapInput` via:
//
//	InstanceClusterMap{ "key": InstanceClusterArgs{...} }
type InstanceClusterMapInput interface {
	pulumi.Input

	ToInstanceClusterMapOutput() InstanceClusterMapOutput
	ToInstanceClusterMapOutputWithContext(context.Context) InstanceClusterMapOutput
}

type InstanceClusterMap map[string]InstanceClusterInput

func (InstanceClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceCluster)(nil)).Elem()
}

func (i InstanceClusterMap) ToInstanceClusterMapOutput() InstanceClusterMapOutput {
	return i.ToInstanceClusterMapOutputWithContext(context.Background())
}

func (i InstanceClusterMap) ToInstanceClusterMapOutputWithContext(ctx context.Context) InstanceClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceClusterMapOutput)
}

type InstanceClusterOutput struct{ *pulumi.OutputState }

func (InstanceClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceCluster)(nil)).Elem()
}

func (o InstanceClusterOutput) ToInstanceClusterOutput() InstanceClusterOutput {
	return o
}

func (o InstanceClusterOutput) ToInstanceClusterOutputWithContext(ctx context.Context) InstanceClusterOutput {
	return o
}

// Cluster type.
func (o InstanceClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// Create time.
func (o InstanceClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The base domain of the cluster.
func (o InstanceClusterOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
func (o InstanceClusterOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The associated environment to the cluster. Defaults to `*`.
func (o InstanceClusterOutput) EnvironmentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringPtrOutput { return v.EnvironmentScope }).(pulumi.StringPtrOutput)
}

// The URL to access the Kubernetes API.
func (o InstanceClusterOutput) KubernetesApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.KubernetesApiUrl }).(pulumi.StringOutput)
}

// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
func (o InstanceClusterOutput) KubernetesAuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringPtrOutput { return v.KubernetesAuthorizationType }).(pulumi.StringPtrOutput)
}

// TLS certificate (needed if API is using a self-signed TLS certificate).
func (o InstanceClusterOutput) KubernetesCaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringPtrOutput { return v.KubernetesCaCert }).(pulumi.StringPtrOutput)
}

// The unique namespace related to the instance.
func (o InstanceClusterOutput) KubernetesNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringPtrOutput { return v.KubernetesNamespace }).(pulumi.StringPtrOutput)
}

// The token to authenticate against Kubernetes. This attribute cannot be read.
func (o InstanceClusterOutput) KubernetesToken() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.KubernetesToken }).(pulumi.StringOutput)
}

// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
func (o InstanceClusterOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.BoolPtrOutput { return v.Managed }).(pulumi.BoolPtrOutput)
}

// The ID of the management project for the cluster.
func (o InstanceClusterOutput) ManagementProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringPtrOutput { return v.ManagementProjectId }).(pulumi.StringPtrOutput)
}

// The name of cluster.
func (o InstanceClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Platform type.
func (o InstanceClusterOutput) PlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.PlatformType }).(pulumi.StringOutput)
}

// Provider type.
func (o InstanceClusterOutput) ProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceCluster) pulumi.StringOutput { return v.ProviderType }).(pulumi.StringOutput)
}

type InstanceClusterArrayOutput struct{ *pulumi.OutputState }

func (InstanceClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceCluster)(nil)).Elem()
}

func (o InstanceClusterArrayOutput) ToInstanceClusterArrayOutput() InstanceClusterArrayOutput {
	return o
}

func (o InstanceClusterArrayOutput) ToInstanceClusterArrayOutputWithContext(ctx context.Context) InstanceClusterArrayOutput {
	return o
}

func (o InstanceClusterArrayOutput) Index(i pulumi.IntInput) InstanceClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceCluster {
		return vs[0].([]*InstanceCluster)[vs[1].(int)]
	}).(InstanceClusterOutput)
}

type InstanceClusterMapOutput struct{ *pulumi.OutputState }

func (InstanceClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceCluster)(nil)).Elem()
}

func (o InstanceClusterMapOutput) ToInstanceClusterMapOutput() InstanceClusterMapOutput {
	return o
}

func (o InstanceClusterMapOutput) ToInstanceClusterMapOutputWithContext(ctx context.Context) InstanceClusterMapOutput {
	return o
}

func (o InstanceClusterMapOutput) MapIndex(k pulumi.StringInput) InstanceClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceCluster {
		return vs[0].(map[string]*InstanceCluster)[vs[1].(string)]
	}).(InstanceClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceClusterInput)(nil)).Elem(), &InstanceCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceClusterArrayInput)(nil)).Elem(), InstanceClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceClusterMapInput)(nil)).Elem(), InstanceClusterMap{})
	pulumi.RegisterOutputType(InstanceClusterOutput{})
	pulumi.RegisterOutputType(InstanceClusterArrayOutput{})
	pulumi.RegisterOutputType(InstanceClusterMapOutput{})
}
