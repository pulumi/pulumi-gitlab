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
    /// The `gitlab.IntegrationJenkins` resource allows to manage the lifecycle of a project integration with Jenkins.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#jenkins)
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
    ///     var jenkins = new GitLab.IntegrationJenkins("jenkins", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         JenkinsUrl = "http://jenkins.example.com",
    ///         ProjectName = "my_project_name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/integrationJenkins:IntegrationJenkins You can import a gitlab_integration_jenkins state using `&lt;resource&gt; &lt;project_id&gt;`:
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/integrationJenkins:IntegrationJenkins jenkins 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/integrationJenkins:IntegrationJenkins")]
    public partial class IntegrationJenkins : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Enable SSL verification. Defaults to `true` (enabled).
        /// </summary>
        [Output("enableSslVerification")]
        public Output<bool> EnableSslVerification { get; private set; } = null!;

        /// <summary>
        /// Jenkins URL like `http://jenkins.example.com`
        /// </summary>
        [Output("jenkinsUrl")]
        public Output<string> JenkinsUrl { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for merge request events.
        /// </summary>
        [Output("mergeRequestEvents")]
        public Output<bool> MergeRequestEvents { get; private set; } = null!;

        /// <summary>
        /// Password for authentication with the Jenkins server, if authentication is required by the server.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URL-friendly project name. Example: `my_project_name`.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Output("pushEvents")]
        public Output<bool> PushEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// Username for authentication with the Jenkins server, if authentication is required by the server.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationJenkins resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationJenkins(string name, IntegrationJenkinsArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationJenkins:IntegrationJenkins", name, args ?? new IntegrationJenkinsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationJenkins(string name, Input<string> id, IntegrationJenkinsState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationJenkins:IntegrationJenkins", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationJenkins resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationJenkins Get(string name, Input<string> id, IntegrationJenkinsState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationJenkins(name, id, state, options);
        }
    }

    public sealed class IntegrationJenkinsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable SSL verification. Defaults to `true` (enabled).
        /// </summary>
        [Input("enableSslVerification")]
        public Input<bool>? EnableSslVerification { get; set; }

        /// <summary>
        /// Jenkins URL like `http://jenkins.example.com`
        /// </summary>
        [Input("jenkinsUrl", required: true)]
        public Input<string> JenkinsUrl { get; set; } = null!;

        /// <summary>
        /// Enable notifications for merge request events.
        /// </summary>
        [Input("mergeRequestEvents")]
        public Input<bool>? MergeRequestEvents { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for authentication with the Jenkins server, if authentication is required by the server.
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
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The URL-friendly project name. Example: `my_project_name`.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// Username for authentication with the Jenkins server, if authentication is required by the server.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationJenkinsArgs()
        {
        }
        public static new IntegrationJenkinsArgs Empty => new IntegrationJenkinsArgs();
    }

    public sealed class IntegrationJenkinsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Enable SSL verification. Defaults to `true` (enabled).
        /// </summary>
        [Input("enableSslVerification")]
        public Input<bool>? EnableSslVerification { get; set; }

        /// <summary>
        /// Jenkins URL like `http://jenkins.example.com`
        /// </summary>
        [Input("jenkinsUrl")]
        public Input<string>? JenkinsUrl { get; set; }

        /// <summary>
        /// Enable notifications for merge request events.
        /// </summary>
        [Input("mergeRequestEvents")]
        public Input<bool>? MergeRequestEvents { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for authentication with the Jenkins server, if authentication is required by the server.
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
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URL-friendly project name. Example: `my_project_name`.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// Username for authentication with the Jenkins server, if authentication is required by the server.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationJenkinsState()
        {
        }
        public static new IntegrationJenkinsState Empty => new IntegrationJenkinsState();
    }
}
