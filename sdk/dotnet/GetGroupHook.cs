// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroupHook
    {
        /// <summary>
        /// The `gitlab.GroupHook` data source allows to retrieve details about a hook in a group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#get-group-hook)
        /// </summary>
        public static Task<GetGroupHookResult> InvokeAsync(GetGroupHookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupHookResult>("gitlab:index/getGroupHook:getGroupHook", args ?? new GetGroupHookArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.GroupHook` data source allows to retrieve details about a hook in a group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#get-group-hook)
        /// </summary>
        public static Output<GetGroupHookResult> Invoke(GetGroupHookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupHookResult>("gitlab:index/getGroupHook:getGroupHook", args ?? new GetGroupHookInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.GroupHook` data source allows to retrieve details about a hook in a group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#get-group-hook)
        /// </summary>
        public static Output<GetGroupHookResult> Invoke(GetGroupHookInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupHookResult>("gitlab:index/getGroupHook:getGroupHook", args ?? new GetGroupHookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupHookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID or full path of the group.
        /// </summary>
        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        /// <summary>
        /// The id of the group hook.
        /// </summary>
        [Input("hookId", required: true)]
        public int HookId { get; set; }

        public GetGroupHookArgs()
        {
        }
        public static new GetGroupHookArgs Empty => new GetGroupHookArgs();
    }

    public sealed class GetGroupHookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID or full path of the group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The id of the group hook.
        /// </summary>
        [Input("hookId", required: true)]
        public Input<int> HookId { get; set; } = null!;

        public GetGroupHookInvokeArgs()
        {
        }
        public static new GetGroupHookInvokeArgs Empty => new GetGroupHookInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupHookResult
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
        /// Set a custom webhook template.
        /// </summary>
        public readonly string CustomWebhookTemplate;
        /// <summary>
        /// Invoke the hook for deployment events.
        /// </summary>
        public readonly bool DeploymentEvents;
        /// <summary>
        /// Enable ssl verification when invoking the hook.
        /// </summary>
        public readonly bool EnableSslVerification;
        /// <summary>
        /// The ID or full path of the group.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The id of the group for the hook.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// The id of the group hook.
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
        /// Invoke the hook for subgroup events.
        /// </summary>
        public readonly bool SubgroupEvents;
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
        private GetGroupHookResult(
            bool confidentialIssuesEvents,

            bool confidentialNoteEvents,

            string customWebhookTemplate,

            bool deploymentEvents,

            bool enableSslVerification,

            string group,

            int groupId,

            int hookId,

            string id,

            bool issuesEvents,

            bool jobEvents,

            bool mergeRequestsEvents,

            bool noteEvents,

            bool pipelineEvents,

            bool pushEvents,

            string pushEventsBranchFilter,

            bool releasesEvents,

            bool subgroupEvents,

            bool tagPushEvents,

            string token,

            string url,

            bool wikiPageEvents)
        {
            ConfidentialIssuesEvents = confidentialIssuesEvents;
            ConfidentialNoteEvents = confidentialNoteEvents;
            CustomWebhookTemplate = customWebhookTemplate;
            DeploymentEvents = deploymentEvents;
            EnableSslVerification = enableSslVerification;
            Group = group;
            GroupId = groupId;
            HookId = hookId;
            Id = id;
            IssuesEvents = issuesEvents;
            JobEvents = jobEvents;
            MergeRequestsEvents = mergeRequestsEvents;
            NoteEvents = noteEvents;
            PipelineEvents = pipelineEvents;
            PushEvents = pushEvents;
            PushEventsBranchFilter = pushEventsBranchFilter;
            ReleasesEvents = releasesEvents;
            SubgroupEvents = subgroupEvents;
            TagPushEvents = tagPushEvents;
            Token = token;
            Url = url;
            WikiPageEvents = wikiPageEvents;
        }
    }
}
