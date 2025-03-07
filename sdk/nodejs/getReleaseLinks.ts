// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getReleaseLinks` data source allows get details of release links.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/releases/links/)
 */
export function getReleaseLinks(args: GetReleaseLinksArgs, opts?: pulumi.InvokeOptions): Promise<GetReleaseLinksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getReleaseLinks:getReleaseLinks", {
        "project": args.project,
        "tagName": args.tagName,
    }, opts);
}

/**
 * A collection of arguments for invoking getReleaseLinks.
 */
export interface GetReleaseLinksArgs {
    /**
     * The ID or full path to the project.
     */
    project: string;
    /**
     * The tag associated with the Release.
     */
    tagName: string;
}

/**
 * A collection of values returned by getReleaseLinks.
 */
export interface GetReleaseLinksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID or full path to the project.
     */
    readonly project: string;
    /**
     * List of release links
     */
    readonly releaseLinks: outputs.GetReleaseLinksReleaseLink[];
    /**
     * The tag associated with the Release.
     */
    readonly tagName: string;
}
/**
 * The `gitlab.getReleaseLinks` data source allows get details of release links.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/releases/links/)
 */
export function getReleaseLinksOutput(args: GetReleaseLinksOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetReleaseLinksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getReleaseLinks:getReleaseLinks", {
        "project": args.project,
        "tagName": args.tagName,
    }, opts);
}

/**
 * A collection of arguments for invoking getReleaseLinks.
 */
export interface GetReleaseLinksOutputArgs {
    /**
     * The ID or full path to the project.
     */
    project: pulumi.Input<string>;
    /**
     * The tag associated with the Release.
     */
    tagName: pulumi.Input<string>;
}
