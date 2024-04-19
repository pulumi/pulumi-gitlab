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

// The `IntegrationMicrosoftTeams` resource allows to manage the lifecycle of a project integration with Microsoft Teams.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#microsoft-teams)
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
//			awesomeProject, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewIntegrationMicrosoftTeams(ctx, "teams", &gitlab.IntegrationMicrosoftTeamsArgs{
//				Project:    awesomeProject.ID(),
//				Webhook:    pulumi.String("https://testurl.com/?token=XYZ"),
//				PushEvents: pulumi.Bool(true),
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
// You can import a gitlab_integration_microsoft_teams state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams teams 1
// ```
type IntegrationMicrosoftTeams struct {
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
	// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
	Webhook pulumi.StringOutput `pulumi:"webhook"`
	// Enable notifications for wiki page events
	WikiPageEvents pulumi.BoolPtrOutput `pulumi:"wikiPageEvents"`
}

// NewIntegrationMicrosoftTeams registers a new resource with the given unique name, arguments, and options.
func NewIntegrationMicrosoftTeams(ctx *pulumi.Context,
	name string, args *IntegrationMicrosoftTeamsArgs, opts ...pulumi.ResourceOption) (*IntegrationMicrosoftTeams, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Webhook == nil {
		return nil, errors.New("invalid value for required argument 'Webhook'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationMicrosoftTeams
	err := ctx.RegisterResource("gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationMicrosoftTeams gets an existing IntegrationMicrosoftTeams resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationMicrosoftTeams(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationMicrosoftTeamsState, opts ...pulumi.ResourceOption) (*IntegrationMicrosoftTeams, error) {
	var resource IntegrationMicrosoftTeams
	err := ctx.ReadResource("gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationMicrosoftTeams resources.
type integrationMicrosoftTeamsState struct {
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
	// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
	Webhook *string `pulumi:"webhook"`
	// Enable notifications for wiki page events
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type IntegrationMicrosoftTeamsState struct {
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
	// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
	Webhook pulumi.StringPtrInput
	// Enable notifications for wiki page events
	WikiPageEvents pulumi.BoolPtrInput
}

func (IntegrationMicrosoftTeamsState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMicrosoftTeamsState)(nil)).Elem()
}

type integrationMicrosoftTeamsArgs struct {
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
	// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
	Webhook string `pulumi:"webhook"`
	// Enable notifications for wiki page events
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a IntegrationMicrosoftTeams resource.
type IntegrationMicrosoftTeamsArgs struct {
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
	// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
	Webhook pulumi.StringInput
	// Enable notifications for wiki page events
	WikiPageEvents pulumi.BoolPtrInput
}

func (IntegrationMicrosoftTeamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationMicrosoftTeamsArgs)(nil)).Elem()
}

type IntegrationMicrosoftTeamsInput interface {
	pulumi.Input

	ToIntegrationMicrosoftTeamsOutput() IntegrationMicrosoftTeamsOutput
	ToIntegrationMicrosoftTeamsOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsOutput
}

func (*IntegrationMicrosoftTeams) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationMicrosoftTeams)(nil)).Elem()
}

func (i *IntegrationMicrosoftTeams) ToIntegrationMicrosoftTeamsOutput() IntegrationMicrosoftTeamsOutput {
	return i.ToIntegrationMicrosoftTeamsOutputWithContext(context.Background())
}

func (i *IntegrationMicrosoftTeams) ToIntegrationMicrosoftTeamsOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMicrosoftTeamsOutput)
}

// IntegrationMicrosoftTeamsArrayInput is an input type that accepts IntegrationMicrosoftTeamsArray and IntegrationMicrosoftTeamsArrayOutput values.
// You can construct a concrete instance of `IntegrationMicrosoftTeamsArrayInput` via:
//
//	IntegrationMicrosoftTeamsArray{ IntegrationMicrosoftTeamsArgs{...} }
type IntegrationMicrosoftTeamsArrayInput interface {
	pulumi.Input

	ToIntegrationMicrosoftTeamsArrayOutput() IntegrationMicrosoftTeamsArrayOutput
	ToIntegrationMicrosoftTeamsArrayOutputWithContext(context.Context) IntegrationMicrosoftTeamsArrayOutput
}

