// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to manage Microsoft Teams integration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		awesomeProject, err := gitlab.NewProject(ctx, "awesomeProject", &gitlab.ProjectArgs{
// 			Description:     pulumi.String("My awesome project."),
// 			VisibilityLevel: pulumi.String("public"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewServiceMicrosoftTeams(ctx, "teams", &gitlab.ServiceMicrosoftTeamsArgs{
// 			Project:    awesomeProject.ID(),
// 			Webhook:    pulumi.String("https://testurl.com/?token=XYZ"),
// 			PushEvents: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # You can import a service_microsoft_teams state using the project ID, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams teams 1
// ```
type ServiceMicrosoftTeams struct {
	pulumi.CustomResourceState

	// Whether the integration is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
	BranchesToBeNotified pulumi.StringPtrOutput `pulumi:"branchesToBeNotified"`
	// Enable notifications for confidential issue events
	ConfidentialIssuesEvents pulumi.BoolPtrOutput `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events
	ConfidentialNoteEvents pulumi.BoolPtrOutput `pulumi:"confidentialNoteEvents"`
	// Create time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Enable notifications for issue events
	IssuesEvents pulumi.BoolPtrOutput `pulumi:"issuesEvents"`
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrOutput `pulumi:"mergeRequestsEvents"`
	// Enable notifications for note events
	NoteEvents pulumi.BoolPtrOutput `pulumi:"noteEvents"`
	// Send notifications for broken pipelines
	NotifyOnlyBrokenPipelines pulumi.BoolPtrOutput `pulumi:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events
	PipelineEvents pulumi.BoolPtrOutput `pulumi:"pipelineEvents"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// Enable notifications for push events
	PushEvents pulumi.BoolPtrOutput `pulumi:"pushEvents"`
	// Enable notifications for tag push events
	TagPushEvents pulumi.BoolPtrOutput `pulumi:"tagPushEvents"`
	// Update time.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The Microsoft Teams webhook. For example, https://outlook.office.com/webhook/...
	Webhook pulumi.StringOutput `pulumi:"webhook"`
	// Enable notifications for wiki page events
	WikiPageEvents pulumi.BoolPtrOutput `pulumi:"wikiPageEvents"`
}

// NewServiceMicrosoftTeams registers a new resource with the given unique name, arguments, and options.
func NewServiceMicrosoftTeams(ctx *pulumi.Context,
	name string, args *ServiceMicrosoftTeamsArgs, opts ...pulumi.ResourceOption) (*ServiceMicrosoftTeams, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Webhook == nil {
		return nil, errors.New("invalid value for required argument 'Webhook'")
	}
	var resource ServiceMicrosoftTeams
	err := ctx.RegisterResource("gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceMicrosoftTeams gets an existing ServiceMicrosoftTeams resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceMicrosoftTeams(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceMicrosoftTeamsState, opts ...pulumi.ResourceOption) (*ServiceMicrosoftTeams, error) {
	var resource ServiceMicrosoftTeams
	err := ctx.ReadResource("gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceMicrosoftTeams resources.
type serviceMicrosoftTeamsState struct {
	// Whether the integration is active.
	Active *bool `pulumi:"active"`
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Enable notifications for confidential issue events
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Create time.
	CreatedAt *string `pulumi:"createdAt"`
	// Enable notifications for issue events
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Enable notifications for merge request events
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Enable notifications for note events
	NoteEvents *bool `pulumi:"noteEvents"`
	// Send notifications for broken pipelines
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// Enable notifications for push events
	PushEvents *bool `pulumi:"pushEvents"`
	// Enable notifications for tag push events
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// Update time.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The Microsoft Teams webhook. For example, https://outlook.office.com/webhook/...
	Webhook *string `pulumi:"webhook"`
	// Enable notifications for wiki page events
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type ServiceMicrosoftTeamsState struct {
	// Whether the integration is active.
	Active pulumi.BoolPtrInput
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
	BranchesToBeNotified pulumi.StringPtrInput
	// Enable notifications for confidential issue events
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Enable notifications for confidential note events
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Create time.
	CreatedAt pulumi.StringPtrInput
	// Enable notifications for issue events
	IssuesEvents pulumi.BoolPtrInput
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrInput
	// Enable notifications for note events
	NoteEvents pulumi.BoolPtrInput
	// Send notifications for broken pipelines
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// Enable notifications for pipeline events
	PipelineEvents pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// Enable notifications for push events
	PushEvents pulumi.BoolPtrInput
	// Enable notifications for tag push events
	TagPushEvents pulumi.BoolPtrInput
	// Update time.
	UpdatedAt pulumi.StringPtrInput
	// The Microsoft Teams webhook. For example, https://outlook.office.com/webhook/...
	Webhook pulumi.StringPtrInput
	// Enable notifications for wiki page events
	WikiPageEvents pulumi.BoolPtrInput
}

func (ServiceMicrosoftTeamsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceMicrosoftTeamsState)(nil)).Elem()
}

type serviceMicrosoftTeamsArgs struct {
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Enable notifications for confidential issue events
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Enable notifications for issue events
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Enable notifications for merge request events
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Enable notifications for note events
	NoteEvents *bool `pulumi:"noteEvents"`
	// Send notifications for broken pipelines
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// Enable notifications for push events
	PushEvents *bool `pulumi:"pushEvents"`
	// Enable notifications for tag push events
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// The Microsoft Teams webhook. For example, https://outlook.office.com/webhook/...
	Webhook string `pulumi:"webhook"`
	// Enable notifications for wiki page events
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a ServiceMicrosoftTeams resource.
type ServiceMicrosoftTeamsArgs struct {
	// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
	BranchesToBeNotified pulumi.StringPtrInput
	// Enable notifications for confidential issue events
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Enable notifications for confidential note events
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Enable notifications for issue events
	IssuesEvents pulumi.BoolPtrInput
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrInput
	// Enable notifications for note events
	NoteEvents pulumi.BoolPtrInput
	// Send notifications for broken pipelines
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// Enable notifications for pipeline events
	PipelineEvents pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// Enable notifications for push events
	PushEvents pulumi.BoolPtrInput
	// Enable notifications for tag push events
	TagPushEvents pulumi.BoolPtrInput
	// The Microsoft Teams webhook. For example, https://outlook.office.com/webhook/...
	Webhook pulumi.StringInput
	// Enable notifications for wiki page events
	WikiPageEvents pulumi.BoolPtrInput
}

func (ServiceMicrosoftTeamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceMicrosoftTeamsArgs)(nil)).Elem()
}

type ServiceMicrosoftTeamsInput interface {
	pulumi.Input

	ToServiceMicrosoftTeamsOutput() ServiceMicrosoftTeamsOutput
	ToServiceMicrosoftTeamsOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsOutput
}

func (*ServiceMicrosoftTeams) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceMicrosoftTeams)(nil)).Elem()
}

func (i *ServiceMicrosoftTeams) ToServiceMicrosoftTeamsOutput() ServiceMicrosoftTeamsOutput {
	return i.ToServiceMicrosoftTeamsOutputWithContext(context.Background())
}

func (i *ServiceMicrosoftTeams) ToServiceMicrosoftTeamsOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMicrosoftTeamsOutput)
}

// ServiceMicrosoftTeamsArrayInput is an input type that accepts ServiceMicrosoftTeamsArray and ServiceMicrosoftTeamsArrayOutput values.
// You can construct a concrete instance of `ServiceMicrosoftTeamsArrayInput` via:
//
//          ServiceMicrosoftTeamsArray{ ServiceMicrosoftTeamsArgs{...} }
type ServiceMicrosoftTeamsArrayInput interface {
	pulumi.Input

	ToServiceMicrosoftTeamsArrayOutput() ServiceMicrosoftTeamsArrayOutput
	ToServiceMicrosoftTeamsArrayOutputWithContext(context.Context) ServiceMicrosoftTeamsArrayOutput
}

type ServiceMicrosoftTeamsArray []ServiceMicrosoftTeamsInput

func (ServiceMicrosoftTeamsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceMicrosoftTeams)(nil)).Elem()
}

func (i ServiceMicrosoftTeamsArray) ToServiceMicrosoftTeamsArrayOutput() ServiceMicrosoftTeamsArrayOutput {
	return i.ToServiceMicrosoftTeamsArrayOutputWithContext(context.Background())
}

func (i ServiceMicrosoftTeamsArray) ToServiceMicrosoftTeamsArrayOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMicrosoftTeamsArrayOutput)
}

// ServiceMicrosoftTeamsMapInput is an input type that accepts ServiceMicrosoftTeamsMap and ServiceMicrosoftTeamsMapOutput values.
// You can construct a concrete instance of `ServiceMicrosoftTeamsMapInput` via:
//
//          ServiceMicrosoftTeamsMap{ "key": ServiceMicrosoftTeamsArgs{...} }
type ServiceMicrosoftTeamsMapInput interface {
	pulumi.Input

	ToServiceMicrosoftTeamsMapOutput() ServiceMicrosoftTeamsMapOutput
	ToServiceMicrosoftTeamsMapOutputWithContext(context.Context) ServiceMicrosoftTeamsMapOutput
}

type ServiceMicrosoftTeamsMap map[string]ServiceMicrosoftTeamsInput

func (ServiceMicrosoftTeamsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceMicrosoftTeams)(nil)).Elem()
}

func (i ServiceMicrosoftTeamsMap) ToServiceMicrosoftTeamsMapOutput() ServiceMicrosoftTeamsMapOutput {
	return i.ToServiceMicrosoftTeamsMapOutputWithContext(context.Background())
}

func (i ServiceMicrosoftTeamsMap) ToServiceMicrosoftTeamsMapOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMicrosoftTeamsMapOutput)
}

type ServiceMicrosoftTeamsOutput struct{ *pulumi.OutputState }

func (ServiceMicrosoftTeamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceMicrosoftTeams)(nil)).Elem()
}

func (o ServiceMicrosoftTeamsOutput) ToServiceMicrosoftTeamsOutput() ServiceMicrosoftTeamsOutput {
	return o
}

func (o ServiceMicrosoftTeamsOutput) ToServiceMicrosoftTeamsOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsOutput {
	return o
}

type ServiceMicrosoftTeamsArrayOutput struct{ *pulumi.OutputState }

func (ServiceMicrosoftTeamsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceMicrosoftTeams)(nil)).Elem()
}

func (o ServiceMicrosoftTeamsArrayOutput) ToServiceMicrosoftTeamsArrayOutput() ServiceMicrosoftTeamsArrayOutput {
	return o
}

func (o ServiceMicrosoftTeamsArrayOutput) ToServiceMicrosoftTeamsArrayOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsArrayOutput {
	return o
}

func (o ServiceMicrosoftTeamsArrayOutput) Index(i pulumi.IntInput) ServiceMicrosoftTeamsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceMicrosoftTeams {
		return vs[0].([]*ServiceMicrosoftTeams)[vs[1].(int)]
	}).(ServiceMicrosoftTeamsOutput)
}

type ServiceMicrosoftTeamsMapOutput struct{ *pulumi.OutputState }

func (ServiceMicrosoftTeamsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceMicrosoftTeams)(nil)).Elem()
}

func (o ServiceMicrosoftTeamsMapOutput) ToServiceMicrosoftTeamsMapOutput() ServiceMicrosoftTeamsMapOutput {
	return o
}

func (o ServiceMicrosoftTeamsMapOutput) ToServiceMicrosoftTeamsMapOutputWithContext(ctx context.Context) ServiceMicrosoftTeamsMapOutput {
	return o
}

func (o ServiceMicrosoftTeamsMapOutput) MapIndex(k pulumi.StringInput) ServiceMicrosoftTeamsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceMicrosoftTeams {
		return vs[0].(map[string]*ServiceMicrosoftTeams)[vs[1].(string)]
	}).(ServiceMicrosoftTeamsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMicrosoftTeamsInput)(nil)).Elem(), &ServiceMicrosoftTeams{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMicrosoftTeamsArrayInput)(nil)).Elem(), ServiceMicrosoftTeamsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMicrosoftTeamsMapInput)(nil)).Elem(), ServiceMicrosoftTeamsMap{})
	pulumi.RegisterOutputType(ServiceMicrosoftTeamsOutput{})
	pulumi.RegisterOutputType(ServiceMicrosoftTeamsArrayOutput{})
	pulumi.RegisterOutputType(ServiceMicrosoftTeamsMapOutput{})
}