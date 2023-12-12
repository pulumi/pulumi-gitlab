// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectTagsTagCommit {
    private String authorEmail;
    private String authorName;
    private String authoredDate;
    private String committedDate;
    private String committerEmail;
    private String committerName;
    private String id;
    private String message;
    private List<String> parentIds;
    private String shortId;
    private String title;

    private GetProjectTagsTagCommit() {}
    public String authorEmail() {
        return this.authorEmail;
    }
    public String authorName() {
        return this.authorName;
    }
    public String authoredDate() {
        return this.authoredDate;
    }
    public String committedDate() {
        return this.committedDate;
    }
    public String committerEmail() {
        return this.committerEmail;
    }
    public String committerName() {
        return this.committerName;
    }
    public String id() {
        return this.id;
    }
    public String message() {
        return this.message;
    }
    public List<String> parentIds() {
        return this.parentIds;
    }
    public String shortId() {
        return this.shortId;
    }
    public String title() {
        return this.title;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectTagsTagCommit defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String authorEmail;
        private String authorName;
        private String authoredDate;
        private String committedDate;
        private String committerEmail;
        private String committerName;
        private String id;
        private String message;
        private List<String> parentIds;
        private String shortId;
        private String title;
        public Builder() {}
        public Builder(GetProjectTagsTagCommit defaults) {
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
        public Builder authorEmail(String authorEmail) {
            this.authorEmail = Objects.requireNonNull(authorEmail);
            return this;
        }
        @CustomType.Setter
        public Builder authorName(String authorName) {
            this.authorName = Objects.requireNonNull(authorName);
            return this;
        }
        @CustomType.Setter
        public Builder authoredDate(String authoredDate) {
            this.authoredDate = Objects.requireNonNull(authoredDate);
            return this;
        }
        @CustomType.Setter
        public Builder committedDate(String committedDate) {
            this.committedDate = Objects.requireNonNull(committedDate);
            return this;
        }
        @CustomType.Setter
        public Builder committerEmail(String committerEmail) {
            this.committerEmail = Objects.requireNonNull(committerEmail);
            return this;
        }
        @CustomType.Setter
        public Builder committerName(String committerName) {
            this.committerName = Objects.requireNonNull(committerName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder message(String message) {
            this.message = Objects.requireNonNull(message);
            return this;
        }
        @CustomType.Setter
        public Builder parentIds(List<String> parentIds) {
            this.parentIds = Objects.requireNonNull(parentIds);
            return this;
        }
        public Builder parentIds(String... parentIds) {
            return parentIds(List.of(parentIds));
        }
        @CustomType.Setter
        public Builder shortId(String shortId) {
            this.shortId = Objects.requireNonNull(shortId);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        public GetProjectTagsTagCommit build() {
            final var _resultValue = new GetProjectTagsTagCommit();
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
