// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserPlainArgs Empty = new GetUserPlainArgs();

    /**
     * The public email address of the user.
     * 
     */
    @Import(name="email")
    private @Nullable String email;

    /**
     * @return The public email address of the user.
     * 
     */
    public Optional<String> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * The ID of the user&#39;s namespace. Requires admin token to access this field.
     * 
     */
    @Import(name="namespaceId")
    private @Nullable Integer namespaceId;

    /**
     * @return The ID of the user&#39;s namespace. Requires admin token to access this field.
     * 
     */
    public Optional<Integer> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }

    /**
     * The ID of the user.
     * 
     */
    @Import(name="userId")
    private @Nullable Integer userId;

    /**
     * @return The ID of the user.
     * 
     */
    public Optional<Integer> userId() {
        return Optional.ofNullable(this.userId);
    }

    /**
     * The username of the user.
     * 
     */
    @Import(name="username")
    private @Nullable String username;

    /**
     * @return The username of the user.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    private GetUserPlainArgs() {}

    private GetUserPlainArgs(GetUserPlainArgs $) {
        this.email = $.email;
        this.namespaceId = $.namespaceId;
        this.userId = $.userId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserPlainArgs $;

        public Builder() {
            $ = new GetUserPlainArgs();
        }

        public Builder(GetUserPlainArgs defaults) {
            $ = new GetUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param email The public email address of the user.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable String email) {
            $.email = email;
            return this;
        }

        /**
         * @param namespaceId The ID of the user&#39;s namespace. Requires admin token to access this field.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(@Nullable Integer namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Integer userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param username The username of the user.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable String username) {
            $.username = username;
            return this;
        }

        public GetUserPlainArgs build() {
            return $;
        }
    }

}
