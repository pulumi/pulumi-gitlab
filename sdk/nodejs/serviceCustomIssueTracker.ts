// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ServiceCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.
 *
 * > This resource is deprecated. use `gitlab.IntegrationCustomIssueTracker`instead!
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#custom-issue-tracker)
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
 * const tracker = new gitlab.ServiceCustomIssueTracker("tracker", {
 *     project: awesomeProject.id,
 *     projectUrl: "https://customtracker.com/issues",
 *     issuesUrl: "https://customtracker.com/TEST-:id",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_custom_issue_tracker`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_service_custom_issue_tracker.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * You can import a gitlab_service_custom_issue_tracker state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker tracker 1
 * ```
 */
export class ServiceCustomIssueTracker extends pulumi.CustomResource {
    /**
     * Get an existing ServiceCustomIssueTracker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceCustomIssueTrackerState, opts?: pulumi.CustomResourceOptions): ServiceCustomIssueTracker {
        return new ServiceCustomIssueTracker(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker';

    /**
     * Returns true if the given object is an instance of ServiceCustomIssueTracker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceCustomIssueTracker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceCustomIssueTracker.__pulumiType;
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
     * The URL to view an issue in the external issue tracker. Must contain :id.
     */
    public readonly issuesUrl!: pulumi.Output<string>;
    /**
     * The ID or full path of the project for the custom issue tracker.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URL to the project in the external issue tracker.
     */
    public readonly projectUrl!: pulumi.Output<string>;
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    public /*out*/ readonly slug!: pulumi.Output<string>;
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ServiceCustomIssueTracker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceCustomIssueTrackerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceCustomIssueTrackerArgs | ServiceCustomIssueTrackerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceCustomIssueTrackerState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["issuesUrl"] = state ? state.issuesUrl : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectUrl"] = state ? state.projectUrl : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ServiceCustomIssueTrackerArgs | undefined;
            if ((!args || args.issuesUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuesUrl'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.projectUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectUrl'");
            }
            resourceInputs["issuesUrl"] = args ? args.issuesUrl : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["projectUrl"] = args ? args.projectUrl : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceCustomIssueTracker.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceCustomIssueTracker resources.
 */
export interface ServiceCustomIssueTrackerState {
    /**
     * Whether the integration is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The URL to view an issue in the external issue tracker. Must contain :id.
     */
    issuesUrl?: pulumi.Input<string>;
    /**
     * The ID or full path of the project for the custom issue tracker.
     */
    project?: pulumi.Input<string>;
    /**
     * The URL to the project in the external issue tracker.
     */
    projectUrl?: pulumi.Input<string>;
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     */
    slug?: pulumi.Input<string>;
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceCustomIssueTracker resource.
 */
export interface ServiceCustomIssueTrackerArgs {
    /**
     * The URL to view an issue in the external issue tracker. Must contain :id.
     */
    issuesUrl: pulumi.Input<string>;
    /**
     * The ID or full path of the project for the custom issue tracker.
     */
    project: pulumi.Input<string>;
    /**
     * The URL to the project in the external issue tracker.
     */
    projectUrl: pulumi.Input<string>;
}
