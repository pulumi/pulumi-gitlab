// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_group\_share\_group
//
// This resource allows you to share a group with another group
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
// 		_, err := gitlab.NewGroupShareGroup(ctx, "test", &gitlab.GroupShareGroupArgs{
// 			GroupId:      pulumi.Any(gitlab_group.Foo.Id),
// 			ShareGroupId: pulumi.Any(gitlab_group.Bar.Id),
// 			GroupAccess:  pulumi.String("guest"),
// 			ExpiresAt:    pulumi.String("2099-01-01"),
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
// GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
// ```
type GroupShareGroup struct {
	pulumi.CustomResourceState

	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// One of five levels of access to the group.
	GroupAccess pulumi.StringOutput `pulumi:"groupAccess"`
	// The id of the main group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The id of an additional group which will be shared with the main group.
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
	// One of five levels of access to the group.
	GroupAccess *string `pulumi:"groupAccess"`
	// The id of the main group.
	GroupId *string `pulumi:"groupId"`
	// The id of an additional group which will be shared with the main group.
	ShareGroupId *int `pulumi:"shareGroupId"`
}

type GroupShareGroupState struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// One of five levels of access to the group.
	GroupAccess pulumi.StringPtrInput
	// The id of the main group.
	GroupId pulumi.StringPtrInput
	// The id of an additional group which will be shared with the main group.
	ShareGroupId pulumi.IntPtrInput
}

func (GroupShareGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupShareGroupState)(nil)).Elem()
}

type groupShareGroupArgs struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt *string `pulumi:"expiresAt"`
	// One of five levels of access to the group.
	GroupAccess string `pulumi:"groupAccess"`
	// The id of the main group.
	GroupId string `pulumi:"groupId"`
	// The id of an additional group which will be shared with the main group.
	ShareGroupId int `pulumi:"shareGroupId"`
}

// The set of arguments for constructing a GroupShareGroup resource.
type GroupShareGroupArgs struct {
	// Share expiration date. Format: `YYYY-MM-DD`
	ExpiresAt pulumi.StringPtrInput
	// One of five levels of access to the group.
	GroupAccess pulumi.StringInput
	// The id of the main group.
	GroupId pulumi.StringInput
	// The id of an additional group which will be shared with the main group.
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

func (GroupShareGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupShareGroup)(nil)).Elem()
}

func (i GroupShareGroup) ToGroupShareGroupOutput() GroupShareGroupOutput {
	return i.ToGroupShareGroupOutputWithContext(context.Background())
}

func (i GroupShareGroup) ToGroupShareGroupOutputWithContext(ctx context.Context) GroupShareGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupShareGroupOutput)
}

type GroupShareGroupOutput struct {
	*pulumi.OutputState
}

func (GroupShareGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupShareGroupOutput)(nil)).Elem()
}

func (o GroupShareGroupOutput) ToGroupShareGroupOutput() GroupShareGroupOutput {
	return o
}

func (o GroupShareGroupOutput) ToGroupShareGroupOutputWithContext(ctx context.Context) GroupShareGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupShareGroupOutput{})
}