// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetGroupProvisionedUsersProvisionedUser extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupProvisionedUsersProvisionedUser Empty = new GetGroupProvisionedUsersProvisionedUser();

    /**
     * The avatar URL of the provisioned user.
     * 
     */
    @Import(name="avatarUrl", required=true)
    private String avatarUrl;

    /**
     * @return The avatar URL of the provisioned user.
     * 
     */
    public String avatarUrl() {
        return this.avatarUrl;
    }

    /**
     * The bio of the provisioned user.
     * 
     */
    @Import(name="bio", required=true)
    private String bio;

    /**
     * @return The bio of the provisioned user.
     * 
     */
    public String bio() {
        return this.bio;
    }

    /**
     * Whether the provisioned user is a bot.
     * 
     */
    @Import(name="bot", required=true)
    private Boolean bot;

    /**
     * @return Whether the provisioned user is a bot.
     * 
     */
    public Boolean bot() {
        return this.bot;
    }

    /**
     * The confirmation date of the provisioned user.
     * 
     */
    @Import(name="confirmedAt", required=true)
    private String confirmedAt;

    /**
     * @return The confirmation date of the provisioned user.
     * 
     */
    public String confirmedAt() {
        return this.confirmedAt;
    }

    /**
     * The creation date of the provisioned user.
     * 
     */
    @Import(name="createdAt", required=true)
    private String createdAt;

    /**
     * @return The creation date of the provisioned user.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }

    /**
     * The email of the provisioned user.
     * 
     */
    @Import(name="email", required=true)
    private String email;

    /**
     * @return The email of the provisioned user.
     * 
     */
    public String email() {
        return this.email;
    }

    /**
     * Whether the provisioned user is external.
     * 
     */
    @Import(name="external", required=true)
    private Boolean external;

    /**
     * @return Whether the provisioned user is external.
     * 
     */
    public Boolean external() {
        return this.external;
    }

    /**
     * The ID of the provisioned user.
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return The ID of the provisioned user.
     * 
     */
    public String id() {
        return this.id;
    }

    /**
     * The job title of the provisioned user.
     * 
     */
    @Import(name="jobTitle", required=true)
    private String jobTitle;

    /**
     * @return The job title of the provisioned user.
     * 
     */
    public String jobTitle() {
        return this.jobTitle;
    }

    /**
     * The last activity date of the provisioned user.
     * 
     */
    @Import(name="lastActivityOn", required=true)
    private String lastActivityOn;

    /**
     * @return The last activity date of the provisioned user.
     * 
     */
    public String lastActivityOn() {
        return this.lastActivityOn;
    }

    /**
     * The last sign-in date of the provisioned user.
     * 
     */
    @Import(name="lastSignInAt", required=true)
    private String lastSignInAt;

    /**
     * @return The last sign-in date of the provisioned user.
     * 
     */
    public String lastSignInAt() {
        return this.lastSignInAt;
    }

    /**
     * The LinkedIn ID of the provisioned user.
     * 
     */
    @Import(name="linkedin", required=true)
    private String linkedin;

    /**
     * @return The LinkedIn ID of the provisioned user.
     * 
     */
    public String linkedin() {
        return this.linkedin;
    }

    /**
     * The location of the provisioned user.
     * 
     */
    @Import(name="location", required=true)
    private String location;

    /**
     * @return The location of the provisioned user.
     * 
     */
    public String location() {
        return this.location;
    }

    /**
     * The name of the provisioned user.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the provisioned user.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The organization of the provisioned user.
     * 
     */
    @Import(name="organization", required=true)
    private String organization;

    /**
     * @return The organization of the provisioned user.
     * 
     */
    public String organization() {
        return this.organization;
    }

    /**
     * Whether the provisioned user has a private profile.
     * 
     */
    @Import(name="privateProfile", required=true)
    private Boolean privateProfile;

    /**
     * @return Whether the provisioned user has a private profile.
     * 
     */
    public Boolean privateProfile() {
        return this.privateProfile;
    }

    /**
     * The pronouns of the provisioned user.
     * 
     */
    @Import(name="pronouns", required=true)
    private String pronouns;

    /**
     * @return The pronouns of the provisioned user.
     * 
     */
    public String pronouns() {
        return this.pronouns;
    }

    /**
     * The public email of the provisioned user.
     * 
     */
    @Import(name="publicEmail", required=true)
    private String publicEmail;

    /**
     * @return The public email of the provisioned user.
     * 
     */
    public String publicEmail() {
        return this.publicEmail;
    }

    /**
     * The Skype ID of the provisioned user.
     * 
     */
    @Import(name="skype", required=true)
    private String skype;

    /**
     * @return The Skype ID of the provisioned user.
     * 
     */
    public String skype() {
        return this.skype;
    }

    /**
     * The state of the provisioned user.
     * 
     */
    @Import(name="state", required=true)
    private String state;

    /**
     * @return The state of the provisioned user.
     * 
     */
    public String state() {
        return this.state;
    }

    /**
     * The Twitter ID of the provisioned user.
     * 
     */
    @Import(name="twitter", required=true)
    private String twitter;

    /**
     * @return The Twitter ID of the provisioned user.
     * 
     */
    public String twitter() {
        return this.twitter;
    }

    /**
     * Whether two-factor authentication is enabled for the provisioned user.
     * 
     */
    @Import(name="twoFactorEnabled", required=true)
    private Boolean twoFactorEnabled;

    /**
     * @return Whether two-factor authentication is enabled for the provisioned user.
     * 
     */
    public Boolean twoFactorEnabled() {
        return this.twoFactorEnabled;
    }

    /**
     * The username of the provisioned user.
     * 
     */
    @Import(name="username", required=true)
    private String username;

    /**
     * @return The username of the provisioned user.
     * 
     */
    public String username() {
        return this.username;
    }

    /**
     * The web URL of the provisioned user.
     * 
     */
    @Import(name="webUrl", required=true)
    private String webUrl;

    /**
     * @return The web URL of the provisioned user.
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }

    /**
     * The website URL of the provisioned user.
     * 
     */
    @Import(name="websiteUrl", required=true)
    private String websiteUrl;

    /**
     * @return The website URL of the provisioned user.
     * 
     */
    public String websiteUrl() {
        return this.websiteUrl;
    }

    private GetGroupProvisionedUsersProvisionedUser() {}

    private GetGroupProvisionedUsersProvisionedUser(GetGroupProvisionedUsersProvisionedUser $) {
        this.avatarUrl = $.avatarUrl;
        this.bio = $.bio;
        this.bot = $.bot;
        this.confirmedAt = $.confirmedAt;
        this.createdAt = $.createdAt;
        this.email = $.email;
        this.external = $.external;
        this.id = $.id;
        this.jobTitle = $.jobTitle;
        this.lastActivityOn = $.lastActivityOn;
        this.lastSignInAt = $.lastSignInAt;
        this.linkedin = $.linkedin;
        this.location = $.location;
        this.name = $.name;
        this.organization = $.organization;
        this.privateProfile = $.privateProfile;
        this.pronouns = $.pronouns;
        this.publicEmail = $.publicEmail;
        this.skype = $.skype;
        this.state = $.state;
        this.twitter = $.twitter;
        this.twoFactorEnabled = $.twoFactorEnabled;
        this.username = $.username;
        this.webUrl = $.webUrl;
        this.websiteUrl = $.websiteUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupProvisionedUsersProvisionedUser defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupProvisionedUsersProvisionedUser $;

        public Builder() {
            $ = new GetGroupProvisionedUsersProvisionedUser();
        }

        public Builder(GetGroupProvisionedUsersProvisionedUser defaults) {
            $ = new GetGroupProvisionedUsersProvisionedUser(Objects.requireNonNull(defaults));
        }

        /**
         * @param avatarUrl The avatar URL of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder avatarUrl(String avatarUrl) {
            $.avatarUrl = avatarUrl;
            return this;
        }

        /**
         * @param bio The bio of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder bio(String bio) {
            $.bio = bio;
            return this;
        }

        /**
         * @param bot Whether the provisioned user is a bot.
         * 
         * @return builder
         * 
         */
        public Builder bot(Boolean bot) {
            $.bot = bot;
            return this;
        }

        /**
         * @param confirmedAt The confirmation date of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder confirmedAt(String confirmedAt) {
            $.confirmedAt = confirmedAt;
            return this;
        }

        /**
         * @param createdAt The creation date of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param email The email of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            $.email = email;
            return this;
        }

        /**
         * @param external Whether the provisioned user is external.
         * 
         * @return builder
         * 
         */
        public Builder external(Boolean external) {
            $.external = external;
            return this;
        }

        /**
         * @param id The ID of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        /**
         * @param jobTitle The job title of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder jobTitle(String jobTitle) {
            $.jobTitle = jobTitle;
            return this;
        }

        /**
         * @param lastActivityOn The last activity date of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder lastActivityOn(String lastActivityOn) {
            $.lastActivityOn = lastActivityOn;
            return this;
        }

        /**
         * @param lastSignInAt The last sign-in date of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder lastSignInAt(String lastSignInAt) {
            $.lastSignInAt = lastSignInAt;
            return this;
        }

        /**
         * @param linkedin The LinkedIn ID of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder linkedin(String linkedin) {
            $.linkedin = linkedin;
            return this;
        }

        /**
         * @param location The location of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            $.location = location;
            return this;
        }

        /**
         * @param name The name of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param organization The organization of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder organization(String organization) {
            $.organization = organization;
            return this;
        }

        /**
         * @param privateProfile Whether the provisioned user has a private profile.
         * 
         * @return builder
         * 
         */
        public Builder privateProfile(Boolean privateProfile) {
            $.privateProfile = privateProfile;
            return this;
        }

        /**
         * @param pronouns The pronouns of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder pronouns(String pronouns) {
            $.pronouns = pronouns;
            return this;
        }

        /**
         * @param publicEmail The public email of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder publicEmail(String publicEmail) {
            $.publicEmail = publicEmail;
            return this;
        }

        /**
         * @param skype The Skype ID of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder skype(String skype) {
            $.skype = skype;
            return this;
        }

        /**
         * @param state The state of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            $.state = state;
            return this;
        }

        /**
         * @param twitter The Twitter ID of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder twitter(String twitter) {
            $.twitter = twitter;
            return this;
        }

        /**
         * @param twoFactorEnabled Whether two-factor authentication is enabled for the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder twoFactorEnabled(Boolean twoFactorEnabled) {
            $.twoFactorEnabled = twoFactorEnabled;
            return this;
        }

        /**
         * @param username The username of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            $.username = username;
            return this;
        }

        /**
         * @param webUrl The web URL of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder webUrl(String webUrl) {
            $.webUrl = webUrl;
            return this;
        }

        /**
         * @param websiteUrl The website URL of the provisioned user.
         * 
         * @return builder
         * 
         */
        public Builder websiteUrl(String websiteUrl) {
            $.websiteUrl = websiteUrl;
            return this;
        }

        public GetGroupProvisionedUsersProvisionedUser build() {
            if ($.avatarUrl == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "avatarUrl");
            }
            if ($.bio == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "bio");
            }
            if ($.bot == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "bot");
            }
            if ($.confirmedAt == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "confirmedAt");
            }
            if ($.createdAt == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "createdAt");
            }
            if ($.email == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "email");
            }
            if ($.external == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "external");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "id");
            }
            if ($.jobTitle == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "jobTitle");
            }
            if ($.lastActivityOn == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "lastActivityOn");
            }
            if ($.lastSignInAt == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "lastSignInAt");
            }
            if ($.linkedin == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "linkedin");
            }
            if ($.location == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "location");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "name");
            }
            if ($.organization == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "organization");
            }
            if ($.privateProfile == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "privateProfile");
            }
            if ($.pronouns == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "pronouns");
            }
            if ($.publicEmail == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "publicEmail");
            }
            if ($.skype == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "skype");
            }
            if ($.state == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "state");
            }
            if ($.twitter == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "twitter");
            }
            if ($.twoFactorEnabled == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "twoFactorEnabled");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "username");
            }
            if ($.webUrl == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "webUrl");
            }
            if ($.websiteUrl == null) {
                throw new MissingRequiredPropertyException("GetGroupProvisionedUsersProvisionedUser", "websiteUrl");
            }
            return $;
        }
    }

}
