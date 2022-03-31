// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := gitlab.NewProject(ctx, "foo", &gitlab.ProjectArgs{
// 			Description:     pulumi.String("Lorem Ipsum"),
// 			VisibilityLevel: pulumi.String("public"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewProjectIssue(ctx, "welcomeIssue", &gitlab.ProjectIssueArgs{
// 			Project: foo.ID(),
// 			Title:   pulumi.String("Welcome!"),
// 			Description: foo.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v", "  Welcome to the ", name, " project!\n", "\n"), nil
// 			}).(pulumi.StringOutput),
// 			DiscussionLocked: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("welcomeIssueWebUrl", data.Gitlab_project_issue.Web_url)
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # You can import this resource with an id made up of `{project-id}:{issue-id}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/projectIssue:ProjectIssue welcome_issue 42:1
// ```
type ProjectIssue struct {
	pulumi.CustomResourceState

	// The IDs of the users to assign the issue to.
	AssigneeIds pulumi.IntArrayOutput `pulumi:"assigneeIds"`
	// The ID of the author of the issue. Use `User` data source to get more information about the user.
	AuthorId pulumi.IntOutput `pulumi:"authorId"`
	// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	ClosedAt pulumi.StringOutput `pulumi:"closedAt"`
	// The ID of the user that closed the issue. Use `User` data source to get more information about the user.
	ClosedByUserId pulumi.IntOutput `pulumi:"closedByUserId"`
	// Set an issue to be confidential.
	Confidential pulumi.BoolPtrOutput `pulumi:"confidential"`
	// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Whether the issue is deleted instead of closed during destroy.
	DeleteOnDestroy pulumi.BoolPtrOutput `pulumi:"deleteOnDestroy"`
	// The description of an issue. Limited to 1,048,576 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the issue is locked for discussions or not.
	DiscussionLocked pulumi.BoolPtrOutput `pulumi:"discussionLocked"`
	// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
	DiscussionToResolve pulumi.StringPtrOutput `pulumi:"discussionToResolve"`
	// The number of downvotes the issue has received.
	Downvotes pulumi.IntOutput `pulumi:"downvotes"`
	// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
	DueDate pulumi.StringPtrOutput `pulumi:"dueDate"`
	// ID of the epic to add the issue to. Valid values are greater than or equal to 0.
	EpicId pulumi.IntOutput `pulumi:"epicId"`
	// The ID of the epic issue.
	EpicIssueId pulumi.IntOutput `pulumi:"epicIssueId"`
	// The external ID of the issue.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// The human-readable time estimate of the issue.
	HumanTimeEstimate pulumi.StringOutput `pulumi:"humanTimeEstimate"`
	// The human-readable total time spent of the issue.
	HumanTotalTimeSpent pulumi.StringOutput `pulumi:"humanTotalTimeSpent"`
	// The internal ID of the project's issue.
	Iid pulumi.IntOutput `pulumi:"iid"`
	// The instance-wide ID of the issue.
	IssueId pulumi.IntOutput `pulumi:"issueId"`
	// The ID of the issue link.
	IssueLinkId pulumi.IntOutput `pulumi:"issueLinkId"`
	// The type of issue. Valid values are: `issue`, `incident`, `testCase`.
	IssueType pulumi.StringPtrOutput `pulumi:"issueType"`
	// The labels of an issue.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The links of the issue.
	Links pulumi.StringMapOutput `pulumi:"links"`
	// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
	MergeRequestToResolveDiscussionsOf pulumi.IntPtrOutput `pulumi:"mergeRequestToResolveDiscussionsOf"`
	// The number of merge requests associated with the issue.
	MergeRequestsCount pulumi.IntOutput `pulumi:"mergeRequestsCount"`
	// The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
	MilestoneId pulumi.IntPtrOutput `pulumi:"milestoneId"`
	// The ID of the issue that was moved to.
	MovedToId pulumi.IntOutput `pulumi:"movedToId"`
	// The name or ID of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// The references of the issue.
	References pulumi.StringMapOutput `pulumi:"references"`
	// The state of the issue. Valid values are: `opened`, `closed`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Whether the authenticated user is subscribed to the issue or not.
	Subscribed pulumi.BoolOutput `pulumi:"subscribed"`
	// The task completion status. It's always a one element list.
	TaskCompletionStatus ProjectIssueTaskCompletionStatusOutput `pulumi:"taskCompletionStatus"`
	// The time estimate of the issue.
	TimeEstimate pulumi.IntOutput `pulumi:"timeEstimate"`
	// The title of the issue.
	Title pulumi.StringOutput `pulumi:"title"`
	// The total time spent of the issue.
	TotalTimeSpent pulumi.IntOutput `pulumi:"totalTimeSpent"`
	// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The number of upvotes the issue has received.
	Upvotes pulumi.IntOutput `pulumi:"upvotes"`
	// The number of user notes on the issue.
	UserNotesCount pulumi.IntOutput `pulumi:"userNotesCount"`
	// The web URL of the issue.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
	// The weight of the issue. Valid values are greater than or equal to 0.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewProjectIssue registers a new resource with the given unique name, arguments, and options.
func NewProjectIssue(ctx *pulumi.Context,
	name string, args *ProjectIssueArgs, opts ...pulumi.ResourceOption) (*ProjectIssue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource ProjectIssue
	err := ctx.RegisterResource("gitlab:index/projectIssue:ProjectIssue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectIssue gets an existing ProjectIssue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectIssue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectIssueState, opts ...pulumi.ResourceOption) (*ProjectIssue, error) {
	var resource ProjectIssue
	err := ctx.ReadResource("gitlab:index/projectIssue:ProjectIssue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectIssue resources.
type projectIssueState struct {
	// The IDs of the users to assign the issue to.
	AssigneeIds []int `pulumi:"assigneeIds"`
	// The ID of the author of the issue. Use `User` data source to get more information about the user.
	AuthorId *int `pulumi:"authorId"`
	// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	ClosedAt *string `pulumi:"closedAt"`
	// The ID of the user that closed the issue. Use `User` data source to get more information about the user.
	ClosedByUserId *int `pulumi:"closedByUserId"`
	// Set an issue to be confidential.
	Confidential *bool `pulumi:"confidential"`
	// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether the issue is deleted instead of closed during destroy.
	DeleteOnDestroy *bool `pulumi:"deleteOnDestroy"`
	// The description of an issue. Limited to 1,048,576 characters.
	Description *string `pulumi:"description"`
	// Whether the issue is locked for discussions or not.
	DiscussionLocked *bool `pulumi:"discussionLocked"`
	// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
	DiscussionToResolve *string `pulumi:"discussionToResolve"`
	// The number of downvotes the issue has received.
	Downvotes *int `pulumi:"downvotes"`
	// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
	DueDate *string `pulumi:"dueDate"`
	// ID of the epic to add the issue to. Valid values are greater than or equal to 0.
	EpicId *int `pulumi:"epicId"`
	// The ID of the epic issue.
	EpicIssueId *int `pulumi:"epicIssueId"`
	// The external ID of the issue.
	ExternalId *string `pulumi:"externalId"`
	// The human-readable time estimate of the issue.
	HumanTimeEstimate *string `pulumi:"humanTimeEstimate"`
	// The human-readable total time spent of the issue.
	HumanTotalTimeSpent *string `pulumi:"humanTotalTimeSpent"`
	// The internal ID of the project's issue.
	Iid *int `pulumi:"iid"`
	// The instance-wide ID of the issue.
	IssueId *int `pulumi:"issueId"`
	// The ID of the issue link.
	IssueLinkId *int `pulumi:"issueLinkId"`
	// The type of issue. Valid values are: `issue`, `incident`, `testCase`.
	IssueType *string `pulumi:"issueType"`
	// The labels of an issue.
	Labels []string `pulumi:"labels"`
	// The links of the issue.
	Links map[string]string `pulumi:"links"`
	// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
	MergeRequestToResolveDiscussionsOf *int `pulumi:"mergeRequestToResolveDiscussionsOf"`
	// The number of merge requests associated with the issue.
	MergeRequestsCount *int `pulumi:"mergeRequestsCount"`
	// The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
	MilestoneId *int `pulumi:"milestoneId"`
	// The ID of the issue that was moved to.
	MovedToId *int `pulumi:"movedToId"`
	// The name or ID of the project.
	Project *string `pulumi:"project"`
	// The references of the issue.
	References map[string]string `pulumi:"references"`
	// The state of the issue. Valid values are: `opened`, `closed`.
	State *string `pulumi:"state"`
	// Whether the authenticated user is subscribed to the issue or not.
	Subscribed *bool `pulumi:"subscribed"`
	// The task completion status. It's always a one element list.
	TaskCompletionStatus *ProjectIssueTaskCompletionStatus `pulumi:"taskCompletionStatus"`
	// The time estimate of the issue.
	TimeEstimate *int `pulumi:"timeEstimate"`
	// The title of the issue.
	Title *string `pulumi:"title"`
	// The total time spent of the issue.
	TotalTimeSpent *int `pulumi:"totalTimeSpent"`
	// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The number of upvotes the issue has received.
	Upvotes *int `pulumi:"upvotes"`
	// The number of user notes on the issue.
	UserNotesCount *int `pulumi:"userNotesCount"`
	// The web URL of the issue.
	WebUrl *string `pulumi:"webUrl"`
	// The weight of the issue. Valid values are greater than or equal to 0.
	Weight *int `pulumi:"weight"`
}

type ProjectIssueState struct {
	// The IDs of the users to assign the issue to.
	AssigneeIds pulumi.IntArrayInput
	// The ID of the author of the issue. Use `User` data source to get more information about the user.
	AuthorId pulumi.IntPtrInput
	// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	ClosedAt pulumi.StringPtrInput
	// The ID of the user that closed the issue. Use `User` data source to get more information about the user.
	ClosedByUserId pulumi.IntPtrInput
	// Set an issue to be confidential.
	Confidential pulumi.BoolPtrInput
	// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
	CreatedAt pulumi.StringPtrInput
	// Whether the issue is deleted instead of closed during destroy.
	DeleteOnDestroy pulumi.BoolPtrInput
	// The description of an issue. Limited to 1,048,576 characters.
	Description pulumi.StringPtrInput
	// Whether the issue is locked for discussions or not.
	DiscussionLocked pulumi.BoolPtrInput
	// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
	DiscussionToResolve pulumi.StringPtrInput
	// The number of downvotes the issue has received.
	Downvotes pulumi.IntPtrInput
	// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
	DueDate pulumi.StringPtrInput
	// ID of the epic to add the issue to. Valid values are greater than or equal to 0.
	EpicId pulumi.IntPtrInput
	// The ID of the epic issue.
	EpicIssueId pulumi.IntPtrInput
	// The external ID of the issue.
	ExternalId pulumi.StringPtrInput
	// The human-readable time estimate of the issue.
	HumanTimeEstimate pulumi.StringPtrInput
	// The human-readable total time spent of the issue.
	HumanTotalTimeSpent pulumi.StringPtrInput
	// The internal ID of the project's issue.
	Iid pulumi.IntPtrInput
	// The instance-wide ID of the issue.
	IssueId pulumi.IntPtrInput
	// The ID of the issue link.
	IssueLinkId pulumi.IntPtrInput
	// The type of issue. Valid values are: `issue`, `incident`, `testCase`.
	IssueType pulumi.StringPtrInput
	// The labels of an issue.
	Labels pulumi.StringArrayInput
	// The links of the issue.
	Links pulumi.StringMapInput
	// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
	MergeRequestToResolveDiscussionsOf pulumi.IntPtrInput
	// The number of merge requests associated with the issue.
	MergeRequestsCount pulumi.IntPtrInput
	// The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
	MilestoneId pulumi.IntPtrInput
	// The ID of the issue that was moved to.
	MovedToId pulumi.IntPtrInput
	// The name or ID of the project.
	Project pulumi.StringPtrInput
	// The references of the issue.
	References pulumi.StringMapInput
	// The state of the issue. Valid values are: `opened`, `closed`.
	State pulumi.StringPtrInput
	// Whether the authenticated user is subscribed to the issue or not.
	Subscribed pulumi.BoolPtrInput
	// The task completion status. It's always a one element list.
	TaskCompletionStatus ProjectIssueTaskCompletionStatusPtrInput
	// The time estimate of the issue.
	TimeEstimate pulumi.IntPtrInput
	// The title of the issue.
	Title pulumi.StringPtrInput
	// The total time spent of the issue.
	TotalTimeSpent pulumi.IntPtrInput
	// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt pulumi.StringPtrInput
	// The number of upvotes the issue has received.
	Upvotes pulumi.IntPtrInput
	// The number of user notes on the issue.
	UserNotesCount pulumi.IntPtrInput
	// The web URL of the issue.
	WebUrl pulumi.StringPtrInput
	// The weight of the issue. Valid values are greater than or equal to 0.
	Weight pulumi.IntPtrInput
}

func (ProjectIssueState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIssueState)(nil)).Elem()
}

type projectIssueArgs struct {
	// The IDs of the users to assign the issue to.
	AssigneeIds []int `pulumi:"assigneeIds"`
	// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	ClosedAt *string `pulumi:"closedAt"`
	// The ID of the user that closed the issue. Use `User` data source to get more information about the user.
	ClosedByUserId *int `pulumi:"closedByUserId"`
	// Set an issue to be confidential.
	Confidential *bool `pulumi:"confidential"`
	// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether the issue is deleted instead of closed during destroy.
	DeleteOnDestroy *bool `pulumi:"deleteOnDestroy"`
	// The description of an issue. Limited to 1,048,576 characters.
	Description *string `pulumi:"description"`
	// Whether the issue is locked for discussions or not.
	DiscussionLocked *bool `pulumi:"discussionLocked"`
	// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
	DiscussionToResolve *string `pulumi:"discussionToResolve"`
	// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
	DueDate *string `pulumi:"dueDate"`
	// The ID of the epic issue.
	EpicIssueId *int `pulumi:"epicIssueId"`
	// The human-readable time estimate of the issue.
	HumanTimeEstimate *string `pulumi:"humanTimeEstimate"`
	// The human-readable total time spent of the issue.
	HumanTotalTimeSpent *string `pulumi:"humanTotalTimeSpent"`
	// The internal ID of the project's issue.
	Iid *int `pulumi:"iid"`
	// The type of issue. Valid values are: `issue`, `incident`, `testCase`.
	IssueType *string `pulumi:"issueType"`
	// The labels of an issue.
	Labels []string `pulumi:"labels"`
	// The links of the issue.
	Links map[string]string `pulumi:"links"`
	// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
	MergeRequestToResolveDiscussionsOf *int `pulumi:"mergeRequestToResolveDiscussionsOf"`
	// The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
	MilestoneId *int `pulumi:"milestoneId"`
	// The name or ID of the project.
	Project string `pulumi:"project"`
	// The references of the issue.
	References map[string]string `pulumi:"references"`
	// The state of the issue. Valid values are: `opened`, `closed`.
	State *string `pulumi:"state"`
	// The task completion status. It's always a one element list.
	TaskCompletionStatus *ProjectIssueTaskCompletionStatus `pulumi:"taskCompletionStatus"`
	// The time estimate of the issue.
	TimeEstimate *int `pulumi:"timeEstimate"`
	// The title of the issue.
	Title string `pulumi:"title"`
	// The total time spent of the issue.
	TotalTimeSpent *int `pulumi:"totalTimeSpent"`
	// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The weight of the issue. Valid values are greater than or equal to 0.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a ProjectIssue resource.
type ProjectIssueArgs struct {
	// The IDs of the users to assign the issue to.
	AssigneeIds pulumi.IntArrayInput
	// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	ClosedAt pulumi.StringPtrInput
	// The ID of the user that closed the issue. Use `User` data source to get more information about the user.
	ClosedByUserId pulumi.IntPtrInput
	// Set an issue to be confidential.
	Confidential pulumi.BoolPtrInput
	// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
	CreatedAt pulumi.StringPtrInput
	// Whether the issue is deleted instead of closed during destroy.
	DeleteOnDestroy pulumi.BoolPtrInput
	// The description of an issue. Limited to 1,048,576 characters.
	Description pulumi.StringPtrInput
	// Whether the issue is locked for discussions or not.
	DiscussionLocked pulumi.BoolPtrInput
	// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
	DiscussionToResolve pulumi.StringPtrInput
	// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
	DueDate pulumi.StringPtrInput
	// The ID of the epic issue.
	EpicIssueId pulumi.IntPtrInput
	// The human-readable time estimate of the issue.
	HumanTimeEstimate pulumi.StringPtrInput
	// The human-readable total time spent of the issue.
	HumanTotalTimeSpent pulumi.StringPtrInput
	// The internal ID of the project's issue.
	Iid pulumi.IntPtrInput
	// The type of issue. Valid values are: `issue`, `incident`, `testCase`.
	IssueType pulumi.StringPtrInput
	// The labels of an issue.
	Labels pulumi.StringArrayInput
	// The links of the issue.
	Links pulumi.StringMapInput
	// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
	MergeRequestToResolveDiscussionsOf pulumi.IntPtrInput
	// The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
	MilestoneId pulumi.IntPtrInput
	// The name or ID of the project.
	Project pulumi.StringInput
	// The references of the issue.
	References pulumi.StringMapInput
	// The state of the issue. Valid values are: `opened`, `closed`.
	State pulumi.StringPtrInput
	// The task completion status. It's always a one element list.
	TaskCompletionStatus ProjectIssueTaskCompletionStatusPtrInput
	// The time estimate of the issue.
	TimeEstimate pulumi.IntPtrInput
	// The title of the issue.
	Title pulumi.StringInput
	// The total time spent of the issue.
	TotalTimeSpent pulumi.IntPtrInput
	// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt pulumi.StringPtrInput
	// The weight of the issue. Valid values are greater than or equal to 0.
	Weight pulumi.IntPtrInput
}

func (ProjectIssueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIssueArgs)(nil)).Elem()
}

type ProjectIssueInput interface {
	pulumi.Input

	ToProjectIssueOutput() ProjectIssueOutput
	ToProjectIssueOutputWithContext(ctx context.Context) ProjectIssueOutput
}

func (*ProjectIssue) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIssue)(nil)).Elem()
}

