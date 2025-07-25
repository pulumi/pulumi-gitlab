// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectVariable` resource allows creating and managing a GitLab project level variable.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_level_variables/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.ProjectVariable("example", {
 *     project: "12345",
 *     key: "project_variable_key",
 *     value: "project_variable_value",
 *     "protected": false,
 * });
 * ```
 *
 * ## Import
 *
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_variable`. For example:
 *
 * terraform
 *
 * import {
 *
 *   to = gitlab_project_variable.example
 *
 *   id = "see CLI command below for ID"
 *
 * }
 *
 * Importing using the CLI is supported with the following syntax:
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
     * The description of the variable.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    public readonly environmentScope!: pulumi.Output<string>;
    /**
     * If set to `true`, the value of the variable will be hidden in the CI/CD User Interface. The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
     */
    public readonly hidden!: pulumi.Output<boolean>;
    /**
     * The name of the variable.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     */
    public readonly masked!: pulumi.Output<boolean>;
    /**
     * The name or id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     */
    public readonly protected!: pulumi.Output<boolean>;
    /**
     * Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
     */
    public readonly raw!: pulumi.Output<boolean>;
    /**
     * The value of the variable.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`.
     */
    public readonly variableType!: pulumi.Output<string>;

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
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environmentScope"] = state ? state.environmentScope : undefined;
            resourceInputs["hidden"] = state ? state.hidden : undefined;
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
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentScope"] = args ? args.environmentScope : undefined;
            resourceInputs["hidden"] = args ? args.hidden : undefined;
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
     * The description of the variable.
     */
    description?: pulumi.Input<string>;
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * If set to `true`, the value of the variable will be hidden in the CI/CD User Interface. The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
     */
    hidden?: pulumi.Input<boolean>;
    /**
     * The name of the variable.
     */
    key?: pulumi.Input<string>;
    /**
     * If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     */
    masked?: pulumi.Input<boolean>;
    /**
     * The name or id of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     */
    protected?: pulumi.Input<boolean>;
    /**
     * Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
     */
    raw?: pulumi.Input<boolean>;
    /**
     * The value of the variable.
     */
    value?: pulumi.Input<string>;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`.
     */
    variableType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectVariable resource.
 */
export interface ProjectVariableArgs {
    /**
     * The description of the variable.
     */
    description?: pulumi.Input<string>;
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * If set to `true`, the value of the variable will be hidden in the CI/CD User Interface. The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
     */
    hidden?: pulumi.Input<boolean>;
    /**
     * The name of the variable.
     */
    key: pulumi.Input<string>;
    /**
     * If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     */
    masked?: pulumi.Input<boolean>;
    /**
     * The name or id of the project.
     */
    project: pulumi.Input<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     */
    protected?: pulumi.Input<boolean>;
    /**
     * Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
     */
    raw?: pulumi.Input<boolean>;
    /**
     * The value of the variable.
     */
    value: pulumi.Input<string>;
    /**
     * The type of a variable. Valid values are: `envVar`, `file`.
     */
    variableType?: pulumi.Input<string>;
}
