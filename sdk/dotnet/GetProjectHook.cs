// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectHook
    {
        /// <summary>
        /// The `gitlab.ProjectHook` data source allows to retrieve details about a hook in a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-project-hook)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var exampleProjectHook = GitLab.GetProjectHook.Invoke(new()
        ///     {
        ///         Project = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         HookId = 1,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectHookResult> InvokeAsync(GetProjectHookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectHookResult>("gitlab:index/getProjectHook:getProjectHook", args ?? new GetProjectHookArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.ProjectHook` data source allows to retrieve details about a hook in a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-project-hook)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var exampleProjectHook = GitLab.GetProjectHook.Invoke(new()
        ///     {
        ///         Project = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         HookId = 1,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectHookResult> Invoke(GetProjectHookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectHookResult>("gitlab:index/getProjectHook:getProjectHook", args ?? new GetProjectHookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectHookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the project hook.
        /// </summary>
        [Input("hookId", required: true)]
        public int HookId { get; set; }

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetProjectHookArgs()
        {
        }
        public static new GetProjectHookArgs Empty => new GetProjectHookArgs();
    }

    public sealed class GetProjectHookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the project hook.
        /// </summary>
        [Input("hookId", required: true)]
        public Input<int> HookId { get; set; } = null!;

        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetProjectHookInvokeArgs()
        {
        }
        public static new GetProjectHookInvokeArgs Empty => new GetProjectHookInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectHookResult
    {
        /// <summary>
        /// Invoke the hook for confidential issues events.
        /// </summary>
        public readonly bool ConfidentialIssuesEvents;
        /// <summary>
        /// Invoke the hook for confidential notes events.
        /// </summary>
        public readonly bool ConfidentialNoteEvents;
        /// <summary>
        /// Invoke the hook for deployment events.
        /// </summary>
        public readonly bool DeploymentEvents;
        /// <summary>
        /// Enable ssl verification when invoking the hook.
        /// </summary>
        public readonly bool EnableSslVerification;
        /// <summary>
        /// The id of the project hook.
        /// </summary>
        public readonly int HookId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Invoke the hook for issues events.
        /// </summary>
        public readonly bool IssuesEvents;
        /// <summary>
        /// Invoke the hook for job events.
        /// </summary>
        public readonly bool JobEvents;
        /// <summary>
        /// Invoke the hook for merge requests.
        /// </summary>
        public readonly bool MergeRequestsEvents;
        /// <summary>
        /// Invoke the hook for notes events.
        /// </summary>
        public readonly bool NoteEvents;
        /// <summary>
        /// Invoke the hook for pipeline events.
        /// </summary>
        public readonly bool PipelineEvents;
        /// <summary>
        /// The name or id of the project to add the hook to.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The id of the project for the hook.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Invoke the hook for push events.
        /// </summary>
        public readonly bool PushEvents;
        /// <summary>
        /// Invoke the hook for push events on matching branches only.
        /// </summary>
        public readonly string PushEventsBranchFilter;
        /// <summary>
        /// Invoke the hook for releases events.
        /// </summary>
        public readonly bool ReleasesEvents;
        /// <summary>
        /// Invoke the hook for tag push events.
        /// </summary>
        public readonly bool TagPushEvents;
        /// <summary>
        /// A token to present when invoking the hook. The token is not available for imported resources.
        /// </summary>
        public readonly string Token;
        /// <summary>
        /// The url of the hook to invoke.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Invoke the hook for wiki page events.
        /// </summary>
        public readonly bool WikiPageEvents;

        [OutputConstructor]
        private GetProjectHookResult(
            bool confidentialIssuesEvents,

            bool confidentialNoteEvents,

            bool deploymentEvents,

            bool enableSslVerification,

            int hookId,

            string id,

            bool issuesEvents,

            bool jobEvents,

            bool mergeRequestsEvents,

            bool noteEvents,

            bool pipelineEvents,

            string project,

            int projectId,

            bool pushEvents,

            string pushEventsBranchFilter,

            bool releasesEvents,

            bool tagPushEvents,

            string token,

            string url,

            bool wikiPageEvents)
        {
            ConfidentialIssuesEvents = confidentialIssuesEvents;
            ConfidentialNoteEvents = confidentialNoteEvents;
            DeploymentEvents = deploymentEvents;
            EnableSslVerification = enableSslVerification;
            HookId = hookId;
            Id = id;
            IssuesEvents = issuesEvents;
            JobEvents = jobEvents;
            MergeRequestsEvents = mergeRequestsEvents;
            NoteEvents = noteEvents;
            PipelineEvents = pipelineEvents;
            Project = project;
            ProjectId = projectId;
            PushEvents = pushEvents;
            PushEventsBranchFilter = pushEventsBranchFilter;
            ReleasesEvents = releasesEvents;
            TagPushEvents = tagPushEvents;
            Token = token;
            Url = url;
            WikiPageEvents = wikiPageEvents;
        }
    }
}