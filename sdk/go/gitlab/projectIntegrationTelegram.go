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

// The `ProjectIntegrationTelegram` resource manages the lifecycle of a project integration with Telegram.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#telegram)
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
//			_, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectIntegrationTelegram(ctx, "default", &gitlab.ProjectIntegrationTelegramArgs{
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
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_integration_telegram`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_integration_telegram.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// You can import a gitlab_project_integration_telegram state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectIntegrationTelegram:ProjectIntegrationTelegram default 1
// ```
type ProjectIntegrationTelegram struct {
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

// NewProjectIntegrationTelegram registers a new resource with the given unique name, arguments, and options.
func NewProjectIntegrationTelegram(ctx *pulumi.Context,
	name string, args *ProjectIntegrationTelegramArgs, opts ...pulumi.ResourceOption) (*ProjectIntegrationTelegram, error) {
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
	var resource ProjectIntegrationTelegram
	err := ctx.RegisterResource("gitlab:index/projectIntegrationTelegram:ProjectIntegrationTelegram", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectIntegrationTelegram gets an existing ProjectIntegrationTelegram resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectIntegrationTelegram(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectIntegrationTelegramState, opts ...pulumi.ResourceOption) (*ProjectIntegrationTelegram, error) {
	var resource ProjectIntegrationTelegram
	err := ctx.ReadResource("gitlab:index/projectIntegrationTelegram:ProjectIntegrationTelegram", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectIntegrationTelegram resources.
type projectIntegrationTelegramState struct {
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

type ProjectIntegrationTelegramState struct {
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

func (ProjectIntegrationTelegramState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIntegrationTelegramState)(nil)).Elem()
}

type projectIntegrationTelegramArgs struct {
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

// The set of arguments for constructing a ProjectIntegrationTelegram resource.
type ProjectIntegrationTelegramArgs struct {
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

func (ProjectIntegrationTelegramArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIntegrationTelegramArgs)(nil)).Elem()
}

type ProjectIntegrationTelegramInput interface {
	pulumi.Input

	ToProjectIntegrationTelegramOutput() ProjectIntegrationTelegramOutput
	ToProjectIntegrationTelegramOutputWithContext(ctx context.Context) ProjectIntegrationTelegramOutput
}

func (*ProjectIntegrationTelegram) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIntegrationTelegram)(nil)).Elem()
}

func (i *ProjectIntegrationTelegram) ToProjectIntegrationTelegramOutput() ProjectIntegrationTelegramOutput {
	return i.ToProjectIntegrationTelegramOutputWithContext(context.Background())
}

func (i *ProjectIntegrationTelegram) ToProjectIntegrationTelegramOutputWithContext(ctx context.Context) ProjectIntegrationTelegramOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIntegrationTelegramOutput)
}

// ProjectIntegrationTelegramArrayInput is an input type that accepts ProjectIntegrationTelegramArray and ProjectIntegrationTelegramArrayOutput values.
// You can construct a concrete instance of `ProjectIntegrationTelegramArrayInput` via:
//
//	ProjectIntegrationTelegramArray{ ProjectIntegrationTelegramArgs{...} }
type ProjectIntegrationTelegramArrayInput interface {
	pulumi.Input

	ToProjectIntegrationTelegramArrayOutput() ProjectIntegrationTelegramArrayOutput
	ToProjectIntegrationTelegramArrayOutputWithContext(context.Context) ProjectIntegrationTelegramArrayOutput
}

type ProjectIntegrationTelegramArray []ProjectIntegrationTelegramInput

func (ProjectIntegrationTelegramArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIntegrationTelegram)(nil)).Elem()
}

func (i ProjectIntegrationTelegramArray) ToProjectIntegrationTelegramArrayOutput() ProjectIntegrationTelegramArrayOutput {
	return i.ToProjectIntegrationTelegramArrayOutputWithContext(context.Background())
}

func (i ProjectIntegrationTelegramArray) ToProjectIntegrationTelegramArrayOutputWithContext(ctx context.Context) ProjectIntegrationTelegramArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIntegrationTelegramArrayOutput)
}

// ProjectIntegrationTelegramMapInput is an input type that accepts ProjectIntegrationTelegramMap and ProjectIntegrationTelegramMapOutput values.
// You can construct a concrete instance of `ProjectIntegrationTelegramMapInput` via:
//
//	ProjectIntegrationTelegramMap{ "key": ProjectIntegrationTelegramArgs{...} }
type ProjectIntegrationTelegramMapInput interface {
	pulumi.Input

	ToProjectIntegrationTelegramMapOutput() ProjectIntegrationTelegramMapOutput
	ToProjectIntegrationTelegramMapOutputWithContext(context.Context) ProjectIntegrationTelegramMapOutput
}

type ProjectIntegrationTelegramMap map[string]ProjectIntegrationTelegramInput

func (ProjectIntegrationTelegramMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIntegrationTelegram)(nil)).Elem()
}

func (i ProjectIntegrationTelegramMap) ToProjectIntegrationTelegramMapOutput() ProjectIntegrationTelegramMapOutput {
	return i.ToProjectIntegrationTelegramMapOutputWithContext(context.Background())
}

