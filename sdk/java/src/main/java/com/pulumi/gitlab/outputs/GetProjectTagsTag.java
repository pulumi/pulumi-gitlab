// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetProjectTagsTagCommit;
import com.pulumi.gitlab.outputs.GetProjectTagsTagRelease;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectTagsTag {
    /**
     * @return The commit associated with the tag.
     * 
     */
    private List<GetProjectTagsTagCommit> commits;
    /**
     * @return The message of the annotated tag.
     * 
     */
    private String message;
    /**
     * @return The name of a tag.
     * 
     */
    private String name;
    /**
     * @return Bool, true if tag has tag protection.
     * 
     */
    private Boolean protected_;
    /**
     * @return The release associated with the tag.
     * 
     */
    private List<GetProjectTagsTagRelease> releases;
    /**
     * @return The unique id assigned to the commit by Gitlab.
     * 
     */
    private String target;

    private GetProjectTagsTag() {}
    /**
     * @return The commit associated with the tag.
     * 
     */
    public List<GetProjectTagsTagCommit> commits() {
        return this.commits;
    }
    /**
     * @return The message of the annotated tag.
     * 
     */
    public String message() {
        return this.message;
    }
    /**
     * @return The name of a tag.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Bool, true if tag has tag protection.
     * 
     */
    public Boolean protected_() {
        return this.protected_;
    }
    /**
     * @return The release associated with the tag.
     * 
     */
    public List<GetProjectTagsTagRelease> releases() {
        return this.releases;
    }
    /**
     * @return The unique id assigned to the commit by Gitlab.
     * 
     */
    public String target() {
        return this.target;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectTagsTag defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetProjectTagsTagCommit> commits;
        private String message;
        private String name;
        private Boolean protected_;
        private List<GetProjectTagsTagRelease> releases;
        private String target;
        public Builder() {}
        public Builder(GetProjectTagsTag defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commits = defaults.commits;
    	      this.message = defaults.message;
    	      this.name = defaults.name;
    	      this.protected_ = defaults.protected_;
    	      this.releases = defaults.releases;
    	      this.target = defaults.target;
        }

        @CustomType.Setter
        public Builder commits(List<GetProjectTagsTagCommit> commits) {
            if (commits == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTag", "commits");
            }
            this.commits = commits;
            return this;
        }
        public Builder commits(GetProjectTagsTagCommit... commits) {
            return commits(List.of(commits));
        }
        @CustomType.Setter
        public Builder message(String message) {
            if (message == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTag", "message");
            }
            this.message = message;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTag", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter("protected")
        public Builder protected_(Boolean protected_) {
            if (protected_ == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTag", "protected_");
            }
            this.protected_ = protected_;
            return this;
        }
        @CustomType.Setter
        public Builder releases(List<GetProjectTagsTagRelease> releases) {
            if (releases == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTag", "releases");
            }
            this.releases = releases;
            return this;
        }
        public Builder releases(GetProjectTagsTagRelease... releases) {
            return releases(List.of(releases));
        }
        @CustomType.Setter
        public Builder target(String target) {
            if (target == null) {
              throw new MissingRequiredPropertyException("GetProjectTagsTag", "target");
            }
            this.target = target;
            return this;
        }
        public GetProjectTagsTag build() {
            final var _resultValue = new GetProjectTagsTag();
            _resultValue.commits = commits;
            _resultValue.message = message;
            _resultValue.name = name;
            _resultValue.protected_ = protected_;
            _resultValue.releases = releases;
            _resultValue.target = target;
            return _resultValue;
        }
    }
}