type IntegrationMicrosoftTeamsArray []IntegrationMicrosoftTeamsInput

func (IntegrationMicrosoftTeamsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationMicrosoftTeams)(nil)).Elem()
}

func (i IntegrationMicrosoftTeamsArray) ToIntegrationMicrosoftTeamsArrayOutput() IntegrationMicrosoftTeamsArrayOutput {
	return i.ToIntegrationMicrosoftTeamsArrayOutputWithContext(context.Background())
}

func (i IntegrationMicrosoftTeamsArray) ToIntegrationMicrosoftTeamsArrayOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMicrosoftTeamsArrayOutput)
}

// IntegrationMicrosoftTeamsMapInput is an input type that accepts IntegrationMicrosoftTeamsMap and IntegrationMicrosoftTeamsMapOutput values.
// You can construct a concrete instance of `IntegrationMicrosoftTeamsMapInput` via:
//
//	IntegrationMicrosoftTeamsMap{ "key": IntegrationMicrosoftTeamsArgs{...} }
type IntegrationMicrosoftTeamsMapInput interface {
	pulumi.Input

	ToIntegrationMicrosoftTeamsMapOutput() IntegrationMicrosoftTeamsMapOutput
	ToIntegrationMicrosoftTeamsMapOutputWithContext(context.Context) IntegrationMicrosoftTeamsMapOutput
}

type IntegrationMicrosoftTeamsMap map[string]IntegrationMicrosoftTeamsInput

func (IntegrationMicrosoftTeamsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationMicrosoftTeams)(nil)).Elem()
}

func (i IntegrationMicrosoftTeamsMap) ToIntegrationMicrosoftTeamsMapOutput() IntegrationMicrosoftTeamsMapOutput {
	return i.ToIntegrationMicrosoftTeamsMapOutputWithContext(context.Background())
}

func (i IntegrationMicrosoftTeamsMap) ToIntegrationMicrosoftTeamsMapOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMicrosoftTeamsMapOutput)
}

type IntegrationMicrosoftTeamsOutput struct{ *pulumi.OutputState }

func (IntegrationMicrosoftTeamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationMicrosoftTeams)(nil)).Elem()
}

func (o IntegrationMicrosoftTeamsOutput) ToIntegrationMicrosoftTeamsOutput() IntegrationMicrosoftTeamsOutput {
	return o
}

func (o IntegrationMicrosoftTeamsOutput) ToIntegrationMicrosoftTeamsOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsOutput {
	return o
}

// Whether the integration is active.
func (o IntegrationMicrosoftTeamsOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Branches to send notifications for. Valid options are “all”, “default”, “protected”, and “default*and*protected”. The default value is “default”
func (o IntegrationMicrosoftTeamsOutput) BranchesToBeNotified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.StringPtrOutput { return v.BranchesToBeNotified }).(pulumi.StringPtrOutput)
}

// Enable notifications for confidential issue events
func (o IntegrationMicrosoftTeamsOutput) ConfidentialIssuesEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolPtrOutput)
}

// Enable notifications for confidential note events
func (o IntegrationMicrosoftTeamsOutput) ConfidentialNoteEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolPtrOutput)
}

// Create time.
func (o IntegrationMicrosoftTeamsOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Enable notifications for issue events
func (o IntegrationMicrosoftTeamsOutput) IssuesEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.IssuesEvents }).(pulumi.BoolPtrOutput)
}

// Enable notifications for merge request events
func (o IntegrationMicrosoftTeamsOutput) MergeRequestsEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.MergeRequestsEvents }).(pulumi.BoolPtrOutput)
}

