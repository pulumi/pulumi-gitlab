// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectMergeRequestAssignee;
import com.pulumi.gitlab.outputs.GetProjectMergeRequestAuthor;
import com.pulumi.gitlab.outputs.GetProjectMergeRequestClosedBy;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectMergeRequestResult {
    /**
     * @return First assignee of the merge request.
     * 
     */
    private GetProjectMergeRequestAssignee assignee;
    /**
     * @return Assignees of the merge request.
     * 
     */
    private List<GetProjectMergeRequestAssignee> assignees;
    /**
     * @return User who created this merge request.
     * 
     */
    private GetProjectMergeRequestAuthor author;
    /**
     * @return Indicates if all discussions are resolved only if all are
     * required before merge request can be merged.
     * 
     */
    private Boolean blockingDiscussionsResolved;
    /**
     * @return Number of changes made on the merge request. Empty when the
     * merge request is created, and populates asynchronously.
     * 
     */
    private String changesCount;
    /**
     * @return Timestamp of when the merge request was closed.
     * 
     */
    private String closedAt;
    /**
     * @return User who closed this merge request.
     * 
     */
    private GetProjectMergeRequestClosedBy closedBy;
    /**
     * @return Timestamp of when the merge request was created.
     * 
     */
    private String createdAt;
    /**
     * @return The unique instance level ID of the merge request.
     * 
     */
    private Integer id;
    /**
     * @return The unique project level ID of the merge request.
     * 
     */
    private Integer iid;
    /**
     * @return The ID or path of the project.
     * 
     */
    private String project;

    private GetProjectMergeRequestResult() {}
    /**
     * @return First assignee of the merge request.
     * 
     */
    public GetProjectMergeRequestAssignee assignee() {
        return this.assignee;
    }
    /**
     * @return Assignees of the merge request.
     * 
     */
    public List<GetProjectMergeRequestAssignee> assignees() {
        return this.assignees;
    }
    /**
     * @return User who created this merge request.
     * 
     */
    public GetProjectMergeRequestAuthor author() {
        return this.author;
    }
    /**
     * @return Indicates if all discussions are resolved only if all are
     * required before merge request can be merged.
     * 
     */
    public Boolean blockingDiscussionsResolved() {
        return this.blockingDiscussionsResolved;
    }
    /**
     * @return Number of changes made on the merge request. Empty when the
     * merge request is created, and populates asynchronously.
     * 
     */
    public String changesCount() {
        return this.changesCount;
    }
    /**
     * @return Timestamp of when the merge request was closed.
     * 
     */
    public String closedAt() {
        return this.closedAt;
    }
    /**
     * @return User who closed this merge request.
     * 
     */
    public GetProjectMergeRequestClosedBy closedBy() {
        return this.closedBy;
    }
    /**
     * @return Timestamp of when the merge request was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The unique instance level ID of the merge request.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The unique project level ID of the merge request.
     * 
     */
    public Integer iid() {
        return this.iid;
    }
    /**
     * @return The ID or path of the project.
     * 
     */
    public String project() {
        return this.project;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectMergeRequestResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private GetProjectMergeRequestAssignee assignee;
        private List<GetProjectMergeRequestAssignee> assignees;
        private GetProjectMergeRequestAuthor author;
        private Boolean blockingDiscussionsResolved;
        private String changesCount;
        private String closedAt;
        private GetProjectMergeRequestClosedBy closedBy;
        private String createdAt;
        private Integer id;
        private Integer iid;
        private String project;
        public Builder() {}
        public Builder(GetProjectMergeRequestResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.assignee = defaults.assignee;
    	      this.assignees = defaults.assignees;
    	      this.author = defaults.author;
    	      this.blockingDiscussionsResolved = defaults.blockingDiscussionsResolved;
    	      this.changesCount = defaults.changesCount;
    	      this.closedAt = defaults.closedAt;
    	      this.closedBy = defaults.closedBy;
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.iid = defaults.iid;
    	      this.project = defaults.project;
        }

        @CustomType.Setter
        public Builder assignee(GetProjectMergeRequestAssignee assignee) {
            if (assignee == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "assignee");
            }
            this.assignee = assignee;
            return this;
        }
        @CustomType.Setter
        public Builder assignees(List<GetProjectMergeRequestAssignee> assignees) {
            if (assignees == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "assignees");
            }
            this.assignees = assignees;
            return this;
        }
        public Builder assignees(GetProjectMergeRequestAssignee... assignees) {
            return assignees(List.of(assignees));
        }
        @CustomType.Setter
        public Builder author(GetProjectMergeRequestAuthor author) {
            if (author == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "author");
            }
            this.author = author;
            return this;
        }
        @CustomType.Setter
        public Builder blockingDiscussionsResolved(Boolean blockingDiscussionsResolved) {
            if (blockingDiscussionsResolved == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "blockingDiscussionsResolved");
            }
            this.blockingDiscussionsResolved = blockingDiscussionsResolved;
            return this;
        }
        @CustomType.Setter
        public Builder changesCount(String changesCount) {
            if (changesCount == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "changesCount");
            }
            this.changesCount = changesCount;
            return this;
        }
        @CustomType.Setter
        public Builder closedAt(String closedAt) {
            if (closedAt == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "closedAt");
            }
            this.closedAt = closedAt;
            return this;
        }
        @CustomType.Setter
        public Builder closedBy(GetProjectMergeRequestClosedBy closedBy) {
            if (closedBy == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "closedBy");
            }
            this.closedBy = closedBy;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder iid(Integer iid) {
            if (iid == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "iid");
            }
            this.iid = iid;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetProjectMergeRequestResult", "project");
            }
            this.project = project;
            return this;
        }
        public GetProjectMergeRequestResult build() {
            final var _resultValue = new GetProjectMergeRequestResult();
            _resultValue.assignee = assignee;
            _resultValue.assignees = assignees;
            _resultValue.author = author;
            _resultValue.blockingDiscussionsResolved = blockingDiscussionsResolved;
            _resultValue.changesCount = changesCount;
            _resultValue.closedAt = closedAt;
            _resultValue.closedBy = closedBy;
            _resultValue.createdAt = createdAt;
            _resultValue.id = id;
            _resultValue.iid = iid;
            _resultValue.project = project;
            return _resultValue;
        }
    }
}
