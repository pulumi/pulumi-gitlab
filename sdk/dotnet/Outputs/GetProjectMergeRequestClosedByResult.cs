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
    public sealed class GetProjectMergeRequestClosedByResult
    {
        /// <summary>
        /// A link to the user's avatar image.
        /// </summary>
        public readonly string AvatarUrl;
        /// <summary>
        /// The internal ID number of the user.
        /// </summary>
        public readonly double Id;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the user account.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The username of the user.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// A link to the user's profile page.
        /// </summary>
        public readonly string WebUrl;

        [OutputConstructor]
        private GetProjectMergeRequestClosedByResult(
            string avatarUrl,

            double id,

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
