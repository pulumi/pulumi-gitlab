// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// The `gitlab.UserSshKey` resource allows to manage the lifecycle of an SSH key assigned to a user.
    /// 
    /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/users.html#single-ssh-key)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleUser = GitLab.GetUser.Invoke(new()
    ///     {
    ///         Username = "example-user",
    ///     });
    /// 
    ///     var exampleUserSshKey = new GitLab.UserSshKey("exampleUserSshKey", new()
    ///     {
    ///         UserId = exampleUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///         Title = "example-key",
    ///         Key = "ssh-ed25519 AAAA...",
    ///         ExpiresAt = "2016-01-21T00:00:00.000Z",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import a user ssh key using an id made up of `{user-id}:{key}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/userSshKey:UserSshKey example 42:1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/userSshKey:UserSshKey")]
    public partial class UserSshKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time when this key was created in GitLab.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The ID of the ssh key.
        /// </summary>
        [Output("keyId")]
        public Output<int> KeyId { get; private set; } = null!;

        /// <summary>
        /// The title of the ssh key.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// The ID or username of the user.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserSshKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserSshKey(string name, UserSshKeyArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/userSshKey:UserSshKey", name, args ?? new UserSshKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserSshKey(string name, Input<string> id, UserSshKeyState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/userSshKey:UserSshKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserSshKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserSshKey Get(string name, Input<string> id, UserSshKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new UserSshKey(name, id, state, options);
        }
    }

    public sealed class UserSshKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The title of the ssh key.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The ID or username of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public UserSshKeyArgs()
        {
        }
        public static new UserSshKeyArgs Empty => new UserSshKeyArgs();
    }

    public sealed class UserSshKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when this key was created in GitLab.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The ID of the ssh key.
        /// </summary>
        [Input("keyId")]
        public Input<int>? KeyId { get; set; }

        /// <summary>
        /// The title of the ssh key.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The ID or username of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public UserSshKeyState()
        {
        }
        public static new UserSshKeyState Empty => new UserSshKeyState();
    }
}
