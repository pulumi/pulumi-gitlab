// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectTagArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectTagArgs Empty = new ProjectTagArgs();

    /**
     * The message of the annotated tag.
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return The message of the annotated tag.
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * The name of a tag.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of a tag.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
     * 
     */
    @Import(name="ref", required=true)
    private Output<String> ref;

    /**
     * @return Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
     * 
     */
    public Output<String> ref() {
        return this.ref;
    }

    private ProjectTagArgs() {}

    private ProjectTagArgs(ProjectTagArgs $) {
        this.message = $.message;
        this.name = $.name;
        this.project = $.project;
        this.ref = $.ref;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectTagArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectTagArgs $;

        public Builder() {
            $ = new ProjectTagArgs();
        }

        public Builder(ProjectTagArgs defaults) {
            $ = new ProjectTagArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param message The message of the annotated tag.
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message The message of the annotated tag.
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param name The name of a tag.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of a tag.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param ref Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder ref(Output<String> ref) {
            $.ref = ref;
            return this;
        }

        /**
         * @param ref Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder ref(String ref) {
            return ref(Output.of(ref));
        }

        public ProjectTagArgs build() {
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            $.ref = Objects.requireNonNull($.ref, "expected parameter 'ref' to be non-null");
            return $;
        }
    }

}