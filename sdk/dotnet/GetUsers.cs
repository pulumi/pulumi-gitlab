// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetUsers
    {
        /// <summary>
        /// The `gitlab.getUsers` data source allows details of multiple users to be retrieved given some optional filter criteria.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; Some available options require administrator privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ce/api/users.html#list-users)
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
        ///     var example = GitLab.GetUsers.Invoke(new()
        ///     {
        ///         Sort = "desc",
        ///         OrderBy = "name",
        ///         CreatedBefore = "2019-01-01",
        ///     });
        /// 
        ///     var example_two = GitLab.GetUsers.Invoke(new()
        ///     {
        ///         Search = "username",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("gitlab:index/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getUsers` data source allows details of multiple users to be retrieved given some optional filter criteria.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; Some available options require administrator privileges.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ce/api/users.html#list-users)
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
        ///     var example = GitLab.GetUsers.Invoke(new()
        ///     {
        ///         Sort = "desc",
        ///         OrderBy = "name",
        ///         CreatedBefore = "2019-01-01",
        ///     });
        /// 
        ///     var example_two = GitLab.GetUsers.Invoke(new()
        ///     {
        ///         Search = "username",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsersResult>("gitlab:index/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter users that are active.
        /// </summary>
        [Input("active")]
        public bool? Active { get; set; }

        /// <summary>
        /// Filter users that are blocked.
        /// </summary>
        [Input("blocked")]
        public bool? Blocked { get; set; }

        /// <summary>
        /// Search for users created after a specific date. (Requires administrator privileges)
        /// </summary>
        [Input("createdAfter")]
        public string? CreatedAfter { get; set; }

        /// <summary>
        /// Search for users created before a specific date. (Requires administrator privileges)
        /// </summary>
        [Input("createdBefore")]
        public string? CreatedBefore { get; set; }

        /// <summary>
        /// Lookup users by external provider. (Requires administrator privileges)
        /// </summary>
        [Input("externProvider")]
        public string? ExternProvider { get; set; }

        /// <summary>
        /// Lookup users by external UID. (Requires administrator privileges)
        /// </summary>
        [Input("externUid")]
        public string? ExternUid { get; set; }

        /// <summary>
        /// Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges)
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Search users by username, name or email.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// Sort users' list in asc or desc order. (Requires administrator privileges)
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        public GetUsersArgs()
        {
        }
        public static new GetUsersArgs Empty => new GetUsersArgs();
    }

    public sealed class GetUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter users that are active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Filter users that are blocked.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// Search for users created after a specific date. (Requires administrator privileges)
        /// </summary>
        [Input("createdAfter")]
        public Input<string>? CreatedAfter { get; set; }

        /// <summary>
        /// Search for users created before a specific date. (Requires administrator privileges)
        /// </summary>
        [Input("createdBefore")]
        public Input<string>? CreatedBefore { get; set; }

        /// <summary>
        /// Lookup users by external provider. (Requires administrator privileges)
        /// </summary>
        [Input("externProvider")]
        public Input<string>? ExternProvider { get; set; }

        /// <summary>
        /// Lookup users by external UID. (Requires administrator privileges)
        /// </summary>
        [Input("externUid")]
        public Input<string>? ExternUid { get; set; }

        /// <summary>
        /// Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges)
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Search users by username, name or email.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// Sort users' list in asc or desc order. (Requires administrator privileges)
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        public GetUsersInvokeArgs()
        {
        }
        public static new GetUsersInvokeArgs Empty => new GetUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        /// <summary>
        /// Filter users that are active.
        /// </summary>
        public readonly bool? Active;
        /// <summary>
        /// Filter users that are blocked.
        /// </summary>
        public readonly bool? Blocked;
        /// <summary>
        /// Search for users created after a specific date. (Requires administrator privileges)
        /// </summary>
        public readonly string? CreatedAfter;
        /// <summary>
        /// Search for users created before a specific date. (Requires administrator privileges)
        /// </summary>
        public readonly string? CreatedBefore;
        /// <summary>
        /// Lookup users by external provider. (Requires administrator privileges)
        /// </summary>
        public readonly string? ExternProvider;
        /// <summary>
        /// Lookup users by external UID. (Requires administrator privileges)
        /// </summary>
        public readonly string? ExternUid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Order the users' list by `id`, `name`, `username`, `created_at` or `updated_at`. (Requires administrator privileges)
        /// </summary>
        public readonly string? OrderBy;
        /// <summary>
        /// Search users by username, name or email.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// Sort users' list in asc or desc order. (Requires administrator privileges)
        /// </summary>
        public readonly string? Sort;
        /// <summary>
        /// The list of users.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersUserResult> Users;

        [OutputConstructor]
        private GetUsersResult(
            bool? active,

            bool? blocked,

            string? createdAfter,

            string? createdBefore,

            string? externProvider,

            string? externUid,

            string id,

            string? orderBy,

            string? search,

            string? sort,

            ImmutableArray<Outputs.GetUsersUserResult> users)
        {
            Active = active;
            Blocked = blocked;
            CreatedAfter = createdAfter;
            CreatedBefore = createdBefore;
            ExternProvider = externProvider;
            ExternUid = externUid;
            Id = id;
            OrderBy = orderBy;
            Search = search;
            Sort = sort;
            Users = users;
        }
    }
}
