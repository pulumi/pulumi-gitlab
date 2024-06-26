// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GroupAccessTokenRotationConfiguration {
    /**
     * @return The duration (in days) the new token should be valid for.
     * 
     */
    private Integer expirationDays;
    /**
     * @return The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
     * 
     */
    private Integer rotateBeforeDays;

    private GroupAccessTokenRotationConfiguration() {}
    /**
     * @return The duration (in days) the new token should be valid for.
     * 
     */
    public Integer expirationDays() {
        return this.expirationDays;
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

    public static Builder builder(GroupAccessTokenRotationConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer expirationDays;
        private Integer rotateBeforeDays;
        public Builder() {}
        public Builder(GroupAccessTokenRotationConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expirationDays = defaults.expirationDays;
    	      this.rotateBeforeDays = defaults.rotateBeforeDays;
        }

        @CustomType.Setter
        public Builder expirationDays(Integer expirationDays) {
            if (expirationDays == null) {
              throw new MissingRequiredPropertyException("GroupAccessTokenRotationConfiguration", "expirationDays");
            }
            this.expirationDays = expirationDays;
            return this;
        }
        @CustomType.Setter
        public Builder rotateBeforeDays(Integer rotateBeforeDays) {
            if (rotateBeforeDays == null) {
              throw new MissingRequiredPropertyException("GroupAccessTokenRotationConfiguration", "rotateBeforeDays");
            }
            this.rotateBeforeDays = rotateBeforeDays;
            return this;
        }
        public GroupAccessTokenRotationConfiguration build() {
            final var _resultValue = new GroupAccessTokenRotationConfiguration();
            _resultValue.expirationDays = expirationDays;
            _resultValue.rotateBeforeDays = rotateBeforeDays;
            return _resultValue;
        }
    }
}
