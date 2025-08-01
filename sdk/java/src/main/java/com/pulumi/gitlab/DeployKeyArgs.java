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


public final class DeployKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final DeployKeyArgs Empty = new DeployKeyArgs();

    /**
     * Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    @Import(name="canPush")
    private @Nullable Output<Boolean> canPush;

    /**
     * @return Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> canPush() {
        return Optional.ofNullable(this.canPush);
    }

    /**
     * Expiration date for the deploy key. Does not expire if no value is provided. Expected in RFC3339 format `(2019-03-15T08:00:00Z)`
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return Expiration date for the deploy key. Does not expire if no value is provided. Expected in RFC3339 format `(2019-03-15T08:00:00Z)`
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The public ssh key body.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The public ssh key body.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * The name or id of the project to add the deploy key to.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The name or id of the project to add the deploy key to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * A title to describe the deploy key with.
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return A title to describe the deploy key with.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private DeployKeyArgs() {}

    private DeployKeyArgs(DeployKeyArgs $) {
        this.canPush = $.canPush;
        this.expiresAt = $.expiresAt;
        this.key = $.key;
        this.project = $.project;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeployKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeployKeyArgs $;

        public Builder() {
            $ = new DeployKeyArgs();
        }

        public Builder(DeployKeyArgs defaults) {
            $ = new DeployKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param canPush Allow this deploy key to be used to push changes to the project. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder canPush(@Nullable Output<Boolean> canPush) {
            $.canPush = canPush;
            return this;
        }

        /**
         * @param canPush Allow this deploy key to be used to push changes to the project. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder canPush(Boolean canPush) {
            return canPush(Output.of(canPush));
        }

        /**
         * @param expiresAt Expiration date for the deploy key. Does not expire if no value is provided. Expected in RFC3339 format `(2019-03-15T08:00:00Z)`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt Expiration date for the deploy key. Does not expire if no value is provided. Expected in RFC3339 format `(2019-03-15T08:00:00Z)`
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param key The public ssh key body.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The public ssh key body.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param project The name or id of the project to add the deploy key to.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project to add the deploy key to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param title A title to describe the deploy key with.
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title A title to describe the deploy key with.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public DeployKeyArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("DeployKeyArgs", "key");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("DeployKeyArgs", "project");
            }
            if ($.title == null) {
                throw new MissingRequiredPropertyException("DeployKeyArgs", "title");
            }
            return $;
        }
    }

}
