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
    /// The `gitlab.ProjectMirror` resource allows to manage the lifecycle of a project mirror.
    /// 
    /// This is for *pushing* changes to a remote repository. *Pull Mirroring* can be configured using a combination of the
    /// import_url, mirror, and mirror_trigger_builds properties on the gitlab.Project resource.
    /// 
    /// &gt; **Warning** By default, the provider sets the `keep_divergent_refs` argument to `True`.
    ///    If you manually set `keep_divergent_refs` to `False`, GitLab mirroring removes branches in the target that aren't in the source.
    ///    This action can result in unexpected branch deletions.
    /// 
    /// &gt; **Destroy Behavior** GitLab 14.10 introduced an API endpoint to delete a project mirror.
    ///    Therefore, for GitLab 14.10 and newer the project mirror will be destroyed when the resource is destroyed.
    ///    For older versions, the mirror will be disabled and the resource will be destroyed.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/remote_mirrors.html)
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
    ///     var foo = new GitLab.ProjectMirror("foo", new()
    ///     {
    ///         Project = "1",
    ///         Url = "https://username:password@github.com/org/repository.git",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectMirror:ProjectMirror foo "12345:1337"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectMirror:ProjectMirror")]
    public partial class ProjectMirror : global::Pulumi.CustomResource
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

        /// <summary>
        /// Mirror ID.
        /// </summary>
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
                AdditionalSecretOutputs =
                {
                    "url",
                },
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

    public sealed class ProjectMirrorArgs : global::Pulumi.ResourceArgs
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

        [Input("url", required: true)]
        private Input<string>? _url;

        /// <summary>
        /// The URL of the remote repository to be mirrored.
        /// </summary>
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProjectMirrorArgs()
        {
        }
        public static new ProjectMirrorArgs Empty => new ProjectMirrorArgs();
    }

    public sealed class ProjectMirrorState : global::Pulumi.ResourceArgs
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
        /// Mirror ID.
        /// </summary>
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

        [Input("url")]
        private Input<string>? _url;

        /// <summary>
        /// The URL of the remote repository to be mirrored.
        /// </summary>
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProjectMirrorState()
        {
        }
        public static new ProjectMirrorState Empty => new ProjectMirrorState();
    }
}
