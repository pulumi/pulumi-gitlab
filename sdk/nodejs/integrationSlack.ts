// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationSlack` resource allows you to manage the lifecycle of a project integration with Slack.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#slack-notifications)
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
 * const slack = new gitlab.IntegrationSlack("slack", {
 *     project: awesomeProject.id,
 *     webhook: "https://webhook.com",
 *     username: "myuser",
 *     pushEvents: true,
 *     pushChannel: "push_chan",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_integration_slack`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_integration_slack.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * You can import a gitlab_integration_slack.slack state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationSlack:IntegrationSlack slack 1
 * ```
 */
export class IntegrationSlack extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationSlack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationSlackState, opts?: pulumi.CustomResourceOptions): IntegrationSlack {
        return new IntegrationSlack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationSlack:IntegrationSlack';

    /**
     * Returns true if the given object is an instance of IntegrationSlack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationSlack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationSlack.__pulumiType;
    }

    /**
     * Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
     */
    public readonly branchesToBeNotified!: pulumi.Output<string>;
    /**
     * The name of the channel to receive confidential issue events notifications.
     */
    public readonly confidentialIssueChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for confidential issues events.
     */
    public readonly confidentialIssuesEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive confidential note events notifications.
     */
    public readonly confidentialNoteChannel!: pulumi.Output<string>;
    /**
     * Enable notifications for confidential note events.
     */
    public readonly confidentialNoteEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive issue events notifications.
     */
    public readonly issueChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to https://gitlab.com/gitlab-org/api/client-go/issues/1354.
     */
    public /*out*/ readonly jobEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive merge request events notifications.
     */
    public readonly mergeRequestChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for merge requests events.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive note events notifications.
     */
    public readonly noteChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for note events.
     */
    public readonly noteEvents!: pulumi.Output<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    public readonly notifyOnlyBrokenPipelines!: pulumi.Output<boolean>;
    /**
     * This parameter has been replaced with `branchesToBeNotified`.
     *
     * @deprecated use 'branches_to_be_notified' argument instead
     */
    public readonly notifyOnlyDefaultBranch!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive pipeline events notifications.
     */
    public readonly pipelineChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the channel to receive push events notifications.
     */
    public readonly pushChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive tag push events notifications.
     */
    public readonly tagPushChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean>;
    /**
     * Username to use.
     */
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
     */
    public readonly webhook!: pulumi.Output<string>;
    /**
     * The name of the channel to receive wiki page events notifications.
     */
    public readonly wikiPageChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for wiki page events.
     */
    public readonly wikiPageEvents!: pulumi.Output<boolean>;

    /**
     * Create a IntegrationSlack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationSlackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationSlackArgs | IntegrationSlackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationSlackState | undefined;
            resourceInputs["branchesToBeNotified"] = state ? state.branchesToBeNotified : undefined;
            resourceInputs["confidentialIssueChannel"] = state ? state.confidentialIssueChannel : undefined;
            resourceInputs["confidentialIssuesEvents"] = state ? state.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteChannel"] = state ? state.confidentialNoteChannel : undefined;
            resourceInputs["confidentialNoteEvents"] = state ? state.confidentialNoteEvents : undefined;
            resourceInputs["issueChannel"] = state ? state.issueChannel : undefined;
            resourceInputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            resourceInputs["jobEvents"] = state ? state.jobEvents : undefined;
            resourceInputs["mergeRequestChannel"] = state ? state.mergeRequestChannel : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["noteChannel"] = state ? state.noteChannel : undefined;
            resourceInputs["noteEvents"] = state ? state.noteEvents : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = state ? state.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["notifyOnlyDefaultBranch"] = state ? state.notifyOnlyDefaultBranch : undefined;
            resourceInputs["pipelineChannel"] = state ? state.pipelineChannel : undefined;
            resourceInputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pushChannel"] = state ? state.pushChannel : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["tagPushChannel"] = state ? state.tagPushChannel : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["webhook"] = state ? state.webhook : undefined;
            resourceInputs["wikiPageChannel"] = state ? state.wikiPageChannel : undefined;
            resourceInputs["wikiPageEvents"] = state ? state.wikiPageEvents : undefined;
        } else {
            const args = argsOrState as IntegrationSlackArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.webhook === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webhook'");
            }
            resourceInputs["branchesToBeNotified"] = args ? args.branchesToBeNotified : undefined;
            resourceInputs["confidentialIssueChannel"] = args ? args.confidentialIssueChannel : undefined;
            resourceInputs["confidentialIssuesEvents"] = args ? args.confidentialIssuesEvents : undefined;
            resourceInputs["confidentialNoteChannel"] = args ? args.confidentialNoteChannel : undefined;
            resourceInputs["confidentialNoteEvents"] = args ? args.confidentialNoteEvents : undefined;
            resourceInputs["issueChannel"] = args ? args.issueChannel : undefined;
            resourceInputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            resourceInputs["mergeRequestChannel"] = args ? args.mergeRequestChannel : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["noteChannel"] = args ? args.noteChannel : undefined;
            resourceInputs["noteEvents"] = args ? args.noteEvents : undefined;
            resourceInputs["notifyOnlyBrokenPipelines"] = args ? args.notifyOnlyBrokenPipelines : undefined;
            resourceInputs["notifyOnlyDefaultBranch"] = args ? args.notifyOnlyDefaultBranch : undefined;
            resourceInputs["pipelineChannel"] = args ? args.pipelineChannel : undefined;
            resourceInputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushChannel"] = args ? args.pushChannel : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["tagPushChannel"] = args ? args.tagPushChannel : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["webhook"] = args ? args.webhook : undefined;
            resourceInputs["wikiPageChannel"] = args ? args.wikiPageChannel : undefined;
            resourceInputs["wikiPageEvents"] = args ? args.wikiPageEvents : undefined;
            resourceInputs["jobEvents"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IntegrationSlack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationSlack resources.
 */
export interface IntegrationSlackState {
    /**
     * Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * The name of the channel to receive confidential issue events notifications.
     */
    confidentialIssueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive confidential note events notifications.
     */
    confidentialNoteChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive issue events notifications.
     */
    issueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for job events. **ATTENTION**: This attribute is currently not being submitted to the GitLab API, due to https://gitlab.com/gitlab-org/api/client-go/issues/1354.
     */
    jobEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive merge request events notifications.
     */
    mergeRequestChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive note events notifications.
     */
    noteChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * This parameter has been replaced with `branchesToBeNotified`.
     *
     * @deprecated use 'branches_to_be_notified' argument instead
     */
    notifyOnlyDefaultBranch?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive pipeline events notifications.
     */
    pipelineChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the channel to receive push events notifications.
     */
    pushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive tag push events notifications.
     */
    tagPushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Username to use.
     */
    username?: pulumi.Input<string>;
    /**
     * Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
     */
    webhook?: pulumi.Input<string>;
    /**
     * The name of the channel to receive wiki page events notifications.
     */
    wikiPageChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for wiki page events.
     */
    wikiPageEvents?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IntegrationSlack resource.
 */
export interface IntegrationSlackArgs {
    /**
     * Branches to send notifications for. Valid options are "all", "default", "protected", and "default*and*protected".
     */
    branchesToBeNotified?: pulumi.Input<string>;
    /**
     * The name of the channel to receive confidential issue events notifications.
     */
    confidentialIssueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive confidential note events notifications.
     */
    confidentialNoteChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential note events.
     */
    confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive issue events notifications.
     */
    issueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive merge request events notifications.
     */
    mergeRequestChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for merge requests events.
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive note events notifications.
     */
    noteChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * This parameter has been replaced with `branchesToBeNotified`.
     *
     * @deprecated use 'branches_to_be_notified' argument instead
     */
    notifyOnlyDefaultBranch?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive pipeline events notifications.
     */
    pipelineChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * The name of the channel to receive push events notifications.
     */
    pushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive tag push events notifications.
     */
    tagPushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Username to use.
     */
    username?: pulumi.Input<string>;
    /**
     * Webhook URL (Example, https://hooks.slack.com/services/...). This value cannot be imported.
     */
    webhook: pulumi.Input<string>;
    /**
     * The name of the channel to receive wiki page events notifications.
     */
    wikiPageChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for wiki page events.
     */
    wikiPageEvents?: pulumi.Input<boolean>;
}
