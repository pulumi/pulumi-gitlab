// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectContainerExpirationPolicy {
    private String cadence;
    private Boolean enabled;
    private Integer keepN;
    /**
     * @deprecated
     * `name_regex` has been deprecated. Use `name_regex_delete` instead.
     * 
     */
    @Deprecated /* `name_regex` has been deprecated. Use `name_regex_delete` instead. */
    private String nameRegex;
    private String nameRegexDelete;
    private String nameRegexKeep;
    private String nextRunAt;
    private String olderThan;

    private GetProjectContainerExpirationPolicy() {}
    public String cadence() {
        return this.cadence;
    }
    public Boolean enabled() {
        return this.enabled;
    }
    public Integer keepN() {
        return this.keepN;
    }
    /**
     * @deprecated
     * `name_regex` has been deprecated. Use `name_regex_delete` instead.
     * 
     */
    @Deprecated /* `name_regex` has been deprecated. Use `name_regex_delete` instead. */
    public String nameRegex() {
        return this.nameRegex;
    }
    public String nameRegexDelete() {
        return this.nameRegexDelete;
    }
    public String nameRegexKeep() {
        return this.nameRegexKeep;
    }
    public String nextRunAt() {
        return this.nextRunAt;
    }
    public String olderThan() {
        return this.olderThan;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectContainerExpirationPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cadence;
        private Boolean enabled;
        private Integer keepN;
        private String nameRegex;
        private String nameRegexDelete;
        private String nameRegexKeep;
        private String nextRunAt;
        private String olderThan;
        public Builder() {}
        public Builder(GetProjectContainerExpirationPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cadence = defaults.cadence;
    	      this.enabled = defaults.enabled;
    	      this.keepN = defaults.keepN;
    	      this.nameRegex = defaults.nameRegex;
    	      this.nameRegexDelete = defaults.nameRegexDelete;
    	      this.nameRegexKeep = defaults.nameRegexKeep;
    	      this.nextRunAt = defaults.nextRunAt;
    	      this.olderThan = defaults.olderThan;
        }

        @CustomType.Setter
        public Builder cadence(String cadence) {
            this.cadence = Objects.requireNonNull(cadence);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder keepN(Integer keepN) {
            this.keepN = Objects.requireNonNull(keepN);
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(String nameRegex) {
            this.nameRegex = Objects.requireNonNull(nameRegex);
            return this;
        }
        @CustomType.Setter
        public Builder nameRegexDelete(String nameRegexDelete) {
            this.nameRegexDelete = Objects.requireNonNull(nameRegexDelete);
            return this;
        }
        @CustomType.Setter
        public Builder nameRegexKeep(String nameRegexKeep) {
            this.nameRegexKeep = Objects.requireNonNull(nameRegexKeep);
            return this;
        }
        @CustomType.Setter
        public Builder nextRunAt(String nextRunAt) {
            this.nextRunAt = Objects.requireNonNull(nextRunAt);
            return this;
        }
        @CustomType.Setter
        public Builder olderThan(String olderThan) {
            this.olderThan = Objects.requireNonNull(olderThan);
            return this;
        }
        public GetProjectContainerExpirationPolicy build() {
            final var _resultValue = new GetProjectContainerExpirationPolicy();
            _resultValue.cadence = cadence;
            _resultValue.enabled = enabled;
            _resultValue.keepN = keepN;
            _resultValue.nameRegex = nameRegex;
            _resultValue.nameRegexDelete = nameRegexDelete;
            _resultValue.nameRegexKeep = nameRegexKeep;
            _resultValue.nextRunAt = nextRunAt;
            _resultValue.olderThan = olderThan;
            return _resultValue;
        }
    }
}
