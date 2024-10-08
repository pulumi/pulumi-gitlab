// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectPushRule {
    /**
     * @return All commit author emails must match this regex, e.g. `{@literal @}my-company.com$`.
     * 
     */
    private String authorEmailRegex;
    /**
     * @return All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     * 
     */
    private String branchNameRegex;
    /**
     * @return Users can only push commits to this repository that were committed with one of their own verified emails.
     * 
     */
    private Boolean commitCommitterCheck;
    /**
     * @return Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
     * 
     */
    private Boolean commitCommitterNameCheck;
    /**
     * @return No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     * 
     */
    private String commitMessageNegativeRegex;
    /**
     * @return All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     * 
     */
    private String commitMessageRegex;
    /**
     * @return Deny deleting a tag.
     * 
     */
    private Boolean denyDeleteTag;
    /**
     * @return All committed filenames must not match this regex, e.g. `(jar|exe)$`.
     * 
     */
    private String fileNameRegex;
    /**
     * @return Maximum file size (MB).
     * 
     */
    private Integer maxFileSize;
    /**
     * @return Restrict commits by author (email) to existing GitLab users.
     * 
     */
    private Boolean memberCheck;
    /**
     * @return GitLab will reject any files that are likely to contain secrets.
     * 
     */
    private Boolean preventSecrets;
    /**
     * @return Reject commit when it’s not DCO certified.
     * 
     */
    private Boolean rejectNonDcoCommits;
    /**
     * @return Reject commit when it’s not signed through GPG.
     * 
     */
    private Boolean rejectUnsignedCommits;

    private GetProjectPushRule() {}
    /**
     * @return All commit author emails must match this regex, e.g. `{@literal @}my-company.com$`.
     * 
     */
    public String authorEmailRegex() {
        return this.authorEmailRegex;
    }
    /**
     * @return All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     * 
     */
    public String branchNameRegex() {
        return this.branchNameRegex;
    }
    /**
     * @return Users can only push commits to this repository that were committed with one of their own verified emails.
     * 
     */
    public Boolean commitCommitterCheck() {
        return this.commitCommitterCheck;
    }
    /**
     * @return Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
     * 
     */
    public Boolean commitCommitterNameCheck() {
        return this.commitCommitterNameCheck;
    }
    /**
     * @return No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     * 
     */
    public String commitMessageNegativeRegex() {
        return this.commitMessageNegativeRegex;
    }
    /**
     * @return All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     * 
     */
    public String commitMessageRegex() {
        return this.commitMessageRegex;
    }
    /**
     * @return Deny deleting a tag.
     * 
     */
    public Boolean denyDeleteTag() {
        return this.denyDeleteTag;
    }
    /**
     * @return All committed filenames must not match this regex, e.g. `(jar|exe)$`.
     * 
     */
    public String fileNameRegex() {
        return this.fileNameRegex;
    }
    /**
     * @return Maximum file size (MB).
     * 
     */
    public Integer maxFileSize() {
        return this.maxFileSize;
    }
    /**
     * @return Restrict commits by author (email) to existing GitLab users.
     * 
     */
    public Boolean memberCheck() {
        return this.memberCheck;
    }
    /**
     * @return GitLab will reject any files that are likely to contain secrets.
     * 
     */
    public Boolean preventSecrets() {
        return this.preventSecrets;
    }
    /**
     * @return Reject commit when it’s not DCO certified.
     * 
     */
    public Boolean rejectNonDcoCommits() {
        return this.rejectNonDcoCommits;
    }
    /**
     * @return Reject commit when it’s not signed through GPG.
     * 
     */
    public Boolean rejectUnsignedCommits() {
        return this.rejectUnsignedCommits;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectPushRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String authorEmailRegex;
        private String branchNameRegex;
        private Boolean commitCommitterCheck;
        private Boolean commitCommitterNameCheck;
        private String commitMessageNegativeRegex;
        private String commitMessageRegex;
        private Boolean denyDeleteTag;
        private String fileNameRegex;
        private Integer maxFileSize;
        private Boolean memberCheck;
        private Boolean preventSecrets;
        private Boolean rejectNonDcoCommits;
        private Boolean rejectUnsignedCommits;
        public Builder() {}
        public Builder(GetProjectPushRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorEmailRegex = defaults.authorEmailRegex;
    	      this.branchNameRegex = defaults.branchNameRegex;
    	      this.commitCommitterCheck = defaults.commitCommitterCheck;
    	      this.commitCommitterNameCheck = defaults.commitCommitterNameCheck;
    	      this.commitMessageNegativeRegex = defaults.commitMessageNegativeRegex;
    	      this.commitMessageRegex = defaults.commitMessageRegex;
    	      this.denyDeleteTag = defaults.denyDeleteTag;
    	      this.fileNameRegex = defaults.fileNameRegex;
    	      this.maxFileSize = defaults.maxFileSize;
    	      this.memberCheck = defaults.memberCheck;
    	      this.preventSecrets = defaults.preventSecrets;
    	      this.rejectNonDcoCommits = defaults.rejectNonDcoCommits;
    	      this.rejectUnsignedCommits = defaults.rejectUnsignedCommits;
        }

        @CustomType.Setter
        public Builder authorEmailRegex(String authorEmailRegex) {
            if (authorEmailRegex == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "authorEmailRegex");
            }
            this.authorEmailRegex = authorEmailRegex;
            return this;
        }
        @CustomType.Setter
        public Builder branchNameRegex(String branchNameRegex) {
            if (branchNameRegex == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "branchNameRegex");
            }
            this.branchNameRegex = branchNameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder commitCommitterCheck(Boolean commitCommitterCheck) {
            if (commitCommitterCheck == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "commitCommitterCheck");
            }
            this.commitCommitterCheck = commitCommitterCheck;
            return this;
        }
        @CustomType.Setter
        public Builder commitCommitterNameCheck(Boolean commitCommitterNameCheck) {
            if (commitCommitterNameCheck == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "commitCommitterNameCheck");
            }
            this.commitCommitterNameCheck = commitCommitterNameCheck;
            return this;
        }
        @CustomType.Setter
        public Builder commitMessageNegativeRegex(String commitMessageNegativeRegex) {
            if (commitMessageNegativeRegex == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "commitMessageNegativeRegex");
            }
            this.commitMessageNegativeRegex = commitMessageNegativeRegex;
            return this;
        }
        @CustomType.Setter
        public Builder commitMessageRegex(String commitMessageRegex) {
            if (commitMessageRegex == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "commitMessageRegex");
            }
            this.commitMessageRegex = commitMessageRegex;
            return this;
        }
        @CustomType.Setter
        public Builder denyDeleteTag(Boolean denyDeleteTag) {
            if (denyDeleteTag == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "denyDeleteTag");
            }
            this.denyDeleteTag = denyDeleteTag;
            return this;
        }
        @CustomType.Setter
        public Builder fileNameRegex(String fileNameRegex) {
            if (fileNameRegex == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "fileNameRegex");
            }
            this.fileNameRegex = fileNameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder maxFileSize(Integer maxFileSize) {
            if (maxFileSize == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "maxFileSize");
            }
            this.maxFileSize = maxFileSize;
            return this;
        }
        @CustomType.Setter
        public Builder memberCheck(Boolean memberCheck) {
            if (memberCheck == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "memberCheck");
            }
            this.memberCheck = memberCheck;
            return this;
        }
        @CustomType.Setter
        public Builder preventSecrets(Boolean preventSecrets) {
            if (preventSecrets == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "preventSecrets");
            }
            this.preventSecrets = preventSecrets;
            return this;
        }
        @CustomType.Setter
        public Builder rejectNonDcoCommits(Boolean rejectNonDcoCommits) {
            if (rejectNonDcoCommits == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "rejectNonDcoCommits");
            }
            this.rejectNonDcoCommits = rejectNonDcoCommits;
            return this;
        }
        @CustomType.Setter
        public Builder rejectUnsignedCommits(Boolean rejectUnsignedCommits) {
            if (rejectUnsignedCommits == null) {
              throw new MissingRequiredPropertyException("GetProjectPushRule", "rejectUnsignedCommits");
            }
            this.rejectUnsignedCommits = rejectUnsignedCommits;
            return this;
        }
        public GetProjectPushRule build() {
            final var _resultValue = new GetProjectPushRule();
            _resultValue.authorEmailRegex = authorEmailRegex;
            _resultValue.branchNameRegex = branchNameRegex;
            _resultValue.commitCommitterCheck = commitCommitterCheck;
            _resultValue.commitCommitterNameCheck = commitCommitterNameCheck;
            _resultValue.commitMessageNegativeRegex = commitMessageNegativeRegex;
            _resultValue.commitMessageRegex = commitMessageRegex;
            _resultValue.denyDeleteTag = denyDeleteTag;
            _resultValue.fileNameRegex = fileNameRegex;
            _resultValue.maxFileSize = maxFileSize;
            _resultValue.memberCheck = memberCheck;
            _resultValue.preventSecrets = preventSecrets;
            _resultValue.rejectNonDcoCommits = rejectNonDcoCommits;
            _resultValue.rejectUnsignedCommits = rejectUnsignedCommits;
            return _resultValue;
        }
    }
}
