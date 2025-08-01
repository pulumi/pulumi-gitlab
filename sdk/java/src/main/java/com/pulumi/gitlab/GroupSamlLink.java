// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupSamlLinkArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupSamlLinkState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupSamlLink` resource allows to manage the lifecycle of an SAML integration with a group.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/saml/#saml-group-links)
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
 * import com.pulumi.gitlab.GroupSamlLink;
 * import com.pulumi.gitlab.GroupSamlLinkArgs;
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
 *         // Basic example
 *         var test = new GroupSamlLink("test", GroupSamlLinkArgs.builder()
 *             .group("12345")
 *             .accessLevel("developer")
 *             .samlGroupName("samlgroupname1")
 *             .build());
 * 
 *         // Example using a Custom Role (Ultimate only)
 *         // When using the custom role, the `access_level` must match the
 *         // base role used to create the custom role.
 *         var testCustomRole = new GroupSamlLink("testCustomRole", GroupSamlLinkArgs.builder()
 *             .group("12345")
 *             .accessLevel("developer")
 *             .samlGroupName("samlgroupname1")
 *             .memberRoleId(123)
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
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_saml_link`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_group_saml_link.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
 * 
 * GitLab group saml links can be imported using an id made up of `group_id:saml_group_name`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupSamlLink:GroupSamlLink test &#34;12345:samlgroupname1&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupSamlLink:GroupSamlLink")
public class GroupSamlLink extends com.pulumi.resources.CustomResource {
    /**
     * Access level for members of the SAML group. Valid values are: `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output<String> accessLevel;

    /**
     * @return Access level for members of the SAML group. Valid values are: `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
    /**
     * The ID or path of the group to add the SAML Group Link to.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or path of the group to add the SAML Group Link to.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * The ID of a custom member role. Only available for Ultimate instances. When using a custom role, the `access_level` must match the base role used to create the custom role.
     * 
     */
    @Export(name="memberRoleId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memberRoleId;

    /**
     * @return The ID of a custom member role. Only available for Ultimate instances. When using a custom role, the `access_level` must match the base role used to create the custom role.
     * 
     */
    public Output<Optional<Integer>> memberRoleId() {
        return Codegen.optional(this.memberRoleId);
    }
    /**
     * The name of the SAML group.
     * 
     */
    @Export(name="samlGroupName", refs={String.class}, tree="[0]")
    private Output<String> samlGroupName;

    /**
     * @return The name of the SAML group.
     * 
     */
    public Output<String> samlGroupName() {
        return this.samlGroupName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupSamlLink(java.lang.String name) {
        this(name, GroupSamlLinkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupSamlLink(java.lang.String name, GroupSamlLinkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupSamlLink(java.lang.String name, GroupSamlLinkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupSamlLink:GroupSamlLink", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupSamlLink(java.lang.String name, Output<java.lang.String> id, @Nullable GroupSamlLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupSamlLink:GroupSamlLink", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupSamlLinkArgs makeArgs(GroupSamlLinkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupSamlLinkArgs.Empty : args;
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
    public static GroupSamlLink get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupSamlLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupSamlLink(name, id, state, options);
    }
}
