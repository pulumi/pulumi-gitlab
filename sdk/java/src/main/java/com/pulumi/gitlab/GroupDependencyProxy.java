// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupDependencyProxyArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupDependencyProxyState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupDependencyProxy` resource allows managing the group docker dependency proxy. More than one dependency proxy per group will conflict with each other.
 * 
 * If you&#39;re looking to manage the project-level package dependency proxy, see the `gitlab_project_package_registry_proxy` resource instead.
 * 
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationupdatedependencyproxysettings)
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
 * import com.pulumi.gitlab.GroupDependencyProxy;
 * import com.pulumi.gitlab.GroupDependencyProxyArgs;
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
 *         var foo = new GroupDependencyProxy("foo", GroupDependencyProxyArgs.builder()
 *             .group("1234")
 *             .enabled(true)
 *             .identity("newidentity")
 *             .secret("somesecret")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_group_dependency_proxy`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_group_dependency_proxy.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import a group dependency proxy using the group id. e.g. `{group-id}`
 * 
 * &#34;secret&#34; will not populate when importing the dependency proxy, but will still
 * 
 * be required in the configuration.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupDependencyProxy:GroupDependencyProxy foo 42
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupDependencyProxy:GroupDependencyProxy")
public class GroupDependencyProxy extends com.pulumi.resources.CustomResource {
    /**
     * Indicates whether the proxy is enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Indicates whether the proxy is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The ID or URL-encoded path of the group.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
     * 
     */
    @Export(name="identity", refs={String.class}, tree="[0]")
    private Output<String> identity;

    /**
     * @return Identity credential used to authenticate with Docker Hub when pulling images. Can be a username (for password or personal access token (PAT)) or organization name (for organization access token (OAT)).
     * 
     */
    public Output<String> identity() {
        return this.identity;
    }
    /**
     * Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return Secret credential used to authenticate with Docker Hub when pulling images. Can be a password, personal access token (PAT), or organization access token (OAT). Cannot be imported.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupDependencyProxy(java.lang.String name) {
        this(name, GroupDependencyProxyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupDependencyProxy(java.lang.String name, GroupDependencyProxyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupDependencyProxy(java.lang.String name, GroupDependencyProxyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupDependencyProxy:GroupDependencyProxy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupDependencyProxy(java.lang.String name, Output<java.lang.String> id, @Nullable GroupDependencyProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupDependencyProxy:GroupDependencyProxy", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupDependencyProxyArgs makeArgs(GroupDependencyProxyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupDependencyProxyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static GroupDependencyProxy get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupDependencyProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupDependencyProxy(name, id, state, options);
    }
}
