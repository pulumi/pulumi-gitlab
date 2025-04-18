// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.PersonalAccessToken` resource allows to manage the lifecycle of a personal access token.
 *
 * > This resource requires administration privileges.
 *
 * > Use of the `timestamp()` function with expiresAt will cause the resource to be re-created with every apply, it's recommended to use `plantimestamp()` or a static value instead.
 *
 * > Observability scopes are in beta and may not work on all instances. See more details in [the documentation](https://docs.gitlab.com/operations/tracing/)
 *
 * > Use `rotationConfiguration` to automatically rotate tokens instead of using `timestamp()` as timestamp will cause changes with every plan. `pulumi up` must still be run to rotate the token.
 *
 * > Due to [Automatic reuse detection](https://docs.gitlab.com/api/personal_access_tokens/#automatic-reuse-detection) it's possible that a new Personal Access Token will immediately be revoked. Check if an old process using the old token is running if this happens.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/personal_access_tokens/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.PersonalAccessToken("example", {
 *     userId: 25,
 *     name: "Example personal access token",
 *     expiresAt: "2020-03-14",
 *     scopes: ["api"],
 * });
 * const exampleProjectVariable = new gitlab.ProjectVariable("example", {
 *     project: exampleGitlabProject.id,
 *     key: "pat",
 *     value: example.token,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_personal_access_token`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_personal_access_token.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * A GitLab Personal Access Token can be imported using a key composed of `<user-id>:<token-id>`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/personalAccessToken:PersonalAccessToken example "12345:1"
 * ```
 *
 * NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 */
export class PersonalAccessToken extends pulumi.CustomResource {
    /**
     * Get an existing PersonalAccessToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PersonalAccessTokenState, opts?: pulumi.CustomResourceOptions): PersonalAccessToken {
        return new PersonalAccessToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/personalAccessToken:PersonalAccessToken';

    /**
     * Returns true if the given object is an instance of PersonalAccessToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PersonalAccessToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PersonalAccessToken.__pulumiType;
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
     * The description of the personal access token.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * When the token will expire, YYYY-MM-DD format. Is automatically set when `rotationConfiguration` is used.
     */
    public readonly expiresAt!: pulumi.Output<string>;
    /**
     * The name of the personal access token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * True if the token is revoked.
     */
    public /*out*/ readonly revoked!: pulumi.Output<boolean>;
    /**
     * The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     */
    public readonly rotationConfiguration!: pulumi.Output<outputs.PersonalAccessTokenRotationConfiguration | undefined>;
    /**
     * The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * The token of the personal access token. **Note**: the token is not available for imported resources.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The ID of the user.
     */
    public readonly userId!: pulumi.Output<number>;

    /**
     * Create a PersonalAccessToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PersonalAccessTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PersonalAccessTokenArgs | PersonalAccessTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PersonalAccessTokenState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["revoked"] = state ? state.revoked : undefined;
            resourceInputs["rotationConfiguration"] = state ? state.rotationConfiguration : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as PersonalAccessTokenArgs | undefined;
            if ((!args || args.scopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopes'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rotationConfiguration"] = args ? args.rotationConfiguration : undefined;
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
        super(PersonalAccessToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PersonalAccessToken resources.
 */
export interface PersonalAccessTokenState {
    /**
     * True if the token is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Time the token has been created, RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the personal access token.
     */
    description?: pulumi.Input<string>;
    /**
     * When the token will expire, YYYY-MM-DD format. Is automatically set when `rotationConfiguration` is used.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The name of the personal access token.
     */
    name?: pulumi.Input<string>;
    /**
     * True if the token is revoked.
     */
    revoked?: pulumi.Input<boolean>;
    /**
     * The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     */
    rotationConfiguration?: pulumi.Input<inputs.PersonalAccessTokenRotationConfiguration>;
    /**
     * The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The token of the personal access token. **Note**: the token is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The ID of the user.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PersonalAccessToken resource.
 */
export interface PersonalAccessTokenArgs {
    /**
     * The description of the personal access token.
     */
    description?: pulumi.Input<string>;
    /**
     * When the token will expire, YYYY-MM-DD format. Is automatically set when `rotationConfiguration` is used.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The name of the personal access token.
     */
    name?: pulumi.Input<string>;
    /**
     * The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     */
    rotationConfiguration?: pulumi.Input<inputs.PersonalAccessTokenRotationConfiguration>;
    /**
     * The scopes of the personal access token. valid values are: `api`, `readUser`, `readApi`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `createRunner`, `manageRunner`, `aiFeatures`, `k8sProxy`, `readServicePing`
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the user.
     */
    userId: pulumi.Input<number>;
}
