// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource allows you to create and manage pipeline schedules.
// For further information on clusters, consult the [gitlab
// documentation](https://docs.gitlab.com/ce/user/project/pipelines/schedules.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/pipeline_schedule.html.markdown.
type PipelineSchedule struct {
	pulumi.CustomResourceState

	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// The cron (e.g. `0 1 * * *`).
	Cron pulumi.StringOutput `pulumi:"cron"`
	// The timezone.
	CronTimezone pulumi.StringPtrOutput `pulumi:"cronTimezone"`
	// The description of the pipeline schedule.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name or id of the project to add the schedule to.
	Project pulumi.StringOutput `pulumi:"project"`
	// The branch/tag name to be triggered.
	Ref pulumi.StringOutput `pulumi:"ref"`
}

// NewPipelineSchedule registers a new resource with the given unique name, arguments, and options.
func NewPipelineSchedule(ctx *pulumi.Context,
	name string, args *PipelineScheduleArgs, opts ...pulumi.ResourceOption) (*PipelineSchedule, error) {
	if args == nil || args.Cron == nil {
		return nil, errors.New("missing required argument 'Cron'")
	}
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Ref == nil {
		return nil, errors.New("missing required argument 'Ref'")
	}
	if args == nil {
		args = &PipelineScheduleArgs{}
	}
	var resource PipelineSchedule
	err := ctx.RegisterResource("gitlab:index/pipelineSchedule:PipelineSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineSchedule gets an existing PipelineSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineScheduleState, opts ...pulumi.ResourceOption) (*PipelineSchedule, error) {
	var resource PipelineSchedule
	err := ctx.ReadResource("gitlab:index/pipelineSchedule:PipelineSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineSchedule resources.
type pipelineScheduleState struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active *bool `pulumi:"active"`
	// The cron (e.g. `0 1 * * *`).
	Cron *string `pulumi:"cron"`
	// The timezone.
	CronTimezone *string `pulumi:"cronTimezone"`
	// The description of the pipeline schedule.
	Description *string `pulumi:"description"`
	// The name or id of the project to add the schedule to.
	Project *string `pulumi:"project"`
	// The branch/tag name to be triggered.
	Ref *string `pulumi:"ref"`
}

type PipelineScheduleState struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active pulumi.BoolPtrInput
	// The cron (e.g. `0 1 * * *`).
	Cron pulumi.StringPtrInput
	// The timezone.
	CronTimezone pulumi.StringPtrInput
	// The description of the pipeline schedule.
	Description pulumi.StringPtrInput
	// The name or id of the project to add the schedule to.
	Project pulumi.StringPtrInput
	// The branch/tag name to be triggered.
	Ref pulumi.StringPtrInput
}

func (PipelineScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleState)(nil)).Elem()
}

type pipelineScheduleArgs struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active *bool `pulumi:"active"`
	// The cron (e.g. `0 1 * * *`).
	Cron string `pulumi:"cron"`
	// The timezone.
	CronTimezone *string `pulumi:"cronTimezone"`
	// The description of the pipeline schedule.
	Description string `pulumi:"description"`
	// The name or id of the project to add the schedule to.
	Project string `pulumi:"project"`
	// The branch/tag name to be triggered.
	Ref string `pulumi:"ref"`
}

// The set of arguments for constructing a PipelineSchedule resource.
type PipelineScheduleArgs struct {
	// The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
	Active pulumi.BoolPtrInput
	// The cron (e.g. `0 1 * * *`).
	Cron pulumi.StringInput
	// The timezone.
	CronTimezone pulumi.StringPtrInput
	// The description of the pipeline schedule.
	Description pulumi.StringInput
	// The name or id of the project to add the schedule to.
	Project pulumi.StringInput
	// The branch/tag name to be triggered.
	Ref pulumi.StringInput
}

func (PipelineScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineScheduleArgs)(nil)).Elem()
}
