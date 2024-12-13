// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetGroupIds
    {
        /// <summary>
        /// The `gitlab.getGroupIds` data source identification information for a given group, allowing a user to translate a full path or ID into the GraphQL ID of the group.
        /// 
        /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#querygroup)
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
        ///     var newGroup = new GitLab.Group("new_group");
        /// 
        ///     // use group IDs to get additional information, such as the GraphQL ID
        ///     // for other resources
        ///     var foo = GitLab.GetGroupIds.Invoke(new()
        ///     {
        ///         Group = "gitlab_group.new_group.id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["graphQLId"] = foo.Apply(getGroupIdsResult =&gt; getGroupIdsResult.GroupGraphqlId),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetGroupIdsResult> InvokeAsync(GetGroupIdsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupIdsResult>("gitlab:index/getGroupIds:getGroupIds", args ?? new GetGroupIdsArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroupIds` data source identification information for a given group, allowing a user to translate a full path or ID into the GraphQL ID of the group.
        /// 
        /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#querygroup)
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
        ///     var newGroup = new GitLab.Group("new_group");
        /// 
        ///     // use group IDs to get additional information, such as the GraphQL ID
        ///     // for other resources
        ///     var foo = GitLab.GetGroupIds.Invoke(new()
        ///     {
        ///         Group = "gitlab_group.new_group.id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["graphQLId"] = foo.Apply(getGroupIdsResult =&gt; getGroupIdsResult.GroupGraphqlId),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupIdsResult> Invoke(GetGroupIdsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupIdsResult>("gitlab:index/getGroupIds:getGroupIds", args ?? new GetGroupIdsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `gitlab.getGroupIds` data source identification information for a given group, allowing a user to translate a full path or ID into the GraphQL ID of the group.
        /// 
        /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#querygroup)
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
        ///     var newGroup = new GitLab.Group("new_group");
        /// 
        ///     // use group IDs to get additional information, such as the GraphQL ID
        ///     // for other resources
        ///     var foo = GitLab.GetGroupIds.Invoke(new()
        ///     {
        ///         Group = "gitlab_group.new_group.id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["graphQLId"] = foo.Apply(getGroupIdsResult =&gt; getGroupIdsResult.GroupGraphqlId),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupIdsResult> Invoke(GetGroupIdsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupIdsResult>("gitlab:index/getGroupIds:getGroupIds", args ?? new GetGroupIdsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupIdsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the group.
        /// </summary>
        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        public GetGroupIdsArgs()
        {
        }
        public static new GetGroupIdsArgs Empty => new GetGroupIdsArgs();
    }

    public sealed class GetGroupIdsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID or URL-encoded path of the group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        public GetGroupIdsInvokeArgs()
        {
        }
        public static new GetGroupIdsInvokeArgs Empty => new GetGroupIdsInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupIdsResult
    {
        /// <summary>
        /// The ID or URL-encoded path of the group.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The full path of the group.
        /// </summary>
        public readonly string GroupFullPath;
        /// <summary>
        /// The GraphQL ID of the group.
        /// </summary>
        public readonly string GroupGraphqlId;
        /// <summary>
        /// The ID of the group.
        /// </summary>
        public readonly string GroupId;
        public readonly string Id;

        [OutputConstructor]
        private GetGroupIdsResult(
            string group,

            string groupFullPath,

            string groupGraphqlId,

            string groupId,

            string id)
        {
            Group = group;
            GroupFullPath = groupFullPath;
            GroupGraphqlId = groupGraphqlId;
            GroupId = groupId;
            Id = id;
        }
    }
}
