// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationTelegram` resource allows to manage the lifecycle of a project integration with Telegram.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#telegram)
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
 * const _default = new gitlab.IntegrationTelegram("default", {
 *     token: "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
 *     room: "-1000000000000000",
 *     notifyOnlyBrokenPipelines: false,
 *     branchesToBeNotified: "all",
 *     pushEvents: false,
 *     issuesEvents: false,
 *     confidentialIssuesEvents: false,
 *     mergeRequestsEvents: false,
 *     tagPushEvents: false,
 *     noteEvents: false,
 *     confidentialNoteEvents: false,
 *     pipelineEvents: false,
 *     wikiPageEvents: false,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_telegram`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_integration_telegram.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * You can import a gitlab_integration_telegram state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationTelegram:IntegrationTelegram default 1
 * ```
 */
export class IntegrationTelegram extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationTelegram resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationTelegramState, opts?: pulumi.CustomResourceOptions): IntegrationTelegram {
        return new IntegrationTelegram(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationTelegram:IntegrationTelegram';

    /**
     * Returns true if the given object is an instance of IntegrationTelegram.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationTelegram {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationTelegram.__pulumiType;
    }

    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
     */
    public readonly branchesToBeNotified!: pulumi.Output<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    public readonly confidentialIssuesEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for confidential note events.
     */
    public readonly confidentialNoteEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for merge requests events.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for note events.
     */
    public readonly noteEvents!: pulumi.Output<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    public readonly notifyOnlyBrokenPipelines!: pulumi.Output<boolean>;
    /**
     * Enable notifications for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean>;
    /**
     * The ID or full path of the project to integrate with Telegram.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Enable notifications for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean>;
    /**
     * Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
     */
    public readonly room!: pulumi.Output<string>;
    /**
     * Enable notifications for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean>;
    /**
     * The Telegram bot token.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * Enable notifications for wiki page events.
     */
    public readonly wikiPageEvents!: pulumi.Output<boolean>;

    /**
     * Create a IntegrationTelegram resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationTelegramArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationTelegramArgs | IntegrationTelegramState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationTelegramState | undefined;
            resourceInputs["branchesToBeNotified"] = state ? state.branchesToBeNotified : undefined;
            resourceInputs["confidentialIssuesEvents"] = state ? state.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteEvents"] = state ? state.confidentialNoteEvents : undefined;
            resourceInputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = state ? state.noteEvents : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = state ? state.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["room"] = state ? state.room : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["wikiPageEvents"] = state ? state.wikiPageEvents : undefined;
        } else {
            const args = argsOrState as IntegrationTelegramArgs | undefined;
            if ((!args || args.confidentialIssuesEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'confidentialIssuesEvents'");
            }
            if ((!args || args.confidentialNoteEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'confidentialNoteEvents'");
            }
            if ((!args || args.issuesEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuesEvents'");
            }
            if ((!args || args.mergeRequestsEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mergeRequestsEvents'");
            }
            if ((!args || args.noteEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'noteEvents'");
            }
            if ((!args || args.pipelineEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineEvents'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.pushEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pushEvents'");
            }
            if ((!args || args.room === undefined) && !opts.urn) {
                throw new Error("Missing required property 'room'");
            }
            if ((!args || args.tagPushEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagPushEvents'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.wikiPageEvents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'wikiPageEvents'");
            }
            resourceInputs["branchesToBeNotified"] = args ? args.branchesToBeNotified : undefined;
            resourceInputs["confidentialIssuesEvents"] = args ? args.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteEvents"] = args ? args.confidentialNoteEvents : undefined;
            resourceInputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = args ? args.noteEvents : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = args ? args.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["room"] = args ? args.room : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["wikiPageEvents"] = args ? args.wikiPageEvents : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IntegrationTelegram.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationTelegram resources.
 */
export interface IntegrationTelegramState {
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * Enable notifications for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * The ID or full path of the project to integrate with Telegram.
     */
    project?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
     */
    room?: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * The Telegram bot token.
     */
    token?: pulumi.Input<string>;
    /**
     * Enable notifications for wiki page events.
     */
    wikiPageEvents?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IntegrationTelegram resource.
 */
export interface IntegrationTelegramArgs {
    /**
     * Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    confidentialIssuesEvents: pulumi.Input<boolean>;
    /**
     * Enable notifications for confidential note events.
     */
    confidentialNoteEvents: pulumi.Input<boolean>;
    /**
     * Enable notifications for issues events.
     */
    issuesEvents: pulumi.Input<boolean>;
    /**
     * Enable notifications for merge requests events.
     */
    mergeRequestsEvents: pulumi.Input<boolean>;
    /**
     * Enable notifications for note events.
     */
    noteEvents: pulumi.Input<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * Enable notifications for pipeline events.
     */
    pipelineEvents: pulumi.Input<boolean>;
    /**
     * The ID or full path of the project to integrate with Telegram.
     */
    project: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents: pulumi.Input<boolean>;
    /**
     * Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
     */
    room: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents: pulumi.Input<boolean>;
    /**
     * The Telegram bot token.
     */
    token: pulumi.Input<string>;
    /**
     * Enable notifications for wiki page events.
     */
    wikiPageEvents: pulumi.Input<boolean>;
}
