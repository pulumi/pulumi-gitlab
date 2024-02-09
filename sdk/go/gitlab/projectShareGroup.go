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

// The `ProjectShareGroup` resource allows to manage the lifecycle of project shared with a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#share-project-with-group)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewProjectShareGroup(ctx, "test", &gitlab.ProjectShareGroupArgs{
//				GroupAccess: pulumi.String("guest"),
//				GroupId:     pulumi.Int(1337),
//				Project:     pulumi.String("12345"),
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
// GitLab project group shares can be imported using an id made up of `projectid:groupid`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectShareGroup:ProjectShareGroup test 12345:1337
// ```
type ProjectShareGroup struct {
	pulumi.CustomResourceState

	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel pulumi.StringPtrOutput `pulumi:"accessLevel"`
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess pulumi.StringPtrOutput `pulumi:"groupAccess"`
	// The id of the group.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// The ID or URL-encoded path of the project.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectShareGroup registers a new resource with the given unique name, arguments, and options.
func NewProjectShareGroup(ctx *pulumi.Context,
	name string, args *ProjectShareGroupArgs, opts ...pulumi.ResourceOption) (*ProjectShareGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectShareGroup
	err := ctx.RegisterResource("gitlab:index/projectShareGroup:ProjectShareGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectShareGroup gets an existing ProjectShareGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectShareGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectShareGroupState, opts ...pulumi.ResourceOption) (*ProjectShareGroup, error) {
	var resource ProjectShareGroup
	err := ctx.ReadResource("gitlab:index/projectShareGroup:ProjectShareGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectShareGroup resources.
type projectShareGroupState struct {
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel *string `pulumi:"accessLevel"`
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess *string `pulumi:"groupAccess"`
	// The id of the group.
	GroupId *int `pulumi:"groupId"`
	// The ID or URL-encoded path of the project.
	Project *string `pulumi:"project"`
}

type ProjectShareGroupState struct {
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel pulumi.StringPtrInput
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess pulumi.StringPtrInput
	// The id of the group.
	GroupId pulumi.IntPtrInput
	// The ID or URL-encoded path of the project.
	Project pulumi.StringPtrInput
}

func (ProjectShareGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectShareGroupState)(nil)).Elem()
}

type projectShareGroupArgs struct {
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel *string `pulumi:"accessLevel"`
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess *string `pulumi:"groupAccess"`
	// The id of the group.
	GroupId int `pulumi:"groupId"`
	// The ID or URL-encoded path of the project.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectShareGroup resource.
type ProjectShareGroupArgs struct {
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel pulumi.StringPtrInput
	// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess pulumi.StringPtrInput
	// The id of the group.
	GroupId pulumi.IntInput
	// The ID or URL-encoded path of the project.
	Project pulumi.StringInput
}

func (ProjectShareGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectShareGroupArgs)(nil)).Elem()
}

type ProjectShareGroupInput interface {
	pulumi.Input

	ToProjectShareGroupOutput() ProjectShareGroupOutput
	ToProjectShareGroupOutputWithContext(ctx context.Context) ProjectShareGroupOutput
}

func (*ProjectShareGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectShareGroup)(nil)).Elem()
}

func (i *ProjectShareGroup) ToProjectShareGroupOutput() ProjectShareGroupOutput {
	return i.ToProjectShareGroupOutputWithContext(context.Background())
}

func (i *ProjectShareGroup) ToProjectShareGroupOutputWithContext(ctx context.Context) ProjectShareGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectShareGroupOutput)
}

// ProjectShareGroupArrayInput is an input type that accepts ProjectShareGroupArray and ProjectShareGroupArrayOutput values.
// You can construct a concrete instance of `ProjectShareGroupArrayInput` via:
//
//	ProjectShareGroupArray{ ProjectShareGroupArgs{...} }
type ProjectShareGroupArrayInput interface {
	pulumi.Input

	ToProjectShareGroupArrayOutput() ProjectShareGroupArrayOutput
	ToProjectShareGroupArrayOutputWithContext(context.Context) ProjectShareGroupArrayOutput
}

type ProjectShareGroupArray []ProjectShareGroupInput

func (ProjectShareGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectShareGroup)(nil)).Elem()
}

