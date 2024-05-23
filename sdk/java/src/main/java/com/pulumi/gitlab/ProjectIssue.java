// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectIssueArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectIssueState;
import com.pulumi.gitlab.outputs.ProjectIssueTaskCompletionStatus;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gitlab.Project;
 * import com.pulumi.gitlab.ProjectArgs;
 * import com.pulumi.gitlab.ProjectIssue;
 * import com.pulumi.gitlab.ProjectIssueArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var foo = new Project("foo", ProjectArgs.builder()
 *             .name("example project")
 *             .description("Lorem Ipsum")
 *             .visibilityLevel("public")
 *             .build());
 * 
 *         var welcomeIssue = new ProjectIssue("welcomeIssue", ProjectIssueArgs.builder()
 *             .project(foo.id())
 *             .title("Welcome!")
 *             .description(foo.name().applyValue(name -> """
 *   Welcome to the %s project!
 * 
 * ", name)))
 *             .discussionLocked(true)
 *             .build());
 * 
 *         ctx.export("welcomeIssueWebUrl", webUrl);
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * You can import this resource with an id made up of `{project-id}:{issue-id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectIssue:ProjectIssue welcome_issue 42:1
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectIssue:ProjectIssue")
public class ProjectIssue extends com.pulumi.resources.CustomResource {
    /**
     * The IDs of the users to assign the issue to.
     * 
     */
    @Export(name="assigneeIds", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> assigneeIds;

    /**
     * @return The IDs of the users to assign the issue to.
     * 
     */
    public Output<Optional<List<Integer>>> assigneeIds() {
        return Codegen.optional(this.assigneeIds);
    }
    /**
     * The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    @Export(name="authorId", refs={Integer.class}, tree="[0]")
    private Output<Integer> authorId;

    /**
     * @return The ID of the author of the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    public Output<Integer> authorId() {
        return this.authorId;
    }
    /**
     * When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    @Export(name="closedAt", refs={String.class}, tree="[0]")
    private Output<String> closedAt;

    /**
     * @return When the issue was closed. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public Output<String> closedAt() {
        return this.closedAt;
    }
    /**
     * The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    @Export(name="closedByUserId", refs={Integer.class}, tree="[0]")
    private Output<Integer> closedByUserId;

    /**
     * @return The ID of the user that closed the issue. Use `gitlab.User` data source to get more information about the user.
     * 
     */
    public Output<Integer> closedByUserId() {
        return this.closedByUserId;
    }
    /**
     * Set an issue to be confidential.
     * 
     */
    @Export(name="confidential", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> confidential;

    /**
     * @return Set an issue to be confidential.
     * 
     */
    public Output<Optional<Boolean>> confidential() {
        return Codegen.optional(this.confidential);
    }
    /**
     * When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Whether the issue is deleted instead of closed during destroy.
     * 
     */
    @Export(name="deleteOnDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteOnDestroy;

    /**
     * @return Whether the issue is deleted instead of closed during destroy.
     * 
     */
    public Output<Optional<Boolean>> deleteOnDestroy() {
        return Codegen.optional(this.deleteOnDestroy);
    }
    /**
     * The description of an issue. Limited to 1,048,576 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of an issue. Limited to 1,048,576 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether the issue is locked for discussions or not.
     * 
     */
    @Export(name="discussionLocked", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> discussionLocked;

    /**
     * @return Whether the issue is locked for discussions or not.
     * 
     */
    public Output<Optional<Boolean>> discussionLocked() {
        return Codegen.optional(this.discussionLocked);
    }
    /**
     * The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     * 
     */
    @Export(name="discussionToResolve", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> discussionToResolve;

    /**
     * @return The ID of a discussion to resolve. This fills out the issue with a default description and mark the discussion as resolved. Use in combination with merge*request*to*resolve*discussions_of.
     * 
     */
    public Output<Optional<String>> discussionToResolve() {
        return Codegen.optional(this.discussionToResolve);
    }
    /**
     * The number of downvotes the issue has received.
     * 
     */
    @Export(name="downvotes", refs={Integer.class}, tree="[0]")
    private Output<Integer> downvotes;

    /**
     * @return The number of downvotes the issue has received.
     * 
     */
    public Output<Integer> downvotes() {
        return this.downvotes;
    }
    /**
     * The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    @Export(name="dueDate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dueDate;

    /**
     * @return The due date. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
     * 
     */
    public Output<Optional<String>> dueDate() {
        return Codegen.optional(this.dueDate);
    }
    /**
     * ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     * 
     */
    @Export(name="epicId", refs={Integer.class}, tree="[0]")
    private Output<Integer> epicId;

    /**
     * @return ID of the epic to add the issue to. Valid values are greater than or equal to 0.
     * 
     */
    public Output<Integer> epicId() {
        return this.epicId;
    }
    /**
     * The ID of the epic issue.
     * 
     */
    @Export(name="epicIssueId", refs={Integer.class}, tree="[0]")
    private Output<Integer> epicIssueId;

    /**
     * @return The ID of the epic issue.
     * 
     */
    public Output<Integer> epicIssueId() {
        return this.epicIssueId;
    }
    /**
     * The external ID of the issue.
     * 
     */
    @Export(name="externalId", refs={String.class}, tree="[0]")
    private Output<String> externalId;

    /**
     * @return The external ID of the issue.
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
    }
    /**
     * The human-readable time estimate of the issue.
     * 
     */
    @Export(name="humanTimeEstimate", refs={String.class}, tree="[0]")
    private Output<String> humanTimeEstimate;

    /**
     * @return The human-readable time estimate of the issue.
     * 
     */
    public Output<String> humanTimeEstimate() {
        return this.humanTimeEstimate;
    }
    /**
     * The human-readable total time spent of the issue.
     * 
     */
    @Export(name="humanTotalTimeSpent", refs={String.class}, tree="[0]")
    private Output<String> humanTotalTimeSpent;

    /**
     * @return The human-readable total time spent of the issue.
     * 
     */
    public Output<String> humanTotalTimeSpent() {
        return this.humanTotalTimeSpent;
    }
    /**
     * The internal ID of the project&#39;s issue.
     * 
     */
    @Export(name="iid", refs={Integer.class}, tree="[0]")
    private Output<Integer> iid;

    /**
     * @return The internal ID of the project&#39;s issue.
     * 
     */
    public Output<Integer> iid() {
        return this.iid;
    }
    /**
     * The instance-wide ID of the issue.
     * 
     */
    @Export(name="issueId", refs={Integer.class}, tree="[0]")
    private Output<Integer> issueId;

    /**
     * @return The instance-wide ID of the issue.
     * 
     */
    public Output<Integer> issueId() {
        return this.issueId;
    }
    /**
     * The ID of the issue link.
     * 
     */
    @Export(name="issueLinkId", refs={Integer.class}, tree="[0]")
    private Output<Integer> issueLinkId;

    /**
     * @return The ID of the issue link.
     * 
     */
    public Output<Integer> issueLinkId() {
        return this.issueLinkId;
    }
    /**
     * The type of issue. Valid values are: `issue`, `incident`, `test_case`.
     * 
     */
    @Export(name="issueType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issueType;

    /**
     * @return The type of issue. Valid values are: `issue`, `incident`, `test_case`.
     * 
     */
    public Output<Optional<String>> issueType() {
        return Codegen.optional(this.issueType);
    }
    /**
     * The labels of an issue.
     * 
     */
    @Export(name="labels", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> labels;

    /**
     * @return The labels of an issue.
     * 
     */
    public Output<Optional<List<String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The links of the issue.
     * 
     */
    @Export(name="links", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> links;

    /**
     * @return The links of the issue.
     * 
     */
    public Output<Map<String,String>> links() {
        return this.links;
    }
    /**
     * The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     * 
     */
    @Export(name="mergeRequestToResolveDiscussionsOf", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> mergeRequestToResolveDiscussionsOf;

    /**
     * @return The IID of a merge request in which to resolve all issues. This fills out the issue with a default description and mark all discussions as resolved. When passing a description or title, these values take precedence over the default values.
     * 
     */
    public Output<Optional<Integer>> mergeRequestToResolveDiscussionsOf() {
        return Codegen.optional(this.mergeRequestToResolveDiscussionsOf);
    }
    /**
     * The number of merge requests associated with the issue.
     * 
     */
    @Export(name="mergeRequestsCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> mergeRequestsCount;

    /**
     * @return The number of merge requests associated with the issue.
     * 
     */
    public Output<Integer> mergeRequestsCount() {
        return this.mergeRequestsCount;
    }
    /**
     * The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue&#39;s details.
     * 
     */
    @Export(name="milestoneId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> milestoneId;

    /**
     * @return The global ID of a milestone to assign issue. To find the milestone_id associated with a milestone, view an issue with the milestone assigned and use the API to retrieve the issue&#39;s details.
     * 
     */
    public Output<Optional<Integer>> milestoneId() {
        return Codegen.optional(this.milestoneId);
    }
    /**
     * The ID of the issue that was moved to.
     * 
     */
    @Export(name="movedToId", refs={Integer.class}, tree="[0]")
    private Output<Integer> movedToId;

    /**
     * @return The ID of the issue that was moved to.
     * 
     */
    public Output<Integer> movedToId() {
        return this.movedToId;
    }
    /**
     * The name or ID of the project.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name or ID of the project.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The references of the issue.
     * 
     */
    @Export(name="references", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> references;

    /**
     * @return The references of the issue.
     * 
     */
    public Output<Map<String,String>> references() {
        return this.references;
    }
    /**
     * The state of the issue. Valid values are: `opened`, `closed`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state of the issue. Valid values are: `opened`, `closed`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * Whether the authenticated user is subscribed to the issue or not.
     * 
     */
    @Export(name="subscribed", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> subscribed;

    /**
     * @return Whether the authenticated user is subscribed to the issue or not.
     * 
     */
    public Output<Boolean> subscribed() {
        return this.subscribed;
    }
    /**
     * The task completion status. It&#39;s always a one element list.
     * 
     */
    @Export(name="taskCompletionStatuses", refs={List.class,ProjectIssueTaskCompletionStatus.class}, tree="[0,1]")
    private Output<List<ProjectIssueTaskCompletionStatus>> taskCompletionStatuses;

    /**
     * @return The task completion status. It&#39;s always a one element list.
     * 
     */
    public Output<List<ProjectIssueTaskCompletionStatus>> taskCompletionStatuses() {
        return this.taskCompletionStatuses;
    }
    /**
     * The time estimate of the issue.
     * 
     */
    @Export(name="timeEstimate", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeEstimate;

    /**
     * @return The time estimate of the issue.
     * 
     */
    public Output<Integer> timeEstimate() {
        return this.timeEstimate;
    }
    /**
     * The title of the issue.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return The title of the issue.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * The total time spent of the issue.
     * 
     */
    @Export(name="totalTimeSpent", refs={Integer.class}, tree="[0]")
    private Output<Integer> totalTimeSpent;

    /**
     * @return The total time spent of the issue.
     * 
     */
    public Output<Integer> totalTimeSpent() {
        return this.totalTimeSpent;
    }
    /**
     * When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return When the issue was updated. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * The number of upvotes the issue has received.
     * 
     */
    @Export(name="upvotes", refs={Integer.class}, tree="[0]")
    private Output<Integer> upvotes;

    /**
     * @return The number of upvotes the issue has received.
     * 
     */
    public Output<Integer> upvotes() {
        return this.upvotes;
    }
    /**
     * The number of user notes on the issue.
     * 
     */
    @Export(name="userNotesCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> userNotesCount;

    /**
     * @return The number of user notes on the issue.
     * 
     */
    public Output<Integer> userNotesCount() {
        return this.userNotesCount;
    }
    /**
     * The web URL of the issue.
     * 
     */
    @Export(name="webUrl", refs={String.class}, tree="[0]")
    private Output<String> webUrl;

    /**
     * @return The web URL of the issue.
     * 
     */
    public Output<String> webUrl() {
        return this.webUrl;
    }
    /**
     * The weight of the issue. Valid values are greater than or equal to 0.
     * 
     */
    @Export(name="weight", refs={Integer.class}, tree="[0]")
    private Output<Integer> weight;

    /**
     * @return The weight of the issue. Valid values are greater than or equal to 0.
     * 
     */
    public Output<Integer> weight() {
        return this.weight;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectIssue(String name) {
        this(name, ProjectIssueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectIssue(String name, ProjectIssueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectIssue(String name, ProjectIssueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectIssue:ProjectIssue", name, args == null ? ProjectIssueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectIssue(String name, Output<String> id, @Nullable ProjectIssueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectIssue:ProjectIssue", name, state, makeResourceOptions(options, id));
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
    public static ProjectIssue get(String name, Output<String> id, @Nullable ProjectIssueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectIssue(name, id, state, options);
    }
}
