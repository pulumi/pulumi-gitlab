// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GroupDefaultBranchProtectionDefaultsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow force push for all users with push access.
        /// </summary>
        [Input("allowForcePush")]
        public Input<bool>? AllowForcePush { get; set; }

        [Input("allowedToMerges")]
        private InputList<string>? _allowedToMerges;

        /// <summary>
        /// An array of access levels allowed to merge. Valid values are: `developer`, `maintainer`.
        /// </summary>
        public InputList<string> AllowedToMerges
        {
            get => _allowedToMerges ?? (_allowedToMerges = new InputList<string>());
            set => _allowedToMerges = value;
        }

        [Input("allowedToPushes")]
        private InputList<string>? _allowedToPushes;

        /// <summary>
        /// An array of access levels allowed to push. Valid values are: `developer`, `maintainer`.
        /// </summary>
        public InputList<string> AllowedToPushes
        {
            get => _allowedToPushes ?? (_allowedToPushes = new InputList<string>());
            set => _allowedToPushes = value;
        }

        /// <summary>
        /// Allow developers to initial push.
        /// </summary>
        [Input("developerCanInitialPush")]
        public Input<bool>? DeveloperCanInitialPush { get; set; }

        public GroupDefaultBranchProtectionDefaultsGetArgs()
        {
        }
        public static new GroupDefaultBranchProtectionDefaultsGetArgs Empty => new GroupDefaultBranchProtectionDefaultsGetArgs();
    }
}