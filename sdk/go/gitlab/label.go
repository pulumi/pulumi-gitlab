// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Label` resource manages the lifecycle of a project label.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/labels/#get-a-single-project-label)
type Label struct {
	pulumi.CustomResourceState

	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringOutput `pulumi:"color"`
	// Read-only, used by the provider to store the API response color. This is always in the 6-digit hex notation with leading '#' sign (e.g. #FFAABB). If `color` contains a color name, this attribute contains the hex notation equivalent. Otherwise, the value of this attribute is the same as `color`.
	ColorHex pulumi.StringOutput `pulumi:"colorHex"`
	// The description of the label.
	Description pulumi.StringOutput `pulumi:"description"`
	// The id of the project label.
	LabelId pulumi.IntOutput `pulumi:"labelId"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Read-only, used by the provider to store the API response color. This is always in the 6-digit hex notation with leading '#' sign (e.g. #FFAABB). If `color` contains a color name, this attribute contains the hex notation equivalent. Otherwise, the value of this attribute is the same as `color`.
	ColorHex *string `pulumi:"colorHex"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The id of the project label.
	LabelId *int `pulumi:"labelId"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project *string `pulumi:"project"`
}

type LabelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringPtrInput
	// Read-only, used by the provider to store the API response color. This is always in the 6-digit hex notation with leading '#' sign (e.g. #FFAABB). If `color` contains a color name, this attribute contains the hex notation equivalent. Otherwise, the value of this attribute is the same as `color`.
	ColorHex pulumi.StringPtrInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The id of the project label.
	LabelId pulumi.IntPtrInput
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
	return reflect.TypeOf((**Label)(nil)).Elem()
}

func (i *Label) ToLabelOutput() LabelOutput {
	return i.ToLabelOutputWithContext(context.Background())
}

func (i *Label) ToLabelOutputWithContext(ctx context.Context) LabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelOutput)
}

// LabelArrayInput is an input type that accepts LabelArray and LabelArrayOutput values.
// You can construct a concrete instance of `LabelArrayInput` via:
//
//	LabelArray{ LabelArgs{...} }
type LabelArrayInput interface {
	pulumi.Input

	ToLabelArrayOutput() LabelArrayOutput
	ToLabelArrayOutputWithContext(context.Context) LabelArrayOutput
}

type LabelArray []LabelInput

func (LabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Label)(nil)).Elem()
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
//	LabelMap{ "key": LabelArgs{...} }
type LabelMapInput interface {
	pulumi.Input

	ToLabelMapOutput() LabelMapOutput
	ToLabelMapOutputWithContext(context.Context) LabelMapOutput
}

type LabelMap map[string]LabelInput

func (LabelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Label)(nil)).Elem()
}

func (i LabelMap) ToLabelMapOutput() LabelMapOutput {
	return i.ToLabelMapOutputWithContext(context.Background())
}

func (i LabelMap) ToLabelMapOutputWithContext(ctx context.Context) LabelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelMapOutput)
}

type LabelOutput struct{ *pulumi.OutputState }

func (LabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Label)(nil)).Elem()
}

func (o LabelOutput) ToLabelOutput() LabelOutput {
	return o
}

func (o LabelOutput) ToLabelOutputWithContext(ctx context.Context) LabelOutput {
	return o
}

// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
func (o LabelOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.Color }).(pulumi.StringOutput)
}

// Read-only, used by the provider to store the API response color. This is always in the 6-digit hex notation with leading '#' sign (e.g. #FFAABB). If `color` contains a color name, this attribute contains the hex notation equivalent. Otherwise, the value of this attribute is the same as `color`.
func (o LabelOutput) ColorHex() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.ColorHex }).(pulumi.StringOutput)
}

// The description of the label.
func (o LabelOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The id of the project label.
func (o LabelOutput) LabelId() pulumi.IntOutput {
	return o.ApplyT(func(v *Label) pulumi.IntOutput { return v.LabelId }).(pulumi.IntOutput)
}

// The name of the label.
func (o LabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name or id of the project to add the label to.
func (o LabelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type LabelArrayOutput struct{ *pulumi.OutputState }

func (LabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Label)(nil)).Elem()
}

func (o LabelArrayOutput) ToLabelArrayOutput() LabelArrayOutput {
	return o
}

func (o LabelArrayOutput) ToLabelArrayOutputWithContext(ctx context.Context) LabelArrayOutput {
	return o
}

func (o LabelArrayOutput) Index(i pulumi.IntInput) LabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Label {
		return vs[0].([]*Label)[vs[1].(int)]
	}).(LabelOutput)
}

type LabelMapOutput struct{ *pulumi.OutputState }

func (LabelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Label)(nil)).Elem()
}

func (o LabelMapOutput) ToLabelMapOutput() LabelMapOutput {
	return o
}

func (o LabelMapOutput) ToLabelMapOutputWithContext(ctx context.Context) LabelMapOutput {
	return o
}

func (o LabelMapOutput) MapIndex(k pulumi.StringInput) LabelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Label {
		return vs[0].(map[string]*Label)[vs[1].(string)]
	}).(LabelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LabelInput)(nil)).Elem(), &Label{})
	pulumi.RegisterInputType(reflect.TypeOf((*LabelArrayInput)(nil)).Elem(), LabelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LabelMapInput)(nil)).Elem(), LabelMap{})
	pulumi.RegisterOutputType(LabelOutput{})
	pulumi.RegisterOutputType(LabelArrayOutput{})
	pulumi.RegisterOutputType(LabelMapOutput{})
}
