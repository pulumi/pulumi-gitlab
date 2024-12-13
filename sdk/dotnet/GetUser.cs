// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetUser
    {
        /// <summary>
        /// The `gitlab.User` data source allows details of a user to be retrieved by either the user ID, username or email address.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
        /// the most related match will prioritize an exact match if one is available.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("gitlab:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.User` data source allows details of a user to be retrieved by either the user ID, username or email address.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
        /// the most related match will prioritize an exact match if one is available.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("gitlab:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.User` data source allows details of a user to be retrieved by either the user ID, username or email address.
        /// 
        /// &gt; Some attributes might not be returned depending on if you're an admin or not.
        /// 
        /// &gt; When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
        /// the most related match will prioritize an exact match if one is available.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("gitlab:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        /// <summary>
        /// The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
        /// </summary>
        [Input("namespaceId")]
        public int? NamespaceId { get; set; }

        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Input("userId")]
        public int? UserId { get; set; }

        /// <summary>
        /// The username of the user.
        /// </summary>
        [Input("username")]
        public string? Username { get; set; }

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
        /// </summary>
        [Input("namespaceId")]
        public Input<int>? NamespaceId { get; set; }

        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        /// <summary>
        /// The username of the user.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The avatar URL of the user.
        /// </summary>
        public readonly string AvatarUrl;
        /// <summary>
        /// The bio of the user.
        /// </summary>
        public readonly string Bio;
        /// <summary>
        /// Whether the user can create groups.
        /// </summary>
        public readonly bool CanCreateGroup;
        /// <summary>
        /// Whether the user can create projects.
        /// </summary>
        public readonly bool CanCreateProject;
        /// <summary>
        /// User's color scheme ID.
        /// </summary>
        public readonly int ColorSchemeId;
        /// <summary>
        /// Date the user was created at.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Current user's sign-in date.
        /// </summary>
        public readonly string CurrentSignInAt;
        /// <summary>
        /// The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The external UID of the user.
        /// </summary>
        public readonly string ExternUid;
        /// <summary>
        /// Whether the user is external.
        /// </summary>
        public readonly bool External;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the user is an admin.
        /// </summary>
        public readonly bool IsAdmin;
        /// <summary>
        /// Whether the user is a bot.
        /// </summary>
        public readonly bool IsBot;
        /// <summary>
        /// Last user's sign-in date.
        /// </summary>
        public readonly string LastSignInAt;
        /// <summary>
        /// LinkedIn profile of the user.
        /// </summary>
        public readonly string Linkedin;
        /// <summary>
        /// The location of the user.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
        /// </summary>
        public readonly int NamespaceId;
        /// <summary>
        /// Admin notes for this user.
        /// </summary>
        public readonly string Note;
        /// <summary>
        /// The organization of the user.
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// Number of projects the user can create.
        /// </summary>
        public readonly int ProjectsLimit;
        /// <summary>
        /// Skype username of the user.
        /// </summary>
        public readonly string Skype;
        /// <summary>
        /// Whether the user is active or blocked.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// User's theme ID.
        /// </summary>
        public readonly int ThemeId;
        /// <summary>
        /// Twitter username of the user.
        /// </summary>
        public readonly string Twitter;
        /// <summary>
        /// Whether user's two-factor auth is enabled.
        /// </summary>
        public readonly bool TwoFactorEnabled;
        /// <summary>
        /// The ID of the user.
        /// </summary>
        public readonly int UserId;
        /// <summary>
        /// The UID provider of the user.
        /// </summary>
        public readonly string UserProvider;
        /// <summary>
        /// The username of the user.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// User's website URL.
        /// </summary>
        public readonly string WebsiteUrl;

        [OutputConstructor]
        private GetUserResult(
            string avatarUrl,

            string bio,

            bool canCreateGroup,

            bool canCreateProject,

            int colorSchemeId,

            string createdAt,

            string currentSignInAt,

            string email,

            string externUid,

            bool external,

            string id,

            bool isAdmin,

            bool isBot,

            string lastSignInAt,

            string linkedin,

            string location,

            string name,

            int namespaceId,

            string note,

            string organization,

            int projectsLimit,

            string skype,

            string state,

            int themeId,

            string twitter,

            bool twoFactorEnabled,

            int userId,

            string userProvider,

            string username,

            string websiteUrl)
        {
            AvatarUrl = avatarUrl;
            Bio = bio;
            CanCreateGroup = canCreateGroup;
            CanCreateProject = canCreateProject;
            ColorSchemeId = colorSchemeId;
            CreatedAt = createdAt;
            CurrentSignInAt = currentSignInAt;
            Email = email;
            ExternUid = externUid;
            External = external;
            Id = id;
            IsAdmin = isAdmin;
            IsBot = isBot;
            LastSignInAt = lastSignInAt;
            Linkedin = linkedin;
            Location = location;
            Name = name;
            NamespaceId = namespaceId;
            Note = note;
            Organization = organization;
            ProjectsLimit = projectsLimit;
            Skype = skype;
            State = state;
            ThemeId = themeId;
            Twitter = twitter;
            TwoFactorEnabled = twoFactorEnabled;
            UserId = userId;
            UserProvider = userProvider;
            Username = username;
            WebsiteUrl = websiteUrl;
        }
    }
}
