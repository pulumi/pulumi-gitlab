// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ReleaseLinkArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ReleaseLinkState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ReleaseLink` resource allows to manage the lifecycle of a release link.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
 * import com.pulumi.gitlab.ReleaseLink;
 * import com.pulumi.gitlab.ReleaseLinkArgs;
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
 *         // Can create release link only to a tag associated with a release
 *         var exampleReleaseLink = new ReleaseLink("exampleReleaseLink", ReleaseLinkArgs.builder()
 *             .project(example.id())
 *             .tagName("tag_name_associated_with_release")
 *             .name("test")
 *             .url("https://test/")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_release_link`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_release_link.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * Gitlab release link can be imported with a key composed of `&lt;project&gt;:&lt;tag_name&gt;:&lt;link_id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/releaseLink:ReleaseLink example &#34;12345:test:2&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/releaseLink:ReleaseLink")
public class ReleaseLink extends com.pulumi.resources.CustomResource {
    /**
     * Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
     * 
     */
    @Export(name="directAssetUrl", refs={String.class}, tree="[0]")
    private Output<String> directAssetUrl;

    /**
     * @return Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
     * 
     */
    public Output<String> directAssetUrl() {
        return this.directAssetUrl;
    }
    /**
     * External or internal link.
     * 
     */
    @Export(name="external", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> external;

    /**
     * @return External or internal link.
     * 
     */
    public Output<Boolean> external() {
        return this.external;
    }
    /**
     * Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
     * 
     */
    @Export(name="filepath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> filepath;

    /**
     * @return Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
     * 
     */
    public Output<Optional<String>> filepath() {
        return Codegen.optional(this.filepath);
    }
    /**
     * The ID of the link.
     * 
     */
    @Export(name="linkId", refs={Integer.class}, tree="[0]")
    private Output<Integer> linkId;

    /**
     * @return The ID of the link.
     * 
     */
    public Output<Integer> linkId() {
        return this.linkId;
    }
    /**
     * The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
     * 
     */
    @Export(name="linkType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> linkType;

    /**
     * @return The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
     * 
     */
    public Output<Optional<String>> linkType() {
        return Codegen.optional(this.linkType);
    }
    /**
     * The name of the link. Link names must be unique within the release.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the link. Link names must be unique within the release.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The tag associated with the Release.
     * 
     */
    @Export(name="tagName", refs={String.class}, tree="[0]")
    private Output<String> tagName;

    /**
     * @return The tag associated with the Release.
     * 
     */
    public Output<String> tagName() {
        return this.tagName;
    }
    /**
     * The URL of the link. Link URLs must be unique within the release.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the link. Link URLs must be unique within the release.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReleaseLink(java.lang.String name) {
        this(name, ReleaseLinkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReleaseLink(java.lang.String name, ReleaseLinkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReleaseLink(java.lang.String name, ReleaseLinkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/releaseLink:ReleaseLink", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReleaseLink(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/releaseLink:ReleaseLink", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseLinkArgs makeArgs(ReleaseLinkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseLinkArgs.Empty : args;
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
    public static ReleaseLink get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReleaseLink(name, id, state, options);
    }
}
