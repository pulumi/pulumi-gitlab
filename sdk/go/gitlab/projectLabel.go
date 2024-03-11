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

// The `ProjectLabel` resource allows to manage the lifecycle of a project label.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/labels.html#project-labels)
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
//			_, err := gitlab.NewProjectLabel(ctx, "fixme", &gitlab.ProjectLabelArgs{
//				Project:     pulumi.String("example"),
//				Description: pulumi.String("issue with failing tests"),
//				Color:       pulumi.String("#ffcc00"),
//			})
//			if err != nil {
//				return err
//			}
//			// Scoped label
//			_, err = gitlab.NewProjectLabel(ctx, "devopsCreate", &gitlab.ProjectLabelArgs{
//				Project:     pulumi.Any(gitlab_project.Example.Id),
//				Description: pulumi.String("issue for creating infrastructure resources"),
//				Color:       pulumi.String("#ffa500"),
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
// Gitlab Project labels can be imported using an id made up of `{project_id}:{group_label_id}`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectLabel:ProjectLabel example 12345:fixme
// ```
type ProjectLabel struct {
	pulumi.CustomResourceState

	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringOutput `pulumi:"color"`
	// The description of the label.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The id of the project label.
	LabelId pulumi.IntOutput `pulumi:"labelId"`
	// The name of the label.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectLabel registers a new resource with the given unique name, arguments, and options.
func NewProjectLabel(ctx *pulumi.Context,
	name string, args *ProjectLabelArgs, opts ...pulumi.ResourceOption) (*ProjectLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Color == nil {
		return nil, errors.New("invalid value for required argument 'Color'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectLabel
	err := ctx.RegisterResource("gitlab:index/projectLabel:ProjectLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectLabel gets an existing ProjectLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectLabelState, opts ...pulumi.ResourceOption) (*ProjectLabel, error) {
	var resource ProjectLabel
	err := ctx.ReadResource("gitlab:index/projectLabel:ProjectLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectLabel resources.
type projectLabelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color *string `pulumi:"color"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The id of the project label.
	LabelId *int `pulumi:"labelId"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project *string `pulumi:"project"`
}

type ProjectLabelState struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringPtrInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The id of the project label.
	LabelId pulumi.IntPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the label to.
	Project pulumi.StringPtrInput
}

func (ProjectLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLabelState)(nil)).Elem()
}

type projectLabelArgs struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color string `pulumi:"color"`
	// The description of the label.
	Description *string `pulumi:"description"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the label to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectLabel resource.
type ProjectLabelArgs struct {
	// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
	Color pulumi.StringInput
	// The description of the label.
	Description pulumi.StringPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the label to.
	Project pulumi.StringInput
}

func (ProjectLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLabelArgs)(nil)).Elem()
}

type ProjectLabelInput interface {
	pulumi.Input

	ToProjectLabelOutput() ProjectLabelOutput
	ToProjectLabelOutputWithContext(ctx context.Context) ProjectLabelOutput
}

func (*ProjectLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLabel)(nil)).Elem()
}

func (i *ProjectLabel) ToProjectLabelOutput() ProjectLabelOutput {
	return i.ToProjectLabelOutputWithContext(context.Background())
}

func (i *ProjectLabel) ToProjectLabelOutputWithContext(ctx context.Context) ProjectLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLabelOutput)
}

// ProjectLabelArrayInput is an input type that accepts ProjectLabelArray and ProjectLabelArrayOutput values.
// You can construct a concrete instance of `ProjectLabelArrayInput` via:
//
//	ProjectLabelArray{ ProjectLabelArgs{...} }
type ProjectLabelArrayInput interface {
	pulumi.Input

	ToProjectLabelArrayOutput() ProjectLabelArrayOutput
	ToProjectLabelArrayOutputWithContext(context.Context) ProjectLabelArrayOutput
}

type ProjectLabelArray []ProjectLabelInput

func (ProjectLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectLabel)(nil)).Elem()
}

func (i ProjectLabelArray) ToProjectLabelArrayOutput() ProjectLabelArrayOutput {
	return i.ToProjectLabelArrayOutputWithContext(context.Background())
}

func (i ProjectLabelArray) ToProjectLabelArrayOutputWithContext(ctx context.Context) ProjectLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLabelArrayOutput)
}

// ProjectLabelMapInput is an input type that accepts ProjectLabelMap and ProjectLabelMapOutput values.
// You can construct a concrete instance of `ProjectLabelMapInput` via:
//
//	ProjectLabelMap{ "key": ProjectLabelArgs{...} }
type ProjectLabelMapInput interface {
	pulumi.Input

	ToProjectLabelMapOutput() ProjectLabelMapOutput
	ToProjectLabelMapOutputWithContext(context.Context) ProjectLabelMapOutput
}

type ProjectLabelMap map[string]ProjectLabelInput

func (ProjectLabelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectLabel)(nil)).Elem()
}

func (i ProjectLabelMap) ToProjectLabelMapOutput() ProjectLabelMapOutput {
	return i.ToProjectLabelMapOutputWithContext(context.Background())
}

func (i ProjectLabelMap) ToProjectLabelMapOutputWithContext(ctx context.Context) ProjectLabelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectLabelMapOutput)
}

type ProjectLabelOutput struct{ *pulumi.OutputState }

func (ProjectLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectLabel)(nil)).Elem()
}

func (o ProjectLabelOutput) ToProjectLabelOutput() ProjectLabelOutput {
	return o
}

func (o ProjectLabelOutput) ToProjectLabelOutputWithContext(ctx context.Context) ProjectLabelOutput {
	return o
}

// The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).
func (o ProjectLabelOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectLabel) pulumi.StringOutput { return v.Color }).(pulumi.StringOutput)
}

// The description of the label.
func (o ProjectLabelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectLabel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The id of the project label.
func (o ProjectLabelOutput) LabelId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectLabel) pulumi.IntOutput { return v.LabelId }).(pulumi.IntOutput)
}

// The name of the label.
func (o ProjectLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectLabel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name or id of the project to add the label to.
func (o ProjectLabelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectLabel) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ProjectLabelArrayOutput struct{ *pulumi.OutputState }

func (ProjectLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectLabel)(nil)).Elem()
}

func (o ProjectLabelArrayOutput) ToProjectLabelArrayOutput() ProjectLabelArrayOutput {
	return o
}

func (o ProjectLabelArrayOutput) ToProjectLabelArrayOutputWithContext(ctx context.Context) ProjectLabelArrayOutput {
	return o
}

func (o ProjectLabelArrayOutput) Index(i pulumi.IntInput) ProjectLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectLabel {
		return vs[0].([]*ProjectLabel)[vs[1].(int)]
	}).(ProjectLabelOutput)
}

type ProjectLabelMapOutput struct{ *pulumi.OutputState }

func (ProjectLabelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectLabel)(nil)).Elem()
}

func (o ProjectLabelMapOutput) ToProjectLabelMapOutput() ProjectLabelMapOutput {
	return o
}

func (o ProjectLabelMapOutput) ToProjectLabelMapOutputWithContext(ctx context.Context) ProjectLabelMapOutput {
	return o
}

func (o ProjectLabelMapOutput) MapIndex(k pulumi.StringInput) ProjectLabelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectLabel {
		return vs[0].(map[string]*ProjectLabel)[vs[1].(string)]
	}).(ProjectLabelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLabelInput)(nil)).Elem(), &ProjectLabel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLabelArrayInput)(nil)).Elem(), ProjectLabelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectLabelMapInput)(nil)).Elem(), ProjectLabelMap{})
	pulumi.RegisterOutputType(ProjectLabelOutput{})
	pulumi.RegisterOutputType(ProjectLabelArrayOutput{})
	pulumi.RegisterOutputType(ProjectLabelMapOutput{})
}
