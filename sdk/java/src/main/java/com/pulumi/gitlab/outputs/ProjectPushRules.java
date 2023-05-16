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
public final class ProjectPushRules {
    /**
     * @return All commit author emails must match this regex, e.g. `@my-company.com$`.
     * 
     */
    private @Nullable String authorEmailRegex;
    /**
     * @return All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     * 
     */
    private @Nullable String branchNameRegex;
    /**
     * @return Users can only push commits to this repository that were committed with one of their own verified emails.
     * 
     */
    private @Nullable Boolean commitCommitterCheck;
    /**
     * @return No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     * 
     */
    private @Nullable String commitMessageNegativeRegex;
    /**
     * @return All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     * 
     */
    private @Nullable String commitMessageRegex;
    /**
     * @return Deny deleting a tag.
     * 
     */
    private @Nullable Boolean denyDeleteTag;
    /**
     * @return All committed filenames must not match this regex, e.g. `(jar|exe)$`.
     * 
     */
    private @Nullable String fileNameRegex;
    /**
     * @return Maximum file size (MB).
     * 
     */
    private @Nullable Integer maxFileSize;
    /**
     * @return Restrict commits by author (email) to existing GitLab users.
     * 
     */
    private @Nullable Boolean memberCheck;
    /**
     * @return GitLab will reject any files that are likely to contain secrets.
     * 
     */
    private @Nullable Boolean preventSecrets;
    /**
     * @return Reject commit when it’s not signed through GPG.
     * 
     */
    private @Nullable Boolean rejectUnsignedCommits;

    private ProjectPushRules() {}
    /**
     * @return All commit author emails must match this regex, e.g. `@my-company.com$`.
     * 
     */
    public Optional<String> authorEmailRegex() {
        return Optional.ofNullable(this.authorEmailRegex);
    }
    /**
     * @return All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     * 
     */
    public Optional<String> branchNameRegex() {
        return Optional.ofNullable(this.branchNameRegex);
    }
    /**
     * @return Users can only push commits to this repository that were committed with one of their own verified emails.
     * 
     */
    public Optional<Boolean> commitCommitterCheck() {
        return Optional.ofNullable(this.commitCommitterCheck);
    }
    /**
     * @return No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     * 
     */
    public Optional<String> commitMessageNegativeRegex() {
        return Optional.ofNullable(this.commitMessageNegativeRegex);
    }
    /**
     * @return All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     * 
     */
    public Optional<String> commitMessageRegex() {
        return Optional.ofNullable(this.commitMessageRegex);
    }
    /**
     * @return Deny deleting a tag.
     * 
     */
    public Optional<Boolean> denyDeleteTag() {
        return Optional.ofNullable(this.denyDeleteTag);
    }
    /**
     * @return All committed filenames must not match this regex, e.g. `(jar|exe)$`.
     * 
     */
    public Optional<String> fileNameRegex() {
        return Optional.ofNullable(this.fileNameRegex);
    }
    /**
     * @return Maximum file size (MB).
     * 
     */
    public Optional<Integer> maxFileSize() {
        return Optional.ofNullable(this.maxFileSize);
    }
    /**
     * @return Restrict commits by author (email) to existing GitLab users.
     * 
     */
    public Optional<Boolean> memberCheck() {
        return Optional.ofNullable(this.memberCheck);
    }
    /**
     * @return GitLab will reject any files that are likely to contain secrets.
     * 
     */
    public Optional<Boolean> preventSecrets() {
        return Optional.ofNullable(this.preventSecrets);
    }
    /**
     * @return Reject commit when it’s not signed through GPG.
     * 
     */
    public Optional<Boolean> rejectUnsignedCommits() {
        return Optional.ofNullable(this.rejectUnsignedCommits);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProjectPushRules defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authorEmailRegex;
        private @Nullable String branchNameRegex;
        private @Nullable Boolean commitCommitterCheck;
        private @Nullable String commitMessageNegativeRegex;
        private @Nullable String commitMessageRegex;
        private @Nullable Boolean denyDeleteTag;
        private @Nullable String fileNameRegex;
        private @Nullable Integer maxFileSize;
        private @Nullable Boolean memberCheck;
        private @Nullable Boolean preventSecrets;
        private @Nullable Boolean rejectUnsignedCommits;
        public Builder() {}
        public Builder(ProjectPushRules defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorEmailRegex = defaults.authorEmailRegex;
    	      this.branchNameRegex = defaults.branchNameRegex;
    	      this.commitCommitterCheck = defaults.commitCommitterCheck;
    	      this.commitMessageNegativeRegex = defaults.commitMessageNegativeRegex;
    	      this.commitMessageRegex = defaults.commitMessageRegex;
    	      this.denyDeleteTag = defaults.denyDeleteTag;
    	      this.fileNameRegex = defaults.fileNameRegex;
    	      this.maxFileSize = defaults.maxFileSize;
    	      this.memberCheck = defaults.memberCheck;
    	      this.preventSecrets = defaults.preventSecrets;
    	      this.rejectUnsignedCommits = defaults.rejectUnsignedCommits;
        }

        @CustomType.Setter
        public Builder authorEmailRegex(@Nullable String authorEmailRegex) {
            this.authorEmailRegex = authorEmailRegex;
            return this;
        }
        @CustomType.Setter
        public Builder branchNameRegex(@Nullable String branchNameRegex) {
            this.branchNameRegex = branchNameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder commitCommitterCheck(@Nullable Boolean commitCommitterCheck) {
            this.commitCommitterCheck = commitCommitterCheck;
            return this;
        }
        @CustomType.Setter
        public Builder commitMessageNegativeRegex(@Nullable String commitMessageNegativeRegex) {
            this.commitMessageNegativeRegex = commitMessageNegativeRegex;
            return this;
        }
        @CustomType.Setter
        public Builder commitMessageRegex(@Nullable String commitMessageRegex) {
            this.commitMessageRegex = commitMessageRegex;
            return this;
        }
        @CustomType.Setter
        public Builder denyDeleteTag(@Nullable Boolean denyDeleteTag) {
            this.denyDeleteTag = denyDeleteTag;
            return this;
        }
        @CustomType.Setter
        public Builder fileNameRegex(@Nullable String fileNameRegex) {
            this.fileNameRegex = fileNameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder maxFileSize(@Nullable Integer maxFileSize) {
            this.maxFileSize = maxFileSize;
            return this;
        }
        @CustomType.Setter
        public Builder memberCheck(@Nullable Boolean memberCheck) {
            this.memberCheck = memberCheck;
            return this;
        }
        @CustomType.Setter
        public Builder preventSecrets(@Nullable Boolean preventSecrets) {
            this.preventSecrets = preventSecrets;
            return this;
        }
        @CustomType.Setter
        public Builder rejectUnsignedCommits(@Nullable Boolean rejectUnsignedCommits) {
            this.rejectUnsignedCommits = rejectUnsignedCommits;
            return this;
        }
        public ProjectPushRules build() {
            final var o = new ProjectPushRules();
            o.authorEmailRegex = authorEmailRegex;
            o.branchNameRegex = branchNameRegex;
            o.commitCommitterCheck = commitCommitterCheck;
            o.commitMessageNegativeRegex = commitMessageNegativeRegex;
            o.commitMessageRegex = commitMessageRegex;
            o.denyDeleteTag = denyDeleteTag;
            o.fileNameRegex = fileNameRegex;
            o.maxFileSize = maxFileSize;
            o.memberCheck = memberCheck;
            o.preventSecrets = preventSecrets;
            o.rejectUnsignedCommits = rejectUnsignedCommits;
            return o;
        }
    }
}
