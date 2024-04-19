// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectBadgeArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectBadgeState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectBadge` resource allows to manage the lifecycle of project badges.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#project-badges)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.ProjectBadge;
 * import com.pulumi.gitlab.ProjectBadgeArgs;
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
 *         var foo = new Project(&#34;foo&#34;, ProjectArgs.builder()        
 *             .name(&#34;foo-project&#34;)
 *             .build());
 * 
 *         var example = new ProjectBadge(&#34;example&#34;, ProjectBadgeArgs.builder()        
 *             .project(foo.id())
 *             .linkUrl(&#34;https://example.com/badge-123&#34;)
 *             .imageUrl(&#34;https://example.com/badge-123.svg&#34;)
 *             .name(&#34;badge-123&#34;)
 *             .build());
 * 
 *         // Pipeline status badges with placeholders will be enabled
 *         var gitlabPipeline = new ProjectBadge(&#34;gitlabPipeline&#34;, ProjectBadgeArgs.builder()        
 *             .project(foo.id())
 *             .linkUrl(&#34;https://gitlab.example.com/%{project_path}/-/pipelines?ref=%{default_branch}&#34;)
 *             .imageUrl(&#34;https://gitlab.example.com/%{project_path}/badges/%{default_branch}/pipeline.svg&#34;)
 *             .name(&#34;badge-pipeline&#34;)
 *             .build());
 * 
 *         // Test coverage report badges with placeholders will be enabled
 *         var gitlabCoverage = new ProjectBadge(&#34;gitlabCoverage&#34;, ProjectBadgeArgs.builder()        
 *             .project(foo.id())
 *             .linkUrl(&#34;https://gitlab.example.com/%{project_path}/-/jobs&#34;)
 *             .imageUrl(&#34;https://gitlab.example.com/%{project_path}/badges/%{default_branch}/coverage.svg&#34;)
 *             .name(&#34;badge-coverage&#34;)
 *             .build());
 * 
 *         // Latest release badges with placeholders will be enabled
 *         var gitlabRelease = new ProjectBadge(&#34;gitlabRelease&#34;, ProjectBadgeArgs.builder()        
 *             .project(foo.id())
 *             .linkUrl(&#34;https://gitlab.example.com/%{project_path}/-/releases&#34;)
 *             .imageUrl(&#34;https://gitlab.example.com/%{project_path}/-/badges/release.svg&#34;)
 *             .name(&#34;badge-release&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitLab project badges can be imported using an id made up of `{project_id}:{badge_id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectBadge:ProjectBadge foo 1:3
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectBadge:ProjectBadge")
public class ProjectBadge extends com.pulumi.resources.CustomResource {
    /**
     * The image url which will be presented on project overview.
     * 
     */
    @Export(name="imageUrl", refs={String.class}, tree="[0]")
    private Output<String> imageUrl;

    /**
     * @return The image url which will be presented on project overview.
     * 
     */
    public Output<String> imageUrl() {
        return this.imageUrl;
    }
    /**
     * The url linked with the badge.
     * 
     */
    @Export(name="linkUrl", refs={String.class}, tree="[0]")
    private Output<String> linkUrl;

    /**
     * @return The url linked with the badge.
     * 
     */
    public Output<String> linkUrl() {
        return this.linkUrl;
    }
    /**
     * The name of the badge.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the badge.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The id of the project to add the badge to.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The id of the project to add the badge to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The image_url argument rendered (in case of use of placeholders).
     * 
     */
    @Export(name="renderedImageUrl", refs={String.class}, tree="[0]")
    private Output<String> renderedImageUrl;

    /**
     * @return The image_url argument rendered (in case of use of placeholders).
     * 
     */
    public Output<String> renderedImageUrl() {
        return this.renderedImageUrl;
    }
    /**
     * The link_url argument rendered (in case of use of placeholders).
     * 
     */
    @Export(name="renderedLinkUrl", refs={String.class}, tree="[0]")
    private Output<String> renderedLinkUrl;

    /**
     * @return The link_url argument rendered (in case of use of placeholders).
     * 
     */
    public Output<String> renderedLinkUrl() {
        return this.renderedLinkUrl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectBadge(String name) {
        this(name, ProjectBadgeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectBadge(String name, ProjectBadgeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectBadge(String name, ProjectBadgeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectBadge:ProjectBadge", name, args == null ? ProjectBadgeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectBadge(String name, Output<String> id, @Nullable ProjectBadgeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectBadge:ProjectBadge", name, state, makeResourceOptions(options, id));
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
    public static ProjectBadge get(String name, Output<String> id, @Nullable ProjectBadgeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectBadge(name, id, state, options);
    }
}
