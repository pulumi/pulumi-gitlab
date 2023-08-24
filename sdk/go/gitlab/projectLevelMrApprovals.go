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
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooProject, err := gitlab.NewProject(ctx, "fooProject", &gitlab.ProjectArgs{
//				Description: pulumi.String("My example project"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectLevelMrApprovals(ctx, "fooProjectLevelMrApprovals", &gitlab.ProjectLevelMrApprovalsArgs{
//				Project:              fooProject.ID(),
//				ResetApprovalsOnPush: pulumi.Bool(true),
//				DisableOverridingApproversPerMergeRequest: pulumi.Bool(false),
//				MergeRequestsAuthorApproval:               pulumi.Bool(false),
//				MergeRequestsDisableCommittersApproval:    pulumi.Bool(true),
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
// ```sh
//
//	$ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals You can import an approval configuration state using `<resource> <project_id>`.
//
// ```
//
// # For example
//
// ```sh
//
//	$ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
//
// ```
type ProjectLevelMrApprovals struct {
	pulumi.CustomResourceState

	// Set to `true` to disable overriding approvers per merge request.
	DisableOverridingApproversPerMergeRequest pulumi.BoolOutput `pulumi:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` to allow merge requests authors to approve their own merge requests.
	MergeRequestsAuthorApproval pulumi.BoolOutput `pulumi:"mergeRequestsAuthorApproval"`
	// Set to `true` to allow merge requests committers to approve their own merge requests.
	MergeRequestsDisableCommittersApproval pulumi.BoolOutput `pulumi:"mergeRequestsDisableCommittersApproval"`
	// The ID or URL-encoded path of a project to change MR approval configuration.
	Project pulumi.StringOutput `pulumi:"project"`
	// Set to `true` to require authentication to approve merge requests.
	RequirePasswordToApprove pulumi.BoolOutput `pulumi:"requirePasswordToApprove"`
	// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
	ResetApprovalsOnPush pulumi.BoolOutput `pulumi:"resetApprovalsOnPush"`
	// Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
	SelectiveCodeOwnerRemovals pulumi.BoolOutput `pulumi:"selectiveCodeOwnerRemovals"`
}

// NewProjectLevelMrApprovals registers a new resource with the given unique name, arguments, and options.
func NewProjectLevelMrApprovals(ctx *pulumi.Context,
	name string, args *ProjectLevelMrApprovalsArgs, opts ...pulumi.ResourceOption) (*ProjectLevelMrApprovals, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Set to `true` to disable overriding approvers per merge request.
	DisableOverridingApproversPerMergeRequest *bool `pulumi:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` to allow merge requests authors to approve their own merge requests.
	MergeRequestsAuthorApproval *bool `pulumi:"mergeRequestsAuthorApproval"`
	// Set to `true` to allow merge requests committers to approve their own merge requests.
	MergeRequestsDisableCommittersApproval *bool `pulumi:"mergeRequestsDisableCommittersApproval"`
	// The ID or URL-encoded path of a project to change MR approval configuration.
	Project *string `pulumi:"project"`
	// Set to `true` to require authentication to approve merge requests.
	RequirePasswordToApprove *bool `pulumi:"requirePasswordToApprove"`
	// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
	ResetApprovalsOnPush *bool `pulumi:"resetApprovalsOnPush"`
	// Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
	SelectiveCodeOwnerRemovals *bool `pulumi:"selectiveCodeOwnerRemovals"`
}

type ProjectLevelMrApprovalsState struct {
	// Set to `true` to disable overriding approvers per merge request.
	DisableOverridingApproversPerMergeRequest pulumi.BoolPtrInput
	// Set to `true` to allow merge requests authors to approve their own merge requests.
	MergeRequestsAuthorApproval pulumi.BoolPtrInput
	// Set to `true` to allow merge requests committers to approve their own merge requests.
	MergeRequestsDisableCommittersApproval pulumi.BoolPtrInput
	// The ID or URL-encoded path of a project to change MR approval configuration.
	Project pulumi.StringPtrInput
	// Set to `true` to require authentication to approve merge requests.
	RequirePasswordToApprove pulumi.BoolPtrInput
	// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
	ResetApprovalsOnPush pulumi.BoolPtrInput
	// Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
	SelectiveCodeOwnerRemovals pulumi.BoolPtrInput
}

func (ProjectLevelMrApprovalsState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLevelMrApprovalsState)(nil)).Elem()
}

