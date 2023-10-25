// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker)
 *
 * ## Import
 *
 * You can import a gitlab_integration_custom_issue_tracker state using the project ID, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker tracker 1
 * ```
 */
export class IntegrationCustomIssueTracker extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationCustomIssueTracker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationCustomIssueTrackerState, opts?: pulumi.CustomResourceOptions): IntegrationCustomIssueTracker {
        return new IntegrationCustomIssueTracker(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker';

    /**
     * Returns true if the given object is an instance of IntegrationCustomIssueTracker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationCustomIssueTracker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationCustomIssueTracker.__pulumiType;
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
     * Create a IntegrationCustomIssueTracker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationCustomIssueTrackerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationCustomIssueTrackerArgs | IntegrationCustomIssueTrackerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationCustomIssueTrackerState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["issuesUrl"] = state ? state.issuesUrl : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectUrl"] = state ? state.projectUrl : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IntegrationCustomIssueTrackerArgs | undefined;
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
        super(IntegrationCustomIssueTracker.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationCustomIssueTracker resources.
 */
export interface IntegrationCustomIssueTrackerState {
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
 * The set of arguments for constructing a IntegrationCustomIssueTracker resource.
 */
export interface IntegrationCustomIssueTrackerArgs {
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
