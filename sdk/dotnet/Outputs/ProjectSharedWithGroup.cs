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
    public sealed class ProjectSharedWithGroup
    {
        /// <summary>
        /// Group's sharing permissions. See [group members permission][group_members_permissions] for more info.
        /// Valid values are `guest`, `reporter`, `developer`, `master`.
        /// </summary>
        public readonly string GroupAccessLevel;
        /// <summary>
        /// Group id of the group you want to share the project with.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// Group's name.
        /// </summary>
        public readonly string? GroupName;

        [OutputConstructor]
        private ProjectSharedWithGroup(
            string groupAccessLevel,

            int groupId,

            string? groupName)
        {
            GroupAccessLevel = groupAccessLevel;
            GroupId = groupId;
            GroupName = groupName;
        }
    }
}
