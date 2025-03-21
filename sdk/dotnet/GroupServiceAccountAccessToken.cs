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
    /// The `gitlab.GroupServiceAccountAccessToken` resource allows to manage the lifecycle of a group service account access token.
    /// 
    /// &gt; Use of the `timestamp()` function with expires_at will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.
    /// 
    /// &gt; Reading the access token status of a service account requires an admin token or a top-level group owner token on gitlab.com. As a result, this resource will ignore permission errors when attempting to read the token status, and will rely on the values in state instead. This can lead to apply-time failures if the token configured for the provider doesn't have permissions to rotate tokens for the service account.
    /// 
    /// &gt; Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.
    /// 
    /// &gt; Due to a limitation in the API, the `rotation_configuration` is unable to set the new expiry date before GitLab 17.9. Instead, when the resource is created, it will default the expiry date to 7 days in the future. On each subsequent apply, the new expiry will be 7 days from the date of the apply.
    /// 
    /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/group_service_accounts/#create-a-personal-access-token-for-a-service-account-user)
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_service_account_access_token`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_group_service_account_access_token.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken You can import a service account access token using `&lt;resource&gt; &lt;id&gt;`. The
    /// ```
    /// 
    /// `id` is in the form of &lt;group_id&gt;:&lt;service_account_id&gt;:&lt;access_token_id&gt;
    /// 
    /// Importing an access token does not import the access token value.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken example 1:2:3
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken")]
    public partial class GroupServiceAccountAccessToken : global::Pulumi.CustomResource
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
        /// The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name of the personal access token.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Output("revoked")]
        public Output<bool> Revoked { get; private set; } = null!;

        /// <summary>
        /// The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        /// </summary>
        [Output("rotationConfiguration")]
        public Output<Outputs.GroupServiceAccountAccessTokenRotationConfiguration?> RotationConfiguration { get; private set; } = null!;

        /// <summary>
        /// The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// The token of the group service account access token. **Note**: the token is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The ID of a service account user.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupServiceAccountAccessToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupServiceAccountAccessToken(string name, GroupServiceAccountAccessTokenArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken", name, args ?? new GroupServiceAccountAccessTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupServiceAccountAccessToken(string name, Input<string> id, GroupServiceAccountAccessTokenState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupServiceAccountAccessToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupServiceAccountAccessToken Get(string name, Input<string> id, GroupServiceAccountAccessTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupServiceAccountAccessToken(name, id, state, options);
        }
    }

    public sealed class GroupServiceAccountAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name of the personal access token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        /// </summary>
        [Input("rotationConfiguration")]
        public Input<Inputs.GroupServiceAccountAccessTokenRotationConfigurationArgs>? RotationConfiguration { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The ID of a service account user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public GroupServiceAccountAccessTokenArgs()
        {
        }
        public static new GroupServiceAccountAccessTokenArgs Empty => new GroupServiceAccountAccessTokenArgs();
    }

    public sealed class GroupServiceAccountAccessTokenState : global::Pulumi.ResourceArgs
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
        /// The service account access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the group containing the service account. Must be a top level group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the personal access token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Input("revoked")]
        public Input<bool>? Revoked { get; set; }

        /// <summary>
        /// The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        /// </summary>
        [Input("rotationConfiguration")]
        public Input<Inputs.GroupServiceAccountAccessTokenRotationConfigurationGetArgs>? RotationConfiguration { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The token of the group service account access token. **Note**: the token is not available for imported resources.
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
        /// The ID of a service account user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GroupServiceAccountAccessTokenState()
        {
        }
        public static new GroupServiceAccountAccessTokenState Empty => new GroupServiceAccountAccessTokenState();
    }
}
