// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getGroups` data source allows details of multiple groups to be retrieved given some optional filter criteria.
 *
 * > Some attributes might not be returned depending on if you're an admin or not.
 *
 * > Some available options require administrator privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-groups)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getGroups({
 *     sort: "desc",
 *     orderBy: "name",
 * });
 * const example_two = gitlab.getGroups({
 *     search: "GitLab",
 * });
 * ```
 */
export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getGroups:getGroups", {
        "orderBy": args.orderBy,
        "search": args.search,
        "sort": args.sort,
        "topLevelOnly": args.topLevelOnly,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    /**
     * Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
     */
    orderBy?: string;
    /**
     * Search groups by name or path.
     */
    search?: string;
    /**
     * Sort groups' list in asc or desc order. (Requires administrator privileges)
     */
    sort?: string;
    /**
     * Limit to top level groups, excluding all subgroups.
     */
    topLevelOnly?: boolean;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    /**
     * The list of groups.
     */
    readonly groups: outputs.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
     */
    readonly orderBy?: string;
    /**
     * Search groups by name or path.
     */
    readonly search?: string;
    /**
     * Sort groups' list in asc or desc order. (Requires administrator privileges)
     */
    readonly sort?: string;
    /**
     * Limit to top level groups, excluding all subgroups.
     */
    readonly topLevelOnly?: boolean;
}
/**
 * The `gitlab.getGroups` data source allows details of multiple groups to be retrieved given some optional filter criteria.
 *
 * > Some attributes might not be returned depending on if you're an admin or not.
 *
 * > Some available options require administrator privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#list-groups)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getGroups({
 *     sort: "desc",
 *     orderBy: "name",
 * });
 * const example_two = gitlab.getGroups({
 *     search: "GitLab",
 * });
 * ```
 */
export function getGroupsOutput(args?: GetGroupsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getGroups:getGroups", {
        "orderBy": args.orderBy,
        "search": args.search,
        "sort": args.sort,
        "topLevelOnly": args.topLevelOnly,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsOutputArgs {
    /**
     * Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Search groups by name or path.
     */
    search?: pulumi.Input<string>;
    /**
     * Sort groups' list in asc or desc order. (Requires administrator privileges)
     */
    sort?: pulumi.Input<string>;
    /**
     * Limit to top level groups, excluding all subgroups.
     */
    topLevelOnly?: pulumi.Input<boolean>;
}
