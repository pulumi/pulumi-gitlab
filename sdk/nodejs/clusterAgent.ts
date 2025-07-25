// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ClusterAgent` resource allows to manage the lifecycle of a GitLab Agent for Kubernetes.
 *
 * > Note that this resource only registers the agent, but doesn't configure it.
 *    The configuration needs to be manually added as described in
 *    [the docs](https://docs.gitlab.com/user/clusters/agent/install/index/#create-an-agent-configuration-file).
 *    However, a `gitlab.RepositoryFile` resource may be used to achieve that.
 *
 * > Requires at least maintainer permissions on the project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/cluster_agents/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * import * as std from "@pulumi/std";
 *
 * const example = new gitlab.ClusterAgent("example", {
 *     project: "12345",
 *     name: "agent-1",
 * });
 * // Optionally, configure the agent as described in
 * // https://docs.gitlab.com/user/clusters/agent/install/index/#create-an-agent-configuration-file
 * const exampleAgentConfig = new gitlab.RepositoryFile("example_agent_config", {
 *     project: example.project,
 *     branch: "main",
 *     filePath: pulumi.interpolate`.gitlab/agents/${example.name}/config.yaml`,
 *     encoding: "base64",
 *     content: std.base64encode({
 *         input: "# the GitLab Agent for Kubernetes configuration goes here ...\n",
 *     }).then(invoke => invoke.result),
 *     authorEmail: "terraform@example.com",
 *     authorName: "Terraform",
 *     commitMessage: pulumi.interpolate`feature: add agent config for ${example.name} [skip ci]`,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_cluster_agent`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_cluster_agent.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * GitLab Agent for Kubernetes can be imported with the following command and the id pattern `<project>:<agent-id>`
 *
 * ```sh
 * $ pulumi import gitlab:index/clusterAgent:ClusterAgent example '12345:42'
 * ```
 */
export class ClusterAgent extends pulumi.CustomResource {
    /**
     * Get an existing ClusterAgent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterAgentState, opts?: pulumi.CustomResourceOptions): ClusterAgent {
        return new ClusterAgent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/clusterAgent:ClusterAgent';

    /**
     * Returns true if the given object is an instance of ClusterAgent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterAgent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterAgent.__pulumiType;
    }

    /**
     * The ID of the agent.
     */
    public /*out*/ readonly agentId!: pulumi.Output<number>;
    /**
     * The ISO8601 datetime when the agent was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The ID of the user who created the agent.
     */
    public /*out*/ readonly createdByUserId!: pulumi.Output<number>;
    /**
     * The Name of the agent.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID or full path of the project maintained by the authenticated user.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ClusterAgent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterAgentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterAgentArgs | ClusterAgentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterAgentState | undefined;
            resourceInputs["agentId"] = state ? state.agentId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["createdByUserId"] = state ? state.createdByUserId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ClusterAgentArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["agentId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdByUserId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterAgent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterAgent resources.
 */
export interface ClusterAgentState {
    /**
     * The ID of the agent.
     */
    agentId?: pulumi.Input<number>;
    /**
     * The ISO8601 datetime when the agent was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The ID of the user who created the agent.
     */
    createdByUserId?: pulumi.Input<number>;
    /**
     * The Name of the agent.
     */
    name?: pulumi.Input<string>;
    /**
     * ID or full path of the project maintained by the authenticated user.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterAgent resource.
 */
export interface ClusterAgentArgs {
    /**
     * The Name of the agent.
     */
    name?: pulumi.Input<string>;
    /**
     * ID or full path of the project maintained by the authenticated user.
     */
    project: pulumi.Input<string>;
}
