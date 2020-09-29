// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ProjectLevelMrApprovals extends pulumi.CustomResource {
    /**
     * Get an existing ProjectLevelMrApprovals resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectLevelMrApprovalsState, opts?: pulumi.CustomResourceOptions): ProjectLevelMrApprovals {
        return new ProjectLevelMrApprovals(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals';

    /**
     * Returns true if the given object is an instance of ProjectLevelMrApprovals.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectLevelMrApprovals {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectLevelMrApprovals.__pulumiType;
    }

    /**
     * By default, users are able to edit the approval rules in merge requests. If set to true,
     * the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
     */
    public readonly disableOverridingApproversPerMergeRequest!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
     * also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
     */
    public readonly mergeRequestsAuthorApproval!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
     */
    public readonly mergeRequestsDisableCommittersApproval!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project to change MR approval configuration.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     */
    public readonly resetApprovalsOnPush!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ProjectLevelMrApprovals resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectLevelMrApprovalsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectLevelMrApprovalsArgs | ProjectLevelMrApprovalsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectLevelMrApprovalsState | undefined;
            inputs["disableOverridingApproversPerMergeRequest"] = state ? state.disableOverridingApproversPerMergeRequest : undefined;
            inputs["mergeRequestsAuthorApproval"] = state ? state.mergeRequestsAuthorApproval : undefined;
            inputs["mergeRequestsDisableCommittersApproval"] = state ? state.mergeRequestsDisableCommittersApproval : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["resetApprovalsOnPush"] = state ? state.resetApprovalsOnPush : undefined;
        } else {
            const args = argsOrState as ProjectLevelMrApprovalsArgs | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["disableOverridingApproversPerMergeRequest"] = args ? args.disableOverridingApproversPerMergeRequest : undefined;
            inputs["mergeRequestsAuthorApproval"] = args ? args.mergeRequestsAuthorApproval : undefined;
            inputs["mergeRequestsDisableCommittersApproval"] = args ? args.mergeRequestsDisableCommittersApproval : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["resetApprovalsOnPush"] = args ? args.resetApprovalsOnPush : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProjectLevelMrApprovals.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectLevelMrApprovals resources.
 */
export interface ProjectLevelMrApprovalsState {
    /**
     * By default, users are able to edit the approval rules in merge requests. If set to true,
     * the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
     */
    readonly disableOverridingApproversPerMergeRequest?: pulumi.Input<boolean>;
    /**
     * Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
     * also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
     */
    readonly mergeRequestsAuthorApproval?: pulumi.Input<boolean>;
    /**
     * Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
     */
    readonly mergeRequestsDisableCommittersApproval?: pulumi.Input<boolean>;
    /**
     * The ID of the project to change MR approval configuration.
     */
    readonly projectId?: pulumi.Input<number>;
    /**
     * Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     */
    readonly resetApprovalsOnPush?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ProjectLevelMrApprovals resource.
 */
export interface ProjectLevelMrApprovalsArgs {
    /**
     * By default, users are able to edit the approval rules in merge requests. If set to true,
     * the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
     */
    readonly disableOverridingApproversPerMergeRequest?: pulumi.Input<boolean>;
    /**
     * Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
     * also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
     */
    readonly mergeRequestsAuthorApproval?: pulumi.Input<boolean>;
    /**
     * Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
     */
    readonly mergeRequestsDisableCommittersApproval?: pulumi.Input<boolean>;
    /**
     * The ID of the project to change MR approval configuration.
     */
    readonly projectId: pulumi.Input<number>;
    /**
     * Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     */
    readonly resetApprovalsOnPush?: pulumi.Input<boolean>;
}