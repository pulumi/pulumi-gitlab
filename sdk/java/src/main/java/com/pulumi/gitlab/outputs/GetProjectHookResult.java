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
public final class GetProjectHookResult {
    /**
     * @return Invoke the hook for confidential issues events.
     * 
     */
    private Boolean confidentialIssuesEvents;
    /**
     * @return Invoke the hook for confidential notes events.
     * 
     */
    private Boolean confidentialNoteEvents;
    /**
     * @return Invoke the hook for deployment events.
     * 
     */
    private Boolean deploymentEvents;
    /**
     * @return Enable ssl verification when invoking the hook.
     * 
     */
    private Boolean enableSslVerification;
    /**
     * @return The id of the project hook.
     * 
     */
    private Integer hookId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Invoke the hook for issues events.
     * 
     */
    private Boolean issuesEvents;
    /**
     * @return Invoke the hook for job events.
     * 
     */
    private Boolean jobEvents;
    /**
     * @return Invoke the hook for merge requests.
     * 
     */
    private Boolean mergeRequestsEvents;
    /**
     * @return Invoke the hook for notes events.
     * 
     */
    private Boolean noteEvents;
    /**
     * @return Invoke the hook for pipeline events.
     * 
     */
    private Boolean pipelineEvents;
    /**
     * @return The name or id of the project to add the hook to.
     * 
     */
    private String project;
    /**
     * @return The id of the project for the hook.
     * 
     */
    private Integer projectId;
    /**
     * @return Invoke the hook for push events.
     * 
     */
    private Boolean pushEvents;
    /**
     * @return Invoke the hook for push events on matching branches only.
     * 
     */
    private String pushEventsBranchFilter;
    /**
     * @return Invoke the hook for releases events.
     * 
     */
    private Boolean releasesEvents;
    /**
     * @return Invoke the hook for tag push events.
     * 
     */
    private Boolean tagPushEvents;
    /**
     * @return A token to present when invoking the hook. The token is not available for imported resources.
     * 
     */
    private String token;
    /**
     * @return The url of the hook to invoke.
     * 
     */
    private String url;
    /**
     * @return Invoke the hook for wiki page events.
     * 
     */
    private Boolean wikiPageEvents;

