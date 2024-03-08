// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectEnvironmentArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectEnvironmentState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.ProjectEnvironment;
 * import com.pulumi.gitlab.ProjectEnvironmentArgs;
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
 *         var thisGroup = new Group(&#34;thisGroup&#34;, GroupArgs.builder()        
 *             .path(&#34;example&#34;)
 *             .description(&#34;An example group&#34;)
 *             .build());
 * 
 *         var thisProject = new Project(&#34;thisProject&#34;, ProjectArgs.builder()        
 *             .namespaceId(thisGroup.id())
 *             .initializeWithReadme(true)
 *             .build());
 * 
 *         var thisProjectEnvironment = new ProjectEnvironment(&#34;thisProjectEnvironment&#34;, ProjectEnvironmentArgs.builder()        
 *             .project(thisProject.id())
 *             .externalUrl(&#34;www.example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitLab project environments can be imported using an id made up of `projectId:environmenId`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectEnvironment:ProjectEnvironment bar 123:321
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectEnvironment:ProjectEnvironment")
public class ProjectEnvironment extends com.pulumi.resources.CustomResource {
    /**
     * The ISO8601 date/time that this environment was created at in UTC.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The ISO8601 date/time that this environment was created at in UTC.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Place to link to for this environment.
     * 
     */
    @Export(name="externalUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> externalUrl;

    /**
     * @return Place to link to for this environment.
     * 
     */
    public Output<Optional<String>> externalUrl() {
        return Codegen.optional(this.externalUrl);
    }
    /**
     * The name of the environment.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the environment.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID or full path of the project to environment is created for.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project to environment is created for.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    @Export(name="slug", refs={String.class}, tree="[0]")
    private Output<String> slug;

    /**
     * @return The name of the environment in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }
    /**
     * State the environment is in. Valid values are `available`, `stopped`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State the environment is in. Valid values are `available`, `stopped`.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Determines whether the environment is attempted to be stopped before the environment is deleted.
     * 
     */
    @Export(name="stopBeforeDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stopBeforeDestroy;

    /**
     * @return Determines whether the environment is attempted to be stopped before the environment is deleted.
     * 
     */
    public Output<Optional<Boolean>> stopBeforeDestroy() {
        return Codegen.optional(this.stopBeforeDestroy);
    }
    /**
     * The ISO8601 date/time that this environment was last updated at in UTC.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The ISO8601 date/time that this environment was last updated at in UTC.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectEnvironment(String name) {
        this(name, ProjectEnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectEnvironment(String name, ProjectEnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectEnvironment(String name, ProjectEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectEnvironment:ProjectEnvironment", name, args == null ? ProjectEnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectEnvironment(String name, Output<String> id, @Nullable ProjectEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectEnvironment:ProjectEnvironment", name, state, makeResourceOptions(options, id));
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
    public static ProjectEnvironment get(String name, Output<String> id, @Nullable ProjectEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectEnvironment(name, id, state, options);
    }
}
