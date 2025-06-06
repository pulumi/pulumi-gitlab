// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class GetGroupProvisionedUsersProvisionedUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The avatar URL of the provisioned user.
        /// </summary>
        [Input("avatarUrl", required: true)]
        public string AvatarUrl { get; set; } = null!;

        /// <summary>
        /// The bio of the provisioned user.
        /// </summary>
        [Input("bio", required: true)]
        public string Bio { get; set; } = null!;

        /// <summary>
        /// Whether the provisioned user is a bot.
        /// </summary>
        [Input("bot", required: true)]
        public bool Bot { get; set; }

        /// <summary>
        /// The confirmation date of the provisioned user.
        /// </summary>
        [Input("confirmedAt", required: true)]
        public string ConfirmedAt { get; set; } = null!;

        /// <summary>
        /// The creation date of the provisioned user.
        /// </summary>
        [Input("createdAt", required: true)]
        public string CreatedAt { get; set; } = null!;

        /// <summary>
        /// The email of the provisioned user.
        /// </summary>
        [Input("email", required: true)]
        public string Email { get; set; } = null!;

        /// <summary>
        /// Whether the provisioned user is external.
        /// </summary>
        [Input("external", required: true)]
        public bool External { get; set; }

        /// <summary>
        /// The ID of the provisioned user.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The job title of the provisioned user.
        /// </summary>
        [Input("jobTitle", required: true)]
        public string JobTitle { get; set; } = null!;

        /// <summary>
        /// The last activity date of the provisioned user.
        /// </summary>
        [Input("lastActivityOn", required: true)]
        public string LastActivityOn { get; set; } = null!;

        /// <summary>
        /// The last sign-in date of the provisioned user.
        /// </summary>
        [Input("lastSignInAt", required: true)]
        public string LastSignInAt { get; set; } = null!;

        /// <summary>
        /// The LinkedIn ID of the provisioned user.
        /// </summary>
        [Input("linkedin", required: true)]
        public string Linkedin { get; set; } = null!;

        /// <summary>
        /// The location of the provisioned user.
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// The name of the provisioned user.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The organization of the provisioned user.
        /// </summary>
        [Input("organization", required: true)]
        public string Organization { get; set; } = null!;

        /// <summary>
        /// Whether the provisioned user has a private profile.
        /// </summary>
        [Input("privateProfile", required: true)]
        public bool PrivateProfile { get; set; }

        /// <summary>
        /// The pronouns of the provisioned user.
        /// </summary>
        [Input("pronouns", required: true)]
        public string Pronouns { get; set; } = null!;

        /// <summary>
        /// The public email of the provisioned user.
        /// </summary>
        [Input("publicEmail", required: true)]
        public string PublicEmail { get; set; } = null!;

        /// <summary>
        /// The Skype ID of the provisioned user.
        /// </summary>
        [Input("skype", required: true)]
        public string Skype { get; set; } = null!;

        /// <summary>
        /// The state of the provisioned user.
        /// </summary>
        [Input("state", required: true)]
        public string State { get; set; } = null!;

        /// <summary>
        /// The Twitter ID of the provisioned user.
        /// </summary>
        [Input("twitter", required: true)]
        public string Twitter { get; set; } = null!;

        /// <summary>
        /// Whether two-factor authentication is enabled for the provisioned user.
        /// </summary>
        [Input("twoFactorEnabled", required: true)]
        public bool TwoFactorEnabled { get; set; }

        /// <summary>
        /// The username of the provisioned user.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        /// <summary>
        /// The web URL of the provisioned user.
        /// </summary>
        [Input("webUrl", required: true)]
        public string WebUrl { get; set; } = null!;

        /// <summary>
        /// The website URL of the provisioned user.
        /// </summary>
        [Input("websiteUrl", required: true)]
        public string WebsiteUrl { get; set; } = null!;

        public GetGroupProvisionedUsersProvisionedUserArgs()
        {
        }
        public static new GetGroupProvisionedUsersProvisionedUserArgs Empty => new GetGroupProvisionedUsersProvisionedUserArgs();
    }
}
