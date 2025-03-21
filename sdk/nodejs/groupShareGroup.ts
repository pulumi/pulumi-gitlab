// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupShareGroup` resource allows managing the lifecycle of a group shared with another group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/groups/#share-groups-with-groups)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const test = new gitlab.GroupShareGroup("test", {
 *     groupId: foo.id,
 *     shareGroupId: bar.id,
 *     groupAccess: "guest",
 *     expiresAt: "2099-01-01",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_share_group`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_group_share_group.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
 * ```
 */
export class GroupShareGroup extends pulumi.CustomResource {
    /**
     * Get an existing GroupShareGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupShareGroupState, opts?: pulumi.CustomResourceOptions): GroupShareGroup {
        return new GroupShareGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupShareGroup:GroupShareGroup';

    /**
     * Returns true if the given object is an instance of GroupShareGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupShareGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupShareGroup.__pulumiType;
    }

    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     */
    public readonly groupAccess!: pulumi.Output<string>;
    /**
     * The id of the main group to be shared.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     */
    public readonly memberRoleId!: pulumi.Output<number>;
    /**
     * The id of the additional group with which the main group will be shared.
     */
    public readonly shareGroupId!: pulumi.Output<number>;

    /**
     * Create a GroupShareGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupShareGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupShareGroupArgs | GroupShareGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupShareGroupState | undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["groupAccess"] = state ? state.groupAccess : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["memberRoleId"] = state ? state.memberRoleId : undefined;
            resourceInputs["shareGroupId"] = state ? state.shareGroupId : undefined;
        } else {
            const args = argsOrState as GroupShareGroupArgs | undefined;
            if ((!args || args.groupAccess === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupAccess'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.shareGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shareGroupId'");
            }
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["groupAccess"] = args ? args.groupAccess : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["memberRoleId"] = args ? args.memberRoleId : undefined;
            resourceInputs["shareGroupId"] = args ? args.shareGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupShareGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupShareGroup resources.
 */
export interface GroupShareGroupState {
    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The id of the main group to be shared.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     */
    memberRoleId?: pulumi.Input<number>;
    /**
     * The id of the additional group with which the main group will be shared.
     */
    shareGroupId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GroupShareGroup resource.
 */
export interface GroupShareGroupArgs {
    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The access level to grant the group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     */
    groupAccess: pulumi.Input<string>;
    /**
     * The id of the main group to be shared.
     */
    groupId: pulumi.Input<string>;
    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     */
    memberRoleId?: pulumi.Input<number>;
    /**
     * The id of the additional group with which the main group will be shared.
     */
    shareGroupId: pulumi.Input<number>;
}
