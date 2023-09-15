// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `GroupMembership` resource allows to manage the lifecycle of a users group membership.
//
// > If a group should grant membership to another group use the `GroupShareGroup` resource instead.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html)
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
//			_, err := gitlab.NewGroupMembership(ctx, "test", &gitlab.GroupMembershipArgs{
//				AccessLevel: pulumi.String("guest"),
//				ExpiresAt:   pulumi.String("2020-12-31"),
//				GroupId:     pulumi.String("12345"),
//				UserId:      pulumi.Int(1337),
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
// GitLab group membership can be imported using an id made up of `group_id:user_id`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/groupMembership:GroupMembership test "12345:1337"
//
// ```
type GroupMembership struct {
	pulumi.CustomResourceState

	// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The id of the group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
	SkipSubresourcesOnDestroy pulumi.BoolPtrOutput `pulumi:"skipSubresourcesOnDestroy"`
	// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
	UnassignIssuablesOnDestroy pulumi.BoolPtrOutput `pulumi:"unassignIssuablesOnDestroy"`
	// The id of the user.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewGroupMembership(ctx *pulumi.Context,
	name string, args *GroupMembershipArgs, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevel'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMembership
	err := ctx.RegisterResource("gitlab:index/groupMembership:GroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembership gets an existing GroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipState, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	var resource GroupMembership
	err := ctx.ReadResource("gitlab:index/groupMembership:GroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembership resources.
type groupMembershipState struct {
	// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
	AccessLevel *string `pulumi:"accessLevel"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// The id of the group.
	GroupId *string `pulumi:"groupId"`
	// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
	SkipSubresourcesOnDestroy *bool `pulumi:"skipSubresourcesOnDestroy"`
	// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
	UnassignIssuablesOnDestroy *bool `pulumi:"unassignIssuablesOnDestroy"`
	// The id of the user.
	UserId *int `pulumi:"userId"`
}

type GroupMembershipState struct {
	// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
	AccessLevel pulumi.StringPtrInput
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// The id of the group.
	GroupId pulumi.StringPtrInput
	// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
	SkipSubresourcesOnDestroy pulumi.BoolPtrInput
	// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
	UnassignIssuablesOnDestroy pulumi.BoolPtrInput
	// The id of the user.
	UserId pulumi.IntPtrInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
	AccessLevel string `pulumi:"accessLevel"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// The id of the group.
	GroupId string `pulumi:"groupId"`
	// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
	SkipSubresourcesOnDestroy *bool `pulumi:"skipSubresourcesOnDestroy"`
	// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
	UnassignIssuablesOnDestroy *bool `pulumi:"unassignIssuablesOnDestroy"`
	// The id of the user.
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
	AccessLevel pulumi.StringInput
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// The id of the group.
	GroupId pulumi.StringInput
	// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
	SkipSubresourcesOnDestroy pulumi.BoolPtrInput
	// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
	UnassignIssuablesOnDestroy pulumi.BoolPtrInput
	// The id of the user.
	UserId pulumi.IntInput
}

func (GroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipArgs)(nil)).Elem()
}

type GroupMembershipInput interface {
	pulumi.Input

	ToGroupMembershipOutput() GroupMembershipOutput
	ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput
}

func (*GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil)).Elem()
}

func (i *GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i *GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

func (i *GroupMembership) ToOutput(ctx context.Context) pulumix.Output[*GroupMembership] {
	return pulumix.Output[*GroupMembership]{
		OutputState: i.ToGroupMembershipOutputWithContext(ctx).OutputState,
	}
}

// GroupMembershipArrayInput is an input type that accepts GroupMembershipArray and GroupMembershipArrayOutput values.
// You can construct a concrete instance of `GroupMembershipArrayInput` via:
//
//	GroupMembershipArray{ GroupMembershipArgs{...} }
type GroupMembershipArrayInput interface {
	pulumi.Input

	ToGroupMembershipArrayOutput() GroupMembershipArrayOutput
	ToGroupMembershipArrayOutputWithContext(context.Context) GroupMembershipArrayOutput
}

type GroupMembershipArray []GroupMembershipInput

func (GroupMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembership)(nil)).Elem()
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return i.ToGroupMembershipArrayOutputWithContext(context.Background())
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipArrayOutput)
}

