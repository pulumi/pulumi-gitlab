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
    /// The `gitlab.ServicePipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.
    /// 
    /// &gt; This resource is deprecated. use `gitlab.IntegrationPipelinesEmail`instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails)
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
    ///     var awesomeProject = new GitLab.Project("awesome_project", new()
    ///     {
    ///         Name = "awesome_project",
    ///         Description = "My awesome project.",
    ///         VisibilityLevel = "public",
    ///     });
    /// 
    ///     var email = new GitLab.ServicePipelinesEmail("email", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Recipients = new[]
    ///         {
    ///             "gitlab@user.create",
    ///         },
    ///         NotifyOnlyBrokenPipelines = true,
    ///         BranchesToBeNotified = "all",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_pipelines_email`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_service_pipelines_email.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// You can import a gitlab_service_pipelines_email state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/servicePipelinesEmail:ServicePipelinesEmail email 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail")]
    public partial class ServicePipelinesEmail : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        /// </summary>
        [Output("branchesToBeNotified")]
        public Output<string?> BranchesToBeNotified { get; private set; } = null!;

        /// <summary>
        /// Notify only broken pipelines. Default is true.
        /// </summary>
        [Output("notifyOnlyBrokenPipelines")]
        public Output<bool?> NotifyOnlyBrokenPipelines { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// ) email addresses where notifications are sent.
        /// </summary>
        [Output("recipients")]
        public Output<ImmutableArray<string>> Recipients { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePipelinesEmail resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePipelinesEmail(string name, ServicePipelinesEmailArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail", name, args ?? new ServicePipelinesEmailArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePipelinesEmail(string name, Input<string> id, ServicePipelinesEmailState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServicePipelinesEmail resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePipelinesEmail Get(string name, Input<string> id, ServicePipelinesEmailState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePipelinesEmail(name, id, state, options);
        }
    }

    public sealed class ServicePipelinesEmailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Notify only broken pipelines. Default is true.
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("recipients", required: true)]
        private InputList<string>? _recipients;

        /// <summary>
        /// ) email addresses where notifications are sent.
        /// </summary>
        public InputList<string> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<string>());
            set => _recipients = value;
        }

        public ServicePipelinesEmailArgs()
        {
        }
        public static new ServicePipelinesEmailArgs Empty => new ServicePipelinesEmailArgs();
    }

    public sealed class ServicePipelinesEmailState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `default_and_protected`. Default is `default`
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Notify only broken pipelines. Default is true.
        /// </summary>
        [Input("notifyOnlyBrokenPipelines")]
        public Input<bool>? NotifyOnlyBrokenPipelines { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("recipients")]
        private InputList<string>? _recipients;

        /// <summary>
        /// ) email addresses where notifications are sent.
        /// </summary>
        public InputList<string> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<string>());
            set => _recipients = value;
        }

        public ServicePipelinesEmailState()
        {
        }
        public static new ServicePipelinesEmailState Empty => new ServicePipelinesEmailState();
    }
}
