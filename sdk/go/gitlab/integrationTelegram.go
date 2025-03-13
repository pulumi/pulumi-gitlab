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

// The `IntegrationTelegram` resource allows to manage the lifecycle of a project integration with Telegram.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#telegram)
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
//			_, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewIntegrationTelegram(ctx, "default", &gitlab.IntegrationTelegramArgs{
//				Token:                     pulumi.String("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"),
//				Room:                      pulumi.String("-1000000000000000"),
//				NotifyOnlyBrokenPipelines: pulumi.Bool(false),
//				BranchesToBeNotified:      pulumi.String("all"),
//				PushEvents:                pulumi.Bool(false),
//				IssuesEvents:              pulumi.Bool(false),
//				ConfidentialIssuesEvents:  pulumi.Bool(false),
//				MergeRequestsEvents:       pulumi.Bool(false),
//				TagPushEvents:             pulumi.Bool(false),
//				NoteEvents:                pulumi.Bool(false),
//				ConfidentialNoteEvents:    pulumi.Bool(false),
//				PipelineEvents:            pulumi.Bool(false),
//				WikiPageEvents:            pulumi.Bool(false),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_integration_telegram`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_integration_telegram.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// You can import a gitlab_integration_telegram state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/integrationTelegram:IntegrationTelegram default 1
// ```
type IntegrationTelegram struct {
	pulumi.CustomResourceState

	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
	BranchesToBeNotified pulumi.StringOutput `pulumi:"branchesToBeNotified"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolOutput `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolOutput `pulumi:"confidentialNoteEvents"`
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolOutput `pulumi:"issuesEvents"`
	// Enable notifications for merge requests events.
	MergeRequestsEvents pulumi.BoolOutput `pulumi:"mergeRequestsEvents"`
	// Enable notifications for note events.
	NoteEvents pulumi.BoolOutput `pulumi:"noteEvents"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines pulumi.BoolOutput `pulumi:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events.
	PipelineEvents pulumi.BoolOutput `pulumi:"pipelineEvents"`
	// The ID or full path of the project to integrate with Telegram.
	Project pulumi.StringOutput `pulumi:"project"`
	// Enable notifications for push events.
	PushEvents pulumi.BoolOutput `pulumi:"pushEvents"`
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
	Room pulumi.StringOutput `pulumi:"room"`
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolOutput `pulumi:"tagPushEvents"`
	// The Telegram bot token.
	Token pulumi.StringOutput `pulumi:"token"`
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolOutput `pulumi:"wikiPageEvents"`
}

