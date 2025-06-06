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
    public sealed class GetProjectsProjectPermissionResult
    {
        /// <summary>
        /// Group access level.
        /// </summary>
        public readonly ImmutableDictionary<string, int> GroupAccess;
        /// <summary>
        /// Project access level.
        /// </summary>
        public readonly ImmutableDictionary<string, int> ProjectAccess;

        [OutputConstructor]
        private GetProjectsProjectPermissionResult(
            ImmutableDictionary<string, int> groupAccess,

            ImmutableDictionary<string, int> projectAccess)
        {
            GroupAccess = groupAccess;
            ProjectAccess = projectAccess;
        }
    }
}
