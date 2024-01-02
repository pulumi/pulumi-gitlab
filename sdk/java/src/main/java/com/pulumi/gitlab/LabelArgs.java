// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LabelArgs extends com.pulumi.resources.ResourceArgs {

    public static final LabelArgs Empty = new LabelArgs();

    /**
     * The color of the label given in 6-digit hex notation with leading &#39;#&#39; sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
     * 
     */
    @Import(name="color", required=true)
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
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the label.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the label.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the label.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name or id of the project to add the label to.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The name or id of the project to add the label to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private LabelArgs() {}

    private LabelArgs(LabelArgs $) {
        this.color = $.color;
        this.description = $.description;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LabelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LabelArgs $;

        public Builder() {
            $ = new LabelArgs();
        }

        public Builder(LabelArgs defaults) {
            $ = new LabelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param color The color of the label given in 6-digit hex notation with leading &#39;#&#39; sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
         * 
         * @return builder
         * 
         */
        public Builder color(Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color The color of the label given in 6-digit hex notation with leading &#39;#&#39; sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
         * 
         * @return builder
         * 
         */
        public Builder color(String color) {
            return color(Output.of(color));
        }

        /**
         * @param description The description of the label.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the label.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the label.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the label.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The name or id of the project to add the label to.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project to add the label to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public LabelArgs build() {
            if ($.color == null) {
                throw new MissingRequiredPropertyException("LabelArgs", "color");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("LabelArgs", "project");
            }
            return $;
        }
    }

}
