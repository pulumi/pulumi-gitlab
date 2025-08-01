// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectAccessTokensAccessToken;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectAccessTokensResult {
    /**
     * @return The list of access tokens returned by the search
     * 
     */
    private List<GetProjectAccessTokensAccessToken> accessTokens;
    private String id;
    /**
     * @return The name or id of the project.
     * 
     */
    private String project;
    /**
     * @return List all project access token that match the specified state. Valid values are `active`, `inactive`. Returns all project access token if not set.
     * 
     */
    private @Nullable String state;

    private GetProjectAccessTokensResult() {}
    /**
     * @return The list of access tokens returned by the search
     * 
     */
    public List<GetProjectAccessTokensAccessToken> accessTokens() {
        return this.accessTokens;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The name or id of the project.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return List all project access token that match the specified state. Valid values are `active`, `inactive`. Returns all project access token if not set.
     * 
     */
    public Optional<String> state() {
        return Optional.ofNullable(this.state);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectAccessTokensResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetProjectAccessTokensAccessToken> accessTokens;
        private String id;
        private String project;
        private @Nullable String state;
        public Builder() {}
        public Builder(GetProjectAccessTokensResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessTokens = defaults.accessTokens;
    	      this.id = defaults.id;
    	      this.project = defaults.project;
    	      this.state = defaults.state;
        }

        @CustomType.Setter
        public Builder accessTokens(List<GetProjectAccessTokensAccessToken> accessTokens) {
            if (accessTokens == null) {
              throw new MissingRequiredPropertyException("GetProjectAccessTokensResult", "accessTokens");
            }
            this.accessTokens = accessTokens;
            return this;
        }
        public Builder accessTokens(GetProjectAccessTokensAccessToken... accessTokens) {
            return accessTokens(List.of(accessTokens));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectAccessTokensResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetProjectAccessTokensResult", "project");
            }
            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder state(@Nullable String state) {

            this.state = state;
            return this;
        }
        public GetProjectAccessTokensResult build() {
            final var _resultValue = new GetProjectAccessTokensResult();
            _resultValue.accessTokens = accessTokens;
            _resultValue.id = id;
            _resultValue.project = project;
            _resultValue.state = state;
            return _resultValue;
        }
    }
}
