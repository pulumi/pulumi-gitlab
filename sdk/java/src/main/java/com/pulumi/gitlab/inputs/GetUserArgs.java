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


public final class GetUserArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserArgs Empty = new GetUserArgs();

    /**
     * The public email address of the user.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The public email address of the user.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * The ID of the user&#39;s namespace. Requires admin token to access this field.
     * 
     */
    @Import(name="namespaceId")
    private @Nullable Output<Integer> namespaceId;

    /**
     * @return The ID of the user&#39;s namespace. Requires admin token to access this field.
     * 
     */
    public Optional<Output<Integer>> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }

    /**
     * The ID of the user.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The ID of the user.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    /**
     * The username of the user.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username of the user.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private GetUserArgs() {}

    private GetUserArgs(GetUserArgs $) {
        this.email = $.email;
        this.namespaceId = $.namespaceId;
        this.userId = $.userId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserArgs $;

        public Builder() {
            $ = new GetUserArgs();
        }

        public Builder(GetUserArgs defaults) {
            $ = new GetUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param email The public email address of the user.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The public email address of the user.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param namespaceId The ID of the user&#39;s namespace. Requires admin token to access this field.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(@Nullable Output<Integer> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId The ID of the user&#39;s namespace. Requires admin token to access this field.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(Integer namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        /**
         * @param username The username of the user.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the user.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GetUserArgs build() {
            return $;
        }
    }

}
