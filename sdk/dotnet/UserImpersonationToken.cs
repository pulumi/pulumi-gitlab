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
    /// The `gitlab.UserImpersonationToken` resource allows to manage impersonation tokens of users.
    /// Requires administrator access. Token values are returned once. You are only able to create impersonation tokens to impersonate the user and perform both API calls and Git reads and writes. The user can’t see these tokens in their profile settings page.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#create-an-impersonation-token)
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
    ///     var @this = new GitLab.UserImpersonationToken("this", new()
    ///     {
    ///         UserId = 12345,
    ///         Name = "token_name",
    ///         Scopes = new[]
    ///         {
    ///             "api",
    ///         },
    ///         ExpiresAt = "2024-08-27",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_user_impersonation_token`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_user_impersonation_token.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// A GitLab User Impersonation Token can be imported using a key composed of `&lt;user-id&gt;:&lt;token-id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/userImpersonationToken:UserImpersonationToken example "12345:1"
    /// ```
    /// 
    /// NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
    /// </summary>
    [GitLabResourceType("gitlab:index/userImpersonationToken:UserImpersonationToken")]
    public partial class UserImpersonationToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// True if the token is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Time the token has been created, RFC3339 format.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Expiration date of the impersonation token in ISO format (YYYY-MM-DD).
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// True as the token is always an impersonation token.
        /// </summary>
        [Output("impersonation")]
        public Output<bool> Impersonation { get; private set; } = null!;

        /// <summary>
        /// The name of the impersonation token.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Output("revoked")]
        public Output<bool> Revoked { get; private set; } = null!;

        /// <summary>
        /// Array of scopes of the impersonation token. valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_service_ping`
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// The token of the user impersonation token. **Note**: the token is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// ID of the impersonation token.
        /// </summary>
        [Output("tokenId")]
        public Output<int> TokenId { get; private set; } = null!;

        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserImpersonationToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserImpersonationToken(string name, UserImpersonationTokenArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/userImpersonationToken:UserImpersonationToken", name, args ?? new UserImpersonationTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserImpersonationToken(string name, Input<string> id, UserImpersonationTokenState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/userImpersonationToken:UserImpersonationToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserImpersonationToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserImpersonationToken Get(string name, Input<string> id, UserImpersonationTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new UserImpersonationToken(name, id, state, options);
        }
    }

    public sealed class UserImpersonationTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration date of the impersonation token in ISO format (YYYY-MM-DD).
        /// </summary>
        [Input("expiresAt", required: true)]
        public Input<string> ExpiresAt { get; set; } = null!;

        /// <summary>
        /// The name of the impersonation token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// Array of scopes of the impersonation token. valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_service_ping`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public UserImpersonationTokenArgs()
        {
        }
        public static new UserImpersonationTokenArgs Empty => new UserImpersonationTokenArgs();
    }

    public sealed class UserImpersonationTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if the token is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Time the token has been created, RFC3339 format.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Expiration date of the impersonation token in ISO format (YYYY-MM-DD).
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// True as the token is always an impersonation token.
        /// </summary>
        [Input("impersonation")]
        public Input<bool>? Impersonation { get; set; }

        /// <summary>
        /// The name of the impersonation token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Input("revoked")]
        public Input<bool>? Revoked { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Array of scopes of the impersonation token. valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_service_ping`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The token of the user impersonation token. **Note**: the token is not available for imported resources.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// ID of the impersonation token.
        /// </summary>
        [Input("tokenId")]
        public Input<int>? TokenId { get; set; }

        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public UserImpersonationTokenState()
        {
        }
        public static new UserImpersonationTokenState Empty => new UserImpersonationTokenState();
    }
}
