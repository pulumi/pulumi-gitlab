// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectVariable` resource allows to manage the lifecycle of a CI/CD variable for a project.
//
// > **Important:** If your GitLab version is older than 13.4, you may see nondeterministic behavior when updating or deleting ProjectVariable resources with non-unique keys, for example if there is another variable with the same key and different environment scope. See [this GitLab issue](https://gitlab.com/gitlab-org/gitlab/-/issues/9912).
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := gitlab.NewProjectVariable(ctx, "example", &gitlab.ProjectVariableArgs{
//				Key:       pulumi.String("project_variable_key"),
//				Project:   pulumi.String("12345"),
//				Protected: pulumi.Bool(false),
//				Value:     pulumi.String("project_variable_value"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// GitLab project variables can be imported using an id made up of `project:key:environment_scope`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectVariable:ProjectVariable example '12345:project_variable_key:*'
// ```
type ProjectVariable struct {
	pulumi.CustomResourceState

	// The description of the variable.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrOutput `pulumi:"environmentScope"`
	// The name of the variable.
	Key pulumi.StringOutput `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked pulumi.BoolPtrOutput `pulumi:"masked"`
	// The name or id of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrOutput `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw pulumi.BoolPtrOutput `pulumi:"raw"`
	// The value of the variable.
	Value pulumi.StringOutput `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType pulumi.StringPtrOutput `pulumi:"variableType"`
}

// NewProjectVariable registers a new resource with the given unique name, arguments, and options.
func NewProjectVariable(ctx *pulumi.Context,
	name string, args *ProjectVariableArgs, opts ...pulumi.ResourceOption) (*ProjectVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectVariable
	err := ctx.RegisterResource("gitlab:index/projectVariable:ProjectVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectVariable gets an existing ProjectVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectVariableState, opts ...pulumi.ResourceOption) (*ProjectVariable, error) {
	var resource ProjectVariable
	err := ctx.ReadResource("gitlab:index/projectVariable:ProjectVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectVariable resources.
type projectVariableState struct {
	// The description of the variable.
	Description *string `pulumi:"description"`
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name of the variable.
	Key *string `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked *bool `pulumi:"masked"`
	// The name or id of the project.
	Project *string `pulumi:"project"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected *bool `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw *bool `pulumi:"raw"`
	// The value of the variable.
	Value *string `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType *string `pulumi:"variableType"`
}

type ProjectVariableState struct {
	// The description of the variable.
	Description pulumi.StringPtrInput
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrInput
	// The name of the variable.
	Key pulumi.StringPtrInput
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked pulumi.BoolPtrInput
	// The name or id of the project.
	Project pulumi.StringPtrInput
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrInput
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw pulumi.BoolPtrInput
	// The value of the variable.
	Value pulumi.StringPtrInput
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType pulumi.StringPtrInput
}

func (ProjectVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectVariableState)(nil)).Elem()
}

type projectVariableArgs struct {
	// The description of the variable.
	Description *string `pulumi:"description"`
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The name of the variable.
	Key string `pulumi:"key"`
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked *bool `pulumi:"masked"`
	// The name or id of the project.
	Project string `pulumi:"project"`
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected *bool `pulumi:"protected"`
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw *bool `pulumi:"raw"`
	// The value of the variable.
	Value string `pulumi:"value"`
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType *string `pulumi:"variableType"`
}

// The set of arguments for constructing a ProjectVariable resource.
type ProjectVariableArgs struct {
	// The description of the variable.
	Description pulumi.StringPtrInput
	// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
	EnvironmentScope pulumi.StringPtrInput
	// The name of the variable.
	Key pulumi.StringInput
	// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
	Masked pulumi.BoolPtrInput
	// The name or id of the project.
	Project pulumi.StringInput
	// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
	Protected pulumi.BoolPtrInput
	// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
	Raw pulumi.BoolPtrInput
	// The value of the variable.
	Value pulumi.StringInput
	// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
	VariableType pulumi.StringPtrInput
}

func (ProjectVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectVariableArgs)(nil)).Elem()
}

type ProjectVariableInput interface {
	pulumi.Input

	ToProjectVariableOutput() ProjectVariableOutput
	ToProjectVariableOutputWithContext(ctx context.Context) ProjectVariableOutput
}

func (*ProjectVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectVariable)(nil)).Elem()
}

func (i *ProjectVariable) ToProjectVariableOutput() ProjectVariableOutput {
	return i.ToProjectVariableOutputWithContext(context.Background())
}

func (i *ProjectVariable) ToProjectVariableOutputWithContext(ctx context.Context) ProjectVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectVariableOutput)
}

// ProjectVariableArrayInput is an input type that accepts ProjectVariableArray and ProjectVariableArrayOutput values.
// You can construct a concrete instance of `ProjectVariableArrayInput` via:
//
//	ProjectVariableArray{ ProjectVariableArgs{...} }
type ProjectVariableArrayInput interface {
	pulumi.Input

	ToProjectVariableArrayOutput() ProjectVariableArrayOutput
	ToProjectVariableArrayOutputWithContext(context.Context) ProjectVariableArrayOutput
}

type ProjectVariableArray []ProjectVariableInput

func (ProjectVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectVariable)(nil)).Elem()
}

func (i ProjectVariableArray) ToProjectVariableArrayOutput() ProjectVariableArrayOutput {
	return i.ToProjectVariableArrayOutputWithContext(context.Background())
}

func (i ProjectVariableArray) ToProjectVariableArrayOutputWithContext(ctx context.Context) ProjectVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectVariableArrayOutput)
}

// ProjectVariableMapInput is an input type that accepts ProjectVariableMap and ProjectVariableMapOutput values.
// You can construct a concrete instance of `ProjectVariableMapInput` via:
//
//	ProjectVariableMap{ "key": ProjectVariableArgs{...} }
type ProjectVariableMapInput interface {
	pulumi.Input

	ToProjectVariableMapOutput() ProjectVariableMapOutput
	ToProjectVariableMapOutputWithContext(context.Context) ProjectVariableMapOutput
}

type ProjectVariableMap map[string]ProjectVariableInput

func (ProjectVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectVariable)(nil)).Elem()
}

func (i ProjectVariableMap) ToProjectVariableMapOutput() ProjectVariableMapOutput {
	return i.ToProjectVariableMapOutputWithContext(context.Background())
}

func (i ProjectVariableMap) ToProjectVariableMapOutputWithContext(ctx context.Context) ProjectVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectVariableMapOutput)
}

type ProjectVariableOutput struct{ *pulumi.OutputState }

func (ProjectVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectVariable)(nil)).Elem()
}

func (o ProjectVariableOutput) ToProjectVariableOutput() ProjectVariableOutput {
	return o
}

func (o ProjectVariableOutput) ToProjectVariableOutputWithContext(ctx context.Context) ProjectVariableOutput {
	return o
}

// The description of the variable.
func (o ProjectVariableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The environment scope of the variable. Defaults to all environment (`*`). Note that in Community Editions of Gitlab, values other than `*` will cause inconsistent plans.
func (o ProjectVariableOutput) EnvironmentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.StringPtrOutput { return v.EnvironmentScope }).(pulumi.StringPtrOutput)
}

// The name of the variable.
func (o ProjectVariableOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// If set to `true`, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to `false`.
func (o ProjectVariableOutput) Masked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.BoolPtrOutput { return v.Masked }).(pulumi.BoolPtrOutput)
}

// The name or id of the project.
func (o ProjectVariableOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
func (o ProjectVariableOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.BoolPtrOutput { return v.Protected }).(pulumi.BoolPtrOutput)
}

// Whether the variable is treated as a raw string. Default: false. When true, variables in the value are not expanded.
func (o ProjectVariableOutput) Raw() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.BoolPtrOutput { return v.Raw }).(pulumi.BoolPtrOutput)
}

// The value of the variable.
func (o ProjectVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// The type of a variable. Valid values are: `envVar`, `file`. Default is `envVar`.
func (o ProjectVariableOutput) VariableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectVariable) pulumi.StringPtrOutput { return v.VariableType }).(pulumi.StringPtrOutput)
}

type ProjectVariableArrayOutput struct{ *pulumi.OutputState }

func (ProjectVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectVariable)(nil)).Elem()
}

func (o ProjectVariableArrayOutput) ToProjectVariableArrayOutput() ProjectVariableArrayOutput {
	return o
}

func (o ProjectVariableArrayOutput) ToProjectVariableArrayOutputWithContext(ctx context.Context) ProjectVariableArrayOutput {
	return o
}

func (o ProjectVariableArrayOutput) Index(i pulumi.IntInput) ProjectVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectVariable {
		return vs[0].([]*ProjectVariable)[vs[1].(int)]
	}).(ProjectVariableOutput)
}

type ProjectVariableMapOutput struct{ *pulumi.OutputState }

func (ProjectVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectVariable)(nil)).Elem()
}

func (o ProjectVariableMapOutput) ToProjectVariableMapOutput() ProjectVariableMapOutput {
	return o
}

func (o ProjectVariableMapOutput) ToProjectVariableMapOutputWithContext(ctx context.Context) ProjectVariableMapOutput {
	return o
}

func (o ProjectVariableMapOutput) MapIndex(k pulumi.StringInput) ProjectVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectVariable {
		return vs[0].(map[string]*ProjectVariable)[vs[1].(string)]
	}).(ProjectVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectVariableInput)(nil)).Elem(), &ProjectVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectVariableArrayInput)(nil)).Elem(), ProjectVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectVariableMapInput)(nil)).Elem(), ProjectVariableMap{})
	pulumi.RegisterOutputType(ProjectVariableOutput{})
	pulumi.RegisterOutputType(ProjectVariableArrayOutput{})
	pulumi.RegisterOutputType(ProjectVariableMapOutput{})
}
