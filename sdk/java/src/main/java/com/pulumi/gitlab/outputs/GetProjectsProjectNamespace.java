// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProjectNamespace {
    private String fullPath;
    private Integer id;
    private String kind;
    private String name;
    private String path;

    private GetProjectsProjectNamespace() {}
    public String fullPath() {
        return this.fullPath;
    }
    public Integer id() {
        return this.id;
    }
    public String kind() {
        return this.kind;
    }
    public String name() {
        return this.name;
    }
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
            this.fullPath = Objects.requireNonNull(fullPath);
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder kind(String kind) {
            this.kind = Objects.requireNonNull(kind);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public GetProjectsProjectNamespace build() {
            final var o = new GetProjectsProjectNamespace();
            o.fullPath = fullPath;
            o.id = id;
            o.kind = kind;
            o.name = name;
            o.path = path;
            return o;
        }
    }
}
