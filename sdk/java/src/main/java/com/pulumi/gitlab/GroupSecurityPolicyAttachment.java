// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupSecurityPolicyAttachmentArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupSecurityPolicyAttachmentState;
import java.lang.String;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.gitlab.GroupSecurityPolicyAttachment;
 * import com.pulumi.gitlab.GroupSecurityPolicyAttachmentArgs;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.RepositoryFile;
 * import com.pulumi.gitlab.RepositoryFileArgs;
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
 *         // This resource can be used to attach a security policy to a pre-existing group
 *         var foo = new GroupSecurityPolicyAttachment("foo", GroupSecurityPolicyAttachmentArgs.builder()
 *             .group(1234)
 *             .policyProject(4567)
 *             .build());
 * 
 *         // Or Terraform can create a new project, add a policy to that project,
 *         // then attach that policy project to other groups.
 *         var my_policy_project = new Project("my-policy-project", ProjectArgs.builder()
 *             .name("security-policy-project")
 *             .build());
 * 
 *         var policy_yml = new RepositoryFile("policy-yml", RepositoryFileArgs.builder()
 *             .project(my_policy_project.id())
 *             .filePath(".gitlab/security-policies/my-policy.yml")
 *             .branch("master")
 *             .encoding("text")
 *             .content("""
 * ---
 * approval_policy:
 * - name: test
 * description: test
 * enabled: true
 * rules:
 * - type: any_merge_request
 *     branch_type: protected
 *     commits: any
 * approval_settings:
 *     block_branch_modification: true
 *     prevent_pushing_and_force_pushing: true
 *     prevent_approval_by_author: true
 *     prevent_approval_by_commit_author: true
 *     remove_approvals_with_new_commit: true
 *     require_password_to_approve: false
 * fallback_behavior:
 *     fail: closed
 * policy_scope:
 *   compliance_frameworks:
 *   - id: 1010101
 *   - id: 0101010
 * actions:
 * - type: send_bot_message
 *     enabled: true
 *             """)
 *             .build());
 * 
 *         // Multiple policies can be attached to a single project by repeating this resource or using a `for_each`
 *         var my_policy = new GroupSecurityPolicyAttachment("my-policy", GroupSecurityPolicyAttachmentArgs.builder()
 *             .group(1234)
 *             .policyProject(my_policy_project.id())
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_security_policy_attachment`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_group_security_policy_attachment.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * GitLab group security policy attachments can be imported using an id made up of `group:policy_project_id` where the policy project ID is the project ID of the policy project, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupSecurityPolicyAttachment:GroupSecurityPolicyAttachment foo 1:2
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupSecurityPolicyAttachment:GroupSecurityPolicyAttachment")
public class GroupSecurityPolicyAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID or Full Path of the group which will have the security policy project assigned to it.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or Full Path of the group which will have the security policy project assigned to it.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * The GraphQL ID of the group to which the security policty project will be attached.
     * 
     */
    @Export(name="groupGraphqlId", refs={String.class}, tree="[0]")
    private Output<String> groupGraphqlId;

    /**
     * @return The GraphQL ID of the group to which the security policty project will be attached.
     * 
     */
    public Output<String> groupGraphqlId() {
        return this.groupGraphqlId;
    }
    /**
     * The ID or Full Path of the security policy project.
     * 
     */
    @Export(name="policyProject", refs={String.class}, tree="[0]")
    private Output<String> policyProject;

    /**
     * @return The ID or Full Path of the security policy project.
     * 
     */
    public Output<String> policyProject() {
        return this.policyProject;
    }
    /**
     * The GraphQL ID of the security policy project.
     * 
     */
    @Export(name="policyProjectGraphqlId", refs={String.class}, tree="[0]")
    private Output<String> policyProjectGraphqlId;

    /**
     * @return The GraphQL ID of the security policy project.
     * 
     */
    public Output<String> policyProjectGraphqlId() {
        return this.policyProjectGraphqlId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupSecurityPolicyAttachment(String name) {
        this(name, GroupSecurityPolicyAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupSecurityPolicyAttachment(String name, GroupSecurityPolicyAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupSecurityPolicyAttachment(String name, GroupSecurityPolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupSecurityPolicyAttachment:GroupSecurityPolicyAttachment", name, args == null ? GroupSecurityPolicyAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupSecurityPolicyAttachment(String name, Output<String> id, @Nullable GroupSecurityPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupSecurityPolicyAttachment:GroupSecurityPolicyAttachment", name, state, makeResourceOptions(options, id));
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
    public static GroupSecurityPolicyAttachment get(String name, Output<String> id, @Nullable GroupSecurityPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupSecurityPolicyAttachment(name, id, state, options);
    }
}
