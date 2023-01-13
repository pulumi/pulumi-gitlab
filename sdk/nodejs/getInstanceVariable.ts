// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.InstanceVariable` data source allows to retrieve details about an instance-level CI/CD variable.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = pulumi.output(gitlab.getInstanceVariable({
 *     key: "foo",
 * }));
 * ```
 */
export function getInstanceVariable(args: GetInstanceVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceVariableResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gitlab:index/getInstanceVariable:getInstanceVariable", {
        "key": args.key,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceVariable.
 */
export interface GetInstanceVariableArgs {
    /**
     * The name of the variable.
     */
    key: string;
}

/**
 * A collection of values returned by getInstanceVariable.
 */
export interface GetInstanceVariableResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the variable.
     */
    readonly key: string;
    /**
     * If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     */
    readonly masked: boolean;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    readonly protected: boolean;
    /**
     * The value of the variable.
     */
    readonly value: string;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
     */
    readonly variableType: string;
}

export function getInstanceVariableOutput(args: GetInstanceVariableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceVariableResult> {
    return pulumi.output(args).apply(a => getInstanceVariable(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceVariable.
 */
export interface GetInstanceVariableOutputArgs {
    /**
     * The name of the variable.
     */
    key: pulumi.Input<string>;
}