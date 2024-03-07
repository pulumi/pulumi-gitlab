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
    /// The `gitlab.GroupProjectFileTemplate` resource allows setting a project from which
    /// custom file templates will be loaded. In order to use this resource, the project selected must be a direct child of
    /// the group selected. After the resource has run, `gitlab_project_template.template_project_id` is available for use.
    /// For more information about which file types are available as templates, view
    /// [GitLab's documentation](https://docs.gitlab.com/ee/user/group/custom_project_templates.html)
    /// 
    /// &gt; This resource requires a GitLab Enterprise instance with a Premium license.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#update-group)
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
    ///     var foo = new GitLab.Group("foo", new()
    ///     {
    ///         Path = "group",
    ///         Description = "An example group",
    ///     });
    /// 
    ///     var bar = new GitLab.Project("bar", new()
    ///     {
    ///         Description = "contains file templates",
    ///         VisibilityLevel = "public",
    ///         NamespaceId = foo.Id,
    ///     });
    /// 
    ///     var templateLink = new GitLab.GroupProjectFileTemplate("templateLink", new()
    ///     {
    ///         GroupId = foo.Id,
    ///         FileTemplateProjectId = bar.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [GitLabResourceType("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate")]
    public partial class GroupProjectFileTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the project that will be used for file templates. This project must be the direct
        /// 			child of the project defined by the group_id
        /// </summary>
        [Output("fileTemplateProjectId")]
        public Output<int> FileTemplateProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the group that will use the file template project. This group must be the direct
        ///             parent of the project defined by project_id
        /// </summary>
        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupProjectFileTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupProjectFileTemplate(string name, GroupProjectFileTemplateArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate", name, args ?? new GroupProjectFileTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupProjectFileTemplate(string name, Input<string> id, GroupProjectFileTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupProjectFileTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupProjectFileTemplate Get(string name, Input<string> id, GroupProjectFileTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupProjectFileTemplate(name, id, state, options);
        }
    }

    public sealed class GroupProjectFileTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the project that will be used for file templates. This project must be the direct
        /// 			child of the project defined by the group_id
        /// </summary>
        [Input("fileTemplateProjectId", required: true)]
        public Input<int> FileTemplateProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the group that will use the file template project. This group must be the direct
        ///             parent of the project defined by project_id
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        public GroupProjectFileTemplateArgs()
        {
        }
        public static new GroupProjectFileTemplateArgs Empty => new GroupProjectFileTemplateArgs();
    }

    public sealed class GroupProjectFileTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the project that will be used for file templates. This project must be the direct
        /// 			child of the project defined by the group_id
        /// </summary>
        [Input("fileTemplateProjectId")]
        public Input<int>? FileTemplateProjectId { get; set; }

        /// <summary>
        /// The ID of the group that will use the file template project. This group must be the direct
        ///             parent of the project defined by project_id
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        public GroupProjectFileTemplateState()
        {
        }
        public static new GroupProjectFileTemplateState Empty => new GroupProjectFileTemplateState();
    }
}
