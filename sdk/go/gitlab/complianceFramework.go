// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `ComplianceFramework` resource allows to manage the lifecycle of a compliance framework on top-level groups.
//
// There can be only one `default` compliance framework. Of all the configured compliance frameworks marked as default, the last one applied will be the default compliance framework.
//
// > This resource requires a GitLab Enterprise instance with a Premium license to create the compliance framework.
//
// > This resource requires a GitLab Enterprise instance with an Ultimate license to specify a compliance pipeline configuration in the compliance framework.
//
// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationcreatecomplianceframework)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewComplianceFramework(ctx, "sample", &gitlab.ComplianceFrameworkArgs{
//				Color:                         pulumi.String("#87BEEF"),
//				Default:                       pulumi.Bool(false),
//				Description:                   pulumi.String("A HIPAA Compliance Framework"),
//				NamespacePath:                 pulumi.String("top-level-group"),
//				PipelineConfigurationFullPath: pulumi.String(".hipaa.yml@top-level-group/compliance-frameworks"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Gitlab compliance frameworks can be imported with a key composed of `<namespace_path>:<framework_id>`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/complianceFramework:ComplianceFramework sample "top-level-group:gid://gitlab/ComplianceManagement::Framework/12345"
//
// ```
type ComplianceFramework struct {
	pulumi.CustomResourceState

	// New color representation of the compliance framework in hex format. e.g. #FCA121.
	Color pulumi.StringOutput `pulumi:"color"`
	// Set this compliance framework as the default framework for the group. Default: `false`
	Default pulumi.BoolOutput `pulumi:"default"`
	// Description for the compliance framework.
	Description pulumi.StringOutput `pulumi:"description"`
	// Globally unique ID of the compliance framework.
	FrameworkId pulumi.StringOutput `pulumi:"frameworkId"`
	// Name for the compliance framework.
	Name pulumi.StringOutput `pulumi:"name"`
	// Full path of the namespace to add the compliance framework to.
	NamespacePath pulumi.StringOutput `pulumi:"namespacePath"`
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	PipelineConfigurationFullPath pulumi.StringPtrOutput `pulumi:"pipelineConfigurationFullPath"`
}

// NewComplianceFramework registers a new resource with the given unique name, arguments, and options.
func NewComplianceFramework(ctx *pulumi.Context,
	name string, args *ComplianceFrameworkArgs, opts ...pulumi.ResourceOption) (*ComplianceFramework, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Color == nil {
		return nil, errors.New("invalid value for required argument 'Color'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.NamespacePath == nil {
		return nil, errors.New("invalid value for required argument 'NamespacePath'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ComplianceFramework
	err := ctx.RegisterResource("gitlab:index/complianceFramework:ComplianceFramework", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComplianceFramework gets an existing ComplianceFramework resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComplianceFramework(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComplianceFrameworkState, opts ...pulumi.ResourceOption) (*ComplianceFramework, error) {
	var resource ComplianceFramework
	err := ctx.ReadResource("gitlab:index/complianceFramework:ComplianceFramework", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComplianceFramework resources.
type complianceFrameworkState struct {
	// New color representation of the compliance framework in hex format. e.g. #FCA121.
	Color *string `pulumi:"color"`
	// Set this compliance framework as the default framework for the group. Default: `false`
	Default *bool `pulumi:"default"`
	// Description for the compliance framework.
	Description *string `pulumi:"description"`
	// Globally unique ID of the compliance framework.
	FrameworkId *string `pulumi:"frameworkId"`
	// Name for the compliance framework.
	Name *string `pulumi:"name"`
	// Full path of the namespace to add the compliance framework to.
	NamespacePath *string `pulumi:"namespacePath"`
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	PipelineConfigurationFullPath *string `pulumi:"pipelineConfigurationFullPath"`
}

type ComplianceFrameworkState struct {
	// New color representation of the compliance framework in hex format. e.g. #FCA121.
	Color pulumi.StringPtrInput
	// Set this compliance framework as the default framework for the group. Default: `false`
	Default pulumi.BoolPtrInput
	// Description for the compliance framework.
	Description pulumi.StringPtrInput
	// Globally unique ID of the compliance framework.
	FrameworkId pulumi.StringPtrInput
	// Name for the compliance framework.
	Name pulumi.StringPtrInput
	// Full path of the namespace to add the compliance framework to.
	NamespacePath pulumi.StringPtrInput
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	PipelineConfigurationFullPath pulumi.StringPtrInput
}

func (ComplianceFrameworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*complianceFrameworkState)(nil)).Elem()
}

type complianceFrameworkArgs struct {
	// New color representation of the compliance framework in hex format. e.g. #FCA121.
	Color string `pulumi:"color"`
	// Set this compliance framework as the default framework for the group. Default: `false`
	Default *bool `pulumi:"default"`
	// Description for the compliance framework.
	Description string `pulumi:"description"`
	// Name for the compliance framework.
	Name *string `pulumi:"name"`
	// Full path of the namespace to add the compliance framework to.
	NamespacePath string `pulumi:"namespacePath"`
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	PipelineConfigurationFullPath *string `pulumi:"pipelineConfigurationFullPath"`
}

// The set of arguments for constructing a ComplianceFramework resource.
type ComplianceFrameworkArgs struct {
	// New color representation of the compliance framework in hex format. e.g. #FCA121.
	Color pulumi.StringInput
	// Set this compliance framework as the default framework for the group. Default: `false`
	Default pulumi.BoolPtrInput
	// Description for the compliance framework.
	Description pulumi.StringInput
	// Name for the compliance framework.
	Name pulumi.StringPtrInput
	// Full path of the namespace to add the compliance framework to.
	NamespacePath pulumi.StringInput
	// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
	PipelineConfigurationFullPath pulumi.StringPtrInput
}

func (ComplianceFrameworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*complianceFrameworkArgs)(nil)).Elem()
}

type ComplianceFrameworkInput interface {
	pulumi.Input

	ToComplianceFrameworkOutput() ComplianceFrameworkOutput
	ToComplianceFrameworkOutputWithContext(ctx context.Context) ComplianceFrameworkOutput
}

func (*ComplianceFramework) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceFramework)(nil)).Elem()
}

