// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationJenkins` resource allows to manage the lifecycle of a project integration with Jenkins.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#jenkins)
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
 * const jenkins = new gitlab.IntegrationJenkins("jenkins", {
 *     project: awesomeProject.id,
 *     jenkinsUrl: "http://jenkins.example.com",
 *     projectName: "my_project_name",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationJenkins:IntegrationJenkins You can import a gitlab_integration_jenkins state using `<resource> <project_id>`:
 * ```
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationJenkins:IntegrationJenkins jenkins 1
 * ```
 */
export class IntegrationJenkins extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationJenkins resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationJenkinsState, opts?: pulumi.CustomResourceOptions): IntegrationJenkins {
        return new IntegrationJenkins(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationJenkins:IntegrationJenkins';

    /**
     * Returns true if the given object is an instance of IntegrationJenkins.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationJenkins {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationJenkins.__pulumiType;
    }

    /**
     * Whether the integration is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * Enable SSL verification. Defaults to `true` (enabled).
     */
    public readonly enableSslVerification!: pulumi.Output<boolean>;
    /**
     * Jenkins URL like `http://jenkins.example.com`
     */
    public readonly jenkinsUrl!: pulumi.Output<string>;
    /**
     * Enable notifications for merge request events.
     */
    public readonly mergeRequestEvents!: pulumi.Output<boolean>;
    /**
     * Password for authentication with the Jenkins server, if authentication is required by the server.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URL-friendly project name. Example: `myProjectName`.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * Enable notifications for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean>;
    /**
     * Username for authentication with the Jenkins server, if authentication is required by the server.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a IntegrationJenkins resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationJenkinsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationJenkinsArgs | IntegrationJenkinsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationJenkinsState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["enableSslVerification"] = state ? state.enableSslVerification : undefined;
            resourceInputs["jenkinsUrl"] = state ? state.jenkinsUrl : undefined;
            resourceInputs["mergeRequestEvents"] = state ? state.mergeRequestEvents : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["pushEvents"] = state ? state.pushEvents : undefined;
            resourceInputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as IntegrationJenkinsArgs | undefined;
            if ((!args || args.jenkinsUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jenkinsUrl'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.projectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectName'");
            }
            resourceInputs["enableSslVerification"] = args ? args.enableSslVerification : undefined;
            resourceInputs["jenkinsUrl"] = args ? args.jenkinsUrl : undefined;
            resourceInputs["mergeRequestEvents"] = args ? args.mergeRequestEvents : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["pushEvents"] = args ? args.pushEvents : undefined;
            resourceInputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["active"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IntegrationJenkins.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationJenkins resources.
 */
export interface IntegrationJenkinsState {
    /**
     * Whether the integration is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Enable SSL verification. Defaults to `true` (enabled).
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Jenkins URL like `http://jenkins.example.com`
     */
    jenkinsUrl?: pulumi.Input<string>;
    /**
     * Enable notifications for merge request events.
     */
    mergeRequestEvents?: pulumi.Input<boolean>;
    /**
     * Password for authentication with the Jenkins server, if authentication is required by the server.
     */
    password?: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * The URL-friendly project name. Example: `myProjectName`.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Username for authentication with the Jenkins server, if authentication is required by the server.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationJenkins resource.
 */
export interface IntegrationJenkinsArgs {
    /**
     * Enable SSL verification. Defaults to `true` (enabled).
     */
    enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Jenkins URL like `http://jenkins.example.com`
     */
    jenkinsUrl: pulumi.Input<string>;
    /**
     * Enable notifications for merge request events.
     */
    mergeRequestEvents?: pulumi.Input<boolean>;
    /**
     * Password for authentication with the Jenkins server, if authentication is required by the server.
     */
    password?: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * The URL-friendly project name. Example: `myProjectName`.
     */
    projectName: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    pushEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for tag push events.
     */
    tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Username for authentication with the Jenkins server, if authentication is required by the server.
     */
    username?: pulumi.Input<string>;
}
