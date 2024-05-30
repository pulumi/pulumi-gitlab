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

// The `ReleaseLink` resource allows to manage the lifecycle of a release link.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/releases/links.html)
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
//			// Create a project
//			example, err := gitlab.NewProject(ctx, "example", &gitlab.ProjectArgs{
//				Name:        pulumi.String("example"),
//				Description: pulumi.String("An example project"),
//			})
//			if err != nil {
//				return err
//			}
//			// Can create release link only to a tag associated with a release
//			_, err = gitlab.NewReleaseLink(ctx, "example", &gitlab.ReleaseLinkArgs{
//				Project: example.ID(),
//				TagName: pulumi.String("tag_name_associated_with_release"),
//				Name:    pulumi.String("test"),
//				Url:     pulumi.String("https://test/"),
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
// Gitlab release link can be imported with a key composed of `<project>:<tag_name>:<link_id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/releaseLink:ReleaseLink example "12345:test:2"
// ```
type ReleaseLink struct {
	pulumi.CustomResourceState

	// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	DirectAssetUrl pulumi.StringOutput `pulumi:"directAssetUrl"`
	// External or internal link.
	External pulumi.BoolOutput `pulumi:"external"`
	// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	Filepath pulumi.StringPtrOutput `pulumi:"filepath"`
	// The ID of the link.
	LinkId pulumi.IntOutput `pulumi:"linkId"`
	// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
	LinkType pulumi.StringPtrOutput `pulumi:"linkType"`
	// The name of the link. Link names must be unique within the release.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project pulumi.StringOutput `pulumi:"project"`
	// The tag associated with the Release.
	TagName pulumi.StringOutput `pulumi:"tagName"`
	// The URL of the link. Link URLs must be unique within the release.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewReleaseLink registers a new resource with the given unique name, arguments, and options.
func NewReleaseLink(ctx *pulumi.Context,
	name string, args *ReleaseLinkArgs, opts ...pulumi.ResourceOption) (*ReleaseLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.TagName == nil {
		return nil, errors.New("invalid value for required argument 'TagName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReleaseLink
	err := ctx.RegisterResource("gitlab:index/releaseLink:ReleaseLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseLink gets an existing ReleaseLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseLinkState, opts ...pulumi.ResourceOption) (*ReleaseLink, error) {
	var resource ReleaseLink
	err := ctx.ReadResource("gitlab:index/releaseLink:ReleaseLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseLink resources.
type releaseLinkState struct {
	// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	DirectAssetUrl *string `pulumi:"directAssetUrl"`
	// External or internal link.
	External *bool `pulumi:"external"`
	// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	Filepath *string `pulumi:"filepath"`
	// The ID of the link.
	LinkId *int `pulumi:"linkId"`
	// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
	LinkType *string `pulumi:"linkType"`
	// The name of the link. Link names must be unique within the release.
	Name *string `pulumi:"name"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project *string `pulumi:"project"`
	// The tag associated with the Release.
	TagName *string `pulumi:"tagName"`
	// The URL of the link. Link URLs must be unique within the release.
	Url *string `pulumi:"url"`
}

type ReleaseLinkState struct {
	// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	DirectAssetUrl pulumi.StringPtrInput
	// External or internal link.
	External pulumi.BoolPtrInput
	// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	Filepath pulumi.StringPtrInput
	// The ID of the link.
	LinkId pulumi.IntPtrInput
	// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
	LinkType pulumi.StringPtrInput
	// The name of the link. Link names must be unique within the release.
	Name pulumi.StringPtrInput
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project pulumi.StringPtrInput
	// The tag associated with the Release.
	TagName pulumi.StringPtrInput
	// The URL of the link. Link URLs must be unique within the release.
	Url pulumi.StringPtrInput
}

func (ReleaseLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseLinkState)(nil)).Elem()
}

type releaseLinkArgs struct {
	// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	Filepath *string `pulumi:"filepath"`
	// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
	LinkType *string `pulumi:"linkType"`
	// The name of the link. Link names must be unique within the release.
	Name *string `pulumi:"name"`
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project string `pulumi:"project"`
	// The tag associated with the Release.
	TagName string `pulumi:"tagName"`
	// The URL of the link. Link URLs must be unique within the release.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ReleaseLink resource.
type ReleaseLinkArgs struct {
	// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
	Filepath pulumi.StringPtrInput
	// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
	LinkType pulumi.StringPtrInput
	// The name of the link. Link names must be unique within the release.
	Name pulumi.StringPtrInput
	// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
	Project pulumi.StringInput
	// The tag associated with the Release.
	TagName pulumi.StringInput
	// The URL of the link. Link URLs must be unique within the release.
	Url pulumi.StringInput
}

func (ReleaseLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseLinkArgs)(nil)).Elem()
}

type ReleaseLinkInput interface {
	pulumi.Input

	ToReleaseLinkOutput() ReleaseLinkOutput
	ToReleaseLinkOutputWithContext(ctx context.Context) ReleaseLinkOutput
}

func (*ReleaseLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseLink)(nil)).Elem()
}

func (i *ReleaseLink) ToReleaseLinkOutput() ReleaseLinkOutput {
	return i.ToReleaseLinkOutputWithContext(context.Background())
}

func (i *ReleaseLink) ToReleaseLinkOutputWithContext(ctx context.Context) ReleaseLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseLinkOutput)
}

// ReleaseLinkArrayInput is an input type that accepts ReleaseLinkArray and ReleaseLinkArrayOutput values.
// You can construct a concrete instance of `ReleaseLinkArrayInput` via:
//
//	ReleaseLinkArray{ ReleaseLinkArgs{...} }
type ReleaseLinkArrayInput interface {
	pulumi.Input

	ToReleaseLinkArrayOutput() ReleaseLinkArrayOutput
	ToReleaseLinkArrayOutputWithContext(context.Context) ReleaseLinkArrayOutput
}

type ReleaseLinkArray []ReleaseLinkInput

func (ReleaseLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseLink)(nil)).Elem()
}

func (i ReleaseLinkArray) ToReleaseLinkArrayOutput() ReleaseLinkArrayOutput {
	return i.ToReleaseLinkArrayOutputWithContext(context.Background())
}

func (i ReleaseLinkArray) ToReleaseLinkArrayOutputWithContext(ctx context.Context) ReleaseLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseLinkArrayOutput)
}

// ReleaseLinkMapInput is an input type that accepts ReleaseLinkMap and ReleaseLinkMapOutput values.
// You can construct a concrete instance of `ReleaseLinkMapInput` via:
//
//	ReleaseLinkMap{ "key": ReleaseLinkArgs{...} }
type ReleaseLinkMapInput interface {
	pulumi.Input

	ToReleaseLinkMapOutput() ReleaseLinkMapOutput
	ToReleaseLinkMapOutputWithContext(context.Context) ReleaseLinkMapOutput
}

type ReleaseLinkMap map[string]ReleaseLinkInput

func (ReleaseLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseLink)(nil)).Elem()
}

func (i ReleaseLinkMap) ToReleaseLinkMapOutput() ReleaseLinkMapOutput {
	return i.ToReleaseLinkMapOutputWithContext(context.Background())
}

func (i ReleaseLinkMap) ToReleaseLinkMapOutputWithContext(ctx context.Context) ReleaseLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseLinkMapOutput)
}

type ReleaseLinkOutput struct{ *pulumi.OutputState }

func (ReleaseLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseLink)(nil)).Elem()
}

func (o ReleaseLinkOutput) ToReleaseLinkOutput() ReleaseLinkOutput {
	return o
}

func (o ReleaseLinkOutput) ToReleaseLinkOutputWithContext(ctx context.Context) ReleaseLinkOutput {
	return o
}

// Full path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
func (o ReleaseLinkOutput) DirectAssetUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringOutput { return v.DirectAssetUrl }).(pulumi.StringOutput)
}

// External or internal link.
func (o ReleaseLinkOutput) External() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.BoolOutput { return v.External }).(pulumi.BoolOutput)
}

// Relative path for a [Direct Asset link](https://docs.gitlab.com/ee/user/project/releases/index.html#permanent-links-to-release-assets).
func (o ReleaseLinkOutput) Filepath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringPtrOutput { return v.Filepath }).(pulumi.StringPtrOutput)
}

// The ID of the link.
func (o ReleaseLinkOutput) LinkId() pulumi.IntOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.IntOutput { return v.LinkId }).(pulumi.IntOutput)
}

// The type of the link. Valid values are `other`, `runbook`, `image`, `package`. Defaults to other.
func (o ReleaseLinkOutput) LinkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringPtrOutput { return v.LinkType }).(pulumi.StringPtrOutput)
}

// The name of the link. Link names must be unique within the release.
func (o ReleaseLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID or [URL-encoded path of the project](https://docs.gitlab.com/ee/api/index.html#namespaced-path-encoding).
func (o ReleaseLinkOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The tag associated with the Release.
func (o ReleaseLinkOutput) TagName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringOutput { return v.TagName }).(pulumi.StringOutput)
}

// The URL of the link. Link URLs must be unique within the release.
func (o ReleaseLinkOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseLink) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ReleaseLinkArrayOutput struct{ *pulumi.OutputState }

func (ReleaseLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseLink)(nil)).Elem()
}

func (o ReleaseLinkArrayOutput) ToReleaseLinkArrayOutput() ReleaseLinkArrayOutput {
	return o
}

func (o ReleaseLinkArrayOutput) ToReleaseLinkArrayOutputWithContext(ctx context.Context) ReleaseLinkArrayOutput {
	return o
}

func (o ReleaseLinkArrayOutput) Index(i pulumi.IntInput) ReleaseLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseLink {
		return vs[0].([]*ReleaseLink)[vs[1].(int)]
	}).(ReleaseLinkOutput)
}

type ReleaseLinkMapOutput struct{ *pulumi.OutputState }

func (ReleaseLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseLink)(nil)).Elem()
}

func (o ReleaseLinkMapOutput) ToReleaseLinkMapOutput() ReleaseLinkMapOutput {
	return o
}

func (o ReleaseLinkMapOutput) ToReleaseLinkMapOutputWithContext(ctx context.Context) ReleaseLinkMapOutput {
	return o
}

func (o ReleaseLinkMapOutput) MapIndex(k pulumi.StringInput) ReleaseLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseLink {
		return vs[0].(map[string]*ReleaseLink)[vs[1].(string)]
	}).(ReleaseLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseLinkInput)(nil)).Elem(), &ReleaseLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseLinkArrayInput)(nil)).Elem(), ReleaseLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseLinkMapInput)(nil)).Elem(), ReleaseLinkMap{})
	pulumi.RegisterOutputType(ReleaseLinkOutput{})
	pulumi.RegisterOutputType(ReleaseLinkArrayOutput{})
	pulumi.RegisterOutputType(ReleaseLinkMapOutput{})
}
