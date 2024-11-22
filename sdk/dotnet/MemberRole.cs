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
    /// The `gitlab.MemberRole` resource allows to manage the lifecycle of a custom member role.
    /// 
    /// Custom roles allow an organization to create user roles with the precise privileges and permissions required for that organization’s needs.
    /// 
    /// &gt; This resource requires an Ultimate license.
    /// 
    /// &gt; Most custom roles are considered billable users that use a seat. [Custom roles billing and seat usage](https://docs.gitlab.com/ee/user/custom_roles.html#billing-and-seat-usage)
    /// 
    /// &gt; There can be only 10 custom roles on your instance or namespace. See [issue 450929](https://gitlab.com/gitlab-org/gitlab/-/issues/450929) for more details.
    /// 
    /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationmemberrolecreate)
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_member_role`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_member_role.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// GitLab member role can be imported using the id made up of `gid://gitlab/MemberRole/&lt;ID&gt;` e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/memberRole:MemberRole example 'gid://gitlab/MemberRole/123'
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/memberRole:MemberRole")]
    public partial class MemberRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
        /// </summary>
        [Output("baseAccessLevel")]
        public Output<string> BaseAccessLevel { get; private set; } = null!;

        /// <summary>
        /// Timestamp of when the member role was created. Only available with GitLab version 17.3 or higher.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description for the member role.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The Web UI path to edit the member role
        /// </summary>
        [Output("editPath")]
        public Output<string> EditPath { get; private set; } = null!;

        /// <summary>
        /// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
        /// </summary>
        [Output("enabledPermissions")]
        public Output<ImmutableArray<string>> EnabledPermissions { get; private set; } = null!;

        /// <summary>
        /// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
        /// </summary>
        [Output("groupPath")]
        public Output<string> GroupPath { get; private set; } = null!;

        /// <summary>
        /// The id integer value extracted from the `id` attribute
        /// </summary>
        [Output("iid")]
        public Output<int> Iid { get; private set; } = null!;

        /// <summary>
        /// Name for the member role.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a MemberRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MemberRole(string name, MemberRoleArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/memberRole:MemberRole", name, args ?? new MemberRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MemberRole(string name, Input<string> id, MemberRoleState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/memberRole:MemberRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MemberRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MemberRole Get(string name, Input<string> id, MemberRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new MemberRole(name, id, state, options);
        }
    }

    public sealed class MemberRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
        /// </summary>
        [Input("baseAccessLevel", required: true)]
        public Input<string> BaseAccessLevel { get; set; } = null!;

        /// <summary>
        /// Description for the member role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabledPermissions", required: true)]
        private InputList<string>? _enabledPermissions;

        /// <summary>
        /// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
        /// </summary>
        public InputList<string> EnabledPermissions
        {
            get => _enabledPermissions ?? (_enabledPermissions = new InputList<string>());
            set => _enabledPermissions = value;
        }

        /// <summary>
        /// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
        /// </summary>
        [Input("groupPath")]
        public Input<string>? GroupPath { get; set; }

        /// <summary>
        /// Name for the member role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MemberRoleArgs()
        {
        }
        public static new MemberRoleArgs Empty => new MemberRoleArgs();
    }

    public sealed class MemberRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
        /// </summary>
        [Input("baseAccessLevel")]
        public Input<string>? BaseAccessLevel { get; set; }

        /// <summary>
        /// Timestamp of when the member role was created. Only available with GitLab version 17.3 or higher.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description for the member role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Web UI path to edit the member role
        /// </summary>
        [Input("editPath")]
        public Input<string>? EditPath { get; set; }

        [Input("enabledPermissions")]
        private InputList<string>? _enabledPermissions;

        /// <summary>
        /// All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
        /// </summary>
        public InputList<string> EnabledPermissions
        {
            get => _enabledPermissions ?? (_enabledPermissions = new InputList<string>());
            set => _enabledPermissions = value;
        }

        /// <summary>
        /// Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
        /// </summary>
        [Input("groupPath")]
        public Input<string>? GroupPath { get; set; }

        /// <summary>
        /// The id integer value extracted from the `id` attribute
        /// </summary>
        [Input("iid")]
        public Input<int>? Iid { get; set; }

        /// <summary>
        /// Name for the member role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MemberRoleState()
        {
        }
        public static new MemberRoleState Empty => new MemberRoleState();
    }
}
