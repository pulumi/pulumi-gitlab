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
//			this, err := gitlab.NewGroup(ctx, "this", &gitlab.GroupArgs{
//				Name:        pulumi.String("example"),
//				Path:        pulumi.String("example"),
//				Description: pulumi.String("An example group"),
//			})
//			if err != nil {
//				return err
//			}
//			thisProject, err := gitlab.NewProject(ctx, "this", &gitlab.ProjectArgs{
//				Name:                 pulumi.String("example"),
//				NamespaceId:          this.ID(),
//				InitializeWithReadme: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectEnvironment(ctx, "this", &gitlab.ProjectEnvironmentArgs{
//				Project:     thisProject.ID(),
//				Name:        pulumi.String("example"),
//				ExternalUrl: pulumi.String("www.example.com"),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_environment`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_environment.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// GitLab project environments can be imported using an id made up of `projectId:environmenId`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectEnvironment:ProjectEnvironment bar 123:321
// ```
type ProjectEnvironment struct {
	pulumi.CustomResourceState

	// The cluster agent to associate with this environment.
	ClusterAgentId pulumi.IntPtrOutput `pulumi:"clusterAgentId"`
	// The ISO8601 date/time that this environment was created at in UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Place to link to for this environment.
	ExternalUrl pulumi.StringPtrOutput `pulumi:"externalUrl"`
	// The Flux resource path to associate with this environment.
	FluxResourcePath pulumi.StringPtrOutput `pulumi:"fluxResourcePath"`
	// The Kubernetes namespace to associate with this environment.
	KubernetesNamespace pulumi.StringPtrOutput `pulumi:"kubernetesNamespace"`
	// The name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID or full path of the project to environment is created for.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// State the environment is in. Valid values are `available`, `stopped`.
	State pulumi.StringOutput `pulumi:"state"`
	// Determines whether the environment is attempted to be stopped before the environment is deleted.
	StopBeforeDestroy pulumi.BoolPtrOutput `pulumi:"stopBeforeDestroy"`
	// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// The ISO8601 date/time that this environment was last updated at in UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewProjectEnvironment registers a new resource with the given unique name, arguments, and options.
func NewProjectEnvironment(ctx *pulumi.Context,
	name string, args *ProjectEnvironmentArgs, opts ...pulumi.ResourceOption) (*ProjectEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectEnvironment
	err := ctx.RegisterResource("gitlab:index/projectEnvironment:ProjectEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectEnvironment gets an existing ProjectEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectEnvironmentState, opts ...pulumi.ResourceOption) (*ProjectEnvironment, error) {
	var resource ProjectEnvironment
	err := ctx.ReadResource("gitlab:index/projectEnvironment:ProjectEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectEnvironment resources.
type projectEnvironmentState struct {
	// The cluster agent to associate with this environment.
	ClusterAgentId *int `pulumi:"clusterAgentId"`
	// The ISO8601 date/time that this environment was created at in UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Place to link to for this environment.
	ExternalUrl *string `pulumi:"externalUrl"`
	// The Flux resource path to associate with this environment.
	FluxResourcePath *string `pulumi:"fluxResourcePath"`
	// The Kubernetes namespace to associate with this environment.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The ID or full path of the project to environment is created for.
	Project *string `pulumi:"project"`
	// The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug *string `pulumi:"slug"`
	// State the environment is in. Valid values are `available`, `stopped`.
	State *string `pulumi:"state"`
	// Determines whether the environment is attempted to be stopped before the environment is deleted.
	StopBeforeDestroy *bool `pulumi:"stopBeforeDestroy"`
	// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Tier *string `pulumi:"tier"`
	// The ISO8601 date/time that this environment was last updated at in UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ProjectEnvironmentState struct {
	// The cluster agent to associate with this environment.
	ClusterAgentId pulumi.IntPtrInput
	// The ISO8601 date/time that this environment was created at in UTC.
	CreatedAt pulumi.StringPtrInput
	// Place to link to for this environment.
	ExternalUrl pulumi.StringPtrInput
	// The Flux resource path to associate with this environment.
	FluxResourcePath pulumi.StringPtrInput
	// The Kubernetes namespace to associate with this environment.
	KubernetesNamespace pulumi.StringPtrInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The ID or full path of the project to environment is created for.
	Project pulumi.StringPtrInput
	// The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringPtrInput
	// State the environment is in. Valid values are `available`, `stopped`.
	State pulumi.StringPtrInput
	// Determines whether the environment is attempted to be stopped before the environment is deleted.
	StopBeforeDestroy pulumi.BoolPtrInput
	// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Tier pulumi.StringPtrInput
	// The ISO8601 date/time that this environment was last updated at in UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (ProjectEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectEnvironmentState)(nil)).Elem()
}

type projectEnvironmentArgs struct {
	// The cluster agent to associate with this environment.
	ClusterAgentId *int `pulumi:"clusterAgentId"`
	// Place to link to for this environment.
	ExternalUrl *string `pulumi:"externalUrl"`
	// The Flux resource path to associate with this environment.
	FluxResourcePath *string `pulumi:"fluxResourcePath"`
	// The Kubernetes namespace to associate with this environment.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The ID or full path of the project to environment is created for.
	Project string `pulumi:"project"`
	// Determines whether the environment is attempted to be stopped before the environment is deleted.
	StopBeforeDestroy *bool `pulumi:"stopBeforeDestroy"`
	// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Tier *string `pulumi:"tier"`
}

// The set of arguments for constructing a ProjectEnvironment resource.
type ProjectEnvironmentArgs struct {
	// The cluster agent to associate with this environment.
	ClusterAgentId pulumi.IntPtrInput
	// Place to link to for this environment.
	ExternalUrl pulumi.StringPtrInput
	// The Flux resource path to associate with this environment.
	FluxResourcePath pulumi.StringPtrInput
	// The Kubernetes namespace to associate with this environment.
	KubernetesNamespace pulumi.StringPtrInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The ID or full path of the project to environment is created for.
	Project pulumi.StringInput
	// Determines whether the environment is attempted to be stopped before the environment is deleted.
	StopBeforeDestroy pulumi.BoolPtrInput
	// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
	Tier pulumi.StringPtrInput
}

func (ProjectEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectEnvironmentArgs)(nil)).Elem()
}

type ProjectEnvironmentInput interface {
	pulumi.Input

	ToProjectEnvironmentOutput() ProjectEnvironmentOutput
	ToProjectEnvironmentOutputWithContext(ctx context.Context) ProjectEnvironmentOutput
}

func (*ProjectEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironment)(nil)).Elem()
}

