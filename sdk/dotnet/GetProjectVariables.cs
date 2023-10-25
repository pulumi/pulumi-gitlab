// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectVariables
    {
        /// <summary>
        /// The `gitlab.getProjectVariables` data source allows to retrieve all project-level CI/CD variables.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
        /// </summary>
        public static Task<GetProjectVariablesResult> InvokeAsync(GetProjectVariablesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectVariablesResult>("gitlab:index/getProjectVariables:getProjectVariables", args ?? new GetProjectVariablesArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectVariables` data source allows to retrieve all project-level CI/CD variables.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
        /// </summary>
        public static Output<GetProjectVariablesResult> Invoke(GetProjectVariablesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectVariablesResult>("gitlab:index/getProjectVariables:getProjectVariables", args ?? new GetProjectVariablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectVariablesArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentScope")]
        public string? EnvironmentScope { get; set; }

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetProjectVariablesArgs()
        {
        }
        public static new GetProjectVariablesArgs Empty => new GetProjectVariablesArgs();
    }

    public sealed class GetProjectVariablesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetProjectVariablesInvokeArgs()
        {
        }
        public static new GetProjectVariablesInvokeArgs Empty => new GetProjectVariablesInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectVariablesResult
    {
        /// <summary>
        /// The environment scope of the variable. Defaults to all environment (`*`).
        /// </summary>
        public readonly string? EnvironmentScope;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name or id of the project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The list of variables returned by the search
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectVariablesVariableResult> Variables;

        [OutputConstructor]
        private GetProjectVariablesResult(
            string? environmentScope,

            string id,

            string project,

            ImmutableArray<Outputs.GetProjectVariablesVariableResult> variables)
        {
            EnvironmentScope = environmentScope;
            Id = id;
            Project = project;
            Variables = variables;
        }
    }
}
