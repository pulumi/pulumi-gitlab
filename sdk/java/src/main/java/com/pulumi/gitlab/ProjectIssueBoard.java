// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectIssueBoardArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectIssueBoardState;
import com.pulumi.gitlab.outputs.ProjectIssueBoardList;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectIssueBoard` resource allows to manage the lifecycle of a Project Issue Board.
 * 
 * &gt; **NOTE:** If the board lists are changed all lists will be recreated.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/boards/)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_issue_board`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_issue_board.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * You can import this resource with an id made up of `{project-id}:{issue-board-id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectIssueBoard:ProjectIssueBoard kanban 42:1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectIssueBoard:ProjectIssueBoard")
public class ProjectIssueBoard extends com.pulumi.resources.CustomResource {
    /**
     * The assignee the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    @Export(name="assigneeId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> assigneeId;

    /**
     * @return The assignee the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Output<Optional<Integer>> assigneeId() {
        return Codegen.optional(this.assigneeId);
    }
    /**
     * The list of label names which the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    @Export(name="labels", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> labels;

    /**
     * @return The list of label names which the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Output<Optional<List<String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The list of issue board lists
     * 
     */
    @Export(name="lists", refs={List.class,ProjectIssueBoardList.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ProjectIssueBoardList>> lists;

    /**
     * @return The list of issue board lists
     * 
     */
    public Output<Optional<List<ProjectIssueBoardList>>> lists() {
        return Codegen.optional(this.lists);
    }
    /**
     * The milestone the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    @Export(name="milestoneId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> milestoneId;

    /**
     * @return The milestone the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Output<Optional<Integer>> milestoneId() {
        return Codegen.optional(this.milestoneId);
    }
    /**
     * The name of the board.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the board.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID or full path of the project maintained by the authenticated user.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project maintained by the authenticated user.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    @Export(name="weight", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> weight;

    /**
     * @return The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
     * 
     */
    public Output<Optional<Integer>> weight() {
        return Codegen.optional(this.weight);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectIssueBoard(String name) {
        this(name, ProjectIssueBoardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectIssueBoard(String name, ProjectIssueBoardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectIssueBoard(String name, ProjectIssueBoardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectIssueBoard:ProjectIssueBoard", name, args == null ? ProjectIssueBoardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectIssueBoard(String name, Output<String> id, @Nullable ProjectIssueBoardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectIssueBoard:ProjectIssueBoard", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static ProjectIssueBoard get(String name, Output<String> id, @Nullable ProjectIssueBoardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectIssueBoard(name, id, state, options);
    }
}
