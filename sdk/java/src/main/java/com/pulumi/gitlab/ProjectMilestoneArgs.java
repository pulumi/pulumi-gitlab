// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectMilestoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectMilestoneArgs Empty = new ProjectMilestoneArgs();

    /**
     * The description of the milestone.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the milestone.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    @Import(name="dueDate")
    private @Nullable Output<String> dueDate;

    /**
     * @return The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public Optional<Output<String>> dueDate() {
        return Optional.ofNullable(this.dueDate);
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
     * The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    @Import(name="startDate")
    private @Nullable Output<String> startDate;

    /**
     * @return The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public Optional<Output<String>> startDate() {
        return Optional.ofNullable(this.startDate);
    }

    /**
     * The state of the milestone. Valid values are: `active`, `closed`.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The state of the milestone. Valid values are: `active`, `closed`.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * The title of a milestone.
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return The title of a milestone.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private ProjectMilestoneArgs() {}

    private ProjectMilestoneArgs(ProjectMilestoneArgs $) {
        this.description = $.description;
        this.dueDate = $.dueDate;
        this.project = $.project;
        this.startDate = $.startDate;
        this.state = $.state;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectMilestoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectMilestoneArgs $;

        public Builder() {
            $ = new ProjectMilestoneArgs();
        }

        public Builder(ProjectMilestoneArgs defaults) {
            $ = new ProjectMilestoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the milestone.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the milestone.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dueDate The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
         * 
         * @return builder
         * 
         */
        public Builder dueDate(@Nullable Output<String> dueDate) {
            $.dueDate = dueDate;
            return this;
        }

        /**
         * @param dueDate The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
         * 
         * @return builder
         * 
         */
        public Builder dueDate(String dueDate) {
            return dueDate(Output.of(dueDate));
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
         * @param startDate The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
         * 
         * @return builder
         * 
         */
        public Builder startDate(@Nullable Output<String> startDate) {
            $.startDate = startDate;
            return this;
        }

        /**
         * @param startDate The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
         * 
         * @return builder
         * 
         */
        public Builder startDate(String startDate) {
            return startDate(Output.of(startDate));
        }

        /**
         * @param state The state of the milestone. Valid values are: `active`, `closed`.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The state of the milestone. Valid values are: `active`, `closed`.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param title The title of a milestone.
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title The title of a milestone.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public ProjectMilestoneArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("ProjectMilestoneArgs", "project");
            }
            if ($.title == null) {
                throw new MissingRequiredPropertyException("ProjectMilestoneArgs", "title");
            }
            return $;
        }
    }

}
