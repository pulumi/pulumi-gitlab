// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ProjectContainerExpirationPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cadence of the policy. Valid values are: `1d`, `7d`, `14d`, `1month`, `3month`.
        /// </summary>
        [Input("cadence")]
        public Input<string>? Cadence { get; set; }

        /// <summary>
        /// If true, the policy is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The number of images to keep.
        /// </summary>
        [Input("keepN")]
        public Input<int>? KeepN { get; set; }

        /// <summary>
        /// The regular expression to match image names to delete.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The regular expression to match image names to delete.
        /// </summary>
        [Input("nameRegexDelete")]
        public Input<string>? NameRegexDelete { get; set; }

        /// <summary>
        /// The regular expression to match image names to keep.
        /// </summary>
        [Input("nameRegexKeep")]
        public Input<string>? NameRegexKeep { get; set; }

        /// <summary>
        /// The next time the policy will run.
        /// </summary>
        [Input("nextRunAt")]
        public Input<string>? NextRunAt { get; set; }

        /// <summary>
        /// The number of days to keep images.
        /// </summary>
        [Input("olderThan")]
        public Input<string>? OlderThan { get; set; }

        public ProjectContainerExpirationPolicyGetArgs()
        {
        }
        public static new ProjectContainerExpirationPolicyGetArgs Empty => new ProjectContainerExpirationPolicyGetArgs();
    }
}
