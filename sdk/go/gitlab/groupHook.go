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

// The `GroupHook` resource allows to manage the lifecycle of a group hook.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#hooks)
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
//			_, err := gitlab.NewGroupHook(ctx, "example", &gitlab.GroupHookArgs{
//				Group:               pulumi.String("example/hooked"),
//				Url:                 pulumi.String("https://example.com/hook/example"),
//				MergeRequestsEvents: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			// Setting all attributes
//			_, err = gitlab.NewGroupHook(ctx, "all_attributes", &gitlab.GroupHookArgs{
//				Group:                    pulumi.String("1"),
//				Url:                      pulumi.String("http://example.com"),
//				Token:                    pulumi.String("supersecret"),
//				EnableSslVerification:    pulumi.Bool(false),
//				PushEvents:               pulumi.Bool(true),
//				PushEventsBranchFilter:   pulumi.String("devel"),
//				IssuesEvents:             pulumi.Bool(false),
//				ConfidentialIssuesEvents: pulumi.Bool(false),
//				MergeRequestsEvents:      pulumi.Bool(true),
//				TagPushEvents:            pulumi.Bool(true),
//				NoteEvents:               pulumi.Bool(true),
//				ConfidentialNoteEvents:   pulumi.Bool(true),
//				JobEvents:                pulumi.Bool(true),
//				PipelineEvents:           pulumi.Bool(true),
//				WikiPageEvents:           pulumi.Bool(true),
//				DeploymentEvents:         pulumi.Bool(true),
//				ReleasesEvents:           pulumi.Bool(true),
//				SubgroupEvents:           pulumi.Bool(true),
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
// A GitLab Group Hook can be imported using a key composed of `<group-id>:<hook-id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/groupHook:GroupHook example "12345:1"
// ```
//
// NOTE: the `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
type GroupHook struct {
	pulumi.CustomResourceState

	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrOutput `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents pulumi.BoolPtrOutput `pulumi:"confidentialNoteEvents"`
	// Set a custom webhook template.
	CustomWebhookTemplate pulumi.StringPtrOutput `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrOutput `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrOutput `pulumi:"enableSslVerification"`
	// The ID or full path of the group.
	Group pulumi.StringOutput `pulumi:"group"`
	// The id of the group for the hook.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// The id of the group hook.
	HookId pulumi.IntOutput `pulumi:"hookId"`
	// Invoke the hook for issues events.
	IssuesEvents pulumi.BoolPtrOutput `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents pulumi.BoolPtrOutput `pulumi:"jobEvents"`
	// Invoke the hook for merge requests.
	MergeRequestsEvents pulumi.BoolPtrOutput `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for notes events.
	NoteEvents pulumi.BoolPtrOutput `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents pulumi.BoolPtrOutput `pulumi:"pipelineEvents"`
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrOutput `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrOutput `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for releases events.
	ReleasesEvents pulumi.BoolPtrOutput `pulumi:"releasesEvents"`
	// Invoke the hook for subgroup events.
	SubgroupEvents pulumi.BoolPtrOutput `pulumi:"subgroupEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrOutput `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// The url of the hook to invoke.
	Url pulumi.StringOutput `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrOutput `pulumi:"wikiPageEvents"`
}

// NewGroupHook registers a new resource with the given unique name, arguments, and options.
func NewGroupHook(ctx *pulumi.Context,
	name string, args *GroupHookArgs, opts ...pulumi.ResourceOption) (*GroupHook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupHook
	err := ctx.RegisterResource("gitlab:index/groupHook:GroupHook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupHook gets an existing GroupHook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupHook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupHookState, opts ...pulumi.ResourceOption) (*GroupHook, error) {
	var resource GroupHook
	err := ctx.ReadResource("gitlab:index/groupHook:GroupHook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupHook resources.
type groupHookState struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Set a custom webhook template.
	CustomWebhookTemplate *string `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents *bool `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification *bool `pulumi:"enableSslVerification"`
	// The ID or full path of the group.
	Group *string `pulumi:"group"`
	// The id of the group for the hook.
	GroupId *int `pulumi:"groupId"`
	// The id of the group hook.
	HookId *int `pulumi:"hookId"`
	// Invoke the hook for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents *bool `pulumi:"jobEvents"`
	// Invoke the hook for merge requests.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for notes events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// Invoke the hook for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter *string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for releases events.
	ReleasesEvents *bool `pulumi:"releasesEvents"`
	// Invoke the hook for subgroup events.
	SubgroupEvents *bool `pulumi:"subgroupEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The url of the hook to invoke.
	Url *string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type GroupHookState struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Set a custom webhook template.
	CustomWebhookTemplate pulumi.StringPtrInput
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrInput
	// Enable ssl verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrInput
	// The ID or full path of the group.
	Group pulumi.StringPtrInput
	// The id of the group for the hook.
	GroupId pulumi.IntPtrInput
	// The id of the group hook.
	HookId pulumi.IntPtrInput
	// Invoke the hook for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for job events.
	JobEvents pulumi.BoolPtrInput
	// Invoke the hook for merge requests.
	MergeRequestsEvents pulumi.BoolPtrInput
	// Invoke the hook for notes events.
	NoteEvents pulumi.BoolPtrInput
	// Invoke the hook for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrInput
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrInput
	// Invoke the hook for releases events.
	ReleasesEvents pulumi.BoolPtrInput
	// Invoke the hook for subgroup events.
	SubgroupEvents pulumi.BoolPtrInput
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The url of the hook to invoke.
	Url pulumi.StringPtrInput
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (GroupHookState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupHookState)(nil)).Elem()
}

