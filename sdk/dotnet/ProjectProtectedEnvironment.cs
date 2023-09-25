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
    /// The `gitlab.ProjectProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a project.
    /// 
    /// &gt; In order to use a user or group in the `deploy_access_levels` configuration,
    ///    you need to make sure that users have access to the project and groups must have this project shared.
    ///    You may use the `gitlab.ProjectMembership` and `gitlab_project_shared_group` resources to achieve this.
    ///    Unfortunately, the GitLab API does not complain about users and groups without access to the project and just ignores those.
    ///    In case this happens you will get perpetual state diffs.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_environments.html)
    /// 
    /// ## Import
    /// 
    /// GitLab protected environments can be imported using an id made up of `projectId:environmentName`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment bar 123:production
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment")]
    public partial class ProjectProtectedEnvironment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Array of approval rules to deploy, with each described by a hash.
        /// </summary>
        [Output("approvalRules")]
        public Output<ImmutableArray<Outputs.ProjectProtectedEnvironmentApprovalRule>> ApprovalRules { get; private set; } = null!;

        /// <summary>
        /// Array of access levels allowed to deploy, with each described by a hash.
        /// </summary>
        [Output("deployAccessLevels")]
        public Output<ImmutableArray<Outputs.ProjectProtectedEnvironmentDeployAccessLevel>> DeployAccessLevels { get; private set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project which the protected environment is created against.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The number of approvals required to deploy to this environment.
        /// </summary>
        [Output("requiredApprovalCount")]
        public Output<int> RequiredApprovalCount { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectProtectedEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectProtectedEnvironment(string name, ProjectProtectedEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment", name, args ?? new ProjectProtectedEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectProtectedEnvironment(string name, Input<string> id, ProjectProtectedEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectProtectedEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectProtectedEnvironment Get(string name, Input<string> id, ProjectProtectedEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectProtectedEnvironment(name, id, state, options);
        }
    }

    public sealed class ProjectProtectedEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        private InputList<Inputs.ProjectProtectedEnvironmentApprovalRuleArgs>? _approvalRules;

        /// <summary>
        /// Array of approval rules to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.ProjectProtectedEnvironmentApprovalRuleArgs> ApprovalRules
        {
            get => _approvalRules ?? (_approvalRules = new InputList<Inputs.ProjectProtectedEnvironmentApprovalRuleArgs>());
            set => _approvalRules = value;
        }

        [Input("deployAccessLevels")]
        private InputList<Inputs.ProjectProtectedEnvironmentDeployAccessLevelArgs>? _deployAccessLevels;

        /// <summary>
        /// Array of access levels allowed to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.ProjectProtectedEnvironmentDeployAccessLevelArgs> DeployAccessLevels
        {
            get => _deployAccessLevels ?? (_deployAccessLevels = new InputList<Inputs.ProjectProtectedEnvironmentDeployAccessLevelArgs>());
            set => _deployAccessLevels = value;
        }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// The ID or full path of the project which the protected environment is created against.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The number of approvals required to deploy to this environment.
        /// </summary>
        [Input("requiredApprovalCount")]
        public Input<int>? RequiredApprovalCount { get; set; }

        public ProjectProtectedEnvironmentArgs()
        {
        }
        public static new ProjectProtectedEnvironmentArgs Empty => new ProjectProtectedEnvironmentArgs();
    }

    public sealed class ProjectProtectedEnvironmentState : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        private InputList<Inputs.ProjectProtectedEnvironmentApprovalRuleGetArgs>? _approvalRules;

        /// <summary>
        /// Array of approval rules to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.ProjectProtectedEnvironmentApprovalRuleGetArgs> ApprovalRules
        {
            get => _approvalRules ?? (_approvalRules = new InputList<Inputs.ProjectProtectedEnvironmentApprovalRuleGetArgs>());
            set => _approvalRules = value;
        }

        [Input("deployAccessLevels")]
        private InputList<Inputs.ProjectProtectedEnvironmentDeployAccessLevelGetArgs>? _deployAccessLevels;

        /// <summary>
        /// Array of access levels allowed to deploy, with each described by a hash.
        /// </summary>
        public InputList<Inputs.ProjectProtectedEnvironmentDeployAccessLevelGetArgs> DeployAccessLevels
        {
            get => _deployAccessLevels ?? (_deployAccessLevels = new InputList<Inputs.ProjectProtectedEnvironmentDeployAccessLevelGetArgs>());
            set => _deployAccessLevels = value;
        }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The ID or full path of the project which the protected environment is created against.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The number of approvals required to deploy to this environment.
        /// </summary>
        [Input("requiredApprovalCount")]
        public Input<int>? RequiredApprovalCount { get; set; }

        public ProjectProtectedEnvironmentState()
        {
        }
        public static new ProjectProtectedEnvironmentState Empty => new ProjectProtectedEnvironmentState();
    }
}
