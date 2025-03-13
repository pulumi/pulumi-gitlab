// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupCustomAttribute` resource allows to manage custom attributes for a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/custom_attributes/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewGroupCustomAttribute(ctx, "attr", &gitlab.GroupCustomAttributeArgs{
//				Group: pulumi.Int(42),
//				Key:   pulumi.String("location"),
//				Value: pulumi.String("Greenland"),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_custom_attribute`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_group_custom_attribute.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// You can import a group custom attribute using the an id made up of `{group-id}:{key}`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/groupCustomAttribute:GroupCustomAttribute attr 42:location
// ```
type GroupCustomAttribute struct {
	pulumi.CustomResourceState

	// The id of the group.
	Group pulumi.IntOutput `pulumi:"group"`
	// Key for the Custom Attribute.
	Key pulumi.StringOutput `pulumi:"key"`
	// Value for the Custom Attribute.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewGroupCustomAttribute registers a new resource with the given unique name, arguments, and options.
func NewGroupCustomAttribute(ctx *pulumi.Context,
	name string, args *GroupCustomAttributeArgs, opts ...pulumi.ResourceOption) (*GroupCustomAttribute, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupCustomAttribute
	err := ctx.RegisterResource("gitlab:index/groupCustomAttribute:GroupCustomAttribute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupCustomAttribute gets an existing GroupCustomAttribute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupCustomAttribute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupCustomAttributeState, opts ...pulumi.ResourceOption) (*GroupCustomAttribute, error) {
	var resource GroupCustomAttribute
	err := ctx.ReadResource("gitlab:index/groupCustomAttribute:GroupCustomAttribute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupCustomAttribute resources.
type groupCustomAttributeState struct {
	// The id of the group.
	Group *int `pulumi:"group"`
	// Key for the Custom Attribute.
	Key *string `pulumi:"key"`
	// Value for the Custom Attribute.
	Value *string `pulumi:"value"`
}

type GroupCustomAttributeState struct {
	// The id of the group.
	Group pulumi.IntPtrInput
	// Key for the Custom Attribute.
	Key pulumi.StringPtrInput
	// Value for the Custom Attribute.
	Value pulumi.StringPtrInput
}

func (GroupCustomAttributeState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupCustomAttributeState)(nil)).Elem()
}

type groupCustomAttributeArgs struct {
	// The id of the group.
	Group int `pulumi:"group"`
	// Key for the Custom Attribute.
	Key string `pulumi:"key"`
	// Value for the Custom Attribute.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a GroupCustomAttribute resource.
type GroupCustomAttributeArgs struct {
	// The id of the group.
	Group pulumi.IntInput
	// Key for the Custom Attribute.
	Key pulumi.StringInput
	// Value for the Custom Attribute.
	Value pulumi.StringInput
}

func (GroupCustomAttributeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupCustomAttributeArgs)(nil)).Elem()
}

type GroupCustomAttributeInput interface {
	pulumi.Input

	ToGroupCustomAttributeOutput() GroupCustomAttributeOutput
	ToGroupCustomAttributeOutputWithContext(ctx context.Context) GroupCustomAttributeOutput
}

func (*GroupCustomAttribute) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupCustomAttribute)(nil)).Elem()
}

func (i *GroupCustomAttribute) ToGroupCustomAttributeOutput() GroupCustomAttributeOutput {
	return i.ToGroupCustomAttributeOutputWithContext(context.Background())
}

func (i *GroupCustomAttribute) ToGroupCustomAttributeOutputWithContext(ctx context.Context) GroupCustomAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupCustomAttributeOutput)
}

// GroupCustomAttributeArrayInput is an input type that accepts GroupCustomAttributeArray and GroupCustomAttributeArrayOutput values.
// You can construct a concrete instance of `GroupCustomAttributeArrayInput` via:
//
//	GroupCustomAttributeArray{ GroupCustomAttributeArgs{...} }
type GroupCustomAttributeArrayInput interface {
	pulumi.Input

	ToGroupCustomAttributeArrayOutput() GroupCustomAttributeArrayOutput
	ToGroupCustomAttributeArrayOutputWithContext(context.Context) GroupCustomAttributeArrayOutput
}

type GroupCustomAttributeArray []GroupCustomAttributeInput

func (GroupCustomAttributeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupCustomAttribute)(nil)).Elem()
}

