// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectComplianceFrameworkArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectComplianceFrameworkState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectComplianceFramework` resource allows to manage the lifecycle of a compliance framework on a project.
 * 
 * &gt; This resource requires a GitLab Enterprise instance with a Premium license to set the compliance framework on a project.
 * 
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/api/graphql/reference/#mutationprojectsetcomplianceframework)
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
 * import com.pulumi.gitlab.ComplianceFramework;
 * import com.pulumi.gitlab.ComplianceFrameworkArgs;
 * import com.pulumi.gitlab.ProjectComplianceFramework;
 * import com.pulumi.gitlab.ProjectComplianceFrameworkArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var sample = new ComplianceFramework("sample", ComplianceFrameworkArgs.builder()
 *             .namespacePath("top-level-group")
 *             .name("HIPAA")
 *             .description("A HIPAA Compliance Framework")
 *             .color("#87BEEF")
 *             .default_(false)
 *             .pipelineConfigurationFullPath(".hipaa.yml}{@literal @}{@code top-level-group/compliance-frameworks")
 *             .build());
 * 
 *         var sampleProjectComplianceFramework = new ProjectComplianceFramework("sampleProjectComplianceFramework", ProjectComplianceFrameworkArgs.builder()
 *             .complianceFrameworkId(sample.frameworkId())
 *             .project("12345678")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_compliance_framework`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_compliance_framework.example
 * 
 *   id = &#34;see CLI command below for ID&#34;
 * 
 * }
 * 
 * Import using the CLI is supported using the following syntax:
 * 
 * Gitlab project compliance frameworks can be imported with a key composed of `&lt;project_id&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import gitlab:index/projectComplianceFramework:ProjectComplianceFramework sample &#34;42&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectComplianceFramework:ProjectComplianceFramework")
public class ProjectComplianceFramework extends com.pulumi.resources.CustomResource {
    /**
     * Globally unique ID of the compliance framework to assign to the project.
     * 
     */
    @Export(name="complianceFrameworkId", refs={String.class}, tree="[0]")
    private Output<String> complianceFrameworkId;

    /**
     * @return Globally unique ID of the compliance framework to assign to the project.
     * 
     */
    public Output<String> complianceFrameworkId() {
        return this.complianceFrameworkId;
    }
    /**
     * The ID or full path of the project to change the compliance framework of.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project to change the compliance framework of.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectComplianceFramework(java.lang.String name) {
        this(name, ProjectComplianceFrameworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectComplianceFramework(java.lang.String name, ProjectComplianceFrameworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectComplianceFramework(java.lang.String name, ProjectComplianceFrameworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectComplianceFramework:ProjectComplianceFramework", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectComplianceFramework(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectComplianceFrameworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectComplianceFramework:ProjectComplianceFramework", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectComplianceFrameworkArgs makeArgs(ProjectComplianceFrameworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectComplianceFrameworkArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ProjectComplianceFramework get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectComplianceFrameworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectComplianceFramework(name, id, state, options);
    }
}
