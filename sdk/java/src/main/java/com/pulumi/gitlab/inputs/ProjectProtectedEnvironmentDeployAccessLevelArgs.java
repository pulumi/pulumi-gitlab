// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectProtectedEnvironmentDeployAccessLevelArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectProtectedEnvironmentDeployAccessLevelArgs Empty = new ProjectProtectedEnvironmentDeployAccessLevelArgs();

    /**
     * Levels of access required to deploy to this protected environment. Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Levels of access required to deploy to this protected environment. Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * Readable description of level of access.
     * 
     */
    @Import(name="accessLevelDescription")
    private @Nullable Output<String> accessLevelDescription;

    /**
     * @return Readable description of level of access.
     * 
     */
    public Optional<Output<String>> accessLevelDescription() {
        return Optional.ofNullable(this.accessLevelDescription);
    }

    /**
     * The ID of the group allowed to deploy to this protected environment. The project must be shared with the group. Mutually exclusive with `access_level` and `user_id`.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return The ID of the group allowed to deploy to this protected environment. The project must be shared with the group. Mutually exclusive with `access_level` and `user_id`.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * Group inheritance allows deploy access levels to take inherited group membership into account. Valid values are `0`, `1`. `0` =&gt; Direct group membership only, `1` =&gt; All inherited groups. Default: `0`
     * 
     */
    @Import(name="groupInheritanceType")
    private @Nullable Output<Integer> groupInheritanceType;

    /**
     * @return Group inheritance allows deploy access levels to take inherited group membership into account. Valid values are `0`, `1`. `0` =&gt; Direct group membership only, `1` =&gt; All inherited groups. Default: `0`
     * 
     */
    public Optional<Output<Integer>> groupInheritanceType() {
        return Optional.ofNullable(this.groupInheritanceType);
    }

    /**
     * The unique ID of the Deploy Access Level object.
     * 
     */
    @Import(name="id")
    private @Nullable Output<Integer> id;

    /**
     * @return The unique ID of the Deploy Access Level object.
     * 
     */
    public Optional<Output<Integer>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The ID of the user allowed to deploy to this protected environment. The user must be a member of the project. Mutually exclusive with `access_level` and `group_id`.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The ID of the user allowed to deploy to this protected environment. The user must be a member of the project. Mutually exclusive with `access_level` and `group_id`.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private ProjectProtectedEnvironmentDeployAccessLevelArgs() {}

    private ProjectProtectedEnvironmentDeployAccessLevelArgs(ProjectProtectedEnvironmentDeployAccessLevelArgs $) {
        this.accessLevel = $.accessLevel;
        this.accessLevelDescription = $.accessLevelDescription;
        this.groupId = $.groupId;
        this.groupInheritanceType = $.groupInheritanceType;
        this.id = $.id;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectProtectedEnvironmentDeployAccessLevelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectProtectedEnvironmentDeployAccessLevelArgs $;

        public Builder() {
            $ = new ProjectProtectedEnvironmentDeployAccessLevelArgs();
        }

        public Builder(ProjectProtectedEnvironmentDeployAccessLevelArgs defaults) {
            $ = new ProjectProtectedEnvironmentDeployAccessLevelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Levels of access required to deploy to this protected environment. Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Levels of access required to deploy to this protected environment. Mutually exclusive with `user_id` and `group_id`. Valid values are `developer`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param accessLevelDescription Readable description of level of access.
         * 
         * @return builder
         * 
         */
        public Builder accessLevelDescription(@Nullable Output<String> accessLevelDescription) {
            $.accessLevelDescription = accessLevelDescription;
            return this;
        }

        /**
         * @param accessLevelDescription Readable description of level of access.
         * 
         * @return builder
         * 
         */
        public Builder accessLevelDescription(String accessLevelDescription) {
            return accessLevelDescription(Output.of(accessLevelDescription));
        }

        /**
         * @param groupId The ID of the group allowed to deploy to this protected environment. The project must be shared with the group. Mutually exclusive with `access_level` and `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the group allowed to deploy to this protected environment. The project must be shared with the group. Mutually exclusive with `access_level` and `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param groupInheritanceType Group inheritance allows deploy access levels to take inherited group membership into account. Valid values are `0`, `1`. `0` =&gt; Direct group membership only, `1` =&gt; All inherited groups. Default: `0`
         * 
         * @return builder
         * 
         */
        public Builder groupInheritanceType(@Nullable Output<Integer> groupInheritanceType) {
            $.groupInheritanceType = groupInheritanceType;
            return this;
        }

        /**
         * @param groupInheritanceType Group inheritance allows deploy access levels to take inherited group membership into account. Valid values are `0`, `1`. `0` =&gt; Direct group membership only, `1` =&gt; All inherited groups. Default: `0`
         * 
         * @return builder
         * 
         */
        public Builder groupInheritanceType(Integer groupInheritanceType) {
            return groupInheritanceType(Output.of(groupInheritanceType));
        }

        /**
         * @param id The unique ID of the Deploy Access Level object.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<Integer> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The unique ID of the Deploy Access Level object.
         * 
         * @return builder
         * 
         */
        public Builder id(Integer id) {
            return id(Output.of(id));
        }

        /**
         * @param userId The ID of the user allowed to deploy to this protected environment. The user must be a member of the project. Mutually exclusive with `access_level` and `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the user allowed to deploy to this protected environment. The user must be a member of the project. Mutually exclusive with `access_level` and `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public ProjectProtectedEnvironmentDeployAccessLevelArgs build() {
            return $;
        }
    }

}
