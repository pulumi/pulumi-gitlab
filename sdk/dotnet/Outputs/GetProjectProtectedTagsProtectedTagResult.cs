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
    public sealed class GetProjectProtectedTagsProtectedTagResult
    {
        /// <summary>
        /// Array of access levels/user(s)/group(s) allowed to create protected tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectProtectedTagsProtectedTagCreateAccessLevelResult> CreateAccessLevels;
        /// <summary>
        /// The name of the protected tag.
        /// </summary>
        public readonly string Tag;

        [OutputConstructor]
        private GetProjectProtectedTagsProtectedTagResult(
            ImmutableArray<Outputs.GetProjectProtectedTagsProtectedTagCreateAccessLevelResult> createAccessLevels,

            string tag)
        {
            CreateAccessLevels = createAccessLevels;
            Tag = tag;
        }
    }
}
