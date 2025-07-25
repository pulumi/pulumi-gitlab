// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a project for the branch to use
    ///     var example = new GitLab.Project("example", new()
    ///     {
    ///         Name = "example",
    ///         Description = "An example project",
    ///         NamespaceId = exampleGitlabGroup.Id,
    ///     });
    /// 
    ///     var exampleBranch = new GitLab.Branch("example", new()
    ///     {
    ///         Name = "example",
    ///         Ref = "main",
    ///         Project = example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_branch`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_branch.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Importing using the CLI is supported with the following syntax:
    /// 
    /// Gitlab branches can be imported with a key composed of `&lt;project_id&gt;:&lt;branch_name&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/branch:Branch example "12345:develop"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/branch:Branch")]
    public partial class Branch : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bool, true if you can push to the branch.
        /// </summary>
        [Output("canPush")]
        public Output<bool> CanPush { get; private set; } = null!;

        /// <summary>
        /// The commit associated with the branch ref.
        /// </summary>
        [Output("commits")]
        public Output<ImmutableArray<Outputs.BranchCommit>> Commits { get; private set; } = null!;

        /// <summary>
        /// Bool, true if branch is the default branch for the project.
        /// </summary>
        [Output("default")]
        public Output<bool> Default { get; private set; } = null!;

        /// <summary>
        /// Bool, true if developer level access allows to merge branch.
        /// </summary>
        [Output("developerCanMerge")]
        public Output<bool> DeveloperCanMerge { get; private set; } = null!;

        /// <summary>
        /// Bool, true if developer level access allows git push.
        /// </summary>
        [Output("developerCanPush")]
        public Output<bool> DeveloperCanPush { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the branch is kept once the resource destroyed (must be applied before a destroy).
        /// </summary>
        [Output("keepOnDestroy")]
        public Output<bool?> KeepOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Bool, true if the branch has been merged into its parent.
        /// </summary>
        [Output("merged")]
        public Output<bool> Merged { get; private set; } = null!;

        /// <summary>
        /// The name for this branch.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project which the branch is created against.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Bool, true if branch has branch protection.
        /// </summary>
        [Output("protected")]
        public Output<bool> Protected { get; private set; } = null!;

        /// <summary>
        /// The ref which the branch is created from.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;

        /// <summary>
        /// The url of the created branch (https).
        /// </summary>
        [Output("webUrl")]
        public Output<string> WebUrl { get; private set; } = null!;


        /// <summary>
        /// Create a Branch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Branch(string name, BranchArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/branch:Branch", name, args ?? new BranchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Branch(string name, Input<string> id, BranchState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/branch:Branch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Branch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Branch Get(string name, Input<string> id, BranchState? state = null, CustomResourceOptions? options = null)
        {
            return new Branch(name, id, state, options);
        }
    }

    public sealed class BranchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the branch is kept once the resource destroyed (must be applied before a destroy).
        /// </summary>
        [Input("keepOnDestroy")]
        public Input<bool>? KeepOnDestroy { get; set; }

        /// <summary>
        /// The name for this branch.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project which the branch is created against.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The ref which the branch is created from.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        public BranchArgs()
        {
        }
        public static new BranchArgs Empty => new BranchArgs();
    }

    public sealed class BranchState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bool, true if you can push to the branch.
        /// </summary>
        [Input("canPush")]
        public Input<bool>? CanPush { get; set; }

        [Input("commits")]
        private InputList<Inputs.BranchCommitGetArgs>? _commits;

        /// <summary>
        /// The commit associated with the branch ref.
        /// </summary>
        public InputList<Inputs.BranchCommitGetArgs> Commits
        {
            get => _commits ?? (_commits = new InputList<Inputs.BranchCommitGetArgs>());
            set => _commits = value;
        }

        /// <summary>
        /// Bool, true if branch is the default branch for the project.
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        /// <summary>
        /// Bool, true if developer level access allows to merge branch.
        /// </summary>
        [Input("developerCanMerge")]
        public Input<bool>? DeveloperCanMerge { get; set; }

        /// <summary>
        /// Bool, true if developer level access allows git push.
        /// </summary>
        [Input("developerCanPush")]
        public Input<bool>? DeveloperCanPush { get; set; }

        /// <summary>
        /// Indicates whether the branch is kept once the resource destroyed (must be applied before a destroy).
        /// </summary>
        [Input("keepOnDestroy")]
        public Input<bool>? KeepOnDestroy { get; set; }

        /// <summary>
        /// Bool, true if the branch has been merged into its parent.
        /// </summary>
        [Input("merged")]
        public Input<bool>? Merged { get; set; }

        /// <summary>
        /// The name for this branch.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project which the branch is created against.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Bool, true if branch has branch protection.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// The ref which the branch is created from.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// The url of the created branch (https).
        /// </summary>
        [Input("webUrl")]
        public Input<string>? WebUrl { get; set; }

        public BranchState()
        {
        }
        public static new BranchState Empty => new BranchState();
    }
}
