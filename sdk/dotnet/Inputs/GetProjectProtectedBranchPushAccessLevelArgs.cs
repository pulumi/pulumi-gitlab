// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GetProjectProtectedBranchPushAccessLevelInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// Readable description of access level.
        /// </summary>
        [Input("accessLevelDescription", required: true)]
        public Input<string> AccessLevelDescription { get; set; } = null!;

        /// <summary>
        /// The ID of a GitLab deploy key allowed to perform the relevant action. Mutually exclusive with `group_id` and `user_id`. This field is read-only until Gitlab 17.5.
        /// </summary>
        [Input("deployKeyId")]
        public Input<int>? DeployKeyId { get; set; }

        /// <summary>
        /// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `user_id`.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `deploy_key_id` and `group_id`.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GetProjectProtectedBranchPushAccessLevelInputArgs()
        {
        }
        public static new GetProjectProtectedBranchPushAccessLevelInputArgs Empty => new GetProjectProtectedBranchPushAccessLevelInputArgs();
    }
}
