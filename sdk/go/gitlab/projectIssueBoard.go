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

// The `ProjectIssueBoard` resource allows to manage the lifecycle of a Project Issue Board.
//
// > **NOTE:** If the board lists are changed all lists will be recreated.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/boards/)
//
// ## Import
//
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_issue_board`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_issue_board.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// You can import this resource with an id made up of `{project-id}:{issue-board-id}`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectIssueBoard:ProjectIssueBoard kanban 42:1
// ```
type ProjectIssueBoard struct {
	pulumi.CustomResourceState

	// The assignee the board should be scoped to. Requires a GitLab EE license.
	AssigneeId pulumi.IntPtrOutput `pulumi:"assigneeId"`
	// The list of label names which the board should be scoped to. Requires a GitLab EE license.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The list of issue board lists
	Lists ProjectIssueBoardListArrayOutput `pulumi:"lists"`
	// The milestone the board should be scoped to. Requires a GitLab EE license.
	MilestoneId pulumi.IntPtrOutput `pulumi:"milestoneId"`
	// The name of the board.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringOutput `pulumi:"project"`
	// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewProjectIssueBoard registers a new resource with the given unique name, arguments, and options.
func NewProjectIssueBoard(ctx *pulumi.Context,
	name string, args *ProjectIssueBoardArgs, opts ...pulumi.ResourceOption) (*ProjectIssueBoard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectIssueBoard
	err := ctx.RegisterResource("gitlab:index/projectIssueBoard:ProjectIssueBoard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectIssueBoard gets an existing ProjectIssueBoard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectIssueBoard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectIssueBoardState, opts ...pulumi.ResourceOption) (*ProjectIssueBoard, error) {
	var resource ProjectIssueBoard
	err := ctx.ReadResource("gitlab:index/projectIssueBoard:ProjectIssueBoard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectIssueBoard resources.
type projectIssueBoardState struct {
	// The assignee the board should be scoped to. Requires a GitLab EE license.
	AssigneeId *int `pulumi:"assigneeId"`
	// The list of label names which the board should be scoped to. Requires a GitLab EE license.
	Labels []string `pulumi:"labels"`
	// The list of issue board lists
	Lists []ProjectIssueBoardList `pulumi:"lists"`
	// The milestone the board should be scoped to. Requires a GitLab EE license.
	MilestoneId *int `pulumi:"milestoneId"`
	// The name of the board.
	Name *string `pulumi:"name"`
	// The ID or full path of the project maintained by the authenticated user.
	Project *string `pulumi:"project"`
	// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
	Weight *int `pulumi:"weight"`
}

type ProjectIssueBoardState struct {
	// The assignee the board should be scoped to. Requires a GitLab EE license.
	AssigneeId pulumi.IntPtrInput
	// The list of label names which the board should be scoped to. Requires a GitLab EE license.
	Labels pulumi.StringArrayInput
	// The list of issue board lists
	Lists ProjectIssueBoardListArrayInput
	// The milestone the board should be scoped to. Requires a GitLab EE license.
	MilestoneId pulumi.IntPtrInput
	// The name of the board.
	Name pulumi.StringPtrInput
	// The ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringPtrInput
	// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
	Weight pulumi.IntPtrInput
}

func (ProjectIssueBoardState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIssueBoardState)(nil)).Elem()
}

type projectIssueBoardArgs struct {
	// The assignee the board should be scoped to. Requires a GitLab EE license.
	AssigneeId *int `pulumi:"assigneeId"`
	// The list of label names which the board should be scoped to. Requires a GitLab EE license.
	Labels []string `pulumi:"labels"`
	// The list of issue board lists
	Lists []ProjectIssueBoardList `pulumi:"lists"`
	// The milestone the board should be scoped to. Requires a GitLab EE license.
	MilestoneId *int `pulumi:"milestoneId"`
	// The name of the board.
	Name *string `pulumi:"name"`
	// The ID or full path of the project maintained by the authenticated user.
	Project string `pulumi:"project"`
	// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a ProjectIssueBoard resource.
type ProjectIssueBoardArgs struct {
	// The assignee the board should be scoped to. Requires a GitLab EE license.
	AssigneeId pulumi.IntPtrInput
	// The list of label names which the board should be scoped to. Requires a GitLab EE license.
	Labels pulumi.StringArrayInput
	// The list of issue board lists
	Lists ProjectIssueBoardListArrayInput
	// The milestone the board should be scoped to. Requires a GitLab EE license.
	MilestoneId pulumi.IntPtrInput
	// The name of the board.
	Name pulumi.StringPtrInput
	// The ID or full path of the project maintained by the authenticated user.
	Project pulumi.StringInput
	// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
	Weight pulumi.IntPtrInput
}

func (ProjectIssueBoardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIssueBoardArgs)(nil)).Elem()
}

type ProjectIssueBoardInput interface {
	pulumi.Input

	ToProjectIssueBoardOutput() ProjectIssueBoardOutput
	ToProjectIssueBoardOutputWithContext(ctx context.Context) ProjectIssueBoardOutput
}

func (*ProjectIssueBoard) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIssueBoard)(nil)).Elem()
}

func (i *ProjectIssueBoard) ToProjectIssueBoardOutput() ProjectIssueBoardOutput {
	return i.ToProjectIssueBoardOutputWithContext(context.Background())
}

func (i *ProjectIssueBoard) ToProjectIssueBoardOutputWithContext(ctx context.Context) ProjectIssueBoardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIssueBoardOutput)
}

