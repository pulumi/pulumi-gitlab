// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationGithubArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationGithubArgs Empty = new IntegrationGithubArgs();

    /**
     * ID of the project you want to activate integration on.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return ID of the project you want to activate integration on.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     * 
     */
    @Import(name="repositoryUrl", required=true)
    private Output<String> repositoryUrl;

    /**
     * @return The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
     * 
     */
    public Output<String> repositoryUrl() {
        return this.repositoryUrl;
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
     * A GitHub personal access token with at least `repo:status` scope.
     * 
     */
    @Import(name="token", required=true)
    private Output<String> token;

    /**
     * @return A GitHub personal access token with at least `repo:status` scope.
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    private IntegrationGithubArgs() {}

    private IntegrationGithubArgs(IntegrationGithubArgs $) {
        this.project = $.project;
        this.repositoryUrl = $.repositoryUrl;
        this.staticContext = $.staticContext;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationGithubArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationGithubArgs $;

        public Builder() {
            $ = new IntegrationGithubArgs();
        }

        public Builder(IntegrationGithubArgs defaults) {
            $ = new IntegrationGithubArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project ID of the project you want to activate integration on.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
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

        /**
         * @param repositoryUrl The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
         * 
         * @return builder
         * 
         */
        public Builder repositoryUrl(Output<String> repositoryUrl) {
            $.repositoryUrl = repositoryUrl;
            return this;
        }

        /**
         * @param repositoryUrl The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
         * 
         * @return builder
         * 
         */
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
         * @param token A GitHub personal access token with at least `repo:status` scope.
         * 
         * @return builder
         * 
         */
        public Builder token(Output<String> token) {
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

        public IntegrationGithubArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("IntegrationGithubArgs", "project");
            }
            if ($.repositoryUrl == null) {
                throw new MissingRequiredPropertyException("IntegrationGithubArgs", "repositoryUrl");
            }
            if ($.token == null) {
                throw new MissingRequiredPropertyException("IntegrationGithubArgs", "token");
            }
            return $;
        }
    }

}
