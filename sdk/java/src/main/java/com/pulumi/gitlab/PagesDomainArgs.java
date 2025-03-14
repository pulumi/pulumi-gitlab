// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PagesDomainArgs extends com.pulumi.resources.ResourceArgs {

    public static final PagesDomainArgs Empty = new PagesDomainArgs();

    /**
     * Enables [automatic generation](https://docs.gitlab.com/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration/) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to &#34;true&#34;, certificate can&#39;t be provided.
     * 
     */
    @Import(name="autoSslEnabled")
    private @Nullable Output<Boolean> autoSslEnabled;

    /**
     * @return Enables [automatic generation](https://docs.gitlab.com/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration/) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to &#34;true&#34;, certificate can&#39;t be provided.
     * 
     */
    public Optional<Output<Boolean>> autoSslEnabled() {
        return Optional.ofNullable(this.autoSslEnabled);
    }

    /**
     * The certificate in PEM format with intermediates following in most specific to least specific order.
     * 
     */
    @Import(name="certificate")
    private @Nullable Output<String> certificate;

    /**
     * @return The certificate in PEM format with intermediates following in most specific to least specific order.
     * 
     */
    public Optional<Output<String>> certificate() {
        return Optional.ofNullable(this.certificate);
    }

    /**
     * The custom domain indicated by the user.
     * 
     */
    @Import(name="domain", required=true)
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
    @Import(name="expired")
    private @Nullable Output<Boolean> expired;

    /**
     * @return Whether the certificate is expired.
     * 
     */
    public Optional<Output<Boolean>> expired() {
        return Optional.ofNullable(this.expired);
    }

    /**
     * The certificate key in PEM format.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The certificate key in PEM format.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding) owned by the authenticated user.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding) owned by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    private PagesDomainArgs() {}

    private PagesDomainArgs(PagesDomainArgs $) {
        this.autoSslEnabled = $.autoSslEnabled;
        this.certificate = $.certificate;
        this.domain = $.domain;
        this.expired = $.expired;
        this.key = $.key;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PagesDomainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PagesDomainArgs $;

        public Builder() {
            $ = new PagesDomainArgs();
        }

        public Builder(PagesDomainArgs defaults) {
            $ = new PagesDomainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoSslEnabled Enables [automatic generation](https://docs.gitlab.com/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration/) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to &#34;true&#34;, certificate can&#39;t be provided.
         * 
         * @return builder
         * 
         */
        public Builder autoSslEnabled(@Nullable Output<Boolean> autoSslEnabled) {
            $.autoSslEnabled = autoSslEnabled;
            return this;
        }

        /**
         * @param autoSslEnabled Enables [automatic generation](https://docs.gitlab.com/user/project/pages/custom_domains_ssl_tls_certification/lets_encrypt_integration/) of SSL certificates issued by Let’s Encrypt for custom domains. When this is set to &#34;true&#34;, certificate can&#39;t be provided.
         * 
         * @return builder
         * 
         */
        public Builder autoSslEnabled(Boolean autoSslEnabled) {
            return autoSslEnabled(Output.of(autoSslEnabled));
        }

        /**
         * @param certificate The certificate in PEM format with intermediates following in most specific to least specific order.
         * 
         * @return builder
         * 
         */
        public Builder certificate(@Nullable Output<String> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate The certificate in PEM format with intermediates following in most specific to least specific order.
         * 
         * @return builder
         * 
         */
        public Builder certificate(String certificate) {
            return certificate(Output.of(certificate));
        }

        /**
         * @param domain The custom domain indicated by the user.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The custom domain indicated by the user.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param expired Whether the certificate is expired.
         * 
         * @return builder
         * 
         */
        public Builder expired(@Nullable Output<Boolean> expired) {
            $.expired = expired;
            return this;
        }

        /**
         * @param expired Whether the certificate is expired.
         * 
         * @return builder
         * 
         */
        public Builder expired(Boolean expired) {
            return expired(Output.of(expired));
        }

        /**
         * @param key The certificate key in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The certificate key in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param project The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding) owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding) owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public PagesDomainArgs build() {
            if ($.domain == null) {
                throw new MissingRequiredPropertyException("PagesDomainArgs", "domain");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("PagesDomainArgs", "project");
            }
            return $;
        }
    }

}
