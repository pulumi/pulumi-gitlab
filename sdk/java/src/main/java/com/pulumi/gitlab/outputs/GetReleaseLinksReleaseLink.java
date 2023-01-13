// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetReleaseLinksReleaseLink {
    private String directAssetUrl;
    private Boolean external;
    private String filepath;
    private Integer linkId;
    private String linkType;
    private String name;
    /**
     * @return The ID or full path to the project.
     * 
     */
    private String project;
    /**
     * @return The tag associated with the Release.
     * 
     */
    private String tagName;
    private String url;

    private GetReleaseLinksReleaseLink() {}
    public String directAssetUrl() {
        return this.directAssetUrl;
    }
    public Boolean external() {
        return this.external;
    }
    public String filepath() {
        return this.filepath;
    }
    public Integer linkId() {
        return this.linkId;
    }
    public String linkType() {
        return this.linkType;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The ID or full path to the project.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return The tag associated with the Release.
     * 
     */
    public String tagName() {
        return this.tagName;
    }
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetReleaseLinksReleaseLink defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String directAssetUrl;
        private Boolean external;
        private String filepath;
        private Integer linkId;
        private String linkType;
        private String name;
        private String project;
        private String tagName;
        private String url;
        public Builder() {}
        public Builder(GetReleaseLinksReleaseLink defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.directAssetUrl = defaults.directAssetUrl;
    	      this.external = defaults.external;
    	      this.filepath = defaults.filepath;
    	      this.linkId = defaults.linkId;
    	      this.linkType = defaults.linkType;
    	      this.name = defaults.name;
    	      this.project = defaults.project;
    	      this.tagName = defaults.tagName;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder directAssetUrl(String directAssetUrl) {
            this.directAssetUrl = Objects.requireNonNull(directAssetUrl);
            return this;
        }
        @CustomType.Setter
        public Builder external(Boolean external) {
            this.external = Objects.requireNonNull(external);
            return this;
        }
        @CustomType.Setter
        public Builder filepath(String filepath) {
            this.filepath = Objects.requireNonNull(filepath);
            return this;
        }
        @CustomType.Setter
        public Builder linkId(Integer linkId) {
            this.linkId = Objects.requireNonNull(linkId);
            return this;
        }
        @CustomType.Setter
        public Builder linkType(String linkType) {
            this.linkType = Objects.requireNonNull(linkType);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        @CustomType.Setter
        public Builder tagName(String tagName) {
            this.tagName = Objects.requireNonNull(tagName);
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public GetReleaseLinksReleaseLink build() {
            final var o = new GetReleaseLinksReleaseLink();
            o.directAssetUrl = directAssetUrl;
            o.external = external;
            o.filepath = filepath;
            o.linkId = linkId;
            o.linkType = linkType;
            o.name = name;
            o.project = project;
            o.tagName = tagName;
            o.url = url;
            return o;
        }
    }
}