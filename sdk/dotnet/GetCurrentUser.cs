// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetCurrentUser
    {
        /// <summary>
        /// The `gitlab.getCurrentUser` data source allows details of the current user (determined by `token` provider attribute) to be retrieved.
        /// 
        /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#querycurrentuser)
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
        ///     var example = GitLab.GetCurrentUser.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCurrentUserResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCurrentUserResult>("gitlab:index/getCurrentUser:getCurrentUser", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// The `gitlab.getCurrentUser` data source allows details of the current user (determined by `token` provider attribute) to be retrieved.
        /// 
        /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#querycurrentuser)
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
        ///     var example = GitLab.GetCurrentUser.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCurrentUserResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCurrentUserResult>("gitlab:index/getCurrentUser:getCurrentUser", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// The `gitlab.getCurrentUser` data source allows details of the current user (determined by `token` provider attribute) to be retrieved.
        /// 
        /// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/index.html#querycurrentuser)
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
        ///     var example = GitLab.GetCurrentUser.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCurrentUserResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCurrentUserResult>("gitlab:index/getCurrentUser:getCurrentUser", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetCurrentUserResult
    {
        /// <summary>
        /// Indicates if the user is a bot.
        /// </summary>
        public readonly bool Bot;
        /// <summary>
        /// Global ID of the user. This is in the form of a GraphQL globally unique ID.
        /// </summary>
        public readonly string GlobalId;
        /// <summary>
        /// Personal namespace of the user. This is in the form of a GraphQL globally unique ID.
        /// </summary>
        public readonly string GlobalNamespaceId;
        /// <summary>
        /// Group count for the user.
        /// </summary>
        public readonly int GroupCount;
        /// <summary>
        /// ID of the user.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Human-readable name of the user. Returns **** if the user is a project bot and the requester does not have permission to view the project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Personal namespace of the user.
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// User’s public email.
        /// </summary>
        public readonly string PublicEmail;
        /// <summary>
        /// Username of the user. Unique within this instance of GitLab.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetCurrentUserResult(
            bool bot,

            string globalId,

            string globalNamespaceId,

            int groupCount,

            string id,

            string name,

            string namespaceId,

            string publicEmail,

            string username)
        {
            Bot = bot;
            GlobalId = globalId;
            GlobalNamespaceId = globalNamespaceId;
            GroupCount = groupCount;
            Id = id;
            Name = name;
            NamespaceId = namespaceId;
            PublicEmail = publicEmail;
            Username = username;
        }
    }
}
