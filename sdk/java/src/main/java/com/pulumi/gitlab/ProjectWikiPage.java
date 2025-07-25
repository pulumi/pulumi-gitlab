// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectWikiPageArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectWikiPageState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectWikiPage` resource allows managing the lifecycle of a project wiki page.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/wikis/)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_wiki_page`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_wiki_page.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
 * 
 * You can import gitlab_project_wiki_page state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectWikiPage:ProjectWikiPage test 12345:my-wiki-page
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectWikiPage:ProjectWikiPage")
public class ProjectWikiPage extends com.pulumi.resources.CustomResource {
    /**
     * Content of the wiki page. Must be at least 1 character long.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return Content of the wiki page. Must be at least 1 character long.
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * The encoding used for the wiki page content.
     * 
     */
    @Export(name="encoding", refs={String.class}, tree="[0]")
    private Output<String> encoding;

    /**
     * @return The encoding used for the wiki page content.
     * 
     */
    public Output<String> encoding() {
        return this.encoding;
    }
    /**
     * Format of the wiki page (auto-generated if not provided). Valid values are: `markdown`, `rdoc`, `asciidoc`, `org`.
     * 
     */
    @Export(name="format", refs={String.class}, tree="[0]")
    private Output<String> format;

    /**
     * @return Format of the wiki page (auto-generated if not provided). Valid values are: `markdown`, `rdoc`, `asciidoc`, `org`.
     * 
     */
    public Output<String> format() {
        return this.format;
    }
    /**
     * The ID or URL-encoded path of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or URL-encoded path of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Slug of the wiki page.
     * 
     */
    @Export(name="slug", refs={String.class}, tree="[0]")
    private Output<String> slug;

    /**
     * @return Slug of the wiki page.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }
    /**
     * Title of the wiki page.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Title of the wiki page.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectWikiPage(java.lang.String name) {
        this(name, ProjectWikiPageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectWikiPage(java.lang.String name, ProjectWikiPageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectWikiPage(java.lang.String name, ProjectWikiPageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectWikiPage:ProjectWikiPage", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectWikiPage(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectWikiPageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectWikiPage:ProjectWikiPage", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectWikiPageArgs makeArgs(ProjectWikiPageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectWikiPageArgs.Empty : args;
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
    public static ProjectWikiPage get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectWikiPageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectWikiPage(name, id, state, options);
    }
}
