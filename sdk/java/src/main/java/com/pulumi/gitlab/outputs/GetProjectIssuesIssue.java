// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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

    private GetProjectIssuesIssue() {}
    public List<Integer> assigneeIds() {
        return this.assigneeIds;
    }
    public Integer authorId() {
        return this.authorId;
    }
    public String closedAt() {
        return this.closedAt;
    }
    public Integer closedByUserId() {
        return this.closedByUserId;
    }
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
    public String issueType() {
        return this.issueType;
    }
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
            if (assigneeIds == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "assigneeIds");
            }
            this.assigneeIds = assigneeIds;
            return this;
        }
        public Builder assigneeIds(Integer... assigneeIds) {
            return assigneeIds(List.of(assigneeIds));
        }
        @CustomType.Setter
        public Builder authorId(Integer authorId) {
            if (authorId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "authorId");
            }
            this.authorId = authorId;
            return this;
        }
        @CustomType.Setter
        public Builder closedAt(String closedAt) {
            if (closedAt == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "closedAt");
            }
            this.closedAt = closedAt;
            return this;
        }
        @CustomType.Setter
        public Builder closedByUserId(Integer closedByUserId) {
            if (closedByUserId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "closedByUserId");
            }
            this.closedByUserId = closedByUserId;
            return this;
        }
        @CustomType.Setter
        public Builder confidential(Boolean confidential) {
            if (confidential == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "confidential");
            }
            this.confidential = confidential;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder discussionLocked(Boolean discussionLocked) {
            if (discussionLocked == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "discussionLocked");
            }
            this.discussionLocked = discussionLocked;
            return this;
        }
        @CustomType.Setter
        public Builder discussionToResolve(String discussionToResolve) {
            if (discussionToResolve == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "discussionToResolve");
            }
            this.discussionToResolve = discussionToResolve;
            return this;
        }
        @CustomType.Setter
        public Builder downvotes(Integer downvotes) {
            if (downvotes == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "downvotes");
            }
            this.downvotes = downvotes;
            return this;
        }
        @CustomType.Setter
        public Builder dueDate(String dueDate) {
            if (dueDate == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "dueDate");
            }
            this.dueDate = dueDate;
            return this;
        }
        @CustomType.Setter
        public Builder epicId(Integer epicId) {
            if (epicId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "epicId");
            }
            this.epicId = epicId;
            return this;
        }
        @CustomType.Setter
        public Builder epicIssueId(Integer epicIssueId) {
            if (epicIssueId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "epicIssueId");
            }
            this.epicIssueId = epicIssueId;
            return this;
        }
        @CustomType.Setter
        public Builder externalId(String externalId) {
            if (externalId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "externalId");
            }
            this.externalId = externalId;
            return this;
        }
        @CustomType.Setter
        public Builder humanTimeEstimate(String humanTimeEstimate) {
            if (humanTimeEstimate == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "humanTimeEstimate");
            }
            this.humanTimeEstimate = humanTimeEstimate;
            return this;
        }
        @CustomType.Setter
        public Builder humanTotalTimeSpent(String humanTotalTimeSpent) {
            if (humanTotalTimeSpent == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "humanTotalTimeSpent");
            }
            this.humanTotalTimeSpent = humanTotalTimeSpent;
            return this;
        }
        @CustomType.Setter
        public Builder iid(Integer iid) {
            if (iid == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "iid");
            }
            this.iid = iid;
            return this;
        }
        @CustomType.Setter
        public Builder issueId(Integer issueId) {
            if (issueId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "issueId");
            }
            this.issueId = issueId;
            return this;
        }
        @CustomType.Setter
        public Builder issueLinkId(Integer issueLinkId) {
            if (issueLinkId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "issueLinkId");
            }
            this.issueLinkId = issueLinkId;
            return this;
        }
        @CustomType.Setter
        public Builder issueType(String issueType) {
            if (issueType == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "issueType");
            }
            this.issueType = issueType;
            return this;
        }
        @CustomType.Setter
        public Builder labels(List<String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "labels");
            }
            this.labels = labels;
            return this;
        }
        public Builder labels(String... labels) {
            return labels(List.of(labels));
        }
        @CustomType.Setter
        public Builder links(Map<String,String> links) {
            if (links == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "links");
            }
            this.links = links;
            return this;
        }
        @CustomType.Setter
        public Builder mergeRequestToResolveDiscussionsOf(Integer mergeRequestToResolveDiscussionsOf) {
            if (mergeRequestToResolveDiscussionsOf == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "mergeRequestToResolveDiscussionsOf");
            }
            this.mergeRequestToResolveDiscussionsOf = mergeRequestToResolveDiscussionsOf;
            return this;
        }
        @CustomType.Setter
        public Builder mergeRequestsCount(Integer mergeRequestsCount) {
            if (mergeRequestsCount == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "mergeRequestsCount");
            }
            this.mergeRequestsCount = mergeRequestsCount;
            return this;
        }
        @CustomType.Setter
        public Builder milestoneId(Integer milestoneId) {
            if (milestoneId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "milestoneId");
            }
            this.milestoneId = milestoneId;
            return this;
        }
        @CustomType.Setter
        public Builder movedToId(Integer movedToId) {
            if (movedToId == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "movedToId");
            }
            this.movedToId = movedToId;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "project");
            }
            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder references(Map<String,String> references) {
            if (references == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "references");
            }
            this.references = references;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder subscribed(Boolean subscribed) {
            if (subscribed == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "subscribed");
            }
            this.subscribed = subscribed;
            return this;
        }
        @CustomType.Setter
        public Builder taskCompletionStatuses(List<GetProjectIssuesIssueTaskCompletionStatus> taskCompletionStatuses) {
            if (taskCompletionStatuses == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "taskCompletionStatuses");
            }
            this.taskCompletionStatuses = taskCompletionStatuses;
            return this;
        }
        public Builder taskCompletionStatuses(GetProjectIssuesIssueTaskCompletionStatus... taskCompletionStatuses) {
            return taskCompletionStatuses(List.of(taskCompletionStatuses));
        }
        @CustomType.Setter
        public Builder timeEstimate(Integer timeEstimate) {
            if (timeEstimate == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "timeEstimate");
            }
            this.timeEstimate = timeEstimate;
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            if (title == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "title");
            }
            this.title = title;
            return this;
        }
        @CustomType.Setter
        public Builder totalTimeSpent(Integer totalTimeSpent) {
            if (totalTimeSpent == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "totalTimeSpent");
            }
            this.totalTimeSpent = totalTimeSpent;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder upvotes(Integer upvotes) {
            if (upvotes == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "upvotes");
            }
            this.upvotes = upvotes;
            return this;
        }
        @CustomType.Setter
        public Builder userNotesCount(Integer userNotesCount) {
            if (userNotesCount == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "userNotesCount");
            }
            this.userNotesCount = userNotesCount;
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            if (webUrl == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "webUrl");
            }
            this.webUrl = webUrl;
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            if (weight == null) {
              throw new MissingRequiredPropertyException("GetProjectIssuesIssue", "weight");
            }
            this.weight = weight;
            return this;
        }
        public GetProjectIssuesIssue build() {
            final var _resultValue = new GetProjectIssuesIssue();
            _resultValue.assigneeIds = assigneeIds;
            _resultValue.authorId = authorId;
            _resultValue.closedAt = closedAt;
            _resultValue.closedByUserId = closedByUserId;
            _resultValue.confidential = confidential;
            _resultValue.createdAt = createdAt;
            _resultValue.description = description;
            _resultValue.discussionLocked = discussionLocked;
            _resultValue.discussionToResolve = discussionToResolve;
            _resultValue.downvotes = downvotes;
            _resultValue.dueDate = dueDate;
            _resultValue.epicId = epicId;
            _resultValue.epicIssueId = epicIssueId;
            _resultValue.externalId = externalId;
            _resultValue.humanTimeEstimate = humanTimeEstimate;
            _resultValue.humanTotalTimeSpent = humanTotalTimeSpent;
            _resultValue.iid = iid;
            _resultValue.issueId = issueId;
            _resultValue.issueLinkId = issueLinkId;
            _resultValue.issueType = issueType;
            _resultValue.labels = labels;
            _resultValue.links = links;
            _resultValue.mergeRequestToResolveDiscussionsOf = mergeRequestToResolveDiscussionsOf;
            _resultValue.mergeRequestsCount = mergeRequestsCount;
            _resultValue.milestoneId = milestoneId;
            _resultValue.movedToId = movedToId;
            _resultValue.project = project;
            _resultValue.references = references;
            _resultValue.state = state;
            _resultValue.subscribed = subscribed;
            _resultValue.taskCompletionStatuses = taskCompletionStatuses;
            _resultValue.timeEstimate = timeEstimate;
            _resultValue.title = title;
            _resultValue.totalTimeSpent = totalTimeSpent;
            _resultValue.updatedAt = updatedAt;
            _resultValue.upvotes = upvotes;
            _resultValue.userNotesCount = userNotesCount;
            _resultValue.webUrl = webUrl;
            _resultValue.weight = weight;
            return _resultValue;
        }
    }
}
