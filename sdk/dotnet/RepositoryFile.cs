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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new GitLab.Group("this", new()
    ///     {
    ///         Name = "example",
    ///         Path = "example",
    ///         Description = "An example group",
    ///     });
    /// 
    ///     var thisProject = new GitLab.Project("this", new()
    ///     {
    ///         Name = "example",
    ///         NamespaceId = @this.Id,
    ///         InitializeWithReadme = true,
    ///     });
    /// 
    ///     var thisRepositoryFile = new GitLab.RepositoryFile("this", new()
    ///     {
    ///         Project = thisProject.Id,
    ///         FilePath = "meow.txt",
    ///         Branch = "main",
    ///         Content = Std.Base64encode.Invoke(new()
    ///         {
    ///             Input = "Meow goes the cat",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         AuthorEmail = "terraform@example.com",
    ///         AuthorName = "Terraform",
    ///         CommitMessage = "feature: add meow file",
    ///     });
    /// 
    ///     var readme = new GitLab.RepositoryFile("readme", new()
    ///     {
    ///         Project = thisProject.Id,
    ///         FilePath = "readme.txt",
    ///         Branch = "main",
    ///         Content = "Meow goes the cat",
    ///         AuthorEmail = "terraform@example.com",
    ///         AuthorName = "Terraform",
    ///         CommitMessage = "feature: add readme file",
    ///     });
    /// 
    ///     var readmeForDogs = new GitLab.RepositoryFile("readme_for_dogs", new()
    ///     {
    ///         Project = thisProject.Id,
    ///         FilePath = "readme.txt",
    ///         Branch = "main",
    ///         Content = "Bark goes the dog",
    ///         AuthorEmail = "terraform@example.com",
    ///         AuthorName = "Terraform",
    ///         CommitMessage = "feature: update readme file",
    ///         OverwriteOnCreate = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_repository_file`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_repository_file.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// A Repository File can be imported using an id made up of `&lt;project-id&gt;:&lt;branch-name&gt;:&lt;file-path&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/repositoryFile:RepositoryFile this 1:main:foo/bar.txt
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/repositoryFile:RepositoryFile")]
    public partial class RepositoryFile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Email of the commit author.
        /// </summary>
        [Output("authorEmail")]
        public Output<string?> AuthorEmail { get; private set; } = null!;

        /// <summary>
        /// Name of the commit author.
        /// </summary>
        [Output("authorName")]
        public Output<string?> AuthorName { get; private set; } = null!;

        /// <summary>
        /// The blob id.
        /// </summary>
        [Output("blobId")]
        public Output<string> BlobId { get; private set; } = null!;

        /// <summary>
        /// Name of the branch to which to commit to.
        /// </summary>
        [Output("branch")]
        public Output<string> Branch { get; private set; } = null!;

        /// <summary>
        /// The commit id.
        /// </summary>
        [Output("commitId")]
        public Output<string> CommitId { get; private set; } = null!;

        /// <summary>
        /// Commit message.
        /// </summary>
        [Output("commitMessage")]
        public Output<string?> CommitMessage { get; private set; } = null!;

        /// <summary>
        /// File content.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// File content sha256 digest.
        /// </summary>
        [Output("contentSha256")]
        public Output<string> ContentSha256 { get; private set; } = null!;

        /// <summary>
        /// Create commit message.
        /// </summary>
        [Output("createCommitMessage")]
        public Output<string?> CreateCommitMessage { get; private set; } = null!;

        /// <summary>
        /// Delete Commit message.
        /// </summary>
        [Output("deleteCommitMessage")]
        public Output<string?> DeleteCommitMessage { get; private set; } = null!;

        /// <summary>
        /// The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
        /// </summary>
        [Output("encoding")]
        public Output<string?> Encoding { get; private set; } = null!;

        /// <summary>
        /// Enables or disables the execute flag on the file.
        /// </summary>
        [Output("executeFilemode")]
        public Output<bool?> ExecuteFilemode { get; private set; } = null!;

        /// <summary>
        /// The filename.
        /// </summary>
        [Output("fileName")]
        public Output<string> FileName { get; private set; } = null!;

        /// <summary>
        /// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        /// </summary>
        [Output("filePath")]
        public Output<string> FilePath { get; private set; } = null!;

        /// <summary>
        /// The last known commit id.
        /// </summary>
        [Output("lastCommitId")]
        public Output<string> LastCommitId { get; private set; } = null!;

        /// <summary>
        /// Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
        /// </summary>
        [Output("overwriteOnCreate")]
        public Output<bool?> OverwriteOnCreate { get; private set; } = null!;

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of branch, tag or commit.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;

        /// <summary>
        /// The file size.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Name of the branch to start the new commit from.
        /// </summary>
        [Output("startBranch")]
        public Output<string?> StartBranch { get; private set; } = null!;

        /// <summary>
        /// Update commit message.
        /// </summary>
        [Output("updateCommitMessage")]
        public Output<string?> UpdateCommitMessage { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryFile(string name, RepositoryFileArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/repositoryFile:RepositoryFile", name, args ?? new RepositoryFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryFile(string name, Input<string> id, RepositoryFileState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/repositoryFile:RepositoryFile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryFile Get(string name, Input<string> id, RepositoryFileState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryFile(name, id, state, options);
        }
    }

    public sealed class RepositoryFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email of the commit author.
        /// </summary>
        [Input("authorEmail")]
        public Input<string>? AuthorEmail { get; set; }

        /// <summary>
        /// Name of the commit author.
        /// </summary>
        [Input("authorName")]
        public Input<string>? AuthorName { get; set; }

        /// <summary>
        /// Name of the branch to which to commit to.
        /// </summary>
        [Input("branch", required: true)]
        public Input<string> Branch { get; set; } = null!;

        /// <summary>
        /// Commit message.
        /// </summary>
        [Input("commitMessage")]
        public Input<string>? CommitMessage { get; set; }

        /// <summary>
        /// File content.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Create commit message.
        /// </summary>
        [Input("createCommitMessage")]
        public Input<string>? CreateCommitMessage { get; set; }

        /// <summary>
        /// Delete Commit message.
        /// </summary>
        [Input("deleteCommitMessage")]
        public Input<string>? DeleteCommitMessage { get; set; }

        /// <summary>
        /// The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Enables or disables the execute flag on the file.
        /// </summary>
        [Input("executeFilemode")]
        public Input<bool>? ExecuteFilemode { get; set; }

        /// <summary>
        /// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        /// </summary>
        [Input("filePath", required: true)]
        public Input<string> FilePath { get; set; } = null!;

        /// <summary>
        /// Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
        /// </summary>
        [Input("overwriteOnCreate")]
        public Input<bool>? OverwriteOnCreate { get; set; }

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Name of the branch to start the new commit from.
        /// </summary>
        [Input("startBranch")]
        public Input<string>? StartBranch { get; set; }

        /// <summary>
        /// Update commit message.
        /// </summary>
        [Input("updateCommitMessage")]
        public Input<string>? UpdateCommitMessage { get; set; }

        public RepositoryFileArgs()
        {
        }
        public static new RepositoryFileArgs Empty => new RepositoryFileArgs();
    }

    public sealed class RepositoryFileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email of the commit author.
        /// </summary>
        [Input("authorEmail")]
        public Input<string>? AuthorEmail { get; set; }

        /// <summary>
        /// Name of the commit author.
        /// </summary>
        [Input("authorName")]
        public Input<string>? AuthorName { get; set; }

        /// <summary>
        /// The blob id.
        /// </summary>
        [Input("blobId")]
        public Input<string>? BlobId { get; set; }

        /// <summary>
        /// Name of the branch to which to commit to.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// The commit id.
        /// </summary>
        [Input("commitId")]
        public Input<string>? CommitId { get; set; }

        /// <summary>
        /// Commit message.
        /// </summary>
        [Input("commitMessage")]
        public Input<string>? CommitMessage { get; set; }

        /// <summary>
        /// File content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// File content sha256 digest.
        /// </summary>
        [Input("contentSha256")]
        public Input<string>? ContentSha256 { get; set; }

        /// <summary>
        /// Create commit message.
        /// </summary>
        [Input("createCommitMessage")]
        public Input<string>? CreateCommitMessage { get; set; }

        /// <summary>
        /// Delete Commit message.
        /// </summary>
        [Input("deleteCommitMessage")]
        public Input<string>? DeleteCommitMessage { get; set; }

        /// <summary>
        /// The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Enables or disables the execute flag on the file.
        /// </summary>
        [Input("executeFilemode")]
        public Input<bool>? ExecuteFilemode { get; set; }

        /// <summary>
        /// The filename.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// The last known commit id.
        /// </summary>
        [Input("lastCommitId")]
        public Input<string>? LastCommitId { get; set; }

        /// <summary>
        /// Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
        /// </summary>
        [Input("overwriteOnCreate")]
        public Input<bool>? OverwriteOnCreate { get; set; }

        /// <summary>
        /// The name or ID of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of branch, tag or commit.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// The file size.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Name of the branch to start the new commit from.
        /// </summary>
        [Input("startBranch")]
        public Input<string>? StartBranch { get; set; }

        /// <summary>
        /// Update commit message.
        /// </summary>
        [Input("updateCommitMessage")]
        public Input<string>? UpdateCommitMessage { get; set; }

        public RepositoryFileState()
        {
        }
        public static new RepositoryFileState Empty => new RepositoryFileState();
    }
}
