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
    public sealed class GetPipelineSchedulesPipelineScheduleOwnerResult
    {
        /// <summary>
        /// Image URL for the user's avatar.
        /// </summary>
        public readonly string AvatarUrl;
        /// <summary>
        /// The user ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User's state, one of: active, blocked.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Username.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// URL to the user's profile.
        /// </summary>
        public readonly string WebUrl;

        [OutputConstructor]
        private GetPipelineSchedulesPipelineScheduleOwnerResult(
            string avatarUrl,

            int id,

            string name,

            string state,

            string username,

            string webUrl)
        {
            AvatarUrl = avatarUrl;
            Id = id;
            Name = name;
            State = state;
            Username = username;
            WebUrl = webUrl;
        }
    }
}
