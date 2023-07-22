// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * A Repository File can be imported using an id made up of `<project-id>:<branch-name>:<file-path>`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/repositoryFile:RepositoryFile this 1:main:foo/bar.txt
 * ```
 */
export class RepositoryFile extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryFileState, opts?: pulumi.CustomResourceOptions): RepositoryFile {
        return new RepositoryFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/repositoryFile:RepositoryFile';

    /**
     * Returns true if the given object is an instance of RepositoryFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryFile.__pulumiType;
    }

    /**
     * Email of the commit author.
     */
    public readonly authorEmail!: pulumi.Output<string | undefined>;
    /**
     * Name of the commit author.
     */
    public readonly authorName!: pulumi.Output<string | undefined>;
    /**
     * The blob id.
     */
    public /*out*/ readonly blobId!: pulumi.Output<string>;
    /**
     * Name of the branch to which to commit to.
     */
    public readonly branch!: pulumi.Output<string>;
    /**
     * The commit id.
     */
    public /*out*/ readonly commitId!: pulumi.Output<string>;
    /**
     * Commit message.
     */
    public readonly commitMessage!: pulumi.Output<string>;
    /**
     * File content.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * File content sha256 digest.
     */
    public /*out*/ readonly contentSha256!: pulumi.Output<string>;
    /**
     * The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
     */
    public readonly encoding!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
     */
    public readonly executeFilemode!: pulumi.Output<boolean | undefined>;
    /**
     * The filename.
     */
    public /*out*/ readonly fileName!: pulumi.Output<string>;
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     */
    public readonly filePath!: pulumi.Output<string>;
    /**
     * The last known commit id.
     */
    public /*out*/ readonly lastCommitId!: pulumi.Output<string>;
    /**
     * Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     */
    public readonly overwriteOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * The name or ID of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of branch, tag or commit.
     */
    public /*out*/ readonly ref!: pulumi.Output<string>;
    /**
     * The file size.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Name of the branch to start the new commit from.
     */
    public readonly startBranch!: pulumi.Output<string | undefined>;

    /**
     * Create a RepositoryFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryFileArgs | RepositoryFileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryFileState | undefined;
            resourceInputs["authorEmail"] = state ? state.authorEmail : undefined;
            resourceInputs["authorName"] = state ? state.authorName : undefined;
            resourceInputs["blobId"] = state ? state.blobId : undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["commitId"] = state ? state.commitId : undefined;
            resourceInputs["commitMessage"] = state ? state.commitMessage : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentSha256"] = state ? state.contentSha256 : undefined;
            resourceInputs["encoding"] = state ? state.encoding : undefined;
            resourceInputs["executeFilemode"] = state ? state.executeFilemode : undefined;
            resourceInputs["fileName"] = state ? state.fileName : undefined;
            resourceInputs["filePath"] = state ? state.filePath : undefined;
            resourceInputs["lastCommitId"] = state ? state.lastCommitId : undefined;
            resourceInputs["overwriteOnCreate"] = state ? state.overwriteOnCreate : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["ref"] = state ? state.ref : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["startBranch"] = state ? state.startBranch : undefined;
        } else {
            const args = argsOrState as RepositoryFileArgs | undefined;
            if ((!args || args.branch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'branch'");
            }
            if ((!args || args.commitMessage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commitMessage'");
            }
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.filePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filePath'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["authorEmail"] = args ? args.authorEmail : undefined;
            resourceInputs["authorName"] = args ? args.authorName : undefined;
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["commitMessage"] = args ? args.commitMessage : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["encoding"] = args ? args.encoding : undefined;
            resourceInputs["executeFilemode"] = args ? args.executeFilemode : undefined;
            resourceInputs["filePath"] = args ? args.filePath : undefined;
            resourceInputs["overwriteOnCreate"] = args ? args.overwriteOnCreate : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["startBranch"] = args ? args.startBranch : undefined;
            resourceInputs["blobId"] = undefined /*out*/;
            resourceInputs["commitId"] = undefined /*out*/;
            resourceInputs["contentSha256"] = undefined /*out*/;
            resourceInputs["fileName"] = undefined /*out*/;
            resourceInputs["lastCommitId"] = undefined /*out*/;
            resourceInputs["ref"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryFile resources.
 */
export interface RepositoryFileState {
    /**
     * Email of the commit author.
     */
    authorEmail?: pulumi.Input<string>;
    /**
     * Name of the commit author.
     */
    authorName?: pulumi.Input<string>;
    /**
     * The blob id.
     */
    blobId?: pulumi.Input<string>;
    /**
     * Name of the branch to which to commit to.
     */
    branch?: pulumi.Input<string>;
    /**
     * The commit id.
     */
    commitId?: pulumi.Input<string>;
    /**
     * Commit message.
     */
    commitMessage?: pulumi.Input<string>;
    /**
     * File content.
     */
    content?: pulumi.Input<string>;
    /**
     * File content sha256 digest.
     */
    contentSha256?: pulumi.Input<string>;
    /**
     * The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
     */
    encoding?: pulumi.Input<string>;
    /**
     * Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
     */
    executeFilemode?: pulumi.Input<boolean>;
    /**
     * The filename.
     */
    fileName?: pulumi.Input<string>;
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     */
    filePath?: pulumi.Input<string>;
    /**
     * The last known commit id.
     */
    lastCommitId?: pulumi.Input<string>;
    /**
     * Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     */
    overwriteOnCreate?: pulumi.Input<boolean>;
    /**
     * The name or ID of the project.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of branch, tag or commit.
     */
    ref?: pulumi.Input<string>;
    /**
     * The file size.
     */
    size?: pulumi.Input<number>;
    /**
     * Name of the branch to start the new commit from.
     */
    startBranch?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryFile resource.
 */
export interface RepositoryFileArgs {
    /**
     * Email of the commit author.
     */
    authorEmail?: pulumi.Input<string>;
    /**
     * Name of the commit author.
     */
    authorName?: pulumi.Input<string>;
    /**
     * Name of the branch to which to commit to.
     */
    branch: pulumi.Input<string>;
    /**
     * Commit message.
     */
    commitMessage: pulumi.Input<string>;
    /**
     * File content.
     */
    content: pulumi.Input<string>;
    /**
     * The file content encoding. Default value is `base64`. Valid values are: `base64`, `text`.
     */
    encoding?: pulumi.Input<string>;
    /**
     * Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
     */
    executeFilemode?: pulumi.Input<boolean>;
    /**
     * The full path of the file. It must be relative to the root of the project without a leading slash `/` or `./`.
     */
    filePath: pulumi.Input<string>;
    /**
     * Enable overwriting existing files, defaults to `false`. This attribute is only used during `create` and must be use carefully. We suggest to use `imports` whenever possible and limit the use of this attribute for when the project was imported on the same `apply`. This attribute is not supported during a resource import.
     */
    overwriteOnCreate?: pulumi.Input<boolean>;
    /**
     * The name or ID of the project.
     */
    project: pulumi.Input<string>;
    /**
     * Name of the branch to start the new commit from.
     */
    startBranch?: pulumi.Input<string>;
}
