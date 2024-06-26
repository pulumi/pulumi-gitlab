// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupMembershipArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupMembershipArgs Empty = new GetGroupMembershipArgs();

    /**
     * Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * The full path of the group.
     * 
     */
    @Import(name="fullPath")
    private @Nullable Output<String> fullPath;

    /**
     * @return The full path of the group.
     * 
     */
    public Optional<Output<String>> fullPath() {
        return Optional.ofNullable(this.fullPath);
    }

    /**
     * The ID of the group.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    /**
     * @return The ID of the group.
     * 
     */
    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * Return all project members including members through ancestor groups.
     * 
     */
    @Import(name="inherited")
    private @Nullable Output<Boolean> inherited;

    /**
     * @return Return all project members including members through ancestor groups.
     * 
     */
    public Optional<Output<Boolean>> inherited() {
        return Optional.ofNullable(this.inherited);
    }

    private GetGroupMembershipArgs() {}

    private GetGroupMembershipArgs(GetGroupMembershipArgs $) {
        this.accessLevel = $.accessLevel;
        this.fullPath = $.fullPath;
        this.groupId = $.groupId;
        this.inherited = $.inherited;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupMembershipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupMembershipArgs $;

        public Builder() {
            $ = new GetGroupMembershipArgs();
        }

        public Builder(GetGroupMembershipArgs defaults) {
            $ = new GetGroupMembershipArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param fullPath The full path of the group.
         * 
         * @return builder
         * 
         */
        public Builder fullPath(@Nullable Output<String> fullPath) {
            $.fullPath = fullPath;
            return this;
        }

        /**
         * @param fullPath The full path of the group.
         * 
         * @return builder
         * 
         */
        public Builder fullPath(String fullPath) {
            return fullPath(Output.of(fullPath));
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param inherited Return all project members including members through ancestor groups.
         * 
         * @return builder
         * 
         */
        public Builder inherited(@Nullable Output<Boolean> inherited) {
            $.inherited = inherited;
            return this;
        }

        /**
         * @param inherited Return all project members including members through ancestor groups.
         * 
         * @return builder
         * 
         */
        public Builder inherited(Boolean inherited) {
            return inherited(Output.of(inherited));
        }

        public GetGroupMembershipArgs build() {
            return $;
        }
    }

}
