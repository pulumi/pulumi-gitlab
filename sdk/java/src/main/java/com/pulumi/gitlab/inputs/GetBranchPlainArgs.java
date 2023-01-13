// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetBranchPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBranchPlainArgs Empty = new GetBranchPlainArgs();

    /**
     * The name of the branch.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the branch.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The full path or id of the project.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The full path or id of the project.
     * 
     */
    public String project() {
        return this.project;
    }

    private GetBranchPlainArgs() {}

    private GetBranchPlainArgs(GetBranchPlainArgs $) {
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBranchPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBranchPlainArgs $;

        public Builder() {
            $ = new GetBranchPlainArgs();
        }

        public Builder(GetBranchPlainArgs defaults) {
            $ = new GetBranchPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the branch.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param project The full path or id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public GetBranchPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            return $;
        }
    }

}