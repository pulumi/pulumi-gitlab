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


public final class RepositoryFileArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryFileArgs Empty = new RepositoryFileArgs();

    /**
     * Email of the commit author.
     * 
     */
    @Import(name="authorEmail")
    private @Nullable Output<String> authorEmail;

    /**
     * @return Email of the commit author.
     * 
     */
    public Optional<Output<String>> authorEmail() {
        return Optional.ofNullable(this.authorEmail);
    }

    /**
     * Name of the commit author.
     * 
     */
    @Import(name="authorName")
    private @Nullable Output<String> authorName;

    /**
     * @return Name of the commit author.
     * 
     */
    public Optional<Output<String>> authorName() {
        return Optional.ofNullable(this.authorName);
    }

    /**
     * Name of the branch to which to commit to.
     * 
     */
    @Import(name="branch", required=true)
    private Output<String> branch;

    /**
     * @return Name of the branch to which to commit to.
     * 
     */
    public Output<String> branch() {
        return this.branch;
    }

    /**
     * Commit message.
     * 
     */
    @Import(name="commitMessage")
    private @Nullable Output<String> commitMessage;

    /**
     * @return Commit message.
     * 
     */
    public Optional<Output<String>> commitMessage() {
        return Optional.ofNullable(this.commitMessage);
    }

    /**
     * File content.
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return File content.
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * Create commit message.
     * 
     */
    @Import(name="createCommitMessage")
    private @Nullable Output<String> createCommitMessage;

    /**
     * @return Create commit message.
     * 
     */
    public Optional<Output<String>> createCommitMessage() {
        return Optional.ofNullable(this.createCommitMessage);
    }

    /**
     * Delete Commit message.
     * 
     */
    @Import(name="deleteCommitMessage")
    private @Nullable Output<String> deleteCommitMessage;

    /**
     * @return Delete Commit message.
     * 
     */
    public Optional<Output<String>> deleteCommitMessage() {
        return Optional.ofNullable(this.deleteCommitMessage);
    }

    /**
     * The file content encoding. Valid values are: `base64`, `text`.
     * 
     */
    @Import(name="encoding", required=true)
    private Output<String> encoding;

    /**
     * @return The file content encoding. Valid values are: `base64`, `text`.
     * 
     */
    public Output<String> encoding() {
        return this.encoding;
    }

    /**
     * Enables or disables the execute flag on the file.
     * 
     */
    @Import(name="executeFilemode")
    private @Nullable Output<Boolean> executeFilemode;

    /**
     * @return Enables or disables the execute flag on the file.
     * 
     */
    public Optional<Output<Boolean>> executeFilemode() {
        return Optional.ofNullable(this.executeFilemode);
    }

    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     * 
     */
    @Import(name="filePath", required=true)
    private Output<String> filePath;

    /**
     * @return The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     * 
     */
    public Output<String> filePath() {
        return this.filePath;
    }

    /**
     * Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     * 
     */
    @Import(name="overwriteOnCreate")
    private @Nullable Output<Boolean> overwriteOnCreate;

    /**
     * @return Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     * 
     */
    public Optional<Output<Boolean>> overwriteOnCreate() {
        return Optional.ofNullable(this.overwriteOnCreate);
    }

    /**
     * The name or ID of the project.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The name or ID of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Name of the branch to start the new commit from.
     * 
     */
    @Import(name="startBranch")
    private @Nullable Output<String> startBranch;

    /**
     * @return Name of the branch to start the new commit from.
     * 
     */
    public Optional<Output<String>> startBranch() {
        return Optional.ofNullable(this.startBranch);
    }

    /**
     * Update commit message.
     * 
     */
    @Import(name="updateCommitMessage")
    private @Nullable Output<String> updateCommitMessage;

    /**
     * @return Update commit message.
     * 
     */
    public Optional<Output<String>> updateCommitMessage() {
        return Optional.ofNullable(this.updateCommitMessage);
    }

    private RepositoryFileArgs() {}

    private RepositoryFileArgs(RepositoryFileArgs $) {
        this.authorEmail = $.authorEmail;
        this.authorName = $.authorName;
        this.branch = $.branch;
        this.commitMessage = $.commitMessage;
        this.content = $.content;
        this.createCommitMessage = $.createCommitMessage;
        this.deleteCommitMessage = $.deleteCommitMessage;
        this.encoding = $.encoding;
        this.executeFilemode = $.executeFilemode;
        this.filePath = $.filePath;
        this.overwriteOnCreate = $.overwriteOnCreate;
        this.project = $.project;
        this.startBranch = $.startBranch;
        this.updateCommitMessage = $.updateCommitMessage;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryFileArgs $;

        public Builder() {
            $ = new RepositoryFileArgs();
        }

        public Builder(RepositoryFileArgs defaults) {
            $ = new RepositoryFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorEmail Email of the commit author.
         * 
         * @return builder
         * 
         */
        public Builder authorEmail(@Nullable Output<String> authorEmail) {
            $.authorEmail = authorEmail;
            return this;
        }

        /**
         * @param authorEmail Email of the commit author.
         * 
         * @return builder
         * 
         */
        public Builder authorEmail(String authorEmail) {
            return authorEmail(Output.of(authorEmail));
        }

        /**
         * @param authorName Name of the commit author.
         * 
         * @return builder
         * 
         */
        public Builder authorName(@Nullable Output<String> authorName) {
            $.authorName = authorName;
            return this;
        }

        /**
         * @param authorName Name of the commit author.
         * 
         * @return builder
         * 
         */
        public Builder authorName(String authorName) {
            return authorName(Output.of(authorName));
        }

        /**
         * @param branch Name of the branch to which to commit to.
         * 
         * @return builder
         * 
         */
        public Builder branch(Output<String> branch) {
            $.branch = branch;
            return this;
        }

        /**
         * @param branch Name of the branch to which to commit to.
         * 
         * @return builder
         * 
         */
        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        /**
         * @param commitMessage Commit message.
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(@Nullable Output<String> commitMessage) {
            $.commitMessage = commitMessage;
            return this;
        }

        /**
         * @param commitMessage Commit message.
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(String commitMessage) {
            return commitMessage(Output.of(commitMessage));
        }

        /**
         * @param content File content.
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content File content.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param createCommitMessage Create commit message.
         * 
         * @return builder
         * 
         */
        public Builder createCommitMessage(@Nullable Output<String> createCommitMessage) {
            $.createCommitMessage = createCommitMessage;
            return this;
        }

        /**
         * @param createCommitMessage Create commit message.
         * 
         * @return builder
         * 
         */
        public Builder createCommitMessage(String createCommitMessage) {
            return createCommitMessage(Output.of(createCommitMessage));
        }

        /**
         * @param deleteCommitMessage Delete Commit message.
         * 
         * @return builder
         * 
         */
        public Builder deleteCommitMessage(@Nullable Output<String> deleteCommitMessage) {
            $.deleteCommitMessage = deleteCommitMessage;
            return this;
        }

        /**
         * @param deleteCommitMessage Delete Commit message.
         * 
         * @return builder
         * 
         */
        public Builder deleteCommitMessage(String deleteCommitMessage) {
            return deleteCommitMessage(Output.of(deleteCommitMessage));
        }

        /**
         * @param encoding The file content encoding. Valid values are: `base64`, `text`.
         * 
         * @return builder
         * 
         */
        public Builder encoding(Output<String> encoding) {
            $.encoding = encoding;
            return this;
        }

        /**
         * @param encoding The file content encoding. Valid values are: `base64`, `text`.
         * 
         * @return builder
         * 
         */
        public Builder encoding(String encoding) {
            return encoding(Output.of(encoding));
        }

        /**
         * @param executeFilemode Enables or disables the execute flag on the file.
         * 
         * @return builder
         * 
         */
        public Builder executeFilemode(@Nullable Output<Boolean> executeFilemode) {
            $.executeFilemode = executeFilemode;
            return this;
        }

        /**
         * @param executeFilemode Enables or disables the execute flag on the file.
         * 
         * @return builder
         * 
         */
        public Builder executeFilemode(Boolean executeFilemode) {
            return executeFilemode(Output.of(executeFilemode));
        }

        /**
         * @param filePath The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
         * 
         * @return builder
         * 
         */
        public Builder filePath(Output<String> filePath) {
            $.filePath = filePath;
            return this;
        }

        /**
         * @param filePath The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
         * 
         * @return builder
         * 
         */
        public Builder filePath(String filePath) {
            return filePath(Output.of(filePath));
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(@Nullable Output<Boolean> overwriteOnCreate) {
            $.overwriteOnCreate = overwriteOnCreate;
            return this;
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(Boolean overwriteOnCreate) {
            return overwriteOnCreate(Output.of(overwriteOnCreate));
        }

        /**
         * @param project The name or ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param startBranch Name of the branch to start the new commit from.
         * 
         * @return builder
         * 
         */
        public Builder startBranch(@Nullable Output<String> startBranch) {
            $.startBranch = startBranch;
            return this;
        }

        /**
         * @param startBranch Name of the branch to start the new commit from.
         * 
         * @return builder
         * 
         */
        public Builder startBranch(String startBranch) {
            return startBranch(Output.of(startBranch));
        }

        /**
         * @param updateCommitMessage Update commit message.
         * 
         * @return builder
         * 
         */
        public Builder updateCommitMessage(@Nullable Output<String> updateCommitMessage) {
            $.updateCommitMessage = updateCommitMessage;
            return this;
        }

        /**
         * @param updateCommitMessage Update commit message.
         * 
         * @return builder
         * 
         */
        public Builder updateCommitMessage(String updateCommitMessage) {
            return updateCommitMessage(Output.of(updateCommitMessage));
        }

        public RepositoryFileArgs build() {
            if ($.branch == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "branch");
            }
            if ($.content == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "content");
            }
            if ($.encoding == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "encoding");
            }
            if ($.filePath == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "filePath");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "project");
            }
            return $;
        }
    }

}
