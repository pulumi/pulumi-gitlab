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
    /// The `gitlab.ProjectAccessToken` resource allows to manage the lifecycle of a project access token.
    /// 
    /// &gt; Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/ee/operations/tracing.html)
    /// 
    /// &gt; Use `rotation_configuration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.
    /// 
    /// &gt; Due to [Automatic reuse detection](https://docs.gitlab.com/ee/api/project_access_tokens.html#automatic-reuse-detection) it's possible that a new Project Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.
    /// 
    /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/project_access_tokens.html)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProjectAccessToken = new GitLab.ProjectAccessToken("exampleProjectAccessToken", new()
    ///     {
    ///         Project = "25",
    ///         ExpiresAt = "2020-03-14",
    ///         AccessLevel = "reporter",
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
    ///         Value = exampleProjectAccessToken.Token,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// A GitLab Project Access Token can be imported using a key composed of `&lt;project-id&gt;:&lt;token-id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectAccessToken:ProjectAccessToken example "12345:1"
    /// ```
    /// 
    /// NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
    /// </summary>
    [GitLabResourceType("gitlab:index/projectAccessToken:ProjectAccessToken")]
    public partial class ProjectAccessToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

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
        /// When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The name of the project access token.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Output("revoked")]
        public Output<bool> Revoked { get; private set; } = null!;

        /// <summary>
        /// The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        /// </summary>
        [Output("rotationConfiguration")]
        public Output<Outputs.ProjectAccessTokenRotationConfiguration?> RotationConfiguration { get; private set; } = null!;

        /// <summary>
        /// The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// The token of the project access token. **Note**: the token is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The user_id associated to the token.
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectAccessToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectAccessToken(string name, ProjectAccessTokenArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectAccessToken:ProjectAccessToken", name, args ?? new ProjectAccessTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectAccessToken(string name, Input<string> id, ProjectAccessTokenState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectAccessToken:ProjectAccessToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectAccessToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectAccessToken Get(string name, Input<string> id, ProjectAccessTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectAccessToken(name, id, state, options);
        }
    }

    public sealed class ProjectAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The name of the project access token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        /// </summary>
        [Input("rotationConfiguration")]
        public Input<Inputs.ProjectAccessTokenRotationConfigurationArgs>? RotationConfiguration { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public ProjectAccessTokenArgs()
        {
        }
        public static new ProjectAccessTokenArgs Empty => new ProjectAccessTokenArgs();
    }

    public sealed class ProjectAccessTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
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
        /// When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The name of the project access token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// True if the token is revoked.
        /// </summary>
        [Input("revoked")]
        public Input<bool>? Revoked { get; set; }

        /// <summary>
        /// The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
        /// </summary>
        [Input("rotationConfiguration")]
        public Input<Inputs.ProjectAccessTokenRotationConfigurationGetArgs>? RotationConfiguration { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The token of the project access token. **Note**: the token is not available for imported resources.
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
        /// The user_id associated to the token.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public ProjectAccessTokenState()
        {
        }
        public static new ProjectAccessTokenState Empty => new ProjectAccessTokenState();
    }
}
