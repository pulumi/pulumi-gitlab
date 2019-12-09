// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// This resource allows you to protect a specific branch by an access level so that the user with less access level cannot Merge/Push to the branch. GitLab EE features to protect by group or user are not supported.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/branch_protection.html.markdown.
    /// </summary>
    public partial class BranchProtection : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the branch.
        /// </summary>
        [Output("branch")]
        public Output<string> Branch { get; private set; } = null!;

        /// <summary>
        /// One of five levels of access to the project.
        /// </summary>
        [Output("mergeAccessLevel")]
        public Output<string> MergeAccessLevel { get; private set; } = null!;

        /// <summary>
        /// The id of the project.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// One of five levels of access to the project.
        /// </summary>
        [Output("pushAccessLevel")]
        public Output<string> PushAccessLevel { get; private set; } = null!;


        /// <summary>
        /// Create a BranchProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BranchProtection(string name, BranchProtectionArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/branchProtection:BranchProtection", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private BranchProtection(string name, Input<string> id, BranchProtectionState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/branchProtection:BranchProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BranchProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BranchProtection Get(string name, Input<string> id, BranchProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new BranchProtection(name, id, state, options);
        }
    }

    public sealed class BranchProtectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the branch.
        /// </summary>
        [Input("branch", required: true)]
        public Input<string> Branch { get; set; } = null!;

        /// <summary>
        /// One of five levels of access to the project.
        /// </summary>
        [Input("mergeAccessLevel", required: true)]
        public Input<string> MergeAccessLevel { get; set; } = null!;

        /// <summary>
        /// The id of the project.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// One of five levels of access to the project.
        /// </summary>
        [Input("pushAccessLevel", required: true)]
        public Input<string> PushAccessLevel { get; set; } = null!;

        public BranchProtectionArgs()
        {
        }
    }

    public sealed class BranchProtectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the branch.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// One of five levels of access to the project.
        /// </summary>
        [Input("mergeAccessLevel")]
        public Input<string>? MergeAccessLevel { get; set; }

        /// <summary>
        /// The id of the project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// One of five levels of access to the project.
        /// </summary>
        [Input("pushAccessLevel")]
        public Input<string>? PushAccessLevel { get; set; }

        public BranchProtectionState()
        {
        }
    }
}
