// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GroupProtectedEnvironmentDeployAccessLevelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Levels of access required to deploy to this protected environment. Valid values are `developer`, `maintainer`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Readable description of level of access.
        /// </summary>
        [Input("accessLevelDescription")]
        public Input<string>? AccessLevelDescription { get; set; }

        /// <summary>
        /// The ID of the group allowed to deploy to this protected environment. The group must be a sub-group under the given group.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// The unique ID of the Deploy Access Level object.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The ID of the user allowed to deploy to this protected environment. The user must be a member of the group with Maintainer role or higher.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GroupProtectedEnvironmentDeployAccessLevelArgs()
        {
        }
        public static new GroupProtectedEnvironmentDeployAccessLevelArgs Empty => new GroupProtectedEnvironmentDeployAccessLevelArgs();
    }
}
