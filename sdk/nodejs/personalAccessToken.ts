// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.PersonalAccessToken` resource allows to manage the lifecycle of a personal access token for a specified user.
 *
 * > This resource requires administration privileges.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/personal_access_tokens.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const examplePersonalAccessToken = new gitlab.PersonalAccessToken("examplePersonalAccessToken", {
 *     userId: 25,
 *     expiresAt: "2020-03-14",
 *     scopes: ["api"],
 * });
 * const exampleProjectVariable = new gitlab.ProjectVariable("exampleProjectVariable", {
 *     project: gitlab_project.example.id,
 *     key: "pat",
 *     value: examplePersonalAccessToken.token,
 * });
 * ```
 *
 * ## Import
 *
 * # A GitLab Personal Access Token can be imported using a key composed of `<user-id>:<token-id>`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/personalAccessToken:PersonalAccessToken example "12345:1"
 * ```
 *
 * # NOTEthe `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
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
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The name of the personal access token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * True if the token is revoked.
     */
    public /*out*/ readonly revoked!: pulumi.Output<boolean>;
    /**
     * The scope for the personal access token. It determines the actions which can be performed when authenticating with this
     * token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`,
     * `write_registry`, `sudo`.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * The personal access token. This is only populated when creating a new personal access token. This attribute is not
     * available for imported resources.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The id of the user.
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
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["revoked"] = state ? state.revoked : undefined;
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
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revoked"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
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
     * The scope for the personal access token. It determines the actions which can be performed when authenticating with this
     * token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`,
     * `write_registry`, `sudo`.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The personal access token. This is only populated when creating a new personal access token. This attribute is not
     * available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The id of the user.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PersonalAccessToken resource.
 */
export interface PersonalAccessTokenArgs {
    /**
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The name of the personal access token.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope for the personal access token. It determines the actions which can be performed when authenticating with this
     * token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`,
     * `write_registry`, `sudo`.
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the user.
     */
    userId: pulumi.Input<number>;
}
