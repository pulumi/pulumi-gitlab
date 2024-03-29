// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess {
    /**
     * @return The creation date of the project. In RFC3339 format.
     * 
     */
    private String createdAt;
    /**
     * @return The description of the project.
     * 
     */
    private String description;
    /**
     * @return The ID of the project.
     * 
     */
    private Integer id;
    /**
     * @return The name of the project.
     * 
     */
    private String name;
    /**
     * @return The name of the project with namespace.
     * 
     */
    private String nameWithNamespace;
    /**
     * @return The path of the project.
     * 
     */
    private String path;
    /**
     * @return The path of the project with namespace.
     * 
     */
    private String pathWithNamespace;

    private GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess() {}
    /**
     * @return The creation date of the project. In RFC3339 format.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The description of the project.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the project.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The name of the project.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The name of the project with namespace.
     * 
     */
    public String nameWithNamespace() {
        return this.nameWithNamespace;
    }
    /**
     * @return The path of the project.
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return The path of the project with namespace.
     * 
     */
    public String pathWithNamespace() {
        return this.pathWithNamespace;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String description;
        private Integer id;
        private String name;
        private String nameWithNamespace;
        private String path;
        private String pathWithNamespace;
        public Builder() {}
        public Builder(GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.nameWithNamespace = defaults.nameWithNamespace;
    	      this.path = defaults.path;
    	      this.pathWithNamespace = defaults.pathWithNamespace;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nameWithNamespace(String nameWithNamespace) {
            if (nameWithNamespace == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "nameWithNamespace");
            }
            this.nameWithNamespace = nameWithNamespace;
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            if (path == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "path");
            }
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder pathWithNamespace(String pathWithNamespace) {
            if (pathWithNamespace == null) {
              throw new MissingRequiredPropertyException("GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess", "pathWithNamespace");
            }
            this.pathWithNamespace = pathWithNamespace;
            return this;
        }
        public GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess build() {
            final var _resultValue = new GetInstanceDeployKeysDeployKeyProjectsWithWriteAccess();
            _resultValue.createdAt = createdAt;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.nameWithNamespace = nameWithNamespace;
            _resultValue.path = path;
            _resultValue.pathWithNamespace = pathWithNamespace;
            return _resultValue;
        }
    }
}