// NewIntegrationTelegram registers a new resource with the given unique name, arguments, and options.
func NewIntegrationTelegram(ctx *pulumi.Context,
	name string, args *IntegrationTelegramArgs, opts ...pulumi.ResourceOption) (*IntegrationTelegram, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfidentialIssuesEvents == nil {
		return nil, errors.New("invalid value for required argument 'ConfidentialIssuesEvents'")
	}
	if args.ConfidentialNoteEvents == nil {
		return nil, errors.New("invalid value for required argument 'ConfidentialNoteEvents'")
	}
	if args.IssuesEvents == nil {
		return nil, errors.New("invalid value for required argument 'IssuesEvents'")
	}
	if args.MergeRequestsEvents == nil {
		return nil, errors.New("invalid value for required argument 'MergeRequestsEvents'")
	}
	if args.NoteEvents == nil {
		return nil, errors.New("invalid value for required argument 'NoteEvents'")
	}
	if args.PipelineEvents == nil {
		return nil, errors.New("invalid value for required argument 'PipelineEvents'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.PushEvents == nil {
		return nil, errors.New("invalid value for required argument 'PushEvents'")
	}
	if args.Room == nil {
		return nil, errors.New("invalid value for required argument 'Room'")
	}
	if args.TagPushEvents == nil {
		return nil, errors.New("invalid value for required argument 'TagPushEvents'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.WikiPageEvents == nil {
		return nil, errors.New("invalid value for required argument 'WikiPageEvents'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationTelegram
	err := ctx.RegisterResource("gitlab:index/integrationTelegram:IntegrationTelegram", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationTelegram gets an existing IntegrationTelegram resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationTelegram(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationTelegramState, opts ...pulumi.ResourceOption) (*IntegrationTelegram, error) {
	var resource IntegrationTelegram
	err := ctx.ReadResource("gitlab:index/integrationTelegram:IntegrationTelegram", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationTelegram resources.
type integrationTelegramState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Enable notifications for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Enable notifications for merge requests events.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Enable notifications for note events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// The ID or full path of the project to integrate with Telegram.
	Project *string `pulumi:"project"`
	// Enable notifications for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
	Room *string `pulumi:"room"`
	// Enable notifications for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// The Telegram bot token.
	Token *string `pulumi:"token"`
	// Enable notifications for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type IntegrationTelegramState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
	BranchesToBeNotified pulumi.StringPtrInput
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// Enable notifications for merge requests events.
	MergeRequestsEvents pulumi.BoolPtrInput
	// Enable notifications for note events.
	NoteEvents pulumi.BoolPtrInput
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// Enable notifications for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// The ID or full path of the project to integrate with Telegram.
	Project pulumi.StringPtrInput
	// Enable notifications for push events.
	PushEvents pulumi.BoolPtrInput
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
	Room pulumi.StringPtrInput
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// The Telegram bot token.
	Token pulumi.StringPtrInput
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (IntegrationTelegramState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationTelegramState)(nil)).Elem()
}

type integrationTelegramArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents bool `pulumi:"confidentialIssuesEvents"`
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents bool `pulumi:"confidentialNoteEvents"`
	// Enable notifications for issues events.
	IssuesEvents bool `pulumi:"issuesEvents"`
	// Enable notifications for merge requests events.
	MergeRequestsEvents bool `pulumi:"mergeRequestsEvents"`
	// Enable notifications for note events.
	NoteEvents bool `pulumi:"noteEvents"`
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// Enable notifications for pipeline events.
	PipelineEvents bool `pulumi:"pipelineEvents"`
	// The ID or full path of the project to integrate with Telegram.
	Project string `pulumi:"project"`
	// Enable notifications for push events.
	PushEvents bool `pulumi:"pushEvents"`
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
	Room string `pulumi:"room"`
	// Enable notifications for tag push events.
	TagPushEvents bool `pulumi:"tagPushEvents"`
	// The Telegram bot token.
	Token string `pulumi:"token"`
	// Enable notifications for wiki page events.
	WikiPageEvents bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a IntegrationTelegram resource.
type IntegrationTelegramArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
	BranchesToBeNotified pulumi.StringPtrInput
	// Enable notifications for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolInput
	// Enable notifications for confidential note events.
	ConfidentialNoteEvents pulumi.BoolInput
	// Enable notifications for issues events.
	IssuesEvents pulumi.BoolInput
	// Enable notifications for merge requests events.
	MergeRequestsEvents pulumi.BoolInput
	// Enable notifications for note events.
	NoteEvents pulumi.BoolInput
	// Send notifications for broken pipelines.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// Enable notifications for pipeline events.
	PipelineEvents pulumi.BoolInput
	// The ID or full path of the project to integrate with Telegram.
	Project pulumi.StringInput
	// Enable notifications for push events.
	PushEvents pulumi.BoolInput
	// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
	Room pulumi.StringInput
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolInput
	// The Telegram bot token.
	Token pulumi.StringInput
	// Enable notifications for wiki page events.
	WikiPageEvents pulumi.BoolInput
}

func (IntegrationTelegramArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationTelegramArgs)(nil)).Elem()
}

type IntegrationTelegramInput interface {
	pulumi.Input

	ToIntegrationTelegramOutput() IntegrationTelegramOutput
	ToIntegrationTelegramOutputWithContext(ctx context.Context) IntegrationTelegramOutput
}

func (*IntegrationTelegram) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationTelegram)(nil)).Elem()
}

func (i *IntegrationTelegram) ToIntegrationTelegramOutput() IntegrationTelegramOutput {
	return i.ToIntegrationTelegramOutputWithContext(context.Background())
}

func (i *IntegrationTelegram) ToIntegrationTelegramOutputWithContext(ctx context.Context) IntegrationTelegramOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationTelegramOutput)
}

// IntegrationTelegramArrayInput is an input type that accepts IntegrationTelegramArray and IntegrationTelegramArrayOutput values.
// You can construct a concrete instance of `IntegrationTelegramArrayInput` via:
//
//	IntegrationTelegramArray{ IntegrationTelegramArgs{...} }
type IntegrationTelegramArrayInput interface {
	pulumi.Input

	ToIntegrationTelegramArrayOutput() IntegrationTelegramArrayOutput
	ToIntegrationTelegramArrayOutputWithContext(context.Context) IntegrationTelegramArrayOutput
}

type IntegrationTelegramArray []IntegrationTelegramInput

func (IntegrationTelegramArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationTelegram)(nil)).Elem()
}

func (i IntegrationTelegramArray) ToIntegrationTelegramArrayOutput() IntegrationTelegramArrayOutput {
	return i.ToIntegrationTelegramArrayOutputWithContext(context.Background())
}

func (i IntegrationTelegramArray) ToIntegrationTelegramArrayOutputWithContext(ctx context.Context) IntegrationTelegramArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationTelegramArrayOutput)
}

