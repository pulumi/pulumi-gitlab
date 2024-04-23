// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.
 *
 * > This data source supports all available filters exposed by the xanzy/go-gitlab package, which might not expose all available filters exposed by the Gitlab APIs.
 *
 * > The owner sub-attributes are only populated if the Gitlab token used has an administrator scope.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // List projects within a group tree
 * const mygroup = gitlab.getGroup({
 *     fullPath: "mygroup",
 * });
 * const groupProjects = mygroup.then(mygroup => gitlab.getProjects({
 *     groupId: mygroup.id,
 *     orderBy: "name",
 *     includeSubgroups: true,
 *     withShared: false,
 * }));
 * // List projects using the search syntax
 * const projects = gitlab.getProjects({
 *     search: "postgresql",
 *     visibility: "private",
 * });
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
        "topics": args.topics,
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
    archived?: boolean;
    /**
     * The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
     */
    groupId?: number;
    /**
     * Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
     */
    includeSubgroups?: boolean;
    /**
     * The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
     */
    maxQueryablePages?: number;
    /**
     * Limit by projects that the current user is a member of.
     */
    membership?: boolean;
    /**
     * Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `groupId`.
     */
    minAccessLevel?: number;
    /**
     * Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     */
    orderBy?: string;
    /**
     * Limit by projects owned by the current user.
     */
    owned?: boolean;
    /**
     * The first page to begin the query on.
     */
    page?: number;
    /**
     * The number of results to return per page.
     */
    perPage?: number;
    /**
     * Return list of authorized projects matching the search criteria.
     */
    search?: string;
    /**
     * Return only the ID, URL, name, and path of each project.
     */
    simple?: boolean;
    /**
     * Return projects sorted in `asc` or `desc` order. Default is `desc`.
     */
    sort?: string;
    /**
     * Limit by projects starred by the current user.
     */
    starred?: boolean;
    /**
     * Include project statistics. Cannot be used with `groupId`.
     */
    statistics?: boolean;
    /**
     * Limit by projects that have all of the given topics.
     */
    topics?: string[];
    /**
     * Limit by visibility `public`, `internal`, or `private`.
     */
    visibility?: string;
    /**
     * Include custom attributes in response *(admins only)*.
     */
    withCustomAttributes?: boolean;
    /**
     * Limit by projects with issues feature enabled. Default is `false`.
     */
    withIssuesEnabled?: boolean;
    /**
     * Limit by projects with merge requests feature enabled. Default is `false`.
     */
    withMergeRequestsEnabled?: boolean;
    /**
     * Limit by projects which use the given programming language. Cannot be used with `groupId`.
     */
    withProgrammingLanguage?: string;
    /**
     * Include projects shared to this group. Default is `true`. Needs `groupId`.
     */
    withShared?: boolean;
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    /**
     * Limit by archived status.
     */
    readonly archived?: boolean;
    /**
     * The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
     */
    readonly groupId?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
     */
    readonly includeSubgroups?: boolean;
    /**
     * The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
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
     * Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     */
    readonly orderBy?: string;
    /**
     * Limit by projects owned by the current user.
     */
    readonly owned?: boolean;
    /**
     * The first page to begin the query on.
     */
    readonly page?: number;
    /**
     * The number of results to return per page.
     */
    readonly perPage?: number;
    /**
     * A list containing the projects matching the supplied arguments
     */
    readonly projects: outputs.GetProjectsProject[];
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
     * Limit by projects that have all of the given topics.
     */
    readonly topics?: string[];
    /**
     * Limit by visibility `public`, `internal`, or `private`.
     */
    readonly visibility?: string;
    /**
     * Include custom attributes in response *(admins only)*.
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
 * The `gitlab.getProjects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.
 *
 * > This data source supports all available filters exposed by the xanzy/go-gitlab package, which might not expose all available filters exposed by the Gitlab APIs.
 *
 * > The owner sub-attributes are only populated if the Gitlab token used has an administrator scope.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // List projects within a group tree
 * const mygroup = gitlab.getGroup({
 *     fullPath: "mygroup",
 * });
 * const groupProjects = mygroup.then(mygroup => gitlab.getProjects({
 *     groupId: mygroup.id,
 *     orderBy: "name",
 *     includeSubgroups: true,
 *     withShared: false,
 * }));
 * // List projects using the search syntax
 * const projects = gitlab.getProjects({
 *     search: "postgresql",
 *     visibility: "private",
 * });
 * ```
 */
export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    return pulumi.output(args).apply((a: any) => getProjects(a, opts))
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsOutputArgs {
    /**
     * Limit by archived status.
     */
    archived?: pulumi.Input<boolean>;
    /**
     * The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
     */
    includeSubgroups?: pulumi.Input<boolean>;
    /**
     * The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
     */
    maxQueryablePages?: pulumi.Input<number>;
    /**
     * Limit by projects that the current user is a member of.
     */
    membership?: pulumi.Input<boolean>;
    /**
     * Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `groupId`.
     */
    minAccessLevel?: pulumi.Input<number>;
    /**
     * Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Limit by projects owned by the current user.
     */
    owned?: pulumi.Input<boolean>;
    /**
     * The first page to begin the query on.
     */
    page?: pulumi.Input<number>;
    /**
     * The number of results to return per page.
     */
    perPage?: pulumi.Input<number>;
    /**
     * Return list of authorized projects matching the search criteria.
     */
    search?: pulumi.Input<string>;
    /**
     * Return only the ID, URL, name, and path of each project.
     */
    simple?: pulumi.Input<boolean>;
    /**
     * Return projects sorted in `asc` or `desc` order. Default is `desc`.
     */
    sort?: pulumi.Input<string>;
    /**
     * Limit by projects starred by the current user.
     */
    starred?: pulumi.Input<boolean>;
    /**
     * Include project statistics. Cannot be used with `groupId`.
     */
    statistics?: pulumi.Input<boolean>;
    /**
     * Limit by projects that have all of the given topics.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Limit by visibility `public`, `internal`, or `private`.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Include custom attributes in response *(admins only)*.
     */
    withCustomAttributes?: pulumi.Input<boolean>;
    /**
     * Limit by projects with issues feature enabled. Default is `false`.
     */
    withIssuesEnabled?: pulumi.Input<boolean>;
    /**
     * Limit by projects with merge requests feature enabled. Default is `false`.
     */
    withMergeRequestsEnabled?: pulumi.Input<boolean>;
    /**
     * Limit by projects which use the given programming language. Cannot be used with `groupId`.
     */
    withProgrammingLanguage?: pulumi.Input<string>;
    /**
     * Include projects shared to this group. Default is `true`. Needs `groupId`.
     */
    withShared?: pulumi.Input<boolean>;
}
