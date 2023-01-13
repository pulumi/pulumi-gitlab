// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectTagsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectTagsArgs Empty = new GetProjectTagsArgs();

    /**
     * Return tags ordered by `name` or `updated` fields. Default is `updated`.
     * 
     */
    @Import(name="orderBy")
    private @Nullable Output<String> orderBy;

    /**
     * @return Return tags ordered by `name` or `updated` fields. Default is `updated`.
     * 
     */
    public Optional<Output<String>> orderBy() {
        return Optional.ofNullable(this.orderBy);
    }

    /**
     * The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or URL-encoded path of the project owned by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
     * 
     */
    @Import(name="search")
    private @Nullable Output<String> search;

    /**
     * @return Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
     * 
     */
    public Optional<Output<String>> search() {
        return Optional.ofNullable(this.search);
    }

    /**
     * Return tags sorted in `asc` or `desc` order. Default is `desc`.
     * 
     */
    @Import(name="sort")
    private @Nullable Output<String> sort;

    /**
     * @return Return tags sorted in `asc` or `desc` order. Default is `desc`.
     * 
     */
    public Optional<Output<String>> sort() {
        return Optional.ofNullable(this.sort);
    }

    private GetProjectTagsArgs() {}

    private GetProjectTagsArgs(GetProjectTagsArgs $) {
        this.orderBy = $.orderBy;
        this.project = $.project;
        this.search = $.search;
        this.sort = $.sort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectTagsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectTagsArgs $;

        public Builder() {
            $ = new GetProjectTagsArgs();
        }

        public Builder(GetProjectTagsArgs defaults) {
            $ = new GetProjectTagsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param orderBy Return tags ordered by `name` or `updated` fields. Default is `updated`.
         * 
         * @return builder
         * 
         */
        public Builder orderBy(@Nullable Output<String> orderBy) {
            $.orderBy = orderBy;
            return this;
        }

        /**
         * @param orderBy Return tags ordered by `name` or `updated` fields. Default is `updated`.
         * 
         * @return builder
         * 
         */
        public Builder orderBy(String orderBy) {
            return orderBy(Output.of(orderBy));
        }

        /**
         * @param project The ID or URL-encoded path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or URL-encoded path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param search Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
         * 
         * @return builder
         * 
         */
        public Builder search(@Nullable Output<String> search) {
            $.search = search;
            return this;
        }

        /**
         * @param search Return list of tags matching the search criteria. You can use `^term` and `term$` to find tags that begin and end with `term` respectively. No other regular expressions are supported.
         * 
         * @return builder
         * 
         */
        public Builder search(String search) {
            return search(Output.of(search));
        }

        /**
         * @param sort Return tags sorted in `asc` or `desc` order. Default is `desc`.
         * 
         * @return builder
         * 
         */
        public Builder sort(@Nullable Output<String> sort) {
            $.sort = sort;
            return this;
        }

        /**
         * @param sort Return tags sorted in `asc` or `desc` order. Default is `desc`.
         * 
         * @return builder
         * 
         */
        public Builder sort(String sort) {
            return sort(Output.of(sort));
        }

        public GetProjectTagsArgs build() {
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            return $;
        }
    }

}