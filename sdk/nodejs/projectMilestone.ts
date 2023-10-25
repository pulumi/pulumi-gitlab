// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectMilestone` resource allows to manage the lifecycle of a project milestone.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
 *
 * ## Import
 *
 * Gitlab project milestone can be imported with a key composed of `<project>:<milestone_id>`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectMilestone:ProjectMilestone example "12345:11"
 * ```
 */
export class ProjectMilestone extends pulumi.CustomResource {
    /**
     * Get an existing ProjectMilestone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectMilestoneState, opts?: pulumi.CustomResourceOptions): ProjectMilestone {
        return new ProjectMilestone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectMilestone:ProjectMilestone';

    /**
     * Returns true if the given object is an instance of ProjectMilestone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectMilestone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectMilestone.__pulumiType;
    }

    /**
     * The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the milestone.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    public readonly dueDate!: pulumi.Output<string | undefined>;
    /**
     * Bool, true if milestone expired.
     */
    public /*out*/ readonly expired!: pulumi.Output<boolean>;
    /**
     * The ID of the project's milestone.
     */
    public /*out*/ readonly iid!: pulumi.Output<number>;
    /**
     * The instance-wide ID of the project’s milestone.
     */
    public /*out*/ readonly milestoneId!: pulumi.Output<number>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The project ID of milestone.
     */
    public /*out*/ readonly projectId!: pulumi.Output<number>;
    /**
     * The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    public readonly startDate!: pulumi.Output<string | undefined>;
    /**
     * The state of the milestone. Valid values are: `active`, `closed`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * The title of a milestone.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The web URL of the milestone.
     */
    public /*out*/ readonly webUrl!: pulumi.Output<string>;

    /**
     * Create a ProjectMilestone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectMilestoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectMilestoneArgs | ProjectMilestoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectMilestoneState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dueDate"] = state ? state.dueDate : undefined;
            resourceInputs["expired"] = state ? state.expired : undefined;
            resourceInputs["iid"] = state ? state.iid : undefined;
            resourceInputs["milestoneId"] = state ? state.milestoneId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["webUrl"] = state ? state.webUrl : undefined;
        } else {
            const args = argsOrState as ProjectMilestoneArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dueDate"] = args ? args.dueDate : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["startDate"] = args ? args.startDate : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["expired"] = undefined /*out*/;
            resourceInputs["iid"] = undefined /*out*/;
            resourceInputs["milestoneId"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["webUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectMilestone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectMilestone resources.
 */
export interface ProjectMilestoneState {
    /**
     * The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the milestone.
     */
    description?: pulumi.Input<string>;
    /**
     * The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    dueDate?: pulumi.Input<string>;
    /**
     * Bool, true if milestone expired.
     */
    expired?: pulumi.Input<boolean>;
    /**
     * The ID of the project's milestone.
     */
    iid?: pulumi.Input<number>;
    /**
     * The instance-wide ID of the project’s milestone.
     */
    milestoneId?: pulumi.Input<number>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project?: pulumi.Input<string>;
    /**
     * The project ID of milestone.
     */
    projectId?: pulumi.Input<number>;
    /**
     * The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    startDate?: pulumi.Input<string>;
    /**
     * The state of the milestone. Valid values are: `active`, `closed`.
     */
    state?: pulumi.Input<string>;
    /**
     * The title of a milestone.
     */
    title?: pulumi.Input<string>;
    /**
     * The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The web URL of the milestone.
     */
    webUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectMilestone resource.
 */
export interface ProjectMilestoneArgs {
    /**
     * The description of the milestone.
     */
    description?: pulumi.Input<string>;
    /**
     * The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    dueDate?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
    /**
     * The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     */
    startDate?: pulumi.Input<string>;
    /**
     * The state of the milestone. Valid values are: `active`, `closed`.
     */
    state?: pulumi.Input<string>;
    /**
     * The title of a milestone.
     */
    title: pulumi.Input<string>;
}
