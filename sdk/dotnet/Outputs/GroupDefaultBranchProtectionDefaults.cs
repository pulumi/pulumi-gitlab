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
    public sealed class GroupDefaultBranchProtectionDefaults
    {
        /// <summary>
        /// Allow force push for all users with push access.
        /// </summary>
        public readonly bool? AllowForcePush;
        /// <summary>
        /// An array of access levels allowed to merge. Valid values are: `developer`, `maintainer`, `no one`.
        /// </summary>
        public readonly ImmutableArray<string> AllowedToMerges;
        /// <summary>
        /// An array of access levels allowed to push. Valid values are: `developer`, `maintainer`, `no one`.
        /// </summary>
        public readonly ImmutableArray<string> AllowedToPushes;
        /// <summary>
        /// Allow developers to initial push.
        /// </summary>
        public readonly bool? DeveloperCanInitialPush;

        [OutputConstructor]
        private GroupDefaultBranchProtectionDefaults(
            bool? allowForcePush,

            ImmutableArray<string> allowedToMerges,

            ImmutableArray<string> allowedToPushes,

            bool? developerCanInitialPush)
        {
            AllowForcePush = allowForcePush;
            AllowedToMerges = allowedToMerges;
            AllowedToPushes = allowedToPushes;
            DeveloperCanInitialPush = developerCanInitialPush;
        }
    }
}
