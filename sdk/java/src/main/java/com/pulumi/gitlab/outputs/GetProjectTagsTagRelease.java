// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectTagsTagRelease {
    /**
     * @return The description of release.
     * 
     */
    private String description;
    /**
     * @return The name of the tag.
     * 
     */
    private String tagName;

    private GetProjectTagsTagRelease() {}
    /**
     * @return The description of release.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The name of the tag.
     * 
     */
    public String tagName() {
        return this.tagName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectTagsTagRelease defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String tagName;
        public Builder() {}
        public Builder(GetProjectTagsTagRelease defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.tagName = defaults.tagName;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTagRelease", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder tagName(String tagName) {
            if (tagName == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTagRelease", "tagName");
            }
            this.tagName = tagName;
            return this;
        }
        public GetProjectTagsTagRelease build() {
            final var _resultValue = new GetProjectTagsTagRelease();
            _resultValue.description = description;
            _resultValue.tagName = tagName;
            return _resultValue;
        }
    }
}
