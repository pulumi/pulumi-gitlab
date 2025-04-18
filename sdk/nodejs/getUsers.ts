// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getUsers` data source allows details of multiple users to be retrieved given some optional filter criteria.
 *
 * > Some attributes might not be returned depending on if you're an admin or not.
 *
 * > Some available options require administrator privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/users/#list-users)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getUsers({
 *     sort: "desc",
 *     orderBy: "name",
 *     createdBefore: "2019-01-01",
 * });
 * const example_two = gitlab.getUsers({
 *     search: "username",
 * });
 * ```
 */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getUsers:getUsers", {
        "active": args.active,
        "blocked": args.blocked,
        "createdAfter": args.createdAfter,
        "createdBefore": args.createdBefore,
        "excludeExternal": args.excludeExternal,
        "excludeInternal": args.excludeInternal,
        "externProvider": args.externProvider,
        "externUid": args.externUid,
        "external": args.external,
        "orderBy": args.orderBy,
        "search": args.search,
        "sort": args.sort,
        "username": args.username,
        "withoutProjectBots": args.withoutProjectBots,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * Filter users that are active.
     */
    active?: boolean;
    /**
     * Filter users that are blocked.
     */
    blocked?: boolean;
    /**
     * Search for users created after a specific date. (Requires administrator privileges)
     */
    createdAfter?: string;
    /**
     * Search for users created before a specific date. (Requires administrator privileges)
     */
    createdBefore?: string;
    /**
     * Filters only non external users.
     */
    excludeExternal?: boolean;
    /**
     * Filters only non internal users.
     */
    excludeInternal?: boolean;
    /**
     * Lookup users by external provider. (Requires administrator privileges)
     */
    externProvider?: string;
    /**
     * Lookup users by external UID. (Requires administrator privileges)
     */
    externUid?: string;
    /**
     * Filters only external users.
     */
    external?: boolean;
    /**
     * Order the users' list by `id`, `name`, `username`, `createdAt` or `updatedAt`. (Requires administrator privileges)
     */
    orderBy?: string;
    /**
     * Search users by username, name or email.
     */
    search?: string;
    /**
     * Sort users' list in asc or desc order. (Requires administrator privileges)
     */
    sort?: string;
    /**
     * Get a single user with a specific username.
     */
    username?: string;
    /**
     * Filters user without project bots.
     */
    withoutProjectBots?: boolean;
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * Filter users that are active.
     */
    readonly active?: boolean;
    /**
     * Filter users that are blocked.
     */
    readonly blocked?: boolean;
    /**
     * Search for users created after a specific date. (Requires administrator privileges)
     */
    readonly createdAfter?: string;
    /**
     * Search for users created before a specific date. (Requires administrator privileges)
     */
    readonly createdBefore?: string;
    /**
     * Filters only non external users.
     */
    readonly excludeExternal?: boolean;
    /**
     * Filters only non internal users.
     */
    readonly excludeInternal?: boolean;
    /**
     * Lookup users by external provider. (Requires administrator privileges)
     */
    readonly externProvider?: string;
    /**
     * Lookup users by external UID. (Requires administrator privileges)
     */
    readonly externUid?: string;
    /**
     * Filters only external users.
     */
    readonly external?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Order the users' list by `id`, `name`, `username`, `createdAt` or `updatedAt`. (Requires administrator privileges)
     */
    readonly orderBy?: string;
    /**
     * Search users by username, name or email.
     */
    readonly search?: string;
    /**
     * Sort users' list in asc or desc order. (Requires administrator privileges)
     */
    readonly sort?: string;
    /**
     * Get a single user with a specific username.
     */
    readonly username?: string;
    /**
     * The list of users.
     */
    readonly users: outputs.GetUsersUser[];
    /**
     * Filters user without project bots.
     */
    readonly withoutProjectBots?: boolean;
}
/**
 * The `gitlab.getUsers` data source allows details of multiple users to be retrieved given some optional filter criteria.
 *
 * > Some attributes might not be returned depending on if you're an admin or not.
 *
 * > Some available options require administrator privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/users/#list-users)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getUsers({
 *     sort: "desc",
 *     orderBy: "name",
 *     createdBefore: "2019-01-01",
 * });
 * const example_two = gitlab.getUsers({
 *     search: "username",
 * });
 * ```
 */
export function getUsersOutput(args?: GetUsersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUsersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getUsers:getUsers", {
        "active": args.active,
        "blocked": args.blocked,
        "createdAfter": args.createdAfter,
        "createdBefore": args.createdBefore,
        "excludeExternal": args.excludeExternal,
        "excludeInternal": args.excludeInternal,
        "externProvider": args.externProvider,
        "externUid": args.externUid,
        "external": args.external,
        "orderBy": args.orderBy,
        "search": args.search,
        "sort": args.sort,
        "username": args.username,
        "withoutProjectBots": args.withoutProjectBots,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * Filter users that are active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Filter users that are blocked.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * Search for users created after a specific date. (Requires administrator privileges)
     */
    createdAfter?: pulumi.Input<string>;
    /**
     * Search for users created before a specific date. (Requires administrator privileges)
     */
    createdBefore?: pulumi.Input<string>;
    /**
     * Filters only non external users.
     */
    excludeExternal?: pulumi.Input<boolean>;
    /**
     * Filters only non internal users.
     */
    excludeInternal?: pulumi.Input<boolean>;
    /**
     * Lookup users by external provider. (Requires administrator privileges)
     */
    externProvider?: pulumi.Input<string>;
    /**
     * Lookup users by external UID. (Requires administrator privileges)
     */
    externUid?: pulumi.Input<string>;
    /**
     * Filters only external users.
     */
    external?: pulumi.Input<boolean>;
    /**
     * Order the users' list by `id`, `name`, `username`, `createdAt` or `updatedAt`. (Requires administrator privileges)
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Search users by username, name or email.
     */
    search?: pulumi.Input<string>;
    /**
     * Sort users' list in asc or desc order. (Requires administrator privileges)
     */
    sort?: pulumi.Input<string>;
    /**
     * Get a single user with a specific username.
     */
    username?: pulumi.Input<string>;
    /**
     * Filters user without project bots.
     */
    withoutProjectBots?: pulumi.Input<boolean>;
}
