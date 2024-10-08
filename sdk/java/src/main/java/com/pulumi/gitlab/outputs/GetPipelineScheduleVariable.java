// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPipelineScheduleVariable {
    /**
     * @return The key of a variable.
     * 
     */
    private String key;
    /**
     * @return The value of a variable.
     * 
     */
    private String value;
    /**
     * @return The type of a variable, one of: env_var and file.
     * 
     */
    private String variableType;

    private GetPipelineScheduleVariable() {}
    /**
     * @return The key of a variable.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return The value of a variable.
     * 
     */
    public String value() {
        return this.value;
    }
    /**
     * @return The type of a variable, one of: env_var and file.
     * 
     */
    public String variableType() {
        return this.variableType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPipelineScheduleVariable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String key;
        private String value;
        private String variableType;
        public Builder() {}
        public Builder(GetPipelineScheduleVariable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.value = defaults.value;
    	      this.variableType = defaults.variableType;
        }

        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleVariable", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleVariable", "value");
            }
            this.value = value;
            return this;
        }
        @CustomType.Setter
        public Builder variableType(String variableType) {
            if (variableType == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleVariable", "variableType");
            }
            this.variableType = variableType;
            return this;
        }
        public GetPipelineScheduleVariable build() {
            final var _resultValue = new GetPipelineScheduleVariable();
            _resultValue.key = key;
            _resultValue.value = value;
            _resultValue.variableType = variableType;
            return _resultValue;
        }
    }
}
