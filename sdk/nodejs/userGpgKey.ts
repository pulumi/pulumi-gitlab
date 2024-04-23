// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.UserGpgKey` resource allows to manage the lifecycle of a GPG key assigned to the current user or a specific user.
 *
 * > Managing GPG keys for arbitrary users requires admin privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#get-a-specific-gpg-key)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = gitlab.getUser({
 *     username: "example-user",
 * });
 * // Manages a GPG key for the specified user. An admin token is required if `user_id` is specified.
 * const exampleUserGpgKey = new gitlab.UserGpgKey("example", {
 *     userId: example.then(example => example.id),
 *     key: `-----BEGIN PGP PUBLIC KEY BLOCK-----
 * ...
 * -----END PGP PUBLIC KEY BLOCK-----`,
 * });
 * // Manages a GPG key for the current user
 * const exampleUser = new gitlab.UserGpgKey("example_user", {key: `-----BEGIN PGP PUBLIC KEY BLOCK-----
 * ...
 * -----END PGP PUBLIC KEY BLOCK-----`});
 * ```
 *
 * ## Import
 *
 * You can import a GPG key for a specific user using an id made up of `{user-id}:{key}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/userGpgKey:UserGpgKey example 42:1
 * ```
 *
 * Alternatively, you can import a GPG key for the current user using an id made up of `{key}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/userGpgKey:UserGpgKey example_user 1
 * ```
 */
export class UserGpgKey extends pulumi.CustomResource {
    /**
     * Get an existing UserGpgKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGpgKeyState, opts?: pulumi.CustomResourceOptions): UserGpgKey {
        return new UserGpgKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/userGpgKey:UserGpgKey';

    /**
     * Returns true if the given object is an instance of UserGpgKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGpgKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGpgKey.__pulumiType;
    }

    /**
     * The time when this key was created in GitLab.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The armored GPG public key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the GPG key.
     */
    public /*out*/ readonly keyId!: pulumi.Output<number>;
    /**
     * The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the specified user, and an admin token is required.
     */
    public readonly userId!: pulumi.Output<number | undefined>;

    /**
     * Create a UserGpgKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserGpgKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGpgKeyArgs | UserGpgKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGpgKeyState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserGpgKeyArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserGpgKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGpgKey resources.
 */
export interface UserGpgKeyState {
    /**
     * The time when this key was created in GitLab.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The armored GPG public key.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the GPG key.
     */
    keyId?: pulumi.Input<number>;
    /**
     * The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the specified user, and an admin token is required.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a UserGpgKey resource.
 */
export interface UserGpgKeyArgs {
    /**
     * The armored GPG public key.
     */
    key: pulumi.Input<string>;
    /**
     * The ID of the user to add the GPG key to. If this field is omitted, this resource manages a GPG key for the current user. Otherwise, this resource manages a GPG key for the specified user, and an admin token is required.
     */
    userId?: pulumi.Input<number>;
}
