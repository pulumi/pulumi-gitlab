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
    /// The `gitlab.IntegrationExternalWiki` resource allows to manage the lifecycle of a project integration with External Wiki Service.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#external-wiki)
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
    ///     var wiki = new GitLab.IntegrationExternalWiki("wiki", new()
    ///     {
    ///         Project = awesomeProject.Id,
    ///         ExternalWikiUrl = "https://MyAwesomeExternalWikiURL.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_external_wiki`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_integration_external_wiki.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// You can import a gitlab_integration_external_wiki state using the project ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/integrationExternalWiki:IntegrationExternalWiki wiki 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/integrationExternalWiki:IntegrationExternalWiki")]
    public partial class IntegrationExternalWiki : global::Pulumi.CustomResource
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
        /// The URL of the external wiki.
        /// </summary>
        [Output("externalWikiUrl")]
        public Output<string> ExternalWikiUrl { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

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
        /// Create a IntegrationExternalWiki resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationExternalWiki(string name, IntegrationExternalWikiArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationExternalWiki:IntegrationExternalWiki", name, args ?? new IntegrationExternalWikiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationExternalWiki(string name, Input<string> id, IntegrationExternalWikiState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/integrationExternalWiki:IntegrationExternalWiki", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationExternalWiki resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationExternalWiki Get(string name, Input<string> id, IntegrationExternalWikiState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationExternalWiki(name, id, state, options);
        }
    }

    public sealed class IntegrationExternalWikiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the external wiki.
        /// </summary>
        [Input("externalWikiUrl", required: true)]
        public Input<string> ExternalWikiUrl { get; set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public IntegrationExternalWikiArgs()
        {
        }
        public static new IntegrationExternalWikiArgs Empty => new IntegrationExternalWikiArgs();
    }

    public sealed class IntegrationExternalWikiState : global::Pulumi.ResourceArgs
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
        /// The URL of the external wiki.
        /// </summary>
        [Input("externalWikiUrl")]
        public Input<string>? ExternalWikiUrl { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

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

        public IntegrationExternalWikiState()
        {
        }
        public static new IntegrationExternalWikiState Empty => new IntegrationExternalWikiState();
    }
}
