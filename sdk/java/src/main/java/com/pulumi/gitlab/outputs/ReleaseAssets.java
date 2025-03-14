// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ReleaseAssets {
    /**
     * @return The total count of assets in this release.
     * 
     */
    private @Nullable Integer count;

    private ReleaseAssets() {}
    /**
     * @return The total count of assets in this release.
     * 
     */
    public Optional<Integer> count() {
        return Optional.ofNullable(this.count);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReleaseAssets defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer count;
        public Builder() {}
        public Builder(ReleaseAssets defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.count = defaults.count;
        }

        @CustomType.Setter
        public Builder count(@Nullable Integer count) {

            this.count = count;
            return this;
        }
        public ReleaseAssets build() {
            final var _resultValue = new ReleaseAssets();
            _resultValue.count = count;
            return _resultValue;
        }
    }
}
