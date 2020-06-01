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
    public sealed class GetProjectsProjectPermissionsResult
    {
        public readonly ImmutableDictionary<string, int> GroupAccess;
        public readonly ImmutableDictionary<string, int> ProjectAccess;

        [OutputConstructor]
        private GetProjectsProjectPermissionsResult(
            ImmutableDictionary<string, int> groupAccess,

            ImmutableDictionary<string, int> projectAccess)
        {
            GroupAccess = groupAccess;
            ProjectAccess = projectAccess;
        }
    }
}
