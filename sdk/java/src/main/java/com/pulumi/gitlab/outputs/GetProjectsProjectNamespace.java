// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProjectNamespace {
    /**
     * @return The full path of the namespace.
     * 
     */
    private String fullPath;
    /**
     * @return The ID of the namespace.
     * 
     */
    private Integer id;
    /**
     * @return The kind of the namespace.
     * 
     */
    private String kind;
    /**
     * @return The name of the namespace.
     * 
     */
    private String name;
    /**
     * @return The path of the namespace.
     * 
     */
    private String path;

    private GetProjectsProjectNamespace() {}
    /**
     * @return The full path of the namespace.
     * 
     */
    public String fullPath() {
        return this.fullPath;
    }
    /**
     * @return The ID of the namespace.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The kind of the namespace.
     * 
     */
    public String kind() {
        return this.kind;
    }
    /**
     * @return The name of the namespace.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The path of the namespace.
     * 
     */
    public String path() {
        return this.path;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProjectNamespace defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fullPath;
        private Integer id;
        private String kind;
        private String name;
        private String path;
        public Builder() {}
        public Builder(GetProjectsProjectNamespace defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullPath = defaults.fullPath;
    	      this.id = defaults.id;
    	      this.kind = defaults.kind;
    	      this.name = defaults.name;
    	      this.path = defaults.path;
        }

        @CustomType.Setter
        public Builder fullPath(String fullPath) {
            if (fullPath == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectNamespace", "fullPath");
            }
            this.fullPath = fullPath;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectNamespace", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder kind(String kind) {
            if (kind == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectNamespace", "kind");
            }
            this.kind = kind;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectNamespace", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            if (path == null) {
              throw new MissingRequiredPropertyException("GetProjectsProjectNamespace", "path");
            }
            this.path = path;
            return this;
        }
        public GetProjectsProjectNamespace build() {
            final var _resultValue = new GetProjectsProjectNamespace();
            _resultValue.fullPath = fullPath;
            _resultValue.id = id;
            _resultValue.kind = kind;
            _resultValue.name = name;
            _resultValue.path = path;
            return _resultValue;
        }
    }
}
