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
    public sealed class GetGroupProvisionedUsersProvisionedUserResult
    {
        /// <summary>
        /// The avatar URL of the provisioned user.
        /// </summary>
        public readonly string AvatarUrl;
        /// <summary>
        /// The bio of the provisioned user.
        /// </summary>
        public readonly string Bio;
        /// <summary>
        /// Whether the provisioned user is a bot.
        /// </summary>
        public readonly bool Bot;
        /// <summary>
        /// The confirmation date of the provisioned user.
        /// </summary>
        public readonly string ConfirmedAt;
        /// <summary>
        /// The creation date of the provisioned user.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The email of the provisioned user.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Whether the provisioned user is external.
        /// </summary>
        public readonly bool External;
        /// <summary>
        /// The ID of the provisioned user.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The job title of the provisioned user.
        /// </summary>
        public readonly string JobTitle;
        /// <summary>
        /// The last activity date of the provisioned user.
        /// </summary>
        public readonly string LastActivityOn;
        /// <summary>
        /// The last sign-in date of the provisioned user.
        /// </summary>
        public readonly string LastSignInAt;
        /// <summary>
        /// The LinkedIn ID of the provisioned user.
        /// </summary>
        public readonly string Linkedin;
        /// <summary>
        /// The location of the provisioned user.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the provisioned user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The organization of the provisioned user.
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// Whether the provisioned user has a private profile.
        /// </summary>
        public readonly bool PrivateProfile;
        /// <summary>
        /// The pronouns of the provisioned user.
        /// </summary>
        public readonly string Pronouns;
        /// <summary>
        /// The public email of the provisioned user.
        /// </summary>
        public readonly string PublicEmail;
        /// <summary>
        /// The Skype ID of the provisioned user.
        /// </summary>
        public readonly string Skype;
        /// <summary>
        /// The state of the provisioned user.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The Twitter ID of the provisioned user.
        /// </summary>
        public readonly string Twitter;
        /// <summary>
        /// Whether two-factor authentication is enabled for the provisioned user.
        /// </summary>
        public readonly bool TwoFactorEnabled;
        /// <summary>
        /// The username of the provisioned user.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The web URL of the provisioned user.
        /// </summary>
        public readonly string WebUrl;
        /// <summary>
        /// The website URL of the provisioned user.
        /// </summary>
        public readonly string WebsiteUrl;

        [OutputConstructor]
        private GetGroupProvisionedUsersProvisionedUserResult(
            string avatarUrl,

            string bio,

            bool bot,

            string confirmedAt,

            string createdAt,

            string email,

            bool external,

            string id,

            string jobTitle,

            string lastActivityOn,

            string lastSignInAt,

            string linkedin,

            string location,

            string name,

            string organization,

            bool privateProfile,

            string pronouns,

            string publicEmail,

            string skype,

            string state,

            string twitter,

            bool twoFactorEnabled,

            string username,

            string webUrl,

            string websiteUrl)
        {
            AvatarUrl = avatarUrl;
            Bio = bio;
            Bot = bot;
            ConfirmedAt = confirmedAt;
            CreatedAt = createdAt;
            Email = email;
            External = external;
            Id = id;
            JobTitle = jobTitle;
            LastActivityOn = lastActivityOn;
            LastSignInAt = lastSignInAt;
            Linkedin = linkedin;
            Location = location;
            Name = name;
            Organization = organization;
            PrivateProfile = privateProfile;
            Pronouns = pronouns;
            PublicEmail = publicEmail;
            Skype = skype;
            State = state;
            Twitter = twitter;
            TwoFactorEnabled = twoFactorEnabled;
            Username = username;
            WebUrl = webUrl;
            WebsiteUrl = websiteUrl;
        }
    }
}
