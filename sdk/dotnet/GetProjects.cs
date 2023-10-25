// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjects
    {
        /// <summary>
        /// The `gitlab.getProjects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.
        /// 
        /// &gt; This data source supports all available filters exposed by the xanzy/go-gitlab package, which might not expose all available filters exposed by the Gitlab APIs.
        /// 
        /// &gt; The owner sub-attributes are only populated if the Gitlab token used has an administrator scope.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects)
        /// </summary>
        public static Task<GetProjectsResult> InvokeAsync(GetProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectsResult>("gitlab:index/getProjects:getProjects", args ?? new GetProjectsArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.
        /// 
        /// &gt; This data source supports all available filters exposed by the xanzy/go-gitlab package, which might not expose all available filters exposed by the Gitlab APIs.
        /// 
        /// &gt; The owner sub-attributes are only populated if the Gitlab token used has an administrator scope.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects)
        /// </summary>
        public static Output<GetProjectsResult> Invoke(GetProjectsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectsResult>("gitlab:index/getProjects:getProjects", args ?? new GetProjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectsArgs : global::Pulumi.InvokeArgs
    {
        [Input("archived")]
        public bool? Archived { get; set; }

        [Input("groupId")]
        public int? GroupId { get; set; }

        /// <summary>
        /// Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
        /// </summary>
        [Input("includeSubgroups")]
        public bool? IncludeSubgroups { get; set; }

        /// <summary>
        /// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
        /// </summary>
        [Input("maxQueryablePages")]
        public int? MaxQueryablePages { get; set; }

        /// <summary>
        /// Limit by projects that the current user is a member of.
        /// </summary>
        [Input("membership")]
        public bool? Membership { get; set; }

        /// <summary>
        /// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
        /// </summary>
        [Input("minAccessLevel")]
        public int? MinAccessLevel { get; set; }

        /// <summary>
        /// Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Limit by projects owned by the current user.
        /// </summary>
        [Input("owned")]
        public bool? Owned { get; set; }

        /// <summary>
        /// The first page to begin the query on.
        /// </summary>
        [Input("page")]
        public int? Page { get; set; }

        /// <summary>
        /// The number of results to return per page.
        /// </summary>
        [Input("perPage")]
        public int? PerPage { get; set; }

        /// <summary>
        /// Return list of authorized projects matching the search criteria.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// Return only the ID, URL, name, and path of each project.
        /// </summary>
        [Input("simple")]
        public bool? Simple { get; set; }

        /// <summary>
        /// Return projects sorted in `asc` or `desc` order. Default is `desc`.
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        /// <summary>
        /// Limit by projects starred by the current user.
        /// </summary>
        [Input("starred")]
        public bool? Starred { get; set; }

        [Input("statistics")]
        public bool? Statistics { get; set; }

        [Input("topics")]
        private List<string>? _topics;

        /// <summary>
        /// Limit by projects that have all of the given topics.
        /// </summary>
        public List<string> Topics
        {
            get => _topics ?? (_topics = new List<string>());
            set => _topics = value;
        }

        [Input("visibility")]
        public string? Visibility { get; set; }

        /// <summary>
        /// Include custom attributes in response *(admins only)*.
        /// </summary>
        [Input("withCustomAttributes")]
        public bool? WithCustomAttributes { get; set; }

        /// <summary>
        /// Limit by projects with issues feature enabled. Default is `false`.
        /// </summary>
        [Input("withIssuesEnabled")]
        public bool? WithIssuesEnabled { get; set; }

        /// <summary>
        /// Limit by projects with merge requests feature enabled. Default is `false`.
        /// </summary>
        [Input("withMergeRequestsEnabled")]
        public bool? WithMergeRequestsEnabled { get; set; }

        /// <summary>
        /// Limit by projects which use the given programming language. Cannot be used with `group_id`.
        /// </summary>
        [Input("withProgrammingLanguage")]
        public string? WithProgrammingLanguage { get; set; }

        /// <summary>
        /// Include projects shared to this group. Default is `true`. Needs `group_id`.
        /// </summary>
        [Input("withShared")]
        public bool? WithShared { get; set; }

        public GetProjectsArgs()
        {
        }
        public static new GetProjectsArgs Empty => new GetProjectsArgs();
    }

    public sealed class GetProjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
        /// </summary>
        [Input("includeSubgroups")]
        public Input<bool>? IncludeSubgroups { get; set; }

        /// <summary>
        /// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
        /// </summary>
        [Input("maxQueryablePages")]
        public Input<int>? MaxQueryablePages { get; set; }

        /// <summary>
        /// Limit by projects that the current user is a member of.
        /// </summary>
        [Input("membership")]
        public Input<bool>? Membership { get; set; }

        /// <summary>
        /// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
        /// </summary>
        [Input("minAccessLevel")]
        public Input<int>? MinAccessLevel { get; set; }

        /// <summary>
        /// Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Limit by projects owned by the current user.
        /// </summary>
        [Input("owned")]
        public Input<bool>? Owned { get; set; }

        /// <summary>
        /// The first page to begin the query on.
        /// </summary>
        [Input("page")]
        public Input<int>? Page { get; set; }

        /// <summary>
        /// The number of results to return per page.
        /// </summary>
        [Input("perPage")]
        public Input<int>? PerPage { get; set; }

        /// <summary>
        /// Return list of authorized projects matching the search criteria.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Return only the ID, URL, name, and path of each project.
        /// </summary>
        [Input("simple")]
        public Input<bool>? Simple { get; set; }

        /// <summary>
        /// Return projects sorted in `asc` or `desc` order. Default is `desc`.
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        /// <summary>
        /// Limit by projects starred by the current user.
        /// </summary>
        [Input("starred")]
        public Input<bool>? Starred { get; set; }

        [Input("statistics")]
        public Input<bool>? Statistics { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// Limit by projects that have all of the given topics.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Include custom attributes in response *(admins only)*.
        /// </summary>
        [Input("withCustomAttributes")]
        public Input<bool>? WithCustomAttributes { get; set; }

        /// <summary>
        /// Limit by projects with issues feature enabled. Default is `false`.
        /// </summary>
        [Input("withIssuesEnabled")]
        public Input<bool>? WithIssuesEnabled { get; set; }

        /// <summary>
        /// Limit by projects with merge requests feature enabled. Default is `false`.
        /// </summary>
        [Input("withMergeRequestsEnabled")]
        public Input<bool>? WithMergeRequestsEnabled { get; set; }

        /// <summary>
        /// Limit by projects which use the given programming language. Cannot be used with `group_id`.
        /// </summary>
        [Input("withProgrammingLanguage")]
        public Input<string>? WithProgrammingLanguage { get; set; }

        /// <summary>
        /// Include projects shared to this group. Default is `true`. Needs `group_id`.
        /// </summary>
        [Input("withShared")]
        public Input<bool>? WithShared { get; set; }

        public GetProjectsInvokeArgs()
        {
        }
        public static new GetProjectsInvokeArgs Empty => new GetProjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectsResult
    {
        /// <summary>
        /// Limit by archived status.
        /// </summary>
        public readonly bool? Archived;
        /// <summary>
        /// The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
        /// </summary>
        public readonly int? GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
        /// </summary>
        public readonly bool? IncludeSubgroups;
        /// <summary>
        /// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
        /// </summary>
        public readonly int? MaxQueryablePages;
        /// <summary>
        /// Limit by projects that the current user is a member of.
        /// </summary>
        public readonly bool? Membership;
        /// <summary>
        /// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
        /// </summary>
        public readonly int? MinAccessLevel;
        /// <summary>
        /// Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
        /// </summary>
        public readonly string? OrderBy;
        /// <summary>
        /// Limit by projects owned by the current user.
        /// </summary>
        public readonly bool? Owned;
        /// <summary>
        /// The first page to begin the query on.
        /// </summary>
        public readonly int? Page;
        /// <summary>
        /// The number of results to return per page.
        /// </summary>
        public readonly int? PerPage;
        /// <summary>
        /// A list containing the projects matching the supplied arguments
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectResult> Projects;
        /// <summary>
        /// Return list of authorized projects matching the search criteria.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// Return only the ID, URL, name, and path of each project.
        /// </summary>
        public readonly bool? Simple;
        /// <summary>
        /// Return projects sorted in `asc` or `desc` order. Default is `desc`.
        /// </summary>
        public readonly string? Sort;
        /// <summary>
        /// Limit by projects starred by the current user.
        /// </summary>
        public readonly bool? Starred;
        /// <summary>
        /// Include project statistics. Cannot be used with `group_id`.
        /// </summary>
        public readonly bool? Statistics;
        /// <summary>
        /// Limit by projects that have all of the given topics.
        /// </summary>
        public readonly ImmutableArray<string> Topics;
        /// <summary>
        /// Limit by visibility `public`, `internal`, or `private`.
        /// </summary>
        public readonly string? Visibility;
        /// <summary>
        /// Include custom attributes in response *(admins only)*.
        /// </summary>
        public readonly bool? WithCustomAttributes;
        /// <summary>
        /// Limit by projects with issues feature enabled. Default is `false`.
        /// </summary>
        public readonly bool? WithIssuesEnabled;
        /// <summary>
        /// Limit by projects with merge requests feature enabled. Default is `false`.
        /// </summary>
        public readonly bool? WithMergeRequestsEnabled;
        /// <summary>
        /// Limit by projects which use the given programming language. Cannot be used with `group_id`.
        /// </summary>
        public readonly string? WithProgrammingLanguage;
        /// <summary>
        /// Include projects shared to this group. Default is `true`. Needs `group_id`.
        /// </summary>
        public readonly bool? WithShared;

        [OutputConstructor]
        private GetProjectsResult(
            bool? archived,

            int? groupId,

            string id,

            bool? includeSubgroups,

            int? maxQueryablePages,

            bool? membership,

            int? minAccessLevel,

            string? orderBy,

            bool? owned,

            int? page,

            int? perPage,

            ImmutableArray<Outputs.GetProjectsProjectResult> projects,

            string? search,

            bool? simple,

            string? sort,

            bool? starred,

            bool? statistics,

            ImmutableArray<string> topics,

            string? visibility,

            bool? withCustomAttributes,

            bool? withIssuesEnabled,

            bool? withMergeRequestsEnabled,

            string? withProgrammingLanguage,

            bool? withShared)
        {
            Archived = archived;
            GroupId = groupId;
            Id = id;
            IncludeSubgroups = includeSubgroups;
            MaxQueryablePages = maxQueryablePages;
            Membership = membership;
            MinAccessLevel = minAccessLevel;
            OrderBy = orderBy;
            Owned = owned;
            Page = page;
            PerPage = perPage;
            Projects = projects;
            Search = search;
            Simple = simple;
            Sort = sort;
            Starred = starred;
            Statistics = statistics;
            Topics = topics;
            Visibility = visibility;
            WithCustomAttributes = withCustomAttributes;
            WithIssuesEnabled = withIssuesEnabled;
            WithMergeRequestsEnabled = withMergeRequestsEnabled;
            WithProgrammingLanguage = withProgrammingLanguage;
            WithShared = withShared;
        }
    }
}