    private GetProjectHookResult() {}
    /**
     * @return Invoke the hook for confidential issues events.
     * 
     */
    public Boolean confidentialIssuesEvents() {
        return this.confidentialIssuesEvents;
    }
    /**
     * @return Invoke the hook for confidential notes events.
     * 
     */
    public Boolean confidentialNoteEvents() {
        return this.confidentialNoteEvents;
    }
    /**
     * @return Invoke the hook for deployment events.
     * 
     */
    public Boolean deploymentEvents() {
        return this.deploymentEvents;
    }
    /**
     * @return Enable ssl verification when invoking the hook.
     * 
     */
    public Boolean enableSslVerification() {
        return this.enableSslVerification;
    }
    /**
     * @return The id of the project hook.
     * 
     */
    public Integer hookId() {
        return this.hookId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Invoke the hook for issues events.
     * 
     */
    public Boolean issuesEvents() {
        return this.issuesEvents;
    }
    /**
     * @return Invoke the hook for job events.
     * 
     */
    public Boolean jobEvents() {
        return this.jobEvents;
    }
    /**
     * @return Invoke the hook for merge requests.
     * 
     */
    public Boolean mergeRequestsEvents() {
        return this.mergeRequestsEvents;
    }
    /**
     * @return Invoke the hook for notes events.
     * 
     */
    public Boolean noteEvents() {
        return this.noteEvents;
    }
    /**
     * @return Invoke the hook for pipeline events.
     * 
     */
    public Boolean pipelineEvents() {
        return this.pipelineEvents;
    }
    /**
     * @return The name or id of the project to add the hook to.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return The id of the project for the hook.
     * 
     */
    public Integer projectId() {
        return this.projectId;
    }
    /**
     * @return Invoke the hook for push events.
     * 
     */
    public Boolean pushEvents() {
        return this.pushEvents;
    }
    /**
     * @return Invoke the hook for push events on matching branches only.
     * 
     */
    public String pushEventsBranchFilter() {
        return this.pushEventsBranchFilter;
    }
    /**
     * @return Invoke the hook for releases events.
     * 
     */
    public Boolean releasesEvents() {
        return this.releasesEvents;
    }
    /**
     * @return Invoke the hook for tag push events.
     * 
     */
    public Boolean tagPushEvents() {
        return this.tagPushEvents;
    }
    /**
     * @return A token to present when invoking the hook. The token is not available for imported resources.
     * 
     */
    public String token() {
        return this.token;
    }
    /**
     * @return The url of the hook to invoke.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Invoke the hook for wiki page events.
     * 
     */
    public Boolean wikiPageEvents() {
        return this.wikiPageEvents;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectHookResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean confidentialIssuesEvents;
        private Boolean confidentialNoteEvents;
        private Boolean deploymentEvents;
        private Boolean enableSslVerification;
        private Integer hookId;
        private String id;
        private Boolean issuesEvents;
        private Boolean jobEvents;
        private Boolean mergeRequestsEvents;
        private Boolean noteEvents;
        private Boolean pipelineEvents;
        private String project;
        private Integer projectId;
        private Boolean pushEvents;
        private String pushEventsBranchFilter;
        private Boolean releasesEvents;
        private Boolean tagPushEvents;
        private String token;
        private String url;
        private Boolean wikiPageEvents;
        public Builder() {}
        public Builder(GetProjectHookResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.confidentialIssuesEvents = defaults.confidentialIssuesEvents;
    	      this.confidentialNoteEvents = defaults.confidentialNoteEvents;
    	      this.deploymentEvents = defaults.deploymentEvents;
    	      this.enableSslVerification = defaults.enableSslVerification;
    	      this.hookId = defaults.hookId;
    	      this.id = defaults.id;
    	      this.issuesEvents = defaults.issuesEvents;
    	      this.jobEvents = defaults.jobEvents;
    	      this.mergeRequestsEvents = defaults.mergeRequestsEvents;
    	      this.noteEvents = defaults.noteEvents;
    	      this.pipelineEvents = defaults.pipelineEvents;
    	      this.project = defaults.project;
    	      this.projectId = defaults.projectId;
    	      this.pushEvents = defaults.pushEvents;
    	      this.pushEventsBranchFilter = defaults.pushEventsBranchFilter;
    	      this.releasesEvents = defaults.releasesEvents;
    	      this.tagPushEvents = defaults.tagPushEvents;
    	      this.token = defaults.token;
    	      this.url = defaults.url;
    	      this.wikiPageEvents = defaults.wikiPageEvents;
        }

        @CustomType.Setter
        public Builder confidentialIssuesEvents(Boolean confidentialIssuesEvents) {
            if (confidentialIssuesEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "confidentialIssuesEvents");
            }
            this.confidentialIssuesEvents = confidentialIssuesEvents;
            return this;
        }
        @CustomType.Setter
        public Builder confidentialNoteEvents(Boolean confidentialNoteEvents) {
            if (confidentialNoteEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "confidentialNoteEvents");
            }
            this.confidentialNoteEvents = confidentialNoteEvents;
            return this;
        }
        @CustomType.Setter
        public Builder deploymentEvents(Boolean deploymentEvents) {
            if (deploymentEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "deploymentEvents");
            }
            this.deploymentEvents = deploymentEvents;
            return this;
        }
        @CustomType.Setter
        public Builder enableSslVerification(Boolean enableSslVerification) {
            if (enableSslVerification == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "enableSslVerification");
            }
            this.enableSslVerification = enableSslVerification;
            return this;
        }
        @CustomType.Setter
        public Builder hookId(Integer hookId) {
            if (hookId == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "hookId");
            }
            this.hookId = hookId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder issuesEvents(Boolean issuesEvents) {
            if (issuesEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "issuesEvents");
            }
            this.issuesEvents = issuesEvents;
            return this;
        }
        @CustomType.Setter
        public Builder jobEvents(Boolean jobEvents) {
            if (jobEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "jobEvents");
            }
            this.jobEvents = jobEvents;
            return this;
        }
        @CustomType.Setter
        public Builder mergeRequestsEvents(Boolean mergeRequestsEvents) {
            if (mergeRequestsEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "mergeRequestsEvents");
            }
            this.mergeRequestsEvents = mergeRequestsEvents;
            return this;
        }
        @CustomType.Setter
        public Builder noteEvents(Boolean noteEvents) {
            if (noteEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "noteEvents");
            }
            this.noteEvents = noteEvents;
            return this;
        }
        @CustomType.Setter
        public Builder pipelineEvents(Boolean pipelineEvents) {
            if (pipelineEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "pipelineEvents");
            }
            this.pipelineEvents = pipelineEvents;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "project");
            }
            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(Integer projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder pushEvents(Boolean pushEvents) {
            if (pushEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "pushEvents");
            }
            this.pushEvents = pushEvents;
            return this;
        }
        @CustomType.Setter
        public Builder pushEventsBranchFilter(String pushEventsBranchFilter) {
            if (pushEventsBranchFilter == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "pushEventsBranchFilter");
            }
            this.pushEventsBranchFilter = pushEventsBranchFilter;
            return this;
        }
        @CustomType.Setter
        public Builder releasesEvents(Boolean releasesEvents) {
            if (releasesEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "releasesEvents");
            }
            this.releasesEvents = releasesEvents;
            return this;
        }
        @CustomType.Setter
        public Builder tagPushEvents(Boolean tagPushEvents) {
            if (tagPushEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "tagPushEvents");
            }
            this.tagPushEvents = tagPushEvents;
            return this;
        }
        @CustomType.Setter
        public Builder token(String token) {
            if (token == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "token");
            }
            this.token = token;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder wikiPageEvents(Boolean wikiPageEvents) {
            if (wikiPageEvents == null) {
              throw new MissingRequiredPropertyException("GetProjectHookResult", "wikiPageEvents");
            }
            this.wikiPageEvents = wikiPageEvents;
            return this;
        }
        public GetProjectHookResult build() {
            final var _resultValue = new GetProjectHookResult();
            _resultValue.confidentialIssuesEvents = confidentialIssuesEvents;
            _resultValue.confidentialNoteEvents = confidentialNoteEvents;
            _resultValue.deploymentEvents = deploymentEvents;
            _resultValue.enableSslVerification = enableSslVerification;
            _resultValue.hookId = hookId;
            _resultValue.id = id;
            _resultValue.issuesEvents = issuesEvents;
            _resultValue.jobEvents = jobEvents;
            _resultValue.mergeRequestsEvents = mergeRequestsEvents;
            _resultValue.noteEvents = noteEvents;
            _resultValue.pipelineEvents = pipelineEvents;
            _resultValue.project = project;
            _resultValue.projectId = projectId;
            _resultValue.pushEvents = pushEvents;
            _resultValue.pushEventsBranchFilter = pushEventsBranchFilter;
            _resultValue.releasesEvents = releasesEvents;
            _resultValue.tagPushEvents = tagPushEvents;
            _resultValue.token = token;
            _resultValue.url = url;
            _resultValue.wikiPageEvents = wikiPageEvents;
            return _resultValue;
        }
    }
}
