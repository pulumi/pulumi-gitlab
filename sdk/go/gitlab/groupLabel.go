// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_group\_label
//
// This resource allows you to create and manage labels for your GitLab groups.
// For further information on labels, consult the [gitlab
// documentation](https://docs.gitlab.com/ee/user/project/labels.html#group-labels).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewGroupLabel(ctx, "fixme", &gitlab.GroupLabelArgs{
// 			Color:       pulumi.String("#ffcc00"),
// 			Description: pulumi.String("issue with failing tests"),
// 			Group:       pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Gitlab group labels can be imported using an id made up of `{group_id}:{group_label_id}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/groupLabel:GroupLabel example 12345:fixme
// ```
type GroupLabel struct {
	pulumi.CustomResourceState

	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringOutput `pulumi:"color"`
	// The description of the label.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name or id of the group to add the label to.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the label.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGroupLabel registers a new resource with the given unique name, arguments, and options.
func NewGroupLabel(ctx *pulumi.Context,
	name string, args *GroupLabelArgs, opts ...pulumi.ResourceOption) (*GroupLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Color == nil {
		return nil, errors.New("invalid value for required argument 'Color'")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	var resource GroupLabel
	err := ctx.RegisterResource("gitlab:index/groupLabel:GroupLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupLabel gets an existing GroupLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupLabelState, opts ...pulumi.ResourceOption) (*GroupLabel, error) {
	var resource GroupLabel
	err := ctx.ReadResource("gitlab:index/groupLabel:GroupLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupLabel resources.
type groupLabelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color *string `pulumi:"color"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The name or id of the group to add the label to.
	Group *string `pulumi:"group"`
	// The name of the label.
	Name *string `pulumi:"name"`
}

type GroupLabelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringPtrInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The name or id of the group to add the label to.
	Group pulumi.StringPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
}

func (GroupLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLabelState)(nil)).Elem()
}

type groupLabelArgs struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color string `pulumi:"color"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The name or id of the group to add the label to.
	Group string `pulumi:"group"`
	// The name of the label.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a GroupLabel resource.
type GroupLabelArgs struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The name or id of the group to add the label to.
	Group pulumi.StringInput
	// The name of the label.
	Name pulumi.StringPtrInput
}

func (GroupLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLabelArgs)(nil)).Elem()
}

type GroupLabelInput interface {
	pulumi.Input

	ToGroupLabelOutput() GroupLabelOutput
	ToGroupLabelOutputWithContext(ctx context.Context) GroupLabelOutput
}

func (*GroupLabel) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupLabel)(nil))
}

func (i *GroupLabel) ToGroupLabelOutput() GroupLabelOutput {
	return i.ToGroupLabelOutputWithContext(context.Background())
}

func (i *GroupLabel) ToGroupLabelOutputWithContext(ctx context.Context) GroupLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLabelOutput)
}

func (i *GroupLabel) ToGroupLabelPtrOutput() GroupLabelPtrOutput {
	return i.ToGroupLabelPtrOutputWithContext(context.Background())
}

func (i *GroupLabel) ToGroupLabelPtrOutputWithContext(ctx context.Context) GroupLabelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLabelPtrOutput)
}

type GroupLabelPtrInput interface {
	pulumi.Input

	ToGroupLabelPtrOutput() GroupLabelPtrOutput
	ToGroupLabelPtrOutputWithContext(ctx context.Context) GroupLabelPtrOutput
}

type groupLabelPtrType GroupLabelArgs

func (*groupLabelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLabel)(nil))
}

func (i *groupLabelPtrType) ToGroupLabelPtrOutput() GroupLabelPtrOutput {
	return i.ToGroupLabelPtrOutputWithContext(context.Background())
}

func (i *groupLabelPtrType) ToGroupLabelPtrOutputWithContext(ctx context.Context) GroupLabelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLabelPtrOutput)
}

// GroupLabelArrayInput is an input type that accepts GroupLabelArray and GroupLabelArrayOutput values.
// You can construct a concrete instance of `GroupLabelArrayInput` via:
//
//          GroupLabelArray{ GroupLabelArgs{...} }
type GroupLabelArrayInput interface {
	pulumi.Input

	ToGroupLabelArrayOutput() GroupLabelArrayOutput
	ToGroupLabelArrayOutputWithContext(context.Context) GroupLabelArrayOutput
}

