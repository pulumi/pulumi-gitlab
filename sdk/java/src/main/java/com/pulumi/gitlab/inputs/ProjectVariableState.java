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


public final class ProjectVariableState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectVariableState Empty = new ProjectVariableState();

    /**
     * The description of the variable.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the variable.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     * 
     */
    @Import(name="environmentScope")
    private @Nullable Output<String> environmentScope;

    /**
     * @return The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
     * 
     */
    public Optional<Output<String>> environmentScope() {
        return Optional.ofNullable(this.environmentScope);
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
     * If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     * 
     */
    @Import(name="masked")
    private @Nullable Output<Boolean> masked;

    /**
     * @return If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
     * 
     */
    public Optional<Output<Boolean>> masked() {
        return Optional.ofNullable(this.masked);
    }

    /**
     * The name or id of the project.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The name or id of the project.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     * 
     */
    @Import(name="protected")
    private @Nullable Output<Boolean> protected_;

    /**
     * @return If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
     * 
     */
    public Optional<Output<Boolean>> protected_() {
        return Optional.ofNullable(this.protected_);
    }

    /**
     * Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
     * 
     */
    @Import(name="raw")
    private @Nullable Output<Boolean> raw;

    /**
     * @return Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
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
     * The type of a variable. Valid values are: `env_var`, `file`.
     * 
     */
    @Import(name="variableType")
    private @Nullable Output<String> variableType;

    /**
     * @return The type of a variable. Valid values are: `env_var`, `file`.
     * 
     */
    public Optional<Output<String>> variableType() {
        return Optional.ofNullable(this.variableType);
    }

    private ProjectVariableState() {}

    private ProjectVariableState(ProjectVariableState $) {
        this.description = $.description;
        this.environmentScope = $.environmentScope;
        this.key = $.key;
        this.masked = $.masked;
        this.project = $.project;
        this.protected_ = $.protected_;
        this.raw = $.raw;
        this.value = $.value;
        this.variableType = $.variableType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectVariableState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectVariableState $;

        public Builder() {
            $ = new ProjectVariableState();
        }

        public Builder(ProjectVariableState defaults) {
            $ = new ProjectVariableState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the variable.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the variable.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param environmentScope The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
         * 
         * @return builder
         * 
         */
        public Builder environmentScope(@Nullable Output<String> environmentScope) {
            $.environmentScope = environmentScope;
            return this;
        }

        /**
         * @param environmentScope The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
         * 
         * @return builder
         * 
         */
        public Builder environmentScope(String environmentScope) {
            return environmentScope(Output.of(environmentScope));
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
         * @param masked If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
         * 
         * @return builder
         * 
         */
        public Builder masked(@Nullable Output<Boolean> masked) {
            $.masked = masked;
            return this;
        }

        /**
         * @param masked If set to `true`, the value of the variable will be masked in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#mask-a-cicd-variable).
         * 
         * @return builder
         * 
         */
        public Builder masked(Boolean masked) {
            return masked(Output.of(masked));
        }

        /**
         * @param project The name or id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param protected_ If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
         * 
         * @return builder
         * 
         */
        public Builder protected_(@Nullable Output<Boolean> protected_) {
            $.protected_ = protected_;
            return this;
        }

        /**
         * @param protected_ If set to `true`, the variable will be passed only to pipelines running on protected branches and tags.
         * 
         * @return builder
         * 
         */
        public Builder protected_(Boolean protected_) {
            return protected_(Output.of(protected_));
        }

        /**
         * @param raw Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
         * 
         * @return builder
         * 
         */
        public Builder raw(@Nullable Output<Boolean> raw) {
            $.raw = raw;
            return this;
        }

        /**
         * @param raw Whether the variable is treated as a raw string. When true, variables in the value are not expanded.
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
         * @param variableType The type of a variable. Valid values are: `env_var`, `file`.
         * 
         * @return builder
         * 
         */
        public Builder variableType(@Nullable Output<String> variableType) {
            $.variableType = variableType;
            return this;
        }

        /**
         * @param variableType The type of a variable. Valid values are: `env_var`, `file`.
         * 
         * @return builder
         * 
         */
        public Builder variableType(String variableType) {
            return variableType(Output.of(variableType));
        }

        public ProjectVariableState build() {
            return $;
        }
    }

}
