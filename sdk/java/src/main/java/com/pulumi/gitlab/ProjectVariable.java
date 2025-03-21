// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectVariableArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectVariableState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectVariable` resource allows creating and managing a GitLab project level variable.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_level_variables/)
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
 * import com.pulumi.gitlab.ProjectVariable;
 * import com.pulumi.gitlab.ProjectVariableArgs;
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
 *         var example = new ProjectVariable("example", ProjectVariableArgs.builder()
 *             .project("12345")
 *             .key("project_variable_key")
 *             .value("project_variable_value")
 *             .protected_(false)
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_variable`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_variable.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * GitLab project variables can be imported using an id made up of `project:key:environment_scope`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectVariable:ProjectVariable example &#39;12345:project_variable_key:*&#39;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectVariable:ProjectVariable")
public class ProjectVariable extends com.pulumi.resources.CustomResource {
    /**
     * The description of the variable.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the variable.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     * 
     */
    @Export(name="environmentScope", refs={String.class}, tree="[0]")
    private Output<String> environmentScope;

    /**
     * @return The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     * 
     */
    public Output<String> environmentScope() {
        return this.environmentScope;
    }
    /**
     * If set to `true`, the value of the variable will be hidden in the CI/CD User Interface. The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
     * 
     */
    @Export(name="hidden", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hidden;

    /**
     * @return If set to `true`, the value of the variable will be hidden in the CI/CD User Interface. The value must meet the [hidden requirements](https://docs.gitlab.com/ci/variables/#hide-a-cicd-variable).
     * 
     */
    public Output<Boolean> hidden() {
        return this.hidden;
    }
    /**
     * The name of the variable.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The name of the variable.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     * 
     */
    @Export(name="masked", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> masked;

    /**
     * @return If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     * 
     */
    public Output<Boolean> masked() {
        return this.masked;
    }
    /**
     * The name or id of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name or id of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     * 
     */
    @Export(name="protected", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> protected_;

    /**
     * @return If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     * 
     */
    public Output<Boolean> protected_() {
        return this.protected_;
    }
    /**
     * Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
     * 
     */
    @Export(name="raw", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> raw;

    /**
     * @return Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
     * 
     */
    public Output<Boolean> raw() {
        return this.raw;
    }
    /**
     * The value of the variable.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The value of the variable.
     * 
     */
    public Output<String> value() {
        return this.value;
    }
    /**
     * The type of a variable. Valid values are: `env_var`, `file`.
     * 
     */
    @Export(name="variableType", refs={String.class}, tree="[0]")
    private Output<String> variableType;

    /**
     * @return The type of a variable. Valid values are: `env_var`, `file`.
     * 
     */
    public Output<String> variableType() {
        return this.variableType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectVariable(java.lang.String name) {
        this(name, ProjectVariableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectVariable(java.lang.String name, ProjectVariableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectVariable(java.lang.String name, ProjectVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectVariable:ProjectVariable", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectVariable(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectVariable:ProjectVariable", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectVariableArgs makeArgs(ProjectVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectVariableArgs.Empty : args;
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
    public static ProjectVariable get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectVariable(name, id, state, options);
    }
}