func (i *ProjectEnvironment) ToProjectEnvironmentOutput() ProjectEnvironmentOutput {
	return i.ToProjectEnvironmentOutputWithContext(context.Background())
}

func (i *ProjectEnvironment) ToProjectEnvironmentOutputWithContext(ctx context.Context) ProjectEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentOutput)
}

// ProjectEnvironmentArrayInput is an input type that accepts ProjectEnvironmentArray and ProjectEnvironmentArrayOutput values.
// You can construct a concrete instance of `ProjectEnvironmentArrayInput` via:
//
//	ProjectEnvironmentArray{ ProjectEnvironmentArgs{...} }
type ProjectEnvironmentArrayInput interface {
	pulumi.Input

	ToProjectEnvironmentArrayOutput() ProjectEnvironmentArrayOutput
	ToProjectEnvironmentArrayOutputWithContext(context.Context) ProjectEnvironmentArrayOutput
}

type ProjectEnvironmentArray []ProjectEnvironmentInput

func (ProjectEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectEnvironment)(nil)).Elem()
}

func (i ProjectEnvironmentArray) ToProjectEnvironmentArrayOutput() ProjectEnvironmentArrayOutput {
	return i.ToProjectEnvironmentArrayOutputWithContext(context.Background())
}

func (i ProjectEnvironmentArray) ToProjectEnvironmentArrayOutputWithContext(ctx context.Context) ProjectEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentArrayOutput)
}

// ProjectEnvironmentMapInput is an input type that accepts ProjectEnvironmentMap and ProjectEnvironmentMapOutput values.
// You can construct a concrete instance of `ProjectEnvironmentMapInput` via:
//
//	ProjectEnvironmentMap{ "key": ProjectEnvironmentArgs{...} }
type ProjectEnvironmentMapInput interface {
	pulumi.Input

	ToProjectEnvironmentMapOutput() ProjectEnvironmentMapOutput
	ToProjectEnvironmentMapOutputWithContext(context.Context) ProjectEnvironmentMapOutput
}

