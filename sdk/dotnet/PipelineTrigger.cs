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
    /// This resource allows you to create and manage pipeline triggers
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new GitLab.PipelineTrigger("example", new GitLab.PipelineTriggerArgs
    ///         {
    ///             Description = "Used to trigger builds",
    ///             Project = "12345",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class PipelineTrigger : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the pipeline trigger.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the trigger to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a PipelineTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PipelineTrigger(string name, PipelineTriggerArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/pipelineTrigger:PipelineTrigger", name, args ?? new PipelineTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PipelineTrigger(string name, Input<string> id, PipelineTriggerState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/pipelineTrigger:PipelineTrigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PipelineTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PipelineTrigger Get(string name, Input<string> id, PipelineTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new PipelineTrigger(name, id, state, options);
        }
    }

    public sealed class PipelineTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the pipeline trigger.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The name or id of the project to add the trigger to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public PipelineTriggerArgs()
        {
        }
    }

    public sealed class PipelineTriggerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the pipeline trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name or id of the project to add the trigger to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("token")]
        public Input<string>? Token { get; set; }

        public PipelineTriggerState()
        {
        }
    }
}