type GroupLabelArray []GroupLabelInput

func (GroupLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GroupLabel)(nil))
}

func (i GroupLabelArray) ToGroupLabelArrayOutput() GroupLabelArrayOutput {
	return i.ToGroupLabelArrayOutputWithContext(context.Background())
}

func (i GroupLabelArray) ToGroupLabelArrayOutputWithContext(ctx context.Context) GroupLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLabelArrayOutput)
}

// GroupLabelMapInput is an input type that accepts GroupLabelMap and GroupLabelMapOutput values.
// You can construct a concrete instance of `GroupLabelMapInput` via:
//
//          GroupLabelMap{ "key": GroupLabelArgs{...} }
type GroupLabelMapInput interface {
	pulumi.Input

	ToGroupLabelMapOutput() GroupLabelMapOutput
	ToGroupLabelMapOutputWithContext(context.Context) GroupLabelMapOutput
}

type GroupLabelMap map[string]GroupLabelInput

func (GroupLabelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GroupLabel)(nil))
}

func (i GroupLabelMap) ToGroupLabelMapOutput() GroupLabelMapOutput {
	return i.ToGroupLabelMapOutputWithContext(context.Background())
}

func (i GroupLabelMap) ToGroupLabelMapOutputWithContext(ctx context.Context) GroupLabelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLabelMapOutput)
}

type GroupLabelOutput struct {
	*pulumi.OutputState
}

func (GroupLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupLabel)(nil))
}

func (o GroupLabelOutput) ToGroupLabelOutput() GroupLabelOutput {
	return o
}

func (o GroupLabelOutput) ToGroupLabelOutputWithContext(ctx context.Context) GroupLabelOutput {
	return o
}

func (o GroupLabelOutput) ToGroupLabelPtrOutput() GroupLabelPtrOutput {
	return o.ToGroupLabelPtrOutputWithContext(context.Background())
}

func (o GroupLabelOutput) ToGroupLabelPtrOutputWithContext(ctx context.Context) GroupLabelPtrOutput {
	return o.ApplyT(func(v GroupLabel) *GroupLabel {
		return &v
	}).(GroupLabelPtrOutput)
}

type GroupLabelPtrOutput struct {
	*pulumi.OutputState
}

func (GroupLabelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLabel)(nil))
}

func (o GroupLabelPtrOutput) ToGroupLabelPtrOutput() GroupLabelPtrOutput {
	return o
}

func (o GroupLabelPtrOutput) ToGroupLabelPtrOutputWithContext(ctx context.Context) GroupLabelPtrOutput {
	return o
}

type GroupLabelArrayOutput struct{ *pulumi.OutputState }

func (GroupLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupLabel)(nil))
}

func (o GroupLabelArrayOutput) ToGroupLabelArrayOutput() GroupLabelArrayOutput {
	return o
}

func (o GroupLabelArrayOutput) ToGroupLabelArrayOutputWithContext(ctx context.Context) GroupLabelArrayOutput {
	return o
}

func (o GroupLabelArrayOutput) Index(i pulumi.IntInput) GroupLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupLabel {
		return vs[0].([]GroupLabel)[vs[1].(int)]
	}).(GroupLabelOutput)
}

type GroupLabelMapOutput struct{ *pulumi.OutputState }

func (GroupLabelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupLabel)(nil))
}

func (o GroupLabelMapOutput) ToGroupLabelMapOutput() GroupLabelMapOutput {
	return o
}

func (o GroupLabelMapOutput) ToGroupLabelMapOutputWithContext(ctx context.Context) GroupLabelMapOutput {
	return o
}

func (o GroupLabelMapOutput) MapIndex(k pulumi.StringInput) GroupLabelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupLabel {
		return vs[0].(map[string]GroupLabel)[vs[1].(string)]
	}).(GroupLabelOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupLabelOutput{})
	pulumi.RegisterOutputType(GroupLabelPtrOutput{})
	pulumi.RegisterOutputType(GroupLabelArrayOutput{})
	pulumi.RegisterOutputType(GroupLabelMapOutput{})
}
