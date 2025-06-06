// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectBranches` data source allows details of the branches of a given project to be retrieved.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/branches/#list-repository-branches)
 */
export function getProjectBranches(args: GetProjectBranchesArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectBranchesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectBranches:getProjectBranches", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectBranches.
 */
export interface GetProjectBranchesArgs {
    /**
     * ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: string;
}

/**
 * A collection of values returned by getProjectBranches.
 */
export interface GetProjectBranchesResult {
    /**
     * The list of branches of the project, as defined below.
     */
    readonly branches: outputs.GetProjectBranchesBranch[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ID or URL-encoded path of the project owned by the authenticated user.
     */
    readonly project: string;
}
/**
 * The `gitlab.getProjectBranches` data source allows details of the branches of a given project to be retrieved.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/branches/#list-repository-branches)
 */
export function getProjectBranchesOutput(args: GetProjectBranchesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetProjectBranchesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getProjectBranches:getProjectBranches", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectBranches.
 */
export interface GetProjectBranchesOutputArgs {
    /**
     * ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
}
