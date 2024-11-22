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
    /// The `gitlab.ServiceJira` resource allows to manage the lifecycle of a project integration with Jira.
    /// 
    /// &gt; This resource is deprecated. use `gitlab.IntegrationJira`instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#jira)
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
    ///     var jira = new GitLab.ServiceJira("jira", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Url = "https://jira.example.com",
    ///         Username = "user",
    ///         Password = "mypass",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_jira`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_service_jira.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// You can import a gitlab_service_jira state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/serviceJira:ServiceJira jira 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/serviceJira:ServiceJira")]
    public partial class ServiceJira : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
        /// </summary>
        [Output("apiUrl")]
        public Output<string> ApiUrl { get; private set; } = null!;

        /// <summary>
        /// Enable comments inside Jira issues on each GitLab event (commit / merge request)
        /// </summary>
        [Output("commentOnEventEnabled")]
        public Output<bool?> CommentOnEventEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for commit events
        /// </summary>
        [Output("commitEvents")]
        public Output<bool> CommitEvents { get; private set; } = null!;

        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Enable viewing Jira issues in GitLab.
        /// </summary>
        [Output("issuesEnabled")]
        public Output<bool?> IssuesEnabled { get; private set; } = null!;

        /// <summary>
        /// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
        /// </summary>
        [Output("jiraAuthType")]
        public Output<int?> JiraAuthType { get; private set; } = null!;

        /// <summary>
        /// Prefix to match Jira issue keys.
        /// </summary>
        [Output("jiraIssuePrefix")]
        public Output<string?> JiraIssuePrefix { get; private set; } = null!;

        /// <summary>
        /// Regular expression to match Jira issue keys.
        /// </summary>
        [Output("jiraIssueRegex")]
        public Output<string?> JiraIssueRegex { get; private set; } = null!;

        /// <summary>
        /// Enable automatic issue transitions. Takes precedence over jira*issue*transition_id if enabled. Defaults to false.
        /// </summary>
        [Output("jiraIssueTransitionAutomatic")]
        public Output<bool?> JiraIssueTransitionAutomatic { get; private set; } = null!;

        /// <summary>
        /// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
        /// </summary>
        [Output("jiraIssueTransitionId")]
        public Output<string?> JiraIssueTransitionId { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool> MergeRequestsEvents { get; private set; } = null!;

        /// <summary>
        /// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
        /// </summary>
        [Output("projectKeys")]
        public Output<ImmutableArray<string>> ProjectKeys { get; private set; } = null!;

        /// <summary>
        /// Title.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Update time.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not to inherit default settings. Defaults to false.
        /// </summary>
        [Output("useInheritedSettings")]
        public Output<bool?> UseInheritedSettings { get; private set; } = null!;

        /// <summary>
        /// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceJira resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceJira(string name, ServiceJiraArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceJira:ServiceJira", name, args ?? new ServiceJiraArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceJira(string name, Input<string> id, ServiceJiraState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceJira:ServiceJira", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceJira resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceJira Get(string name, Input<string> id, ServiceJiraState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceJira(name, id, state, options);
        }
    }

    public sealed class ServiceJiraArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
        /// </summary>
        [Input("apiUrl")]
        public Input<string>? ApiUrl { get; set; }

        /// <summary>
        /// Enable comments inside Jira issues on each GitLab event (commit / merge request)
        /// </summary>
        [Input("commentOnEventEnabled")]
        public Input<bool>? CommentOnEventEnabled { get; set; }

        /// <summary>
        /// Enable notifications for commit events
        /// </summary>
        [Input("commitEvents")]
        public Input<bool>? CommitEvents { get; set; }

        /// <summary>
        /// Enable viewing Jira issues in GitLab.
        /// </summary>
        [Input("issuesEnabled")]
        public Input<bool>? IssuesEnabled { get; set; }

        /// <summary>
        /// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
        /// </summary>
        [Input("jiraAuthType")]
        public Input<int>? JiraAuthType { get; set; }

        /// <summary>
        /// Prefix to match Jira issue keys.
        /// </summary>
        [Input("jiraIssuePrefix")]
        public Input<string>? JiraIssuePrefix { get; set; }

        /// <summary>
        /// Regular expression to match Jira issue keys.
        /// </summary>
        [Input("jiraIssueRegex")]
        public Input<string>? JiraIssueRegex { get; set; }

        /// <summary>
        /// Enable automatic issue transitions. Takes precedence over jira*issue*transition_id if enabled. Defaults to false.
        /// </summary>
        [Input("jiraIssueTransitionAutomatic")]
        public Input<bool>? JiraIssueTransitionAutomatic { get; set; }

        /// <summary>
        /// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
        /// </summary>
        [Input("jiraIssueTransitionId")]
        public Input<string>? JiraIssueTransitionId { get; set; }

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
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
        /// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("projectKeys")]
        private InputList<string>? _projectKeys;

        /// <summary>
        /// Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
        /// </summary>
        public InputList<string> ProjectKeys
        {
            get => _projectKeys ?? (_projectKeys = new InputList<string>());
            set => _projectKeys = value;
        }

        /// <summary>
        /// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Indicates whether or not to inherit default settings. Defaults to false.
        /// </summary>
        [Input("useInheritedSettings")]
        public Input<bool>? UseInheritedSettings { get; set; }

        /// <summary>
        /// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceJiraArgs()
        {
        }
        public static new ServiceJiraArgs Empty => new ServiceJiraArgs();
    }

    public sealed class ServiceJiraState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
        /// </summary>
        [Input("apiUrl")]
        public Input<string>? ApiUrl { get; set; }

        /// <summary>
        /// Enable comments inside Jira issues on each GitLab event (commit / merge request)
        /// </summary>
        [Input("commentOnEventEnabled")]
        public Input<bool>? CommentOnEventEnabled { get; set; }

        /// <summary>
        /// Enable notifications for commit events
        /// </summary>
        [Input("commitEvents")]
        public Input<bool>? CommitEvents { get; set; }

        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Enable viewing Jira issues in GitLab.
        /// </summary>
        [Input("issuesEnabled")]
        public Input<bool>? IssuesEnabled { get; set; }

        /// <summary>
        /// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
        /// </summary>
        [Input("jiraAuthType")]
        public Input<int>? JiraAuthType { get; set; }

        /// <summary>
        /// Prefix to match Jira issue keys.
        /// </summary>
        [Input("jiraIssuePrefix")]
        public Input<string>? JiraIssuePrefix { get; set; }

        /// <summary>
        /// Regular expression to match Jira issue keys.
        /// </summary>
        [Input("jiraIssueRegex")]
        public Input<string>? JiraIssueRegex { get; set; }

        /// <summary>
        /// Enable automatic issue transitions. Takes precedence over jira*issue*transition_id if enabled. Defaults to false.
        /// </summary>
        [Input("jiraIssueTransitionAutomatic")]
        public Input<bool>? JiraIssueTransitionAutomatic { get; set; }

        /// <summary>
        /// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
        /// </summary>
        [Input("jiraIssueTransitionId")]
        public Input<string>? JiraIssueTransitionId { get; set; }

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
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
        /// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("projectKeys")]
        private InputList<string>? _projectKeys;

        /// <summary>
        /// Keys of Jira projects. When issues_enabled is true, this setting specifies which Jira projects to view issues from in GitLab.
        /// </summary>
        public InputList<string> ProjectKeys
        {
            get => _projectKeys ?? (_projectKeys = new InputList<string>());
            set => _projectKeys = value;
        }

        /// <summary>
        /// Title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Update time.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Indicates whether or not to inherit default settings. Defaults to false.
        /// </summary>
        [Input("useInheritedSettings")]
        public Input<bool>? UseInheritedSettings { get; set; }

        /// <summary>
        /// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceJiraState()
        {
        }
        public static new ServiceJiraState Empty => new ServiceJiraState();
    }
}
