// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationPipelinesEmail` resource manages the lifecycle of a project integration with the Pipeline Emails Service.
 *
 * > This resource is deprecated and will be removed in 19.0. Use `gitlab.ProjectIntegrationPipelinesEmail`instead!
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#pipeline-status-emails)
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
 * const email = new gitlab.IntegrationPipelinesEmail("email", {
 *     project: awesomeProject.id,
 *     recipients: ["gitlab@user.create"],
 *     notifyOnlyBrokenPipelines: true,
 *     branchesToBeNotified: "all",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_integration_pipelines_email`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_integration_pipelines_email.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * You can import a gitlab_integration_pipelines_email state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail email 1
 * ```
 */
export class IntegrationPipelinesEmail extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationPipelinesEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationPipelinesEmailState, opts?: pulumi.CustomResourceOptions): IntegrationPipelinesEmail {
        return new IntegrationPipelinesEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail';

    /**
     * Returns true if the given object is an instance of IntegrationPipelinesEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationPipelinesEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationPipelinesEmail.__pulumiType;
    }

    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
     */
    public readonly branchesToBeNotified!: pulumi.Output<string | undefined>;
    /**
     * Notify only broken pipelines. Default is true.
     */
    public readonly notifyOnlyBrokenPipelines!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * ) email addresses where notifications are sent.
     */
    public readonly recipients!: pulumi.Output<string[]>;

    /**
     * Create a IntegrationPipelinesEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationPipelinesEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationPipelinesEmailArgs | IntegrationPipelinesEmailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationPipelinesEmailState | undefined;
            resourceInputs["branchesToBeNotified"] = state ? state.branchesToBeNotified : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = state ? state.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["recipients"] = state ? state.recipients : undefined;
        } else {
            const args = argsOrState as IntegrationPipelinesEmailArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.recipients === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recipients'");
            }
            resourceInputs["branchesToBeNotified"] = args ? args.branchesToBeNotified : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = args ? args.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["recipients"] = args ? args.recipients : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IntegrationPipelinesEmail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationPipelinesEmail resources.
 */
export interface IntegrationPipelinesEmailState {
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * Notify only broken pipelines. Default is true.
     */
    notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * ) email addresses where notifications are sent.
     */
    recipients?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a IntegrationPipelinesEmail resource.
 */
export interface IntegrationPipelinesEmailArgs {
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * Notify only broken pipelines. Default is true.
     */
    notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * ) email addresses where notifications are sent.
     */
    recipients: pulumi.Input<pulumi.Input<string>[]>;
}
