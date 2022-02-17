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
    public sealed class GetProjectProtectedBranchMergeAccessLevelResult
    {
        public readonly string AccessLevel;
        public readonly string AccessLevelDescription;
        public readonly int GroupId;
        public readonly int UserId;

        [OutputConstructor]
        private GetProjectProtectedBranchMergeAccessLevelResult(
            string accessLevel,

            string accessLevelDescription,

            int groupId,

            int userId)
        {
            AccessLevel = accessLevel;
            AccessLevelDescription = accessLevelDescription;
            GroupId = groupId;
            UserId = userId;
        }
    }
}