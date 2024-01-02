// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBranchCommit {
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

    private GetBranchCommit() {}
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

    public static Builder builder(GetBranchCommit defaults) {
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
        public Builder(GetBranchCommit defaults) {
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
            if (authorEmail == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "authorEmail");
            }
            this.authorEmail = authorEmail;
            return this;
        }
        @CustomType.Setter
        public Builder authorName(String authorName) {
            if (authorName == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "authorName");
            }
            this.authorName = authorName;
            return this;
        }
        @CustomType.Setter
        public Builder authoredDate(String authoredDate) {
            if (authoredDate == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "authoredDate");
            }
            this.authoredDate = authoredDate;
            return this;
        }
        @CustomType.Setter
        public Builder committedDate(String committedDate) {
            if (committedDate == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "committedDate");
            }
            this.committedDate = committedDate;
            return this;
        }
        @CustomType.Setter
        public Builder committerEmail(String committerEmail) {
            if (committerEmail == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "committerEmail");
            }
            this.committerEmail = committerEmail;
            return this;
        }
        @CustomType.Setter
        public Builder committerName(String committerName) {
            if (committerName == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "committerName");
            }
            this.committerName = committerName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder message(String message) {
            if (message == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "message");
            }
            this.message = message;
            return this;
        }
        @CustomType.Setter
        public Builder parentIds(List<String> parentIds) {
            if (parentIds == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "parentIds");
            }
            this.parentIds = parentIds;
            return this;
        }
        public Builder parentIds(String... parentIds) {
            return parentIds(List.of(parentIds));
        }
        @CustomType.Setter
        public Builder shortId(String shortId) {
            if (shortId == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "shortId");
            }
            this.shortId = shortId;
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            if (title == null) {
              throw new MissingRequiredPropertyException("GetBranchCommit", "title");
            }
            this.title = title;
            return this;
        }
        public GetBranchCommit build() {
            final var _resultValue = new GetBranchCommit();
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
