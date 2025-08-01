// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectFreezePeriod` resource allows to manage the lifecycle of a freeze period for a project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/freeze_periods/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewProjectFreezePeriod(ctx, "schedule", &gitlab.ProjectFreezePeriodArgs{
//				Project:      pulumi.Any(foo.Id),
//				FreezeStart:  pulumi.String("0 23 * * 5"),
//				FreezeEnd:    pulumi.String("0 7 * * 1"),
//				CronTimezone: pulumi.String("UTC"),
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
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_freeze_period`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_freeze_period.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// GitLab project freeze periods can be imported using an id made up of `project_id:freeze_period_id`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectFreezePeriod:ProjectFreezePeriod schedule "12345:1337"
// ```
type ProjectFreezePeriod struct {
	pulumi.CustomResourceState

	// The timezone.
	CronTimezone pulumi.StringPtrOutput `pulumi:"cronTimezone"`
	// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
	FreezeEnd pulumi.StringOutput `pulumi:"freezeEnd"`
	// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
	FreezeStart pulumi.StringOutput `pulumi:"freezeStart"`
	// The ID or URL-encoded path of the project to add the schedule to.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectFreezePeriod registers a new resource with the given unique name, arguments, and options.
func NewProjectFreezePeriod(ctx *pulumi.Context,
	name string, args *ProjectFreezePeriodArgs, opts ...pulumi.ResourceOption) (*ProjectFreezePeriod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FreezeEnd == nil {
		return nil, errors.New("invalid value for required argument 'FreezeEnd'")
	}
	if args.FreezeStart == nil {
		return nil, errors.New("invalid value for required argument 'FreezeStart'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectFreezePeriod
	err := ctx.RegisterResource("gitlab:index/projectFreezePeriod:ProjectFreezePeriod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectFreezePeriod gets an existing ProjectFreezePeriod resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectFreezePeriod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectFreezePeriodState, opts ...pulumi.ResourceOption) (*ProjectFreezePeriod, error) {
	var resource ProjectFreezePeriod
	err := ctx.ReadResource("gitlab:index/projectFreezePeriod:ProjectFreezePeriod", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectFreezePeriod resources.
type projectFreezePeriodState struct {
	// The timezone.
	CronTimezone *string `pulumi:"cronTimezone"`
	// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
	FreezeEnd *string `pulumi:"freezeEnd"`
	// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
	FreezeStart *string `pulumi:"freezeStart"`
	// The ID or URL-encoded path of the project to add the schedule to.
	Project *string `pulumi:"project"`
}

type ProjectFreezePeriodState struct {
	// The timezone.
	CronTimezone pulumi.StringPtrInput
	// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
	FreezeEnd pulumi.StringPtrInput
	// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
	FreezeStart pulumi.StringPtrInput
	// The ID or URL-encoded path of the project to add the schedule to.
	Project pulumi.StringPtrInput
}

func (ProjectFreezePeriodState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFreezePeriodState)(nil)).Elem()
}

type projectFreezePeriodArgs struct {
	// The timezone.
	CronTimezone *string `pulumi:"cronTimezone"`
	// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
	FreezeEnd string `pulumi:"freezeEnd"`
	// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
	FreezeStart string `pulumi:"freezeStart"`
	// The ID or URL-encoded path of the project to add the schedule to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectFreezePeriod resource.
type ProjectFreezePeriodArgs struct {
	// The timezone.
	CronTimezone pulumi.StringPtrInput
	// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
	FreezeEnd pulumi.StringInput
	// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
	FreezeStart pulumi.StringInput
	// The ID or URL-encoded path of the project to add the schedule to.
	Project pulumi.StringInput
}

func (ProjectFreezePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFreezePeriodArgs)(nil)).Elem()
}

type ProjectFreezePeriodInput interface {
	pulumi.Input

	ToProjectFreezePeriodOutput() ProjectFreezePeriodOutput
	ToProjectFreezePeriodOutputWithContext(ctx context.Context) ProjectFreezePeriodOutput
}

func (*ProjectFreezePeriod) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFreezePeriod)(nil)).Elem()
}

func (i *ProjectFreezePeriod) ToProjectFreezePeriodOutput() ProjectFreezePeriodOutput {
	return i.ToProjectFreezePeriodOutputWithContext(context.Background())
}

func (i *ProjectFreezePeriod) ToProjectFreezePeriodOutputWithContext(ctx context.Context) ProjectFreezePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFreezePeriodOutput)
}

// ProjectFreezePeriodArrayInput is an input type that accepts ProjectFreezePeriodArray and ProjectFreezePeriodArrayOutput values.
// You can construct a concrete instance of `ProjectFreezePeriodArrayInput` via:
//
//	ProjectFreezePeriodArray{ ProjectFreezePeriodArgs{...} }
type ProjectFreezePeriodArrayInput interface {
	pulumi.Input

	ToProjectFreezePeriodArrayOutput() ProjectFreezePeriodArrayOutput
	ToProjectFreezePeriodArrayOutputWithContext(context.Context) ProjectFreezePeriodArrayOutput
}

type ProjectFreezePeriodArray []ProjectFreezePeriodInput

