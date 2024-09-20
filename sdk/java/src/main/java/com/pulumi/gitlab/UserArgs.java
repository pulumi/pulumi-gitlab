// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    /**
     * Boolean, defaults to false. Whether to allow the user to create groups.
     * 
     */
    @Import(name="canCreateGroup")
    private @Nullable Output<Boolean> canCreateGroup;

    /**
     * @return Boolean, defaults to false. Whether to allow the user to create groups.
     * 
     */
    public Optional<Output<Boolean>> canCreateGroup() {
        return Optional.ofNullable(this.canCreateGroup);
    }

    /**
     * The e-mail address of the user.
     * 
     */
    @Import(name="email", required=true)
    private Output<String> email;

    /**
     * @return The e-mail address of the user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }

    /**
     * String, a specific external authentication provider UID.
     * 
     */
    @Import(name="externUid")
    private @Nullable Output<String> externUid;

    /**
     * @return String, a specific external authentication provider UID.
     * 
     */
    public Optional<Output<String>> externUid() {
        return Optional.ofNullable(this.externUid);
    }

    /**
     * String, the external provider.
     * 
     */
    @Import(name="externalProvider")
    private @Nullable Output<String> externalProvider;

    /**
     * @return String, the external provider.
     * 
     */
    public Optional<Output<String>> externalProvider() {
        return Optional.ofNullable(this.externalProvider);
    }

    /**
     * Boolean, defaults to false.  Whether to enable administrative privileges
     * 
     */
    @Import(name="isAdmin")
    private @Nullable Output<Boolean> isAdmin;

    /**
     * @return Boolean, defaults to false.  Whether to enable administrative privileges
     * 
     */
    public Optional<Output<Boolean>> isAdmin() {
        return Optional.ofNullable(this.isAdmin);
    }

    /**
     * Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     * 
     */
    @Import(name="isExternal")
    private @Nullable Output<Boolean> isExternal;

    /**
     * @return Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
     * 
     */
    public Optional<Output<Boolean>> isExternal() {
        return Optional.ofNullable(this.isExternal);
    }

    /**
     * The name of the user.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the user.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the user&#39;s namespace. Available since GitLab 14.10.
     * 
     */
    @Import(name="namespaceId")
    private @Nullable Output<Integer> namespaceId;

    /**
     * @return The ID of the user&#39;s namespace. Available since GitLab 14.10.
     * 
     */
    public Optional<Output<Integer>> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }

    /**
     * The note associated to the user.
     * 
     */
    @Import(name="note")
    private @Nullable Output<String> note;

    /**
     * @return The note associated to the user.
     * 
     */
    public Optional<Output<String>> note() {
        return Optional.ofNullable(this.note);
    }

    /**
     * The password of the user.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password of the user.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Integer, defaults to 0.  Number of projects user can create.
     * 
     */
    @Import(name="projectsLimit")
    private @Nullable Output<Integer> projectsLimit;

    /**
     * @return Integer, defaults to 0.  Number of projects user can create.
     * 
     */
    public Optional<Output<Integer>> projectsLimit() {
        return Optional.ofNullable(this.projectsLimit);
    }

    /**
     * Boolean, defaults to false. Send user password reset link.
     * 
     */
    @Import(name="resetPassword")
    private @Nullable Output<Boolean> resetPassword;

    /**
     * @return Boolean, defaults to false. Send user password reset link.
     * 
     */
    public Optional<Output<Boolean>> resetPassword() {
        return Optional.ofNullable(this.resetPassword);
    }

    /**
     * Boolean, defaults to true. Whether to skip confirmation.
     * 
     */
    @Import(name="skipConfirmation")
    private @Nullable Output<Boolean> skipConfirmation;

    /**
     * @return Boolean, defaults to true. Whether to skip confirmation.
     * 
     */
    public Optional<Output<Boolean>> skipConfirmation() {
        return Optional.ofNullable(this.skipConfirmation);
    }

    /**
     * String, defaults to &#39;active&#39;. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return String, defaults to &#39;active&#39;. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * The username of the user.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the user.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private UserArgs() {}

    private UserArgs(UserArgs $) {
        this.canCreateGroup = $.canCreateGroup;
        this.email = $.email;
        this.externUid = $.externUid;
        this.externalProvider = $.externalProvider;
        this.isAdmin = $.isAdmin;
        this.isExternal = $.isExternal;
        this.name = $.name;
        this.namespaceId = $.namespaceId;
        this.note = $.note;
        this.password = $.password;
        this.projectsLimit = $.projectsLimit;
        this.resetPassword = $.resetPassword;
        this.skipConfirmation = $.skipConfirmation;
        this.state = $.state;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserArgs $;

        public Builder() {
            $ = new UserArgs();
        }

        public Builder(UserArgs defaults) {
            $ = new UserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param canCreateGroup Boolean, defaults to false. Whether to allow the user to create groups.
         * 
         * @return builder
         * 
         */
        public Builder canCreateGroup(@Nullable Output<Boolean> canCreateGroup) {
            $.canCreateGroup = canCreateGroup;
            return this;
        }

        /**
         * @param canCreateGroup Boolean, defaults to false. Whether to allow the user to create groups.
         * 
         * @return builder
         * 
         */
        public Builder canCreateGroup(Boolean canCreateGroup) {
            return canCreateGroup(Output.of(canCreateGroup));
        }

        /**
         * @param email The e-mail address of the user.
         * 
         * @return builder
         * 
         */
        public Builder email(Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The e-mail address of the user.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param externUid String, a specific external authentication provider UID.
         * 
         * @return builder
         * 
         */
        public Builder externUid(@Nullable Output<String> externUid) {
            $.externUid = externUid;
            return this;
        }

        /**
         * @param externUid String, a specific external authentication provider UID.
         * 
         * @return builder
         * 
         */
        public Builder externUid(String externUid) {
            return externUid(Output.of(externUid));
        }

        /**
         * @param externalProvider String, the external provider.
         * 
         * @return builder
         * 
         */
        public Builder externalProvider(@Nullable Output<String> externalProvider) {
            $.externalProvider = externalProvider;
            return this;
        }

        /**
         * @param externalProvider String, the external provider.
         * 
         * @return builder
         * 
         */
        public Builder externalProvider(String externalProvider) {
            return externalProvider(Output.of(externalProvider));
        }

        /**
         * @param isAdmin Boolean, defaults to false.  Whether to enable administrative privileges
         * 
         * @return builder
         * 
         */
        public Builder isAdmin(@Nullable Output<Boolean> isAdmin) {
            $.isAdmin = isAdmin;
            return this;
        }

        /**
         * @param isAdmin Boolean, defaults to false.  Whether to enable administrative privileges
         * 
         * @return builder
         * 
         */
        public Builder isAdmin(Boolean isAdmin) {
            return isAdmin(Output.of(isAdmin));
        }

        /**
         * @param isExternal Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
         * 
         * @return builder
         * 
         */
        public Builder isExternal(@Nullable Output<Boolean> isExternal) {
            $.isExternal = isExternal;
            return this;
        }

        /**
         * @param isExternal Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.
         * 
         * @return builder
         * 
         */
        public Builder isExternal(Boolean isExternal) {
            return isExternal(Output.of(isExternal));
        }

        /**
         * @param name The name of the user.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespaceId The ID of the user&#39;s namespace. Available since GitLab 14.10.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(@Nullable Output<Integer> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId The ID of the user&#39;s namespace. Available since GitLab 14.10.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(Integer namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param note The note associated to the user.
         * 
         * @return builder
         * 
         */
        public Builder note(@Nullable Output<String> note) {
            $.note = note;
            return this;
        }

        /**
         * @param note The note associated to the user.
         * 
         * @return builder
         * 
         */
        public Builder note(String note) {
            return note(Output.of(note));
        }

        /**
         * @param password The password of the user.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the user.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param projectsLimit Integer, defaults to 0.  Number of projects user can create.
         * 
         * @return builder
         * 
         */
        public Builder projectsLimit(@Nullable Output<Integer> projectsLimit) {
            $.projectsLimit = projectsLimit;
            return this;
        }

        /**
         * @param projectsLimit Integer, defaults to 0.  Number of projects user can create.
         * 
         * @return builder
         * 
         */
        public Builder projectsLimit(Integer projectsLimit) {
            return projectsLimit(Output.of(projectsLimit));
        }

        /**
         * @param resetPassword Boolean, defaults to false. Send user password reset link.
         * 
         * @return builder
         * 
         */
        public Builder resetPassword(@Nullable Output<Boolean> resetPassword) {
            $.resetPassword = resetPassword;
            return this;
        }

        /**
         * @param resetPassword Boolean, defaults to false. Send user password reset link.
         * 
         * @return builder
         * 
         */
        public Builder resetPassword(Boolean resetPassword) {
            return resetPassword(Output.of(resetPassword));
        }

        /**
         * @param skipConfirmation Boolean, defaults to true. Whether to skip confirmation.
         * 
         * @return builder
         * 
         */
        public Builder skipConfirmation(@Nullable Output<Boolean> skipConfirmation) {
            $.skipConfirmation = skipConfirmation;
            return this;
        }

        /**
         * @param skipConfirmation Boolean, defaults to true. Whether to skip confirmation.
         * 
         * @return builder
         * 
         */
        public Builder skipConfirmation(Boolean skipConfirmation) {
            return skipConfirmation(Output.of(skipConfirmation));
        }

        /**
         * @param state String, defaults to &#39;active&#39;. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state String, defaults to &#39;active&#39;. The state of the user account. Valid values are `active`, `deactivated`, `blocked`.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param username The username of the user.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the user.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public UserArgs build() {
            if ($.email == null) {
                throw new MissingRequiredPropertyException("UserArgs", "email");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("UserArgs", "username");
            }
            return $;
        }
    }

}
