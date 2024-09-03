// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupProjectFileTemplateArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupProjectFileTemplateState;
import java.lang.Integer;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupProjectFileTemplate` resource allows setting a project from which
 * custom file templates will be loaded. In order to use this resource, the project selected must be a direct child of
 * the group selected. After the resource has run, `gitlab_project_template.template_project_id` is available for use.
 * For more information about which file types are available as templates, view
 * [GitLab&#39;s documentation](https://docs.gitlab.com/ee/user/group/custom_project_templates.html)
 * 
 * &gt; This resource requires a GitLab Enterprise instance with a Premium license.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#update-group)
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
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.GroupProjectFileTemplate;
 * import com.pulumi.gitlab.GroupProjectFileTemplateArgs;
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
 *         var foo = new Group("foo", GroupArgs.builder()
 *             .name("group")
 *             .path("group")
 *             .description("An example group")
 *             .build());
 * 
 *         var bar = new Project("bar", ProjectArgs.builder()
 *             .name("template project")
 *             .description("contains file templates")
 *             .visibilityLevel("public")
 *             .namespaceId(foo.id())
 *             .build());
 * 
 *         var templateLink = new GroupProjectFileTemplate("templateLink", GroupProjectFileTemplateArgs.builder()
 *             .groupId(foo.id())
 *             .fileTemplateProjectId(bar.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate")
public class GroupProjectFileTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the project that will be used for file templates. This project must be the direct
     * child of the project defined by the group_id
     * 
     */
    @Export(name="fileTemplateProjectId", refs={Integer.class}, tree="[0]")
    private Output<Integer> fileTemplateProjectId;

    /**
     * @return The ID of the project that will be used for file templates. This project must be the direct
     * child of the project defined by the group_id
     * 
     */
    public Output<Integer> fileTemplateProjectId() {
        return this.fileTemplateProjectId;
    }
    /**
     * The ID of the group that will use the file template project. This group must be the direct
     * parent of the project defined by project_id
     * 
     */
    @Export(name="groupId", refs={Integer.class}, tree="[0]")
    private Output<Integer> groupId;

    /**
     * @return The ID of the group that will use the file template project. This group must be the direct
     * parent of the project defined by project_id
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupProjectFileTemplate(java.lang.String name) {
        this(name, GroupProjectFileTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupProjectFileTemplate(java.lang.String name, GroupProjectFileTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupProjectFileTemplate(java.lang.String name, GroupProjectFileTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupProjectFileTemplate(java.lang.String name, Output<java.lang.String> id, @Nullable GroupProjectFileTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupProjectFileTemplateArgs makeArgs(GroupProjectFileTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupProjectFileTemplateArgs.Empty : args;
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
    public static GroupProjectFileTemplate get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupProjectFileTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupProjectFileTemplate(name, id, state, options);
    }
}
