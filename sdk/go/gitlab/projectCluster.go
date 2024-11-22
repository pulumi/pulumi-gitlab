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

// The `ProjectCluster` resource allows to manage the lifecycle of a project cluster.
//
// > This is deprecated GitLab feature since 14.5
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_clusters.html)
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
//			foo, err := gitlab.NewProject(ctx, "foo", &gitlab.ProjectArgs{
//				Name: pulumi.String("foo-project"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectCluster(ctx, "bar", &gitlab.ProjectClusterArgs{
//				Project:                     foo.ID(),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_cluster`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_cluster.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// GitLab project clusters can be imported using an id made up of `projectid:clusterid`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectCluster:ProjectCluster bar 123:321
// ```
type ProjectCluster struct {
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
	// The unique namespace related to the project.
	KubernetesNamespace pulumi.StringPtrOutput `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes.
	KubernetesToken pulumi.StringOutput `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrOutput `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrOutput `pulumi:"managementProjectId"`
	// The name of cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Platform type.
	PlatformType pulumi.StringOutput `pulumi:"platformType"`
	// The id of the project to add the cluster to.
	Project pulumi.StringOutput `pulumi:"project"`
	// Provider type.
	ProviderType pulumi.StringOutput `pulumi:"providerType"`
}

// NewProjectCluster registers a new resource with the given unique name, arguments, and options.
func NewProjectCluster(ctx *pulumi.Context,
	name string, args *ProjectClusterArgs, opts ...pulumi.ResourceOption) (*ProjectCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KubernetesApiUrl == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesApiUrl'")
	}
	if args.KubernetesToken == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesToken'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.KubernetesToken != nil {
		args.KubernetesToken = pulumi.ToSecret(args.KubernetesToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"kubernetesToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectCluster
	err := ctx.RegisterResource("gitlab:index/projectCluster:ProjectCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectCluster gets an existing ProjectCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectClusterState, opts ...pulumi.ResourceOption) (*ProjectCluster, error) {
	var resource ProjectCluster
	err := ctx.ReadResource("gitlab:index/projectCluster:ProjectCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectCluster resources.
type projectClusterState struct {
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
	// The unique namespace related to the project.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes.
	KubernetesToken *string `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed *bool `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId *string `pulumi:"managementProjectId"`
	// The name of cluster.
	Name *string `pulumi:"name"`
	// Platform type.
	PlatformType *string `pulumi:"platformType"`
	// The id of the project to add the cluster to.
	Project *string `pulumi:"project"`
	// Provider type.
	ProviderType *string `pulumi:"providerType"`
}

type ProjectClusterState struct {
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
	// The unique namespace related to the project.
	KubernetesNamespace pulumi.StringPtrInput
	// The token to authenticate against Kubernetes.
	KubernetesToken pulumi.StringPtrInput
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrInput
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrInput
	// The name of cluster.
	Name pulumi.StringPtrInput
	// Platform type.
	PlatformType pulumi.StringPtrInput
	// The id of the project to add the cluster to.
	Project pulumi.StringPtrInput
	// Provider type.
	ProviderType pulumi.StringPtrInput
}

func (ProjectClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectClusterState)(nil)).Elem()
}

type projectClusterArgs struct {
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
	// The unique namespace related to the project.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes.
	KubernetesToken string `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed *bool `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId *string `pulumi:"managementProjectId"`
	// The name of cluster.
	Name *string `pulumi:"name"`
	// The id of the project to add the cluster to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectCluster resource.
type ProjectClusterArgs struct {
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
	// The unique namespace related to the project.
	KubernetesNamespace pulumi.StringPtrInput
	// The token to authenticate against Kubernetes.
	KubernetesToken pulumi.StringInput
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrInput
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrInput
	// The name of cluster.
	Name pulumi.StringPtrInput
	// The id of the project to add the cluster to.
	Project pulumi.StringInput
}

func (ProjectClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectClusterArgs)(nil)).Elem()
}

type ProjectClusterInput interface {
	pulumi.Input

	ToProjectClusterOutput() ProjectClusterOutput
	ToProjectClusterOutputWithContext(ctx context.Context) ProjectClusterOutput
}

func (*ProjectCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectCluster)(nil)).Elem()
}

func (i *ProjectCluster) ToProjectClusterOutput() ProjectClusterOutput {
	return i.ToProjectClusterOutputWithContext(context.Background())
}

func (i *ProjectCluster) ToProjectClusterOutputWithContext(ctx context.Context) ProjectClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectClusterOutput)
}

// ProjectClusterArrayInput is an input type that accepts ProjectClusterArray and ProjectClusterArrayOutput values.
// You can construct a concrete instance of `ProjectClusterArrayInput` via:
//
//	ProjectClusterArray{ ProjectClusterArgs{...} }
type ProjectClusterArrayInput interface {
	pulumi.Input

	ToProjectClusterArrayOutput() ProjectClusterArrayOutput
	ToProjectClusterArrayOutputWithContext(context.Context) ProjectClusterArrayOutput
}

type ProjectClusterArray []ProjectClusterInput

func (ProjectClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectCluster)(nil)).Elem()
}

func (i ProjectClusterArray) ToProjectClusterArrayOutput() ProjectClusterArrayOutput {
	return i.ToProjectClusterArrayOutputWithContext(context.Background())
}

func (i ProjectClusterArray) ToProjectClusterArrayOutputWithContext(ctx context.Context) ProjectClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectClusterArrayOutput)
}

// ProjectClusterMapInput is an input type that accepts ProjectClusterMap and ProjectClusterMapOutput values.
// You can construct a concrete instance of `ProjectClusterMapInput` via:
//
//	ProjectClusterMap{ "key": ProjectClusterArgs{...} }
type ProjectClusterMapInput interface {
	pulumi.Input

	ToProjectClusterMapOutput() ProjectClusterMapOutput
	ToProjectClusterMapOutputWithContext(context.Context) ProjectClusterMapOutput
}

type ProjectClusterMap map[string]ProjectClusterInput

func (ProjectClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectCluster)(nil)).Elem()
}

func (i ProjectClusterMap) ToProjectClusterMapOutput() ProjectClusterMapOutput {
	return i.ToProjectClusterMapOutputWithContext(context.Background())
}

func (i ProjectClusterMap) ToProjectClusterMapOutputWithContext(ctx context.Context) ProjectClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectClusterMapOutput)
}

type ProjectClusterOutput struct{ *pulumi.OutputState }

func (ProjectClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectCluster)(nil)).Elem()
}

func (o ProjectClusterOutput) ToProjectClusterOutput() ProjectClusterOutput {
	return o
}

func (o ProjectClusterOutput) ToProjectClusterOutputWithContext(ctx context.Context) ProjectClusterOutput {
	return o
}

// Cluster type.
func (o ProjectClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// Create time.
func (o ProjectClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The base domain of the cluster.
func (o ProjectClusterOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
func (o ProjectClusterOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The associated environment to the cluster. Defaults to `*`.
func (o ProjectClusterOutput) EnvironmentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringPtrOutput { return v.EnvironmentScope }).(pulumi.StringPtrOutput)
}

// The URL to access the Kubernetes API.
func (o ProjectClusterOutput) KubernetesApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.KubernetesApiUrl }).(pulumi.StringOutput)
}

// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
func (o ProjectClusterOutput) KubernetesAuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringPtrOutput { return v.KubernetesAuthorizationType }).(pulumi.StringPtrOutput)
}

// TLS certificate (needed if API is using a self-signed TLS certificate).
func (o ProjectClusterOutput) KubernetesCaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringPtrOutput { return v.KubernetesCaCert }).(pulumi.StringPtrOutput)
}

// The unique namespace related to the project.
func (o ProjectClusterOutput) KubernetesNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringPtrOutput { return v.KubernetesNamespace }).(pulumi.StringPtrOutput)
}

// The token to authenticate against Kubernetes.
func (o ProjectClusterOutput) KubernetesToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.KubernetesToken }).(pulumi.StringOutput)
}

// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
func (o ProjectClusterOutput) Managed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.BoolPtrOutput { return v.Managed }).(pulumi.BoolPtrOutput)
}

// The ID of the management project for the cluster.
func (o ProjectClusterOutput) ManagementProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringPtrOutput { return v.ManagementProjectId }).(pulumi.StringPtrOutput)
}

// The name of cluster.
func (o ProjectClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Platform type.
func (o ProjectClusterOutput) PlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.PlatformType }).(pulumi.StringOutput)
}

// The id of the project to add the cluster to.
func (o ProjectClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Provider type.
func (o ProjectClusterOutput) ProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCluster) pulumi.StringOutput { return v.ProviderType }).(pulumi.StringOutput)
}

type ProjectClusterArrayOutput struct{ *pulumi.OutputState }

func (ProjectClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectCluster)(nil)).Elem()
}

func (o ProjectClusterArrayOutput) ToProjectClusterArrayOutput() ProjectClusterArrayOutput {
	return o
}

func (o ProjectClusterArrayOutput) ToProjectClusterArrayOutputWithContext(ctx context.Context) ProjectClusterArrayOutput {
	return o
}

func (o ProjectClusterArrayOutput) Index(i pulumi.IntInput) ProjectClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectCluster {
		return vs[0].([]*ProjectCluster)[vs[1].(int)]
	}).(ProjectClusterOutput)
}

type ProjectClusterMapOutput struct{ *pulumi.OutputState }

func (ProjectClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectCluster)(nil)).Elem()
}

func (o ProjectClusterMapOutput) ToProjectClusterMapOutput() ProjectClusterMapOutput {
	return o
}

func (o ProjectClusterMapOutput) ToProjectClusterMapOutputWithContext(ctx context.Context) ProjectClusterMapOutput {
	return o
}

func (o ProjectClusterMapOutput) MapIndex(k pulumi.StringInput) ProjectClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectCluster {
		return vs[0].(map[string]*ProjectCluster)[vs[1].(string)]
	}).(ProjectClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectClusterInput)(nil)).Elem(), &ProjectCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectClusterArrayInput)(nil)).Elem(), ProjectClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectClusterMapInput)(nil)).Elem(), ProjectClusterMap{})
	pulumi.RegisterOutputType(ProjectClusterOutput{})
	pulumi.RegisterOutputType(ProjectClusterArrayOutput{})
	pulumi.RegisterOutputType(ProjectClusterMapOutput{})
}
