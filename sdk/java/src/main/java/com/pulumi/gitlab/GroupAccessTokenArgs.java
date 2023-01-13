// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupAccessTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupAccessTokenArgs Empty = new GroupAccessTokenArgs();

    /**
     * The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The ID or path of the group to add the group access token to.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or path of the group to add the group access token to.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * The name of the group access token.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the group access token.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`.
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<String>> scopes;

    /**
     * @return The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`.
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }

    private GroupAccessTokenArgs() {}

    private GroupAccessTokenArgs(GroupAccessTokenArgs $) {
        this.accessLevel = $.accessLevel;
        this.expiresAt = $.expiresAt;
        this.group = $.group;
        this.name = $.name;
        this.scopes = $.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupAccessTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupAccessTokenArgs $;

        public Builder() {
            $ = new GroupAccessTokenArgs();
        }

        public Builder(GroupAccessTokenArgs defaults) {
            $ = new GroupAccessTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param expiresAt The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param group The ID or path of the group to add the group access token to.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or path of the group to add the group access token to.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param name The name of the group access token.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the group access token.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param scopes The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        public GroupAccessTokenArgs build() {
            $.group = Objects.requireNonNull($.group, "expected parameter 'group' to be non-null");
            $.scopes = Objects.requireNonNull($.scopes, "expected parameter 'scopes' to be non-null");
            return $;
        }
    }

}