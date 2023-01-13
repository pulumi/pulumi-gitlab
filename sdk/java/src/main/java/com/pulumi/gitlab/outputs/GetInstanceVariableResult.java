// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstanceVariableResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the variable.
     * 
     */
    private String key;
    /**
     * @return If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     * 
     */
    private Boolean masked;
    /**
     * @return If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     * 
     */
    private Boolean protected_;
    /**
     * @return The value of the variable.
     * 
     */
    private String value;
    /**
     * @return The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    private String variableType;

    private GetInstanceVariableResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the variable.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
     * 
     */
    public Boolean masked() {
        return this.masked;
    }
    /**
     * @return If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     * 
     */
    public Boolean protected_() {
        return this.protected_;
    }
    /**
     * @return The value of the variable.
     * 
     */
    public String value() {
        return this.value;
    }
    /**
     * @return The type of a variable. Valid values are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    public String variableType() {
        return this.variableType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceVariableResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String key;
        private Boolean masked;
        private Boolean protected_;
        private String value;
        private String variableType;
        public Builder() {}
        public Builder(GetInstanceVariableResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.masked = defaults.masked;
    	      this.protected_ = defaults.protected_;
    	      this.value = defaults.value;
    	      this.variableType = defaults.variableType;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder masked(Boolean masked) {
            this.masked = Objects.requireNonNull(masked);
            return this;
        }
        @CustomType.Setter("protected")
        public Builder protected_(Boolean protected_) {
            this.protected_ = Objects.requireNonNull(protected_);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        @CustomType.Setter
        public Builder variableType(String variableType) {
            this.variableType = Objects.requireNonNull(variableType);
            return this;
        }
        public GetInstanceVariableResult build() {
            final var o = new GetInstanceVariableResult();
            o.id = id;
            o.key = key;
            o.masked = masked;
            o.protected_ = protected_;
            o.value = value;
            o.variableType = variableType;
            return o;
        }
    }
}