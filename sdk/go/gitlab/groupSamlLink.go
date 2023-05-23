// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupSamlLink` resource allows to manage the lifecycle of an SAML integration with a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#saml-group-links)
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
//			_, err := gitlab.NewGroupSamlLink(ctx, "test", &gitlab.GroupSamlLinkArgs{
//				AccessLevel:   pulumi.String("developer"),
//				Group:         pulumi.String("12345"),
//				SamlGroupName: pulumi.String("samlgroupname1"),
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
// GitLab group saml links can be imported using an id made up of `group_id:saml_group_name`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/groupSamlLink:GroupSamlLink test "12345:samlgroupname1"
//
// ```
type GroupSamlLink struct {
	pulumi.CustomResourceState

	// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// The ID or path of the group to add the SAML Group Link to.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the SAML group.
	SamlGroupName pulumi.StringOutput `pulumi:"samlGroupName"`
}

// NewGroupSamlLink registers a new resource with the given unique name, arguments, and options.
func NewGroupSamlLink(ctx *pulumi.Context,
	name string, args *GroupSamlLinkArgs, opts ...pulumi.ResourceOption) (*GroupSamlLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevel'")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.SamlGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SamlGroupName'")
	}
	var resource GroupSamlLink
	err := ctx.RegisterResource("gitlab:index/groupSamlLink:GroupSamlLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupSamlLink gets an existing GroupSamlLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupSamlLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupSamlLinkState, opts ...pulumi.ResourceOption) (*GroupSamlLink, error) {
	var resource GroupSamlLink
	err := ctx.ReadResource("gitlab:index/groupSamlLink:GroupSamlLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupSamlLink resources.
type groupSamlLinkState struct {
	// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The ID or path of the group to add the SAML Group Link to.
	Group *string `pulumi:"group"`
	// The name of the SAML group.
	SamlGroupName *string `pulumi:"samlGroupName"`
}

type GroupSamlLinkState struct {
	// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringPtrInput
	// The ID or path of the group to add the SAML Group Link to.
	Group pulumi.StringPtrInput
	// The name of the SAML group.
	SamlGroupName pulumi.StringPtrInput
}

func (GroupSamlLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupSamlLinkState)(nil)).Elem()
}

type groupSamlLinkArgs struct {
	// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel string `pulumi:"accessLevel"`
	// The ID or path of the group to add the SAML Group Link to.
	Group string `pulumi:"group"`
	// The name of the SAML group.
	SamlGroupName string `pulumi:"samlGroupName"`
}

// The set of arguments for constructing a GroupSamlLink resource.
type GroupSamlLinkArgs struct {
	// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringInput
	// The ID or path of the group to add the SAML Group Link to.
	Group pulumi.StringInput
	// The name of the SAML group.
	SamlGroupName pulumi.StringInput
}

func (GroupSamlLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupSamlLinkArgs)(nil)).Elem()
}

type GroupSamlLinkInput interface {
	pulumi.Input

	ToGroupSamlLinkOutput() GroupSamlLinkOutput
	ToGroupSamlLinkOutputWithContext(ctx context.Context) GroupSamlLinkOutput
}

func (*GroupSamlLink) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupSamlLink)(nil)).Elem()
}

func (i *GroupSamlLink) ToGroupSamlLinkOutput() GroupSamlLinkOutput {
	return i.ToGroupSamlLinkOutputWithContext(context.Background())
}

func (i *GroupSamlLink) ToGroupSamlLinkOutputWithContext(ctx context.Context) GroupSamlLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupSamlLinkOutput)
}

// GroupSamlLinkArrayInput is an input type that accepts GroupSamlLinkArray and GroupSamlLinkArrayOutput values.
// You can construct a concrete instance of `GroupSamlLinkArrayInput` via:
//
//	GroupSamlLinkArray{ GroupSamlLinkArgs{...} }
type GroupSamlLinkArrayInput interface {
	pulumi.Input

	ToGroupSamlLinkArrayOutput() GroupSamlLinkArrayOutput
	ToGroupSamlLinkArrayOutputWithContext(context.Context) GroupSamlLinkArrayOutput
}

type GroupSamlLinkArray []GroupSamlLinkInput

func (GroupSamlLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupSamlLink)(nil)).Elem()
}

