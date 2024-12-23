// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetGroupBillableMemberMembershipsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupBillableMemberMembershipsPlainArgs Empty = new GetGroupBillableMemberMembershipsPlainArgs();

    /**
     * The ID of the group.
     * 
     */
    @Import(name="groupId", required=true)
    private String groupId;

    /**
     * @return The ID of the group.
     * 
     */
    public String groupId() {
        return this.groupId;
    }

    /**
     * The ID of the user.
     * 
     */
    @Import(name="userId", required=true)
    private Integer userId;

    /**
     * @return The ID of the user.
     * 
     */
    public Integer userId() {
        return this.userId;
    }

    private GetGroupBillableMemberMembershipsPlainArgs() {}

    private GetGroupBillableMemberMembershipsPlainArgs(GetGroupBillableMemberMembershipsPlainArgs $) {
        this.groupId = $.groupId;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupBillableMemberMembershipsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupBillableMemberMembershipsPlainArgs $;

        public Builder() {
            $ = new GetGroupBillableMemberMembershipsPlainArgs();
        }

        public Builder(GetGroupBillableMemberMembershipsPlainArgs defaults) {
            $ = new GetGroupBillableMemberMembershipsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId The ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param userId The ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Integer userId) {
            $.userId = userId;
            return this;
        }

        public GetGroupBillableMemberMembershipsPlainArgs build() {
            if ($.groupId == null) {
                throw new MissingRequiredPropertyException("GetGroupBillableMemberMembershipsPlainArgs", "groupId");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("GetGroupBillableMemberMembershipsPlainArgs", "userId");
            }
            return $;
        }
    }

}
