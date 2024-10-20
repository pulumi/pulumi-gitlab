// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupServiceAccountAccessToken` resource allows to manage the lifecycle of a group service account access token.
 *
 * > Use of the `timestamp()` function with expiresAt will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/group_service_accounts.html#create-a-personal-access-token-for-a-service-account-user)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.Group("example", {
 *     name: "example",
 *     path: "example",
 *     description: "An example group",
 * });
 * const example_sa = new gitlab.GroupServiceAccount("example-sa", {
 *     group: example.id,
 *     name: "example-name",
 *     username: "example-username",
 * });
 * const example_sa_token = new gitlab.GroupServiceAccountAccessToken("example-sa-token", {
 *     group: example.id,
 *     userId: example_sa.id,
 *     name: "Example personal access token",
 *     expiresAt: "2020-03-14",
 *     scopes: ["api"],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken You can import a service account access token using `<resource> <id>`. The
 * ```
 *
 * `id` is in the form of <group_id>:<service_account_id>:<access_token_id>
 *
 * Importing an access token does not import the access token value.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken example 1:2:3
 * ```
 */
export class GroupServiceAccountAccessToken extends pulumi.CustomResource {
    /**
     * Get an existing GroupServiceAccountAccessToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupServiceAccountAccessTokenState, opts?: pulumi.CustomResourceOptions): GroupServiceAccountAccessToken {
        return new GroupServiceAccountAccessToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken';

    /**
     * Returns true if the given object is an instance of GroupServiceAccountAccessToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupServiceAccountAccessToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupServiceAccountAccessToken.__pulumiType;
    }

    /**
     * True if the token is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * Time the token has been created, RFC3339 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     */
    public readonly expiresAt!: pulumi.Output<string>;
    /**
     * The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The name of the personal access token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * True if the token is revoked.
     */
    public /*out*/ readonly revoked!: pulumi.Output<boolean>;
    /**
     * The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * The token of the group service account access token. **Note**: the token is not available for imported resources.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The ID of a service account user.
     */
    public readonly userId!: pulumi.Output<number>;

    /**
     * Create a GroupServiceAccountAccessToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupServiceAccountAccessTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupServiceAccountAccessTokenArgs | GroupServiceAccountAccessTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupServiceAccountAccessTokenState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["revoked"] = state ? state.revoked : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as GroupServiceAccountAccessTokenArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.scopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopes'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revoked"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(GroupServiceAccountAccessToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupServiceAccountAccessToken resources.
 */
export interface GroupServiceAccountAccessTokenState {
    /**
     * True if the token is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Time the token has been created, RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     */
    group?: pulumi.Input<string>;
    /**
     * The name of the personal access token.
     */
    name?: pulumi.Input<string>;
    /**
     * True if the token is revoked.
     */
    revoked?: pulumi.Input<boolean>;
    /**
     * The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The token of the group service account access token. **Note**: the token is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The ID of a service account user.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GroupServiceAccountAccessToken resource.
 */
export interface GroupServiceAccountAccessTokenArgs {
    /**
     * The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     */
    group: pulumi.Input<string>;
    /**
     * The name of the personal access token.
     */
    name?: pulumi.Input<string>;
    /**
     * The scopes of the group service account access token. valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readObservability`, `writeObservability`
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of a service account user.
     */
    userId: pulumi.Input<number>;
}
