// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetGroupHookArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupHookArgs Empty = new GetGroupHookArgs();

    /**
     * The ID or full path of the group.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or full path of the group.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * The id of the group hook.
     * 
     */
    @Import(name="hookId", required=true)
    private Output<Integer> hookId;

    /**
     * @return The id of the group hook.
     * 
     */
    public Output<Integer> hookId() {
        return this.hookId;
    }

    private GetGroupHookArgs() {}

    private GetGroupHookArgs(GetGroupHookArgs $) {
        this.group = $.group;
        this.hookId = $.hookId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupHookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupHookArgs $;

        public Builder() {
            $ = new GetGroupHookArgs();
        }

        public Builder(GetGroupHookArgs defaults) {
            $ = new GetGroupHookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param group The ID or full path of the group.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or full path of the group.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param hookId The id of the group hook.
         * 
         * @return builder
         * 
         */
        public Builder hookId(Output<Integer> hookId) {
            $.hookId = hookId;
            return this;
        }

        /**
         * @param hookId The id of the group hook.
         * 
         * @return builder
         * 
         */
        public Builder hookId(Integer hookId) {
            return hookId(Output.of(hookId));
        }

        public GetGroupHookArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GetGroupHookArgs", "group");
            }
            if ($.hookId == null) {
                throw new MissingRequiredPropertyException("GetGroupHookArgs", "hookId");
            }
            return $;
        }
    }

}
