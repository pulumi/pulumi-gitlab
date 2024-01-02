// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectsProject;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectsResult {
    /**
     * @return Limit by archived status.
     * 
     */
    private @Nullable Boolean archived;
    /**
     * @return The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
     * 
     */
    private @Nullable Integer groupId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
     * 
     */
    private @Nullable Boolean includeSubgroups;
    /**
     * @return The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
     * 
     */
    private @Nullable Integer maxQueryablePages;
    /**
     * @return Limit by projects that the current user is a member of.
     * 
     */
    private @Nullable Boolean membership;
    /**
     * @return Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
     * 
     */
    private @Nullable Integer minAccessLevel;
    /**
     * @return Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     * 
     */
    private @Nullable String orderBy;
    /**
     * @return Limit by projects owned by the current user.
     * 
     */
    private @Nullable Boolean owned;
    /**
     * @return The first page to begin the query on.
     * 
     */
    private @Nullable Integer page;
    /**
     * @return The number of results to return per page.
     * 
     */
    private @Nullable Integer perPage;
    /**
     * @return A list containing the projects matching the supplied arguments
     * 
     */
    private List<GetProjectsProject> projects;
    /**
     * @return Return list of authorized projects matching the search criteria.
     * 
     */
    private @Nullable String search;
    /**
     * @return Return only the ID, URL, name, and path of each project.
     * 
     */
    private @Nullable Boolean simple;
    /**
     * @return Return projects sorted in `asc` or `desc` order. Default is `desc`.
     * 
     */
    private @Nullable String sort;
    /**
     * @return Limit by projects starred by the current user.
     * 
     */
    private @Nullable Boolean starred;
    /**
     * @return Include project statistics. Cannot be used with `group_id`.
     * 
     */
    private @Nullable Boolean statistics;
    /**
     * @return Limit by projects that have all of the given topics.
     * 
     */
    private @Nullable List<String> topics;
    /**
     * @return Limit by visibility `public`, `internal`, or `private`.
     * 
     */
    private @Nullable String visibility;
    /**
     * @return Include custom attributes in response *(admins only)*.
     * 
     */
    private @Nullable Boolean withCustomAttributes;
    /**
     * @return Limit by projects with issues feature enabled. Default is `false`.
     * 
     */
    private @Nullable Boolean withIssuesEnabled;
    /**
     * @return Limit by projects with merge requests feature enabled. Default is `false`.
     * 
     */
    private @Nullable Boolean withMergeRequestsEnabled;
    /**
     * @return Limit by projects which use the given programming language. Cannot be used with `group_id`.
     * 
     */
    private @Nullable String withProgrammingLanguage;
    /**
     * @return Include projects shared to this group. Default is `true`. Needs `group_id`.
     * 
     */
    private @Nullable Boolean withShared;

    private GetProjectsResult() {}
    /**
     * @return Limit by archived status.
     * 
     */
    public Optional<Boolean> archived() {
        return Optional.ofNullable(this.archived);
    }
    /**
     * @return The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
     * 
     */
    public Optional<Integer> groupId() {
        return Optional.ofNullable(this.groupId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
     * 
     */
    public Optional<Boolean> includeSubgroups() {
        return Optional.ofNullable(this.includeSubgroups);
    }
    /**
     * @return The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
     * 
     */
    public Optional<Integer> maxQueryablePages() {
        return Optional.ofNullable(this.maxQueryablePages);
    }
    /**
     * @return Limit by projects that the current user is a member of.
     * 
     */
    public Optional<Boolean> membership() {
        return Optional.ofNullable(this.membership);
    }
    /**
     * @return Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
     * 
     */
    public Optional<Integer> minAccessLevel() {
        return Optional.ofNullable(this.minAccessLevel);
    }
    /**
     * @return Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     * 
     */
    public Optional<String> orderBy() {
        return Optional.ofNullable(this.orderBy);
    }
    /**
     * @return Limit by projects owned by the current user.
     * 
     */
    public Optional<Boolean> owned() {
        return Optional.ofNullable(this.owned);
    }
    /**
     * @return The first page to begin the query on.
     * 
     */
    public Optional<Integer> page() {
        return Optional.ofNullable(this.page);
    }
    /**
     * @return The number of results to return per page.
     * 
     */
    public Optional<Integer> perPage() {
        return Optional.ofNullable(this.perPage);
    }
    /**
     * @return A list containing the projects matching the supplied arguments
     * 
     */
    public List<GetProjectsProject> projects() {
        return this.projects;
    }
    /**
     * @return Return list of authorized projects matching the search criteria.
     * 
     */
    public Optional<String> search() {
        return Optional.ofNullable(this.search);
    }
    /**
     * @return Return only the ID, URL, name, and path of each project.
     * 
     */
    public Optional<Boolean> simple() {
        return Optional.ofNullable(this.simple);
    }
    /**
     * @return Return projects sorted in `asc` or `desc` order. Default is `desc`.
     * 
     */
    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }
    /**
     * @return Limit by projects starred by the current user.
     * 
     */
    public Optional<Boolean> starred() {
        return Optional.ofNullable(this.starred);
    }
    /**
     * @return Include project statistics. Cannot be used with `group_id`.
     * 
     */
    public Optional<Boolean> statistics() {
        return Optional.ofNullable(this.statistics);
    }
    /**
     * @return Limit by projects that have all of the given topics.
     * 
     */
    public List<String> topics() {
        return this.topics == null ? List.of() : this.topics;
    }
    /**
     * @return Limit by visibility `public`, `internal`, or `private`.
     * 
     */
    public Optional<String> visibility() {
        return Optional.ofNullable(this.visibility);
    }
    /**
     * @return Include custom attributes in response *(admins only)*.
     * 
     */
    public Optional<Boolean> withCustomAttributes() {
        return Optional.ofNullable(this.withCustomAttributes);
    }
    /**
     * @return Limit by projects with issues feature enabled. Default is `false`.
     * 
     */
    public Optional<Boolean> withIssuesEnabled() {
        return Optional.ofNullable(this.withIssuesEnabled);
    }
    /**
     * @return Limit by projects with merge requests feature enabled. Default is `false`.
     * 
     */
    public Optional<Boolean> withMergeRequestsEnabled() {
        return Optional.ofNullable(this.withMergeRequestsEnabled);
    }
    /**
     * @return Limit by projects which use the given programming language. Cannot be used with `group_id`.
     * 
     */
    public Optional<String> withProgrammingLanguage() {
        return Optional.ofNullable(this.withProgrammingLanguage);
    }
    /**
     * @return Include projects shared to this group. Default is `true`. Needs `group_id`.
     * 
     */
    public Optional<Boolean> withShared() {
        return Optional.ofNullable(this.withShared);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean archived;
        private @Nullable Integer groupId;
        private String id;
        private @Nullable Boolean includeSubgroups;
        private @Nullable Integer maxQueryablePages;
        private @Nullable Boolean membership;
        private @Nullable Integer minAccessLevel;
        private @Nullable String orderBy;
        private @Nullable Boolean owned;
        private @Nullable Integer page;
        private @Nullable Integer perPage;
        private List<GetProjectsProject> projects;
        private @Nullable String search;
        private @Nullable Boolean simple;
        private @Nullable String sort;
        private @Nullable Boolean starred;
        private @Nullable Boolean statistics;
        private @Nullable List<String> topics;
        private @Nullable String visibility;
        private @Nullable Boolean withCustomAttributes;
        private @Nullable Boolean withIssuesEnabled;
        private @Nullable Boolean withMergeRequestsEnabled;
        private @Nullable String withProgrammingLanguage;
        private @Nullable Boolean withShared;
        public Builder() {}
        public Builder(GetProjectsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.archived = defaults.archived;
    	      this.groupId = defaults.groupId;
    	      this.id = defaults.id;
    	      this.includeSubgroups = defaults.includeSubgroups;
    	      this.maxQueryablePages = defaults.maxQueryablePages;
    	      this.membership = defaults.membership;
    	      this.minAccessLevel = defaults.minAccessLevel;
    	      this.orderBy = defaults.orderBy;
    	      this.owned = defaults.owned;
    	      this.page = defaults.page;
    	      this.perPage = defaults.perPage;
    	      this.projects = defaults.projects;
    	      this.search = defaults.search;
    	      this.simple = defaults.simple;
    	      this.sort = defaults.sort;
    	      this.starred = defaults.starred;
    	      this.statistics = defaults.statistics;
    	      this.topics = defaults.topics;
    	      this.visibility = defaults.visibility;
    	      this.withCustomAttributes = defaults.withCustomAttributes;
    	      this.withIssuesEnabled = defaults.withIssuesEnabled;
    	      this.withMergeRequestsEnabled = defaults.withMergeRequestsEnabled;
    	      this.withProgrammingLanguage = defaults.withProgrammingLanguage;
    	      this.withShared = defaults.withShared;
        }

        @CustomType.Setter
        public Builder archived(@Nullable Boolean archived) {

            this.archived = archived;
            return this;
        }
        @CustomType.Setter
        public Builder groupId(@Nullable Integer groupId) {

            this.groupId = groupId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includeSubgroups(@Nullable Boolean includeSubgroups) {

            this.includeSubgroups = includeSubgroups;
            return this;
        }
        @CustomType.Setter
        public Builder maxQueryablePages(@Nullable Integer maxQueryablePages) {

            this.maxQueryablePages = maxQueryablePages;
            return this;
        }
        @CustomType.Setter
        public Builder membership(@Nullable Boolean membership) {

            this.membership = membership;
            return this;
        }
        @CustomType.Setter
        public Builder minAccessLevel(@Nullable Integer minAccessLevel) {

            this.minAccessLevel = minAccessLevel;
            return this;
        }
        @CustomType.Setter
        public Builder orderBy(@Nullable String orderBy) {

            this.orderBy = orderBy;
            return this;
        }
        @CustomType.Setter
        public Builder owned(@Nullable Boolean owned) {

            this.owned = owned;
            return this;
        }
        @CustomType.Setter
        public Builder page(@Nullable Integer page) {

            this.page = page;
            return this;
        }
        @CustomType.Setter
        public Builder perPage(@Nullable Integer perPage) {

            this.perPage = perPage;
            return this;
        }
        @CustomType.Setter
        public Builder projects(List<GetProjectsProject> projects) {
            if (projects == null) {
              throw new MissingRequiredPropertyException("GetProjectsResult", "projects");
            }
            this.projects = projects;
            return this;
        }
        public Builder projects(GetProjectsProject... projects) {
            return projects(List.of(projects));
        }
        @CustomType.Setter
        public Builder search(@Nullable String search) {

            this.search = search;
            return this;
        }
        @CustomType.Setter
        public Builder simple(@Nullable Boolean simple) {

            this.simple = simple;
            return this;
        }
        @CustomType.Setter
        public Builder sort(@Nullable String sort) {

            this.sort = sort;
            return this;
        }
        @CustomType.Setter
        public Builder starred(@Nullable Boolean starred) {

            this.starred = starred;
            return this;
        }
        @CustomType.Setter
        public Builder statistics(@Nullable Boolean statistics) {

            this.statistics = statistics;
            return this;
        }
        @CustomType.Setter
        public Builder topics(@Nullable List<String> topics) {

            this.topics = topics;
            return this;
        }
        public Builder topics(String... topics) {
            return topics(List.of(topics));
        }
        @CustomType.Setter
        public Builder visibility(@Nullable String visibility) {

            this.visibility = visibility;
            return this;
        }
        @CustomType.Setter
        public Builder withCustomAttributes(@Nullable Boolean withCustomAttributes) {

            this.withCustomAttributes = withCustomAttributes;
            return this;
        }
        @CustomType.Setter
        public Builder withIssuesEnabled(@Nullable Boolean withIssuesEnabled) {

            this.withIssuesEnabled = withIssuesEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder withMergeRequestsEnabled(@Nullable Boolean withMergeRequestsEnabled) {

            this.withMergeRequestsEnabled = withMergeRequestsEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder withProgrammingLanguage(@Nullable String withProgrammingLanguage) {

            this.withProgrammingLanguage = withProgrammingLanguage;
            return this;
        }
        @CustomType.Setter
        public Builder withShared(@Nullable Boolean withShared) {

            this.withShared = withShared;
            return this;
        }
        public GetProjectsResult build() {
            final var _resultValue = new GetProjectsResult();
            _resultValue.archived = archived;
            _resultValue.groupId = groupId;
            _resultValue.id = id;
            _resultValue.includeSubgroups = includeSubgroups;
            _resultValue.maxQueryablePages = maxQueryablePages;
            _resultValue.membership = membership;
            _resultValue.minAccessLevel = minAccessLevel;
            _resultValue.orderBy = orderBy;
            _resultValue.owned = owned;
            _resultValue.page = page;
            _resultValue.perPage = perPage;
            _resultValue.projects = projects;
            _resultValue.search = search;
            _resultValue.simple = simple;
            _resultValue.sort = sort;
            _resultValue.starred = starred;
            _resultValue.statistics = statistics;
            _resultValue.topics = topics;
            _resultValue.visibility = visibility;
            _resultValue.withCustomAttributes = withCustomAttributes;
            _resultValue.withIssuesEnabled = withIssuesEnabled;
            _resultValue.withMergeRequestsEnabled = withMergeRequestsEnabled;
            _resultValue.withProgrammingLanguage = withProgrammingLanguage;
            _resultValue.withShared = withShared;
            return _resultValue;
        }
    }
}
