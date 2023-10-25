// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectIssues` data source allows to retrieve details about issues in a project.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)
 */
export function getProjectIssues(args: GetProjectIssuesArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectIssuesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectIssues:getProjectIssues", {
        "assigneeId": args.assigneeId,
        "assigneeUsername": args.assigneeUsername,
        "authorId": args.authorId,
        "confidential": args.confidential,
        "createdAfter": args.createdAfter,
        "createdBefore": args.createdBefore,
        "dueDate": args.dueDate,
        "iids": args.iids,
        "issueType": args.issueType,
        "labels": args.labels,
        "milestone": args.milestone,
        "myReactionEmoji": args.myReactionEmoji,
        "notAssigneeIds": args.notAssigneeIds,
        "notAuthorIds": args.notAuthorIds,
        "notLabels": args.notLabels,
        "notMilestone": args.notMilestone,
        "notMyReactionEmojis": args.notMyReactionEmojis,
        "orderBy": args.orderBy,
        "project": args.project,
        "scope": args.scope,
        "search": args.search,
        "sort": args.sort,
        "updatedAfter": args.updatedAfter,
        "updatedBefore": args.updatedBefore,
        "weight": args.weight,
        "withLabelsDetails": args.withLabelsDetails,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectIssues.
 */
export interface GetProjectIssuesArgs {
    /**
     * Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
     */
    assigneeId?: number;
    /**
     * Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
     */
    assigneeUsername?: string;
    authorId?: number;
    confidential?: boolean;
    /**
     * Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    createdAfter?: string;
    /**
     * Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    createdBefore?: string;
    dueDate?: string;
    /**
     * Return only the issues having the given iid
     */
    iids?: number[];
    issueType?: string;
    labels?: string[];
    /**
     * The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
     */
    milestone?: string;
    /**
     * Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
     */
    myReactionEmoji?: string;
    /**
     * Return issues that do not match the assignee id.
     */
    notAssigneeIds?: number[];
    /**
     * Return issues that do not match the author id.
     */
    notAuthorIds?: number[];
    /**
     * Return issues that do not match the labels.
     */
    notLabels?: string[];
    /**
     * Return issues that do not match the milestone.
     */
    notMilestone?: string;
    /**
     * Return issues not reacted by the authenticated user by the given emoji.
     */
    notMyReactionEmojis?: string[];
    /**
     * Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
     */
    orderBy?: string;
    project: string;
    /**
     * Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
     */
    scope?: string;
    /**
     * Search project issues against their title and description
     */
    search?: string;
    /**
     * Return issues sorted in asc or desc order. Default is desc
     */
    sort?: string;
    /**
     * Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    updatedAfter?: string;
    /**
     * Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    updatedBefore?: string;
    weight?: number;
    /**
     * If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false. descriptionHtml was introduced in GitLab 12.7
     */
    withLabelsDetails?: boolean;
}

/**
 * A collection of values returned by getProjectIssues.
 */
export interface GetProjectIssuesResult {
    /**
     * Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
     */
    readonly assigneeId?: number;
    /**
     * Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
     */
    readonly assigneeUsername?: string;
    /**
     * Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
     */
    readonly authorId?: number;
    /**
     * Filter confidential or public issues.
     */
    readonly confidential?: boolean;
    /**
     * Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    readonly createdAfter?: string;
    /**
     * Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    readonly createdBefore?: string;
    /**
     * Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
     */
    readonly dueDate?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Return only the issues having the given iid
     */
    readonly iids?: number[];
    /**
     * Filter to a given type of issue. Valid values are [issue incident testCase]. (Introduced in GitLab 13.12)
     */
    readonly issueType?: string;
    /**
     * The list of issues returned by the search.
     */
    readonly issues: outputs.GetProjectIssuesIssue[];
    /**
     * Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
     */
    readonly labels?: string[];
    /**
     * The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
     */
    readonly milestone?: string;
    /**
     * Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
     */
    readonly myReactionEmoji?: string;
    /**
     * Return issues that do not match the assignee id.
     */
    readonly notAssigneeIds?: number[];
    /**
     * Return issues that do not match the author id.
     */
    readonly notAuthorIds?: number[];
    /**
     * Return issues that do not match the labels.
     */
    readonly notLabels?: string[];
    /**
     * Return issues that do not match the milestone.
     */
    readonly notMilestone?: string;
    /**
     * Return issues not reacted by the authenticated user by the given emoji.
     */
    readonly notMyReactionEmojis?: string[];
    /**
     * Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
     */
    readonly orderBy?: string;
    /**
     * The name or id of the project.
     */
    readonly project: string;
    /**
     * Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
     */
    readonly scope?: string;
    /**
     * Search project issues against their title and description
     */
    readonly search?: string;
    /**
     * Return issues sorted in asc or desc order. Default is desc
     */
    readonly sort?: string;
    /**
     * Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    readonly updatedAfter?: string;
    /**
     * Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    readonly updatedBefore?: string;
    /**
     * Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
     */
    readonly weight?: number;
    /**
     * If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false. descriptionHtml was introduced in GitLab 12.7
     */
    readonly withLabelsDetails?: boolean;
}
/**
 * The `gitlab.getProjectIssues` data source allows to retrieve details about issues in a project.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/issues.html)
 */
export function getProjectIssuesOutput(args: GetProjectIssuesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectIssuesResult> {
    return pulumi.output(args).apply((a: any) => getProjectIssues(a, opts))
}

/**
 * A collection of arguments for invoking getProjectIssues.
 */
export interface GetProjectIssuesOutputArgs {
    /**
     * Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
     */
    assigneeId?: pulumi.Input<number>;
    /**
     * Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assigneeUsername array should only contain a single value. Otherwise, an invalid parameter error is returned.
     */
    assigneeUsername?: pulumi.Input<string>;
    authorId?: pulumi.Input<number>;
    confidential?: pulumi.Input<boolean>;
    /**
     * Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    createdAfter?: pulumi.Input<string>;
    /**
     * Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    createdBefore?: pulumi.Input<string>;
    dueDate?: pulumi.Input<string>;
    /**
     * Return only the issues having the given iid
     */
    iids?: pulumi.Input<pulumi.Input<number>[]>;
    issueType?: pulumi.Input<string>;
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
     */
    milestone?: pulumi.Input<string>;
    /**
     * Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
     */
    myReactionEmoji?: pulumi.Input<string>;
    /**
     * Return issues that do not match the assignee id.
     */
    notAssigneeIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Return issues that do not match the author id.
     */
    notAuthorIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Return issues that do not match the labels.
     */
    notLabels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Return issues that do not match the milestone.
     */
    notMilestone?: pulumi.Input<string>;
    /**
     * Return issues not reacted by the authenticated user by the given emoji.
     */
    notMyReactionEmojis?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Return issues ordered by. Valid values are `createdAt`, `updatedAt`, `priority`, `dueDate`, `relativePosition`, `labelPriority`, `milestoneDue`, `popularity`, `weight`. Default is created_at
     */
    orderBy?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * Return issues for the given scope. Valid values are `createdByMe`, `assignedToMe`, `all`. Defaults to all.
     */
    scope?: pulumi.Input<string>;
    /**
     * Search project issues against their title and description
     */
    search?: pulumi.Input<string>;
    /**
     * Return issues sorted in asc or desc order. Default is desc
     */
    sort?: pulumi.Input<string>;
    /**
     * Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    updatedAfter?: pulumi.Input<string>;
    /**
     * Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     */
    updatedBefore?: pulumi.Input<string>;
    weight?: pulumi.Input<number>;
    /**
     * If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false. descriptionHtml was introduced in GitLab 12.7
     */
    withLabelsDetails?: pulumi.Input<boolean>;
}
