// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ValueStreamAnalyticsStageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean whether the stage is customized. If false, it assigns a built-in default stage by name.
        /// </summary>
        [Input("custom")]
        public Input<bool>? Custom { get; set; }

        /// <summary>
        /// End event identifier. Valid values are: `CODE_STAGE_START`, `ISSUE_CLOSED`, `ISSUE_CREATED`, `ISSUE_DEPLOYED_TO_PRODUCTION`, `ISSUE_FIRST_ADDED_TO_BOARD`, `ISSUE_FIRST_ADDED_TO_ITERATION`, `ISSUE_FIRST_ASSIGNED_AT`, `ISSUE_FIRST_ASSOCIATED_WITH_MILESTONE`, `ISSUE_FIRST_MENTIONED_IN_COMMIT`, `ISSUE_LABEL_ADDED`, `ISSUE_LABEL_REMOVED`, `ISSUE_LAST_EDITED`, `ISSUE_STAGE_END`, `MERGE_REQUEST_CLOSED`, `MERGE_REQUEST_CREATED`, `MERGE_REQUEST_FIRST_ASSIGNED_AT`, `MERGE_REQUEST_FIRST_COMMIT_AT`, `MERGE_REQUEST_FIRST_DEPLOYED_TO_PRODUCTION`, `MERGE_REQUEST_LABEL_ADDED`, `MERGE_REQUEST_LABEL_REMOVED`, `MERGE_REQUEST_LAST_BUILD_FINISHED`, `MERGE_REQUEST_LAST_BUILD_STARTED`, `MERGE_REQUEST_LAST_EDITED`, `MERGE_REQUEST_MERGED`, `MERGE_REQUEST_REVIEWER_FIRST_ASSIGNED`, `MERGE_REQUEST_PLAN_STAGE_START`
        /// </summary>
        [Input("endEventIdentifier")]
        public Input<string>? EndEventIdentifier { get; set; }

        /// <summary>
        /// Label ID associated with the end event identifier. In the format of `gid://gitlab/GroupLabel/&lt;id&gt;` or `gid://gitlab/ProjectLabel/&lt;id&gt;`
        /// </summary>
        [Input("endEventLabelId")]
        public Input<string>? EndEventLabelId { get; set; }

        /// <summary>
        /// Boolean whether the stage is hidden, GitLab provided default stages are hidden by default.
        /// </summary>
        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        /// <summary>
        /// The ID of the value stream stage.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the value stream stage.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Start event identifier. Valid values are: `CODE_STAGE_START`, `ISSUE_CLOSED`, `ISSUE_CREATED`, `ISSUE_DEPLOYED_TO_PRODUCTION`, `ISSUE_FIRST_ADDED_TO_BOARD`, `ISSUE_FIRST_ADDED_TO_ITERATION`, `ISSUE_FIRST_ASSIGNED_AT`, `ISSUE_FIRST_ASSOCIATED_WITH_MILESTONE`, `ISSUE_FIRST_MENTIONED_IN_COMMIT`, `ISSUE_LABEL_ADDED`, `ISSUE_LABEL_REMOVED`, `ISSUE_LAST_EDITED`, `ISSUE_STAGE_END`, `MERGE_REQUEST_CLOSED`, `MERGE_REQUEST_CREATED`, `MERGE_REQUEST_FIRST_ASSIGNED_AT`, `MERGE_REQUEST_FIRST_COMMIT_AT`, `MERGE_REQUEST_FIRST_DEPLOYED_TO_PRODUCTION`, `MERGE_REQUEST_LABEL_ADDED`, `MERGE_REQUEST_LABEL_REMOVED`, `MERGE_REQUEST_LAST_BUILD_FINISHED`, `MERGE_REQUEST_LAST_BUILD_STARTED`, `MERGE_REQUEST_LAST_EDITED`, `MERGE_REQUEST_MERGED`, `MERGE_REQUEST_REVIEWER_FIRST_ASSIGNED`, `MERGE_REQUEST_PLAN_STAGE_START`
        /// </summary>
        [Input("startEventIdentifier")]
        public Input<string>? StartEventIdentifier { get; set; }

        /// <summary>
        /// Label ID associated with the start event identifier. In the format of `gid://gitlab/GroupLabel/&lt;id&gt;` or `gid://gitlab/ProjectLabel/&lt;id&gt;`
        /// </summary>
        [Input("startEventLabelId")]
        public Input<string>? StartEventLabelId { get; set; }

        public ValueStreamAnalyticsStageGetArgs()
        {
        }
        public static new ValueStreamAnalyticsStageGetArgs Empty => new ValueStreamAnalyticsStageGetArgs();
    }
}
