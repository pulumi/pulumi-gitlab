// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectClusterArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectClusterState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectCluster` resource allows to manage the lifecycle of a project cluster.
 * 
 * &gt; This is deprecated GitLab feature since 14.5
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_clusters.html)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectCluster;
 * import com.pulumi.gitlab.ProjectClusterArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var foo = new Project(&#34;foo&#34;);
 * 
 *         var bar = new ProjectCluster(&#34;bar&#34;, ProjectClusterArgs.builder()        
 *             .project(foo.id())
 *             .domain(&#34;example.com&#34;)
 *             .enabled(true)
 *             .kubernetesApiUrl(&#34;https://124.124.124&#34;)
 *             .kubernetesToken(&#34;some-token&#34;)
 *             .kubernetesCaCert(&#34;some-cert&#34;)
 *             .kubernetesNamespace(&#34;namespace&#34;)
 *             .kubernetesAuthorizationType(&#34;rbac&#34;)
 *             .environmentScope(&#34;*&#34;)
 *             .managementProjectId(&#34;123456&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitLab project clusters can be imported using an id made up of `projectid:clusterid`, e.g.
 * 
 * ```sh
 *  $ pulumi import gitlab:index/projectCluster:ProjectCluster bar 123:321
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectCluster:ProjectCluster")
public class ProjectCluster extends com.pulumi.resources.CustomResource {
    /**
     * Cluster type.
     * 
     */
    @Export(name="clusterType", refs={String.class}, tree="[0]")
    private Output<String> clusterType;

    /**
     * @return Cluster type.
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }
    /**
     * Create time.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Create time.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The base domain of the cluster.
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domain;

    /**
     * @return The base domain of the cluster.
     * 
     */
    public Output<Optional<String>> domain() {
        return Codegen.optional(this.domain);
    }
    /**
     * Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The associated environment to the cluster. Defaults to `*`.
     * 
     */
    @Export(name="environmentScope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> environmentScope;

    /**
     * @return The associated environment to the cluster. Defaults to `*`.
     * 
     */
    public Output<Optional<String>> environmentScope() {
        return Codegen.optional(this.environmentScope);
    }
    /**
     * The URL to access the Kubernetes API.
     * 
     */
    @Export(name="kubernetesApiUrl", refs={String.class}, tree="[0]")
    private Output<String> kubernetesApiUrl;

    /**
     * @return The URL to access the Kubernetes API.
     * 
     */
    public Output<String> kubernetesApiUrl() {
        return this.kubernetesApiUrl;
    }
    /**
     * The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
     * 
     */
    @Export(name="kubernetesAuthorizationType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kubernetesAuthorizationType;

    /**
     * @return The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
     * 
     */
    public Output<Optional<String>> kubernetesAuthorizationType() {
        return Codegen.optional(this.kubernetesAuthorizationType);
    }
    /**
     * TLS certificate (needed if API is using a self-signed TLS certificate).
     * 
     */
    @Export(name="kubernetesCaCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kubernetesCaCert;

    /**
     * @return TLS certificate (needed if API is using a self-signed TLS certificate).
     * 
     */
    public Output<Optional<String>> kubernetesCaCert() {
        return Codegen.optional(this.kubernetesCaCert);
    }
    /**
     * The unique namespace related to the project.
     * 
     */
    @Export(name="kubernetesNamespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kubernetesNamespace;

    /**
     * @return The unique namespace related to the project.
     * 
     */
    public Output<Optional<String>> kubernetesNamespace() {
        return Codegen.optional(this.kubernetesNamespace);
    }
    /**
     * The token to authenticate against Kubernetes.
     * 
     */
    @Export(name="kubernetesToken", refs={String.class}, tree="[0]")
    private Output<String> kubernetesToken;

    /**
     * @return The token to authenticate against Kubernetes.
     * 
     */
    public Output<String> kubernetesToken() {
        return this.kubernetesToken;
    }
    /**
     * Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    @Export(name="managed", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> managed;

    /**
     * @return Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    public Output<Optional<Boolean>> managed() {
        return Codegen.optional(this.managed);
    }
    /**
     * The ID of the management project for the cluster.
     * 
     */
    @Export(name="managementProjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> managementProjectId;

    /**
     * @return The ID of the management project for the cluster.
     * 
     */
    public Output<Optional<String>> managementProjectId() {
        return Codegen.optional(this.managementProjectId);
    }
    /**
     * The name of cluster.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of cluster.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Platform type.
     * 
     */
    @Export(name="platformType", refs={String.class}, tree="[0]")
    private Output<String> platformType;

    /**
     * @return Platform type.
     * 
     */
    public Output<String> platformType() {
        return this.platformType;
    }
    /**
     * The id of the project to add the cluster to.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The id of the project to add the cluster to.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Provider type.
     * 
     */
    @Export(name="providerType", refs={String.class}, tree="[0]")
    private Output<String> providerType;

    /**
     * @return Provider type.
     * 
     */
    public Output<String> providerType() {
        return this.providerType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectCluster(String name) {
        this(name, ProjectClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectCluster(String name, ProjectClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectCluster(String name, ProjectClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectCluster:ProjectCluster", name, args == null ? ProjectClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectCluster(String name, Output<String> id, @Nullable ProjectClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectCluster:ProjectCluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "kubernetesToken"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ProjectCluster get(String name, Output<String> id, @Nullable ProjectClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectCluster(name, id, state, options);
    }
}
