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
    /// **NOTE**: requires either EE (self-hosted) or Silver and above (GitLab.com).
    /// 
    /// This resource manages a [GitHub integration](https://docs.gitlab.com/ee/user/project/integrations/github.html) that updates pipeline statuses on a GitHub repo's pull requests.
    /// </summary>
    public partial class ServiceGithub : Pulumi.CustomResource
    {
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("repositoryUrl")]
        public Output<string> RepositoryUrl { get; private set; } = null!;

        /// <summary>
        /// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more. 
        /// </summary>
        [Output("staticContext")]
        public Output<bool?> StaticContext { get; private set; } = null!;

        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// A GitHub personal access token with at least `repo:status` scope.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceGithub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceGithub(string name, ServiceGithubArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceGithub:ServiceGithub", name, args ?? new ServiceGithubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceGithub(string name, Input<string> id, ServiceGithubState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/serviceGithub:ServiceGithub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceGithub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceGithub Get(string name, Input<string> id, ServiceGithubState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceGithub(name, id, state, options);
        }
    }

    public sealed class ServiceGithubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("repositoryUrl", required: true)]
        public Input<string> RepositoryUrl { get; set; } = null!;

        /// <summary>
        /// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more. 
        /// </summary>
        [Input("staticContext")]
        public Input<bool>? StaticContext { get; set; }

        /// <summary>
        /// A GitHub personal access token with at least `repo:status` scope.
        /// </summary>
        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        public ServiceGithubArgs()
        {
        }
    }

    public sealed class ServiceGithubState : Pulumi.ResourceArgs
    {
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// ID of the project you want to activate integration on.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("repositoryUrl")]
        public Input<string>? RepositoryUrl { get; set; }

        /// <summary>
        /// Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more. 
        /// </summary>
        [Input("staticContext")]
        public Input<bool>? StaticContext { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// A GitHub personal access token with at least `repo:status` scope.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ServiceGithubState()
        {
        }
    }
}