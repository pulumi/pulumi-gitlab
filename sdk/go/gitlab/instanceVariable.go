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

// The `InstanceVariable` resource allows to manage the lifecycle of an instance-level CI/CD variable.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := gitlab.NewInstanceVariable(ctx, "example", &gitlab.InstanceVariableArgs{
//				Key:       pulumi.String("instance_variable_key"),
//				Value:     pulumi.String("instance_variable_value"),
//				Protected: pulumi.Bool(false),
//				Masked:    pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// GitLab instance variables can be imported using an id made up of `variablename`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/instanceVariable:InstanceVariable example instance_variable_key
// ```
type InstanceVariable struct {
	pulumi.CustomResourceState

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

// NewInstanceVariable registers a new resource with the given unique name, arguments, and options.
func NewInstanceVariable(ctx *pulumi.Context,
	name string, args *InstanceVariableArgs, opts ...pulumi.ResourceOption) (*InstanceVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceVariable
	err := ctx.RegisterResource("gitlab:index/instanceVariable:InstanceVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceVariable gets an existing InstanceVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceVariableState, opts ...pulumi.ResourceOption) (*InstanceVariable, error) {
	var resource InstanceVariable
	err := ctx.ReadResource("gitlab:index/instanceVariable:InstanceVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceVariable resources.
type instanceVariableState struct {
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

type InstanceVariableState struct {
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

func (InstanceVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceVariableState)(nil)).Elem()
}

type instanceVariableArgs struct {
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

// The set of arguments for constructing a InstanceVariable resource.
type InstanceVariableArgs struct {
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

func (InstanceVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceVariableArgs)(nil)).Elem()
}

type InstanceVariableInput interface {
	pulumi.Input

	ToInstanceVariableOutput() InstanceVariableOutput
	ToInstanceVariableOutputWithContext(ctx context.Context) InstanceVariableOutput
}

func (*InstanceVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceVariable)(nil)).Elem()
}

func (i *InstanceVariable) ToInstanceVariableOutput() InstanceVariableOutput {
	return i.ToInstanceVariableOutputWithContext(context.Background())
}

func (i *InstanceVariable) ToInstanceVariableOutputWithContext(ctx context.Context) InstanceVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceVariableOutput)
}

// InstanceVariableArrayInput is an input type that accepts InstanceVariableArray and InstanceVariableArrayOutput values.
// You can construct a concrete instance of `InstanceVariableArrayInput` via:
//
//	InstanceVariableArray{ InstanceVariableArgs{...} }
type InstanceVariableArrayInput interface {
	pulumi.Input

	ToInstanceVariableArrayOutput() InstanceVariableArrayOutput
	ToInstanceVariableArrayOutputWithContext(context.Context) InstanceVariableArrayOutput
}

type InstanceVariableArray []InstanceVariableInput

func (InstanceVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceVariable)(nil)).Elem()
}

func (i InstanceVariableArray) ToInstanceVariableArrayOutput() InstanceVariableArrayOutput {
	return i.ToInstanceVariableArrayOutputWithContext(context.Background())
}

func (i InstanceVariableArray) ToInstanceVariableArrayOutputWithContext(ctx context.Context) InstanceVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceVariableArrayOutput)
}

// InstanceVariableMapInput is an input type that accepts InstanceVariableMap and InstanceVariableMapOutput values.
// You can construct a concrete instance of `InstanceVariableMapInput` via:
//
//	InstanceVariableMap{ "key": InstanceVariableArgs{...} }
type InstanceVariableMapInput interface {
	pulumi.Input

	ToInstanceVariableMapOutput() InstanceVariableMapOutput
	ToInstanceVariableMapOutputWithContext(context.Context) InstanceVariableMapOutput
}

type InstanceVariableMap map[string]InstanceVariableInput

func (InstanceVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceVariable)(nil)).Elem()
}

func (i InstanceVariableMap) ToInstanceVariableMapOutput() InstanceVariableMapOutput {
	return i.ToInstanceVariableMapOutputWithContext(context.Background())
}

func (i InstanceVariableMap) ToInstanceVariableMapOutputWithContext(ctx context.Context) InstanceVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceVariableMapOutput)
}

type InstanceVariableOutput struct{ *pulumi.OutputState }

func (InstanceVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceVariable)(nil)).Elem()
}

func (o InstanceVariableOutput) ToInstanceVariableOutput() InstanceVariableOutput {
	return o
}

func (o InstanceVariableOutput) ToInstanceVariableOutputWithContext(ctx context.Context) InstanceVariableOutput {
	return o
}

// The name of the variable.
func (o InstanceVariableOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceVariable) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
func (o InstanceVariableOutput) Masked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceVariable) pulumi.BoolPtrOutput { return v.Masked }).(pulumi.BoolPtrOutput)
}

// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
func (o InstanceVariableOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceVariable) pulumi.BoolPtrOutput { return v.Protected }).(pulumi.BoolPtrOutput)
}

// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
func (o InstanceVariableOutput) Raw() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceVariable) pulumi.BoolPtrOutput { return v.Raw }).(pulumi.BoolPtrOutput)
}

// The value of the variable.
func (o InstanceVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceVariable) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
func (o InstanceVariableOutput) VariableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceVariable) pulumi.StringPtrOutput { return v.VariableType }).(pulumi.StringPtrOutput)
}

type InstanceVariableArrayOutput struct{ *pulumi.OutputState }

func (InstanceVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceVariable)(nil)).Elem()
}

func (o InstanceVariableArrayOutput) ToInstanceVariableArrayOutput() InstanceVariableArrayOutput {
	return o
}

func (o InstanceVariableArrayOutput) ToInstanceVariableArrayOutputWithContext(ctx context.Context) InstanceVariableArrayOutput {
	return o
}

func (o InstanceVariableArrayOutput) Index(i pulumi.IntInput) InstanceVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceVariable {
		return vs[0].([]*InstanceVariable)[vs[1].(int)]
	}).(InstanceVariableOutput)
}

type InstanceVariableMapOutput struct{ *pulumi.OutputState }

func (InstanceVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceVariable)(nil)).Elem()
}

func (o InstanceVariableMapOutput) ToInstanceVariableMapOutput() InstanceVariableMapOutput {
	return o
}

func (o InstanceVariableMapOutput) ToInstanceVariableMapOutputWithContext(ctx context.Context) InstanceVariableMapOutput {
	return o
}

func (o InstanceVariableMapOutput) MapIndex(k pulumi.StringInput) InstanceVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceVariable {
		return vs[0].(map[string]*InstanceVariable)[vs[1].(string)]
	}).(InstanceVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceVariableInput)(nil)).Elem(), &InstanceVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceVariableArrayInput)(nil)).Elem(), InstanceVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceVariableMapInput)(nil)).Elem(), InstanceVariableMap{})
	pulumi.RegisterOutputType(InstanceVariableOutput{})
	pulumi.RegisterOutputType(InstanceVariableArrayOutput{})
	pulumi.RegisterOutputType(InstanceVariableMapOutput{})
}
