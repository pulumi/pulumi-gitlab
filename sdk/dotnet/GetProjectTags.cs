// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectTags
    {
        /// <summary>
        /// The `gitlab.getProjectTags` data source allows details of project tags to be retrieved by some search criteria.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)
        /// </summary>
        public static Task<GetProjectTagsResult> InvokeAsync(GetProjectTagsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectTagsResult>("gitlab:index/getProjectTags:getProjectTags", args ?? new GetProjectTagsArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectTags` data source allows details of project tags to be retrieved by some search criteria.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)
        /// </summary>
        public static Output<GetProjectTagsResult> Invoke(GetProjectTagsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectTagsResult>("gitlab:index/getProjectTags:getProjectTags", args ?? new GetProjectTagsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectTags` data source allows details of project tags to be retrieved by some search criteria.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)
        /// </summary>
        public static Output<GetProjectTagsResult> Invoke(GetProjectTagsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectTagsResult>("gitlab:index/getProjectTags:getProjectTags", args ?? new GetProjectTagsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectTagsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return tags ordered by `name` or `updated` fields. Default is `updated`.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// Return tags sorted in `asc` or `desc` order. Default is `desc`.
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        public GetProjectTagsArgs()
        {
        }
        public static new GetProjectTagsArgs Empty => new GetProjectTagsArgs();
    }

    public sealed class GetProjectTagsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return tags ordered by `name` or `updated` fields. Default is `updated`.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Return tags sorted in `asc` or `desc` order. Default is `desc`.
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        public GetProjectTagsInvokeArgs()
        {
        }
        public static new GetProjectTagsInvokeArgs Empty => new GetProjectTagsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectTagsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Return tags ordered by `name` or `updated` fields. Default is `updated`.
        /// </summary>
        public readonly string? OrderBy;
        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// Return tags sorted in `asc` or `desc` order. Default is `desc`.
        /// </summary>
        public readonly string? Sort;
        /// <summary>
        /// List of repository tags from a project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectTagsTagResult> Tags;

        [OutputConstructor]
        private GetProjectTagsResult(
            string id,

            string? orderBy,

            string project,

            string? search,

            string? sort,

            ImmutableArray<Outputs.GetProjectTagsTagResult> tags)
        {
            Id = id;
            OrderBy = orderBy;
            Project = project;
            Search = search;
            Sort = sort;
            Tags = tags;
        }
    }
}
