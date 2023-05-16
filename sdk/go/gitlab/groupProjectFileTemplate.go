// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupProjectFileTemplate` resource allows setting a project from which
// custom file templates will be loaded. In order to use this resource, the project selected must be a direct child of
// the group selected. After the resource has run, `gitlab_project_template.template_project_id` is available for use.
// For more information about which file types are available as templates, view
// [GitLab's documentation](https://docs.gitlab.com/ee/user/group/custom_project_templates.html)
//
// > This resource requires a GitLab Enterprise instance with a Premium license.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#update-group)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v5/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			foo, err := gitlab.NewGroup(ctx, "foo", &gitlab.GroupArgs{
//				Path:        pulumi.String("group"),
//				Description: pulumi.String("An example group"),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := gitlab.NewProject(ctx, "bar", &gitlab.ProjectArgs{
//				Description:     pulumi.String("contains file templates"),
//				VisibilityLevel: pulumi.String("public"),
//				NamespaceId:     foo.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewGroupProjectFileTemplate(ctx, "templateLink", &gitlab.GroupProjectFileTemplateArgs{
//				GroupId:               foo.ID(),
//				FileTemplateProjectId: bar.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type GroupProjectFileTemplate struct {
	pulumi.CustomResourceState

	// The ID of the project that will be used for file templates. This project must be the direct
	// 			child of the project defined by the group_id
	FileTemplateProjectId pulumi.IntOutput `pulumi:"fileTemplateProjectId"`
	// The ID of the group that will use the file template project. This group must be the direct
	//             parent of the project defined by project_id
	GroupId pulumi.IntOutput `pulumi:"groupId"`
}

// NewGroupProjectFileTemplate registers a new resource with the given unique name, arguments, and options.
func NewGroupProjectFileTemplate(ctx *pulumi.Context,
	name string, args *GroupProjectFileTemplateArgs, opts ...pulumi.ResourceOption) (*GroupProjectFileTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileTemplateProjectId == nil {
		return nil, errors.New("invalid value for required argument 'FileTemplateProjectId'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	var resource GroupProjectFileTemplate
	err := ctx.RegisterResource("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupProjectFileTemplate gets an existing GroupProjectFileTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupProjectFileTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupProjectFileTemplateState, opts ...pulumi.ResourceOption) (*GroupProjectFileTemplate, error) {
	var resource GroupProjectFileTemplate
	err := ctx.ReadResource("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupProjectFileTemplate resources.
type groupProjectFileTemplateState struct {
	// The ID of the project that will be used for file templates. This project must be the direct
	// 			child of the project defined by the group_id
	FileTemplateProjectId *int `pulumi:"fileTemplateProjectId"`
	// The ID of the group that will use the file template project. This group must be the direct
	//             parent of the project defined by project_id
	GroupId *int `pulumi:"groupId"`
}

type GroupProjectFileTemplateState struct {
	// The ID of the project that will be used for file templates. This project must be the direct
	// 			child of the project defined by the group_id
	FileTemplateProjectId pulumi.IntPtrInput
	// The ID of the group that will use the file template project. This group must be the direct
	//             parent of the project defined by project_id
	GroupId pulumi.IntPtrInput
}

func (GroupProjectFileTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupProjectFileTemplateState)(nil)).Elem()
}

type groupProjectFileTemplateArgs struct {
	// The ID of the project that will be used for file templates. This project must be the direct
	// 			child of the project defined by the group_id
	FileTemplateProjectId int `pulumi:"fileTemplateProjectId"`
	// The ID of the group that will use the file template project. This group must be the direct
	//             parent of the project defined by project_id
	GroupId int `pulumi:"groupId"`
}

// The set of arguments for constructing a GroupProjectFileTemplate resource.
type GroupProjectFileTemplateArgs struct {
	// The ID of the project that will be used for file templates. This project must be the direct
	// 			child of the project defined by the group_id
	FileTemplateProjectId pulumi.IntInput
	// The ID of the group that will use the file template project. This group must be the direct
	//             parent of the project defined by project_id
	GroupId pulumi.IntInput
}

func (GroupProjectFileTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupProjectFileTemplateArgs)(nil)).Elem()
}

type GroupProjectFileTemplateInput interface {
	pulumi.Input

	ToGroupProjectFileTemplateOutput() GroupProjectFileTemplateOutput
	ToGroupProjectFileTemplateOutputWithContext(ctx context.Context) GroupProjectFileTemplateOutput
}

func (*GroupProjectFileTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupProjectFileTemplate)(nil)).Elem()
}

func (i *GroupProjectFileTemplate) ToGroupProjectFileTemplateOutput() GroupProjectFileTemplateOutput {
	return i.ToGroupProjectFileTemplateOutputWithContext(context.Background())
}

func (i *GroupProjectFileTemplate) ToGroupProjectFileTemplateOutputWithContext(ctx context.Context) GroupProjectFileTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupProjectFileTemplateOutput)
}

// GroupProjectFileTemplateArrayInput is an input type that accepts GroupProjectFileTemplateArray and GroupProjectFileTemplateArrayOutput values.
// You can construct a concrete instance of `GroupProjectFileTemplateArrayInput` via:
//
//	GroupProjectFileTemplateArray{ GroupProjectFileTemplateArgs{...} }
type GroupProjectFileTemplateArrayInput interface {
	pulumi.Input

	ToGroupProjectFileTemplateArrayOutput() GroupProjectFileTemplateArrayOutput
	ToGroupProjectFileTemplateArrayOutputWithContext(context.Context) GroupProjectFileTemplateArrayOutput
}