type groupHookArgs struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Set a custom webhook template.
	CustomWebhookTemplate *string `pulumi:"customWebhookTemplate"`
	// Invoke the hook for deployment events.
	DeploymentEvents *bool `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification *bool `pulumi:"enableSslVerification"`
	// The ID or full path of the group.
	Group string `pulumi:"group"`
	// Invoke the hook for issues events.
	IssuesEvents *bool `pulumi:"issuesEvents"`
	// Invoke the hook for job events.
	JobEvents *bool `pulumi:"jobEvents"`
	// Invoke the hook for merge requests.
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// Invoke the hook for notes events.
	NoteEvents *bool `pulumi:"noteEvents"`
	// Invoke the hook for pipeline events.
	PipelineEvents *bool `pulumi:"pipelineEvents"`
	// Invoke the hook for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter *string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for releases events.
	ReleasesEvents *bool `pulumi:"releasesEvents"`
	// Invoke the hook for subgroup events.
	SubgroupEvents *bool `pulumi:"subgroupEvents"`
	// Invoke the hook for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token *string `pulumi:"token"`
	// The url of the hook to invoke.
	Url string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a GroupHook resource.
type GroupHookArgs struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Set a custom webhook template.
	CustomWebhookTemplate pulumi.StringPtrInput
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrInput
	// Enable ssl verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrInput
	// The ID or full path of the group.
	Group pulumi.StringInput
	// Invoke the hook for issues events.
	IssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for job events.
	JobEvents pulumi.BoolPtrInput
	// Invoke the hook for merge requests.
	MergeRequestsEvents pulumi.BoolPtrInput
	// Invoke the hook for notes events.
	NoteEvents pulumi.BoolPtrInput
	// Invoke the hook for pipeline events.
	PipelineEvents pulumi.BoolPtrInput
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrInput
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrInput
	// Invoke the hook for releases events.
	ReleasesEvents pulumi.BoolPtrInput
	// Invoke the hook for subgroup events.
	SubgroupEvents pulumi.BoolPtrInput
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// A token to present when invoking the hook. The token is not available for imported resources.
	Token pulumi.StringPtrInput
	// The url of the hook to invoke.
	Url pulumi.StringInput
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (GroupHookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupHookArgs)(nil)).Elem()
}

type GroupHookInput interface {
	pulumi.Input

	ToGroupHookOutput() GroupHookOutput
	ToGroupHookOutputWithContext(ctx context.Context) GroupHookOutput
}

func (*GroupHook) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupHook)(nil)).Elem()
}

func (i *GroupHook) ToGroupHookOutput() GroupHookOutput {
	return i.ToGroupHookOutputWithContext(context.Background())
}

func (i *GroupHook) ToGroupHookOutputWithContext(ctx context.Context) GroupHookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupHookOutput)
}

// GroupHookArrayInput is an input type that accepts GroupHookArray and GroupHookArrayOutput values.
// You can construct a concrete instance of `GroupHookArrayInput` via:
//
//	GroupHookArray{ GroupHookArgs{...} }
type GroupHookArrayInput interface {
	pulumi.Input

	ToGroupHookArrayOutput() GroupHookArrayOutput
	ToGroupHookArrayOutputWithContext(context.Context) GroupHookArrayOutput
}

type GroupHookArray []GroupHookInput

func (GroupHookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupHook)(nil)).Elem()
}

func (i GroupHookArray) ToGroupHookArrayOutput() GroupHookArrayOutput {
	return i.ToGroupHookArrayOutputWithContext(context.Background())
}

func (i GroupHookArray) ToGroupHookArrayOutputWithContext(ctx context.Context) GroupHookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupHookArrayOutput)
}

// GroupHookMapInput is an input type that accepts GroupHookMap and GroupHookMapOutput values.
// You can construct a concrete instance of `GroupHookMapInput` via:
//
//	GroupHookMap{ "key": GroupHookArgs{...} }
type GroupHookMapInput interface {
	pulumi.Input

	ToGroupHookMapOutput() GroupHookMapOutput
	ToGroupHookMapOutputWithContext(context.Context) GroupHookMapOutput
}

type GroupHookMap map[string]GroupHookInput

func (GroupHookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupHook)(nil)).Elem()
}

func (i GroupHookMap) ToGroupHookMapOutput() GroupHookMapOutput {
	return i.ToGroupHookMapOutputWithContext(context.Background())
}

func (i GroupHookMap) ToGroupHookMapOutputWithContext(ctx context.Context) GroupHookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupHookMapOutput)
}

type GroupHookOutput struct{ *pulumi.OutputState }

func (GroupHookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupHook)(nil)).Elem()
}

func (o GroupHookOutput) ToGroupHookOutput() GroupHookOutput {
	return o
}

func (o GroupHookOutput) ToGroupHookOutputWithContext(ctx context.Context) GroupHookOutput {
	return o
}

// Invoke the hook for confidential issues events.
func (o GroupHookOutput) ConfidentialIssuesEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.ConfidentialIssuesEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for confidential notes events.
func (o GroupHookOutput) ConfidentialNoteEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.ConfidentialNoteEvents }).(pulumi.BoolPtrOutput)
}

// Set a custom webhook template.
func (o GroupHookOutput) CustomWebhookTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.StringPtrOutput { return v.CustomWebhookTemplate }).(pulumi.StringPtrOutput)
}

// Invoke the hook for deployment events.
func (o GroupHookOutput) DeploymentEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.DeploymentEvents }).(pulumi.BoolPtrOutput)
}

// Enable ssl verification when invoking the hook.
func (o GroupHookOutput) EnableSslVerification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.EnableSslVerification }).(pulumi.BoolPtrOutput)
}

// The ID or full path of the group.
func (o GroupHookOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The id of the group for the hook.
func (o GroupHookOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

// The id of the group hook.
func (o GroupHookOutput) HookId() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.IntOutput { return v.HookId }).(pulumi.IntOutput)
}

// Invoke the hook for issues events.
func (o GroupHookOutput) IssuesEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.IssuesEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for job events.
func (o GroupHookOutput) JobEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.JobEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for merge requests.
func (o GroupHookOutput) MergeRequestsEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.MergeRequestsEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for notes events.
func (o GroupHookOutput) NoteEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.NoteEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for pipeline events.
func (o GroupHookOutput) PipelineEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.PipelineEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for push events.
func (o GroupHookOutput) PushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.PushEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for push events on matching branches only.
func (o GroupHookOutput) PushEventsBranchFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.StringPtrOutput { return v.PushEventsBranchFilter }).(pulumi.StringPtrOutput)
}

// Invoke the hook for releases events.
func (o GroupHookOutput) ReleasesEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.ReleasesEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for subgroup events.
func (o GroupHookOutput) SubgroupEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.SubgroupEvents }).(pulumi.BoolPtrOutput)
}

// Invoke the hook for tag push events.
func (o GroupHookOutput) TagPushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.TagPushEvents }).(pulumi.BoolPtrOutput)
}

// A token to present when invoking the hook. The token is not available for imported resources.
func (o GroupHookOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

// The url of the hook to invoke.
func (o GroupHookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Invoke the hook for wiki page events.
func (o GroupHookOutput) WikiPageEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupHook) pulumi.BoolPtrOutput { return v.WikiPageEvents }).(pulumi.BoolPtrOutput)
}

type GroupHookArrayOutput struct{ *pulumi.OutputState }

func (GroupHookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupHook)(nil)).Elem()
}

func (o GroupHookArrayOutput) ToGroupHookArrayOutput() GroupHookArrayOutput {
	return o
}

func (o GroupHookArrayOutput) ToGroupHookArrayOutputWithContext(ctx context.Context) GroupHookArrayOutput {
	return o
}

func (o GroupHookArrayOutput) Index(i pulumi.IntInput) GroupHookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupHook {
		return vs[0].([]*GroupHook)[vs[1].(int)]
	}).(GroupHookOutput)
}

type GroupHookMapOutput struct{ *pulumi.OutputState }

func (GroupHookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupHook)(nil)).Elem()
}

func (o GroupHookMapOutput) ToGroupHookMapOutput() GroupHookMapOutput {
	return o
}

func (o GroupHookMapOutput) ToGroupHookMapOutputWithContext(ctx context.Context) GroupHookMapOutput {
	return o
}

func (o GroupHookMapOutput) MapIndex(k pulumi.StringInput) GroupHookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupHook {
		return vs[0].(map[string]*GroupHook)[vs[1].(string)]
	}).(GroupHookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupHookInput)(nil)).Elem(), &GroupHook{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupHookArrayInput)(nil)).Elem(), GroupHookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupHookMapInput)(nil)).Elem(), GroupHookMap{})
	pulumi.RegisterOutputType(GroupHookOutput{})
	pulumi.RegisterOutputType(GroupHookArrayOutput{})
	pulumi.RegisterOutputType(GroupHookMapOutput{})
}
