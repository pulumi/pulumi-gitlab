// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
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
 * const mygroup = gitlab.getGroup({
 *     fullPath: "mygroup",
 * });
 * const groupProjects = mygroup.then(mygroup => gitlab.getProjects({
 *     groupId: mygroup.id,
 *     orderBy: "name",
 *     includeSubgroups: true,
 *     withShared: false,
 * }));
 * const projects = gitlab.getProjects({
 *     search: "postgresql",
 *     visibility: "private",
 * });
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
    archived?: boolean;
    groupId?: number;
    includeSubgroups?: boolean;
    maxQueryablePages?: number;
    membership?: boolean;
    minAccessLevel?: number;
    orderBy?: string;
    owned?: boolean;
    page?: number;
    perPage?: number;
    search?: string;
    simple?: boolean;
    sort?: string;
    starred?: boolean;
    statistics?: boolean;
    visibility?: string;
    withCustomAttributes?: boolean;
    withIssuesEnabled?: boolean;
    withMergeRequestsEnabled?: boolean;
    withProgrammingLanguage?: string;
    withShared?: boolean;
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
    readonly projects: outputs.GetProjectsProject[];
    readonly search?: string;
    readonly simple?: boolean;
    readonly sort?: string;
    readonly starred?: boolean;
    readonly statistics?: boolean;
    readonly visibility?: string;
    readonly withCustomAttributes?: boolean;
    readonly withIssuesEnabled?: boolean;
    readonly withMergeRequestsEnabled?: boolean;
    readonly withProgrammingLanguage?: string;
    readonly withShared?: boolean;
}

export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    return pulumi.output(args).apply(a => getProjects(a, opts))
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsOutputArgs {
    archived?: pulumi.Input<boolean>;
    groupId?: pulumi.Input<number>;
    includeSubgroups?: pulumi.Input<boolean>;
    maxQueryablePages?: pulumi.Input<number>;
    membership?: pulumi.Input<boolean>;
    minAccessLevel?: pulumi.Input<number>;
    orderBy?: pulumi.Input<string>;
    owned?: pulumi.Input<boolean>;
    page?: pulumi.Input<number>;
    perPage?: pulumi.Input<number>;
    search?: pulumi.Input<string>;
    simple?: pulumi.Input<boolean>;
    sort?: pulumi.Input<string>;
    starred?: pulumi.Input<boolean>;
    statistics?: pulumi.Input<boolean>;
    visibility?: pulumi.Input<string>;
    withCustomAttributes?: pulumi.Input<boolean>;
    withIssuesEnabled?: pulumi.Input<boolean>;
    withMergeRequestsEnabled?: pulumi.Input<boolean>;
    withProgrammingLanguage?: pulumi.Input<string>;
    withShared?: pulumi.Input<boolean>;
}
