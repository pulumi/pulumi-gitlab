// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.User` resource allows to manage the lifecycle of a user.
 *
 * > the provider needs to be configured with admin-level access for this resource to work.
 *
 * > You must specify either password or reset_password.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/users/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.User("example", {
 *     name: "Example Foo",
 *     username: "example",
 *     password: "superPassword",
 *     email: "gitlab@user.create",
 *     isAdmin: true,
 *     projectsLimit: 4,
 *     canCreateGroup: false,
 *     isExternal: true,
 *     resetPassword: false,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_user`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_user.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * ```sh
 * $ pulumi import gitlab:index/user:User You can import a user to terraform state using `<resource> <id>`.
 * ```
 *
 * The `id` must be an integer for the id of the user you want to import,
 *
 * for example:
 *
 * ```sh
 * $ pulumi import gitlab:index/user:User example 42
 * ```
 */
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
     * String, a specific external authentication provider UID.
     *
     * @deprecated To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     */
    public readonly externUid!: pulumi.Output<string | undefined>;
    /**
     * String, the external provider.
     *
     * @deprecated To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     */
    public readonly externalProvider!: pulumi.Output<string | undefined>;
    /**
     * Set user password to a random value
     */
    public readonly forceRandomPassword!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, defaults to false.  Whether to enable administrative privileges
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
     * The ID of the user's namespace.
     */
    public readonly namespaceId!: pulumi.Output<number>;
    /**
     * The note associated to the user.
     */
    public readonly note!: pulumi.Output<string | undefined>;
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
     * String, defaults to 'active'. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["canCreateGroup"] = state ? state.canCreateGroup : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["externUid"] = state ? state.externUid : undefined;
            resourceInputs["externalProvider"] = state ? state.externalProvider : undefined;
            resourceInputs["forceRandomPassword"] = state ? state.forceRandomPassword : undefined;
            resourceInputs["isAdmin"] = state ? state.isAdmin : undefined;
            resourceInputs["isExternal"] = state ? state.isExternal : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectsLimit"] = state ? state.projectsLimit : undefined;
            resourceInputs["resetPassword"] = state ? state.resetPassword : undefined;
            resourceInputs["skipConfirmation"] = state ? state.skipConfirmation : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["canCreateGroup"] = args ? args.canCreateGroup : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["externUid"] = args ? args.externUid : undefined;
            resourceInputs["externalProvider"] = args ? args.externalProvider : undefined;
            resourceInputs["forceRandomPassword"] = args ? args.forceRandomPassword : undefined;
            resourceInputs["isAdmin"] = args ? args.isAdmin : undefined;
            resourceInputs["isExternal"] = args ? args.isExternal : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectsLimit"] = args ? args.projectsLimit : undefined;
            resourceInputs["resetPassword"] = args ? args.resetPassword : undefined;
            resourceInputs["skipConfirmation"] = args ? args.skipConfirmation : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     */
    canCreateGroup?: pulumi.Input<boolean>;
    /**
     * The e-mail address of the user.
     */
    email?: pulumi.Input<string>;
    /**
     * String, a specific external authentication provider UID.
     *
     * @deprecated To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     */
    externUid?: pulumi.Input<string>;
    /**
     * String, the external provider.
     *
     * @deprecated To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     */
    externalProvider?: pulumi.Input<string>;
    /**
     * Set user password to a random value
     */
    forceRandomPassword?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to false.  Whether to enable administrative privileges
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     */
    isExternal?: pulumi.Input<boolean>;
    /**
     * The name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the user's namespace.
     */
    namespaceId?: pulumi.Input<number>;
    /**
     * The note associated to the user.
     */
    note?: pulumi.Input<string>;
    /**
     * The password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * Integer, defaults to 0.  Number of projects user can create.
     */
    projectsLimit?: pulumi.Input<number>;
    /**
     * Boolean, defaults to false. Send user password reset link.
     */
    resetPassword?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     */
    skipConfirmation?: pulumi.Input<boolean>;
    /**
     * String, defaults to 'active'. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     */
    state?: pulumi.Input<string>;
    /**
     * The username of the user.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     */
    canCreateGroup?: pulumi.Input<boolean>;
    /**
     * The e-mail address of the user.
     */
    email: pulumi.Input<string>;
    /**
     * String, a specific external authentication provider UID.
     *
     * @deprecated To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     */
    externUid?: pulumi.Input<string>;
    /**
     * String, the external provider.
     *
     * @deprecated To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     */
    externalProvider?: pulumi.Input<string>;
    /**
     * Set user password to a random value
     */
    forceRandomPassword?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to false.  Whether to enable administrative privileges
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     */
    isExternal?: pulumi.Input<boolean>;
    /**
     * The name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the user's namespace.
     */
    namespaceId?: pulumi.Input<number>;
    /**
     * The note associated to the user.
     */
    note?: pulumi.Input<string>;
    /**
     * The password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * Integer, defaults to 0.  Number of projects user can create.
     */
    projectsLimit?: pulumi.Input<number>;
    /**
     * Boolean, defaults to false. Send user password reset link.
     */
    resetPassword?: pulumi.Input<boolean>;
    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     */
    skipConfirmation?: pulumi.Input<boolean>;
    /**
     * String, defaults to 'active'. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     */
    state?: pulumi.Input<string>;
    /**
     * The username of the user.
     */
    username: pulumi.Input<string>;
}
