// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProjectContainerExpirationPolicy {
    /**
     * @return The cadence of the policy. Valid values are: `1d`, `7d`, `14d`, `1month`, `3month`.
     * 
     */
    private @Nullable String cadence;
    /**
     * @return If true, the policy is enabled.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return The number of images to keep.
     * 
     */
    private @Nullable Integer keepN;
    /**
     * @return The regular expression to match image names to delete.
     * 
     */
    private @Nullable String nameRegex;
    /**
     * @return The regular expression to match image names to delete.
     * 
     * @deprecated
     * `name_regex_delete` has been deprecated. Use `name_regex` instead.
     * 
     */
    @Deprecated /* `name_regex_delete` has been deprecated. Use `name_regex` instead. */
    private @Nullable String nameRegexDelete;
    /**
     * @return The regular expression to match image names to keep.
     * 
     */
    private @Nullable String nameRegexKeep;
    /**
     * @return The next time the policy will run.
     * 
     */
    private @Nullable String nextRunAt;
    /**
     * @return The number of days to keep images.
     * 
     */
    private @Nullable String olderThan;

    private ProjectContainerExpirationPolicy() {}
    /**
     * @return The cadence of the policy. Valid values are: `1d`, `7d`, `14d`, `1month`, `3month`.
     * 
     */
    public Optional<String> cadence() {
        return Optional.ofNullable(this.cadence);
    }
    /**
     * @return If true, the policy is enabled.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return The number of images to keep.
     * 
     */
    public Optional<Integer> keepN() {
        return Optional.ofNullable(this.keepN);
    }
    /**
     * @return The regular expression to match image names to delete.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return The regular expression to match image names to delete.
     * 
     * @deprecated
     * `name_regex_delete` has been deprecated. Use `name_regex` instead.
     * 
     */
    @Deprecated /* `name_regex_delete` has been deprecated. Use `name_regex` instead. */
    public Optional<String> nameRegexDelete() {
        return Optional.ofNullable(this.nameRegexDelete);
    }
    /**
     * @return The regular expression to match image names to keep.
     * 
     */
    public Optional<String> nameRegexKeep() {
        return Optional.ofNullable(this.nameRegexKeep);
    }
    /**
     * @return The next time the policy will run.
     * 
     */
    public Optional<String> nextRunAt() {
        return Optional.ofNullable(this.nextRunAt);
    }
    /**
     * @return The number of days to keep images.
     * 
     */
    public Optional<String> olderThan() {
        return Optional.ofNullable(this.olderThan);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProjectContainerExpirationPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cadence;
        private @Nullable Boolean enabled;
        private @Nullable Integer keepN;
        private @Nullable String nameRegex;
        private @Nullable String nameRegexDelete;
        private @Nullable String nameRegexKeep;
        private @Nullable String nextRunAt;
        private @Nullable String olderThan;
        public Builder() {}
        public Builder(ProjectContainerExpirationPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cadence = defaults.cadence;
    	      this.enabled = defaults.enabled;
    	      this.keepN = defaults.keepN;
    	      this.nameRegex = defaults.nameRegex;
    	      this.nameRegexDelete = defaults.nameRegexDelete;
    	      this.nameRegexKeep = defaults.nameRegexKeep;
    	      this.nextRunAt = defaults.nextRunAt;
    	      this.olderThan = defaults.olderThan;
        }

        @CustomType.Setter
        public Builder cadence(@Nullable String cadence) {
            this.cadence = cadence;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder keepN(@Nullable Integer keepN) {
            this.keepN = keepN;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegexDelete(@Nullable String nameRegexDelete) {
            this.nameRegexDelete = nameRegexDelete;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegexKeep(@Nullable String nameRegexKeep) {
            this.nameRegexKeep = nameRegexKeep;
            return this;
        }
        @CustomType.Setter
        public Builder nextRunAt(@Nullable String nextRunAt) {
            this.nextRunAt = nextRunAt;
            return this;
        }
        @CustomType.Setter
        public Builder olderThan(@Nullable String olderThan) {
            this.olderThan = olderThan;
            return this;
        }
        public ProjectContainerExpirationPolicy build() {
            final var o = new ProjectContainerExpirationPolicy();
            o.cadence = cadence;
            o.enabled = enabled;
            o.keepN = keepN;
            o.nameRegex = nameRegex;
            o.nameRegexDelete = nameRegexDelete;
            o.nameRegexKeep = nameRegexKeep;
            o.nextRunAt = nextRunAt;
            o.olderThan = olderThan;
            return o;
        }
    }
}
