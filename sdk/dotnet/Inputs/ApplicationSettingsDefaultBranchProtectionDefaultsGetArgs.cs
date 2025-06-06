// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ApplicationSettingsDefaultBranchProtectionDefaultsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow force push for all users with push access.
        /// </summary>
        [Input("allowForcePush")]
        public Input<bool>? AllowForcePush { get; set; }

        [Input("allowedToMerges")]
        private InputList<int>? _allowedToMerges;

        /// <summary>
        /// An array of access levels allowed to merge. Supports Developer (30) or Maintainer (40).
        /// </summary>
        public InputList<int> AllowedToMerges
        {
            get => _allowedToMerges ?? (_allowedToMerges = new InputList<int>());
            set => _allowedToMerges = value;
        }

        [Input("allowedToPushes")]
        private InputList<int>? _allowedToPushes;

        /// <summary>
        /// An array of access levels allowed to push. Supports Developer (30) or Maintainer (40).
        /// </summary>
        public InputList<int> AllowedToPushes
        {
            get => _allowedToPushes ?? (_allowedToPushes = new InputList<int>());
            set => _allowedToPushes = value;
        }

        /// <summary>
        /// Allow developers to initial push.
        /// </summary>
        [Input("developerCanInitialPush")]
        public Input<bool>? DeveloperCanInitialPush { get; set; }

        public ApplicationSettingsDefaultBranchProtectionDefaultsGetArgs()
        {
        }
        public static new ApplicationSettingsDefaultBranchProtectionDefaultsGetArgs Empty => new ApplicationSettingsDefaultBranchProtectionDefaultsGetArgs();
    }
}
