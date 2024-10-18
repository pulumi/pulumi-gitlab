// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupServiceAccountAccessTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupServiceAccountAccessTokenArgs Empty = new GroupServiceAccountAccessTokenArgs();

    /**
     * The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group containing the service account. Must be a top level group.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * The name of the personal access token.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the personal access token.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<String>> scopes;

    /**
     * @return The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }

    /**
     * The ID of a service account user.
     * 
     */
    @Import(name="userId", required=true)
    private Output<Integer> userId;

    /**
     * @return The ID of a service account user.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    private GroupServiceAccountAccessTokenArgs() {}

    private GroupServiceAccountAccessTokenArgs(GroupServiceAccountAccessTokenArgs $) {
        this.expiresAt = $.expiresAt;
        this.group = $.group;
        this.name = $.name;
        this.scopes = $.scopes;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupServiceAccountAccessTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupServiceAccountAccessTokenArgs $;

        public Builder() {
            $ = new GroupServiceAccountAccessTokenArgs();
        }

        public Builder(GroupServiceAccountAccessTokenArgs defaults) {
            $ = new GroupServiceAccountAccessTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param expiresAt The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt The personal access token expiry date. When left blank, the token follows the standard rule of expiry for personal access tokens.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param group The ID or URL-encoded path of the group containing the service account. Must be a top level group.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or URL-encoded path of the group containing the service account. Must be a top level group.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param name The name of the personal access token.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the personal access token.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param scopes The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes The scopes of the group service account access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        /**
         * @param userId The ID of a service account user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of a service account user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public GroupServiceAccountAccessTokenArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GroupServiceAccountAccessTokenArgs", "group");
            }
            if ($.scopes == null) {
                throw new MissingRequiredPropertyException("GroupServiceAccountAccessTokenArgs", "scopes");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("GroupServiceAccountAccessTokenArgs", "userId");
            }
            return $;
        }
    }

}
