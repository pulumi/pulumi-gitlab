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
    /// The `gitlab.IntegrationGithub` resource allows to manage the lifecycle of a project integration with GitHub.
    /// 
    /// &gt; This resource requires a GitLab Enterprise instance.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#github)
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
    ///     var awesomeProject = new GitLab.Project("awesome_project", new()
    ///     {
    ///         Name = "awesome_project",
    ///         Description = "My awesome project.",
    ///         VisibilityLevel = "public",
    ///     });
    /// 
    ///     var github = new GitLab.IntegrationGithub("github", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Token = "REDACTED",
    ///         RepositoryUrl = "https://github.com/gitlabhq/terraform-provider-gitlab",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_github`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_integration_github.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `&lt;resource&gt; &lt;project_id&gt;`:
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/integrationGithub:IntegrationGithub")]
    public partial class IntegrationGithub : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("repositoryUrl")]
        public Output<string> RepositoryUrl { get; private set; } = null!;

        /// <summary>
        /// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        /// </summary>
        [Output("staticContext")]
        public Output<bool?> StaticContext { get; private set; } = null!;

        /// <summary>
        /// Title.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// A GitHub personal access token with at least `repo:status` scope.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// Update time.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationGithub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationGithub(string name, IntegrationGithubArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationGithub:IntegrationGithub", name, args ?? new IntegrationGithubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationGithub(string name, Input<string> id, IntegrationGithubState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationGithub:IntegrationGithub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationGithub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationGithub Get(string name, Input<string> id, IntegrationGithubState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationGithub(name, id, state, options);
        }
    }

    public sealed class IntegrationGithubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("repositoryUrl", required: true)]
        public Input<string> RepositoryUrl { get; set; } = null!;

        /// <summary>
        /// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        /// </summary>
        [Input("staticContext")]
        public Input<bool>? StaticContext { get; set; }

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// A GitHub personal access token with at least `repo:status` scope.
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

        public IntegrationGithubArgs()
        {
        }
        public static new IntegrationGithubArgs Empty => new IntegrationGithubArgs();
    }

    public sealed class IntegrationGithubState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("repositoryUrl")]
        public Input<string>? RepositoryUrl { get; set; }

        /// <summary>
        /// Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
        /// </summary>
        [Input("staticContext")]
        public Input<bool>? StaticContext { get; set; }

        /// <summary>
        /// Title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// A GitHub personal access token with at least `repo:status` scope.
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
        /// Update time.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public IntegrationGithubState()
        {
        }
        public static new IntegrationGithubState Empty => new IntegrationGithubState();
    }
}
