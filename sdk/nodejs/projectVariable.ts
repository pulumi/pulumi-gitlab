// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectVariable` resource allows to manage the lifecycle of a CI/CD variable for a project.
 *
 * > **Important:** If your GitLab version is older than 13.4, you may see nondeterministic behavior when updating or deleting gitlab.ProjectVariable resources with non-unique keys, for example if there is another variable with the same key and different environment scope. See [this GitLab issue](https://gitlab.com/gitlab-org/gitlab/-/issues/9912).
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.ProjectVariable("example", {
 *     key: "project_variable_key",
 *     project: "12345",
 *     "protected": false,
 *     value: "project_variable_value",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab project variables can be imported using an id made up of `project:key:environment_scope`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectVariable:ProjectVariable example '12345:project_variable_key:*'
 * ```
 */
export class ProjectVariable extends pulumi.CustomResource {
    /**
     * Get an existing ProjectVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectVariableState, opts?: pulumi.CustomResourceOptions): ProjectVariable {
        return new ProjectVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectVariable:ProjectVariable';

    /**
     * Returns true if the given object is an instance of ProjectVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectVariable.__pulumiType;
    }

    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    public readonly environmentScope!: pulumi.Output<string | undefined>;
    /**
     * The name of the variable.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     */
    public readonly masked!: pulumi.Output<boolean | undefined>;
    /**
     * The name or id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    public readonly protected!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     */
    public readonly raw!: pulumi.Output<boolean | undefined>;
    /**
     * The value of the variable.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
     */
    public readonly variableType!: pulumi.Output<string | undefined>;

    /**
     * Create a ProjectVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectVariableArgs | ProjectVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectVariableState | undefined;
            resourceInputs["environmentScope"] = state ? state.environmentScope : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["masked"] = state ? state.masked : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protected"] = state ? state.protected : undefined;
            resourceInputs["raw"] = state ? state.raw : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["variableType"] = state ? state.variableType : undefined;
        } else {
            const args = argsOrState as ProjectVariableArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["environmentScope"] = args ? args.environmentScope : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["masked"] = args ? args.masked : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protected"] = args ? args.protected : undefined;
            resourceInputs["raw"] = args ? args.raw : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["variableType"] = args ? args.variableType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectVariable resources.
 */
export interface ProjectVariableState {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * The name of the variable.
     */
    key?: pulumi.Input<string>;
    /**
     * If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     */
    masked?: pulumi.Input<boolean>;
    /**
     * The name or id of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    protected?: pulumi.Input<boolean>;
    /**
     * Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     */
    raw?: pulumi.Input<boolean>;
    /**
     * The value of the variable.
     */
    value?: pulumi.Input<string>;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
     */
    variableType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectVariable resource.
 */
export interface ProjectVariableArgs {
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * The name of the variable.
     */
    key: pulumi.Input<string>;
    /**
     * If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     */
    masked?: pulumi.Input<boolean>;
    /**
     * The name or id of the project.
     */
    project: pulumi.Input<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    protected?: pulumi.Input<boolean>;
    /**
     * Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     */
    raw?: pulumi.Input<boolean>;
    /**
     * The value of the variable.
     */
    value: pulumi.Input<string>;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
     */
    variableType?: pulumi.Input<string>;
}
