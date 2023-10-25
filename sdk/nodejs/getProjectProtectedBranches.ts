// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectProtectedBranches` data source allows details of the protected branches of a given project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#list-protected-branches)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProjectProtectedBranches({
 *     projectId: "foo/bar/baz",
 * });
 * ```
 */
export function getProjectProtectedBranches(args: GetProjectProtectedBranchesArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectProtectedBranchesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectProtectedBranches:getProjectProtectedBranches", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectProtectedBranches.
 */
export interface GetProjectProtectedBranchesArgs {
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    projectId: string;
}

/**
 * A collection of values returned by getProjectProtectedBranches.
 */
export interface GetProjectProtectedBranchesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    readonly projectId: string;
    /**
     * A list of protected branches, as defined below.
     */
    readonly protectedBranches: outputs.GetProjectProtectedBranchesProtectedBranch[];
}
/**
 * The `gitlab.getProjectProtectedBranches` data source allows details of the protected branches of a given project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#list-protected-branches)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getProjectProtectedBranches({
 *     projectId: "foo/bar/baz",
 * });
 * ```
 */
export function getProjectProtectedBranchesOutput(args: GetProjectProtectedBranchesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectProtectedBranchesResult> {
    return pulumi.output(args).apply((a: any) => getProjectProtectedBranches(a, opts))
}

/**
 * A collection of arguments for invoking getProjectProtectedBranches.
 */
export interface GetProjectProtectedBranchesOutputArgs {
    /**
     * The integer or path with namespace that uniquely identifies the project.
     */
    projectId: pulumi.Input<string>;
}
