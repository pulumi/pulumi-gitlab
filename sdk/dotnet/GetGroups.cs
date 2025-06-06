// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroups
    {
        /// <summary>
        /// The `gitlab.getGroups` data source allows details of multiple groups to be retrieved given some optional filter criteria.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; Some available options require administrator privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-groups)
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
        ///     var example = GitLab.GetGroups.Invoke(new()
        ///     {
        ///         Sort = "desc",
        ///         OrderBy = "name",
        ///     });
        /// 
        ///     var example_two = GitLab.GetGroups.Invoke(new()
        ///     {
        ///         Search = "GitLab",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("gitlab:index/getGroups:getGroups", args ?? new GetGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroups` data source allows details of multiple groups to be retrieved given some optional filter criteria.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; Some available options require administrator privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-groups)
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
        ///     var example = GitLab.GetGroups.Invoke(new()
        ///     {
        ///         Sort = "desc",
        ///         OrderBy = "name",
        ///     });
        /// 
        ///     var example_two = GitLab.GetGroups.Invoke(new()
        ///     {
        ///         Search = "GitLab",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupsResult> Invoke(GetGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupsResult>("gitlab:index/getGroups:getGroups", args ?? new GetGroupsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroups` data source allows details of multiple groups to be retrieved given some optional filter criteria.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; Some available options require administrator privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-groups)
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
        ///     var example = GitLab.GetGroups.Invoke(new()
        ///     {
        ///         Sort = "desc",
        ///         OrderBy = "name",
        ///     });
        /// 
        ///     var example_two = GitLab.GetGroups.Invoke(new()
        ///     {
        ///         Search = "GitLab",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupsResult> Invoke(GetGroupsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupsResult>("gitlab:index/getGroups:getGroups", args ?? new GetGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Search groups by name or path.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// Sort groups' list in asc or desc order. (Requires administrator privileges)
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        /// <summary>
        /// Limit to top level groups, excluding all subgroups.
        /// </summary>
        [Input("topLevelOnly")]
        public bool? TopLevelOnly { get; set; }

        public GetGroupsArgs()
        {
        }
        public static new GetGroupsArgs Empty => new GetGroupsArgs();
    }

    public sealed class GetGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Search groups by name or path.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Sort groups' list in asc or desc order. (Requires administrator privileges)
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        /// <summary>
        /// Limit to top level groups, excluding all subgroups.
        /// </summary>
        [Input("topLevelOnly")]
        public Input<bool>? TopLevelOnly { get; set; }

        public GetGroupsInvokeArgs()
        {
        }
        public static new GetGroupsInvokeArgs Empty => new GetGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        /// <summary>
        /// The list of groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
        /// </summary>
        public readonly string? OrderBy;
        /// <summary>
        /// Search groups by name or path.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// Sort groups' list in asc or desc order. (Requires administrator privileges)
        /// </summary>
        public readonly string? Sort;
        /// <summary>
        /// Limit to top level groups, excluding all subgroups.
        /// </summary>
        public readonly bool? TopLevelOnly;

        [OutputConstructor]
        private GetGroupsResult(
            ImmutableArray<Outputs.GetGroupsGroupResult> groups,

            string id,

            string? orderBy,

            string? search,

            string? sort,

            bool? topLevelOnly)
        {
            Groups = groups;
            Id = id;
            OrderBy = orderBy;
            Search = search;
            Sort = sort;
            TopLevelOnly = topLevelOnly;
        }
    }
}
