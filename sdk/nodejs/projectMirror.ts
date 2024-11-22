// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectMirror` resource allows to manage the lifecycle of a project mirror.
 *
 * This is for *pushing* changes to a remote repository. *Pull Mirroring* can be configured using a combination of the
 * import_url, mirror, and mirrorTriggerBuilds properties on the gitlab.Project resource.
 *
 * > **Warning** By default, the provider sets the `keepDivergentRefs` argument to `True`.
 *    If you manually set `keepDivergentRefs` to `False`, GitLab mirroring removes branches in the target that aren't in the source.
 *    This action can result in unexpected branch deletions.
 *
 * > **Destroy Behavior** GitLab 14.10 introduced an API endpoint to delete a project mirror.
 *    Therefore, for GitLab 14.10 and newer the project mirror will be destroyed when the resource is destroyed.
 *    For older versions, the mirror will be disabled and the resource will be destroyed.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/remote_mirrors.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.ProjectMirror("foo", {
 *     project: "1",
 *     url: "https://username:password@github.com/org/repository.git",
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_mirror`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_mirror.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Import using the CLI is supported using the following syntax:
 *
 * GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectMirror:ProjectMirror foo "12345:1337"
 * ```
 */
export class ProjectMirror extends pulumi.CustomResource {
    /**
     * Get an existing ProjectMirror resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectMirrorState, opts?: pulumi.CustomResourceOptions): ProjectMirror {
        return new ProjectMirror(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectMirror:ProjectMirror';

    /**
     * Returns true if the given object is an instance of ProjectMirror.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectMirror {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectMirror.__pulumiType;
    }

    /**
     * Determines if the mirror is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Determines if divergent refs are skipped.
     */
    public readonly keepDivergentRefs!: pulumi.Output<boolean | undefined>;
    /**
     * Mirror ID.
     */
    public /*out*/ readonly mirrorId!: pulumi.Output<number>;
    /**
     * Determines if only protected branches are mirrored.
     */
    public readonly onlyProtectedBranches!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URL of the remote repository to be mirrored.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ProjectMirror resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectMirrorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectMirrorArgs | ProjectMirrorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectMirrorState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["keepDivergentRefs"] = state ? state.keepDivergentRefs : undefined;
            resourceInputs["mirrorId"] = state ? state.mirrorId : undefined;
            resourceInputs["onlyProtectedBranches"] = state ? state.onlyProtectedBranches : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ProjectMirrorArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["keepDivergentRefs"] = args ? args.keepDivergentRefs : undefined;
            resourceInputs["onlyProtectedBranches"] = args ? args.onlyProtectedBranches : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["mirrorId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["url"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProjectMirror.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectMirror resources.
 */
export interface ProjectMirrorState {
    /**
     * Determines if the mirror is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Determines if divergent refs are skipped.
     */
    keepDivergentRefs?: pulumi.Input<boolean>;
    /**
     * Mirror ID.
     */
    mirrorId?: pulumi.Input<number>;
    /**
     * Determines if only protected branches are mirrored.
     */
    onlyProtectedBranches?: pulumi.Input<boolean>;
    /**
     * The id of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * The URL of the remote repository to be mirrored.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectMirror resource.
 */
export interface ProjectMirrorArgs {
    /**
     * Determines if the mirror is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Determines if divergent refs are skipped.
     */
    keepDivergentRefs?: pulumi.Input<boolean>;
    /**
     * Determines if only protected branches are mirrored.
     */
    onlyProtectedBranches?: pulumi.Input<boolean>;
    /**
     * The id of the project.
     */
    project: pulumi.Input<string>;
    /**
     * The URL of the remote repository to be mirrored.
     */
    url: pulumi.Input<string>;
}
