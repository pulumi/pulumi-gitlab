// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     */
    public readonly canCreateGroup!: pulumi.Output<boolean | undefined>;
    /**
     * The e-mail address of the user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Boolean, defaults to false.  Whether to enable administrative priviledges
     * for the user.
     */
    public readonly isAdmin!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     */
    public readonly isExternal!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password of the user.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Integer, defaults to 0.  Number of projects user can create.
     */
    public readonly projectsLimit!: pulumi.Output<number | undefined>;
    /**
     * Boolean, defaults to false. Send user password reset link.
     */
    public readonly resetPassword!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     */
    public readonly skipConfirmation!: pulumi.Output<boolean | undefined>;
    /**
     * The username of the user.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["canCreateGroup"] = state ? state.canCreateGroup : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["isAdmin"] = state ? state.isAdmin : undefined;
            inputs["isExternal"] = state ? state.isExternal : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["projectsLimit"] = state ? state.projectsLimit : undefined;
            inputs["resetPassword"] = state ? state.resetPassword : undefined;
            inputs["skipConfirmation"] = state ? state.skipConfirmation : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            inputs["canCreateGroup"] = args ? args.canCreateGroup : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["isAdmin"] = args ? args.isAdmin : undefined;
            inputs["isExternal"] = args ? args.isExternal : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["projectsLimit"] = args ? args.projectsLimit : undefined;
            inputs["resetPassword"] = args ? args.resetPassword : undefined;
            inputs["skipConfirmation"] = args ? args.skipConfirmation : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     */
    readonly canCreateGroup?: pulumi.Input<boolean>;
    /**
     * The e-mail address of the user.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Boolean, defaults to false.  Whether to enable administrative priviledges
     * for the user.
     */
    readonly isAdmin?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     */
    readonly isExternal?: pulumi.Input<boolean>;
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password of the user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Integer, defaults to 0.  Number of projects user can create.
     */
    readonly projectsLimit?: pulumi.Input<number>;
    /**
     * Boolean, defaults to false. Send user password reset link.
     */
    readonly resetPassword?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     */
    readonly skipConfirmation?: pulumi.Input<boolean>;
    /**
     * The username of the user.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     */
    readonly canCreateGroup?: pulumi.Input<boolean>;
    /**
     * The e-mail address of the user.
     */
    readonly email: pulumi.Input<string>;
    /**
     * Boolean, defaults to false.  Whether to enable administrative priviledges
     * for the user.
     */
    readonly isAdmin?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     */
    readonly isExternal?: pulumi.Input<boolean>;
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password of the user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Integer, defaults to 0.  Number of projects user can create.
     */
    readonly projectsLimit?: pulumi.Input<number>;
    /**
     * Boolean, defaults to false. Send user password reset link.
     */
    readonly resetPassword?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     */
    readonly skipConfirmation?: pulumi.Input<boolean>;
    /**
     * The username of the user.
     */
    readonly username: pulumi.Input<string>;
}
