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
 *     accessLevel: "developer",
 *     cn: "testuser",
 *     groupId: "12345",
 *     ldapProvider: "ldapmain",
 * });
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
     * Acceptable values are: guest, reporter, developer, maintainer, owner.
     */
    public readonly accessLevel!: pulumi.Output<string>;
    /**
     * The CN of the LDAP group to link with.
     */
    public readonly cn!: pulumi.Output<string>;
    public readonly force!: pulumi.Output<boolean | undefined>;
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupLdapLinkState | undefined;
            inputs["accessLevel"] = state ? state.accessLevel : undefined;
            inputs["cn"] = state ? state.cn : undefined;
            inputs["force"] = state ? state.force : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["ldapProvider"] = state ? state.ldapProvider : undefined;
        } else {
            const args = argsOrState as GroupLdapLinkArgs | undefined;
            if (!args || args.accessLevel === undefined) {
                throw new Error("Missing required property 'accessLevel'");
            }
            if (!args || args.cn === undefined) {
                throw new Error("Missing required property 'cn'");
            }
            if (!args || args.groupId === undefined) {
                throw new Error("Missing required property 'groupId'");
            }
            if (!args || args.ldapProvider === undefined) {
                throw new Error("Missing required property 'ldapProvider'");
            }
            inputs["accessLevel"] = args ? args.accessLevel : undefined;
            inputs["cn"] = args ? args.cn : undefined;
            inputs["force"] = args ? args.force : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["ldapProvider"] = args ? args.ldapProvider : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupLdapLink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupLdapLink resources.
 */
export interface GroupLdapLinkState {
    /**
     * Acceptable values are: guest, reporter, developer, maintainer, owner.
     */
    readonly accessLevel?: pulumi.Input<string>;
    /**
     * The CN of the LDAP group to link with.
     */
    readonly cn?: pulumi.Input<string>;
    readonly force?: pulumi.Input<boolean>;
    /**
     * The id of the GitLab group.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database.
     */
    readonly ldapProvider?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupLdapLink resource.
 */
export interface GroupLdapLinkArgs {
    /**
     * Acceptable values are: guest, reporter, developer, maintainer, owner.
     */
    readonly accessLevel: pulumi.Input<string>;
    /**
     * The CN of the LDAP group to link with.
     */
    readonly cn: pulumi.Input<string>;
    readonly force?: pulumi.Input<boolean>;
    /**
     * The id of the GitLab group.
     */
    readonly groupId: pulumi.Input<string>;
    /**
     * The name of the LDAP provider as stored in the GitLab database.
     */
    readonly ldapProvider: pulumi.Input<string>;
}
