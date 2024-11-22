// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a group.
 *
 * > In order to use a userId in the `deployAccessLevels` configuration,
 *    you need to make sure that users have access to the group with Maintainer role or higher.
 *    In order to use a groupId in the `deployAccessLevels` configuration,
 *    the groupId must be a sub-group under the given group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_protected_environments.html)
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_protected_environment`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_group_protected_environment.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * GitLab group protected environments can be imported using an id made up of `groupId:environmentName`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment bar 123:production
 * ```
 */
export class GroupProtectedEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing GroupProtectedEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupProtectedEnvironmentState, opts?: pulumi.CustomResourceOptions): GroupProtectedEnvironment {
        return new GroupProtectedEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment';

    /**
     * Returns true if the given object is an instance of GroupProtectedEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupProtectedEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupProtectedEnvironment.__pulumiType;
    }

    /**
     * Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
     */
    public readonly approvalRules!: pulumi.Output<outputs.GroupProtectedEnvironmentApprovalRule[]>;
    /**
     * Array of access levels allowed to deploy, with each described by a hash. Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
     */
    public readonly deployAccessLevels!: pulumi.Output<outputs.GroupProtectedEnvironmentDeployAccessLevel[]>;
    /**
     * The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * The ID or full path of the group which the protected environment is created against.
     */
    public readonly group!: pulumi.Output<string>;

    /**
     * Create a GroupProtectedEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupProtectedEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupProtectedEnvironmentArgs | GroupProtectedEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupProtectedEnvironmentState | undefined;
            resourceInputs["approvalRules"] = state ? state.approvalRules : undefined;
            resourceInputs["deployAccessLevels"] = state ? state.deployAccessLevels : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
        } else {
            const args = argsOrState as GroupProtectedEnvironmentArgs | undefined;
            if ((!args || args.deployAccessLevels === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deployAccessLevels'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            resourceInputs["approvalRules"] = args ? args.approvalRules : undefined;
            resourceInputs["deployAccessLevels"] = args ? args.deployAccessLevels : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupProtectedEnvironment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupProtectedEnvironment resources.
 */
export interface GroupProtectedEnvironmentState {
    /**
     * Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
     */
    approvalRules?: pulumi.Input<pulumi.Input<inputs.GroupProtectedEnvironmentApprovalRule>[]>;
    /**
     * Array of access levels allowed to deploy, with each described by a hash. Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
     */
    deployAccessLevels?: pulumi.Input<pulumi.Input<inputs.GroupProtectedEnvironmentDeployAccessLevel>[]>;
    /**
     * The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
     */
    environment?: pulumi.Input<string>;
    /**
     * The ID or full path of the group which the protected environment is created against.
     */
    group?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupProtectedEnvironment resource.
 */
export interface GroupProtectedEnvironmentArgs {
    /**
     * Array of approval rules to deploy, with each described by a hash. Elements in the `approvalRules` should be one of `userId`, `groupId` or `accessLevel`.
     */
    approvalRules?: pulumi.Input<pulumi.Input<inputs.GroupProtectedEnvironmentApprovalRule>[]>;
    /**
     * Array of access levels allowed to deploy, with each described by a hash. Elements in the `deployAccessLevels` should be one of `userId`, `groupId` or `accessLevel`.
     */
    deployAccessLevels: pulumi.Input<pulumi.Input<inputs.GroupProtectedEnvironmentDeployAccessLevel>[]>;
    /**
     * The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
     */
    environment: pulumi.Input<string>;
    /**
     * The ID or full path of the group which the protected environment is created against.
     */
    group: pulumi.Input<string>;
}
