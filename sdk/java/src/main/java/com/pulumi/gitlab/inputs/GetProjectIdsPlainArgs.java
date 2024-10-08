// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetProjectIdsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectIdsPlainArgs Empty = new GetProjectIdsPlainArgs();

    /**
     * The ID or URL-encoded path of the project.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The ID or URL-encoded path of the project.
     * 
     */
    public String project() {
        return this.project;
    }

    private GetProjectIdsPlainArgs() {}

    private GetProjectIdsPlainArgs(GetProjectIdsPlainArgs $) {
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectIdsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectIdsPlainArgs $;

        public Builder() {
            $ = new GetProjectIdsPlainArgs();
        }

        public Builder(GetProjectIdsPlainArgs defaults) {
            $ = new GetProjectIdsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The ID or URL-encoded path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public GetProjectIdsPlainArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectIdsPlainArgs", "project");
            }
            return $;
        }
    }

}
