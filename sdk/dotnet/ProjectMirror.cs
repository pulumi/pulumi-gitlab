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
    /// ## # gitlab\_project_mirror
    /// 
    /// This resource allows you to add a mirror target for the repository, all changes will be synced to the remote target.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new GitLab.ProjectMirror("foo", new GitLab.ProjectMirrorArgs
    ///         {
    ///             Project = "1",
    ///             Url = "https://username:password@github.com/org/repository.git",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectMirror:ProjectMirror foo "12345:1337"
    /// ```
    /// </summary>
    public partial class ProjectMirror : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines if the mirror is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Determines if divergent refs are skipped.
        /// </summary>
        [Output("keepDivergentRefs")]
        public Output<bool?> KeepDivergentRefs { get; private set; } = null!;

        [Output("mirrorId")]
        public Output<int> MirrorId { get; private set; } = null!;

        /// <summary>
        /// Determines if only protected branches are mirrored.
        /// </summary>
        [Output("onlyProtectedBranches")]
        public Output<bool?> OnlyProtectedBranches { get; private set; } = null!;

        /// <summary>
        /// The id of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URL of the remote repository to be mirrored.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectMirror resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectMirror(string name, ProjectMirrorArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectMirror:ProjectMirror", name, args ?? new ProjectMirrorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectMirror(string name, Input<string> id, ProjectMirrorState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectMirror:ProjectMirror", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectMirror resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectMirror Get(string name, Input<string> id, ProjectMirrorState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectMirror(name, id, state, options);
        }
    }

    public sealed class ProjectMirrorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the mirror is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Determines if divergent refs are skipped.
        /// </summary>
        [Input("keepDivergentRefs")]
        public Input<bool>? KeepDivergentRefs { get; set; }

        /// <summary>
        /// Determines if only protected branches are mirrored.
        /// </summary>
        [Input("onlyProtectedBranches")]
        public Input<bool>? OnlyProtectedBranches { get; set; }

        /// <summary>
        /// The id of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The URL of the remote repository to be mirrored.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ProjectMirrorArgs()
        {
        }
    }

    public sealed class ProjectMirrorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the mirror is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Determines if divergent refs are skipped.
        /// </summary>
        [Input("keepDivergentRefs")]
        public Input<bool>? KeepDivergentRefs { get; set; }

        [Input("mirrorId")]
        public Input<int>? MirrorId { get; set; }

        /// <summary>
        /// Determines if only protected branches are mirrored.
        /// </summary>
        [Input("onlyProtectedBranches")]
        public Input<bool>? OnlyProtectedBranches { get; set; }

        /// <summary>
        /// The id of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URL of the remote repository to be mirrored.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProjectMirrorState()
        {
        }
    }
}
