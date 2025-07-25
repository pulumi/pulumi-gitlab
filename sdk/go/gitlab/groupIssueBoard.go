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

// The `GroupIssueBoard` resource manages the lifecycle of an issue board in a group.
//
// > Multiple issue boards on one group requires a GitLab Premium or above License.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_boards/)
//
// ## Import
//
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_issue_board`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_group_issue_board.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// Gitlab group issue boards can be imported with a key composed of `<group-id>:<issue-board-id>`, for example:
//
// ```sh
// $ pulumi import gitlab:index/groupIssueBoard:GroupIssueBoard example "12345:1"
// ```
type GroupIssueBoard struct {
	pulumi.CustomResourceState

	// The ID or URL-encoded path of the group owned by the authenticated user.
	Group pulumi.StringOutput `pulumi:"group"`
	// The list of label names which the board should be scoped to.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The list of issue board lists.
	Lists GroupIssueBoardListArrayOutput `pulumi:"lists"`
	// The milestone the board should be scoped to.
	MilestoneId pulumi.IntPtrOutput `pulumi:"milestoneId"`
	// The name of the board.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGroupIssueBoard registers a new resource with the given unique name, arguments, and options.
func NewGroupIssueBoard(ctx *pulumi.Context,
	name string, args *GroupIssueBoardArgs, opts ...pulumi.ResourceOption) (*GroupIssueBoard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupIssueBoard
	err := ctx.RegisterResource("gitlab:index/groupIssueBoard:GroupIssueBoard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupIssueBoard gets an existing GroupIssueBoard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupIssueBoard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupIssueBoardState, opts ...pulumi.ResourceOption) (*GroupIssueBoard, error) {
	var resource GroupIssueBoard
	err := ctx.ReadResource("gitlab:index/groupIssueBoard:GroupIssueBoard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupIssueBoard resources.
type groupIssueBoardState struct {
	// The ID or URL-encoded path of the group owned by the authenticated user.
	Group *string `pulumi:"group"`
	// The list of label names which the board should be scoped to.
	Labels []string `pulumi:"labels"`
	// The list of issue board lists.
	Lists []GroupIssueBoardList `pulumi:"lists"`
	// The milestone the board should be scoped to.
	MilestoneId *int `pulumi:"milestoneId"`
	// The name of the board.
	Name *string `pulumi:"name"`
}

type GroupIssueBoardState struct {
	// The ID or URL-encoded path of the group owned by the authenticated user.
	Group pulumi.StringPtrInput
	// The list of label names which the board should be scoped to.
	Labels pulumi.StringArrayInput
	// The list of issue board lists.
	Lists GroupIssueBoardListArrayInput
	// The milestone the board should be scoped to.
	MilestoneId pulumi.IntPtrInput
	// The name of the board.
	Name pulumi.StringPtrInput
}

func (GroupIssueBoardState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupIssueBoardState)(nil)).Elem()
}

type groupIssueBoardArgs struct {
	// The ID or URL-encoded path of the group owned by the authenticated user.
	Group string `pulumi:"group"`
	// The list of label names which the board should be scoped to.
	Labels []string `pulumi:"labels"`
	// The list of issue board lists.
	Lists []GroupIssueBoardList `pulumi:"lists"`
	// The milestone the board should be scoped to.
	MilestoneId *int `pulumi:"milestoneId"`
	// The name of the board.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a GroupIssueBoard resource.
type GroupIssueBoardArgs struct {
	// The ID or URL-encoded path of the group owned by the authenticated user.
	Group pulumi.StringInput
	// The list of label names which the board should be scoped to.
	Labels pulumi.StringArrayInput
	// The list of issue board lists.
	Lists GroupIssueBoardListArrayInput
	// The milestone the board should be scoped to.
	MilestoneId pulumi.IntPtrInput
	// The name of the board.
	Name pulumi.StringPtrInput
}

func (GroupIssueBoardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupIssueBoardArgs)(nil)).Elem()
}

type GroupIssueBoardInput interface {
	pulumi.Input

	ToGroupIssueBoardOutput() GroupIssueBoardOutput
	ToGroupIssueBoardOutputWithContext(ctx context.Context) GroupIssueBoardOutput
}

func (*GroupIssueBoard) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupIssueBoard)(nil)).Elem()
}

func (i *GroupIssueBoard) ToGroupIssueBoardOutput() GroupIssueBoardOutput {
	return i.ToGroupIssueBoardOutputWithContext(context.Background())
}

func (i *GroupIssueBoard) ToGroupIssueBoardOutputWithContext(ctx context.Context) GroupIssueBoardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupIssueBoardOutput)
}

// GroupIssueBoardArrayInput is an input type that accepts GroupIssueBoardArray and GroupIssueBoardArrayOutput values.
// You can construct a concrete instance of `GroupIssueBoardArrayInput` via:
//
//	GroupIssueBoardArray{ GroupIssueBoardArgs{...} }
type GroupIssueBoardArrayInput interface {
	pulumi.Input

	ToGroupIssueBoardArrayOutput() GroupIssueBoardArrayOutput
	ToGroupIssueBoardArrayOutputWithContext(context.Context) GroupIssueBoardArrayOutput
}

