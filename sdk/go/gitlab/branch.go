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

// The `Branch` resource allows to manage the lifecycle of a repository branch.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/branches.html)
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
//			// Create a project for the branch to use
//			example, err := gitlab.NewProject(ctx, "example", &gitlab.ProjectArgs{
//				Name:        pulumi.String("example"),
//				Description: pulumi.String("An example project"),
//				NamespaceId: pulumi.Any(exampleGitlabGroup.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewBranch(ctx, "example", &gitlab.BranchArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Gitlab branches can be imported with a key composed of `<project_id>:<branch_name>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/branch:Branch example "12345:develop"
// ```
type Branch struct {
	pulumi.CustomResourceState

	// Bool, true if you can push to the branch.
	CanPush pulumi.BoolOutput `pulumi:"canPush"`
	// The commit associated with the branch ref.
	Commits BranchCommitArrayOutput `pulumi:"commits"`
	// Bool, true if branch is the default branch for the project.
	Default pulumi.BoolOutput `pulumi:"default"`
	// Bool, true if developer level access allows to merge branch.
	DeveloperCanMerge pulumi.BoolOutput `pulumi:"developerCanMerge"`
	// Bool, true if developer level access allows git push.
	DeveloperCanPush pulumi.BoolOutput `pulumi:"developerCanPush"`
	// Bool, true if the branch has been merged into it's parent.
	Merged pulumi.BoolOutput `pulumi:"merged"`
	// The name for this branch.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID or full path of the project which the branch is created against.
	Project pulumi.StringOutput `pulumi:"project"`
	// Bool, true if branch has branch protection.
	Protected pulumi.BoolOutput `pulumi:"protected"`
	// The ref which the branch is created from.
	Ref pulumi.StringOutput `pulumi:"ref"`
	// The url of the created branch (https).
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
}

// NewBranch registers a new resource with the given unique name, arguments, and options.
func NewBranch(ctx *pulumi.Context,
	name string, args *BranchArgs, opts ...pulumi.ResourceOption) (*Branch, error) {
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
	var resource Branch
	err := ctx.RegisterResource("gitlab:index/branch:Branch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranch gets an existing Branch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchState, opts ...pulumi.ResourceOption) (*Branch, error) {
	var resource Branch
	err := ctx.ReadResource("gitlab:index/branch:Branch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Branch resources.
type branchState struct {
	// Bool, true if you can push to the branch.
	CanPush *bool `pulumi:"canPush"`
	// The commit associated with the branch ref.
	Commits []BranchCommit `pulumi:"commits"`
	// Bool, true if branch is the default branch for the project.
	Default *bool `pulumi:"default"`
	// Bool, true if developer level access allows to merge branch.
	DeveloperCanMerge *bool `pulumi:"developerCanMerge"`
	// Bool, true if developer level access allows git push.
	DeveloperCanPush *bool `pulumi:"developerCanPush"`
	// Bool, true if the branch has been merged into it's parent.
	Merged *bool `pulumi:"merged"`
	// The name for this branch.
	Name *string `pulumi:"name"`
	// The ID or full path of the project which the branch is created against.
	Project *string `pulumi:"project"`
	// Bool, true if branch has branch protection.
	Protected *bool `pulumi:"protected"`
	// The ref which the branch is created from.
	Ref *string `pulumi:"ref"`
	// The url of the created branch (https).
	WebUrl *string `pulumi:"webUrl"`
}

type BranchState struct {
	// Bool, true if you can push to the branch.
	CanPush pulumi.BoolPtrInput
	// The commit associated with the branch ref.
	Commits BranchCommitArrayInput
	// Bool, true if branch is the default branch for the project.
	Default pulumi.BoolPtrInput
	// Bool, true if developer level access allows to merge branch.
	DeveloperCanMerge pulumi.BoolPtrInput
	// Bool, true if developer level access allows git push.
	DeveloperCanPush pulumi.BoolPtrInput
	// Bool, true if the branch has been merged into it's parent.
	Merged pulumi.BoolPtrInput
	// The name for this branch.
	Name pulumi.StringPtrInput
	// The ID or full path of the project which the branch is created against.
	Project pulumi.StringPtrInput
	// Bool, true if branch has branch protection.
	Protected pulumi.BoolPtrInput
	// The ref which the branch is created from.
	Ref pulumi.StringPtrInput
	// The url of the created branch (https).
	WebUrl pulumi.StringPtrInput
}

func (BranchState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchState)(nil)).Elem()
}

type branchArgs struct {
	// The name for this branch.
	Name *string `pulumi:"name"`
	// The ID or full path of the project which the branch is created against.
	Project string `pulumi:"project"`
	// The ref which the branch is created from.
	Ref string `pulumi:"ref"`
}

// The set of arguments for constructing a Branch resource.
type BranchArgs struct {
	// The name for this branch.
	Name pulumi.StringPtrInput
	// The ID or full path of the project which the branch is created against.
	Project pulumi.StringInput
	// The ref which the branch is created from.
	Ref pulumi.StringInput
}

func (BranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchArgs)(nil)).Elem()
}

