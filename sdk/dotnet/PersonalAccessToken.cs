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
    /// The `gitlab.PersonalAccessToken` resource allows to manage the lifecycle of a personal access token for a specified user.
    /// 
    /// &gt; This resource requires administration privileges.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/personal_access_tokens.html)
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
    ///     var examplePersonalAccessToken = new GitLab.PersonalAccessToken("examplePersonalAccessToken", new()
    ///     {
    ///         UserId = 25,
    ///         ExpiresAt = "2020-03-14",
    ///         Scopes = new[]
    ///         {
    ///             "api",
    ///         },
    ///     });
    /// 
    ///     var exampleProjectVariable = new GitLab.ProjectVariable("exampleProjectVariable", new()
    ///     {
    ///         Project = gitlab_project.Example.Id,
    ///         Key = "pat",
    ///         Value = examplePersonalAccessToken.Token,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A GitLab Personal Access Token can be imported using a key composed of `&lt;user-id&gt;:&lt;token-id&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/personalAccessToken:PersonalAccessToken example "12345:1"
    /// ```
    /// 
    ///  NOTEthe `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
    /// </summary>
    [GitLabResourceType("gitlab:index/personalAccessToken:PersonalAccessToken")]
    public partial class PersonalAccessToken : global::Pulumi.CustomResource
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
        /// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

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
        /// The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a PersonalAccessToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PersonalAccessToken(string name, PersonalAccessTokenArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/personalAccessToken:PersonalAccessToken", name, args ?? new PersonalAccessTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PersonalAccessToken(string name, Input<string> id, PersonalAccessTokenState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/personalAccessToken:PersonalAccessToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PersonalAccessToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PersonalAccessToken Get(string name, Input<string> id, PersonalAccessTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new PersonalAccessToken(name, id, state, options);
        }
    }

    public sealed class PersonalAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
        /// </summary>
        [Input("expiresAt", required: true)]
        public Input<string> ExpiresAt { get; set; } = null!;

        /// <summary>
        /// The name of the personal access token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The id of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public PersonalAccessTokenArgs()
        {
        }
        public static new PersonalAccessTokenArgs Empty => new PersonalAccessTokenArgs();
    }

    public sealed class PersonalAccessTokenState : global::Pulumi.ResourceArgs
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
        /// The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

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

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
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
        /// The id of the user.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public PersonalAccessTokenState()
        {
        }
        public static new PersonalAccessTokenState Empty => new PersonalAccessTokenState();
    }
}
