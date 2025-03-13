// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupPushRulesArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupPushRulesArgs Empty = new GroupPushRulesArgs();

    /**
     * All commit author emails must match this regex, e.g. `{@literal @}my-company.com$`.
     * 
     */
    @Import(name="authorEmailRegex")
    private @Nullable Output<String> authorEmailRegex;

    /**
     * @return All commit author emails must match this regex, e.g. `{@literal @}my-company.com$`.
     * 
     */
    public Optional<Output<String>> authorEmailRegex() {
        return Optional.ofNullable(this.authorEmailRegex);
    }

    /**
     * All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     * 
     */
    @Import(name="branchNameRegex")
    private @Nullable Output<String> branchNameRegex;

    /**
     * @return All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     * 
     */
    public Optional<Output<String>> branchNameRegex() {
        return Optional.ofNullable(this.branchNameRegex);
    }

    /**
     * Only commits pushed using verified emails are allowed.
     * 
     */
    @Import(name="commitCommitterCheck")
    private @Nullable Output<Boolean> commitCommitterCheck;

    /**
     * @return Only commits pushed using verified emails are allowed.
     * 
     */
    public Optional<Output<Boolean>> commitCommitterCheck() {
        return Optional.ofNullable(this.commitCommitterCheck);
    }

    /**
     * Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
     * 
     */
    @Import(name="commitCommitterNameCheck")
    private @Nullable Output<Boolean> commitCommitterNameCheck;

    /**
     * @return Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
     * 
     */
    public Optional<Output<Boolean>> commitCommitterNameCheck() {
        return Optional.ofNullable(this.commitCommitterNameCheck);
    }

    /**
     * No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     * 
     */
    @Import(name="commitMessageNegativeRegex")
    private @Nullable Output<String> commitMessageNegativeRegex;

    /**
     * @return No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     * 
     */
    public Optional<Output<String>> commitMessageNegativeRegex() {
        return Optional.ofNullable(this.commitMessageNegativeRegex);
    }

    /**
     * All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     * 
     */
    @Import(name="commitMessageRegex")
    private @Nullable Output<String> commitMessageRegex;

    /**
     * @return All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     * 
     */
    public Optional<Output<String>> commitMessageRegex() {
        return Optional.ofNullable(this.commitMessageRegex);
    }

    /**
     * Deny deleting a tag.
     * 
     */
    @Import(name="denyDeleteTag")
    private @Nullable Output<Boolean> denyDeleteTag;

    /**
     * @return Deny deleting a tag.
     * 
     */
    public Optional<Output<Boolean>> denyDeleteTag() {
        return Optional.ofNullable(this.denyDeleteTag);
    }

    /**
     * Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
     * 
     */
    @Import(name="fileNameRegex")
    private @Nullable Output<String> fileNameRegex;

    /**
     * @return Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
     * 
     */
    public Optional<Output<String>> fileNameRegex() {
        return Optional.ofNullable(this.fileNameRegex);
    }

    /**
     * Maximum file size (MB) allowed.
     * 
     */
    @Import(name="maxFileSize")
    private @Nullable Output<Integer> maxFileSize;

    /**
     * @return Maximum file size (MB) allowed.
     * 
     */
    public Optional<Output<Integer>> maxFileSize() {
        return Optional.ofNullable(this.maxFileSize);
    }

    /**
     * Allows only GitLab users to author commits.
     * 
     */
    @Import(name="memberCheck")
    private @Nullable Output<Boolean> memberCheck;

    /**
     * @return Allows only GitLab users to author commits.
     * 
     */
    public Optional<Output<Boolean>> memberCheck() {
        return Optional.ofNullable(this.memberCheck);
    }

    /**
     * GitLab will reject any files that are likely to contain secrets.
     * 
     */
    @Import(name="preventSecrets")
    private @Nullable Output<Boolean> preventSecrets;

    /**
     * @return GitLab will reject any files that are likely to contain secrets.
     * 
     */
    public Optional<Output<Boolean>> preventSecrets() {
        return Optional.ofNullable(this.preventSecrets);
    }

    /**
     * Reject commit when it’s not DCO certified.
     * 
     */
    @Import(name="rejectNonDcoCommits")
    private @Nullable Output<Boolean> rejectNonDcoCommits;

    /**
     * @return Reject commit when it’s not DCO certified.
     * 
     */
    public Optional<Output<Boolean>> rejectNonDcoCommits() {
        return Optional.ofNullable(this.rejectNonDcoCommits);
    }

    /**
     * Only commits signed through GPG are allowed.
     * 
     */
    @Import(name="rejectUnsignedCommits")
    private @Nullable Output<Boolean> rejectUnsignedCommits;

    /**
     * @return Only commits signed through GPG are allowed.
     * 
     */
    public Optional<Output<Boolean>> rejectUnsignedCommits() {
        return Optional.ofNullable(this.rejectUnsignedCommits);
    }

    private GroupPushRulesArgs() {}

    private GroupPushRulesArgs(GroupPushRulesArgs $) {
        this.authorEmailRegex = $.authorEmailRegex;
        this.branchNameRegex = $.branchNameRegex;
        this.commitCommitterCheck = $.commitCommitterCheck;
        this.commitCommitterNameCheck = $.commitCommitterNameCheck;
        this.commitMessageNegativeRegex = $.commitMessageNegativeRegex;
        this.commitMessageRegex = $.commitMessageRegex;
        this.denyDeleteTag = $.denyDeleteTag;
        this.fileNameRegex = $.fileNameRegex;
        this.maxFileSize = $.maxFileSize;
        this.memberCheck = $.memberCheck;
        this.preventSecrets = $.preventSecrets;
        this.rejectNonDcoCommits = $.rejectNonDcoCommits;
        this.rejectUnsignedCommits = $.rejectUnsignedCommits;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupPushRulesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupPushRulesArgs $;

        public Builder() {
            $ = new GroupPushRulesArgs();
        }

        public Builder(GroupPushRulesArgs defaults) {
            $ = new GroupPushRulesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorEmailRegex All commit author emails must match this regex, e.g. `{@literal @}my-company.com$`.
         * 
         * @return builder
         * 
         */
        public Builder authorEmailRegex(@Nullable Output<String> authorEmailRegex) {
            $.authorEmailRegex = authorEmailRegex;
            return this;
        }

        /**
         * @param authorEmailRegex All commit author emails must match this regex, e.g. `{@literal @}my-company.com$`.
         * 
         * @return builder
         * 
         */
        public Builder authorEmailRegex(String authorEmailRegex) {
            return authorEmailRegex(Output.of(authorEmailRegex));
        }

        /**
         * @param branchNameRegex All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
         * 
         * @return builder
         * 
         */
        public Builder branchNameRegex(@Nullable Output<String> branchNameRegex) {
            $.branchNameRegex = branchNameRegex;
            return this;
        }

        /**
         * @param branchNameRegex All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
         * 
         * @return builder
         * 
         */
        public Builder branchNameRegex(String branchNameRegex) {
            return branchNameRegex(Output.of(branchNameRegex));
        }

        /**
         * @param commitCommitterCheck Only commits pushed using verified emails are allowed.
         * 
         * @return builder
         * 
         */
        public Builder commitCommitterCheck(@Nullable Output<Boolean> commitCommitterCheck) {
            $.commitCommitterCheck = commitCommitterCheck;
            return this;
        }

        /**
         * @param commitCommitterCheck Only commits pushed using verified emails are allowed.
         * 
         * @return builder
         * 
         */
        public Builder commitCommitterCheck(Boolean commitCommitterCheck) {
            return commitCommitterCheck(Output.of(commitCommitterCheck));
        }

        /**
         * @param commitCommitterNameCheck Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
         * 
         * @return builder
         * 
         */
        public Builder commitCommitterNameCheck(@Nullable Output<Boolean> commitCommitterNameCheck) {
            $.commitCommitterNameCheck = commitCommitterNameCheck;
            return this;
        }

        /**
         * @param commitCommitterNameCheck Users can only push commits to this repository if the commit author name is consistent with their GitLab account name.
         * 
         * @return builder
         * 
         */
        public Builder commitCommitterNameCheck(Boolean commitCommitterNameCheck) {
            return commitCommitterNameCheck(Output.of(commitCommitterNameCheck));
        }

        /**
         * @param commitMessageNegativeRegex No commit message is allowed to match this regex, for example `ssh\:\/\/`.
         * 
         * @return builder
         * 
         */
        public Builder commitMessageNegativeRegex(@Nullable Output<String> commitMessageNegativeRegex) {
            $.commitMessageNegativeRegex = commitMessageNegativeRegex;
            return this;
        }

        /**
         * @param commitMessageNegativeRegex No commit message is allowed to match this regex, for example `ssh\:\/\/`.
         * 
         * @return builder
         * 
         */
        public Builder commitMessageNegativeRegex(String commitMessageNegativeRegex) {
            return commitMessageNegativeRegex(Output.of(commitMessageNegativeRegex));
        }

        /**
         * @param commitMessageRegex All commit messages must match this regex, e.g. `Fixed \d+\..*`.
         * 
         * @return builder
         * 
         */
        public Builder commitMessageRegex(@Nullable Output<String> commitMessageRegex) {
            $.commitMessageRegex = commitMessageRegex;
            return this;
        }

        /**
         * @param commitMessageRegex All commit messages must match this regex, e.g. `Fixed \d+\..*`.
         * 
         * @return builder
         * 
         */
        public Builder commitMessageRegex(String commitMessageRegex) {
            return commitMessageRegex(Output.of(commitMessageRegex));
        }

        /**
         * @param denyDeleteTag Deny deleting a tag.
         * 
         * @return builder
         * 
         */
        public Builder denyDeleteTag(@Nullable Output<Boolean> denyDeleteTag) {
            $.denyDeleteTag = denyDeleteTag;
            return this;
        }

        /**
         * @param denyDeleteTag Deny deleting a tag.
         * 
         * @return builder
         * 
         */
        public Builder denyDeleteTag(Boolean denyDeleteTag) {
            return denyDeleteTag(Output.of(denyDeleteTag));
        }

        /**
         * @param fileNameRegex Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
         * 
         * @return builder
         * 
         */
        public Builder fileNameRegex(@Nullable Output<String> fileNameRegex) {
            $.fileNameRegex = fileNameRegex;
            return this;
        }

        /**
         * @param fileNameRegex Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
         * 
         * @return builder
         * 
         */
        public Builder fileNameRegex(String fileNameRegex) {
            return fileNameRegex(Output.of(fileNameRegex));
        }

        /**
         * @param maxFileSize Maximum file size (MB) allowed.
         * 
         * @return builder
         * 
         */
        public Builder maxFileSize(@Nullable Output<Integer> maxFileSize) {
            $.maxFileSize = maxFileSize;
            return this;
        }

        /**
         * @param maxFileSize Maximum file size (MB) allowed.
         * 
         * @return builder
         * 
         */
        public Builder maxFileSize(Integer maxFileSize) {
            return maxFileSize(Output.of(maxFileSize));
        }

        /**
         * @param memberCheck Allows only GitLab users to author commits.
         * 
         * @return builder
         * 
         */
        public Builder memberCheck(@Nullable Output<Boolean> memberCheck) {
            $.memberCheck = memberCheck;
            return this;
        }

        /**
         * @param memberCheck Allows only GitLab users to author commits.
         * 
         * @return builder
         * 
         */
        public Builder memberCheck(Boolean memberCheck) {
            return memberCheck(Output.of(memberCheck));
        }

        /**
         * @param preventSecrets GitLab will reject any files that are likely to contain secrets.
         * 
         * @return builder
         * 
         */
        public Builder preventSecrets(@Nullable Output<Boolean> preventSecrets) {
            $.preventSecrets = preventSecrets;
            return this;
        }

        /**
         * @param preventSecrets GitLab will reject any files that are likely to contain secrets.
         * 
         * @return builder
         * 
         */
        public Builder preventSecrets(Boolean preventSecrets) {
            return preventSecrets(Output.of(preventSecrets));
        }

        /**
         * @param rejectNonDcoCommits Reject commit when it’s not DCO certified.
         * 
         * @return builder
         * 
         */
        public Builder rejectNonDcoCommits(@Nullable Output<Boolean> rejectNonDcoCommits) {
            $.rejectNonDcoCommits = rejectNonDcoCommits;
            return this;
        }

        /**
         * @param rejectNonDcoCommits Reject commit when it’s not DCO certified.
         * 
         * @return builder
         * 
         */
        public Builder rejectNonDcoCommits(Boolean rejectNonDcoCommits) {
            return rejectNonDcoCommits(Output.of(rejectNonDcoCommits));
        }

        /**
         * @param rejectUnsignedCommits Only commits signed through GPG are allowed.
         * 
         * @return builder
         * 
         */
        public Builder rejectUnsignedCommits(@Nullable Output<Boolean> rejectUnsignedCommits) {
            $.rejectUnsignedCommits = rejectUnsignedCommits;
            return this;
        }

        /**
         * @param rejectUnsignedCommits Only commits signed through GPG are allowed.
         * 
         * @return builder
         * 
         */
        public Builder rejectUnsignedCommits(Boolean rejectUnsignedCommits) {
            return rejectUnsignedCommits(Output.of(rejectUnsignedCommits));
        }

        public GroupPushRulesArgs build() {
            return $;
        }
    }

}
