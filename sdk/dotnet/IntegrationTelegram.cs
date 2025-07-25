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
    /// The `gitlab.IntegrationTelegram` resource manages the lifecycle of a project integration with Telegram.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#telegram)
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
    ///     var @default = new GitLab.IntegrationTelegram("default", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Token = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
    ///         Room = "-1000000000000000",
    ///         NotifyOnlyBrokenPipelines = false,
    ///         BranchesToBeNotified = "all",
    ///         PushEvents = false,
    ///         IssuesEvents = false,
    ///         ConfidentialIssuesEvents = false,
    ///         MergeRequestsEvents = false,
    ///         TagPushEvents = false,
    ///         NoteEvents = false,
    ///         ConfidentialNoteEvents = false,
    ///         PipelineEvents = false,
    ///         WikiPageEvents = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_integration_telegram`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_integration_telegram.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Importing using the CLI is supported with the following syntax:
    /// 
    /// You can import a gitlab_integration_telegram state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/integrationTelegram:IntegrationTelegram default 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/integrationTelegram:IntegrationTelegram")]
    public partial class IntegrationTelegram : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`.
        /// </summary>
        [Output("branchesToBeNotified")]
        public Output<string> BranchesToBeNotified { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for confidential issues events.
        /// </summary>
        [Output("confidentialIssuesEvents")]
        public Output<bool> ConfidentialIssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for confidential note events.
        /// </summary>
        [Output("confidentialNoteEvents")]
        public Output<bool> ConfidentialNoteEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for issues events.
        /// </summary>
        [Output("issuesEvents")]
        public Output<bool> IssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for merge requests events.
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool> MergeRequestsEvents { get; private set; } = null!;

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
        /// Enable notifications for pipeline events.
        /// </summary>
        [Output("pipelineEvents")]
        public Output<bool> PipelineEvents { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project to integrate with Telegram.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Output("pushEvents")]
        public Output<bool> PushEvents { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
        /// </summary>
        [Output("room")]
        public Output<string> Room { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// The Telegram bot token.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for wiki page events.
        /// </summary>
        [Output("wikiPageEvents")]
        public Output<bool> WikiPageEvents { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationTelegram resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationTelegram(string name, IntegrationTelegramArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationTelegram:IntegrationTelegram", name, args ?? new IntegrationTelegramArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationTelegram(string name, Input<string> id, IntegrationTelegramState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationTelegram:IntegrationTelegram", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationTelegram resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationTelegram Get(string name, Input<string> id, IntegrationTelegramState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationTelegram(name, id, state, options);
        }
    }

    public sealed class IntegrationTelegramArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`.
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Enable notifications for confidential issues events.
        /// </summary>
        [Input("confidentialIssuesEvents", required: true)]
        public Input<bool> ConfidentialIssuesEvents { get; set; } = null!;

        /// <summary>
        /// Enable notifications for confidential note events.
        /// </summary>
        [Input("confidentialNoteEvents", required: true)]
        public Input<bool> ConfidentialNoteEvents { get; set; } = null!;

        /// <summary>
        /// Enable notifications for issues events.
        /// </summary>
        [Input("issuesEvents", required: true)]
        public Input<bool> IssuesEvents { get; set; } = null!;

        /// <summary>
        /// Enable notifications for merge requests events.
        /// </summary>
        [Input("mergeRequestsEvents", required: true)]
        public Input<bool> MergeRequestsEvents { get; set; } = null!;

        /// <summary>
        /// Enable notifications for note events.
        /// </summary>
        [Input("noteEvents", required: true)]
        public Input<bool> NoteEvents { get; set; } = null!;

        /// <summary>
        /// Send notifications for broken pipelines.
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// Enable notifications for pipeline events.
        /// </summary>
        [Input("pipelineEvents", required: true)]
        public Input<bool> PipelineEvents { get; set; } = null!;

        /// <summary>
        /// The ID or full path of the project to integrate with Telegram.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents", required: true)]
        public Input<bool> PushEvents { get; set; } = null!;

        /// <summary>
        /// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
        /// </summary>
        [Input("room", required: true)]
        public Input<string> Room { get; set; } = null!;

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents", required: true)]
        public Input<bool> TagPushEvents { get; set; } = null!;

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// The Telegram bot token.
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
        /// Enable notifications for wiki page events.
        /// </summary>
        [Input("wikiPageEvents", required: true)]
        public Input<bool> WikiPageEvents { get; set; } = null!;

        public IntegrationTelegramArgs()
        {
        }
        public static new IntegrationTelegramArgs Empty => new IntegrationTelegramArgs();
    }

    public sealed class IntegrationTelegramState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`.
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Enable notifications for confidential issues events.
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for confidential note events.
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// Enable notifications for issues events.
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// Enable notifications for merge requests events.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

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
        /// Enable notifications for pipeline events.
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// The ID or full path of the project to integrate with Telegram.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
        /// </summary>
        [Input("room")]
        public Input<string>? Room { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The Telegram bot token.
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
        /// Enable notifications for wiki page events.
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public IntegrationTelegramState()
        {
        }
        public static new IntegrationTelegramState Empty => new IntegrationTelegramState();
    }
}
