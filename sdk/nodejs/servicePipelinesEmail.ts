// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ServicePipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.
 *
 * > This resource is deprecated. use `gitlab.IntegrationPipelinesEmail`instead!
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const awesomeProject = new gitlab.Project("awesomeProject", {
 *     description: "My awesome project.",
 *     visibilityLevel: "public",
 * });
 * const email = new gitlab.ServicePipelinesEmail("email", {
 *     project: awesomeProject.id,
 *     recipients: ["gitlab@user.create"],
 *     notifyOnlyBrokenPipelines: true,
 *     branchesToBeNotified: "all",
 * });
 * ```
 *
 * ## Import
 *
 * You can import a gitlab_service_pipelines_email state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/servicePipelinesEmail:ServicePipelinesEmail email 1
 * ```
 */
export class ServicePipelinesEmail extends pulumi.CustomResource {
    /**
     * Get an existing ServicePipelinesEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePipelinesEmailState, opts?: pulumi.CustomResourceOptions): ServicePipelinesEmail {
        return new ServicePipelinesEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/servicePipelinesEmail:ServicePipelinesEmail';

    /**
     * Returns true if the given object is an instance of ServicePipelinesEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePipelinesEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePipelinesEmail.__pulumiType;
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
     * Create a ServicePipelinesEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePipelinesEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePipelinesEmailArgs | ServicePipelinesEmailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePipelinesEmailState | undefined;
            resourceInputs["branchesToBeNotified"] = state ? state.branchesToBeNotified : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = state ? state.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["recipients"] = state ? state.recipients : undefined;
        } else {
            const args = argsOrState as ServicePipelinesEmailArgs | undefined;
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
        super(ServicePipelinesEmail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePipelinesEmail resources.
 */
export interface ServicePipelinesEmailState {
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
 * The set of arguments for constructing a ServicePipelinesEmail resource.
 */
export interface ServicePipelinesEmailArgs {
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
