// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ReleaseLinksArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL of the release's closed issues.
        /// </summary>
        [Input("closedIssuesUrl")]
        public Input<string>? ClosedIssuesUrl { get; set; }

        /// <summary>
        /// URL of the release's closed merge requests.
        /// </summary>
        [Input("closedMergeRequestsUrl")]
        public Input<string>? ClosedMergeRequestsUrl { get; set; }

        /// <summary>
        /// URL of the release's edit page.
        /// </summary>
        [Input("editUrl")]
        public Input<string>? EditUrl { get; set; }

        /// <summary>
        /// URL of the release's merged merge requests.
        /// </summary>
        [Input("mergedMergeRequestsUrl")]
        public Input<string>? MergedMergeRequestsUrl { get; set; }

        /// <summary>
        /// URL of the release's open issues.
        /// </summary>
        [Input("openedIssuesUrl")]
        public Input<string>? OpenedIssuesUrl { get; set; }

        /// <summary>
        /// URL of the release's open merge requests.
        /// </summary>
        [Input("openedMergeRequestsUrl")]
        public Input<string>? OpenedMergeRequestsUrl { get; set; }

        /// <summary>
        /// URL of the release.
        /// </summary>
        [Input("self")]
        public Input<string>? Self { get; set; }

        public ReleaseLinksArgs()
        {
        }
        public static new ReleaseLinksArgs Empty => new ReleaseLinksArgs();
    }
}
