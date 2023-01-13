// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GroupCustomAttributeArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupCustomAttributeArgs Empty = new GroupCustomAttributeArgs();

    /**
     * The id of the group.
     * 
     */
    @Import(name="group", required=true)
    private Output<Integer> group;

    /**
     * @return The id of the group.
     * 
     */
    public Output<Integer> group() {
        return this.group;
    }

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

    private GroupCustomAttributeArgs() {}

    private GroupCustomAttributeArgs(GroupCustomAttributeArgs $) {
        this.group = $.group;
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupCustomAttributeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupCustomAttributeArgs $;

        public Builder() {
            $ = new GroupCustomAttributeArgs();
        }

        public Builder(GroupCustomAttributeArgs defaults) {
            $ = new GroupCustomAttributeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param group The id of the group.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<Integer> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The id of the group.
         * 
         * @return builder
         * 
         */
        public Builder group(Integer group) {
            return group(Output.of(group));
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

        public GroupCustomAttributeArgs build() {
            $.group = Objects.requireNonNull($.group, "expected parameter 'group' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}