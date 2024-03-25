// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GetReleaseAssetsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The number of assets for a release
        /// </summary>
        [Input("count", required: true)]
        public int Count { get; set; }

        [Input("links")]
        private List<Inputs.GetReleaseAssetsLinkArgs>? _links;

        /// <summary>
        /// The links for a release
        /// </summary>
        public List<Inputs.GetReleaseAssetsLinkArgs> Links
        {
            get => _links ?? (_links = new List<Inputs.GetReleaseAssetsLinkArgs>());
            set => _links = value;
        }

        [Input("sources")]
        private List<Inputs.GetReleaseAssetsSourceArgs>? _sources;

        /// <summary>
        /// The sources for a release
        /// </summary>
        public List<Inputs.GetReleaseAssetsSourceArgs> Sources
        {
            get => _sources ?? (_sources = new List<Inputs.GetReleaseAssetsSourceArgs>());
            set => _sources = value;
        }

        public GetReleaseAssetsArgs()
        {
        }
        public static new GetReleaseAssetsArgs Empty => new GetReleaseAssetsArgs();
    }
}
