// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ServiceExternalWiki` resource allows to manage the lifecycle of a project integration with External Wiki Service.
 *
 * > This resource is deprecated. use `gitlab.IntegrationExternalWiki`instead!
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#external-wiki)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const awesomeProject = new gitlab.Project("awesome_project", {
 *     name: "awesome_project",
 *     description: "My awesome project.",
 *     visibilityLevel: "public",
 * });
 * const wiki = new gitlab.ServiceExternalWiki("wiki", {
 *     project: awesomeProject.id,
 *     externalWikiUrl: "https://MyAwesomeExternalWikiURL.com",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_external_wiki`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_service_external_wiki.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * You can import a gitlab_service_external_wiki state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/serviceExternalWiki:ServiceExternalWiki wiki 1
 * ```
 */
export class ServiceExternalWiki extends pulumi.CustomResource {
    /**
     * Get an existing ServiceExternalWiki resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceExternalWikiState, opts?: pulumi.CustomResourceOptions): ServiceExternalWiki {
        return new ServiceExternalWiki(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/serviceExternalWiki:ServiceExternalWiki';

    /**
     * Returns true if the given object is an instance of ServiceExternalWiki.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceExternalWiki {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceExternalWiki.__pulumiType;
    }

    /**
     * Whether the integration is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The URL of the external wiki.
     */
    public readonly externalWikiUrl!: pulumi.Output<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    public /*out*/ readonly slug!: pulumi.Output<string>;
    /**
     * Title of the integration.
     */
    public /*out*/ readonly title!: pulumi.Output<string>;
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ServiceExternalWiki resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceExternalWikiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceExternalWikiArgs | ServiceExternalWikiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceExternalWikiState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["externalWikiUrl"] = state ? state.externalWikiUrl : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ServiceExternalWikiArgs | undefined;
            if ((!args || args.externalWikiUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalWikiUrl'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["externalWikiUrl"] = args ? args.externalWikiUrl : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceExternalWiki.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceExternalWiki resources.
 */
export interface ServiceExternalWikiState {
    /**
     * Whether the integration is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The URL of the external wiki.
     */
    externalWikiUrl?: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    slug?: pulumi.Input<string>;
    /**
     * Title of the integration.
     */
    title?: pulumi.Input<string>;
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceExternalWiki resource.
 */
export interface ServiceExternalWikiArgs {
    /**
     * The URL of the external wiki.
     */
    externalWikiUrl: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
}