func (i GroupMembershipArray) ToOutput(ctx context.Context) pulumix.Output[[]*GroupMembership] {
	return pulumix.Output[[]*GroupMembership]{
		OutputState: i.ToGroupMembershipArrayOutputWithContext(ctx).OutputState,
	}
}

// GroupMembershipMapInput is an input type that accepts GroupMembershipMap and GroupMembershipMapOutput values.
// You can construct a concrete instance of `GroupMembershipMapInput` via:
//
//	GroupMembershipMap{ "key": GroupMembershipArgs{...} }
type GroupMembershipMapInput interface {
	pulumi.Input

	ToGroupMembershipMapOutput() GroupMembershipMapOutput
	ToGroupMembershipMapOutputWithContext(context.Context) GroupMembershipMapOutput
}

type GroupMembershipMap map[string]GroupMembershipInput

func (GroupMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembership)(nil)).Elem()
}

func (i GroupMembershipMap) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return i.ToGroupMembershipMapOutputWithContext(context.Background())
}

func (i GroupMembershipMap) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipMapOutput)
}

func (i GroupMembershipMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupMembership] {
	return pulumix.Output[map[string]*GroupMembership]{
		OutputState: i.ToGroupMembershipMapOutputWithContext(ctx).OutputState,
	}
}

type GroupMembershipOutput struct{ *pulumi.OutputState }

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil)).Elem()
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupMembership] {
	return pulumix.Output[*GroupMembership]{
		OutputState: o.OutputState,
	}
}

// Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`.
func (o GroupMembershipOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// Expiration date for the group membership. Format: `YYYY-MM-DD`
func (o GroupMembershipOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The id of the group.
func (o GroupMembershipOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
func (o GroupMembershipOutput) SkipSubresourcesOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.BoolPtrOutput { return v.SkipSubresourcesOnDestroy }).(pulumi.BoolPtrOutput)
}

// Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
func (o GroupMembershipOutput) UnassignIssuablesOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.BoolPtrOutput { return v.UnassignIssuablesOnDestroy }).(pulumi.BoolPtrOutput)
}

// The id of the user.
func (o GroupMembershipOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type GroupMembershipArrayOutput struct{ *pulumi.OutputState }

func (GroupMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembership)(nil)).Elem()
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GroupMembership] {
	return pulumix.Output[[]*GroupMembership]{
		OutputState: o.OutputState,
	}
}

func (o GroupMembershipArrayOutput) Index(i pulumi.IntInput) GroupMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMembership {
		return vs[0].([]*GroupMembership)[vs[1].(int)]
	}).(GroupMembershipOutput)
}

type GroupMembershipMapOutput struct{ *pulumi.OutputState }

func (GroupMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembership)(nil)).Elem()
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupMembership] {
	return pulumix.Output[map[string]*GroupMembership]{
		OutputState: o.OutputState,
	}
}

func (o GroupMembershipMapOutput) MapIndex(k pulumi.StringInput) GroupMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMembership {
		return vs[0].(map[string]*GroupMembership)[vs[1].(string)]
	}).(GroupMembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipInput)(nil)).Elem(), &GroupMembership{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipArrayInput)(nil)).Elem(), GroupMembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipMapInput)(nil)).Elem(), GroupMembershipMap{})
	pulumi.RegisterOutputType(GroupMembershipOutput{})
	pulumi.RegisterOutputType(GroupMembershipArrayOutput{})
	pulumi.RegisterOutputType(GroupMembershipMapOutput{})
}
