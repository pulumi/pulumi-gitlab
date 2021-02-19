// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProject
    {
        /// <summary>
        /// ## # gitlab\_project
        /// 
        /// Provides details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(GitLab.GetProject.InvokeAsync(new GitLab.GetProjectArgs
        ///         {
        ///             Id = "30",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(GitLab.GetProject.InvokeAsync(new GitLab.GetProjectArgs
        ///         {
        ///             Id = "foo/bar/baz",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("gitlab:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithVersion());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project within the gitlab install.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetProjectArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// Whether the project is in read-only mode (archived).
        /// </summary>
        public readonly bool Archived;
        /// <summary>
        /// The default branch for the project.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// A description of the project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// repository via HTTP.
        /// </summary>
        public readonly string HttpUrlToRepo;
        /// <summary>
        /// Integer that uniquely identifies the project within the gitlab install.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable issue tracking for the project.
        /// </summary>
        public readonly bool IssuesEnabled;
        /// <summary>
        /// Enable LFS for the project.
        /// </summary>
        public readonly bool LfsEnabled;
        /// <summary>
        /// Enable merge requests for the project.
        /// </summary>
        public readonly bool MergeRequestsEnabled;
        public readonly string Name;
        /// <summary>
        /// The namespace (group or user) of the project. Defaults to your user.
        /// See `gitlab.Group` for an example.
        /// </summary>
        public readonly int NamespaceId;
        /// <summary>
        /// The path of the repository.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        public readonly string PathWithNamespace;
        /// <summary>
        /// Enable pipelines for the project.
        /// </summary>
        public readonly bool PipelinesEnabled;
        public readonly Outputs.GetProjectPushRulesResult PushRules;
        /// <summary>
        /// Enable `Delete source branch` option by default for all new merge requests
        /// </summary>
        public readonly bool RemoveSourceBranchAfterMerge;
        /// <summary>
        /// Allow users to request member access.
        /// </summary>
        public readonly bool RequestAccessEnabled;
        /// <summary>
        /// Registration token to use during runner setup.
        /// </summary>
        public readonly string RunnersToken;
        /// <summary>
        /// Enable snippets for the project.
        /// </summary>
        public readonly bool SnippetsEnabled;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// repository via SSH.
        /// </summary>
        public readonly string SshUrlToRepo;
        /// <summary>
        /// Repositories are created as private by default.
        /// </summary>
        public readonly string VisibilityLevel;
        /// <summary>
        /// URL that can be used to find the project in a browser.
        /// </summary>
        public readonly string WebUrl;
        /// <summary>
        /// Enable wiki for the project.
        /// </summary>
        public readonly bool WikiEnabled;

        [OutputConstructor]
        private GetProjectResult(
            bool archived,

            string defaultBranch,

            string description,

            string httpUrlToRepo,

            string id,

            bool issuesEnabled,

            bool lfsEnabled,

            bool mergeRequestsEnabled,

            string name,

            int namespaceId,

            string path,

            string pathWithNamespace,

            bool pipelinesEnabled,

            Outputs.GetProjectPushRulesResult pushRules,

            bool removeSourceBranchAfterMerge,

            bool requestAccessEnabled,

            string runnersToken,

            bool snippetsEnabled,

            string sshUrlToRepo,

            string visibilityLevel,

            string webUrl,

            bool wikiEnabled)
        {
            Archived = archived;
            DefaultBranch = defaultBranch;
            Description = description;
            HttpUrlToRepo = httpUrlToRepo;
            Id = id;
            IssuesEnabled = issuesEnabled;
            LfsEnabled = lfsEnabled;
            MergeRequestsEnabled = mergeRequestsEnabled;
            Name = name;
            NamespaceId = namespaceId;
            Path = path;
            PathWithNamespace = pathWithNamespace;
            PipelinesEnabled = pipelinesEnabled;
            PushRules = pushRules;
            RemoveSourceBranchAfterMerge = removeSourceBranchAfterMerge;
            RequestAccessEnabled = requestAccessEnabled;
            RunnersToken = runnersToken;
            SnippetsEnabled = snippetsEnabled;
            SshUrlToRepo = sshUrlToRepo;
            VisibilityLevel = visibilityLevel;
            WebUrl = webUrl;
            WikiEnabled = wikiEnabled;
        }
    }
}
