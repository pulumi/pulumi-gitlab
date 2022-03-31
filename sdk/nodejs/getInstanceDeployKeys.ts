// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The `gitlab.getInstanceDeployKeys` data source allows to retrieve a list of deploy keys for a GitLab instance.
 *
 * > This data source requires administration privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#list-all-deploy-keys)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // only public deploy keys
 * const example = pulumi.output(gitlab.getInstanceDeployKeys({
 *     public: true,
 * }));
 * ```
 */
export function getInstanceDeployKeys(args?: GetInstanceDeployKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceDeployKeysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", {
        "public": args.public,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceDeployKeys.
 */
export interface GetInstanceDeployKeysArgs {
    public?: boolean;
}

/**
 * A collection of values returned by getInstanceDeployKeys.
 */
export interface GetInstanceDeployKeysResult {
    readonly deployKeys: outputs.GetInstanceDeployKeysDeployKey[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly public?: boolean;
}

export function getInstanceDeployKeysOutput(args?: GetInstanceDeployKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceDeployKeysResult> {
    return pulumi.output(args).apply(a => getInstanceDeployKeys(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceDeployKeys.
 */
export interface GetInstanceDeployKeysOutputArgs {
    public?: pulumi.Input<boolean>;
}