type GroupIssueBoardArray []GroupIssueBoardInput

func (GroupIssueBoardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupIssueBoard)(nil)).Elem()
}

func (i GroupIssueBoardArray) ToGroupIssueBoardArrayOutput() GroupIssueBoardArrayOutput {
	return i.ToGroupIssueBoardArrayOutputWithContext(context.Background())
}

func (i GroupIssueBoardArray) ToGroupIssueBoardArrayOutputWithContext(ctx context.Context) GroupIssueBoardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupIssueBoardArrayOutput)
}

// GroupIssueBoardMapInput is an input type that accepts GroupIssueBoardMap and GroupIssueBoardMapOutput values.
// You can construct a concrete instance of `GroupIssueBoardMapInput` via:
//
//	GroupIssueBoardMap{ "key": GroupIssueBoardArgs{...} }
type GroupIssueBoardMapInput interface {
	pulumi.Input

	ToGroupIssueBoardMapOutput() GroupIssueBoardMapOutput
	ToGroupIssueBoardMapOutputWithContext(context.Context) GroupIssueBoardMapOutput
}

type GroupIssueBoardMap map[string]GroupIssueBoardInput

func (GroupIssueBoardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupIssueBoard)(nil)).Elem()
}

func (i GroupIssueBoardMap) ToGroupIssueBoardMapOutput() GroupIssueBoardMapOutput {
	return i.ToGroupIssueBoardMapOutputWithContext(context.Background())
}

func (i GroupIssueBoardMap) ToGroupIssueBoardMapOutputWithContext(ctx context.Context) GroupIssueBoardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupIssueBoardMapOutput)
}

type GroupIssueBoardOutput struct{ *pulumi.OutputState }

func (GroupIssueBoardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupIssueBoard)(nil)).Elem()
}

func (o GroupIssueBoardOutput) ToGroupIssueBoardOutput() GroupIssueBoardOutput {
	return o
}

func (o GroupIssueBoardOutput) ToGroupIssueBoardOutputWithContext(ctx context.Context) GroupIssueBoardOutput {
	return o
}

// The ID or URL-encoded path of the group owned by the authenticated user.
func (o GroupIssueBoardOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupIssueBoard) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The list of label names which the board should be scoped to.
func (o GroupIssueBoardOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupIssueBoard) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The list of issue board lists.
func (o GroupIssueBoardOutput) Lists() GroupIssueBoardListArrayOutput {
	return o.ApplyT(func(v *GroupIssueBoard) GroupIssueBoardListArrayOutput { return v.Lists }).(GroupIssueBoardListArrayOutput)
}

// The milestone the board should be scoped to.
func (o GroupIssueBoardOutput) MilestoneId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GroupIssueBoard) pulumi.IntPtrOutput { return v.MilestoneId }).(pulumi.IntPtrOutput)
}

// The name of the board.
func (o GroupIssueBoardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupIssueBoard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type GroupIssueBoardArrayOutput struct{ *pulumi.OutputState }

func (GroupIssueBoardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupIssueBoard)(nil)).Elem()
}

func (o GroupIssueBoardArrayOutput) ToGroupIssueBoardArrayOutput() GroupIssueBoardArrayOutput {
	return o
}

func (o GroupIssueBoardArrayOutput) ToGroupIssueBoardArrayOutputWithContext(ctx context.Context) GroupIssueBoardArrayOutput {
	return o
}

func (o GroupIssueBoardArrayOutput) Index(i pulumi.IntInput) GroupIssueBoardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupIssueBoard {
		return vs[0].([]*GroupIssueBoard)[vs[1].(int)]
	}).(GroupIssueBoardOutput)
}

type GroupIssueBoardMapOutput struct{ *pulumi.OutputState }

func (GroupIssueBoardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupIssueBoard)(nil)).Elem()
}

func (o GroupIssueBoardMapOutput) ToGroupIssueBoardMapOutput() GroupIssueBoardMapOutput {
	return o
}

func (o GroupIssueBoardMapOutput) ToGroupIssueBoardMapOutputWithContext(ctx context.Context) GroupIssueBoardMapOutput {
	return o
}

func (o GroupIssueBoardMapOutput) MapIndex(k pulumi.StringInput) GroupIssueBoardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupIssueBoard {
		return vs[0].(map[string]*GroupIssueBoard)[vs[1].(string)]
	}).(GroupIssueBoardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupIssueBoardInput)(nil)).Elem(), &GroupIssueBoard{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupIssueBoardArrayInput)(nil)).Elem(), GroupIssueBoardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupIssueBoardMapInput)(nil)).Elem(), GroupIssueBoardMap{})
	pulumi.RegisterOutputType(GroupIssueBoardOutput{})
	pulumi.RegisterOutputType(GroupIssueBoardArrayOutput{})
	pulumi.RegisterOutputType(GroupIssueBoardMapOutput{})
}
