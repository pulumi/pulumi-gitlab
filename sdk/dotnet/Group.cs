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
    /// The `gitlab.Group` resource allows to manage the lifecycle of a group.
    /// 
    /// &gt; On GitLab SaaS, you must use the GitLab UI to create groups without a parent group. You cannot use this provider nor the API to do this.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new GitLab.Group("example", new()
    ///     {
    ///         Name = "example",
    ///         Path = "example",
    ///         Description = "An example group",
    ///     });
    /// 
    ///     // Create a project in the example group
    ///     var exampleProject = new GitLab.Project("example", new()
    ///     {
    ///         Name = "example",
    ///         Description = "An example project",
    ///         NamespaceId = example.Id,
    ///     });
    /// 
    ///     // Group with custom push rules
    ///     var example_two = new GitLab.Group("example-two", new()
    ///     {
    ///         Name = "example-two",
    ///         Path = "example-two",
    ///         Description = "An example group with push rules",
    ///         PushRules = new GitLab.Inputs.GroupPushRulesArgs
    ///         {
    ///             AuthorEmailRegex = "@example\\.com$",
    ///             CommitCommitterCheck = true,
    ///             MemberCheck = true,
    ///             PreventSecrets = true,
    ///         },
    ///     });
    /// 
    ///     // Group with custom default branch protection defaults
    ///     var example_three = new GitLab.Group("example-three", new()
    ///     {
    ///         Name = "example-three",
    ///         Path = "example-three",
    ///         Description = "An example group with default branch protection defaults",
    ///         DefaultBranchProtectionDefaults = new GitLab.Inputs.GroupDefaultBranchProtectionDefaultsArgs
    ///         {
    ///             AllowedToPushes = new[]
    ///             {
    ///                 "developer",
    ///             },
    ///             AllowForcePush = true,
    ///             AllowedToMerges = new[]
    ///             {
    ///                 "developer",
    ///                 "maintainer",
    ///             },
    ///             DeveloperCanInitialPush = true,
    ///         },
    ///     });
    /// 
    ///     // Group with custom default branch protection defaults
    ///     var example_four = new GitLab.Group("example-four", new()
    ///     {
    ///         Name = "example-four",
    ///         Path = "example-four",
    ///         Description = "An example group with default branch protection defaults",
    ///         DefaultBranchProtectionDefaults = new GitLab.Inputs.GroupDefaultBranchProtectionDefaultsArgs
    ///         {
    ///             AllowedToPushes = new[]
    ///             {
    ///                 "no one",
    ///             },
    ///             AllowForcePush = true,
    ///             AllowedToMerges = new[]
    ///             {
    ///                 "no one",
    ///             },
    ///             DeveloperCanInitialPush = true,
    ///         },
    ///     });
    /// 
    ///     // Group with a default branch name specified
    ///     var example_five = new GitLab.Group("example-five", new()
    ///     {
    ///         Name = "example",
    ///         Path = "example",
    ///         DefaultBranch = "develop",
    ///         Description = "An example group with a default branch name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_group.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/group:Group You can import a group state using `&lt;resource&gt; &lt;id&gt;`. The
    /// ```
    /// 
    /// `id` can be whatever the [details of a group][details_of_a_group] api takes for
    /// 
    /// its `:id` value, so for example:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/group:Group example example
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
        /// </summary>
        [Output("allowedEmailDomainsLists")]
        public Output<ImmutableArray<string>> AllowedEmailDomainsLists { get; private set; } = null!;

        /// <summary>
        /// Default to Auto DevOps pipeline for all projects within this group.
        /// </summary>
        [Output("autoDevopsEnabled")]
        public Output<bool> AutoDevopsEnabled { get; private set; } = null!;

        /// <summary>
        /// A local path to the avatar image to upload. **Note**: not available for imported resources.
        /// </summary>
        [Output("avatar")]
        public Output<string?> Avatar { get; private set; } = null!;

        /// <summary>
        /// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
        /// </summary>
        [Output("avatarHash")]
        public Output<string> AvatarHash { get; private set; } = null!;

        /// <summary>
        /// The URL of the avatar image.
        /// </summary>
        [Output("avatarUrl")]
        public Output<string> AvatarUrl { get; private set; } = null!;

        /// <summary>
        /// Initial default branch name.
        /// </summary>
        [Output("defaultBranch")]
        public Output<string?> DefaultBranch { get; private set; } = null!;

        /// <summary>
        /// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
        /// </summary>
        [Output("defaultBranchProtection")]
        public Output<int> DefaultBranchProtection { get; private set; } = null!;

        /// <summary>
        /// The default branch protection defaults
        /// </summary>
        [Output("defaultBranchProtectionDefaults")]
        public Output<Outputs.GroupDefaultBranchProtectionDefaults> DefaultBranchProtectionDefaults { get; private set; } = null!;

        /// <summary>
        /// The group's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Enable email notifications.
        /// </summary>
        [Output("emailsEnabled")]
        public Output<bool> EmailsEnabled { get; private set; } = null!;

        /// <summary>
        /// Can be set by administrators only. Additional CI/CD minutes for this group.
        /// </summary>
        [Output("extraSharedRunnersMinutesLimit")]
        public Output<int> ExtraSharedRunnersMinutesLimit { get; private set; } = null!;

        /// <summary>
        /// The full name of the group.
        /// </summary>
        [Output("fullName")]
        public Output<string> FullName { get; private set; } = null!;

        /// <summary>
        /// The full path of the group.
        /// </summary>
        [Output("fullPath")]
        public Output<string> FullPath { get; private set; } = null!;

        /// <summary>
        /// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
        /// </summary>
        [Output("ipRestrictionRanges")]
        public Output<ImmutableArray<string>> IpRestrictionRanges { get; private set; } = null!;

        /// <summary>
        /// Enable/disable Large File Storage (LFS) for the projects in this group.
        /// </summary>
        [Output("lfsEnabled")]
        public Output<bool> LfsEnabled { get; private set; } = null!;

        /// <summary>
        /// Users cannot be added to projects in this group.
        /// </summary>
        [Output("membershipLock")]
        public Output<bool?> MembershipLock { get; private set; } = null!;

        /// <summary>
        /// Disable the capability of a group from getting mentioned.
        /// </summary>
        [Output("mentionsDisabled")]
        public Output<bool> MentionsDisabled { get; private set; } = null!;

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Id of the parent group (creates a nested group).
        /// </summary>
        [Output("parentId")]
        public Output<int> ParentId { get; private set; } = null!;

        /// <summary>
        /// The path of the group.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
        /// </summary>
        [Output("permanentlyRemoveOnDelete")]
        public Output<bool?> PermanentlyRemoveOnDelete { get; private set; } = null!;

        /// <summary>
        /// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
        /// </summary>
        [Output("preventForkingOutsideGroup")]
        public Output<bool> PreventForkingOutsideGroup { get; private set; } = null!;

        /// <summary>
        /// Determine if developers can create projects in the group. Valid values are: `noone`, `owner`, `maintainer`, `developer`
        /// </summary>
        [Output("projectCreationLevel")]
        public Output<string> ProjectCreationLevel { get; private set; } = null!;

        /// <summary>
        /// Push rules for the group.
        /// </summary>
        [Output("pushRules")]
        public Output<Outputs.GroupPushRules> PushRules { get; private set; } = null!;

        /// <summary>
        /// Allow users to request member access.
        /// </summary>
        [Output("requestAccessEnabled")]
        public Output<bool> RequestAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// Require all users in this group to setup Two-factor authentication.
        /// </summary>
        [Output("requireTwoFactorAuthentication")]
        public Output<bool> RequireTwoFactorAuthentication { get; private set; } = null!;

        /// <summary>
        /// The group level registration token to use during runner setup.
        /// </summary>
        [Output("runnersToken")]
        public Output<string> RunnersToken { get; private set; } = null!;

        /// <summary>
        /// Prevent sharing a project with another group within this group.
        /// </summary>
        [Output("shareWithGroupLock")]
        public Output<bool> ShareWithGroupLock { get; private set; } = null!;

        /// <summary>
        /// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
        /// </summary>
        [Output("sharedRunnersMinutesLimit")]
        public Output<int> SharedRunnersMinutesLimit { get; private set; } = null!;

        /// <summary>
        /// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
        /// </summary>
        [Output("sharedRunnersSetting")]
        public Output<string> SharedRunnersSetting { get; private set; } = null!;

        /// <summary>
        /// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
        /// </summary>
        [Output("subgroupCreationLevel")]
        public Output<string> SubgroupCreationLevel { get; private set; } = null!;

        /// <summary>
        /// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
        /// </summary>
        [Output("twoFactorGracePeriod")]
        public Output<int> TwoFactorGracePeriod { get; private set; } = null!;

        /// <summary>
        /// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
        /// </summary>
        [Output("visibilityLevel")]
        public Output<string> VisibilityLevel { get; private set; } = null!;

        /// <summary>
        /// Web URL of the group.
        /// </summary>
        [Output("webUrl")]
        public Output<string> WebUrl { get; private set; } = null!;

        /// <summary>
        /// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        [Output("wikiAccessLevel")]
        public Output<string> WikiAccessLevel { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/group:Group", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "runnersToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedEmailDomainsLists")]
        private InputList<string>? _allowedEmailDomainsLists;

        /// <summary>
        /// A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
        /// </summary>
        public InputList<string> AllowedEmailDomainsLists
        {
            get => _allowedEmailDomainsLists ?? (_allowedEmailDomainsLists = new InputList<string>());
            set => _allowedEmailDomainsLists = value;
        }

        /// <summary>
        /// Default to Auto DevOps pipeline for all projects within this group.
        /// </summary>
        [Input("autoDevopsEnabled")]
        public Input<bool>? AutoDevopsEnabled { get; set; }

        /// <summary>
        /// A local path to the avatar image to upload. **Note**: not available for imported resources.
        /// </summary>
        [Input("avatar")]
        public Input<string>? Avatar { get; set; }

        /// <summary>
        /// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
        /// </summary>
        [Input("avatarHash")]
        public Input<string>? AvatarHash { get; set; }

        /// <summary>
        /// Initial default branch name.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
        /// </summary>
        [Input("defaultBranchProtection")]
        public Input<int>? DefaultBranchProtection { get; set; }

        /// <summary>
        /// The default branch protection defaults
        /// </summary>
        [Input("defaultBranchProtectionDefaults")]
        public Input<Inputs.GroupDefaultBranchProtectionDefaultsArgs>? DefaultBranchProtectionDefaults { get; set; }

        /// <summary>
        /// The group's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable email notifications.
        /// </summary>
        [Input("emailsEnabled")]
        public Input<bool>? EmailsEnabled { get; set; }

        /// <summary>
        /// Can be set by administrators only. Additional CI/CD minutes for this group.
        /// </summary>
        [Input("extraSharedRunnersMinutesLimit")]
        public Input<int>? ExtraSharedRunnersMinutesLimit { get; set; }

        [Input("ipRestrictionRanges")]
        private InputList<string>? _ipRestrictionRanges;

        /// <summary>
        /// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
        /// </summary>
        public InputList<string> IpRestrictionRanges
        {
            get => _ipRestrictionRanges ?? (_ipRestrictionRanges = new InputList<string>());
            set => _ipRestrictionRanges = value;
        }

        /// <summary>
        /// Enable/disable Large File Storage (LFS) for the projects in this group.
        /// </summary>
        [Input("lfsEnabled")]
        public Input<bool>? LfsEnabled { get; set; }

        /// <summary>
        /// Users cannot be added to projects in this group.
        /// </summary>
        [Input("membershipLock")]
        public Input<bool>? MembershipLock { get; set; }

        /// <summary>
        /// Disable the capability of a group from getting mentioned.
        /// </summary>
        [Input("mentionsDisabled")]
        public Input<bool>? MentionsDisabled { get; set; }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Id of the parent group (creates a nested group).
        /// </summary>
        [Input("parentId")]
        public Input<int>? ParentId { get; set; }

        /// <summary>
        /// The path of the group.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
        /// </summary>
        [Input("permanentlyRemoveOnDelete")]
        public Input<bool>? PermanentlyRemoveOnDelete { get; set; }

        /// <summary>
        /// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
        /// </summary>
        [Input("preventForkingOutsideGroup")]
        public Input<bool>? PreventForkingOutsideGroup { get; set; }

        /// <summary>
        /// Determine if developers can create projects in the group. Valid values are: `noone`, `owner`, `maintainer`, `developer`
        /// </summary>
        [Input("projectCreationLevel")]
        public Input<string>? ProjectCreationLevel { get; set; }

        /// <summary>
        /// Push rules for the group.
        /// </summary>
        [Input("pushRules")]
        public Input<Inputs.GroupPushRulesArgs>? PushRules { get; set; }

        /// <summary>
        /// Allow users to request member access.
        /// </summary>
        [Input("requestAccessEnabled")]
        public Input<bool>? RequestAccessEnabled { get; set; }

        /// <summary>
        /// Require all users in this group to setup Two-factor authentication.
        /// </summary>
        [Input("requireTwoFactorAuthentication")]
        public Input<bool>? RequireTwoFactorAuthentication { get; set; }

        /// <summary>
        /// Prevent sharing a project with another group within this group.
        /// </summary>
        [Input("shareWithGroupLock")]
        public Input<bool>? ShareWithGroupLock { get; set; }

        /// <summary>
        /// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
        /// </summary>
        [Input("sharedRunnersMinutesLimit")]
        public Input<int>? SharedRunnersMinutesLimit { get; set; }

        /// <summary>
        /// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
        /// </summary>
        [Input("sharedRunnersSetting")]
        public Input<string>? SharedRunnersSetting { get; set; }

        /// <summary>
        /// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
        /// </summary>
        [Input("subgroupCreationLevel")]
        public Input<string>? SubgroupCreationLevel { get; set; }

        /// <summary>
        /// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
        /// </summary>
        [Input("twoFactorGracePeriod")]
        public Input<int>? TwoFactorGracePeriod { get; set; }

        /// <summary>
        /// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
        /// </summary>
        [Input("visibilityLevel")]
        public Input<string>? VisibilityLevel { get; set; }

        /// <summary>
        /// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        [Input("wikiAccessLevel")]
        public Input<string>? WikiAccessLevel { get; set; }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        [Input("allowedEmailDomainsLists")]
        private InputList<string>? _allowedEmailDomainsLists;

        /// <summary>
        /// A list of email address domains to allow group access. Will be concatenated together into a comma separated string.
        /// </summary>
        public InputList<string> AllowedEmailDomainsLists
        {
            get => _allowedEmailDomainsLists ?? (_allowedEmailDomainsLists = new InputList<string>());
            set => _allowedEmailDomainsLists = value;
        }

        /// <summary>
        /// Default to Auto DevOps pipeline for all projects within this group.
        /// </summary>
        [Input("autoDevopsEnabled")]
        public Input<bool>? AutoDevopsEnabled { get; set; }

        /// <summary>
        /// A local path to the avatar image to upload. **Note**: not available for imported resources.
        /// </summary>
        [Input("avatar")]
        public Input<string>? Avatar { get; set; }

        /// <summary>
        /// The hash of the avatar image. Use `filesha256("path/to/avatar.png")` whenever possible. **Note**: this is used to trigger an update of the avatar. If it's not given, but an avatar is given, the avatar will be updated each time.
        /// </summary>
        [Input("avatarHash")]
        public Input<string>? AvatarHash { get; set; }

        /// <summary>
        /// The URL of the avatar image.
        /// </summary>
        [Input("avatarUrl")]
        public Input<string>? AvatarUrl { get; set; }

        /// <summary>
        /// Initial default branch name.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// See https://docs.gitlab.com/ee/api/groups.html#options-for-default*branch*protection. Valid values are: `0`, `1`, `2`, `3`, `4`.
        /// </summary>
        [Input("defaultBranchProtection")]
        public Input<int>? DefaultBranchProtection { get; set; }

        /// <summary>
        /// The default branch protection defaults
        /// </summary>
        [Input("defaultBranchProtectionDefaults")]
        public Input<Inputs.GroupDefaultBranchProtectionDefaultsGetArgs>? DefaultBranchProtectionDefaults { get; set; }

        /// <summary>
        /// The group's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable email notifications.
        /// </summary>
        [Input("emailsEnabled")]
        public Input<bool>? EmailsEnabled { get; set; }

        /// <summary>
        /// Can be set by administrators only. Additional CI/CD minutes for this group.
        /// </summary>
        [Input("extraSharedRunnersMinutesLimit")]
        public Input<int>? ExtraSharedRunnersMinutesLimit { get; set; }

        /// <summary>
        /// The full name of the group.
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// The full path of the group.
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        [Input("ipRestrictionRanges")]
        private InputList<string>? _ipRestrictionRanges;

        /// <summary>
        /// A list of IP addresses or subnet masks to restrict group access. Will be concatenated together into a comma separated string. Only allowed on top level groups.
        /// </summary>
        public InputList<string> IpRestrictionRanges
        {
            get => _ipRestrictionRanges ?? (_ipRestrictionRanges = new InputList<string>());
            set => _ipRestrictionRanges = value;
        }

        /// <summary>
        /// Enable/disable Large File Storage (LFS) for the projects in this group.
        /// </summary>
        [Input("lfsEnabled")]
        public Input<bool>? LfsEnabled { get; set; }

        /// <summary>
        /// Users cannot be added to projects in this group.
        /// </summary>
        [Input("membershipLock")]
        public Input<bool>? MembershipLock { get; set; }

        /// <summary>
        /// Disable the capability of a group from getting mentioned.
        /// </summary>
        [Input("mentionsDisabled")]
        public Input<bool>? MentionsDisabled { get; set; }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Id of the parent group (creates a nested group).
        /// </summary>
        [Input("parentId")]
        public Input<int>? ParentId { get; set; }

        /// <summary>
        /// The path of the group.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Whether the group should be permanently removed during a `delete` operation. This only works with subgroups. Must be configured via an `apply` before the `destroy` is run.
        /// </summary>
        [Input("permanentlyRemoveOnDelete")]
        public Input<bool>? PermanentlyRemoveOnDelete { get; set; }

        /// <summary>
        /// Defaults to false. When enabled, users can not fork projects from this group to external namespaces.
        /// </summary>
        [Input("preventForkingOutsideGroup")]
        public Input<bool>? PreventForkingOutsideGroup { get; set; }

        /// <summary>
        /// Determine if developers can create projects in the group. Valid values are: `noone`, `owner`, `maintainer`, `developer`
        /// </summary>
        [Input("projectCreationLevel")]
        public Input<string>? ProjectCreationLevel { get; set; }

        /// <summary>
        /// Push rules for the group.
        /// </summary>
        [Input("pushRules")]
        public Input<Inputs.GroupPushRulesGetArgs>? PushRules { get; set; }

        /// <summary>
        /// Allow users to request member access.
        /// </summary>
        [Input("requestAccessEnabled")]
        public Input<bool>? RequestAccessEnabled { get; set; }

        /// <summary>
        /// Require all users in this group to setup Two-factor authentication.
        /// </summary>
        [Input("requireTwoFactorAuthentication")]
        public Input<bool>? RequireTwoFactorAuthentication { get; set; }

        [Input("runnersToken")]
        private Input<string>? _runnersToken;

        /// <summary>
        /// The group level registration token to use during runner setup.
        /// </summary>
        public Input<string>? RunnersToken
        {
            get => _runnersToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _runnersToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Prevent sharing a project with another group within this group.
        /// </summary>
        [Input("shareWithGroupLock")]
        public Input<bool>? ShareWithGroupLock { get; set; }

        /// <summary>
        /// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or &gt; 0.
        /// </summary>
        [Input("sharedRunnersMinutesLimit")]
        public Input<int>? SharedRunnersMinutesLimit { get; set; }

        /// <summary>
        /// Enable or disable shared runners for a group’s subgroups and projects. Valid values are: `enabled`, `disabled_and_overridable`, `disabled_and_unoverridable`, `disabled_with_override`.
        /// </summary>
        [Input("sharedRunnersSetting")]
        public Input<string>? SharedRunnersSetting { get; set; }

        /// <summary>
        /// Allowed to create subgroups. Valid values are: `owner`, `maintainer`.
        /// </summary>
        [Input("subgroupCreationLevel")]
        public Input<string>? SubgroupCreationLevel { get; set; }

        /// <summary>
        /// Defaults to 48. Time before Two-factor authentication is enforced (in hours).
        /// </summary>
        [Input("twoFactorGracePeriod")]
        public Input<int>? TwoFactorGracePeriod { get; set; }

        /// <summary>
        /// The group's visibility. Can be `private`, `internal`, or `public`. Valid values are: `private`, `internal`, `public`.
        /// </summary>
        [Input("visibilityLevel")]
        public Input<string>? VisibilityLevel { get; set; }

        /// <summary>
        /// Web URL of the group.
        /// </summary>
        [Input("webUrl")]
        public Input<string>? WebUrl { get; set; }

        /// <summary>
        /// The group's wiki access level. Only available on Premium and Ultimate plans. Valid values are `disabled`, `private`, `enabled`.
        /// </summary>
        [Input("wikiAccessLevel")]
        public Input<string>? WikiAccessLevel { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
