// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gitlab
{
    /// <summary>
    /// This resource allows you to create and manage pipeline schedules.
    /// For further information on clusters, consult the [gitlab
    /// documentation](https://docs.gitlab.com/ce/user/project/pipelines/schedules.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/pipeline_schedule.html.markdown.
    /// </summary>
    public partial class PipelineSchedule : Pulumi.CustomResource
    {
        /// <summary>
        /// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// The cron (e.g. `0 1 * * *`).
        /// </summary>
        [Output("cron")]
        public Output<string> Cron { get; private set; } = null!;

        /// <summary>
        /// The timezone.
        /// </summary>
        [Output("cronTimezone")]
        public Output<string?> CronTimezone { get; private set; } = null!;

        /// <summary>
        /// The description of the pipeline schedule.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the schedule to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The branch/tag name to be triggered.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;


        /// <summary>
        /// Create a PipelineSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PipelineSchedule(string name, PipelineScheduleArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/pipelineSchedule:PipelineSchedule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private PipelineSchedule(string name, Input<string> id, PipelineScheduleState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/pipelineSchedule:PipelineSchedule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PipelineSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PipelineSchedule Get(string name, Input<string> id, PipelineScheduleState? state = null, CustomResourceOptions? options = null)
        {
            return new PipelineSchedule(name, id, state, options);
        }
    }

    public sealed class PipelineScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The cron (e.g. `0 1 * * *`).
        /// </summary>
        [Input("cron", required: true)]
        public Input<string> Cron { get; set; } = null!;

        /// <summary>
        /// The timezone.
        /// </summary>
        [Input("cronTimezone")]
        public Input<string>? CronTimezone { get; set; }

        /// <summary>
        /// The description of the pipeline schedule.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The name or id of the project to add the schedule to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The branch/tag name to be triggered.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        public PipelineScheduleArgs()
        {
        }
    }

    public sealed class PipelineScheduleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The cron (e.g. `0 1 * * *`).
        /// </summary>
        [Input("cron")]
        public Input<string>? Cron { get; set; }

        /// <summary>
        /// The timezone.
        /// </summary>
        [Input("cronTimezone")]
        public Input<string>? CronTimezone { get; set; }

        /// <summary>
        /// The description of the pipeline schedule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name or id of the project to add the schedule to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The branch/tag name to be triggered.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        public PipelineScheduleState()
        {
        }
    }
}
