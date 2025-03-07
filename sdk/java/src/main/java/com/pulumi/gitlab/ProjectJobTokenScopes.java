// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectJobTokenScopesArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectJobTokenScopesState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectJobTokenScopes` resource allows to manage the CI/CD Job Token scopes in a project.
 * Any project not within the defined set in this attribute will be removed, which allows this resource to be used as an explicit deny.
 * 
 * &gt; Conflicts with the use of `gitlab.ProjectJobTokenScope` when used on the same project. Use one or the other to ensure the desired state.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_job_token_scopes/)
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
 * import com.pulumi.gitlab.ProjectJobTokenScopes;
 * import com.pulumi.gitlab.ProjectJobTokenScopesArgs;
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
 *         var allowedSingleProject = new ProjectJobTokenScopes("allowedSingleProject", ProjectJobTokenScopesArgs.builder()
 *             .project("111")
 *             .targetProjectIds(123)
 *             .build());
 * 
 *         var allowedMultipleProject = new ProjectJobTokenScopes("allowedMultipleProject", ProjectJobTokenScopesArgs.builder()
 *             .project("111")
 *             .targetProjectIds(            
 *                 123,
 *                 456,
 *                 789)
 *             .build());
 * 
 *         var allowedMultipleGroups = new ProjectJobTokenScopes("allowedMultipleGroups", ProjectJobTokenScopesArgs.builder()
 *             .projectId(111)
 *             .targetProjectIds()
 *             .targetGroupIds(            
 *                 321,
 *                 654)
 *             .build());
 * 
 *         // This will remove all job token scopes, even if added outside of TF.
 *         var explicitDeny = new ProjectJobTokenScopes("explicitDeny", ProjectJobTokenScopesArgs.builder()
 *             .project("111")
 *             .targetProjectIds()
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_job_token_scopes`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_job_token_scopes.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * GitLab project job token scopes can be imported using an id made up of just the `project_id`
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectJobTokenScopes:ProjectJobTokenScopes bar 123
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectJobTokenScopes:ProjectJobTokenScopes")
public class ProjectJobTokenScopes extends com.pulumi.resources.CustomResource {
    /**
     * The ID or full path of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The ID of the project.
     * 
     * @deprecated
     * `project_id` has been deprecated. Use `project` instead.
     * 
     */
    @Deprecated /* `project_id` has been deprecated. Use `project` instead. */
    @Export(name="projectId", refs={Integer.class}, tree="[0]")
    private Output<Integer> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<Integer> projectId() {
        return this.projectId;
    }
    /**
     * A set of group IDs that are in the CI/CD job token inbound allowlist.
     * 
     */
    @Export(name="targetGroupIds", refs={List.class,Integer.class}, tree="[0,1]")
    private Output<List<Integer>> targetGroupIds;

    /**
     * @return A set of group IDs that are in the CI/CD job token inbound allowlist.
     * 
     */
    public Output<List<Integer>> targetGroupIds() {
        return this.targetGroupIds;
    }
    /**
     * A set of project IDs that are in the CI/CD job token inbound allowlist.
     * 
     */
    @Export(name="targetProjectIds", refs={List.class,Integer.class}, tree="[0,1]")
    private Output<List<Integer>> targetProjectIds;

    /**
     * @return A set of project IDs that are in the CI/CD job token inbound allowlist.
     * 
     */
    public Output<List<Integer>> targetProjectIds() {
        return this.targetProjectIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectJobTokenScopes(String name) {
        this(name, ProjectJobTokenScopesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectJobTokenScopes(String name, @Nullable ProjectJobTokenScopesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectJobTokenScopes(String name, @Nullable ProjectJobTokenScopesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectJobTokenScopes:ProjectJobTokenScopes", name, args == null ? ProjectJobTokenScopesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectJobTokenScopes(String name, Output<String> id, @Nullable ProjectJobTokenScopesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectJobTokenScopes:ProjectJobTokenScopes", name, state, makeResourceOptions(options, id));
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
    public static ProjectJobTokenScopes get(String name, Output<String> id, @Nullable ProjectJobTokenScopesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectJobTokenScopes(name, id, state, options);
    }
}
