// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetGroupSubgroupsSubgroup {
    private Boolean autoDevopsEnabled;
    private String avatarUrl;
    private String createdAt;
    private Integer defaultBranchProtection;
    private String description;
    private Boolean emailsDisabled;
    private Integer fileTemplateProjectId;
    private String fullName;
    private String fullPath;
    /**
     * @return The ID of the group.
     * 
     */
    private Integer groupId;
    private String ipRestrictionRanges;
    private Boolean lfsEnabled;
    private Boolean mentionsDisabled;
    private String name;
    private Integer parentId;
    private String path;
    private String projectCreationLevel;
    private Boolean requestAccessEnabled;
    private Boolean requireTwoFactorAuthentication;
    private Boolean shareWithGroupLock;
    /**
     * @return Include group statistics (administrators only).
     * 
     */
    private Map<String,String> statistics;
    private String subgroupCreationLevel;
    private Integer twoFactorGracePeriod;
    private String visibility;
    private String webUrl;

    private GetGroupSubgroupsSubgroup() {}
    public Boolean autoDevopsEnabled() {
        return this.autoDevopsEnabled;
    }
    public String avatarUrl() {
        return this.avatarUrl;
    }
    public String createdAt() {
        return this.createdAt;
    }
    public Integer defaultBranchProtection() {
        return this.defaultBranchProtection;
    }
    public String description() {
        return this.description;
    }
    public Boolean emailsDisabled() {
        return this.emailsDisabled;
    }
    public Integer fileTemplateProjectId() {
        return this.fileTemplateProjectId;
    }
    public String fullName() {
        return this.fullName;
    }
    public String fullPath() {
        return this.fullPath;
    }
    /**
     * @return The ID of the group.
     * 
     */
    public Integer groupId() {
        return this.groupId;
    }
    public String ipRestrictionRanges() {
        return this.ipRestrictionRanges;
    }
    public Boolean lfsEnabled() {
        return this.lfsEnabled;
    }
    public Boolean mentionsDisabled() {
        return this.mentionsDisabled;
    }
    public String name() {
        return this.name;
    }
    public Integer parentId() {
        return this.parentId;
    }
    public String path() {
        return this.path;
    }
    public String projectCreationLevel() {
        return this.projectCreationLevel;
    }
    public Boolean requestAccessEnabled() {
        return this.requestAccessEnabled;
    }
    public Boolean requireTwoFactorAuthentication() {
        return this.requireTwoFactorAuthentication;
    }
    public Boolean shareWithGroupLock() {
        return this.shareWithGroupLock;
    }
    /**
     * @return Include group statistics (administrators only).
     * 
     */
    public Map<String,String> statistics() {
        return this.statistics;
    }
    public String subgroupCreationLevel() {
        return this.subgroupCreationLevel;
    }
    public Integer twoFactorGracePeriod() {
        return this.twoFactorGracePeriod;
    }
    public String visibility() {
        return this.visibility;
    }
    public String webUrl() {
        return this.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupSubgroupsSubgroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean autoDevopsEnabled;
        private String avatarUrl;
        private String createdAt;
        private Integer defaultBranchProtection;
        private String description;
        private Boolean emailsDisabled;
        private Integer fileTemplateProjectId;
        private String fullName;
        private String fullPath;
        private Integer groupId;
        private String ipRestrictionRanges;
        private Boolean lfsEnabled;
        private Boolean mentionsDisabled;
        private String name;
        private Integer parentId;
        private String path;
        private String projectCreationLevel;
        private Boolean requestAccessEnabled;
        private Boolean requireTwoFactorAuthentication;
        private Boolean shareWithGroupLock;
        private Map<String,String> statistics;
        private String subgroupCreationLevel;
        private Integer twoFactorGracePeriod;
        private String visibility;
        private String webUrl;
        public Builder() {}
        public Builder(GetGroupSubgroupsSubgroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoDevopsEnabled = defaults.autoDevopsEnabled;
    	      this.avatarUrl = defaults.avatarUrl;
    	      this.createdAt = defaults.createdAt;
    	      this.defaultBranchProtection = defaults.defaultBranchProtection;
    	      this.description = defaults.description;
    	      this.emailsDisabled = defaults.emailsDisabled;
    	      this.fileTemplateProjectId = defaults.fileTemplateProjectId;
    	      this.fullName = defaults.fullName;
    	      this.fullPath = defaults.fullPath;
    	      this.groupId = defaults.groupId;
    	      this.ipRestrictionRanges = defaults.ipRestrictionRanges;
    	      this.lfsEnabled = defaults.lfsEnabled;
    	      this.mentionsDisabled = defaults.mentionsDisabled;
    	      this.name = defaults.name;
    	      this.parentId = defaults.parentId;
    	      this.path = defaults.path;
    	      this.projectCreationLevel = defaults.projectCreationLevel;
    	      this.requestAccessEnabled = defaults.requestAccessEnabled;
    	      this.requireTwoFactorAuthentication = defaults.requireTwoFactorAuthentication;
    	      this.shareWithGroupLock = defaults.shareWithGroupLock;
    	      this.statistics = defaults.statistics;
    	      this.subgroupCreationLevel = defaults.subgroupCreationLevel;
    	      this.twoFactorGracePeriod = defaults.twoFactorGracePeriod;
    	      this.visibility = defaults.visibility;
    	      this.webUrl = defaults.webUrl;
        }

        @CustomType.Setter
        public Builder autoDevopsEnabled(Boolean autoDevopsEnabled) {
            this.autoDevopsEnabled = Objects.requireNonNull(autoDevopsEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder avatarUrl(String avatarUrl) {
            this.avatarUrl = Objects.requireNonNull(avatarUrl);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder defaultBranchProtection(Integer defaultBranchProtection) {
            this.defaultBranchProtection = Objects.requireNonNull(defaultBranchProtection);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder emailsDisabled(Boolean emailsDisabled) {
            this.emailsDisabled = Objects.requireNonNull(emailsDisabled);
            return this;
        }
        @CustomType.Setter
        public Builder fileTemplateProjectId(Integer fileTemplateProjectId) {
            this.fileTemplateProjectId = Objects.requireNonNull(fileTemplateProjectId);
            return this;
        }
        @CustomType.Setter
        public Builder fullName(String fullName) {
            this.fullName = Objects.requireNonNull(fullName);
            return this;
        }
        @CustomType.Setter
        public Builder fullPath(String fullPath) {
            this.fullPath = Objects.requireNonNull(fullPath);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(Integer groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder ipRestrictionRanges(String ipRestrictionRanges) {
            this.ipRestrictionRanges = Objects.requireNonNull(ipRestrictionRanges);
            return this;
        }
        @CustomType.Setter
        public Builder lfsEnabled(Boolean lfsEnabled) {
            this.lfsEnabled = Objects.requireNonNull(lfsEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder mentionsDisabled(Boolean mentionsDisabled) {
            this.mentionsDisabled = Objects.requireNonNull(mentionsDisabled);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder parentId(Integer parentId) {
            this.parentId = Objects.requireNonNull(parentId);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder projectCreationLevel(String projectCreationLevel) {
            this.projectCreationLevel = Objects.requireNonNull(projectCreationLevel);
            return this;
        }
        @CustomType.Setter
        public Builder requestAccessEnabled(Boolean requestAccessEnabled) {
            this.requestAccessEnabled = Objects.requireNonNull(requestAccessEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder requireTwoFactorAuthentication(Boolean requireTwoFactorAuthentication) {
            this.requireTwoFactorAuthentication = Objects.requireNonNull(requireTwoFactorAuthentication);
            return this;
        }
        @CustomType.Setter
        public Builder shareWithGroupLock(Boolean shareWithGroupLock) {
            this.shareWithGroupLock = Objects.requireNonNull(shareWithGroupLock);
            return this;
        }
        @CustomType.Setter
        public Builder statistics(Map<String,String> statistics) {
            this.statistics = Objects.requireNonNull(statistics);
            return this;
        }
        @CustomType.Setter
        public Builder subgroupCreationLevel(String subgroupCreationLevel) {
            this.subgroupCreationLevel = Objects.requireNonNull(subgroupCreationLevel);
            return this;
        }
        @CustomType.Setter
        public Builder twoFactorGracePeriod(Integer twoFactorGracePeriod) {
            this.twoFactorGracePeriod = Objects.requireNonNull(twoFactorGracePeriod);
            return this;
        }
        @CustomType.Setter
        public Builder visibility(String visibility) {
            this.visibility = Objects.requireNonNull(visibility);
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            this.webUrl = Objects.requireNonNull(webUrl);
            return this;
        }
        public GetGroupSubgroupsSubgroup build() {
            final var o = new GetGroupSubgroupsSubgroup();
            o.autoDevopsEnabled = autoDevopsEnabled;
            o.avatarUrl = avatarUrl;
            o.createdAt = createdAt;
            o.defaultBranchProtection = defaultBranchProtection;
            o.description = description;
            o.emailsDisabled = emailsDisabled;
            o.fileTemplateProjectId = fileTemplateProjectId;
            o.fullName = fullName;
            o.fullPath = fullPath;
            o.groupId = groupId;
            o.ipRestrictionRanges = ipRestrictionRanges;
            o.lfsEnabled = lfsEnabled;
            o.mentionsDisabled = mentionsDisabled;
            o.name = name;
            o.parentId = parentId;
            o.path = path;
            o.projectCreationLevel = projectCreationLevel;
            o.requestAccessEnabled = requestAccessEnabled;
            o.requireTwoFactorAuthentication = requireTwoFactorAuthentication;
            o.shareWithGroupLock = shareWithGroupLock;
            o.statistics = statistics;
            o.subgroupCreationLevel = subgroupCreationLevel;
            o.twoFactorGracePeriod = twoFactorGracePeriod;
            o.visibility = visibility;
            o.webUrl = webUrl;
            return o;
        }
    }
}
