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

// The `GroupShareGroup` resource allows managing the lifecycle of a group shared with another group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#share-groups-with-groups)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewGroupShareGroup(ctx, "test", &gitlab.GroupShareGroupArgs{
//				GroupId:      pulumi.Any(foo.Id),
//				ShareGroupId: pulumi.Any(bar.Id),
//				GroupAccess:  pulumi.String("guest"),
//				ExpiresAt:    pulumi.String("2099-01-01"),
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
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_share_group`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_group_share_group.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
// ```
type GroupShareGroup struct {
	pulumi.CustomResourceState

	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess pulumi.StringOutput `pulumi:"groupAccess"`
	// The id of the main group to be shared.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The ID of a custom member role. Only available for Ultimate instances. If `memberRoleId` is removed from the config, the group share will revert to a base role.
	MemberRoleId pulumi.IntPtrOutput `pulumi:"memberRoleId"`
	// The id of the additional group with which the main group will be shared.
	ShareGroupId pulumi.IntOutput `pulumi:"shareGroupId"`
}

// NewGroupShareGroup registers a new resource with the given unique name, arguments, and options.
func NewGroupShareGroup(ctx *pulumi.Context,
	name string, args *GroupShareGroupArgs, opts ...pulumi.ResourceOption) (*GroupShareGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupAccess == nil {
		return nil, errors.New("invalid value for required argument 'GroupAccess'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.ShareGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ShareGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupShareGroup
	err := ctx.RegisterResource("gitlab:index/groupShareGroup:GroupShareGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupShareGroup gets an existing GroupShareGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupShareGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupShareGroupState, opts ...pulumi.ResourceOption) (*GroupShareGroup, error) {
	var resource GroupShareGroup
	err := ctx.ReadResource("gitlab:index/groupShareGroup:GroupShareGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupShareGroup resources.
type groupShareGroupState struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess *string `pulumi:"groupAccess"`
	// The id of the main group to be shared.
	GroupId *string `pulumi:"groupId"`
	// The ID of a custom member role. Only available for Ultimate instances. If `memberRoleId` is removed from the config, the group share will revert to a base role.
	MemberRoleId *int `pulumi:"memberRoleId"`
	// The id of the additional group with which the main group will be shared.
	ShareGroupId *int `pulumi:"shareGroupId"`
}

type GroupShareGroupState struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess pulumi.StringPtrInput
	// The id of the main group to be shared.
	GroupId pulumi.StringPtrInput
	// The ID of a custom member role. Only available for Ultimate instances. If `memberRoleId` is removed from the config, the group share will revert to a base role.
	MemberRoleId pulumi.IntPtrInput
	// The id of the additional group with which the main group will be shared.
	ShareGroupId pulumi.IntPtrInput
}

func (GroupShareGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupShareGroupState)(nil)).Elem()
}

type groupShareGroupArgs struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess string `pulumi:"groupAccess"`
	// The id of the main group to be shared.
	GroupId string `pulumi:"groupId"`
	// The ID of a custom member role. Only available for Ultimate instances. If `memberRoleId` is removed from the config, the group share will revert to a base role.
	MemberRoleId *int `pulumi:"memberRoleId"`
	// The id of the additional group with which the main group will be shared.
	ShareGroupId int `pulumi:"shareGroupId"`
}

// The set of arguments for constructing a GroupShareGroup resource.
type GroupShareGroupArgs struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
	GroupAccess pulumi.StringInput
	// The id of the main group to be shared.
	GroupId pulumi.StringInput
	// The ID of a custom member role. Only available for Ultimate instances. If `memberRoleId` is removed from the config, the group share will revert to a base role.
	MemberRoleId pulumi.IntPtrInput
	// The id of the additional group with which the main group will be shared.
	ShareGroupId pulumi.IntInput
}

func (GroupShareGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupShareGroupArgs)(nil)).Elem()
}

type GroupShareGroupInput interface {
	pulumi.Input

	ToGroupShareGroupOutput() GroupShareGroupOutput
	ToGroupShareGroupOutputWithContext(ctx context.Context) GroupShareGroupOutput
}

func (*GroupShareGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupShareGroup)(nil)).Elem()
}

func (i *GroupShareGroup) ToGroupShareGroupOutput() GroupShareGroupOutput {
	return i.ToGroupShareGroupOutputWithContext(context.Background())
}

func (i *GroupShareGroup) ToGroupShareGroupOutputWithContext(ctx context.Context) GroupShareGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupShareGroupOutput)
}

// GroupShareGroupArrayInput is an input type that accepts GroupShareGroupArray and GroupShareGroupArrayOutput values.
// You can construct a concrete instance of `GroupShareGroupArrayInput` via:
//
//	GroupShareGroupArray{ GroupShareGroupArgs{...} }
type GroupShareGroupArrayInput interface {
	pulumi.Input

	ToGroupShareGroupArrayOutput() GroupShareGroupArrayOutput
	ToGroupShareGroupArrayOutputWithContext(context.Context) GroupShareGroupArrayOutput
}

