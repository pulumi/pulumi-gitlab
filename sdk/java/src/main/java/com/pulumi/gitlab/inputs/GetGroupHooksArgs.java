// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetGroupHooksArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupHooksArgs Empty = new GetGroupHooksArgs();

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

    private GetGroupHooksArgs() {}

    private GetGroupHooksArgs(GetGroupHooksArgs $) {
        this.group = $.group;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupHooksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupHooksArgs $;

        public Builder() {
            $ = new GetGroupHooksArgs();
        }

        public Builder(GetGroupHooksArgs defaults) {
            $ = new GetGroupHooksArgs(Objects.requireNonNull(defaults));
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

        public GetGroupHooksArgs build() {
            $.group = Objects.requireNonNull($.group, "expected parameter 'group' to be non-null");
            return $;
        }
    }

}