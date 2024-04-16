// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ApplicationArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ApplicationState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The `gitlab.Application` resource allows to manage the lifecycle of applications in gitlab.
 * 
 * &gt; In order to use a user for a user to create an application, they must have admin privileges at the instance level.
 * To create an OIDC application, a scope of &#34;openid&#34;.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/applications.html)
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
 * import com.pulumi.gitlab.Application;
 * import com.pulumi.gitlab.ApplicationArgs;
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
 *         var oidc = new Application(&#34;oidc&#34;, ApplicationArgs.builder()        
 *             .confidential(true)
 *             .scopes(&#34;openid&#34;)
 *             .name(&#34;company_oidc&#34;)
 *             .redirectUrl(&#34;https://mycompany.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Gitlab applications can be imported with their id, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/application:Application example &#34;1&#34;
 * ```
 * 
 * NOTE: the secret and scopes cannot be imported
 * 
 */
@ResourceType(type="gitlab:index/application:Application")
public class Application extends com.pulumi.resources.CustomResource {
    /**
     * Internal name of the application.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return Internal name of the application.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
     * 
     */
    @Export(name="confidential", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> confidential;

    /**
     * @return The application is used where the client secret can be kept confidential. Native mobile apps and Single Page Apps are considered non-confidential. Defaults to true if not supplied
     * 
     */
    public Output<Boolean> confidential() {
        return this.confidential;
    }
    /**
     * Name of the application.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the application.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The URL gitlab should send the user to after authentication.
     * 
     */
    @Export(name="redirectUrl", refs={String.class}, tree="[0]")
    private Output<String> redirectUrl;

    /**
     * @return The URL gitlab should send the user to after authentication.
     * 
     */
    public Output<String> redirectUrl() {
        return this.redirectUrl;
    }
    /**
     * Scopes of the application. Use &#34;openid&#34; if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
     * This is only populated when creating a new application. This attribute is not available for imported resources
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return Scopes of the application. Use &#34;openid&#34; if you plan to use this as an oidc authentication application. Valid options are: `api`, `read_api`, `read_user`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`, `openid`, `profile`, `email`.
     * This is only populated when creating a new application. This attribute is not available for imported resources
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return Application secret. Sensitive and must be kept secret. This is only populated when creating a new application. This attribute is not available for imported resources.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Application(String name) {
        this(name, ApplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Application(String name, ApplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Application(String name, ApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/application:Application", name, args == null ? ApplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Application(String name, Output<String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/application:Application", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
            ))
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
    public static Application get(String name, Output<String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Application(name, id, state, options);
    }
}
