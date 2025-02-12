// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getMetadata` data source retrieves the metadata of the GitLab instance.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/metadata.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const _this = gitlab.getMetadata({});
 * ```
 */
export function getMetadata(opts?: pulumi.InvokeOptions): Promise<GetMetadataResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getMetadata:getMetadata", {
    }, opts);
}

/**
 * A collection of values returned by getMetadata.
 */
export interface GetMetadataResult {
    /**
     * If the GitLab instance is an enterprise instance or not. Supported for GitLab 15.6 onwards.
     */
    readonly enterprise: boolean;
    /**
     * The id of the data source. It will always be `1`
     */
    readonly id: string;
    /**
     * Metadata about the GitLab agent server for Kubernetes (KAS).
     */
    readonly kas: outputs.GetMetadataKas;
    /**
     * Revision of the GitLab instance.
     */
    readonly revision: string;
    /**
     * Version of the GitLab instance.
     */
    readonly version: string;
}
/**
 * The `gitlab.getMetadata` data source retrieves the metadata of the GitLab instance.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/metadata.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const _this = gitlab.getMetadata({});
 * ```
 */
export function getMetadataOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMetadataResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getMetadata:getMetadata", {
    }, opts);
}
