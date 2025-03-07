// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.LabelArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.LabelState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.Label` resource allows to manage the lifecycle of a project label.
 * 
 * &gt; This resource is deprecated. use `gitlab.ProjectLabel`instead!
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/labels/#project-labels)
 * 
 */
@ResourceType(type="gitlab:index/label:Label")
public class Label extends com.pulumi.resources.CustomResource {
    /**
     * The color of the label given in 6-digit hex notation with leading &#39;#&#39; sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     * 
     */
    @Export(name="color", refs={String.class}, tree="[0]")
    private Output<String> color;

    /**
     * @return The color of the label given in 6-digit hex notation with leading &#39;#&#39; sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     * 
     */
    public Output<String> color() {
        return this.color;
    }
    /**
     * The description of the label.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the label.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The id of the project label.
     * 
     */
    @Export(name="labelId", refs={Integer.class}, tree="[0]")
    private Output<Integer> labelId;

    /**
     * @return The id of the project label.
     * 
     */
    public Output<Integer> labelId() {
        return this.labelId;
    }
    /**
     * The name of the label.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the label.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name or id of the project to add the label to.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name or id of the project to add the label to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Label(java.lang.String name) {
        this(name, LabelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Label(java.lang.String name, LabelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Label(java.lang.String name, LabelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/label:Label", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Label(java.lang.String name, Output<java.lang.String> id, @Nullable LabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/label:Label", name, state, makeResourceOptions(options, id), false);
    }

    private static LabelArgs makeArgs(LabelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LabelArgs.Empty : args;
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
    public static Label get(java.lang.String name, Output<java.lang.String> id, @Nullable LabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Label(name, id, state, options);
    }
}
