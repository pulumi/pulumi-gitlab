// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectVariable` data source allows to retrieve details about a project-level CI/CD variable.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = gitlab.getProjectVariable({
 *     key: "foo",
 *     project: "my/example/project",
 * });
 * const bar = gitlab.getProjectVariable({
 *     environmentScope: "staging/*",
 *     key: "bar",
 *     project: "my/example/project",
 * });
 * ```
 */
export function getProjectVariable(args: GetProjectVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectVariableResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectVariable:getProjectVariable", {
        "environmentScope": args.environmentScope,
        "key": args.key,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectVariable.
 */
export interface GetProjectVariableArgs {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    environmentScope?: string;
    /**
     * The name of the variable.
     */
    key: string;
    /**
     * The name or id of the project.
     */
    project: string;
}

/**
 * A collection of values returned by getProjectVariable.
 */
export interface GetProjectVariableResult {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    readonly environmentScope: string;
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
     * The name or id of the project.
     */
    readonly project: string;
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
/**
 * The `gitlab.ProjectVariable` data source allows to retrieve details about a project-level CI/CD variable.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = gitlab.getProjectVariable({
 *     key: "foo",
 *     project: "my/example/project",
 * });
 * const bar = gitlab.getProjectVariable({
 *     environmentScope: "staging/*",
 *     key: "bar",
 *     project: "my/example/project",
 * });
 * ```
 */
export function getProjectVariableOutput(args: GetProjectVariableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectVariableResult> {
    return pulumi.output(args).apply((a: any) => getProjectVariable(a, opts))
}

/**
 * A collection of arguments for invoking getProjectVariable.
 */
export interface GetProjectVariableOutputArgs {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * The name of the variable.
     */
    key: pulumi.Input<string>;
    /**
     * The name or id of the project.
     */
    project: pulumi.Input<string>;
}
