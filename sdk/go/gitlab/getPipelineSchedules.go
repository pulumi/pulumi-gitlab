// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `PipelineSchedule` data source retrieves information about a gitlab pipeline schedule for a project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/pipeline_schedules.html)
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
//			_, err := gitlab.GetPipelineSchedules(ctx, &gitlab.GetPipelineSchedulesArgs{
//				Project: "12345",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPipelineSchedules(ctx *pulumi.Context, args *GetPipelineSchedulesArgs, opts ...pulumi.InvokeOption) (*GetPipelineSchedulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPipelineSchedulesResult
	err := ctx.Invoke("gitlab:index/getPipelineSchedules:getPipelineSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPipelineSchedules.
type GetPipelineSchedulesArgs struct {
	// The name or id of the project to add the schedule to.
	Project string `pulumi:"project"`
}

// A collection of values returned by getPipelineSchedules.
type GetPipelineSchedulesResult struct {
	Id string `pulumi:"id"`
	// The list of pipeline schedules.
	PipelineSchedules []GetPipelineSchedulesPipelineSchedule `pulumi:"pipelineSchedules"`
	// The name or id of the project to add the schedule to.
	Project string `pulumi:"project"`
}

func GetPipelineSchedulesOutput(ctx *pulumi.Context, args GetPipelineSchedulesOutputArgs, opts ...pulumi.InvokeOption) GetPipelineSchedulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPipelineSchedulesResult, error) {
			args := v.(GetPipelineSchedulesArgs)
			r, err := GetPipelineSchedules(ctx, &args, opts...)
			var s GetPipelineSchedulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPipelineSchedulesResultOutput)
}

// A collection of arguments for invoking getPipelineSchedules.
type GetPipelineSchedulesOutputArgs struct {
	// The name or id of the project to add the schedule to.
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetPipelineSchedulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPipelineSchedulesArgs)(nil)).Elem()
}

// A collection of values returned by getPipelineSchedules.
type GetPipelineSchedulesResultOutput struct{ *pulumi.OutputState }

func (GetPipelineSchedulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPipelineSchedulesResult)(nil)).Elem()
}

func (o GetPipelineSchedulesResultOutput) ToGetPipelineSchedulesResultOutput() GetPipelineSchedulesResultOutput {
	return o
}

func (o GetPipelineSchedulesResultOutput) ToGetPipelineSchedulesResultOutputWithContext(ctx context.Context) GetPipelineSchedulesResultOutput {
	return o
}

func (o GetPipelineSchedulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPipelineSchedulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of pipeline schedules.
func (o GetPipelineSchedulesResultOutput) PipelineSchedules() GetPipelineSchedulesPipelineScheduleArrayOutput {
	return o.ApplyT(func(v GetPipelineSchedulesResult) []GetPipelineSchedulesPipelineSchedule { return v.PipelineSchedules }).(GetPipelineSchedulesPipelineScheduleArrayOutput)
}

// The name or id of the project to add the schedule to.
func (o GetPipelineSchedulesResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetPipelineSchedulesResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPipelineSchedulesResultOutput{})
}
