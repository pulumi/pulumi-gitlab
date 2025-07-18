// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectLevelNotifications` resource manages notifications for a project.
 *
 * > While the API supports both groups and projects, this resource only supports projects currently.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/notification_settings/#group--project-level-notification-settings)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.Project("example", {
 *     name: "example project",
 *     description: "Lorem Ipsum",
 *     visibilityLevel: "public",
 * });
 * // Basic example
 * const notifications = new gitlab.ProjectLevelNotifications("notifications", {
 *     project: example.id,
 *     level: "global",
 * });
 * // Custom notification example
 * const custom = new gitlab.ProjectLevelNotifications("custom", {
 *     project: example.id,
 *     level: "custom",
 *     newMergeRequest: true,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_level_notifications`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_level_notifications.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * A GitLab Project level notification can be imported using a key composed of `<project-id>`, for example:
 *
 * ```sh
 * $ pulumi import gitlab:index/projectLevelNotifications:ProjectLevelNotifications example "12345"
 * ```
 */
export class ProjectLevelNotifications extends pulumi.CustomResource {
    /**
     * Get an existing ProjectLevelNotifications resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectLevelNotificationsState, opts?: pulumi.CustomResourceOptions): ProjectLevelNotifications {
        return new ProjectLevelNotifications(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectLevelNotifications:ProjectLevelNotifications';

    /**
     * Returns true if the given object is an instance of ProjectLevelNotifications.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectLevelNotifications {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectLevelNotifications.__pulumiType;
    }

    /**
     * Enable notifications for closed issues. Can only be used when `level` is `custom`.
     */
    public readonly closeIssue!: pulumi.Output<boolean>;
    /**
     * Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
     */
    public readonly closeMergeRequest!: pulumi.Output<boolean>;
    /**
     * Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
     */
    public readonly failedPipeline!: pulumi.Output<boolean>;
    /**
     * Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
     */
    public readonly fixedPipeline!: pulumi.Output<boolean>;
    /**
     * Enable notifications for due issues. Can only be used when `level` is `custom`.
     */
    public readonly issueDue!: pulumi.Output<boolean>;
    /**
     * The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
     */
    public readonly level!: pulumi.Output<string>;
    /**
     * Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
     */
    public readonly mergeMergeRequest!: pulumi.Output<boolean>;
    /**
     * Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
     */
    public readonly mergeWhenPipelineSucceeds!: pulumi.Output<boolean>;
    /**
     * Enable notifications for moved projects. Can only be used when `level` is `custom`.
     */
    public readonly movedProject!: pulumi.Output<boolean>;
    /**
     * Enable notifications for new issues. Can only be used when `level` is `custom`.
     */
    public readonly newIssue!: pulumi.Output<boolean>;
    /**
     * Enable notifications for new merge requests. Can only be used when `level` is `custom`.
     */
    public readonly newMergeRequest!: pulumi.Output<boolean>;
    /**
     * Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
     */
    public readonly newNote!: pulumi.Output<boolean>;
    /**
     * The ID or URL-encoded path of a project where notifications will be configured.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
     */
    public readonly pushToMergeRequest!: pulumi.Output<boolean>;
    /**
     * Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
     */
    public readonly reassignIssue!: pulumi.Output<boolean>;
    /**
     * Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
     */
    public readonly reassignMergeRequest!: pulumi.Output<boolean>;
    /**
     * Enable notifications for reopened issues. Can only be used when `level` is `custom`.
     */
    public readonly reopenIssue!: pulumi.Output<boolean>;
    /**
     * Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
     */
    public readonly reopenMergeRequest!: pulumi.Output<boolean>;
    /**
     * Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
     */
    public readonly successPipeline!: pulumi.Output<boolean>;

    /**
     * Create a ProjectLevelNotifications resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectLevelNotificationsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectLevelNotificationsArgs | ProjectLevelNotificationsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectLevelNotificationsState | undefined;
            resourceInputs["closeIssue"] = state ? state.closeIssue : undefined;
            resourceInputs["closeMergeRequest"] = state ? state.closeMergeRequest : undefined;
            resourceInputs["failedPipeline"] = state ? state.failedPipeline : undefined;
            resourceInputs["fixedPipeline"] = state ? state.fixedPipeline : undefined;
            resourceInputs["issueDue"] = state ? state.issueDue : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["mergeMergeRequest"] = state ? state.mergeMergeRequest : undefined;
            resourceInputs["mergeWhenPipelineSucceeds"] = state ? state.mergeWhenPipelineSucceeds : undefined;
            resourceInputs["movedProject"] = state ? state.movedProject : undefined;
            resourceInputs["newIssue"] = state ? state.newIssue : undefined;
            resourceInputs["newMergeRequest"] = state ? state.newMergeRequest : undefined;
            resourceInputs["newNote"] = state ? state.newNote : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pushToMergeRequest"] = state ? state.pushToMergeRequest : undefined;
            resourceInputs["reassignIssue"] = state ? state.reassignIssue : undefined;
            resourceInputs["reassignMergeRequest"] = state ? state.reassignMergeRequest : undefined;
            resourceInputs["reopenIssue"] = state ? state.reopenIssue : undefined;
            resourceInputs["reopenMergeRequest"] = state ? state.reopenMergeRequest : undefined;
            resourceInputs["successPipeline"] = state ? state.successPipeline : undefined;
        } else {
            const args = argsOrState as ProjectLevelNotificationsArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["closeIssue"] = args ? args.closeIssue : undefined;
            resourceInputs["closeMergeRequest"] = args ? args.closeMergeRequest : undefined;
            resourceInputs["failedPipeline"] = args ? args.failedPipeline : undefined;
            resourceInputs["fixedPipeline"] = args ? args.fixedPipeline : undefined;
            resourceInputs["issueDue"] = args ? args.issueDue : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["mergeMergeRequest"] = args ? args.mergeMergeRequest : undefined;
            resourceInputs["mergeWhenPipelineSucceeds"] = args ? args.mergeWhenPipelineSucceeds : undefined;
            resourceInputs["movedProject"] = args ? args.movedProject : undefined;
            resourceInputs["newIssue"] = args ? args.newIssue : undefined;
            resourceInputs["newMergeRequest"] = args ? args.newMergeRequest : undefined;
            resourceInputs["newNote"] = args ? args.newNote : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushToMergeRequest"] = args ? args.pushToMergeRequest : undefined;
            resourceInputs["reassignIssue"] = args ? args.reassignIssue : undefined;
            resourceInputs["reassignMergeRequest"] = args ? args.reassignMergeRequest : undefined;
            resourceInputs["reopenIssue"] = args ? args.reopenIssue : undefined;
            resourceInputs["reopenMergeRequest"] = args ? args.reopenMergeRequest : undefined;
            resourceInputs["successPipeline"] = args ? args.successPipeline : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectLevelNotifications.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectLevelNotifications resources.
 */
export interface ProjectLevelNotificationsState {
    /**
     * Enable notifications for closed issues. Can only be used when `level` is `custom`.
     */
    closeIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
     */
    closeMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
     */
    failedPipeline?: pulumi.Input<boolean>;
    /**
     * Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
     */
    fixedPipeline?: pulumi.Input<boolean>;
    /**
     * Enable notifications for due issues. Can only be used when `level` is `custom`.
     */
    issueDue?: pulumi.Input<boolean>;
    /**
     * The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
     */
    level?: pulumi.Input<string>;
    /**
     * Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
     */
    mergeMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
     */
    mergeWhenPipelineSucceeds?: pulumi.Input<boolean>;
    /**
     * Enable notifications for moved projects. Can only be used when `level` is `custom`.
     */
    movedProject?: pulumi.Input<boolean>;
    /**
     * Enable notifications for new issues. Can only be used when `level` is `custom`.
     */
    newIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for new merge requests. Can only be used when `level` is `custom`.
     */
    newMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
     */
    newNote?: pulumi.Input<boolean>;
    /**
     * The ID or URL-encoded path of a project where notifications will be configured.
     */
    project?: pulumi.Input<string>;
    /**
     * Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
     */
    pushToMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
     */
    reassignIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
     */
    reassignMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for reopened issues. Can only be used when `level` is `custom`.
     */
    reopenIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
     */
    reopenMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
     */
    successPipeline?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ProjectLevelNotifications resource.
 */
export interface ProjectLevelNotificationsArgs {
    /**
     * Enable notifications for closed issues. Can only be used when `level` is `custom`.
     */
    closeIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
     */
    closeMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
     */
    failedPipeline?: pulumi.Input<boolean>;
    /**
     * Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
     */
    fixedPipeline?: pulumi.Input<boolean>;
    /**
     * Enable notifications for due issues. Can only be used when `level` is `custom`.
     */
    issueDue?: pulumi.Input<boolean>;
    /**
     * The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
     */
    level?: pulumi.Input<string>;
    /**
     * Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
     */
    mergeMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
     */
    mergeWhenPipelineSucceeds?: pulumi.Input<boolean>;
    /**
     * Enable notifications for moved projects. Can only be used when `level` is `custom`.
     */
    movedProject?: pulumi.Input<boolean>;
    /**
     * Enable notifications for new issues. Can only be used when `level` is `custom`.
     */
    newIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for new merge requests. Can only be used when `level` is `custom`.
     */
    newMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
     */
    newNote?: pulumi.Input<boolean>;
    /**
     * The ID or URL-encoded path of a project where notifications will be configured.
     */
    project: pulumi.Input<string>;
    /**
     * Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
     */
    pushToMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
     */
    reassignIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
     */
    reassignMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for reopened issues. Can only be used when `level` is `custom`.
     */
    reopenIssue?: pulumi.Input<boolean>;
    /**
     * Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
     */
    reopenMergeRequest?: pulumi.Input<boolean>;
    /**
     * Enable notifications for successful pipelines. Can only be used when `level` is `custom`.
     */
    successPipeline?: pulumi.Input<boolean>;
}