func (i *ProjectIssue) ToProjectIssueOutput() ProjectIssueOutput {
	return i.ToProjectIssueOutputWithContext(context.Background())
}

func (i *ProjectIssue) ToProjectIssueOutputWithContext(ctx context.Context) ProjectIssueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIssueOutput)
}

// ProjectIssueArrayInput is an input type that accepts ProjectIssueArray and ProjectIssueArrayOutput values.
// You can construct a concrete instance of `ProjectIssueArrayInput` via:
//
//          ProjectIssueArray{ ProjectIssueArgs{...} }
type ProjectIssueArrayInput interface {
	pulumi.Input

	ToProjectIssueArrayOutput() ProjectIssueArrayOutput
	ToProjectIssueArrayOutputWithContext(context.Context) ProjectIssueArrayOutput
}

type ProjectIssueArray []ProjectIssueInput

func (ProjectIssueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIssue)(nil)).Elem()
}

func (i ProjectIssueArray) ToProjectIssueArrayOutput() ProjectIssueArrayOutput {
	return i.ToProjectIssueArrayOutputWithContext(context.Background())
}

func (i ProjectIssueArray) ToProjectIssueArrayOutputWithContext(ctx context.Context) ProjectIssueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIssueArrayOutput)
}

// ProjectIssueMapInput is an input type that accepts ProjectIssueMap and ProjectIssueMapOutput values.
// You can construct a concrete instance of `ProjectIssueMapInput` via:
//
//          ProjectIssueMap{ "key": ProjectIssueArgs{...} }
type ProjectIssueMapInput interface {
	pulumi.Input

	ToProjectIssueMapOutput() ProjectIssueMapOutput
	ToProjectIssueMapOutputWithContext(context.Context) ProjectIssueMapOutput
}

