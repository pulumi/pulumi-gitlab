// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceGithubState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceGithubState Empty = new ServiceGithubState();

    /**
     * Whether the integration is active.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return Whether the integration is active.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * Create time.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Create time.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    @Import(name="repositoryUrl")
    private @Nullable Output<String> repositoryUrl;

    public Optional<Output<String>> repositoryUrl() {
        return Optional.ofNullable(this.repositoryUrl);
    }

    /**
     * Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     * 
     */
    @Import(name="staticContext")
    private @Nullable Output<Boolean> staticContext;

    /**
     * @return Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
     * 
     */
    public Optional<Output<Boolean>> staticContext() {
        return Optional.ofNullable(this.staticContext);
    }

    /**
     * Title.
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return Title.
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    /**
     * A GitHub personal access token with at least `repo:status` scope.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return A GitHub personal access token with at least `repo:status` scope.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Update time.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Update time.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private ServiceGithubState() {}

    private ServiceGithubState(ServiceGithubState $) {
        this.active = $.active;
        this.createdAt = $.createdAt;
        this.project = $.project;
        this.repositoryUrl = $.repositoryUrl;
        this.staticContext = $.staticContext;
        this.title = $.title;
        this.token = $.token;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceGithubState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceGithubState $;

        public Builder() {
            $ = new ServiceGithubState();
        }

        public Builder(ServiceGithubState defaults) {
            $ = new ServiceGithubState(Objects.requireNonNull(defaults));
        }

        /**
         * @param active Whether the integration is active.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active Whether the integration is active.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param createdAt Create time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Create time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param project ID of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project ID of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public Builder repositoryUrl(@Nullable Output<String> repositoryUrl) {
            $.repositoryUrl = repositoryUrl;
            return this;
        }

        public Builder repositoryUrl(String repositoryUrl) {
            return repositoryUrl(Output.of(repositoryUrl));
        }

        /**
         * @param staticContext Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
         * 
         * @return builder
         * 
         */
        public Builder staticContext(@Nullable Output<Boolean> staticContext) {
            $.staticContext = staticContext;
            return this;
        }

        /**
         * @param staticContext Append instance name instead of branch to the status. Must enable to set a GitLab status check as *required* in GitHub. See [Static / dynamic status check names] to learn more.
         * 
         * @return builder
         * 
         */
        public Builder staticContext(Boolean staticContext) {
            return staticContext(Output.of(staticContext));
        }

        /**
         * @param title Title.
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Title.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        /**
         * @param token A GitHub personal access token with at least `repo:status` scope.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token A GitHub personal access token with at least `repo:status` scope.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param updatedAt Update time.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Update time.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public ServiceGithubState build() {
            return $;
        }
    }

}
