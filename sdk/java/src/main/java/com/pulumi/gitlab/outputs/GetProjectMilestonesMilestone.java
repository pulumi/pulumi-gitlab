// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectMilestonesMilestone {
    /**
     * @return The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    private String createdAt;
    /**
     * @return The description of the milestone.
     * 
     */
    private String description;
    /**
     * @return The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    private String dueDate;
    /**
     * @return Bool, true if milestone expired.
     * 
     */
    private Boolean expired;
    /**
     * @return The ID of the project&#39;s milestone.
     * 
     */
    private Integer iid;
    /**
     * @return The instance-wide ID of the project’s milestone.
     * 
     */
    private Integer milestoneId;
    /**
     * @return The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    private String project;
    /**
     * @return The project ID of milestone.
     * 
     */
    private Integer projectId;
    /**
     * @return The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    private String startDate;
    /**
     * @return The state of the milestone. Valid values are: `active`, `closed`.
     * 
     */
    private String state;
    /**
     * @return The title of a milestone.
     * 
     */
    private String title;
    /**
     * @return The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    private String updatedAt;
    /**
     * @return The web URL of the milestone.
     * 
     */
    private String webUrl;

    private GetProjectMilestonesMilestone() {}
    /**
     * @return The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The description of the milestone.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public String dueDate() {
        return this.dueDate;
    }
    /**
     * @return Bool, true if milestone expired.
     * 
     */
    public Boolean expired() {
        return this.expired;
    }
    /**
     * @return The ID of the project&#39;s milestone.
     * 
     */
    public Integer iid() {
        return this.iid;
    }
    /**
     * @return The instance-wide ID of the project’s milestone.
     * 
     */
    public Integer milestoneId() {
        return this.milestoneId;
    }
    /**
     * @return The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return The project ID of milestone.
     * 
     */
    public Integer projectId() {
        return this.projectId;
    }
    /**
     * @return The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public String startDate() {
        return this.startDate;
    }
    /**
     * @return The state of the milestone. Valid values are: `active`, `closed`.
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return The title of a milestone.
     * 
     */
    public String title() {
        return this.title;
    }
    /**
     * @return The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The web URL of the milestone.
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectMilestonesMilestone defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String description;
        private String dueDate;
        private Boolean expired;
        private Integer iid;
        private Integer milestoneId;
        private String project;
        private Integer projectId;
        private String startDate;
        private String state;
        private String title;
        private String updatedAt;
        private String webUrl;
        public Builder() {}
        public Builder(GetProjectMilestonesMilestone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.dueDate = defaults.dueDate;
    	      this.expired = defaults.expired;
    	      this.iid = defaults.iid;
    	      this.milestoneId = defaults.milestoneId;
    	      this.project = defaults.project;
    	      this.projectId = defaults.projectId;
    	      this.startDate = defaults.startDate;
    	      this.state = defaults.state;
    	      this.title = defaults.title;
    	      this.updatedAt = defaults.updatedAt;
    	      this.webUrl = defaults.webUrl;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder dueDate(String dueDate) {
            if (dueDate == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "dueDate");
            }
            this.dueDate = dueDate;
            return this;
        }
        @CustomType.Setter
        public Builder expired(Boolean expired) {
            if (expired == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "expired");
            }
            this.expired = expired;
            return this;
        }
        @CustomType.Setter
        public Builder iid(Integer iid) {
            if (iid == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "iid");
            }
            this.iid = iid;
            return this;
        }
        @CustomType.Setter
        public Builder milestoneId(Integer milestoneId) {
            if (milestoneId == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "milestoneId");
            }
            this.milestoneId = milestoneId;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "project");
            }
            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(Integer projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder startDate(String startDate) {
            if (startDate == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "startDate");
            }
            this.startDate = startDate;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            if (title == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "title");
            }
            this.title = title;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            if (webUrl == null) {
              throw new MissingRequiredPropertyException("GetProjectMilestonesMilestone", "webUrl");
            }
            this.webUrl = webUrl;
            return this;
        }
        public GetProjectMilestonesMilestone build() {
            final var _resultValue = new GetProjectMilestonesMilestone();
            _resultValue.createdAt = createdAt;
            _resultValue.description = description;
            _resultValue.dueDate = dueDate;
            _resultValue.expired = expired;
            _resultValue.iid = iid;
            _resultValue.milestoneId = milestoneId;
            _resultValue.project = project;
            _resultValue.projectId = projectId;
            _resultValue.startDate = startDate;
            _resultValue.state = state;
            _resultValue.title = title;
            _resultValue.updatedAt = updatedAt;
            _resultValue.webUrl = webUrl;
            return _resultValue;
        }
    }
}
