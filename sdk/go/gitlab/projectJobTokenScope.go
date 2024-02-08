// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectJobTokenScope` resource allows to manage the CI/CD Job Token scope in a project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_job_token_scopes.html)
//
// ## Import
//
// GitLab project environments can be imported using an id made up of `projectId:targetProjectId`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectJobTokenScope:ProjectJobTokenScope bar 123:321
// ```
type ProjectJobTokenScope struct {
	pulumi.CustomResourceState

	// The ID or full path of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// The ID of the project that is in the CI/CD job token inbound allowlist.
	TargetProjectId pulumi.IntOutput `pulumi:"targetProjectId"`
}

// NewProjectJobTokenScope registers a new resource with the given unique name, arguments, and options.
func NewProjectJobTokenScope(ctx *pulumi.Context,
	name string, args *ProjectJobTokenScopeArgs, opts ...pulumi.ResourceOption) (*ProjectJobTokenScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.TargetProjectId == nil {
		return nil, errors.New("invalid value for required argument 'TargetProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectJobTokenScope
	err := ctx.RegisterResource("gitlab:index/projectJobTokenScope:ProjectJobTokenScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectJobTokenScope gets an existing ProjectJobTokenScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectJobTokenScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectJobTokenScopeState, opts ...pulumi.ResourceOption) (*ProjectJobTokenScope, error) {
	var resource ProjectJobTokenScope
	err := ctx.ReadResource("gitlab:index/projectJobTokenScope:ProjectJobTokenScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectJobTokenScope resources.
type projectJobTokenScopeState struct {
	// The ID or full path of the project.
	Project *string `pulumi:"project"`
	// The ID of the project that is in the CI/CD job token inbound allowlist.
	TargetProjectId *int `pulumi:"targetProjectId"`
}

type ProjectJobTokenScopeState struct {
	// The ID or full path of the project.
	Project pulumi.StringPtrInput
	// The ID of the project that is in the CI/CD job token inbound allowlist.
	TargetProjectId pulumi.IntPtrInput
}

func (ProjectJobTokenScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectJobTokenScopeState)(nil)).Elem()
}

type projectJobTokenScopeArgs struct {
	// The ID or full path of the project.
	Project string `pulumi:"project"`
	// The ID of the project that is in the CI/CD job token inbound allowlist.
	TargetProjectId int `pulumi:"targetProjectId"`
}

// The set of arguments for constructing a ProjectJobTokenScope resource.
type ProjectJobTokenScopeArgs struct {
	// The ID or full path of the project.
	Project pulumi.StringInput
	// The ID of the project that is in the CI/CD job token inbound allowlist.
	TargetProjectId pulumi.IntInput
}

func (ProjectJobTokenScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectJobTokenScopeArgs)(nil)).Elem()
}

type ProjectJobTokenScopeInput interface {
	pulumi.Input

	ToProjectJobTokenScopeOutput() ProjectJobTokenScopeOutput
	ToProjectJobTokenScopeOutputWithContext(ctx context.Context) ProjectJobTokenScopeOutput
}

func (*ProjectJobTokenScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectJobTokenScope)(nil)).Elem()
}

func (i *ProjectJobTokenScope) ToProjectJobTokenScopeOutput() ProjectJobTokenScopeOutput {
	return i.ToProjectJobTokenScopeOutputWithContext(context.Background())
}

func (i *ProjectJobTokenScope) ToProjectJobTokenScopeOutputWithContext(ctx context.Context) ProjectJobTokenScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectJobTokenScopeOutput)
}

// ProjectJobTokenScopeArrayInput is an input type that accepts ProjectJobTokenScopeArray and ProjectJobTokenScopeArrayOutput values.
// You can construct a concrete instance of `ProjectJobTokenScopeArrayInput` via:
//
//	ProjectJobTokenScopeArray{ ProjectJobTokenScopeArgs{...} }
type ProjectJobTokenScopeArrayInput interface {
	pulumi.Input

	ToProjectJobTokenScopeArrayOutput() ProjectJobTokenScopeArrayOutput
	ToProjectJobTokenScopeArrayOutputWithContext(context.Context) ProjectJobTokenScopeArrayOutput
}

type ProjectJobTokenScopeArray []ProjectJobTokenScopeInput

func (ProjectJobTokenScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectJobTokenScope)(nil)).Elem()
}

func (i ProjectJobTokenScopeArray) ToProjectJobTokenScopeArrayOutput() ProjectJobTokenScopeArrayOutput {
	return i.ToProjectJobTokenScopeArrayOutputWithContext(context.Background())
}

