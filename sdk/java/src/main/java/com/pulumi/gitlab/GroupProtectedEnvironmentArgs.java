// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.inputs.GroupProtectedEnvironmentApprovalRuleArgs;
import com.pulumi.gitlab.inputs.GroupProtectedEnvironmentDeployAccessLevelArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupProtectedEnvironmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupProtectedEnvironmentArgs Empty = new GroupProtectedEnvironmentArgs();

    /**
     * Array of approval rules to deploy, with each described by a hash.
     * 
     */
    @Import(name="approvalRules")
    private @Nullable Output<List<GroupProtectedEnvironmentApprovalRuleArgs>> approvalRules;

    /**
     * @return Array of approval rules to deploy, with each described by a hash.
     * 
     */
    public Optional<Output<List<GroupProtectedEnvironmentApprovalRuleArgs>>> approvalRules() {
        return Optional.ofNullable(this.approvalRules);
    }

    /**
     * Array of access levels allowed to deploy, with each described by a hash.
     * 
     */
    @Import(name="deployAccessLevels", required=true)
    private Output<List<GroupProtectedEnvironmentDeployAccessLevelArgs>> deployAccessLevels;

    /**
     * @return Array of access levels allowed to deploy, with each described by a hash.
     * 
     */
    public Output<List<GroupProtectedEnvironmentDeployAccessLevelArgs>> deployAccessLevels() {
        return this.deployAccessLevels;
    }

    /**
     * The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
     * 
     */
    @Import(name="environment", required=true)
    private Output<String> environment;

    /**
     * @return The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }

    /**
     * The ID or full path of the group which the protected environment is created against.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or full path of the group which the protected environment is created against.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    private GroupProtectedEnvironmentArgs() {}

    private GroupProtectedEnvironmentArgs(GroupProtectedEnvironmentArgs $) {
        this.approvalRules = $.approvalRules;
        this.deployAccessLevels = $.deployAccessLevels;
        this.environment = $.environment;
        this.group = $.group;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupProtectedEnvironmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupProtectedEnvironmentArgs $;

        public Builder() {
            $ = new GroupProtectedEnvironmentArgs();
        }

        public Builder(GroupProtectedEnvironmentArgs defaults) {
            $ = new GroupProtectedEnvironmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param approvalRules Array of approval rules to deploy, with each described by a hash.
         * 
         * @return builder
         * 
         */
        public Builder approvalRules(@Nullable Output<List<GroupProtectedEnvironmentApprovalRuleArgs>> approvalRules) {
            $.approvalRules = approvalRules;
            return this;
        }

        /**
         * @param approvalRules Array of approval rules to deploy, with each described by a hash.
         * 
         * @return builder
         * 
         */
        public Builder approvalRules(List<GroupProtectedEnvironmentApprovalRuleArgs> approvalRules) {
            return approvalRules(Output.of(approvalRules));
        }

        /**
         * @param approvalRules Array of approval rules to deploy, with each described by a hash.
         * 
         * @return builder
         * 
         */
        public Builder approvalRules(GroupProtectedEnvironmentApprovalRuleArgs... approvalRules) {
            return approvalRules(List.of(approvalRules));
        }

        /**
         * @param deployAccessLevels Array of access levels allowed to deploy, with each described by a hash.
         * 
         * @return builder
         * 
         */
        public Builder deployAccessLevels(Output<List<GroupProtectedEnvironmentDeployAccessLevelArgs>> deployAccessLevels) {
            $.deployAccessLevels = deployAccessLevels;
            return this;
        }

        /**
         * @param deployAccessLevels Array of access levels allowed to deploy, with each described by a hash.
         * 
         * @return builder
         * 
         */
        public Builder deployAccessLevels(List<GroupProtectedEnvironmentDeployAccessLevelArgs> deployAccessLevels) {
            return deployAccessLevels(Output.of(deployAccessLevels));
        }

        /**
         * @param deployAccessLevels Array of access levels allowed to deploy, with each described by a hash.
         * 
         * @return builder
         * 
         */
        public Builder deployAccessLevels(GroupProtectedEnvironmentDeployAccessLevelArgs... deployAccessLevels) {
            return deployAccessLevels(List.of(deployAccessLevels));
        }

        /**
         * @param environment The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
         * 
         * @return builder
         * 
         */
        public Builder environment(Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment The deployment tier of the environment.  Valid values are `production`, `staging`, `testing`, `development`, `other`.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param group The ID or full path of the group which the protected environment is created against.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or full path of the group which the protected environment is created against.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        public GroupProtectedEnvironmentArgs build() {
            if ($.deployAccessLevels == null) {
                throw new MissingRequiredPropertyException("GroupProtectedEnvironmentArgs", "deployAccessLevels");
            }
            if ($.environment == null) {
                throw new MissingRequiredPropertyException("GroupProtectedEnvironmentArgs", "environment");
            }
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GroupProtectedEnvironmentArgs", "group");
            }
            return $;
        }
    }

}
