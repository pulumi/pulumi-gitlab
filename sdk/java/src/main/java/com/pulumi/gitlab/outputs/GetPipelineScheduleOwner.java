// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPipelineScheduleOwner {
    /**
     * @return Image URL for the user&#39;s avatar.
     * 
     */
    private String avatarUrl;
    /**
     * @return The user ID.
     * 
     */
    private Integer id;
    /**
     * @return Name.
     * 
     */
    private String name;
    /**
     * @return User&#39;s state, one of: active, blocked.
     * 
     */
    private String state;
    /**
     * @return Username.
     * 
     */
    private String username;
    /**
     * @return URL to the user&#39;s profile.
     * 
     */
    private String webUrl;

    private GetPipelineScheduleOwner() {}
    /**
     * @return Image URL for the user&#39;s avatar.
     * 
     */
    public String avatarUrl() {
        return this.avatarUrl;
    }
    /**
     * @return The user ID.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return Name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return User&#39;s state, one of: active, blocked.
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return Username.
     * 
     */
    public String username() {
        return this.username;
    }
    /**
     * @return URL to the user&#39;s profile.
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPipelineScheduleOwner defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String avatarUrl;
        private Integer id;
        private String name;
        private String state;
        private String username;
        private String webUrl;
        public Builder() {}
        public Builder(GetPipelineScheduleOwner defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.avatarUrl = defaults.avatarUrl;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.state = defaults.state;
    	      this.username = defaults.username;
    	      this.webUrl = defaults.webUrl;
        }

        @CustomType.Setter
        public Builder avatarUrl(String avatarUrl) {
            if (avatarUrl == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleOwner", "avatarUrl");
            }
            this.avatarUrl = avatarUrl;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleOwner", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleOwner", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleOwner", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleOwner", "username");
            }
            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            if (webUrl == null) {
              throw new MissingRequiredPropertyException("GetPipelineScheduleOwner", "webUrl");
            }
            this.webUrl = webUrl;
            return this;
        }
        public GetPipelineScheduleOwner build() {
            final var _resultValue = new GetPipelineScheduleOwner();
            _resultValue.avatarUrl = avatarUrl;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.state = state;
            _resultValue.username = username;
            _resultValue.webUrl = webUrl;
            return _resultValue;
        }
    }
}