type ProjectEnvironmentMap map[string]ProjectEnvironmentInput

func (ProjectEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectEnvironment)(nil)).Elem()
}

func (i ProjectEnvironmentMap) ToProjectEnvironmentMapOutput() ProjectEnvironmentMapOutput {
	return i.ToProjectEnvironmentMapOutputWithContext(context.Background())
}

func (i ProjectEnvironmentMap) ToProjectEnvironmentMapOutputWithContext(ctx context.Context) ProjectEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentMapOutput)
}

type ProjectEnvironmentOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironment)(nil)).Elem()
}

func (o ProjectEnvironmentOutput) ToProjectEnvironmentOutput() ProjectEnvironmentOutput {
	return o
}

func (o ProjectEnvironmentOutput) ToProjectEnvironmentOutputWithContext(ctx context.Context) ProjectEnvironmentOutput {
	return o
}

// The cluster agent to associate with this environment.
func (o ProjectEnvironmentOutput) ClusterAgentId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.IntPtrOutput { return v.ClusterAgentId }).(pulumi.IntPtrOutput)
}

// The ISO8601 date/time that this environment was created at in UTC.
func (o ProjectEnvironmentOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Place to link to for this environment.
func (o ProjectEnvironmentOutput) ExternalUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringPtrOutput { return v.ExternalUrl }).(pulumi.StringPtrOutput)
}

// The Flux resource path to associate with this environment.
func (o ProjectEnvironmentOutput) FluxResourcePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringPtrOutput { return v.FluxResourcePath }).(pulumi.StringPtrOutput)
}

// The Kubernetes namespace to associate with this environment.
func (o ProjectEnvironmentOutput) KubernetesNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringPtrOutput { return v.KubernetesNamespace }).(pulumi.StringPtrOutput)
}

// The name of the environment.
func (o ProjectEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID or full path of the project to environment is created for.
func (o ProjectEnvironmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
func (o ProjectEnvironmentOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// State the environment is in. Valid values are `available`, `stopped`.
func (o ProjectEnvironmentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Determines whether the environment is attempted to be stopped before the environment is deleted.
func (o ProjectEnvironmentOutput) StopBeforeDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.BoolPtrOutput { return v.StopBeforeDestroy }).(pulumi.BoolPtrOutput)
}

// The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
func (o ProjectEnvironmentOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

// The ISO8601 date/time that this environment was last updated at in UTC.
func (o ProjectEnvironmentOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironment) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ProjectEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectEnvironment)(nil)).Elem()
}

func (o ProjectEnvironmentArrayOutput) ToProjectEnvironmentArrayOutput() ProjectEnvironmentArrayOutput {
	return o
}

func (o ProjectEnvironmentArrayOutput) ToProjectEnvironmentArrayOutputWithContext(ctx context.Context) ProjectEnvironmentArrayOutput {
	return o
}

func (o ProjectEnvironmentArrayOutput) Index(i pulumi.IntInput) ProjectEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectEnvironment {
		return vs[0].([]*ProjectEnvironment)[vs[1].(int)]
	}).(ProjectEnvironmentOutput)
}

type ProjectEnvironmentMapOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectEnvironment)(nil)).Elem()
}

func (o ProjectEnvironmentMapOutput) ToProjectEnvironmentMapOutput() ProjectEnvironmentMapOutput {
	return o
}

func (o ProjectEnvironmentMapOutput) ToProjectEnvironmentMapOutputWithContext(ctx context.Context) ProjectEnvironmentMapOutput {
	return o
}

func (o ProjectEnvironmentMapOutput) MapIndex(k pulumi.StringInput) ProjectEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectEnvironment {
		return vs[0].(map[string]*ProjectEnvironment)[vs[1].(string)]
	}).(ProjectEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectEnvironmentInput)(nil)).Elem(), &ProjectEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectEnvironmentArrayInput)(nil)).Elem(), ProjectEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectEnvironmentMapInput)(nil)).Elem(), ProjectEnvironmentMap{})
	pulumi.RegisterOutputType(ProjectEnvironmentOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentMapOutput{})
}
