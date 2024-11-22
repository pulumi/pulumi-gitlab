// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectComplianceFrameworks` resource allows to manage the lifecycle of compliance frameworks on a project.
//
// > This resource requires a GitLab Enterprise instance with a Premium license to set the compliance frameworks on a project.
//
// **Upstream API**: [GitLab GraphQL API docs](https://docs.gitlab.com/ee/api/graphql/reference/#mutationprojectupdatecomplianceframeworks)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			alpha, err := gitlab.NewComplianceFramework(ctx, "alpha", &gitlab.ComplianceFrameworkArgs{
//				NamespacePath: pulumi.String("top-level-group"),
//				Name:          pulumi.String("HIPAA"),
//				Description:   pulumi.String("A HIPAA Compliance Framework"),
//				Color:         pulumi.String("#87BEEF"),
//				Default:       pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			beta, err := gitlab.NewComplianceFramework(ctx, "beta", &gitlab.ComplianceFrameworkArgs{
//				NamespacePath: pulumi.String("top-level-group"),
//				Name:          pulumi.String("SOC"),
//				Description:   pulumi.String("A SOC Compliance Framework"),
//				Color:         pulumi.String("#223344"),
//				Default:       pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectComplianceFrameworks(ctx, "sample", &gitlab.ProjectComplianceFrameworksArgs{
//				ComplianceFrameworkIds: pulumi.StringArray{
//					alpha.FrameworkId,
//					beta.FrameworkId,
//				},
//				Project: pulumi.String("12345678"),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_compliance_frameworks`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_compliance_frameworks.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// Gitlab project compliance frameworks can be imported with a key composed of `<project_id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks sample "42"
// ```
type ProjectComplianceFrameworks struct {
	pulumi.CustomResourceState

	// Globally unique IDs of the compliance frameworks to assign to the project.
	ComplianceFrameworkIds pulumi.StringArrayOutput `pulumi:"complianceFrameworkIds"`
	// The ID or full path of the project to change the compliance frameworks of.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectComplianceFrameworks registers a new resource with the given unique name, arguments, and options.
func NewProjectComplianceFrameworks(ctx *pulumi.Context,
	name string, args *ProjectComplianceFrameworksArgs, opts ...pulumi.ResourceOption) (*ProjectComplianceFrameworks, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComplianceFrameworkIds == nil {
		return nil, errors.New("invalid value for required argument 'ComplianceFrameworkIds'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectComplianceFrameworks
	err := ctx.RegisterResource("gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectComplianceFrameworks gets an existing ProjectComplianceFrameworks resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectComplianceFrameworks(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectComplianceFrameworksState, opts ...pulumi.ResourceOption) (*ProjectComplianceFrameworks, error) {
	var resource ProjectComplianceFrameworks
	err := ctx.ReadResource("gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectComplianceFrameworks resources.
type projectComplianceFrameworksState struct {
	// Globally unique IDs of the compliance frameworks to assign to the project.
	ComplianceFrameworkIds []string `pulumi:"complianceFrameworkIds"`
	// The ID or full path of the project to change the compliance frameworks of.
	Project *string `pulumi:"project"`
}

type ProjectComplianceFrameworksState struct {
	// Globally unique IDs of the compliance frameworks to assign to the project.
	ComplianceFrameworkIds pulumi.StringArrayInput
	// The ID or full path of the project to change the compliance frameworks of.
	Project pulumi.StringPtrInput
}

func (ProjectComplianceFrameworksState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectComplianceFrameworksState)(nil)).Elem()
}

type projectComplianceFrameworksArgs struct {
	// Globally unique IDs of the compliance frameworks to assign to the project.
	ComplianceFrameworkIds []string `pulumi:"complianceFrameworkIds"`
	// The ID or full path of the project to change the compliance frameworks of.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectComplianceFrameworks resource.
type ProjectComplianceFrameworksArgs struct {
	// Globally unique IDs of the compliance frameworks to assign to the project.
	ComplianceFrameworkIds pulumi.StringArrayInput
	// The ID or full path of the project to change the compliance frameworks of.
	Project pulumi.StringInput
}

func (ProjectComplianceFrameworksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectComplianceFrameworksArgs)(nil)).Elem()
}

type ProjectComplianceFrameworksInput interface {
	pulumi.Input

	ToProjectComplianceFrameworksOutput() ProjectComplianceFrameworksOutput
	ToProjectComplianceFrameworksOutputWithContext(ctx context.Context) ProjectComplianceFrameworksOutput
}

func (*ProjectComplianceFrameworks) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectComplianceFrameworks)(nil)).Elem()
}

func (i *ProjectComplianceFrameworks) ToProjectComplianceFrameworksOutput() ProjectComplianceFrameworksOutput {
	return i.ToProjectComplianceFrameworksOutputWithContext(context.Background())
}

func (i *ProjectComplianceFrameworks) ToProjectComplianceFrameworksOutputWithContext(ctx context.Context) ProjectComplianceFrameworksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectComplianceFrameworksOutput)
}

// ProjectComplianceFrameworksArrayInput is an input type that accepts ProjectComplianceFrameworksArray and ProjectComplianceFrameworksArrayOutput values.
// You can construct a concrete instance of `ProjectComplianceFrameworksArrayInput` via:
//
//	ProjectComplianceFrameworksArray{ ProjectComplianceFrameworksArgs{...} }
type ProjectComplianceFrameworksArrayInput interface {
	pulumi.Input

	ToProjectComplianceFrameworksArrayOutput() ProjectComplianceFrameworksArrayOutput
	ToProjectComplianceFrameworksArrayOutputWithContext(context.Context) ProjectComplianceFrameworksArrayOutput
}

