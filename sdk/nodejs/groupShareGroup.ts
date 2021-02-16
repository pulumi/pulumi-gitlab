// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_group\_share\_group
 *
 * This resource allows you to share a group with another group
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const test = new gitlab.GroupShareGroup("test", {
 *     groupId: gitlab_group.foo.id,
 *     shareGroupId: gitlab_group.bar.id,
 *     groupAccess: "guest",
 *     expiresAt: "2099-01-01",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab group shares can be imported using an id made up of `mainGroupId:shareGroupId`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/groupShareGroup:GroupShareGroup test 12345:1337
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
     * One of five levels of access to the group.
     */
    public readonly groupAccess!: pulumi.Output<string>;
    /**
     * The id of the main group.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The id of an additional group which will be shared with the main group.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupShareGroupState | undefined;
            inputs["expiresAt"] = state ? state.expiresAt : undefined;
            inputs["groupAccess"] = state ? state.groupAccess : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["shareGroupId"] = state ? state.shareGroupId : undefined;
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
            inputs["expiresAt"] = args ? args.expiresAt : undefined;
            inputs["groupAccess"] = args ? args.groupAccess : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["shareGroupId"] = args ? args.shareGroupId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupShareGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupShareGroup resources.
 */
export interface GroupShareGroupState {
    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     */
    readonly expiresAt?: pulumi.Input<string>;
    /**
     * One of five levels of access to the group.
     */
    readonly groupAccess?: pulumi.Input<string>;
    /**
     * The id of the main group.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The id of an additional group which will be shared with the main group.
     */
    readonly shareGroupId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GroupShareGroup resource.
 */
export interface GroupShareGroupArgs {
    /**
     * Share expiration date. Format: `YYYY-MM-DD`
     */
    readonly expiresAt?: pulumi.Input<string>;
    /**
     * One of five levels of access to the group.
     */
    readonly groupAccess: pulumi.Input<string>;
    /**
     * The id of the main group.
     */
    readonly groupId: pulumi.Input<string>;
    /**
     * The id of an additional group which will be shared with the main group.
     */
    readonly shareGroupId: pulumi.Input<number>;
}
