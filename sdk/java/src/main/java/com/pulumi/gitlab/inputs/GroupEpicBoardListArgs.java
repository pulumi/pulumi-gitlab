// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupEpicBoardListArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupEpicBoardListArgs Empty = new GroupEpicBoardListArgs();

    /**
     * The ID of the list.
     * 
     */
    @Import(name="id")
    private @Nullable Output<Integer> id;

    /**
     * @return The ID of the list.
     * 
     */
    public Optional<Output<Integer>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The ID of the label the list should be scoped to.
     * 
     */
    @Import(name="labelId")
    private @Nullable Output<Integer> labelId;

    /**
     * @return The ID of the label the list should be scoped to.
     * 
     */
    public Optional<Output<Integer>> labelId() {
        return Optional.ofNullable(this.labelId);
    }

    /**
     * The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
     * 
     */
    @Import(name="position")
    private @Nullable Output<Integer> position;

    /**
     * @return The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
     * 
     */
    public Optional<Output<Integer>> position() {
        return Optional.ofNullable(this.position);
    }

    private GroupEpicBoardListArgs() {}

    private GroupEpicBoardListArgs(GroupEpicBoardListArgs $) {
        this.id = $.id;
        this.labelId = $.labelId;
        this.position = $.position;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupEpicBoardListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupEpicBoardListArgs $;

        public Builder() {
            $ = new GroupEpicBoardListArgs();
        }

        public Builder(GroupEpicBoardListArgs defaults) {
            $ = new GroupEpicBoardListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The ID of the list.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<Integer> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The ID of the list.
         * 
         * @return builder
         * 
         */
        public Builder id(Integer id) {
            return id(Output.of(id));
        }

        /**
         * @param labelId The ID of the label the list should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder labelId(@Nullable Output<Integer> labelId) {
            $.labelId = labelId;
            return this;
        }

        /**
         * @param labelId The ID of the label the list should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder labelId(Integer labelId) {
            return labelId(Output.of(labelId));
        }

        /**
         * @param position The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
         * 
         * @return builder
         * 
         */
        public Builder position(@Nullable Output<Integer> position) {
            $.position = position;
            return this;
        }

        /**
         * @param position The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
         * 
         * @return builder
         * 
         */
        public Builder position(Integer position) {
            return position(Output.of(position));
        }

        public GroupEpicBoardListArgs build() {
            return $;
        }
    }

}
