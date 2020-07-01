// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides details about a specific user in the gitlab provider. Especially the ability to lookup the id for linking to other resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = pulumi.output(gitlab.getUser({
 *     username: "myuser",
 * }, { async: true }));
 * ```
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gitlab:index/getUser:getUser", {
        "email": args.email,
        "userId": args.userId,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The e-mail address of the user. (Requires administrator privileges)
     */
    readonly email?: string;
    /**
     * The ID of the user.
     */
    readonly userId?: number;
    /**
     * The username of the user.
     */
    readonly username?: string;
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
     * The e-mail address of the user.
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
     * Last user's sign-in date.
     */
    readonly lastSignInAt: string;
    /**
     * Linkedin profile of the user.
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
     * Whether user's two factor auth is enabled.
     */
    readonly twoFactorEnabled: boolean;
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
