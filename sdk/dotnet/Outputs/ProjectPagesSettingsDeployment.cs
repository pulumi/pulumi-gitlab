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
    public sealed class ProjectPagesSettingsDeployment
    {
        /// <summary>
        /// Date the deployment was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The path prefix of the deployment when using parallel deployments.
        /// </summary>
        public readonly string? PathPrefix;
        /// <summary>
        /// The root directory of the deployment.
        /// </summary>
        public readonly string? RootDirectory;
        /// <summary>
        /// The URL of the deployment.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private ProjectPagesSettingsDeployment(
            string? createdAt,

            string? pathPrefix,

            string? rootDirectory,

            string? url)
        {
            CreatedAt = createdAt;
            PathPrefix = pathPrefix;
            RootDirectory = rootDirectory;
            Url = url;
        }
    }
}
