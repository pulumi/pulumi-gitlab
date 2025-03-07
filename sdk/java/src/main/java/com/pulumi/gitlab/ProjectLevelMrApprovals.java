// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectLevelMrApprovalsArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectLevelMrApprovalsState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab_project_level_mr_approval_rule` resource allows to manage the lifecycle of a Merge Request-level approval rule.
 * 
 * &gt; This resource requires a GitLab Enterprise instance.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/merge_request_approvals/#merge-request-level-mr-approvals)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.ProjectLevelMrApprovals;
 * import com.pulumi.gitlab.ProjectLevelMrApprovalsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var foo = new Project("foo", ProjectArgs.builder()
 *             .name("Example")
 *             .description("My example project")
 *             .build());
 * 
 *         var fooProjectLevelMrApprovals = new ProjectLevelMrApprovals("fooProjectLevelMrApprovals", ProjectLevelMrApprovalsArgs.builder()
 *             .project(foo.id())
 *             .resetApprovalsOnPush(true)
 *             .disableOverridingApproversPerMergeRequest(false)
 *             .mergeRequestsAuthorApproval(false)
 *             .mergeRequestsDisableCommittersApproval(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_level_mr_approvals`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_level_mr_approvals.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals You can import an approval configuration state using `&lt;resource&gt; &lt;project_id&gt;`.
 * ```
 * 
 * # 
 * 
 * For example:
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals")
public class ProjectLevelMrApprovals extends com.pulumi.resources.CustomResource {
    /**
     * Set to `true` to disable overriding approvers per merge request.
     * 
     */
    @Export(name="disableOverridingApproversPerMergeRequest", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableOverridingApproversPerMergeRequest;

    /**
     * @return Set to `true` to disable overriding approvers per merge request.
     * 
     */
    public Output<Boolean> disableOverridingApproversPerMergeRequest() {
        return this.disableOverridingApproversPerMergeRequest;
    }
    /**
     * Set to `true` to allow merge requests authors to approve their own merge requests.
     * 
     */
    @Export(name="mergeRequestsAuthorApproval", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeRequestsAuthorApproval;

    /**
     * @return Set to `true` to allow merge requests authors to approve their own merge requests.
     * 
     */
    public Output<Boolean> mergeRequestsAuthorApproval() {
        return this.mergeRequestsAuthorApproval;
    }
    /**
     * Set to `true` to disable merge request committers from approving their own merge requests.
     * 
     */
    @Export(name="mergeRequestsDisableCommittersApproval", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mergeRequestsDisableCommittersApproval;

    /**
     * @return Set to `true` to disable merge request committers from approving their own merge requests.
     * 
     */
    public Output<Boolean> mergeRequestsDisableCommittersApproval() {
        return this.mergeRequestsDisableCommittersApproval;
    }
    /**
     * The ID or URL-encoded path of a project to change MR approval configuration.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or URL-encoded path of a project to change MR approval configuration.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Set to `true` to require authentication to approve merge requests.
     * 
     */
    @Export(name="requirePasswordToApprove", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> requirePasswordToApprove;

    /**
     * @return Set to `true` to require authentication to approve merge requests.
     * 
     */
    public Output<Boolean> requirePasswordToApprove() {
        return this.requirePasswordToApprove;
    }
    /**
     * Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     * 
     */
    @Export(name="resetApprovalsOnPush", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> resetApprovalsOnPush;

    /**
     * @return Set to `true` to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
     * 
     */
    public Output<Boolean> resetApprovalsOnPush() {
        return this.resetApprovalsOnPush;
    }
    /**
     * Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
     * 
     */
    @Export(name="selectiveCodeOwnerRemovals", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> selectiveCodeOwnerRemovals;

    /**
     * @return Reset approvals from Code Owners if their files changed. Can be enabled only if reset*approvals*on_push is disabled.
     * 
     */
    public Output<Boolean> selectiveCodeOwnerRemovals() {
        return this.selectiveCodeOwnerRemovals;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectLevelMrApprovals(String name) {
        this(name, ProjectLevelMrApprovalsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectLevelMrApprovals(String name, ProjectLevelMrApprovalsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectLevelMrApprovals(String name, ProjectLevelMrApprovalsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, args == null ? ProjectLevelMrApprovalsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectLevelMrApprovals(String name, Output<String> id, @Nullable ProjectLevelMrApprovalsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ProjectLevelMrApprovals get(String name, Output<String> id, @Nullable ProjectLevelMrApprovalsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectLevelMrApprovals(name, id, state, options);
    }
}