// Enable notifications for note events
func (o IntegrationMicrosoftTeamsOutput) NoteEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.NoteEvents }).(pulumi.BoolPtrOutput)
}

// Send notifications for broken pipelines
func (o IntegrationMicrosoftTeamsOutput) NotifyOnlyBrokenPipelines() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolPtrOutput)
}

// Enable notifications for pipeline events
func (o IntegrationMicrosoftTeamsOutput) PipelineEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.PipelineEvents }).(pulumi.BoolPtrOutput)
}

// ID of the project you want to activate integration on.
func (o IntegrationMicrosoftTeamsOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Enable notifications for push events
func (o IntegrationMicrosoftTeamsOutput) PushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.PushEvents }).(pulumi.BoolPtrOutput)
}

// Enable notifications for tag push events
func (o IntegrationMicrosoftTeamsOutput) TagPushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.TagPushEvents }).(pulumi.BoolPtrOutput)
}

// Update time.
func (o IntegrationMicrosoftTeamsOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The Microsoft Teams webhook (Example, https://outlook.office.com/webhook/...). This value cannot be imported.
func (o IntegrationMicrosoftTeamsOutput) Webhook() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.StringOutput { return v.Webhook }).(pulumi.StringOutput)
}

// Enable notifications for wiki page events
func (o IntegrationMicrosoftTeamsOutput) WikiPageEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationMicrosoftTeams) pulumi.BoolPtrOutput { return v.WikiPageEvents }).(pulumi.BoolPtrOutput)
}

type IntegrationMicrosoftTeamsArrayOutput struct{ *pulumi.OutputState }

func (IntegrationMicrosoftTeamsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationMicrosoftTeams)(nil)).Elem()
}

func (o IntegrationMicrosoftTeamsArrayOutput) ToIntegrationMicrosoftTeamsArrayOutput() IntegrationMicrosoftTeamsArrayOutput {
	return o
}

func (o IntegrationMicrosoftTeamsArrayOutput) ToIntegrationMicrosoftTeamsArrayOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsArrayOutput {
	return o
}

func (o IntegrationMicrosoftTeamsArrayOutput) Index(i pulumi.IntInput) IntegrationMicrosoftTeamsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationMicrosoftTeams {
		return vs[0].([]*IntegrationMicrosoftTeams)[vs[1].(int)]
	}).(IntegrationMicrosoftTeamsOutput)
}

type IntegrationMicrosoftTeamsMapOutput struct{ *pulumi.OutputState }

func (IntegrationMicrosoftTeamsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationMicrosoftTeams)(nil)).Elem()
}

func (o IntegrationMicrosoftTeamsMapOutput) ToIntegrationMicrosoftTeamsMapOutput() IntegrationMicrosoftTeamsMapOutput {
	return o
}

func (o IntegrationMicrosoftTeamsMapOutput) ToIntegrationMicrosoftTeamsMapOutputWithContext(ctx context.Context) IntegrationMicrosoftTeamsMapOutput {
	return o
}

func (o IntegrationMicrosoftTeamsMapOutput) MapIndex(k pulumi.StringInput) IntegrationMicrosoftTeamsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationMicrosoftTeams {
		return vs[0].(map[string]*IntegrationMicrosoftTeams)[vs[1].(string)]
	}).(IntegrationMicrosoftTeamsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMicrosoftTeamsInput)(nil)).Elem(), &IntegrationMicrosoftTeams{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMicrosoftTeamsArrayInput)(nil)).Elem(), IntegrationMicrosoftTeamsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMicrosoftTeamsMapInput)(nil)).Elem(), IntegrationMicrosoftTeamsMap{})
	pulumi.RegisterOutputType(IntegrationMicrosoftTeamsOutput{})
	pulumi.RegisterOutputType(IntegrationMicrosoftTeamsArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMicrosoftTeamsMapOutput{})
}
