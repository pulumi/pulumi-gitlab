// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface ProjectPushRules {
    /**
     * All commit author emails must match this regex, e.g. `@my-company.com$`.
     */
    authorEmailRegex?: pulumi.Input<string>;
    /**
     * All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     */
    branchNameRegex?: pulumi.Input<string>;
    /**
     * Users can only push commits to this repository that were committed with one of their own verified emails.
     */
    commitCommitterCheck?: pulumi.Input<boolean>;
    /**
     * No commit message is allowed to match this regex, for example `ssh\:\/\/`.
     */
    commitMessageNegativeRegex?: pulumi.Input<string>;
    /**
     * All commit messages must match this regex, e.g. `Fixed \d+\..*`.
     */
    commitMessageRegex?: pulumi.Input<string>;
    /**
     * Deny deleting a tag.
     */
    denyDeleteTag?: pulumi.Input<boolean>;
    /**
     * All commited filenames must not match this regex, e.g. `(jar|exe)$`.
     */
    fileNameRegex?: pulumi.Input<string>;
    /**
     * Maximum file size (MB).
     */
    maxFileSize?: pulumi.Input<number>;
    /**
     * Restrict commits by author (email) to existing GitLab users.
     */
    memberCheck?: pulumi.Input<boolean>;
    /**
     * GitLab will reject any files that are likely to contain secrets.
     */
    preventSecrets?: pulumi.Input<boolean>;
    /**
     * Reject commit when it’s not signed through GPG.
     */
    rejectUnsignedCommits?: pulumi.Input<boolean>;
}
