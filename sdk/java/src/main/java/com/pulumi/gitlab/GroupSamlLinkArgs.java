// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GroupSamlLinkArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupSamlLinkArgs Empty = new GroupSamlLinkArgs();

    /**
     * Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    @Import(name="accessLevel", required=true)
    private Output<String> accessLevel;

    /**
     * @return Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }

    /**
     * The ID or path of the group to add the SAML Group Link to.
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return The ID or path of the group to add the SAML Group Link to.
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * The name of the SAML group.
     * 
     */
    @Import(name="samlGroupName", required=true)
    private Output<String> samlGroupName;

    /**
     * @return The name of the SAML group.
     * 
     */
    public Output<String> samlGroupName() {
        return this.samlGroupName;
    }

    private GroupSamlLinkArgs() {}

    private GroupSamlLinkArgs(GroupSamlLinkArgs $) {
        this.accessLevel = $.accessLevel;
        this.group = $.group;
        this.samlGroupName = $.samlGroupName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupSamlLinkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupSamlLinkArgs $;

        public Builder() {
            $ = new GroupSamlLinkArgs();
        }

        public Builder(GroupSamlLinkArgs defaults) {
            $ = new GroupSamlLinkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel Access level for members of the SAML group. Valid values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param group The ID or path of the group to add the SAML Group Link to.
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or path of the group to add the SAML Group Link to.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param samlGroupName The name of the SAML group.
         * 
         * @return builder
         * 
         */
        public Builder samlGroupName(Output<String> samlGroupName) {
            $.samlGroupName = samlGroupName;
            return this;
        }

        /**
         * @param samlGroupName The name of the SAML group.
         * 
         * @return builder
         * 
         */
        public Builder samlGroupName(String samlGroupName) {
            return samlGroupName(Output.of(samlGroupName));
        }

        public GroupSamlLinkArgs build() {
            if ($.accessLevel == null) {
                throw new MissingRequiredPropertyException("GroupSamlLinkArgs", "accessLevel");
            }
            if ($.group == null) {
                throw new MissingRequiredPropertyException("GroupSamlLinkArgs", "group");
            }
            if ($.samlGroupName == null) {
                throw new MissingRequiredPropertyException("GroupSamlLinkArgs", "samlGroupName");
            }
            return $;
        }
    }

}
