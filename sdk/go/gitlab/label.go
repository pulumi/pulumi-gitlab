// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_label
//
// This resource allows you to create and manage labels for your GitLab projects.
// For further information on labels, consult the [gitlab
// documentation](https://docs.gitlab.com/ee/user/project/labels.html#project-labels).
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
// 		_, err := gitlab.NewLabel(ctx, "fixme", &gitlab.LabelArgs{
// 			Color:       pulumi.String("#ffcc00"),
// 			Description: pulumi.String("issue with failing tests"),
// 			Project:     pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Label struct {
	pulumi.CustomResourceState

	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringOutput `pulumi:"color"`
	// The description of the label.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the label.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewLabel registers a new resource with the given unique name, arguments, and options.
func NewLabel(ctx *pulumi.Context,
	name string, args *LabelArgs, opts ...pulumi.ResourceOption) (*Label, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Color == nil {
		return nil, errors.New("invalid value for required argument 'Color'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Label
	err := ctx.RegisterResource("gitlab:index/label:Label", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabel gets an existing Label resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelState, opts ...pulumi.ResourceOption) (*Label, error) {
	var resource Label
	err := ctx.ReadResource("gitlab:index/label:Label", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Label resources.
type labelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color *string `pulumi:"color"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project *string `pulumi:"project"`
}

type LabelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringPtrInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the label to.
	Project pulumi.StringPtrInput
}

func (LabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelState)(nil)).Elem()
}

type labelArgs struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color string `pulumi:"color"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a Label resource.
type LabelArgs struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the label to.
	Project pulumi.StringInput
}

func (LabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelArgs)(nil)).Elem()
}

type LabelInput interface {
	pulumi.Input

	ToLabelOutput() LabelOutput
	ToLabelOutputWithContext(ctx context.Context) LabelOutput
}

func (*Label) ElementType() reflect.Type {
	return reflect.TypeOf((*Label)(nil))
}

func (i *Label) ToLabelOutput() LabelOutput {
	return i.ToLabelOutputWithContext(context.Background())
}

func (i *Label) ToLabelOutputWithContext(ctx context.Context) LabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelOutput)
}

func (i *Label) ToLabelPtrOutput() LabelPtrOutput {
	return i.ToLabelPtrOutputWithContext(context.Background())
}

func (i *Label) ToLabelPtrOutputWithContext(ctx context.Context) LabelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelPtrOutput)
}

type LabelPtrInput interface {
	pulumi.Input

	ToLabelPtrOutput() LabelPtrOutput
	ToLabelPtrOutputWithContext(ctx context.Context) LabelPtrOutput
}

type labelPtrType LabelArgs

func (*labelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Label)(nil))
}

func (i *labelPtrType) ToLabelPtrOutput() LabelPtrOutput {
	return i.ToLabelPtrOutputWithContext(context.Background())
}

func (i *labelPtrType) ToLabelPtrOutputWithContext(ctx context.Context) LabelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelPtrOutput)
}

// LabelArrayInput is an input type that accepts LabelArray and LabelArrayOutput values.
// You can construct a concrete instance of `LabelArrayInput` via:
//
//          LabelArray{ LabelArgs{...} }
type LabelArrayInput interface {
	pulumi.Input

	ToLabelArrayOutput() LabelArrayOutput
	ToLabelArrayOutputWithContext(context.Context) LabelArrayOutput
}

type LabelArray []LabelInput

func (LabelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Label)(nil))
}

func (i LabelArray) ToLabelArrayOutput() LabelArrayOutput {
	return i.ToLabelArrayOutputWithContext(context.Background())
}

func (i LabelArray) ToLabelArrayOutputWithContext(ctx context.Context) LabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelArrayOutput)
}

// LabelMapInput is an input type that accepts LabelMap and LabelMapOutput values.
// You can construct a concrete instance of `LabelMapInput` via:
//
//          LabelMap{ "key": LabelArgs{...} }
type LabelMapInput interface {
	pulumi.Input

	ToLabelMapOutput() LabelMapOutput
	ToLabelMapOutputWithContext(context.Context) LabelMapOutput
}

type LabelMap map[string]LabelInput

func (LabelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Label)(nil))
}

func (i LabelMap) ToLabelMapOutput() LabelMapOutput {
	return i.ToLabelMapOutputWithContext(context.Background())
}

func (i LabelMap) ToLabelMapOutputWithContext(ctx context.Context) LabelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelMapOutput)
}

type LabelOutput struct {
	*pulumi.OutputState
}

func (LabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Label)(nil))
}

func (o LabelOutput) ToLabelOutput() LabelOutput {
	return o
}

func (o LabelOutput) ToLabelOutputWithContext(ctx context.Context) LabelOutput {
	return o
}

func (o LabelOutput) ToLabelPtrOutput() LabelPtrOutput {
	return o.ToLabelPtrOutputWithContext(context.Background())
}

func (o LabelOutput) ToLabelPtrOutputWithContext(ctx context.Context) LabelPtrOutput {
	return o.ApplyT(func(v Label) *Label {
		return &v
	}).(LabelPtrOutput)
}

type LabelPtrOutput struct {
	*pulumi.OutputState
}

func (LabelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Label)(nil))
}

func (o LabelPtrOutput) ToLabelPtrOutput() LabelPtrOutput {
	return o
}

func (o LabelPtrOutput) ToLabelPtrOutputWithContext(ctx context.Context) LabelPtrOutput {
	return o
}

type LabelArrayOutput struct{ *pulumi.OutputState }

func (LabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Label)(nil))
}

func (o LabelArrayOutput) ToLabelArrayOutput() LabelArrayOutput {
	return o
}

func (o LabelArrayOutput) ToLabelArrayOutputWithContext(ctx context.Context) LabelArrayOutput {
	return o
}

func (o LabelArrayOutput) Index(i pulumi.IntInput) LabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Label {
		return vs[0].([]Label)[vs[1].(int)]
	}).(LabelOutput)
}

type LabelMapOutput struct{ *pulumi.OutputState }

func (LabelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Label)(nil))
}

func (o LabelMapOutput) ToLabelMapOutput() LabelMapOutput {
	return o
}

func (o LabelMapOutput) ToLabelMapOutputWithContext(ctx context.Context) LabelMapOutput {
	return o
}

func (o LabelMapOutput) MapIndex(k pulumi.StringInput) LabelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Label {
		return vs[0].(map[string]Label)[vs[1].(string)]
	}).(LabelOutput)
}

func init() {
	pulumi.RegisterOutputType(LabelOutput{})
	pulumi.RegisterOutputType(LabelPtrOutput{})
	pulumi.RegisterOutputType(LabelArrayOutput{})
	pulumi.RegisterOutputType(LabelMapOutput{})
}
