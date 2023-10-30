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
    public sealed class GetGroupsGroupResult
    {
        public readonly int DefaultBranchProtection;
        public readonly string Description;
        public readonly string FullName;
        public readonly string FullPath;
        public readonly int GroupId;
        public readonly bool LfsEnabled;
        public readonly string Name;
        public readonly int ParentId;
        public readonly string Path;
        public readonly bool PreventForkingOutsideGroup;
        public readonly bool RequestAccessEnabled;
        public readonly string RunnersToken;
        public readonly string SharedRunnersSetting;
        public readonly string VisibilityLevel;
        public readonly string WebUrl;
        public readonly string WikiAccessLevel;

        [OutputConstructor]
        private GetGroupsGroupResult(
            int defaultBranchProtection,

            string description,

            string fullName,

            string fullPath,

            int groupId,

            bool lfsEnabled,

            string name,

            int parentId,

            string path,

            bool preventForkingOutsideGroup,

            bool requestAccessEnabled,

            string runnersToken,

            string sharedRunnersSetting,

            string visibilityLevel,

            string webUrl,

            string wikiAccessLevel)
        {
            DefaultBranchProtection = defaultBranchProtection;
            Description = description;
            FullName = fullName;
            FullPath = fullPath;
            GroupId = groupId;
            LfsEnabled = lfsEnabled;
            Name = name;
            ParentId = parentId;
            Path = path;
            PreventForkingOutsideGroup = preventForkingOutsideGroup;
            RequestAccessEnabled = requestAccessEnabled;
            RunnersToken = runnersToken;
            SharedRunnersSetting = sharedRunnersSetting;
            VisibilityLevel = visibilityLevel;
            WebUrl = webUrl;
            WikiAccessLevel = wikiAccessLevel;
        }
    }
}
