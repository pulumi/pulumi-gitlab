// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetReleaseLinkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReleaseLinkPlainArgs Empty = new GetReleaseLinkPlainArgs();

    /**
     * The ID of the link.
     * 
     */
    @Import(name="linkId", required=true)
    private Integer linkId;

    /**
     * @return The ID of the link.
     * 
     */
    public Integer linkId() {
        return this.linkId;
    }

    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
     * 
     */
    @Import(name="project", required=true)
    private String project;

    /**
     * @return The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
     * 
     */
    public String project() {
        return this.project;
    }

    /**
     * The tag associated with the Release.
     * 
     */
    @Import(name="tagName", required=true)
    private String tagName;

    /**
     * @return The tag associated with the Release.
     * 
     */
    public String tagName() {
        return this.tagName;
    }

    private GetReleaseLinkPlainArgs() {}

    private GetReleaseLinkPlainArgs(GetReleaseLinkPlainArgs $) {
        this.linkId = $.linkId;
        this.project = $.project;
        this.tagName = $.tagName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReleaseLinkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReleaseLinkPlainArgs $;

        public Builder() {
            $ = new GetReleaseLinkPlainArgs();
        }

        public Builder(GetReleaseLinkPlainArgs defaults) {
            $ = new GetReleaseLinkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param linkId The ID of the link.
         * 
         * @return builder
         * 
         */
        public Builder linkId(Integer linkId) {
            $.linkId = linkId;
            return this;
        }

        /**
         * @param project The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding).
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            $.project = project;
            return this;
        }

        /**
         * @param tagName The tag associated with the Release.
         * 
         * @return builder
         * 
         */
        public Builder tagName(String tagName) {
            $.tagName = tagName;
            return this;
        }

        public GetReleaseLinkPlainArgs build() {
            if ($.linkId == null) {
                throw new MissingRequiredPropertyException("GetReleaseLinkPlainArgs", "linkId");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetReleaseLinkPlainArgs", "project");
            }
            if ($.tagName == null) {
                throw new MissingRequiredPropertyException("GetReleaseLinkPlainArgs", "tagName");
            }
            return $;
        }
    }

}
