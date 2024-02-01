// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupMembershipMember {
    /**
     * @return The level of access to the group.
     * 
     */
    private String accessLevel;
    /**
     * @return The avatar URL of the user.
     * 
     */
    private String avatarUrl;
    /**
     * @return Expiration date for the group membership.
     * 
     */
    private String expiresAt;
    /**
     * @return The unique id assigned to the user by the gitlab server.
     * 
     */
    private Integer id;
    /**
     * @return The name of the user.
     * 
     */
    private String name;
    /**
     * @return Whether the user is active or blocked.
     * 
     */
    private String state;
    /**
     * @return The username of the user.
     * 
     */
    private String username;
    /**
     * @return User&#39;s website URL.
     * 
     */
    private String webUrl;

    private GetGroupMembershipMember() {}
    /**
     * @return The level of access to the group.
     * 
     */
    public String accessLevel() {
        return this.accessLevel;
    }
    /**
     * @return The avatar URL of the user.
     * 
     */
    public String avatarUrl() {
        return this.avatarUrl;
    }
    /**
     * @return Expiration date for the group membership.
     * 
     */
    public String expiresAt() {
        return this.expiresAt;
    }
    /**
     * @return The unique id assigned to the user by the gitlab server.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The name of the user.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Whether the user is active or blocked.
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return The username of the user.
     * 
     */
    public String username() {
        return this.username;
    }
    /**
     * @return User&#39;s website URL.
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupMembershipMember defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessLevel;
        private String avatarUrl;
        private String expiresAt;
        private Integer id;
        private String name;
        private String state;
        private String username;
        private String webUrl;
        public Builder() {}
        public Builder(GetGroupMembershipMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLevel = defaults.accessLevel;
    	      this.avatarUrl = defaults.avatarUrl;
    	      this.expiresAt = defaults.expiresAt;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.state = defaults.state;
    	      this.username = defaults.username;
    	      this.webUrl = defaults.webUrl;
        }

        @CustomType.Setter
        public Builder accessLevel(String accessLevel) {
            if (accessLevel == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "accessLevel");
            }
            this.accessLevel = accessLevel;
            return this;
        }
        @CustomType.Setter
        public Builder avatarUrl(String avatarUrl) {
            if (avatarUrl == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "avatarUrl");
            }
            this.avatarUrl = avatarUrl;
            return this;
        }
        @CustomType.Setter
        public Builder expiresAt(String expiresAt) {
            if (expiresAt == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "expiresAt");
            }
            this.expiresAt = expiresAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "username");
            }
            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            if (webUrl == null) {
              throw new MissingRequiredPropertyException("GetGroupMembershipMember", "webUrl");
            }
            this.webUrl = webUrl;
            return this;
        }
        public GetGroupMembershipMember build() {
            final var _resultValue = new GetGroupMembershipMember();
            _resultValue.accessLevel = accessLevel;
            _resultValue.avatarUrl = avatarUrl;
            _resultValue.expiresAt = expiresAt;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.state = state;
            _resultValue.username = username;
            _resultValue.webUrl = webUrl;
            return _resultValue;
        }
    }
}