func (i *ComplianceFramework) ToComplianceFrameworkOutput() ComplianceFrameworkOutput {
	return i.ToComplianceFrameworkOutputWithContext(context.Background())
}

func (i *ComplianceFramework) ToComplianceFrameworkOutputWithContext(ctx context.Context) ComplianceFrameworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceFrameworkOutput)
}

func (i *ComplianceFramework) ToOutput(ctx context.Context) pulumix.Output[*ComplianceFramework] {
	return pulumix.Output[*ComplianceFramework]{
		OutputState: i.ToComplianceFrameworkOutputWithContext(ctx).OutputState,
	}
}

// ComplianceFrameworkArrayInput is an input type that accepts ComplianceFrameworkArray and ComplianceFrameworkArrayOutput values.
// You can construct a concrete instance of `ComplianceFrameworkArrayInput` via:
//
//	ComplianceFrameworkArray{ ComplianceFrameworkArgs{...} }
type ComplianceFrameworkArrayInput interface {
	pulumi.Input

	ToComplianceFrameworkArrayOutput() ComplianceFrameworkArrayOutput
	ToComplianceFrameworkArrayOutputWithContext(context.Context) ComplianceFrameworkArrayOutput
}

type ComplianceFrameworkArray []ComplianceFrameworkInput

func (ComplianceFrameworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComplianceFramework)(nil)).Elem()
}

func (i ComplianceFrameworkArray) ToComplianceFrameworkArrayOutput() ComplianceFrameworkArrayOutput {
	return i.ToComplianceFrameworkArrayOutputWithContext(context.Background())
}

func (i ComplianceFrameworkArray) ToComplianceFrameworkArrayOutputWithContext(ctx context.Context) ComplianceFrameworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceFrameworkArrayOutput)
}

func (i ComplianceFrameworkArray) ToOutput(ctx context.Context) pulumix.Output[[]*ComplianceFramework] {
	return pulumix.Output[[]*ComplianceFramework]{
		OutputState: i.ToComplianceFrameworkArrayOutputWithContext(ctx).OutputState,
	}
}

// ComplianceFrameworkMapInput is an input type that accepts ComplianceFrameworkMap and ComplianceFrameworkMapOutput values.
// You can construct a concrete instance of `ComplianceFrameworkMapInput` via:
//
//	ComplianceFrameworkMap{ "key": ComplianceFrameworkArgs{...} }
type ComplianceFrameworkMapInput interface {
	pulumi.Input

	ToComplianceFrameworkMapOutput() ComplianceFrameworkMapOutput
	ToComplianceFrameworkMapOutputWithContext(context.Context) ComplianceFrameworkMapOutput
}

type ComplianceFrameworkMap map[string]ComplianceFrameworkInput

func (ComplianceFrameworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComplianceFramework)(nil)).Elem()
}

func (i ComplianceFrameworkMap) ToComplianceFrameworkMapOutput() ComplianceFrameworkMapOutput {
	return i.ToComplianceFrameworkMapOutputWithContext(context.Background())
}