type ProjectIssueMap map[string]ProjectIssueInput

func (ProjectIssueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIssue)(nil)).Elem()
}

func (i ProjectIssueMap) ToProjectIssueMapOutput() ProjectIssueMapOutput {
	return i.ToProjectIssueMapOutputWithContext(context.Background())
}

func (i ProjectIssueMap) ToProjectIssueMapOutputWithContext(ctx context.Context) ProjectIssueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIssueMapOutput)
}

type ProjectIssueOutput struct{ *pulumi.OutputState }

func (ProjectIssueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIssue)(nil)).Elem()
}

func (o ProjectIssueOutput) ToProjectIssueOutput() ProjectIssueOutput {
	return o
}

func (o ProjectIssueOutput) ToProjectIssueOutputWithContext(ctx context.Context) ProjectIssueOutput {
	return o
}

type ProjectIssueArrayOutput struct{ *pulumi.OutputState }

func (ProjectIssueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIssue)(nil)).Elem()
}

func (o ProjectIssueArrayOutput) ToProjectIssueArrayOutput() ProjectIssueArrayOutput {
	return o
}

func (o ProjectIssueArrayOutput) ToProjectIssueArrayOutputWithContext(ctx context.Context) ProjectIssueArrayOutput {
	return o
}

func (o ProjectIssueArrayOutput) Index(i pulumi.IntInput) ProjectIssueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectIssue {
		return vs[0].([]*ProjectIssue)[vs[1].(int)]
	}).(ProjectIssueOutput)
}

type ProjectIssueMapOutput struct{ *pulumi.OutputState }

func (ProjectIssueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIssue)(nil)).Elem()
}

func (o ProjectIssueMapOutput) ToProjectIssueMapOutput() ProjectIssueMapOutput {
	return o
}

func (o ProjectIssueMapOutput) ToProjectIssueMapOutputWithContext(ctx context.Context) ProjectIssueMapOutput {
	return o
}

func (o ProjectIssueMapOutput) MapIndex(k pulumi.StringInput) ProjectIssueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectIssue {
		return vs[0].(map[string]*ProjectIssue)[vs[1].(string)]
	}).(ProjectIssueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIssueInput)(nil)).Elem(), &ProjectIssue{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIssueArrayInput)(nil)).Elem(), ProjectIssueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIssueMapInput)(nil)).Elem(), ProjectIssueMap{})
	pulumi.RegisterOutputType(ProjectIssueOutput{})
	pulumi.RegisterOutputType(ProjectIssueArrayOutput{})
	pulumi.RegisterOutputType(ProjectIssueMapOutput{})
}
