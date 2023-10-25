// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `ProjectBadge` resource allows to manage the lifecycle of project badges.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#project-badges)
//
// ## Import
//
// GitLab project badges can be imported using an id made up of `{project_id}:{badge_id}`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/projectBadge:ProjectBadge foo 1:3
//
// ```
type ProjectBadge struct {
	pulumi.CustomResourceState

	// The image url which will be presented on project overview.
	ImageUrl pulumi.StringOutput `pulumi:"imageUrl"`
	// The url linked with the badge.
	LinkUrl pulumi.StringOutput `pulumi:"linkUrl"`
	// The name of the badge.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the project to add the badge to.
	Project pulumi.StringOutput `pulumi:"project"`
	// The imageUrl argument rendered (in case of use of placeholders).
	RenderedImageUrl pulumi.StringOutput `pulumi:"renderedImageUrl"`
	// The linkUrl argument rendered (in case of use of placeholders).
	RenderedLinkUrl pulumi.StringOutput `pulumi:"renderedLinkUrl"`
}

// NewProjectBadge registers a new resource with the given unique name, arguments, and options.
func NewProjectBadge(ctx *pulumi.Context,
	name string, args *ProjectBadgeArgs, opts ...pulumi.ResourceOption) (*ProjectBadge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageUrl == nil {
		return nil, errors.New("invalid value for required argument 'ImageUrl'")
	}
	if args.LinkUrl == nil {
		return nil, errors.New("invalid value for required argument 'LinkUrl'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectBadge
	err := ctx.RegisterResource("gitlab:index/projectBadge:ProjectBadge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectBadge gets an existing ProjectBadge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectBadge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectBadgeState, opts ...pulumi.ResourceOption) (*ProjectBadge, error) {
	var resource ProjectBadge
	err := ctx.ReadResource("gitlab:index/projectBadge:ProjectBadge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectBadge resources.
type projectBadgeState struct {
	// The image url which will be presented on project overview.
	ImageUrl *string `pulumi:"imageUrl"`
	// The url linked with the badge.
	LinkUrl *string `pulumi:"linkUrl"`
	// The name of the badge.
	Name *string `pulumi:"name"`
	// The id of the project to add the badge to.
	Project *string `pulumi:"project"`
	// The imageUrl argument rendered (in case of use of placeholders).
	RenderedImageUrl *string `pulumi:"renderedImageUrl"`
	// The linkUrl argument rendered (in case of use of placeholders).
	RenderedLinkUrl *string `pulumi:"renderedLinkUrl"`
}

type ProjectBadgeState struct {
	// The image url which will be presented on project overview.
	ImageUrl pulumi.StringPtrInput
	// The url linked with the badge.
	LinkUrl pulumi.StringPtrInput
	// The name of the badge.
	Name pulumi.StringPtrInput
	// The id of the project to add the badge to.
	Project pulumi.StringPtrInput
	// The imageUrl argument rendered (in case of use of placeholders).
	RenderedImageUrl pulumi.StringPtrInput
	// The linkUrl argument rendered (in case of use of placeholders).
	RenderedLinkUrl pulumi.StringPtrInput
}

func (ProjectBadgeState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBadgeState)(nil)).Elem()
}

type projectBadgeArgs struct {
	// The image url which will be presented on project overview.
	ImageUrl string `pulumi:"imageUrl"`
	// The url linked with the badge.
	LinkUrl string `pulumi:"linkUrl"`
	// The name of the badge.
	Name *string `pulumi:"name"`
	// The id of the project to add the badge to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectBadge resource.
type ProjectBadgeArgs struct {
	// The image url which will be presented on project overview.
	ImageUrl pulumi.StringInput
	// The url linked with the badge.
	LinkUrl pulumi.StringInput
	// The name of the badge.
	Name pulumi.StringPtrInput
	// The id of the project to add the badge to.
	Project pulumi.StringInput
}

func (ProjectBadgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBadgeArgs)(nil)).Elem()
}

type ProjectBadgeInput interface {
	pulumi.Input

	ToProjectBadgeOutput() ProjectBadgeOutput
	ToProjectBadgeOutputWithContext(ctx context.Context) ProjectBadgeOutput
}

func (*ProjectBadge) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectBadge)(nil)).Elem()
}

func (i *ProjectBadge) ToProjectBadgeOutput() ProjectBadgeOutput {
	return i.ToProjectBadgeOutputWithContext(context.Background())
}

func (i *ProjectBadge) ToProjectBadgeOutputWithContext(ctx context.Context) ProjectBadgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBadgeOutput)
}

func (i *ProjectBadge) ToOutput(ctx context.Context) pulumix.Output[*ProjectBadge] {
	return pulumix.Output[*ProjectBadge]{
		OutputState: i.ToProjectBadgeOutputWithContext(ctx).OutputState,
	}
}

// ProjectBadgeArrayInput is an input type that accepts ProjectBadgeArray and ProjectBadgeArrayOutput values.
// You can construct a concrete instance of `ProjectBadgeArrayInput` via:
//
//	ProjectBadgeArray{ ProjectBadgeArgs{...} }
type ProjectBadgeArrayInput interface {
	pulumi.Input

	ToProjectBadgeArrayOutput() ProjectBadgeArrayOutput
	ToProjectBadgeArrayOutputWithContext(context.Context) ProjectBadgeArrayOutput
}

type ProjectBadgeArray []ProjectBadgeInput

func (ProjectBadgeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectBadge)(nil)).Elem()
}

