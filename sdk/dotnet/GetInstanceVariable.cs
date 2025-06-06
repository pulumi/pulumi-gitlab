// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetInstanceVariable
    {
        /// <summary>
        /// The `gitlab.InstanceVariable` data source allows to retrieve details about an instance-level CI/CD variable.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/instance_level_ci_variables/)
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
        ///     var foo = GitLab.GetInstanceVariable.Invoke(new()
        ///     {
        ///         Key = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetInstanceVariableResult> InvokeAsync(GetInstanceVariableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceVariableResult>("gitlab:index/getInstanceVariable:getInstanceVariable", args ?? new GetInstanceVariableArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.InstanceVariable` data source allows to retrieve details about an instance-level CI/CD variable.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/instance_level_ci_variables/)
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
        ///     var foo = GitLab.GetInstanceVariable.Invoke(new()
        ///     {
        ///         Key = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstanceVariableResult> Invoke(GetInstanceVariableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceVariableResult>("gitlab:index/getInstanceVariable:getInstanceVariable", args ?? new GetInstanceVariableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.InstanceVariable` data source allows to retrieve details about an instance-level CI/CD variable.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/instance_level_ci_variables/)
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
        ///     var foo = GitLab.GetInstanceVariable.Invoke(new()
        ///     {
        ///         Key = "foo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstanceVariableResult> Invoke(GetInstanceVariableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceVariableResult>("gitlab:index/getInstanceVariable:getInstanceVariable", args ?? new GetInstanceVariableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceVariableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        public GetInstanceVariableArgs()
        {
        }
        public static new GetInstanceVariableArgs Empty => new GetInstanceVariableArgs();
    }

    public sealed class GetInstanceVariableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the variable.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public GetInstanceVariableInvokeArgs()
        {
        }
        public static new GetInstanceVariableInvokeArgs Empty => new GetInstanceVariableInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceVariableResult
    {
        /// <summary>
        /// The description of the variable. Maximum of 255 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the variable.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#masked-variables). Defaults to `false`.
        /// </summary>
        public readonly bool Masked;
        /// <summary>
        /// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
        /// </summary>
        public readonly bool Protected;
        /// <summary>
        /// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
        /// </summary>
        public readonly bool Raw;
        /// <summary>
        /// The value of the variable.
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
        /// </summary>
        public readonly string VariableType;

        [OutputConstructor]
        private GetInstanceVariableResult(
            string description,

            string id,

            string key,

            bool masked,

            bool @protected,

            bool raw,

            string value,

            string variableType)
        {
            Description = description;
            Id = id;
            Key = key;
            Masked = masked;
            Protected = @protected;
            Raw = raw;
            Value = value;
            VariableType = variableType;
        }
    }
}
