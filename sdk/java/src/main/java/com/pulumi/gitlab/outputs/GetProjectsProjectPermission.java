// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetProjectsProjectPermission {
    /**
     * @return Group access level.
     * 
     */
    private Map<String,Integer> groupAccess;
    /**
     * @return Project access level.
     * 
     */
    private Map<String,Integer> projectAccess;

    private GetProjectsProjectPermission() {}
    /**
     * @return Group access level.
     * 
     */
    public Map<String,Integer> groupAccess() {
        return this.groupAccess;
    }
    /**
     * @return Project access level.
     * 
     */
    public Map<String,Integer> projectAccess() {
        return this.projectAccess;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProjectPermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,Integer> groupAccess;
        private Map<String,Integer> projectAccess;
        public Builder() {}
        public Builder(GetProjectsProjectPermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupAccess = defaults.groupAccess;
    	      this.projectAccess = defaults.projectAccess;
        }

        @CustomType.Setter
        public Builder groupAccess(Map<String,Integer> groupAccess) {
            if (groupAccess == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectPermission", "groupAccess");
            }
            this.groupAccess = groupAccess;
            return this;
        }
        @CustomType.Setter
        public Builder projectAccess(Map<String,Integer> projectAccess) {
            if (projectAccess == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectPermission", "projectAccess");
            }
            this.projectAccess = projectAccess;
            return this;
        }
        public GetProjectsProjectPermission build() {
            final var _resultValue = new GetProjectsProjectPermission();
            _resultValue.groupAccess = groupAccess;
            _resultValue.projectAccess = projectAccess;
            return _resultValue;
        }
    }
}
