// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProjectProtectedBranches
    {
        /// <summary>
        /// The `gitlab.getProjectProtectedBranches` data source allows details of the protected branches of a given project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/protected_branches/#list-protected-branches)
        /// </summary>
        public static Task<GetProjectProtectedBranchesResult> InvokeAsync(GetProjectProtectedBranchesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectProtectedBranchesResult>("gitlab:index/getProjectProtectedBranches:getProjectProtectedBranches", args ?? new GetProjectProtectedBranchesArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectProtectedBranches` data source allows details of the protected branches of a given project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/protected_branches/#list-protected-branches)
        /// </summary>
        public static Output<GetProjectProtectedBranchesResult> Invoke(GetProjectProtectedBranchesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectProtectedBranchesResult>("gitlab:index/getProjectProtectedBranches:getProjectProtectedBranches", args ?? new GetProjectProtectedBranchesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getProjectProtectedBranches` data source allows details of the protected branches of a given project.
        /// 
        /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/protected_branches/#list-protected-branches)
        /// </summary>
        public static Output<GetProjectProtectedBranchesResult> Invoke(GetProjectProtectedBranchesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectProtectedBranchesResult>("gitlab:index/getProjectProtectedBranches:getProjectProtectedBranches", args ?? new GetProjectProtectedBranchesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectProtectedBranchesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        [Input("protectedBranches")]
        private List<Inputs.GetProjectProtectedBranchesProtectedBranchArgs>? _protectedBranches;

        /// <summary>
        /// A list of protected branches, as defined below.
        /// </summary>
        public List<Inputs.GetProjectProtectedBranchesProtectedBranchArgs> ProtectedBranches
        {
            get => _protectedBranches ?? (_protectedBranches = new List<Inputs.GetProjectProtectedBranchesProtectedBranchArgs>());
            set => _protectedBranches = value;
        }

        public GetProjectProtectedBranchesArgs()
        {
        }
        public static new GetProjectProtectedBranchesArgs Empty => new GetProjectProtectedBranchesArgs();
    }

    public sealed class GetProjectProtectedBranchesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("protectedBranches")]
        private InputList<Inputs.GetProjectProtectedBranchesProtectedBranchInputArgs>? _protectedBranches;

        /// <summary>
        /// A list of protected branches, as defined below.
        /// </summary>
        public InputList<Inputs.GetProjectProtectedBranchesProtectedBranchInputArgs> ProtectedBranches
        {
            get => _protectedBranches ?? (_protectedBranches = new InputList<Inputs.GetProjectProtectedBranchesProtectedBranchInputArgs>());
            set => _protectedBranches = value;
        }

        public GetProjectProtectedBranchesInvokeArgs()
        {
        }
        public static new GetProjectProtectedBranchesInvokeArgs Empty => new GetProjectProtectedBranchesInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectProtectedBranchesResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// A list of protected branches, as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectProtectedBranchesProtectedBranchResult> ProtectedBranches;

        [OutputConstructor]
        private GetProjectProtectedBranchesResult(
            int id,

            string projectId,

            ImmutableArray<Outputs.GetProjectProtectedBranchesProtectedBranchResult> protectedBranches)
        {
            Id = id;
            ProjectId = projectId;
            ProtectedBranches = protectedBranches;
        }
    }
}
