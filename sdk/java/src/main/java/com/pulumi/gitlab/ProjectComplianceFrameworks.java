// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gitlab.ProjectComplianceFrameworksArgs;
import com.pulumi.gitlab.Utilities;
import com.pulumi.gitlab.inputs.ProjectComplianceFrameworksState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The `gitlab.ProjectComplianceFrameworks` resource allows to manage the lifecycle of compliance frameworks on a project.
 * 
 * &gt; This resource requires a GitLab Enterprise instance with a Premium license to set the compliance frameworks on a project.
 * 
 * **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationprojectupdatecomplianceframeworks)
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
 * import com.pulumi.gitlab.ProjectComplianceFrameworks;
 * import com.pulumi.gitlab.ProjectComplianceFrameworksArgs;
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
 *         var alpha = new ComplianceFramework("alpha", ComplianceFrameworkArgs.builder()
 *             .namespacePath("top-level-group")
 *             .name("HIPAA")
 *             .description("A HIPAA Compliance Framework")
 *             .color("#87BEEF")
 *             .default_(false)
 *             .build());
 * 
 *         var beta = new ComplianceFramework("beta", ComplianceFrameworkArgs.builder()
 *             .namespacePath("top-level-group")
 *             .name("SOC")
 *             .description("A SOC Compliance Framework")
 *             .color("#223344")
 *             .default_(false)
 *             .build());
 * 
 *         var sample = new ProjectComplianceFrameworks("sample", ProjectComplianceFrameworksArgs.builder()
 *             .complianceFrameworkIds(            
 *                 alpha.frameworkId(),
 *                 beta.frameworkId())
 *             .project("12345678")
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
 * Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_compliance_frameworks`. For example:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = gitlab_project_compliance_frameworks.example
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
 * $ pulumi import gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks sample &#34;42&#34;
 * ```
 * 
 */
@ResourceType(type="gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks")
public class ProjectComplianceFrameworks extends com.pulumi.resources.CustomResource {
    /**
     * Globally unique IDs of the compliance frameworks to assign to the project.
     * 
     */
    @Export(name="complianceFrameworkIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> complianceFrameworkIds;

    /**
     * @return Globally unique IDs of the compliance frameworks to assign to the project.
     * 
     */
    public Output<List<String>> complianceFrameworkIds() {
        return this.complianceFrameworkIds;
    }
    /**
     * The ID or full path of the project to change the compliance frameworks of.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID or full path of the project to change the compliance frameworks of.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectComplianceFrameworks(java.lang.String name) {
        this(name, ProjectComplianceFrameworksArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectComplianceFrameworks(java.lang.String name, ProjectComplianceFrameworksArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectComplianceFrameworks(java.lang.String name, ProjectComplianceFrameworksArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProjectComplianceFrameworks(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectComplianceFrameworksState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks", name, state, makeResourceOptions(options, id), false);
    }

    private static ProjectComplianceFrameworksArgs makeArgs(ProjectComplianceFrameworksArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProjectComplianceFrameworksArgs.Empty : args;
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
    public static ProjectComplianceFrameworks get(java.lang.String name, Output<java.lang.String> id, @Nullable ProjectComplianceFrameworksState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectComplianceFrameworks(name, id, state, options);
    }
}
