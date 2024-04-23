// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectIssuesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectIssuesPlainArgs Empty = new GetProjectIssuesPlainArgs();

    /**
     * Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
     * 
     */
    @Import(name="assigneeId")
    private @Nullable Integer assigneeId;

    /**
     * @return Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
     * 
     */
    public Optional<Integer> assigneeId() {
        return Optional.ofNullable(this.assigneeId);
    }

    /**
     * Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
     * 
     */
    @Import(name="assigneeUsername")
    private @Nullable String assigneeUsername;

    /**
     * @return Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
     * 
     */
    public Optional<String> assigneeUsername() {
        return Optional.ofNullable(this.assigneeUsername);
    }

    /**
     * Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
     * 
     */
    @Import(name="authorId")
    private @Nullable Integer authorId;

    /**
     * @return Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
     * 
     */
    public Optional<Integer> authorId() {
        return Optional.ofNullable(this.authorId);
    }

    /**
     * Filter confidential or public issues.
     * 
     */
    @Import(name="confidential")
    private @Nullable Boolean confidential;

    /**
     * @return Filter confidential or public issues.
     * 
     */
    public Optional<Boolean> confidential() {
        return Optional.ofNullable(this.confidential);
    }

    /**
     * Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    @Import(name="createdAfter")
    private @Nullable String createdAfter;

    /**
     * @return Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    public Optional<String> createdAfter() {
        return Optional.ofNullable(this.createdAfter);
    }

    /**
     * Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    @Import(name="createdBefore")
    private @Nullable String createdBefore;

    /**
     * @return Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    public Optional<String> createdBefore() {
        return Optional.ofNullable(this.createdBefore);
    }

    /**
     * Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
     * 
     */
    @Import(name="dueDate")
    private @Nullable String dueDate;

    /**
     * @return Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
     * 
     */
    public Optional<String> dueDate() {
        return Optional.ofNullable(this.dueDate);
    }

    /**
     * Return only the issues having the given iid
     * 
     */
    @Import(name="iids")
    private @Nullable List<Integer> iids;

    /**
     * @return Return only the issues having the given iid
     * 
     */
    public Optional<List<Integer>> iids() {
        return Optional.ofNullable(this.iids);
    }

    /**
     * Filter to a given type of issue. Valid values are [issue incident test_case]. (Introduced in GitLab 13.12)
     * 
     */
    @Import(name="issueType")
    private @Nullable String issueType;

    /**
     * @return Filter to a given type of issue. Valid values are [issue incident test_case]. (Introduced in GitLab 13.12)
     * 
     */
    public Optional<String> issueType() {
        return Optional.ofNullable(this.issueType);
    }

    /**
     * Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
     * 
     */
    @Import(name="labels")
    private @Nullable List<String> labels;

    /**
     * @return Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
     * 
     */
    public Optional<List<String>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
     * 
     */
    @Import(name="milestone")
    private @Nullable String milestone;

    /**
     * @return The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
     * 
     */
    public Optional<String> milestone() {
        return Optional.ofNullable(this.milestone);
    }

    /**
     * Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
     * 
     */
    @Import(name="myReactionEmoji")
    private @Nullable String myReactionEmoji;

    /**
     * @return Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
     * 
     */
    public Optional<String> myReactionEmoji() {
        return Optional.ofNullable(this.myReactionEmoji);
    }

    /**
     * Return issues that do not match the assignee id.
     * 
     */
    @Import(name="notAssigneeIds")
    private @Nullable List<Integer> notAssigneeIds;

    /**
     * @return Return issues that do not match the assignee id.
     * 
     */
    public Optional<List<Integer>> notAssigneeIds() {
        return Optional.ofNullable(this.notAssigneeIds);
    }

    /**
     * Return issues that do not match the author id.
     * 
     */
    @Import(name="notAuthorIds")
    private @Nullable List<Integer> notAuthorIds;

    /**
     * @return Return issues that do not match the author id.
     * 
     */
    public Optional<List<Integer>> notAuthorIds() {
        return Optional.ofNullable(this.notAuthorIds);
    }

    /**
     * Return issues that do not match the labels.
     * 
     */
    @Import(name="notLabels")
    private @Nullable List<String> notLabels;

    /**
     * @return Return issues that do not match the labels.
     * 
     */
    public Optional<List<String>> notLabels() {
        return Optional.ofNullable(this.notLabels);
    }

    /**
     * Return issues that do not match the milestone.
     * 
     */
    @Import(name="notMilestone")
    private @Nullable String notMilestone;

    /**
     * @return Return issues that do not match the milestone.
     * 
     */
    public Optional<String> notMilestone() {
        return Optional.ofNullable(this.notMilestone);
    }

    /**
     * Return issues not reacted by the authenticated user by the given emoji.
     * 
     */
    @Import(name="notMyReactionEmojis")
    private @Nullable List<String> notMyReactionEmojis;

    /**
     * @return Return issues not reacted by the authenticated user by the given emoji.
     * 
     */
    public Optional<List<String>> notMyReactionEmojis() {
        return Optional.ofNullable(this.notMyReactionEmojis);
    }

    /**
     * Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
     * 
     */
    @Import(name="orderBy")
    private @Nullable String orderBy;

    /**
     * @return Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
     * 
     */
    public Optional<String> orderBy() {
        return Optional.ofNullable(this.orderBy);
    }

    /**
     * The name or id of the project.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The name or id of the project.
     * 
     */
    public String project() {
        return this.project;
    }

    /**
     * Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
     * 
     */
    @Import(name="scope")
    private @Nullable String scope;

    /**
     * @return Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
     * 
     */
    public Optional<String> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * Search project issues against their title and description
     * 
     */
    @Import(name="search")
    private @Nullable String search;

    /**
     * @return Search project issues against their title and description
     * 
     */
    public Optional<String> search() {
        return Optional.ofNullable(this.search);
    }

    /**
     * Return issues sorted in asc or desc order. Default is desc
     * 
     */
    @Import(name="sort")
    private @Nullable String sort;

    /**
     * @return Return issues sorted in asc or desc order. Default is desc
     * 
     */
    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }

    /**
     * Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    @Import(name="updatedAfter")
    private @Nullable String updatedAfter;

    /**
     * @return Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    public Optional<String> updatedAfter() {
        return Optional.ofNullable(this.updatedAfter);
    }

    /**
     * Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    @Import(name="updatedBefore")
    private @Nullable String updatedBefore;

    /**
     * @return Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
     * 
     */
    public Optional<String> updatedBefore() {
        return Optional.ofNullable(this.updatedBefore);
    }

    /**
     * Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
     * 
     */
    @Import(name="weight")
    private @Nullable Integer weight;

    /**
     * @return Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
     * 
     */
    public Optional<Integer> weight() {
        return Optional.ofNullable(this.weight);
    }

    /**
     * If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false. description_html was introduced in GitLab 12.7
     * 
     */
    @Import(name="withLabelsDetails")
    private @Nullable Boolean withLabelsDetails;

    /**
     * @return If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false. description_html was introduced in GitLab 12.7
     * 
     */
    public Optional<Boolean> withLabelsDetails() {
        return Optional.ofNullable(this.withLabelsDetails);
    }

    private GetProjectIssuesPlainArgs() {}

    private GetProjectIssuesPlainArgs(GetProjectIssuesPlainArgs $) {
        this.assigneeId = $.assigneeId;
        this.assigneeUsername = $.assigneeUsername;
        this.authorId = $.authorId;
        this.confidential = $.confidential;
        this.createdAfter = $.createdAfter;
        this.createdBefore = $.createdBefore;
        this.dueDate = $.dueDate;
        this.iids = $.iids;
        this.issueType = $.issueType;
        this.labels = $.labels;
        this.milestone = $.milestone;
        this.myReactionEmoji = $.myReactionEmoji;
        this.notAssigneeIds = $.notAssigneeIds;
        this.notAuthorIds = $.notAuthorIds;
        this.notLabels = $.notLabels;
        this.notMilestone = $.notMilestone;
        this.notMyReactionEmojis = $.notMyReactionEmojis;
        this.orderBy = $.orderBy;
        this.project = $.project;
        this.scope = $.scope;
        this.search = $.search;
        this.sort = $.sort;
        this.updatedAfter = $.updatedAfter;
        this.updatedBefore = $.updatedBefore;
        this.weight = $.weight;
        this.withLabelsDetails = $.withLabelsDetails;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectIssuesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectIssuesPlainArgs $;

        public Builder() {
            $ = new GetProjectIssuesPlainArgs();
        }

        public Builder(GetProjectIssuesPlainArgs defaults) {
            $ = new GetProjectIssuesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assigneeId Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
         * 
         * @return builder
         * 
         */
        public Builder assigneeId(@Nullable Integer assigneeId) {
            $.assigneeId = assigneeId;
            return this;
        }

        /**
         * @param assigneeUsername Return issues assigned to the given username. Similar to assignee*id and mutually exclusive with assignee*id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
         * 
         * @return builder
         * 
         */
        public Builder assigneeUsername(@Nullable String assigneeUsername) {
            $.assigneeUsername = assigneeUsername;
            return this;
        }

        /**
         * @param authorId Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
         * 
         * @return builder
         * 
         */
        public Builder authorId(@Nullable Integer authorId) {
            $.authorId = authorId;
            return this;
        }

        /**
         * @param confidential Filter confidential or public issues.
         * 
         * @return builder
         * 
         */
        public Builder confidential(@Nullable Boolean confidential) {
            $.confidential = confidential;
            return this;
        }

        /**
         * @param createdAfter Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
         * 
         * @return builder
         * 
         */
        public Builder createdAfter(@Nullable String createdAfter) {
            $.createdAfter = createdAfter;
            return this;
        }

        /**
         * @param createdBefore Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
         * 
         * @return builder
         * 
         */
        public Builder createdBefore(@Nullable String createdBefore) {
            $.createdBefore = createdBefore;
            return this;
        }

        /**
         * @param dueDate Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
         * 
         * @return builder
         * 
         */
        public Builder dueDate(@Nullable String dueDate) {
            $.dueDate = dueDate;
            return this;
        }

        /**
         * @param iids Return only the issues having the given iid
         * 
         * @return builder
         * 
         */
        public Builder iids(@Nullable List<Integer> iids) {
            $.iids = iids;
            return this;
        }

        /**
         * @param iids Return only the issues having the given iid
         * 
         * @return builder
         * 
         */
        public Builder iids(Integer... iids) {
            return iids(List.of(iids));
        }

        /**
         * @param issueType Filter to a given type of issue. Valid values are [issue incident test_case]. (Introduced in GitLab 13.12)
         * 
         * @return builder
         * 
         */
        public Builder issueType(@Nullable String issueType) {
            $.issueType = issueType;
            return this;
        }

        /**
         * @param labels Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable List<String> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
         * 
         * @return builder
         * 
         */
        public Builder labels(String... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param milestone The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone.
         * 
         * @return builder
         * 
         */
        public Builder milestone(@Nullable String milestone) {
            $.milestone = milestone;
            return this;
        }

        /**
         * @param myReactionEmoji Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
         * 
         * @return builder
         * 
         */
        public Builder myReactionEmoji(@Nullable String myReactionEmoji) {
            $.myReactionEmoji = myReactionEmoji;
            return this;
        }

        /**
         * @param notAssigneeIds Return issues that do not match the assignee id.
         * 
         * @return builder
         * 
         */
        public Builder notAssigneeIds(@Nullable List<Integer> notAssigneeIds) {
            $.notAssigneeIds = notAssigneeIds;
            return this;
        }

        /**
         * @param notAssigneeIds Return issues that do not match the assignee id.
         * 
         * @return builder
         * 
         */
        public Builder notAssigneeIds(Integer... notAssigneeIds) {
            return notAssigneeIds(List.of(notAssigneeIds));
        }

        /**
         * @param notAuthorIds Return issues that do not match the author id.
         * 
         * @return builder
         * 
         */
        public Builder notAuthorIds(@Nullable List<Integer> notAuthorIds) {
            $.notAuthorIds = notAuthorIds;
            return this;
        }

        /**
         * @param notAuthorIds Return issues that do not match the author id.
         * 
         * @return builder
         * 
         */
        public Builder notAuthorIds(Integer... notAuthorIds) {
            return notAuthorIds(List.of(notAuthorIds));
        }

        /**
         * @param notLabels Return issues that do not match the labels.
         * 
         * @return builder
         * 
         */
        public Builder notLabels(@Nullable List<String> notLabels) {
            $.notLabels = notLabels;
            return this;
        }

        /**
         * @param notLabels Return issues that do not match the labels.
         * 
         * @return builder
         * 
         */
        public Builder notLabels(String... notLabels) {
            return notLabels(List.of(notLabels));
        }

        /**
         * @param notMilestone Return issues that do not match the milestone.
         * 
         * @return builder
         * 
         */
        public Builder notMilestone(@Nullable String notMilestone) {
            $.notMilestone = notMilestone;
            return this;
        }

        /**
         * @param notMyReactionEmojis Return issues not reacted by the authenticated user by the given emoji.
         * 
         * @return builder
         * 
         */
        public Builder notMyReactionEmojis(@Nullable List<String> notMyReactionEmojis) {
            $.notMyReactionEmojis = notMyReactionEmojis;
            return this;
        }

        /**
         * @param notMyReactionEmojis Return issues not reacted by the authenticated user by the given emoji.
         * 
         * @return builder
         * 
         */
        public Builder notMyReactionEmojis(String... notMyReactionEmojis) {
            return notMyReactionEmojis(List.of(notMyReactionEmojis));
        }

        /**
         * @param orderBy Return issues ordered by. Valid values are `created_at`, `updated_at`, `priority`, `due_date`, `relative_position`, `label_priority`, `milestone_due`, `popularity`, `weight`. Default is created_at
         * 
         * @return builder
         * 
         */
        public Builder orderBy(@Nullable String orderBy) {
            $.orderBy = orderBy;
            return this;
        }

        /**
         * @param project The name or id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        /**
         * @param scope Return issues for the given scope. Valid values are `created_by_me`, `assigned_to_me`, `all`. Defaults to all.
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable String scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param search Search project issues against their title and description
         * 
         * @return builder
         * 
         */
        public Builder search(@Nullable String search) {
            $.search = search;
            return this;
        }

        /**
         * @param sort Return issues sorted in asc or desc order. Default is desc
         * 
         * @return builder
         * 
         */
        public Builder sort(@Nullable String sort) {
            $.sort = sort;
            return this;
        }

        /**
         * @param updatedAfter Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
         * 
         * @return builder
         * 
         */
        public Builder updatedAfter(@Nullable String updatedAfter) {
            $.updatedAfter = updatedAfter;
            return this;
        }

        /**
         * @param updatedBefore Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)
         * 
         * @return builder
         * 
         */
        public Builder updatedBefore(@Nullable String updatedBefore) {
            $.updatedBefore = updatedBefore;
            return this;
        }

        /**
         * @param weight Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Integer weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param withLabelsDetails If true, the response returns more details for each label in labels field: :name, :color, :description, :description*html, :text*color. Default is false. description_html was introduced in GitLab 12.7
         * 
         * @return builder
         * 
         */
        public Builder withLabelsDetails(@Nullable Boolean withLabelsDetails) {
            $.withLabelsDetails = withLabelsDetails;
            return this;
        }

        public GetProjectIssuesPlainArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectIssuesPlainArgs", "project");
            }
            return $;
        }
    }

}
