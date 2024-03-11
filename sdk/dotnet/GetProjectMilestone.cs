// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectMilestone
    {
        /// <summary>
        /// The `gitlab.ProjectMilestone` data source allows get details of a project milestone.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
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
        ///     var example = GitLab.GetProjectMilestone.Invoke(new()
        ///     {
        ///         MilestoneId = 10,
        ///         Project = "foo/bar",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetProjectMilestoneResult> InvokeAsync(GetProjectMilestoneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectMilestoneResult>("gitlab:index/getProjectMilestone:getProjectMilestone", args ?? new GetProjectMilestoneArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.ProjectMilestone` data source allows get details of a project milestone.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
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
        ///     var example = GitLab.GetProjectMilestone.Invoke(new()
        ///     {
        ///         MilestoneId = 10,
        ///         Project = "foo/bar",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetProjectMilestoneResult> Invoke(GetProjectMilestoneInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectMilestoneResult>("gitlab:index/getProjectMilestone:getProjectMilestone", args ?? new GetProjectMilestoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectMilestoneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance-wide ID of the project’s milestone.
        /// </summary>
        [Input("milestoneId", required: true)]
        public int MilestoneId { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetProjectMilestoneArgs()
        {
        }
        public static new GetProjectMilestoneArgs Empty => new GetProjectMilestoneArgs();
    }

    public sealed class GetProjectMilestoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance-wide ID of the project’s milestone.
        /// </summary>
        [Input("milestoneId", required: true)]
        public Input<int> MilestoneId { get; set; } = null!;

        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetProjectMilestoneInvokeArgs()
        {
        }
        public static new GetProjectMilestoneInvokeArgs Empty => new GetProjectMilestoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectMilestoneResult
    {
        /// <summary>
        /// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the milestone.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        /// </summary>
        public readonly string DueDate;
        /// <summary>
        /// Bool, true if milestone expired.
        /// </summary>
        public readonly bool Expired;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the project's milestone.
        /// </summary>
        public readonly int Iid;
        /// <summary>
        /// The instance-wide ID of the project’s milestone.
        /// </summary>
        public readonly int MilestoneId;
        /// <summary>
        /// The ID or URL-encoded path of the project owned by the authenticated user.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The project ID of milestone.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
        /// </summary>
        public readonly string StartDate;
        /// <summary>
        /// The state of the milestone. Valid values are: `active`, `closed`.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The title of a milestone.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The web URL of the milestone.
        /// </summary>
        public readonly string WebUrl;

        [OutputConstructor]
        private GetProjectMilestoneResult(
            string createdAt,

            string description,

            string dueDate,

            bool expired,

            string id,

            int iid,

            int milestoneId,

            string project,

            int projectId,

            string startDate,

            string state,

            string title,

            string updatedAt,

            string webUrl)
        {
            CreatedAt = createdAt;
            Description = description;
            DueDate = dueDate;
            Expired = expired;
            Id = id;
            Iid = iid;
            MilestoneId = milestoneId;
            Project = project;
            ProjectId = projectId;
            StartDate = startDate;
            State = state;
            Title = title;
            UpdatedAt = updatedAt;
            WebUrl = webUrl;
        }
    }
}
