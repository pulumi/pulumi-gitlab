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
    /// This resource allows you to manage Jira integration.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/service_jira.html.markdown.
    /// </summary>
    public partial class ServiceJira : Pulumi.CustomResource
    {
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Enable comments inside Jira issues on each GitLab event (commit / merge request)
        /// </summary>
        [Output("commentOnEventEnabled")]
        public Output<bool> CommentOnEventEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for commit events
        /// </summary>
        [Output("commitEvents")]
        public Output<bool> CommitEvents { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
        /// </summary>
        [Output("jiraIssueTransitionId")]
        public Output<string?> JiraIssueTransitionId { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool> MergeRequestsEvents { get; private set; } = null!;

        /// <summary>
        /// The password of the user created to be used with GitLab/JIRA.
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

        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The username of the user created to be used with GitLab/JIRA.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceJira resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceJira(string name, ServiceJiraArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceJira:ServiceJira", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class ServiceJiraArgs : Pulumi.ResourceArgs
    {
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
        /// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
        /// </summary>
        [Input("jiraIssueTransitionId")]
        public Input<string>? JiraIssueTransitionId { get; set; }

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// The password of the user created to be used with GitLab/JIRA.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

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

        /// <summary>
        /// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// The username of the user created to be used with GitLab/JIRA.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ServiceJiraArgs()
        {
        }
    }

    public sealed class ServiceJiraState : Pulumi.ResourceArgs
    {
        [Input("active")]
        public Input<bool>? Active { get; set; }

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

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration &gt; Issues &gt; Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
        /// </summary>
        [Input("jiraIssueTransitionId")]
        public Input<string>? JiraIssueTransitionId { get; set; }

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// The password of the user created to be used with GitLab/JIRA.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

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

        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The username of the user created to be used with GitLab/JIRA.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceJiraState()
        {
        }
    }
}
