// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupVariable` resource allows to manage the lifecycle of a CI/CD variable for a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
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
//			_, err := gitlab.NewGroupVariable(ctx, "example", &gitlab.GroupVariableArgs{
//				EnvironmentScope: pulumi.String("*"),
//				Group:            pulumi.String("12345"),
//				Key:              pulumi.String("group_variable_key"),
//				Masked:           pulumi.Bool(false),
//				Protected:        pulumi.Bool(false),
//				Value:            pulumi.String("group_variable_value"),
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
// GitLab group variables can be imported using an id made up of `groupid:variablename:scope`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/groupVariable:GroupVariable example 12345:group_variable_key:*
//
// ```
type GroupVariable struct {
	pulumi.CustomResourceState

	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrOutput `pulumi:"environmentScope"`
	// The name or id of the group.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the variable.
	Key pulumi.StringOutput `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked pulumi.BoolPtrOutput `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrOutput `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw pulumi.BoolPtrOutput `pulumi:"raw"`
	// The value of the variable.
	Value pulumi.StringOutput `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType pulumi.StringPtrOutput `pulumi:"variableType"`
}

// NewGroupVariable registers a new resource with the given unique name, arguments, and options.
func NewGroupVariable(ctx *pulumi.Context,
	name string, args *GroupVariableArgs, opts ...pulumi.ResourceOption) (*GroupVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource GroupVariable
	err := ctx.RegisterResource("gitlab:index/groupVariable:GroupVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupVariable gets an existing GroupVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupVariableState, opts ...pulumi.ResourceOption) (*GroupVariable, error) {
	var resource GroupVariable
	err := ctx.ReadResource("gitlab:index/groupVariable:GroupVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupVariable resources.
type groupVariableState struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name or id of the group.
	Group *string `pulumi:"group"`
	// The name of the variable.
	Key *string `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked *bool `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected *bool `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw *bool `pulumi:"raw"`
	// The value of the variable.
	Value *string `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType *string `pulumi:"variableType"`
}

type GroupVariableState struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrInput
	// The name or id of the group.
	Group pulumi.StringPtrInput
	// The name of the variable.
	Key pulumi.StringPtrInput
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked pulumi.BoolPtrInput
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrInput
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw pulumi.BoolPtrInput
	// The value of the variable.
	Value pulumi.StringPtrInput
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType pulumi.StringPtrInput
}

func (GroupVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupVariableState)(nil)).Elem()
}

type groupVariableArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name or id of the group.
	Group string `pulumi:"group"`
	// The name of the variable.
	Key string `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked *bool `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected *bool `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw *bool `pulumi:"raw"`
	// The value of the variable.
	Value string `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType *string `pulumi:"variableType"`
}

// The set of arguments for constructing a GroupVariable resource.
type GroupVariableArgs struct {
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrInput
	// The name or id of the group.
	Group pulumi.StringInput
	// The name of the variable.
	Key pulumi.StringInput
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked pulumi.BoolPtrInput
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrInput
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw pulumi.BoolPtrInput
	// The value of the variable.
	Value pulumi.StringInput
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType pulumi.StringPtrInput
}

func (GroupVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupVariableArgs)(nil)).Elem()
}

type GroupVariableInput interface {
	pulumi.Input

	ToGroupVariableOutput() GroupVariableOutput
	ToGroupVariableOutputWithContext(ctx context.Context) GroupVariableOutput
}

func (*GroupVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupVariable)(nil)).Elem()
}

func (i *GroupVariable) ToGroupVariableOutput() GroupVariableOutput {
	return i.ToGroupVariableOutputWithContext(context.Background())
}

func (i *GroupVariable) ToGroupVariableOutputWithContext(ctx context.Context) GroupVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariableOutput)
}

// GroupVariableArrayInput is an input type that accepts GroupVariableArray and GroupVariableArrayOutput values.
// You can construct a concrete instance of `GroupVariableArrayInput` via:
//
//	GroupVariableArray{ GroupVariableArgs{...} }
type GroupVariableArrayInput interface {
	pulumi.Input

	ToGroupVariableArrayOutput() GroupVariableArrayOutput
	ToGroupVariableArrayOutputWithContext(context.Context) GroupVariableArrayOutput
}

