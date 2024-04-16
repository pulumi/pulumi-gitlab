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
    /// The `gitlab.ProjectTag` resource allows to manage the lifecycle of a tag in a project.
    /// 
    /// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/tags.html)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a project for the tag to use
    ///     var example = new GitLab.Project("example", new()
    ///     {
    ///         Name = "example",
    ///         Description = "An example project",
    ///         NamespaceId = exampleGitlabGroup.Id,
    ///     });
    /// 
    ///     var exampleProjectTag = new GitLab.ProjectTag("example", new()
    ///     {
    ///         Name = "example",
    ///         Ref = "main",
    ///         Project = example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Gitlab project tags can be imported with a key composed of `&lt;project_id&gt;:&lt;tag_name&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectTag:ProjectTag example "12345:develop"
    /// ```
    /// 
    /// NOTE: the `ref` attribute won't be available for imported `gitlab_project_tag` resources.
    /// </summary>
    [GitLabResourceType("gitlab:index/projectTag:ProjectTag")]
    public partial class ProjectTag : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The commit associated with the tag.
        /// </summary>
        [Output("commits")]
        public Output<ImmutableArray<Outputs.ProjectTagCommit>> Commits { get; private set; } = null!;

        /// <summary>
        /// The message of the annotated tag.
        /// </summary>
        [Output("message")]
        public Output<string?> Message { get; private set; } = null!;

        /// <summary>
        /// The name of a tag.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Bool, true if tag has tag protection.
        /// </summary>
        [Output("protected")]
        public Output<bool> Protected { get; private set; } = null!;

        /// <summary>
        /// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;

        /// <summary>
        /// The release associated with the tag.
        /// </summary>
        [Output("releases")]
        public Output<ImmutableArray<Outputs.ProjectTagRelease>> Releases { get; private set; } = null!;

        /// <summary>
        /// The unique id assigned to the commit by Gitlab.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectTag(string name, ProjectTagArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectTag:ProjectTag", name, args ?? new ProjectTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectTag(string name, Input<string> id, ProjectTagState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectTag:ProjectTag", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectTag Get(string name, Input<string> id, ProjectTagState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectTag(name, id, state, options);
        }
    }

    public sealed class ProjectTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The message of the annotated tag.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The name of a tag.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        public ProjectTagArgs()
        {
        }
        public static new ProjectTagArgs Empty => new ProjectTagArgs();
    }

    public sealed class ProjectTagState : global::Pulumi.ResourceArgs
    {
        [Input("commits")]
        private InputList<Inputs.ProjectTagCommitGetArgs>? _commits;

        /// <summary>
        /// The commit associated with the tag.
        /// </summary>
        public InputList<Inputs.ProjectTagCommitGetArgs> Commits
        {
            get => _commits ?? (_commits = new InputList<Inputs.ProjectTagCommitGetArgs>());
            set => _commits = value;
        }

        /// <summary>
        /// The message of the annotated tag.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The name of a tag.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Bool, true if tag has tag protection.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        [Input("releases")]
        private InputList<Inputs.ProjectTagReleaseGetArgs>? _releases;

        /// <summary>
        /// The release associated with the tag.
        /// </summary>
        public InputList<Inputs.ProjectTagReleaseGetArgs> Releases
        {
            get => _releases ?? (_releases = new InputList<Inputs.ProjectTagReleaseGetArgs>());
            set => _releases = value;
        }

        /// <summary>
        /// The unique id assigned to the commit by Gitlab.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public ProjectTagState()
        {
        }
        public static new ProjectTagState Empty => new ProjectTagState();
    }
}
