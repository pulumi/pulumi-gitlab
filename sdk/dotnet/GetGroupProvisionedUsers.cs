// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroupProvisionedUsers
    {
        /// <summary>
        /// The `gitlab.getGroupProvisionedUsers` data source allows details of the provisioned users of a given group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-provisioned-users)
        /// </summary>
        public static Task<GetGroupProvisionedUsersResult> InvokeAsync(GetGroupProvisionedUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupProvisionedUsersResult>("gitlab:index/getGroupProvisionedUsers:getGroupProvisionedUsers", args ?? new GetGroupProvisionedUsersArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroupProvisionedUsers` data source allows details of the provisioned users of a given group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-provisioned-users)
        /// </summary>
        public static Output<GetGroupProvisionedUsersResult> Invoke(GetGroupProvisionedUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupProvisionedUsersResult>("gitlab:index/getGroupProvisionedUsers:getGroupProvisionedUsers", args ?? new GetGroupProvisionedUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupProvisionedUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return only active provisioned users.
        /// </summary>
        [Input("active")]
        public bool? Active { get; set; }

        /// <summary>
        /// Return only blocked provisioned users.
        /// </summary>
        [Input("blocked")]
        public bool? Blocked { get; set; }

        /// <summary>
        /// Return only provisioned users created on or after the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("createdAfter")]
        public string? CreatedAfter { get; set; }

        /// <summary>
        /// Return only provisioned users created on or before the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("createdBefore")]
        public string? CreatedBefore { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the group.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("provisionedUsers")]
        private List<Inputs.GetGroupProvisionedUsersProvisionedUserArgs>? _provisionedUsers;

        /// <summary>
        /// The list of provisioned users.
        /// </summary>
        public List<Inputs.GetGroupProvisionedUsersProvisionedUserArgs> ProvisionedUsers
        {
            get => _provisionedUsers ?? (_provisionedUsers = new List<Inputs.GetGroupProvisionedUsersProvisionedUserArgs>());
            set => _provisionedUsers = value;
        }

        /// <summary>
        /// The search query to filter the provisioned users.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// The username of the provisioned user.
        /// </summary>
        [Input("username")]
        public string? Username { get; set; }

        public GetGroupProvisionedUsersArgs()
        {
        }
        public static new GetGroupProvisionedUsersArgs Empty => new GetGroupProvisionedUsersArgs();
    }

    public sealed class GetGroupProvisionedUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return only active provisioned users.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Return only blocked provisioned users.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// Return only provisioned users created on or after the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("createdAfter")]
        public Input<string>? CreatedAfter { get; set; }

        /// <summary>
        /// Return only provisioned users created on or before the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        [Input("createdBefore")]
        public Input<string>? CreatedBefore { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the group.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("provisionedUsers")]
        private InputList<Inputs.GetGroupProvisionedUsersProvisionedUserInputArgs>? _provisionedUsers;

        /// <summary>
        /// The list of provisioned users.
        /// </summary>
        public InputList<Inputs.GetGroupProvisionedUsersProvisionedUserInputArgs> ProvisionedUsers
        {
            get => _provisionedUsers ?? (_provisionedUsers = new InputList<Inputs.GetGroupProvisionedUsersProvisionedUserInputArgs>());
            set => _provisionedUsers = value;
        }

        /// <summary>
        /// The search query to filter the provisioned users.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// The username of the provisioned user.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GetGroupProvisionedUsersInvokeArgs()
        {
        }
        public static new GetGroupProvisionedUsersInvokeArgs Empty => new GetGroupProvisionedUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupProvisionedUsersResult
    {
        /// <summary>
        /// Return only active provisioned users.
        /// </summary>
        public readonly bool? Active;
        /// <summary>
        /// Return only blocked provisioned users.
        /// </summary>
        public readonly bool? Blocked;
        /// <summary>
        /// Return only provisioned users created on or after the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        public readonly string? CreatedAfter;
        /// <summary>
        /// Return only provisioned users created on or before the specified date. Expected in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
        /// </summary>
        public readonly string? CreatedBefore;
        /// <summary>
        /// The ID or URL-encoded path of the group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of provisioned users.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupProvisionedUsersProvisionedUserResult> ProvisionedUsers;
        /// <summary>
        /// The search query to filter the provisioned users.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// The username of the provisioned user.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private GetGroupProvisionedUsersResult(
            bool? active,

            bool? blocked,

            string? createdAfter,

            string? createdBefore,

            string id,

            ImmutableArray<Outputs.GetGroupProvisionedUsersProvisionedUserResult> provisionedUsers,

            string? search,

            string? username)
        {
            Active = active;
            Blocked = blocked;
            CreatedAfter = createdAfter;
            CreatedBefore = createdBefore;
            Id = id;
            ProvisionedUsers = provisionedUsers;
            Search = search;
            Username = username;
        }
    }
}
