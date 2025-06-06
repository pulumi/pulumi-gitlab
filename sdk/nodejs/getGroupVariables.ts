// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getGroupVariables` data source allows to retrieve all group-level CI/CD variables.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_level_variables/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const vars = gitlab.getGroupVariables({
 *     group: "my/example/group",
 * });
 * // Using an environment scope
 * const stagingVars = gitlab.getGroupVariables({
 *     group: "my/example/group",
 *     environmentScope: "staging/*",
 * });
 * ```
 */
export function getGroupVariables(args: GetGroupVariablesArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupVariablesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getGroupVariables:getGroupVariables", {
        "environmentScope": args.environmentScope,
        "group": args.group,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupVariables.
 */
export interface GetGroupVariablesArgs {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`).
     */
    environmentScope?: string;
    /**
     * The name or id of the group.
     */
    group: string;
}

/**
 * A collection of values returned by getGroupVariables.
 */
export interface GetGroupVariablesResult {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`).
     */
    readonly environmentScope?: string;
    /**
     * The name or id of the group.
     */
    readonly group: string;
    readonly id: string;
    /**
     * The list of variables returned by the search
     */
    readonly variables: outputs.GetGroupVariablesVariable[];
}
/**
 * The `gitlab.getGroupVariables` data source allows to retrieve all group-level CI/CD variables.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_level_variables/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const vars = gitlab.getGroupVariables({
 *     group: "my/example/group",
 * });
 * // Using an environment scope
 * const stagingVars = gitlab.getGroupVariables({
 *     group: "my/example/group",
 *     environmentScope: "staging/*",
 * });
 * ```
 */
export function getGroupVariablesOutput(args: GetGroupVariablesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupVariablesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getGroupVariables:getGroupVariables", {
        "environmentScope": args.environmentScope,
        "group": args.group,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupVariables.
 */
export interface GetGroupVariablesOutputArgs {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`).
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * The name or id of the group.
     */
    group: pulumi.Input<string>;
}
