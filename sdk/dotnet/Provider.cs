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
    /// The provider type for the gitlab package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [GitLabResourceType("pulumi:providers:gitlab")]
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
        /// Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
        /// the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        /// </summary>
        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
        /// CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        /// </summary>
        [Output("cacertFile")]
        public Output<string?> CacertFile { get; private set; } = null!;

        /// <summary>
        /// File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        /// </summary>
        [Output("clientCert")]
        public Output<string?> ClientCert { get; private set; } = null!;

        /// <summary>
        /// File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
        /// `client_cert` is set.
        /// </summary>
        [Output("clientKey")]
        public Output<string?> ClientKey { get; private set; } = null!;

        /// <summary>
        /// The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
        /// used in this provider for authentication (using Bearer authorization token). See
        /// https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
        /// variable.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("gitlab", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
        /// Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
        /// the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
        /// CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
        /// </summary>
        [Input("cacertFile")]
        public Input<string>? CacertFile { get; set; }

        /// <summary>
        /// File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
        /// `client_cert` is set.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// (Experimental) By default the provider does a dummy request to get the current user in order to verify that the provider
        /// configuration is correct and the GitLab API is reachable. Turn it off, to skip this check. This may be useful if the
        /// GitLab instance does not yet exist and is created within the same terraform module. This is an experimental feature and
        /// may change in the future. Please make sure to always keep backups of your state.
        /// </summary>
        [Input("earlyAuthCheck", json: true)]
        public Input<bool>? EarlyAuthCheck { get; set; }

        /// <summary>
        /// When set to true this disables SSL verification of the connection to the GitLab instance.
        /// </summary>
        [Input("insecure", json: true)]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
        /// used in this provider for authentication (using Bearer authorization token). See
        /// https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
        /// variable.
        /// </summary>
        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        public ProviderArgs()
        {
        }
    }
}
