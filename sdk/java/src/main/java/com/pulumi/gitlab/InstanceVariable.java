// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.InstanceVariableArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.InstanceVariableState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.InstanceVariable` resource allows to manage the lifecycle of an instance-level CI/CD variable.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)
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
 * import com.pulumi.gitlab.InstanceVariable;
 * import com.pulumi.gitlab.InstanceVariableArgs;
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
 *         var example = new InstanceVariable("example", InstanceVariableArgs.builder()        
 *             .key("instance_variable_key")
 *             .value("instance_variable_value")
 *             .protected_(false)
 *             .masked(false)
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
 * GitLab instance variables can be imported using an id made up of `variablename`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/instanceVariable:InstanceVariable example instance_variable_key
 * ```
 * 
 */
@ResourceType(type="gitlab:index/instanceVariable:InstanceVariable")
public class InstanceVariable extends com.pulumi.resources.CustomResource {
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
     * If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     * 
     */
    @Export(name="masked", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> masked;

    /**
     * @return If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> masked() {
        return Codegen.optional(this.masked);
    }
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     * 
     */
    @Export(name="protected", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> protected_;

    /**
     * @return If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> protected_() {
        return Codegen.optional(this.protected_);
    }
    /**
     * Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     * 
     */
    @Export(name="raw", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> raw;

    /**
     * @return Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     * 
     */
    public Output<Optional<Boolean>> raw() {
        return Codegen.optional(this.raw);
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
     * The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    @Export(name="variableType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> variableType;

    /**
     * @return The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    public Output<Optional<String>> variableType() {
        return Codegen.optional(this.variableType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceVariable(String name) {
        this(name, InstanceVariableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceVariable(String name, InstanceVariableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceVariable(String name, InstanceVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/instanceVariable:InstanceVariable", name, args == null ? InstanceVariableArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceVariable(String name, Output<String> id, @Nullable InstanceVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/instanceVariable:InstanceVariable", name, state, makeResourceOptions(options, id));
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
    public static InstanceVariable get(String name, Output<String> id, @Nullable InstanceVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceVariable(name, id, state, options);
    }
}
