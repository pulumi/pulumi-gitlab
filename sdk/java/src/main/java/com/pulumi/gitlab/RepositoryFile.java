// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.RepositoryFileArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.RepositoryFileState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
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
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
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
 *         var this_ = new Group("this", GroupArgs.builder()
 *             .name("example")
 *             .path("example")
 *             .description("An example group")
 *             .build());
 * 
 *         var thisProject = new Project("thisProject", ProjectArgs.builder()
 *             .name("example")
 *             .namespaceId(this_.id())
 *             .initializeWithReadme(true)
 *             .build());
 * 
 *         var thisRepositoryFile = new RepositoryFile("thisRepositoryFile", RepositoryFileArgs.builder()
 *             .project(thisProject.id())
 *             .filePath("meow.txt")
 *             .branch("main")
 *             .content(StdFunctions.base64encode(Base64encodeArgs.builder()
 *                 .input("Meow goes the cat")
 *                 .build()).result())
 *             .authorEmail("terraform{@literal @}example.com")
 *             .authorName("Terraform")
 *             .commitMessage("feature: add meow file")
 *             .build());
 * 
 *         var readme = new RepositoryFile("readme", RepositoryFileArgs.builder()
 *             .project(thisProject.id())
 *             .filePath("readme.txt")
 *             .branch("main")
 *             .content("Meow goes the cat")
 *             .authorEmail("terraform{@literal @}example.com")
 *             .authorName("Terraform")
 *             .commitMessage("feature: add readme file")
 *             .build());
 * 
 *         var readmeForDogs = new RepositoryFile("readmeForDogs", RepositoryFileArgs.builder()
 *             .project(thisProject.id())
 *             .filePath("readme.txt")
 *             .branch("main")
 *             .content("Bark goes the dog")
 *             .authorEmail("terraform{@literal @}example.com")
 *             .authorName("Terraform")
 *             .commitMessage("feature: update readme file")
 *             .overwriteOnCreate(true)
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
 * A Repository File can be imported using an id made up of `&lt;project-id&gt;:&lt;branch-name&gt;:&lt;file-path&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/repositoryFile:RepositoryFile this 1:main:foo/bar.txt
 * ```
 * 
 */
@ResourceType(type="gitlab:index/repositoryFile:RepositoryFile")
public class RepositoryFile extends com.pulumi.resources.CustomResource {
    /**
     * Email of the commit author.
     * 
     */
    @Export(name="authorEmail", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authorEmail;

    /**
     * @return Email of the commit author.
     * 
     */
    public Output<Optional<String>> authorEmail() {
        return Codegen.optional(this.authorEmail);
    }
    /**
     * Name of the commit author.
     * 
     */
    @Export(name="authorName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authorName;

    /**
     * @return Name of the commit author.
     * 
     */
    public Output<Optional<String>> authorName() {
        return Codegen.optional(this.authorName);
    }
    /**
     * The blob id.
     * 
     */
    @Export(name="blobId", refs={String.class}, tree="[0]")
    private Output<String> blobId;

    /**
     * @return The blob id.
     * 
     */
    public Output<String> blobId() {
        return this.blobId;
    }
    /**
     * Name of the branch to which to commit to.
     * 
     */
    @Export(name="branch", refs={String.class}, tree="[0]")
    private Output<String> branch;

    /**
     * @return Name of the branch to which to commit to.
     * 
     */
    public Output<String> branch() {
        return this.branch;
    }
    /**
     * The commit id.
     * 
     */
    @Export(name="commitId", refs={String.class}, tree="[0]")
    private Output<String> commitId;

    /**
     * @return The commit id.
     * 
     */
    public Output<String> commitId() {
        return this.commitId;
    }
    /**
     * Commit message.
     * 
     */
    @Export(name="commitMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> commitMessage;

    /**
     * @return Commit message.
     * 
     */
    public Output<Optional<String>> commitMessage() {
        return Codegen.optional(this.commitMessage);
    }
    /**
     * File content.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return File content.
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * File content sha256 digest.
     * 
     */
    @Export(name="contentSha256", refs={String.class}, tree="[0]")
    private Output<String> contentSha256;

    /**
     * @return File content sha256 digest.
     * 
     */
    public Output<String> contentSha256() {
        return this.contentSha256;
    }
    /**
     * Create commit message.
     * 
     */
    @Export(name="createCommitMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> createCommitMessage;

    /**
     * @return Create commit message.
     * 
     */
    public Output<Optional<String>> createCommitMessage() {
        return Codegen.optional(this.createCommitMessage);
    }
    /**
     * Delete Commit message.
     * 
     */
    @Export(name="deleteCommitMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deleteCommitMessage;

    /**
     * @return Delete Commit message.
     * 
     */
    public Output<Optional<String>> deleteCommitMessage() {
        return Codegen.optional(this.deleteCommitMessage);
    }
    /**
     * The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
     * 
     */
    @Export(name="encoding", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encoding;

    /**
     * @return The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
     * 
     */
    public Output<Optional<String>> encoding() {
        return Codegen.optional(this.encoding);
    }
    /**
     * Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
     * 
     */
    @Export(name="executeFilemode", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> executeFilemode;

    /**
     * @return Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
     * 
     */
    public Output<Optional<Boolean>> executeFilemode() {
        return Codegen.optional(this.executeFilemode);
    }
    /**
     * The filename.
     * 
     */
    @Export(name="fileName", refs={String.class}, tree="[0]")
    private Output<String> fileName;

    /**
     * @return The filename.
     * 
     */
    public Output<String> fileName() {
        return this.fileName;
    }
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     * 
     */
    @Export(name="filePath", refs={String.class}, tree="[0]")
    private Output<String> filePath;

    /**
     * @return The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     * 
     */
    public Output<String> filePath() {
        return this.filePath;
    }
    /**
     * The last known commit id.
     * 
     */
    @Export(name="lastCommitId", refs={String.class}, tree="[0]")
    private Output<String> lastCommitId;

    /**
     * @return The last known commit id.
     * 
     */
    public Output<String> lastCommitId() {
        return this.lastCommitId;
    }
    /**
     * Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     * 
     */
    @Export(name="overwriteOnCreate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> overwriteOnCreate;

    /**
     * @return Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     * 
     */
    public Output<Optional<Boolean>> overwriteOnCreate() {
        return Codegen.optional(this.overwriteOnCreate);
    }
    /**
     * The name or ID of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name or ID of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The name of branch, tag or commit.
     * 
     */
    @Export(name="ref", refs={String.class}, tree="[0]")
    private Output<String> ref;

    /**
     * @return The name of branch, tag or commit.
     * 
     */
    public Output<String> ref() {
        return this.ref;
    }
    /**
     * The file size.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The file size.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * Name of the branch to start the new commit from.
     * 
     */
    @Export(name="startBranch", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> startBranch;

    /**
     * @return Name of the branch to start the new commit from.
     * 
     */
    public Output<Optional<String>> startBranch() {
        return Codegen.optional(this.startBranch);
    }
    /**
     * Update commit message.
     * 
     */
    @Export(name="updateCommitMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> updateCommitMessage;

    /**
     * @return Update commit message.
     * 
     */
    public Output<Optional<String>> updateCommitMessage() {
        return Codegen.optional(this.updateCommitMessage);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryFile(String name) {
        this(name, RepositoryFileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryFile(String name, RepositoryFileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryFile(String name, RepositoryFileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/repositoryFile:RepositoryFile", name, args == null ? RepositoryFileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryFile(String name, Output<String> id, @Nullable RepositoryFileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/repositoryFile:RepositoryFile", name, state, makeResourceOptions(options, id));
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
    public static RepositoryFile get(String name, Output<String> id, @Nullable RepositoryFileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryFile(name, id, state, options);
    }
}
