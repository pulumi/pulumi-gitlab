// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.IntegrationGithub` resource allows to manage the lifecycle of a project integration with GitHub.
 *
 * > This resource requires a GitLab Enterprise instance.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#github)
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
 * const github = new gitlab.IntegrationGithub("github", {
 *     project: awesomeProject.id,
 *     token: "REDACTED",
 *     repositoryUrl: "https://github.com/gitlabhq/terraform-provider-gitlab",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationGithub:IntegrationGithub You can import a gitlab_integration_github state using `<resource> <project_id>`:
 * ```
 *
 * ```sh
 * $ pulumi import gitlab:index/integrationGithub:IntegrationGithub github 1
 * ```
 */
export class IntegrationGithub extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationGithub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationGithubState, opts?: pulumi.CustomResourceOptions): IntegrationGithub {
        return new IntegrationGithub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/integrationGithub:IntegrationGithub';

    /**
     * Returns true if the given object is an instance of IntegrationGithub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationGithub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationGithub.__pulumiType;
    }

    /**
     * Whether the integration is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * Create time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     */
    public readonly repositoryUrl!: pulumi.Output<string>;
    /**
     * Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     */
    public readonly staticContext!: pulumi.Output<boolean | undefined>;
    /**
     * Title.
     */
    public /*out*/ readonly title!: pulumi.Output<string>;
    /**
     * A GitHub personal access token with at least `repo:status` scope.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * Update time.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IntegrationGithub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationGithubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationGithubArgs | IntegrationGithubState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationGithubState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["repositoryUrl"] = state ? state.repositoryUrl : undefined;
            resourceInputs["staticContext"] = state ? state.staticContext : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IntegrationGithubArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.repositoryUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryUrl'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["repositoryUrl"] = args ? args.repositoryUrl : undefined;
            resourceInputs["staticContext"] = args ? args.staticContext : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IntegrationGithub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationGithub resources.
 */
export interface IntegrationGithubState {
    /**
     * Whether the integration is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Create time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * ID of the project you want to activate integration on.
     */
    project?: pulumi.Input<string>;
    /**
     * The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     */
    repositoryUrl?: pulumi.Input<string>;
    /**
     * Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     */
    staticContext?: pulumi.Input<boolean>;
    /**
     * Title.
     */
    title?: pulumi.Input<string>;
    /**
     * A GitHub personal access token with at least `repo:status` scope.
     */
    token?: pulumi.Input<string>;
    /**
     * Update time.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationGithub resource.
 */
export interface IntegrationGithubArgs {
    /**
     * ID of the project you want to activate integration on.
     */
    project: pulumi.Input<string>;
    /**
     * The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     */
    repositoryUrl: pulumi.Input<string>;
    /**
     * Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     */
    staticContext?: pulumi.Input<boolean>;
    /**
     * A GitHub personal access token with at least `repo:status` scope.
     */
    token: pulumi.Input<string>;
}
