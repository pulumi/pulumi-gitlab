// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectVariablesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectVariablesPlainArgs Empty = new GetProjectVariablesPlainArgs();

    /**
     * The environment scope of the variable. Defaults to all environment (`*`).
     * 
     */
    @Import(name="environmentScope")
    private @Nullable String environmentScope;

    /**
     * @return The environment scope of the variable. Defaults to all environment (`*`).
     * 
     */
    public Optional<String> environmentScope() {
        return Optional.ofNullable(this.environmentScope);
    }

    /**
     * The name or id of the project.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The name or id of the project.
     * 
     */
    public String project() {
        return this.project;
    }

    private GetProjectVariablesPlainArgs() {}

    private GetProjectVariablesPlainArgs(GetProjectVariablesPlainArgs $) {
        this.environmentScope = $.environmentScope;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectVariablesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectVariablesPlainArgs $;

        public Builder() {
            $ = new GetProjectVariablesPlainArgs();
        }

        public Builder(GetProjectVariablesPlainArgs defaults) {
            $ = new GetProjectVariablesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param environmentScope The environment scope of the variable. Defaults to all environment (`*`).
         * 
         * @return builder
         * 
         */
        public Builder environmentScope(@Nullable String environmentScope) {
            $.environmentScope = environmentScope;
            return this;
        }

        /**
         * @param project The name or id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public GetProjectVariablesPlainArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectVariablesPlainArgs", "project");
            }
            return $;
        }
    }

}