type GroupVariableArray []GroupVariableInput

func (GroupVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupVariable)(nil)).Elem()
}

func (i GroupVariableArray) ToGroupVariableArrayOutput() GroupVariableArrayOutput {
	return i.ToGroupVariableArrayOutputWithContext(context.Background())
}

func (i GroupVariableArray) ToGroupVariableArrayOutputWithContext(ctx context.Context) GroupVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariableArrayOutput)
}

// GroupVariableMapInput is an input type that accepts GroupVariableMap and GroupVariableMapOutput values.
// You can construct a concrete instance of `GroupVariableMapInput` via:
//
//	GroupVariableMap{ "key": GroupVariableArgs{...} }
type GroupVariableMapInput interface {
	pulumi.Input

	ToGroupVariableMapOutput() GroupVariableMapOutput
	ToGroupVariableMapOutputWithContext(context.Context) GroupVariableMapOutput
}

type GroupVariableMap map[string]GroupVariableInput

func (GroupVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupVariable)(nil)).Elem()
}

func (i GroupVariableMap) ToGroupVariableMapOutput() GroupVariableMapOutput {
	return i.ToGroupVariableMapOutputWithContext(context.Background())
}

func (i GroupVariableMap) ToGroupVariableMapOutputWithContext(ctx context.Context) GroupVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariableMapOutput)
}

type GroupVariableOutput struct{ *pulumi.OutputState }

func (GroupVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupVariable)(nil)).Elem()
}

func (o GroupVariableOutput) ToGroupVariableOutput() GroupVariableOutput {
	return o
}

func (o GroupVariableOutput) ToGroupVariableOutputWithContext(ctx context.Context) GroupVariableOutput {
	return o
}

// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
func (o GroupVariableOutput) EnvironmentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.StringPtrOutput { return v.EnvironmentScope }).(pulumi.StringPtrOutput)
}

// The name or id of the group.
func (o GroupVariableOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The name of the variable.
func (o GroupVariableOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
func (o GroupVariableOutput) Masked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.BoolPtrOutput { return v.Masked }).(pulumi.BoolPtrOutput)
}

// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
func (o GroupVariableOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.BoolPtrOutput { return v.Protected }).(pulumi.BoolPtrOutput)
}

// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
func (o GroupVariableOutput) Raw() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.BoolPtrOutput { return v.Raw }).(pulumi.BoolPtrOutput)
}

// The value of the variable.
func (o GroupVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
func (o GroupVariableOutput) VariableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupVariable) pulumi.StringPtrOutput { return v.VariableType }).(pulumi.StringPtrOutput)
}

type GroupVariableArrayOutput struct{ *pulumi.OutputState }

func (GroupVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupVariable)(nil)).Elem()
}

func (o GroupVariableArrayOutput) ToGroupVariableArrayOutput() GroupVariableArrayOutput {
	return o
}

func (o GroupVariableArrayOutput) ToGroupVariableArrayOutputWithContext(ctx context.Context) GroupVariableArrayOutput {
	return o
}

func (o GroupVariableArrayOutput) Index(i pulumi.IntInput) GroupVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupVariable {
		return vs[0].([]*GroupVariable)[vs[1].(int)]
	}).(GroupVariableOutput)
}

type GroupVariableMapOutput struct{ *pulumi.OutputState }

func (GroupVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupVariable)(nil)).Elem()
}

func (o GroupVariableMapOutput) ToGroupVariableMapOutput() GroupVariableMapOutput {
	return o
}

func (o GroupVariableMapOutput) ToGroupVariableMapOutputWithContext(ctx context.Context) GroupVariableMapOutput {
	return o
}

func (o GroupVariableMapOutput) MapIndex(k pulumi.StringInput) GroupVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupVariable {
		return vs[0].(map[string]*GroupVariable)[vs[1].(string)]
	}).(GroupVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupVariableInput)(nil)).Elem(), &GroupVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupVariableArrayInput)(nil)).Elem(), GroupVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupVariableMapInput)(nil)).Elem(), GroupVariableMap{})
	pulumi.RegisterOutputType(GroupVariableOutput{})
	pulumi.RegisterOutputType(GroupVariableArrayOutput{})
	pulumi.RegisterOutputType(GroupVariableMapOutput{})
}
