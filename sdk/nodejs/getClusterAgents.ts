// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.getClusterAgents` data source allows details of GitLab Agents for Kubernetes in a project.
 *
 * > Requires at least GitLab 14.10
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const agents = gitlab.getClusterAgents({
 *     project: "12345",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getClusterAgents(args: GetClusterAgentsArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterAgentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getClusterAgents:getClusterAgents", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterAgents.
 */
export interface GetClusterAgentsArgs {
    project: string;
}

/**
 * A collection of values returned by getClusterAgents.
 */
export interface GetClusterAgentsResult {
    /**
     * List of the registered agents.
     */
    readonly clusterAgents: outputs.GetClusterAgentsClusterAgent[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID or full path of the project owned by the authenticated user.
     */
    readonly project: string;
}
/**
 * The `gitlab.getClusterAgents` data source allows details of GitLab Agents for Kubernetes in a project.
 *
 * > Requires at least GitLab 14.10
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/cluster_agents.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const agents = gitlab.getClusterAgents({
 *     project: "12345",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getClusterAgentsOutput(args: GetClusterAgentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterAgentsResult> {
    return pulumi.output(args).apply((a: any) => getClusterAgents(a, opts))
}

/**
 * A collection of arguments for invoking getClusterAgents.
 */
export interface GetClusterAgentsOutputArgs {
    project: pulumi.Input<string>;
}