func (i ProjectJobTokenScopeArray) ToProjectJobTokenScopeArrayOutputWithContext(ctx context.Context) ProjectJobTokenScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectJobTokenScopeArrayOutput)
}

// ProjectJobTokenScopeMapInput is an input type that accepts ProjectJobTokenScopeMap and ProjectJobTokenScopeMapOutput values.
// You can construct a concrete instance of `ProjectJobTokenScopeMapInput` via:
//
//	ProjectJobTokenScopeMap{ "key": ProjectJobTokenScopeArgs{...} }
type ProjectJobTokenScopeMapInput interface {
	pulumi.Input

	ToProjectJobTokenScopeMapOutput() ProjectJobTokenScopeMapOutput
	ToProjectJobTokenScopeMapOutputWithContext(context.Context) ProjectJobTokenScopeMapOutput
}

type ProjectJobTokenScopeMap map[string]ProjectJobTokenScopeInput

func (ProjectJobTokenScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectJobTokenScope)(nil)).Elem()
}

func (i ProjectJobTokenScopeMap) ToProjectJobTokenScopeMapOutput() ProjectJobTokenScopeMapOutput {
	return i.ToProjectJobTokenScopeMapOutputWithContext(context.Background())
}

func (i ProjectJobTokenScopeMap) ToProjectJobTokenScopeMapOutputWithContext(ctx context.Context) ProjectJobTokenScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectJobTokenScopeMapOutput)
}

type ProjectJobTokenScopeOutput struct{ *pulumi.OutputState }

func (ProjectJobTokenScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectJobTokenScope)(nil)).Elem()
}

func (o ProjectJobTokenScopeOutput) ToProjectJobTokenScopeOutput() ProjectJobTokenScopeOutput {
	return o
}

func (o ProjectJobTokenScopeOutput) ToProjectJobTokenScopeOutputWithContext(ctx context.Context) ProjectJobTokenScopeOutput {
	return o
}

// The ID or full path of the project.
func (o ProjectJobTokenScopeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectJobTokenScope) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The ID of the project that is in the CI/CD job token inbound allowlist.
func (o ProjectJobTokenScopeOutput) TargetProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectJobTokenScope) pulumi.IntOutput { return v.TargetProjectId }).(pulumi.IntOutput)
}

type ProjectJobTokenScopeArrayOutput struct{ *pulumi.OutputState }

func (ProjectJobTokenScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectJobTokenScope)(nil)).Elem()
}

func (o ProjectJobTokenScopeArrayOutput) ToProjectJobTokenScopeArrayOutput() ProjectJobTokenScopeArrayOutput {
	return o
}

func (o ProjectJobTokenScopeArrayOutput) ToProjectJobTokenScopeArrayOutputWithContext(ctx context.Context) ProjectJobTokenScopeArrayOutput {
	return o
}

func (o ProjectJobTokenScopeArrayOutput) Index(i pulumi.IntInput) ProjectJobTokenScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectJobTokenScope {
		return vs[0].([]*ProjectJobTokenScope)[vs[1].(int)]
	}).(ProjectJobTokenScopeOutput)
}

type ProjectJobTokenScopeMapOutput struct{ *pulumi.OutputState }

func (ProjectJobTokenScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectJobTokenScope)(nil)).Elem()
}

func (o ProjectJobTokenScopeMapOutput) ToProjectJobTokenScopeMapOutput() ProjectJobTokenScopeMapOutput {
	return o
}

func (o ProjectJobTokenScopeMapOutput) ToProjectJobTokenScopeMapOutputWithContext(ctx context.Context) ProjectJobTokenScopeMapOutput {
	return o
}

func (o ProjectJobTokenScopeMapOutput) MapIndex(k pulumi.StringInput) ProjectJobTokenScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectJobTokenScope {
		return vs[0].(map[string]*ProjectJobTokenScope)[vs[1].(string)]
	}).(ProjectJobTokenScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectJobTokenScopeInput)(nil)).Elem(), &ProjectJobTokenScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectJobTokenScopeArrayInput)(nil)).Elem(), ProjectJobTokenScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectJobTokenScopeMapInput)(nil)).Elem(), ProjectJobTokenScopeMap{})
	pulumi.RegisterOutputType(ProjectJobTokenScopeOutput{})
	pulumi.RegisterOutputType(ProjectJobTokenScopeArrayOutput{})
	pulumi.RegisterOutputType(ProjectJobTokenScopeMapOutput{})
}
