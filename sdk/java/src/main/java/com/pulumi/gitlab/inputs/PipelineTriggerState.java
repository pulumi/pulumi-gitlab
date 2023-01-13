// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipelineTriggerState extends com.pulumi.resources.ResourceArgs {

    public static final PipelineTriggerState Empty = new PipelineTriggerState();

    /**
     * The description of the pipeline trigger.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the pipeline trigger.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name or id of the project to add the trigger to.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The name or id of the project to add the trigger to.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The pipeline trigger token.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The pipeline trigger token.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    private PipelineTriggerState() {}

    private PipelineTriggerState(PipelineTriggerState $) {
        this.description = $.description;
        this.project = $.project;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipelineTriggerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipelineTriggerState $;

        public Builder() {
            $ = new PipelineTriggerState();
        }

        public Builder(PipelineTriggerState defaults) {
            $ = new PipelineTriggerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the pipeline trigger.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the pipeline trigger.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param project The name or id of the project to add the trigger to.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project to add the trigger to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param token The pipeline trigger token.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The pipeline trigger token.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public PipelineTriggerState build() {
            return $;
        }
    }

}