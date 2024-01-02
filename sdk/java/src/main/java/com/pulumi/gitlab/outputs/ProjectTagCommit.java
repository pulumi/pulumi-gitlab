// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProjectTagCommit {
    private @Nullable String authorEmail;
    private @Nullable String authorName;
    private @Nullable String authoredDate;
    private @Nullable String committedDate;
    private @Nullable String committerEmail;
    private @Nullable String committerName;
    private @Nullable String id;
    private @Nullable String message;
    private @Nullable List<String> parentIds;
    private @Nullable String shortId;
    private @Nullable String title;

    private ProjectTagCommit() {}
    public Optional<String> authorEmail() {
        return Optional.ofNullable(this.authorEmail);
    }
    public Optional<String> authorName() {
        return Optional.ofNullable(this.authorName);
    }
    public Optional<String> authoredDate() {
        return Optional.ofNullable(this.authoredDate);
    }
    public Optional<String> committedDate() {
        return Optional.ofNullable(this.committedDate);
    }
    public Optional<String> committerEmail() {
        return Optional.ofNullable(this.committerEmail);
    }
    public Optional<String> committerName() {
        return Optional.ofNullable(this.committerName);
    }
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    public Optional<String> message() {
        return Optional.ofNullable(this.message);
    }
    public List<String> parentIds() {
        return this.parentIds == null ? List.of() : this.parentIds;
    }
    public Optional<String> shortId() {
        return Optional.ofNullable(this.shortId);
    }
    public Optional<String> title() {
        return Optional.ofNullable(this.title);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProjectTagCommit defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authorEmail;
        private @Nullable String authorName;
        private @Nullable String authoredDate;
        private @Nullable String committedDate;
        private @Nullable String committerEmail;
        private @Nullable String committerName;
        private @Nullable String id;
        private @Nullable String message;
        private @Nullable List<String> parentIds;
        private @Nullable String shortId;
        private @Nullable String title;
        public Builder() {}
        public Builder(ProjectTagCommit defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorEmail = defaults.authorEmail;
    	      this.authorName = defaults.authorName;
    	      this.authoredDate = defaults.authoredDate;
    	      this.committedDate = defaults.committedDate;
    	      this.committerEmail = defaults.committerEmail;
    	      this.committerName = defaults.committerName;
    	      this.id = defaults.id;
    	      this.message = defaults.message;
    	      this.parentIds = defaults.parentIds;
    	      this.shortId = defaults.shortId;
    	      this.title = defaults.title;
        }

        @CustomType.Setter
        public Builder authorEmail(@Nullable String authorEmail) {

            this.authorEmail = authorEmail;
            return this;
        }
        @CustomType.Setter
        public Builder authorName(@Nullable String authorName) {

            this.authorName = authorName;
            return this;
        }
        @CustomType.Setter
        public Builder authoredDate(@Nullable String authoredDate) {

            this.authoredDate = authoredDate;
            return this;
        }
        @CustomType.Setter
        public Builder committedDate(@Nullable String committedDate) {

            this.committedDate = committedDate;
            return this;
        }
        @CustomType.Setter
        public Builder committerEmail(@Nullable String committerEmail) {

            this.committerEmail = committerEmail;
            return this;
        }
        @CustomType.Setter
        public Builder committerName(@Nullable String committerName) {

            this.committerName = committerName;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder message(@Nullable String message) {

            this.message = message;
            return this;
        }
        @CustomType.Setter
        public Builder parentIds(@Nullable List<String> parentIds) {

            this.parentIds = parentIds;
            return this;
        }
        public Builder parentIds(String... parentIds) {
            return parentIds(List.of(parentIds));
        }
        @CustomType.Setter
        public Builder shortId(@Nullable String shortId) {

            this.shortId = shortId;
            return this;
        }
        @CustomType.Setter
        public Builder title(@Nullable String title) {

            this.title = title;
            return this;
        }
        public ProjectTagCommit build() {
            final var _resultValue = new ProjectTagCommit();
            _resultValue.authorEmail = authorEmail;
            _resultValue.authorName = authorName;
            _resultValue.authoredDate = authoredDate;
            _resultValue.committedDate = committedDate;
            _resultValue.committerEmail = committerEmail;
            _resultValue.committerName = committerName;
            _resultValue.id = id;
            _resultValue.message = message;
            _resultValue.parentIds = parentIds;
            _resultValue.shortId = shortId;
            _resultValue.title = title;
            return _resultValue;
        }
    }
}
