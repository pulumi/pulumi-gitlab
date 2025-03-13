// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectShareGroupArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectShareGroupState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectShareGroup` resource allows to manage the lifecycle of project shared with a group.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#share-project-with-group)
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
 * import com.pulumi.gitlab.ProjectShareGroup;
 * import com.pulumi.gitlab.ProjectShareGroupArgs;
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
 *         var test = new ProjectShareGroup("test", ProjectShareGroupArgs.builder()
 *             .project("12345")
 *             .groupId(1337)
 *             .groupAccess("guest")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_share_group`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_share_group.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * GitLab project group shares can be imported using an id made up of `projectid:groupid`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectShareGroup:ProjectShareGroup test 12345:1337
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectShareGroup:ProjectShareGroup")
public class ProjectShareGroup extends com.pulumi.resources.CustomResource {
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     * @deprecated
     * Use `group_access` instead of the `access_level` attribute.
     * 
     */
    @Deprecated /* Use `group_access` instead of the `access_level` attribute. */
    @Export(name="accessLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessLevel;

    /**
     * @return The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    public Output<Optional<String>> accessLevel() {
        return Codegen.optional(this.accessLevel);
    }
    /**
     * The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    @Export(name="groupAccess", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupAccess;

    /**
     * @return The access level to grant the group for the project. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    public Output<Optional<String>> groupAccess() {
        return Codegen.optional(this.groupAccess);
    }
    /**
     * The id of the group.
     * 
     */
    @Export(name="groupId", refs={Integer.class}, tree="[0]")
    private Output<Integer> groupId;

    /**
     * @return The id of the group.
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectShareGroup(java.lang.String name) {
        this(name, ProjectShareGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectShareGroup(java.lang.String name, ProjectShareGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectShareGroup(java.lang.String name, ProjectShareGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectShareGroup:ProjectShareGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectShareGroup(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectShareGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectShareGroup:ProjectShareGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectShareGroupArgs makeArgs(ProjectShareGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectShareGroupArgs.Empty : args;
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
    public static ProjectShareGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectShareGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectShareGroup(name, id, state, options);
    }
}
