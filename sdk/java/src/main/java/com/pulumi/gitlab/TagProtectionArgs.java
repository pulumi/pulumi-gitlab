// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.TagProtectionAllowedToCreateArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TagProtectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final TagProtectionArgs Empty = new TagProtectionArgs();

    /**
     * User or group which are allowed to create.
     * 
     */
    @Import(name="allowedToCreates")
    private @Nullable Output<List<TagProtectionAllowedToCreateArgs>> allowedToCreates;

    /**
     * @return User or group which are allowed to create.
     * 
     */
    public Optional<Output<List<TagProtectionAllowedToCreateArgs>>> allowedToCreates() {
        return Optional.ofNullable(this.allowedToCreates);
    }

    /**
     * Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    @Import(name="createAccessLevel", required=true)
    private Output<String> createAccessLevel;

    /**
     * @return Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
     * 
     */
    public Output<String> createAccessLevel() {
        return this.createAccessLevel;
    }

    /**
     * The id of the project.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The id of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Name of the tag or wildcard.
     * 
     */
    @Import(name="tag", required=true)
    private Output<String> tag;

    /**
     * @return Name of the tag or wildcard.
     * 
     */
    public Output<String> tag() {
        return this.tag;
    }

    private TagProtectionArgs() {}

    private TagProtectionArgs(TagProtectionArgs $) {
        this.allowedToCreates = $.allowedToCreates;
        this.createAccessLevel = $.createAccessLevel;
        this.project = $.project;
        this.tag = $.tag;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TagProtectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TagProtectionArgs $;

        public Builder() {
            $ = new TagProtectionArgs();
        }

        public Builder(TagProtectionArgs defaults) {
            $ = new TagProtectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedToCreates User or group which are allowed to create.
         * 
         * @return builder
         * 
         */
        public Builder allowedToCreates(@Nullable Output<List<TagProtectionAllowedToCreateArgs>> allowedToCreates) {
            $.allowedToCreates = allowedToCreates;
            return this;
        }

        /**
         * @param allowedToCreates User or group which are allowed to create.
         * 
         * @return builder
         * 
         */
        public Builder allowedToCreates(List<TagProtectionAllowedToCreateArgs> allowedToCreates) {
            return allowedToCreates(Output.of(allowedToCreates));
        }

        /**
         * @param allowedToCreates User or group which are allowed to create.
         * 
         * @return builder
         * 
         */
        public Builder allowedToCreates(TagProtectionAllowedToCreateArgs... allowedToCreates) {
            return allowedToCreates(List.of(allowedToCreates));
        }

        /**
         * @param createAccessLevel Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder createAccessLevel(Output<String> createAccessLevel) {
            $.createAccessLevel = createAccessLevel;
            return this;
        }

        /**
         * @param createAccessLevel Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
         * 
         * @return builder
         * 
         */
        public Builder createAccessLevel(String createAccessLevel) {
            return createAccessLevel(Output.of(createAccessLevel));
        }

        /**
         * @param project The id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The id of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param tag Name of the tag or wildcard.
         * 
         * @return builder
         * 
         */
        public Builder tag(Output<String> tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tag Name of the tag or wildcard.
         * 
         * @return builder
         * 
         */
        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        public TagProtectionArgs build() {
            if ($.createAccessLevel == null) {
                throw new MissingRequiredPropertyException("TagProtectionArgs", "createAccessLevel");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("TagProtectionArgs", "project");
            }
            if ($.tag == null) {
                throw new MissingRequiredPropertyException("TagProtectionArgs", "tag");
            }
            return $;
        }
    }

}
