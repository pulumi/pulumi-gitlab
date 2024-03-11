// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.IntegrationExternalWikiArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.IntegrationExternalWikiState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.IntegrationExternalWiki` resource allows to manage the lifecycle of a project integration with External Wiki Service.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#external-wiki)
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
 * import com.pulumi.gitlab.IntegrationExternalWiki;
 * import com.pulumi.gitlab.IntegrationExternalWikiArgs;
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
 *         var awesomeProject = new Project(&#34;awesomeProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;My awesome project.&#34;)
 *             .visibilityLevel(&#34;public&#34;)
 *             .build());
 * 
 *         var wiki = new IntegrationExternalWiki(&#34;wiki&#34;, IntegrationExternalWikiArgs.builder()        
 *             .project(awesomeProject.id())
 *             .externalWikiUrl(&#34;https://MyAwesomeExternalWikiURL.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * You can import a gitlab_integration_external_wiki state using the project ID, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/integrationExternalWiki:IntegrationExternalWiki wiki 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/integrationExternalWiki:IntegrationExternalWiki")
public class IntegrationExternalWiki extends com.pulumi.resources.CustomResource {
    /**
     * Whether the integration is active.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> active;

    /**
     * @return Whether the integration is active.
     * 
     */
    public Output<Boolean> active() {
        return this.active;
    }
    /**
     * The ISO8601 date/time that this integration was activated at in UTC.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The ISO8601 date/time that this integration was activated at in UTC.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The URL of the external wiki.
     * 
     */
    @Export(name="externalWikiUrl", refs={String.class}, tree="[0]")
    private Output<String> externalWikiUrl;

    /**
     * @return The URL of the external wiki.
     * 
     */
    public Output<String> externalWikiUrl() {
        return this.externalWikiUrl;
    }
    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    @Export(name="slug", refs={String.class}, tree="[0]")
    private Output<String> slug;

    /**
     * @return The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }
    /**
     * Title of the integration.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Title of the integration.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * The ISO8601 date/time that this integration was last updated at in UTC.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The ISO8601 date/time that this integration was last updated at in UTC.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationExternalWiki(String name) {
        this(name, IntegrationExternalWikiArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationExternalWiki(String name, IntegrationExternalWikiArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationExternalWiki(String name, IntegrationExternalWikiArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationExternalWiki:IntegrationExternalWiki", name, args == null ? IntegrationExternalWikiArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IntegrationExternalWiki(String name, Output<String> id, @Nullable IntegrationExternalWikiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/integrationExternalWiki:IntegrationExternalWiki", name, state, makeResourceOptions(options, id));
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
    public static IntegrationExternalWiki get(String name, Output<String> id, @Nullable IntegrationExternalWikiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationExternalWiki(name, id, state, options);
    }
}