// ProjectIssueBoardArrayInput is an input type that accepts ProjectIssueBoardArray and ProjectIssueBoardArrayOutput values.
// You can construct a concrete instance of `ProjectIssueBoardArrayInput` via:
//
//	ProjectIssueBoardArray{ ProjectIssueBoardArgs{...} }
type ProjectIssueBoardArrayInput interface {
	pulumi.Input

	ToProjectIssueBoardArrayOutput() ProjectIssueBoardArrayOutput
	ToProjectIssueBoardArrayOutputWithContext(context.Context) ProjectIssueBoardArrayOutput
}

type ProjectIssueBoardArray []ProjectIssueBoardInput

func (ProjectIssueBoardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIssueBoard)(nil)).Elem()
}

func (i ProjectIssueBoardArray) ToProjectIssueBoardArrayOutput() ProjectIssueBoardArrayOutput {
	return i.ToProjectIssueBoardArrayOutputWithContext(context.Background())
}

func (i ProjectIssueBoardArray) ToProjectIssueBoardArrayOutputWithContext(ctx context.Context) ProjectIssueBoardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIssueBoardArrayOutput)
}

// ProjectIssueBoardMapInput is an input type that accepts ProjectIssueBoardMap and ProjectIssueBoardMapOutput values.
// You can construct a concrete instance of `ProjectIssueBoardMapInput` via:
//
//	ProjectIssueBoardMap{ "key": ProjectIssueBoardArgs{...} }
type ProjectIssueBoardMapInput interface {
	pulumi.Input

	ToProjectIssueBoardMapOutput() ProjectIssueBoardMapOutput
	ToProjectIssueBoardMapOutputWithContext(context.Context) ProjectIssueBoardMapOutput
}

type ProjectIssueBoardMap map[string]ProjectIssueBoardInput

func (ProjectIssueBoardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIssueBoard)(nil)).Elem()
}

func (i ProjectIssueBoardMap) ToProjectIssueBoardMapOutput() ProjectIssueBoardMapOutput {
	return i.ToProjectIssueBoardMapOutputWithContext(context.Background())
}

func (i ProjectIssueBoardMap) ToProjectIssueBoardMapOutputWithContext(ctx context.Context) ProjectIssueBoardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIssueBoardMapOutput)
}

type ProjectIssueBoardOutput struct{ *pulumi.OutputState }

func (ProjectIssueBoardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIssueBoard)(nil)).Elem()
}

func (o ProjectIssueBoardOutput) ToProjectIssueBoardOutput() ProjectIssueBoardOutput {
	return o
}

func (o ProjectIssueBoardOutput) ToProjectIssueBoardOutputWithContext(ctx context.Context) ProjectIssueBoardOutput {
	return o
}

// The assignee the board should be scoped to. Requires a GitLab EE license.
func (o ProjectIssueBoardOutput) AssigneeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) pulumi.IntPtrOutput { return v.AssigneeId }).(pulumi.IntPtrOutput)
}

// The list of label names which the board should be scoped to. Requires a GitLab EE license.
func (o ProjectIssueBoardOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The list of issue board lists
func (o ProjectIssueBoardOutput) Lists() ProjectIssueBoardListArrayOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) ProjectIssueBoardListArrayOutput { return v.Lists }).(ProjectIssueBoardListArrayOutput)
}

// The milestone the board should be scoped to. Requires a GitLab EE license.
func (o ProjectIssueBoardOutput) MilestoneId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) pulumi.IntPtrOutput { return v.MilestoneId }).(pulumi.IntPtrOutput)
}

// The name of the board.
func (o ProjectIssueBoardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID or full path of the project maintained by the authenticated user.
func (o ProjectIssueBoardOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The weight range from 0 to 9, to which the board should be scoped to. Requires a GitLab EE license.
func (o ProjectIssueBoardOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectIssueBoard) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

type ProjectIssueBoardArrayOutput struct{ *pulumi.OutputState }

func (ProjectIssueBoardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIssueBoard)(nil)).Elem()
}

func (o ProjectIssueBoardArrayOutput) ToProjectIssueBoardArrayOutput() ProjectIssueBoardArrayOutput {
	return o
}

func (o ProjectIssueBoardArrayOutput) ToProjectIssueBoardArrayOutputWithContext(ctx context.Context) ProjectIssueBoardArrayOutput {
	return o
}

func (o ProjectIssueBoardArrayOutput) Index(i pulumi.IntInput) ProjectIssueBoardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectIssueBoard {
		return vs[0].([]*ProjectIssueBoard)[vs[1].(int)]
	}).(ProjectIssueBoardOutput)
}

type ProjectIssueBoardMapOutput struct{ *pulumi.OutputState }

func (ProjectIssueBoardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIssueBoard)(nil)).Elem()
}

func (o ProjectIssueBoardMapOutput) ToProjectIssueBoardMapOutput() ProjectIssueBoardMapOutput {
	return o
}

func (o ProjectIssueBoardMapOutput) ToProjectIssueBoardMapOutputWithContext(ctx context.Context) ProjectIssueBoardMapOutput {
	return o
}

func (o ProjectIssueBoardMapOutput) MapIndex(k pulumi.StringInput) ProjectIssueBoardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectIssueBoard {
		return vs[0].(map[string]*ProjectIssueBoard)[vs[1].(string)]
	}).(ProjectIssueBoardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIssueBoardInput)(nil)).Elem(), &ProjectIssueBoard{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIssueBoardArrayInput)(nil)).Elem(), ProjectIssueBoardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIssueBoardMapInput)(nil)).Elem(), ProjectIssueBoardMap{})
	pulumi.RegisterOutputType(ProjectIssueBoardOutput{})
	pulumi.RegisterOutputType(ProjectIssueBoardArrayOutput{})
	pulumi.RegisterOutputType(ProjectIssueBoardMapOutput{})
}
