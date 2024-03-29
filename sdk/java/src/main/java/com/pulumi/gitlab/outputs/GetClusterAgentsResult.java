// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetClusterAgentsClusterAgent;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetClusterAgentsResult {
    /**
     * @return List of the registered agents.
     * 
     */
    private List<GetClusterAgentsClusterAgent> clusterAgents;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The ID or full path of the project owned by the authenticated user.
     * 
     */
    private String project;

    private GetClusterAgentsResult() {}
    /**
     * @return List of the registered agents.
     * 
     */
    public List<GetClusterAgentsClusterAgent> clusterAgents() {
        return this.clusterAgents;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID or full path of the project owned by the authenticated user.
     * 
     */
    public String project() {
        return this.project;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClusterAgentsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetClusterAgentsClusterAgent> clusterAgents;
        private String id;
        private String project;
        public Builder() {}
        public Builder(GetClusterAgentsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterAgents = defaults.clusterAgents;
    	      this.id = defaults.id;
    	      this.project = defaults.project;
        }

        @CustomType.Setter
        public Builder clusterAgents(List<GetClusterAgentsClusterAgent> clusterAgents) {
            if (clusterAgents == null) {
              throw new MissingRequiredPropertyException("GetClusterAgentsResult", "clusterAgents");
            }
            this.clusterAgents = clusterAgents;
            return this;
        }
        public Builder clusterAgents(GetClusterAgentsClusterAgent... clusterAgents) {
            return clusterAgents(List.of(clusterAgents));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetClusterAgentsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            if (project == null) {
              throw new MissingRequiredPropertyException("GetClusterAgentsResult", "project");
            }
            this.project = project;
            return this;
        }
        public GetClusterAgentsResult build() {
            final var _resultValue = new GetClusterAgentsResult();
            _resultValue.clusterAgents = clusterAgents;
            _resultValue.id = id;
            _resultValue.project = project;
            return _resultValue;
        }
    }
}
