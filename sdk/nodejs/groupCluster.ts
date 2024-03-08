// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.GroupCluster` resource allows to manage the lifecycle of a group cluster.
 *
 * > This is deprecated GitLab feature since 14.5
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_clusters.html)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.Group("foo", {path: "foo-path"});
 * const bar = new gitlab.GroupCluster("bar", {
 *     group: foo.id,
 *     domain: "example.com",
 *     enabled: true,
 *     kubernetesApiUrl: "https://124.124.124",
 *     kubernetesToken: "some-token",
 *     kubernetesCaCert: "some-cert",
 *     kubernetesAuthorizationType: "rbac",
 *     environmentScope: "*",
 *     managementProjectId: "123456",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GitLab group clusters can be imported using an id made up of `groupid:clusterid`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/groupCluster:GroupCluster bar 123:321
 * ```
 */
export class GroupCluster extends pulumi.CustomResource {
    /**
     * Get an existing GroupCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupClusterState, opts?: pulumi.CustomResourceOptions): GroupCluster {
        return new GroupCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupCluster:GroupCluster';

    /**
     * Returns true if the given object is an instance of GroupCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupCluster.__pulumiType;
    }

    /**
     * Cluster type.
     */
    public /*out*/ readonly clusterType!: pulumi.Output<string>;
    /**
     * Create time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The base domain of the cluster.
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    /**
     * Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The associated environment to the cluster. Defaults to `*`.
     */
    public readonly environmentScope!: pulumi.Output<string | undefined>;
    /**
     * The id of the group to add the cluster to.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The URL to access the Kubernetes API.
     */
    public readonly kubernetesApiUrl!: pulumi.Output<string>;
    /**
     * The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
     */
    public readonly kubernetesAuthorizationType!: pulumi.Output<string | undefined>;
    /**
     * TLS certificate (needed if API is using a self-signed TLS certificate).
     */
    public readonly kubernetesCaCert!: pulumi.Output<string | undefined>;
    /**
     * The token to authenticate against Kubernetes.
     */
    public readonly kubernetesToken!: pulumi.Output<string>;
    /**
     * Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     */
    public readonly managed!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the management project for the cluster.
     */
    public readonly managementProjectId!: pulumi.Output<string | undefined>;
    /**
     * The name of cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Platform type.
     */
    public /*out*/ readonly platformType!: pulumi.Output<string>;
    /**
     * Provider type.
     */
    public /*out*/ readonly providerType!: pulumi.Output<string>;

    /**
     * Create a GroupCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupClusterArgs | GroupClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupClusterState | undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["environmentScope"] = state ? state.environmentScope : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["kubernetesApiUrl"] = state ? state.kubernetesApiUrl : undefined;
            resourceInputs["kubernetesAuthorizationType"] = state ? state.kubernetesAuthorizationType : undefined;
            resourceInputs["kubernetesCaCert"] = state ? state.kubernetesCaCert : undefined;
            resourceInputs["kubernetesToken"] = state ? state.kubernetesToken : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["managementProjectId"] = state ? state.managementProjectId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platformType"] = state ? state.platformType : undefined;
            resourceInputs["providerType"] = state ? state.providerType : undefined;
        } else {
            const args = argsOrState as GroupClusterArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.kubernetesApiUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesApiUrl'");
            }
            if ((!args || args.kubernetesToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesToken'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["environmentScope"] = args ? args.environmentScope : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["kubernetesApiUrl"] = args ? args.kubernetesApiUrl : undefined;
            resourceInputs["kubernetesAuthorizationType"] = args ? args.kubernetesAuthorizationType : undefined;
            resourceInputs["kubernetesCaCert"] = args ? args.kubernetesCaCert : undefined;
            resourceInputs["kubernetesToken"] = args?.kubernetesToken ? pulumi.secret(args.kubernetesToken) : undefined;
            resourceInputs["managed"] = args ? args.managed : undefined;
            resourceInputs["managementProjectId"] = args ? args.managementProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["clusterType"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["platformType"] = undefined /*out*/;
            resourceInputs["providerType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["kubernetesToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(GroupCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupCluster resources.
 */
export interface GroupClusterState {
    /**
     * Cluster type.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * Create time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The base domain of the cluster.
     */
    domain?: pulumi.Input<string>;
    /**
     * Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The associated environment to the cluster. Defaults to `*`.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * The id of the group to add the cluster to.
     */
    group?: pulumi.Input<string>;
    /**
     * The URL to access the Kubernetes API.
     */
    kubernetesApiUrl?: pulumi.Input<string>;
    /**
     * The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
     */
    kubernetesAuthorizationType?: pulumi.Input<string>;
    /**
     * TLS certificate (needed if API is using a self-signed TLS certificate).
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * The token to authenticate against Kubernetes.
     */
    kubernetesToken?: pulumi.Input<string>;
    /**
     * Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     */
    managed?: pulumi.Input<boolean>;
    /**
     * The ID of the management project for the cluster.
     */
    managementProjectId?: pulumi.Input<string>;
    /**
     * The name of cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Platform type.
     */
    platformType?: pulumi.Input<string>;
    /**
     * Provider type.
     */
    providerType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupCluster resource.
 */
export interface GroupClusterArgs {
    /**
     * The base domain of the cluster.
     */
    domain?: pulumi.Input<string>;
    /**
     * Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The associated environment to the cluster. Defaults to `*`.
     */
    environmentScope?: pulumi.Input<string>;
    /**
     * The id of the group to add the cluster to.
     */
    group: pulumi.Input<string>;
    /**
     * The URL to access the Kubernetes API.
     */
    kubernetesApiUrl: pulumi.Input<string>;
    /**
     * The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
     */
    kubernetesAuthorizationType?: pulumi.Input<string>;
    /**
     * TLS certificate (needed if API is using a self-signed TLS certificate).
     */
    kubernetesCaCert?: pulumi.Input<string>;
    /**
     * The token to authenticate against Kubernetes.
     */
    kubernetesToken: pulumi.Input<string>;
    /**
     * Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     */
    managed?: pulumi.Input<boolean>;
    /**
     * The ID of the management project for the cluster.
     */
    managementProjectId?: pulumi.Input<string>;
    /**
     * The name of cluster.
     */
    name?: pulumi.Input<string>;
}
