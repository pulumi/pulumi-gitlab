// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_group\_variable
//
// This resource allows you to create and manage CI/CD variables for your GitLab groups.
// For further information on variables, consult the [gitlab
// documentation](https://docs.gitlab.com/ce/ci/variables/README.html#variables).
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
// 		_, err := gitlab.NewGroupVariable(ctx, "example", &gitlab.GroupVariableArgs{
// 			Group:     pulumi.String("12345"),
// 			Key:       pulumi.String("group_variable_key"),
// 			Masked:    pulumi.Bool(false),
// 			Protected: pulumi.Bool(false),
// 			Value:     pulumi.String("group_variable_value"),
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
// GitLab group variables can be imported using an id made up of `groupid:variablename`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/groupVariable:GroupVariable example 12345:group_variable_key
// ```
type GroupVariable struct {
	pulumi.CustomResourceState

	// The name or id of the group to add the hook to.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the variable.
	Key    pulumi.StringOutput  `pulumi:"key"`
	Masked pulumi.BoolPtrOutput `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrOutput `pulumi:"protected"`
	// The value of the variable.
	Value pulumi.StringOutput `pulumi:"value"`
	// The type of a variable. Available types are: envVar (default) and file.
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
	// The name or id of the group to add the hook to.
	Group *string `pulumi:"group"`
	// The name of the variable.
	Key    *string `pulumi:"key"`
	Masked *bool   `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected *bool `pulumi:"protected"`
	// The value of the variable.
	Value *string `pulumi:"value"`
	// The type of a variable. Available types are: envVar (default) and file.
	VariableType *string `pulumi:"variableType"`
}

type GroupVariableState struct {
	// The name or id of the group to add the hook to.
	Group pulumi.StringPtrInput
	// The name of the variable.
	Key    pulumi.StringPtrInput
	Masked pulumi.BoolPtrInput
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrInput
	// The value of the variable.
	Value pulumi.StringPtrInput
	// The type of a variable. Available types are: envVar (default) and file.
	VariableType pulumi.StringPtrInput
}

func (GroupVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupVariableState)(nil)).Elem()
}

