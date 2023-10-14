// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryTreeTree {
    private String id;
    private String mode;
    private String name;
    private String path;
    private String type;

    private GetRepositoryTreeTree() {}
    public String id() {
        return this.id;
    }
    public String mode() {
        return this.mode;
    }
    public String name() {
        return this.name;
    }
    public String path() {
        return this.path;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryTreeTree defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String mode;
        private String name;
        private String path;
        private String type;
        public Builder() {}
        public Builder(GetRepositoryTreeTree defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.path = defaults.path;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            this.mode = Objects.requireNonNull(mode);
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
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetRepositoryTreeTree build() {
            final var o = new GetRepositoryTreeTree();
            o.id = id;
            o.mode = mode;
            o.name = name;
            o.path = path;
            o.type = type;
            return o;
        }
    }
}
