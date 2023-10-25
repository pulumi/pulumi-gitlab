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
    /// The `gitlab.ServiceMicrosoftTeams` resource allows to manage the lifecycle of a project integration with Microsoft Teams.
    /// 
    /// &gt; This resource is deprecated. use `gitlab.IntegrationMicrosoftTeams`instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#microsoft-teams)
    /// 
    /// ## Import
    /// 
    /// You can import a gitlab_service_microsoft_teams state using the project ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams teams 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams")]
    public partial class ServiceMicrosoftTeams : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
        /// </summary>
        [Output("branchesToBeNotified")]
        public Output<string?> BranchesToBeNotified { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for confidential issue events
        /// </summary>
        [Output("confidentialIssuesEvents")]
        public Output<bool?> ConfidentialIssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for confidential note events
        /// </summary>
        [Output("confidentialNoteEvents")]
        public Output<bool?> ConfidentialNoteEvents { get; private set; } = null!;

        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for issue events
        /// </summary>
        [Output("issuesEvents")]
        public Output<bool?> IssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool?> MergeRequestsEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for note events
        /// </summary>
        [Output("noteEvents")]
        public Output<bool?> NoteEvents { get; private set; } = null!;

        /// <summary>
        /// Send notifications for broken pipelines
        /// </summary>
        [Output("notifyOnlyBrokenPipelines")]
        public Output<bool?> NotifyOnlyBrokenPipelines { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for pipeline events
        /// </summary>
        [Output("pipelineEvents")]
        public Output<bool?> PipelineEvents { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for push events
        /// </summary>
        [Output("pushEvents")]
        public Output<bool?> PushEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for tag push events
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool?> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// Update time.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
        /// </summary>
        [Output("webhook")]
        public Output<string> Webhook { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for wiki page events
        /// </summary>
        [Output("wikiPageEvents")]
        public Output<bool?> WikiPageEvents { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceMicrosoftTeams resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceMicrosoftTeams(string name, ServiceMicrosoftTeamsArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams", name, args ?? new ServiceMicrosoftTeamsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceMicrosoftTeams(string name, Input<string> id, ServiceMicrosoftTeamsState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceMicrosoftTeams resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceMicrosoftTeams Get(string name, Input<string> id, ServiceMicrosoftTeamsState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceMicrosoftTeams(name, id, state, options);
        }
    }

    public sealed class ServiceMicrosoftTeamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Enable notifications for confidential issue events
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for confidential note events
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// Enable notifications for issue events
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// Enable notifications for note events
        /// </summary>
        [Input("noteEvents")]
        public Input<bool>? NoteEvents { get; set; }

        /// <summary>
        /// Send notifications for broken pipelines
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// Enable notifications for pipeline events
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Enable notifications for push events
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Enable notifications for tag push events
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
        /// </summary>
        [Input("webhook", required: true)]
        public Input<string> Webhook { get; set; } = null!;

        /// <summary>
        /// Enable notifications for wiki page events
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public ServiceMicrosoftTeamsArgs()
        {
        }
        public static new ServiceMicrosoftTeamsArgs Empty => new ServiceMicrosoftTeamsArgs();
    }

    public sealed class ServiceMicrosoftTeamsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Enable notifications for confidential issue events
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for confidential note events
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Enable notifications for issue events
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for merge request events
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// Enable notifications for note events
        /// </summary>
        [Input("noteEvents")]
        public Input<bool>? NoteEvents { get; set; }

        /// <summary>
        /// Send notifications for broken pipelines
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// Enable notifications for pipeline events
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Enable notifications for push events
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Enable notifications for tag push events
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// Update time.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
        /// </summary>
        [Input("webhook")]
        public Input<string>? Webhook { get; set; }

        /// <summary>
        /// Enable notifications for wiki page events
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public ServiceMicrosoftTeamsState()
        {
        }
        public static new ServiceMicrosoftTeamsState Empty => new ServiceMicrosoftTeamsState();
    }
}
