// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gitlab.inputs.ProjectAccessTokenRotationConfigurationArgs;
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
     * The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
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
     * The description of the project access token.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the project access token.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The name of the project access token.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the project access token.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID or full path of the project.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID or full path of the project.
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
     * The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     * 
     */
    @Import(name="rotationConfiguration")
    private @Nullable Output<ProjectAccessTokenRotationConfigurationArgs> rotationConfiguration;

    /**
     * @return The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
     * 
     */
    public Optional<Output<ProjectAccessTokenRotationConfigurationArgs>> rotationConfiguration() {
        return Optional.ofNullable(this.rotationConfiguration);
    }

    /**
     * The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    @Import(name="scopes")
    private @Nullable Output<List<String>> scopes;

    /**
     * @return The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    public Optional<Output<List<String>>> scopes() {
        return Optional.ofNullable(this.scopes);
    }

    /**
     * The token of the project access token. **Note**: the token is not available for imported resources.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The token of the project access token. **Note**: the token is not available for imported resources.
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
        this.description = $.description;
        this.expiresAt = $.expiresAt;
        this.name = $.name;
        this.project = $.project;
        this.revoked = $.revoked;
        this.rotationConfiguration = $.rotationConfiguration;
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
         * @param accessLevel The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `planner`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
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
         * @param description The description of the project access token.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the project access token.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param expiresAt When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt When the token will expire, YYYY-MM-DD format. Is automatically set when `rotation_configuration` is used.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param name The name of the project access token.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the project access token.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID or full path of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or full path of the project.
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
         * @param rotationConfiguration The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
         * 
         * @return builder
         * 
         */
        public Builder rotationConfiguration(@Nullable Output<ProjectAccessTokenRotationConfigurationArgs> rotationConfiguration) {
            $.rotationConfiguration = rotationConfiguration;
            return this;
        }

        /**
         * @param rotationConfiguration The configuration for when to rotate a token automatically. Will not rotate a token until `pulumi up` is run.
         * 
         * @return builder
         * 
         */
        public Builder rotationConfiguration(ProjectAccessTokenRotationConfigurationArgs rotationConfiguration) {
            return rotationConfiguration(Output.of(rotationConfiguration));
        }

        /**
         * @param scopes The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(@Nullable Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes The scopes of the project access token. valid values are: `api`, `read_api`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `manage_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        /**
         * @param token The token of the project access token. **Note**: the token is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The token of the project access token. **Note**: the token is not available for imported resources.
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
