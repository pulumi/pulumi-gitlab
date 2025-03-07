// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.UserCustomAttributeArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.UserCustomAttributeState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.UserCustomAttribute` resource allows to manage custom attributes for a user.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/custom_attributes/)
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
 * import com.pulumi.gitlab.UserCustomAttribute;
 * import com.pulumi.gitlab.UserCustomAttributeArgs;
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
 *         var attr = new UserCustomAttribute("attr", UserCustomAttributeArgs.builder()
 *             .user("42")
 *             .key("location")
 *             .value("Greenland")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_user_custom_attribute`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_user_custom_attribute.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import a user custom attribute using an id made up of `{user-id}:{key}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/userCustomAttribute:UserCustomAttribute attr 42:location
 * ```
 * 
 */
@ResourceType(type="gitlab:index/userCustomAttribute:UserCustomAttribute")
public class UserCustomAttribute extends com.pulumi.resources.CustomResource {
    /**
     * Key for the Custom Attribute.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Key for the Custom Attribute.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The id of the user.
     * 
     */
    @Export(name="user", refs={Integer.class}, tree="[0]")
    private Output<Integer> user;

    /**
     * @return The id of the user.
     * 
     */
    public Output<Integer> user() {
        return this.user;
    }
    /**
     * Value for the Custom Attribute.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return Value for the Custom Attribute.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserCustomAttribute(java.lang.String name) {
        this(name, UserCustomAttributeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserCustomAttribute(java.lang.String name, UserCustomAttributeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserCustomAttribute(java.lang.String name, UserCustomAttributeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/userCustomAttribute:UserCustomAttribute", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private UserCustomAttribute(java.lang.String name, Output<java.lang.String> id, @Nullable UserCustomAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/userCustomAttribute:UserCustomAttribute", name, state, makeResourceOptions(options, id), false);
    }

    private static UserCustomAttributeArgs makeArgs(UserCustomAttributeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserCustomAttributeArgs.Empty : args;
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
    public static UserCustomAttribute get(java.lang.String name, Output<java.lang.String> id, @Nullable UserCustomAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserCustomAttribute(name, id, state, options);
    }
}