type projectLevelMrApprovalsArgs struct {
	// Set to `true` to disable overriding approvers per merge request.
	DisableOverridingApproversPerMergeRequest *bool `pulumi:"disableOverridingApproversPerMergeRequest"`
	// Set to `true` to allow merge requests authors to approve their own merge requests.
	MergeRequestsAuthorApproval *bool `pulumi:"mergeRequestsAuthorApproval"`
	// Set to `true` to allow merge requests committers to approve their own merge requests.
	MergeRequestsDisableCommittersApproval *bool `pulumi:"mergeRequestsDisableCommittersApproval"`
	// The ID or URL-encoded path of a project to change MR approval configuration.
	Project string `pulumi:"project"`
	// Set to `true` to require authentication to approve merge requests.
	RequirePasswordToApprove *bool `pulumi:"requirePasswordToApprove"`
	// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
	ResetApprovalsOnPush *bool `pulumi:"resetApprovalsOnPush"`
	// Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
	SelectiveCodeOwnerRemovals *bool `pulumi:"selectiveCodeOwnerRemovals"`
}

// The set of arguments for constructing a ProjectLevelMrApprovals resource.
type ProjectLevelMrApprovalsArgs struct {
	// Set to `true` to disable overriding approvers per merge request.
	DisableOverridingApproversPerMergeRequest pulumi.BoolPtrInput
	// Set to `true` to allow merge requests authors to approve their own merge requests.
	MergeRequestsAuthorApproval pulumi.BoolPtrInput
	// Set to `true` to allow merge requests committers to approve their own merge requests.
	MergeRequestsDisableCommittersApproval pulumi.BoolPtrInput
	// The ID or URL-encoded path of a project to change MR approval configuration.
	Project pulumi.StringInput
	// Set to `true` to require authentication to approve merge requests.
	RequirePasswordToApprove pulumi.BoolPtrInput
	// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
	ResetApprovalsOnPush pulumi.BoolPtrInput
	// Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
	SelectiveCodeOwnerRemovals pulumi.BoolPtrInput
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
//	ProjectLevelMrApprovalsArray{ ProjectLevelMrApprovalsArgs{...} }
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
//	ProjectLevelMrApprovalsMap{ "key": ProjectLevelMrApprovalsArgs{...} }
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

// Set to `true` to disable overriding approvers per merge request.
func (o ProjectLevelMrApprovalsOutput) DisableOverridingApproversPerMergeRequest() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.BoolOutput { return v.DisableOverridingApproversPerMergeRequest }).(pulumi.BoolOutput)
}

// Set to `true` to allow merge requests authors to approve their own merge requests.
func (o ProjectLevelMrApprovalsOutput) MergeRequestsAuthorApproval() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.BoolOutput { return v.MergeRequestsAuthorApproval }).(pulumi.BoolOutput)
}

// Set to `true` to allow merge requests committers to approve their own merge requests.
func (o ProjectLevelMrApprovalsOutput) MergeRequestsDisableCommittersApproval() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.BoolOutput { return v.MergeRequestsDisableCommittersApproval }).(pulumi.BoolOutput)
}

// The ID or URL-encoded path of a project to change MR approval configuration.
func (o ProjectLevelMrApprovalsOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Set to `true` to require authentication to approve merge requests.
func (o ProjectLevelMrApprovalsOutput) RequirePasswordToApprove() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.BoolOutput { return v.RequirePasswordToApprove }).(pulumi.BoolOutput)
}

// Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
func (o ProjectLevelMrApprovalsOutput) ResetApprovalsOnPush() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.BoolOutput { return v.ResetApprovalsOnPush }).(pulumi.BoolOutput)
}

// Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
func (o ProjectLevelMrApprovalsOutput) SelectiveCodeOwnerRemovals() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectLevelMrApprovals) pulumi.BoolOutput { return v.SelectiveCodeOwnerRemovals }).(pulumi.BoolOutput)
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
