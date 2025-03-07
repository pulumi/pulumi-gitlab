// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectCustomAttributeArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectCustomAttributeState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectCustomAttribute` resource allows to manage custom attributes for a project.
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
 * import com.pulumi.gitlab.ProjectCustomAttribute;
 * import com.pulumi.gitlab.ProjectCustomAttributeArgs;
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
 *         var attr = new ProjectCustomAttribute("attr", ProjectCustomAttributeArgs.builder()
 *             .project("42")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_custom_attribute`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_custom_attribute.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import a project custom attribute using an id made up of `{project-id}:{key}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectCustomAttribute:ProjectCustomAttribute attr 42:location
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectCustomAttribute:ProjectCustomAttribute")
public class ProjectCustomAttribute extends com.pulumi.resources.CustomResource {
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
     * The id of the project.
     * 
     */
    @Export(name="project", refs={Integer.class}, tree="[0]")
    private Output<Integer> project;

    /**
     * @return The id of the project.
     * 
     */
    public Output<Integer> project() {
        return this.project;
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
    public ProjectCustomAttribute(String name) {
        this(name, ProjectCustomAttributeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectCustomAttribute(String name, ProjectCustomAttributeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectCustomAttribute(String name, ProjectCustomAttributeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectCustomAttribute:ProjectCustomAttribute", name, args == null ? ProjectCustomAttributeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectCustomAttribute(String name, Output<String> id, @Nullable ProjectCustomAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectCustomAttribute:ProjectCustomAttribute", name, state, makeResourceOptions(options, id));
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
    public static ProjectCustomAttribute get(String name, Output<String> id, @Nullable ProjectCustomAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectCustomAttribute(name, id, state, options);
    }
}
