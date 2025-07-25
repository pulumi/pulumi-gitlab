// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_pages_settings`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_pages_settings.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * Gitlab project pages settings can be imported using the project ID, for example:
 *
 * ```sh
 * $ pulumi import gitlab:index/projectPagesSettings:ProjectPagesSettings example 12345
 * ```
 */
export class ProjectPagesSettings extends pulumi.CustomResource {
    /**
     * Get an existing ProjectPagesSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectPagesSettingsState, opts?: pulumi.CustomResourceOptions): ProjectPagesSettings {
        return new ProjectPagesSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectPagesSettings:ProjectPagesSettings';

    /**
     * Returns true if the given object is an instance of ProjectPagesSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectPagesSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectPagesSettings.__pulumiType;
    }

    /**
     * List of current active deployments.
     */
    public /*out*/ readonly deployments!: pulumi.Output<outputs.ProjectPagesSettingsDeployment[]>;
    /**
     * Boolean indicating if the project is set to force https. Requires `externalHttps` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
     */
    public readonly forceHttps!: pulumi.Output<boolean>;
    /**
     * Boolean indicating if a unique domain is enabled.
     */
    public readonly isUniqueDomainEnabled!: pulumi.Output<boolean>;
    public readonly keepSettingsOnDestroy!: pulumi.Output<boolean>;
    /**
     * The project ID or path.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URL to access the project pages.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a ProjectPagesSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectPagesSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectPagesSettingsArgs | ProjectPagesSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectPagesSettingsState | undefined;
            resourceInputs["deployments"] = state ? state.deployments : undefined;
            resourceInputs["forceHttps"] = state ? state.forceHttps : undefined;
            resourceInputs["isUniqueDomainEnabled"] = state ? state.isUniqueDomainEnabled : undefined;
            resourceInputs["keepSettingsOnDestroy"] = state ? state.keepSettingsOnDestroy : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ProjectPagesSettingsArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["forceHttps"] = args ? args.forceHttps : undefined;
            resourceInputs["isUniqueDomainEnabled"] = args ? args.isUniqueDomainEnabled : undefined;
            resourceInputs["keepSettingsOnDestroy"] = args ? args.keepSettingsOnDestroy : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["deployments"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectPagesSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectPagesSettings resources.
 */
export interface ProjectPagesSettingsState {
    /**
     * List of current active deployments.
     */
    deployments?: pulumi.Input<pulumi.Input<inputs.ProjectPagesSettingsDeployment>[]>;
    /**
     * Boolean indicating if the project is set to force https. Requires `externalHttps` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
     */
    forceHttps?: pulumi.Input<boolean>;
    /**
     * Boolean indicating if a unique domain is enabled.
     */
    isUniqueDomainEnabled?: pulumi.Input<boolean>;
    keepSettingsOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID or path.
     */
    project?: pulumi.Input<string>;
    /**
     * The URL to access the project pages.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectPagesSettings resource.
 */
export interface ProjectPagesSettingsArgs {
    /**
     * Boolean indicating if the project is set to force https. Requires `externalHttps` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
     */
    forceHttps?: pulumi.Input<boolean>;
    /**
     * Boolean indicating if a unique domain is enabled.
     */
    isUniqueDomainEnabled?: pulumi.Input<boolean>;
    keepSettingsOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID or path.
     */
    project: pulumi.Input<string>;
}
