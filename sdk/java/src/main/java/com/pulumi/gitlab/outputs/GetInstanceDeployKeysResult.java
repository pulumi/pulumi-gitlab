// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetInstanceDeployKeysDeployKey;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceDeployKeysResult {
    /**
     * @return The list of all deploy keys across all projects of the GitLab instance.
     * 
     */
    private List<GetInstanceDeployKeysDeployKey> deployKeys;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Only return deploy keys that are public.
     * 
     */
    private @Nullable Boolean public_;

    private GetInstanceDeployKeysResult() {}
    /**
     * @return The list of all deploy keys across all projects of the GitLab instance.
     * 
     */
    public List<GetInstanceDeployKeysDeployKey> deployKeys() {
        return this.deployKeys;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Only return deploy keys that are public.
     * 
     */
    public Optional<Boolean> public_() {
        return Optional.ofNullable(this.public_);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceDeployKeysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetInstanceDeployKeysDeployKey> deployKeys;
        private String id;
        private @Nullable Boolean public_;
        public Builder() {}
        public Builder(GetInstanceDeployKeysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.deployKeys = defaults.deployKeys;
    	      this.id = defaults.id;
    	      this.public_ = defaults.public_;
        }

        @CustomType.Setter
        public Builder deployKeys(List<GetInstanceDeployKeysDeployKey> deployKeys) {
            this.deployKeys = Objects.requireNonNull(deployKeys);
            return this;
        }
        public Builder deployKeys(GetInstanceDeployKeysDeployKey... deployKeys) {
            return deployKeys(List.of(deployKeys));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter("public")
        public Builder public_(@Nullable Boolean public_) {
            this.public_ = public_;
            return this;
        }
        public GetInstanceDeployKeysResult build() {
            final var _resultValue = new GetInstanceDeployKeysResult();
            _resultValue.deployKeys = deployKeys;
            _resultValue.id = id;
            _resultValue.public_ = public_;
            return _resultValue;
        }
    }
}
