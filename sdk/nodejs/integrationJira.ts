// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationJira` resource allows to manage the lifecycle of a project integration with Jira.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/services.html#jira)
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
 * const jira = new gitlab.IntegrationJira("jira", {
 *     project: awesomeProject.id,
 *     url: "https://jira.example.com",
 *     username: "user",
 *     password: "mypass",
 * });
 * ```
 *
 * ## Import
 *
 * You can import a gitlab_integration_jira state using the project ID, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/integrationJira:IntegrationJira jira 1
 * ```
 */
export class IntegrationJira extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationJira resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationJiraState, opts?: pulumi.CustomResourceOptions): IntegrationJira {
        return new IntegrationJira(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationJira:IntegrationJira';

    /**
     * Returns true if the given object is an instance of IntegrationJira.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationJira {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationJira.__pulumiType;
    }

    /**
     * Whether the integration is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
     */
    public readonly apiUrl!: pulumi.Output<string>;
    /**
     * Enable comments inside Jira issues on each GitLab event (commit / merge request)
     */
    public readonly commentOnEventEnabled!: pulumi.Output<boolean>;
    /**
     * Enable notifications for commit events
     */
    public readonly commitEvents!: pulumi.Output<boolean>;
    /**
     * Create time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Enable notifications for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean>;
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     */
    public readonly jiraIssueTransitionId!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for job events.
     */
    public readonly jobEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for merge request events
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for note events.
     */
    public readonly noteEvents!: pulumi.Output<boolean>;
    /**
     * The password of the user created to be used with GitLab/JIRA.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Enable notifications for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for tagPush events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean>;
    /**
     * Title.
     */
    public /*out*/ readonly title!: pulumi.Output<string>;
    /**
     * Update time.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The username of the user created to be used with GitLab/JIRA.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a IntegrationJira resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationJiraArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationJiraArgs | IntegrationJiraState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationJiraState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["apiUrl"] = state ? state.apiUrl : undefined;
            resourceInputs["commentOnEventEnabled"] = state ? state.commentOnEventEnabled : undefined;
            resourceInputs["commitEvents"] = state ? state.commitEvents : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            resourceInputs["jiraIssueTransitionId"] = state ? state.jiraIssueTransitionId : undefined;
            resourceInputs["jobEvents"] = state ? state.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = state ? state.noteEvents : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as IntegrationJiraArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["apiUrl"] = args ? args.apiUrl : undefined;
            resourceInputs["commentOnEventEnabled"] = args ? args.commentOnEventEnabled : undefined;
            resourceInputs["commitEvents"] = args ? args.commitEvents : undefined;
            resourceInputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            resourceInputs["jiraIssueTransitionId"] = args ? args.jiraIssueTransitionId : undefined;
            resourceInputs["jobEvents"] = args ? args.jobEvents : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["noteEvents"] = args ? args.noteEvents : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IntegrationJira.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationJira resources.
 */
export interface IntegrationJiraState {
    /**
     * Whether the integration is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
     */
    apiUrl?: pulumi.Input<string>;
    /**
     * Enable comments inside Jira issues on each GitLab event (commit / merge request)
     */
    commentOnEventEnabled?: pulumi.Input<boolean>;
    /**
     * Enable notifications for commit events
     */
    commitEvents?: pulumi.Input<boolean>;
    /**
     * Create time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Enable notifications for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     */
    jiraIssueTransitionId?: pulumi.Input<string>;
    /**
     * Enable notifications for job events.
     */
    jobEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merge request events
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * The password of the user created to be used with GitLab/JIRA.
     */
    password?: pulumi.Input<string>;
    /**
     * Enable notifications for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for tagPush events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Title.
     */
    title?: pulumi.Input<string>;
    /**
     * Update time.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
     */
    url?: pulumi.Input<string>;
    /**
     * The username of the user created to be used with GitLab/JIRA.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationJira resource.
 */
export interface IntegrationJiraArgs {
    /**
     * The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
     */
    apiUrl?: pulumi.Input<string>;
    /**
     * Enable comments inside Jira issues on each GitLab event (commit / merge request)
     */
    commentOnEventEnabled?: pulumi.Input<boolean>;
    /**
     * Enable notifications for commit events
     */
    commitEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for issues events.
     */
    issuesEvents?: pulumi.Input<boolean>;
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. *Note**: importing this field is only supported since GitLab 15.2.
     */
    jiraIssueTransitionId?: pulumi.Input<string>;
    /**
     * Enable notifications for job events.
     */
    jobEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for merge request events
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for note events.
     */
    noteEvents?: pulumi.Input<boolean>;
    /**
     * The password of the user created to be used with GitLab/JIRA.
     */
    password: pulumi.Input<string>;
    /**
     * Enable notifications for pipeline events.
     */
    pipelineEvents?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * The short identifier for your JIRA project, all uppercase, e.g., PROJ.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for tagPush events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
     */
    url: pulumi.Input<string>;
    /**
     * The username of the user created to be used with GitLab/JIRA.
     */
    username: pulumi.Input<string>;
}
