// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectsArgs Empty = new GetProjectsArgs();

    @Import(name="archived")
    private @Nullable Output<Boolean> archived;

    public Optional<Output<Boolean>> archived() {
        return Optional.ofNullable(this.archived);
    }

    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
     * 
     */
    @Import(name="includeSubgroups")
    private @Nullable Output<Boolean> includeSubgroups;

    /**
     * @return Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
     * 
     */
    public Optional<Output<Boolean>> includeSubgroups() {
        return Optional.ofNullable(this.includeSubgroups);
    }

    /**
     * The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
     * 
     */
    @Import(name="maxQueryablePages")
    private @Nullable Output<Integer> maxQueryablePages;

    /**
     * @return The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
     * 
     */
    public Optional<Output<Integer>> maxQueryablePages() {
        return Optional.ofNullable(this.maxQueryablePages);
    }

    /**
     * Limit by projects that the current user is a member of.
     * 
     */
    @Import(name="membership")
    private @Nullable Output<Boolean> membership;

    /**
     * @return Limit by projects that the current user is a member of.
     * 
     */
    public Optional<Output<Boolean>> membership() {
        return Optional.ofNullable(this.membership);
    }

    /**
     * Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
     * 
     */
    @Import(name="minAccessLevel")
    private @Nullable Output<Integer> minAccessLevel;

    /**
     * @return Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
     * 
     */
    public Optional<Output<Integer>> minAccessLevel() {
        return Optional.ofNullable(this.minAccessLevel);
    }

    /**
     * Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     * 
     */
    @Import(name="orderBy")
    private @Nullable Output<String> orderBy;

    /**
     * @return Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
     * 
     */
    public Optional<Output<String>> orderBy() {
        return Optional.ofNullable(this.orderBy);
    }

    /**
     * Limit by projects owned by the current user.
     * 
     */
    @Import(name="owned")
    private @Nullable Output<Boolean> owned;

    /**
     * @return Limit by projects owned by the current user.
     * 
     */
    public Optional<Output<Boolean>> owned() {
        return Optional.ofNullable(this.owned);
    }

    /**
     * The first page to begin the query on.
     * 
     */
    @Import(name="page")
    private @Nullable Output<Integer> page;

    /**
     * @return The first page to begin the query on.
     * 
     */
    public Optional<Output<Integer>> page() {
        return Optional.ofNullable(this.page);
    }

    /**
     * The number of results to return per page.
     * 
     */
    @Import(name="perPage")
    private @Nullable Output<Integer> perPage;

    /**
     * @return The number of results to return per page.
     * 
     */
    public Optional<Output<Integer>> perPage() {
        return Optional.ofNullable(this.perPage);
    }

    /**
     * Return list of authorized projects matching the search criteria.
     * 
     */
    @Import(name="search")
    private @Nullable Output<String> search;

    /**
     * @return Return list of authorized projects matching the search criteria.
     * 
     */
    public Optional<Output<String>> search() {
        return Optional.ofNullable(this.search);
    }

    /**
     * Return only the ID, URL, name, and path of each project.
     * 
     */
    @Import(name="simple")
    private @Nullable Output<Boolean> simple;

    /**
     * @return Return only the ID, URL, name, and path of each project.
     * 
     */
    public Optional<Output<Boolean>> simple() {
        return Optional.ofNullable(this.simple);
    }

    /**
     * Return projects sorted in `asc` or `desc` order. Default is `desc`.
     * 
     */
    @Import(name="sort")
    private @Nullable Output<String> sort;

    /**
     * @return Return projects sorted in `asc` or `desc` order. Default is `desc`.
     * 
     */
    public Optional<Output<String>> sort() {
        return Optional.ofNullable(this.sort);
    }

    /**
     * Limit by projects starred by the current user.
     * 
     */
    @Import(name="starred")
    private @Nullable Output<Boolean> starred;

    /**
     * @return Limit by projects starred by the current user.
     * 
     */
    public Optional<Output<Boolean>> starred() {
        return Optional.ofNullable(this.starred);
    }

    @Import(name="statistics")
    private @Nullable Output<Boolean> statistics;

    public Optional<Output<Boolean>> statistics() {
        return Optional.ofNullable(this.statistics);
    }

    /**
     * Limit by projects that have all of the given topics.
     * 
     */
    @Import(name="topics")
    private @Nullable Output<List<String>> topics;

    /**
     * @return Limit by projects that have all of the given topics.
     * 
     */
    public Optional<Output<List<String>>> topics() {
        return Optional.ofNullable(this.topics);
    }

    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    /**
     * Include custom attributes in response *(admins only)*.
     * 
     */
    @Import(name="withCustomAttributes")
    private @Nullable Output<Boolean> withCustomAttributes;

    /**
     * @return Include custom attributes in response *(admins only)*.
     * 
     */
    public Optional<Output<Boolean>> withCustomAttributes() {
        return Optional.ofNullable(this.withCustomAttributes);
    }

    /**
     * Limit by projects with issues feature enabled. Default is `false`.
     * 
     */
    @Import(name="withIssuesEnabled")
    private @Nullable Output<Boolean> withIssuesEnabled;

    /**
     * @return Limit by projects with issues feature enabled. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> withIssuesEnabled() {
        return Optional.ofNullable(this.withIssuesEnabled);
    }

    /**
     * Limit by projects with merge requests feature enabled. Default is `false`.
     * 
     */
    @Import(name="withMergeRequestsEnabled")
    private @Nullable Output<Boolean> withMergeRequestsEnabled;

    /**
     * @return Limit by projects with merge requests feature enabled. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> withMergeRequestsEnabled() {
        return Optional.ofNullable(this.withMergeRequestsEnabled);
    }

    /**
     * Limit by projects which use the given programming language. Cannot be used with `group_id`.
     * 
     */
    @Import(name="withProgrammingLanguage")
    private @Nullable Output<String> withProgrammingLanguage;

    /**
     * @return Limit by projects which use the given programming language. Cannot be used with `group_id`.
     * 
     */
    public Optional<Output<String>> withProgrammingLanguage() {
        return Optional.ofNullable(this.withProgrammingLanguage);
    }

    /**
     * Include projects shared to this group. Default is `true`. Needs `group_id`.
     * 
     */
    @Import(name="withShared")
    private @Nullable Output<Boolean> withShared;

    /**
     * @return Include projects shared to this group. Default is `true`. Needs `group_id`.
     * 
     */
    public Optional<Output<Boolean>> withShared() {
        return Optional.ofNullable(this.withShared);
    }

    private GetProjectsArgs() {}

    private GetProjectsArgs(GetProjectsArgs $) {
        this.archived = $.archived;
        this.groupId = $.groupId;
        this.includeSubgroups = $.includeSubgroups;
        this.maxQueryablePages = $.maxQueryablePages;
        this.membership = $.membership;
        this.minAccessLevel = $.minAccessLevel;
        this.orderBy = $.orderBy;
        this.owned = $.owned;
        this.page = $.page;
        this.perPage = $.perPage;
        this.search = $.search;
        this.simple = $.simple;
        this.sort = $.sort;
        this.starred = $.starred;
        this.statistics = $.statistics;
        this.topics = $.topics;
        this.visibility = $.visibility;
        this.withCustomAttributes = $.withCustomAttributes;
        this.withIssuesEnabled = $.withIssuesEnabled;
        this.withMergeRequestsEnabled = $.withMergeRequestsEnabled;
        this.withProgrammingLanguage = $.withProgrammingLanguage;
        this.withShared = $.withShared;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectsArgs $;

        public Builder() {
            $ = new GetProjectsArgs();
        }

        public Builder(GetProjectsArgs defaults) {
            $ = new GetProjectsArgs(Objects.requireNonNull(defaults));
        }

        public Builder archived(@Nullable Output<Boolean> archived) {
            $.archived = archived;
            return this;
        }

        public Builder archived(Boolean archived) {
            return archived(Output.of(archived));
        }

        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param includeSubgroups Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder includeSubgroups(@Nullable Output<Boolean> includeSubgroups) {
            $.includeSubgroups = includeSubgroups;
            return this;
        }

        /**
         * @param includeSubgroups Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder includeSubgroups(Boolean includeSubgroups) {
            return includeSubgroups(Output.of(includeSubgroups));
        }

        /**
         * @param maxQueryablePages The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
         * 
         * @return builder
         * 
         */
        public Builder maxQueryablePages(@Nullable Output<Integer> maxQueryablePages) {
            $.maxQueryablePages = maxQueryablePages;
            return this;
        }

        /**
         * @param maxQueryablePages The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
         * 
         * @return builder
         * 
         */
        public Builder maxQueryablePages(Integer maxQueryablePages) {
            return maxQueryablePages(Output.of(maxQueryablePages));
        }

        /**
         * @param membership Limit by projects that the current user is a member of.
         * 
         * @return builder
         * 
         */
        public Builder membership(@Nullable Output<Boolean> membership) {
            $.membership = membership;
            return this;
        }

        /**
         * @param membership Limit by projects that the current user is a member of.
         * 
         * @return builder
         * 
         */
        public Builder membership(Boolean membership) {
            return membership(Output.of(membership));
        }

        /**
         * @param minAccessLevel Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder minAccessLevel(@Nullable Output<Integer> minAccessLevel) {
            $.minAccessLevel = minAccessLevel;
            return this;
        }

        /**
         * @param minAccessLevel Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder minAccessLevel(Integer minAccessLevel) {
            return minAccessLevel(Output.of(minAccessLevel));
        }

        /**
         * @param orderBy Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
         * 
         * @return builder
         * 
         */
        public Builder orderBy(@Nullable Output<String> orderBy) {
            $.orderBy = orderBy;
            return this;
        }

        /**
         * @param orderBy Return projects ordered ordered by: `id`, `name`, `path`, `created_at`, `updated_at`, `last_activity_at`, `similarity`, `repository_size`, `storage_size`, `packages_size`, `wiki_size`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects) for details.
         * 
         * @return builder
         * 
         */
        public Builder orderBy(String orderBy) {
            return orderBy(Output.of(orderBy));
        }

        /**
         * @param owned Limit by projects owned by the current user.
         * 
         * @return builder
         * 
         */
        public Builder owned(@Nullable Output<Boolean> owned) {
            $.owned = owned;
            return this;
        }

        /**
         * @param owned Limit by projects owned by the current user.
         * 
         * @return builder
         * 
         */
        public Builder owned(Boolean owned) {
            return owned(Output.of(owned));
        }

        /**
         * @param page The first page to begin the query on.
         * 
         * @return builder
         * 
         */
        public Builder page(@Nullable Output<Integer> page) {
            $.page = page;
            return this;
        }

        /**
         * @param page The first page to begin the query on.
         * 
         * @return builder
         * 
         */
        public Builder page(Integer page) {
            return page(Output.of(page));
        }

        /**
         * @param perPage The number of results to return per page.
         * 
         * @return builder
         * 
         */
        public Builder perPage(@Nullable Output<Integer> perPage) {
            $.perPage = perPage;
            return this;
        }

        /**
         * @param perPage The number of results to return per page.
         * 
         * @return builder
         * 
         */
        public Builder perPage(Integer perPage) {
            return perPage(Output.of(perPage));
        }

        /**
         * @param search Return list of authorized projects matching the search criteria.
         * 
         * @return builder
         * 
         */
        public Builder search(@Nullable Output<String> search) {
            $.search = search;
            return this;
        }

        /**
         * @param search Return list of authorized projects matching the search criteria.
         * 
         * @return builder
         * 
         */
        public Builder search(String search) {
            return search(Output.of(search));
        }

        /**
         * @param simple Return only the ID, URL, name, and path of each project.
         * 
         * @return builder
         * 
         */
        public Builder simple(@Nullable Output<Boolean> simple) {
            $.simple = simple;
            return this;
        }

        /**
         * @param simple Return only the ID, URL, name, and path of each project.
         * 
         * @return builder
         * 
         */
        public Builder simple(Boolean simple) {
            return simple(Output.of(simple));
        }

        /**
         * @param sort Return projects sorted in `asc` or `desc` order. Default is `desc`.
         * 
         * @return builder
         * 
         */
        public Builder sort(@Nullable Output<String> sort) {
            $.sort = sort;
            return this;
        }

        /**
         * @param sort Return projects sorted in `asc` or `desc` order. Default is `desc`.
         * 
         * @return builder
         * 
         */
        public Builder sort(String sort) {
            return sort(Output.of(sort));
        }

        /**
         * @param starred Limit by projects starred by the current user.
         * 
         * @return builder
         * 
         */
        public Builder starred(@Nullable Output<Boolean> starred) {
            $.starred = starred;
            return this;
        }

        /**
         * @param starred Limit by projects starred by the current user.
         * 
         * @return builder
         * 
         */
        public Builder starred(Boolean starred) {
            return starred(Output.of(starred));
        }

        public Builder statistics(@Nullable Output<Boolean> statistics) {
            $.statistics = statistics;
            return this;
        }

        public Builder statistics(Boolean statistics) {
            return statistics(Output.of(statistics));
        }

        /**
         * @param topics Limit by projects that have all of the given topics.
         * 
         * @return builder
         * 
         */
        public Builder topics(@Nullable Output<List<String>> topics) {
            $.topics = topics;
            return this;
        }

        /**
         * @param topics Limit by projects that have all of the given topics.
         * 
         * @return builder
         * 
         */
        public Builder topics(List<String> topics) {
            return topics(Output.of(topics));
        }

        /**
         * @param topics Limit by projects that have all of the given topics.
         * 
         * @return builder
         * 
         */
        public Builder topics(String... topics) {
            return topics(List.of(topics));
        }

        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        /**
         * @param withCustomAttributes Include custom attributes in response *(admins only)*.
         * 
         * @return builder
         * 
         */
        public Builder withCustomAttributes(@Nullable Output<Boolean> withCustomAttributes) {
            $.withCustomAttributes = withCustomAttributes;
            return this;
        }

        /**
         * @param withCustomAttributes Include custom attributes in response *(admins only)*.
         * 
         * @return builder
         * 
         */
        public Builder withCustomAttributes(Boolean withCustomAttributes) {
            return withCustomAttributes(Output.of(withCustomAttributes));
        }

        /**
         * @param withIssuesEnabled Limit by projects with issues feature enabled. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder withIssuesEnabled(@Nullable Output<Boolean> withIssuesEnabled) {
            $.withIssuesEnabled = withIssuesEnabled;
            return this;
        }

        /**
         * @param withIssuesEnabled Limit by projects with issues feature enabled. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder withIssuesEnabled(Boolean withIssuesEnabled) {
            return withIssuesEnabled(Output.of(withIssuesEnabled));
        }

        /**
         * @param withMergeRequestsEnabled Limit by projects with merge requests feature enabled. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder withMergeRequestsEnabled(@Nullable Output<Boolean> withMergeRequestsEnabled) {
            $.withMergeRequestsEnabled = withMergeRequestsEnabled;
            return this;
        }

        /**
         * @param withMergeRequestsEnabled Limit by projects with merge requests feature enabled. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder withMergeRequestsEnabled(Boolean withMergeRequestsEnabled) {
            return withMergeRequestsEnabled(Output.of(withMergeRequestsEnabled));
        }

        /**
         * @param withProgrammingLanguage Limit by projects which use the given programming language. Cannot be used with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder withProgrammingLanguage(@Nullable Output<String> withProgrammingLanguage) {
            $.withProgrammingLanguage = withProgrammingLanguage;
            return this;
        }

        /**
         * @param withProgrammingLanguage Limit by projects which use the given programming language. Cannot be used with `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder withProgrammingLanguage(String withProgrammingLanguage) {
            return withProgrammingLanguage(Output.of(withProgrammingLanguage));
        }

        /**
         * @param withShared Include projects shared to this group. Default is `true`. Needs `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder withShared(@Nullable Output<Boolean> withShared) {
            $.withShared = withShared;
            return this;
        }

        /**
         * @param withShared Include projects shared to this group. Default is `true`. Needs `group_id`.
         * 
         * @return builder
         * 
         */
        public Builder withShared(Boolean withShared) {
            return withShared(Output.of(withShared));
        }

        public GetProjectsArgs build() {
            return $;
        }
    }

}
