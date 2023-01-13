// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetProjectHooksPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectHooksPlainArgs Empty = new GetProjectHooksPlainArgs();

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

    private GetProjectHooksPlainArgs() {}

    private GetProjectHooksPlainArgs(GetProjectHooksPlainArgs $) {
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectHooksPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectHooksPlainArgs $;

        public Builder() {
            $ = new GetProjectHooksPlainArgs();
        }

        public Builder(GetProjectHooksPlainArgs defaults) {
            $ = new GetProjectHooksPlainArgs(Objects.requireNonNull(defaults));
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

        public GetProjectHooksPlainArgs build() {
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            return $;
        }
    }

}