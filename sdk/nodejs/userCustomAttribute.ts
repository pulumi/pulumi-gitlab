// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to set custom attributes for a user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const attr = new gitlab.UserCustomAttribute("attr", {
 *     key: "location",
 *     user: 42,
 *     value: "Greenland",
 * });
 * ```
 *
 * ## Import
 *
 * # You can import a user custom attribute using an id made up of `{user-id}:{key}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/userCustomAttribute:UserCustomAttribute attr 42:location
 * ```
 */
export class UserCustomAttribute extends pulumi.CustomResource {
    /**
     * Get an existing UserCustomAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserCustomAttributeState, opts?: pulumi.CustomResourceOptions): UserCustomAttribute {
        return new UserCustomAttribute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/userCustomAttribute:UserCustomAttribute';

    /**
     * Returns true if the given object is an instance of UserCustomAttribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserCustomAttribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserCustomAttribute.__pulumiType;
    }

    /**
     * Key for the Custom Attribute.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The id of the user.
     */
    public readonly user!: pulumi.Output<number>;
    /**
     * Value for the Custom Attribute.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a UserCustomAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserCustomAttributeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserCustomAttributeArgs | UserCustomAttributeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserCustomAttributeState | undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as UserCustomAttributeArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserCustomAttribute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserCustomAttribute resources.
 */
export interface UserCustomAttributeState {
    /**
     * Key for the Custom Attribute.
     */
    key?: pulumi.Input<string>;
    /**
     * The id of the user.
     */
    user?: pulumi.Input<number>;
    /**
     * Value for the Custom Attribute.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserCustomAttribute resource.
 */
export interface UserCustomAttributeArgs {
    /**
     * Key for the Custom Attribute.
     */
    key: pulumi.Input<string>;
    /**
     * The id of the user.
     */
    user: pulumi.Input<number>;
    /**
     * Value for the Custom Attribute.
     */
    value: pulumi.Input<string>;
}