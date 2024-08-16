// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGroupIdsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupIdsPlainArgs Empty = new GetGroupIdsPlainArgs();

    /**
     * The ID or URL-encoded path of the group.
     * 
     */
    @Import(name="group", required=true)
    private String group;

    /**
     * @return The ID or URL-encoded path of the group.
     * 
     */
    public String group() {
        return this.group;
    }

    private GetGroupIdsPlainArgs() {}

    private GetGroupIdsPlainArgs(GetGroupIdsPlainArgs $) {
        this.group = $.group;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupIdsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupIdsPlainArgs $;

        public Builder() {
            $ = new GetGroupIdsPlainArgs();
        }

        public Builder(GetGroupIdsPlainArgs defaults) {
            $ = new GetGroupIdsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param group The ID or URL-encoded path of the group.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            $.group = group;
            return this;
        }

        public GetGroupIdsPlainArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GetGroupIdsPlainArgs", "group");
            }
            return $;
        }
    }

}
