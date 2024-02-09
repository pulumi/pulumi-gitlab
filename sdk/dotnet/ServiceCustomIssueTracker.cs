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
    /// The `gitlab.ServiceCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.
    /// 
    /// &gt; This resource is deprecated. use `gitlab.IntegrationCustomIssueTracker`instead!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker)
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
    ///     var awesomeProject = new GitLab.Project("awesomeProject", new()
    ///     {
    ///         Description = "My awesome project.",
    ///         VisibilityLevel = "public",
    ///     });
    /// 
    ///     var tracker = new GitLab.ServiceCustomIssueTracker("tracker", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         ProjectUrl = "https://customtracker.com/issues",
    ///         IssuesUrl = "https://customtracker.com/TEST-:id",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import a gitlab_service_custom_issue_tracker state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker tracker 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker")]
    public partial class ServiceCustomIssueTracker : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 date/time that this integration was activated at in UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The URL to view an issue in the external issue tracker. Must contain :id.
        /// </summary>
        [Output("issuesUrl")]
        public Output<string> IssuesUrl { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project for the custom issue tracker.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URL to the project in the external issue tracker.
        /// </summary>
        [Output("projectUrl")]
        public Output<string> ProjectUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 date/time that this integration was last updated at in UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceCustomIssueTracker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceCustomIssueTracker(string name, ServiceCustomIssueTrackerArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker", name, args ?? new ServiceCustomIssueTrackerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceCustomIssueTracker(string name, Input<string> id, ServiceCustomIssueTrackerState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceCustomIssueTracker resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceCustomIssueTracker Get(string name, Input<string> id, ServiceCustomIssueTrackerState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceCustomIssueTracker(name, id, state, options);
        }
    }

    public sealed class ServiceCustomIssueTrackerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to view an issue in the external issue tracker. Must contain :id.
        /// </summary>
        [Input("issuesUrl", required: true)]
        public Input<string> IssuesUrl { get; set; } = null!;

        /// <summary>
        /// The ID or full path of the project for the custom issue tracker.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The URL to the project in the external issue tracker.
        /// </summary>
        [Input("projectUrl", required: true)]
        public Input<string> ProjectUrl { get; set; } = null!;

        public ServiceCustomIssueTrackerArgs()
        {
        }
        public static new ServiceCustomIssueTrackerArgs Empty => new ServiceCustomIssueTrackerArgs();
    }

    public sealed class ServiceCustomIssueTrackerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is active.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The ISO8601 date/time that this integration was activated at in UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The URL to view an issue in the external issue tracker. Must contain :id.
        /// </summary>
        [Input("issuesUrl")]
        public Input<string>? IssuesUrl { get; set; }

        /// <summary>
        /// The ID or full path of the project for the custom issue tracker.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URL to the project in the external issue tracker.
        /// </summary>
        [Input("projectUrl")]
        public Input<string>? ProjectUrl { get; set; }

        /// <summary>
        /// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// The ISO8601 date/time that this integration was last updated at in UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ServiceCustomIssueTrackerState()
        {
        }
        public static new ServiceCustomIssueTrackerState Empty => new ServiceCustomIssueTrackerState();
    }
}
