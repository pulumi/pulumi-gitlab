// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource allows you to create and manage variables for pipeline schedules.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/pipeline_schedule_variable.html.markdown.
type PipelineScheduleVariable struct {
	s *pulumi.ResourceState
}

// NewPipelineScheduleVariable registers a new resource with the given unique name, arguments, and options.
func NewPipelineScheduleVariable(ctx *pulumi.Context,
	name string, args *PipelineScheduleVariableArgs, opts ...pulumi.ResourceOpt) (*PipelineScheduleVariable, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.PipelineScheduleId == nil {
		return nil, errors.New("missing required argument 'PipelineScheduleId'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["key"] = nil
		inputs["pipelineScheduleId"] = nil
		inputs["project"] = nil
		inputs["value"] = nil
	} else {
		inputs["key"] = args.Key
		inputs["pipelineScheduleId"] = args.PipelineScheduleId
		inputs["project"] = args.Project
		inputs["value"] = args.Value
	}
	s, err := ctx.RegisterResource("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PipelineScheduleVariable{s: s}, nil
}

// GetPipelineScheduleVariable gets an existing PipelineScheduleVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineScheduleVariable(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PipelineScheduleVariableState, opts ...pulumi.ResourceOpt) (*PipelineScheduleVariable, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["key"] = state.Key
		inputs["pipelineScheduleId"] = state.PipelineScheduleId
		inputs["project"] = state.Project
		inputs["value"] = state.Value
	}
	s, err := ctx.ReadResource("gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PipelineScheduleVariable{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PipelineScheduleVariable) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PipelineScheduleVariable) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Name of the variable.
func (r *PipelineScheduleVariable) Key() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["key"])
}

// The id of the pipeline schedule.
func (r *PipelineScheduleVariable) PipelineScheduleId() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["pipelineScheduleId"])
}

// The id of the project to add the schedule to.
func (r *PipelineScheduleVariable) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Value of the variable.
func (r *PipelineScheduleVariable) Value() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["value"])
}

// Input properties used for looking up and filtering PipelineScheduleVariable resources.
type PipelineScheduleVariableState struct {
	// Name of the variable.
	Key interface{}
	// The id of the pipeline schedule.
	PipelineScheduleId interface{}
	// The id of the project to add the schedule to.
	Project interface{}
	// Value of the variable.
	Value interface{}
}

// The set of arguments for constructing a PipelineScheduleVariable resource.
type PipelineScheduleVariableArgs struct {
	// Name of the variable.
	Key interface{}
	// The id of the pipeline schedule.
	PipelineScheduleId interface{}
	// The id of the project to add the schedule to.
	Project interface{}
	// Value of the variable.
	Value interface{}
}
