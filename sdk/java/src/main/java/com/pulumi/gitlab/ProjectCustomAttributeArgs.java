// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class ProjectCustomAttributeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectCustomAttributeArgs Empty = new ProjectCustomAttributeArgs();

    /**
     * Key for the Custom Attribute.
     * 
     */
    @Import(name="key", required=true)
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
    @Import(name="project", required=true)
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
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Value for the Custom Attribute.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private ProjectCustomAttributeArgs() {}

    private ProjectCustomAttributeArgs(ProjectCustomAttributeArgs $) {
        this.key = $.key;
        this.project = $.project;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectCustomAttributeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectCustomAttributeArgs $;

        public Builder() {
            $ = new ProjectCustomAttributeArgs();
        }

        public Builder(ProjectCustomAttributeArgs defaults) {
            $ = new ProjectCustomAttributeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key Key for the Custom Attribute.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Key for the Custom Attribute.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param project The id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<Integer> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(Integer project) {
            return project(Output.of(project));
        }

        /**
         * @param value Value for the Custom Attribute.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value for the Custom Attribute.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ProjectCustomAttributeArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}