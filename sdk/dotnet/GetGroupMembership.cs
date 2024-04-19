// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroupMembership
    {
        /// <summary>
        /// The `gitlab.GroupMembership` data source allows to list and filter all members of a group specified by either its id or full path.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
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
        ///     var example = GitLab.GetGroupMembership.Invoke(new()
        ///     {
        ///         FullPath = "foo/bar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGroupMembershipResult> InvokeAsync(GetGroupMembershipArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupMembershipResult>("gitlab:index/getGroupMembership:getGroupMembership", args ?? new GetGroupMembershipArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.GroupMembership` data source allows to list and filter all members of a group specified by either its id or full path.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
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
        ///     var example = GitLab.GetGroupMembership.Invoke(new()
        ///     {
        ///         FullPath = "foo/bar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupMembershipResult> Invoke(GetGroupMembershipInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupMembershipResult>("gitlab:index/getGroupMembership:getGroupMembership", args ?? new GetGroupMembershipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupMembershipArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessLevel")]
        public string? AccessLevel { get; set; }

        /// <summary>
        /// The full path of the group.
        /// </summary>
        [Input("fullPath")]
        public string? FullPath { get; set; }

        /// <summary>
        /// The ID of the group.
        /// </summary>
        [Input("groupId")]
        public int? GroupId { get; set; }

        /// <summary>
        /// Return all project members including members through ancestor groups.
        /// </summary>
        [Input("inherited")]
        public bool? Inherited { get; set; }

        public GetGroupMembershipArgs()
        {
        }
        public static new GetGroupMembershipArgs Empty => new GetGroupMembershipArgs();
    }

    public sealed class GetGroupMembershipInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The full path of the group.
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// The ID of the group.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Return all project members including members through ancestor groups.
        /// </summary>
        [Input("inherited")]
        public Input<bool>? Inherited { get; set; }

        public GetGroupMembershipInvokeArgs()
        {
        }
        public static new GetGroupMembershipInvokeArgs Empty => new GetGroupMembershipInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupMembershipResult
    {
        /// <summary>
        /// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        public readonly string AccessLevel;
        /// <summary>
        /// The full path of the group.
        /// </summary>
        public readonly string FullPath;
        /// <summary>
        /// The ID of the group.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Return all project members including members through ancestor groups.
        /// </summary>
        public readonly bool? Inherited;
        /// <summary>
        /// The list of group members.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMembershipMemberResult> Members;

        [OutputConstructor]
        private GetGroupMembershipResult(
            string accessLevel,

            string fullPath,

            int groupId,

            string id,

            bool? inherited,

            ImmutableArray<Outputs.GetGroupMembershipMemberResult> members)
        {
            AccessLevel = accessLevel;
            FullPath = fullPath;
            GroupId = groupId;
            Id = id;
            Inherited = inherited;
            Members = members;
        }
    }
}
