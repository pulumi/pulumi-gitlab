// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetPipelineSchedules
    {
        /// <summary>
        /// The `gitlab.PipelineSchedule` data source retrieves information about a gitlab pipeline schedule for a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html)
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
        ///     var example = GitLab.GetPipelineSchedules.Invoke(new()
        ///     {
        ///         Project = "12345",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPipelineSchedulesResult> InvokeAsync(GetPipelineSchedulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPipelineSchedulesResult>("gitlab:index/getPipelineSchedules:getPipelineSchedules", args ?? new GetPipelineSchedulesArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.PipelineSchedule` data source retrieves information about a gitlab pipeline schedule for a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html)
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
        ///     var example = GitLab.GetPipelineSchedules.Invoke(new()
        ///     {
        ///         Project = "12345",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPipelineSchedulesResult> Invoke(GetPipelineSchedulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPipelineSchedulesResult>("gitlab:index/getPipelineSchedules:getPipelineSchedules", args ?? new GetPipelineSchedulesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.PipelineSchedule` data source retrieves information about a gitlab pipeline schedule for a project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html)
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
        ///     var example = GitLab.GetPipelineSchedules.Invoke(new()
        ///     {
        ///         Project = "12345",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPipelineSchedulesResult> Invoke(GetPipelineSchedulesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPipelineSchedulesResult>("gitlab:index/getPipelineSchedules:getPipelineSchedules", args ?? new GetPipelineSchedulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPipelineSchedulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name or id of the project to add the schedule to.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetPipelineSchedulesArgs()
        {
        }
        public static new GetPipelineSchedulesArgs Empty => new GetPipelineSchedulesArgs();
    }

    public sealed class GetPipelineSchedulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name or id of the project to add the schedule to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetPipelineSchedulesInvokeArgs()
        {
        }
        public static new GetPipelineSchedulesInvokeArgs Empty => new GetPipelineSchedulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetPipelineSchedulesResult
    {
        public readonly string Id;
        /// <summary>
        /// The list of pipeline schedules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPipelineSchedulesPipelineScheduleResult> PipelineSchedules;
        /// <summary>
        /// The name or id of the project to add the schedule to.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private GetPipelineSchedulesResult(
            string id,

            ImmutableArray<Outputs.GetPipelineSchedulesPipelineScheduleResult> pipelineSchedules,

            string project)
        {
            Id = id;
            PipelineSchedules = pipelineSchedules;
            Project = project;
        }
    }
}
