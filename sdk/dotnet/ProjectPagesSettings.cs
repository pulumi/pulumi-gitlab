// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_pages_settings`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_project_pages_settings.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Importing using the CLI is supported with the following syntax:
    /// 
    /// Gitlab project pages settings can be imported using the project ID, for example:
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectPagesSettings:ProjectPagesSettings example 12345
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectPagesSettings:ProjectPagesSettings")]
    public partial class ProjectPagesSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of current active deployments.
        /// </summary>
        [Output("deployments")]
        public Output<ImmutableArray<Outputs.ProjectPagesSettingsDeployment>> Deployments { get; private set; } = null!;

        /// <summary>
        /// Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
        /// </summary>
        [Output("forceHttps")]
        public Output<bool> ForceHttps { get; private set; } = null!;

        /// <summary>
        /// Boolean indicating if a unique domain is enabled.
        /// </summary>
        [Output("isUniqueDomainEnabled")]
        public Output<bool> IsUniqueDomainEnabled { get; private set; } = null!;

        [Output("keepSettingsOnDestroy")]
        public Output<bool> KeepSettingsOnDestroy { get; private set; } = null!;

        /// <summary>
        /// The project ID or path.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URL to access the project pages.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectPagesSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectPagesSettings(string name, ProjectPagesSettingsArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectPagesSettings:ProjectPagesSettings", name, args ?? new ProjectPagesSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectPagesSettings(string name, Input<string> id, ProjectPagesSettingsState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectPagesSettings:ProjectPagesSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectPagesSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectPagesSettings Get(string name, Input<string> id, ProjectPagesSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectPagesSettings(name, id, state, options);
        }
    }

    public sealed class ProjectPagesSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
        /// </summary>
        [Input("forceHttps")]
        public Input<bool>? ForceHttps { get; set; }

        /// <summary>
        /// Boolean indicating if a unique domain is enabled.
        /// </summary>
        [Input("isUniqueDomainEnabled")]
        public Input<bool>? IsUniqueDomainEnabled { get; set; }

        [Input("keepSettingsOnDestroy")]
        public Input<bool>? KeepSettingsOnDestroy { get; set; }

        /// <summary>
        /// The project ID or path.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ProjectPagesSettingsArgs()
        {
        }
        public static new ProjectPagesSettingsArgs Empty => new ProjectPagesSettingsArgs();
    }

    public sealed class ProjectPagesSettingsState : global::Pulumi.ResourceArgs
    {
        [Input("deployments")]
        private InputList<Inputs.ProjectPagesSettingsDeploymentGetArgs>? _deployments;

        /// <summary>
        /// List of current active deployments.
        /// </summary>
        public InputList<Inputs.ProjectPagesSettingsDeploymentGetArgs> Deployments
        {
            get => _deployments ?? (_deployments = new InputList<Inputs.ProjectPagesSettingsDeploymentGetArgs>());
            set => _deployments = value;
        }

        /// <summary>
        /// Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
        /// </summary>
        [Input("forceHttps")]
        public Input<bool>? ForceHttps { get; set; }

        /// <summary>
        /// Boolean indicating if a unique domain is enabled.
        /// </summary>
        [Input("isUniqueDomainEnabled")]
        public Input<bool>? IsUniqueDomainEnabled { get; set; }

        [Input("keepSettingsOnDestroy")]
        public Input<bool>? KeepSettingsOnDestroy { get; set; }

        /// <summary>
        /// The project ID or path.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URL to access the project pages.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProjectPagesSettingsState()
        {
        }
        public static new ProjectPagesSettingsState Empty => new ProjectPagesSettingsState();
    }
}
