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

// The `ProjectTag` resource allows to manage the lifecycle of a tag in a project.
//
// **Upstream API**: [GitLab API docs](https://docs.gitlab.com/api/tags/)
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
//			// Create a project for the tag to use
//			example, err := gitlab.NewProject(ctx, "example", &gitlab.ProjectArgs{
//				Name:        pulumi.String("example"),
//				Description: pulumi.String("An example project"),
//				NamespaceId: pulumi.Any(exampleGitlabGroup.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewProjectTag(ctx, "example", &gitlab.ProjectTagArgs{
//				Name:    pulumi.String("example"),
//				Ref:     pulumi.String("main"),
//				Project: example.ID(),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_tag`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_tag.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// Gitlab project tags can be imported with a key composed of `<project_id>:<tag_name>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectTag:ProjectTag example "12345:develop"
// ```
//
// NOTE: the `ref` attribute won't be available for imported `gitlab_project_tag` resources.
type ProjectTag struct {
	pulumi.CustomResourceState

	// The commit associated with the tag.
	Commits ProjectTagCommitArrayOutput `pulumi:"commits"`
	// The message of the annotated tag.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// The name of a tag.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringOutput `pulumi:"project"`
	// Bool, true if tag has tag protection.
	Protected pulumi.BoolOutput `pulumi:"protected"`
	// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
	Ref pulumi.StringOutput `pulumi:"ref"`
	// The release associated with the tag.
	Releases ProjectTagReleaseArrayOutput `pulumi:"releases"`
	// The unique id assigned to the commit by Gitlab.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewProjectTag registers a new resource with the given unique name, arguments, and options.
func NewProjectTag(ctx *pulumi.Context,
	name string, args *ProjectTagArgs, opts ...pulumi.ResourceOption) (*ProjectTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Ref == nil {
		return nil, errors.New("invalid value for required argument 'Ref'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectTag
	err := ctx.RegisterResource("gitlab:index/projectTag:ProjectTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectTag gets an existing ProjectTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectTagState, opts ...pulumi.ResourceOption) (*ProjectTag, error) {
	var resource ProjectTag
	err := ctx.ReadResource("gitlab:index/projectTag:ProjectTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectTag resources.
type projectTagState struct {
	// The commit associated with the tag.
	Commits []ProjectTagCommit `pulumi:"commits"`
	// The message of the annotated tag.
	Message *string `pulumi:"message"`
	// The name of a tag.
	Name *string `pulumi:"name"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project *string `pulumi:"project"`
	// Bool, true if tag has tag protection.
	Protected *bool `pulumi:"protected"`
	// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
	Ref *string `pulumi:"ref"`
	// The release associated with the tag.
	Releases []ProjectTagRelease `pulumi:"releases"`
	// The unique id assigned to the commit by Gitlab.
	Target *string `pulumi:"target"`
}

type ProjectTagState struct {
	// The commit associated with the tag.
	Commits ProjectTagCommitArrayInput
	// The message of the annotated tag.
	Message pulumi.StringPtrInput
	// The name of a tag.
	Name pulumi.StringPtrInput
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringPtrInput
	// Bool, true if tag has tag protection.
	Protected pulumi.BoolPtrInput
	// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
	Ref pulumi.StringPtrInput
	// The release associated with the tag.
	Releases ProjectTagReleaseArrayInput
	// The unique id assigned to the commit by Gitlab.
	Target pulumi.StringPtrInput
}

func (ProjectTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectTagState)(nil)).Elem()
}

type projectTagArgs struct {
	// The message of the annotated tag.
	Message *string `pulumi:"message"`
	// The name of a tag.
	Name *string `pulumi:"name"`
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project string `pulumi:"project"`
	// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
	Ref string `pulumi:"ref"`
}

// The set of arguments for constructing a ProjectTag resource.
type ProjectTagArgs struct {
	// The message of the annotated tag.
	Message pulumi.StringPtrInput
	// The name of a tag.
	Name pulumi.StringPtrInput
	// The ID or URL-encoded path of the project owned by the authenticated user.
	Project pulumi.StringInput
	// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
	Ref pulumi.StringInput
}

func (ProjectTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectTagArgs)(nil)).Elem()
}

type ProjectTagInput interface {
	pulumi.Input

	ToProjectTagOutput() ProjectTagOutput
	ToProjectTagOutputWithContext(ctx context.Context) ProjectTagOutput
}

func (*ProjectTag) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectTag)(nil)).Elem()
}

func (i *ProjectTag) ToProjectTagOutput() ProjectTagOutput {
	return i.ToProjectTagOutputWithContext(context.Background())
}

func (i *ProjectTag) ToProjectTagOutputWithContext(ctx context.Context) ProjectTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectTagOutput)
}

// ProjectTagArrayInput is an input type that accepts ProjectTagArray and ProjectTagArrayOutput values.
// You can construct a concrete instance of `ProjectTagArrayInput` via:
//
//	ProjectTagArray{ ProjectTagArgs{...} }
type ProjectTagArrayInput interface {
	pulumi.Input

	ToProjectTagArrayOutput() ProjectTagArrayOutput
	ToProjectTagArrayOutputWithContext(context.Context) ProjectTagArrayOutput
}

type ProjectTagArray []ProjectTagInput

func (ProjectTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectTag)(nil)).Elem()
}

func (i ProjectTagArray) ToProjectTagArrayOutput() ProjectTagArrayOutput {
	return i.ToProjectTagArrayOutputWithContext(context.Background())
}

func (i ProjectTagArray) ToProjectTagArrayOutputWithContext(ctx context.Context) ProjectTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectTagArrayOutput)
}

// ProjectTagMapInput is an input type that accepts ProjectTagMap and ProjectTagMapOutput values.
// You can construct a concrete instance of `ProjectTagMapInput` via:
//
//	ProjectTagMap{ "key": ProjectTagArgs{...} }
type ProjectTagMapInput interface {
	pulumi.Input

	ToProjectTagMapOutput() ProjectTagMapOutput
	ToProjectTagMapOutputWithContext(context.Context) ProjectTagMapOutput
}

type ProjectTagMap map[string]ProjectTagInput

func (ProjectTagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectTag)(nil)).Elem()
}

func (i ProjectTagMap) ToProjectTagMapOutput() ProjectTagMapOutput {
	return i.ToProjectTagMapOutputWithContext(context.Background())
}

func (i ProjectTagMap) ToProjectTagMapOutputWithContext(ctx context.Context) ProjectTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectTagMapOutput)
}

type ProjectTagOutput struct{ *pulumi.OutputState }

func (ProjectTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectTag)(nil)).Elem()
}

func (o ProjectTagOutput) ToProjectTagOutput() ProjectTagOutput {
	return o
}

func (o ProjectTagOutput) ToProjectTagOutputWithContext(ctx context.Context) ProjectTagOutput {
	return o
}

// The commit associated with the tag.
func (o ProjectTagOutput) Commits() ProjectTagCommitArrayOutput {
	return o.ApplyT(func(v *ProjectTag) ProjectTagCommitArrayOutput { return v.Commits }).(ProjectTagCommitArrayOutput)
}

// The message of the annotated tag.
func (o ProjectTagOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectTag) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// The name of a tag.
func (o ProjectTagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectTag) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID or URL-encoded path of the project owned by the authenticated user.
func (o ProjectTagOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectTag) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Bool, true if tag has tag protection.
func (o ProjectTagOutput) Protected() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectTag) pulumi.BoolOutput { return v.Protected }).(pulumi.BoolOutput)
}

// Create tag using commit SHA, another tag name, or branch name. This attribute is not available for imported resources.
func (o ProjectTagOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectTag) pulumi.StringOutput { return v.Ref }).(pulumi.StringOutput)
}

// The release associated with the tag.
func (o ProjectTagOutput) Releases() ProjectTagReleaseArrayOutput {
	return o.ApplyT(func(v *ProjectTag) ProjectTagReleaseArrayOutput { return v.Releases }).(ProjectTagReleaseArrayOutput)
}

// The unique id assigned to the commit by Gitlab.
func (o ProjectTagOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectTag) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

type ProjectTagArrayOutput struct{ *pulumi.OutputState }

func (ProjectTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectTag)(nil)).Elem()
}

func (o ProjectTagArrayOutput) ToProjectTagArrayOutput() ProjectTagArrayOutput {
	return o
}

func (o ProjectTagArrayOutput) ToProjectTagArrayOutputWithContext(ctx context.Context) ProjectTagArrayOutput {
	return o
}

func (o ProjectTagArrayOutput) Index(i pulumi.IntInput) ProjectTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectTag {
		return vs[0].([]*ProjectTag)[vs[1].(int)]
	}).(ProjectTagOutput)
}

type ProjectTagMapOutput struct{ *pulumi.OutputState }

func (ProjectTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectTag)(nil)).Elem()
}

func (o ProjectTagMapOutput) ToProjectTagMapOutput() ProjectTagMapOutput {
	return o
}

func (o ProjectTagMapOutput) ToProjectTagMapOutputWithContext(ctx context.Context) ProjectTagMapOutput {
	return o
}

func (o ProjectTagMapOutput) MapIndex(k pulumi.StringInput) ProjectTagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectTag {
		return vs[0].(map[string]*ProjectTag)[vs[1].(string)]
	}).(ProjectTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectTagInput)(nil)).Elem(), &ProjectTag{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectTagArrayInput)(nil)).Elem(), ProjectTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectTagMapInput)(nil)).Elem(), ProjectTagMap{})
	pulumi.RegisterOutputType(ProjectTagOutput{})
	pulumi.RegisterOutputType(ProjectTagArrayOutput{})
	pulumi.RegisterOutputType(ProjectTagMapOutput{})
}
