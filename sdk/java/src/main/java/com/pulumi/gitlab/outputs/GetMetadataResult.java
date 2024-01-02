// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gitlab.outputs.GetMetadataKas;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMetadataResult {
    /**
     * @return If the GitLab instance is an enterprise instance or not. Supported for GitLab 15.6 onwards.
     * 
     */
    private Boolean enterprise;
    /**
     * @return The id of the data source. It will always be `1`
     * 
     */
    private String id;
    /**
     * @return Metadata about the GitLab agent server for Kubernetes (KAS).
     * 
     */
    private GetMetadataKas kas;
    /**
     * @return Revision of the GitLab instance.
     * 
     */
    private String revision;
    /**
     * @return Version of the GitLab instance.
     * 
     */
    private String version;

    private GetMetadataResult() {}
    /**
     * @return If the GitLab instance is an enterprise instance or not. Supported for GitLab 15.6 onwards.
     * 
     */
    public Boolean enterprise() {
        return this.enterprise;
    }
    /**
     * @return The id of the data source. It will always be `1`
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Metadata about the GitLab agent server for Kubernetes (KAS).
     * 
     */
    public GetMetadataKas kas() {
        return this.kas;
    }
    /**
     * @return Revision of the GitLab instance.
     * 
     */
    public String revision() {
        return this.revision;
    }
    /**
     * @return Version of the GitLab instance.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMetadataResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enterprise;
        private String id;
        private GetMetadataKas kas;
        private String revision;
        private String version;
        public Builder() {}
        public Builder(GetMetadataResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enterprise = defaults.enterprise;
    	      this.id = defaults.id;
    	      this.kas = defaults.kas;
    	      this.revision = defaults.revision;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder enterprise(Boolean enterprise) {
            if (enterprise == null) {
              throw new MissingRequiredPropertyException("GetMetadataResult", "enterprise");
            }
            this.enterprise = enterprise;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMetadataResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder kas(GetMetadataKas kas) {
            if (kas == null) {
              throw new MissingRequiredPropertyException("GetMetadataResult", "kas");
            }
            this.kas = kas;
            return this;
        }
        @CustomType.Setter
        public Builder revision(String revision) {
            if (revision == null) {
              throw new MissingRequiredPropertyException("GetMetadataResult", "revision");
            }
            this.revision = revision;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("GetMetadataResult", "version");
            }
            this.version = version;
            return this;
        }
        public GetMetadataResult build() {
            final var _resultValue = new GetMetadataResult();
            _resultValue.enterprise = enterprise;
            _resultValue.id = id;
            _resultValue.kas = kas;
            _resultValue.revision = revision;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
