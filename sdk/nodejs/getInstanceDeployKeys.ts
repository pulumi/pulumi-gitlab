// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getInstanceDeployKeys` data source allows to retrieve a list of deploy keys for a GitLab instance.
 *
 * > This data source requires administration privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_keys/#list-all-deploy-keys)
 */
export function getInstanceDeployKeys(args?: GetInstanceDeployKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceDeployKeysResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", {
        "public": args.public,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceDeployKeys.
 */
export interface GetInstanceDeployKeysArgs {
    /**
     * Only return deploy keys that are public.
     */
    public?: boolean;
}

/**
 * A collection of values returned by getInstanceDeployKeys.
 */
export interface GetInstanceDeployKeysResult {
    /**
     * The list of all deploy keys across all projects of the GitLab instance.
     */
    readonly deployKeys: outputs.GetInstanceDeployKeysDeployKey[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Only return deploy keys that are public.
     */
    readonly public?: boolean;
}
/**
 * The `gitlab.getInstanceDeployKeys` data source allows to retrieve a list of deploy keys for a GitLab instance.
 *
 * > This data source requires administration privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/deploy_keys/#list-all-deploy-keys)
 */
export function getInstanceDeployKeysOutput(args?: GetInstanceDeployKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceDeployKeysResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", {
        "public": args.public,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceDeployKeys.
 */
export interface GetInstanceDeployKeysOutputArgs {
    /**
     * Only return deploy keys that are public.
     */
    public?: pulumi.Input<boolean>;
}
