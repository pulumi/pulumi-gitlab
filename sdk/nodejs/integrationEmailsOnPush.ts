// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#emails-on-push)
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
 * const emails = new gitlab.IntegrationEmailsOnPush("emails", {
 *     project: awesomeProject.id,
 *     recipients: "myrecipient@example.com myotherrecipient@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_emails_on_push`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_integration_emails_on_push.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * You can import a gitlab_integration_emails_on_push state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush emails 1
 * ```
 */
export class IntegrationEmailsOnPush extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationEmailsOnPush resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationEmailsOnPushState, opts?: pulumi.CustomResourceOptions): IntegrationEmailsOnPush {
        return new IntegrationEmailsOnPush(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush';

    /**
     * Returns true if the given object is an instance of IntegrationEmailsOnPush.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationEmailsOnPush {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationEmailsOnPush.__pulumiType;
    }

    /**
     * Whether the integration is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
     */
    public readonly branchesToBeNotified!: pulumi.Output<string | undefined>;
    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Disable code diffs.
     */
    public readonly disableDiffs!: pulumi.Output<boolean | undefined>;
    /**
     * ID or full-path of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Enable notifications for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Emails separated by whitespace.
     */
    public readonly recipients!: pulumi.Output<string>;
    /**
     * Send from committer.
     */
    public readonly sendFromCommitterEmail!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    public /*out*/ readonly slug!: pulumi.Output<string>;
    /**
     * Enable notifications for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Title of the integration.
     */
    public /*out*/ readonly title!: pulumi.Output<string>;
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IntegrationEmailsOnPush resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationEmailsOnPushArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationEmailsOnPushArgs | IntegrationEmailsOnPushState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationEmailsOnPushState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["branchesToBeNotified"] = state ? state.branchesToBeNotified : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["disableDiffs"] = state ? state.disableDiffs : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["recipients"] = state ? state.recipients : undefined;
            resourceInputs["sendFromCommitterEmail"] = state ? state.sendFromCommitterEmail : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IntegrationEmailsOnPushArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.recipients === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recipients'");
            }
            resourceInputs["branchesToBeNotified"] = args ? args.branchesToBeNotified : undefined;
            resourceInputs["disableDiffs"] = args ? args.disableDiffs : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["recipients"] = args ? args.recipients : undefined;
            resourceInputs["sendFromCommitterEmail"] = args ? args.sendFromCommitterEmail : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IntegrationEmailsOnPush.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationEmailsOnPush resources.
 */
export interface IntegrationEmailsOnPushState {
    /**
     * Whether the integration is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Disable code diffs.
     */
    disableDiffs?: pulumi.Input<boolean>;
    /**
     * ID or full-path of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Emails separated by whitespace.
     */
    recipients?: pulumi.Input<string>;
    /**
     * Send from committer.
     */
    sendFromCommitterEmail?: pulumi.Input<boolean>;
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    slug?: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
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
 * The set of arguments for constructing a IntegrationEmailsOnPush resource.
 */
export interface IntegrationEmailsOnPushArgs {
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * Disable code diffs.
     */
    disableDiffs?: pulumi.Input<boolean>;
    /**
     * ID or full-path of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Emails separated by whitespace.
     */
    recipients: pulumi.Input<string>;
    /**
     * Send from committer.
     */
    sendFromCommitterEmail?: pulumi.Input<boolean>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
}