// IntegrationTelegramMapInput is an input type that accepts IntegrationTelegramMap and IntegrationTelegramMapOutput values.
// You can construct a concrete instance of `IntegrationTelegramMapInput` via:
//
//	IntegrationTelegramMap{ "key": IntegrationTelegramArgs{...} }
type IntegrationTelegramMapInput interface {
	pulumi.Input

	ToIntegrationTelegramMapOutput() IntegrationTelegramMapOutput
	ToIntegrationTelegramMapOutputWithContext(context.Context) IntegrationTelegramMapOutput
}

type IntegrationTelegramMap map[string]IntegrationTelegramInput

func (IntegrationTelegramMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationTelegram)(nil)).Elem()
}

func (i IntegrationTelegramMap) ToIntegrationTelegramMapOutput() IntegrationTelegramMapOutput {
	return i.ToIntegrationTelegramMapOutputWithContext(context.Background())
}

func (i IntegrationTelegramMap) ToIntegrationTelegramMapOutputWithContext(ctx context.Context) IntegrationTelegramMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationTelegramMapOutput)
}

type IntegrationTelegramOutput struct{ *pulumi.OutputState }

func (IntegrationTelegramOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationTelegram)(nil)).Elem()
}

func (o IntegrationTelegramOutput) ToIntegrationTelegramOutput() IntegrationTelegramOutput {
	return o
}

func (o IntegrationTelegramOutput) ToIntegrationTelegramOutputWithContext(ctx context.Context) IntegrationTelegramOutput {
	return o
}

// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
func (o IntegrationTelegramOutput) BranchesToBeNotified() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.StringOutput { return v.BranchesToBeNotified }).(pulumi.StringOutput)
}

// Enable notifications for confidential issues events.
func (o IntegrationTelegramOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// Enable notifications for confidential note events.
func (o IntegrationTelegramOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// Enable notifications for issues events.
func (o IntegrationTelegramOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// Enable notifications for merge requests events.
func (o IntegrationTelegramOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// Enable notifications for note events.
func (o IntegrationTelegramOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Send notifications for broken pipelines.
func (o IntegrationTelegramOutput) NotifyOnlyBrokenPipelines() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolOutput)
}

// Enable notifications for pipeline events.
func (o IntegrationTelegramOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// The ID or full path of the project to integrate with Telegram.
func (o IntegrationTelegramOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Enable notifications for push events.
func (o IntegrationTelegramOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.PushEvents }).(pulumi.BoolOutput)
}

// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
func (o IntegrationTelegramOutput) Room() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.StringOutput { return v.Room }).(pulumi.StringOutput)
}

// Enable notifications for tag push events.
func (o IntegrationTelegramOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// The Telegram bot token.
func (o IntegrationTelegramOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Enable notifications for wiki page events.
func (o IntegrationTelegramOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationTelegram) pulumi.BoolOutput { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

type IntegrationTelegramArrayOutput struct{ *pulumi.OutputState }

func (IntegrationTelegramArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationTelegram)(nil)).Elem()
}

func (o IntegrationTelegramArrayOutput) ToIntegrationTelegramArrayOutput() IntegrationTelegramArrayOutput {
	return o
}

func (o IntegrationTelegramArrayOutput) ToIntegrationTelegramArrayOutputWithContext(ctx context.Context) IntegrationTelegramArrayOutput {
	return o
}

func (o IntegrationTelegramArrayOutput) Index(i pulumi.IntInput) IntegrationTelegramOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationTelegram {
		return vs[0].([]*IntegrationTelegram)[vs[1].(int)]
	}).(IntegrationTelegramOutput)
}

type IntegrationTelegramMapOutput struct{ *pulumi.OutputState }

func (IntegrationTelegramMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationTelegram)(nil)).Elem()
}

func (o IntegrationTelegramMapOutput) ToIntegrationTelegramMapOutput() IntegrationTelegramMapOutput {
	return o
}

func (o IntegrationTelegramMapOutput) ToIntegrationTelegramMapOutputWithContext(ctx context.Context) IntegrationTelegramMapOutput {
	return o
}

func (o IntegrationTelegramMapOutput) MapIndex(k pulumi.StringInput) IntegrationTelegramOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationTelegram {
		return vs[0].(map[string]*IntegrationTelegram)[vs[1].(string)]
	}).(IntegrationTelegramOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationTelegramInput)(nil)).Elem(), &IntegrationTelegram{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationTelegramArrayInput)(nil)).Elem(), IntegrationTelegramArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationTelegramMapInput)(nil)).Elem(), IntegrationTelegramMap{})
	pulumi.RegisterOutputType(IntegrationTelegramOutput{})
	pulumi.RegisterOutputType(IntegrationTelegramArrayOutput{})
	pulumi.RegisterOutputType(IntegrationTelegramMapOutput{})
}
