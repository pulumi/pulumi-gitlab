// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoryTreePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryTreePlainArgs Empty = new GetRepositoryTreePlainArgs();

    @Import(name="path")
    private @Nullable String path;

    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The ID or full path of the project owned by the authenticated user.
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The ID or full path of the project owned by the authenticated user.
     * 
     */
    public String project() {
        return this.project;
    }

    /**
     * Boolean value used to get a recursive tree (false by default).
     * 
     */
    @Import(name="recursive")
    private @Nullable Boolean recursive;

    /**
     * @return Boolean value used to get a recursive tree (false by default).
     * 
     */
    public Optional<Boolean> recursive() {
        return Optional.ofNullable(this.recursive);
    }

    /**
     * The name of a repository branch or tag.
     * 
     */
    @Import(name="ref", required=true)
    private String ref;

    /**
     * @return The name of a repository branch or tag.
     * 
     */
    public String ref() {
        return this.ref;
    }

    private GetRepositoryTreePlainArgs() {}

    private GetRepositoryTreePlainArgs(GetRepositoryTreePlainArgs $) {
        this.path = $.path;
        this.project = $.project;
        this.recursive = $.recursive;
        this.ref = $.ref;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryTreePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryTreePlainArgs $;

        public Builder() {
            $ = new GetRepositoryTreePlainArgs();
        }

        public Builder(GetRepositoryTreePlainArgs defaults) {
            $ = new GetRepositoryTreePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder path(@Nullable String path) {
            $.path = path;
            return this;
        }

        /**
         * @param project The ID or full path of the project owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        /**
         * @param recursive Boolean value used to get a recursive tree (false by default).
         * 
         * @return builder
         * 
         */
        public Builder recursive(@Nullable Boolean recursive) {
            $.recursive = recursive;
            return this;
        }

        /**
         * @param ref The name of a repository branch or tag.
         * 
         * @return builder
         * 
         */
        public Builder ref(String ref) {
            $.ref = ref;
            return this;
        }

        public GetRepositoryTreePlainArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetRepositoryTreePlainArgs", "project");
            }
            if ($.ref == null) {
                throw new MissingRequiredPropertyException("GetRepositoryTreePlainArgs", "ref");
            }
            return $;
        }
    }

}
