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
    public sealed class GetProjectSharedWithGroupResult
    {
        /// <summary>
        /// The access_level permission level of the shared group.
        /// </summary>
        public readonly int GroupAccessLevel;
        /// <summary>
        /// The full path of the group shared with.
        /// </summary>
        public readonly string GroupFullPath;
        /// <summary>
        /// The ID of the group shared with.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// The name of the group shared with.
        /// </summary>
        public readonly string GroupName;

        [OutputConstructor]
        private GetProjectSharedWithGroupResult(
            int groupAccessLevel,

            string groupFullPath,

            int groupId,

            string groupName)
        {
            GroupAccessLevel = groupAccessLevel;
            GroupFullPath = groupFullPath;
            GroupId = groupId;
            GroupName = groupName;
        }
    }
}
