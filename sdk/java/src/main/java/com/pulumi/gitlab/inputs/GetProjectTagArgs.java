// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetProjectTagArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectTagArgs Empty = new GetProjectTagArgs();

    /**
     * The name of a tag.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of a tag.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private GetProjectTagArgs() {}

    private GetProjectTagArgs(GetProjectTagArgs $) {
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectTagArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectTagArgs $;

        public Builder() {
            $ = new GetProjectTagArgs();
        }

        public Builder(GetProjectTagArgs defaults) {
            $ = new GetProjectTagArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of a tag.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of a tag.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID or URL-encoded path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or URL-encoded path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetProjectTagArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetProjectTagArgs", "name");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectTagArgs", "project");
            }
            return $;
        }
    }

}
