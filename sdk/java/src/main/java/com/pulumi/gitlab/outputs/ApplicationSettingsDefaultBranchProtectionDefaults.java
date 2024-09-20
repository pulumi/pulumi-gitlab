// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationSettingsDefaultBranchProtectionDefaults {
    /**
     * @return Allow force push for all users with push access.
     * 
     */
    private @Nullable Boolean allowForcePush;
    /**
     * @return An array of access levels allowed to merge. Supports Developer (30) or Maintainer (40).
     * 
     */
    private @Nullable List<Object> allowedToMerges;
    /**
     * @return An array of access levels allowed to push. Supports Developer (30) or Maintainer (40).
     * 
     */
    private @Nullable List<Object> allowedToPushes;
    /**
     * @return Allow developers to initial push.
     * 
     */
    private @Nullable Boolean developerCanInitialPush;

    private ApplicationSettingsDefaultBranchProtectionDefaults() {}
    /**
     * @return Allow force push for all users with push access.
     * 
     */
    public Optional<Boolean> allowForcePush() {
        return Optional.ofNullable(this.allowForcePush);
    }
    /**
     * @return An array of access levels allowed to merge. Supports Developer (30) or Maintainer (40).
     * 
     */
    public List<Object> allowedToMerges() {
        return this.allowedToMerges == null ? List.of() : this.allowedToMerges;
    }
    /**
     * @return An array of access levels allowed to push. Supports Developer (30) or Maintainer (40).
     * 
     */
    public List<Object> allowedToPushes() {
        return this.allowedToPushes == null ? List.of() : this.allowedToPushes;
    }
    /**
     * @return Allow developers to initial push.
     * 
     */
    public Optional<Boolean> developerCanInitialPush() {
        return Optional.ofNullable(this.developerCanInitialPush);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationSettingsDefaultBranchProtectionDefaults defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowForcePush;
        private @Nullable List<Object> allowedToMerges;
        private @Nullable List<Object> allowedToPushes;
        private @Nullable Boolean developerCanInitialPush;
        public Builder() {}
        public Builder(ApplicationSettingsDefaultBranchProtectionDefaults defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowForcePush = defaults.allowForcePush;
    	      this.allowedToMerges = defaults.allowedToMerges;
    	      this.allowedToPushes = defaults.allowedToPushes;
    	      this.developerCanInitialPush = defaults.developerCanInitialPush;
        }

        @CustomType.Setter
        public Builder allowForcePush(@Nullable Boolean allowForcePush) {

            this.allowForcePush = allowForcePush;
            return this;
        }
        @CustomType.Setter
        public Builder allowedToMerges(@Nullable List<Object> allowedToMerges) {

            this.allowedToMerges = allowedToMerges;
            return this;
        }
        public Builder allowedToMerges(Object... allowedToMerges) {
            return allowedToMerges(List.of(allowedToMerges));
        }
        @CustomType.Setter
        public Builder allowedToPushes(@Nullable List<Object> allowedToPushes) {

            this.allowedToPushes = allowedToPushes;
            return this;
        }
        public Builder allowedToPushes(Object... allowedToPushes) {
            return allowedToPushes(List.of(allowedToPushes));
        }
        @CustomType.Setter
        public Builder developerCanInitialPush(@Nullable Boolean developerCanInitialPush) {

            this.developerCanInitialPush = developerCanInitialPush;
            return this;
        }
        public ApplicationSettingsDefaultBranchProtectionDefaults build() {
            final var _resultValue = new ApplicationSettingsDefaultBranchProtectionDefaults();
            _resultValue.allowForcePush = allowForcePush;
            _resultValue.allowedToMerges = allowedToMerges;
            _resultValue.allowedToPushes = allowedToPushes;
            _resultValue.developerCanInitialPush = developerCanInitialPush;
            return _resultValue;
        }
    }
}