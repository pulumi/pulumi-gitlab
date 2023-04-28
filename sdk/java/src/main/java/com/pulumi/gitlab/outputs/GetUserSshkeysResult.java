// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetUserSshkeysKey;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetUserSshkeysResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The user&#39;s keys.
     * 
     */
    private List<GetUserSshkeysKey> keys;
    /**
     * @return ID of the user to get the SSH keys for.
     * 
     */
    private Integer userId;
    /**
     * @return Username of the user to get the SSH keys for.
     * 
     */
    private String username;

    private GetUserSshkeysResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The user&#39;s keys.
     * 
     */
    public List<GetUserSshkeysKey> keys() {
        return this.keys;
    }
    /**
     * @return ID of the user to get the SSH keys for.
     * 
     */
    public Integer userId() {
        return this.userId;
    }
    /**
     * @return Username of the user to get the SSH keys for.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserSshkeysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetUserSshkeysKey> keys;
        private Integer userId;
        private String username;
        public Builder() {}
        public Builder(GetUserSshkeysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.keys = defaults.keys;
    	      this.userId = defaults.userId;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder keys(List<GetUserSshkeysKey> keys) {
            this.keys = Objects.requireNonNull(keys);
            return this;
        }
        public Builder keys(GetUserSshkeysKey... keys) {
            return keys(List.of(keys));
        }
        @CustomType.Setter
        public Builder userId(Integer userId) {
            this.userId = Objects.requireNonNull(userId);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public GetUserSshkeysResult build() {
            final var o = new GetUserSshkeysResult();
            o.id = id;
            o.keys = keys;
            o.userId = userId;
            o.username = username;
            return o;
        }
    }
}
