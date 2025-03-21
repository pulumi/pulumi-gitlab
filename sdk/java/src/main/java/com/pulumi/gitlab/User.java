// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.UserArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.UserState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.User` resource allows to manage the lifecycle of a user.
 * 
 * &gt; the provider needs to be configured with admin-level access for this resource to work.
 * 
 * &gt; You must specify either password or reset_password.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/users/)
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
 * import com.pulumi.gitlab.User;
 * import com.pulumi.gitlab.UserArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var example = new User("example", UserArgs.builder()
 *             .name("Example Foo")
 *             .username("example")
 *             .password("superPassword")
 *             .email("gitlab}{@literal @}{@code user.create")
 *             .isAdmin(true)
 *             .projectsLimit(4)
 *             .canCreateGroup(false)
 *             .isExternal(true)
 *             .resetPassword(false)
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_user`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_user.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * ```sh
 * $ pulumi import gitlab:index/user:User You can import a user to terraform state using `&lt;resource&gt; &lt;id&gt;`.
 * ```
 * 
 * The `id` must be an integer for the id of the user you want to import,
 * 
 * for example:
 * 
 * ```sh
 * $ pulumi import gitlab:index/user:User example 42
 * ```
 * 
 */
@ResourceType(type="gitlab:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     * 
     */
    @Export(name="canCreateGroup", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> canCreateGroup;

    /**
     * @return Boolean, defaults to false. Whether to allow the user to create groups.
     * 
     */
    public Output<Optional<Boolean>> canCreateGroup() {
        return Codegen.optional(this.canCreateGroup);
    }
    /**
     * The e-mail address of the user.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    /**
     * @return The e-mail address of the user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * String, a specific external authentication provider UID.
     * 
     * @deprecated
     * To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     * 
     */
    @Deprecated /* To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295 */
    @Export(name="externUid", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> externUid;

    /**
     * @return String, a specific external authentication provider UID.
     * 
     */
    public Output<Optional<String>> externUid() {
        return Codegen.optional(this.externUid);
    }
    /**
     * String, the external provider.
     * 
     * @deprecated
     * To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295
     * 
     */
    @Deprecated /* To be removed in 18.0. Use gitlab.UserIdentity resource instead. See https://gitlab.com/gitlab-org/terraform-provider-gitlab/-/issues/1295 */
    @Export(name="externalProvider", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> externalProvider;

    /**
     * @return String, the external provider.
     * 
     */
    public Output<Optional<String>> externalProvider() {
        return Codegen.optional(this.externalProvider);
    }
    /**
     * Set user password to a random value
     * 
     */
    @Export(name="forceRandomPassword", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceRandomPassword;

    /**
     * @return Set user password to a random value
     * 
     */
    public Output<Optional<Boolean>> forceRandomPassword() {
        return Codegen.optional(this.forceRandomPassword);
    }
    /**
     * Boolean, defaults to false.  Whether to enable administrative privileges
     * 
     */
    @Export(name="isAdmin", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isAdmin;

    /**
     * @return Boolean, defaults to false.  Whether to enable administrative privileges
     * 
     */
    public Output<Optional<Boolean>> isAdmin() {
        return Codegen.optional(this.isAdmin);
    }
    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     * 
     */
    @Export(name="isExternal", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isExternal;

    /**
     * @return Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     * 
     */
    public Output<Optional<Boolean>> isExternal() {
        return Codegen.optional(this.isExternal);
    }
    /**
     * The name of the user.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the user.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the user&#39;s namespace.
     * 
     */
    @Export(name="namespaceId", refs={Integer.class}, tree="[0]")
    private Output<Integer> namespaceId;

    /**
     * @return The ID of the user&#39;s namespace.
     * 
     */
    public Output<Integer> namespaceId() {
        return this.namespaceId;
    }
    /**
     * The note associated to the user.
     * 
     */
    @Export(name="note", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> note;

    /**
     * @return The note associated to the user.
     * 
     */
    public Output<Optional<String>> note() {
        return Codegen.optional(this.note);
    }
    /**
     * The password of the user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password of the user.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Integer, defaults to 0.  Number of projects user can create.
     * 
     */
    @Export(name="projectsLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> projectsLimit;

    /**
     * @return Integer, defaults to 0.  Number of projects user can create.
     * 
     */
    public Output<Optional<Integer>> projectsLimit() {
        return Codegen.optional(this.projectsLimit);
    }
    /**
     * Boolean, defaults to false. Send user password reset link.
     * 
     */
    @Export(name="resetPassword", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> resetPassword;

    /**
     * @return Boolean, defaults to false. Send user password reset link.
     * 
     */
    public Output<Optional<Boolean>> resetPassword() {
        return Codegen.optional(this.resetPassword);
    }
    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     * 
     */
    @Export(name="skipConfirmation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipConfirmation;

    /**
     * @return Boolean, defaults to true. Whether to skip confirmation.
     * 
     */
    public Output<Optional<Boolean>> skipConfirmation() {
        return Codegen.optional(this.skipConfirmation);
    }
    /**
     * String, defaults to &#39;active&#39;. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return String, defaults to &#39;active&#39;. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * The username of the user.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The username of the user.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(java.lang.String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(java.lang.String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(java.lang.String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/user:User", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private User(java.lang.String name, Output<java.lang.String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/user:User", name, state, makeResourceOptions(options, id), false);
    }

    private static UserArgs makeArgs(UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static User get(java.lang.String name, Output<java.lang.String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
