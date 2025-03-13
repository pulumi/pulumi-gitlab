// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectEnvironments
    {
        /// <summary>
        /// The `gitlab.getProjectEnvironments` data source retrieves information about all environments of the given project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/environments/#list-environments)
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
        ///     var thisProject = new GitLab.Project("this", new()
        ///     {
        ///         Name = "example",
        ///         InitializeWithReadme = true,
        ///     });
        /// 
        ///     var @this = GitLab.GetProjectEnvironments.Invoke(new()
        ///     {
        ///         Project = thisProject.PathWithNamespace,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProjectEnvironmentsResult> InvokeAsync(GetProjectEnvironmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectEnvironmentsResult>("gitlab:index/getProjectEnvironments:getProjectEnvironments", args ?? new GetProjectEnvironmentsArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectEnvironments` data source retrieves information about all environments of the given project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/environments/#list-environments)
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
        ///     var thisProject = new GitLab.Project("this", new()
        ///     {
        ///         Name = "example",
        ///         InitializeWithReadme = true,
        ///     });
        /// 
        ///     var @this = GitLab.GetProjectEnvironments.Invoke(new()
        ///     {
        ///         Project = thisProject.PathWithNamespace,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProjectEnvironmentsResult> Invoke(GetProjectEnvironmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectEnvironmentsResult>("gitlab:index/getProjectEnvironments:getProjectEnvironments", args ?? new GetProjectEnvironmentsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectEnvironments` data source retrieves information about all environments of the given project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/environments/#list-environments)
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
        ///     var thisProject = new GitLab.Project("this", new()
        ///     {
        ///         Name = "example",
        ///         InitializeWithReadme = true,
        ///     });
        /// 
        ///     var @this = GitLab.GetProjectEnvironments.Invoke(new()
        ///     {
        ///         Project = thisProject.PathWithNamespace,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProjectEnvironmentsResult> Invoke(GetProjectEnvironmentsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectEnvironmentsResult>("gitlab:index/getProjectEnvironments:getProjectEnvironments", args ?? new GetProjectEnvironmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectEnvironmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return the environment with this name. Mutually exclusive with search.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// Return list of environments matching the search criteria. Mutually exclusive with name. Must be at least 3 characters long.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// List all environments that match the specified state. Valid values are `available`, `stopping`, `stopped`. Returns all environments if not set.
        /// </summary>
        [Input("states")]
        public string? States { get; set; }

        public GetProjectEnvironmentsArgs()
        {
        }
        public static new GetProjectEnvironmentsArgs Empty => new GetProjectEnvironmentsArgs();
    }

    public sealed class GetProjectEnvironmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return the environment with this name. Mutually exclusive with search.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Return list of environments matching the search criteria. Mutually exclusive with name. Must be at least 3 characters long.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// List all environments that match the specified state. Valid values are `available`, `stopping`, `stopped`. Returns all environments if not set.
        /// </summary>
        [Input("states")]
        public Input<string>? States { get; set; }

        public GetProjectEnvironmentsInvokeArgs()
        {
        }
        public static new GetProjectEnvironmentsInvokeArgs Empty => new GetProjectEnvironmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectEnvironmentsResult
    {
        /// <summary>
        /// The list of environments.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectEnvironmentsEnvironmentResult> Environments;
        public readonly string Id;
        /// <summary>
        /// Return the environment with this name. Mutually exclusive with search.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Return list of environments matching the search criteria. Mutually exclusive with name. Must be at least 3 characters long.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// List all environments that match the specified state. Valid values are `available`, `stopping`, `stopped`. Returns all environments if not set.
        /// </summary>
        public readonly string? States;

        [OutputConstructor]
        private GetProjectEnvironmentsResult(
            ImmutableArray<Outputs.GetProjectEnvironmentsEnvironmentResult> environments,

            string id,

            string? name,

            string project,

            string? search,

            string? states)
        {
            Environments = environments;
            Id = id;
            Name = name;
            Project = project;
            Search = search;
            States = states;
        }
    }
}
