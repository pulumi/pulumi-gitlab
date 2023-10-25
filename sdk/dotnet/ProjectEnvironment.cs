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
    /// ## Import
    /// 
    /// GitLab project environments can be imported using an id made up of `projectId:environmenId`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectEnvironment:ProjectEnvironment bar 123:321
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectEnvironment:ProjectEnvironment")]
    public partial class ProjectEnvironment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ISO8601 date/time that this environment was created at in UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Place to link to for this environment.
        /// </summary>
        [Output("externalUrl")]
        public Output<string?> ExternalUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID or full path of the project to environment is created for.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// State the environment is in. Valid values are `available`, `stopped`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Determines whether the environment is attempted to be stopped before the environment is deleted.
        /// </summary>
        [Output("stopBeforeDestroy")]
        public Output<bool?> StopBeforeDestroy { get; private set; } = null!;

        /// <summary>
        /// The ISO8601 date/time that this environment was last updated at in UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectEnvironment(string name, ProjectEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectEnvironment:ProjectEnvironment", name, args ?? new ProjectEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectEnvironment(string name, Input<string> id, ProjectEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectEnvironment:ProjectEnvironment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectEnvironment Get(string name, Input<string> id, ProjectEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectEnvironment(name, id, state, options);
        }
    }

    public sealed class ProjectEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Place to link to for this environment.
        /// </summary>
        [Input("externalUrl")]
        public Input<string>? ExternalUrl { get; set; }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project to environment is created for.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Determines whether the environment is attempted to be stopped before the environment is deleted.
        /// </summary>
        [Input("stopBeforeDestroy")]
        public Input<bool>? StopBeforeDestroy { get; set; }

        public ProjectEnvironmentArgs()
        {
        }
        public static new ProjectEnvironmentArgs Empty => new ProjectEnvironmentArgs();
    }

    public sealed class ProjectEnvironmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO8601 date/time that this environment was created at in UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Place to link to for this environment.
        /// </summary>
        [Input("externalUrl")]
        public Input<string>? ExternalUrl { get; set; }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID or full path of the project to environment is created for.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// State the environment is in. Valid values are `available`, `stopped`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Determines whether the environment is attempted to be stopped before the environment is deleted.
        /// </summary>
        [Input("stopBeforeDestroy")]
        public Input<bool>? StopBeforeDestroy { get; set; }

        /// <summary>
        /// The ISO8601 date/time that this environment was last updated at in UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ProjectEnvironmentState()
        {
        }
        public static new ProjectEnvironmentState Empty => new ProjectEnvironmentState();
    }
}
