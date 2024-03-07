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
    /// The `gitlab.Runner` resource allows to manage the lifecycle of a runner.
    /// 
    /// A runner can either be registered at an instance level or group level.
    /// The runner will be registered at a group level if the token used is from a group, or at an instance level if the token used is for the instance.
    /// 
    /// ~ &gt; Using this resource will register a runner using the deprecated `registration_token` flow. To use the new `authentication_token` flow instead,
    /// use the `gitlab.UserRunner` resource!
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#register-a-new-runner)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// using Local = Pulumi.Local;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Basic GitLab Group Runner
    ///     var myGroup = new GitLab.Group("myGroup", new()
    ///     {
    ///         Description = "group that holds the runners",
    ///     });
    /// 
    ///     var basicRunner = new GitLab.Runner("basicRunner", new()
    ///     {
    ///         RegistrationToken = myGroup.RunnersToken,
    ///     });
    /// 
    ///     // GitLab Runner that runs only tagged jobs
    ///     var taggedOnly = new GitLab.Runner("taggedOnly", new()
    ///     {
    ///         RegistrationToken = myGroup.RunnersToken,
    ///         Description = "I only run tagged jobs",
    ///         RunUntagged = false,
    ///         TagLists = new[]
    ///         {
    ///             "tag_one",
    ///             "tag_two",
    ///         },
    ///     });
    /// 
    ///     // GitLab Runner that only runs on protected branches
    ///     var @protected = new GitLab.Runner("protected", new()
    ///     {
    ///         RegistrationToken = myGroup.RunnersToken,
    ///         Description = "I only run protected jobs",
    ///         AccessLevel = "ref_protected",
    ///     });
    /// 
    ///     // Generate a `config.toml` file that you can use to create a runner
    ///     // This is the typical workflow for this resource, using it to create an authentication_token which can then be used
    ///     // to generate the `config.toml` file to prevent re-registering the runner every time new hardware is created.
    ///     var myCustomGroup = new GitLab.Group("myCustomGroup", new()
    ///     {
    ///         Description = "group that holds the custom runners",
    ///     });
    /// 
    ///     var myRunner = new GitLab.Runner("myRunner", new()
    ///     {
    ///         RegistrationToken = myCustomGroup.RunnersToken,
    ///     });
    /// 
    ///     // This creates a configuration for a local "shell" runner, but can be changed to generate whatever is needed.
    ///     // Place this configuration file on a server at `/etc/gitlab-runner/config.toml`, then run `gitlab-runner start`.
    ///     // See https://docs.gitlab.com/runner/configuration/advanced-configuration.html for more information.
    ///     var config = new Local.File("config", new()
    ///     {
    ///         Filename = $"{path.Module}/config.toml",
    ///         Content = myRunner.AuthenticationToken.Apply(authenticationToken =&gt; @$"  concurrent = 1
    /// 
    ///   [[runners]]
    ///     name = ""Hello Terraform""
    ///     url = ""https://example.gitlab.com/""
    ///     token = ""{authenticationToken}""
    ///     executor = ""shell""
    ///     
    /// "),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// A GitLab Runner can be imported using the runner's ID, eg
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/runner:Runner this 1
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/runner:Runner")]
    public partial class Runner : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// The authentication token used for building a config.toml file. This value is not present when imported.
        /// </summary>
        [Output("authenticationToken")]
        public Output<string> AuthenticationToken { get; private set; } = null!;

        /// <summary>
        /// The runner's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the runner should be locked for current project.
        /// </summary>
        [Output("locked")]
        public Output<bool> Locked { get; private set; } = null!;

        /// <summary>
        /// Maximum timeout set when this runner handles the job.
        /// </summary>
        [Output("maximumTimeout")]
        public Output<int?> MaximumTimeout { get; private set; } = null!;

        /// <summary>
        /// Whether the runner should ignore new jobs.
        /// </summary>
        [Output("paused")]
        public Output<bool> Paused { get; private set; } = null!;

        /// <summary>
        /// The registration token used to register the runner.
        /// </summary>
        [Output("registrationToken")]
        public Output<string> RegistrationToken { get; private set; } = null!;

        /// <summary>
        /// Whether the runner should handle untagged jobs.
        /// </summary>
        [Output("runUntagged")]
        public Output<bool> RunUntagged { get; private set; } = null!;

        /// <summary>
        /// The status of runners to show, one of: online and offline. active and paused are also possible values
        /// 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// List of runner’s tags.
        /// </summary>
        [Output("tagLists")]
        public Output<ImmutableArray<string>> TagLists { get; private set; } = null!;


        /// <summary>
        /// Create a Runner resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Runner(string name, RunnerArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/runner:Runner", name, args ?? new RunnerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Runner(string name, Input<string> id, RunnerState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/runner:Runner", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "authenticationToken",
                    "registrationToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Runner resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Runner Get(string name, Input<string> id, RunnerState? state = null, CustomResourceOptions? options = null)
        {
            return new Runner(name, id, state, options);
        }
    }

    public sealed class RunnerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The runner's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the runner should be locked for current project.
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// Maximum timeout set when this runner handles the job.
        /// </summary>
        [Input("maximumTimeout")]
        public Input<int>? MaximumTimeout { get; set; }

        /// <summary>
        /// Whether the runner should ignore new jobs.
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        [Input("registrationToken", required: true)]
        private Input<string>? _registrationToken;

        /// <summary>
        /// The registration token used to register the runner.
        /// </summary>
        public Input<string>? RegistrationToken
        {
            get => _registrationToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _registrationToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Whether the runner should handle untagged jobs.
        /// </summary>
        [Input("runUntagged")]
        public Input<bool>? RunUntagged { get; set; }

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// List of runner’s tags.
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        public RunnerArgs()
        {
        }
        public static new RunnerArgs Empty => new RunnerArgs();
    }

    public sealed class RunnerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access_level of the runner. Valid values are: `not_protected`, `ref_protected`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        [Input("authenticationToken")]
        private Input<string>? _authenticationToken;

        /// <summary>
        /// The authentication token used for building a config.toml file. This value is not present when imported.
        /// </summary>
        public Input<string>? AuthenticationToken
        {
            get => _authenticationToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authenticationToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The runner's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the runner should be locked for current project.
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// Maximum timeout set when this runner handles the job.
        /// </summary>
        [Input("maximumTimeout")]
        public Input<int>? MaximumTimeout { get; set; }

        /// <summary>
        /// Whether the runner should ignore new jobs.
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        [Input("registrationToken")]
        private Input<string>? _registrationToken;

        /// <summary>
        /// The registration token used to register the runner.
        /// </summary>
        public Input<string>? RegistrationToken
        {
            get => _registrationToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _registrationToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Whether the runner should handle untagged jobs.
        /// </summary>
        [Input("runUntagged")]
        public Input<bool>? RunUntagged { get; set; }

        /// <summary>
        /// The status of runners to show, one of: online and offline. active and paused are also possible values
        /// 			              which were deprecated in GitLab 14.8 and will be removed in GitLab 16.0.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// List of runner’s tags.
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        public RunnerState()
        {
        }
        public static new RunnerState Empty => new RunnerState();
    }
}