func (i ProjectShareGroupArray) ToProjectShareGroupArrayOutput() ProjectShareGroupArrayOutput {
	return i.ToProjectShareGroupArrayOutputWithContext(context.Background())
}

func (i ProjectShareGroupArray) ToProjectShareGroupArrayOutputWithContext(ctx context.Context) ProjectShareGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectShareGroupArrayOutput)
}

// ProjectShareGroupMapInput is an input type that accepts ProjectShareGroupMap and ProjectShareGroupMapOutput values.
// You can construct a concrete instance of `ProjectShareGroupMapInput` via:
//
//	ProjectShareGroupMap{ "key": ProjectShareGroupArgs{...} }
type ProjectShareGroupMapInput interface {
	pulumi.Input

	ToProjectShareGroupMapOutput() ProjectShareGroupMapOutput
	ToProjectShareGroupMapOutputWithContext(context.Context) ProjectShareGroupMapOutput
}

type ProjectShareGroupMap map[string]ProjectShareGroupInput

func (ProjectShareGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectShareGroup)(nil)).Elem()
}

func (i ProjectShareGroupMap) ToProjectShareGroupMapOutput() ProjectShareGroupMapOutput {
	return i.ToProjectShareGroupMapOutputWithContext(context.Background())
}

func (i ProjectShareGroupMap) ToProjectShareGroupMapOutputWithContext(ctx context.Context) ProjectShareGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectShareGroupMapOutput)
}

type ProjectShareGroupOutput struct{ *pulumi.OutputState }

func (ProjectShareGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectShareGroup)(nil)).Elem()
}

func (o ProjectShareGroupOutput) ToProjectShareGroupOutput() ProjectShareGroupOutput {
	return o
}

func (o ProjectShareGroupOutput) ToProjectShareGroupOutputWithContext(ctx context.Context) ProjectShareGroupOutput {
	return o
}

// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
//
// Deprecated: Use `group_access` instead of the `access_level` attribute.
func (o ProjectShareGroupOutput) AccessLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectShareGroup) pulumi.StringPtrOutput { return v.AccessLevel }).(pulumi.StringPtrOutput)
}

// The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
func (o ProjectShareGroupOutput) GroupAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectShareGroup) pulumi.StringPtrOutput { return v.GroupAccess }).(pulumi.StringPtrOutput)
}

// The id of the group.
func (o ProjectShareGroupOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectShareGroup) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

// The ID or URL-encoded path of the project.
func (o ProjectShareGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectShareGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ProjectShareGroupArrayOutput struct{ *pulumi.OutputState }

func (ProjectShareGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectShareGroup)(nil)).Elem()
}

func (o ProjectShareGroupArrayOutput) ToProjectShareGroupArrayOutput() ProjectShareGroupArrayOutput {
	return o
}

func (o ProjectShareGroupArrayOutput) ToProjectShareGroupArrayOutputWithContext(ctx context.Context) ProjectShareGroupArrayOutput {
	return o
}

func (o ProjectShareGroupArrayOutput) Index(i pulumi.IntInput) ProjectShareGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectShareGroup {
		return vs[0].([]*ProjectShareGroup)[vs[1].(int)]
	}).(ProjectShareGroupOutput)
}

type ProjectShareGroupMapOutput struct{ *pulumi.OutputState }

func (ProjectShareGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectShareGroup)(nil)).Elem()
}

func (o ProjectShareGroupMapOutput) ToProjectShareGroupMapOutput() ProjectShareGroupMapOutput {
	return o
}

func (o ProjectShareGroupMapOutput) ToProjectShareGroupMapOutputWithContext(ctx context.Context) ProjectShareGroupMapOutput {
	return o
}

func (o ProjectShareGroupMapOutput) MapIndex(k pulumi.StringInput) ProjectShareGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectShareGroup {
		return vs[0].(map[string]*ProjectShareGroup)[vs[1].(string)]
	}).(ProjectShareGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectShareGroupInput)(nil)).Elem(), &ProjectShareGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectShareGroupArrayInput)(nil)).Elem(), ProjectShareGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectShareGroupMapInput)(nil)).Elem(), ProjectShareGroupMap{})
	pulumi.RegisterOutputType(ProjectShareGroupOutput{})
	pulumi.RegisterOutputType(ProjectShareGroupArrayOutput{})
	pulumi.RegisterOutputType(ProjectShareGroupMapOutput{})
}
