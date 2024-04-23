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
    /// The `gitlab.DeployKeyEnable` resource allows to enable an already existing deploy key (see `gitlab.DeployKey resource`) for a specific project.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#enable-a-deploy-key)
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
    ///     // A repo to host the deployment key
    ///     var parent = new GitLab.Project("parent", new()
    ///     {
    ///         Name = "parent_project",
    ///     });
    /// 
    ///     // A second repo to use the deployment key from the parent project
    ///     var foo = new GitLab.Project("foo", new()
    ///     {
    ///         Name = "foo_project",
    ///     });
    /// 
    ///     // Upload a deployment key for the parent repo
    ///     var parentDeployKey = new GitLab.DeployKey("parent", new()
    ///     {
    ///         Project = parent.Id,
    ///         Title = "Example deploy key",
    ///         Key = "ssh-ed25519 AAAA...",
    ///     });
    /// 
    ///     // Enable the deployment key on the second repo
    ///     var fooDeployKeyEnable = new GitLab.DeployKeyEnable("foo", new()
    ///     {
    ///         Project = foo.Id,
    ///         KeyId = parentDeployKey.DeployKeyId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.
    /// 
    /// `project_id` can be whatever the [get single project api][get_single_project] takes for
    /// 
    /// its `:id` value, so for example:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example richardc/example:67890
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/deployKeyEnable:DeployKeyEnable")]
    public partial class DeployKeyEnable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Can deploy key push to the project's repository.
        /// </summary>
        [Output("canPush")]
        public Output<bool?> CanPush { get; private set; } = null!;

        /// <summary>
        /// Deploy key.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The Gitlab key id for the pre-existing deploy key
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the deploy key to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Deploy key's title.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a DeployKeyEnable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeployKeyEnable(string name, DeployKeyEnableArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/deployKeyEnable:DeployKeyEnable", name, args ?? new DeployKeyEnableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeployKeyEnable(string name, Input<string> id, DeployKeyEnableState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/deployKeyEnable:DeployKeyEnable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeployKeyEnable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeployKeyEnable Get(string name, Input<string> id, DeployKeyEnableState? state = null, CustomResourceOptions? options = null)
        {
            return new DeployKeyEnable(name, id, state, options);
        }
    }

    public sealed class DeployKeyEnableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can deploy key push to the project's repository.
        /// </summary>
        [Input("canPush")]
        public Input<bool>? CanPush { get; set; }

        /// <summary>
        /// Deploy key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The Gitlab key id for the pre-existing deploy key
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        /// <summary>
        /// The name or id of the project to add the deploy key to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Deploy key's title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public DeployKeyEnableArgs()
        {
        }
        public static new DeployKeyEnableArgs Empty => new DeployKeyEnableArgs();
    }

    public sealed class DeployKeyEnableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can deploy key push to the project's repository.
        /// </summary>
        [Input("canPush")]
        public Input<bool>? CanPush { get; set; }

        /// <summary>
        /// Deploy key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The Gitlab key id for the pre-existing deploy key
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The name or id of the project to add the deploy key to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Deploy key's title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public DeployKeyEnableState()
        {
        }
        public static new DeployKeyEnableState Empty => new DeployKeyEnableState();
    }
}
