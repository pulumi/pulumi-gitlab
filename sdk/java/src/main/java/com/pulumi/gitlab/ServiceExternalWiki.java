// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ServiceExternalWikiArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ServiceExternalWikiState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ServiceExternalWiki` resource allows to manage the lifecycle of a project integration with External Wiki Service.
 * 
 * &gt; This resource is deprecated. use `gitlab.IntegrationExternalWiki`instead!
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#external-wiki)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.ServiceExternalWiki;
 * import com.pulumi.gitlab.ServiceExternalWikiArgs;
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
 *         var wiki = new ServiceExternalWiki(&#34;wiki&#34;, ServiceExternalWikiArgs.builder()        
 *             .project(awesomeProject.id())
 *             .externalWikiUrl(&#34;https://MyAwesomeExternalWikiURL.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * You can import a gitlab_service_external_wiki state using the project ID, e.g.
 * 
 * ```sh
 *  $ pulumi import gitlab:index/serviceExternalWiki:ServiceExternalWiki wiki 1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/serviceExternalWiki:ServiceExternalWiki")
public class ServiceExternalWiki extends com.pulumi.resources.CustomResource {
    /**
     * Whether the integration is active.
     * 
     */
    @Export(name="active", type=Boolean.class, parameters={})
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
    @Export(name="createdAt", type=String.class, parameters={})
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
    @Export(name="externalWikiUrl", type=String.class, parameters={})
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
    @Export(name="project", type=String.class, parameters={})
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
    @Export(name="slug", type=String.class, parameters={})
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
    @Export(name="title", type=String.class, parameters={})
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
    @Export(name="updatedAt", type=String.class, parameters={})
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
    public ServiceExternalWiki(String name) {
        this(name, ServiceExternalWikiArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceExternalWiki(String name, ServiceExternalWikiArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceExternalWiki(String name, ServiceExternalWikiArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/serviceExternalWiki:ServiceExternalWiki", name, args == null ? ServiceExternalWikiArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceExternalWiki(String name, Output<String> id, @Nullable ServiceExternalWikiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/serviceExternalWiki:ServiceExternalWiki", name, state, makeResourceOptions(options, id));
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
    public static ServiceExternalWiki get(String name, Output<String> id, @Nullable ServiceExternalWikiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceExternalWiki(name, id, state, options);
    }
}
