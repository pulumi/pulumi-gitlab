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

// The `ProjectMilestone` resource allows to manage the lifecycle of a project milestone.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/milestones.html)
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
//			// Create a project for the milestone to use
//			example, err := gitlab.NewProject(ctx, "example", &gitlab.ProjectArgs{
//				Name:        pulumi.String("example"),
//				Description: pulumi.String("An example project"),
//				NamespaceId: pulumi.Any(exampleGitlabGroup.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectMilestone(ctx, "example", &gitlab.ProjectMilestoneArgs{
//				Project: example.ID(),
//				Title:   pulumi.String("example"),
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
// Gitlab project milestone can be imported with a key composed of `<project>:<milestone_id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectMilestone:ProjectMilestone example "12345:11"
// ```
type ProjectMilestone struct {
	pulumi.CustomResourceState

	// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the milestone.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	DueDate pulumi.StringPtrOutput `pulumi:"dueDate"`
	// Bool, true if milestone expired.
	Expired pulumi.BoolOutput `pulumi:"expired"`
	// The ID of the project's milestone.
	Iid pulumi.IntOutput `pulumi:"iid"`
	// The instance-wide ID of the project’s milestone.
	MilestoneId pulumi.IntOutput `pulumi:"milestoneId"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringOutput `pulumi:"project"`
	// The project ID of milestone.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// The state of the milestone. Valid values are: `active`, `closed`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The title of a milestone.
	Title pulumi.StringOutput `pulumi:"title"`
	// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The web URL of the milestone.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
}

// NewProjectMilestone registers a new resource with the given unique name, arguments, and options.
func NewProjectMilestone(ctx *pulumi.Context,
	name string, args *ProjectMilestoneArgs, opts ...pulumi.ResourceOption) (*ProjectMilestone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectMilestone
	err := ctx.RegisterResource("gitlab:index/projectMilestone:ProjectMilestone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMilestone gets an existing ProjectMilestone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMilestone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMilestoneState, opts ...pulumi.ResourceOption) (*ProjectMilestone, error) {
	var resource ProjectMilestone
	err := ctx.ReadResource("gitlab:index/projectMilestone:ProjectMilestone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMilestone resources.
type projectMilestoneState struct {
	// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the milestone.
	Description *string `pulumi:"description"`
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	DueDate *string `pulumi:"dueDate"`
	// Bool, true if milestone expired.
	Expired *bool `pulumi:"expired"`
	// The ID of the project's milestone.
	Iid *int `pulumi:"iid"`
	// The instance-wide ID of the project’s milestone.
	MilestoneId *int `pulumi:"milestoneId"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project *string `pulumi:"project"`
	// The project ID of milestone.
	ProjectId *int `pulumi:"projectId"`
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	StartDate *string `pulumi:"startDate"`
	// The state of the milestone. Valid values are: `active`, `closed`.
	State *string `pulumi:"state"`
	// The title of a milestone.
	Title *string `pulumi:"title"`
	// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The web URL of the milestone.
	WebUrl *string `pulumi:"webUrl"`
}

type ProjectMilestoneState struct {
	// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	CreatedAt pulumi.StringPtrInput
	// The description of the milestone.
	Description pulumi.StringPtrInput
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	DueDate pulumi.StringPtrInput
	// Bool, true if milestone expired.
	Expired pulumi.BoolPtrInput
	// The ID of the project's milestone.
	Iid pulumi.IntPtrInput
	// The instance-wide ID of the project’s milestone.
	MilestoneId pulumi.IntPtrInput
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringPtrInput
	// The project ID of milestone.
	ProjectId pulumi.IntPtrInput
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	StartDate pulumi.StringPtrInput
	// The state of the milestone. Valid values are: `active`, `closed`.
	State pulumi.StringPtrInput
	// The title of a milestone.
	Title pulumi.StringPtrInput
	// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
	UpdatedAt pulumi.StringPtrInput
	// The web URL of the milestone.
	WebUrl pulumi.StringPtrInput
}

func (ProjectMilestoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMilestoneState)(nil)).Elem()
}

