// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetProjectHookPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectHookPlainArgs Empty = new GetProjectHookPlainArgs();

    /**
     * The id of the project hook.
     * 
     */
    @Import(name="hookId", required=true)
    private Integer hookId;

    /**
     * @return The id of the project hook.
     * 
     */
    public Integer hookId() {
        return this.hookId;
    }

    /**
     * The name or id of the project to add the hook to.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The name or id of the project to add the hook to.
     * 
     */
    public String project() {
        return this.project;
    }

    private GetProjectHookPlainArgs() {}

    private GetProjectHookPlainArgs(GetProjectHookPlainArgs $) {
        this.hookId = $.hookId;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectHookPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectHookPlainArgs $;

        public Builder() {
            $ = new GetProjectHookPlainArgs();
        }

        public Builder(GetProjectHookPlainArgs defaults) {
            $ = new GetProjectHookPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hookId The id of the project hook.
         * 
         * @return builder
         * 
         */
        public Builder hookId(Integer hookId) {
            $.hookId = hookId;
            return this;
        }

        /**
         * @param project The name or id of the project to add the hook to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public GetProjectHookPlainArgs build() {
            if ($.hookId == null) {
                throw new MissingRequiredPropertyException("GetProjectHookPlainArgs", "hookId");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectHookPlainArgs", "project");
            }
            return $;
        }
    }

}
