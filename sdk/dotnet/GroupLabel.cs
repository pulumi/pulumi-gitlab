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
    /// ## # gitlab\_group\_label
    /// 
    /// This resource allows you to create and manage labels for your GitLab groups.
    /// For further information on labels, consult the [gitlab
    /// documentation](https://docs.gitlab.com/ee/user/project/labels.html#group-labels).
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
    ///         var fixme = new GitLab.GroupLabel("fixme", new GitLab.GroupLabelArgs
    ///         {
    ///             Color = "#ffcc00",
    ///             Description = "issue with failing tests",
    ///             Group = "example",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Gitlab group labels can be imported using an id made up of `{group_id}:{group_label_id}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/groupLabel:GroupLabel example 12345:fixme
    /// ```
    /// </summary>
    public partial class GroupLabel : Pulumi.CustomResource
    {
        /// <summary>
        /// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        /// </summary>
        [Output("color")]
        public Output<string> Color { get; private set; } = null!;

        /// <summary>
        /// The description of the label.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name or id of the group to add the label to.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a GroupLabel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupLabel(string name, GroupLabelArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupLabel:GroupLabel", name, args ?? new GroupLabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupLabel(string name, Input<string> id, GroupLabelState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupLabel:GroupLabel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupLabel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupLabel Get(string name, Input<string> id, GroupLabelState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupLabel(name, id, state, options);
        }
    }

    public sealed class GroupLabelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// The description of the label.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name or id of the group to add the label to.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupLabelArgs()
        {
        }
    }

    public sealed class GroupLabelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The description of the label.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name or id of the group to add the label to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupLabelState()
        {
        }
    }
}
