// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getInstanceVariables` data source allows to retrieve all instance-level CI/CD variables.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/instance_level_ci_variables/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const vars = gitlab.getInstanceVariables({});
 * ```
 */
export function getInstanceVariables(opts?: pulumi.InvokeOptions): Promise<GetInstanceVariablesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getInstanceVariables:getInstanceVariables", {
    }, opts);
}

/**
 * A collection of values returned by getInstanceVariables.
 */
export interface GetInstanceVariablesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of variables returned by the search
     */
    readonly variables: outputs.GetInstanceVariablesVariable[];
}
/**
 * The `gitlab.getInstanceVariables` data source allows to retrieve all instance-level CI/CD variables.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/instance_level_ci_variables/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const vars = gitlab.getInstanceVariables({});
 * ```
 */
export function getInstanceVariablesOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceVariablesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getInstanceVariables:getInstanceVariables", {
    }, opts);
}
