// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceVariableState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceVariableState Empty = new InstanceVariableState();

    /**
     * The description of the variable. Maximum of 255 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the variable. Maximum of 255 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the variable.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The name of the variable.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#masked-variables). Defaults to `false`.
     * 
     */
    @Import(name="masked")
    private @Nullable Output<Boolean> masked;

    /**
     * @return If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#masked-variables). Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> masked() {
        return Optional.ofNullable(this.masked);
    }

    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     * 
     */
    @Import(name="protected")
    private @Nullable Output<Boolean> protected_;

    /**
     * @return If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> protected_() {
        return Optional.ofNullable(this.protected_);
    }

    /**
     * Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     * 
     */
    @Import(name="raw")
    private @Nullable Output<Boolean> raw;

    /**
     * @return Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     * 
     */
    public Optional<Output<Boolean>> raw() {
        return Optional.ofNullable(this.raw);
    }

    /**
     * The value of the variable.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value of the variable.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    /**
     * The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    @Import(name="variableType")
    private @Nullable Output<String> variableType;

    /**
     * @return The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    public Optional<Output<String>> variableType() {
        return Optional.ofNullable(this.variableType);
    }

    private InstanceVariableState() {}

    private InstanceVariableState(InstanceVariableState $) {
        this.description = $.description;
        this.key = $.key;
        this.masked = $.masked;
        this.protected_ = $.protected_;
        this.raw = $.raw;
        this.value = $.value;
        this.variableType = $.variableType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceVariableState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceVariableState $;

        public Builder() {
            $ = new InstanceVariableState();
        }

        public Builder(InstanceVariableState defaults) {
            $ = new InstanceVariableState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the variable. Maximum of 255 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the variable. Maximum of 255 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param key The name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param masked If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#masked-variables). Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder masked(@Nullable Output<Boolean> masked) {
            $.masked = masked;
            return this;
        }

        /**
         * @param masked If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ci/variables/#masked-variables). Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder masked(Boolean masked) {
            return masked(Output.of(masked));
        }

        /**
         * @param protected_ If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder protected_(@Nullable Output<Boolean> protected_) {
            $.protected_ = protected_;
            return this;
        }

        /**
         * @param protected_ If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder protected_(Boolean protected_) {
            return protected_(Output.of(protected_));
        }

        /**
         * @param raw Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
         * 
         * @return builder
         * 
         */
        public Builder raw(@Nullable Output<Boolean> raw) {
            $.raw = raw;
            return this;
        }

        /**
         * @param raw Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
         * 
         * @return builder
         * 
         */
        public Builder raw(Boolean raw) {
            return raw(Output.of(raw));
        }

        /**
         * @param value The value of the variable.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of the variable.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        /**
         * @param variableType The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
         * 
         * @return builder
         * 
         */
        public Builder variableType(@Nullable Output<String> variableType) {
            $.variableType = variableType;
            return this;
        }

        /**
         * @param variableType The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
         * 
         * @return builder
         * 
         */
        public Builder variableType(String variableType) {
            return variableType(Output.of(variableType));
        }

        public InstanceVariableState build() {
            return $;
        }
    }

}
