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
    public sealed class GetProjectsProjectOwnerResult
    {
        /// <summary>
        /// The avatar url of the owner.
        /// </summary>
        public readonly string AvatarUrl;
        /// <summary>
        /// The ID of the owner.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The name of the owner.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the owner.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The username of the owner.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The website url of the owner.
        /// </summary>
        public readonly string WebsiteUrl;

        [OutputConstructor]
        private GetProjectsProjectOwnerResult(
            string avatarUrl,

            int id,

            string name,

            string state,

            string username,

            string websiteUrl)
        {
            AvatarUrl = avatarUrl;
            Id = id;
            Name = name;
            State = state;
            Username = username;
            WebsiteUrl = websiteUrl;
        }
    }
}
