// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// ## Import
    /// 
    /// GitLab project approval rules can be imported using a key composed of `&lt;project-id&gt;:&lt;rule-id&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectApprovalRule:ProjectApprovalRule example "12345:6"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectApprovalRule:ProjectApprovalRule")]
    public partial class ProjectApprovalRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the rule is applied to all protected branches. If set to 'true', the value of `protected_branch_ids` is ignored. Default is 'false'.
        /// </summary>
        [Output("appliesToAllProtectedBranches")]
        public Output<bool?> AppliesToAllProtectedBranches { get; private set; } = null!;

        /// <summary>
        /// The number of approvals required for this rule.
        /// </summary>
        [Output("approvalsRequired")]
        public Output<int> ApprovalsRequired { get; private set; } = null!;

        /// <summary>
        /// When this flag is set, the default `any_approver` rule will not be imported if present.
        /// </summary>
        [Output("disableImportingDefaultAnyApproverRuleOnCreate")]
        public Output<bool?> DisableImportingDefaultAnyApproverRuleOnCreate { get; private set; } = null!;

        /// <summary>
        /// A list of group IDs whose members can approve of the merge request.
        /// </summary>
        [Output("groupIds")]
        public Output<ImmutableArray<int>> GroupIds { get; private set; } = null!;

        /// <summary>
        /// The name of the approval rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name or id of the project to add the approval rules.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A list of protected branch IDs (not branch names) for which the rule applies.
        /// </summary>
        [Output("protectedBranchIds")]
        public Output<ImmutableArray<int>> ProtectedBranchIds { get; private set; } = null!;

        /// <summary>
        /// String, defaults to 'regular'. The type of rule. `any_approver` is a pre-configured default rule with `approvals_required` at `0`. Valid values are `regular`, `any_approver`.
        /// </summary>
        [Output("ruleType")]
        public Output<string> RuleType { get; private set; } = null!;

        /// <summary>
        /// A list of specific User IDs to add to the list of approvers.
        /// </summary>
        [Output("userIds")]
        public Output<ImmutableArray<int>> UserIds { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectApprovalRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectApprovalRule(string name, ProjectApprovalRuleArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectApprovalRule:ProjectApprovalRule", name, args ?? new ProjectApprovalRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectApprovalRule(string name, Input<string> id, ProjectApprovalRuleState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectApprovalRule:ProjectApprovalRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProjectApprovalRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectApprovalRule Get(string name, Input<string> id, ProjectApprovalRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectApprovalRule(name, id, state, options);
        }
    }

    public sealed class ProjectApprovalRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the rule is applied to all protected branches. If set to 'true', the value of `protected_branch_ids` is ignored. Default is 'false'.
        /// </summary>
        [Input("appliesToAllProtectedBranches")]
        public Input<bool>? AppliesToAllProtectedBranches { get; set; }

        /// <summary>
        /// The number of approvals required for this rule.
        /// </summary>
        [Input("approvalsRequired", required: true)]
        public Input<int> ApprovalsRequired { get; set; } = null!;

        /// <summary>
        /// When this flag is set, the default `any_approver` rule will not be imported if present.
        /// </summary>
        [Input("disableImportingDefaultAnyApproverRuleOnCreate")]
        public Input<bool>? DisableImportingDefaultAnyApproverRuleOnCreate { get; set; }

        [Input("groupIds")]
        private InputList<int>? _groupIds;

        /// <summary>
        /// A list of group IDs whose members can approve of the merge request.
        /// </summary>
        public InputList<int> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<int>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The name of the approval rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or id of the project to add the approval rules.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("protectedBranchIds")]
        private InputList<int>? _protectedBranchIds;

        /// <summary>
        /// A list of protected branch IDs (not branch names) for which the rule applies.
        /// </summary>
        public InputList<int> ProtectedBranchIds
        {
            get => _protectedBranchIds ?? (_protectedBranchIds = new InputList<int>());
            set => _protectedBranchIds = value;
        }

        /// <summary>
        /// String, defaults to 'regular'. The type of rule. `any_approver` is a pre-configured default rule with `approvals_required` at `0`. Valid values are `regular`, `any_approver`.
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        [Input("userIds")]
        private InputList<int>? _userIds;

        /// <summary>
        /// A list of specific User IDs to add to the list of approvers.
        /// </summary>
        public InputList<int> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<int>());
            set => _userIds = value;
        }

        public ProjectApprovalRuleArgs()
        {
        }
        public static new ProjectApprovalRuleArgs Empty => new ProjectApprovalRuleArgs();
    }

    public sealed class ProjectApprovalRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the rule is applied to all protected branches. If set to 'true', the value of `protected_branch_ids` is ignored. Default is 'false'.
        /// </summary>
        [Input("appliesToAllProtectedBranches")]
        public Input<bool>? AppliesToAllProtectedBranches { get; set; }

        /// <summary>
        /// The number of approvals required for this rule.
        /// </summary>
        [Input("approvalsRequired")]
        public Input<int>? ApprovalsRequired { get; set; }

        /// <summary>
        /// When this flag is set, the default `any_approver` rule will not be imported if present.
        /// </summary>
        [Input("disableImportingDefaultAnyApproverRuleOnCreate")]
        public Input<bool>? DisableImportingDefaultAnyApproverRuleOnCreate { get; set; }

        [Input("groupIds")]
        private InputList<int>? _groupIds;

        /// <summary>
        /// A list of group IDs whose members can approve of the merge request.
        /// </summary>
        public InputList<int> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<int>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The name of the approval rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or id of the project to add the approval rules.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("protectedBranchIds")]
        private InputList<int>? _protectedBranchIds;

        /// <summary>
        /// A list of protected branch IDs (not branch names) for which the rule applies.
        /// </summary>
        public InputList<int> ProtectedBranchIds
        {
            get => _protectedBranchIds ?? (_protectedBranchIds = new InputList<int>());
            set => _protectedBranchIds = value;
        }

        /// <summary>
        /// String, defaults to 'regular'. The type of rule. `any_approver` is a pre-configured default rule with `approvals_required` at `0`. Valid values are `regular`, `any_approver`.
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        [Input("userIds")]
        private InputList<int>? _userIds;

        /// <summary>
        /// A list of specific User IDs to add to the list of approvers.
        /// </summary>
        public InputList<int> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<int>());
            set => _userIds = value;
        }

        public ProjectApprovalRuleState()
        {
        }
        public static new ProjectApprovalRuleState Empty => new ProjectApprovalRuleState();
    }
}
