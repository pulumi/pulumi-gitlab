// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupServiceAccountPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupServiceAccountPlainArgs Empty = new GetGroupServiceAccountPlainArgs();

    /**
     * The ID or URL-encoded path of the target group. Must be a top-level group.
     * 
     */
    @Import(name="group", required=true)
    private String group;

    /**
     * @return The ID or URL-encoded path of the target group. Must be a top-level group.
     * 
     */
    public String group() {
        return this.group;
    }

    /**
     * The name of the user. If not specified, the default Service account user name is used.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the user. If not specified, the default Service account user name is used.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The service account id.
     * 
     */
    @Import(name="serviceAccountId", required=true)
    private String serviceAccountId;

    /**
     * @return The service account id.
     * 
     */
    public String serviceAccountId() {
        return this.serviceAccountId;
    }

    /**
     * The username of the user. If not specified, it&#39;s automatically generated.
     * 
     */
    @Import(name="username")
    private @Nullable String username;

    /**
     * @return The username of the user. If not specified, it&#39;s automatically generated.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    private GetGroupServiceAccountPlainArgs() {}

    private GetGroupServiceAccountPlainArgs(GetGroupServiceAccountPlainArgs $) {
        this.group = $.group;
        this.name = $.name;
        this.serviceAccountId = $.serviceAccountId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupServiceAccountPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupServiceAccountPlainArgs $;

        public Builder() {
            $ = new GetGroupServiceAccountPlainArgs();
        }

        public Builder(GetGroupServiceAccountPlainArgs defaults) {
            $ = new GetGroupServiceAccountPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param group The ID or URL-encoded path of the target group. Must be a top-level group.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            $.group = group;
            return this;
        }

        /**
         * @param name The name of the user. If not specified, the default Service account user name is used.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param serviceAccountId The service account id.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountId(String serviceAccountId) {
            $.serviceAccountId = serviceAccountId;
            return this;
        }

        /**
         * @param username The username of the user. If not specified, it&#39;s automatically generated.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable String username) {
            $.username = username;
            return this;
        }

        public GetGroupServiceAccountPlainArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GetGroupServiceAccountPlainArgs", "group");
            }
            if ($.serviceAccountId == null) {
                throw new MissingRequiredPropertyException("GetGroupServiceAccountPlainArgs", "serviceAccountId");
            }
            return $;
        }
    }

}