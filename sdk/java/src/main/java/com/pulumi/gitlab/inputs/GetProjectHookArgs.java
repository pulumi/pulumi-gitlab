// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetProjectHookArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectHookArgs Empty = new GetProjectHookArgs();

    /**
     * The id of the project hook.
     * 
     */
    @Import(name="hookId", required=true)
    private Output<Integer> hookId;

    /**
     * @return The id of the project hook.
     * 
     */
    public Output<Integer> hookId() {
        return this.hookId;
    }

    /**
     * The name or id of the project to add the hook to.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The name or id of the project to add the hook to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private GetProjectHookArgs() {}

    private GetProjectHookArgs(GetProjectHookArgs $) {
        this.hookId = $.hookId;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectHookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectHookArgs $;

        public Builder() {
            $ = new GetProjectHookArgs();
        }

        public Builder(GetProjectHookArgs defaults) {
            $ = new GetProjectHookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hookId The id of the project hook.
         * 
         * @return builder
         * 
         */
        public Builder hookId(Output<Integer> hookId) {
            $.hookId = hookId;
            return this;
        }

        /**
         * @param hookId The id of the project hook.
         * 
         * @return builder
         * 
         */
        public Builder hookId(Integer hookId) {
            return hookId(Output.of(hookId));
        }

        /**
         * @param project The name or id of the project to add the hook to.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project to add the hook to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetProjectHookArgs build() {
            $.hookId = Objects.requireNonNull($.hookId, "expected parameter 'hookId' to be non-null");
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            return $;
        }
    }

}