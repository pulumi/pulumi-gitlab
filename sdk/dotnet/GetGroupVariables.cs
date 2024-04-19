// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroupVariables
    {
        /// <summary>
        /// The `gitlab.getGroupVariables` data source allows to retrieve all group-level CI/CD variables.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
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
        ///     var vars = GitLab.GetGroupVariables.Invoke(new()
        ///     {
        ///         Group = "my/example/group",
        ///     });
        /// 
        ///     var stagingVars = GitLab.GetGroupVariables.Invoke(new()
        ///     {
        ///         EnvironmentScope = "staging/*",
        ///         Group = "my/example/group",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGroupVariablesResult> InvokeAsync(GetGroupVariablesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupVariablesResult>("gitlab:index/getGroupVariables:getGroupVariables", args ?? new GetGroupVariablesArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroupVariables` data source allows to retrieve all group-level CI/CD variables.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)
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
        ///     var vars = GitLab.GetGroupVariables.Invoke(new()
        ///     {
        ///         Group = "my/example/group",
        ///     });
        /// 
        ///     var stagingVars = GitLab.GetGroupVariables.Invoke(new()
        ///     {
        ///         EnvironmentScope = "staging/*",
        ///         Group = "my/example/group",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupVariablesResult> Invoke(GetGroupVariablesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupVariablesResult>("gitlab:index/getGroupVariables:getGroupVariables", args ?? new GetGroupVariablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupVariablesArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentScope")]
        public string? EnvironmentScope { get; set; }

        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        public GetGroupVariablesArgs()
        {
        }
        public static new GetGroupVariablesArgs Empty => new GetGroupVariablesArgs();
    }

    public sealed class GetGroupVariablesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        public GetGroupVariablesInvokeArgs()
        {
        }
        public static new GetGroupVariablesInvokeArgs Empty => new GetGroupVariablesInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupVariablesResult
    {
        /// <summary>
        /// The environment scope of the variable. Defaults to all environment (`*`).
        /// </summary>
        public readonly string? EnvironmentScope;
        /// <summary>
        /// The name or id of the group.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of variables returned by the search
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupVariablesVariableResult> Variables;

        [OutputConstructor]
        private GetGroupVariablesResult(
            string? environmentScope,

            string group,

            string id,

            ImmutableArray<Outputs.GetGroupVariablesVariableResult> variables)
        {
            EnvironmentScope = environmentScope;
            Group = group;
            Id = id;
            Variables = variables;
        }
    }
}
