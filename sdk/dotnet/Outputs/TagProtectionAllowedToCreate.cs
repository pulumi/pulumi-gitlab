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
    public sealed class TagProtectionAllowedToCreate
    {
        /// <summary>
        /// Access levels allowed to create protected tags. Valid values are: `no one`, `developer`, `maintainer`.
        /// </summary>
        public readonly string? AccessLevel;
        /// <summary>
        /// Readable description of access level.
        /// </summary>
        public readonly string? AccessLevelDescription;
        /// <summary>
        /// The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
        /// </summary>
        public readonly int? GroupId;
        /// <summary>
        /// The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
        /// </summary>
        public readonly int? UserId;

        [OutputConstructor]
        private TagProtectionAllowedToCreate(
            string? accessLevel,

            string? accessLevelDescription,

            int? groupId,

            int? userId)
        {
            AccessLevel = accessLevel;
            AccessLevelDescription = accessLevelDescription;
            GroupId = groupId;
            UserId = userId;
        }
    }
}
