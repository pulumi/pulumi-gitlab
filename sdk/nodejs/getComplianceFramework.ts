// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ComplianceFramework` data source allows details of a compliance framework to be retrieved by its name and the namespace it belongs to.
 *
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#querynamespace)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getComplianceFramework({
 *     namespacePath: "top-level-group",
 *     name: "HIPAA",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getComplianceFramework(args: GetComplianceFrameworkArgs, opts?: pulumi.InvokeOptions): Promise<GetComplianceFrameworkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getComplianceFramework:getComplianceFramework", {
        "name": args.name,
        "namespacePath": args.namespacePath,
    }, opts);
}

/**
 * A collection of arguments for invoking getComplianceFramework.
 */
export interface GetComplianceFrameworkArgs {
    /**
     * Name for the compliance framework.
     */
    name: string;
    /**
     * Full path of the namespace to where the compliance framework is.
     */
    namespacePath: string;
}

/**
 * A collection of values returned by getComplianceFramework.
 */
export interface GetComplianceFrameworkResult {
    /**
     * Color representation of the compliance framework in hex format. e.g. #FCA121.
     */
    readonly color: string;
    /**
     * Is the compliance framework the default framework for the group.
     */
    readonly default: boolean;
    /**
     * Description for the compliance framework.
     */
    readonly description: string;
    /**
     * Globally unique ID of the compliance framework.
     */
    readonly frameworkId: string;
    readonly id: string;
    /**
     * Name for the compliance framework.
     */
    readonly name: string;
    /**
     * Full path of the namespace to where the compliance framework is.
     */
    readonly namespacePath: string;
    /**
     * Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
     */
    readonly pipelineConfigurationFullPath: string;
}
/**
 * The `gitlab.ComplianceFramework` data source allows details of a compliance framework to be retrieved by its name and the namespace it belongs to.
 *
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#querynamespace)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getComplianceFramework({
 *     namespacePath: "top-level-group",
 *     name: "HIPAA",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getComplianceFrameworkOutput(args: GetComplianceFrameworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComplianceFrameworkResult> {
    return pulumi.output(args).apply((a: any) => getComplianceFramework(a, opts))
}

/**
 * A collection of arguments for invoking getComplianceFramework.
 */
export interface GetComplianceFrameworkOutputArgs {
    /**
     * Name for the compliance framework.
     */
    name: pulumi.Input<string>;
    /**
     * Full path of the namespace to where the compliance framework is.
     */
    namespacePath: pulumi.Input<string>;
}
