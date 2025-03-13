// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.UserRunner` resource allows creating a GitLab runner using the new [GitLab Runner Registration Flow](https://docs.gitlab.com/ci/runners/new_creation_workflow/).
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/users/#create-a-runner)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * // Create a project runner
 * const projectRunner = new gitlab.UserRunner("project_runner", {
 *     runnerType: "project_type",
 *     projectId: 123456,
 *     description: "A runner created using a user access token instead of a registration token",
 *     tagLists: [
 *         "a-tag",
 *         "other-tag",
 *     ],
 *     untagged: true,
 * });
 * // Create a group runner
 * const groupRunner = new gitlab.UserRunner("group_runner", {
 *     runnerType: "group_type",
 *     groupId: 123456,
 * });
 * // Create a instance runner
 * const instanceRunner = new gitlab.UserRunner("instance_runner", {runnerType: "instance_type"});
 * const configToml = pulumi.interpolate`concurrent = 1
 * check_interval = 0
 *
 * [session_server]
 *   session_timeout = 1800
 *
 * [[runners]]
 *   name = "my_gitlab_runner"
 *   url = "https://example.gitlab.com"
 *   token = "${groupRunner.token}"
 *   executor = "docker"
 *
 *   [runners.custom_build_dir]
 *   [runners.cache]
 *     [runners.cache.s3]
 *     [runners.cache.gcs]
 *     [runners.cache.azure]
 *   [runners.docker]
 *     tls_verify = false
 *     image = "ubuntu"
 *     privileged = true
 *     disable_entrypoint_overwrite = false
 *     oom_kill_disable = false
 *     disable_cache = false
 *     volumes = ["/cache", "/certs/client"]
 *     shm_size = 0
 * `;
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_user_runner`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_user_runner.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * You can import a gitlab runner using its ID
 *
 * Note: Importing a runner will not provide access to the `token` attribute
 *
 * ```sh
 * $ pulumi import gitlab:index/userRunner:UserRunner example 12345
 * ```
 */
export class UserRunner extends pulumi.CustomResource {
    /**
     * Get an existing UserRunner resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserRunnerState, opts?: pulumi.CustomResourceOptions): UserRunner {
        return new UserRunner(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/userRunner:UserRunner';

    /**
     * Returns true if the given object is an instance of UserRunner.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserRunner {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserRunner.__pulumiType;
    }

    /**
     * The access level of the runner. Valid values are: `notProtected`, `refProtected`.
     */
    public readonly accessLevel!: pulumi.Output<string>;
    /**
     * Description of the runner.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ID of the group that the runner is created in. Required if runner*type is group*type.
     */
    public readonly groupId!: pulumi.Output<number | undefined>;
    /**
     * Specifies if the runner should be locked for the current project.
     */
    public readonly locked!: pulumi.Output<boolean>;
    /**
     * Free-form maintenance notes for the runner (1024 characters)
     */
    public readonly maintenanceNote!: pulumi.Output<string>;
    /**
     * Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
     */
    public readonly maximumTimeout!: pulumi.Output<number>;
    /**
     * Specifies if the runner should ignore new jobs.
     */
    public readonly paused!: pulumi.Output<boolean>;
    /**
     * The ID of the project that the runner is created in. Required if runner*type is project*type.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
     */
    public readonly runnerType!: pulumi.Output<string>;
    /**
     * A list of runner tags.
     */
    public readonly tagLists!: pulumi.Output<string[]>;
    /**
     * The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * Specifies if the runner should handle untagged jobs.
     */
    public readonly untagged!: pulumi.Output<boolean>;

    /**
     * Create a UserRunner resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserRunnerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserRunnerArgs | UserRunnerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserRunnerState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["locked"] = state ? state.locked : undefined;
            resourceInputs["maintenanceNote"] = state ? state.maintenanceNote : undefined;
            resourceInputs["maximumTimeout"] = state ? state.maximumTimeout : undefined;
            resourceInputs["paused"] = state ? state.paused : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["runnerType"] = state ? state.runnerType : undefined;
            resourceInputs["tagLists"] = state ? state.tagLists : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["untagged"] = state ? state.untagged : undefined;
        } else {
            const args = argsOrState as UserRunnerArgs | undefined;
            if ((!args || args.runnerType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runnerType'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["locked"] = args ? args.locked : undefined;
            resourceInputs["maintenanceNote"] = args ? args.maintenanceNote : undefined;
            resourceInputs["maximumTimeout"] = args ? args.maximumTimeout : undefined;
            resourceInputs["paused"] = args ? args.paused : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["runnerType"] = args ? args.runnerType : undefined;
            resourceInputs["tagLists"] = args ? args.tagLists : undefined;
            resourceInputs["untagged"] = args ? args.untagged : undefined;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(UserRunner.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserRunner resources.
 */
export interface UserRunnerState {
    /**
     * The access level of the runner. Valid values are: `notProtected`, `refProtected`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Description of the runner.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the group that the runner is created in. Required if runner*type is group*type.
     */
    groupId?: pulumi.Input<number>;
    /**
     * Specifies if the runner should be locked for the current project.
     */
    locked?: pulumi.Input<boolean>;
    /**
     * Free-form maintenance notes for the runner (1024 characters)
     */
    maintenanceNote?: pulumi.Input<string>;
    /**
     * Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
     */
    maximumTimeout?: pulumi.Input<number>;
    /**
     * Specifies if the runner should ignore new jobs.
     */
    paused?: pulumi.Input<boolean>;
    /**
     * The ID of the project that the runner is created in. Required if runner*type is project*type.
     */
    projectId?: pulumi.Input<number>;
    /**
     * The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
     */
    runnerType?: pulumi.Input<string>;
    /**
     * A list of runner tags.
     */
    tagLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The authentication token to use when setting up a new runner with this configuration. This value cannot be imported.
     */
    token?: pulumi.Input<string>;
    /**
     * Specifies if the runner should handle untagged jobs.
     */
    untagged?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a UserRunner resource.
 */
export interface UserRunnerArgs {
    /**
     * The access level of the runner. Valid values are: `notProtected`, `refProtected`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Description of the runner.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the group that the runner is created in. Required if runner*type is group*type.
     */
    groupId?: pulumi.Input<number>;
    /**
     * Specifies if the runner should be locked for the current project.
     */
    locked?: pulumi.Input<boolean>;
    /**
     * Free-form maintenance notes for the runner (1024 characters)
     */
    maintenanceNote?: pulumi.Input<string>;
    /**
     * Maximum timeout that limits the amount of time (in seconds) that runners can run jobs. Must be at least 600 (10 minutes).
     */
    maximumTimeout?: pulumi.Input<number>;
    /**
     * Specifies if the runner should ignore new jobs.
     */
    paused?: pulumi.Input<boolean>;
    /**
     * The ID of the project that the runner is created in. Required if runner*type is project*type.
     */
    projectId?: pulumi.Input<number>;
    /**
     * The scope of the runner. Valid values are: `instanceType`, `groupType`, `projectType`.
     */
    runnerType: pulumi.Input<string>;
    /**
     * A list of runner tags.
     */
    tagLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies if the runner should handle untagged jobs.
     */
    untagged?: pulumi.Input<boolean>;
}
