// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetInstanceServiceAccountPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceServiceAccountPlainArgs Empty = new GetInstanceServiceAccountPlainArgs();

    /**
     * The service account id.
     * 
     */
    @Import(name="serviceAccountId", required=true)
    private String serviceAccountId;

    /**
     * @return The service account id.
     * 
     */
    public String serviceAccountId() {
        return this.serviceAccountId;
    }

    private GetInstanceServiceAccountPlainArgs() {}

    private GetInstanceServiceAccountPlainArgs(GetInstanceServiceAccountPlainArgs $) {
        this.serviceAccountId = $.serviceAccountId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceServiceAccountPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceServiceAccountPlainArgs $;

        public Builder() {
            $ = new GetInstanceServiceAccountPlainArgs();
        }

        public Builder(GetInstanceServiceAccountPlainArgs defaults) {
            $ = new GetInstanceServiceAccountPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceAccountId The service account id.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            $.serviceAccountId = serviceAccountId;
            return this;
        }

        public GetInstanceServiceAccountPlainArgs build() {
            if ($.serviceAccountId == null) {
                throw new MissingRequiredPropertyException("GetInstanceServiceAccountPlainArgs", "serviceAccountId");
            }
            return $;
        }
    }

}
