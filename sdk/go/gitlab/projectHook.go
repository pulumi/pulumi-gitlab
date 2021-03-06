// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # gitlab\_project\_hook
//
// This resource allows you to create and manage hooks for your GitLab projects.
// For further information on hooks, consult the [gitlab
// documentation](https://docs.gitlab.com/ce/user/project/integrations/webhooks.html).
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
// 		_, err := gitlab.NewProjectHook(ctx, "example", &gitlab.ProjectHookArgs{
// 			MergeRequestsEvents: pulumi.Bool(true),
// 			Project:             pulumi.String("example/hooked"),
// 			Url:                 pulumi.String("https://example.com/hook/example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectHook struct {
	pulumi.CustomResourceState

	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrOutput `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents pulumi.BoolPtrOutput `pulumi:"confidentialNoteEvents"`
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrOutput `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrOutput `pulumi:"enableSslVerification"`
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
	// The name or id of the project to add the hook to.
	Project pulumi.StringOutput `pulumi:"project"`
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrOutput `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrOutput `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrOutput `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// The url of the hook to invoke.
	Url pulumi.StringOutput `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrOutput `pulumi:"wikiPageEvents"`
}

// NewProjectHook registers a new resource with the given unique name, arguments, and options.
func NewProjectHook(ctx *pulumi.Context,
	name string, args *ProjectHookArgs, opts ...pulumi.ResourceOption) (*ProjectHook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource ProjectHook
	err := ctx.RegisterResource("gitlab:index/projectHook:ProjectHook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectHook gets an existing ProjectHook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectHook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectHookState, opts ...pulumi.ResourceOption) (*ProjectHook, error) {
	var resource ProjectHook
	err := ctx.ReadResource("gitlab:index/projectHook:ProjectHook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectHook resources.
type projectHookState struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Invoke the hook for deployment events.
	DeploymentEvents *bool `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification *bool `pulumi:"enableSslVerification"`
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
	// The name or id of the project to add the hook to.
	Project *string `pulumi:"project"`
	// Invoke the hook for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter *string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook.
	Token *string `pulumi:"token"`
	// The url of the hook to invoke.
	Url *string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

type ProjectHookState struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrInput
	// Enable ssl verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrInput
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
	// The name or id of the project to add the hook to.
	Project pulumi.StringPtrInput
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrInput
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrInput
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// A token to present when invoking the hook.
	Token pulumi.StringPtrInput
	// The url of the hook to invoke.
	Url pulumi.StringPtrInput
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (ProjectHookState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectHookState)(nil)).Elem()
}

type projectHookArgs struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents *bool `pulumi:"confidentialIssuesEvents"`
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents *bool `pulumi:"confidentialNoteEvents"`
	// Invoke the hook for deployment events.
	DeploymentEvents *bool `pulumi:"deploymentEvents"`
	// Enable ssl verification when invoking the hook.
	EnableSslVerification *bool `pulumi:"enableSslVerification"`
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
	// The name or id of the project to add the hook to.
	Project string `pulumi:"project"`
	// Invoke the hook for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter *string `pulumi:"pushEventsBranchFilter"`
	// Invoke the hook for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// A token to present when invoking the hook.
	Token *string `pulumi:"token"`
	// The url of the hook to invoke.
	Url string `pulumi:"url"`
	// Invoke the hook for wiki page events.
	WikiPageEvents *bool `pulumi:"wikiPageEvents"`
}

// The set of arguments for constructing a ProjectHook resource.
type ProjectHookArgs struct {
	// Invoke the hook for confidential issues events.
	ConfidentialIssuesEvents pulumi.BoolPtrInput
	// Invoke the hook for confidential notes events.
	ConfidentialNoteEvents pulumi.BoolPtrInput
	// Invoke the hook for deployment events.
	DeploymentEvents pulumi.BoolPtrInput
	// Enable ssl verification when invoking the hook.
	EnableSslVerification pulumi.BoolPtrInput
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
	// The name or id of the project to add the hook to.
	Project pulumi.StringInput
	// Invoke the hook for push events.
	PushEvents pulumi.BoolPtrInput
	// Invoke the hook for push events on matching branches only.
	PushEventsBranchFilter pulumi.StringPtrInput
	// Invoke the hook for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// A token to present when invoking the hook.
	Token pulumi.StringPtrInput
	// The url of the hook to invoke.
	Url pulumi.StringInput
	// Invoke the hook for wiki page events.
	WikiPageEvents pulumi.BoolPtrInput
}

func (ProjectHookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectHookArgs)(nil)).Elem()
}

type ProjectHookInput interface {
	pulumi.Input

	ToProjectHookOutput() ProjectHookOutput
	ToProjectHookOutputWithContext(ctx context.Context) ProjectHookOutput
}

func (*ProjectHook) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectHook)(nil))
}

func (i *ProjectHook) ToProjectHookOutput() ProjectHookOutput {
	return i.ToProjectHookOutputWithContext(context.Background())
}

func (i *ProjectHook) ToProjectHookOutputWithContext(ctx context.Context) ProjectHookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookOutput)
}

func (i *ProjectHook) ToProjectHookPtrOutput() ProjectHookPtrOutput {
	return i.ToProjectHookPtrOutputWithContext(context.Background())
}

func (i *ProjectHook) ToProjectHookPtrOutputWithContext(ctx context.Context) ProjectHookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookPtrOutput)
}

type ProjectHookPtrInput interface {
	pulumi.Input

	ToProjectHookPtrOutput() ProjectHookPtrOutput
	ToProjectHookPtrOutputWithContext(ctx context.Context) ProjectHookPtrOutput
}

type projectHookPtrType ProjectHookArgs

func (*projectHookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectHook)(nil))
}

func (i *projectHookPtrType) ToProjectHookPtrOutput() ProjectHookPtrOutput {
	return i.ToProjectHookPtrOutputWithContext(context.Background())
}

func (i *projectHookPtrType) ToProjectHookPtrOutputWithContext(ctx context.Context) ProjectHookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookPtrOutput)
}

// ProjectHookArrayInput is an input type that accepts ProjectHookArray and ProjectHookArrayOutput values.
// You can construct a concrete instance of `ProjectHookArrayInput` via:
//
//          ProjectHookArray{ ProjectHookArgs{...} }
type ProjectHookArrayInput interface {
	pulumi.Input

	ToProjectHookArrayOutput() ProjectHookArrayOutput
	ToProjectHookArrayOutputWithContext(context.Context) ProjectHookArrayOutput
}

type ProjectHookArray []ProjectHookInput

func (ProjectHookArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ProjectHook)(nil))
}

func (i ProjectHookArray) ToProjectHookArrayOutput() ProjectHookArrayOutput {
	return i.ToProjectHookArrayOutputWithContext(context.Background())
}

func (i ProjectHookArray) ToProjectHookArrayOutputWithContext(ctx context.Context) ProjectHookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookArrayOutput)
}

// ProjectHookMapInput is an input type that accepts ProjectHookMap and ProjectHookMapOutput values.
// You can construct a concrete instance of `ProjectHookMapInput` via:
//
//          ProjectHookMap{ "key": ProjectHookArgs{...} }
type ProjectHookMapInput interface {
	pulumi.Input

	ToProjectHookMapOutput() ProjectHookMapOutput
	ToProjectHookMapOutputWithContext(context.Context) ProjectHookMapOutput
}

type ProjectHookMap map[string]ProjectHookInput

func (ProjectHookMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ProjectHook)(nil))
}

func (i ProjectHookMap) ToProjectHookMapOutput() ProjectHookMapOutput {
	return i.ToProjectHookMapOutputWithContext(context.Background())
}

func (i ProjectHookMap) ToProjectHookMapOutputWithContext(ctx context.Context) ProjectHookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectHookMapOutput)
}

type ProjectHookOutput struct {
	*pulumi.OutputState
}

func (ProjectHookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectHook)(nil))
}

func (o ProjectHookOutput) ToProjectHookOutput() ProjectHookOutput {
	return o
}

func (o ProjectHookOutput) ToProjectHookOutputWithContext(ctx context.Context) ProjectHookOutput {
	return o
}

func (o ProjectHookOutput) ToProjectHookPtrOutput() ProjectHookPtrOutput {
	return o.ToProjectHookPtrOutputWithContext(context.Background())
}

func (o ProjectHookOutput) ToProjectHookPtrOutputWithContext(ctx context.Context) ProjectHookPtrOutput {
	return o.ApplyT(func(v ProjectHook) *ProjectHook {
		return &v
	}).(ProjectHookPtrOutput)
}

type ProjectHookPtrOutput struct {
	*pulumi.OutputState
}

func (ProjectHookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectHook)(nil))
}

func (o ProjectHookPtrOutput) ToProjectHookPtrOutput() ProjectHookPtrOutput {
	return o
}

func (o ProjectHookPtrOutput) ToProjectHookPtrOutputWithContext(ctx context.Context) ProjectHookPtrOutput {
	return o
}

type ProjectHookArrayOutput struct{ *pulumi.OutputState }

func (ProjectHookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectHook)(nil))
}

func (o ProjectHookArrayOutput) ToProjectHookArrayOutput() ProjectHookArrayOutput {
	return o
}

func (o ProjectHookArrayOutput) ToProjectHookArrayOutputWithContext(ctx context.Context) ProjectHookArrayOutput {
	return o
}

func (o ProjectHookArrayOutput) Index(i pulumi.IntInput) ProjectHookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectHook {
		return vs[0].([]ProjectHook)[vs[1].(int)]
	}).(ProjectHookOutput)
}

type ProjectHookMapOutput struct{ *pulumi.OutputState }

func (ProjectHookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProjectHook)(nil))
}

func (o ProjectHookMapOutput) ToProjectHookMapOutput() ProjectHookMapOutput {
	return o
}

func (o ProjectHookMapOutput) ToProjectHookMapOutputWithContext(ctx context.Context) ProjectHookMapOutput {
	return o
}

func (o ProjectHookMapOutput) MapIndex(k pulumi.StringInput) ProjectHookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProjectHook {
		return vs[0].(map[string]ProjectHook)[vs[1].(string)]
	}).(ProjectHookOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectHookOutput{})
	pulumi.RegisterOutputType(ProjectHookPtrOutput{})
	pulumi.RegisterOutputType(ProjectHookArrayOutput{})
	pulumi.RegisterOutputType(ProjectHookMapOutput{})
}
