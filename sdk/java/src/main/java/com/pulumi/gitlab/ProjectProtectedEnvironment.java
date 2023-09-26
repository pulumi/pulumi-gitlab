// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectProtectedEnvironmentArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectProtectedEnvironmentState;
import com.pulumi.gitlab.outputs.ProjectProtectedEnvironmentApprovalRule;
import com.pulumi.gitlab.outputs.ProjectProtectedEnvironmentDeployAccessLevel;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectProtectedEnvironment` resource allows to manage the lifecycle of a protected environment in a project.
 * 
 * &gt; In order to use a user or group in the `deploy_access_levels` configuration,
 *    you need to make sure that users have access to the project and groups must have this project shared.
 *    You may use the `gitlab.ProjectMembership` and `gitlab_project_shared_group` resources to achieve this.
 *    Unfortunately, the GitLab API does not complain about users and groups without access to the project and just ignores those.
 *    In case this happens you will get perpetual state diffs.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_environments.html)
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.ProjectEnvironment;
 * import com.pulumi.gitlab.ProjectEnvironmentArgs;
 * import com.pulumi.gitlab.ProjectProtectedEnvironment;
 * import com.pulumi.gitlab.ProjectProtectedEnvironmentArgs;
 * import com.pulumi.gitlab.inputs.ProjectProtectedEnvironmentDeployAccessLevelArgs;
 * import com.pulumi.gitlab.inputs.ProjectProtectedEnvironmentApprovalRuleArgs;
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
 *         var this_ = new ProjectEnvironment(&#34;this&#34;, ProjectEnvironmentArgs.builder()        
 *             .project(123)
 *             .externalUrl(&#34;www.example.com&#34;)
 *             .build());
 * 
 *         var exampleWithAccessLevel = new ProjectProtectedEnvironment(&#34;exampleWithAccessLevel&#34;, ProjectProtectedEnvironmentArgs.builder()        
 *             .project(this_.project())
 *             .requiredApprovalCount(1)
 *             .environment(this_.name())
 *             .deployAccessLevels(ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                 .accessLevel(&#34;developer&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleWithGroup = new ProjectProtectedEnvironment(&#34;exampleWithGroup&#34;, ProjectProtectedEnvironmentArgs.builder()        
 *             .project(this_.project())
 *             .environment(this_.name())
 *             .deployAccessLevels(ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                 .groupId(456)
 *                 .build())
 *             .build());
 * 
 *         var exampleWithUser = new ProjectProtectedEnvironment(&#34;exampleWithUser&#34;, ProjectProtectedEnvironmentArgs.builder()        
 *             .project(this_.project())
 *             .environment(this_.name())
 *             .deployAccessLevels(ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                 .userId(789)
 *                 .build())
 *             .build());
 * 
 *         var exampleWithMultipleProjectProtectedEnvironment = new ProjectProtectedEnvironment(&#34;exampleWithMultipleProjectProtectedEnvironment&#34;, ProjectProtectedEnvironmentArgs.builder()        
 *             .project(this_.project())
 *             .requiredApprovalCount(2)
 *             .environment(this_.name())
 *             .deployAccessLevels(            
 *                 ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                     .accessLevel(&#34;developer&#34;)
 *                     .build(),
 *                 ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                     .groupId(456)
 *                     .build(),
 *                 ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                     .userId(789)
 *                     .build())
 *             .build());
 * 
 *         var exampleWithMultipleIndex_projectProtectedEnvironmentProjectProtectedEnvironment = new ProjectProtectedEnvironment(&#34;exampleWithMultipleIndex/projectProtectedEnvironmentProjectProtectedEnvironment&#34;, ProjectProtectedEnvironmentArgs.builder()        
 *             .project(this_.project())
 *             .requiredApprovalCount(2)
 *             .environment(this_.name())
 *             .deployAccessLevels(ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                 .accessLevel(&#34;developer&#34;)
 *                 .build())
 *             .approvalRules(ProjectProtectedEnvironmentApprovalRuleArgs.builder()
 *                 .access_level(&#34;developer&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleWithMultipleGitlabIndex_projectProtectedEnvironmentProjectProtectedEnvironment = new ProjectProtectedEnvironment(&#34;exampleWithMultipleGitlabIndex/projectProtectedEnvironmentProjectProtectedEnvironment&#34;, ProjectProtectedEnvironmentArgs.builder()        
 *             .project(this_.project())
 *             .requiredApprovalCount(2)
 *             .environment(this_.name())
 *             .deployAccessLevels(ProjectProtectedEnvironmentDeployAccessLevelArgs.builder()
 *                 .accessLevel(&#34;developer&#34;)
 *                 .build())
 *             .approvalRules(            
 *                 ProjectProtectedEnvironmentApprovalRuleArgs.builder()
 *                     .user_id(789)
 *                     .build(),
 *                 ProjectProtectedEnvironmentApprovalRuleArgs.builder()
 *                     .access_level(&#34;developer&#34;)
 *                     .build(),
 *                 ProjectProtectedEnvironmentApprovalRuleArgs.builder()
 *                     .group_id(456)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitLab protected environments can be imported using an id made up of `projectId:environmentName`, e.g.
 * 
 * ```sh
 *  $ pulumi import gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment bar 123:production
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment")
public class ProjectProtectedEnvironment extends com.pulumi.resources.CustomResource {
    /**
     * Array of approval rules to deploy, with each described by a hash.
     * 
     */
    @Export(name="approvalRules", refs={List.class,ProjectProtectedEnvironmentApprovalRule.class}, tree="[0,1]")
    private Output<List<ProjectProtectedEnvironmentApprovalRule>> approvalRules;

    /**
     * @return Array of approval rules to deploy, with each described by a hash.
     * 
     */
    public Output<List<ProjectProtectedEnvironmentApprovalRule>> approvalRules() {
        return this.approvalRules;
    }
    /**
     * Array of access levels allowed to deploy, with each described by a hash.
     * 
     */
    @Export(name="deployAccessLevels", refs={List.class,ProjectProtectedEnvironmentDeployAccessLevel.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ProjectProtectedEnvironmentDeployAccessLevel>> deployAccessLevels;

    /**
     * @return Array of access levels allowed to deploy, with each described by a hash.
     * 
     */
    public Output<Optional<List<ProjectProtectedEnvironmentDeployAccessLevel>>> deployAccessLevels() {
        return Codegen.optional(this.deployAccessLevels);
    }
    /**
     * The name of the environment.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output<String> environment;

    /**
     * @return The name of the environment.
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }
    /**
     * The ID or full path of the project which the protected environment is created against.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project which the protected environment is created against.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The number of approvals required to deploy to this environment.
     * 
     */
    @Export(name="requiredApprovalCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> requiredApprovalCount;

    /**
     * @return The number of approvals required to deploy to this environment.
     * 
     */
    public Output<Integer> requiredApprovalCount() {
        return this.requiredApprovalCount;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectProtectedEnvironment(String name) {
        this(name, ProjectProtectedEnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectProtectedEnvironment(String name, ProjectProtectedEnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectProtectedEnvironment(String name, ProjectProtectedEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment", name, args == null ? ProjectProtectedEnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectProtectedEnvironment(String name, Output<String> id, @Nullable ProjectProtectedEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment", name, state, makeResourceOptions(options, id));
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
    public static ProjectProtectedEnvironment get(String name, Output<String> id, @Nullable ProjectProtectedEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectProtectedEnvironment(name, id, state, options);
    }
}
