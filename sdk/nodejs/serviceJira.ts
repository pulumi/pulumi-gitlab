// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ServiceJira` resource allows to manage the lifecycle of a project integration with Jira.
 *
 * > This resource is deprecated. use `gitlab.IntegrationJira`instead!
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#jira-issues)
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
 * const jira = new gitlab.ServiceJira("jira", {
 *     project: awesomeProject.id,
 *     url: "https://jira.example.com",
 *     username: "user",
 *     password: "mypass",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_jira`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_service_jira.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * You can import a gitlab_service_jira state using the project ID, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/serviceJira:ServiceJira jira 1
 * ```
 */
export class ServiceJira extends pulumi.CustomResource {
    /**
     * Get an existing ServiceJira resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceJiraState, opts?: pulumi.CustomResourceOptions): ServiceJira {
        return new ServiceJira(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/serviceJira:ServiceJira';

    /**
     * Returns true if the given object is an instance of ServiceJira.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceJira {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceJira.__pulumiType;
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
     * Enable viewing Jira issues in GitLab.
     */
    public readonly issuesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
     */
    public readonly jiraAuthType!: pulumi.Output<number | undefined>;
    /**
     * Prefix to match Jira issue keys.
     */
    public readonly jiraIssuePrefix!: pulumi.Output<string | undefined>;
    /**
     * Regular expression to match Jira issue keys.
     */
    public readonly jiraIssueRegex!: pulumi.Output<string | undefined>;
    public readonly jiraIssueTransitionAutomatic!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
     */
    public readonly jiraIssueTransitionId!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for merge request events
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean>;
    /**
     * The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The short identifier for your JIRA project. Must be all uppercase. For example, `PROJ`.
     *
     * @deprecated `projectKey` is deprecated. Use `projectKeys` instead.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
     */
    public readonly projectKeys!: pulumi.Output<string[] | undefined>;
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
     * Indicates whether or not to inherit default settings. Defaults to false.
     */
    public readonly useInheritedSettings!: pulumi.Output<boolean | undefined>;
    /**
     * The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceJira resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceJiraArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceJiraArgs | ServiceJiraState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceJiraState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["apiUrl"] = state ? state.apiUrl : undefined;
            resourceInputs["commentOnEventEnabled"] = state ? state.commentOnEventEnabled : undefined;
            resourceInputs["commitEvents"] = state ? state.commitEvents : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["issuesEnabled"] = state ? state.issuesEnabled : undefined;
            resourceInputs["jiraAuthType"] = state ? state.jiraAuthType : undefined;
            resourceInputs["jiraIssuePrefix"] = state ? state.jiraIssuePrefix : undefined;
            resourceInputs["jiraIssueRegex"] = state ? state.jiraIssueRegex : undefined;
            resourceInputs["jiraIssueTransitionAutomatic"] = state ? state.jiraIssueTransitionAutomatic : undefined;
            resourceInputs["jiraIssueTransitionId"] = state ? state.jiraIssueTransitionId : undefined;
            resourceInputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["projectKeys"] = state ? state.projectKeys : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["useInheritedSettings"] = state ? state.useInheritedSettings : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceJiraArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["apiUrl"] = args ? args.apiUrl : undefined;
            resourceInputs["commentOnEventEnabled"] = args ? args.commentOnEventEnabled : undefined;
            resourceInputs["commitEvents"] = args ? args.commitEvents : undefined;
            resourceInputs["issuesEnabled"] = args ? args.issuesEnabled : undefined;
            resourceInputs["jiraAuthType"] = args ? args.jiraAuthType : undefined;
            resourceInputs["jiraIssuePrefix"] = args ? args.jiraIssuePrefix : undefined;
            resourceInputs["jiraIssueRegex"] = args ? args.jiraIssueRegex : undefined;
            resourceInputs["jiraIssueTransitionAutomatic"] = args ? args.jiraIssueTransitionAutomatic : undefined;
            resourceInputs["jiraIssueTransitionId"] = args ? args.jiraIssueTransitionId : undefined;
            resourceInputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["projectKeys"] = args ? args.projectKeys : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["useInheritedSettings"] = args ? args.useInheritedSettings : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceJira.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceJira resources.
 */
export interface ServiceJiraState {
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
     * Enable viewing Jira issues in GitLab.
     */
    issuesEnabled?: pulumi.Input<boolean>;
    /**
     * The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
     */
    jiraAuthType?: pulumi.Input<number>;
    /**
     * Prefix to match Jira issue keys.
     */
    jiraIssuePrefix?: pulumi.Input<string>;
    /**
     * Regular expression to match Jira issue keys.
     */
    jiraIssueRegex?: pulumi.Input<string>;
    jiraIssueTransitionAutomatic?: pulumi.Input<boolean>;
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
     */
    jiraIssueTransitionId?: pulumi.Input<string>;
    /**
     * Enable notifications for merge request events
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
     */
    password?: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * The short identifier for your JIRA project. Must be all uppercase. For example, `PROJ`.
     *
     * @deprecated `projectKey` is deprecated. Use `projectKeys` instead.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
     */
    projectKeys?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Indicates whether or not to inherit default settings. Defaults to false.
     */
    useInheritedSettings?: pulumi.Input<boolean>;
    /**
     * The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceJira resource.
 */
export interface ServiceJiraArgs {
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
     * Enable viewing Jira issues in GitLab.
     */
    issuesEnabled?: pulumi.Input<boolean>;
    /**
     * The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
     */
    jiraAuthType?: pulumi.Input<number>;
    /**
     * Prefix to match Jira issue keys.
     */
    jiraIssuePrefix?: pulumi.Input<string>;
    /**
     * Regular expression to match Jira issue keys.
     */
    jiraIssueRegex?: pulumi.Input<string>;
    jiraIssueTransitionAutomatic?: pulumi.Input<boolean>;
    /**
     * The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
     */
    jiraIssueTransitionId?: pulumi.Input<string>;
    /**
     * Enable notifications for merge request events
     */
    mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
     */
    password: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * The short identifier for your JIRA project. Must be all uppercase. For example, `PROJ`.
     *
     * @deprecated `projectKey` is deprecated. Use `projectKeys` instead.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
     */
    projectKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
     */
    url: pulumi.Input<string>;
    /**
     * Indicates whether or not to inherit default settings. Defaults to false.
     */
    useInheritedSettings?: pulumi.Input<boolean>;
    /**
     * The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
     */
    username?: pulumi.Input<string>;
}
