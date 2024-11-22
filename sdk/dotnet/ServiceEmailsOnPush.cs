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
    /// The `gitlab.ServiceEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.
    /// 
    /// &gt; This resource is deprecated. Please use `gitlab.IntegrationEmailsOnPush` instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#emails-on-push)
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
    ///     var emails = new GitLab.ServiceEmailsOnPush("emails", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         Recipients = "myrecipient@example.com myotherrecipient@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_emails_on_push`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_service_emails_on_push.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// You can import a gitlab_service_emails_on_push state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush emails 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush")]
    public partial class ServiceEmailsOnPush : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        /// </summary>
        [Output("branchesToBeNotified")]
        public Output<string?> BranchesToBeNotified { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 date/time that this integration was activated at in UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Disable code diffs.
        /// </summary>
        [Output("disableDiffs")]
        public Output<bool?> DisableDiffs { get; private set; } = null!;

        /// <summary>
        /// ID or full-path of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Output("pushEvents")]
        public Output<bool?> PushEvents { get; private set; } = null!;

        /// <summary>
        /// Emails separated by whitespace.
        /// </summary>
        [Output("recipients")]
        public Output<string> Recipients { get; private set; } = null!;

        /// <summary>
        /// Send from committer.
        /// </summary>
        [Output("sendFromCommitterEmail")]
        public Output<bool?> SendFromCommitterEmail { get; private set; } = null!;

        /// <summary>
        /// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool?> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// Title of the integration.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 date/time that this integration was last updated at in UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEmailsOnPush resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEmailsOnPush(string name, ServiceEmailsOnPushArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush", name, args ?? new ServiceEmailsOnPushArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEmailsOnPush(string name, Input<string> id, ServiceEmailsOnPushState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEmailsOnPush resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEmailsOnPush Get(string name, Input<string> id, ServiceEmailsOnPushState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEmailsOnPush(name, id, state, options);
        }
    }

    public sealed class ServiceEmailsOnPushArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// Disable code diffs.
        /// </summary>
        [Input("disableDiffs")]
        public Input<bool>? DisableDiffs { get; set; }

        /// <summary>
        /// ID or full-path of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Emails separated by whitespace.
        /// </summary>
        [Input("recipients", required: true)]
        public Input<string> Recipients { get; set; } = null!;

        /// <summary>
        /// Send from committer.
        /// </summary>
        [Input("sendFromCommitterEmail")]
        public Input<bool>? SendFromCommitterEmail { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        public ServiceEmailsOnPushArgs()
        {
        }
        public static new ServiceEmailsOnPushArgs Empty => new ServiceEmailsOnPushArgs();
    }

    public sealed class ServiceEmailsOnPushState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
        /// </summary>
        [Input("branchesToBeNotified")]
        public Input<string>? BranchesToBeNotified { get; set; }

        /// <summary>
        /// The ISO8601 date/time that this integration was activated at in UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Disable code diffs.
        /// </summary>
        [Input("disableDiffs")]
        public Input<bool>? DisableDiffs { get; set; }

        /// <summary>
        /// ID or full-path of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Enable notifications for push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Emails separated by whitespace.
        /// </summary>
        [Input("recipients")]
        public Input<string>? Recipients { get; set; }

        /// <summary>
        /// Send from committer.
        /// </summary>
        [Input("sendFromCommitterEmail")]
        public Input<bool>? SendFromCommitterEmail { get; set; }

        /// <summary>
        /// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// Enable notifications for tag push events.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        /// <summary>
        /// Title of the integration.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The ISO8601 date/time that this integration was last updated at in UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ServiceEmailsOnPushState()
        {
        }
        public static new ServiceEmailsOnPushState Empty => new ServiceEmailsOnPushState();
    }
}
