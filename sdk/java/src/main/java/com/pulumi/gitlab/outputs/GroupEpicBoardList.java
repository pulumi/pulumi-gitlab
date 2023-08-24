// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupEpicBoardList {
    /**
     * @return The ID of the list.
     * 
     */
    private @Nullable Integer id;
    /**
     * @return The ID of the label the list should be scoped to.
     * 
     */
    private @Nullable Integer labelId;
    /**
     * @return The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
     * 
     */
    private @Nullable Integer position;

    private GroupEpicBoardList() {}
    /**
     * @return The ID of the list.
     * 
     */
    public Optional<Integer> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The ID of the label the list should be scoped to.
     * 
     */
    public Optional<Integer> labelId() {
        return Optional.ofNullable(this.labelId);
    }
    /**
     * @return The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
     * 
     */
    public Optional<Integer> position() {
        return Optional.ofNullable(this.position);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupEpicBoardList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer id;
        private @Nullable Integer labelId;
        private @Nullable Integer position;
        public Builder() {}
        public Builder(GroupEpicBoardList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.labelId = defaults.labelId;
    	      this.position = defaults.position;
        }

        @CustomType.Setter
        public Builder id(@Nullable Integer id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder labelId(@Nullable Integer labelId) {
            this.labelId = labelId;
            return this;
        }
        @CustomType.Setter
        public Builder position(@Nullable Integer position) {
            this.position = position;
            return this;
        }
        public GroupEpicBoardList build() {
            final var o = new GroupEpicBoardList();
            o.id = id;
            o.labelId = labelId;
            o.position = position;
            return o;
        }
    }
}
