// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gitlab.inputs.GroupIssueBoardListArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupIssueBoardState extends com.pulumi.resources.ResourceArgs {

    public static final GroupIssueBoardState Empty = new GroupIssueBoardState();

    /**
     * The ID or URL-encoded path of the group owned by the authenticated user.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return The ID or URL-encoded path of the group owned by the authenticated user.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * The list of label names which the board should be scoped to.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<String>> labels;

    /**
     * @return The list of label names which the board should be scoped to.
     * 
     */
    public Optional<Output<List<String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The list of issue board lists.
     * 
     */
    @Import(name="lists")
    private @Nullable Output<List<GroupIssueBoardListArgs>> lists;

    /**
     * @return The list of issue board lists.
     * 
     */
    public Optional<Output<List<GroupIssueBoardListArgs>>> lists() {
        return Optional.ofNullable(this.lists);
    }

    /**
     * The milestone the board should be scoped to.
     * 
     */
    @Import(name="milestoneId")
    private @Nullable Output<Integer> milestoneId;

    /**
     * @return The milestone the board should be scoped to.
     * 
     */
    public Optional<Output<Integer>> milestoneId() {
        return Optional.ofNullable(this.milestoneId);
    }

    /**
     * The name of the board.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the board.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GroupIssueBoardState() {}

    private GroupIssueBoardState(GroupIssueBoardState $) {
        this.group = $.group;
        this.labels = $.labels;
        this.lists = $.lists;
        this.milestoneId = $.milestoneId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupIssueBoardState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupIssueBoardState $;

        public Builder() {
            $ = new GroupIssueBoardState();
        }

        public Builder(GroupIssueBoardState defaults) {
            $ = new GroupIssueBoardState(Objects.requireNonNull(defaults));
        }

        /**
         * @param group The ID or URL-encoded path of the group owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group The ID or URL-encoded path of the group owned by the authenticated user.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param labels The list of label names which the board should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels The list of label names which the board should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder labels(List<String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels The list of label names which the board should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder labels(String... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param lists The list of issue board lists.
         * 
         * @return builder
         * 
         */
        public Builder lists(@Nullable Output<List<GroupIssueBoardListArgs>> lists) {
            $.lists = lists;
            return this;
        }

        /**
         * @param lists The list of issue board lists.
         * 
         * @return builder
         * 
         */
        public Builder lists(List<GroupIssueBoardListArgs> lists) {
            return lists(Output.of(lists));
        }

        /**
         * @param lists The list of issue board lists.
         * 
         * @return builder
         * 
         */
        public Builder lists(GroupIssueBoardListArgs... lists) {
            return lists(List.of(lists));
        }

        /**
         * @param milestoneId The milestone the board should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder milestoneId(@Nullable Output<Integer> milestoneId) {
            $.milestoneId = milestoneId;
            return this;
        }

        /**
         * @param milestoneId The milestone the board should be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder milestoneId(Integer milestoneId) {
            return milestoneId(Output.of(milestoneId));
        }

        /**
         * @param name The name of the board.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the board.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GroupIssueBoardState build() {
            return $;
        }
    }

}
