// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ProjectTargetBranchRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectTargetBranchRuleArgs Empty = new ProjectTargetBranchRuleArgs();

    /**
     * The ID or URL-encoded path of the project.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or URL-encoded path of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * A pattern matching the branch name for which the merge request should have a default target branch configured.
     * 
     */
    @Import(name="sourceBranchPattern", required=true)
    private Output<String> sourceBranchPattern;

    /**
     * @return A pattern matching the branch name for which the merge request should have a default target branch configured.
     * 
     */
    public Output<String> sourceBranchPattern() {
        return this.sourceBranchPattern;
    }

    /**
     * The name of the branch to which the merge request should be addressed.
     * 
     */
    @Import(name="targetBranchName", required=true)
    private Output<String> targetBranchName;

    /**
     * @return The name of the branch to which the merge request should be addressed.
     * 
     */
    public Output<String> targetBranchName() {
        return this.targetBranchName;
    }

    private ProjectTargetBranchRuleArgs() {}

    private ProjectTargetBranchRuleArgs(ProjectTargetBranchRuleArgs $) {
        this.project = $.project;
        this.sourceBranchPattern = $.sourceBranchPattern;
        this.targetBranchName = $.targetBranchName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectTargetBranchRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectTargetBranchRuleArgs $;

        public Builder() {
            $ = new ProjectTargetBranchRuleArgs();
        }

        public Builder(ProjectTargetBranchRuleArgs defaults) {
            $ = new ProjectTargetBranchRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The ID or URL-encoded path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or URL-encoded path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param sourceBranchPattern A pattern matching the branch name for which the merge request should have a default target branch configured.
         * 
         * @return builder
         * 
         */
        public Builder sourceBranchPattern(Output<String> sourceBranchPattern) {
            $.sourceBranchPattern = sourceBranchPattern;
            return this;
        }

        /**
         * @param sourceBranchPattern A pattern matching the branch name for which the merge request should have a default target branch configured.
         * 
         * @return builder
         * 
         */
        public Builder sourceBranchPattern(String sourceBranchPattern) {
            return sourceBranchPattern(Output.of(sourceBranchPattern));
        }

        /**
         * @param targetBranchName The name of the branch to which the merge request should be addressed.
         * 
         * @return builder
         * 
         */
        public Builder targetBranchName(Output<String> targetBranchName) {
            $.targetBranchName = targetBranchName;
            return this;
        }

        /**
         * @param targetBranchName The name of the branch to which the merge request should be addressed.
         * 
         * @return builder
         * 
         */
        public Builder targetBranchName(String targetBranchName) {
            return targetBranchName(Output.of(targetBranchName));
        }

        public ProjectTargetBranchRuleArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectTargetBranchRuleArgs", "project");
            }
            if ($.sourceBranchPattern == null) {
                throw new MissingRequiredPropertyException("ProjectTargetBranchRuleArgs", "sourceBranchPattern");
            }
            if ($.targetBranchName == null) {
                throw new MissingRequiredPropertyException("ProjectTargetBranchRuleArgs", "targetBranchName");
            }
            return $;
        }
    }

}
