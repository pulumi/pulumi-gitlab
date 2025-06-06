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
    public sealed class GetUserSshkeysKeyResult
    {
        /// <summary>
        /// The time when this key was created in GitLab.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
        /// </summary>
        public readonly string ExpiresAt;
        /// <summary>
        /// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The ID of the ssh key.
        /// </summary>
        public readonly int KeyId;
        /// <summary>
        /// The title of the ssh key.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The ID or username of the user. If this field is omitted, this resource manages a SSH key for the current user. Otherwise, this resource manages a SSH key for the specified user, and an admin token is required.
        /// </summary>
        public readonly int UserId;

        [OutputConstructor]
        private GetUserSshkeysKeyResult(
            string createdAt,

            string expiresAt,

            string key,

            int keyId,

            string title,

            int userId)
        {
            CreatedAt = createdAt;
            ExpiresAt = expiresAt;
            Key = key;
            KeyId = keyId;
            Title = title;
            UserId = userId;
        }
    }
}
