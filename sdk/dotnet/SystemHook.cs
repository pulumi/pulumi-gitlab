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
    /// The `gitlab.SystemHook` resource allows to manage the lifecycle of a system hook.
    /// 
    /// &gt; This resource requires GitLab 14.9
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/system_hooks.html)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new GitLab.SystemHook("example", new()
    ///     {
    ///         Url = "https://example.com/hook-%d",
    ///         Token = "secret-token",
    ///         PushEvents = true,
    ///         TagPushEvents = true,
    ///         MergeRequestsEvents = true,
    ///         RepositoryUpdateEvents = true,
    ///         EnableSslVerification = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// You can import a system hook using the hook id `{hook-id}`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/systemHook:SystemHook example 42
    /// ```
    /// 
    /// NOTE: the `token` attribute won't be available for imported resources.
    /// </summary>
    [GitLabResourceType("gitlab:index/systemHook:SystemHook")]
    public partial class SystemHook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time the hook was created in ISO8601 format.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Do SSL verification when triggering the hook.
        /// </summary>
        [Output("enableSslVerification")]
        public Output<bool?> EnableSslVerification { get; private set; } = null!;

        /// <summary>
        /// Trigger hook on merge requests events.
        /// </summary>
        [Output("mergeRequestsEvents")]
        public Output<bool?> MergeRequestsEvents { get; private set; } = null!;

        /// <summary>
        /// When true, the hook fires on push events.
        /// </summary>
        [Output("pushEvents")]
        public Output<bool?> PushEvents { get; private set; } = null!;

        /// <summary>
        /// Trigger hook on repository update events.
        /// </summary>
        [Output("repositoryUpdateEvents")]
        public Output<bool?> RepositoryUpdateEvents { get; private set; } = null!;

        /// <summary>
        /// When true, the hook fires on new tags being pushed.
        /// </summary>
        [Output("tagPushEvents")]
        public Output<bool?> TagPushEvents { get; private set; } = null!;

        /// <summary>
        /// Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// The hook URL.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a SystemHook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemHook(string name, SystemHookArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/systemHook:SystemHook", name, args ?? new SystemHookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemHook(string name, Input<string> id, SystemHookState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/systemHook:SystemHook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SystemHook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemHook Get(string name, Input<string> id, SystemHookState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemHook(name, id, state, options);
        }
    }

    public sealed class SystemHookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Do SSL verification when triggering the hook.
        /// </summary>
        [Input("enableSslVerification")]
        public Input<bool>? EnableSslVerification { get; set; }

        /// <summary>
        /// Trigger hook on merge requests events.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// When true, the hook fires on push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Trigger hook on repository update events.
        /// </summary>
        [Input("repositoryUpdateEvents")]
        public Input<bool>? RepositoryUpdateEvents { get; set; }

        /// <summary>
        /// When true, the hook fires on new tags being pushed.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The hook URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public SystemHookArgs()
        {
        }
        public static new SystemHookArgs Empty => new SystemHookArgs();
    }

    public sealed class SystemHookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time the hook was created in ISO8601 format.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Do SSL verification when triggering the hook.
        /// </summary>
        [Input("enableSslVerification")]
        public Input<bool>? EnableSslVerification { get; set; }

        /// <summary>
        /// Trigger hook on merge requests events.
        /// </summary>
        [Input("mergeRequestsEvents")]
        public Input<bool>? MergeRequestsEvents { get; set; }

        /// <summary>
        /// When true, the hook fires on push events.
        /// </summary>
        [Input("pushEvents")]
        public Input<bool>? PushEvents { get; set; }

        /// <summary>
        /// Trigger hook on repository update events.
        /// </summary>
        [Input("repositoryUpdateEvents")]
        public Input<bool>? RepositoryUpdateEvents { get; set; }

        /// <summary>
        /// When true, the hook fires on new tags being pushed.
        /// </summary>
        [Input("tagPushEvents")]
        public Input<bool>? TagPushEvents { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Secret token to validate received payloads; this isn’t returned in the response. This attribute is not available for imported resources.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The hook URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public SystemHookState()
        {
        }
        public static new SystemHookState Empty => new SystemHookState();
    }
}
