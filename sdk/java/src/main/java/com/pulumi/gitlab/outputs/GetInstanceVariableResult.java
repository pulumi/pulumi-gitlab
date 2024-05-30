// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstanceVariableResult {
    /**
     * @return The description of the variable. Maximum of 255 characters.
     * 
     */
    private String description;
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
     * @return Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     * 
     */
    private Boolean raw;
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
     * @return The description of the variable. Maximum of 255 characters.
     * 
     */
    public String description() {
        return this.description;
    }
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
     * @return Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
     * 
     */
    public Boolean raw() {
        return this.raw;
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
        private String description;
        private String id;
        private String key;
        private Boolean masked;
        private Boolean protected_;
        private Boolean raw;
        private String value;
        private String variableType;
        public Builder() {}
        public Builder(GetInstanceVariableResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.masked = defaults.masked;
    	      this.protected_ = defaults.protected_;
    	      this.raw = defaults.raw;
    	      this.value = defaults.value;
    	      this.variableType = defaults.variableType;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder masked(Boolean masked) {
            if (masked == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "masked");
            }
            this.masked = masked;
            return this;
        }
        @CustomType.Setter("protected")
        public Builder protected_(Boolean protected_) {
            if (protected_ == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "protected_");
            }
            this.protected_ = protected_;
            return this;
        }
        @CustomType.Setter
        public Builder raw(Boolean raw) {
            if (raw == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "raw");
            }
            this.raw = raw;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "value");
            }
            this.value = value;
            return this;
        }
        @CustomType.Setter
        public Builder variableType(String variableType) {
            if (variableType == null) {
              throw new MissingRequiredPropertyException("GetInstanceVariableResult", "variableType");
            }
            this.variableType = variableType;
            return this;
        }
        public GetInstanceVariableResult build() {
            final var _resultValue = new GetInstanceVariableResult();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.key = key;
            _resultValue.masked = masked;
            _resultValue.protected_ = protected_;
            _resultValue.raw = raw;
            _resultValue.value = value;
            _resultValue.variableType = variableType;
            return _resultValue;
        }
    }
}
