// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupMembership` resource allows to manage the lifecycle of a users group membership.
 *
 * > If a group should grant membership to another group use the `gitlab.GroupShareGroup` resource instead.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/members/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const test = new gitlab.GroupMembership("test", {
 *     groupId: 12345,
 *     userId: 1337,
 *     accessLevel: "guest",
 *     expiresAt: "2020-12-31",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_membership`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_group_membership.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * GitLab group membership can be imported using an id made up of `group_id:user_id`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupMembership:GroupMembership test "12345:1337"
 * ```
 */
export class GroupMembership extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembershipState, opts?: pulumi.CustomResourceOptions): GroupMembership {
        return new GroupMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupMembership:GroupMembership';

    /**
     * Returns true if the given object is an instance of GroupMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembership.__pulumiType;
    }

    /**
     * Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
     */
    public readonly accessLevel!: pulumi.Output<string>;
    /**
     * Expiration date for the group membership. Format: `YYYY-MM-DD`
     */
    public readonly expiresAt!: pulumi.Output<string>;
    /**
     * The ID of the group.
     */
    public readonly groupId!: pulumi.Output<number>;
    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     */
    public readonly memberRoleId!: pulumi.Output<number>;
    /**
     * Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     */
    public readonly skipSubresourcesOnDestroy!: pulumi.Output<boolean>;
    /**
     * Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     */
    public readonly unassignIssuablesOnDestroy!: pulumi.Output<boolean>;
    /**
     * The ID of the user.
     */
    public readonly userId!: pulumi.Output<number>;

    /**
     * Create a GroupMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMembershipArgs | GroupMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMembershipState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["memberRoleId"] = state ? state.memberRoleId : undefined;
            resourceInputs["skipSubresourcesOnDestroy"] = state ? state.skipSubresourcesOnDestroy : undefined;
            resourceInputs["unassignIssuablesOnDestroy"] = state ? state.unassignIssuablesOnDestroy : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as GroupMembershipArgs | undefined;
            if ((!args || args.accessLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessLevel'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["memberRoleId"] = args ? args.memberRoleId : undefined;
            resourceInputs["skipSubresourcesOnDestroy"] = args ? args.skipSubresourcesOnDestroy : undefined;
            resourceInputs["unassignIssuablesOnDestroy"] = args ? args.unassignIssuablesOnDestroy : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembership resources.
 */
export interface GroupMembershipState {
    /**
     * Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Expiration date for the group membership. Format: `YYYY-MM-DD`
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID of the group.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     */
    memberRoleId?: pulumi.Input<number>;
    /**
     * Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     */
    skipSubresourcesOnDestroy?: pulumi.Input<boolean>;
    /**
     * Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     */
    unassignIssuablesOnDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the user.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GroupMembership resource.
 */
export interface GroupMembershipArgs {
    /**
     * Access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
     */
    accessLevel: pulumi.Input<string>;
    /**
     * Expiration date for the group membership. Format: `YYYY-MM-DD`
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID of the group.
     */
    groupId: pulumi.Input<number>;
    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     */
    memberRoleId?: pulumi.Input<number>;
    /**
     * Whether the deletion of direct memberships of the removed member in subgroups and projects should be skipped. Only used during a destroy.
     */
    skipSubresourcesOnDestroy?: pulumi.Input<boolean>;
    /**
     * Whether the removed member should be unassigned from any issues or merge requests inside a given group or project. Only used during a destroy.
     */
    unassignIssuablesOnDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the user.
     */
    userId: pulumi.Input<number>;
}
