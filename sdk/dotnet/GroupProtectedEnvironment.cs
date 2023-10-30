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
    /// The `gitlab.GroupProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a group.
    /// 
    /// &gt; In order to use a user_id in the `deploy_access_levels` configuration,
    ///    you need to make sure that users have access to the group with Maintainer role or higher.
    ///    In order to use a group_id in the `deploy_access_levels` configuration,
    ///    the group_id must be a sub-group under the given group.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_protected_environments.html)
    /// 
    /// ## Import
    /// 
    /// GitLab group protected environments can be imported using an id made up of `groupId:environmentName`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment bar 123:production
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment")]
    public partial class GroupProtectedEnvironment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Array of approval rules to deploy, with each described by a hash.
        /// </summary>
        [Output("approvalRules")]
        public Output<ImmutableArray<Outputs.GroupProtectedEnvironmentApprovalRule>> ApprovalRules { get; private set; } = null!;

        /// <summary>
        /// Array of access levels allowed to deploy, with each described by a hash.
        /// </summary>
        [Output("deployAccessLevels")]
        public Output<ImmutableArray<Outputs.GroupProtectedEnvironmentDeployAccessLevel>> DeployAccessLevels { get; private set; } = null!;

        /// <summary>
        /// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the group which the protected environment is created against.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The number of approvals required to deploy to this environment.
        /// </summary>
        [Output("requiredApprovalCount")]
        public Output<int> RequiredApprovalCount { get; private set; } = null!;


        /// <summary>
        /// Create a GroupProtectedEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupProtectedEnvironment(string name, GroupProtectedEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment", name, args ?? new GroupProtectedEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupProtectedEnvironment(string name, Input<string> id, GroupProtectedEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupProtectedEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupProtectedEnvironment Get(string name, Input<string> id, GroupProtectedEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupProtectedEnvironment(name, id, state, options);
        }
    }

    public sealed class GroupProtectedEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        private InputList<Inputs.GroupProtectedEnvironmentApprovalRuleArgs>? _approvalRules;

        /// <summary>
        /// Array of approval rules to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.GroupProtectedEnvironmentApprovalRuleArgs> ApprovalRules
        {
            get => _approvalRules ?? (_approvalRules = new InputList<Inputs.GroupProtectedEnvironmentApprovalRuleArgs>());
            set => _approvalRules = value;
        }

        [Input("deployAccessLevels", required: true)]
        private InputList<Inputs.GroupProtectedEnvironmentDeployAccessLevelArgs>? _deployAccessLevels;

        /// <summary>
        /// Array of access levels allowed to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.GroupProtectedEnvironmentDeployAccessLevelArgs> DeployAccessLevels
        {
            get => _deployAccessLevels ?? (_deployAccessLevels = new InputList<Inputs.GroupProtectedEnvironmentDeployAccessLevelArgs>());
            set => _deployAccessLevels = value;
        }

        /// <summary>
        /// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// The ID or full path of the group which the protected environment is created against.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The number of approvals required to deploy to this environment.
        /// </summary>
        [Input("requiredApprovalCount")]
        public Input<int>? RequiredApprovalCount { get; set; }

        public GroupProtectedEnvironmentArgs()
        {
        }
        public static new GroupProtectedEnvironmentArgs Empty => new GroupProtectedEnvironmentArgs();
    }

    public sealed class GroupProtectedEnvironmentState : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        private InputList<Inputs.GroupProtectedEnvironmentApprovalRuleGetArgs>? _approvalRules;

        /// <summary>
        /// Array of approval rules to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.GroupProtectedEnvironmentApprovalRuleGetArgs> ApprovalRules
        {
            get => _approvalRules ?? (_approvalRules = new InputList<Inputs.GroupProtectedEnvironmentApprovalRuleGetArgs>());
            set => _approvalRules = value;
        }

        [Input("deployAccessLevels")]
        private InputList<Inputs.GroupProtectedEnvironmentDeployAccessLevelGetArgs>? _deployAccessLevels;

        /// <summary>
        /// Array of access levels allowed to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.GroupProtectedEnvironmentDeployAccessLevelGetArgs> DeployAccessLevels
        {
            get => _deployAccessLevels ?? (_deployAccessLevels = new InputList<Inputs.GroupProtectedEnvironmentDeployAccessLevelGetArgs>());
            set => _deployAccessLevels = value;
        }

        /// <summary>
        /// The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The ID or full path of the group which the protected environment is created against.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The number of approvals required to deploy to this environment.
        /// </summary>
        [Input("requiredApprovalCount")]
        public Input<int>? RequiredApprovalCount { get; set; }

        public GroupProtectedEnvironmentState()
        {
        }
        public static new GroupProtectedEnvironmentState Empty => new GroupProtectedEnvironmentState();
    }
}
