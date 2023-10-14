// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectMilestonesMilestone {
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

    private GetProjectMilestonesMilestone() {}
    public String createdAt() {
        return this.createdAt;
    }
    public String description() {
        return this.description;
    }
    public String dueDate() {
        return this.dueDate;
    }
    public Boolean expired() {
        return this.expired;
    }
    public Integer iid() {
        return this.iid;
    }
    public Integer milestoneId() {
        return this.milestoneId;
    }
    public String project() {
        return this.project;
    }
    public Integer projectId() {
        return this.projectId;
    }
    public String startDate() {
        return this.startDate;
    }
    public String state() {
        return this.state;
    }
    public String title() {
        return this.title;
    }
    public String updatedAt() {
        return this.updatedAt;
    }
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
        public GetProjectMilestonesMilestone build() {
            final var o = new GetProjectMilestonesMilestone();
            o.createdAt = createdAt;
            o.description = description;
            o.dueDate = dueDate;
            o.expired = expired;
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
