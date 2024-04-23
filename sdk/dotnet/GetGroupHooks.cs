// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroupHooks
    {
        /// <summary>
        /// The `gitlab.getGroupHooks` data source allows to retrieve details about hooks in a group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-group-hooks)
        /// </summary>
        public static Task<GetGroupHooksResult> InvokeAsync(GetGroupHooksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupHooksResult>("gitlab:index/getGroupHooks:getGroupHooks", args ?? new GetGroupHooksArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroupHooks` data source allows to retrieve details about hooks in a group.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-group-hooks)
        /// </summary>
        public static Output<GetGroupHooksResult> Invoke(GetGroupHooksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupHooksResult>("gitlab:index/getGroupHooks:getGroupHooks", args ?? new GetGroupHooksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupHooksArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID or full path of the group.
        /// </summary>
        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        public GetGroupHooksArgs()
        {
        }
        public static new GetGroupHooksArgs Empty => new GetGroupHooksArgs();
    }

    public sealed class GetGroupHooksInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID or full path of the group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        public GetGroupHooksInvokeArgs()
        {
        }
        public static new GetGroupHooksInvokeArgs Empty => new GetGroupHooksInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupHooksResult
    {
        /// <summary>
        /// The ID or full path of the group.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The list of hooks.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupHooksHookResult> Hooks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupHooksResult(
            string group,

            ImmutableArray<Outputs.GetGroupHooksHookResult> hooks,

            string id)
        {
            Group = group;
            Hooks = hooks;
            Id = id;
        }
    }
}
