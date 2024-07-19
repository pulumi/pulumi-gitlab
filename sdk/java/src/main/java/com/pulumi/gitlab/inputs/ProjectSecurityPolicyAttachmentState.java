// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectSecurityPolicyAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectSecurityPolicyAttachmentState Empty = new ProjectSecurityPolicyAttachmentState();

    /**
     * The ID or Full Path of the security policy project.
     * 
     */
    @Import(name="policyProject")
    private @Nullable Output<String> policyProject;

    /**
     * @return The ID or Full Path of the security policy project.
     * 
     */
    public Optional<Output<String>> policyProject() {
        return Optional.ofNullable(this.policyProject);
    }

    /**
     * The GraphQL ID of the security policy project.
     * 
     */
    @Import(name="policyProjectGraphqlId")
    private @Nullable Output<String> policyProjectGraphqlId;

    /**
     * @return The GraphQL ID of the security policy project.
     * 
     */
    public Optional<Output<String>> policyProjectGraphqlId() {
        return Optional.ofNullable(this.policyProjectGraphqlId);
    }

    /**
     * The ID or Full Path of the project which will have the security policy project assigned to it.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID or Full Path of the project which will have the security policy project assigned to it.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The GraphQL ID of the project to which the security policty project will be attached.
     * 
     */
    @Import(name="projectGraphqlId")
    private @Nullable Output<String> projectGraphqlId;

    /**
     * @return The GraphQL ID of the project to which the security policty project will be attached.
     * 
     */
    public Optional<Output<String>> projectGraphqlId() {
        return Optional.ofNullable(this.projectGraphqlId);
    }

    private ProjectSecurityPolicyAttachmentState() {}

    private ProjectSecurityPolicyAttachmentState(ProjectSecurityPolicyAttachmentState $) {
        this.policyProject = $.policyProject;
        this.policyProjectGraphqlId = $.policyProjectGraphqlId;
        this.project = $.project;
        this.projectGraphqlId = $.projectGraphqlId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectSecurityPolicyAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectSecurityPolicyAttachmentState $;

        public Builder() {
            $ = new ProjectSecurityPolicyAttachmentState();
        }

        public Builder(ProjectSecurityPolicyAttachmentState defaults) {
            $ = new ProjectSecurityPolicyAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param policyProject The ID or Full Path of the security policy project.
         * 
         * @return builder
         * 
         */
        public Builder policyProject(@Nullable Output<String> policyProject) {
            $.policyProject = policyProject;
            return this;
        }

        /**
         * @param policyProject The ID or Full Path of the security policy project.
         * 
         * @return builder
         * 
         */
        public Builder policyProject(String policyProject) {
            return policyProject(Output.of(policyProject));
        }

        /**
         * @param policyProjectGraphqlId The GraphQL ID of the security policy project.
         * 
         * @return builder
         * 
         */
        public Builder policyProjectGraphqlId(@Nullable Output<String> policyProjectGraphqlId) {
            $.policyProjectGraphqlId = policyProjectGraphqlId;
            return this;
        }

        /**
         * @param policyProjectGraphqlId The GraphQL ID of the security policy project.
         * 
         * @return builder
         * 
         */
        public Builder policyProjectGraphqlId(String policyProjectGraphqlId) {
            return policyProjectGraphqlId(Output.of(policyProjectGraphqlId));
        }

        /**
         * @param project The ID or Full Path of the project which will have the security policy project assigned to it.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or Full Path of the project which will have the security policy project assigned to it.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param projectGraphqlId The GraphQL ID of the project to which the security policty project will be attached.
         * 
         * @return builder
         * 
         */
        public Builder projectGraphqlId(@Nullable Output<String> projectGraphqlId) {
            $.projectGraphqlId = projectGraphqlId;
            return this;
        }

        /**
         * @param projectGraphqlId The GraphQL ID of the project to which the security policty project will be attached.
         * 
         * @return builder
         * 
         */
        public Builder projectGraphqlId(String projectGraphqlId) {
            return projectGraphqlId(Output.of(projectGraphqlId));
        }

        public ProjectSecurityPolicyAttachmentState build() {
            return $;
        }
    }

}
