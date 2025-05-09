// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectMembershipArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectMembershipArgs Empty = new ProjectMembershipArgs();

    /**
     * The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    @Import(name="accessLevel", required=true)
    private Output<String> accessLevel;

    /**
     * @return The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }

    /**
     * Expiration date for the project membership. Format: `YYYY-MM-DD`
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return Expiration date for the project membership. Format: `YYYY-MM-DD`
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The ID of a custom member role. Only available for Ultimate instances.
     * 
     */
    @Import(name="memberRoleId")
    private @Nullable Output<Integer> memberRoleId;

    /**
     * @return The ID of a custom member role. Only available for Ultimate instances.
     * 
     */
    public Optional<Output<Integer>> memberRoleId() {
        return Optional.ofNullable(this.memberRoleId);
    }

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
     * The id of the user.
     * 
     */
    @Import(name="userId", required=true)
    private Output<Integer> userId;

    /**
     * @return The id of the user.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    private ProjectMembershipArgs() {}

    private ProjectMembershipArgs(ProjectMembershipArgs $) {
        this.accessLevel = $.accessLevel;
        this.expiresAt = $.expiresAt;
        this.memberRoleId = $.memberRoleId;
        this.project = $.project;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectMembershipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectMembershipArgs $;

        public Builder() {
            $ = new ProjectMembershipArgs();
        }

        public Builder(ProjectMembershipArgs defaults) {
            $ = new ProjectMembershipArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel The access level for the member. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param expiresAt Expiration date for the project membership. Format: `YYYY-MM-DD`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt Expiration date for the project membership. Format: `YYYY-MM-DD`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(@Nullable Output<Integer> memberRoleId) {
            $.memberRoleId = memberRoleId;
            return this;
        }

        /**
         * @param memberRoleId The ID of a custom member role. Only available for Ultimate instances.
         * 
         * @return builder
         * 
         */
        public Builder memberRoleId(Integer memberRoleId) {
            return memberRoleId(Output.of(memberRoleId));
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
         * @param userId The id of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The id of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public ProjectMembershipArgs build() {
            if ($.accessLevel == null) {
                throw new MissingRequiredPropertyException("ProjectMembershipArgs", "accessLevel");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectMembershipArgs", "project");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("ProjectMembershipArgs", "userId");
            }
            return $;
        }
    }

}
