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


public final class ProjectAccessTokenState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectAccessTokenState Empty = new ProjectAccessTokenState();

    /**
     * The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

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
     * Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * A name to describe the project access token.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name to describe the project access token.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The id of the project to add the project access token to.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The id of the project to add the project access token to.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
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
     * Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
     * 
     */
    @Import(name="scopes")
    private @Nullable Output<List<String>> scopes;

    /**
     * @return Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
     * 
     */
    public Optional<Output<List<String>>> scopes() {
        return Optional.ofNullable(this.scopes);
    }

    /**
     * The secret token. **Note**: the token is not available for imported resources.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The secret token. **Note**: the token is not available for imported resources.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * The user_id associated to the token.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<Integer> userId;

    /**
     * @return The user_id associated to the token.
     * 
     */
    public Optional<Output<Integer>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private ProjectAccessTokenState() {}

    private ProjectAccessTokenState(ProjectAccessTokenState $) {
        this.accessLevel = $.accessLevel;
        this.active = $.active;
        this.createdAt = $.createdAt;
        this.expiresAt = $.expiresAt;
        this.name = $.name;
        this.project = $.project;
        this.revoked = $.revoked;
        this.scopes = $.scopes;
        this.token = $.token;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectAccessTokenState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectAccessTokenState $;

        public Builder() {
            $ = new ProjectAccessTokenState();
        }

        public Builder(ProjectAccessTokenState defaults) {
            $ = new ProjectAccessTokenState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`. Default is `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
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
         * @param expiresAt Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt Time the token will expire it, YYYY-MM-DD format. Will not expire per default.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param name A name to describe the project access token.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name to describe the project access token.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The id of the project to add the project access token to.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The id of the project to add the project access token to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
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
         * @param scopes Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(@Nullable Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        /**
         * @param token The secret token. **Note**: the token is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The secret token. **Note**: the token is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param userId The user_id associated to the token.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The user_id associated to the token.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public ProjectAccessTokenState build() {
            return $;
        }
    }

}
