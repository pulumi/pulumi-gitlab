// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides details about a list of projects in the Gitlab provider. Listing all projects and group projects with [project filtering](https://docs.gitlab.com/ee/api/projects.html#list-user-projects) or [group project filtering](https://docs.gitlab.com/ee/api/groups.html#list-a-groups-projects) is supported.
 *
 * > NOTE: This data source supports all available filters exposed by the `xanzy/go-gitlab` package, which might not expose all available filters exposed by the Gitlab APIs.   
 *
 * ## Example Usage
 *
 * ### List projects within a group tree
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const mygroup = gitlab.getGroup({
 *     fullPath: "mygroup",
 * });
 * const groupProjects = mygroup.then(mygroup => gitlab.getProjects({
 *     groupId: mygroup.id,
 *     orderBy: "name",
 *     includeSubgroups: true,
 *     withShared: false,
 * }));
 * ```
 *
 * ### List projects using the search syntax
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const projects = pulumi.output(gitlab.getProjects({
 *     search: "postgresql",
 *     visibility: "private",
 * }, { async: true }));
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gitlab:index/getProjects:getProjects", {
        "archived": args.archived,
        "groupId": args.groupId,
        "includeSubgroups": args.includeSubgroups,
        "maxQueryablePages": args.maxQueryablePages,
        "membership": args.membership,
        "minAccessLevel": args.minAccessLevel,
        "orderBy": args.orderBy,
        "owned": args.owned,
        "page": args.page,
        "perPage": args.perPage,
        "search": args.search,
        "simple": args.simple,
        "sort": args.sort,
        "starred": args.starred,
        "statistics": args.statistics,
        "visibility": args.visibility,
        "withCustomAttributes": args.withCustomAttributes,
        "withIssuesEnabled": args.withIssuesEnabled,
        "withMergeRequestsEnabled": args.withMergeRequestsEnabled,
        "withProgrammingLanguage": args.withProgrammingLanguage,
        "withShared": args.withShared,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * Limit by archived status.
     */
    readonly archived?: boolean;
    /**
     * The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
     */
    readonly groupId?: number;
    /**
     * Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
     */
    readonly includeSubgroups?: boolean;
    /**
     * Prevents overloading your Gitlab instance in case of a misconfiguration. Default is `10`.
     */
    readonly maxQueryablePages?: number;
    /**
     * Limit by projects that the current user is a member of.
     */
    readonly membership?: boolean;
    /**
     * Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `groupId`.
     */
    readonly minAccessLevel?: number;
    /**
     * Return projects ordered by `id`, `name`, `path`, `createdAt`, `updatedAt`, or `lastActivityAt` fields. Default is `createdAt`.
     */
    readonly orderBy?: string;
    /**
     * Limit by projects owned by the current user.
     */
    readonly owned?: boolean;
    readonly page?: number;
    readonly perPage?: number;
    /**
     * Return list of authorized projects matching the search criteria.
     */
    readonly search?: string;
    /**
     * Return only the ID, URL, name, and path of each project.
     */
    readonly simple?: boolean;
    /**
     * Return projects sorted in `asc` or `desc` order. Default is `desc`.
     */
    readonly sort?: string;
    /**
     * Limit by projects starred by the current user.
     */
    readonly starred?: boolean;
    /**
     * Include project statistics. Cannot be used with `groupId`.
     */
    readonly statistics?: boolean;
    /**
     * Limit by visibility `public`, `internal`, or `private`.
     */
    readonly visibility?: string;
    /**
     * Include custom attributes in response _(admins only)_.
     */
    readonly withCustomAttributes?: boolean;
    /**
     * Limit by projects with issues feature enabled. Default is `false`.
     */
    readonly withIssuesEnabled?: boolean;
    /**
     * Limit by projects with merge requests feature enabled. Default is `false`.
     */
    readonly withMergeRequestsEnabled?: boolean;
    /**
     * Limit by projects which use the given programming language. Cannot be used with `groupId`.
     */
    readonly withProgrammingLanguage?: string;
    /**
     * Include projects shared to this group. Default is `true`. Needs `groupId`.
     */
    readonly withShared?: boolean;
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    readonly archived?: boolean;
    readonly groupId?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeSubgroups?: boolean;
    readonly maxQueryablePages?: number;
    readonly membership?: boolean;
    readonly minAccessLevel?: number;
    readonly orderBy?: string;
    readonly owned?: boolean;
    readonly page?: number;
    readonly perPage?: number;
    /**
     * A list containing the projects matching the supplied arguments
     */
    readonly projects: outputs.GetProjectsProject[];
    readonly search?: string;
    readonly simple?: boolean;
    readonly sort?: string;
    readonly starred?: boolean;
    readonly statistics?: boolean;
    /**
     * The visibility of the project.
     */
    readonly visibility?: string;
    readonly withCustomAttributes?: boolean;
    readonly withIssuesEnabled?: boolean;
    readonly withMergeRequestsEnabled?: boolean;
    readonly withProgrammingLanguage?: string;
    readonly withShared?: boolean;
}
