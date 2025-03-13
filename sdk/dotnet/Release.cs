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
    /// The `gitlab.Release` resource allows to manage the lifecycle of releases in gitlab.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/releases/)
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
    ///     // Create a project
    ///     var example = new GitLab.Project("example", new()
    ///     {
    ///         Name = "example",
    ///         Description = "An example project",
    ///     });
    /// 
    ///     // Create a release
    ///     var exampleRelease = new GitLab.Release("example", new()
    ///     {
    ///         Project = example.Id,
    ///         Name = "test-release",
    ///         TagName = "v1.0.0",
    ///         Description = "Test release description",
    ///         Ref = "main",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_release`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_release.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// Gitlab release link can be imported with a key composed of `&lt;project&gt;:&lt;tag_name&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/release:Release example "12345:test"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/release:Release")]
    public partial class Release : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The release assets.
        /// </summary>
        [Output("assets")]
        public Output<Outputs.ReleaseAssets> Assets { get; private set; } = null!;

        /// <summary>
        /// The author of the release.
        /// </summary>
        [Output("author")]
        public Output<Outputs.ReleaseAuthor> Author { get; private set; } = null!;

        /// <summary>
        /// The release commit.
        /// </summary>
        [Output("commit")]
        public Output<Outputs.ReleaseCommit> Commit { get; private set; } = null!;

        /// <summary>
        /// The path to the commit
        /// </summary>
        [Output("commitPath")]
        public Output<string> CommitPath { get; private set; } = null!;

        /// <summary>
        /// Date and time the release was created. In ISO 8601 format (2019-03-15T08:00:00Z).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the release. You can use Markdown.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// HTML rendered Markdown of the release description.
        /// </summary>
        [Output("descriptionHtml")]
        public Output<string> DescriptionHtml { get; private set; } = null!;

        /// <summary>
        /// Links of the release
        /// </summary>
        [Output("links")]
        public Output<Outputs.ReleaseLinks> Links { get; private set; } = null!;

        /// <summary>
        /// The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones.
        /// </summary>
        [Output("milestones")]
        public Output<ImmutableArray<string>> Milestones { get; private set; } = null!;

        /// <summary>
        /// The name of the release.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// If a tag specified in tag*name doesn't exist, the release is created from ref and tagged with tag*name. It can be a commit SHA, another tag name, or a branch name.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;

        /// <summary>
        /// Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (2019-03-15T08:00:00Z). Only provide this field if creating an upcoming or historical release.
        /// </summary>
        [Output("releasedAt")]
        public Output<string> ReleasedAt { get; private set; } = null!;

        /// <summary>
        /// Message to use if creating a new annotated tag.
        /// </summary>
        [Output("tagMessage")]
        public Output<string?> TagMessage { get; private set; } = null!;

        /// <summary>
        /// The tag where the release is created from.
        /// </summary>
        [Output("tagName")]
        public Output<string> TagName { get; private set; } = null!;

        /// <summary>
        /// The path to the tag.
        /// </summary>
        [Output("tagPath")]
        public Output<string> TagPath { get; private set; } = null!;

        /// <summary>
        /// Whether the release_at attribute is set to a future date.
        /// </summary>
        [Output("upcomingRelease")]
        public Output<bool> UpcomingRelease { get; private set; } = null!;


        /// <summary>
        /// Create a Release resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Release(string name, ReleaseArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/release:Release", name, args ?? new ReleaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Release(string name, Input<string> id, ReleaseState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/release:Release", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Release resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Release Get(string name, Input<string> id, ReleaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Release(name, id, state, options);
        }
    }

    public sealed class ReleaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The release assets.
        /// </summary>
        [Input("assets")]
        public Input<Inputs.ReleaseAssetsArgs>? Assets { get; set; }

        /// <summary>
        /// The description of the release. You can use Markdown.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("milestones")]
        private InputList<string>? _milestones;

        /// <summary>
        /// The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones.
        /// </summary>
        public InputList<string> Milestones
        {
            get => _milestones ?? (_milestones = new InputList<string>());
            set => _milestones = value;
        }

        /// <summary>
        /// The name of the release.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// If a tag specified in tag*name doesn't exist, the release is created from ref and tagged with tag*name. It can be a commit SHA, another tag name, or a branch name.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (2019-03-15T08:00:00Z). Only provide this field if creating an upcoming or historical release.
        /// </summary>
        [Input("releasedAt")]
        public Input<string>? ReleasedAt { get; set; }

        /// <summary>
        /// Message to use if creating a new annotated tag.
        /// </summary>
        [Input("tagMessage")]
        public Input<string>? TagMessage { get; set; }

        /// <summary>
        /// The tag where the release is created from.
        /// </summary>
        [Input("tagName", required: true)]
        public Input<string> TagName { get; set; } = null!;

        public ReleaseArgs()
        {
        }
        public static new ReleaseArgs Empty => new ReleaseArgs();
    }

    public sealed class ReleaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The release assets.
        /// </summary>
        [Input("assets")]
        public Input<Inputs.ReleaseAssetsGetArgs>? Assets { get; set; }

        /// <summary>
        /// The author of the release.
        /// </summary>
        [Input("author")]
        public Input<Inputs.ReleaseAuthorGetArgs>? Author { get; set; }

        /// <summary>
        /// The release commit.
        /// </summary>
        [Input("commit")]
        public Input<Inputs.ReleaseCommitGetArgs>? Commit { get; set; }

        /// <summary>
        /// The path to the commit
        /// </summary>
        [Input("commitPath")]
        public Input<string>? CommitPath { get; set; }

        /// <summary>
        /// Date and time the release was created. In ISO 8601 format (2019-03-15T08:00:00Z).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The description of the release. You can use Markdown.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// HTML rendered Markdown of the release description.
        /// </summary>
        [Input("descriptionHtml")]
        public Input<string>? DescriptionHtml { get; set; }

        /// <summary>
        /// Links of the release
        /// </summary>
        [Input("links")]
        public Input<Inputs.ReleaseLinksGetArgs>? Links { get; set; }

        [Input("milestones")]
        private InputList<string>? _milestones;

        /// <summary>
        /// The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones.
        /// </summary>
        public InputList<string> Milestones
        {
            get => _milestones ?? (_milestones = new InputList<string>());
            set => _milestones = value;
        }

        /// <summary>
        /// The name of the release.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// If a tag specified in tag*name doesn't exist, the release is created from ref and tagged with tag*name. It can be a commit SHA, another tag name, or a branch name.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (2019-03-15T08:00:00Z). Only provide this field if creating an upcoming or historical release.
        /// </summary>
        [Input("releasedAt")]
        public Input<string>? ReleasedAt { get; set; }

        /// <summary>
        /// Message to use if creating a new annotated tag.
        /// </summary>
        [Input("tagMessage")]
        public Input<string>? TagMessage { get; set; }

        /// <summary>
        /// The tag where the release is created from.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        /// <summary>
        /// The path to the tag.
        /// </summary>
        [Input("tagPath")]
        public Input<string>? TagPath { get; set; }

        /// <summary>
        /// Whether the release_at attribute is set to a future date.
        /// </summary>
        [Input("upcomingRelease")]
        public Input<bool>? UpcomingRelease { get; set; }

        public ReleaseState()
        {
        }
        public static new ReleaseState Empty => new ReleaseState();
    }
}
