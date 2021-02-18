// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_deploy\_token
 *
 * This resource allows you to create and manage deploy token for your GitLab projects and groups.
 *
 * ## Example Usage
 * ### Project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.DeployToken("example", {
 *     expiresAt: "2020-03-14T00:00:00.000Z",
 *     project: "example/deploying",
 *     scopes: [
 *         "read_repository",
 *         "read_registry",
 *     ],
 *     username: "example-username",
 * });
 * ```
 * ### Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.DeployToken("example", {
 *     group: "example/deploying",
 *     scopes: ["read_repository"],
 * });
 * ```
 */
export class DeployToken extends pulumi.CustomResource {
    /**
     * Get an existing DeployToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeployTokenState, opts?: pulumi.CustomResourceOptions): DeployToken {
        return new DeployToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/deployToken:DeployToken';

    /**
     * Returns true if the given object is an instance of DeployToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeployToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeployToken.__pulumiType;
    }

    /**
     * Time the token will expire it, RFC3339 format. Will not expire per default.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The name or id of the group to add the deploy token to.
     * Either `project` or `group` must be set.
     */
    public readonly group!: pulumi.Output<string | undefined>;
    /**
     * A name to describe the deploy token with.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the deploy token to.
     * Either `project` or `group` must be set.
     */
    public readonly project!: pulumi.Output<string | undefined>;
    /**
     * Valid values: `readRepository`, `readRegistry`.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * The secret token. This is only populated when creating a new deploy token.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a DeployToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeployTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeployTokenArgs | DeployTokenState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeployTokenState | undefined;
            inputs["expiresAt"] = state ? state.expiresAt : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["token"] = state ? state.token : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as DeployTokenArgs | undefined;
            if ((!args || args.scopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopes'");
            }
            inputs["expiresAt"] = args ? args.expiresAt : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["token"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeployToken.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployToken resources.
 */
export interface DeployTokenState {
    /**
     * Time the token will expire it, RFC3339 format. Will not expire per default.
     */
    readonly expiresAt?: pulumi.Input<string>;
    /**
     * The name or id of the group to add the deploy token to.
     * Either `project` or `group` must be set.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * A name to describe the deploy token with.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy token to.
     * Either `project` or `group` must be set.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Valid values: `readRepository`, `readRegistry`.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The secret token. This is only populated when creating a new deploy token.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeployToken resource.
 */
export interface DeployTokenArgs {
    /**
     * Time the token will expire it, RFC3339 format. Will not expire per default.
     */
    readonly expiresAt?: pulumi.Input<string>;
    /**
     * The name or id of the group to add the deploy token to.
     * Either `project` or `group` must be set.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * A name to describe the deploy token with.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the deploy token to.
     * Either `project` or `group` must be set.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Valid values: `readRepository`, `readRegistry`.
     */
    readonly scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
     */
    readonly username?: pulumi.Input<string>;
}
