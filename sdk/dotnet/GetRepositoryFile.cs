// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetRepositoryFile
    {
        /// <summary>
        /// The `gitlab.RepositoryFile` data source allows details of a file in a repository to be retrieved.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/repository_files/)
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
        ///     var example = GitLab.GetRepositoryFile.Invoke(new()
        ///     {
        ///         Project = "example",
        ///         Ref = "main",
        ///         FilePath = "README.md",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryFileResult> InvokeAsync(GetRepositoryFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryFileResult>("gitlab:index/getRepositoryFile:getRepositoryFile", args ?? new GetRepositoryFileArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.RepositoryFile` data source allows details of a file in a repository to be retrieved.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/repository_files/)
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
        ///     var example = GitLab.GetRepositoryFile.Invoke(new()
        ///     {
        ///         Project = "example",
        ///         Ref = "main",
        ///         FilePath = "README.md",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryFileResult> Invoke(GetRepositoryFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryFileResult>("gitlab:index/getRepositoryFile:getRepositoryFile", args ?? new GetRepositoryFileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.RepositoryFile` data source allows details of a file in a repository to be retrieved.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/repository_files/)
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
        ///     var example = GitLab.GetRepositoryFile.Invoke(new()
        ///     {
        ///         Project = "example",
        ///         Ref = "main",
        ///         FilePath = "README.md",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryFileResult> Invoke(GetRepositoryFileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryFileResult>("gitlab:index/getRepositoryFile:getRepositoryFile", args ?? new GetRepositoryFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryFileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        /// </summary>
        [Input("filePath", required: true)]
        public string FilePath { get; set; } = null!;

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        /// <summary>
        /// The name of branch, tag or commit.
        /// </summary>
        [Input("ref", required: true)]
        public string Ref { get; set; } = null!;

        public GetRepositoryFileArgs()
        {
        }
        public static new GetRepositoryFileArgs Empty => new GetRepositoryFileArgs();
    }

    public sealed class GetRepositoryFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        /// </summary>
        [Input("filePath", required: true)]
        public Input<string> FilePath { get; set; } = null!;

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The name of branch, tag or commit.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        public GetRepositoryFileInvokeArgs()
        {
        }
        public static new GetRepositoryFileInvokeArgs Empty => new GetRepositoryFileInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryFileResult
    {
        /// <summary>
        /// The blob id.
        /// </summary>
        public readonly string BlobId;
        /// <summary>
        /// The commit id.
        /// </summary>
        public readonly string CommitId;
        /// <summary>
        /// File content.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// File content sha256 digest.
        /// </summary>
        public readonly string ContentSha256;
        /// <summary>
        /// The file content encoding.
        /// </summary>
        public readonly string Encoding;
        /// <summary>
        /// Enables or disables the execute flag on the file.
        /// </summary>
        public readonly bool ExecuteFilemode;
        /// <summary>
        /// The filename.
        /// </summary>
        public readonly string FileName;
        /// <summary>
        /// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        /// </summary>
        public readonly string FilePath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last known commit id.
        /// </summary>
        public readonly string LastCommitId;
        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The name of branch, tag or commit.
        /// </summary>
        public readonly string Ref;
        /// <summary>
        /// The file size.
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private GetRepositoryFileResult(
            string blobId,

            string commitId,

            string content,

            string contentSha256,

            string encoding,

            bool executeFilemode,

            string fileName,

            string filePath,

            string id,

            string lastCommitId,

            string project,

            string @ref,

            int size)
        {
            BlobId = blobId;
            CommitId = commitId;
            Content = content;
            ContentSha256 = contentSha256;
            Encoding = encoding;
            ExecuteFilemode = executeFilemode;
            FileName = fileName;
            FilePath = filePath;
            Id = id;
            LastCommitId = lastCommitId;
            Project = project;
            Ref = @ref;
            Size = size;
        }
    }
}
