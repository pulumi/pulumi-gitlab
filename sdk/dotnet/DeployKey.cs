// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gitlab
{
    /// <summary>
    /// This resource allows you to create and manage deploy keys for your GitLab projects.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/deploy_key.html.markdown.
    /// </summary>
    public partial class DeployKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
        /// </summary>
        [Output("canPush")]
        public Output<bool?> CanPush { get; private set; } = null!;

        /// <summary>
        /// The public ssh key body.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the deploy key to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A title to describe the deploy key with.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a DeployKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeployKey(string name, DeployKeyArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/deployKey:DeployKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DeployKey(string name, Input<string> id, DeployKeyState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/deployKey:DeployKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeployKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeployKey Get(string name, Input<string> id, DeployKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new DeployKey(name, id, state, options);
        }
    }

    public sealed class DeployKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
        /// </summary>
        [Input("canPush")]
        public Input<bool>? CanPush { get; set; }

        /// <summary>
        /// The public ssh key body.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The name or id of the project to add the deploy key to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// A title to describe the deploy key with.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public DeployKeyArgs()
        {
        }
    }

    public sealed class DeployKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow this deploy key to be used to push changes to the project.  Defaults to `false`. **NOTE::** this cannot currently be managed.
        /// </summary>
        [Input("canPush")]
        public Input<bool>? CanPush { get; set; }

        /// <summary>
        /// The public ssh key body.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The name or id of the project to add the deploy key to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A title to describe the deploy key with.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public DeployKeyState()
        {
        }
    }
}
