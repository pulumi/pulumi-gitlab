// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to add an LDAP link to an existing GitLab group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const test = new gitlab.GroupLdapLink("test", {
 *     cn: "testuser",
 *     groupAccess: "developer",
 *     groupId: "12345",
 *     ldapProvider: "ldapmain",
 * });
 * ```
 *
 * ## Import
 *
 * # GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testuser"
 * ```
 */
export class GroupLdapLink extends pulumi.CustomResource {
    /**
     * Get an existing GroupLdapLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupLdapLinkState, opts?: pulumi.CustomResourceOptions): GroupLdapLink {
        return new GroupLdapLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupLdapLink:GroupLdapLink';

    /**
     * Returns true if the given object is an instance of GroupLdapLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupLdapLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupLdapLink.__pulumiType;
    }

    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     *
     * @deprecated Use `group_access` instead of the `access_level` attribute.
     */
    public readonly accessLevel!: pulumi.Output<string | undefined>;
    /**
     * The CN of the LDAP group to link with.
     */
    public readonly cn!: pulumi.Output<string>;
    /**
     * If true, then delete and replace an existing LDAP link if one exists.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     */
    public readonly groupAccess!: pulumi.Output<string | undefined>;
    /**
     * The id of the GitLab group.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database.
     */
    public readonly ldapProvider!: pulumi.Output<string>;

    /**
     * Create a GroupLdapLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupLdapLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupLdapLinkArgs | GroupLdapLinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupLdapLinkState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["cn"] = state ? state.cn : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["groupAccess"] = state ? state.groupAccess : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["ldapProvider"] = state ? state.ldapProvider : undefined;
        } else {
            const args = argsOrState as GroupLdapLinkArgs | undefined;
            if ((!args || args.cn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cn'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.ldapProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapProvider'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["cn"] = args ? args.cn : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["groupAccess"] = args ? args.groupAccess : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["ldapProvider"] = args ? args.ldapProvider : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupLdapLink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupLdapLink resources.
 */
export interface GroupLdapLinkState {
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     *
     * @deprecated Use `group_access` instead of the `access_level` attribute.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The CN of the LDAP group to link with.
     */
    cn?: pulumi.Input<string>;
    /**
     * If true, then delete and replace an existing LDAP link if one exists.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The id of the GitLab group.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database.
     */
    ldapProvider?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupLdapLink resource.
 */
export interface GroupLdapLinkArgs {
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     *
     * @deprecated Use `group_access` instead of the `access_level` attribute.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The CN of the LDAP group to link with.
     */
    cn: pulumi.Input<string>;
    /**
     * If true, then delete and replace an existing LDAP link if one exists.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The id of the GitLab group.
     */
    groupId: pulumi.Input<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database.
     */
    ldapProvider: pulumi.Input<string>;
}
