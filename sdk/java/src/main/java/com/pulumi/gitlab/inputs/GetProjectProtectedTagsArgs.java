// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetProjectProtectedTagsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectProtectedTagsArgs Empty = new GetProjectProtectedTagsArgs();

    /**
     * The integer or path with namespace that uniquely identifies the project.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The integer or path with namespace that uniquely identifies the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private GetProjectProtectedTagsArgs() {}

    private GetProjectProtectedTagsArgs(GetProjectProtectedTagsArgs $) {
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectProtectedTagsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectProtectedTagsArgs $;

        public Builder() {
            $ = new GetProjectProtectedTagsArgs();
        }

        public Builder(GetProjectProtectedTagsArgs defaults) {
            $ = new GetProjectProtectedTagsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The integer or path with namespace that uniquely identifies the project.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The integer or path with namespace that uniquely identifies the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetProjectProtectedTagsArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedTagsArgs", "project");
            }
            return $;
        }
    }

}
