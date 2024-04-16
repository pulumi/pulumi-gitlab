// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getProjectVariables` data source allows to retrieve all project-level CI/CD variables.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const vars = gitlab.getProjectVariables({
 *     project: "my/example/project",
 * });
 * // Using an environment scope
 * const stagingVars = gitlab.getProjectVariables({
 *     project: "my/example/project",
 *     environmentScope: "staging/*",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getProjectVariables(args: GetProjectVariablesArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectVariablesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getProjectVariables:getProjectVariables", {
        "environmentScope": args.environmentScope,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectVariables.
 */
export interface GetProjectVariablesArgs {
    environmentScope?: string;
    project: string;
}

/**
 * A collection of values returned by getProjectVariables.
 */
export interface GetProjectVariablesResult {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`).
     */
    readonly environmentScope?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name or id of the project.
     */
    readonly project: string;
    /**
     * The list of variables returned by the search
     */
    readonly variables: outputs.GetProjectVariablesVariable[];
}
/**
 * The `gitlab.getProjectVariables` data source allows to retrieve all project-level CI/CD variables.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const vars = gitlab.getProjectVariables({
 *     project: "my/example/project",
 * });
 * // Using an environment scope
 * const stagingVars = gitlab.getProjectVariables({
 *     project: "my/example/project",
 *     environmentScope: "staging/*",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getProjectVariablesOutput(args: GetProjectVariablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectVariablesResult> {
    return pulumi.output(args).apply((a: any) => getProjectVariables(a, opts))
}

/**
 * A collection of arguments for invoking getProjectVariables.
 */
export interface GetProjectVariablesOutputArgs {
    environmentScope?: pulumi.Input<string>;
    project: pulumi.Input<string>;
}
