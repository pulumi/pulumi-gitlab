// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupClusterState extends com.pulumi.resources.ResourceArgs {

    public static final GroupClusterState Empty = new GroupClusterState();

    /**
     * Cluster type.
     * 
     */
    @Import(name="clusterType")
    private @Nullable Output<String> clusterType;

    /**
     * @return Cluster type.
     * 
     */
    public Optional<Output<String>> clusterType() {
        return Optional.ofNullable(this.clusterType);
    }

    /**
     * Create time.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Create time.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The base domain of the cluster.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return The base domain of the cluster.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The associated environment to the cluster. Defaults to `*`.
     * 
     */
    @Import(name="environmentScope")
    private @Nullable Output<String> environmentScope;

    /**
     * @return The associated environment to the cluster. Defaults to `*`.
     * 
     */
    public Optional<Output<String>> environmentScope() {
        return Optional.ofNullable(this.environmentScope);
    }

    /**
     * The id of the group to add the cluster to.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return The id of the group to add the cluster to.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * The URL to access the Kubernetes API.
     * 
     */
    @Import(name="kubernetesApiUrl")
    private @Nullable Output<String> kubernetesApiUrl;

    /**
     * @return The URL to access the Kubernetes API.
     * 
     */
    public Optional<Output<String>> kubernetesApiUrl() {
        return Optional.ofNullable(this.kubernetesApiUrl);
    }

    /**
     * The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
     * 
     */
    @Import(name="kubernetesAuthorizationType")
    private @Nullable Output<String> kubernetesAuthorizationType;

    /**
     * @return The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
     * 
     */
    public Optional<Output<String>> kubernetesAuthorizationType() {
        return Optional.ofNullable(this.kubernetesAuthorizationType);
    }

    /**
     * TLS certificate (needed if API is using a self-signed TLS certificate).
     * 
     */
    @Import(name="kubernetesCaCert")
    private @Nullable Output<String> kubernetesCaCert;

    /**
     * @return TLS certificate (needed if API is using a self-signed TLS certificate).
     * 
     */
    public Optional<Output<String>> kubernetesCaCert() {
        return Optional.ofNullable(this.kubernetesCaCert);
    }

    /**
     * The token to authenticate against Kubernetes.
     * 
     */
    @Import(name="kubernetesToken")
    private @Nullable Output<String> kubernetesToken;

    /**
     * @return The token to authenticate against Kubernetes.
     * 
     */
    public Optional<Output<String>> kubernetesToken() {
        return Optional.ofNullable(this.kubernetesToken);
    }

    /**
     * Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    @Import(name="managed")
    private @Nullable Output<Boolean> managed;

    /**
     * @return Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
     * 
     */
    public Optional<Output<Boolean>> managed() {
        return Optional.ofNullable(this.managed);
    }

    /**
     * The ID of the management project for the cluster.
     * 
     */
    @Import(name="managementProjectId")
    private @Nullable Output<String> managementProjectId;

    /**
     * @return The ID of the management project for the cluster.
     * 
     */
    public Optional<Output<String>> managementProjectId() {
        return Optional.ofNullable(this.managementProjectId);
    }

    /**
     * The name of cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Platform type.
     * 
     */
    @Import(name="platformType")
    private @Nullable Output<String> platformType;

    /**
     * @return Platform type.
     * 
     */
    public Optional<Output<String>> platformType() {
        return Optional.ofNullable(this.platformType);
    }

    /**
     * Provider type.
     * 
     */
    @Import(name="providerType")
    private @Nullable Output<String> providerType;

    /**
     * @return Provider type.
     * 
     */
    public Optional<Output<String>> providerType() {
        return Optional.ofNullable(this.providerType);
    }

    private GroupClusterState() {}

    private GroupClusterState(GroupClusterState $) {
        this.clusterType = $.clusterType;
        this.createdAt = $.createdAt;
        this.domain = $.domain;
        this.enabled = $.enabled;
        this.environmentScope = $.environmentScope;
        this.group = $.group;
        this.kubernetesApiUrl = $.kubernetesApiUrl;
        this.kubernetesAuthorizationType = $.kubernetesAuthorizationType;
        this.kubernetesCaCert = $.kubernetesCaCert;
        this.kubernetesToken = $.kubernetesToken;
        this.managed = $.managed;
        this.managementProjectId = $.managementProjectId;
        this.name = $.name;
        this.platformType = $.platformType;
        this.providerType = $.providerType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupClusterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupClusterState $;

        public Builder() {
            $ = new GroupClusterState();
        }

        public Builder(GroupClusterState defaults) {
            $ = new GroupClusterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterType Cluster type.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(@Nullable Output<String> clusterType) {
            $.clusterType = clusterType;
            return this;
        }

        /**
         * @param clusterType Cluster type.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(String clusterType) {
            return clusterType(Output.of(clusterType));
        }

        /**
         * @param createdAt Create time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Create time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param domain The base domain of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The base domain of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param enabled Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param environmentScope The associated environment to the cluster. Defaults to `*`.
         * 
         * @return builder
         * 
         */
        public Builder environmentScope(@Nullable Output<String> environmentScope) {
            $.environmentScope = environmentScope;
            return this;
        }

        /**
         * @param environmentScope The associated environment to the cluster. Defaults to `*`.
         * 
         * @return builder
         * 
         */
        public Builder environmentScope(String environmentScope) {
            return environmentScope(Output.of(environmentScope));
        }

        /**
         * @param group The id of the group to add the cluster to.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The id of the group to add the cluster to.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param kubernetesApiUrl The URL to access the Kubernetes API.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesApiUrl(@Nullable Output<String> kubernetesApiUrl) {
            $.kubernetesApiUrl = kubernetesApiUrl;
            return this;
        }

        /**
         * @param kubernetesApiUrl The URL to access the Kubernetes API.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesApiUrl(String kubernetesApiUrl) {
            return kubernetesApiUrl(Output.of(kubernetesApiUrl));
        }

        /**
         * @param kubernetesAuthorizationType The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesAuthorizationType(@Nullable Output<String> kubernetesAuthorizationType) {
            $.kubernetesAuthorizationType = kubernetesAuthorizationType;
            return this;
        }

        /**
         * @param kubernetesAuthorizationType The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesAuthorizationType(String kubernetesAuthorizationType) {
            return kubernetesAuthorizationType(Output.of(kubernetesAuthorizationType));
        }

        /**
         * @param kubernetesCaCert TLS certificate (needed if API is using a self-signed TLS certificate).
         * 
         * @return builder
         * 
         */
        public Builder kubernetesCaCert(@Nullable Output<String> kubernetesCaCert) {
            $.kubernetesCaCert = kubernetesCaCert;
            return this;
        }

        /**
         * @param kubernetesCaCert TLS certificate (needed if API is using a self-signed TLS certificate).
         * 
         * @return builder
         * 
         */
        public Builder kubernetesCaCert(String kubernetesCaCert) {
            return kubernetesCaCert(Output.of(kubernetesCaCert));
        }

        /**
         * @param kubernetesToken The token to authenticate against Kubernetes.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesToken(@Nullable Output<String> kubernetesToken) {
            $.kubernetesToken = kubernetesToken;
            return this;
        }

        /**
         * @param kubernetesToken The token to authenticate against Kubernetes.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesToken(String kubernetesToken) {
            return kubernetesToken(Output.of(kubernetesToken));
        }

        /**
         * @param managed Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
         * 
         * @return builder
         * 
         */
        public Builder managed(@Nullable Output<Boolean> managed) {
            $.managed = managed;
            return this;
        }

        /**
         * @param managed Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
         * 
         * @return builder
         * 
         */
        public Builder managed(Boolean managed) {
            return managed(Output.of(managed));
        }

        /**
         * @param managementProjectId The ID of the management project for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder managementProjectId(@Nullable Output<String> managementProjectId) {
            $.managementProjectId = managementProjectId;
            return this;
        }

        /**
         * @param managementProjectId The ID of the management project for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder managementProjectId(String managementProjectId) {
            return managementProjectId(Output.of(managementProjectId));
        }

        /**
         * @param name The name of cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param platformType Platform type.
         * 
         * @return builder
         * 
         */
        public Builder platformType(@Nullable Output<String> platformType) {
            $.platformType = platformType;
            return this;
        }

        /**
         * @param platformType Platform type.
         * 
         * @return builder
         * 
         */
        public Builder platformType(String platformType) {
            return platformType(Output.of(platformType));
        }

        /**
         * @param providerType Provider type.
         * 
         * @return builder
         * 
         */
        public Builder providerType(@Nullable Output<String> providerType) {
            $.providerType = providerType;
            return this;
        }

        /**
         * @param providerType Provider type.
         * 
         * @return builder
         * 
         */
        public Builder providerType(String providerType) {
            return providerType(Output.of(providerType));
        }

        public GroupClusterState build() {
            return $;
        }
    }

}
