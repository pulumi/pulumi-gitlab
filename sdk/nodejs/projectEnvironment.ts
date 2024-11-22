// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const _this = new gitlab.Group("this", {
 *     name: "example",
 *     path: "example",
 *     description: "An example group",
 * });
 * const thisProject = new gitlab.Project("this", {
 *     name: "example",
 *     namespaceId: _this.id,
 *     initializeWithReadme: true,
 * });
 * const thisProjectEnvironment = new gitlab.ProjectEnvironment("this", {
 *     project: thisProject.id,
 *     name: "example",
 *     externalUrl: "www.example.com",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_environment`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_environment.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * GitLab project environments can be imported using an id made up of `projectId:environmenId`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectEnvironment:ProjectEnvironment bar 123:321
 * ```
 */
export class ProjectEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing ProjectEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectEnvironmentState, opts?: pulumi.CustomResourceOptions): ProjectEnvironment {
        return new ProjectEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectEnvironment:ProjectEnvironment';

    /**
     * Returns true if the given object is an instance of ProjectEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectEnvironment.__pulumiType;
    }

    /**
     * The cluster agent to associate with this environment.
     */
    public readonly clusterAgentId!: pulumi.Output<number | undefined>;
    /**
     * The ISO8601 date/time that this environment was created at in UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Place to link to for this environment.
     */
    public readonly externalUrl!: pulumi.Output<string | undefined>;
    /**
     * The Flux resource path to associate with this environment.
     */
    public readonly fluxResourcePath!: pulumi.Output<string | undefined>;
    /**
     * The Kubernetes namespace to associate with this environment.
     */
    public readonly kubernetesNamespace!: pulumi.Output<string | undefined>;
    /**
     * The name of the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID or full path of the project to environment is created for.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    public /*out*/ readonly slug!: pulumi.Output<string>;
    /**
     * State the environment is in. Valid values are `available`, `stopped`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Determines whether the environment is attempted to be stopped before the environment is deleted.
     */
    public readonly stopBeforeDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * The ISO8601 date/time that this environment was last updated at in UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ProjectEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectEnvironmentArgs | ProjectEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectEnvironmentState | undefined;
            resourceInputs["clusterAgentId"] = state ? state.clusterAgentId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["externalUrl"] = state ? state.externalUrl : undefined;
            resourceInputs["fluxResourcePath"] = state ? state.fluxResourcePath : undefined;
            resourceInputs["kubernetesNamespace"] = state ? state.kubernetesNamespace : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["stopBeforeDestroy"] = state ? state.stopBeforeDestroy : undefined;
            resourceInputs["tier"] = state ? state.tier : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ProjectEnvironmentArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["clusterAgentId"] = args ? args.clusterAgentId : undefined;
            resourceInputs["externalUrl"] = args ? args.externalUrl : undefined;
            resourceInputs["fluxResourcePath"] = args ? args.fluxResourcePath : undefined;
            resourceInputs["kubernetesNamespace"] = args ? args.kubernetesNamespace : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["stopBeforeDestroy"] = args ? args.stopBeforeDestroy : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectEnvironment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectEnvironment resources.
 */
export interface ProjectEnvironmentState {
    /**
     * The cluster agent to associate with this environment.
     */
    clusterAgentId?: pulumi.Input<number>;
    /**
     * The ISO8601 date/time that this environment was created at in UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Place to link to for this environment.
     */
    externalUrl?: pulumi.Input<string>;
    /**
     * The Flux resource path to associate with this environment.
     */
    fluxResourcePath?: pulumi.Input<string>;
    /**
     * The Kubernetes namespace to associate with this environment.
     */
    kubernetesNamespace?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID or full path of the project to environment is created for.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    slug?: pulumi.Input<string>;
    /**
     * State the environment is in. Valid values are `available`, `stopped`.
     */
    state?: pulumi.Input<string>;
    /**
     * Determines whether the environment is attempted to be stopped before the environment is deleted.
     */
    stopBeforeDestroy?: pulumi.Input<boolean>;
    /**
     * The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
     */
    tier?: pulumi.Input<string>;
    /**
     * The ISO8601 date/time that this environment was last updated at in UTC.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectEnvironment resource.
 */
export interface ProjectEnvironmentArgs {
    /**
     * The cluster agent to associate with this environment.
     */
    clusterAgentId?: pulumi.Input<number>;
    /**
     * Place to link to for this environment.
     */
    externalUrl?: pulumi.Input<string>;
    /**
     * The Flux resource path to associate with this environment.
     */
    fluxResourcePath?: pulumi.Input<string>;
    /**
     * The Kubernetes namespace to associate with this environment.
     */
    kubernetesNamespace?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID or full path of the project to environment is created for.
     */
    project: pulumi.Input<string>;
    /**
     * Determines whether the environment is attempted to be stopped before the environment is deleted.
     */
    stopBeforeDestroy?: pulumi.Input<boolean>;
    /**
     * The tier of the new environment. Valid values are `production`, `staging`, `testing`, `development`, `other`.
     */
    tier?: pulumi.Input<string>;
}