type groupVariableArgs struct {
	// The name or id of the group to add the hook to.
	Group string `pulumi:"group"`
	// The name of the variable.
	Key    string `pulumi:"key"`
	Masked *bool  `pulumi:"masked"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected *bool `pulumi:"protected"`
	// The value of the variable.
	Value string `pulumi:"value"`
	// The type of a variable. Available types are: envVar (default) and file.
	VariableType *string `pulumi:"variableType"`
}

// The set of arguments for constructing a GroupVariable resource.
type GroupVariableArgs struct {
	// The name or id of the group to add the hook to.
	Group pulumi.StringInput
	// The name of the variable.
	Key    pulumi.StringInput
	Masked pulumi.BoolPtrInput
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrInput
	// The value of the variable.
	Value pulumi.StringInput
	// The type of a variable. Available types are: envVar (default) and file.
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
	return reflect.TypeOf((*GroupVariable)(nil))
}

func (i *GroupVariable) ToGroupVariableOutput() GroupVariableOutput {
	return i.ToGroupVariableOutputWithContext(context.Background())
}

func (i *GroupVariable) ToGroupVariableOutputWithContext(ctx context.Context) GroupVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariableOutput)
}

func (i *GroupVariable) ToGroupVariablePtrOutput() GroupVariablePtrOutput {
	return i.ToGroupVariablePtrOutputWithContext(context.Background())
}

func (i *GroupVariable) ToGroupVariablePtrOutputWithContext(ctx context.Context) GroupVariablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariablePtrOutput)
}

type GroupVariablePtrInput interface {
	pulumi.Input

	ToGroupVariablePtrOutput() GroupVariablePtrOutput
	ToGroupVariablePtrOutputWithContext(ctx context.Context) GroupVariablePtrOutput
}

type groupVariablePtrType GroupVariableArgs

func (*groupVariablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupVariable)(nil))
}

func (i *groupVariablePtrType) ToGroupVariablePtrOutput() GroupVariablePtrOutput {
	return i.ToGroupVariablePtrOutputWithContext(context.Background())
}

func (i *groupVariablePtrType) ToGroupVariablePtrOutputWithContext(ctx context.Context) GroupVariablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariablePtrOutput)
}

// GroupVariableArrayInput is an input type that accepts GroupVariableArray and GroupVariableArrayOutput values.
// You can construct a concrete instance of `GroupVariableArrayInput` via:
//
//          GroupVariableArray{ GroupVariableArgs{...} }
type GroupVariableArrayInput interface {
	pulumi.Input

	ToGroupVariableArrayOutput() GroupVariableArrayOutput
	ToGroupVariableArrayOutputWithContext(context.Context) GroupVariableArrayOutput
}

type GroupVariableArray []GroupVariableInput

func (GroupVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GroupVariable)(nil))
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
//          GroupVariableMap{ "key": GroupVariableArgs{...} }
type GroupVariableMapInput interface {
	pulumi.Input

	ToGroupVariableMapOutput() GroupVariableMapOutput
	ToGroupVariableMapOutputWithContext(context.Context) GroupVariableMapOutput
}

type GroupVariableMap map[string]GroupVariableInput

func (GroupVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GroupVariable)(nil))
}

func (i GroupVariableMap) ToGroupVariableMapOutput() GroupVariableMapOutput {
	return i.ToGroupVariableMapOutputWithContext(context.Background())
}

func (i GroupVariableMap) ToGroupVariableMapOutputWithContext(ctx context.Context) GroupVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupVariableMapOutput)
}

type GroupVariableOutput struct {
	*pulumi.OutputState
}

func (GroupVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupVariable)(nil))
}

func (o GroupVariableOutput) ToGroupVariableOutput() GroupVariableOutput {
	return o
}

func (o GroupVariableOutput) ToGroupVariableOutputWithContext(ctx context.Context) GroupVariableOutput {
	return o
}

func (o GroupVariableOutput) ToGroupVariablePtrOutput() GroupVariablePtrOutput {
	return o.ToGroupVariablePtrOutputWithContext(context.Background())
}

func (o GroupVariableOutput) ToGroupVariablePtrOutputWithContext(ctx context.Context) GroupVariablePtrOutput {
	return o.ApplyT(func(v GroupVariable) *GroupVariable {
		return &v
	}).(GroupVariablePtrOutput)
}

type GroupVariablePtrOutput struct {
	*pulumi.OutputState
}

func (GroupVariablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupVariable)(nil))
}

func (o GroupVariablePtrOutput) ToGroupVariablePtrOutput() GroupVariablePtrOutput {
	return o
}

func (o GroupVariablePtrOutput) ToGroupVariablePtrOutputWithContext(ctx context.Context) GroupVariablePtrOutput {
	return o
}

type GroupVariableArrayOutput struct{ *pulumi.OutputState }

func (GroupVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupVariable)(nil))
}

func (o GroupVariableArrayOutput) ToGroupVariableArrayOutput() GroupVariableArrayOutput {
	return o
}

func (o GroupVariableArrayOutput) ToGroupVariableArrayOutputWithContext(ctx context.Context) GroupVariableArrayOutput {
	return o
}

func (o GroupVariableArrayOutput) Index(i pulumi.IntInput) GroupVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupVariable {
		return vs[0].([]GroupVariable)[vs[1].(int)]
	}).(GroupVariableOutput)
}

type GroupVariableMapOutput struct{ *pulumi.OutputState }

func (GroupVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupVariable)(nil))
}

func (o GroupVariableMapOutput) ToGroupVariableMapOutput() GroupVariableMapOutput {
	return o
}

func (o GroupVariableMapOutput) ToGroupVariableMapOutputWithContext(ctx context.Context) GroupVariableMapOutput {
	return o
}

func (o GroupVariableMapOutput) MapIndex(k pulumi.StringInput) GroupVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupVariable {
		return vs[0].(map[string]GroupVariable)[vs[1].(string)]
	}).(GroupVariableOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupVariableOutput{})
	pulumi.RegisterOutputType(GroupVariablePtrOutput{})
	pulumi.RegisterOutputType(GroupVariableArrayOutput{})
	pulumi.RegisterOutputType(GroupVariableMapOutput{})
}
