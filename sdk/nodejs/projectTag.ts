// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectTag` resource allows to manage the lifecycle of a tag in a project.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/tags.html)
 *
 * ## Import
 *
 * Gitlab project tags can be imported with a key composed of `<project_id>:<tag_name>`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectTag:ProjectTag example "12345:develop"
 * ```
 *
 *  NOTEthe `ref` attribute won't be available for imported `gitlab_project_tag` resources.
 */
export class ProjectTag extends pulumi.CustomResource {
    /**
     * Get an existing ProjectTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectTagState, opts?: pulumi.CustomResourceOptions): ProjectTag {
        return new ProjectTag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectTag:ProjectTag';

    /**
     * Returns true if the given object is an instance of ProjectTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectTag.__pulumiType;
    }

    /**
     * The commit associated with the tag.
     */
    public /*out*/ readonly commits!: pulumi.Output<outputs.ProjectTagCommit[]>;
    /**
     * The message of the annotated tag.
     */
    public readonly message!: pulumi.Output<string | undefined>;
    /**
     * The name of a tag.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Bool, true if tag has tag protection.
     */
    public /*out*/ readonly protected!: pulumi.Output<boolean>;
    /**
     * Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
     */
    public readonly ref!: pulumi.Output<string>;
    /**
     * The release associated with the tag.
     */
    public /*out*/ readonly releases!: pulumi.Output<outputs.ProjectTagRelease[]>;
    /**
     * The unique id assigned to the commit by Gitlab.
     */
    public /*out*/ readonly target!: pulumi.Output<string>;

    /**
     * Create a ProjectTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectTagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectTagArgs | ProjectTagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectTagState | undefined;
            resourceInputs["commits"] = state ? state.commits : undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protected"] = state ? state.protected : undefined;
            resourceInputs["ref"] = state ? state.ref : undefined;
            resourceInputs["releases"] = state ? state.releases : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as ProjectTagArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.ref === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ref'");
            }
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["ref"] = args ? args.ref : undefined;
            resourceInputs["commits"] = undefined /*out*/;
            resourceInputs["protected"] = undefined /*out*/;
            resourceInputs["releases"] = undefined /*out*/;
            resourceInputs["target"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectTag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectTag resources.
 */
export interface ProjectTagState {
    /**
     * The commit associated with the tag.
     */
    commits?: pulumi.Input<pulumi.Input<inputs.ProjectTagCommit>[]>;
    /**
     * The message of the annotated tag.
     */
    message?: pulumi.Input<string>;
    /**
     * The name of a tag.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project?: pulumi.Input<string>;
    /**
     * Bool, true if tag has tag protection.
     */
    protected?: pulumi.Input<boolean>;
    /**
     * Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
     */
    ref?: pulumi.Input<string>;
    /**
     * The release associated with the tag.
     */
    releases?: pulumi.Input<pulumi.Input<inputs.ProjectTagRelease>[]>;
    /**
     * The unique id assigned to the commit by Gitlab.
     */
    target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectTag resource.
 */
export interface ProjectTagArgs {
    /**
     * The message of the annotated tag.
     */
    message?: pulumi.Input<string>;
    /**
     * The name of a tag.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     */
    project: pulumi.Input<string>;
    /**
     * Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
     */
    ref: pulumi.Input<string>;
}
