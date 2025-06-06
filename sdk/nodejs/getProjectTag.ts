// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectTag` data source allows details of a project tag to be retrieved by its name.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/tags/)
 */
export function getProjectTag(args: GetProjectTagArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectTagResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectTag:getProjectTag", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectTag.
 */
export interface GetProjectTagArgs {
    /**
     * The name of a tag.
     */
    name: string;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: string;
}

/**
 * A collection of values returned by getProjectTag.
 */
export interface GetProjectTagResult {
    /**
     * The commit associated with the tag.
     */
    readonly commits: outputs.GetProjectTagCommit[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The message of the annotated tag.
     */
    readonly message: string;
    /**
     * The name of a tag.
     */
    readonly name: string;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    readonly project: string;
    /**
     * Bool, true if tag has tag protection.
     */
    readonly protected: boolean;
    /**
     * The release associated with the tag.
     */
    readonly releases: outputs.GetProjectTagRelease[];
    /**
     * The unique id assigned to the commit by Gitlab.
     */
    readonly target: string;
}
/**
 * The `gitlab.ProjectTag` data source allows details of a project tag to be retrieved by its name.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/tags/)
 */
export function getProjectTagOutput(args: GetProjectTagOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProjectTagResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getProjectTag:getProjectTag", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectTag.
 */
export interface GetProjectTagOutputArgs {
    /**
     * The name of a tag.
     */
    name: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
}
