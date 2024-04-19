// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Outputs
{

    [OutputType]
    public sealed class ProjectProtectedEnvironmentApprovalRule
    {
        /// <summary>
        /// Levels of access allowed to approve a deployment to this protected environment. Valid values are `developer`, `maintainer`.
        /// </summary>
        public readonly string? AccessLevel;
        /// <summary>
        /// Readable description of level of access.
        /// </summary>
        public readonly string? AccessLevelDescription;
        /// <summary>
        /// The ID of the group allowed to approve a deployment to this protected environment. The project must be shared with the group. This is mutually exclusive with user_id.
        /// </summary>
        public readonly int? GroupId;
        /// <summary>
        /// Group inheritance allows deploy access levels to take inherited group membership into account. Valid values are `0`, `1`. `0` =&gt; Direct group membership only, `1` =&gt; All inherited groups. Default: `0`
        /// </summary>
        public readonly int? GroupInheritanceType;
        /// <summary>
        /// The unique ID of the Approval Rules object.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// The number of approval required to allow deployment to this protected environment. This is mutually exclusive with user_id.
        /// </summary>
        public readonly int? RequiredApprovals;
        /// <summary>
        /// The ID of the user allowed to approve a deployment to this protected environment. The user must be a member of the project. This is mutually exclusive with group*id and required*approvals.
        /// </summary>
        public readonly int? UserId;

        [OutputConstructor]
        private ProjectProtectedEnvironmentApprovalRule(
            string? accessLevel,

            string? accessLevelDescription,

            int? groupId,

            int? groupInheritanceType,

            int? id,

            int? requiredApprovals,

            int? userId)
        {
            AccessLevel = accessLevel;
            AccessLevelDescription = accessLevelDescription;
            GroupId = groupId;
            GroupInheritanceType = groupInheritanceType;
            Id = id;
            RequiredApprovals = requiredApprovals;
            UserId = userId;
        }
    }
}
