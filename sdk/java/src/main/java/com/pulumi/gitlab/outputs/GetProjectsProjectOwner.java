// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProjectOwner {
    private String avatarUrl;
    private Integer id;
    private String name;
    private String state;
    private String username;
    private String websiteUrl;

    private GetProjectsProjectOwner() {}
    public String avatarUrl() {
        return this.avatarUrl;
    }
    public Integer id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public String state() {
        return this.state;
    }
    public String username() {
        return this.username;
    }
    public String websiteUrl() {
        return this.websiteUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProjectOwner defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String avatarUrl;
        private Integer id;
        private String name;
        private String state;
        private String username;
        private String websiteUrl;
        public Builder() {}
        public Builder(GetProjectsProjectOwner defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.avatarUrl = defaults.avatarUrl;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.state = defaults.state;
    	      this.username = defaults.username;
    	      this.websiteUrl = defaults.websiteUrl;
        }

        @CustomType.Setter
        public Builder avatarUrl(String avatarUrl) {
            this.avatarUrl = Objects.requireNonNull(avatarUrl);
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        @CustomType.Setter
        public Builder websiteUrl(String websiteUrl) {
            this.websiteUrl = Objects.requireNonNull(websiteUrl);
            return this;
        }
        public GetProjectsProjectOwner build() {
            final var _resultValue = new GetProjectsProjectOwner();
            _resultValue.avatarUrl = avatarUrl;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.state = state;
            _resultValue.username = username;
            _resultValue.websiteUrl = websiteUrl;
            return _resultValue;
        }
    }
}
