// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectIssue
    {
        /// <summary>
        /// The `gitlab.ProjectIssue` data source allows to retrieve details about an issue in a project.
        /// 
        /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)
        /// </summary>
        public static Task<GetProjectIssueResult> InvokeAsync(GetProjectIssueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectIssueResult>("gitlab:index/getProjectIssue:getProjectIssue", args ?? new GetProjectIssueArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.ProjectIssue` data source allows to retrieve details about an issue in a project.
        /// 
        /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)
        /// </summary>
        public static Output<GetProjectIssueResult> Invoke(GetProjectIssueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectIssueResult>("gitlab:index/getProjectIssue:getProjectIssue", args ?? new GetProjectIssueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectIssueArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The internal ID of the project's issue.
        /// </summary>
        [Input("iid", required: true)]
        public int Iid { get; set; }

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetProjectIssueArgs()
        {
        }
        public static new GetProjectIssueArgs Empty => new GetProjectIssueArgs();
    }

    public sealed class GetProjectIssueInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The internal ID of the project's issue.
        /// </summary>
        [Input("iid", required: true)]
        public Input<int> Iid { get; set; } = null!;

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetProjectIssueInvokeArgs()
        {
        }
        public static new GetProjectIssueInvokeArgs Empty => new GetProjectIssueInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectIssueResult
    {
        /// <summary>
        /// The IDs of the users to assign the issue to.
        /// </summary>
        public readonly ImmutableArray<int> AssigneeIds;
        /// <summary>
        /// The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
        /// </summary>
        public readonly int AuthorId;
        /// <summary>
        /// When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        public readonly string ClosedAt;
        /// <summary>
        /// The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
        /// </summary>
        public readonly int ClosedByUserId;
        /// <summary>
        /// Set an issue to be confidential.
        /// </summary>
        public readonly bool Confidential;
        /// <summary>
        /// When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of an issue. Limited to 1,048,576 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the issue is locked for discussions or not.
        /// </summary>
        public readonly bool DiscussionLocked;
        /// <summary>
        /// The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
        /// </summary>
        public readonly string DiscussionToResolve;
        /// <summary>
        /// The number of downvotes the issue has received.
        /// </summary>
        public readonly int Downvotes;
        /// <summary>
        /// The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        /// </summary>
        public readonly string DueDate;
        /// <summary>
        /// ID of the epic to add the issue to. Valid values are greater than or equal to 0.
        /// </summary>
        public readonly int EpicId;
        /// <summary>
        /// The ID of the epic issue.
        /// </summary>
        public readonly int EpicIssueId;
        /// <summary>
        /// The external ID of the issue.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The human-readable time estimate of the issue.
        /// </summary>
        public readonly string HumanTimeEstimate;
        /// <summary>
        /// The human-readable total time spent of the issue.
        /// </summary>
        public readonly string HumanTotalTimeSpent;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The internal ID of the project's issue.
        /// </summary>
        public readonly int Iid;
        /// <summary>
        /// The instance-wide ID of the issue.
        /// </summary>
        public readonly int IssueId;
        /// <summary>
        /// The ID of the issue link.
        /// </summary>
        public readonly int IssueLinkId;
        /// <summary>
        /// The type of issue. Valid values are: `issue`, `incident`, `test_case`.
        /// </summary>
        public readonly string IssueType;
        /// <summary>
        /// The labels of an issue.
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// The links of the issue.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Links;
        /// <summary>
        /// The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
        /// </summary>
        public readonly int MergeRequestToResolveDiscussionsOf;
        /// <summary>
        /// The number of merge requests associated with the issue.
        /// </summary>
        public readonly int MergeRequestsCount;
        /// <summary>
        /// The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
        /// </summary>
        public readonly int MilestoneId;
        /// <summary>
        /// The ID of the issue that was moved to.
        /// </summary>
        public readonly int MovedToId;
        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The references of the issue.
        /// </summary>
        public readonly ImmutableDictionary<string, string> References;
        /// <summary>
        /// The state of the issue. Valid values are: `opened`, `closed`.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Whether the authenticated user is subscribed to the issue or not.
        /// </summary>
        public readonly bool Subscribed;
        /// <summary>
        /// The task completion status. It's always a one element list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectIssueTaskCompletionStatusResult> TaskCompletionStatuses;
        /// <summary>
        /// The time estimate of the issue.
        /// </summary>
        public readonly int TimeEstimate;
        /// <summary>
        /// The title of the issue.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The total time spent of the issue.
        /// </summary>
        public readonly int TotalTimeSpent;
        /// <summary>
        /// When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The number of upvotes the issue has received.
        /// </summary>
        public readonly int Upvotes;
        /// <summary>
        /// The number of user notes on the issue.
        /// </summary>
        public readonly int UserNotesCount;
        /// <summary>
        /// The web URL of the issue.
        /// </summary>
        public readonly string WebUrl;
        /// <summary>
        /// The weight of the issue. Valid values are greater than or equal to 0.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private GetProjectIssueResult(
            ImmutableArray<int> assigneeIds,

            int authorId,

            string closedAt,

            int closedByUserId,

            bool confidential,

            string createdAt,

            string description,

            bool discussionLocked,

            string discussionToResolve,

            int downvotes,

            string dueDate,

            int epicId,

            int epicIssueId,

            string externalId,

            string humanTimeEstimate,

            string humanTotalTimeSpent,

            string id,

            int iid,

            int issueId,

            int issueLinkId,

            string issueType,

            ImmutableArray<string> labels,

            ImmutableDictionary<string, string> links,

            int mergeRequestToResolveDiscussionsOf,

            int mergeRequestsCount,

            int milestoneId,

            int movedToId,

            string project,

            ImmutableDictionary<string, string> references,

            string state,

            bool subscribed,

            ImmutableArray<Outputs.GetProjectIssueTaskCompletionStatusResult> taskCompletionStatuses,

            int timeEstimate,

            string title,

            int totalTimeSpent,

            string updatedAt,

            int upvotes,

            int userNotesCount,

            string webUrl,

            int weight)
        {
            AssigneeIds = assigneeIds;
            AuthorId = authorId;
            ClosedAt = closedAt;
            ClosedByUserId = closedByUserId;
            Confidential = confidential;
            CreatedAt = createdAt;
            Description = description;
            DiscussionLocked = discussionLocked;
            DiscussionToResolve = discussionToResolve;
            Downvotes = downvotes;
            DueDate = dueDate;
            EpicId = epicId;
            EpicIssueId = epicIssueId;
            ExternalId = externalId;
            HumanTimeEstimate = humanTimeEstimate;
            HumanTotalTimeSpent = humanTotalTimeSpent;
            Id = id;
            Iid = iid;
            IssueId = issueId;
            IssueLinkId = issueLinkId;
            IssueType = issueType;
            Labels = labels;
            Links = links;
            MergeRequestToResolveDiscussionsOf = mergeRequestToResolveDiscussionsOf;
            MergeRequestsCount = mergeRequestsCount;
            MilestoneId = milestoneId;
            MovedToId = movedToId;
            Project = project;
            References = references;
            State = state;
            Subscribed = subscribed;
            TaskCompletionStatuses = taskCompletionStatuses;
            TimeEstimate = timeEstimate;
            Title = title;
            TotalTimeSpent = totalTimeSpent;
            UpdatedAt = updatedAt;
            Upvotes = upvotes;
            UserNotesCount = userNotesCount;
            WebUrl = webUrl;
            Weight = weight;
        }
    }
}