type projectMilestoneArgs struct {
	// The description of the milestone.
	Description *string `pulumi:"description"`
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	DueDate *string `pulumi:"dueDate"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	StartDate *string `pulumi:"startDate"`
	// The state of the milestone. Valid values are: `active`, `closed`.
	State *string `pulumi:"state"`
	// The title of a milestone.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a ProjectMilestone resource.
type ProjectMilestoneArgs struct {
	// The description of the milestone.
	Description pulumi.StringPtrInput
	// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	DueDate pulumi.StringPtrInput
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringInput
	// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
	StartDate pulumi.StringPtrInput
	// The state of the milestone. Valid values are: `active`, `closed`.
	State pulumi.StringPtrInput
	// The title of a milestone.
	Title pulumi.StringInput
}

func (ProjectMilestoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMilestoneArgs)(nil)).Elem()
}

type ProjectMilestoneInput interface {
	pulumi.Input

	ToProjectMilestoneOutput() ProjectMilestoneOutput
	ToProjectMilestoneOutputWithContext(ctx context.Context) ProjectMilestoneOutput
}

func (*ProjectMilestone) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMilestone)(nil)).Elem()
}

func (i *ProjectMilestone) ToProjectMilestoneOutput() ProjectMilestoneOutput {
	return i.ToProjectMilestoneOutputWithContext(context.Background())
}

func (i *ProjectMilestone) ToProjectMilestoneOutputWithContext(ctx context.Context) ProjectMilestoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMilestoneOutput)
}

// ProjectMilestoneArrayInput is an input type that accepts ProjectMilestoneArray and ProjectMilestoneArrayOutput values.
// You can construct a concrete instance of `ProjectMilestoneArrayInput` via:
//
//	ProjectMilestoneArray{ ProjectMilestoneArgs{...} }
type ProjectMilestoneArrayInput interface {
	pulumi.Input

	ToProjectMilestoneArrayOutput() ProjectMilestoneArrayOutput
	ToProjectMilestoneArrayOutputWithContext(context.Context) ProjectMilestoneArrayOutput
}

type ProjectMilestoneArray []ProjectMilestoneInput

func (ProjectMilestoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMilestone)(nil)).Elem()
}

func (i ProjectMilestoneArray) ToProjectMilestoneArrayOutput() ProjectMilestoneArrayOutput {
	return i.ToProjectMilestoneArrayOutputWithContext(context.Background())
}

func (i ProjectMilestoneArray) ToProjectMilestoneArrayOutputWithContext(ctx context.Context) ProjectMilestoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMilestoneArrayOutput)
}

// ProjectMilestoneMapInput is an input type that accepts ProjectMilestoneMap and ProjectMilestoneMapOutput values.
// You can construct a concrete instance of `ProjectMilestoneMapInput` via:
//
//	ProjectMilestoneMap{ "key": ProjectMilestoneArgs{...} }
type ProjectMilestoneMapInput interface {
	pulumi.Input

	ToProjectMilestoneMapOutput() ProjectMilestoneMapOutput
	ToProjectMilestoneMapOutputWithContext(context.Context) ProjectMilestoneMapOutput
}

type ProjectMilestoneMap map[string]ProjectMilestoneInput

func (ProjectMilestoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMilestone)(nil)).Elem()
}

func (i ProjectMilestoneMap) ToProjectMilestoneMapOutput() ProjectMilestoneMapOutput {
	return i.ToProjectMilestoneMapOutputWithContext(context.Background())
}

func (i ProjectMilestoneMap) ToProjectMilestoneMapOutputWithContext(ctx context.Context) ProjectMilestoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMilestoneMapOutput)
}

type ProjectMilestoneOutput struct{ *pulumi.OutputState }

func (ProjectMilestoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMilestone)(nil)).Elem()
}

func (o ProjectMilestoneOutput) ToProjectMilestoneOutput() ProjectMilestoneOutput {
	return o
}

func (o ProjectMilestoneOutput) ToProjectMilestoneOutputWithContext(ctx context.Context) ProjectMilestoneOutput {
	return o
}

// The time of creation of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
func (o ProjectMilestoneOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the milestone.
func (o ProjectMilestoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The due date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
func (o ProjectMilestoneOutput) DueDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringPtrOutput { return v.DueDate }).(pulumi.StringPtrOutput)
}

// Bool, true if milestone expired.
func (o ProjectMilestoneOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.BoolOutput { return v.Expired }).(pulumi.BoolOutput)
}

// The ID of the project's milestone.
func (o ProjectMilestoneOutput) Iid() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.IntOutput { return v.Iid }).(pulumi.IntOutput)
}

// The instance-wide ID of the project’s milestone.
func (o ProjectMilestoneOutput) MilestoneId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.IntOutput { return v.MilestoneId }).(pulumi.IntOutput)
}

// The ID or URL-encoded path of the project owned by the authenticated user.
func (o ProjectMilestoneOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The project ID of milestone.
func (o ProjectMilestoneOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// The start date of the milestone. Date time string in the format YYYY-MM-DD, for example 2016-03-11.
func (o ProjectMilestoneOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

// The state of the milestone. Valid values are: `active`, `closed`.
func (o ProjectMilestoneOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The title of a milestone.
func (o ProjectMilestoneOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The last update time of the milestone. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z.
func (o ProjectMilestoneOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The web URL of the milestone.
func (o ProjectMilestoneOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMilestone) pulumi.StringOutput { return v.WebUrl }).(pulumi.StringOutput)
}

type ProjectMilestoneArrayOutput struct{ *pulumi.OutputState }

func (ProjectMilestoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMilestone)(nil)).Elem()
}

func (o ProjectMilestoneArrayOutput) ToProjectMilestoneArrayOutput() ProjectMilestoneArrayOutput {
	return o
}

func (o ProjectMilestoneArrayOutput) ToProjectMilestoneArrayOutputWithContext(ctx context.Context) ProjectMilestoneArrayOutput {
	return o
}

func (o ProjectMilestoneArrayOutput) Index(i pulumi.IntInput) ProjectMilestoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectMilestone {
		return vs[0].([]*ProjectMilestone)[vs[1].(int)]
	}).(ProjectMilestoneOutput)
}

type ProjectMilestoneMapOutput struct{ *pulumi.OutputState }

func (ProjectMilestoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMilestone)(nil)).Elem()
}

func (o ProjectMilestoneMapOutput) ToProjectMilestoneMapOutput() ProjectMilestoneMapOutput {
	return o
}

func (o ProjectMilestoneMapOutput) ToProjectMilestoneMapOutputWithContext(ctx context.Context) ProjectMilestoneMapOutput {
	return o
}

func (o ProjectMilestoneMapOutput) MapIndex(k pulumi.StringInput) ProjectMilestoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectMilestone {
		return vs[0].(map[string]*ProjectMilestone)[vs[1].(string)]
	}).(ProjectMilestoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMilestoneInput)(nil)).Elem(), &ProjectMilestone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMilestoneArrayInput)(nil)).Elem(), ProjectMilestoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMilestoneMapInput)(nil)).Elem(), ProjectMilestoneMap{})
	pulumi.RegisterOutputType(ProjectMilestoneOutput{})
	pulumi.RegisterOutputType(ProjectMilestoneArrayOutput{})
	pulumi.RegisterOutputType(ProjectMilestoneMapOutput{})
}
