// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.User` data source allows details of a user to be retrieved by either the user ID, username or email address.
 *
 * > Some attributes might not be returned depending on if you're an admin or not.
 *
 * > When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
 * the most related match will prioritize an exact match if one is available.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getUser:getUser", {
        "email": args.email,
        "namespaceId": args.namespaceId,
        "userId": args.userId,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
     */
    email?: string;
    /**
     * The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
     */
    namespaceId?: number;
    /**
     * The ID of the user.
     */
    userId?: number;
    /**
     * The username of the user.
     */
    username?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * The avatar URL of the user.
     */
    readonly avatarUrl: string;
    /**
     * The bio of the user.
     */
    readonly bio: string;
    /**
     * Whether the user can create groups.
     */
    readonly canCreateGroup: boolean;
    /**
     * Whether the user can create projects.
     */
    readonly canCreateProject: boolean;
    /**
     * User's color scheme ID.
     */
    readonly colorSchemeId: number;
    /**
     * Date the user was created at.
     */
    readonly createdAt: string;
    /**
     * Current user's sign-in date.
     */
    readonly currentSignInAt: string;
    /**
     * The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
     */
    readonly email: string;
    /**
     * The external UID of the user.
     */
    readonly externUid: string;
    /**
     * Whether the user is external.
     */
    readonly external: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether the user is an admin.
     */
    readonly isAdmin: boolean;
    /**
     * Whether the user is a bot.
     */
    readonly isBot: boolean;
    /**
     * Last user's sign-in date.
     */
    readonly lastSignInAt: string;
    /**
     * LinkedIn profile of the user.
     */
    readonly linkedin: string;
    /**
     * The location of the user.
     */
    readonly location: string;
    /**
     * The name of the user.
     */
    readonly name: string;
    /**
     * The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
     */
    readonly namespaceId: number;
    /**
     * Admin notes for this user.
     */
    readonly note: string;
    /**
     * The organization of the user.
     */
    readonly organization: string;
    /**
     * Number of projects the user can create.
     */
    readonly projectsLimit: number;
    /**
     * Skype username of the user.
     */
    readonly skype: string;
    /**
     * Whether the user is active or blocked.
     */
    readonly state: string;
    /**
     * User's theme ID.
     */
    readonly themeId: number;
    /**
     * Twitter username of the user.
     */
    readonly twitter: string;
    /**
     * Whether user's two-factor auth is enabled.
     */
    readonly twoFactorEnabled: boolean;
    /**
     * The ID of the user.
     */
    readonly userId: number;
    /**
     * The UID provider of the user.
     */
    readonly userProvider: string;
    /**
     * The username of the user.
     */
    readonly username: string;
    /**
     * User's website URL.
     */
    readonly websiteUrl: string;
}
/**
 * The `gitlab.User` data source allows details of a user to be retrieved by either the user ID, username or email address.
 *
 * > Some attributes might not be returned depending on if you're an admin or not.
 *
 * > When using the `email` attribute, an exact match is not guaranteed. The most related match will be returned. Starting with GitLab 16.6,
 * the most related match will prioritize an exact match if one is available.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#single-user)
 */
export function getUserOutput(args?: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply((a: any) => getUser(a, opts))
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * The public email address of the user. **Note**: before GitLab 14.8 the lookup was based on the users primary email address.
     */
    email?: pulumi.Input<string>;
    /**
     * The ID of the user's namespace. Requires admin token to access this field. Available since GitLab 14.10.
     */
    namespaceId?: pulumi.Input<number>;
    /**
     * The ID of the user.
     */
    userId?: pulumi.Input<number>;
    /**
     * The username of the user.
     */
    username?: pulumi.Input<string>;
}
