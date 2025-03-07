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


public final class PagesDomainState extends com.pulumi.resources.ResourceArgs {

    public static final PagesDomainState Empty = new PagesDomainState();

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
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return The custom domain indicated by the user.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
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
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID or [URL-encoded path of the project](https://docs.gitlab.com/api/index/#namespaced-path-encoding) owned by the authenticated user.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The URL for the given domain.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL for the given domain.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * The verification code for the domain.
     * 
     */
    @Import(name="verificationCode")
    private @Nullable Output<String> verificationCode;

    /**
     * @return The verification code for the domain.
     * 
     */
    public Optional<Output<String>> verificationCode() {
        return Optional.ofNullable(this.verificationCode);
    }

    /**
     * The certificate data.
     * 
     */
    @Import(name="verified")
    private @Nullable Output<Boolean> verified;

    /**
     * @return The certificate data.
     * 
     */
    public Optional<Output<Boolean>> verified() {
        return Optional.ofNullable(this.verified);
    }

    private PagesDomainState() {}

    private PagesDomainState(PagesDomainState $) {
        this.autoSslEnabled = $.autoSslEnabled;
        this.certificate = $.certificate;
        this.domain = $.domain;
        this.expired = $.expired;
        this.key = $.key;
        this.project = $.project;
        this.url = $.url;
        this.verificationCode = $.verificationCode;
        this.verified = $.verified;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PagesDomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PagesDomainState $;

        public Builder() {
            $ = new PagesDomainState();
        }

        public Builder(PagesDomainState defaults) {
            $ = new PagesDomainState(Objects.requireNonNull(defaults));
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
        public Builder domain(@Nullable Output<String> domain) {
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
        public Builder project(@Nullable Output<String> project) {
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

        /**
         * @param url The URL for the given domain.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL for the given domain.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param verificationCode The verification code for the domain.
         * 
         * @return builder
         * 
         */
        public Builder verificationCode(@Nullable Output<String> verificationCode) {
            $.verificationCode = verificationCode;
            return this;
        }

        /**
         * @param verificationCode The verification code for the domain.
         * 
         * @return builder
         * 
         */
        public Builder verificationCode(String verificationCode) {
            return verificationCode(Output.of(verificationCode));
        }

        /**
         * @param verified The certificate data.
         * 
         * @return builder
         * 
         */
        public Builder verified(@Nullable Output<Boolean> verified) {
            $.verified = verified;
            return this;
        }

        /**
         * @param verified The certificate data.
         * 
         * @return builder
         * 
         */
        public Builder verified(Boolean verified) {
            return verified(Output.of(verified));
        }

        public PagesDomainState build() {
            return $;
        }
    }

}
