// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_project\_approval\_rule
 *
 * This resource allows you to create and manage multiple approval rules for your GitLab
 * projects. For further information on approval rules, consult the [gitlab
 * documentation](https://docs.gitlab.com/ee/api/merge_request_approvals.html#project-level-mr-approvals).
 *
 * > This feature requires a GitLab Starter account or above.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example_one = new gitlab.ProjectApprovalRule("example-one", {
 *     approvalsRequired: 3,
 *     groupIds: [51],
 *     project: "5",
 *     userIds: [
 *         50,
 *         500,
 *     ],
 * });
 * const example_two = new gitlab.ProjectApprovalRule("example-two", {
 *     approvalsRequired: 1,
 *     groupIds: [52],
 *     project: "5",
 *     userIds: [],
 * });
 * ```
 *
 * ## Import
 *
 * GitLab project approval rules can be imported using an id consisting of `project-id:rule-id`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectApprovalRule:ProjectApprovalRule example "12345:6"
 * ```
 */
export class ProjectApprovalRule extends pulumi.CustomResource {
    /**
     * Get an existing ProjectApprovalRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectApprovalRuleState, opts?: pulumi.CustomResourceOptions): ProjectApprovalRule {
        return new ProjectApprovalRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectApprovalRule:ProjectApprovalRule';

    /**
     * Returns true if the given object is an instance of ProjectApprovalRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectApprovalRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectApprovalRule.__pulumiType;
    }

    /**
     * The number of approvals required for this rule.
     */
    public readonly approvalsRequired!: pulumi.Output<number>;
    /**
     * A list of group IDs who's members can approve of the merge request
     */
    public readonly groupIds!: pulumi.Output<number[] | undefined>;
    /**
     * The name of the approval rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the approval rules.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A list of specific User IDs to add to the list of approvers.
     */
    public readonly userIds!: pulumi.Output<number[] | undefined>;

    /**
     * Create a ProjectApprovalRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectApprovalRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectApprovalRuleArgs | ProjectApprovalRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectApprovalRuleState | undefined;
            inputs["approvalsRequired"] = state ? state.approvalsRequired : undefined;
            inputs["groupIds"] = state ? state.groupIds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["userIds"] = state ? state.userIds : undefined;
        } else {
            const args = argsOrState as ProjectApprovalRuleArgs | undefined;
            if ((!args || args.approvalsRequired === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'approvalsRequired'");
            }
            if ((!args || args.project === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'project'");
            }
            inputs["approvalsRequired"] = args ? args.approvalsRequired : undefined;
            inputs["groupIds"] = args ? args.groupIds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["userIds"] = args ? args.userIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProjectApprovalRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectApprovalRule resources.
 */
export interface ProjectApprovalRuleState {
    /**
     * The number of approvals required for this rule.
     */
    readonly approvalsRequired?: pulumi.Input<number>;
    /**
     * A list of group IDs who's members can approve of the merge request
     */
    readonly groupIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The name of the approval rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the approval rules.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A list of specific User IDs to add to the list of approvers.
     */
    readonly userIds?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a ProjectApprovalRule resource.
 */
export interface ProjectApprovalRuleArgs {
    /**
     * The number of approvals required for this rule.
     */
    readonly approvalsRequired: pulumi.Input<number>;
    /**
     * A list of group IDs who's members can approve of the merge request
     */
    readonly groupIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The name of the approval rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the approval rules.
     */
    readonly project: pulumi.Input<string>;
    /**
     * A list of specific User IDs to add to the list of approvers.
     */
    readonly userIds?: pulumi.Input<pulumi.Input<number>[]>;
}