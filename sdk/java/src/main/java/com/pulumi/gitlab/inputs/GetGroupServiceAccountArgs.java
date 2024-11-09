// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupServiceAccountArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupServiceAccountArgs Empty = new GetGroupServiceAccountArgs();

    /**
     * The ID or URL-encoded path of the target group. Must be a top-level group.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the target group. Must be a top-level group.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * The name of the user. If not specified, the default Service account user name is used.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the user. If not specified, the default Service account user name is used.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The service account id.
     * 
     */
    @Import(name="serviceAccountId", required=true)
    private Output<String> serviceAccountId;

    /**
     * @return The service account id.
     * 
     */
    public Output<String> serviceAccountId() {
        return this.serviceAccountId;
    }

    /**
     * The username of the user. If not specified, it&#39;s automatically generated.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username of the user. If not specified, it&#39;s automatically generated.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private GetGroupServiceAccountArgs() {}

    private GetGroupServiceAccountArgs(GetGroupServiceAccountArgs $) {
        this.group = $.group;
        this.name = $.name;
        this.serviceAccountId = $.serviceAccountId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupServiceAccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupServiceAccountArgs $;

        public Builder() {
            $ = new GetGroupServiceAccountArgs();
        }

        public Builder(GetGroupServiceAccountArgs defaults) {
            $ = new GetGroupServiceAccountArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param group The ID or URL-encoded path of the target group. Must be a top-level group.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or URL-encoded path of the target group. Must be a top-level group.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param name The name of the user. If not specified, the default Service account user name is used.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the user. If not specified, the default Service account user name is used.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param serviceAccountId The service account id.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(Output<String> serviceAccountId) {
            $.serviceAccountId = serviceAccountId;
            return this;
        }

        /**
         * @param serviceAccountId The service account id.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            return serviceAccountId(Output.of(serviceAccountId));
        }

        /**
         * @param username The username of the user. If not specified, it&#39;s automatically generated.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the user. If not specified, it&#39;s automatically generated.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GetGroupServiceAccountArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GetGroupServiceAccountArgs", "group");
            }
            if ($.serviceAccountId == null) {
                throw new MissingRequiredPropertyException("GetGroupServiceAccountArgs", "serviceAccountId");
            }
            return $;
        }
    }

}