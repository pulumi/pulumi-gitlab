// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface BranchCommit {
    authorEmail?: pulumi.Input<string>;
    authorName?: pulumi.Input<string>;
    authoredDate?: pulumi.Input<string>;
    committedDate?: pulumi.Input<string>;
    committerEmail?: pulumi.Input<string>;
    committerName?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    message?: pulumi.Input<string>;
    parentIds?: pulumi.Input<pulumi.Input<string>[]>;
    shortId?: pulumi.Input<string>;
    title?: pulumi.Input<string>;
}

export interface BranchProtectionAllowedToMerge {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface BranchProtectionAllowedToPush {
    /**
     * Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface BranchProtectionAllowedToUnprotect {
    /**
     * Access levels allowed to unprotect push to protected branch. Valid values are: `developer`, `maintainer`, `admin`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface GetProjectProtectedBranchMergeAccessLevel {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: string;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: string;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: number;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: number;
}

export interface GetProjectProtectedBranchMergeAccessLevelArgs {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface GetProjectProtectedBranchPushAccessLevel {
    /**
     * Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: string;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: string;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: number;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: number;
}

export interface GetProjectProtectedBranchPushAccessLevelArgs {
    /**
     * Access levels allowed to push to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface GetProjectProtectedBranchesProtectedBranch {
    /**
     * Whether force push is allowed.
     */
    allowForcePush?: boolean;
    /**
     * Reject code pushes that change files listed in the CODEOWNERS file.
     */
    codeOwnerApprovalRequired?: boolean;
    /**
     * The ID of this resource.
     */
    id?: number;
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    mergeAccessLevels?: inputs.GetProjectProtectedBranchesProtectedBranchMergeAccessLevel[];
    /**
     * The name of the protected branch.
     */
    name?: string;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    pushAccessLevels?: inputs.GetProjectProtectedBranchesProtectedBranchPushAccessLevel[];
}

export interface GetProjectProtectedBranchesProtectedBranchArgs {
    /**
     * Whether force push is allowed.
     */
    allowForcePush?: pulumi.Input<boolean>;
    /**
     * Reject code pushes that change files listed in the CODEOWNERS file.
     */
    codeOwnerApprovalRequired?: pulumi.Input<boolean>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<number>;
    /**
     * Array of access levels and user(s)/group(s) allowed to merge to protected branch.
     */
    mergeAccessLevels?: pulumi.Input<pulumi.Input<inputs.GetProjectProtectedBranchesProtectedBranchMergeAccessLevelArgs>[]>;
    /**
     * The name of the protected branch.
     */
    name?: pulumi.Input<string>;
    /**
     * Array of access levels and user(s)/group(s) allowed to push to protected branch.
     */
    pushAccessLevels?: pulumi.Input<pulumi.Input<inputs.GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs>[]>;
}

export interface GetProjectProtectedBranchesProtectedBranchMergeAccessLevel {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: string;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: string;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: number;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: number;
}

export interface GetProjectProtectedBranchesProtectedBranchMergeAccessLevelArgs {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface GetProjectProtectedBranchesProtectedBranchPushAccessLevel {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: string;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: string;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: number;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: number;
}

export interface GetProjectProtectedBranchesProtectedBranchPushAccessLevelArgs {
    /**
     * Access levels allowed to merge to protected branch. Valid values are: `no one`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of access level.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}

export interface GroupEpicBoardList {
    /**
     * The ID of the list.
     */
    id?: pulumi.Input<number>;
    /**
     * The ID of the label the list should be scoped to.
     */
    labelId?: pulumi.Input<number>;
    /**
     * The position of the list within the board. The position for the list is sed on the its position in the `lists` array.
     */
    position?: pulumi.Input<number>;
}

export interface GroupIssueBoardList {
    /**
     * The ID of the list.
     */
    id?: pulumi.Input<number>;
    /**
     * The ID of the label the list should be scoped to.
     */
    labelId?: pulumi.Input<number>;
    /**
     * The explicit position of the list within the board, zero based.
     */
    position?: pulumi.Input<number>;
}

export interface GroupProtectedEnvironmentApprovalRule {
    /**
     * Levels of access allowed to approve a deployment to this protected environment. Valid values are `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of level of access.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of the group allowed to approve a deployment to this protected environment. TThe group must be a sub-group under the given group. This is mutually exclusive with user_id.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The unique ID of the Approval Rules object.
     */
    id?: pulumi.Input<number>;
    /**
     * The number of approval required to allow deployment to this protected environment. This is mutually exclusive with user_id.
     */
    requiredApprovals?: pulumi.Input<number>;
    /**
     * The ID of the user allowed to approve a deployment to this protected environment. The user must be a member of the group with Maintainer role or higher. This is mutually exclusive with group*id and required*approvals.
     */
    userId?: pulumi.Input<number>;
}

export interface GroupProtectedEnvironmentDeployAccessLevel {
    /**
     * Levels of access required to deploy to this protected environment. Valid values are `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of level of access.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of the group allowed to deploy to this protected environment. The group must be a sub-group under the given group.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The unique ID of the Deploy Access Level object.
     */
    id?: pulumi.Input<number>;
    /**
     * The ID of the user allowed to deploy to this protected environment. The user must be a member of the group with Maintainer role or higher.
     */
    userId?: pulumi.Input<number>;
}

export interface GroupPushRules {
    /**
     * All commit author emails must match this regex, e.g. `@my-company.com$`.
     */
    authorEmailRegex?: pulumi.Input<string>;
    /**
     * All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
     */
    branchNameRegex?: pulumi.Input<string>;
    /**
     * Only commits pushed using verified emails are allowed.  **Note** This attribute is only supported in GitLab versions >= 16.4.
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
     * Filenames matching the regular expression provided in this attribute are not allowed, for example, `(jar|exe)$`.
     */
    fileNameRegex?: pulumi.Input<string>;
    /**
     * Maximum file size (MB) allowed.
     */
    maxFileSize?: pulumi.Input<number>;
    /**
     * Allows only GitLab users to author commits.
     */
    memberCheck?: pulumi.Input<boolean>;
    /**
     * GitLab will reject any files that are likely to contain secrets.
     */
    preventSecrets?: pulumi.Input<boolean>;
    /**
     * Only commits signed through GPG are allowed.  **Note** This attribute is only supported in GitLab versions >= 16.4.
     */
    rejectUnsignedCommits?: pulumi.Input<boolean>;
}

export interface ProjectContainerExpirationPolicy {
    /**
     * The cadence of the policy. Valid values are: `1d`, `7d`, `14d`, `1month`, `3month`.
     */
    cadence?: pulumi.Input<string>;
    /**
     * If true, the policy is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The number of images to keep.
     */
    keepN?: pulumi.Input<number>;
    /**
     * The regular expression to match image names to delete.
     *
     * @deprecated `name_regex` has been deprecated. Use `name_regex_delete` instead.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The regular expression to match image names to delete.
     */
    nameRegexDelete?: pulumi.Input<string>;
    /**
     * The regular expression to match image names to keep.
     */
    nameRegexKeep?: pulumi.Input<string>;
    /**
     * The next time the policy will run.
     */
    nextRunAt?: pulumi.Input<string>;
    /**
     * The number of days to keep images.
     */
    olderThan?: pulumi.Input<string>;
}

export interface ProjectIssueBoardList {
    /**
     * The ID of the assignee the list should be scoped to. Requires a GitLab EE license.
     */
    assigneeId?: pulumi.Input<number>;
    /**
     * The ID of the list
     */
    id?: pulumi.Input<number>;
    /**
     * The ID of the iteration the list should be scoped to. Requires a GitLab EE license.
     */
    iterationId?: pulumi.Input<number>;
    /**
     * The ID of the label the list should be scoped to. Requires a GitLab EE license.
     */
    labelId?: pulumi.Input<number>;
    /**
     * The ID of the milestone the list should be scoped to. Requires a GitLab EE license.
     */
    milestoneId?: pulumi.Input<number>;
    /**
     * The position of the list within the board. The position for the list is based on the its position in the `lists` array.
     */
    position?: pulumi.Input<number>;
}

export interface ProjectIssueTaskCompletionStatus {
    completedCount?: pulumi.Input<number>;
    count?: pulumi.Input<number>;
}

export interface ProjectProtectedEnvironmentApprovalRule {
    /**
     * Levels of access allowed to approve a deployment to this protected environment. Valid values are `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of level of access.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of the group allowed to approve a deployment to this protected environment. The project must be shared with the group. This is mutually exclusive with user_id.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The unique ID of the Approval Rules object.
     */
    id?: pulumi.Input<number>;
    /**
     * The number of approval required to allow deployment to this protected environment. This is mutually exclusive with user_id.
     */
    requiredApprovals?: pulumi.Input<number>;
    /**
     * The ID of the user allowed to approve a deployment to this protected environment. The user must be a member of the project. This is mutually exclusive with group*id and required*approvals.
     */
    userId?: pulumi.Input<number>;
}

export interface ProjectProtectedEnvironmentDeployAccessLevel {
    /**
     * Levels of access required to deploy to this protected environment. Valid values are `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of level of access.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of the group allowed to deploy to this protected environment. The project must be shared with the group.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The unique ID of the Deploy Access Level object.
     */
    id?: pulumi.Input<number>;
    /**
     * The ID of the user allowed to deploy to this protected environment. The user must be a member of the project.
     */
    userId?: pulumi.Input<number>;
}

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
     * All committed filenames must not match this regex, e.g. `(jar|exe)$`.
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

export interface ProjectTagCommit {
    authorEmail?: pulumi.Input<string>;
    authorName?: pulumi.Input<string>;
    authoredDate?: pulumi.Input<string>;
    committedDate?: pulumi.Input<string>;
    committerEmail?: pulumi.Input<string>;
    committerName?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    message?: pulumi.Input<string>;
    parentIds?: pulumi.Input<pulumi.Input<string>[]>;
    shortId?: pulumi.Input<string>;
    title?: pulumi.Input<string>;
}

export interface ProjectTagRelease {
    description?: pulumi.Input<string>;
    tagName?: pulumi.Input<string>;
}

export interface TagProtectionAllowedToCreate {
    /**
     * Level of access.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * Readable description of level of access.
     */
    accessLevelDescription?: pulumi.Input<string>;
    /**
     * The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `userId`.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `groupId`.
     */
    userId?: pulumi.Input<number>;
}
