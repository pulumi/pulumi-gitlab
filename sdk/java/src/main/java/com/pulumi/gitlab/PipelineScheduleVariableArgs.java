// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipelineScheduleVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipelineScheduleVariableArgs Empty = new PipelineScheduleVariableArgs();

    /**
     * Name of the variable.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return Name of the variable.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * The id of the pipeline schedule.
     * 
     */
    @Import(name="pipelineScheduleId", required=true)
    private Output<Integer> pipelineScheduleId;

    /**
     * @return The id of the pipeline schedule.
     * 
     */
    public Output<Integer> pipelineScheduleId() {
        return this.pipelineScheduleId;
    }

    /**
     * The id of the project to add the schedule to.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The id of the project to add the schedule to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Value of the variable.
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Value of the variable.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     * The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    @Import(name="variableType")
    private @Nullable Output<String> variableType;

    /**
     * @return The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
     * 
     */
    public Optional<Output<String>> variableType() {
        return Optional.ofNullable(this.variableType);
    }

    private PipelineScheduleVariableArgs() {}

    private PipelineScheduleVariableArgs(PipelineScheduleVariableArgs $) {
        this.key = $.key;
        this.pipelineScheduleId = $.pipelineScheduleId;
        this.project = $.project;
        this.value = $.value;
        this.variableType = $.variableType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipelineScheduleVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipelineScheduleVariableArgs $;

        public Builder() {
            $ = new PipelineScheduleVariableArgs();
        }

        public Builder(PipelineScheduleVariableArgs defaults) {
            $ = new PipelineScheduleVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key Name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param pipelineScheduleId The id of the pipeline schedule.
         * 
         * @return builder
         * 
         */
        public Builder pipelineScheduleId(Output<Integer> pipelineScheduleId) {
            $.pipelineScheduleId = pipelineScheduleId;
            return this;
        }

        /**
         * @param pipelineScheduleId The id of the pipeline schedule.
         * 
         * @return builder
         * 
         */
        public Builder pipelineScheduleId(Integer pipelineScheduleId) {
            return pipelineScheduleId(Output.of(pipelineScheduleId));
        }

        /**
         * @param project The id of the project to add the schedule to.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The id of the project to add the schedule to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param value Value of the variable.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value of the variable.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        /**
         * @param variableType The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
         * 
         * @return builder
         * 
         */
        public Builder variableType(@Nullable Output<String> variableType) {
            $.variableType = variableType;
            return this;
        }

        /**
         * @param variableType The type of a variable. Available types are: `env_var`, `file`. Default is `env_var`.
         * 
         * @return builder
         * 
         */
        public Builder variableType(String variableType) {
            return variableType(Output.of(variableType));
        }

        public PipelineScheduleVariableArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("PipelineScheduleVariableArgs", "key");
            }
            if ($.pipelineScheduleId == null) {
                throw new MissingRequiredPropertyException("PipelineScheduleVariableArgs", "pipelineScheduleId");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("PipelineScheduleVariableArgs", "project");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("PipelineScheduleVariableArgs", "value");
            }
            return $;
        }
    }

}
