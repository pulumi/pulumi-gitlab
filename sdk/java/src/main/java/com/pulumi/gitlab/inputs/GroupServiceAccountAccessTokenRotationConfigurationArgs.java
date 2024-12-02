// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;


public final class GroupServiceAccountAccessTokenRotationConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupServiceAccountAccessTokenRotationConfigurationArgs Empty = new GroupServiceAccountAccessTokenRotationConfigurationArgs();

    /**
     * The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
     * 
     */
    @Import(name="rotateBeforeDays", required=true)
    private Output<Integer> rotateBeforeDays;

    /**
     * @return The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
     * 
     */
    public Output<Integer> rotateBeforeDays() {
        return this.rotateBeforeDays;
    }

    private GroupServiceAccountAccessTokenRotationConfigurationArgs() {}

    private GroupServiceAccountAccessTokenRotationConfigurationArgs(GroupServiceAccountAccessTokenRotationConfigurationArgs $) {
        this.rotateBeforeDays = $.rotateBeforeDays;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupServiceAccountAccessTokenRotationConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupServiceAccountAccessTokenRotationConfigurationArgs $;

        public Builder() {
            $ = new GroupServiceAccountAccessTokenRotationConfigurationArgs();
        }

        public Builder(GroupServiceAccountAccessTokenRotationConfigurationArgs defaults) {
            $ = new GroupServiceAccountAccessTokenRotationConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param rotateBeforeDays The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
         * 
         * @return builder
         * 
         */
        public Builder rotateBeforeDays(Output<Integer> rotateBeforeDays) {
            $.rotateBeforeDays = rotateBeforeDays;
            return this;
        }

        /**
         * @param rotateBeforeDays The duration (in days) before the expiration when the token should be rotated. As an example, if set to 7 days, the token will rotate 7 days before the expiration date, but only when `pulumi up` is run in that timeframe.
         * 
         * @return builder
         * 
         */
        public Builder rotateBeforeDays(Integer rotateBeforeDays) {
            return rotateBeforeDays(Output.of(rotateBeforeDays));
        }

        public GroupServiceAccountAccessTokenRotationConfigurationArgs build() {
            if ($.rotateBeforeDays == null) {
                throw new MissingRequiredPropertyException("GroupServiceAccountAccessTokenRotationConfigurationArgs", "rotateBeforeDays");
            }
            return $;
        }
    }

}
