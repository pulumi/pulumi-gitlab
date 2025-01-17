// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectMembership
    {
        /// <summary>
        /// The `gitlab.ProjectMembership` data source allows to list and filter all members of a project specified by either its id or full path.
        /// 
        /// &gt; **Note** exactly one of project_id or full_path must be provided.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
        /// </summary>
        public static Task<GetProjectMembershipResult> InvokeAsync(GetProjectMembershipArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectMembershipResult>("gitlab:index/getProjectMembership:getProjectMembership", args ?? new GetProjectMembershipArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.ProjectMembership` data source allows to list and filter all members of a project specified by either its id or full path.
        /// 
        /// &gt; **Note** exactly one of project_id or full_path must be provided.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
        /// </summary>
        public static Output<GetProjectMembershipResult> Invoke(GetProjectMembershipInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectMembershipResult>("gitlab:index/getProjectMembership:getProjectMembership", args ?? new GetProjectMembershipInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.ProjectMembership` data source allows to list and filter all members of a project specified by either its id or full path.
        /// 
        /// &gt; **Note** exactly one of project_id or full_path must be provided.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
        /// </summary>
        public static Output<GetProjectMembershipResult> Invoke(GetProjectMembershipInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectMembershipResult>("gitlab:index/getProjectMembership:getProjectMembership", args ?? new GetProjectMembershipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectMembershipArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full path of the project.
        /// </summary>
        [Input("fullPath")]
        public string? FullPath { get; set; }

        /// <summary>
        /// Return all project members including members through ancestor groups
        /// </summary>
        [Input("inherited")]
        public bool? Inherited { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public int? ProjectId { get; set; }

        /// <summary>
        /// A query string to search for members
        /// </summary>
        [Input("query")]
        public string? Query { get; set; }

        [Input("userIds")]
        private List<int>? _userIds;

        /// <summary>
        /// List of user ids to filter members by
        /// </summary>
        public List<int> UserIds
        {
            get => _userIds ?? (_userIds = new List<int>());
            set => _userIds = value;
        }

        public GetProjectMembershipArgs()
        {
        }
        public static new GetProjectMembershipArgs Empty => new GetProjectMembershipArgs();
    }

    public sealed class GetProjectMembershipInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full path of the project.
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// Return all project members including members through ancestor groups
        /// </summary>
        [Input("inherited")]
        public Input<bool>? Inherited { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// A query string to search for members
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        [Input("userIds")]
        private InputList<int>? _userIds;

        /// <summary>
        /// List of user ids to filter members by
        /// </summary>
        public InputList<int> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<int>());
            set => _userIds = value;
        }

        public GetProjectMembershipInvokeArgs()
        {
        }
        public static new GetProjectMembershipInvokeArgs Empty => new GetProjectMembershipInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectMembershipResult
    {
        /// <summary>
        /// The full path of the project.
        /// </summary>
        public readonly string FullPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Return all project members including members through ancestor groups
        /// </summary>
        public readonly bool? Inherited;
        /// <summary>
        /// The list of project members.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectMembershipMemberResult> Members;
        /// <summary>
        /// The ID of the project.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// A query string to search for members
        /// </summary>
        public readonly string? Query;
        /// <summary>
        /// List of user ids to filter members by
        /// </summary>
        public readonly ImmutableArray<int> UserIds;

        [OutputConstructor]
        private GetProjectMembershipResult(
            string fullPath,

            string id,

            bool? inherited,

            ImmutableArray<Outputs.GetProjectMembershipMemberResult> members,

            int projectId,

            string? query,

            ImmutableArray<int> userIds)
        {
            FullPath = fullPath;
            Id = id;
            Inherited = inherited;
            Members = members;
            ProjectId = projectId;
            Query = query;
            UserIds = userIds;
        }
    }
}
