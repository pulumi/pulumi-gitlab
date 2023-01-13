// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectHooks` data source allows to retrieve details about hooks in a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-project-hooks)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProject({
 *     id: "foo/bar/baz",
 * });
 * const examples = example.then(example => gitlab.getProjectHooks({
 *     project: example.id,
 * }));
 * ```
 */
export function getProjectHooks(args: GetProjectHooksArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectHooksResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getProjectHooks:getProjectHooks", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectHooks.
 */
export interface GetProjectHooksArgs {
    /**
     * The name or id of the project.
     */
    project: string;
}

/**
 * A collection of values returned by getProjectHooks.
 */
export interface GetProjectHooksResult {
    /**
     * The list of hooks.
     */
    readonly hooks: outputs.GetProjectHooksHook[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name or id of the project.
     */
    readonly project: string;
}

export function getProjectHooksOutput(args: GetProjectHooksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectHooksResult> {
    return pulumi.output(args).apply(a => getProjectHooks(a, opts))
}

/**
 * A collection of arguments for invoking getProjectHooks.
 */
export interface GetProjectHooksOutputArgs {
    /**
     * The name or id of the project.
     */
    project: pulumi.Input<string>;
}