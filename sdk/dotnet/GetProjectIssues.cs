// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectIssues
    {
        /// <summary>
        /// The `gitlab.getProjectIssues` data source allows to retrieve details about issues in a project.
        /// 
        /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/issues/)
        /// 
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
        ///     var foo = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var allWithFoo = GitLab.GetProjectIssues.Invoke(new()
        ///     {
        ///         Project = foo.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Search = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProjectIssuesResult> InvokeAsync(GetProjectIssuesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectIssuesResult>("gitlab:index/getProjectIssues:getProjectIssues", args ?? new GetProjectIssuesArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectIssues` data source allows to retrieve details about issues in a project.
        /// 
        /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/issues/)
        /// 
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
        ///     var foo = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var allWithFoo = GitLab.GetProjectIssues.Invoke(new()
        ///     {
        ///         Project = foo.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Search = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProjectIssuesResult> Invoke(GetProjectIssuesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectIssuesResult>("gitlab:index/getProjectIssues:getProjectIssues", args ?? new GetProjectIssuesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectIssues` data source allows to retrieve details about issues in a project.
        /// 
        /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/issues/)
        /// 
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
        ///     var foo = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var allWithFoo = GitLab.GetProjectIssues.Invoke(new()
        ///     {
        ///         Project = foo.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Search = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProjectIssuesResult> Invoke(GetProjectIssuesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectIssuesResult>("gitlab:index/getProjectIssues:getProjectIssues", args ?? new GetProjectIssuesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectIssuesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
        /// </summary>
        [Input("assigneeId")]
        public int? AssigneeId { get; set; }

        /// <summary>
        /// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
        /// </summary>
        [Input("assigneeUsername")]
        public string? AssigneeUsername { get; set; }

        /// <summary>
        /// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
        /// </summary>
        [Input("authorId")]
        public int? AuthorId { get; set; }

        /// <summary>
        /// Filter confidential or public issues.
        /// </summary>
        [Input("confidential")]
        public bool? Confidential { get; set; }

        /// <summary>
        /// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("createdAfter")]
        public string? CreatedAfter { get; set; }

        /// <summary>
        /// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("createdBefore")]
        public string? CreatedBefore { get; set; }

        /// <summary>
        /// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
        /// </summary>
        [Input("dueDate")]
        public string? DueDate { get; set; }

        [Input("iids")]
        private List<int>? _iids;

        /// <summary>
        /// Return only the issues having the given iid
        /// </summary>
        public List<int> Iids
        {
            get => _iids ?? (_iids = new List<int>());
            set => _iids = value;
        }

        /// <summary>
        /// Filter to a given type of issue. Valid values are [issue incident test_case].
        /// </summary>
        [Input("issueType")]
        public string? IssueType { get; set; }

        [Input("labels")]
        private List<string>? _labels;

        /// <summary>
        /// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
        /// </summary>
        public List<string> Labels
        {
            get => _labels ?? (_labels = new List<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
        /// </summary>
        [Input("milestone")]
        public string? Milestone { get; set; }

        /// <summary>
        /// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
        /// </summary>
        [Input("myReactionEmoji")]
        public string? MyReactionEmoji { get; set; }

        [Input("notAssigneeIds")]
        private List<int>? _notAssigneeIds;

        /// <summary>
        /// Return issues that do not match the assignee id.
        /// </summary>
        public List<int> NotAssigneeIds
        {
            get => _notAssigneeIds ?? (_notAssigneeIds = new List<int>());
            set => _notAssigneeIds = value;
        }

        [Input("notAuthorIds")]
        private List<int>? _notAuthorIds;

        /// <summary>
        /// Return issues that do not match the author id.
        /// </summary>
        public List<int> NotAuthorIds
        {
            get => _notAuthorIds ?? (_notAuthorIds = new List<int>());
            set => _notAuthorIds = value;
        }

        [Input("notLabels")]
        private List<string>? _notLabels;

        /// <summary>
        /// Return issues that do not match the labels.
        /// </summary>
        public List<string> NotLabels
        {
            get => _notLabels ?? (_notLabels = new List<string>());
            set => _notLabels = value;
        }

        /// <summary>
        /// Return issues that do not match the milestone.
        /// </summary>
        [Input("notMilestone")]
        public string? NotMilestone { get; set; }

        [Input("notMyReactionEmojis")]
        private List<string>? _notMyReactionEmojis;

        /// <summary>
        /// Return issues not reacted by the authenticated user by the given emoji.
        /// </summary>
        public List<string> NotMyReactionEmojis
        {
            get => _notMyReactionEmojis ?? (_notMyReactionEmojis = new List<string>());
            set => _notMyReactionEmojis = value;
        }

        /// <summary>
        /// Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// The name or id of the project.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
        /// </summary>
        [Input("scope")]
        public string? Scope { get; set; }

        /// <summary>
        /// Search project issues against their title and description
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// Return issues sorted in asc or desc order. Default is desc
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        /// <summary>
        /// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("updatedAfter")]
        public string? UpdatedAfter { get; set; }

        /// <summary>
        /// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("updatedBefore")]
        public string? UpdatedBefore { get; set; }

        /// <summary>
        /// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
        /// </summary>
        [Input("weight")]
        public int? Weight { get; set; }

        /// <summary>
        /// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
        /// </summary>
        [Input("withLabelsDetails")]
        public bool? WithLabelsDetails { get; set; }

        public GetProjectIssuesArgs()
        {
        }
        public static new GetProjectIssuesArgs Empty => new GetProjectIssuesArgs();
    }

    public sealed class GetProjectIssuesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
        /// </summary>
        [Input("assigneeId")]
        public Input<int>? AssigneeId { get; set; }

        /// <summary>
        /// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
        /// </summary>
        [Input("assigneeUsername")]
        public Input<string>? AssigneeUsername { get; set; }

        /// <summary>
        /// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
        /// </summary>
        [Input("authorId")]
        public Input<int>? AuthorId { get; set; }

        /// <summary>
        /// Filter confidential or public issues.
        /// </summary>
        [Input("confidential")]
        public Input<bool>? Confidential { get; set; }

        /// <summary>
        /// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("createdAfter")]
        public Input<string>? CreatedAfter { get; set; }

        /// <summary>
        /// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("createdBefore")]
        public Input<string>? CreatedBefore { get; set; }

        /// <summary>
        /// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
        /// </summary>
        [Input("dueDate")]
        public Input<string>? DueDate { get; set; }

        [Input("iids")]
        private InputList<int>? _iids;

        /// <summary>
        /// Return only the issues having the given iid
        /// </summary>
        public InputList<int> Iids
        {
            get => _iids ?? (_iids = new InputList<int>());
            set => _iids = value;
        }

        /// <summary>
        /// Filter to a given type of issue. Valid values are [issue incident test_case].
        /// </summary>
        [Input("issueType")]
        public Input<string>? IssueType { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
        /// </summary>
        [Input("milestone")]
        public Input<string>? Milestone { get; set; }

        /// <summary>
        /// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
        /// </summary>
        [Input("myReactionEmoji")]
        public Input<string>? MyReactionEmoji { get; set; }

        [Input("notAssigneeIds")]
        private InputList<int>? _notAssigneeIds;

        /// <summary>
        /// Return issues that do not match the assignee id.
        /// </summary>
        public InputList<int> NotAssigneeIds
        {
            get => _notAssigneeIds ?? (_notAssigneeIds = new InputList<int>());
            set => _notAssigneeIds = value;
        }

        [Input("notAuthorIds")]
        private InputList<int>? _notAuthorIds;

        /// <summary>
        /// Return issues that do not match the author id.
        /// </summary>
        public InputList<int> NotAuthorIds
        {
            get => _notAuthorIds ?? (_notAuthorIds = new InputList<int>());
            set => _notAuthorIds = value;
        }

        [Input("notLabels")]
        private InputList<string>? _notLabels;

        /// <summary>
        /// Return issues that do not match the labels.
        /// </summary>
        public InputList<string> NotLabels
        {
            get => _notLabels ?? (_notLabels = new InputList<string>());
            set => _notLabels = value;
        }

        /// <summary>
        /// Return issues that do not match the milestone.
        /// </summary>
        [Input("notMilestone")]
        public Input<string>? NotMilestone { get; set; }

        [Input("notMyReactionEmojis")]
        private InputList<string>? _notMyReactionEmojis;

        /// <summary>
        /// Return issues not reacted by the authenticated user by the given emoji.
        /// </summary>
        public InputList<string> NotMyReactionEmojis
        {
            get => _notMyReactionEmojis ?? (_notMyReactionEmojis = new InputList<string>());
            set => _notMyReactionEmojis = value;
        }

        /// <summary>
        /// Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// The name or id of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Search project issues against their title and description
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Return issues sorted in asc or desc order. Default is desc
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        /// <summary>
        /// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("updatedAfter")]
        public Input<string>? UpdatedAfter { get; set; }

        /// <summary>
        /// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        [Input("updatedBefore")]
        public Input<string>? UpdatedBefore { get; set; }

        /// <summary>
        /// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        /// <summary>
        /// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
        /// </summary>
        [Input("withLabelsDetails")]
        public Input<bool>? WithLabelsDetails { get; set; }

        public GetProjectIssuesInvokeArgs()
        {
        }
        public static new GetProjectIssuesInvokeArgs Empty => new GetProjectIssuesInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectIssuesResult
    {
        /// <summary>
        /// Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
        /// </summary>
        public readonly int? AssigneeId;
        /// <summary>
        /// Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
        /// </summary>
        public readonly string? AssigneeUsername;
        /// <summary>
        /// Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
        /// </summary>
        public readonly int? AuthorId;
        /// <summary>
        /// Filter confidential or public issues.
        /// </summary>
        public readonly bool? Confidential;
        /// <summary>
        /// Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        public readonly string? CreatedAfter;
        /// <summary>
        /// Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        public readonly string? CreatedBefore;
        /// <summary>
        /// Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
        /// </summary>
        public readonly string? DueDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Return only the issues having the given iid
        /// </summary>
        public readonly ImmutableArray<int> Iids;
        /// <summary>
        /// Filter to a given type of issue. Valid values are [issue incident test_case].
        /// </summary>
        public readonly string? IssueType;
        /// <summary>
        /// The list of issues returned by the search.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectIssuesIssueResult> Issues;
        /// <summary>
        /// Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
        /// </summary>
        public readonly string? Milestone;
        /// <summary>
        /// Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
        /// </summary>
        public readonly string? MyReactionEmoji;
        /// <summary>
        /// Return issues that do not match the assignee id.
        /// </summary>
        public readonly ImmutableArray<int> NotAssigneeIds;
        /// <summary>
        /// Return issues that do not match the author id.
        /// </summary>
        public readonly ImmutableArray<int> NotAuthorIds;
        /// <summary>
        /// Return issues that do not match the labels.
        /// </summary>
        public readonly ImmutableArray<string> NotLabels;
        /// <summary>
        /// Return issues that do not match the milestone.
        /// </summary>
        public readonly string? NotMilestone;
        /// <summary>
        /// Return issues not reacted by the authenticated user by the given emoji.
        /// </summary>
        public readonly ImmutableArray<string> NotMyReactionEmojis;
        /// <summary>
        /// Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
        /// </summary>
        public readonly string? OrderBy;
        /// <summary>
        /// The name or id of the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Search project issues against their title and description
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// Return issues sorted in asc or desc order. Default is desc
        /// </summary>
        public readonly string? Sort;
        /// <summary>
        /// Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        public readonly string? UpdatedAfter;
        /// <summary>
        /// Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
        /// </summary>
        public readonly string? UpdatedBefore;
        /// <summary>
        /// Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
        /// </summary>
        public readonly int? Weight;
        /// <summary>
        /// If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false.
        /// </summary>
        public readonly bool? WithLabelsDetails;

        [OutputConstructor]
        private GetProjectIssuesResult(
            int? assigneeId,

            string? assigneeUsername,

            int? authorId,

            bool? confidential,

            string? createdAfter,

            string? createdBefore,

            string? dueDate,

            string id,

            ImmutableArray<int> iids,

            string? issueType,

            ImmutableArray<Outputs.GetProjectIssuesIssueResult> issues,

            ImmutableArray<string> labels,

            string? milestone,

            string? myReactionEmoji,

            ImmutableArray<int> notAssigneeIds,

            ImmutableArray<int> notAuthorIds,

            ImmutableArray<string> notLabels,

            string? notMilestone,

            ImmutableArray<string> notMyReactionEmojis,

            string? orderBy,

            string project,

            string? scope,

            string? search,

            string? sort,

            string? updatedAfter,

            string? updatedBefore,

            int? weight,

            bool? withLabelsDetails)
        {
            AssigneeId = assigneeId;
            AssigneeUsername = assigneeUsername;
            AuthorId = authorId;
            Confidential = confidential;
            CreatedAfter = createdAfter;
            CreatedBefore = createdBefore;
            DueDate = dueDate;
            Id = id;
            Iids = iids;
            IssueType = issueType;
            Issues = issues;
            Labels = labels;
            Milestone = milestone;
            MyReactionEmoji = myReactionEmoji;
            NotAssigneeIds = notAssigneeIds;
            NotAuthorIds = notAuthorIds;
            NotLabels = notLabels;
            NotMilestone = notMilestone;
            NotMyReactionEmojis = notMyReactionEmojis;
            OrderBy = orderBy;
            Project = project;
            Scope = scope;
            Search = search;
            Sort = sort;
            UpdatedAfter = updatedAfter;
            UpdatedBefore = updatedBefore;
            Weight = weight;
            WithLabelsDetails = withLabelsDetails;
        }
    }
}
