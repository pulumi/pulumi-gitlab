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
    /// The `gitlab.ProjectHook` resource allows to manage the lifecycle of a project hook.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#hooks)
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
    ///     var example = new GitLab.ProjectHook("example", new()
    ///     {
    ///         MergeRequestsEvents = true,
    ///         Project = "example/hooked",
    ///         Url = "https://example.com/hook/example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A GitLab Project Hook can be imported using a key composed of `&lt;project-id&gt;:&lt;hook-id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectHook:ProjectHook example "12345:1"
    /// ```
    /// 
    /// NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
    /// </summary>
    [GitLabResourceType("gitlab:index/projectHook:ProjectHook")]
    public partial class ProjectHook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Invoke the hook for confidential issues events.
        /// </summary>
        [Output("confidentialIssuesEvents")]
        public Output<bool?> ConfidentialIssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for confidential notes events.
        /// </summary>
        [Output("confidentialNoteEvents")]
        public Output<bool?> ConfidentialNoteEvents { get; private set; } = null!;

        /// <summary>
        /// Set a custom webhook template.
        /// </summary>
        [Output("customWebhookTemplate")]
        public Output<string?> CustomWebhookTemplate { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for deployment events.
        /// </summary>
        [Output("deploymentEvents")]
        public Output<bool?> DeploymentEvents { get; private set; } = null!;

        /// <summary>
        /// Enable ssl verification when invoking the hook.
        /// </summary>
        [Output("enableSslVerification")]
        public Output<bool?> EnableSslVerification { get; private set; } = null!;

        /// <summary>
        /// The id of the project hook.
        /// </summary>
        [Output("hookId")]
        public Output<int> HookId { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for issues events.
        /// </summary>
        [Output("issuesEvents")]
        public Output<bool?> IssuesEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for job events.
        /// </summary>
        [Output("jobEvents")]
        public Output<bool?> JobEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for merge requests.
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool?> MergeRequestsEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for notes events.
        /// </summary>
        [Output("noteEvents")]
        public Output<bool?> NoteEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for pipeline events.
        /// </summary>
        [Output("pipelineEvents")]
        public Output<bool?> PipelineEvents { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The id of the project for the hook.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for push events.
        /// </summary>
        [Output("pushEvents")]
        public Output<bool?> PushEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for push events on matching branches only.
        /// </summary>
        [Output("pushEventsBranchFilter")]
        public Output<string?> PushEventsBranchFilter { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for releases events.
        /// </summary>
        [Output("releasesEvents")]
        public Output<bool?> ReleasesEvents { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for tag push events.
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool?> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// A token to present when invoking the hook. The token is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// The url of the hook to invoke. Forces re-creation to preserve `token`.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Invoke the hook for wiki page events.
        /// </summary>
        [Output("wikiPageEvents")]
        public Output<bool?> WikiPageEvents { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectHook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectHook(string name, ProjectHookArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectHook:ProjectHook", name, args ?? new ProjectHookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectHook(string name, Input<string> id, ProjectHookState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectHook:ProjectHook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectHook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectHook Get(string name, Input<string> id, ProjectHookState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectHook(name, id, state, options);
        }
    }

    public sealed class ProjectHookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Invoke the hook for confidential issues events.
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// Invoke the hook for confidential notes events.
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// Set a custom webhook template.
        /// </summary>
        [Input("customWebhookTemplate")]
        public Input<string>? CustomWebhookTemplate { get; set; }

        /// <summary>
        /// Invoke the hook for deployment events.
        /// </summary>
        [Input("deploymentEvents")]
        public Input<bool>? DeploymentEvents { get; set; }

        /// <summary>
        /// Enable ssl verification when invoking the hook.
        /// </summary>
        [Input("enableSslVerification")]
        public Input<bool>? EnableSslVerification { get; set; }

        /// <summary>
        /// Invoke the hook for issues events.
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// Invoke the hook for job events.
        /// </summary>
        [Input("jobEvents")]
        public Input<bool>? JobEvents { get; set; }

        /// <summary>
        /// Invoke the hook for merge requests.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// Invoke the hook for notes events.
        /// </summary>
        [Input("noteEvents")]
        public Input<bool>? NoteEvents { get; set; }

        /// <summary>
        /// Invoke the hook for pipeline events.
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Invoke the hook for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Invoke the hook for push events on matching branches only.
        /// </summary>
        [Input("pushEventsBranchFilter")]
        public Input<string>? PushEventsBranchFilter { get; set; }

        /// <summary>
        /// Invoke the hook for releases events.
        /// </summary>
        [Input("releasesEvents")]
        public Input<bool>? ReleasesEvents { get; set; }

        /// <summary>
        /// Invoke the hook for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// A token to present when invoking the hook. The token is not available for imported resources.
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
        /// The url of the hook to invoke. Forces re-creation to preserve `token`.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Invoke the hook for wiki page events.
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public ProjectHookArgs()
        {
        }
        public static new ProjectHookArgs Empty => new ProjectHookArgs();
    }

    public sealed class ProjectHookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Invoke the hook for confidential issues events.
        /// </summary>
        [Input("confidentialIssuesEvents")]
        public Input<bool>? ConfidentialIssuesEvents { get; set; }

        /// <summary>
        /// Invoke the hook for confidential notes events.
        /// </summary>
        [Input("confidentialNoteEvents")]
        public Input<bool>? ConfidentialNoteEvents { get; set; }

        /// <summary>
        /// Set a custom webhook template.
        /// </summary>
        [Input("customWebhookTemplate")]
        public Input<string>? CustomWebhookTemplate { get; set; }

        /// <summary>
        /// Invoke the hook for deployment events.
        /// </summary>
        [Input("deploymentEvents")]
        public Input<bool>? DeploymentEvents { get; set; }

        /// <summary>
        /// Enable ssl verification when invoking the hook.
        /// </summary>
        [Input("enableSslVerification")]
        public Input<bool>? EnableSslVerification { get; set; }

        /// <summary>
        /// The id of the project hook.
        /// </summary>
        [Input("hookId")]
        public Input<int>? HookId { get; set; }

        /// <summary>
        /// Invoke the hook for issues events.
        /// </summary>
        [Input("issuesEvents")]
        public Input<bool>? IssuesEvents { get; set; }

        /// <summary>
        /// Invoke the hook for job events.
        /// </summary>
        [Input("jobEvents")]
        public Input<bool>? JobEvents { get; set; }

        /// <summary>
        /// Invoke the hook for merge requests.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// Invoke the hook for notes events.
        /// </summary>
        [Input("noteEvents")]
        public Input<bool>? NoteEvents { get; set; }

        /// <summary>
        /// Invoke the hook for pipeline events.
        /// </summary>
        [Input("pipelineEvents")]
        public Input<bool>? PipelineEvents { get; set; }

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The id of the project for the hook.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Invoke the hook for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Invoke the hook for push events on matching branches only.
        /// </summary>
        [Input("pushEventsBranchFilter")]
        public Input<string>? PushEventsBranchFilter { get; set; }

        /// <summary>
        /// Invoke the hook for releases events.
        /// </summary>
        [Input("releasesEvents")]
        public Input<bool>? ReleasesEvents { get; set; }

        /// <summary>
        /// Invoke the hook for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// A token to present when invoking the hook. The token is not available for imported resources.
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
        /// The url of the hook to invoke. Forces re-creation to preserve `token`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Invoke the hook for wiki page events.
        /// </summary>
        [Input("wikiPageEvents")]
        public Input<bool>? WikiPageEvents { get; set; }

        public ProjectHookState()
        {
        }
        public static new ProjectHookState Empty => new ProjectHookState();
    }
}
