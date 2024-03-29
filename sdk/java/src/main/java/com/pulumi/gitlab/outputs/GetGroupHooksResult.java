// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetGroupHooksHook;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetGroupHooksResult {
    /**
     * @return The ID or full path of the group.
     * 
     */
    private String group;
    /**
     * @return The list of hooks.
     * 
     */
    private List<GetGroupHooksHook> hooks;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetGroupHooksResult() {}
    /**
     * @return The ID or full path of the group.
     * 
     */
    public String group() {
        return this.group;
    }
    /**
     * @return The list of hooks.
     * 
     */
    public List<GetGroupHooksHook> hooks() {
        return this.hooks;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupHooksResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String group;
        private List<GetGroupHooksHook> hooks;
        private String id;
        public Builder() {}
        public Builder(GetGroupHooksResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.group = defaults.group;
    	      this.hooks = defaults.hooks;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder group(String group) {
            if (group == null) {
              throw new MissingRequiredPropertyException("GetGroupHooksResult", "group");
            }
            this.group = group;
            return this;
        }
        @CustomType.Setter
        public Builder hooks(List<GetGroupHooksHook> hooks) {
            if (hooks == null) {
              throw new MissingRequiredPropertyException("GetGroupHooksResult", "hooks");
            }
            this.hooks = hooks;
            return this;
        }
        public Builder hooks(GetGroupHooksHook... hooks) {
            return hooks(List.of(hooks));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupHooksResult", "id");
            }
            this.id = id;
            return this;
        }
        public GetGroupHooksResult build() {
            final var _resultValue = new GetGroupHooksResult();
            _resultValue.group = group;
            _resultValue.hooks = hooks;
            _resultValue.id = id;
            return _resultValue;
        }
    }
}
