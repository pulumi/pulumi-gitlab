// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectIssue` data source allows to retrieve details about an issue in a project.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = gitlab.getProject({
 *     pathWithNamespace: "foo/bar/baz",
 * });
 * const welcomeIssue = foo.then(foo => gitlab.getProjectIssue({
 *     project: foo.id,
 *     iid: 1,
 * }));
 * export const welcomeIssueWebUrl = webUrl;
 * ```
 */
export function getProjectIssue(args: GetProjectIssueArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectIssueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectIssue:getProjectIssue", {
        "iid": args.iid,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectIssue.
 */
export interface GetProjectIssueArgs {
    /**
     * The internal ID of the project's issue.
     */
    iid: number;
    /**
     * The name or ID of the project.
     */
    project: string;
}

/**
 * A collection of values returned by getProjectIssue.
 */
export interface GetProjectIssueResult {
    /**
     * The IDs of the users to assign the issue to.
     */
    readonly assigneeIds: number[];
    /**
     * The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     */
    readonly authorId: number;
    /**
     * When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    readonly closedAt: string;
    /**
     * The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     */
    readonly closedByUserId: number;
    /**
     * Set an issue to be confidential.
     */
    readonly confidential: boolean;
    /**
     * When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     */
    readonly createdAt: string;
    /**
     * The description of an issue. Limited to 1,048,576 characters.
     */
    readonly description: string;
    /**
     * Whether the issue is locked for discussions or not.
     */
    readonly discussionLocked: boolean;
    /**
     * The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     */
    readonly discussionToResolve: string;
    /**
     * The number of downvotes the issue has received.
     */
    readonly downvotes: number;
    /**
     * The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    readonly dueDate: string;
    /**
     * ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     */
    readonly epicId: number;
    /**
     * The ID of the epic issue.
     */
    readonly epicIssueId: number;
    /**
     * The external ID of the issue.
     */
    readonly externalId: string;
    /**
     * The human-readable time estimate of the issue.
     */
    readonly humanTimeEstimate: string;
    /**
     * The human-readable total time spent of the issue.
     */
    readonly humanTotalTimeSpent: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The internal ID of the project's issue.
     */
    readonly iid: number;
    /**
     * The instance-wide ID of the issue.
     */
    readonly issueId: number;
    /**
     * The ID of the issue link.
     */
    readonly issueLinkId: number;
    /**
     * The type of issue. Valid values are: `issue`, `incident`, `testCase`.
     */
    readonly issueType: string;
    /**
     * The labels of an issue.
     */
    readonly labels: string[];
    /**
     * The links of the issue.
     */
    readonly links: {[key: string]: string};
    /**
     * The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     */
    readonly mergeRequestToResolveDiscussionsOf: number;
    /**
     * The number of merge requests associated with the issue.
     */
    readonly mergeRequestsCount: number;
    /**
     * The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
     */
    readonly milestoneId: number;
    /**
     * The ID of the issue that was moved to.
     */
    readonly movedToId: number;
    /**
     * The name or ID of the project.
     */
    readonly project: string;
    /**
     * The references of the issue.
     */
    readonly references: {[key: string]: string};
    /**
     * The state of the issue. Valid values are: `opened`, `closed`.
     */
    readonly state: string;
    /**
     * Whether the authenticated user is subscribed to the issue or not.
     */
    readonly subscribed: boolean;
    /**
     * The task completion status. It's always a one element list.
     */
    readonly taskCompletionStatuses: outputs.GetProjectIssueTaskCompletionStatus[];
    /**
     * The time estimate of the issue.
     */
    readonly timeEstimate: number;
    /**
     * The title of the issue.
     */
    readonly title: string;
    /**
     * The total time spent of the issue.
     */
    readonly totalTimeSpent: number;
    /**
     * When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    readonly updatedAt: string;
    /**
     * The number of upvotes the issue has received.
     */
    readonly upvotes: number;
    /**
     * The number of user notes on the issue.
     */
    readonly userNotesCount: number;
    /**
     * The web URL of the issue.
     */
    readonly webUrl: string;
    /**
     * The weight of the issue. Valid values are greater than or equal to 0.
     */
    readonly weight: number;
}
/**
 * The `gitlab.ProjectIssue` data source allows to retrieve details about an issue in a project.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = gitlab.getProject({
 *     pathWithNamespace: "foo/bar/baz",
 * });
 * const welcomeIssue = foo.then(foo => gitlab.getProjectIssue({
 *     project: foo.id,
 *     iid: 1,
 * }));
 * export const welcomeIssueWebUrl = webUrl;
 * ```
 */
export function getProjectIssueOutput(args: GetProjectIssueOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectIssueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getProjectIssue:getProjectIssue", {
        "iid": args.iid,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectIssue.
 */
export interface GetProjectIssueOutputArgs {
    /**
     * The internal ID of the project's issue.
     */
    iid: pulumi.Input<number>;
    /**
     * The name or ID of the project.
     */
    project: pulumi.Input<string>;
}
