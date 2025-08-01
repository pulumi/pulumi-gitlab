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


public final class ProjectPagesSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectPagesSettingsArgs Empty = new ProjectPagesSettingsArgs();

    /**
     * Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
     * 
     */
    @Import(name="forceHttps")
    private @Nullable Output<Boolean> forceHttps;

    /**
     * @return Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
     * 
     */
    public Optional<Output<Boolean>> forceHttps() {
        return Optional.ofNullable(this.forceHttps);
    }

    /**
     * Boolean indicating if a unique domain is enabled.
     * 
     */
    @Import(name="isUniqueDomainEnabled")
    private @Nullable Output<Boolean> isUniqueDomainEnabled;

    /**
     * @return Boolean indicating if a unique domain is enabled.
     * 
     */
    public Optional<Output<Boolean>> isUniqueDomainEnabled() {
        return Optional.ofNullable(this.isUniqueDomainEnabled);
    }

    @Import(name="keepSettingsOnDestroy")
    private @Nullable Output<Boolean> keepSettingsOnDestroy;

    public Optional<Output<Boolean>> keepSettingsOnDestroy() {
        return Optional.ofNullable(this.keepSettingsOnDestroy);
    }

    /**
     * The project ID or path.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The project ID or path.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private ProjectPagesSettingsArgs() {}

    private ProjectPagesSettingsArgs(ProjectPagesSettingsArgs $) {
        this.forceHttps = $.forceHttps;
        this.isUniqueDomainEnabled = $.isUniqueDomainEnabled;
        this.keepSettingsOnDestroy = $.keepSettingsOnDestroy;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectPagesSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectPagesSettingsArgs $;

        public Builder() {
            $ = new ProjectPagesSettingsArgs();
        }

        public Builder(ProjectPagesSettingsArgs defaults) {
            $ = new ProjectPagesSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param forceHttps Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
         * 
         * @return builder
         * 
         */
        public Builder forceHttps(@Nullable Output<Boolean> forceHttps) {
            $.forceHttps = forceHttps;
            return this;
        }

        /**
         * @param forceHttps Boolean indicating if the project is set to force https. Requires `external_https` to be configured in the GitLab instance: https://docs.gitlab.com/administration/pages/#custom-domains-with-tls-support.
         * 
         * @return builder
         * 
         */
        public Builder forceHttps(Boolean forceHttps) {
            return forceHttps(Output.of(forceHttps));
        }

        /**
         * @param isUniqueDomainEnabled Boolean indicating if a unique domain is enabled.
         * 
         * @return builder
         * 
         */
        public Builder isUniqueDomainEnabled(@Nullable Output<Boolean> isUniqueDomainEnabled) {
            $.isUniqueDomainEnabled = isUniqueDomainEnabled;
            return this;
        }

        /**
         * @param isUniqueDomainEnabled Boolean indicating if a unique domain is enabled.
         * 
         * @return builder
         * 
         */
        public Builder isUniqueDomainEnabled(Boolean isUniqueDomainEnabled) {
            return isUniqueDomainEnabled(Output.of(isUniqueDomainEnabled));
        }

        public Builder keepSettingsOnDestroy(@Nullable Output<Boolean> keepSettingsOnDestroy) {
            $.keepSettingsOnDestroy = keepSettingsOnDestroy;
            return this;
        }

        public Builder keepSettingsOnDestroy(Boolean keepSettingsOnDestroy) {
            return keepSettingsOnDestroy(Output.of(keepSettingsOnDestroy));
        }

        /**
         * @param project The project ID or path.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project ID or path.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public ProjectPagesSettingsArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectPagesSettingsArgs", "project");
            }
            return $;
        }
    }

}
