// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// The `gitlab.ProjectIntegrationHarbor` resource manages the lifecycle of a project integration with Harbor.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#harbor)
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
    ///     var harbor = new GitLab.ProjectIntegrationHarbor("harbor", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Url = "http://harbor.example.com",
    ///         ProjectName = "my_project_name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_integration_harbor`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_project_integration_harbor.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Importing using the CLI is supported with the following syntax:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectIntegrationHarbor:ProjectIntegrationHarbor You can import a gitlab_project_integration_harbor state using `&lt;resource&gt; &lt;project_id&gt;`:
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectIntegrationHarbor:ProjectIntegrationHarbor harbor 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectIntegrationHarbor:ProjectIntegrationHarbor")]
    public partial class ProjectIntegrationHarbor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Password for authentication with the Harbor server, if authentication is required by the server.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// ID of the GitLab project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URL-friendly Harbor project name. This project needs to already exist in Harbor. Example: `my_project_name`.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Harbor URL. Example: `http://harbor.example.com`
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not to inherit default settings. Defaults to false.
        /// </summary>
        [Output("useInheritedSettings")]
        public Output<bool> UseInheritedSettings { get; private set; } = null!;

        /// <summary>
        /// Username for authentication with the Harbor server, if authentication is required by the server.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectIntegrationHarbor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectIntegrationHarbor(string name, ProjectIntegrationHarborArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectIntegrationHarbor:ProjectIntegrationHarbor", name, args ?? new ProjectIntegrationHarborArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectIntegrationHarbor(string name, Input<string> id, ProjectIntegrationHarborState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectIntegrationHarbor:ProjectIntegrationHarbor", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProjectIntegrationHarbor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectIntegrationHarbor Get(string name, Input<string> id, ProjectIntegrationHarborState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectIntegrationHarbor(name, id, state, options);
        }
    }

    public sealed class ProjectIntegrationHarborArgs : global::Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password for authentication with the Harbor server, if authentication is required by the server.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// ID of the GitLab project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The URL-friendly Harbor project name. This project needs to already exist in Harbor. Example: `my_project_name`.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Harbor URL. Example: `http://harbor.example.com`
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Indicates whether or not to inherit default settings. Defaults to false.
        /// </summary>
        [Input("useInheritedSettings")]
        public Input<bool>? UseInheritedSettings { get; set; }

        /// <summary>
        /// Username for authentication with the Harbor server, if authentication is required by the server.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ProjectIntegrationHarborArgs()
        {
        }
        public static new ProjectIntegrationHarborArgs Empty => new ProjectIntegrationHarborArgs();
    }

    public sealed class ProjectIntegrationHarborState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for authentication with the Harbor server, if authentication is required by the server.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// ID of the GitLab project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URL-friendly Harbor project name. This project needs to already exist in Harbor. Example: `my_project_name`.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Harbor URL. Example: `http://harbor.example.com`
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Indicates whether or not to inherit default settings. Defaults to false.
        /// </summary>
        [Input("useInheritedSettings")]
        public Input<bool>? UseInheritedSettings { get; set; }

        /// <summary>
        /// Username for authentication with the Harbor server, if authentication is required by the server.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProjectIntegrationHarborState()
        {
        }
        public static new ProjectIntegrationHarborState Empty => new ProjectIntegrationHarborState();
    }
}
