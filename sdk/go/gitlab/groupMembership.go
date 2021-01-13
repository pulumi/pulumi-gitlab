// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_group_membership
//
// This resource allows you to add a user to an existing group.
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
// 		_, err := gitlab.NewGroupMembership(ctx, "test", &gitlab.GroupMembershipArgs{
// 			AccessLevel: pulumi.String("guest"),
// 			ExpiresAt:   pulumi.String("2020-12-31"),
// 			GroupId:     pulumi.String("12345"),
// 			UserId:      pulumi.Int(1337),
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
// GitLab group membership can be imported using an id made up of `group_id:user_id`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/groupMembership:GroupMembership test "12345:1337"
// ```
type GroupMembership struct {
	pulumi.CustomResourceState

	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The id of the group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
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
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel *string `pulumi:"accessLevel"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// The id of the group.
	GroupId *string `pulumi:"groupId"`
	// The id of the user.
	UserId *int `pulumi:"userId"`
}

type GroupMembershipState struct {
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel pulumi.StringPtrInput
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// The id of the group.
	GroupId pulumi.StringPtrInput
	// The id of the user.
	UserId pulumi.IntPtrInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel string `pulumi:"accessLevel"`
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// The id of the group.
	GroupId string `pulumi:"groupId"`
	// The id of the user.
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// Acceptable values are: guest, reporter, developer, maintainer, owner.
	AccessLevel pulumi.StringInput
	// Expiration date for the group membership. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// The id of the group.
	GroupId pulumi.StringInput
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

func (GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembership)(nil)).Elem()
}

func (i GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

type GroupMembershipOutput struct {
	*pulumi.OutputState
}

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembershipOutput)(nil)).Elem()
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupMembershipOutput{})
}
