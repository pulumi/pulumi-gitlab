// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new GitLab.Project("foo", new()
    ///     {
    ///         Description = "Lorem Ipsum",
    ///         VisibilityLevel = "public",
    ///     });
    /// 
    ///     var welcomeIssue = new GitLab.ProjectIssue("welcomeIssue", new()
    ///     {
    ///         Project = foo.Id,
    ///         Title = "Welcome!",
    ///         Description = foo.Name.Apply(name =&gt; @$"  Welcome to the {name} project!
    /// 
    /// "),
    ///         DiscussionLocked = true,
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["welcomeIssueWebUrl"] = data.Gitlab_project_issue.Web_url,
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import this resource with an id made up of `{project-id}:{issue-id}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectIssue:ProjectIssue welcome_issue 42:1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectIssue:ProjectIssue")]
    public partial class ProjectIssue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IDs of the users to assign the issue to.
        /// </summary>
        [Output("assigneeIds")]
        public Output<ImmutableArray<int>> AssigneeIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
        /// </summary>
        [Output("authorId")]
        public Output<int> AuthorId { get; private set; } = null!;

        /// <summary>
        /// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        [Output("closedAt")]
        public Output<string> ClosedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
        /// </summary>
        [Output("closedByUserId")]
        public Output<int> ClosedByUserId { get; private set; } = null!;

        /// <summary>
        /// Set an issue to be confidential.
        /// </summary>
        [Output("confidential")]
        public Output<bool?> Confidential { get; private set; } = null!;

        /// <summary>
        /// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Whether the issue is deleted instead of closed during destroy.
        /// </summary>
        [Output("deleteOnDestroy")]
        public Output<bool?> DeleteOnDestroy { get; private set; } = null!;

        /// <summary>
        /// The description of an issue. Limited to 1,048,576 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the issue is locked for discussions or not.
        /// </summary>
        [Output("discussionLocked")]
        public Output<bool?> DiscussionLocked { get; private set; } = null!;

        /// <summary>
        /// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
        /// </summary>
        [Output("discussionToResolve")]
        public Output<string?> DiscussionToResolve { get; private set; } = null!;

        /// <summary>
        /// The number of downvotes the issue has received.
        /// </summary>
        [Output("downvotes")]
        public Output<int> Downvotes { get; private set; } = null!;

        /// <summary>
        /// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        /// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
        /// </summary>
        [Output("dueDate")]
        public Output<string?> DueDate { get; private set; } = null!;

        /// <summary>
        /// ID of the epic to add the issue to. Valid values are greater than or equal to 0.
        /// </summary>
        [Output("epicId")]
        public Output<int> EpicId { get; private set; } = null!;

        /// <summary>
        /// The ID of the epic issue.
        /// </summary>
        [Output("epicIssueId")]
        public Output<int> EpicIssueId { get; private set; } = null!;

        /// <summary>
        /// The external ID of the issue.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// The human-readable time estimate of the issue.
        /// </summary>
        [Output("humanTimeEstimate")]
        public Output<string> HumanTimeEstimate { get; private set; } = null!;

        /// <summary>
        /// The human-readable total time spent of the issue.
        /// </summary>
        [Output("humanTotalTimeSpent")]
        public Output<string> HumanTotalTimeSpent { get; private set; } = null!;

        /// <summary>
        /// The internal ID of the project's issue.
        /// </summary>
        [Output("iid")]
        public Output<int> Iid { get; private set; } = null!;

        /// <summary>
        /// The instance-wide ID of the issue.
        /// </summary>
        [Output("issueId")]
        public Output<int> IssueId { get; private set; } = null!;

        /// <summary>
        /// The ID of the issue link.
        /// </summary>
        [Output("issueLinkId")]
        public Output<int> IssueLinkId { get; private set; } = null!;

        /// <summary>
        /// The type of issue. Valid values are: `issue`, `incident`, `test_case`.
        /// </summary>
        [Output("issueType")]
        public Output<string?> IssueType { get; private set; } = null!;

        /// <summary>
        /// The labels of an issue.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The links of the issue.
        /// </summary>
        [Output("links")]
        public Output<ImmutableDictionary<string, string>> Links { get; private set; } = null!;

        /// <summary>
        /// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
        /// </summary>
        [Output("mergeRequestToResolveDiscussionsOf")]
        public Output<int?> MergeRequestToResolveDiscussionsOf { get; private set; } = null!;

        /// <summary>
        /// The number of merge requests associated with the issue.
        /// </summary>
        [Output("mergeRequestsCount")]
        public Output<int> MergeRequestsCount { get; private set; } = null!;

        /// <summary>
        /// The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
        /// </summary>
        [Output("milestoneId")]
        public Output<int?> MilestoneId { get; private set; } = null!;

        /// <summary>
        /// The ID of the issue that was moved to.
        /// </summary>
        [Output("movedToId")]
        public Output<int> MovedToId { get; private set; } = null!;

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The references of the issue.
        /// </summary>
        [Output("references")]
        public Output<ImmutableDictionary<string, string>> References { get; private set; } = null!;

        /// <summary>
        /// The state of the issue. Valid values are: `opened`, `closed`.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Whether the authenticated user is subscribed to the issue or not.
        /// </summary>
        [Output("subscribed")]
        public Output<bool> Subscribed { get; private set; } = null!;

        /// <summary>
        /// The task completion status. It's always a one element list.
        /// </summary>
        [Output("taskCompletionStatuses")]
        public Output<ImmutableArray<Outputs.ProjectIssueTaskCompletionStatus>> TaskCompletionStatuses { get; private set; } = null!;

        /// <summary>
        /// The time estimate of the issue.
        /// </summary>
        [Output("timeEstimate")]
        public Output<int> TimeEstimate { get; private set; } = null!;

        /// <summary>
        /// The title of the issue.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// The total time spent of the issue.
        /// </summary>
        [Output("totalTimeSpent")]
        public Output<int> TotalTimeSpent { get; private set; } = null!;

        /// <summary>
        /// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The number of upvotes the issue has received.
        /// </summary>
        [Output("upvotes")]
        public Output<int> Upvotes { get; private set; } = null!;

        /// <summary>
        /// The number of user notes on the issue.
        /// </summary>
        [Output("userNotesCount")]
        public Output<int> UserNotesCount { get; private set; } = null!;

        /// <summary>
        /// The web URL of the issue.
        /// </summary>
        [Output("webUrl")]
        public Output<string> WebUrl { get; private set; } = null!;

        /// <summary>
        /// The weight of the issue. Valid values are greater than or equal to 0.
        /// </summary>
        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectIssue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectIssue(string name, ProjectIssueArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectIssue:ProjectIssue", name, args ?? new ProjectIssueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectIssue(string name, Input<string> id, ProjectIssueState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectIssue:ProjectIssue", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProjectIssue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectIssue Get(string name, Input<string> id, ProjectIssueState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectIssue(name, id, state, options);
        }
    }

    public sealed class ProjectIssueArgs : global::Pulumi.ResourceArgs
    {
        [Input("assigneeIds")]
        private InputList<int>? _assigneeIds;

        /// <summary>
        /// The IDs of the users to assign the issue to.
        /// </summary>
        public InputList<int> AssigneeIds
        {
            get => _assigneeIds ?? (_assigneeIds = new InputList<int>());
            set => _assigneeIds = value;
        }

        /// <summary>
        /// Set an issue to be confidential.
        /// </summary>
        [Input("confidential")]
        public Input<bool>? Confidential { get; set; }

        /// <summary>
        /// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether the issue is deleted instead of closed during destroy.
        /// </summary>
        [Input("deleteOnDestroy")]
        public Input<bool>? DeleteOnDestroy { get; set; }

        /// <summary>
        /// The description of an issue. Limited to 1,048,576 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the issue is locked for discussions or not.
        /// </summary>
        [Input("discussionLocked")]
        public Input<bool>? DiscussionLocked { get; set; }

        /// <summary>
        /// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
        /// </summary>
        [Input("discussionToResolve")]
        public Input<string>? DiscussionToResolve { get; set; }

        /// <summary>
        /// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        /// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
        /// </summary>
        [Input("dueDate")]
        public Input<string>? DueDate { get; set; }

        /// <summary>
        /// The ID of the epic issue.
        /// </summary>
        [Input("epicIssueId")]
        public Input<int>? EpicIssueId { get; set; }

        /// <summary>
        /// The internal ID of the project's issue.
        /// </summary>
        [Input("iid")]
        public Input<int>? Iid { get; set; }

        /// <summary>
        /// The type of issue. Valid values are: `issue`, `incident`, `test_case`.
        /// </summary>
        [Input("issueType")]
        public Input<string>? IssueType { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// The labels of an issue.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
        /// </summary>
        [Input("mergeRequestToResolveDiscussionsOf")]
        public Input<int>? MergeRequestToResolveDiscussionsOf { get; set; }

        /// <summary>
        /// The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
        /// </summary>
        [Input("milestoneId")]
        public Input<int>? MilestoneId { get; set; }

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The state of the issue. Valid values are: `opened`, `closed`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The title of the issue.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The weight of the issue. Valid values are greater than or equal to 0.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ProjectIssueArgs()
        {
        }
        public static new ProjectIssueArgs Empty => new ProjectIssueArgs();
    }

    public sealed class ProjectIssueState : global::Pulumi.ResourceArgs
    {
        [Input("assigneeIds")]
        private InputList<int>? _assigneeIds;

        /// <summary>
        /// The IDs of the users to assign the issue to.
        /// </summary>
        public InputList<int> AssigneeIds
        {
            get => _assigneeIds ?? (_assigneeIds = new InputList<int>());
            set => _assigneeIds = value;
        }

        /// <summary>
        /// The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
        /// </summary>
        [Input("authorId")]
        public Input<int>? AuthorId { get; set; }

        /// <summary>
        /// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        [Input("closedAt")]
        public Input<string>? ClosedAt { get; set; }

        /// <summary>
        /// The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
        /// </summary>
        [Input("closedByUserId")]
        public Input<int>? ClosedByUserId { get; set; }

        /// <summary>
        /// Set an issue to be confidential.
        /// </summary>
        [Input("confidential")]
        public Input<bool>? Confidential { get; set; }

        /// <summary>
        /// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether the issue is deleted instead of closed during destroy.
        /// </summary>
        [Input("deleteOnDestroy")]
        public Input<bool>? DeleteOnDestroy { get; set; }

        /// <summary>
        /// The description of an issue. Limited to 1,048,576 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the issue is locked for discussions or not.
        /// </summary>
        [Input("discussionLocked")]
        public Input<bool>? DiscussionLocked { get; set; }

        /// <summary>
        /// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
        /// </summary>
        [Input("discussionToResolve")]
        public Input<string>? DiscussionToResolve { get; set; }

        /// <summary>
        /// The number of downvotes the issue has received.
        /// </summary>
        [Input("downvotes")]
        public Input<int>? Downvotes { get; set; }

        /// <summary>
        /// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        /// **Note:** removing a due date is currently not supported, see https://github.com/xanzy/go-gitlab/issues/1384 for details.
        /// </summary>
        [Input("dueDate")]
        public Input<string>? DueDate { get; set; }

        /// <summary>
        /// ID of the epic to add the issue to. Valid values are greater than or equal to 0.
        /// </summary>
        [Input("epicId")]
        public Input<int>? EpicId { get; set; }

        /// <summary>
        /// The ID of the epic issue.
        /// </summary>
        [Input("epicIssueId")]
        public Input<int>? EpicIssueId { get; set; }

        /// <summary>
        /// The external ID of the issue.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The human-readable time estimate of the issue.
        /// </summary>
        [Input("humanTimeEstimate")]
        public Input<string>? HumanTimeEstimate { get; set; }

        /// <summary>
        /// The human-readable total time spent of the issue.
        /// </summary>
        [Input("humanTotalTimeSpent")]
        public Input<string>? HumanTotalTimeSpent { get; set; }

        /// <summary>
        /// The internal ID of the project's issue.
        /// </summary>
        [Input("iid")]
        public Input<int>? Iid { get; set; }

        /// <summary>
        /// The instance-wide ID of the issue.
        /// </summary>
        [Input("issueId")]
        public Input<int>? IssueId { get; set; }

        /// <summary>
        /// The ID of the issue link.
        /// </summary>
        [Input("issueLinkId")]
        public Input<int>? IssueLinkId { get; set; }

        /// <summary>
        /// The type of issue. Valid values are: `issue`, `incident`, `test_case`.
        /// </summary>
        [Input("issueType")]
        public Input<string>? IssueType { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// The labels of an issue.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("links")]
        private InputMap<string>? _links;

        /// <summary>
        /// The links of the issue.
        /// </summary>
        public InputMap<string> Links
        {
            get => _links ?? (_links = new InputMap<string>());
            set => _links = value;
        }

        /// <summary>
        /// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
        /// </summary>
        [Input("mergeRequestToResolveDiscussionsOf")]
        public Input<int>? MergeRequestToResolveDiscussionsOf { get; set; }

        /// <summary>
        /// The number of merge requests associated with the issue.
        /// </summary>
        [Input("mergeRequestsCount")]
        public Input<int>? MergeRequestsCount { get; set; }

        /// <summary>
        /// The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
        /// </summary>
        [Input("milestoneId")]
        public Input<int>? MilestoneId { get; set; }

        /// <summary>
        /// The ID of the issue that was moved to.
        /// </summary>
        [Input("movedToId")]
        public Input<int>? MovedToId { get; set; }

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("references")]
        private InputMap<string>? _references;

        /// <summary>
        /// The references of the issue.
        /// </summary>
        public InputMap<string> References
        {
            get => _references ?? (_references = new InputMap<string>());
            set => _references = value;
        }

        /// <summary>
        /// The state of the issue. Valid values are: `opened`, `closed`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Whether the authenticated user is subscribed to the issue or not.
        /// </summary>
        [Input("subscribed")]
        public Input<bool>? Subscribed { get; set; }

        [Input("taskCompletionStatuses")]
        private InputList<Inputs.ProjectIssueTaskCompletionStatusGetArgs>? _taskCompletionStatuses;

        /// <summary>
        /// The task completion status. It's always a one element list.
        /// </summary>
        public InputList<Inputs.ProjectIssueTaskCompletionStatusGetArgs> TaskCompletionStatuses
        {
            get => _taskCompletionStatuses ?? (_taskCompletionStatuses = new InputList<Inputs.ProjectIssueTaskCompletionStatusGetArgs>());
            set => _taskCompletionStatuses = value;
        }

        /// <summary>
        /// The time estimate of the issue.
        /// </summary>
        [Input("timeEstimate")]
        public Input<int>? TimeEstimate { get; set; }

        /// <summary>
        /// The title of the issue.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The total time spent of the issue.
        /// </summary>
        [Input("totalTimeSpent")]
        public Input<int>? TotalTimeSpent { get; set; }

        /// <summary>
        /// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The number of upvotes the issue has received.
        /// </summary>
        [Input("upvotes")]
        public Input<int>? Upvotes { get; set; }

        /// <summary>
        /// The number of user notes on the issue.
        /// </summary>
        [Input("userNotesCount")]
        public Input<int>? UserNotesCount { get; set; }

        /// <summary>
        /// The web URL of the issue.
        /// </summary>
        [Input("webUrl")]
        public Input<string>? WebUrl { get; set; }

        /// <summary>
        /// The weight of the issue. Valid values are greater than or equal to 0.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ProjectIssueState()
        {
        }
        public static new ProjectIssueState Empty => new ProjectIssueState();
    }
}
