// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.ProjectCluster` resource allows to manage the lifecycle of a project cluster.
 *
 * > This is deprecated GitLab feature since 14.5
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_clusters.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.Project("foo", {});
 * const bar = new gitlab.ProjectCluster("bar", {
 *     project: foo.id,
 *     domain: "example.com",
 *     enabled: true,
 *     kubernetesApiUrl: "https://124.124.124",
 *     kubernetesToken: "some-token",
 *     kubernetesCaCert: "some-cert",
 *     kubernetesNamespace: "namespace",
 *     kubernetesAuthorizationType: "rbac",
 *     environmentScope: "*",
 *     managementProjectId: "123456",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab project clusters can be imported using an id made up of `projectid:clusterid`, e.g.
 *
 * ```sh
 * $ pulumi import gitlab:index/projectCluster:ProjectCluster bar 123:321
 * ```
 */
export class ProjectCluster extends pulumi.CustomResource {
    /**
     * Get an existing ProjectCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectClusterState, opts?: pulumi.CustomResourceOptions): ProjectCluster {
        return new ProjectCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectCluster:ProjectCluster';

    /**
     * Returns true if the given object is an instance of ProjectCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectCluster.__pulumiType;
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
     * The unique namespace related to the project.
     */
    public readonly kubernetesNamespace!: pulumi.Output<string | undefined>;
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
     * The id of the project to add the cluster to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Provider type.
     */
    public /*out*/ readonly providerType!: pulumi.Output<string>;

    /**
     * Create a ProjectCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectClusterArgs | ProjectClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectClusterState | undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["environmentScope"] = state ? state.environmentScope : undefined;
            resourceInputs["kubernetesApiUrl"] = state ? state.kubernetesApiUrl : undefined;
            resourceInputs["kubernetesAuthorizationType"] = state ? state.kubernetesAuthorizationType : undefined;
            resourceInputs["kubernetesCaCert"] = state ? state.kubernetesCaCert : undefined;
            resourceInputs["kubernetesNamespace"] = state ? state.kubernetesNamespace : undefined;
            resourceInputs["kubernetesToken"] = state ? state.kubernetesToken : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["managementProjectId"] = state ? state.managementProjectId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platformType"] = state ? state.platformType : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["providerType"] = state ? state.providerType : undefined;
        } else {
            const args = argsOrState as ProjectClusterArgs | undefined;
            if ((!args || args.kubernetesApiUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesApiUrl'");
            }
            if ((!args || args.kubernetesToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesToken'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["environmentScope"] = args ? args.environmentScope : undefined;
            resourceInputs["kubernetesApiUrl"] = args ? args.kubernetesApiUrl : undefined;
            resourceInputs["kubernetesAuthorizationType"] = args ? args.kubernetesAuthorizationType : undefined;
            resourceInputs["kubernetesCaCert"] = args ? args.kubernetesCaCert : undefined;
            resourceInputs["kubernetesNamespace"] = args ? args.kubernetesNamespace : undefined;
            resourceInputs["kubernetesToken"] = args?.kubernetesToken ? pulumi.secret(args.kubernetesToken) : undefined;
            resourceInputs["managed"] = args ? args.managed : undefined;
            resourceInputs["managementProjectId"] = args ? args.managementProjectId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["clusterType"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["platformType"] = undefined /*out*/;
            resourceInputs["providerType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["kubernetesToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProjectCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectCluster resources.
 */
export interface ProjectClusterState {
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
     * The unique namespace related to the project.
     */
    kubernetesNamespace?: pulumi.Input<string>;
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
     * The id of the project to add the cluster to.
     */
    project?: pulumi.Input<string>;
    /**
     * Provider type.
     */
    providerType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectCluster resource.
 */
export interface ProjectClusterArgs {
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
     * The unique namespace related to the project.
     */
    kubernetesNamespace?: pulumi.Input<string>;
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
    /**
     * The id of the project to add the cluster to.
     */
    project: pulumi.Input<string>;
}
