// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectMirrorPublicKeyResult {
    private String id;
    /**
     * @return The id of the remote mirror.
     * 
     */
    private Integer mirrorId;
    /**
     * @return The integer or path with namespace that uniquely identifies the project.
     * 
     */
    private String projectId;
    /**
     * @return Public key of the remote mirror.
     * 
     */
    private String publicKey;

    private GetProjectMirrorPublicKeyResult() {}
    public String id() {
        return this.id;
    }
    /**
     * @return The id of the remote mirror.
     * 
     */
    public Integer mirrorId() {
        return this.mirrorId;
    }
    /**
     * @return The integer or path with namespace that uniquely identifies the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Public key of the remote mirror.
     * 
     */
    public String publicKey() {
        return this.publicKey;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectMirrorPublicKeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Integer mirrorId;
        private String projectId;
        private String publicKey;
        public Builder() {}
        public Builder(GetProjectMirrorPublicKeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.mirrorId = defaults.mirrorId;
    	      this.projectId = defaults.projectId;
    	      this.publicKey = defaults.publicKey;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectMirrorPublicKeyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mirrorId(Integer mirrorId) {
            if (mirrorId == null) {
              throw new MissingRequiredPropertyException("GetProjectMirrorPublicKeyResult", "mirrorId");
            }
            this.mirrorId = mirrorId;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetProjectMirrorPublicKeyResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder publicKey(String publicKey) {
            if (publicKey == null) {
              throw new MissingRequiredPropertyException("GetProjectMirrorPublicKeyResult", "publicKey");
            }
            this.publicKey = publicKey;
            return this;
        }
        public GetProjectMirrorPublicKeyResult build() {
            final var _resultValue = new GetProjectMirrorPublicKeyResult();
            _resultValue.id = id;
            _resultValue.mirrorId = mirrorId;
            _resultValue.projectId = projectId;
            _resultValue.publicKey = publicKey;
            return _resultValue;
        }
    }
}
