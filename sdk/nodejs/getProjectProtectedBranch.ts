// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectProtectedBranch` data source allows details of a protected branch to be retrieved by its name and the project it belongs to.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProjectProtectedBranch({
 *     name: "main",
 *     projectId: "foo/bar/baz",
 * });
 * ```
 */
export function getProjectProtectedBranch(args: GetProjectProtectedBranchArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectProtectedBranchResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectProtectedBranch:getProjectProtectedBranch", {
        "mergeAccessLevels": args.mergeAccessLevels,
        "name": args.name,
        "projectId": args.projectId,
        "pushAccessLevels": args.pushAccessLevels,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectProtectedBranch.
 */
export interface GetProjectProtectedBranchArgs {
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    mergeAccessLevels?: inputs.GetProjectProtectedBranchMergeAccessLevel[];
    /**
     * The name of the protected branch.
     */
    name: string;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    projectId: string;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    pushAccessLevels?: inputs.GetProjectProtectedBranchPushAccessLevel[];
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
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    readonly mergeAccessLevels?: outputs.GetProjectProtectedBranchMergeAccessLevel[];
    /**
     * The name of the protected branch.
     */
    readonly name: string;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    readonly projectId: string;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    readonly pushAccessLevels?: outputs.GetProjectProtectedBranchPushAccessLevel[];
}
/**
 * The `gitlab.getProjectProtectedBranch` data source allows details of a protected branch to be retrieved by its name and the project it belongs to.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProjectProtectedBranch({
 *     name: "main",
 *     projectId: "foo/bar/baz",
 * });
 * ```
 */
export function getProjectProtectedBranchOutput(args: GetProjectProtectedBranchOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectProtectedBranchResult> {
    return pulumi.output(args).apply((a: any) => getProjectProtectedBranch(a, opts))
}

/**
 * A collection of arguments for invoking getProjectProtectedBranch.
 */
export interface GetProjectProtectedBranchOutputArgs {
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    mergeAccessLevels?: pulumi.Input<pulumi.Input<inputs.GetProjectProtectedBranchMergeAccessLevelArgs>[]>;
    /**
     * The name of the protected branch.
     */
    name: pulumi.Input<string>;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    pushAccessLevels?: pulumi.Input<pulumi.Input<inputs.GetProjectProtectedBranchPushAccessLevelArgs>[]>;
}
