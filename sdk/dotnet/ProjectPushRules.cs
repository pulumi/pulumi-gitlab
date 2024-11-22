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
    ///     var sample = new GitLab.ProjectPushRules("sample", new()
    ///     {
    ///         Project = "42",
    ///         AuthorEmailRegex = "@gitlab.com$",
    ///         BranchNameRegex = "(feat|fix)\\/*",
    ///         CommitCommitterCheck = true,
    ///         CommitCommitterNameCheck = true,
    ///         CommitMessageNegativeRegex = "ssh\\:\\/\\/",
    ///         CommitMessageRegex = "(feat|fix):.*",
    ///         DenyDeleteTag = false,
    ///         FileNameRegex = "(jar|exe)$",
    ///         MaxFileSize = 4,
    ///         MemberCheck = true,
    ///         PreventSecrets = true,
    ///         RejectUnsignedCommits = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_push_rules`. For example:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = gitlab_project_push_rules.example
    /// 
    ///   id = "see CLI command below for ID"
    /// 
    /// }
    /// 
    /// Import using the CLI is supported using the following syntax:
    /// 
    /// Gitlab project push rules can be imported with a key composed of `&lt;project_id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gitlab:index/projectPushRules:ProjectPushRules sample "42"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectPushRules:ProjectPushRules")]
    public partial class ProjectPushRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// All commit author emails must match this regex, e.g. `@my-company.com$`.
        /// </summary>
        [Output("authorEmailRegex")]
        public Output<string> AuthorEmailRegex { get; private set; } = null!;

        /// <summary>
        /// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        /// </summary>
        [Output("branchNameRegex")]
        public Output<string> BranchNameRegex { get; private set; } = null!;

        /// <summary>
        /// Users can only push commits to this repository that were committed with one of their own verified emails.
        /// </summary>
        [Output("commitCommitterCheck")]
        public Output<bool> CommitCommitterCheck { get; private set; } = null!;

        /// <summary>
        /// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
        /// </summary>
        [Output("commitCommitterNameCheck")]
        public Output<bool> CommitCommitterNameCheck { get; private set; } = null!;

        /// <summary>
        /// No commit message is allowed to match this regex, e.g. `ssh\:\/\/`.
        /// </summary>
        [Output("commitMessageNegativeRegex")]
        public Output<string> CommitMessageNegativeRegex { get; private set; } = null!;

        /// <summary>
        /// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        /// </summary>
        [Output("commitMessageRegex")]
        public Output<string> CommitMessageRegex { get; private set; } = null!;

        /// <summary>
        /// Deny deleting a tag.
        /// </summary>
        [Output("denyDeleteTag")]
        public Output<bool> DenyDeleteTag { get; private set; } = null!;

        /// <summary>
        /// All committed filenames must not match this regex, e.g. `(jar|exe)$`.
        /// </summary>
        [Output("fileNameRegex")]
        public Output<string> FileNameRegex { get; private set; } = null!;

        /// <summary>
        /// Maximum file size (MB).
        /// </summary>
        [Output("maxFileSize")]
        public Output<int> MaxFileSize { get; private set; } = null!;

        /// <summary>
        /// Restrict commits by author (email) to existing GitLab users.
        /// </summary>
        [Output("memberCheck")]
        public Output<bool> MemberCheck { get; private set; } = null!;

        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        [Output("preventSecrets")]
        public Output<bool> PreventSecrets { get; private set; } = null!;

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Reject commit when it’s not DCO certified.
        /// </summary>
        [Output("rejectNonDcoCommits")]
        public Output<bool> RejectNonDcoCommits { get; private set; } = null!;

        /// <summary>
        /// Reject commit when it’s not signed.
        /// </summary>
        [Output("rejectUnsignedCommits")]
        public Output<bool> RejectUnsignedCommits { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectPushRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectPushRules(string name, ProjectPushRulesArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectPushRules:ProjectPushRules", name, args ?? new ProjectPushRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectPushRules(string name, Input<string> id, ProjectPushRulesState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectPushRules:ProjectPushRules", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// Get an existing ProjectPushRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectPushRules Get(string name, Input<string> id, ProjectPushRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectPushRules(name, id, state, options);
        }
    }

    public sealed class ProjectPushRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// All commit author emails must match this regex, e.g. `@my-company.com$`.
        /// </summary>
        [Input("authorEmailRegex")]
        public Input<string>? AuthorEmailRegex { get; set; }

        /// <summary>
        /// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        /// </summary>
        [Input("branchNameRegex")]
        public Input<string>? BranchNameRegex { get; set; }

        /// <summary>
        /// Users can only push commits to this repository that were committed with one of their own verified emails.
        /// </summary>
        [Input("commitCommitterCheck")]
        public Input<bool>? CommitCommitterCheck { get; set; }

        /// <summary>
        /// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
        /// </summary>
        [Input("commitCommitterNameCheck")]
        public Input<bool>? CommitCommitterNameCheck { get; set; }

        /// <summary>
        /// No commit message is allowed to match this regex, e.g. `ssh\:\/\/`.
        /// </summary>
        [Input("commitMessageNegativeRegex")]
        public Input<string>? CommitMessageNegativeRegex { get; set; }

        /// <summary>
        /// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        /// </summary>
        [Input("commitMessageRegex")]
        public Input<string>? CommitMessageRegex { get; set; }

        /// <summary>
        /// Deny deleting a tag.
        /// </summary>
        [Input("denyDeleteTag")]
        public Input<bool>? DenyDeleteTag { get; set; }

        /// <summary>
        /// All committed filenames must not match this regex, e.g. `(jar|exe)$`.
        /// </summary>
        [Input("fileNameRegex")]
        public Input<string>? FileNameRegex { get; set; }

        /// <summary>
        /// Maximum file size (MB).
        /// </summary>
        [Input("maxFileSize")]
        public Input<int>? MaxFileSize { get; set; }

        /// <summary>
        /// Restrict commits by author (email) to existing GitLab users.
        /// </summary>
        [Input("memberCheck")]
        public Input<bool>? MemberCheck { get; set; }

        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        [Input("preventSecrets")]
        public Input<bool>? PreventSecrets { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Reject commit when it’s not DCO certified.
        /// </summary>
        [Input("rejectNonDcoCommits")]
        public Input<bool>? RejectNonDcoCommits { get; set; }

        /// <summary>
        /// Reject commit when it’s not signed.
        /// </summary>
        [Input("rejectUnsignedCommits")]
        public Input<bool>? RejectUnsignedCommits { get; set; }

        public ProjectPushRulesArgs()
        {
        }
        public static new ProjectPushRulesArgs Empty => new ProjectPushRulesArgs();
    }

    public sealed class ProjectPushRulesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// All commit author emails must match this regex, e.g. `@my-company.com$`.
        /// </summary>
        [Input("authorEmailRegex")]
        public Input<string>? AuthorEmailRegex { get; set; }

        /// <summary>
        /// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        /// </summary>
        [Input("branchNameRegex")]
        public Input<string>? BranchNameRegex { get; set; }

        /// <summary>
        /// Users can only push commits to this repository that were committed with one of their own verified emails.
        /// </summary>
        [Input("commitCommitterCheck")]
        public Input<bool>? CommitCommitterCheck { get; set; }

        /// <summary>
        /// Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
        /// </summary>
        [Input("commitCommitterNameCheck")]
        public Input<bool>? CommitCommitterNameCheck { get; set; }

        /// <summary>
        /// No commit message is allowed to match this regex, e.g. `ssh\:\/\/`.
        /// </summary>
        [Input("commitMessageNegativeRegex")]
        public Input<string>? CommitMessageNegativeRegex { get; set; }

        /// <summary>
        /// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        /// </summary>
        [Input("commitMessageRegex")]
        public Input<string>? CommitMessageRegex { get; set; }

        /// <summary>
        /// Deny deleting a tag.
        /// </summary>
        [Input("denyDeleteTag")]
        public Input<bool>? DenyDeleteTag { get; set; }

        /// <summary>
        /// All committed filenames must not match this regex, e.g. `(jar|exe)$`.
        /// </summary>
        [Input("fileNameRegex")]
        public Input<string>? FileNameRegex { get; set; }

        /// <summary>
        /// Maximum file size (MB).
        /// </summary>
        [Input("maxFileSize")]
        public Input<int>? MaxFileSize { get; set; }

        /// <summary>
        /// Restrict commits by author (email) to existing GitLab users.
        /// </summary>
        [Input("memberCheck")]
        public Input<bool>? MemberCheck { get; set; }

        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        [Input("preventSecrets")]
        public Input<bool>? PreventSecrets { get; set; }

        /// <summary>
        /// The ID or URL-encoded path of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Reject commit when it’s not DCO certified.
        /// </summary>
        [Input("rejectNonDcoCommits")]
        public Input<bool>? RejectNonDcoCommits { get; set; }

        /// <summary>
        /// Reject commit when it’s not signed.
        /// </summary>
        [Input("rejectUnsignedCommits")]
        public Input<bool>? RejectUnsignedCommits { get; set; }

        public ProjectPushRulesState()
        {
        }
        public static new ProjectPushRulesState Empty => new ProjectPushRulesState();
    }
}
