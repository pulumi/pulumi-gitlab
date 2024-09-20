// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetPipelineSchedulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPipelineSchedulesPlainArgs Empty = new GetPipelineSchedulesPlainArgs();

    /**
     * The name or id of the project to add the schedule to.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The name or id of the project to add the schedule to.
     * 
     */
    public String project() {
        return this.project;
    }

    private GetPipelineSchedulesPlainArgs() {}

    private GetPipelineSchedulesPlainArgs(GetPipelineSchedulesPlainArgs $) {
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPipelineSchedulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPipelineSchedulesPlainArgs $;

        public Builder() {
            $ = new GetPipelineSchedulesPlainArgs();
        }

        public Builder(GetPipelineSchedulesPlainArgs defaults) {
            $ = new GetPipelineSchedulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The name or id of the project to add the schedule to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public GetPipelineSchedulesPlainArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetPipelineSchedulesPlainArgs", "project");
            }
            return $;
        }
    }

}