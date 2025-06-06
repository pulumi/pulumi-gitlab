// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getGroupSubgroups` data source allows to get subgroups of a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-subgroups)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * export = async () => {
 *     const subgroups = await gitlab.getGroupSubgroups({
 *         groupId: 123456,
 *     });
 *     return {
 *         subgroups: subgroups,
 *     };
 * }
 * ```
 */
export function getGroupSubgroups(args: GetGroupSubgroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupSubgroupsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getGroupSubgroups:getGroupSubgroups", {
        "allAvailable": args.allAvailable,
        "groupId": args.groupId,
        "minAccessLevel": args.minAccessLevel,
        "orderBy": args.orderBy,
        "owned": args.owned,
        "search": args.search,
        "skipGroups": args.skipGroups,
        "sort": args.sort,
        "statistics": args.statistics,
        "withCustomAttributes": args.withCustomAttributes,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupSubgroups.
 */
export interface GetGroupSubgroupsArgs {
    /**
     * Show all the groups you have access to.
     */
    allAvailable?: boolean;
    /**
     * The ID of the group.
     */
    groupId: number;
    /**
     * Limit to groups where current user has at least this access level.
     */
    minAccessLevel?: string;
    /**
     * Order groups by name, path or id.
     */
    orderBy?: string;
    /**
     * Limit to groups explicitly owned by the current user.
     */
    owned?: boolean;
    /**
     * Return the list of authorized groups matching the search criteria.
     */
    search?: string;
    /**
     * Skip the group IDs passed.
     */
    skipGroups?: number[];
    /**
     * Order groups in asc or desc order.
     */
    sort?: string;
    /**
     * Include group statistics (administrators only).
     */
    statistics?: boolean;
    /**
     * Include custom attributes in response (administrators only).
     */
    withCustomAttributes?: boolean;
}

/**
 * A collection of values returned by getGroupSubgroups.
 */
export interface GetGroupSubgroupsResult {
    /**
     * Show all the groups you have access to.
     */
    readonly allAvailable: boolean;
    /**
     * The ID of the group.
     */
    readonly groupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Limit to groups where current user has at least this access level.
     */
    readonly minAccessLevel: string;
    /**
     * Order groups by name, path or id.
     */
    readonly orderBy: string;
    /**
     * Limit to groups explicitly owned by the current user.
     */
    readonly owned: boolean;
    /**
     * Return the list of authorized groups matching the search criteria.
     */
    readonly search: string;
    /**
     * Skip the group IDs passed.
     */
    readonly skipGroups: number[];
    /**
     * Order groups in asc or desc order.
     */
    readonly sort: string;
    /**
     * Include group statistics (administrators only).
     */
    readonly statistics: boolean;
    /**
     * Subgroups of the parent group.
     */
    readonly subgroups: outputs.GetGroupSubgroupsSubgroup[];
    /**
     * Include custom attributes in response (administrators only).
     */
    readonly withCustomAttributes: boolean;
}
/**
 * The `gitlab.getGroupSubgroups` data source allows to get subgroups of a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-subgroups)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * export = async () => {
 *     const subgroups = await gitlab.getGroupSubgroups({
 *         groupId: 123456,
 *     });
 *     return {
 *         subgroups: subgroups,
 *     };
 * }
 * ```
 */
export function getGroupSubgroupsOutput(args: GetGroupSubgroupsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupSubgroupsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getGroupSubgroups:getGroupSubgroups", {
        "allAvailable": args.allAvailable,
        "groupId": args.groupId,
        "minAccessLevel": args.minAccessLevel,
        "orderBy": args.orderBy,
        "owned": args.owned,
        "search": args.search,
        "skipGroups": args.skipGroups,
        "sort": args.sort,
        "statistics": args.statistics,
        "withCustomAttributes": args.withCustomAttributes,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupSubgroups.
 */
export interface GetGroupSubgroupsOutputArgs {
    /**
     * Show all the groups you have access to.
     */
    allAvailable?: pulumi.Input<boolean>;
    /**
     * The ID of the group.
     */
    groupId: pulumi.Input<number>;
    /**
     * Limit to groups where current user has at least this access level.
     */
    minAccessLevel?: pulumi.Input<string>;
    /**
     * Order groups by name, path or id.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Limit to groups explicitly owned by the current user.
     */
    owned?: pulumi.Input<boolean>;
    /**
     * Return the list of authorized groups matching the search criteria.
     */
    search?: pulumi.Input<string>;
    /**
     * Skip the group IDs passed.
     */
    skipGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Order groups in asc or desc order.
     */
    sort?: pulumi.Input<string>;
    /**
     * Include group statistics (administrators only).
     */
    statistics?: pulumi.Input<boolean>;
    /**
     * Include custom attributes in response (administrators only).
     */
    withCustomAttributes?: pulumi.Input<boolean>;
}
