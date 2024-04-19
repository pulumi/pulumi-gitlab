// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;


public final class GroupAccessTokenRotationConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupAccessTokenRotationConfigurationArgs Empty = new GroupAccessTokenRotationConfigurationArgs();

    /**
     * The duration (in days) the new token should be valid for.
     * 
     */
    @Import(name="expirationDays", required=true)
    private Output<Integer> expirationDays;

    /**
     * @return The duration (in days) the new token should be valid for.
     * 
     */
    public Output<Integer> expirationDays() {
        return this.expirationDays;
    }

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

    private GroupAccessTokenRotationConfigurationArgs() {}

    private GroupAccessTokenRotationConfigurationArgs(GroupAccessTokenRotationConfigurationArgs $) {
        this.expirationDays = $.expirationDays;
        this.rotateBeforeDays = $.rotateBeforeDays;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupAccessTokenRotationConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupAccessTokenRotationConfigurationArgs $;

        public Builder() {
            $ = new GroupAccessTokenRotationConfigurationArgs();
        }

        public Builder(GroupAccessTokenRotationConfigurationArgs defaults) {
            $ = new GroupAccessTokenRotationConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param expirationDays The duration (in days) the new token should be valid for.
         * 
         * @return builder
         * 
         */
        public Builder expirationDays(Output<Integer> expirationDays) {
            $.expirationDays = expirationDays;
            return this;
        }

        /**
         * @param expirationDays The duration (in days) the new token should be valid for.
         * 
         * @return builder
         * 
         */
        public Builder expirationDays(Integer expirationDays) {
            return expirationDays(Output.of(expirationDays));
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

        public GroupAccessTokenRotationConfigurationArgs build() {
            if ($.expirationDays == null) {
                throw new MissingRequiredPropertyException("GroupAccessTokenRotationConfigurationArgs", "expirationDays");
            }
            if ($.rotateBeforeDays == null) {
                throw new MissingRequiredPropertyException("GroupAccessTokenRotationConfigurationArgs", "rotateBeforeDays");
            }
            return $;
        }
    }

}
