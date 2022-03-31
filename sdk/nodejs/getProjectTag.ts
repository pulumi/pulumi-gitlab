// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectTag` data source allows details of a project tag to be retrieved by its name.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/tags.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // By project full path
 * const foo = pulumi.output(gitlab.getProjectTag({
 *     name: "example",
 *     project: "foo/bar",
 * }));
 * ```
 */
export function getProjectTag(args: GetProjectTagArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectTagResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getProjectTag:getProjectTag", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectTag.
 */
export interface GetProjectTagArgs {
    name: string;
    project: string;
}

/**
 * A collection of values returned by getProjectTag.
 */
export interface GetProjectTagResult {
    readonly commits: outputs.GetProjectTagCommit[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly message: string;
    readonly name: string;
    readonly project: string;
    readonly protected: boolean;
    readonly releases: outputs.GetProjectTagRelease[];
    readonly target: string;
}

export function getProjectTagOutput(args: GetProjectTagOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectTagResult> {
    return pulumi.output(args).apply(a => getProjectTag(a, opts))
}

/**
 * A collection of arguments for invoking getProjectTag.
 */
export interface GetProjectTagOutputArgs {
    name: pulumi.Input<string>;
    project: pulumi.Input<string>;
}