func (i GroupSamlLinkArray) ToGroupSamlLinkArrayOutput() GroupSamlLinkArrayOutput {
	return i.ToGroupSamlLinkArrayOutputWithContext(context.Background())
}

func (i GroupSamlLinkArray) ToGroupSamlLinkArrayOutputWithContext(ctx context.Context) GroupSamlLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupSamlLinkArrayOutput)
}

// GroupSamlLinkMapInput is an input type that accepts GroupSamlLinkMap and GroupSamlLinkMapOutput values.
// You can construct a concrete instance of `GroupSamlLinkMapInput` via:
//
//	GroupSamlLinkMap{ "key": GroupSamlLinkArgs{...} }
type GroupSamlLinkMapInput interface {
	pulumi.Input

	ToGroupSamlLinkMapOutput() GroupSamlLinkMapOutput
	ToGroupSamlLinkMapOutputWithContext(context.Context) GroupSamlLinkMapOutput
}

type GroupSamlLinkMap map[string]GroupSamlLinkInput

func (GroupSamlLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupSamlLink)(nil)).Elem()
}

func (i GroupSamlLinkMap) ToGroupSamlLinkMapOutput() GroupSamlLinkMapOutput {
	return i.ToGroupSamlLinkMapOutputWithContext(context.Background())
}

func (i GroupSamlLinkMap) ToGroupSamlLinkMapOutputWithContext(ctx context.Context) GroupSamlLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupSamlLinkMapOutput)
}

type GroupSamlLinkOutput struct{ *pulumi.OutputState }

func (GroupSamlLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupSamlLink)(nil)).Elem()
}

func (o GroupSamlLinkOutput) ToGroupSamlLinkOutput() GroupSamlLinkOutput {
	return o
}

func (o GroupSamlLinkOutput) ToGroupSamlLinkOutputWithContext(ctx context.Context) GroupSamlLinkOutput {
	return o
}

// Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
func (o GroupSamlLinkOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupSamlLink) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// The ID or path of the group to add the SAML Group Link to.
func (o GroupSamlLinkOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupSamlLink) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The name of the SAML group.
func (o GroupSamlLinkOutput) SamlGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupSamlLink) pulumi.StringOutput { return v.SamlGroupName }).(pulumi.StringOutput)
}

type GroupSamlLinkArrayOutput struct{ *pulumi.OutputState }

func (GroupSamlLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupSamlLink)(nil)).Elem()
}

func (o GroupSamlLinkArrayOutput) ToGroupSamlLinkArrayOutput() GroupSamlLinkArrayOutput {
	return o
}

func (o GroupSamlLinkArrayOutput) ToGroupSamlLinkArrayOutputWithContext(ctx context.Context) GroupSamlLinkArrayOutput {
	return o
}

func (o GroupSamlLinkArrayOutput) Index(i pulumi.IntInput) GroupSamlLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupSamlLink {
		return vs[0].([]*GroupSamlLink)[vs[1].(int)]
	}).(GroupSamlLinkOutput)
}

type GroupSamlLinkMapOutput struct{ *pulumi.OutputState }

func (GroupSamlLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupSamlLink)(nil)).Elem()
}

func (o GroupSamlLinkMapOutput) ToGroupSamlLinkMapOutput() GroupSamlLinkMapOutput {
	return o
}

func (o GroupSamlLinkMapOutput) ToGroupSamlLinkMapOutputWithContext(ctx context.Context) GroupSamlLinkMapOutput {
	return o
}

func (o GroupSamlLinkMapOutput) MapIndex(k pulumi.StringInput) GroupSamlLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupSamlLink {
		return vs[0].(map[string]*GroupSamlLink)[vs[1].(string)]
	}).(GroupSamlLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupSamlLinkInput)(nil)).Elem(), &GroupSamlLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupSamlLinkArrayInput)(nil)).Elem(), GroupSamlLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupSamlLinkMapInput)(nil)).Elem(), GroupSamlLinkMap{})
	pulumi.RegisterOutputType(GroupSamlLinkOutput{})
	pulumi.RegisterOutputType(GroupSamlLinkArrayOutput{})
	pulumi.RegisterOutputType(GroupSamlLinkMapOutput{})
}
