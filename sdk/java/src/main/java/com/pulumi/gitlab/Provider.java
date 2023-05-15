// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProviderArgs;
import com.pulumi.gitlab.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the gitlab package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:gitlab")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
     * Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
     * the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
     * 
     */
    @Export(name="baseUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> baseUrl;

    /**
     * @return This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
     * Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
     * the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
     * 
     */
    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    /**
     * This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
     * CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
     * 
     */
    @Export(name="cacertFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cacertFile;

    /**
     * @return This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
     * CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
     * 
     */
    public Output<Optional<String>> cacertFile() {
        return Codegen.optional(this.cacertFile);
    }
    /**
     * File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
     * 
     */
    @Export(name="clientCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCert;

    /**
     * @return File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
     * 
     */
    public Output<Optional<String>> clientCert() {
        return Codegen.optional(this.clientCert);
    }
    /**
     * File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
     * `client_cert` is set.
     * 
     */
    @Export(name="clientKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientKey;

    /**
     * @return File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
     * `client_cert` is set.
     * 
     */
    public Output<Optional<String>> clientKey() {
        return Codegen.optional(this.clientKey);
    }
    /**
     * The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
     * used in this provider for authentication (using Bearer authorization token). See
     * https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
     * variable.
     * 
     */
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    /**
     * @return The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
     * used in this provider for authentication (using Bearer authorization token). See
     * https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
     * variable.
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
