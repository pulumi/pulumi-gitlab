// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupClusterArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupClusterState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupCluster` resource allows to manage the lifecycle of a group cluster.
 * 
 * &gt; This is deprecated, due for removal in GitLab 19.0.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_clusters/)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.GroupCluster;
 * import com.pulumi.gitlab.GroupClusterArgs;
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
 *         var foo = new Group("foo", GroupArgs.builder()
 *             .name("foo-group")
 *             .path("foo-path")
 *             .build());
 * 
 *         var bar = new GroupCluster("bar", GroupClusterArgs.builder()
 *             .group(foo.id())
 *             .name("bar-cluster")
 *             .domain("example.com")
 *             .enabled(true)
 *             .kubernetesApiUrl("https://124.124.124")
 *             .kubernetesToken("some-token")
 *             .kubernetesCaCert("some-cert")
 *             .kubernetesAuthorizationType("rbac")
 *             .environmentScope("*")
 *             .managementProjectId("123456")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0, you can use an import block to import `gitlab_group_cluster`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_group_cluster.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Importing using the CLI is supported with the following syntax:
 * 
 * GitLab group clusters can be imported using an id made up of `groupid:clusterid`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupCluster:GroupCluster bar 123:321
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupCluster:GroupCluster")
public class GroupCluster extends com.pulumi.resources.CustomResource {
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
     * The id of the group to add the cluster to.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The id of the group to add the cluster to.
     * 
     */
    public Output<String> group() {
        return this.group;
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
    public GroupCluster(java.lang.String name) {
        this(name, GroupClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupCluster(java.lang.String name, GroupClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupCluster(java.lang.String name, GroupClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupCluster:GroupCluster", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupCluster(java.lang.String name, Output<java.lang.String> id, @Nullable GroupClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupCluster:GroupCluster", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupClusterArgs makeArgs(GroupClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupClusterArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static GroupCluster get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupCluster(name, id, state, options);
    }
}