type BranchInput interface {
	pulumi.Input

	ToBranchOutput() BranchOutput
	ToBranchOutputWithContext(ctx context.Context) BranchOutput
}

func (*Branch) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil)).Elem()
}

func (i *Branch) ToBranchOutput() BranchOutput {
	return i.ToBranchOutputWithContext(context.Background())
}

func (i *Branch) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchOutput)
}

// BranchArrayInput is an input type that accepts BranchArray and BranchArrayOutput values.
// You can construct a concrete instance of `BranchArrayInput` via:
//
//	BranchArray{ BranchArgs{...} }
type BranchArrayInput interface {
	pulumi.Input

	ToBranchArrayOutput() BranchArrayOutput
	ToBranchArrayOutputWithContext(context.Context) BranchArrayOutput
}

type BranchArray []BranchInput

func (BranchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Branch)(nil)).Elem()
}

func (i BranchArray) ToBranchArrayOutput() BranchArrayOutput {
	return i.ToBranchArrayOutputWithContext(context.Background())
}

func (i BranchArray) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchArrayOutput)
}

// BranchMapInput is an input type that accepts BranchMap and BranchMapOutput values.
// You can construct a concrete instance of `BranchMapInput` via:
//
//	BranchMap{ "key": BranchArgs{...} }
type BranchMapInput interface {
	pulumi.Input

	ToBranchMapOutput() BranchMapOutput
	ToBranchMapOutputWithContext(context.Context) BranchMapOutput
}

type BranchMap map[string]BranchInput

func (BranchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Branch)(nil)).Elem()
}

func (i BranchMap) ToBranchMapOutput() BranchMapOutput {
	return i.ToBranchMapOutputWithContext(context.Background())
}

func (i BranchMap) ToBranchMapOutputWithContext(ctx context.Context) BranchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchMapOutput)
}

type BranchOutput struct{ *pulumi.OutputState }

func (BranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil)).Elem()
}

func (o BranchOutput) ToBranchOutput() BranchOutput {
	return o
}

func (o BranchOutput) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return o
}

// Bool, true if you can push to the branch.
func (o BranchOutput) CanPush() pulumi.BoolOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolOutput { return v.CanPush }).(pulumi.BoolOutput)
}

// The commit associated with the branch ref.
func (o BranchOutput) Commits() BranchCommitArrayOutput {
	return o.ApplyT(func(v *Branch) BranchCommitArrayOutput { return v.Commits }).(BranchCommitArrayOutput)
}

// Bool, true if branch is the default branch for the project.
func (o BranchOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

// Bool, true if developer level access allows to merge branch.
func (o BranchOutput) DeveloperCanMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolOutput { return v.DeveloperCanMerge }).(pulumi.BoolOutput)
}

// Bool, true if developer level access allows git push.
func (o BranchOutput) DeveloperCanPush() pulumi.BoolOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolOutput { return v.DeveloperCanPush }).(pulumi.BoolOutput)
}

// Bool, true if the branch has been merged into it's parent.
func (o BranchOutput) Merged() pulumi.BoolOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolOutput { return v.Merged }).(pulumi.BoolOutput)
}

// The name for this branch.
func (o BranchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID or full path of the project which the branch is created against.
func (o BranchOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Bool, true if branch has branch protection.
func (o BranchOutput) Protected() pulumi.BoolOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolOutput { return v.Protected }).(pulumi.BoolOutput)
}

// The ref which the branch is created from.
func (o BranchOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.Ref }).(pulumi.StringOutput)
}

// The url of the created branch (https).
func (o BranchOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.WebUrl }).(pulumi.StringOutput)
}

type BranchArrayOutput struct{ *pulumi.OutputState }

func (BranchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Branch)(nil)).Elem()
}

func (o BranchArrayOutput) ToBranchArrayOutput() BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) Index(i pulumi.IntInput) BranchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Branch {
		return vs[0].([]*Branch)[vs[1].(int)]
	}).(BranchOutput)
}

type BranchMapOutput struct{ *pulumi.OutputState }

func (BranchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Branch)(nil)).Elem()
}

func (o BranchMapOutput) ToBranchMapOutput() BranchMapOutput {
	return o
}

func (o BranchMapOutput) ToBranchMapOutputWithContext(ctx context.Context) BranchMapOutput {
	return o
}

func (o BranchMapOutput) MapIndex(k pulumi.StringInput) BranchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Branch {
		return vs[0].(map[string]*Branch)[vs[1].(string)]
	}).(BranchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchInput)(nil)).Elem(), &Branch{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchArrayInput)(nil)).Elem(), BranchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchMapInput)(nil)).Elem(), BranchMap{})
	pulumi.RegisterOutputType(BranchOutput{})
	pulumi.RegisterOutputType(BranchArrayOutput{})
	pulumi.RegisterOutputType(BranchMapOutput{})
}
