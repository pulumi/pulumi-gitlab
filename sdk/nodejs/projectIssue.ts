// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.Project("foo", {
 *     name: "example project",
 *     description: "Lorem Ipsum",
 *     visibilityLevel: "public",
 * });
 * const welcomeIssue = new gitlab.ProjectIssue("welcome_issue", {
 *     project: foo.id,
 *     title: "Welcome!",
 *     description: pulumi.interpolate`  Welcome to the ${foo.name} project!
 *
 * `,
 *     discussionLocked: true,
 * });
 * export const welcomeIssueWebUrl = webUrl;
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_issue`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_issue.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
 *
 * You can import this resource with an id made up of `{project-id}:{issue-id}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectIssue:ProjectIssue welcome_issue 42:1
 * ```
 */
export class ProjectIssue extends pulumi.CustomResource {
    /**
     * Get an existing ProjectIssue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectIssueState, opts?: pulumi.CustomResourceOptions): ProjectIssue {
        return new ProjectIssue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectIssue:ProjectIssue';

    /**
     * Returns true if the given object is an instance of ProjectIssue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectIssue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectIssue.__pulumiType;
    }

    /**
     * The IDs of the users to assign the issue to.
     */
    public readonly assigneeIds!: pulumi.Output<number[] | undefined>;
    /**
     * The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     */
    public /*out*/ readonly authorId!: pulumi.Output<number>;
    /**
     * When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    public /*out*/ readonly closedAt!: pulumi.Output<string>;
    /**
     * The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     */
    public /*out*/ readonly closedByUserId!: pulumi.Output<number>;
    /**
     * Set an issue to be confidential.
     */
    public readonly confidential!: pulumi.Output<boolean | undefined>;
    /**
     * When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     */
    public readonly createdAt!: pulumi.Output<string>;
    /**
     * Whether the issue is deleted instead of closed during destroy.
     */
    public readonly deleteOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The description of an issue. Limited to 1,048,576 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the issue is locked for discussions or not.
     */
    public readonly discussionLocked!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     */
    public readonly discussionToResolve!: pulumi.Output<string | undefined>;
    /**
     * The number of downvotes the issue has received.
     */
    public /*out*/ readonly downvotes!: pulumi.Output<number>;
    /**
     * The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    public readonly dueDate!: pulumi.Output<string | undefined>;
    /**
     * ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     */
    public /*out*/ readonly epicId!: pulumi.Output<number>;
    /**
     * The ID of the epic issue.
     */
    public readonly epicIssueId!: pulumi.Output<number>;
    /**
     * The external ID of the issue.
     */
    public /*out*/ readonly externalId!: pulumi.Output<string>;
    /**
     * The human-readable time estimate of the issue.
     */
    public /*out*/ readonly humanTimeEstimate!: pulumi.Output<string>;
    /**
     * The human-readable total time spent of the issue.
     */
    public /*out*/ readonly humanTotalTimeSpent!: pulumi.Output<string>;
    /**
     * The internal ID of the project's issue.
     */
    public readonly iid!: pulumi.Output<number>;
    /**
     * The instance-wide ID of the issue.
     */
    public /*out*/ readonly issueId!: pulumi.Output<number>;
    /**
     * The ID of the issue link.
     */
    public /*out*/ readonly issueLinkId!: pulumi.Output<number>;
    /**
     * The type of issue. Valid values are: `issue`, `incident`, `testCase`.
     */
    public readonly issueType!: pulumi.Output<string | undefined>;
    /**
     * The labels of an issue.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The links of the issue.
     */
    public /*out*/ readonly links!: pulumi.Output<{[key: string]: string}>;
    /**
     * The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     */
    public readonly mergeRequestToResolveDiscussionsOf!: pulumi.Output<number | undefined>;
    /**
     * The number of merge requests associated with the issue.
     */
    public /*out*/ readonly mergeRequestsCount!: pulumi.Output<number>;
    /**
     * The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
     */
    public readonly milestoneId!: pulumi.Output<number | undefined>;
    /**
     * The ID of the issue that was moved to.
     */
    public /*out*/ readonly movedToId!: pulumi.Output<number>;
    /**
     * The name or ID of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The references of the issue.
     */
    public /*out*/ readonly references!: pulumi.Output<{[key: string]: string}>;
    /**
     * The state of the issue. Valid values are: `opened`, `closed`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * Whether the authenticated user is subscribed to the issue or not.
     */
    public /*out*/ readonly subscribed!: pulumi.Output<boolean>;
    /**
     * The task completion status. It's always a one element list.
     */
    public /*out*/ readonly taskCompletionStatuses!: pulumi.Output<outputs.ProjectIssueTaskCompletionStatus[]>;
    /**
     * The time estimate of the issue.
     */
    public /*out*/ readonly timeEstimate!: pulumi.Output<number>;
    /**
     * The title of the issue.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The total time spent of the issue.
     */
    public /*out*/ readonly totalTimeSpent!: pulumi.Output<number>;
    /**
     * When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    public readonly updatedAt!: pulumi.Output<string>;
    /**
     * The number of upvotes the issue has received.
     */
    public /*out*/ readonly upvotes!: pulumi.Output<number>;
    /**
     * The number of user notes on the issue.
     */
    public /*out*/ readonly userNotesCount!: pulumi.Output<number>;
    /**
     * The web URL of the issue.
     */
    public /*out*/ readonly webUrl!: pulumi.Output<string>;
    /**
     * The weight of the issue. Valid values are greater than or equal to 0.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a ProjectIssue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectIssueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectIssueArgs | ProjectIssueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectIssueState | undefined;
            resourceInputs["assigneeIds"] = state ? state.assigneeIds : undefined;
            resourceInputs["authorId"] = state ? state.authorId : undefined;
            resourceInputs["closedAt"] = state ? state.closedAt : undefined;
            resourceInputs["closedByUserId"] = state ? state.closedByUserId : undefined;
            resourceInputs["confidential"] = state ? state.confidential : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deleteOnDestroy"] = state ? state.deleteOnDestroy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["discussionLocked"] = state ? state.discussionLocked : undefined;
            resourceInputs["discussionToResolve"] = state ? state.discussionToResolve : undefined;
            resourceInputs["downvotes"] = state ? state.downvotes : undefined;
            resourceInputs["dueDate"] = state ? state.dueDate : undefined;
            resourceInputs["epicId"] = state ? state.epicId : undefined;
            resourceInputs["epicIssueId"] = state ? state.epicIssueId : undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["humanTimeEstimate"] = state ? state.humanTimeEstimate : undefined;
            resourceInputs["humanTotalTimeSpent"] = state ? state.humanTotalTimeSpent : undefined;
            resourceInputs["iid"] = state ? state.iid : undefined;
            resourceInputs["issueId"] = state ? state.issueId : undefined;
            resourceInputs["issueLinkId"] = state ? state.issueLinkId : undefined;
            resourceInputs["issueType"] = state ? state.issueType : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["mergeRequestToResolveDiscussionsOf"] = state ? state.mergeRequestToResolveDiscussionsOf : undefined;
            resourceInputs["mergeRequestsCount"] = state ? state.mergeRequestsCount : undefined;
            resourceInputs["milestoneId"] = state ? state.milestoneId : undefined;
            resourceInputs["movedToId"] = state ? state.movedToId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["references"] = state ? state.references : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["subscribed"] = state ? state.subscribed : undefined;
            resourceInputs["taskCompletionStatuses"] = state ? state.taskCompletionStatuses : undefined;
            resourceInputs["timeEstimate"] = state ? state.timeEstimate : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["totalTimeSpent"] = state ? state.totalTimeSpent : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["upvotes"] = state ? state.upvotes : undefined;
            resourceInputs["userNotesCount"] = state ? state.userNotesCount : undefined;
            resourceInputs["webUrl"] = state ? state.webUrl : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as ProjectIssueArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["assigneeIds"] = args ? args.assigneeIds : undefined;
            resourceInputs["confidential"] = args ? args.confidential : undefined;
            resourceInputs["createdAt"] = args ? args.createdAt : undefined;
            resourceInputs["deleteOnDestroy"] = args ? args.deleteOnDestroy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["discussionLocked"] = args ? args.discussionLocked : undefined;
            resourceInputs["discussionToResolve"] = args ? args.discussionToResolve : undefined;
            resourceInputs["dueDate"] = args ? args.dueDate : undefined;
            resourceInputs["epicIssueId"] = args ? args.epicIssueId : undefined;
            resourceInputs["iid"] = args ? args.iid : undefined;
            resourceInputs["issueType"] = args ? args.issueType : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["mergeRequestToResolveDiscussionsOf"] = args ? args.mergeRequestToResolveDiscussionsOf : undefined;
            resourceInputs["milestoneId"] = args ? args.milestoneId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["updatedAt"] = args ? args.updatedAt : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["authorId"] = undefined /*out*/;
            resourceInputs["closedAt"] = undefined /*out*/;
            resourceInputs["closedByUserId"] = undefined /*out*/;
            resourceInputs["downvotes"] = undefined /*out*/;
            resourceInputs["epicId"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["humanTimeEstimate"] = undefined /*out*/;
            resourceInputs["humanTotalTimeSpent"] = undefined /*out*/;
            resourceInputs["issueId"] = undefined /*out*/;
            resourceInputs["issueLinkId"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["mergeRequestsCount"] = undefined /*out*/;
            resourceInputs["movedToId"] = undefined /*out*/;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["subscribed"] = undefined /*out*/;
            resourceInputs["taskCompletionStatuses"] = undefined /*out*/;
            resourceInputs["timeEstimate"] = undefined /*out*/;
            resourceInputs["totalTimeSpent"] = undefined /*out*/;
            resourceInputs["upvotes"] = undefined /*out*/;
            resourceInputs["userNotesCount"] = undefined /*out*/;
            resourceInputs["webUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectIssue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectIssue resources.
 */
export interface ProjectIssueState {
    /**
     * The IDs of the users to assign the issue to.
     */
    assigneeIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     */
    authorId?: pulumi.Input<number>;
    /**
     * When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    closedAt?: pulumi.Input<string>;
    /**
     * The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     */
    closedByUserId?: pulumi.Input<number>;
    /**
     * Set an issue to be confidential.
     */
    confidential?: pulumi.Input<boolean>;
    /**
     * When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Whether the issue is deleted instead of closed during destroy.
     */
    deleteOnDestroy?: pulumi.Input<boolean>;
    /**
     * The description of an issue. Limited to 1,048,576 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the issue is locked for discussions or not.
     */
    discussionLocked?: pulumi.Input<boolean>;
    /**
     * The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     */
    discussionToResolve?: pulumi.Input<string>;
    /**
     * The number of downvotes the issue has received.
     */
    downvotes?: pulumi.Input<number>;
    /**
     * The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    dueDate?: pulumi.Input<string>;
    /**
     * ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     */
    epicId?: pulumi.Input<number>;
    /**
     * The ID of the epic issue.
     */
    epicIssueId?: pulumi.Input<number>;
    /**
     * The external ID of the issue.
     */
    externalId?: pulumi.Input<string>;
    /**
     * The human-readable time estimate of the issue.
     */
    humanTimeEstimate?: pulumi.Input<string>;
    /**
     * The human-readable total time spent of the issue.
     */
    humanTotalTimeSpent?: pulumi.Input<string>;
    /**
     * The internal ID of the project's issue.
     */
    iid?: pulumi.Input<number>;
    /**
     * The instance-wide ID of the issue.
     */
    issueId?: pulumi.Input<number>;
    /**
     * The ID of the issue link.
     */
    issueLinkId?: pulumi.Input<number>;
    /**
     * The type of issue. Valid values are: `issue`, `incident`, `testCase`.
     */
    issueType?: pulumi.Input<string>;
    /**
     * The labels of an issue.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The links of the issue.
     */
    links?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     */
    mergeRequestToResolveDiscussionsOf?: pulumi.Input<number>;
    /**
     * The number of merge requests associated with the issue.
     */
    mergeRequestsCount?: pulumi.Input<number>;
    /**
     * The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
     */
    milestoneId?: pulumi.Input<number>;
    /**
     * The ID of the issue that was moved to.
     */
    movedToId?: pulumi.Input<number>;
    /**
     * The name or ID of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * The references of the issue.
     */
    references?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The state of the issue. Valid values are: `opened`, `closed`.
     */
    state?: pulumi.Input<string>;
    /**
     * Whether the authenticated user is subscribed to the issue or not.
     */
    subscribed?: pulumi.Input<boolean>;
    /**
     * The task completion status. It's always a one element list.
     */
    taskCompletionStatuses?: pulumi.Input<pulumi.Input<inputs.ProjectIssueTaskCompletionStatus>[]>;
    /**
     * The time estimate of the issue.
     */
    timeEstimate?: pulumi.Input<number>;
    /**
     * The title of the issue.
     */
    title?: pulumi.Input<string>;
    /**
     * The total time spent of the issue.
     */
    totalTimeSpent?: pulumi.Input<number>;
    /**
     * When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The number of upvotes the issue has received.
     */
    upvotes?: pulumi.Input<number>;
    /**
     * The number of user notes on the issue.
     */
    userNotesCount?: pulumi.Input<number>;
    /**
     * The web URL of the issue.
     */
    webUrl?: pulumi.Input<string>;
    /**
     * The weight of the issue. Valid values are greater than or equal to 0.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectIssue resource.
 */
export interface ProjectIssueArgs {
    /**
     * The IDs of the users to assign the issue to.
     */
    assigneeIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Set an issue to be confidential.
     */
    confidential?: pulumi.Input<boolean>;
    /**
     * When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Whether the issue is deleted instead of closed during destroy.
     */
    deleteOnDestroy?: pulumi.Input<boolean>;
    /**
     * The description of an issue. Limited to 1,048,576 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the issue is locked for discussions or not.
     */
    discussionLocked?: pulumi.Input<boolean>;
    /**
     * The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     */
    discussionToResolve?: pulumi.Input<string>;
    /**
     * The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    dueDate?: pulumi.Input<string>;
    /**
     * The ID of the epic issue.
     */
    epicIssueId?: pulumi.Input<number>;
    /**
     * The internal ID of the project's issue.
     */
    iid?: pulumi.Input<number>;
    /**
     * The type of issue. Valid values are: `issue`, `incident`, `testCase`.
     */
    issueType?: pulumi.Input<string>;
    /**
     * The labels of an issue.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     */
    mergeRequestToResolveDiscussionsOf?: pulumi.Input<number>;
    /**
     * The global ID of a milestone to assign issue. To find the milestoneId associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue's details.
     */
    milestoneId?: pulumi.Input<number>;
    /**
     * The name or ID of the project.
     */
    project: pulumi.Input<string>;
    /**
     * The state of the issue. Valid values are: `opened`, `closed`.
     */
    state?: pulumi.Input<string>;
    /**
     * The title of the issue.
     */
    title: pulumi.Input<string>;
    /**
     * When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The weight of the issue. Valid values are greater than or equal to 0.
     */
    weight?: pulumi.Input<number>;
}