func (i ComplianceFrameworkMap) ToComplianceFrameworkMapOutputWithContext(ctx context.Context) ComplianceFrameworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComplianceFrameworkMapOutput)
}

func (i ComplianceFrameworkMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ComplianceFramework] {
	return pulumix.Output[map[string]*ComplianceFramework]{
		OutputState: i.ToComplianceFrameworkMapOutputWithContext(ctx).OutputState,
	}
}

type ComplianceFrameworkOutput struct{ *pulumi.OutputState }

func (ComplianceFrameworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceFramework)(nil)).Elem()
}

func (o ComplianceFrameworkOutput) ToComplianceFrameworkOutput() ComplianceFrameworkOutput {
	return o
}

func (o ComplianceFrameworkOutput) ToComplianceFrameworkOutputWithContext(ctx context.Context) ComplianceFrameworkOutput {
	return o
}

func (o ComplianceFrameworkOutput) ToOutput(ctx context.Context) pulumix.Output[*ComplianceFramework] {
	return pulumix.Output[*ComplianceFramework]{
		OutputState: o.OutputState,
	}
}

// New color representation of the compliance framework in hex format. e.g. #FCA121.
func (o ComplianceFrameworkOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.StringOutput { return v.Color }).(pulumi.StringOutput)
}

// Set this compliance framework as the default framework for the group. Default: `false`
func (o ComplianceFrameworkOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

// Description for the compliance framework.
func (o ComplianceFrameworkOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Globally unique ID of the compliance framework.
func (o ComplianceFrameworkOutput) FrameworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.StringOutput { return v.FrameworkId }).(pulumi.StringOutput)
}

// Name for the compliance framework.
func (o ComplianceFrameworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Full path of the namespace to add the compliance framework to.
func (o ComplianceFrameworkOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.StringOutput { return v.NamespacePath }).(pulumi.StringOutput)
}

// Full path of the compliance pipeline configuration stored in a project repository, such as `.gitlab/.compliance-gitlab-ci.yml@compliance/hipaa`. Required format: `path/file.y[a]ml@group-name/project-name` **Note**: Ultimate license required.
func (o ComplianceFrameworkOutput) PipelineConfigurationFullPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComplianceFramework) pulumi.StringPtrOutput { return v.PipelineConfigurationFullPath }).(pulumi.StringPtrOutput)
}

type ComplianceFrameworkArrayOutput struct{ *pulumi.OutputState }

func (ComplianceFrameworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComplianceFramework)(nil)).Elem()
}

func (o ComplianceFrameworkArrayOutput) ToComplianceFrameworkArrayOutput() ComplianceFrameworkArrayOutput {
	return o
}

func (o ComplianceFrameworkArrayOutput) ToComplianceFrameworkArrayOutputWithContext(ctx context.Context) ComplianceFrameworkArrayOutput {
	return o
}

func (o ComplianceFrameworkArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ComplianceFramework] {
	return pulumix.Output[[]*ComplianceFramework]{
		OutputState: o.OutputState,
	}
}

func (o ComplianceFrameworkArrayOutput) Index(i pulumi.IntInput) ComplianceFrameworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ComplianceFramework {
		return vs[0].([]*ComplianceFramework)[vs[1].(int)]
	}).(ComplianceFrameworkOutput)
}

type ComplianceFrameworkMapOutput struct{ *pulumi.OutputState }

func (ComplianceFrameworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComplianceFramework)(nil)).Elem()
}

func (o ComplianceFrameworkMapOutput) ToComplianceFrameworkMapOutput() ComplianceFrameworkMapOutput {
	return o
}

func (o ComplianceFrameworkMapOutput) ToComplianceFrameworkMapOutputWithContext(ctx context.Context) ComplianceFrameworkMapOutput {
	return o
}

func (o ComplianceFrameworkMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ComplianceFramework] {
	return pulumix.Output[map[string]*ComplianceFramework]{
		OutputState: o.OutputState,
	}
}

func (o ComplianceFrameworkMapOutput) MapIndex(k pulumi.StringInput) ComplianceFrameworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ComplianceFramework {
		return vs[0].(map[string]*ComplianceFramework)[vs[1].(string)]
	}).(ComplianceFrameworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComplianceFrameworkInput)(nil)).Elem(), &ComplianceFramework{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComplianceFrameworkArrayInput)(nil)).Elem(), ComplianceFrameworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComplianceFrameworkMapInput)(nil)).Elem(), ComplianceFrameworkMap{})
	pulumi.RegisterOutputType(ComplianceFrameworkOutput{})
	pulumi.RegisterOutputType(ComplianceFrameworkArrayOutput{})
	pulumi.RegisterOutputType(ComplianceFrameworkMapOutput{})
}