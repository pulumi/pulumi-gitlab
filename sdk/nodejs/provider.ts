// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the gitlab package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'gitlab';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
     * Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
     * the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
     */
    public readonly baseUrl!: pulumi.Output<string | undefined>;
    /**
     * This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
     * CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
     */
    public readonly cacertFile!: pulumi.Output<string | undefined>;
    /**
     * File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
     */
    public readonly clientCert!: pulumi.Output<string | undefined>;
    /**
     * File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
     * `client_cert` is set.
     */
    public readonly clientKey!: pulumi.Output<string | undefined>;
    /**
     * The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
     * used in this provider for authentication (using Bearer authorization token). See
     * https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
     * variable.
     */
    public readonly token!: pulumi.Output<string>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            resourceInputs["baseUrl"] = args ? args.baseUrl : undefined;
            resourceInputs["cacertFile"] = args ? args.cacertFile : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientKey"] = args ? args.clientKey : undefined;
            resourceInputs["earlyAuthCheck"] = pulumi.output(args ? args.earlyAuthCheck : undefined).apply(JSON.stringify);
            resourceInputs["insecure"] = pulumi.output(args ? args.insecure : undefined).apply(JSON.stringify);
            resourceInputs["token"] = args ? args.token : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * This is the target GitLab base API endpoint. Providing a value is a requirement when working with GitLab CE or GitLab
     * Enterprise e.g. `https://my.gitlab.server/api/v4/`. It is optional to provide this value and it can also be sourced from
     * the `GITLAB_BASE_URL` environment variable. The value must end with a slash.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * This is a file containing the ca cert to verify the gitlab instance. This is available for use when working with GitLab
     * CE or Gitlab Enterprise with a locally-issued or self-signed certificate chain.
     */
    cacertFile?: pulumi.Input<string>;
    /**
     * File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data. Required when
     * `client_cert` is set.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * (Experimental) By default the provider does a dummy request to get the current user in order to verify that the provider
     * configuration is correct and the GitLab API is reachable. Turn it off, to skip this check. This may be useful if the
     * GitLab instance does not yet exist and is created within the same terraform module. This is an experimental feature and
     * may change in the future. Please make sure to always keep backups of your state.
     */
    earlyAuthCheck?: pulumi.Input<boolean>;
    /**
     * When set to true this disables SSL verification of the connection to the GitLab instance.
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * The OAuth2 Token, Project, Group, Personal Access Token or CI Job Token used to connect to GitLab. The OAuth method is
     * used in this provider for authentication (using Bearer authorization token). See
     * https://docs.gitlab.com/ee/api/#authentication for details. It may be sourced from the `GITLAB_TOKEN` environment
     * variable.
     */
    token: pulumi.Input<string>;
}
