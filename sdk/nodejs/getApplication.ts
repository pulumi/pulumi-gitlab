// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `gitlab.Application` data source retrieves information about a gitlab application.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/applications/)
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitlab:index/getApplication:getApplication", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationArgs {
    id: string;
}

/**
 * A collection of values returned by getApplication.
 */
export interface GetApplicationResult {
    /**
     * Internal GitLab application id.
     */
    readonly applicationId: string;
    /**
     * Indicates if the application is kept confidential.
     */
    readonly confidential: boolean;
    readonly id: string;
    /**
     * The name of the GitLab application.
     */
    readonly name: string;
    /**
     * The redirect url of the application.
     */
    readonly redirectUrl: string;
}
/**
 * The `gitlab.Application` data source retrieves information about a gitlab application.
 *
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/applications/)
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("gitlab:index/getApplication:getApplication", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationOutputArgs {
    id: pulumi.Input<string>;
}
