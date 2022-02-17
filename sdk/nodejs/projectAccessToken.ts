// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage Project Access Token for your GitLab projects.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const exampleProjectAccessToken = new gitlab.ProjectAccessToken("exampleProjectAccessToken", {
 *     project: "25",
 *     expiresAt: "2020-03-14",
 *     scopes: ["api"],
 * });
 * const exampleProjectVariable = new gitlab.ProjectVariable("exampleProjectVariable", {
 *     project: gitlab_project.example.id,
 *     key: "pat",
 *     value: exampleProjectAccessToken.token,
 * });
 * ```
 */
export class ProjectAccessToken extends pulumi.CustomResource {
    /**
     * Get an existing ProjectAccessToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectAccessTokenState, opts?: pulumi.CustomResourceOptions): ProjectAccessToken {
        return new ProjectAccessToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectAccessToken:ProjectAccessToken';

    /**
     * Returns true if the given object is an instance of ProjectAccessToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectAccessToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectAccessToken.__pulumiType;
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
     * Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * A name to describe the project access token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the project to add the project access token to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * True if the token is revoked.
     */
    public /*out*/ readonly revoked!: pulumi.Output<boolean>;
    /**
     * Valid values: `api`, `readApi`, `readRepository`, `writeRepository`.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * The secret token. This is only populated when creating a new project access token.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The userId associated to the token.
     */
    public /*out*/ readonly userId!: pulumi.Output<number>;

    /**
     * Create a ProjectAccessToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectAccessTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectAccessTokenArgs | ProjectAccessTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectAccessTokenState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["revoked"] = state ? state.revoked : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as ProjectAccessTokenArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.scopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopes'");
            }
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revoked"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectAccessToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectAccessToken resources.
 */
export interface ProjectAccessTokenState {
    /**
     * True if the token is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Time the token has been created, RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * A name to describe the project access token.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the project to add the project access token to.
     */
    project?: pulumi.Input<string>;
    /**
     * True if the token is revoked.
     */
    revoked?: pulumi.Input<boolean>;
    /**
     * Valid values: `api`, `readApi`, `readRepository`, `writeRepository`.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The secret token. This is only populated when creating a new project access token.
     */
    token?: pulumi.Input<string>;
    /**
     * The userId associated to the token.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectAccessToken resource.
 */
export interface ProjectAccessTokenArgs {
    /**
     * Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * A name to describe the project access token.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the project to add the project access token to.
     */
    project: pulumi.Input<string>;
    /**
     * Valid values: `api`, `readApi`, `readRepository`, `writeRepository`.
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
}