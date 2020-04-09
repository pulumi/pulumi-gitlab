// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides details about a specific group in the gitlab provider.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/group.html.markdown.
        /// </summary>
        [Obsolete("Use GetGroup.InvokeAsync() instead")]
        public static Task<GetGroupResult> GetGroup(GetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("gitlab:index/getGroup:getGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetGroup
    {
        /// <summary>
        /// Provides details about a specific group in the gitlab provider.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/group.html.markdown.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("gitlab:index/getGroup:getGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full path of the group.
        /// </summary>
        [Input("fullPath")]
        public string? FullPath { get; set; }

        /// <summary>
        /// The ID of the group.
        /// </summary>
        [Input("groupId")]
        public int? GroupId { get; set; }

        public GetGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The description of the group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The full name of the group.
        /// </summary>
        public readonly string FullName;
        /// <summary>
        /// The full path of the group.
        /// </summary>
        public readonly string FullPath;
        public readonly int GroupId;
        /// <summary>
        /// Boolean, is LFS enabled for projects in this group.
        /// </summary>
        public readonly bool LfsEnabled;
        /// <summary>
        /// The name of this group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Integer, ID of the parent group.
        /// </summary>
        public readonly int ParentId;
        /// <summary>
        /// The path of the group.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Boolean, is request for access enabled to the group.
        /// </summary>
        public readonly bool RequestAccessEnabled;
        /// <summary>
        /// The group level registration token to use during runner setup.
        /// </summary>
        public readonly string RunnersToken;
        /// <summary>
        /// Visibility level of the group. Possible values are `private`, `internal`, `public`.
        /// </summary>
        public readonly string VisibilityLevel;
        /// <summary>
        /// Web URL of the group.
        /// </summary>
        public readonly string WebUrl;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupResult(
            string description,
            string fullName,
            string fullPath,
            int groupId,
            bool lfsEnabled,
            string name,
            int parentId,
            string path,
            bool requestAccessEnabled,
            string runnersToken,
            string visibilityLevel,
            string webUrl,
            string id)
        {
            Description = description;
            FullName = fullName;
            FullPath = fullPath;
            GroupId = groupId;
            LfsEnabled = lfsEnabled;
            Name = name;
            ParentId = parentId;
            Path = path;
            RequestAccessEnabled = requestAccessEnabled;
            RunnersToken = runnersToken;
            VisibilityLevel = visibilityLevel;
            WebUrl = webUrl;
            Id = id;
        }
    }
}
