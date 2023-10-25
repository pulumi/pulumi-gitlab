// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getRepositoryTree` data source allows details of directories and files in a repository to be retrieved.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repositories.html#list-repository-tree)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const this = gitlab.getRepositoryTree({
 *     path: "ExampleSubFolder",
 *     project: "example",
 *     recursive: true,
 *     ref: "main",
 * });
 * ```
 */
export function getRepositoryTree(args: GetRepositoryTreeArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryTreeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getRepositoryTree:getRepositoryTree", {
        "path": args.path,
        "project": args.project,
        "recursive": args.recursive,
        "ref": args.ref,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryTree.
 */
export interface GetRepositoryTreeArgs {
    path?: string;
    /**
     * The ID or full path of the project owned by the authenticated user.
     */
    project: string;
    /**
     * Boolean value used to get a recursive tree (false by default).
     */
    recursive?: boolean;
    /**
     * The name of a repository branch or tag.
     */
    ref: string;
}

/**
 * A collection of values returned by getRepositoryTree.
 */
export interface GetRepositoryTreeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The path inside repository. Used to get content of subdirectories.
     */
    readonly path?: string;
    /**
     * The ID or full path of the project owned by the authenticated user.
     */
    readonly project: string;
    /**
     * Boolean value used to get a recursive tree (false by default).
     */
    readonly recursive?: boolean;
    /**
     * The name of a repository branch or tag.
     */
    readonly ref: string;
    /**
     * The list of files/directories returned by the search
     */
    readonly trees: outputs.GetRepositoryTreeTree[];
}
/**
 * The `gitlab.getRepositoryTree` data source allows details of directories and files in a repository to be retrieved.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/repositories.html#list-repository-tree)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const this = gitlab.getRepositoryTree({
 *     path: "ExampleSubFolder",
 *     project: "example",
 *     recursive: true,
 *     ref: "main",
 * });
 * ```
 */
export function getRepositoryTreeOutput(args: GetRepositoryTreeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryTreeResult> {
    return pulumi.output(args).apply((a: any) => getRepositoryTree(a, opts))
}

/**
 * A collection of arguments for invoking getRepositoryTree.
 */
export interface GetRepositoryTreeOutputArgs {
    path?: pulumi.Input<string>;
    /**
     * The ID or full path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
    /**
     * Boolean value used to get a recursive tree (false by default).
     */
    recursive?: pulumi.Input<boolean>;
    /**
     * The name of a repository branch or tag.
     */
    ref: pulumi.Input<string>;
}