type GroupProjectFileTemplateArray []GroupProjectFileTemplateInput

func (GroupProjectFileTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupProjectFileTemplate)(nil)).Elem()
}

func (i GroupProjectFileTemplateArray) ToGroupProjectFileTemplateArrayOutput() GroupProjectFileTemplateArrayOutput {
	return i.ToGroupProjectFileTemplateArrayOutputWithContext(context.Background())
}

func (i GroupProjectFileTemplateArray) ToGroupProjectFileTemplateArrayOutputWithContext(ctx context.Context) GroupProjectFileTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupProjectFileTemplateArrayOutput)
}

// GroupProjectFileTemplateMapInput is an input type that accepts GroupProjectFileTemplateMap and GroupProjectFileTemplateMapOutput values.
// You can construct a concrete instance of `GroupProjectFileTemplateMapInput` via:
//
//	GroupProjectFileTemplateMap{ "key": GroupProjectFileTemplateArgs{...} }
type GroupProjectFileTemplateMapInput interface {
	pulumi.Input

	ToGroupProjectFileTemplateMapOutput() GroupProjectFileTemplateMapOutput
	ToGroupProjectFileTemplateMapOutputWithContext(context.Context) GroupProjectFileTemplateMapOutput
}

type GroupProjectFileTemplateMap map[string]GroupProjectFileTemplateInput

func (GroupProjectFileTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupProjectFileTemplate)(nil)).Elem()
}

func (i GroupProjectFileTemplateMap) ToGroupProjectFileTemplateMapOutput() GroupProjectFileTemplateMapOutput {
	return i.ToGroupProjectFileTemplateMapOutputWithContext(context.Background())
}

func (i GroupProjectFileTemplateMap) ToGroupProjectFileTemplateMapOutputWithContext(ctx context.Context) GroupProjectFileTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupProjectFileTemplateMapOutput)
}

type GroupProjectFileTemplateOutput struct{ *pulumi.OutputState }

func (GroupProjectFileTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupProjectFileTemplate)(nil)).Elem()
}

func (o GroupProjectFileTemplateOutput) ToGroupProjectFileTemplateOutput() GroupProjectFileTemplateOutput {
	return o
}

func (o GroupProjectFileTemplateOutput) ToGroupProjectFileTemplateOutputWithContext(ctx context.Context) GroupProjectFileTemplateOutput {
	return o
}

// The ID of the project that will be used for file templates. This project must be the direct
//
//	child of the project defined by the group_id
func (o GroupProjectFileTemplateOutput) FileTemplateProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupProjectFileTemplate) pulumi.IntOutput { return v.FileTemplateProjectId }).(pulumi.IntOutput)
}

// The ID of the group that will use the file template project. This group must be the direct
//
//	parent of the project defined by project_id
func (o GroupProjectFileTemplateOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupProjectFileTemplate) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

type GroupProjectFileTemplateArrayOutput struct{ *pulumi.OutputState }

func (GroupProjectFileTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupProjectFileTemplate)(nil)).Elem()
}

func (o GroupProjectFileTemplateArrayOutput) ToGroupProjectFileTemplateArrayOutput() GroupProjectFileTemplateArrayOutput {
	return o
}

func (o GroupProjectFileTemplateArrayOutput) ToGroupProjectFileTemplateArrayOutputWithContext(ctx context.Context) GroupProjectFileTemplateArrayOutput {
	return o
}

func (o GroupProjectFileTemplateArrayOutput) Index(i pulumi.IntInput) GroupProjectFileTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupProjectFileTemplate {
		return vs[0].([]*GroupProjectFileTemplate)[vs[1].(int)]
	}).(GroupProjectFileTemplateOutput)
}

type GroupProjectFileTemplateMapOutput struct{ *pulumi.OutputState }

func (GroupProjectFileTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupProjectFileTemplate)(nil)).Elem()
}

func (o GroupProjectFileTemplateMapOutput) ToGroupProjectFileTemplateMapOutput() GroupProjectFileTemplateMapOutput {
	return o
}

func (o GroupProjectFileTemplateMapOutput) ToGroupProjectFileTemplateMapOutputWithContext(ctx context.Context) GroupProjectFileTemplateMapOutput {
	return o
}

func (o GroupProjectFileTemplateMapOutput) MapIndex(k pulumi.StringInput) GroupProjectFileTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupProjectFileTemplate {
		return vs[0].(map[string]*GroupProjectFileTemplate)[vs[1].(string)]
	}).(GroupProjectFileTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupProjectFileTemplateInput)(nil)).Elem(), &GroupProjectFileTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupProjectFileTemplateArrayInput)(nil)).Elem(), GroupProjectFileTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupProjectFileTemplateMapInput)(nil)).Elem(), GroupProjectFileTemplateMap{})
	pulumi.RegisterOutputType(GroupProjectFileTemplateOutput{})
	pulumi.RegisterOutputType(GroupProjectFileTemplateArrayOutput{})
	pulumi.RegisterOutputType(GroupProjectFileTemplateMapOutput{})
}
