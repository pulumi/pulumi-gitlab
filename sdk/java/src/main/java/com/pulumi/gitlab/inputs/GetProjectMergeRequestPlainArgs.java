// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetProjectMergeRequestPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectMergeRequestPlainArgs Empty = new GetProjectMergeRequestPlainArgs();

    /**
     * The unique project level ID of the merge request.
     * 
     */
    @Import(name="iid", required=true)
    private Integer iid;

    /**
     * @return The unique project level ID of the merge request.
     * 
     */
    public Integer iid() {
        return this.iid;
    }

    /**
     * The ID or path of the project.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The ID or path of the project.
     * 
     */
    public String project() {
        return this.project;
    }

    private GetProjectMergeRequestPlainArgs() {}

    private GetProjectMergeRequestPlainArgs(GetProjectMergeRequestPlainArgs $) {
        this.iid = $.iid;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectMergeRequestPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectMergeRequestPlainArgs $;

        public Builder() {
            $ = new GetProjectMergeRequestPlainArgs();
        }

        public Builder(GetProjectMergeRequestPlainArgs defaults) {
            $ = new GetProjectMergeRequestPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param iid The unique project level ID of the merge request.
         * 
         * @return builder
         * 
         */
        public Builder iid(Integer iid) {
            $.iid = iid;
            return this;
        }

        /**
         * @param project The ID or path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public GetProjectMergeRequestPlainArgs build() {
            if ($.iid == null) {
                throw new MissingRequiredPropertyException("GetProjectMergeRequestPlainArgs", "iid");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectMergeRequestPlainArgs", "project");
            }
            return $;
        }
    }

}
