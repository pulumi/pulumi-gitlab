// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseLinkArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseLinkArgs Empty = new ReleaseLinkArgs();

    /**
     * Relative path for a [Direct Asset link](https://docs.gitlab.com/user/project/releases/index/#permanent-links-to-release-assets).
     * 
     */
    @Import(name="filepath")
    private @Nullable Output<String> filepath;

    /**
     * @return Relative path for a [Direct Asset link](https://docs.gitlab.com/user/project/releases/index/#permanent-links-to-release-assets).
     * 
     */
    public Optional<Output<String>> filepath() {
        return Optional.ofNullable(this.filepath);
    }

    /**
     * The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
     * 
     */
    @Import(name="linkType")
    private @Nullable Output<String> linkType;

    /**
     * @return The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
     * 
     */
    public Optional<Output<String>> linkType() {
        return Optional.ofNullable(this.linkType);
    }

    /**
     * The name of the link. Link names must be unique within the release.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the link. Link names must be unique within the release.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * The tag associated with the Release.
     * 
     */
    @Import(name="tagName", required=true)
    private Output<String> tagName;

    /**
     * @return The tag associated with the Release.
     * 
     */
    public Output<String> tagName() {
        return this.tagName;
    }

    /**
     * The URL of the link. Link URLs must be unique within the release.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The URL of the link. Link URLs must be unique within the release.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private ReleaseLinkArgs() {}

    private ReleaseLinkArgs(ReleaseLinkArgs $) {
        this.filepath = $.filepath;
        this.linkType = $.linkType;
        this.name = $.name;
        this.project = $.project;
        this.tagName = $.tagName;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseLinkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseLinkArgs $;

        public Builder() {
            $ = new ReleaseLinkArgs();
        }

        public Builder(ReleaseLinkArgs defaults) {
            $ = new ReleaseLinkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filepath Relative path for a [Direct Asset link](https://docs.gitlab.com/user/project/releases/index/#permanent-links-to-release-assets).
         * 
         * @return builder
         * 
         */
        public Builder filepath(@Nullable Output<String> filepath) {
            $.filepath = filepath;
            return this;
        }

        /**
         * @param filepath Relative path for a [Direct Asset link](https://docs.gitlab.com/user/project/releases/index/#permanent-links-to-release-assets).
         * 
         * @return builder
         * 
         */
        public Builder filepath(String filepath) {
            return filepath(Output.of(filepath));
        }

        /**
         * @param linkType The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
         * 
         * @return builder
         * 
         */
        public Builder linkType(@Nullable Output<String> linkType) {
            $.linkType = linkType;
            return this;
        }

        /**
         * @param linkType The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
         * 
         * @return builder
         * 
         */
        public Builder linkType(String linkType) {
            return linkType(Output.of(linkType));
        }

        /**
         * @param name The name of the link. Link names must be unique within the release.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the link. Link names must be unique within the release.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param tagName The tag associated with the Release.
         * 
         * @return builder
         * 
         */
        public Builder tagName(Output<String> tagName) {
            $.tagName = tagName;
            return this;
        }

        /**
         * @param tagName The tag associated with the Release.
         * 
         * @return builder
         * 
         */
        public Builder tagName(String tagName) {
            return tagName(Output.of(tagName));
        }

        /**
         * @param url The URL of the link. Link URLs must be unique within the release.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL of the link. Link URLs must be unique within the release.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ReleaseLinkArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ReleaseLinkArgs", "project");
            }
            if ($.tagName == null) {
                throw new MissingRequiredPropertyException("ReleaseLinkArgs", "tagName");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ReleaseLinkArgs", "url");
            }
            return $;
        }
    }

}
