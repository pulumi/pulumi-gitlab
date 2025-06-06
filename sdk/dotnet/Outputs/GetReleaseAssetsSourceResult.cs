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
    public sealed class GetReleaseAssetsSourceResult
    {
        /// <summary>
        /// The format of the source
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// The URL of the source
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetReleaseAssetsSourceResult(
            string format,

            string url)
        {
            Format = format;
            Url = url;
        }
    }
}
