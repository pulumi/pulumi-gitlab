// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectHooks
    {
        /// <summary>
        /// The `gitlab.getProjectHooks` data source allows to retrieve details about hooks in a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-project-hooks)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var examples = GitLab.GetProjectHooks.Invoke(new()
        ///     {
        ///         Project = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectHooksResult> InvokeAsync(GetProjectHooksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectHooksResult>("gitlab:index/getProjectHooks:getProjectHooks", args ?? new GetProjectHooksArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectHooks` data source allows to retrieve details about hooks in a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-project-hooks)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = GitLab.GetProject.Invoke(new()
        ///     {
        ///         Id = "foo/bar/baz",
        ///     });
        /// 
        ///     var examples = GitLab.GetProjectHooks.Invoke(new()
        ///     {
        ///         Project = example.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectHooksResult> Invoke(GetProjectHooksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectHooksResult>("gitlab:index/getProjectHooks:getProjectHooks", args ?? new GetProjectHooksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectHooksArgs : global::Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetProjectHooksArgs()
        {
        }
        public static new GetProjectHooksArgs Empty => new GetProjectHooksArgs();
    }

    public sealed class GetProjectHooksInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetProjectHooksInvokeArgs()
        {
        }
        public static new GetProjectHooksInvokeArgs Empty => new GetProjectHooksInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectHooksResult
    {
        /// <summary>
        /// The list of hooks.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectHooksHookResult> Hooks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name or id of the project.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private GetProjectHooksResult(
            ImmutableArray<Outputs.GetProjectHooksHookResult> hooks,

            string id,

            string project)
        {
            Hooks = hooks;
            Id = id;
            Project = project;
        }
    }
}
