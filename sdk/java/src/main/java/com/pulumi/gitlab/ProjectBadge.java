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
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/user/project/badges/#project-badges)
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
 *         var foo = new Project("foo", ProjectArgs.builder()
 *             .name("foo-project")
 *             .build());
 * 
 *         var example = new ProjectBadge("example", ProjectBadgeArgs.builder()
 *             .project(foo.id())
 *             .linkUrl("https://example.com/badge-123")
 *             .imageUrl("https://example.com/badge-123.svg")
 *             .name("badge-123")
 *             .build());
 * 
 *         // Pipeline status badges with placeholders will be enabled
 *         var gitlabPipeline = new ProjectBadge("gitlabPipeline", ProjectBadgeArgs.builder()
 *             .project(foo.id())
 *             .linkUrl("https://gitlab.example.com/%{project_path}/-/pipelines?ref=%{default_branch}")
 *             .imageUrl("https://gitlab.example.com/%{project_path}/badges/%{default_branch}/pipeline.svg")
 *             .name("badge-pipeline")
 *             .build());
 * 
 *         // Test coverage report badges with placeholders will be enabled
 *         var gitlabCoverage = new ProjectBadge("gitlabCoverage", ProjectBadgeArgs.builder()
 *             .project(foo.id())
 *             .linkUrl("https://gitlab.example.com/%{project_path}/-/jobs")
 *             .imageUrl("https://gitlab.example.com/%{project_path}/badges/%{default_branch}/coverage.svg")
 *             .name("badge-coverage")
 *             .build());
 * 
 *         // Latest release badges with placeholders will be enabled
 *         var gitlabRelease = new ProjectBadge("gitlabRelease", ProjectBadgeArgs.builder()
 *             .project(foo.id())
 *             .linkUrl("https://gitlab.example.com/%{project_path}/-/releases")
 *             .imageUrl("https://gitlab.example.com/%{project_path}/-/badges/release.svg")
 *             .name("badge-release")
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
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_badge`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_badge.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
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
    public ProjectBadge(java.lang.String name) {
        this(name, ProjectBadgeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectBadge(java.lang.String name, ProjectBadgeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectBadge(java.lang.String name, ProjectBadgeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectBadge:ProjectBadge", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectBadge(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectBadgeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectBadge:ProjectBadge", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectBadgeArgs makeArgs(ProjectBadgeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectBadgeArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ProjectBadge get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectBadgeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectBadge(name, id, state, options);
    }
}
