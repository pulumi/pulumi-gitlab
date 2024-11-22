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
    /// The `gitlab.PipelineScheduleVariable` resource allows to manage the lifecycle of a variable for a pipeline schedule.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html#pipeline-schedule-variables)
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
    ///     var example = new GitLab.PipelineSchedule("example", new()
    ///     {
    ///         Project = "12345",
    ///         Description = "Used to schedule builds",
    ///         Ref = "master",
    ///         Cron = "0 1 * * *",
    ///     });
    /// 
    ///     var examplePipelineScheduleVariable = new GitLab.PipelineScheduleVariable("example", new()
    ///     {
    ///         Project = example.Project,
    ///         PipelineScheduleId = example.PipelineScheduleId,
    ///         Key = "EXAMPLE_KEY",
    ///         Value = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_pipeline_schedule_variable`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_pipeline_schedule_variable.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// Pipeline schedule variables can be imported using an id made up of `project_id:pipeline_schedule_id:key`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable example 123456789:13:mykey
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable")]
    public partial class PipelineScheduleVariable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the variable.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The id of the pipeline schedule.
        /// </summary>
        [Output("pipelineScheduleId")]
        public Output<int> PipelineScheduleId { get; private set; } = null!;

        /// <summary>
        /// The id of the project to add the schedule to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Value of the variable.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        [Output("variableType")]
        public Output<string> VariableType { get; private set; } = null!;


        /// <summary>
        /// Create a PipelineScheduleVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PipelineScheduleVariable(string name, PipelineScheduleVariableArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, args ?? new PipelineScheduleVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PipelineScheduleVariable(string name, Input<string> id, PipelineScheduleVariableState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PipelineScheduleVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PipelineScheduleVariable Get(string name, Input<string> id, PipelineScheduleVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new PipelineScheduleVariable(name, id, state, options);
        }
    }

    public sealed class PipelineScheduleVariableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the variable.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The id of the pipeline schedule.
        /// </summary>
        [Input("pipelineScheduleId", required: true)]
        public Input<int> PipelineScheduleId { get; set; } = null!;

        /// <summary>
        /// The id of the project to add the schedule to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Value of the variable.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        [Input("variableType")]
        public Input<string>? VariableType { get; set; }

        public PipelineScheduleVariableArgs()
        {
        }
        public static new PipelineScheduleVariableArgs Empty => new PipelineScheduleVariableArgs();
    }

    public sealed class PipelineScheduleVariableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the variable.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The id of the pipeline schedule.
        /// </summary>
        [Input("pipelineScheduleId")]
        public Input<int>? PipelineScheduleId { get; set; }

        /// <summary>
        /// The id of the project to add the schedule to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Value of the variable.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        [Input("variableType")]
        public Input<string>? VariableType { get; set; }

        public PipelineScheduleVariableState()
        {
        }
        public static new PipelineScheduleVariableState Empty => new PipelineScheduleVariableState();
    }
}
