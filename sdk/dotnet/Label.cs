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
    /// The `gitlab.Label` resource allows to manage the lifecycle of a project label.
    /// 
    /// &gt; This resource is deprecated. use `gitlab.ProjectLabel`instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/labels.html#project-labels)
    /// </summary>
    [GitLabResourceType("gitlab:index/label:Label")]
    public partial class Label : global::Pulumi.CustomResource
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
        /// The id of the project label.
        /// </summary>
        [Output("labelId")]
        public Output<int> LabelId { get; private set; } = null!;

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the label to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Label resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Label(string name, LabelArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/label:Label", name, args ?? new LabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Label(string name, Input<string> id, LabelState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/label:Label", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Label resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Label Get(string name, Input<string> id, LabelState? state = null, CustomResourceOptions? options = null)
        {
            return new Label(name, id, state, options);
        }
    }

    public sealed class LabelArgs : global::Pulumi.ResourceArgs
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
        /// The name of the label.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or id of the project to add the label to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public LabelArgs()
        {
        }
        public static new LabelArgs Empty => new LabelArgs();
    }

    public sealed class LabelState : global::Pulumi.ResourceArgs
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
        /// The id of the project label.
        /// </summary>
        [Input("labelId")]
        public Input<int>? LabelId { get; set; }

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or id of the project to add the label to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public LabelState()
        {
        }
        public static new LabelState Empty => new LabelState();
    }
}
