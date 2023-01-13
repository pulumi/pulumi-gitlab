// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ReleaseLink` data source allows get details of a release link.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // By project full path
 * const example = pulumi.output(gitlab.getReleaseLink({
 *     linkId: 11,
 *     project: "foo/bar",
 *     tagName: "v1.0.1",
 * }));
 * ```
 */
export function getReleaseLink(args: GetReleaseLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetReleaseLinkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getReleaseLink:getReleaseLink", {
        "linkId": args.linkId,
        "project": args.project,
        "tagName": args.tagName,
    }, opts);
}

/**
 * A collection of arguments for invoking getReleaseLink.
 */
export interface GetReleaseLinkArgs {
    /**
     * The ID of the link.
     */
    linkId: number;
    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
     */
    project: string;
    /**
     * The tag associated with the Release.
     */
    tagName: string;
}

/**
 * A collection of values returned by getReleaseLink.
 */
export interface GetReleaseLinkResult {
    /**
     * Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
     */
    readonly directAssetUrl: string;
    /**
     * External or internal link.
     */
    readonly external: boolean;
    /**
     * Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
     */
    readonly filepath: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of the link.
     */
    readonly linkId: number;
    /**
     * The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
     */
    readonly linkType: string;
    /**
     * The name of the link. Link names must be unique within the release.
     */
    readonly name: string;
    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
     */
    readonly project: string;
    /**
     * The tag associated with the Release.
     */
    readonly tagName: string;
    /**
     * The URL of the link. Link URLs must be unique within the release.
     */
    readonly url: string;
}

export function getReleaseLinkOutput(args: GetReleaseLinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReleaseLinkResult> {
    return pulumi.output(args).apply(a => getReleaseLink(a, opts))
}

/**
 * A collection of arguments for invoking getReleaseLink.
 */
export interface GetReleaseLinkOutputArgs {
    /**
     * The ID of the link.
     */
    linkId: pulumi.Input<number>;
    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
     */
    project: pulumi.Input<string>;
    /**
     * The tag associated with the Release.
     */
    tagName: pulumi.Input<string>;
}