type ProjectComplianceFrameworksArray []ProjectComplianceFrameworksInput

func (ProjectComplianceFrameworksArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectComplianceFrameworks)(nil)).Elem()
}

func (i ProjectComplianceFrameworksArray) ToProjectComplianceFrameworksArrayOutput() ProjectComplianceFrameworksArrayOutput {
	return i.ToProjectComplianceFrameworksArrayOutputWithContext(context.Background())
}

func (i ProjectComplianceFrameworksArray) ToProjectComplianceFrameworksArrayOutputWithContext(ctx context.Context) ProjectComplianceFrameworksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectComplianceFrameworksArrayOutput)
}

// ProjectComplianceFrameworksMapInput is an input type that accepts ProjectComplianceFrameworksMap and ProjectComplianceFrameworksMapOutput values.
// You can construct a concrete instance of `ProjectComplianceFrameworksMapInput` via:
//
//	ProjectComplianceFrameworksMap{ "key": ProjectComplianceFrameworksArgs{...} }
type ProjectComplianceFrameworksMapInput interface {
	pulumi.Input

	ToProjectComplianceFrameworksMapOutput() ProjectComplianceFrameworksMapOutput
	ToProjectComplianceFrameworksMapOutputWithContext(context.Context) ProjectComplianceFrameworksMapOutput
}

type ProjectComplianceFrameworksMap map[string]ProjectComplianceFrameworksInput

func (ProjectComplianceFrameworksMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectComplianceFrameworks)(nil)).Elem()
}

func (i ProjectComplianceFrameworksMap) ToProjectComplianceFrameworksMapOutput() ProjectComplianceFrameworksMapOutput {
	return i.ToProjectComplianceFrameworksMapOutputWithContext(context.Background())
}

func (i ProjectComplianceFrameworksMap) ToProjectComplianceFrameworksMapOutputWithContext(ctx context.Context) ProjectComplianceFrameworksMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectComplianceFrameworksMapOutput)
}

type ProjectComplianceFrameworksOutput struct{ *pulumi.OutputState }

func (ProjectComplianceFrameworksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectComplianceFrameworks)(nil)).Elem()
}

func (o ProjectComplianceFrameworksOutput) ToProjectComplianceFrameworksOutput() ProjectComplianceFrameworksOutput {
	return o
}

func (o ProjectComplianceFrameworksOutput) ToProjectComplianceFrameworksOutputWithContext(ctx context.Context) ProjectComplianceFrameworksOutput {
	return o
}

// Globally unique IDs of the compliance frameworks to assign to the project.
func (o ProjectComplianceFrameworksOutput) ComplianceFrameworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectComplianceFrameworks) pulumi.StringArrayOutput { return v.ComplianceFrameworkIds }).(pulumi.StringArrayOutput)
}

// The ID or full path of the project to change the compliance frameworks of.
func (o ProjectComplianceFrameworksOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectComplianceFrameworks) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ProjectComplianceFrameworksArrayOutput struct{ *pulumi.OutputState }

func (ProjectComplianceFrameworksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectComplianceFrameworks)(nil)).Elem()
}

func (o ProjectComplianceFrameworksArrayOutput) ToProjectComplianceFrameworksArrayOutput() ProjectComplianceFrameworksArrayOutput {
	return o
}

func (o ProjectComplianceFrameworksArrayOutput) ToProjectComplianceFrameworksArrayOutputWithContext(ctx context.Context) ProjectComplianceFrameworksArrayOutput {
	return o
}

func (o ProjectComplianceFrameworksArrayOutput) Index(i pulumi.IntInput) ProjectComplianceFrameworksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectComplianceFrameworks {
		return vs[0].([]*ProjectComplianceFrameworks)[vs[1].(int)]
	}).(ProjectComplianceFrameworksOutput)
}

type ProjectComplianceFrameworksMapOutput struct{ *pulumi.OutputState }

func (ProjectComplianceFrameworksMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectComplianceFrameworks)(nil)).Elem()
}

func (o ProjectComplianceFrameworksMapOutput) ToProjectComplianceFrameworksMapOutput() ProjectComplianceFrameworksMapOutput {
	return o
}

func (o ProjectComplianceFrameworksMapOutput) ToProjectComplianceFrameworksMapOutputWithContext(ctx context.Context) ProjectComplianceFrameworksMapOutput {
	return o
}

func (o ProjectComplianceFrameworksMapOutput) MapIndex(k pulumi.StringInput) ProjectComplianceFrameworksOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectComplianceFrameworks {
		return vs[0].(map[string]*ProjectComplianceFrameworks)[vs[1].(string)]
	}).(ProjectComplianceFrameworksOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectComplianceFrameworksInput)(nil)).Elem(), &ProjectComplianceFrameworks{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectComplianceFrameworksArrayInput)(nil)).Elem(), ProjectComplianceFrameworksArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectComplianceFrameworksMapInput)(nil)).Elem(), ProjectComplianceFrameworksMap{})
	pulumi.RegisterOutputType(ProjectComplianceFrameworksOutput{})
	pulumi.RegisterOutputType(ProjectComplianceFrameworksArrayOutput{})
	pulumi.RegisterOutputType(ProjectComplianceFrameworksMapOutput{})
}
