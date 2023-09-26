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
    /// The `gitlab_group_access`token resource allows to manage the lifecycle of a group access token.
    /// 
    /// &gt; Group Access Token were introduced in GitLab 14.7
    /// 
    /// **Upstream API**: [GitLab REST API](https://docs.gitlab.com/ee/api/group_access_tokens.html)
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
    ///     var exampleGroupAccessToken = new GitLab.GroupAccessToken("exampleGroupAccessToken", new()
    ///     {
    ///         Group = "25",
    ///         ExpiresAt = "2020-03-14",
    ///         AccessLevel = "developer",
    ///         Scopes = new[]
    ///         {
    ///             "api",
    ///         },
    ///     });
    /// 
    ///     var exampleGroupVariable = new GitLab.GroupVariable("exampleGroupVariable", new()
    ///     {
    ///         Group = "25",
    ///         Key = "gat",
    ///         Value = exampleGroupAccessToken.Token,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A GitLab Group Access Token can be imported using a key composed of `&lt;group-id&gt;:&lt;token-id&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/groupAccessToken:GroupAccessToken example "12345:1"
    /// ```
    /// 
    ///  ATTENTIONthe `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
    /// </summary>
    [GitLabResourceType("gitlab:index/groupAccessToken:GroupAccessToken")]
    public partial class GroupAccessToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string?> AccessLevel { get; private set; } = null!;

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
        /// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ID or path of the group to add the group access token to.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name of the group access token.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Output("revoked")]
        public Output<bool> Revoked { get; private set; } = null!;

        /// <summary>
        /// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The user id associated to the token.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupAccessToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupAccessToken(string name, GroupAccessTokenArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupAccessToken:GroupAccessToken", name, args ?? new GroupAccessTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupAccessToken(string name, Input<string> id, GroupAccessTokenState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupAccessToken:GroupAccessToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupAccessToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupAccessToken Get(string name, Input<string> id, GroupAccessTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupAccessToken(name, id, state, options);
        }
    }

    public sealed class GroupAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
        /// </summary>
        [Input("expiresAt", required: true)]
        public Input<string> ExpiresAt { get; set; } = null!;

        /// <summary>
        /// The ID or path of the group to add the group access token to.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name of the group access token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public GroupAccessTokenArgs()
        {
        }
        public static new GroupAccessTokenArgs Empty => new GroupAccessTokenArgs();
    }

    public sealed class GroupAccessTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

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
        /// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID or path of the group to add the group access token to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the group access token.
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
        /// The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
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
        /// The user id associated to the token.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GroupAccessTokenState()
        {
        }
        public static new GroupAccessTokenState Empty => new GroupAccessTokenState();
    }
}