func (ProjectFreezePeriodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectFreezePeriod)(nil)).Elem()
}

func (i ProjectFreezePeriodArray) ToProjectFreezePeriodArrayOutput() ProjectFreezePeriodArrayOutput {
	return i.ToProjectFreezePeriodArrayOutputWithContext(context.Background())
}

func (i ProjectFreezePeriodArray) ToProjectFreezePeriodArrayOutputWithContext(ctx context.Context) ProjectFreezePeriodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFreezePeriodArrayOutput)
}

// ProjectFreezePeriodMapInput is an input type that accepts ProjectFreezePeriodMap and ProjectFreezePeriodMapOutput values.
// You can construct a concrete instance of `ProjectFreezePeriodMapInput` via:
//
//	ProjectFreezePeriodMap{ "key": ProjectFreezePeriodArgs{...} }
type ProjectFreezePeriodMapInput interface {
	pulumi.Input

	ToProjectFreezePeriodMapOutput() ProjectFreezePeriodMapOutput
	ToProjectFreezePeriodMapOutputWithContext(context.Context) ProjectFreezePeriodMapOutput
}

type ProjectFreezePeriodMap map[string]ProjectFreezePeriodInput

func (ProjectFreezePeriodMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectFreezePeriod)(nil)).Elem()
}

func (i ProjectFreezePeriodMap) ToProjectFreezePeriodMapOutput() ProjectFreezePeriodMapOutput {
	return i.ToProjectFreezePeriodMapOutputWithContext(context.Background())
}

func (i ProjectFreezePeriodMap) ToProjectFreezePeriodMapOutputWithContext(ctx context.Context) ProjectFreezePeriodMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFreezePeriodMapOutput)
}

type ProjectFreezePeriodOutput struct{ *pulumi.OutputState }

func (ProjectFreezePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFreezePeriod)(nil)).Elem()
}

func (o ProjectFreezePeriodOutput) ToProjectFreezePeriodOutput() ProjectFreezePeriodOutput {
	return o
}

func (o ProjectFreezePeriodOutput) ToProjectFreezePeriodOutputWithContext(ctx context.Context) ProjectFreezePeriodOutput {
	return o
}

// The timezone.
func (o ProjectFreezePeriodOutput) CronTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectFreezePeriod) pulumi.StringPtrOutput { return v.CronTimezone }).(pulumi.StringPtrOutput)
}

// End of the Freeze Period in cron format (e.g. `0 2 * * *`).
func (o ProjectFreezePeriodOutput) FreezeEnd() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectFreezePeriod) pulumi.StringOutput { return v.FreezeEnd }).(pulumi.StringOutput)
}

// Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
func (o ProjectFreezePeriodOutput) FreezeStart() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectFreezePeriod) pulumi.StringOutput { return v.FreezeStart }).(pulumi.StringOutput)
}

// The ID or URL-encoded path of the project to add the schedule to.
func (o ProjectFreezePeriodOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectFreezePeriod) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ProjectFreezePeriodArrayOutput struct{ *pulumi.OutputState }

func (ProjectFreezePeriodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectFreezePeriod)(nil)).Elem()
}

func (o ProjectFreezePeriodArrayOutput) ToProjectFreezePeriodArrayOutput() ProjectFreezePeriodArrayOutput {
	return o
}

func (o ProjectFreezePeriodArrayOutput) ToProjectFreezePeriodArrayOutputWithContext(ctx context.Context) ProjectFreezePeriodArrayOutput {
	return o
}

func (o ProjectFreezePeriodArrayOutput) Index(i pulumi.IntInput) ProjectFreezePeriodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectFreezePeriod {
		return vs[0].([]*ProjectFreezePeriod)[vs[1].(int)]
	}).(ProjectFreezePeriodOutput)
}

type ProjectFreezePeriodMapOutput struct{ *pulumi.OutputState }

func (ProjectFreezePeriodMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectFreezePeriod)(nil)).Elem()
}

func (o ProjectFreezePeriodMapOutput) ToProjectFreezePeriodMapOutput() ProjectFreezePeriodMapOutput {
	return o
}

func (o ProjectFreezePeriodMapOutput) ToProjectFreezePeriodMapOutputWithContext(ctx context.Context) ProjectFreezePeriodMapOutput {
	return o
}

func (o ProjectFreezePeriodMapOutput) MapIndex(k pulumi.StringInput) ProjectFreezePeriodOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectFreezePeriod {
		return vs[0].(map[string]*ProjectFreezePeriod)[vs[1].(string)]
	}).(ProjectFreezePeriodOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFreezePeriodInput)(nil)).Elem(), &ProjectFreezePeriod{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFreezePeriodArrayInput)(nil)).Elem(), ProjectFreezePeriodArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFreezePeriodMapInput)(nil)).Elem(), ProjectFreezePeriodMap{})
	pulumi.RegisterOutputType(ProjectFreezePeriodOutput{})
	pulumi.RegisterOutputType(ProjectFreezePeriodArrayOutput{})
	pulumi.RegisterOutputType(ProjectFreezePeriodMapOutput{})
}
