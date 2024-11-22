// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.Application` resource allows to manage the lifecycle of applications in gitlab.
 *
 * > In order to use a user for a user to create an application, they must have admin privileges at the instance level.
 * To create an OIDC application, a scope of "openid".
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const oidc = new gitlab.Application("oidc", {
 *     confidential: true,
 *     scopes: ["openid"],
 *     name: "company_oidc",
 *     redirectUrl: "https://mycompany.com",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_application`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_application.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * Gitlab applications can be imported with their id, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/application:Application example "1"
 * ```
 *
 * NOTE: the secret and scopes cannot be imported
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Internal name of the application.
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
     */
    public readonly confidential!: pulumi.Output<boolean>;
    /**
     * Name of the application.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL gitlab should send the user to after authentication.
     */
    public readonly redirectUrl!: pulumi.Output<string>;
    /**
     * Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
     * This is only populated when creating a new application. This attribute is not available for imported resources
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["confidential"] = state ? state.confidential : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["redirectUrl"] = state ? state.redirectUrl : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.redirectUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'redirectUrl'");
            }
            if ((!args || args.scopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopes'");
            }
            resourceInputs["confidential"] = args ? args.confidential : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["redirectUrl"] = args ? args.redirectUrl : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * Internal name of the application.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
     */
    confidential?: pulumi.Input<boolean>;
    /**
     * Name of the application.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL gitlab should send the user to after authentication.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
     * This is only populated when creating a new application. This attribute is not available for imported resources
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
     */
    secret?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
     */
    confidential?: pulumi.Input<boolean>;
    /**
     * Name of the application.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL gitlab should send the user to after authentication.
     */
    redirectUrl: pulumi.Input<string>;
    /**
     * Scopes of the application. Use "openid" if you plan to use this as an oidc authentication application. Valid options are: `api`, `readApi`, `readUser`, `readRepository`, `writeRepository`, `readRegistry`, `writeRegistry`, `sudo`, `adminMode`, `openid`, `profile`, `email`.
     * This is only populated when creating a new application. This attribute is not available for imported resources
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
}
