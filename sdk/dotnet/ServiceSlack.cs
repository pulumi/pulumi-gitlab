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
    /// The `gitlab.ServiceSlack` resource allows to manage the lifecycle of a project integration with Slack.
    /// 
    /// &gt; This resource is deprecated. use `gitlab.IntegrationSlack`instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#slack-notifications)
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
    ///     var slack = new GitLab.ServiceSlack("slack", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Webhook = "https://webhook.com",
    ///         Username = "myuser",
    ///         PushEvents = true,
    ///         PushChannel = "push_chan",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import a gitlab_service_slack.slack state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/serviceSlack:ServiceSlack email 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/serviceSlack:ServiceSlack")]
    public partial class ServiceSlack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
        /// </summary>
        [Output("branchesToBeNotified")]
        public Output<string> BranchesToBeNotified { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive confidential issue events notifications.
        /// </summary>
        [Output("confidentialIssueChannel")]
        public Output<string?> ConfidentialIssueChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for confidential issues events.
        /// </summary>
        [Output("confidentialIssuesEvents")]
        public Output<bool> ConfidentialIssuesEvents { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive confidential note events notifications.
        /// </summary>
        [Output("confidentialNoteChannel")]
        public Output<string> ConfidentialNoteChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for confidential note events.
        /// </summary>
        [Output("confidentialNoteEvents")]
        public Output<bool> ConfidentialNoteEvents { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive issue events notifications.
        /// </summary>
        [Output("issueChannel")]
        public Output<string?> IssueChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for issues events.
        /// </summary>
        [Output("issuesEvents")]
        public Output<bool> IssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to [https://github.com/xanzy/go-gitlab/issues/1354](https://github.com/xanzy/go-gitlab/issues/1354).
        /// </summary>
        [Output("jobEvents")]
        public Output<bool> JobEvents { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive merge request events notifications.
        /// </summary>
        [Output("mergeRequestChannel")]
        public Output<string?> MergeRequestChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for merge requests events.
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool> MergeRequestsEvents { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive note events notifications.
        /// </summary>
        [Output("noteChannel")]
        public Output<string?> NoteChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for note events.
        /// </summary>
        [Output("noteEvents")]
        public Output<bool> NoteEvents { get; private set; } = null!;

        /// <summary>
        /// Send notifications for broken pipelines.
        /// </summary>
        [Output("notifyOnlyBrokenPipelines")]
        public Output<bool> NotifyOnlyBrokenPipelines { get; private set; } = null!;

        /// <summary>
        /// This parameter has been replaced with `branches_to_be_notified`.
        /// </summary>
        [Output("notifyOnlyDefaultBranch")]
        public Output<bool> NotifyOnlyDefaultBranch { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive pipeline events notifications.
        /// </summary>
        [Output("pipelineChannel")]
        public Output<string?> PipelineChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for pipeline events.
        /// </summary>
        [Output("pipelineEvents")]
        public Output<bool> PipelineEvents { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive push events notifications.
        /// </summary>
        [Output("pushChannel")]
        public Output<string?> PushChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Output("pushEvents")]
        public Output<bool> PushEvents { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive tag push events notifications.
        /// </summary>
        [Output("tagPushChannel")]
        public Output<string?> TagPushChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// Username to use.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
        /// </summary>
        [Output("webhook")]
        public Output<string> Webhook { get; private set; } = null!;

        /// <summary>
        /// The name of the channel to receive wiki page events notifications.
        /// </summary>
        [Output("wikiPageChannel")]
        public Output<string?> WikiPageChannel { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for wiki page events.
        /// </summary>
        [Output("wikiPageEvents")]
        public Output<bool> WikiPageEvents { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceSlack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceSlack(string name, ServiceSlackArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceSlack:ServiceSlack", name, args ?? new ServiceSlackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceSlack(string name, Input<string> id, ServiceSlackState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceSlack:ServiceSlack", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceSlack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceSlack Get(string name, Input<string> id, ServiceSlackState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceSlack(name, id, state, options);
        }
    }

    public sealed class ServiceSlackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// The name of the channel to receive confidential issue events notifications.
        /// </summary>
        [Input("confidentialIssueChannel")]
        public Input<string>? ConfidentialIssueChannel { get; set; }

        /// <summary>
        /// Enable notifications for confidential issues events.
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive confidential note events notifications.
        /// </summary>
        [Input("confidentialNoteChannel")]
        public Input<string>? ConfidentialNoteChannel { get; set; }

        /// <summary>
        /// Enable notifications for confidential note events.
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive issue events notifications.
        /// </summary>
        [Input("issueChannel")]
        public Input<string>? IssueChannel { get; set; }

        /// <summary>
        /// Enable notifications for issues events.
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive merge request events notifications.
        /// </summary>
        [Input("mergeRequestChannel")]
        public Input<string>? MergeRequestChannel { get; set; }

        /// <summary>
        /// Enable notifications for merge requests events.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive note events notifications.
        /// </summary>
        [Input("noteChannel")]
        public Input<string>? NoteChannel { get; set; }

        /// <summary>
        /// Enable notifications for note events.
        /// </summary>
        [Input("noteEvents")]
        public Input<bool>? NoteEvents { get; set; }

        /// <summary>
        /// Send notifications for broken pipelines.
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// This parameter has been replaced with `branches_to_be_notified`.
        /// </summary>
        [Input("notifyOnlyDefaultBranch")]
        public Input<bool>? NotifyOnlyDefaultBranch { get; set; }

        /// <summary>
        /// The name of the channel to receive pipeline events notifications.
        /// </summary>
        [Input("pipelineChannel")]
        public Input<string>? PipelineChannel { get; set; }

        /// <summary>
        /// Enable notifications for pipeline events.
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The name of the channel to receive push events notifications.
        /// </summary>
        [Input("pushChannel")]
        public Input<string>? PushChannel { get; set; }

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive tag push events notifications.
        /// </summary>
        [Input("tagPushChannel")]
        public Input<string>? TagPushChannel { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// Username to use.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
        /// </summary>
        [Input("webhook", required: true)]
        public Input<string> Webhook { get; set; } = null!;

        /// <summary>
        /// The name of the channel to receive wiki page events notifications.
        /// </summary>
        [Input("wikiPageChannel")]
        public Input<string>? WikiPageChannel { get; set; }

        /// <summary>
        /// Enable notifications for wiki page events.
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public ServiceSlackArgs()
        {
        }
        public static new ServiceSlackArgs Empty => new ServiceSlackArgs();
    }

    public sealed class ServiceSlackState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// The name of the channel to receive confidential issue events notifications.
        /// </summary>
        [Input("confidentialIssueChannel")]
        public Input<string>? ConfidentialIssueChannel { get; set; }

        /// <summary>
        /// Enable notifications for confidential issues events.
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive confidential note events notifications.
        /// </summary>
        [Input("confidentialNoteChannel")]
        public Input<string>? ConfidentialNoteChannel { get; set; }

        /// <summary>
        /// Enable notifications for confidential note events.
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive issue events notifications.
        /// </summary>
        [Input("issueChannel")]
        public Input<string>? IssueChannel { get; set; }

        /// <summary>
        /// Enable notifications for issues events.
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to [https://github.com/xanzy/go-gitlab/issues/1354](https://github.com/xanzy/go-gitlab/issues/1354).
        /// </summary>
        [Input("jobEvents")]
        public Input<bool>? JobEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive merge request events notifications.
        /// </summary>
        [Input("mergeRequestChannel")]
        public Input<string>? MergeRequestChannel { get; set; }

        /// <summary>
        /// Enable notifications for merge requests events.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive note events notifications.
        /// </summary>
        [Input("noteChannel")]
        public Input<string>? NoteChannel { get; set; }

        /// <summary>
        /// Enable notifications for note events.
        /// </summary>
        [Input("noteEvents")]
        public Input<bool>? NoteEvents { get; set; }

        /// <summary>
        /// Send notifications for broken pipelines.
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// This parameter has been replaced with `branches_to_be_notified`.
        /// </summary>
        [Input("notifyOnlyDefaultBranch")]
        public Input<bool>? NotifyOnlyDefaultBranch { get; set; }

        /// <summary>
        /// The name of the channel to receive pipeline events notifications.
        /// </summary>
        [Input("pipelineChannel")]
        public Input<string>? PipelineChannel { get; set; }

        /// <summary>
        /// Enable notifications for pipeline events.
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the channel to receive push events notifications.
        /// </summary>
        [Input("pushChannel")]
        public Input<string>? PushChannel { get; set; }

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// The name of the channel to receive tag push events notifications.
        /// </summary>
        [Input("tagPushChannel")]
        public Input<string>? TagPushChannel { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// Username to use.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Webhook URL (Example, [https://hooks.slack.com/services/...)](https://hooks.slack.com/services/...)). This value cannot be imported.
        /// </summary>
        [Input("webhook")]
        public Input<string>? Webhook { get; set; }

        /// <summary>
        /// The name of the channel to receive wiki page events notifications.
        /// </summary>
        [Input("wikiPageChannel")]
        public Input<string>? WikiPageChannel { get; set; }

        /// <summary>
        /// Enable notifications for wiki page events.
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public ServiceSlackState()
        {
        }
        public static new ServiceSlackState Empty => new ServiceSlackState();
    }
}
