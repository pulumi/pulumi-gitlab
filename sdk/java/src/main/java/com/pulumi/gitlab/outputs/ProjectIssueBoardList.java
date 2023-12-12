// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProjectIssueBoardList {
    /**
     * @return The ID of the assignee the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    private @Nullable Integer assigneeId;
    /**
     * @return The ID of the list
     * 
     */
    private @Nullable Integer id;
    /**
     * @return The ID of the iteration the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    private @Nullable Integer iterationId;
    /**
     * @return The ID of the label the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    private @Nullable Integer labelId;
    /**
     * @return The ID of the milestone the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    private @Nullable Integer milestoneId;
    /**
     * @return The position of the list within the board. The position for the list is based on the its position in the `lists` array.
     * 
     */
    private @Nullable Integer position;

    private ProjectIssueBoardList() {}
    /**
     * @return The ID of the assignee the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Optional<Integer> assigneeId() {
        return Optional.ofNullable(this.assigneeId);
    }
    /**
     * @return The ID of the list
     * 
     */
    public Optional<Integer> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The ID of the iteration the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Optional<Integer> iterationId() {
        return Optional.ofNullable(this.iterationId);
    }
    /**
     * @return The ID of the label the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Optional<Integer> labelId() {
        return Optional.ofNullable(this.labelId);
    }
    /**
     * @return The ID of the milestone the list should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Optional<Integer> milestoneId() {
        return Optional.ofNullable(this.milestoneId);
    }
    /**
     * @return The position of the list within the board. The position for the list is based on the its position in the `lists` array.
     * 
     */
    public Optional<Integer> position() {
        return Optional.ofNullable(this.position);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProjectIssueBoardList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer assigneeId;
        private @Nullable Integer id;
        private @Nullable Integer iterationId;
        private @Nullable Integer labelId;
        private @Nullable Integer milestoneId;
        private @Nullable Integer position;
        public Builder() {}
        public Builder(ProjectIssueBoardList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.assigneeId = defaults.assigneeId;
    	      this.id = defaults.id;
    	      this.iterationId = defaults.iterationId;
    	      this.labelId = defaults.labelId;
    	      this.milestoneId = defaults.milestoneId;
    	      this.position = defaults.position;
        }

        @CustomType.Setter
        public Builder assigneeId(@Nullable Integer assigneeId) {
            this.assigneeId = assigneeId;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable Integer id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder iterationId(@Nullable Integer iterationId) {
            this.iterationId = iterationId;
            return this;
        }
        @CustomType.Setter
        public Builder labelId(@Nullable Integer labelId) {
            this.labelId = labelId;
            return this;
        }
        @CustomType.Setter
        public Builder milestoneId(@Nullable Integer milestoneId) {
            this.milestoneId = milestoneId;
            return this;
        }
        @CustomType.Setter
        public Builder position(@Nullable Integer position) {
            this.position = position;
            return this;
        }
        public ProjectIssueBoardList build() {
            final var _resultValue = new ProjectIssueBoardList();
            _resultValue.assigneeId = assigneeId;
            _resultValue.id = id;
            _resultValue.iterationId = iterationId;
            _resultValue.labelId = labelId;
            _resultValue.milestoneId = milestoneId;
            _resultValue.position = position;
            return _resultValue;
        }
    }
}
