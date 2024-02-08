// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupLabelArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupLabelState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupLabel` resource allows to manage the lifecycle of labels within a group.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/labels.html#group-labels)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.GroupLabel;
 * import com.pulumi.gitlab.GroupLabelArgs;
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
 *         var fixme = new GroupLabel(&#34;fixme&#34;, GroupLabelArgs.builder()        
 *             .color(&#34;#ffcc00&#34;)
 *             .description(&#34;issue with failing tests&#34;)
 *             .group(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Gitlab group labels can be imported using an id made up of `{group_id}:{group_label_id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupLabel:GroupLabel example 12345:fixme
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupLabel:GroupLabel")
public class GroupLabel extends com.pulumi.resources.CustomResource {
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
     * The name or id of the group to add the label to.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The name or id of the group to add the label to.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * The id of the group label.
     * 
     */
    @Export(name="labelId", refs={Integer.class}, tree="[0]")
    private Output<Integer> labelId;

    /**
     * @return The id of the group label.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupLabel(String name) {
        this(name, GroupLabelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupLabel(String name, GroupLabelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupLabel(String name, GroupLabelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupLabel:GroupLabel", name, args == null ? GroupLabelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupLabel(String name, Output<String> id, @Nullable GroupLabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupLabel:GroupLabel", name, state, makeResourceOptions(options, id));
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
    public static GroupLabel get(String name, Output<String> id, @Nullable GroupLabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupLabel(name, id, state, options);
    }
}
