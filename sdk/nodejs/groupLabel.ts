// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage labels for your GitLab groups.
 * For further information on labels, consult the [gitlab
 * documentation](https://docs.gitlab.com/ee/user/project/labels.html#group-labels).
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const fixme = new gitlab.GroupLabel("fixme", {
 *     color: "#ffcc00",
 *     description: "issue with failing tests",
 *     group: "example",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/group_label.html.markdown.
 */
export class GroupLabel extends pulumi.CustomResource {
    /**
     * Get an existing GroupLabel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupLabelState, opts?: pulumi.CustomResourceOptions): GroupLabel {
        return new GroupLabel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupLabel:GroupLabel';

    /**
     * Returns true if the given object is an instance of GroupLabel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupLabel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupLabel.__pulumiType;
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
     * The name or id of the group to add the label to.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The name of the label.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a GroupLabel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupLabelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupLabelArgs | GroupLabelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupLabelState | undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as GroupLabelArgs | undefined;
            if (!args || args.color === undefined) {
                throw new Error("Missing required property 'color'");
            }
            if (!args || args.group === undefined) {
                throw new Error("Missing required property 'group'");
            }
            inputs["color"] = args ? args.color : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupLabel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupLabel resources.
 */
export interface GroupLabelState {
    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    readonly color?: pulumi.Input<string>;
    /**
     * The description of the label.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name or id of the group to add the label to.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupLabel resource.
 */
export interface GroupLabelArgs {
    /**
     * The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     */
    readonly color: pulumi.Input<string>;
    /**
     * The description of the label.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name or id of the group to add the label to.
     */
    readonly group: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    readonly name?: pulumi.Input<string>;
}