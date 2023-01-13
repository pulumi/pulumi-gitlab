// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetBranchCommit;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBranchResult {
    /**
     * @return Bool, true if you can push to the branch.
     * 
     */
    private Boolean canPush;
    /**
     * @return The commit associated with the branch ref.
     * 
     */
    private List<GetBranchCommit> commits;
    /**
     * @return Bool, true if branch is the default branch for the project.
     * 
     */
    private Boolean default_;
    /**
     * @return Bool, true if developer level access allows to merge branch.
     * 
     */
    private Boolean developerCanMerge;
    /**
     * @return Bool, true if developer level access allows git push.
     * 
     */
    private Boolean developerCanPush;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Bool, true if the branch has been merged into it&#39;s parent.
     * 
     */
    private Boolean merged;
    /**
     * @return The name of the branch.
     * 
     */
    private String name;
    /**
     * @return The full path or id of the project.
     * 
     */
    private String project;
    /**
     * @return Bool, true if branch has branch protection.
     * 
     */
    private Boolean protected_;
    /**
     * @return The url of the created branch (https.)
     * 
     */
    private String webUrl;

    private GetBranchResult() {}
    /**
     * @return Bool, true if you can push to the branch.
     * 
     */
    public Boolean canPush() {
        return this.canPush;
    }
    /**
     * @return The commit associated with the branch ref.
     * 
     */
    public List<GetBranchCommit> commits() {
        return this.commits;
    }
    /**
     * @return Bool, true if branch is the default branch for the project.
     * 
     */
    public Boolean default_() {
        return this.default_;
    }
    /**
     * @return Bool, true if developer level access allows to merge branch.
     * 
     */
    public Boolean developerCanMerge() {
        return this.developerCanMerge;
    }
    /**
     * @return Bool, true if developer level access allows git push.
     * 
     */
    public Boolean developerCanPush() {
        return this.developerCanPush;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Bool, true if the branch has been merged into it&#39;s parent.
     * 
     */
    public Boolean merged() {
        return this.merged;
    }
    /**
     * @return The name of the branch.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The full path or id of the project.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return Bool, true if branch has branch protection.
     * 
     */
    public Boolean protected_() {
        return this.protected_;
    }
    /**
     * @return The url of the created branch (https.)
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBranchResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean canPush;
        private List<GetBranchCommit> commits;
        private Boolean default_;
        private Boolean developerCanMerge;
        private Boolean developerCanPush;
        private String id;
        private Boolean merged;
        private String name;
        private String project;
        private Boolean protected_;
        private String webUrl;
        public Builder() {}
        public Builder(GetBranchResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.canPush = defaults.canPush;
    	      this.commits = defaults.commits;
    	      this.default_ = defaults.default_;
    	      this.developerCanMerge = defaults.developerCanMerge;
    	      this.developerCanPush = defaults.developerCanPush;
    	      this.id = defaults.id;
    	      this.merged = defaults.merged;
    	      this.name = defaults.name;
    	      this.project = defaults.project;
    	      this.protected_ = defaults.protected_;
    	      this.webUrl = defaults.webUrl;
        }

        @CustomType.Setter
        public Builder canPush(Boolean canPush) {
            this.canPush = Objects.requireNonNull(canPush);
            return this;
        }
        @CustomType.Setter
        public Builder commits(List<GetBranchCommit> commits) {
            this.commits = Objects.requireNonNull(commits);
            return this;
        }
        public Builder commits(GetBranchCommit... commits) {
            return commits(List.of(commits));
        }
        @CustomType.Setter("default")
        public Builder default_(Boolean default_) {
            this.default_ = Objects.requireNonNull(default_);
            return this;
        }
        @CustomType.Setter
        public Builder developerCanMerge(Boolean developerCanMerge) {
            this.developerCanMerge = Objects.requireNonNull(developerCanMerge);
            return this;
        }
        @CustomType.Setter
        public Builder developerCanPush(Boolean developerCanPush) {
            this.developerCanPush = Objects.requireNonNull(developerCanPush);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder merged(Boolean merged) {
            this.merged = Objects.requireNonNull(merged);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        @CustomType.Setter("protected")
        public Builder protected_(Boolean protected_) {
            this.protected_ = Objects.requireNonNull(protected_);
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            this.webUrl = Objects.requireNonNull(webUrl);
            return this;
        }
        public GetBranchResult build() {
            final var o = new GetBranchResult();
            o.canPush = canPush;
            o.commits = commits;
            o.default_ = default_;
            o.developerCanMerge = developerCanMerge;
            o.developerCanPush = developerCanPush;
            o.id = id;
            o.merged = merged;
            o.name = name;
            o.project = project;
            o.protected_ = protected_;
            o.webUrl = webUrl;
            return o;
        }
    }
}