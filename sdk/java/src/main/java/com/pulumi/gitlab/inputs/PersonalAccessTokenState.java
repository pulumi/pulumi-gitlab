// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PersonalAccessTokenState extends com.pulumi.resources.ResourceArgs {

    public static final PersonalAccessTokenState Empty = new PersonalAccessTokenState();

    /**
     * True if the token is active.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return True if the token is active.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * Time the token has been created, RFC3339 format.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Time the token has been created, RFC3339 format.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
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
     * True if the token is revoked.
     * 
     */
    @Import(name="revoked")
    private @Nullable Output<Boolean> revoked;

    /**
     * @return True if the token is revoked.
     * 
     */
    public Optional<Output<Boolean>> revoked() {
        return Optional.ofNullable(this.revoked);
    }

    /**
     * The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
     * 
     */
    @Import(name="scopes")
    private @Nullable Output<List<String>> scopes;

    /**
     * @return The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
     * 
     */
    public Optional<Output<List<String>>> scopes() {
        return Optional.ofNullable(this.scopes);
    }

    /**
     * The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * The id of the user.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The id of the user.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private PersonalAccessTokenState() {}

    private PersonalAccessTokenState(PersonalAccessTokenState $) {
        this.active = $.active;
        this.createdAt = $.createdAt;
        this.expiresAt = $.expiresAt;
        this.name = $.name;
        this.revoked = $.revoked;
        this.scopes = $.scopes;
        this.token = $.token;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PersonalAccessTokenState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PersonalAccessTokenState $;

        public Builder() {
            $ = new PersonalAccessTokenState();
        }

        public Builder(PersonalAccessTokenState defaults) {
            $ = new PersonalAccessTokenState(Objects.requireNonNull(defaults));
        }

        /**
         * @param active True if the token is active.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active True if the token is active.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param createdAt Time the token has been created, RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Time the token has been created, RFC3339 format.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
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
         * @param revoked True if the token is revoked.
         * 
         * @return builder
         * 
         */
        public Builder revoked(@Nullable Output<Boolean> revoked) {
            $.revoked = revoked;
            return this;
        }

        /**
         * @param revoked True if the token is revoked.
         * 
         * @return builder
         * 
         */
        public Builder revoked(Boolean revoked) {
            return revoked(Output.of(revoked));
        }

        /**
         * @param scopes The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(@Nullable Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes The scope for the personal access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `read_user`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`, `sudo`, `admin_mode`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        /**
         * @param token The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The personal access token. This is only populated when creating a new personal access token. This attribute is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param userId The id of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
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

        public PersonalAccessTokenState build() {
            return $;
        }
    }

}
