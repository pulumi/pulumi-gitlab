// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectTags` data source allows details of project tags to be retrieved by some search criteria.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)
 */
export function getProjectTags(args: GetProjectTagsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectTagsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectTags:getProjectTags", {
        "orderBy": args.orderBy,
        "project": args.project,
        "search": args.search,
        "sort": args.sort,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectTags.
 */
export interface GetProjectTagsArgs {
    /**
     * Return tags ordered by `name` or `updated` fields. Default is `updated`.
     */
    orderBy?: string;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: string;
    /**
     * Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
     */
    search?: string;
    /**
     * Return tags sorted in `asc` or `desc` order. Default is `desc`.
     */
    sort?: string;
}

/**
 * A collection of values returned by getProjectTags.
 */
export interface GetProjectTagsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Return tags ordered by `name` or `updated` fields. Default is `updated`.
     */
    readonly orderBy?: string;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    readonly project: string;
    /**
     * Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
     */
    readonly search?: string;
    /**
     * Return tags sorted in `asc` or `desc` order. Default is `desc`.
     */
    readonly sort?: string;
    /**
     * List of repository tags from a project.
     */
    readonly tags: outputs.GetProjectTagsTag[];
}
/**
 * The `gitlab.getProjectTags` data source allows details of project tags to be retrieved by some search criteria.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/tags.html#list-project-repository-tags)
 */
export function getProjectTagsOutput(args: GetProjectTagsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProjectTagsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getProjectTags:getProjectTags", {
        "orderBy": args.orderBy,
        "project": args.project,
        "search": args.search,
        "sort": args.sort,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectTags.
 */
export interface GetProjectTagsOutputArgs {
    /**
     * Return tags ordered by `name` or `updated` fields. Default is `updated`.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
    /**
     * Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
     */
    search?: pulumi.Input<string>;
    /**
     * Return tags sorted in `asc` or `desc` order. Default is `desc`.
     */
    sort?: pulumi.Input<string>;
}
