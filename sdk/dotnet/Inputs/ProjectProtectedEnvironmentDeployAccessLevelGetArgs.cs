// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ProjectProtectedEnvironmentDeployAccessLevelGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Levels of access required to deploy to this protected environment. Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// Readable description of level of access.
        /// </summary>
        [Input("accessLevelDescription")]
        public Input<string>? AccessLevelDescription { get; set; }

        /// <summary>
        /// The ID of the group allowed to deploy to this protected environment. The project must be shared with the group. Mutually exclusive with `access_level` and `user_id`.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Group inheritance allows deploy access levels to take inherited group membership into account. Valid values are `0`, `1`. `0` =&gt; Direct group membership only, `1` =&gt; All inherited groups. Default: `0`
        /// </summary>
        [Input("groupInheritanceType")]
        public Input<int>? GroupInheritanceType { get; set; }

        /// <summary>
        /// The unique ID of the Deploy Access Level object.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The ID of the user allowed to deploy to this protected environment. The user must be a member of the project. Mutually exclusive with `access_level` and `group_id`.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public ProjectProtectedEnvironmentDeployAccessLevelGetArgs()
        {
        }
        public static new ProjectProtectedEnvironmentDeployAccessLevelGetArgs Empty => new ProjectProtectedEnvironmentDeployAccessLevelGetArgs();
    }
}
