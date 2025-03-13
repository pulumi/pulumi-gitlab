// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ValueStreamAnalyticsStage {
    /**
     * @return Boolean whether the stage is customized. If false, it assigns a built-in default stage by name.
     * 
     */
    private @Nullable Boolean custom;
    /**
     * @return End event identifier. Valid values are: `CODE_STAGE_START`, `ISSUE_CLOSED`, `ISSUE_CREATED`, `ISSUE_DEPLOYED_TO_PRODUCTION`, `ISSUE_FIRST_ADDED_TO_BOARD`, `ISSUE_FIRST_ADDED_TO_ITERATION`, `ISSUE_FIRST_ASSIGNED_AT`, `ISSUE_FIRST_ASSOCIATED_WITH_MILESTONE`, `ISSUE_FIRST_MENTIONED_IN_COMMIT`, `ISSUE_LABEL_ADDED`, `ISSUE_LABEL_REMOVED`, `ISSUE_LAST_EDITED`, `ISSUE_STAGE_END`, `MERGE_REQUEST_CLOSED`, `MERGE_REQUEST_CREATED`, `MERGE_REQUEST_FIRST_ASSIGNED_AT`, `MERGE_REQUEST_FIRST_COMMIT_AT`, `MERGE_REQUEST_FIRST_DEPLOYED_TO_PRODUCTION`, `MERGE_REQUEST_LABEL_ADDED`, `MERGE_REQUEST_LABEL_REMOVED`, `MERGE_REQUEST_LAST_BUILD_FINISHED`, `MERGE_REQUEST_LAST_BUILD_STARTED`, `MERGE_REQUEST_LAST_EDITED`, `MERGE_REQUEST_MERGED`, `MERGE_REQUEST_REVIEWER_FIRST_ASSIGNED`, `MERGE_REQUEST_PLAN_STAGE_START`
     * 
     */
    private @Nullable String endEventIdentifier;
    /**
     * @return Label ID associated with the end event identifier. In the format of `gid://gitlab/GroupLabel/&lt;id&gt;` or `gid://gitlab/ProjectLabel/&lt;id&gt;`
     * 
     */
    private @Nullable String endEventLabelId;
    /**
     * @return Boolean whether the stage is hidden, GitLab provided default stages are hidden by default.
     * 
     */
    private @Nullable Boolean hidden;
    /**
     * @return The ID of the value stream stage.
     * 
     */
    private @Nullable String id;
    /**
     * @return The name of the value stream stage.
     * 
     */
    private String name;
    /**
     * @return Start event identifier. Valid values are: `CODE_STAGE_START`, `ISSUE_CLOSED`, `ISSUE_CREATED`, `ISSUE_DEPLOYED_TO_PRODUCTION`, `ISSUE_FIRST_ADDED_TO_BOARD`, `ISSUE_FIRST_ADDED_TO_ITERATION`, `ISSUE_FIRST_ASSIGNED_AT`, `ISSUE_FIRST_ASSOCIATED_WITH_MILESTONE`, `ISSUE_FIRST_MENTIONED_IN_COMMIT`, `ISSUE_LABEL_ADDED`, `ISSUE_LABEL_REMOVED`, `ISSUE_LAST_EDITED`, `ISSUE_STAGE_END`, `MERGE_REQUEST_CLOSED`, `MERGE_REQUEST_CREATED`, `MERGE_REQUEST_FIRST_ASSIGNED_AT`, `MERGE_REQUEST_FIRST_COMMIT_AT`, `MERGE_REQUEST_FIRST_DEPLOYED_TO_PRODUCTION`, `MERGE_REQUEST_LABEL_ADDED`, `MERGE_REQUEST_LABEL_REMOVED`, `MERGE_REQUEST_LAST_BUILD_FINISHED`, `MERGE_REQUEST_LAST_BUILD_STARTED`, `MERGE_REQUEST_LAST_EDITED`, `MERGE_REQUEST_MERGED`, `MERGE_REQUEST_REVIEWER_FIRST_ASSIGNED`, `MERGE_REQUEST_PLAN_STAGE_START`
     * 
     */
    private @Nullable String startEventIdentifier;
    /**
     * @return Label ID associated with the start event identifier. In the format of `gid://gitlab/GroupLabel/&lt;id&gt;` or `gid://gitlab/ProjectLabel/&lt;id&gt;`
     * 
     */
    private @Nullable String startEventLabelId;

    private ValueStreamAnalyticsStage() {}
    /**
     * @return Boolean whether the stage is customized. If false, it assigns a built-in default stage by name.
     * 
     */
    public Optional<Boolean> custom() {
        return Optional.ofNullable(this.custom);
    }
    /**
     * @return End event identifier. Valid values are: `CODE_STAGE_START`, `ISSUE_CLOSED`, `ISSUE_CREATED`, `ISSUE_DEPLOYED_TO_PRODUCTION`, `ISSUE_FIRST_ADDED_TO_BOARD`, `ISSUE_FIRST_ADDED_TO_ITERATION`, `ISSUE_FIRST_ASSIGNED_AT`, `ISSUE_FIRST_ASSOCIATED_WITH_MILESTONE`, `ISSUE_FIRST_MENTIONED_IN_COMMIT`, `ISSUE_LABEL_ADDED`, `ISSUE_LABEL_REMOVED`, `ISSUE_LAST_EDITED`, `ISSUE_STAGE_END`, `MERGE_REQUEST_CLOSED`, `MERGE_REQUEST_CREATED`, `MERGE_REQUEST_FIRST_ASSIGNED_AT`, `MERGE_REQUEST_FIRST_COMMIT_AT`, `MERGE_REQUEST_FIRST_DEPLOYED_TO_PRODUCTION`, `MERGE_REQUEST_LABEL_ADDED`, `MERGE_REQUEST_LABEL_REMOVED`, `MERGE_REQUEST_LAST_BUILD_FINISHED`, `MERGE_REQUEST_LAST_BUILD_STARTED`, `MERGE_REQUEST_LAST_EDITED`, `MERGE_REQUEST_MERGED`, `MERGE_REQUEST_REVIEWER_FIRST_ASSIGNED`, `MERGE_REQUEST_PLAN_STAGE_START`
     * 
     */
    public Optional<String> endEventIdentifier() {
        return Optional.ofNullable(this.endEventIdentifier);
    }
    /**
     * @return Label ID associated with the end event identifier. In the format of `gid://gitlab/GroupLabel/&lt;id&gt;` or `gid://gitlab/ProjectLabel/&lt;id&gt;`
     * 
     */
    public Optional<String> endEventLabelId() {
        return Optional.ofNullable(this.endEventLabelId);
    }
    /**
     * @return Boolean whether the stage is hidden, GitLab provided default stages are hidden by default.
     * 
     */
    public Optional<Boolean> hidden() {
        return Optional.ofNullable(this.hidden);
    }
    /**
     * @return The ID of the value stream stage.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The name of the value stream stage.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Start event identifier. Valid values are: `CODE_STAGE_START`, `ISSUE_CLOSED`, `ISSUE_CREATED`, `ISSUE_DEPLOYED_TO_PRODUCTION`, `ISSUE_FIRST_ADDED_TO_BOARD`, `ISSUE_FIRST_ADDED_TO_ITERATION`, `ISSUE_FIRST_ASSIGNED_AT`, `ISSUE_FIRST_ASSOCIATED_WITH_MILESTONE`, `ISSUE_FIRST_MENTIONED_IN_COMMIT`, `ISSUE_LABEL_ADDED`, `ISSUE_LABEL_REMOVED`, `ISSUE_LAST_EDITED`, `ISSUE_STAGE_END`, `MERGE_REQUEST_CLOSED`, `MERGE_REQUEST_CREATED`, `MERGE_REQUEST_FIRST_ASSIGNED_AT`, `MERGE_REQUEST_FIRST_COMMIT_AT`, `MERGE_REQUEST_FIRST_DEPLOYED_TO_PRODUCTION`, `MERGE_REQUEST_LABEL_ADDED`, `MERGE_REQUEST_LABEL_REMOVED`, `MERGE_REQUEST_LAST_BUILD_FINISHED`, `MERGE_REQUEST_LAST_BUILD_STARTED`, `MERGE_REQUEST_LAST_EDITED`, `MERGE_REQUEST_MERGED`, `MERGE_REQUEST_REVIEWER_FIRST_ASSIGNED`, `MERGE_REQUEST_PLAN_STAGE_START`
     * 
     */
    public Optional<String> startEventIdentifier() {
        return Optional.ofNullable(this.startEventIdentifier);
    }
    /**
     * @return Label ID associated with the start event identifier. In the format of `gid://gitlab/GroupLabel/&lt;id&gt;` or `gid://gitlab/ProjectLabel/&lt;id&gt;`
     * 
     */
    public Optional<String> startEventLabelId() {
        return Optional.ofNullable(this.startEventLabelId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ValueStreamAnalyticsStage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean custom;
        private @Nullable String endEventIdentifier;
        private @Nullable String endEventLabelId;
        private @Nullable Boolean hidden;
        private @Nullable String id;
        private String name;
        private @Nullable String startEventIdentifier;
        private @Nullable String startEventLabelId;
        public Builder() {}
        public Builder(ValueStreamAnalyticsStage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.custom = defaults.custom;
    	      this.endEventIdentifier = defaults.endEventIdentifier;
    	      this.endEventLabelId = defaults.endEventLabelId;
    	      this.hidden = defaults.hidden;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.startEventIdentifier = defaults.startEventIdentifier;
    	      this.startEventLabelId = defaults.startEventLabelId;
        }

        @CustomType.Setter
        public Builder custom(@Nullable Boolean custom) {

            this.custom = custom;
            return this;
        }
        @CustomType.Setter
        public Builder endEventIdentifier(@Nullable String endEventIdentifier) {

            this.endEventIdentifier = endEventIdentifier;
            return this;
        }
        @CustomType.Setter
        public Builder endEventLabelId(@Nullable String endEventLabelId) {

            this.endEventLabelId = endEventLabelId;
            return this;
        }
        @CustomType.Setter
        public Builder hidden(@Nullable Boolean hidden) {

            this.hidden = hidden;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("ValueStreamAnalyticsStage", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder startEventIdentifier(@Nullable String startEventIdentifier) {

            this.startEventIdentifier = startEventIdentifier;
            return this;
        }
        @CustomType.Setter
        public Builder startEventLabelId(@Nullable String startEventLabelId) {

            this.startEventLabelId = startEventLabelId;
            return this;
        }
        public ValueStreamAnalyticsStage build() {
            final var _resultValue = new ValueStreamAnalyticsStage();
            _resultValue.custom = custom;
            _resultValue.endEventIdentifier = endEventIdentifier;
            _resultValue.endEventLabelId = endEventLabelId;
            _resultValue.hidden = hidden;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.startEventIdentifier = startEventIdentifier;
            _resultValue.startEventLabelId = startEventLabelId;
            return _resultValue;
        }
    }
}
