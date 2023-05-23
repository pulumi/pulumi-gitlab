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


public final class ProjectAccessTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectAccessTokenArgs Empty = new ProjectAccessTokenArgs();

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
     * Time the token will expire it, YYYY-MM-DD format.
     * 
     */
    @Import(name="expiresAt", required=true)
    private Output<String> expiresAt;

    /**
     * @return Time the token will expire it, YYYY-MM-DD format.
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
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
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The id of the project to add the project access token to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<String>> scopes;

    /**
     * @return Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
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
         * @param expiresAt Time the token will expire it, YYYY-MM-DD format.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt Time the token will expire it, YYYY-MM-DD format.
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
        public Builder project(Output<String> project) {
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
         * @param scopes Valid values: `api`, `read_api`, `read_repository`, `write_repository`, `read_registry`, `write_registry`.
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<String>> scopes) {
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

        public ProjectAccessTokenArgs build() {
            $.expiresAt = Objects.requireNonNull($.expiresAt, "expected parameter 'expiresAt' to be non-null");
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            $.scopes = Objects.requireNonNull($.scopes, "expected parameter 'scopes' to be non-null");
            return $;
        }
    }

}
