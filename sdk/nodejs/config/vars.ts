// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("gitlab");

/**
 * The GitLab Base API URL
 */
export declare const baseUrl: string | undefined;
Object.defineProperty(exports, "baseUrl", {
    get() {
        return __config.get("baseUrl");
    },
    enumerable: true,
});

/**
 * A file containing the ca certificate to use in case ssl certificate is not from a standard chain
 */
export declare const cacertFile: string | undefined;
Object.defineProperty(exports, "cacertFile", {
    get() {
        return __config.get("cacertFile");
    },
    enumerable: true,
});

/**
 * File path to client certificate when GitLab instance is behind company proxy. File must contain PEM encoded data.
 */
export declare const clientCert: string | undefined;
Object.defineProperty(exports, "clientCert", {
    get() {
        return __config.get("clientCert");
    },
    enumerable: true,
});

/**
 * File path to client key when GitLab instance is behind company proxy. File must contain PEM encoded data.
 */
export declare const clientKey: string | undefined;
Object.defineProperty(exports, "clientKey", {
    get() {
        return __config.get("clientKey");
    },
    enumerable: true,
});

/**
 * Disable SSL verification of API calls
 */
export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure");
    },
    enumerable: true,
});

/**
 * The OAuth2 token or project/personal access token used to connect to GitLab.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

