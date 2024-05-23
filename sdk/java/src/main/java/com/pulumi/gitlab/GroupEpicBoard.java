// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.GroupEpicBoardArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.GroupEpicBoardState;
import com.pulumi.gitlab.outputs.GroupEpicBoardList;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `gitlab.GroupEpicBoard` resource allows to manage the lifecycle of a epic board in a group.
 * 
 * &gt; Multiple epic boards on one group requires a GitLab Premium or above License.
 * 
 * **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_boards.html)
 * 
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
 * import com.pulumi.gitlab.Group;
 * import com.pulumi.gitlab.GroupArgs;
 * import com.pulumi.gitlab.GroupLabel;
 * import com.pulumi.gitlab.GroupLabelArgs;
 * import com.pulumi.gitlab.GroupEpicBoard;
 * import com.pulumi.gitlab.GroupEpicBoardArgs;
 * import com.pulumi.gitlab.inputs.GroupEpicBoardListArgs;
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
 *         var example = new Group("example", GroupArgs.builder()
 *             .name("test_group")
 *             .path("test_group")
 *             .description("An example group")
 *             .build());
 * 
 *         var label1 = new GroupLabel("label1", GroupLabelArgs.builder()
 *             .group(example.id())
 *             .color("#FF0000")
 *             .name("red-label")
 *             .build());
 * 
 *         var label3 = new GroupLabel("label3", GroupLabelArgs.builder()
 *             .group(example.id())
 *             .name("label-3")
 *             .color("#003000")
 *             .build());
 * 
 *         var epicBoard = new GroupEpicBoard("epicBoard", GroupEpicBoardArgs.builder()
 *             .name("epic board 6")
 *             .group(example.path())
 *             .lists(GroupEpicBoardListArgs.builder()
 *                 .labelId(label1.labelId())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * You can import this resource with an id made up of `{group-id}:{epic-board-id}`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/groupEpicBoard:GroupEpicBoard agile 70:156
 * ```
 * 
 */
@ResourceType(type="gitlab:index/groupEpicBoard:GroupEpicBoard")
public class GroupEpicBoard extends com.pulumi.resources.CustomResource {
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
     * The list of epic board lists.
     * 
     */
    @Export(name="lists", refs={List.class,GroupEpicBoardList.class}, tree="[0,1]")
    private Output</* @Nullable */ List<GroupEpicBoardList>> lists;

    /**
     * @return The list of epic board lists.
     * 
     */
    public Output<Optional<List<GroupEpicBoardList>>> lists() {
        return Codegen.optional(this.lists);
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
    public GroupEpicBoard(String name) {
        this(name, GroupEpicBoardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupEpicBoard(String name, GroupEpicBoardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupEpicBoard(String name, GroupEpicBoardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupEpicBoard:GroupEpicBoard", name, args == null ? GroupEpicBoardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupEpicBoard(String name, Output<String> id, @Nullable GroupEpicBoardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/groupEpicBoard:GroupEpicBoard", name, state, makeResourceOptions(options, id));
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
    public static GroupEpicBoard get(String name, Output<String> id, @Nullable GroupEpicBoardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupEpicBoard(name, id, state, options);
    }
}
