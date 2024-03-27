// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetReleaseAssetsLink {
    /**
     * @return The ID of the link
     * 
     */
    private Integer id;
    /**
     * @return The type of the link
     * 
     */
    private String linkType;
    /**
     * @return The name of the link
     * 
     */
    private String name;
    /**
     * @return The URL of the link
     * 
     */
    private String url;

    private GetReleaseAssetsLink() {}
    /**
     * @return The ID of the link
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The type of the link
     * 
     */
    public String linkType() {
        return this.linkType;
    }
    /**
     * @return The name of the link
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The URL of the link
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetReleaseAssetsLink defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer id;
        private String linkType;
        private String name;
        private String url;
        public Builder() {}
        public Builder(GetReleaseAssetsLink defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.linkType = defaults.linkType;
    	      this.name = defaults.name;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetReleaseAssetsLink", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder linkType(String linkType) {
            if (linkType == null) {
              throw new MissingRequiredPropertyException("GetReleaseAssetsLink", "linkType");
            }
            this.linkType = linkType;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetReleaseAssetsLink", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetReleaseAssetsLink", "url");
            }
            this.url = url;
            return this;
        }
        public GetReleaseAssetsLink build() {
            final var _resultValue = new GetReleaseAssetsLink();
            _resultValue.id = id;
            _resultValue.linkType = linkType;
            _resultValue.name = name;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}