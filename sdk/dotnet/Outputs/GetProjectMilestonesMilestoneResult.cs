// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Outputs
{

    [OutputType]
    public sealed class GetProjectMilestonesMilestoneResult
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
        private GetProjectMilestonesMilestoneResult(
            string createdAt,

            string description,

            string dueDate,

            bool expired,

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
