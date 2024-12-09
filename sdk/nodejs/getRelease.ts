// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getRelease` data source retrieves information about a gitlab release for a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // By project ID and tag_name
 * const example = gitlab.getRelease({
 *     projectId: "1234",
 *     tagName: "v1.0",
 * });
 * ```
 */
export function getRelease(args: GetReleaseArgs, opts?: pulumi.InvokeOptions): Promise<GetReleaseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getRelease:getRelease", {
        "assets": args.assets,
        "projectId": args.projectId,
        "tagName": args.tagName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRelease.
 */
export interface GetReleaseArgs {
    /**
     * The assets for a release
     */
    assets?: inputs.GetReleaseAssets;
    /**
     * The ID or URL-encoded path of the project.
     */
    projectId: string;
    /**
     * The Git tag the release is associated with.
     */
    tagName: string;
}

/**
 * A collection of values returned by getRelease.
 */
export interface GetReleaseResult {
    /**
     * The assets for a release
     */
    readonly assets?: outputs.GetReleaseAssets;
    /**
     * The date the release was created.
     */
    readonly createdAt: string;
    /**
     * An HTML rendered description of the release.
     */
    readonly description: string;
    readonly id: string;
    /**
     * The name of the release.
     */
    readonly name: string;
    /**
     * The ID or URL-encoded path of the project.
     */
    readonly projectId: string;
    /**
     * The date the release was created.
     */
    readonly releasedAt: string;
    /**
     * The Git tag the release is associated with.
     */
    readonly tagName: string;
}
/**
 * The `gitlab.getRelease` data source retrieves information about a gitlab release for a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // By project ID and tag_name
 * const example = gitlab.getRelease({
 *     projectId: "1234",
 *     tagName: "v1.0",
 * });
 * ```
 */
export function getReleaseOutput(args: GetReleaseOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetReleaseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getRelease:getRelease", {
        "assets": args.assets,
        "projectId": args.projectId,
        "tagName": args.tagName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRelease.
 */
export interface GetReleaseOutputArgs {
    /**
     * The assets for a release
     */
    assets?: pulumi.Input<inputs.GetReleaseAssetsArgs>;
    /**
     * The ID or URL-encoded path of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Git tag the release is associated with.
     */
    tagName: pulumi.Input<string>;
}
