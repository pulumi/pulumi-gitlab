// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
        /// </summary>
        [Input("accessLevel", required: true)]
        public string AccessLevel { get; set; } = null!;

        /// <summary>
        /// Readable description of access level.
        /// </summary>
        [Input("accessLevelDescription", required: true)]
        public string AccessLevelDescription { get; set; } = null!;

        /// <summary>
        /// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
        /// </summary>
        [Input("groupId")]
        public int? GroupId { get; set; }

        /// <summary>
        /// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
        /// </summary>
        [Input("userId")]
        public int? UserId { get; set; }

        public GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs()
        {
        }
        public static new GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs Empty => new GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs();
    }
}
