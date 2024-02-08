// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectRunnerEnablementArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectRunnerEnablementState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectRunnerEnablement` resource allows to enable a runner in a project.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/runners.html#enable-a-runner-in-project)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.ProjectRunnerEnablement;
 * import com.pulumi.gitlab.ProjectRunnerEnablementArgs;
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
 *         var foo = new ProjectRunnerEnablement(&#34;foo&#34;, ProjectRunnerEnablementArgs.builder()        
 *             .project(5)
 *             .runnerId(7)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitLab project runners can be imported using an id made up of `project:runner_id`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement foo 5:7
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement")
public class ProjectRunnerEnablement extends com.pulumi.resources.CustomResource {
    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The ID of a runner to enable for the project.
     * 
     */
    @Export(name="runnerId", refs={Integer.class}, tree="[0]")
    private Output<Integer> runnerId;

    /**
     * @return The ID of a runner to enable for the project.
     * 
     */
    public Output<Integer> runnerId() {
        return this.runnerId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectRunnerEnablement(String name) {
        this(name, ProjectRunnerEnablementArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectRunnerEnablement(String name, ProjectRunnerEnablementArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectRunnerEnablement(String name, ProjectRunnerEnablementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement", name, args == null ? ProjectRunnerEnablementArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectRunnerEnablement(String name, Output<String> id, @Nullable ProjectRunnerEnablementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement", name, state, makeResourceOptions(options, id));
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
    public static ProjectRunnerEnablement get(String name, Output<String> id, @Nullable ProjectRunnerEnablementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectRunnerEnablement(name, id, state, options);
    }
}
