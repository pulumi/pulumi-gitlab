// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage labels for your GitLab projects.
 * For further information on labels, consult the [gitlab
 * documentation](https://docs.gitlab.com/ee/user/project/labels.htm).
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const fixme = new gitlab.Label("fixme", {
 *     color: "#ffcc00",
 *     description: "issue with failing tests",
 *     project: "example",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/label.html.markdown.
 */
export class Label extends pulumi.CustomResource {
    /**
     * Get an existing Label resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LabelState, opts?: pulumi.CustomResourceOptions): Label {
        return new Label(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/label:Label';

    /**
     * Returns true if the given object is an instance of Label.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Label {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Label.__pulumiType;
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
     * The name of the label.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the label to.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Label resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LabelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LabelArgs | LabelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LabelState | undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as LabelArgs | undefined;
            if (!args || args.color === undefined) {
                throw new Error("Missing required property 'color'");
            }
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            inputs["color"] = args ? args.color : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Label.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Label resources.
 */
export interface LabelState {
    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    readonly color?: pulumi.Input<string>;
    /**
     * The description of the label.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the label to.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Label resource.
 */
export interface LabelArgs {
    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    readonly color: pulumi.Input<string>;
    /**
     * The description of the label.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the label to.
     */
    readonly project: pulumi.Input<string>;
}
