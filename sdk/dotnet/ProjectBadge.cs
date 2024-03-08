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
    /// The `gitlab.ProjectBadge` resource allows to manage the lifecycle of project badges.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#project-badges)
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
    ///     var foo = new GitLab.Project("foo");
    /// 
    ///     var example = new GitLab.ProjectBadge("example", new()
    ///     {
    ///         Project = foo.Id,
    ///         LinkUrl = "https://example.com/badge-123",
    ///         ImageUrl = "https://example.com/badge-123.svg",
    ///     });
    /// 
    ///     // Pipeline status badges with placeholders will be enabled
    ///     var gitlabPipeline = new GitLab.ProjectBadge("gitlabPipeline", new()
    ///     {
    ///         Project = foo.Id,
    ///         LinkUrl = "https://gitlab.example.com/%{project_path}/-/pipelines?ref=%{default_branch}",
    ///         ImageUrl = "https://gitlab.example.com/%{project_path}/badges/%{default_branch}/pipeline.svg",
    ///     });
    /// 
    ///     // Test coverage report badges with placeholders will be enabled
    ///     var gitlabCoverage = new GitLab.ProjectBadge("gitlabCoverage", new()
    ///     {
    ///         Project = foo.Id,
    ///         LinkUrl = "https://gitlab.example.com/%{project_path}/-/jobs",
    ///         ImageUrl = "https://gitlab.example.com/%{project_path}/badges/%{default_branch}/coverage.svg",
    ///     });
    /// 
    ///     // Latest release badges with placeholders will be enabled
    ///     var gitlabRelease = new GitLab.ProjectBadge("gitlabRelease", new()
    ///     {
    ///         Project = foo.Id,
    ///         LinkUrl = "https://gitlab.example.com/%{project_path}/-/releases",
    ///         ImageUrl = "https://gitlab.example.com/%{project_path}/-/badges/release.svg",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// GitLab project badges can be imported using an id made up of `{project_id}:{badge_id}`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectBadge:ProjectBadge foo 1:3
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectBadge:ProjectBadge")]
    public partial class ProjectBadge : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The image url which will be presented on project overview.
        /// </summary>
        [Output("imageUrl")]
        public Output<string> ImageUrl { get; private set; } = null!;

        /// <summary>
        /// The url linked with the badge.
        /// </summary>
        [Output("linkUrl")]
        public Output<string> LinkUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the badge.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the project to add the badge to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The image_url argument rendered (in case of use of placeholders).
        /// </summary>
        [Output("renderedImageUrl")]
        public Output<string> RenderedImageUrl { get; private set; } = null!;

        /// <summary>
        /// The link_url argument rendered (in case of use of placeholders).
        /// </summary>
        [Output("renderedLinkUrl")]
        public Output<string> RenderedLinkUrl { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectBadge resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectBadge(string name, ProjectBadgeArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectBadge:ProjectBadge", name, args ?? new ProjectBadgeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectBadge(string name, Input<string> id, ProjectBadgeState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectBadge:ProjectBadge", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectBadge resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectBadge Get(string name, Input<string> id, ProjectBadgeState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectBadge(name, id, state, options);
        }
    }

    public sealed class ProjectBadgeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The image url which will be presented on project overview.
        /// </summary>
        [Input("imageUrl", required: true)]
        public Input<string> ImageUrl { get; set; } = null!;

        /// <summary>
        /// The url linked with the badge.
        /// </summary>
        [Input("linkUrl", required: true)]
        public Input<string> LinkUrl { get; set; } = null!;

        /// <summary>
        /// The name of the badge.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the project to add the badge to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ProjectBadgeArgs()
        {
        }
        public static new ProjectBadgeArgs Empty => new ProjectBadgeArgs();
    }

    public sealed class ProjectBadgeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The image url which will be presented on project overview.
        /// </summary>
        [Input("imageUrl")]
        public Input<string>? ImageUrl { get; set; }

        /// <summary>
        /// The url linked with the badge.
        /// </summary>
        [Input("linkUrl")]
        public Input<string>? LinkUrl { get; set; }

        /// <summary>
        /// The name of the badge.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the project to add the badge to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The image_url argument rendered (in case of use of placeholders).
        /// </summary>
        [Input("renderedImageUrl")]
        public Input<string>? RenderedImageUrl { get; set; }

        /// <summary>
        /// The link_url argument rendered (in case of use of placeholders).
        /// </summary>
        [Input("renderedLinkUrl")]
        public Input<string>? RenderedLinkUrl { get; set; }

        public ProjectBadgeState()
        {
        }
        public static new ProjectBadgeState Empty => new ProjectBadgeState();
    }
}
