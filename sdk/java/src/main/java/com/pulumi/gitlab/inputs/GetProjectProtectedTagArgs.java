// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetProjectProtectedTagArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectProtectedTagArgs Empty = new GetProjectProtectedTagArgs();

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

    /**
     * The name of the protected tag.
     * 
     */
    @Import(name="tag", required=true)
    private Output<String> tag;

    /**
     * @return The name of the protected tag.
     * 
     */
    public Output<String> tag() {
        return this.tag;
    }

    private GetProjectProtectedTagArgs() {}

    private GetProjectProtectedTagArgs(GetProjectProtectedTagArgs $) {
        this.project = $.project;
        this.tag = $.tag;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectProtectedTagArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectProtectedTagArgs $;

        public Builder() {
            $ = new GetProjectProtectedTagArgs();
        }

        public Builder(GetProjectProtectedTagArgs defaults) {
            $ = new GetProjectProtectedTagArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param tag The name of the protected tag.
         * 
         * @return builder
         * 
         */
        public Builder tag(Output<String> tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tag The name of the protected tag.
         * 
         * @return builder
         * 
         */
        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        public GetProjectProtectedTagArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedTagArgs", "project");
            }
            if ($.tag == null) {
                throw new MissingRequiredPropertyException("GetProjectProtectedTagArgs", "tag");
            }
            return $;
        }
    }

}
