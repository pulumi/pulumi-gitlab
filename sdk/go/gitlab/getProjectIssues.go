// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjectIssues` data source allows to retrieve details about issues in a project.
//
// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/issues/)
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
//			foo, err := gitlab.LookupProject(ctx, &gitlab.LookupProjectArgs{
//				Id: pulumi.StringRef("foo/bar/baz"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.GetProjectIssues(ctx, &gitlab.GetProjectIssuesArgs{
//				Project: foo.Id,
//				Search:  pulumi.StringRef("foo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjectIssues(ctx *pulumi.Context, args *GetProjectIssuesArgs, opts ...pulumi.InvokeOption) (*GetProjectIssuesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectIssuesResult
	err := ctx.Invoke("gitlab:index/getProjectIssues:getProjectIssues", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectIssues.
type GetProjectIssuesArgs struct {
	// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
	AssigneeId *int `pulumi:"assigneeId"`
	// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
	AssigneeUsername *string `pulumi:"assigneeUsername"`
	// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
	AuthorId *int `pulumi:"authorId"`
	// Filter confidential or public issues.
	Confidential *bool `pulumi:"confidential"`
	// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	CreatedAfter *string `pulumi:"createdAfter"`
	// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	CreatedBefore *string `pulumi:"createdBefore"`
	// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
	DueDate *string `pulumi:"dueDate"`
	// Return only the issues having the given iid
	Iids []int `pulumi:"iids"`
	// Filter to a given type of issue. Valid values are [issue incident testCase].
	IssueType *string `pulumi:"issueType"`
	// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
	Labels []string `pulumi:"labels"`
	// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
	Milestone *string `pulumi:"milestone"`
	// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
	MyReactionEmoji *string `pulumi:"myReactionEmoji"`
	// Return issues that do not match the assignee id.
	NotAssigneeId *int `pulumi:"notAssigneeId"`
	// Return issues that do not match the author id.
	NotAuthorId *int `pulumi:"notAuthorId"`
	// Return issues that do not match the labels.
	NotLabels []string `pulumi:"notLabels"`
	// Return issues that do not match the milestone.
	NotMilestone *string `pulumi:"notMilestone"`
	// Return issues not reacted by the authenticated user by the given emoji.
	NotMyReactionEmoji *string `pulumi:"notMyReactionEmoji"`
	// Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
	OrderBy *string `pulumi:"orderBy"`
	// The name or id of the project.
	Project string `pulumi:"project"`
	// Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
	Scope *string `pulumi:"scope"`
	// Search project issues against their title and description
	Search *string `pulumi:"search"`
	// Return issues sorted in asc or desc order. Default is desc
	Sort *string `pulumi:"sort"`
	// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	UpdatedAfter *string `pulumi:"updatedAfter"`
	// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	UpdatedBefore *string `pulumi:"updatedBefore"`
	// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
	Weight *int `pulumi:"weight"`
	// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
	WithLabelsDetails *bool `pulumi:"withLabelsDetails"`
}

// A collection of values returned by getProjectIssues.
type GetProjectIssuesResult struct {
	// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
	AssigneeId *int `pulumi:"assigneeId"`
	// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
	AssigneeUsername *string `pulumi:"assigneeUsername"`
	// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
	AuthorId *int `pulumi:"authorId"`
	// Filter confidential or public issues.
	Confidential *bool `pulumi:"confidential"`
	// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	CreatedAfter *string `pulumi:"createdAfter"`
	// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	CreatedBefore *string `pulumi:"createdBefore"`
	// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
	DueDate *string `pulumi:"dueDate"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Return only the issues having the given iid
	Iids []int `pulumi:"iids"`
	// Filter to a given type of issue. Valid values are [issue incident testCase].
	IssueType *string `pulumi:"issueType"`
	// The list of issues returned by the search.
	Issues []GetProjectIssuesIssue `pulumi:"issues"`
	// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
	Labels []string `pulumi:"labels"`
	// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
	Milestone *string `pulumi:"milestone"`
	// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
	MyReactionEmoji *string `pulumi:"myReactionEmoji"`
	// Return issues that do not match the assignee id.
	NotAssigneeId *int `pulumi:"notAssigneeId"`
	// Return issues that do not match the author id.
	NotAuthorId *int `pulumi:"notAuthorId"`
	// Return issues that do not match the labels.
	NotLabels []string `pulumi:"notLabels"`
	// Return issues that do not match the milestone.
	NotMilestone *string `pulumi:"notMilestone"`
	// Return issues not reacted by the authenticated user by the given emoji.
	NotMyReactionEmoji *string `pulumi:"notMyReactionEmoji"`
	// Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
	OrderBy *string `pulumi:"orderBy"`
	// The name or id of the project.
	Project string `pulumi:"project"`
	// Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
	Scope *string `pulumi:"scope"`
	// Search project issues against their title and description
	Search *string `pulumi:"search"`
	// Return issues sorted in asc or desc order. Default is desc
	Sort *string `pulumi:"sort"`
	// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	UpdatedAfter *string `pulumi:"updatedAfter"`
	// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	UpdatedBefore *string `pulumi:"updatedBefore"`
	// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
	Weight *int `pulumi:"weight"`
	// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
	WithLabelsDetails *bool `pulumi:"withLabelsDetails"`
}

func GetProjectIssuesOutput(ctx *pulumi.Context, args GetProjectIssuesOutputArgs, opts ...pulumi.InvokeOption) GetProjectIssuesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectIssuesResultOutput, error) {
			args := v.(GetProjectIssuesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gitlab:index/getProjectIssues:getProjectIssues", args, GetProjectIssuesResultOutput{}, options).(GetProjectIssuesResultOutput), nil
		}).(GetProjectIssuesResultOutput)
}

// A collection of arguments for invoking getProjectIssues.
type GetProjectIssuesOutputArgs struct {
	// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
	AssigneeId pulumi.IntPtrInput `pulumi:"assigneeId"`
	// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
	AssigneeUsername pulumi.StringPtrInput `pulumi:"assigneeUsername"`
	// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
	AuthorId pulumi.IntPtrInput `pulumi:"authorId"`
	// Filter confidential or public issues.
	Confidential pulumi.BoolPtrInput `pulumi:"confidential"`
	// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	CreatedAfter pulumi.StringPtrInput `pulumi:"createdAfter"`
	// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	CreatedBefore pulumi.StringPtrInput `pulumi:"createdBefore"`
	// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
	DueDate pulumi.StringPtrInput `pulumi:"dueDate"`
	// Return only the issues having the given iid
	Iids pulumi.IntArrayInput `pulumi:"iids"`
	// Filter to a given type of issue. Valid values are [issue incident testCase].
	IssueType pulumi.StringPtrInput `pulumi:"issueType"`
	// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
	Labels pulumi.StringArrayInput `pulumi:"labels"`
	// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
	Milestone pulumi.StringPtrInput `pulumi:"milestone"`
	// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
	MyReactionEmoji pulumi.StringPtrInput `pulumi:"myReactionEmoji"`
	// Return issues that do not match the assignee id.
	NotAssigneeId pulumi.IntPtrInput `pulumi:"notAssigneeId"`
	// Return issues that do not match the author id.
	NotAuthorId pulumi.IntPtrInput `pulumi:"notAuthorId"`
	// Return issues that do not match the labels.
	NotLabels pulumi.StringArrayInput `pulumi:"notLabels"`
	// Return issues that do not match the milestone.
	NotMilestone pulumi.StringPtrInput `pulumi:"notMilestone"`
	// Return issues not reacted by the authenticated user by the given emoji.
	NotMyReactionEmoji pulumi.StringPtrInput `pulumi:"notMyReactionEmoji"`
	// Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// The name or id of the project.
	Project pulumi.StringInput `pulumi:"project"`
	// Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
	Scope pulumi.StringPtrInput `pulumi:"scope"`
	// Search project issues against their title and description
	Search pulumi.StringPtrInput `pulumi:"search"`
	// Return issues sorted in asc or desc order. Default is desc
	Sort pulumi.StringPtrInput `pulumi:"sort"`
	// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	UpdatedAfter pulumi.StringPtrInput `pulumi:"updatedAfter"`
	// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	UpdatedBefore pulumi.StringPtrInput `pulumi:"updatedBefore"`
	// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
	Weight pulumi.IntPtrInput `pulumi:"weight"`
	// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
	WithLabelsDetails pulumi.BoolPtrInput `pulumi:"withLabelsDetails"`
}

func (GetProjectIssuesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectIssuesArgs)(nil)).Elem()
}

// A collection of values returned by getProjectIssues.
type GetProjectIssuesResultOutput struct{ *pulumi.OutputState }

func (GetProjectIssuesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectIssuesResult)(nil)).Elem()
}

func (o GetProjectIssuesResultOutput) ToGetProjectIssuesResultOutput() GetProjectIssuesResultOutput {
	return o
}

func (o GetProjectIssuesResultOutput) ToGetProjectIssuesResultOutputWithContext(ctx context.Context) GetProjectIssuesResultOutput {
	return o
}

// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
func (o GetProjectIssuesResultOutput) AssigneeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *int { return v.AssigneeId }).(pulumi.IntPtrOutput)
}

// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
func (o GetProjectIssuesResultOutput) AssigneeUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.AssigneeUsername }).(pulumi.StringPtrOutput)
}

// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
func (o GetProjectIssuesResultOutput) AuthorId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *int { return v.AuthorId }).(pulumi.IntPtrOutput)
}

// Filter confidential or public issues.
func (o GetProjectIssuesResultOutput) Confidential() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *bool { return v.Confidential }).(pulumi.BoolPtrOutput)
}

// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
func (o GetProjectIssuesResultOutput) CreatedAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.CreatedAfter }).(pulumi.StringPtrOutput)
}

// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
func (o GetProjectIssuesResultOutput) CreatedBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.CreatedBefore }).(pulumi.StringPtrOutput)
}

// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
func (o GetProjectIssuesResultOutput) DueDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.DueDate }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectIssuesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Return only the issues having the given iid
func (o GetProjectIssuesResultOutput) Iids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) []int { return v.Iids }).(pulumi.IntArrayOutput)
}

// Filter to a given type of issue. Valid values are [issue incident testCase].
func (o GetProjectIssuesResultOutput) IssueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.IssueType }).(pulumi.StringPtrOutput)
}

// The list of issues returned by the search.
func (o GetProjectIssuesResultOutput) Issues() GetProjectIssuesIssueArrayOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) []GetProjectIssuesIssue { return v.Issues }).(GetProjectIssuesIssueArrayOutput)
}

// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
func (o GetProjectIssuesResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
func (o GetProjectIssuesResultOutput) Milestone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.Milestone }).(pulumi.StringPtrOutput)
}

// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
func (o GetProjectIssuesResultOutput) MyReactionEmoji() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.MyReactionEmoji }).(pulumi.StringPtrOutput)
}

// Return issues that do not match the assignee id.
func (o GetProjectIssuesResultOutput) NotAssigneeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *int { return v.NotAssigneeId }).(pulumi.IntPtrOutput)
}

// Return issues that do not match the author id.
func (o GetProjectIssuesResultOutput) NotAuthorId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *int { return v.NotAuthorId }).(pulumi.IntPtrOutput)
}

// Return issues that do not match the labels.
func (o GetProjectIssuesResultOutput) NotLabels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) []string { return v.NotLabels }).(pulumi.StringArrayOutput)
}

// Return issues that do not match the milestone.
func (o GetProjectIssuesResultOutput) NotMilestone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.NotMilestone }).(pulumi.StringPtrOutput)
}

// Return issues not reacted by the authenticated user by the given emoji.
func (o GetProjectIssuesResultOutput) NotMyReactionEmoji() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.NotMyReactionEmoji }).(pulumi.StringPtrOutput)
}

// Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
func (o GetProjectIssuesResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

// The name or id of the project.
func (o GetProjectIssuesResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) string { return v.Project }).(pulumi.StringOutput)
}

// Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
func (o GetProjectIssuesResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Search project issues against their title and description
func (o GetProjectIssuesResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

// Return issues sorted in asc or desc order. Default is desc
func (o GetProjectIssuesResultOutput) Sort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.Sort }).(pulumi.StringPtrOutput)
}

// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
func (o GetProjectIssuesResultOutput) UpdatedAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.UpdatedAfter }).(pulumi.StringPtrOutput)
}

// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
func (o GetProjectIssuesResultOutput) UpdatedBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *string { return v.UpdatedBefore }).(pulumi.StringPtrOutput)
}

// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
func (o GetProjectIssuesResultOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
func (o GetProjectIssuesResultOutput) WithLabelsDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectIssuesResult) *bool { return v.WithLabelsDetails }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectIssuesResultOutput{})
}
