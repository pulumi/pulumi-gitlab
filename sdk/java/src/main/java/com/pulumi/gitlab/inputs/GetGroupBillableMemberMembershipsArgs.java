// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetGroupBillableMemberMembershipsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupBillableMemberMembershipsArgs Empty = new GetGroupBillableMemberMembershipsArgs();

    /**
     * The ID of the group.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return The ID of the group.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * The ID of the user.
     * 
     */
    @Import(name="userId", required=true)
    private Output<Integer> userId;

    /**
     * @return The ID of the user.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    private GetGroupBillableMemberMembershipsArgs() {}

    private GetGroupBillableMemberMembershipsArgs(GetGroupBillableMemberMembershipsArgs $) {
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupBillableMemberMembershipsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupBillableMemberMembershipsArgs $;

        public Builder() {
            $ = new GetGroupBillableMemberMembershipsArgs();
        }

        public Builder(GetGroupBillableMemberMembershipsArgs defaults) {
            $ = new GetGroupBillableMemberMembershipsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Output<Integer> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            return userId(Output.of(userId));
        }

        public GetGroupBillableMemberMembershipsArgs build() {
            if ($.groupId == null) {
                throw new MissingRequiredPropertyException("GetGroupBillableMemberMembershipsArgs", "groupId");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("GetGroupBillableMemberMembershipsArgs", "userId");
            }
            return $;
        }
    }

}
