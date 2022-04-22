// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `gitlabProjectLevelMrApprovalRule` resource allows to manage the lifecycle of a Merge Request-level approval rule.
//
// > This resource requires a GitLab Enterprise instance.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/merge_request_approvals.html#merge-request-level-mr-approvals)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooProject, err := gitlab.NewProject(ctx, "fooProject", &gitlab.ProjectArgs{
// 			Description: pulumi.String("My example project"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewProjectLevelMrApprovals(ctx, "fooProjectLevelMrApprovals", &gitlab.ProjectLevelMrApprovalsArgs{
// 			ProjectId:            fooProject.ID(),
// 			ResetApprovalsOnPush: pulumi.Bool(true),
// 			DisableOverridingApproversPerMergeRequest: pulumi.Bool(false),
// 			MergeRequestsAuthorApproval:               pulumi.Bool(false),
// 			MergeRequestsDisableCommittersApproval:    pulumi.Bool(true),
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
// ```sh
//  $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals # You can import an approval configuration state using `<resource> <project_id>`.
// ```
//
// # # For example
//
// ```sh
//  $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
// ```
type ProjectLevelMrApprovals struct {
	pulumi.CustomResourceState

	// By default, users are able to edit the approval rules in merge requests. If set to true,
	DisableOverridingApproversPerMergeRequest pulumi.BoolPtrOutput `pulumi:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
	MergeRequestsAuthorApproval pulumi.BoolPtrOutput `pulumi:"mergeRequestsAuthorApproval"`
	// Set to `true` if you want to prevent approval of merge requests by merge request committers.
	MergeRequestsDisableCommittersApproval pulumi.BoolPtrOutput `pulumi:"mergeRequestsDisableCommittersApproval"`
	// The ID of the project to change MR approval configuration.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Set to `true` if you want to require authentication when approving a merge request.
	RequirePasswordToApprove pulumi.BoolPtrOutput `pulumi:"requirePasswordToApprove"`
	// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch.
	// Default is `true`.
	ResetApprovalsOnPush pulumi.BoolPtrOutput `pulumi:"resetApprovalsOnPush"`
}

// NewProjectLevelMrApprovals registers a new resource with the given unique name, arguments, and options.
func NewProjectLevelMrApprovals(ctx *pulumi.Context,
	name string, args *ProjectLevelMrApprovalsArgs, opts ...pulumi.ResourceOption) (*ProjectLevelMrApprovals, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource ProjectLevelMrApprovals
	err := ctx.RegisterResource("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectLevelMrApprovals gets an existing ProjectLevelMrApprovals resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectLevelMrApprovals(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectLevelMrApprovalsState, opts ...pulumi.ResourceOption) (*ProjectLevelMrApprovals, error) {
	var resource ProjectLevelMrApprovals
	err := ctx.ReadResource("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectLevelMrApprovals resources.
type projectLevelMrApprovalsState struct {
	// By default, users are able to edit the approval rules in merge requests. If set to true,
	DisableOverridingApproversPerMergeRequest *bool `pulumi:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
	MergeRequestsAuthorApproval *bool `pulumi:"mergeRequestsAuthorApproval"`
	// Set to `true` if you want to prevent approval of merge requests by merge request committers.
	MergeRequestsDisableCommittersApproval *bool `pulumi:"mergeRequestsDisableCommittersApproval"`
	// The ID of the project to change MR approval configuration.
	ProjectId *int `pulumi:"projectId"`
	// Set to `true` if you want to require authentication when approving a merge request.
	RequirePasswordToApprove *bool `pulumi:"requirePasswordToApprove"`
	// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch.
	// Default is `true`.
	ResetApprovalsOnPush *bool `pulumi:"resetApprovalsOnPush"`
}

type ProjectLevelMrApprovalsState struct {
	// By default, users are able to edit the approval rules in merge requests. If set to true,
	DisableOverridingApproversPerMergeRequest pulumi.BoolPtrInput
	// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
	MergeRequestsAuthorApproval pulumi.BoolPtrInput
	// Set to `true` if you want to prevent approval of merge requests by merge request committers.
	MergeRequestsDisableCommittersApproval pulumi.BoolPtrInput
	// The ID of the project to change MR approval configuration.
	ProjectId pulumi.IntPtrInput
	// Set to `true` if you want to require authentication when approving a merge request.
	RequirePasswordToApprove pulumi.BoolPtrInput
	// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch.
	// Default is `true`.
	ResetApprovalsOnPush pulumi.BoolPtrInput
}

func (ProjectLevelMrApprovalsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLevelMrApprovalsState)(nil)).Elem()
}

type projectLevelMrApprovalsArgs struct {
	// By default, users are able to edit the approval rules in merge requests. If set to true,
	DisableOverridingApproversPerMergeRequest *bool `pulumi:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
	MergeRequestsAuthorApproval *bool `pulumi:"mergeRequestsAuthorApproval"`
	// Set to `true` if you want to prevent approval of merge requests by merge request committers.
	MergeRequestsDisableCommittersApproval *bool `pulumi:"mergeRequestsDisableCommittersApproval"`
	// The ID of the project to change MR approval configuration.
	ProjectId int `pulumi:"projectId"`
	// Set to `true` if you want to require authentication when approving a merge request.
	RequirePasswordToApprove *bool `pulumi:"requirePasswordToApprove"`
	// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch.
	// Default is `true`.
	ResetApprovalsOnPush *bool `pulumi:"resetApprovalsOnPush"`
}

// The set of arguments for constructing a ProjectLevelMrApprovals resource.
type ProjectLevelMrApprovalsArgs struct {
	// By default, users are able to edit the approval rules in merge requests. If set to true,
	DisableOverridingApproversPerMergeRequest pulumi.BoolPtrInput
	// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
	MergeRequestsAuthorApproval pulumi.BoolPtrInput
	// Set to `true` if you want to prevent approval of merge requests by merge request committers.
	MergeRequestsDisableCommittersApproval pulumi.BoolPtrInput
	// The ID of the project to change MR approval configuration.
	ProjectId pulumi.IntInput
	// Set to `true` if you want to require authentication when approving a merge request.
	RequirePasswordToApprove pulumi.BoolPtrInput
	// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch.
	// Default is `true`.
	ResetApprovalsOnPush pulumi.BoolPtrInput
}

func (ProjectLevelMrApprovalsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLevelMrApprovalsArgs)(nil)).Elem()
}

type ProjectLevelMrApprovalsInput interface {
	pulumi.Input

	ToProjectLevelMrApprovalsOutput() ProjectLevelMrApprovalsOutput
	ToProjectLevelMrApprovalsOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsOutput
}

func (*ProjectLevelMrApprovals) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLevelMrApprovals)(nil)).Elem()
}

func (i *ProjectLevelMrApprovals) ToProjectLevelMrApprovalsOutput() ProjectLevelMrApprovalsOutput {
	return i.ToProjectLevelMrApprovalsOutputWithContext(context.Background())
}

func (i *ProjectLevelMrApprovals) ToProjectLevelMrApprovalsOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLevelMrApprovalsOutput)
}

// ProjectLevelMrApprovalsArrayInput is an input type that accepts ProjectLevelMrApprovalsArray and ProjectLevelMrApprovalsArrayOutput values.
// You can construct a concrete instance of `ProjectLevelMrApprovalsArrayInput` via:
//
//          ProjectLevelMrApprovalsArray{ ProjectLevelMrApprovalsArgs{...} }
type ProjectLevelMrApprovalsArrayInput interface {
	pulumi.Input

	ToProjectLevelMrApprovalsArrayOutput() ProjectLevelMrApprovalsArrayOutput
	ToProjectLevelMrApprovalsArrayOutputWithContext(context.Context) ProjectLevelMrApprovalsArrayOutput
}

type ProjectLevelMrApprovalsArray []ProjectLevelMrApprovalsInput

func (ProjectLevelMrApprovalsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectLevelMrApprovals)(nil)).Elem()
}

func (i ProjectLevelMrApprovalsArray) ToProjectLevelMrApprovalsArrayOutput() ProjectLevelMrApprovalsArrayOutput {
	return i.ToProjectLevelMrApprovalsArrayOutputWithContext(context.Background())
}

func (i ProjectLevelMrApprovalsArray) ToProjectLevelMrApprovalsArrayOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLevelMrApprovalsArrayOutput)
}

// ProjectLevelMrApprovalsMapInput is an input type that accepts ProjectLevelMrApprovalsMap and ProjectLevelMrApprovalsMapOutput values.
// You can construct a concrete instance of `ProjectLevelMrApprovalsMapInput` via:
//
//          ProjectLevelMrApprovalsMap{ "key": ProjectLevelMrApprovalsArgs{...} }
type ProjectLevelMrApprovalsMapInput interface {
	pulumi.Input

	ToProjectLevelMrApprovalsMapOutput() ProjectLevelMrApprovalsMapOutput
	ToProjectLevelMrApprovalsMapOutputWithContext(context.Context) ProjectLevelMrApprovalsMapOutput
}

type ProjectLevelMrApprovalsMap map[string]ProjectLevelMrApprovalsInput

func (ProjectLevelMrApprovalsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectLevelMrApprovals)(nil)).Elem()
}

func (i ProjectLevelMrApprovalsMap) ToProjectLevelMrApprovalsMapOutput() ProjectLevelMrApprovalsMapOutput {
	return i.ToProjectLevelMrApprovalsMapOutputWithContext(context.Background())
}

func (i ProjectLevelMrApprovalsMap) ToProjectLevelMrApprovalsMapOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLevelMrApprovalsMapOutput)
}

type ProjectLevelMrApprovalsOutput struct{ *pulumi.OutputState }

func (ProjectLevelMrApprovalsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLevelMrApprovals)(nil)).Elem()
}

func (o ProjectLevelMrApprovalsOutput) ToProjectLevelMrApprovalsOutput() ProjectLevelMrApprovalsOutput {
	return o
}

func (o ProjectLevelMrApprovalsOutput) ToProjectLevelMrApprovalsOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsOutput {
	return o
}

type ProjectLevelMrApprovalsArrayOutput struct{ *pulumi.OutputState }

func (ProjectLevelMrApprovalsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectLevelMrApprovals)(nil)).Elem()
}

func (o ProjectLevelMrApprovalsArrayOutput) ToProjectLevelMrApprovalsArrayOutput() ProjectLevelMrApprovalsArrayOutput {
	return o
}

func (o ProjectLevelMrApprovalsArrayOutput) ToProjectLevelMrApprovalsArrayOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsArrayOutput {
	return o
}

func (o ProjectLevelMrApprovalsArrayOutput) Index(i pulumi.IntInput) ProjectLevelMrApprovalsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectLevelMrApprovals {
		return vs[0].([]*ProjectLevelMrApprovals)[vs[1].(int)]
	}).(ProjectLevelMrApprovalsOutput)
}

type ProjectLevelMrApprovalsMapOutput struct{ *pulumi.OutputState }

func (ProjectLevelMrApprovalsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectLevelMrApprovals)(nil)).Elem()
}

func (o ProjectLevelMrApprovalsMapOutput) ToProjectLevelMrApprovalsMapOutput() ProjectLevelMrApprovalsMapOutput {
	return o
}

func (o ProjectLevelMrApprovalsMapOutput) ToProjectLevelMrApprovalsMapOutputWithContext(ctx context.Context) ProjectLevelMrApprovalsMapOutput {
	return o
}

func (o ProjectLevelMrApprovalsMapOutput) MapIndex(k pulumi.StringInput) ProjectLevelMrApprovalsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectLevelMrApprovals {
		return vs[0].(map[string]*ProjectLevelMrApprovals)[vs[1].(string)]
	}).(ProjectLevelMrApprovalsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLevelMrApprovalsInput)(nil)).Elem(), &ProjectLevelMrApprovals{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLevelMrApprovalsArrayInput)(nil)).Elem(), ProjectLevelMrApprovalsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLevelMrApprovalsMapInput)(nil)).Elem(), ProjectLevelMrApprovalsMap{})
	pulumi.RegisterOutputType(ProjectLevelMrApprovalsOutput{})
	pulumi.RegisterOutputType(ProjectLevelMrApprovalsArrayOutput{})
	pulumi.RegisterOutputType(ProjectLevelMrApprovalsMapOutput{})
}
