// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides details about a specific protected branch in a given project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = pulumi.output(gitlab.getProjectProtectedBranch({
 *     name: "main",
 *     projectId: "foo/bar/baz",
 * }));
 * ```
 */
export function getProjectProtectedBranch(args: GetProjectProtectedBranchArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectProtectedBranchResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getProjectProtectedBranch:getProjectProtectedBranch", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectProtectedBranch.
 */
export interface GetProjectProtectedBranchArgs {
    /**
     * The name of the protected branch.
     */
    name: string;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    projectId: string;
}

/**
 * A collection of values returned by getProjectProtectedBranch.
 */
export interface GetProjectProtectedBranchResult {
    /**
     * Whether force push is allowed.
     */
    readonly allowForcePush: boolean;
    /**
     * Reject code pushes that change files listed in the CODEOWNERS file.
     */
    readonly codeOwnerApprovalRequired: boolean;
    /**
     * The ID of this resource.
     */
    readonly id: number;
    /**
     * Describes which access levels, users, or groups are allowed to perform the action.
     */
    readonly mergeAccessLevels: outputs.GetProjectProtectedBranchMergeAccessLevel[];
    /**
     * The name of the protected branch.
     */
    readonly name: string;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    readonly projectId: string;
    /**
     * Describes which access levels, users, or groups are allowed to perform the action.
     */
    readonly pushAccessLevels: outputs.GetProjectProtectedBranchPushAccessLevel[];
}

export function getProjectProtectedBranchOutput(args: GetProjectProtectedBranchOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectProtectedBranchResult> {
    return pulumi.output(args).apply(a => getProjectProtectedBranch(a, opts))
}

/**
 * A collection of arguments for invoking getProjectProtectedBranch.
 */
export interface GetProjectProtectedBranchOutputArgs {
    /**
     * The name of the protected branch.
     */
    name: pulumi.Input<string>;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    projectId: pulumi.Input<string>;
}
