// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.ProjectAccessTokenRotationConfigurationArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectAccessTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectAccessTokenArgs Empty = new ProjectAccessTokenArgs();

    /**
     * The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
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
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or full path of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
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
     * The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<String>> scopes;

    /**
     * @return The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }

    private ProjectAccessTokenArgs() {}

    private ProjectAccessTokenArgs(ProjectAccessTokenArgs $) {
        this.accessLevel = $.accessLevel;
        this.expiresAt = $.expiresAt;
        this.name = $.name;
        this.project = $.project;
        this.rotationConfiguration = $.rotationConfiguration;
        this.scopes = $.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectAccessTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectAccessTokenArgs $;

        public Builder() {
            $ = new ProjectAccessTokenArgs();
        }

        public Builder(ProjectAccessTokenArgs defaults) {
            $ = new ProjectAccessTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel The access level for the project access token. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`. Default is `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
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
        public Builder project(Output<String> project) {
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
         * @param scopes The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes The scopes of the project access token. valid values are: `api`, `read_api`, `read_user`, `k8s_proxy`, `read_registry`, `write_registry`, `read_repository`, `write_repository`, `create_runner`, `ai_features`, `k8s_proxy`, `read_observability`, `write_observability`
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        public ProjectAccessTokenArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectAccessTokenArgs", "project");
            }
            if ($.scopes == null) {
                throw new MissingRequiredPropertyException("ProjectAccessTokenArgs", "scopes");
            }
            return $;
        }
    }

}
