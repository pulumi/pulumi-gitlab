// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupIssueBoardArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupIssueBoardState;
import com.pulumi.gitlab.outputs.GroupIssueBoardList;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupIssueBoard` resource allows to manage the lifecycle of a issue board in a group.
 * 
 * &gt; Multiple issue boards on one group requires a GitLab Premium or above License.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)
 * 
 */
@ResourceType(type="gitlab:index/groupIssueBoard:GroupIssueBoard")
public class GroupIssueBoard extends com.pulumi.resources.CustomResource {
    /**
     * The ID or URL-encoded path of the group owned by the authenticated user.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group owned by the authenticated user.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * The list of label names which the board should be scoped to.
     * 
     */
    @Export(name="labels", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> labels;

    /**
     * @return The list of label names which the board should be scoped to.
     * 
     */
    public Output<Optional<List<String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The list of issue board lists.
     * 
     */
    @Export(name="lists", refs={List.class,GroupIssueBoardList.class}, tree="[0,1]")
    private Output</* @Nullable */ List<GroupIssueBoardList>> lists;

    /**
     * @return The list of issue board lists.
     * 
     */
    public Output<Optional<List<GroupIssueBoardList>>> lists() {
        return Codegen.optional(this.lists);
    }
    /**
     * The milestone the board should be scoped to.
     * 
     */
    @Export(name="milestoneId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> milestoneId;

    /**
     * @return The milestone the board should be scoped to.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupIssueBoard(String name) {
        this(name, GroupIssueBoardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupIssueBoard(String name, GroupIssueBoardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupIssueBoard(String name, GroupIssueBoardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupIssueBoard:GroupIssueBoard", name, args == null ? GroupIssueBoardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupIssueBoard(String name, Output<String> id, @Nullable GroupIssueBoardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupIssueBoard:GroupIssueBoard", name, state, makeResourceOptions(options, id));
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
    public static GroupIssueBoard get(String name, Output<String> id, @Nullable GroupIssueBoardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupIssueBoard(name, id, state, options);
    }
}