type GroupShareGroupArray []GroupShareGroupInput

func (GroupShareGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupShareGroup)(nil)).Elem()
}

func (i GroupShareGroupArray) ToGroupShareGroupArrayOutput() GroupShareGroupArrayOutput {
	return i.ToGroupShareGroupArrayOutputWithContext(context.Background())
}

func (i GroupShareGroupArray) ToGroupShareGroupArrayOutputWithContext(ctx context.Context) GroupShareGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupShareGroupArrayOutput)
}

// GroupShareGroupMapInput is an input type that accepts GroupShareGroupMap and GroupShareGroupMapOutput values.
// You can construct a concrete instance of `GroupShareGroupMapInput` via:
//
//	GroupShareGroupMap{ "key": GroupShareGroupArgs{...} }
type GroupShareGroupMapInput interface {
	pulumi.Input

	ToGroupShareGroupMapOutput() GroupShareGroupMapOutput
	ToGroupShareGroupMapOutputWithContext(context.Context) GroupShareGroupMapOutput
}

type GroupShareGroupMap map[string]GroupShareGroupInput

func (GroupShareGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupShareGroup)(nil)).Elem()
}

func (i GroupShareGroupMap) ToGroupShareGroupMapOutput() GroupShareGroupMapOutput {
	return i.ToGroupShareGroupMapOutputWithContext(context.Background())
}

func (i GroupShareGroupMap) ToGroupShareGroupMapOutputWithContext(ctx context.Context) GroupShareGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupShareGroupMapOutput)
}

type GroupShareGroupOutput struct{ *pulumi.OutputState }

func (GroupShareGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupShareGroup)(nil)).Elem()
}

func (o GroupShareGroupOutput) ToGroupShareGroupOutput() GroupShareGroupOutput {
	return o
}

func (o GroupShareGroupOutput) ToGroupShareGroupOutputWithContext(ctx context.Context) GroupShareGroupOutput {
	return o
}

// Share expiration date. Format: `YYYY-MM-DD`
func (o GroupShareGroupOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupShareGroup) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
func (o GroupShareGroupOutput) GroupAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupShareGroup) pulumi.StringOutput { return v.GroupAccess }).(pulumi.StringOutput)
}

// The id of the main group to be shared.
func (o GroupShareGroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupShareGroup) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The ID of a custom member role. Only available for Ultimate instances. If `memberRoleId` is removed from the config, the group share will revert to a base role.
func (o GroupShareGroupOutput) MemberRoleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GroupShareGroup) pulumi.IntPtrOutput { return v.MemberRoleId }).(pulumi.IntPtrOutput)
}

// The id of the additional group with which the main group will be shared.
func (o GroupShareGroupOutput) ShareGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupShareGroup) pulumi.IntOutput { return v.ShareGroupId }).(pulumi.IntOutput)
}

type GroupShareGroupArrayOutput struct{ *pulumi.OutputState }

func (GroupShareGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupShareGroup)(nil)).Elem()
}

func (o GroupShareGroupArrayOutput) ToGroupShareGroupArrayOutput() GroupShareGroupArrayOutput {
	return o
}

func (o GroupShareGroupArrayOutput) ToGroupShareGroupArrayOutputWithContext(ctx context.Context) GroupShareGroupArrayOutput {
	return o
}

func (o GroupShareGroupArrayOutput) Index(i pulumi.IntInput) GroupShareGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupShareGroup {
		return vs[0].([]*GroupShareGroup)[vs[1].(int)]
	}).(GroupShareGroupOutput)
}

type GroupShareGroupMapOutput struct{ *pulumi.OutputState }

func (GroupShareGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupShareGroup)(nil)).Elem()
}

func (o GroupShareGroupMapOutput) ToGroupShareGroupMapOutput() GroupShareGroupMapOutput {
	return o
}

func (o GroupShareGroupMapOutput) ToGroupShareGroupMapOutputWithContext(ctx context.Context) GroupShareGroupMapOutput {
	return o
}

func (o GroupShareGroupMapOutput) MapIndex(k pulumi.StringInput) GroupShareGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupShareGroup {
		return vs[0].(map[string]*GroupShareGroup)[vs[1].(string)]
	}).(GroupShareGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupShareGroupInput)(nil)).Elem(), &GroupShareGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupShareGroupArrayInput)(nil)).Elem(), GroupShareGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupShareGroupMapInput)(nil)).Elem(), GroupShareGroupMap{})
	pulumi.RegisterOutputType(GroupShareGroupOutput{})
	pulumi.RegisterOutputType(GroupShareGroupArrayOutput{})
	pulumi.RegisterOutputType(GroupShareGroupMapOutput{})
}
