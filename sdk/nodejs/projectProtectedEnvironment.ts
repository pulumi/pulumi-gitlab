// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a project.
 *
 * > In order to use a user or group in the `deployAccessLevels` configuration,
 *    you need to make sure that users have access to the project and groups must have this project shared.
 *    You may use the `gitlab.ProjectMembership` and `gitlabProjectSharedGroup` resources to achieve this.
 *    Unfortunately, the GitLab API does not complain about users and groups without access to the project and just ignores those.
 *    In case this happens you will get perpetual state diffs.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_environments.html)
 *
 * ## Import
 *
 * GitLab protected environments can be imported using an id made up of `projectId:environmentName`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment bar 123:production
 * ```
 */
export class ProjectProtectedEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing ProjectProtectedEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectProtectedEnvironmentState, opts?: pulumi.CustomResourceOptions): ProjectProtectedEnvironment {
        return new ProjectProtectedEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment';

    /**
     * Returns true if the given object is an instance of ProjectProtectedEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectProtectedEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectProtectedEnvironment.__pulumiType;
    }

    /**
     * Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
     */
    public readonly approvalRules!: pulumi.Output<outputs.ProjectProtectedEnvironmentApprovalRule[]>;
    /**
     * Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
     */
    public readonly deployAccessLevels!: pulumi.Output<outputs.ProjectProtectedEnvironmentDeployAccessLevel[] | undefined>;
    /**
     * The name of the environment.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * The ID or full path of the project which the protected environment is created against.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ProjectProtectedEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectProtectedEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectProtectedEnvironmentArgs | ProjectProtectedEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectProtectedEnvironmentState | undefined;
            resourceInputs["approvalRules"] = state ? state.approvalRules : undefined;
            resourceInputs["deployAccessLevels"] = state ? state.deployAccessLevels : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectProtectedEnvironmentArgs | undefined;
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["approvalRules"] = args ? args.approvalRules : undefined;
            resourceInputs["deployAccessLevels"] = args ? args.deployAccessLevels : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectProtectedEnvironment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectProtectedEnvironment resources.
 */
export interface ProjectProtectedEnvironmentState {
    /**
     * Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
     */
    approvalRules?: pulumi.Input<pulumi.Input<inputs.ProjectProtectedEnvironmentApprovalRule>[]>;
    /**
     * Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
     */
    deployAccessLevels?: pulumi.Input<pulumi.Input<inputs.ProjectProtectedEnvironmentDeployAccessLevel>[]>;
    /**
     * The name of the environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * The ID or full path of the project which the protected environment is created against.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectProtectedEnvironment resource.
 */
export interface ProjectProtectedEnvironmentArgs {
    /**
     * Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
     */
    approvalRules?: pulumi.Input<pulumi.Input<inputs.ProjectProtectedEnvironmentApprovalRule>[]>;
    /**
     * Array of access levels allowed to deploy, with each described by a hash.  Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
     */
    deployAccessLevels?: pulumi.Input<pulumi.Input<inputs.ProjectProtectedEnvironmentDeployAccessLevel>[]>;
    /**
     * The name of the environment.
     */
    environment: pulumi.Input<string>;
    /**
     * The ID or full path of the project which the protected environment is created against.
     */
    project: pulumi.Input<string>;
}
