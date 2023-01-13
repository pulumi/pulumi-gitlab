// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectMilestoneResult {
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
     * @return Bool, true if milestore expired.
     * 
     */
    private Boolean expired;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
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

    private GetProjectMilestoneResult() {}
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
     * @return Bool, true if milestore expired.
     * 
     */
    public Boolean expired() {
        return this.expired;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
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

    public static Builder builder(GetProjectMilestoneResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String description;
        private String dueDate;
        private Boolean expired;
        private String id;
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
        public Builder(GetProjectMilestoneResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.dueDate = defaults.dueDate;
    	      this.expired = defaults.expired;
    	      this.id = defaults.id;
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
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder dueDate(String dueDate) {
            this.dueDate = Objects.requireNonNull(dueDate);
            return this;
        }
        @CustomType.Setter
        public Builder expired(Boolean expired) {
            this.expired = Objects.requireNonNull(expired);
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
        public Builder milestoneId(Integer milestoneId) {
            this.milestoneId = Objects.requireNonNull(milestoneId);
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(Integer projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder startDate(String startDate) {
            this.startDate = Objects.requireNonNull(startDate);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            this.webUrl = Objects.requireNonNull(webUrl);
            return this;
        }
        public GetProjectMilestoneResult build() {
            final var o = new GetProjectMilestoneResult();
            o.createdAt = createdAt;
            o.description = description;
            o.dueDate = dueDate;
            o.expired = expired;
            o.id = id;
            o.iid = iid;
            o.milestoneId = milestoneId;
            o.project = project;
            o.projectId = projectId;
            o.startDate = startDate;
            o.state = state;
            o.title = title;
            o.updatedAt = updatedAt;
            o.webUrl = webUrl;
            return o;
        }
    }
}