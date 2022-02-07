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
        /// Provide details about a list of group members in the gitlab provider. The results include id, username, name and more about the requested members.
        /// 
        /// &gt; **Note**: exactly one of group_id or full_path must be provided.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(GitLab.GetGroupMembership.InvokeAsync(new GitLab.GetGroupMembershipArgs
        ///         {
        ///             FullPath = "foo/bar",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupMembershipResult> InvokeAsync(GetGroupMembershipArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupMembershipResult>("gitlab:index/getGroupMembership:getGroupMembership", args ?? new GetGroupMembershipArgs(), options.WithDefaults());

        /// <summary>
        /// Provide details about a list of group members in the gitlab provider. The results include id, username, name and more about the requested members.
        /// 
        /// &gt; **Note**: exactly one of group_id or full_path must be provided.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(GitLab.GetGroupMembership.InvokeAsync(new GitLab.GetGroupMembershipArgs
        ///         {
        ///             FullPath = "foo/bar",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupMembershipResult> Invoke(GetGroupMembershipInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupMembershipResult>("gitlab:index/getGroupMembership:getGroupMembership", args ?? new GetGroupMembershipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupMembershipArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
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

        public GetGroupMembershipArgs()
        {
        }
    }

    public sealed class GetGroupMembershipInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
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

        public GetGroupMembershipInvokeArgs()
        {
        }
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
        /// The list of group members.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMembershipMemberResult> Members;

        [OutputConstructor]
        private GetGroupMembershipResult(
            string accessLevel,

            string fullPath,

            int groupId,

            string id,

            ImmutableArray<Outputs.GetGroupMembershipMemberResult> members)
        {
            AccessLevel = accessLevel;
            FullPath = fullPath;
            GroupId = groupId;
            Id = id;
            Members = members;
        }
    }
}
