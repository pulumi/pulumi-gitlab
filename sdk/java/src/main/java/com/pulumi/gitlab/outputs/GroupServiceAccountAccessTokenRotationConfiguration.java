// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupServiceAccountAccessTokenRotationConfiguration {
    /**
     * @return The duration (in days) the new token should be valid for.
     * 
     */
    private @Nullable Integer expirationDays;
    /**
     * @return The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
     * 
     */
    private Integer rotateBeforeDays;

    private GroupServiceAccountAccessTokenRotationConfiguration() {}
    /**
     * @return The duration (in days) the new token should be valid for.
     * 
     */
    public Optional<Integer> expirationDays() {
        return Optional.ofNullable(this.expirationDays);
    }
    /**
     * @return The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
     * 
     */
    public Integer rotateBeforeDays() {
        return this.rotateBeforeDays;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupServiceAccountAccessTokenRotationConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer expirationDays;
        private Integer rotateBeforeDays;
        public Builder() {}
        public Builder(GroupServiceAccountAccessTokenRotationConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expirationDays = defaults.expirationDays;
    	      this.rotateBeforeDays = defaults.rotateBeforeDays;
        }

        @CustomType.Setter
        public Builder expirationDays(@Nullable Integer expirationDays) {

            this.expirationDays = expirationDays;
            return this;
        }
        @CustomType.Setter
        public Builder rotateBeforeDays(Integer rotateBeforeDays) {
            if (rotateBeforeDays == null) {
              throw new MissingRequiredPropertyException("GroupServiceAccountAccessTokenRotationConfiguration", "rotateBeforeDays");
            }
            this.rotateBeforeDays = rotateBeforeDays;
            return this;
        }
        public GroupServiceAccountAccessTokenRotationConfiguration build() {
            final var _resultValue = new GroupServiceAccountAccessTokenRotationConfiguration();
            _resultValue.expirationDays = expirationDays;
            _resultValue.rotateBeforeDays = rotateBeforeDays;
            return _resultValue;
        }
    }
}
