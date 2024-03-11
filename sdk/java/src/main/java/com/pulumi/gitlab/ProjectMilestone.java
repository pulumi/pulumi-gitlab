// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectMilestoneArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectMilestoneState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectMilestone` resource allows to manage the lifecycle of a project milestone.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
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
 * import com.pulumi.gitlab.ProjectMilestone;
 * import com.pulumi.gitlab.ProjectMilestoneArgs;
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;An example project&#34;)
 *             .namespaceId(gitlab_group.example().id())
 *             .build());
 * 
 *         var exampleProjectMilestone = new ProjectMilestone(&#34;exampleProjectMilestone&#34;, ProjectMilestoneArgs.builder()        
 *             .project(exampleProject.id())
 *             .title(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Gitlab project milestone can be imported with a key composed of `&lt;project&gt;:&lt;milestone_id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectMilestone:ProjectMilestone example &#34;12345:11&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectMilestone:ProjectMilestone")
public class ProjectMilestone extends com.pulumi.resources.CustomResource {
    /**
     * The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The description of the milestone.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the milestone.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    @Export(name="dueDate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dueDate;

    /**
     * @return The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public Output<Optional<String>> dueDate() {
        return Codegen.optional(this.dueDate);
    }
    /**
     * Bool, true if milestone expired.
     * 
     */
    @Export(name="expired", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> expired;

    /**
     * @return Bool, true if milestone expired.
     * 
     */
    public Output<Boolean> expired() {
        return this.expired;
    }
    /**
     * The ID of the project&#39;s milestone.
     * 
     */
    @Export(name="iid", refs={Integer.class}, tree="[0]")
    private Output<Integer> iid;

    /**
     * @return The ID of the project&#39;s milestone.
     * 
     */
    public Output<Integer> iid() {
        return this.iid;
    }
    /**
     * The instance-wide ID of the project’s milestone.
     * 
     */
    @Export(name="milestoneId", refs={Integer.class}, tree="[0]")
    private Output<Integer> milestoneId;

    /**
     * @return The instance-wide ID of the project’s milestone.
     * 
     */
    public Output<Integer> milestoneId() {
        return this.milestoneId;
    }
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
     * The project ID of milestone.
     * 
     */
    @Export(name="projectId", refs={Integer.class}, tree="[0]")
    private Output<Integer> projectId;

    /**
     * @return The project ID of milestone.
     * 
     */
    public Output<Integer> projectId() {
        return this.projectId;
    }
    /**
     * The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    @Export(name="startDate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> startDate;

    /**
     * @return The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public Output<Optional<String>> startDate() {
        return Codegen.optional(this.startDate);
    }
    /**
     * The state of the milestone. Valid values are: `active`, `closed`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state of the milestone. Valid values are: `active`, `closed`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * The title of a milestone.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return The title of a milestone.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * The web URL of the milestone.
     * 
     */
    @Export(name="webUrl", refs={String.class}, tree="[0]")
    private Output<String> webUrl;

    /**
     * @return The web URL of the milestone.
     * 
     */
    public Output<String> webUrl() {
        return this.webUrl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectMilestone(String name) {
        this(name, ProjectMilestoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectMilestone(String name, ProjectMilestoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectMilestone(String name, ProjectMilestoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectMilestone:ProjectMilestone", name, args == null ? ProjectMilestoneArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectMilestone(String name, Output<String> id, @Nullable ProjectMilestoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectMilestone:ProjectMilestone", name, state, makeResourceOptions(options, id));
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
    public static ProjectMilestone get(String name, Output<String> id, @Nullable ProjectMilestoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectMilestone(name, id, state, options);
    }
}
