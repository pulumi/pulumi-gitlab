// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ReleaseArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ReleaseState;
import com.pulumi.gitlab.outputs.ReleaseAssets;
import com.pulumi.gitlab.outputs.ReleaseAuthor;
import com.pulumi.gitlab.outputs.ReleaseCommit;
import com.pulumi.gitlab.outputs.ReleaseLinks;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.Release` resource allows to manage the lifecycle of releases in gitlab.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/releases/)
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
 * import com.pulumi.gitlab.Release;
 * import com.pulumi.gitlab.ReleaseArgs;
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
 *         // Create a project
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name("example")
 *             .description("An example project")
 *             .build());
 * 
 *         // Create a release
 *         var exampleRelease = new Release("exampleRelease", ReleaseArgs.builder()
 *             .project(example.id())
 *             .name("test-release")
 *             .tagName("v1.0.0")
 *             .description("Test release description")
 *             .ref("main")
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
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_release`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_release.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
 * 
 * Gitlab release link can be imported with a key composed of `&lt;project&gt;:&lt;tag_name&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/release:Release example &#34;12345:test&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/release:Release")
public class Release extends com.pulumi.resources.CustomResource {
    /**
     * The release assets.
     * 
     */
    @Export(name="assets", refs={ReleaseAssets.class}, tree="[0]")
    private Output<ReleaseAssets> assets;

    /**
     * @return The release assets.
     * 
     */
    public Output<ReleaseAssets> assets() {
        return this.assets;
    }
    /**
     * The author of the release.
     * 
     */
    @Export(name="author", refs={ReleaseAuthor.class}, tree="[0]")
    private Output<ReleaseAuthor> author;

    /**
     * @return The author of the release.
     * 
     */
    public Output<ReleaseAuthor> author() {
        return this.author;
    }
    /**
     * The release commit.
     * 
     */
    @Export(name="commit", refs={ReleaseCommit.class}, tree="[0]")
    private Output<ReleaseCommit> commit;

    /**
     * @return The release commit.
     * 
     */
    public Output<ReleaseCommit> commit() {
        return this.commit;
    }
    /**
     * The path to the commit
     * 
     */
    @Export(name="commitPath", refs={String.class}, tree="[0]")
    private Output<String> commitPath;

    /**
     * @return The path to the commit
     * 
     */
    public Output<String> commitPath() {
        return this.commitPath;
    }
    /**
     * Date and time the release was created. In ISO 8601 format (2019-03-15T08:00:00Z).
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date and time the release was created. In ISO 8601 format (2019-03-15T08:00:00Z).
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The description of the release. You can use Markdown.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the release. You can use Markdown.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * HTML rendered Markdown of the release description.
     * 
     */
    @Export(name="descriptionHtml", refs={String.class}, tree="[0]")
    private Output<String> descriptionHtml;

    /**
     * @return HTML rendered Markdown of the release description.
     * 
     */
    public Output<String> descriptionHtml() {
        return this.descriptionHtml;
    }
    /**
     * Links of the release
     * 
     */
    @Export(name="links", refs={ReleaseLinks.class}, tree="[0]")
    private Output<ReleaseLinks> links;

    /**
     * @return Links of the release
     * 
     */
    public Output<ReleaseLinks> links() {
        return this.links;
    }
    /**
     * The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones.
     * 
     */
    @Export(name="milestones", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> milestones;

    /**
     * @return The title of each milestone the release is associated with. GitLab Premium customers can specify group milestones.
     * 
     */
    public Output<Optional<List<String>>> milestones() {
        return Codegen.optional(this.milestones);
    }
    /**
     * The name of the release.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the release.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
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
     * If a tag specified in tag*name doesn&#39;t exist, the release is created from ref and tagged with tag*name. It can be a commit SHA, another tag name, or a branch name.
     * 
     */
    @Export(name="ref", refs={String.class}, tree="[0]")
    private Output<String> ref;

    /**
     * @return If a tag specified in tag*name doesn&#39;t exist, the release is created from ref and tagged with tag*name. It can be a commit SHA, another tag name, or a branch name.
     * 
     */
    public Output<String> ref() {
        return this.ref;
    }
    /**
     * Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (2019-03-15T08:00:00Z). Only provide this field if creating an upcoming or historical release.
     * 
     */
    @Export(name="releasedAt", refs={String.class}, tree="[0]")
    private Output<String> releasedAt;

    /**
     * @return Date and time for the release. Defaults to the current time. Expected in ISO 8601 format (2019-03-15T08:00:00Z). Only provide this field if creating an upcoming or historical release.
     * 
     */
    public Output<String> releasedAt() {
        return this.releasedAt;
    }
    /**
     * Message to use if creating a new annotated tag.
     * 
     */
    @Export(name="tagMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tagMessage;

    /**
     * @return Message to use if creating a new annotated tag.
     * 
     */
    public Output<Optional<String>> tagMessage() {
        return Codegen.optional(this.tagMessage);
    }
    /**
     * The tag where the release is created from.
     * 
     */
    @Export(name="tagName", refs={String.class}, tree="[0]")
    private Output<String> tagName;

    /**
     * @return The tag where the release is created from.
     * 
     */
    public Output<String> tagName() {
        return this.tagName;
    }
    /**
     * The path to the tag.
     * 
     */
    @Export(name="tagPath", refs={String.class}, tree="[0]")
    private Output<String> tagPath;

    /**
     * @return The path to the tag.
     * 
     */
    public Output<String> tagPath() {
        return this.tagPath;
    }
    /**
     * Whether the release_at attribute is set to a future date.
     * 
     */
    @Export(name="upcomingRelease", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> upcomingRelease;

    /**
     * @return Whether the release_at attribute is set to a future date.
     * 
     */
    public Output<Boolean> upcomingRelease() {
        return this.upcomingRelease;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Release(java.lang.String name) {
        this(name, ReleaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Release(java.lang.String name, ReleaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Release(java.lang.String name, ReleaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/release:Release", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Release(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/release:Release", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseArgs makeArgs(ReleaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseArgs.Empty : args;
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
    public static Release get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Release(name, id, state, options);
    }
}
