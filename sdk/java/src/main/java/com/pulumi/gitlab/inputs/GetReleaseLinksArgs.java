// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetReleaseLinksArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReleaseLinksArgs Empty = new GetReleaseLinksArgs();

    @Import(name="project", required=true)
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }

    @Import(name="tagName", required=true)
    private Output<String> tagName;

    public Output<String> tagName() {
        return this.tagName;
    }

    private GetReleaseLinksArgs() {}

    private GetReleaseLinksArgs(GetReleaseLinksArgs $) {
        this.project = $.project;
        this.tagName = $.tagName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReleaseLinksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReleaseLinksArgs $;

        public Builder() {
            $ = new GetReleaseLinksArgs();
        }

        public Builder(GetReleaseLinksArgs defaults) {
            $ = new GetReleaseLinksArgs(Objects.requireNonNull(defaults));
        }

        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        public Builder project(String project) {
            return project(Output.of(project));
        }

        public Builder tagName(Output<String> tagName) {
            $.tagName = tagName;
            return this;
        }

        public Builder tagName(String tagName) {
            return tagName(Output.of(tagName));
        }

        public GetReleaseLinksArgs build() {
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            $.tagName = Objects.requireNonNull($.tagName, "expected parameter 'tagName' to be non-null");
            return $;
        }
    }

}
