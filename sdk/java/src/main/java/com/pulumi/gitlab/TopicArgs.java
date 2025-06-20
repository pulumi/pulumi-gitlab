// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TopicArgs extends com.pulumi.resources.ResourceArgs {

    public static final TopicArgs Empty = new TopicArgs();

    /**
     * A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    @Import(name="avatar")
    private @Nullable Output<String> avatar;

    /**
     * @return A local path to the avatar image to upload. **Note**: not available for imported resources.
     * 
     */
    public Optional<Output<String>> avatar() {
        return Optional.ofNullable(this.avatar);
    }

    /**
     * The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    @Import(name="avatarHash")
    private @Nullable Output<String> avatarHash;

    /**
     * @return The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
     * 
     */
    public Optional<Output<String>> avatarHash() {
        return Optional.ofNullable(this.avatarHash);
    }

    /**
     * A text describing the topic.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A text describing the topic.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The topic&#39;s name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The topic&#39;s name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The topic&#39;s description.
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return The topic&#39;s description.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private TopicArgs() {}

    private TopicArgs(TopicArgs $) {
        this.avatar = $.avatar;
        this.avatarHash = $.avatarHash;
        this.description = $.description;
        this.name = $.name;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TopicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TopicArgs $;

        public Builder() {
            $ = new TopicArgs();
        }

        public Builder(TopicArgs defaults) {
            $ = new TopicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param avatar A local path to the avatar image to upload. **Note**: not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder avatar(@Nullable Output<String> avatar) {
            $.avatar = avatar;
            return this;
        }

        /**
         * @param avatar A local path to the avatar image to upload. **Note**: not available for imported resources.
         * 
         * @return builder
         * 
         */
        public Builder avatar(String avatar) {
            return avatar(Output.of(avatar));
        }

        /**
         * @param avatarHash The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
         * 
         * @return builder
         * 
         */
        public Builder avatarHash(@Nullable Output<String> avatarHash) {
            $.avatarHash = avatarHash;
            return this;
        }

        /**
         * @param avatarHash The hash of the avatar image. Use `filesha256(&#34;path/to/avatar.png&#34;)` whenever possible. **Note**: this is used to trigger an update of the avatar. If it&#39;s not given, but an avatar is given, the avatar will be updated each time.
         * 
         * @return builder
         * 
         */
        public Builder avatarHash(String avatarHash) {
            return avatarHash(Output.of(avatarHash));
        }

        /**
         * @param description A text describing the topic.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A text describing the topic.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The topic&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The topic&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param title The topic&#39;s description.
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title The topic&#39;s description.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public TopicArgs build() {
            if ($.title == null) {
                throw new MissingRequiredPropertyException("TopicArgs", "title");
            }
            return $;
        }
    }

}