func (i ProjectBadgeArray) ToProjectBadgeArrayOutput() ProjectBadgeArrayOutput {
	return i.ToProjectBadgeArrayOutputWithContext(context.Background())
}

func (i ProjectBadgeArray) ToProjectBadgeArrayOutputWithContext(ctx context.Context) ProjectBadgeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBadgeArrayOutput)
}

func (i ProjectBadgeArray) ToOutput(ctx context.Context) pulumix.Output[[]*ProjectBadge] {
	return pulumix.Output[[]*ProjectBadge]{
		OutputState: i.ToProjectBadgeArrayOutputWithContext(ctx).OutputState,
	}
}

// ProjectBadgeMapInput is an input type that accepts ProjectBadgeMap and ProjectBadgeMapOutput values.
// You can construct a concrete instance of `ProjectBadgeMapInput` via:
//
//	ProjectBadgeMap{ "key": ProjectBadgeArgs{...} }
type ProjectBadgeMapInput interface {
	pulumi.Input

	ToProjectBadgeMapOutput() ProjectBadgeMapOutput
	ToProjectBadgeMapOutputWithContext(context.Context) ProjectBadgeMapOutput
}

type ProjectBadgeMap map[string]ProjectBadgeInput

func (ProjectBadgeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectBadge)(nil)).Elem()
}

func (i ProjectBadgeMap) ToProjectBadgeMapOutput() ProjectBadgeMapOutput {
	return i.ToProjectBadgeMapOutputWithContext(context.Background())
}

func (i ProjectBadgeMap) ToProjectBadgeMapOutputWithContext(ctx context.Context) ProjectBadgeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBadgeMapOutput)
}

func (i ProjectBadgeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProjectBadge] {
	return pulumix.Output[map[string]*ProjectBadge]{
		OutputState: i.ToProjectBadgeMapOutputWithContext(ctx).OutputState,
	}
}

type ProjectBadgeOutput struct{ *pulumi.OutputState }

func (ProjectBadgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectBadge)(nil)).Elem()
}

func (o ProjectBadgeOutput) ToProjectBadgeOutput() ProjectBadgeOutput {
	return o
}

func (o ProjectBadgeOutput) ToProjectBadgeOutputWithContext(ctx context.Context) ProjectBadgeOutput {
	return o
}

func (o ProjectBadgeOutput) ToOutput(ctx context.Context) pulumix.Output[*ProjectBadge] {
	return pulumix.Output[*ProjectBadge]{
		OutputState: o.OutputState,
	}
}

// The image url which will be presented on project overview.
func (o ProjectBadgeOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBadge) pulumi.StringOutput { return v.ImageUrl }).(pulumi.StringOutput)
}

// The url linked with the badge.
func (o ProjectBadgeOutput) LinkUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBadge) pulumi.StringOutput { return v.LinkUrl }).(pulumi.StringOutput)
}

// The name of the badge.
func (o ProjectBadgeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBadge) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the project to add the badge to.
func (o ProjectBadgeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBadge) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The imageUrl argument rendered (in case of use of placeholders).
func (o ProjectBadgeOutput) RenderedImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBadge) pulumi.StringOutput { return v.RenderedImageUrl }).(pulumi.StringOutput)
}

// The linkUrl argument rendered (in case of use of placeholders).
func (o ProjectBadgeOutput) RenderedLinkUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBadge) pulumi.StringOutput { return v.RenderedLinkUrl }).(pulumi.StringOutput)
}

type ProjectBadgeArrayOutput struct{ *pulumi.OutputState }

func (ProjectBadgeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectBadge)(nil)).Elem()
}

func (o ProjectBadgeArrayOutput) ToProjectBadgeArrayOutput() ProjectBadgeArrayOutput {
	return o
}

func (o ProjectBadgeArrayOutput) ToProjectBadgeArrayOutputWithContext(ctx context.Context) ProjectBadgeArrayOutput {
	return o
}

func (o ProjectBadgeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ProjectBadge] {
	return pulumix.Output[[]*ProjectBadge]{
		OutputState: o.OutputState,
	}
}

func (o ProjectBadgeArrayOutput) Index(i pulumi.IntInput) ProjectBadgeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectBadge {
		return vs[0].([]*ProjectBadge)[vs[1].(int)]
	}).(ProjectBadgeOutput)
}

type ProjectBadgeMapOutput struct{ *pulumi.OutputState }

func (ProjectBadgeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectBadge)(nil)).Elem()
}

func (o ProjectBadgeMapOutput) ToProjectBadgeMapOutput() ProjectBadgeMapOutput {
	return o
}

func (o ProjectBadgeMapOutput) ToProjectBadgeMapOutputWithContext(ctx context.Context) ProjectBadgeMapOutput {
	return o
}

func (o ProjectBadgeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ProjectBadge] {
	return pulumix.Output[map[string]*ProjectBadge]{
		OutputState: o.OutputState,
	}
}

func (o ProjectBadgeMapOutput) MapIndex(k pulumi.StringInput) ProjectBadgeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectBadge {
		return vs[0].(map[string]*ProjectBadge)[vs[1].(string)]
	}).(ProjectBadgeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBadgeInput)(nil)).Elem(), &ProjectBadge{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBadgeArrayInput)(nil)).Elem(), ProjectBadgeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBadgeMapInput)(nil)).Elem(), ProjectBadgeMap{})
	pulumi.RegisterOutputType(ProjectBadgeOutput{})
	pulumi.RegisterOutputType(ProjectBadgeArrayOutput{})
	pulumi.RegisterOutputType(ProjectBadgeMapOutput{})
}
