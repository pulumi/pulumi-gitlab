// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Outputs
{

    [OutputType]
    public sealed class GetProjectProtectedBranchesProtectedBranchResult
    {
        /// <summary>
        /// Whether force push is allowed.
        /// </summary>
        public readonly bool AllowForcePush;
        /// <summary>
        /// Reject code pushes that change files listed in the CODEOWNERS file.
        /// </summary>
        public readonly bool CodeOwnerApprovalRequired;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectProtectedBranchesProtectedBranchMergeAccessLevelResult> MergeAccessLevels;
        /// <summary>
        /// The name of the protected branch.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Array of access levels and user(s)/group(s) allowed to push to protected branch.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectProtectedBranchesProtectedBranchPushAccessLevelResult> PushAccessLevels;

        [OutputConstructor]
        private GetProjectProtectedBranchesProtectedBranchResult(
            bool allowForcePush,

            bool codeOwnerApprovalRequired,

            int id,

            ImmutableArray<Outputs.GetProjectProtectedBranchesProtectedBranchMergeAccessLevelResult> mergeAccessLevels,

            string name,

            ImmutableArray<Outputs.GetProjectProtectedBranchesProtectedBranchPushAccessLevelResult> pushAccessLevels)
        {
            AllowForcePush = allowForcePush;
            CodeOwnerApprovalRequired = codeOwnerApprovalRequired;
            Id = id;
            MergeAccessLevels = mergeAccessLevels;
            Name = name;
            PushAccessLevels = pushAccessLevels;
        }
    }
}
