// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectProtectedTagCreateAccessLevel;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectProtectedTagResult {
    /**
     * @return Array of access levels/user(s)/group(s) allowed to create protected tags.
     * 
     */
    private List<GetProjectProtectedTagCreateAccessLevel> createAccessLevels;
    /**
     * @return The ID of this resource. In the format of `&lt;tag&gt;`.
     * 
     */
    private String id;
    /**
     * @return The integer or path with namespace that uniquely identifies the project.
     * 
     */
    private String project;
    /**
     * @return The name of the protected tag.
     * 
     */
    private String tag;

    private GetProjectProtectedTagResult() {}
    /**
     * @return Array of access levels/user(s)/group(s) allowed to create protected tags.
     * 
     */
    public List<GetProjectProtectedTagCreateAccessLevel> createAccessLevels() {
        return this.createAccessLevels;
    }
    /**
     * @return The ID of this resource. In the format of `&lt;tag&gt;`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The integer or path with namespace that uniquely identifies the project.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return The name of the protected tag.
     * 
     */
    public String tag() {
        return this.tag;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectProtectedTagResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetProjectProtectedTagCreateAccessLevel> createAccessLevels;
        private String id;
        private String project;
        private String tag;
        public Builder() {}
        public Builder(GetProjectProtectedTagResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createAccessLevels = defaults.createAccessLevels;
    	      this.id = defaults.id;
    	      this.project = defaults.project;
    	      this.tag = defaults.tag;
        }

        @CustomType.Setter
        public Builder createAccessLevels(List<GetProjectProtectedTagCreateAccessLevel> createAccessLevels) {
            if (createAccessLevels == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedTagResult", "createAccessLevels");
            }
            this.createAccessLevels = createAccessLevels;
            return this;
        }
        public Builder createAccessLevels(GetProjectProtectedTagCreateAccessLevel... createAccessLevels) {
            return createAccessLevels(List.of(createAccessLevels));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedTagResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedTagResult", "project");
            }
            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder tag(String tag) {
            if (tag == null) {
              throw new MissingRequiredPropertyException("GetProjectProtectedTagResult", "tag");
            }
            this.tag = tag;
            return this;
        }
        public GetProjectProtectedTagResult build() {
            final var _resultValue = new GetProjectProtectedTagResult();
            _resultValue.createAccessLevels = createAccessLevels;
            _resultValue.id = id;
            _resultValue.project = project;
            _resultValue.tag = tag;
            return _resultValue;
        }
    }
}