func (i ProjectIntegrationTelegramMap) ToProjectIntegrationTelegramMapOutputWithContext(ctx context.Context) ProjectIntegrationTelegramMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIntegrationTelegramMapOutput)
}

type ProjectIntegrationTelegramOutput struct{ *pulumi.OutputState }

func (ProjectIntegrationTelegramOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIntegrationTelegram)(nil)).Elem()
}

func (o ProjectIntegrationTelegramOutput) ToProjectIntegrationTelegramOutput() ProjectIntegrationTelegramOutput {
	return o
}

func (o ProjectIntegrationTelegramOutput) ToProjectIntegrationTelegramOutputWithContext(ctx context.Context) ProjectIntegrationTelegramOutput {
	return o
}

// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`.
func (o ProjectIntegrationTelegramOutput) BranchesToBeNotified() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.StringOutput { return v.BranchesToBeNotified }).(pulumi.StringOutput)
}

// Enable notifications for confidential issues events.
func (o ProjectIntegrationTelegramOutput) ConfidentialIssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolOutput)
}

// Enable notifications for confidential note events.
func (o ProjectIntegrationTelegramOutput) ConfidentialNoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolOutput)
}

// Enable notifications for issues events.
func (o ProjectIntegrationTelegramOutput) IssuesEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.IssuesEvents }).(pulumi.BoolOutput)
}

// Enable notifications for merge requests events.
func (o ProjectIntegrationTelegramOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// Enable notifications for note events.
func (o ProjectIntegrationTelegramOutput) NoteEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.NoteEvents }).(pulumi.BoolOutput)
}

// Send notifications for broken pipelines.
func (o ProjectIntegrationTelegramOutput) NotifyOnlyBrokenPipelines() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolOutput)
}

// Enable notifications for pipeline events.
func (o ProjectIntegrationTelegramOutput) PipelineEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.PipelineEvents }).(pulumi.BoolOutput)
}

// The ID or full path of the project to integrate with Telegram.
func (o ProjectIntegrationTelegramOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Enable notifications for push events.
func (o ProjectIntegrationTelegramOutput) PushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.PushEvents }).(pulumi.BoolOutput)
}

// Unique identifier for the target chat or the username of the target channel (in the format `@channelusername`)
func (o ProjectIntegrationTelegramOutput) Room() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.StringOutput { return v.Room }).(pulumi.StringOutput)
}

// Enable notifications for tag push events.
func (o ProjectIntegrationTelegramOutput) TagPushEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.TagPushEvents }).(pulumi.BoolOutput)
}

// The Telegram bot token.
func (o ProjectIntegrationTelegramOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Enable notifications for wiki page events.
func (o ProjectIntegrationTelegramOutput) WikiPageEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationTelegram) pulumi.BoolOutput { return v.WikiPageEvents }).(pulumi.BoolOutput)
}

type ProjectIntegrationTelegramArrayOutput struct{ *pulumi.OutputState }

func (ProjectIntegrationTelegramArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIntegrationTelegram)(nil)).Elem()
}

func (o ProjectIntegrationTelegramArrayOutput) ToProjectIntegrationTelegramArrayOutput() ProjectIntegrationTelegramArrayOutput {
	return o
}

func (o ProjectIntegrationTelegramArrayOutput) ToProjectIntegrationTelegramArrayOutputWithContext(ctx context.Context) ProjectIntegrationTelegramArrayOutput {
	return o
}

func (o ProjectIntegrationTelegramArrayOutput) Index(i pulumi.IntInput) ProjectIntegrationTelegramOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectIntegrationTelegram {
		return vs[0].([]*ProjectIntegrationTelegram)[vs[1].(int)]
	}).(ProjectIntegrationTelegramOutput)
}

type ProjectIntegrationTelegramMapOutput struct{ *pulumi.OutputState }

func (ProjectIntegrationTelegramMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIntegrationTelegram)(nil)).Elem()
}

func (o ProjectIntegrationTelegramMapOutput) ToProjectIntegrationTelegramMapOutput() ProjectIntegrationTelegramMapOutput {
	return o
}

func (o ProjectIntegrationTelegramMapOutput) ToProjectIntegrationTelegramMapOutputWithContext(ctx context.Context) ProjectIntegrationTelegramMapOutput {
	return o
}

func (o ProjectIntegrationTelegramMapOutput) MapIndex(k pulumi.StringInput) ProjectIntegrationTelegramOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectIntegrationTelegram {
		return vs[0].(map[string]*ProjectIntegrationTelegram)[vs[1].(string)]
	}).(ProjectIntegrationTelegramOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIntegrationTelegramInput)(nil)).Elem(), &ProjectIntegrationTelegram{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIntegrationTelegramArrayInput)(nil)).Elem(), ProjectIntegrationTelegramArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIntegrationTelegramMapInput)(nil)).Elem(), ProjectIntegrationTelegramMap{})
	pulumi.RegisterOutputType(ProjectIntegrationTelegramOutput{})
	pulumi.RegisterOutputType(ProjectIntegrationTelegramArrayOutput{})
	pulumi.RegisterOutputType(ProjectIntegrationTelegramMapOutput{})
}
