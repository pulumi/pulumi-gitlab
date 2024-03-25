// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.GetReleaseAssets;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetReleasePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReleasePlainArgs Empty = new GetReleasePlainArgs();

    /**
     * The assets for a release
     * 
     */
    @Import(name="assets")
    private @Nullable GetReleaseAssets assets;

    /**
     * @return The assets for a release
     * 
     */
    public Optional<GetReleaseAssets> assets() {
        return Optional.ofNullable(this.assets);
    }

    /**
     * The ID or URL-encoded path of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The ID or URL-encoded path of the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    /**
     * The Git tag the release is associated with.
     * 
     */
    @Import(name="tagName", required=true)
    private String tagName;

    /**
     * @return The Git tag the release is associated with.
     * 
     */
    public String tagName() {
        return this.tagName;
    }

    private GetReleasePlainArgs() {}

    private GetReleasePlainArgs(GetReleasePlainArgs $) {
        this.assets = $.assets;
        this.projectId = $.projectId;
        this.tagName = $.tagName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReleasePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReleasePlainArgs $;

        public Builder() {
            $ = new GetReleasePlainArgs();
        }

        public Builder(GetReleasePlainArgs defaults) {
            $ = new GetReleasePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assets The assets for a release
         * 
         * @return builder
         * 
         */
        public Builder assets(@Nullable GetReleaseAssets assets) {
            $.assets = assets;
            return this;
        }

        /**
         * @param projectId The ID or URL-encoded path of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param tagName The Git tag the release is associated with.
         * 
         * @return builder
         * 
         */
        public Builder tagName(String tagName) {
            $.tagName = tagName;
            return this;
        }

        public GetReleasePlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetReleasePlainArgs", "projectId");
            }
            if ($.tagName == null) {
                throw new MissingRequiredPropertyException("GetReleasePlainArgs", "tagName");
            }
            return $;
        }
    }

}
