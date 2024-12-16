// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetRepositoryTree
    {
        /// <summary>
        /// The `gitlab.getRepositoryTree` data source allows details of directories and files in a repository to be retrieved.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repositories.html#list-repository-tree)
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
        ///     var @this = GitLab.GetRepositoryTree.Invoke(new()
        ///     {
        ///         Project = "example",
        ///         Ref = "main",
        ///         Path = "ExampleSubFolder",
        ///         Recursive = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryTreeResult> InvokeAsync(GetRepositoryTreeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryTreeResult>("gitlab:index/getRepositoryTree:getRepositoryTree", args ?? new GetRepositoryTreeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getRepositoryTree` data source allows details of directories and files in a repository to be retrieved.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repositories.html#list-repository-tree)
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
        ///     var @this = GitLab.GetRepositoryTree.Invoke(new()
        ///     {
        ///         Project = "example",
        ///         Ref = "main",
        ///         Path = "ExampleSubFolder",
        ///         Recursive = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryTreeResult> Invoke(GetRepositoryTreeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryTreeResult>("gitlab:index/getRepositoryTree:getRepositoryTree", args ?? new GetRepositoryTreeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getRepositoryTree` data source allows details of directories and files in a repository to be retrieved.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repositories.html#list-repository-tree)
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
        ///     var @this = GitLab.GetRepositoryTree.Invoke(new()
        ///     {
        ///         Project = "example",
        ///         Ref = "main",
        ///         Path = "ExampleSubFolder",
        ///         Recursive = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryTreeResult> Invoke(GetRepositoryTreeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryTreeResult>("gitlab:index/getRepositoryTree:getRepositoryTree", args ?? new GetRepositoryTreeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryTreeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path inside repository. Used to get content of subdirectories.
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        /// <summary>
        /// The ID or full path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// Boolean value used to get a recursive tree (false by default).
        /// </summary>
        [Input("recursive")]
        public bool? Recursive { get; set; }

        /// <summary>
        /// The name of a repository branch or tag.
        /// </summary>
        [Input("ref", required: true)]
        public string Ref { get; set; } = null!;

        public GetRepositoryTreeArgs()
        {
        }
        public static new GetRepositoryTreeArgs Empty => new GetRepositoryTreeArgs();
    }

    public sealed class GetRepositoryTreeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path inside repository. Used to get content of subdirectories.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The ID or full path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Boolean value used to get a recursive tree (false by default).
        /// </summary>
        [Input("recursive")]
        public Input<bool>? Recursive { get; set; }

        /// <summary>
        /// The name of a repository branch or tag.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        public GetRepositoryTreeInvokeArgs()
        {
        }
        public static new GetRepositoryTreeInvokeArgs Empty => new GetRepositoryTreeInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryTreeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The path inside repository. Used to get content of subdirectories.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The ID or full path of the project owned by the authenticated user.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Boolean value used to get a recursive tree (false by default).
        /// </summary>
        public readonly bool? Recursive;
        /// <summary>
        /// The name of a repository branch or tag.
        /// </summary>
        public readonly string Ref;
        /// <summary>
        /// The list of files/directories returned by the search
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryTreeTreeResult> Trees;

        [OutputConstructor]
        private GetRepositoryTreeResult(
            string id,

            string? path,

            string project,

            bool? recursive,

            string @ref,

            ImmutableArray<Outputs.GetRepositoryTreeTreeResult> trees)
        {
            Id = id;
            Path = path;
            Project = project;
            Recursive = recursive;
            Ref = @ref;
            Trees = trees;
        }
    }
}
