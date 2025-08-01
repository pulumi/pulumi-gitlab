// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetUserResult {
    /**
     * @return The avatar URL of the user.
     * 
     */
    private String avatarUrl;
    /**
     * @return The bio of the user.
     * 
     */
    private String bio;
    /**
     * @return Whether the user can create groups.
     * 
     */
    private Boolean canCreateGroup;
    /**
     * @return Whether the user can create projects.
     * 
     */
    private Boolean canCreateProject;
    /**
     * @return User&#39;s color scheme ID.
     * 
     */
    private Integer colorSchemeId;
    /**
     * @return Date the user was created at.
     * 
     */
    private String createdAt;
    /**
     * @return Current user&#39;s sign-in date.
     * 
     */
    private String currentSignInAt;
    /**
     * @return The public email address of the user.
     * 
     */
    private String email;
    /**
     * @return (Experimental) If true, returns only an exact match. Otherwise, fuzzy matching might return the closest result. If no exact match is available, the data source returns an error.
     * 
     */
    private @Nullable Boolean emailExactMatch;
    /**
     * @return The external UID of the user.
     * 
     */
    private String externUid;
    /**
     * @return Whether the user is external.
     * 
     */
    private Boolean external;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Whether the user is an admin.
     * 
     */
    private Boolean isAdmin;
    /**
     * @return Whether the user is a bot.
     * 
     */
    private Boolean isBot;
    /**
     * @return Last user&#39;s sign-in date.
     * 
     */
    private String lastSignInAt;
    /**
     * @return LinkedIn profile of the user.
     * 
     */
    private String linkedin;
    /**
     * @return The location of the user.
     * 
     */
    private String location;
    /**
     * @return The name of the user.
     * 
     */
    private String name;
    /**
     * @return The ID of the user&#39;s namespace. Requires admin token to access this field.
     * 
     */
    private Integer namespaceId;
    /**
     * @return Admin notes for this user.
     * 
     */
    private String note;
    /**
     * @return The organization of the user.
     * 
     */
    private String organization;
    /**
     * @return Number of projects the user can create.
     * 
     */
    private Integer projectsLimit;
    /**
     * @return Skype username of the user.
     * 
     */
    private String skype;
    /**
     * @return Whether the user is active or blocked.
     * 
     */
    private String state;
    /**
     * @return User&#39;s theme ID.
     * 
     */
    private Integer themeId;
    /**
     * @return Twitter username of the user.
     * 
     */
    private String twitter;
    /**
     * @return Whether user&#39;s two-factor auth is enabled.
     * 
     */
    private Boolean twoFactorEnabled;
    /**
     * @return The ID of the user.
     * 
     */
    private Integer userId;
    /**
     * @return The UID provider of the user.
     * 
     */
    private String userProvider;
    /**
     * @return The username of the user.
     * 
     */
    private String username;
    /**
     * @return User&#39;s website URL.
     * 
     */
    private String websiteUrl;

    private GetUserResult() {}
    /**
     * @return The avatar URL of the user.
     * 
     */
    public String avatarUrl() {
        return this.avatarUrl;
    }
    /**
     * @return The bio of the user.
     * 
     */
    public String bio() {
        return this.bio;
    }
    /**
     * @return Whether the user can create groups.
     * 
     */
    public Boolean canCreateGroup() {
        return this.canCreateGroup;
    }
    /**
     * @return Whether the user can create projects.
     * 
     */
    public Boolean canCreateProject() {
        return this.canCreateProject;
    }
    /**
     * @return User&#39;s color scheme ID.
     * 
     */
    public Integer colorSchemeId() {
        return this.colorSchemeId;
    }
    /**
     * @return Date the user was created at.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Current user&#39;s sign-in date.
     * 
     */
    public String currentSignInAt() {
        return this.currentSignInAt;
    }
    /**
     * @return The public email address of the user.
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return (Experimental) If true, returns only an exact match. Otherwise, fuzzy matching might return the closest result. If no exact match is available, the data source returns an error.
     * 
     */
    public Optional<Boolean> emailExactMatch() {
        return Optional.ofNullable(this.emailExactMatch);
    }
    /**
     * @return The external UID of the user.
     * 
     */
    public String externUid() {
        return this.externUid;
    }
    /**
     * @return Whether the user is external.
     * 
     */
    public Boolean external() {
        return this.external;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether the user is an admin.
     * 
     */
    public Boolean isAdmin() {
        return this.isAdmin;
    }
    /**
     * @return Whether the user is a bot.
     * 
     */
    public Boolean isBot() {
        return this.isBot;
    }
    /**
     * @return Last user&#39;s sign-in date.
     * 
     */
    public String lastSignInAt() {
        return this.lastSignInAt;
    }
    /**
     * @return LinkedIn profile of the user.
     * 
     */
    public String linkedin() {
        return this.linkedin;
    }
    /**
     * @return The location of the user.
     * 
     */
    public String location() {
        return this.location;
    }
    /**
     * @return The name of the user.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The ID of the user&#39;s namespace. Requires admin token to access this field.
     * 
     */
    public Integer namespaceId() {
        return this.namespaceId;
    }
    /**
     * @return Admin notes for this user.
     * 
     */
    public String note() {
        return this.note;
    }
    /**
     * @return The organization of the user.
     * 
     */
    public String organization() {
        return this.organization;
    }
    /**
     * @return Number of projects the user can create.
     * 
     */
    public Integer projectsLimit() {
        return this.projectsLimit;
    }
    /**
     * @return Skype username of the user.
     * 
     */
    public String skype() {
        return this.skype;
    }
    /**
     * @return Whether the user is active or blocked.
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return User&#39;s theme ID.
     * 
     */
    public Integer themeId() {
        return this.themeId;
    }
    /**
     * @return Twitter username of the user.
     * 
     */
    public String twitter() {
        return this.twitter;
    }
    /**
     * @return Whether user&#39;s two-factor auth is enabled.
     * 
     */
    public Boolean twoFactorEnabled() {
        return this.twoFactorEnabled;
    }
    /**
     * @return The ID of the user.
     * 
     */
    public Integer userId() {
        return this.userId;
    }
    /**
     * @return The UID provider of the user.
     * 
     */
    public String userProvider() {
        return this.userProvider;
    }
    /**
     * @return The username of the user.
     * 
     */
    public String username() {
        return this.username;
    }
    /**
     * @return User&#39;s website URL.
     * 
     */
    public String websiteUrl() {
        return this.websiteUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String avatarUrl;
        private String bio;
        private Boolean canCreateGroup;
        private Boolean canCreateProject;
        private Integer colorSchemeId;
        private String createdAt;
        private String currentSignInAt;
        private String email;
        private @Nullable Boolean emailExactMatch;
        private String externUid;
        private Boolean external;
        private String id;
        private Boolean isAdmin;
        private Boolean isBot;
        private String lastSignInAt;
        private String linkedin;
        private String location;
        private String name;
        private Integer namespaceId;
        private String note;
        private String organization;
        private Integer projectsLimit;
        private String skype;
        private String state;
        private Integer themeId;
        private String twitter;
        private Boolean twoFactorEnabled;
        private Integer userId;
        private String userProvider;
        private String username;
        private String websiteUrl;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.avatarUrl = defaults.avatarUrl;
    	      this.bio = defaults.bio;
    	      this.canCreateGroup = defaults.canCreateGroup;
    	      this.canCreateProject = defaults.canCreateProject;
    	      this.colorSchemeId = defaults.colorSchemeId;
    	      this.createdAt = defaults.createdAt;
    	      this.currentSignInAt = defaults.currentSignInAt;
    	      this.email = defaults.email;
    	      this.emailExactMatch = defaults.emailExactMatch;
    	      this.externUid = defaults.externUid;
    	      this.external = defaults.external;
    	      this.id = defaults.id;
    	      this.isAdmin = defaults.isAdmin;
    	      this.isBot = defaults.isBot;
    	      this.lastSignInAt = defaults.lastSignInAt;
    	      this.linkedin = defaults.linkedin;
    	      this.location = defaults.location;
    	      this.name = defaults.name;
    	      this.namespaceId = defaults.namespaceId;
    	      this.note = defaults.note;
    	      this.organization = defaults.organization;
    	      this.projectsLimit = defaults.projectsLimit;
    	      this.skype = defaults.skype;
    	      this.state = defaults.state;
    	      this.themeId = defaults.themeId;
    	      this.twitter = defaults.twitter;
    	      this.twoFactorEnabled = defaults.twoFactorEnabled;
    	      this.userId = defaults.userId;
    	      this.userProvider = defaults.userProvider;
    	      this.username = defaults.username;
    	      this.websiteUrl = defaults.websiteUrl;
        }

        @CustomType.Setter
        public Builder avatarUrl(String avatarUrl) {
            if (avatarUrl == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "avatarUrl");
            }
            this.avatarUrl = avatarUrl;
            return this;
        }
        @CustomType.Setter
        public Builder bio(String bio) {
            if (bio == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "bio");
            }
            this.bio = bio;
            return this;
        }
        @CustomType.Setter
        public Builder canCreateGroup(Boolean canCreateGroup) {
            if (canCreateGroup == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "canCreateGroup");
            }
            this.canCreateGroup = canCreateGroup;
            return this;
        }
        @CustomType.Setter
        public Builder canCreateProject(Boolean canCreateProject) {
            if (canCreateProject == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "canCreateProject");
            }
            this.canCreateProject = canCreateProject;
            return this;
        }
        @CustomType.Setter
        public Builder colorSchemeId(Integer colorSchemeId) {
            if (colorSchemeId == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "colorSchemeId");
            }
            this.colorSchemeId = colorSchemeId;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder currentSignInAt(String currentSignInAt) {
            if (currentSignInAt == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "currentSignInAt");
            }
            this.currentSignInAt = currentSignInAt;
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            if (email == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "email");
            }
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder emailExactMatch(@Nullable Boolean emailExactMatch) {

            this.emailExactMatch = emailExactMatch;
            return this;
        }
        @CustomType.Setter
        public Builder externUid(String externUid) {
            if (externUid == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "externUid");
            }
            this.externUid = externUid;
            return this;
        }
        @CustomType.Setter
        public Builder external(Boolean external) {
            if (external == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "external");
            }
            this.external = external;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isAdmin(Boolean isAdmin) {
            if (isAdmin == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "isAdmin");
            }
            this.isAdmin = isAdmin;
            return this;
        }
        @CustomType.Setter
        public Builder isBot(Boolean isBot) {
            if (isBot == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "isBot");
            }
            this.isBot = isBot;
            return this;
        }
        @CustomType.Setter
        public Builder lastSignInAt(String lastSignInAt) {
            if (lastSignInAt == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "lastSignInAt");
            }
            this.lastSignInAt = lastSignInAt;
            return this;
        }
        @CustomType.Setter
        public Builder linkedin(String linkedin) {
            if (linkedin == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "linkedin");
            }
            this.linkedin = linkedin;
            return this;
        }
        @CustomType.Setter
        public Builder location(String location) {
            if (location == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "location");
            }
            this.location = location;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder namespaceId(Integer namespaceId) {
            if (namespaceId == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "namespaceId");
            }
            this.namespaceId = namespaceId;
            return this;
        }
        @CustomType.Setter
        public Builder note(String note) {
            if (note == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "note");
            }
            this.note = note;
            return this;
        }
        @CustomType.Setter
        public Builder organization(String organization) {
            if (organization == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "organization");
            }
            this.organization = organization;
            return this;
        }
        @CustomType.Setter
        public Builder projectsLimit(Integer projectsLimit) {
            if (projectsLimit == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "projectsLimit");
            }
            this.projectsLimit = projectsLimit;
            return this;
        }
        @CustomType.Setter
        public Builder skype(String skype) {
            if (skype == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "skype");
            }
            this.skype = skype;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder themeId(Integer themeId) {
            if (themeId == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "themeId");
            }
            this.themeId = themeId;
            return this;
        }
        @CustomType.Setter
        public Builder twitter(String twitter) {
            if (twitter == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "twitter");
            }
            this.twitter = twitter;
            return this;
        }
        @CustomType.Setter
        public Builder twoFactorEnabled(Boolean twoFactorEnabled) {
            if (twoFactorEnabled == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "twoFactorEnabled");
            }
            this.twoFactorEnabled = twoFactorEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder userId(Integer userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "userId");
            }
            this.userId = userId;
            return this;
        }
        @CustomType.Setter
        public Builder userProvider(String userProvider) {
            if (userProvider == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "userProvider");
            }
            this.userProvider = userProvider;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "username");
            }
            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder websiteUrl(String websiteUrl) {
            if (websiteUrl == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "websiteUrl");
            }
            this.websiteUrl = websiteUrl;
            return this;
        }
        public GetUserResult build() {
            final var _resultValue = new GetUserResult();
            _resultValue.avatarUrl = avatarUrl;
            _resultValue.bio = bio;
            _resultValue.canCreateGroup = canCreateGroup;
            _resultValue.canCreateProject = canCreateProject;
            _resultValue.colorSchemeId = colorSchemeId;
            _resultValue.createdAt = createdAt;
            _resultValue.currentSignInAt = currentSignInAt;
            _resultValue.email = email;
            _resultValue.emailExactMatch = emailExactMatch;
            _resultValue.externUid = externUid;
            _resultValue.external = external;
            _resultValue.id = id;
            _resultValue.isAdmin = isAdmin;
            _resultValue.isBot = isBot;
            _resultValue.lastSignInAt = lastSignInAt;
            _resultValue.linkedin = linkedin;
            _resultValue.location = location;
            _resultValue.name = name;
            _resultValue.namespaceId = namespaceId;
            _resultValue.note = note;
            _resultValue.organization = organization;
            _resultValue.projectsLimit = projectsLimit;
            _resultValue.skype = skype;
            _resultValue.state = state;
            _resultValue.themeId = themeId;
            _resultValue.twitter = twitter;
            _resultValue.twoFactorEnabled = twoFactorEnabled;
            _resultValue.userId = userId;
            _resultValue.userProvider = userProvider;
            _resultValue.username = username;
            _resultValue.websiteUrl = websiteUrl;
            return _resultValue;
        }
    }
}
