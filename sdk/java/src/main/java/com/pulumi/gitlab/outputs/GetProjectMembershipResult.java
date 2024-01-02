// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectMembershipMember;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectMembershipResult {
    /**
     * @return The full path of the project.
     * 
     */
    private String fullPath;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Return all project members including members through ancestor groups
     * 
     */
    private @Nullable Boolean inherited;
    /**
     * @return The list of project members.
     * 
     */
    private List<GetProjectMembershipMember> members;
    /**
     * @return The ID of the project.
     * 
     */
    private Integer projectId;
    /**
     * @return A query string to search for members
     * 
     */
    private @Nullable String query;

    private GetProjectMembershipResult() {}
    /**
     * @return The full path of the project.
     * 
     */
    public String fullPath() {
        return this.fullPath;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Return all project members including members through ancestor groups
     * 
     */
    public Optional<Boolean> inherited() {
        return Optional.ofNullable(this.inherited);
    }
    /**
     * @return The list of project members.
     * 
     */
    public List<GetProjectMembershipMember> members() {
        return this.members;
    }
    /**
     * @return The ID of the project.
     * 
     */
    public Integer projectId() {
        return this.projectId;
    }
    /**
     * @return A query string to search for members
     * 
     */
    public Optional<String> query() {
        return Optional.ofNullable(this.query);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectMembershipResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fullPath;
        private String id;
        private @Nullable Boolean inherited;
        private List<GetProjectMembershipMember> members;
        private Integer projectId;
        private @Nullable String query;
        public Builder() {}
        public Builder(GetProjectMembershipResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullPath = defaults.fullPath;
    	      this.id = defaults.id;
    	      this.inherited = defaults.inherited;
    	      this.members = defaults.members;
    	      this.projectId = defaults.projectId;
    	      this.query = defaults.query;
        }

        @CustomType.Setter
        public Builder fullPath(String fullPath) {
            if (fullPath == null) {
              throw new MissingRequiredPropertyException("GetProjectMembershipResult", "fullPath");
            }
            this.fullPath = fullPath;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectMembershipResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder inherited(@Nullable Boolean inherited) {

            this.inherited = inherited;
            return this;
        }
        @CustomType.Setter
        public Builder members(List<GetProjectMembershipMember> members) {
            if (members == null) {
              throw new MissingRequiredPropertyException("GetProjectMembershipResult", "members");
            }
            this.members = members;
            return this;
        }
        public Builder members(GetProjectMembershipMember... members) {
            return members(List.of(members));
        }
        @CustomType.Setter
        public Builder projectId(Integer projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetProjectMembershipResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder query(@Nullable String query) {

            this.query = query;
            return this;
        }
        public GetProjectMembershipResult build() {
            final var _resultValue = new GetProjectMembershipResult();
            _resultValue.fullPath = fullPath;
            _resultValue.id = id;
            _resultValue.inherited = inherited;
            _resultValue.members = members;
            _resultValue.projectId = projectId;
            _resultValue.query = query;
            return _resultValue;
        }
    }
}