func (i GroupCustomAttributeArray) ToGroupCustomAttributeArrayOutput() GroupCustomAttributeArrayOutput {
	return i.ToGroupCustomAttributeArrayOutputWithContext(context.Background())
}

func (i GroupCustomAttributeArray) ToGroupCustomAttributeArrayOutputWithContext(ctx context.Context) GroupCustomAttributeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupCustomAttributeArrayOutput)
}

// GroupCustomAttributeMapInput is an input type that accepts GroupCustomAttributeMap and GroupCustomAttributeMapOutput values.
// You can construct a concrete instance of `GroupCustomAttributeMapInput` via:
//
//	GroupCustomAttributeMap{ "key": GroupCustomAttributeArgs{...} }
type GroupCustomAttributeMapInput interface {
	pulumi.Input

	ToGroupCustomAttributeMapOutput() GroupCustomAttributeMapOutput
	ToGroupCustomAttributeMapOutputWithContext(context.Context) GroupCustomAttributeMapOutput
}

type GroupCustomAttributeMap map[string]GroupCustomAttributeInput

func (GroupCustomAttributeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupCustomAttribute)(nil)).Elem()
}

func (i GroupCustomAttributeMap) ToGroupCustomAttributeMapOutput() GroupCustomAttributeMapOutput {
	return i.ToGroupCustomAttributeMapOutputWithContext(context.Background())
}

func (i GroupCustomAttributeMap) ToGroupCustomAttributeMapOutputWithContext(ctx context.Context) GroupCustomAttributeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupCustomAttributeMapOutput)
}

type GroupCustomAttributeOutput struct{ *pulumi.OutputState }

func (GroupCustomAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupCustomAttribute)(nil)).Elem()
}

func (o GroupCustomAttributeOutput) ToGroupCustomAttributeOutput() GroupCustomAttributeOutput {
	return o
}

func (o GroupCustomAttributeOutput) ToGroupCustomAttributeOutputWithContext(ctx context.Context) GroupCustomAttributeOutput {
	return o
}

// The id of the group.
func (o GroupCustomAttributeOutput) Group() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupCustomAttribute) pulumi.IntOutput { return v.Group }).(pulumi.IntOutput)
}

// Key for the Custom Attribute.
func (o GroupCustomAttributeOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupCustomAttribute) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Value for the Custom Attribute.
func (o GroupCustomAttributeOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupCustomAttribute) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type GroupCustomAttributeArrayOutput struct{ *pulumi.OutputState }

func (GroupCustomAttributeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupCustomAttribute)(nil)).Elem()
}

func (o GroupCustomAttributeArrayOutput) ToGroupCustomAttributeArrayOutput() GroupCustomAttributeArrayOutput {
	return o
}

func (o GroupCustomAttributeArrayOutput) ToGroupCustomAttributeArrayOutputWithContext(ctx context.Context) GroupCustomAttributeArrayOutput {
	return o
}

func (o GroupCustomAttributeArrayOutput) Index(i pulumi.IntInput) GroupCustomAttributeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupCustomAttribute {
		return vs[0].([]*GroupCustomAttribute)[vs[1].(int)]
	}).(GroupCustomAttributeOutput)
}

type GroupCustomAttributeMapOutput struct{ *pulumi.OutputState }

func (GroupCustomAttributeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupCustomAttribute)(nil)).Elem()
}

func (o GroupCustomAttributeMapOutput) ToGroupCustomAttributeMapOutput() GroupCustomAttributeMapOutput {
	return o
}

func (o GroupCustomAttributeMapOutput) ToGroupCustomAttributeMapOutputWithContext(ctx context.Context) GroupCustomAttributeMapOutput {
	return o
}

func (o GroupCustomAttributeMapOutput) MapIndex(k pulumi.StringInput) GroupCustomAttributeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupCustomAttribute {
		return vs[0].(map[string]*GroupCustomAttribute)[vs[1].(string)]
	}).(GroupCustomAttributeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupCustomAttributeInput)(nil)).Elem(), &GroupCustomAttribute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupCustomAttributeArrayInput)(nil)).Elem(), GroupCustomAttributeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupCustomAttributeMapInput)(nil)).Elem(), GroupCustomAttributeMap{})
	pulumi.RegisterOutputType(GroupCustomAttributeOutput{})
	pulumi.RegisterOutputType(GroupCustomAttributeArrayOutput{})
	pulumi.RegisterOutputType(GroupCustomAttributeMapOutput{})
}
