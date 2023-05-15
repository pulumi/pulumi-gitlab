// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetProjectIssueTaskCompletionStatus;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetProjectIssueResult {
    /**
     * @return The IDs of the users to assign the issue to.
     * 
     */
    private List<Integer> assigneeIds;
    /**
     * @return The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    private Integer authorId;
    /**
     * @return When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    private String closedAt;
    /**
     * @return The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    private Integer closedByUserId;
    /**
     * @return Set an issue to be confidential.
     * 
     */
    private Boolean confidential;
    /**
     * @return When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     * 
     */
    private String createdAt;
    /**
     * @return The description of an issue. Limited to 1,048,576 characters.
     * 
     */
    private String description;
    /**
     * @return Whether the issue is locked for discussions or not.
     * 
     */
    private Boolean discussionLocked;
    /**
     * @return The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     * 
     */
    private String discussionToResolve;
    /**
     * @return The number of downvotes the issue has received.
     * 
     */
    private Integer downvotes;
    /**
     * @return The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    private String dueDate;
    /**
     * @return ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     * 
     */
    private Integer epicId;
    /**
     * @return The ID of the epic issue.
     * 
     */
    private Integer epicIssueId;
    /**
     * @return The external ID of the issue.
     * 
     */
    private String externalId;
    /**
     * @return The human-readable time estimate of the issue.
     * 
     */
    private String humanTimeEstimate;
    /**
     * @return The human-readable total time spent of the issue.
     * 
     */
    private String humanTotalTimeSpent;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The internal ID of the project&#39;s issue.
     * 
     */
    private Integer iid;
    /**
     * @return The instance-wide ID of the issue.
     * 
     */
    private Integer issueId;
    /**
     * @return The ID of the issue link.
     * 
     */
    private Integer issueLinkId;
    /**
     * @return The type of issue. Valid values are: `issue`, `incident`, `test_case`.
     * 
     */
    private String issueType;
    /**
     * @return The labels of an issue.
     * 
     */
    private List<String> labels;
    /**
     * @return The links of the issue.
     * 
     */
    private Map<String,String> links;
    /**
     * @return The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     * 
     */
    private Integer mergeRequestToResolveDiscussionsOf;
    /**
     * @return The number of merge requests associated with the issue.
     * 
     */
    private Integer mergeRequestsCount;
    /**
     * @return The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue&#39;s details.
     * 
     */
    private Integer milestoneId;
    /**
     * @return The ID of the issue that was moved to.
     * 
     */
    private Integer movedToId;
    /**
     * @return The name or ID of the project.
     * 
     */
    private String project;
    /**
     * @return The references of the issue.
     * 
     */
    private Map<String,String> references;
    /**
     * @return The state of the issue. Valid values are: `opened`, `closed`.
     * 
     */
    private String state;
    /**
     * @return Whether the authenticated user is subscribed to the issue or not.
     * 
     */
    private Boolean subscribed;
    /**
     * @return The task completion status. It&#39;s always a one element list.
     * 
     */
    private List<GetProjectIssueTaskCompletionStatus> taskCompletionStatuses;
    /**
     * @return The time estimate of the issue.
     * 
     */
    private Integer timeEstimate;
    /**
     * @return The title of the issue.
     * 
     */
    private String title;
    /**
     * @return The total time spent of the issue.
     * 
     */
    private Integer totalTimeSpent;
    /**
     * @return When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    private String updatedAt;
    /**
     * @return The number of upvotes the issue has received.
     * 
     */
    private Integer upvotes;
    /**
     * @return The number of user notes on the issue.
     * 
     */
    private Integer userNotesCount;
    /**
     * @return The web URL of the issue.
     * 
     */
    private String webUrl;
    /**
     * @return The weight of the issue. Valid values are greater than or equal to 0.
     * 
     */
    private Integer weight;

    private GetProjectIssueResult() {}
    /**
     * @return The IDs of the users to assign the issue to.
     * 
     */
    public List<Integer> assigneeIds() {
        return this.assigneeIds;
    }
    /**
     * @return The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    public Integer authorId() {
        return this.authorId;
    }
    /**
     * @return When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public String closedAt() {
        return this.closedAt;
    }
    /**
     * @return The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    public Integer closedByUserId() {
        return this.closedByUserId;
    }
    /**
     * @return Set an issue to be confidential.
     * 
     */
    public Boolean confidential() {
        return this.confidential;
    }
    /**
     * @return When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The description of an issue. Limited to 1,048,576 characters.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Whether the issue is locked for discussions or not.
     * 
     */
    public Boolean discussionLocked() {
        return this.discussionLocked;
    }
    /**
     * @return The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     * 
     */
    public String discussionToResolve() {
        return this.discussionToResolve;
    }
    /**
     * @return The number of downvotes the issue has received.
     * 
     */
    public Integer downvotes() {
        return this.downvotes;
    }
    /**
     * @return The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public String dueDate() {
        return this.dueDate;
    }
    /**
     * @return ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     * 
     */
    public Integer epicId() {
        return this.epicId;
    }
    /**
     * @return The ID of the epic issue.
     * 
     */
    public Integer epicIssueId() {
        return this.epicIssueId;
    }
    /**
     * @return The external ID of the issue.
     * 
     */
    public String externalId() {
        return this.externalId;
    }
    /**
     * @return The human-readable time estimate of the issue.
     * 
     */
    public String humanTimeEstimate() {
        return this.humanTimeEstimate;
    }
    /**
     * @return The human-readable total time spent of the issue.
     * 
     */
    public String humanTotalTimeSpent() {
        return this.humanTotalTimeSpent;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The internal ID of the project&#39;s issue.
     * 
     */
    public Integer iid() {
        return this.iid;
    }
    /**
     * @return The instance-wide ID of the issue.
     * 
     */
    public Integer issueId() {
        return this.issueId;
    }
    /**
     * @return The ID of the issue link.
     * 
     */
    public Integer issueLinkId() {
        return this.issueLinkId;
    }
    /**
     * @return The type of issue. Valid values are: `issue`, `incident`, `test_case`.
     * 
     */
    public String issueType() {
        return this.issueType;
    }
    /**
     * @return The labels of an issue.
     * 
     */
    public List<String> labels() {
        return this.labels;
    }
    /**
     * @return The links of the issue.
     * 
     */
    public Map<String,String> links() {
        return this.links;
    }
    /**
     * @return The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     * 
     */
    public Integer mergeRequestToResolveDiscussionsOf() {
        return this.mergeRequestToResolveDiscussionsOf;
    }
    /**
     * @return The number of merge requests associated with the issue.
     * 
     */
    public Integer mergeRequestsCount() {
        return this.mergeRequestsCount;
    }
    /**
     * @return The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue&#39;s details.
     * 
     */
    public Integer milestoneId() {
        return this.milestoneId;
    }
    /**
     * @return The ID of the issue that was moved to.
     * 
     */
    public Integer movedToId() {
        return this.movedToId;
    }
    /**
     * @return The name or ID of the project.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return The references of the issue.
     * 
     */
    public Map<String,String> references() {
        return this.references;
    }
    /**
     * @return The state of the issue. Valid values are: `opened`, `closed`.
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return Whether the authenticated user is subscribed to the issue or not.
     * 
     */
    public Boolean subscribed() {
        return this.subscribed;
    }
    /**
     * @return The task completion status. It&#39;s always a one element list.
     * 
     */
    public List<GetProjectIssueTaskCompletionStatus> taskCompletionStatuses() {
        return this.taskCompletionStatuses;
    }
    /**
     * @return The time estimate of the issue.
     * 
     */
    public Integer timeEstimate() {
        return this.timeEstimate;
    }
    /**
     * @return The title of the issue.
     * 
     */
    public String title() {
        return this.title;
    }
    /**
     * @return The total time spent of the issue.
     * 
     */
    public Integer totalTimeSpent() {
        return this.totalTimeSpent;
    }
    /**
     * @return When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The number of upvotes the issue has received.
     * 
     */
    public Integer upvotes() {
        return this.upvotes;
    }
    /**
     * @return The number of user notes on the issue.
     * 
     */
    public Integer userNotesCount() {
        return this.userNotesCount;
    }
    /**
     * @return The web URL of the issue.
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }
    /**
     * @return The weight of the issue. Valid values are greater than or equal to 0.
     * 
     */
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectIssueResult defaults) {
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
        private String id;
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
        private List<GetProjectIssueTaskCompletionStatus> taskCompletionStatuses;
        private Integer timeEstimate;
        private String title;
        private Integer totalTimeSpent;
        private String updatedAt;
        private Integer upvotes;
        private Integer userNotesCount;
        private String webUrl;
        private Integer weight;
        public Builder() {}
        public Builder(GetProjectIssueResult defaults) {
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
    	      this.id = defaults.id;
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
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
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
        public Builder taskCompletionStatuses(List<GetProjectIssueTaskCompletionStatus> taskCompletionStatuses) {
            this.taskCompletionStatuses = Objects.requireNonNull(taskCompletionStatuses);
            return this;
        }
        public Builder taskCompletionStatuses(GetProjectIssueTaskCompletionStatus... taskCompletionStatuses) {
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
        public GetProjectIssueResult build() {
            final var o = new GetProjectIssueResult();
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
            o.id = id;
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
