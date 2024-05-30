// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipelineScheduleState extends com.pulumi.resources.ResourceArgs {

    public static final PipelineScheduleState Empty = new PipelineScheduleState();

    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * The cron (e.g. `0 1 * * *`).
     * 
     */
    @Import(name="cron")
    private @Nullable Output<String> cron;

    /**
     * @return The cron (e.g. `0 1 * * *`).
     * 
     */
    public Optional<Output<String>> cron() {
        return Optional.ofNullable(this.cron);
    }

    /**
     * The timezone.
     * 
     */
    @Import(name="cronTimezone")
    private @Nullable Output<String> cronTimezone;

    /**
     * @return The timezone.
     * 
     */
    public Optional<Output<String>> cronTimezone() {
        return Optional.ofNullable(this.cronTimezone);
    }

    /**
     * The description of the pipeline schedule.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the pipeline schedule.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the user that owns the pipeline schedule.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<Integer> owner;

    /**
     * @return The ID of the user that owns the pipeline schedule.
     * 
     */
    public Optional<Output<Integer>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * The pipeline schedule id.
     * 
     */
    @Import(name="pipelineScheduleId")
    private @Nullable Output<Integer> pipelineScheduleId;

    /**
     * @return The pipeline schedule id.
     * 
     */
    public Optional<Output<Integer>> pipelineScheduleId() {
        return Optional.ofNullable(this.pipelineScheduleId);
    }

    /**
     * The name or id of the project to add the schedule to.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The name or id of the project to add the schedule to.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The branch/tag name to be triggered. This must be the full branch reference, for example: `refs/heads/main`, not `main`.
     * 
     */
    @Import(name="ref")
    private @Nullable Output<String> ref;

    /**
     * @return The branch/tag name to be triggered. This must be the full branch reference, for example: `refs/heads/main`, not `main`.
     * 
     */
    public Optional<Output<String>> ref() {
        return Optional.ofNullable(this.ref);
    }

    @Import(name="takeOwnership")
    private @Nullable Output<Boolean> takeOwnership;

    public Optional<Output<Boolean>> takeOwnership() {
        return Optional.ofNullable(this.takeOwnership);
    }

    private PipelineScheduleState() {}

    private PipelineScheduleState(PipelineScheduleState $) {
        this.active = $.active;
        this.cron = $.cron;
        this.cronTimezone = $.cronTimezone;
        this.description = $.description;
        this.owner = $.owner;
        this.pipelineScheduleId = $.pipelineScheduleId;
        this.project = $.project;
        this.ref = $.ref;
        this.takeOwnership = $.takeOwnership;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipelineScheduleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipelineScheduleState $;

        public Builder() {
            $ = new PipelineScheduleState();
        }

        public Builder(PipelineScheduleState defaults) {
            $ = new PipelineScheduleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param active The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param cron The cron (e.g. `0 1 * * *`).
         * 
         * @return builder
         * 
         */
        public Builder cron(@Nullable Output<String> cron) {
            $.cron = cron;
            return this;
        }

        /**
         * @param cron The cron (e.g. `0 1 * * *`).
         * 
         * @return builder
         * 
         */
        public Builder cron(String cron) {
            return cron(Output.of(cron));
        }

        /**
         * @param cronTimezone The timezone.
         * 
         * @return builder
         * 
         */
        public Builder cronTimezone(@Nullable Output<String> cronTimezone) {
            $.cronTimezone = cronTimezone;
            return this;
        }

        /**
         * @param cronTimezone The timezone.
         * 
         * @return builder
         * 
         */
        public Builder cronTimezone(String cronTimezone) {
            return cronTimezone(Output.of(cronTimezone));
        }

        /**
         * @param description The description of the pipeline schedule.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the pipeline schedule.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param owner The ID of the user that owns the pipeline schedule.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<Integer> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The ID of the user that owns the pipeline schedule.
         * 
         * @return builder
         * 
         */
        public Builder owner(Integer owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param pipelineScheduleId The pipeline schedule id.
         * 
         * @return builder
         * 
         */
        public Builder pipelineScheduleId(@Nullable Output<Integer> pipelineScheduleId) {
            $.pipelineScheduleId = pipelineScheduleId;
            return this;
        }

        /**
         * @param pipelineScheduleId The pipeline schedule id.
         * 
         * @return builder
         * 
         */
        public Builder pipelineScheduleId(Integer pipelineScheduleId) {
            return pipelineScheduleId(Output.of(pipelineScheduleId));
        }

        /**
         * @param project The name or id of the project to add the schedule to.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project to add the schedule to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param ref The branch/tag name to be triggered. This must be the full branch reference, for example: `refs/heads/main`, not `main`.
         * 
         * @return builder
         * 
         */
        public Builder ref(@Nullable Output<String> ref) {
            $.ref = ref;
            return this;
        }

        /**
         * @param ref The branch/tag name to be triggered. This must be the full branch reference, for example: `refs/heads/main`, not `main`.
         * 
         * @return builder
         * 
         */
        public Builder ref(String ref) {
            return ref(Output.of(ref));
        }

        public Builder takeOwnership(@Nullable Output<Boolean> takeOwnership) {
            $.takeOwnership = takeOwnership;
            return this;
        }

        public Builder takeOwnership(Boolean takeOwnership) {
            return takeOwnership(Output.of(takeOwnership));
        }

        public PipelineScheduleState build() {
            return $;
        }
    }

}
