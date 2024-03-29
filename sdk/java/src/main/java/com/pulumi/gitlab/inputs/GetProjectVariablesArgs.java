// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProjectVariablesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectVariablesArgs Empty = new GetProjectVariablesArgs();

    @Import(name="environmentScope")
    private @Nullable Output<String> environmentScope;

    public Optional<Output<String>> environmentScope() {
        return Optional.ofNullable(this.environmentScope);
    }

    @Import(name="project", required=true)
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }

    private GetProjectVariablesArgs() {}

    private GetProjectVariablesArgs(GetProjectVariablesArgs $) {
        this.environmentScope = $.environmentScope;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectVariablesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectVariablesArgs $;

        public Builder() {
            $ = new GetProjectVariablesArgs();
        }

        public Builder(GetProjectVariablesArgs defaults) {
            $ = new GetProjectVariablesArgs(Objects.requireNonNull(defaults));
        }

        public Builder environmentScope(@Nullable Output<String> environmentScope) {
            $.environmentScope = environmentScope;
            return this;
        }

        public Builder environmentScope(String environmentScope) {
            return environmentScope(Output.of(environmentScope));
        }

        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetProjectVariablesArgs build() {
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetProjectVariablesArgs", "project");
            }
            return $;
        }
    }

}
