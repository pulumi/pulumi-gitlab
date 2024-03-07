// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectLabel` resource allows to manage the lifecycle of a project label.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/labels.html#project-labels)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const fixme = new gitlab.ProjectLabel("fixme", {
 *     project: "example",
 *     description: "issue with failing tests",
 *     color: "#ffcc00",
 * });
 * // Scoped label
 * const devopsCreate = new gitlab.ProjectLabel("devopsCreate", {
 *     project: gitlab_project.example.id,
 *     description: "issue for creating infrastructure resources",
 *     color: "#ffa500",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Gitlab Project labels can be imported using an id made up of `{project_id}:{group_label_id}`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectLabel:ProjectLabel example 12345:fixme
 * ```
 */
export class ProjectLabel extends pulumi.CustomResource {
    /**
     * Get an existing ProjectLabel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectLabelState, opts?: pulumi.CustomResourceOptions): ProjectLabel {
        return new ProjectLabel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectLabel:ProjectLabel';

    /**
     * Returns true if the given object is an instance of ProjectLabel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectLabel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectLabel.__pulumiType;
    }

    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    public readonly color!: pulumi.Output<string>;
    /**
     * The description of the label.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The id of the project label.
     */
    public /*out*/ readonly labelId!: pulumi.Output<number>;
    /**
     * The name of the label.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the label to.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ProjectLabel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectLabelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectLabelArgs | ProjectLabelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectLabelState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["labelId"] = state ? state.labelId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectLabelArgs | undefined;
            if ((!args || args.color === undefined) && !opts.urn) {
                throw new Error("Missing required property 'color'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["labelId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectLabel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectLabel resources.
 */
export interface ProjectLabelState {
    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    color?: pulumi.Input<string>;
    /**
     * The description of the label.
     */
    description?: pulumi.Input<string>;
    /**
     * The id of the project label.
     */
    labelId?: pulumi.Input<number>;
    /**
     * The name of the label.
     */
    name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the label to.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectLabel resource.
 */
export interface ProjectLabelArgs {
    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    color: pulumi.Input<string>;
    /**
     * The description of the label.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the label to.
     */
    project: pulumi.Input<string>;
}
