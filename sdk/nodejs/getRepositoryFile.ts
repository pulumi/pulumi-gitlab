// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.RepositoryFile` data source allows details of a file in a repository to be retrieved.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/repository_files/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getRepositoryFile({
 *     project: "example",
 *     ref: "main",
 *     filePath: "README.md",
 * });
 * ```
 */
export function getRepositoryFile(args: GetRepositoryFileArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryFileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getRepositoryFile:getRepositoryFile", {
        "filePath": args.filePath,
        "project": args.project,
        "ref": args.ref,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryFile.
 */
export interface GetRepositoryFileArgs {
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     */
    filePath: string;
    /**
     * The name or ID of the project.
     */
    project: string;
    /**
     * The name of branch, tag or commit.
     */
    ref: string;
}

/**
 * A collection of values returned by getRepositoryFile.
 */
export interface GetRepositoryFileResult {
    /**
     * The blob id.
     */
    readonly blobId: string;
    /**
     * The commit id.
     */
    readonly commitId: string;
    /**
     * File content.
     */
    readonly content: string;
    /**
     * File content sha256 digest.
     */
    readonly contentSha256: string;
    /**
     * The file content encoding.
     */
    readonly encoding: string;
    /**
     * Enables or disables the execute flag on the file.
     */
    readonly executeFilemode: boolean;
    /**
     * The filename.
     */
    readonly fileName: string;
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     */
    readonly filePath: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The last known commit id.
     */
    readonly lastCommitId: string;
    /**
     * The name or ID of the project.
     */
    readonly project: string;
    /**
     * The name of branch, tag or commit.
     */
    readonly ref: string;
    /**
     * The file size.
     */
    readonly size: number;
}
/**
 * The `gitlab.RepositoryFile` data source allows details of a file in a repository to be retrieved.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/repository_files/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getRepositoryFile({
 *     project: "example",
 *     ref: "main",
 *     filePath: "README.md",
 * });
 * ```
 */
export function getRepositoryFileOutput(args: GetRepositoryFileOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRepositoryFileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getRepositoryFile:getRepositoryFile", {
        "filePath": args.filePath,
        "project": args.project,
        "ref": args.ref,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryFile.
 */
export interface GetRepositoryFileOutputArgs {
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     */
    filePath: pulumi.Input<string>;
    /**
     * The name or ID of the project.
     */
    project: pulumi.Input<string>;
    /**
     * The name of branch, tag or commit.
     */
    ref: pulumi.Input<string>;
}
