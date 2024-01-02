// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComplianceFrameworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComplianceFrameworkArgs Empty = new ComplianceFrameworkArgs();

    /**
     * New color representation of the compliance framework in hex format. e.g. #FCA121.
     * 
     */
    @Import(name="color", required=true)
    private Output<String> color;

    /**
     * @return New color representation of the compliance framework in hex format. e.g. #FCA121.
     * 
     */
    public Output<String> color() {
        return this.color;
    }

    /**
     * Set this compliance framework as the default framework for the group. Default: `false`
     * 
     */
    @Import(name="default")
    private @Nullable Output<Boolean> default_;

    /**
     * @return Set this compliance framework as the default framework for the group. Default: `false`
     * 
     */
    public Optional<Output<Boolean>> default_() {
        return Optional.ofNullable(this.default_);
    }

    /**
     * Description for the compliance framework.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return Description for the compliance framework.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * Name for the compliance framework.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name for the compliance framework.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Full path of the namespace to add the compliance framework to.
     * 
     */
    @Import(name="namespacePath", required=true)
    private Output<String> namespacePath;

    /**
     * @return Full path of the namespace to add the compliance framework to.
     * 
     */
    public Output<String> namespacePath() {
        return this.namespacePath;
    }

    /**
     * Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
     * 
     */
    @Import(name="pipelineConfigurationFullPath")
    private @Nullable Output<String> pipelineConfigurationFullPath;

    /**
     * @return Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
     * 
     */
    public Optional<Output<String>> pipelineConfigurationFullPath() {
        return Optional.ofNullable(this.pipelineConfigurationFullPath);
    }

    private ComplianceFrameworkArgs() {}

    private ComplianceFrameworkArgs(ComplianceFrameworkArgs $) {
        this.color = $.color;
        this.default_ = $.default_;
        this.description = $.description;
        this.name = $.name;
        this.namespacePath = $.namespacePath;
        this.pipelineConfigurationFullPath = $.pipelineConfigurationFullPath;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComplianceFrameworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComplianceFrameworkArgs $;

        public Builder() {
            $ = new ComplianceFrameworkArgs();
        }

        public Builder(ComplianceFrameworkArgs defaults) {
            $ = new ComplianceFrameworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param color New color representation of the compliance framework in hex format. e.g. #FCA121.
         * 
         * @return builder
         * 
         */
        public Builder color(Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color New color representation of the compliance framework in hex format. e.g. #FCA121.
         * 
         * @return builder
         * 
         */
        public Builder color(String color) {
            return color(Output.of(color));
        }

        /**
         * @param default_ Set this compliance framework as the default framework for the group. Default: `false`
         * 
         * @return builder
         * 
         */
        public Builder default_(@Nullable Output<Boolean> default_) {
            $.default_ = default_;
            return this;
        }

        /**
         * @param default_ Set this compliance framework as the default framework for the group. Default: `false`
         * 
         * @return builder
         * 
         */
        public Builder default_(Boolean default_) {
            return default_(Output.of(default_));
        }

        /**
         * @param description Description for the compliance framework.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the compliance framework.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name for the compliance framework.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name for the compliance framework.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespacePath Full path of the namespace to add the compliance framework to.
         * 
         * @return builder
         * 
         */
        public Builder namespacePath(Output<String> namespacePath) {
            $.namespacePath = namespacePath;
            return this;
        }

        /**
         * @param namespacePath Full path of the namespace to add the compliance framework to.
         * 
         * @return builder
         * 
         */
        public Builder namespacePath(String namespacePath) {
            return namespacePath(Output.of(namespacePath));
        }

        /**
         * @param pipelineConfigurationFullPath Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
         * 
         * @return builder
         * 
         */
        public Builder pipelineConfigurationFullPath(@Nullable Output<String> pipelineConfigurationFullPath) {
            $.pipelineConfigurationFullPath = pipelineConfigurationFullPath;
            return this;
        }

        /**
         * @param pipelineConfigurationFullPath Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
         * 
         * @return builder
         * 
         */
        public Builder pipelineConfigurationFullPath(String pipelineConfigurationFullPath) {
            return pipelineConfigurationFullPath(Output.of(pipelineConfigurationFullPath));
        }

        public ComplianceFrameworkArgs build() {
            if ($.color == null) {
                throw new MissingRequiredPropertyException("ComplianceFrameworkArgs", "color");
            }
            if ($.description == null) {
                throw new MissingRequiredPropertyException("ComplianceFrameworkArgs", "description");
            }
            if ($.namespacePath == null) {
                throw new MissingRequiredPropertyException("ComplianceFrameworkArgs", "namespacePath");
            }
            return $;
        }
    }

}
