// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MemberRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final MemberRoleArgs Empty = new MemberRoleArgs();

    /**
     * The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
     * 
     */
    @Import(name="baseAccessLevel", required=true)
    private Output<String> baseAccessLevel;

    /**
     * @return The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
     * 
     */
    public Output<String> baseAccessLevel() {
        return this.baseAccessLevel;
    }

    /**
     * Description for the member role.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the member role.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
     * 
     */
    @Import(name="enabledPermissions", required=true)
    private Output<List<String>> enabledPermissions;

    /**
     * @return All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
     * 
     */
    public Output<List<String>> enabledPermissions() {
        return this.enabledPermissions;
    }

    /**
     * Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
     * 
     */
    @Import(name="groupPath")
    private @Nullable Output<String> groupPath;

    /**
     * @return Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
     * 
     */
    public Optional<Output<String>> groupPath() {
        return Optional.ofNullable(this.groupPath);
    }

    /**
     * Name for the member role.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name for the member role.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private MemberRoleArgs() {}

    private MemberRoleArgs(MemberRoleArgs $) {
        this.baseAccessLevel = $.baseAccessLevel;
        this.description = $.description;
        this.enabledPermissions = $.enabledPermissions;
        this.groupPath = $.groupPath;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MemberRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MemberRoleArgs $;

        public Builder() {
            $ = new MemberRoleArgs();
        }

        public Builder(MemberRoleArgs defaults) {
            $ = new MemberRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param baseAccessLevel The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
         * 
         * @return builder
         * 
         */
        public Builder baseAccessLevel(Output<String> baseAccessLevel) {
            $.baseAccessLevel = baseAccessLevel;
            return this;
        }

        /**
         * @param baseAccessLevel The base access level for the custom role. Valid values are: `DEVELOPER`, `GUEST`, `MAINTAINER`, `MINIMAL_ACCESS`, `OWNER`, `REPORTER`
         * 
         * @return builder
         * 
         */
        public Builder baseAccessLevel(String baseAccessLevel) {
            return baseAccessLevel(Output.of(baseAccessLevel));
        }

        /**
         * @param description Description for the member role.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the member role.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enabledPermissions All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
         * 
         * @return builder
         * 
         */
        public Builder enabledPermissions(Output<List<String>> enabledPermissions) {
            $.enabledPermissions = enabledPermissions;
            return this;
        }

        /**
         * @param enabledPermissions All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
         * 
         * @return builder
         * 
         */
        public Builder enabledPermissions(List<String> enabledPermissions) {
            return enabledPermissions(Output.of(enabledPermissions));
        }

        /**
         * @param enabledPermissions All permissions enabled for the custom role. Valid values are: `ADMIN_CICD_VARIABLES`, `ADMIN_COMPLIANCE_FRAMEWORK`, `ADMIN_GROUP_MEMBER`, `ADMIN_INTEGRATIONS`, `ADMIN_MERGE_REQUEST`, `ADMIN_PUSH_RULES`, `ADMIN_RUNNERS`, `ADMIN_TERRAFORM_STATE`, `ADMIN_VULNERABILITY`, `ADMIN_WEB_HOOK`, `ARCHIVE_PROJECT`, `MANAGE_DEPLOY_TOKENS`, `MANAGE_GROUP_ACCESS_TOKENS`, `MANAGE_MERGE_REQUEST_SETTINGS`, `MANAGE_PROJECT_ACCESS_TOKENS`, `MANAGE_SECURITY_POLICY_LINK`, `READ_CODE`, `READ_CRM_CONTACT`, `READ_DEPENDENCY`, `READ_RUNNERS`, `READ_VULNERABILITY`, `REMOVE_GROUP`, `REMOVE_PROJECT`
         * 
         * @return builder
         * 
         */
        public Builder enabledPermissions(String... enabledPermissions) {
            return enabledPermissions(List.of(enabledPermissions));
        }

        /**
         * @param groupPath Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
         * 
         * @return builder
         * 
         */
        public Builder groupPath(@Nullable Output<String> groupPath) {
            $.groupPath = groupPath;
            return this;
        }

        /**
         * @param groupPath Full path of the namespace to create the member role in. **Required for SAAS** **Not allowed for self-managed**
         * 
         * @return builder
         * 
         */
        public Builder groupPath(String groupPath) {
            return groupPath(Output.of(groupPath));
        }

        /**
         * @param name Name for the member role.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name for the member role.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public MemberRoleArgs build() {
            if ($.baseAccessLevel == null) {
                throw new MissingRequiredPropertyException("MemberRoleArgs", "baseAccessLevel");
            }
            if ($.enabledPermissions == null) {
                throw new MissingRequiredPropertyException("MemberRoleArgs", "enabledPermissions");
            }
            return $;
        }
    }

}
