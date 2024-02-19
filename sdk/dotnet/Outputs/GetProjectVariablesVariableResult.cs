// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Outputs
{

    [OutputType]
    public sealed class GetProjectVariablesVariableResult
    {
        /// <summary>
        /// The description of the variable.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
        /// </summary>
        public readonly string EnvironmentScope;
        /// <summary>
        /// The name of the variable.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
        /// </summary>
        public readonly bool Masked;
        /// <summary>
        /// The name or id of the project.
        /// </summary>
        public readonly string Project;
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
        private GetProjectVariablesVariableResult(
            string description,

            string environmentScope,

            string key,

            bool masked,

            string project,

            bool @protected,

            bool raw,

            string value,

            string variableType)
        {
            Description = description;
            EnvironmentScope = environmentScope;
            Key = key;
            Masked = masked;
            Project = project;
            Protected = @protected;
            Raw = raw;
            Value = value;
            VariableType = variableType;
        }
    }
}
