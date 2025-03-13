// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectRunnerEnablement` resource allows to enable a runner in a project.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/runners/#enable-a-runner-in-project)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.ProjectRunnerEnablement("foo", {
 *     project: "5",
 *     runnerId: 7,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_runner_enablement`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_runner_enablement.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * GitLab project runners can be imported using an id made up of `project:runner_id`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement foo 5:7
 * ```
 */
export class ProjectRunnerEnablement extends pulumi.CustomResource {
    /**
     * Get an existing ProjectRunnerEnablement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectRunnerEnablementState, opts?: pulumi.CustomResourceOptions): ProjectRunnerEnablement {
        return new ProjectRunnerEnablement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement';

    /**
     * Returns true if the given object is an instance of ProjectRunnerEnablement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectRunnerEnablement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectRunnerEnablement.__pulumiType;
    }

    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The ID of a runner to enable for the project.
     */
    public readonly runnerId!: pulumi.Output<number>;

    /**
     * Create a ProjectRunnerEnablement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectRunnerEnablementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectRunnerEnablementArgs | ProjectRunnerEnablementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectRunnerEnablementState | undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["runnerId"] = state ? state.runnerId : undefined;
        } else {
            const args = argsOrState as ProjectRunnerEnablementArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.runnerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runnerId'");
            }
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["runnerId"] = args ? args.runnerId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectRunnerEnablement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectRunnerEnablement resources.
 */
export interface ProjectRunnerEnablementState {
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project?: pulumi.Input<string>;
    /**
     * The ID of a runner to enable for the project.
     */
    runnerId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectRunnerEnablement resource.
 */
export interface ProjectRunnerEnablementArgs {
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
    /**
     * The ID of a runner to enable for the project.
     */
    runnerId: pulumi.Input<number>;
}
