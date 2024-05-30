// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)
 *
 * ## Import
 *
 * GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn:filter`. CN and Filter are mutually exclusive, so one will be missing.
 *
 * If using the CN for the group link, the ID will end with a blank filter (":"). e.g.,
 *
 * ```sh
 * $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testcn:"
 * ```
 *
 * If using the Filter for the group link, the ID will have two "::" in the middle due to having a blank CN. e.g.,
 *
 * ```sh
 * $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain::testfilter"
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
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     *
     * @deprecated Use `groupAccess` instead of the `accessLevel` attribute.
     */
    public readonly accessLevel!: pulumi.Output<string | undefined>;
    /**
     * The CN of the LDAP group to link with. Required if `filter` is not provided.
     */
    public readonly cn!: pulumi.Output<string>;
    /**
     * The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * The ID or URL-encoded path of the group
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     */
    public readonly groupAccess!: pulumi.Output<string | undefined>;
    /**
     * The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
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
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["groupAccess"] = state ? state.groupAccess : undefined;
            resourceInputs["ldapProvider"] = state ? state.ldapProvider : undefined;
        } else {
            const args = argsOrState as GroupLdapLinkArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.ldapProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapProvider'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["cn"] = args ? args.cn : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["groupAccess"] = args ? args.groupAccess : undefined;
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
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     *
     * @deprecated Use `groupAccess` instead of the `accessLevel` attribute.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The CN of the LDAP group to link with. Required if `filter` is not provided.
     */
    cn?: pulumi.Input<string>;
    /**
     * The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     */
    filter?: pulumi.Input<string>;
    /**
     * If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The ID or URL-encoded path of the group
     */
    group?: pulumi.Input<string>;
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     */
    ldapProvider?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupLdapLink resource.
 */
export interface GroupLdapLinkArgs {
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     *
     * @deprecated Use `groupAccess` instead of the `accessLevel` attribute.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The CN of the LDAP group to link with. Required if `filter` is not provided.
     */
    cn?: pulumi.Input<string>;
    /**
     * The LDAP filter for the group. Required if `cn` is not provided. Requires GitLab Premium or above.
     */
    filter?: pulumi.Input<string>;
    /**
     * If true, then delete and replace an existing LDAP link if one exists. Will also remove an LDAP link if the parent group is not found.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The ID or URL-encoded path of the group
     */
    group: pulumi.Input<string>;
    /**
     * Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`
     */
    groupAccess?: pulumi.Input<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
     */
    ldapProvider: pulumi.Input<string>;
}
