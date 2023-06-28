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


public final class TagProtectionAllowedToCreateArgs extends com.pulumi.resources.ResourceArgs {

    public static final TagProtectionAllowedToCreateArgs Empty = new TagProtectionAllowedToCreateArgs();

    /**
     * Level of access.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Level of access.
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
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private TagProtectionAllowedToCreateArgs() {}

    private TagProtectionAllowedToCreateArgs(TagProtectionAllowedToCreateArgs $) {
        this.accessLevel = $.accessLevel;
        this.accessLevelDescription = $.accessLevelDescription;
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TagProtectionAllowedToCreateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TagProtectionAllowedToCreateArgs $;

        public Builder() {
            $ = new TagProtectionAllowedToCreateArgs();
        }

        public Builder(TagProtectionAllowedToCreateArgs defaults) {
            $ = new TagProtectionAllowedToCreateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Level of access.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Level of access.
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
         * @param groupId The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param userId The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public TagProtectionAllowedToCreateArgs build() {
            return $;
        }
    }

}
