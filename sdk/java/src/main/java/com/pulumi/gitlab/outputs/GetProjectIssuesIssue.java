// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetProjectIssuesIssueTaskCompletionStatus;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetProjectIssuesIssue {
    private List<Integer> assigneeIds;
    /**
     * @return Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
     * 
     */
    private Integer authorId;
    private String closedAt;
    private Integer closedByUserId;
    /**
     * @return Filter confidential or public issues.
     * 
     */
    private Boolean confidential;
    private String createdAt;
    private String description;
    private Boolean discussionLocked;
    private String discussionToResolve;
    private Integer downvotes;
    /**
     * @return Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
     * 
     */
    private String dueDate;
    private Integer epicId;
    private Integer epicIssueId;
    private String externalId;
    private String humanTimeEstimate;
    private String humanTotalTimeSpent;
    private Integer iid;
    private Integer issueId;
    private Integer issueLinkId;
    /**
     * @return Filter to a given type of issue. Valid values are [issue incident test_case]. (Introduced in GitLab 13.12)
     * 
     */
    private String issueType;
    /**
     * @return Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
     * 
     */
    private List<String> labels;
    private Map<String,String> links;
    private Integer mergeRequestToResolveDiscussionsOf;
    private Integer mergeRequestsCount;
    private Integer milestoneId;
    private Integer movedToId;
    /**
     * @return The name or id of the project.
     * 
     */
    private String project;
    private Map<String,String> references;
    private String state;
    private Boolean subscribed;
    private List<GetProjectIssuesIssueTaskCompletionStatus> taskCompletionStatuses;
    private Integer timeEstimate;
    private String title;
    private Integer totalTimeSpent;
    private String updatedAt;
    private Integer upvotes;
    private Integer userNotesCount;
    private String webUrl;
    /**
     * @return Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
     * 
     */
    private Integer weight;

    private GetProjectIssuesIssue() {}
    public List<Integer> assigneeIds() {
        return this.assigneeIds;
    }
    /**
     * @return Return issues created by the given user id. Combine with scope=all or scope=assigned*to*me.
     * 
     */
    public Integer authorId() {
        return this.authorId;
    }
    public String closedAt() {
        return this.closedAt;
    }
    public Integer closedByUserId() {
        return this.closedByUserId;
    }
    /**
     * @return Filter confidential or public issues.
     * 
     */
    public Boolean confidential() {
        return this.confidential;
    }
    public String createdAt() {
        return this.createdAt;
    }
    public String description() {
        return this.description;
    }
    public Boolean discussionLocked() {
        return this.discussionLocked;
    }
    public String discussionToResolve() {
        return this.discussionToResolve;
    }
    public Integer downvotes() {
        return this.downvotes;
    }
    /**
     * @return Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next*month*and*previous*two_weeks.
     * 
     */
    public String dueDate() {
        return this.dueDate;
    }
    public Integer epicId() {
        return this.epicId;
    }
    public Integer epicIssueId() {
        return this.epicIssueId;
    }
    public String externalId() {
        return this.externalId;
    }
    public String humanTimeEstimate() {
        return this.humanTimeEstimate;
    }
    public String humanTotalTimeSpent() {
        return this.humanTotalTimeSpent;
    }
    public Integer iid() {
        return this.iid;
    }
    public Integer issueId() {
        return this.issueId;
    }
    public Integer issueLinkId() {
        return this.issueLinkId;
    }
    /**
     * @return Filter to a given type of issue. Valid values are [issue incident test_case]. (Introduced in GitLab 13.12)
     * 
     */
    public String issueType() {
        return this.issueType;
    }
    /**
     * @return Return issues with labels. Issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
     * 
     */
    public List<String> labels() {
        return this.labels;
    }
    public Map<String,String> links() {
        return this.links;
    }
    public Integer mergeRequestToResolveDiscussionsOf() {
        return this.mergeRequestToResolveDiscussionsOf;
    }
    public Integer mergeRequestsCount() {
        return this.mergeRequestsCount;
    }
    public Integer milestoneId() {
        return this.milestoneId;
    }
    public Integer movedToId() {
        return this.movedToId;
    }
    /**
     * @return The name or id of the project.
     * 
     */
    public String project() {
        return this.project;
    }
    public Map<String,String> references() {
        return this.references;
    }
    public String state() {
        return this.state;
    }
    public Boolean subscribed() {
        return this.subscribed;
    }
    public List<GetProjectIssuesIssueTaskCompletionStatus> taskCompletionStatuses() {
        return this.taskCompletionStatuses;
    }
    public Integer timeEstimate() {
        return this.timeEstimate;
    }
    public String title() {
        return this.title;
    }
    public Integer totalTimeSpent() {
        return this.totalTimeSpent;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public Integer upvotes() {
        return this.upvotes;
    }
    public Integer userNotesCount() {
        return this.userNotesCount;
    }
    public String webUrl() {
        return this.webUrl;
    }
    /**
     * @return Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
     * 
     */
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectIssuesIssue defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Integer> assigneeIds;
        private Integer authorId;
        private String closedAt;
        private Integer closedByUserId;
        private Boolean confidential;
        private String createdAt;
        private String description;
        private Boolean discussionLocked;
        private String discussionToResolve;
        private Integer downvotes;
        private String dueDate;
        private Integer epicId;
        private Integer epicIssueId;
        private String externalId;
        private String humanTimeEstimate;
        private String humanTotalTimeSpent;
        private Integer iid;
        private Integer issueId;
        private Integer issueLinkId;
        private String issueType;
        private List<String> labels;
        private Map<String,String> links;
        private Integer mergeRequestToResolveDiscussionsOf;
        private Integer mergeRequestsCount;
        private Integer milestoneId;
        private Integer movedToId;
        private String project;
        private Map<String,String> references;
        private String state;
        private Boolean subscribed;
        private List<GetProjectIssuesIssueTaskCompletionStatus> taskCompletionStatuses;
        private Integer timeEstimate;
        private String title;
        private Integer totalTimeSpent;
        private String updatedAt;
        private Integer upvotes;
        private Integer userNotesCount;
        private String webUrl;
        private Integer weight;
        public Builder() {}
        public Builder(GetProjectIssuesIssue defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.assigneeIds = defaults.assigneeIds;
    	      this.authorId = defaults.authorId;
    	      this.closedAt = defaults.closedAt;
    	      this.closedByUserId = defaults.closedByUserId;
    	      this.confidential = defaults.confidential;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.discussionLocked = defaults.discussionLocked;
    	      this.discussionToResolve = defaults.discussionToResolve;
    	      this.downvotes = defaults.downvotes;
    	      this.dueDate = defaults.dueDate;
    	      this.epicId = defaults.epicId;
    	      this.epicIssueId = defaults.epicIssueId;
    	      this.externalId = defaults.externalId;
    	      this.humanTimeEstimate = defaults.humanTimeEstimate;
    	      this.humanTotalTimeSpent = defaults.humanTotalTimeSpent;
    	      this.iid = defaults.iid;
    	      this.issueId = defaults.issueId;
    	      this.issueLinkId = defaults.issueLinkId;
    	      this.issueType = defaults.issueType;
    	      this.labels = defaults.labels;
    	      this.links = defaults.links;
    	      this.mergeRequestToResolveDiscussionsOf = defaults.mergeRequestToResolveDiscussionsOf;
    	      this.mergeRequestsCount = defaults.mergeRequestsCount;
    	      this.milestoneId = defaults.milestoneId;
    	      this.movedToId = defaults.movedToId;
    	      this.project = defaults.project;
    	      this.references = defaults.references;
    	      this.state = defaults.state;
    	      this.subscribed = defaults.subscribed;
    	      this.taskCompletionStatuses = defaults.taskCompletionStatuses;
    	      this.timeEstimate = defaults.timeEstimate;
    	      this.title = defaults.title;
    	      this.totalTimeSpent = defaults.totalTimeSpent;
    	      this.updatedAt = defaults.updatedAt;
    	      this.upvotes = defaults.upvotes;
    	      this.userNotesCount = defaults.userNotesCount;
    	      this.webUrl = defaults.webUrl;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder assigneeIds(List<Integer> assigneeIds) {
            this.assigneeIds = Objects.requireNonNull(assigneeIds);
            return this;
        }
        public Builder assigneeIds(Integer... assigneeIds) {
            return assigneeIds(List.of(assigneeIds));
        }
        @CustomType.Setter
        public Builder authorId(Integer authorId) {
            this.authorId = Objects.requireNonNull(authorId);
            return this;
        }
        @CustomType.Setter
        public Builder closedAt(String closedAt) {
            this.closedAt = Objects.requireNonNull(closedAt);
            return this;
        }
        @CustomType.Setter
        public Builder closedByUserId(Integer closedByUserId) {
            this.closedByUserId = Objects.requireNonNull(closedByUserId);
            return this;
        }
        @CustomType.Setter
        public Builder confidential(Boolean confidential) {
            this.confidential = Objects.requireNonNull(confidential);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder discussionLocked(Boolean discussionLocked) {
            this.discussionLocked = Objects.requireNonNull(discussionLocked);
            return this;
        }
        @CustomType.Setter
        public Builder discussionToResolve(String discussionToResolve) {
            this.discussionToResolve = Objects.requireNonNull(discussionToResolve);
            return this;
        }
        @CustomType.Setter
        public Builder downvotes(Integer downvotes) {
            this.downvotes = Objects.requireNonNull(downvotes);
            return this;
        }
        @CustomType.Setter
        public Builder dueDate(String dueDate) {
            this.dueDate = Objects.requireNonNull(dueDate);
            return this;
        }
        @CustomType.Setter
        public Builder epicId(Integer epicId) {
            this.epicId = Objects.requireNonNull(epicId);
            return this;
        }
        @CustomType.Setter
        public Builder epicIssueId(Integer epicIssueId) {
            this.epicIssueId = Objects.requireNonNull(epicIssueId);
            return this;
        }
        @CustomType.Setter
        public Builder externalId(String externalId) {
            this.externalId = Objects.requireNonNull(externalId);
            return this;
        }
        @CustomType.Setter
        public Builder humanTimeEstimate(String humanTimeEstimate) {
            this.humanTimeEstimate = Objects.requireNonNull(humanTimeEstimate);
            return this;
        }
        @CustomType.Setter
        public Builder humanTotalTimeSpent(String humanTotalTimeSpent) {
            this.humanTotalTimeSpent = Objects.requireNonNull(humanTotalTimeSpent);
            return this;
        }
        @CustomType.Setter
        public Builder iid(Integer iid) {
            this.iid = Objects.requireNonNull(iid);
            return this;
        }
        @CustomType.Setter
        public Builder issueId(Integer issueId) {
            this.issueId = Objects.requireNonNull(issueId);
            return this;
        }
        @CustomType.Setter
        public Builder issueLinkId(Integer issueLinkId) {
            this.issueLinkId = Objects.requireNonNull(issueLinkId);
            return this;
        }
        @CustomType.Setter
        public Builder issueType(String issueType) {
            this.issueType = Objects.requireNonNull(issueType);
            return this;
        }
        @CustomType.Setter
        public Builder labels(List<String> labels) {
            this.labels = Objects.requireNonNull(labels);
            return this;
        }
        public Builder labels(String... labels) {
            return labels(List.of(labels));
        }
        @CustomType.Setter
        public Builder links(Map<String,String> links) {
            this.links = Objects.requireNonNull(links);
            return this;
        }
        @CustomType.Setter
        public Builder mergeRequestToResolveDiscussionsOf(Integer mergeRequestToResolveDiscussionsOf) {
            this.mergeRequestToResolveDiscussionsOf = Objects.requireNonNull(mergeRequestToResolveDiscussionsOf);
            return this;
        }
        @CustomType.Setter
        public Builder mergeRequestsCount(Integer mergeRequestsCount) {
            this.mergeRequestsCount = Objects.requireNonNull(mergeRequestsCount);
            return this;
        }
        @CustomType.Setter
        public Builder milestoneId(Integer milestoneId) {
            this.milestoneId = Objects.requireNonNull(milestoneId);
            return this;
        }
        @CustomType.Setter
        public Builder movedToId(Integer movedToId) {
            this.movedToId = Objects.requireNonNull(movedToId);
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        @CustomType.Setter
        public Builder references(Map<String,String> references) {
            this.references = Objects.requireNonNull(references);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder subscribed(Boolean subscribed) {
            this.subscribed = Objects.requireNonNull(subscribed);
            return this;
        }
        @CustomType.Setter
        public Builder taskCompletionStatuses(List<GetProjectIssuesIssueTaskCompletionStatus> taskCompletionStatuses) {
            this.taskCompletionStatuses = Objects.requireNonNull(taskCompletionStatuses);
            return this;
        }
        public Builder taskCompletionStatuses(GetProjectIssuesIssueTaskCompletionStatus... taskCompletionStatuses) {
            return taskCompletionStatuses(List.of(taskCompletionStatuses));
        }
        @CustomType.Setter
        public Builder timeEstimate(Integer timeEstimate) {
            this.timeEstimate = Objects.requireNonNull(timeEstimate);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        @CustomType.Setter
        public Builder totalTimeSpent(Integer totalTimeSpent) {
            this.totalTimeSpent = Objects.requireNonNull(totalTimeSpent);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder upvotes(Integer upvotes) {
            this.upvotes = Objects.requireNonNull(upvotes);
            return this;
        }
        @CustomType.Setter
        public Builder userNotesCount(Integer userNotesCount) {
            this.userNotesCount = Objects.requireNonNull(userNotesCount);
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            this.webUrl = Objects.requireNonNull(webUrl);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetProjectIssuesIssue build() {
            final var o = new GetProjectIssuesIssue();
            o.assigneeIds = assigneeIds;
            o.authorId = authorId;
            o.closedAt = closedAt;
            o.closedByUserId = closedByUserId;
            o.confidential = confidential;
            o.createdAt = createdAt;
            o.description = description;
            o.discussionLocked = discussionLocked;
            o.discussionToResolve = discussionToResolve;
            o.downvotes = downvotes;
            o.dueDate = dueDate;
            o.epicId = epicId;
            o.epicIssueId = epicIssueId;
            o.externalId = externalId;
            o.humanTimeEstimate = humanTimeEstimate;
            o.humanTotalTimeSpent = humanTotalTimeSpent;
            o.iid = iid;
            o.issueId = issueId;
            o.issueLinkId = issueLinkId;
            o.issueType = issueType;
            o.labels = labels;
            o.links = links;
            o.mergeRequestToResolveDiscussionsOf = mergeRequestToResolveDiscussionsOf;
            o.mergeRequestsCount = mergeRequestsCount;
            o.milestoneId = milestoneId;
            o.movedToId = movedToId;
            o.project = project;
            o.references = references;
            o.state = state;
            o.subscribed = subscribed;
            o.taskCompletionStatuses = taskCompletionStatuses;
            o.timeEstimate = timeEstimate;
            o.title = title;
            o.totalTimeSpent = totalTimeSpent;
            o.updatedAt = updatedAt;
            o.upvotes = upvotes;
            o.userNotesCount = userNotesCount;
            o.webUrl = webUrl;
            o.weight = weight;
            return o;
        }
    }
}