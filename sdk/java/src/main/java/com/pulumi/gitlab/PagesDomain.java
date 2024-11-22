// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.PagesDomainArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.PagesDomainState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.PagesDomain` resource allows connecting custom domains and TLS certificates in GitLab Pages.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pages_domains.html)
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_pages_domain`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_pages_domain.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * GitLab pages domain can be imported using an id made up of `projectId:domain` _without_ the http protocol, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/pagesDomain:PagesDomain this 123:example.com
 * ```
 * 
 */
@ResourceType(type="gitlab:index/pagesDomain:PagesDomain")
public class PagesDomain extends com.pulumi.resources.CustomResource {
    /**
     * Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to &#34;true&#34;, certificate can&#39;t be provided.
     * 
     */
    @Export(name="autoSslEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoSslEnabled;

    /**
     * @return Enables [automatic generation](https://docs.gitlab.com/ee/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration.html) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to &#34;true&#34;, certificate can&#39;t be provided.
     * 
     */
    public Output<Boolean> autoSslEnabled() {
        return this.autoSslEnabled;
    }
    /**
     * The certificate in PEM format with intermediates following in most specific to least specific order.
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return The certificate in PEM format with intermediates following in most specific to least specific order.
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * The custom domain indicated by the user.
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return The custom domain indicated by the user.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * Whether the certificate is expired.
     * 
     */
    @Export(name="expired", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> expired;

    /**
     * @return Whether the certificate is expired.
     * 
     */
    public Output<Boolean> expired() {
        return this.expired;
    }
    /**
     * The certificate key in PEM format.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> key;

    /**
     * @return The certificate key in PEM format.
     * 
     */
    public Output<Optional<String>> key() {
        return Codegen.optional(this.key);
    }
    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding) owned by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URL for the given domain.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL for the given domain.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The verification code for the domain.
     * 
     */
    @Export(name="verificationCode", refs={String.class}, tree="[0]")
    private Output<String> verificationCode;

    /**
     * @return The verification code for the domain.
     * 
     */
    public Output<String> verificationCode() {
        return this.verificationCode;
    }
    /**
     * The certificate data.
     * 
     */
    @Export(name="verified", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> verified;

    /**
     * @return The certificate data.
     * 
     */
    public Output<Boolean> verified() {
        return this.verified;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PagesDomain(java.lang.String name) {
        this(name, PagesDomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PagesDomain(java.lang.String name, PagesDomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PagesDomain(java.lang.String name, PagesDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/pagesDomain:PagesDomain", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PagesDomain(java.lang.String name, Output<java.lang.String> id, @Nullable PagesDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/pagesDomain:PagesDomain", name, state, makeResourceOptions(options, id), false);
    }

    private static PagesDomainArgs makeArgs(PagesDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PagesDomainArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "verificationCode"
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
    public static PagesDomain get(java.lang.String name, Output<java.lang.String> id, @Nullable PagesDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PagesDomain(name, id, state, options);
    }
}
