// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_branch_protection
//
// This resource allows you to protect a specific branch by an access level so that the user with less access level cannot Merge/Push to the branch. GitLab EE features to protect by group or user are not supported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewBranchProtection(ctx, "branchProtect", &gitlab.BranchProtectionArgs{
// 			Branch:           pulumi.String("BranchProtected"),
// 			MergeAccessLevel: pulumi.String("developer"),
// 			Project:          pulumi.String("12345"),
// 			PushAccessLevel:  pulumi.String("developer"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type BranchProtection struct {
	pulumi.CustomResourceState

	// Name of the branch.
	Branch pulumi.StringOutput `pulumi:"branch"`
	// Bool, defaults to false. Can be set to true to require code owner approval before merging.
	CodeOwnerApprovalRequired pulumi.BoolPtrOutput `pulumi:"codeOwnerApprovalRequired"`
	// One of five levels of access to the project.
	MergeAccessLevel pulumi.StringOutput `pulumi:"mergeAccessLevel"`
	// The id of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// One of five levels of access to the project.
	PushAccessLevel pulumi.StringOutput `pulumi:"pushAccessLevel"`
}

// NewBranchProtection registers a new resource with the given unique name, arguments, and options.
func NewBranchProtection(ctx *pulumi.Context,
	name string, args *BranchProtectionArgs, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	if args == nil || args.Branch == nil {
		return nil, errors.New("missing required argument 'Branch'")
	}
	if args == nil || args.MergeAccessLevel == nil {
		return nil, errors.New("missing required argument 'MergeAccessLevel'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.PushAccessLevel == nil {
		return nil, errors.New("missing required argument 'PushAccessLevel'")
	}
	if args == nil {
		args = &BranchProtectionArgs{}
	}
	var resource BranchProtection
	err := ctx.RegisterResource("gitlab:index/branchProtection:BranchProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchProtection gets an existing BranchProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchProtectionState, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	var resource BranchProtection
	err := ctx.ReadResource("gitlab:index/branchProtection:BranchProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchProtection resources.
type branchProtectionState struct {
	// Name of the branch.
	Branch *string `pulumi:"branch"`
	// Bool, defaults to false. Can be set to true to require code owner approval before merging.
	CodeOwnerApprovalRequired *bool `pulumi:"codeOwnerApprovalRequired"`
	// One of five levels of access to the project.
	MergeAccessLevel *string `pulumi:"mergeAccessLevel"`
	// The id of the project.
	Project *string `pulumi:"project"`
	// One of five levels of access to the project.
	PushAccessLevel *string `pulumi:"pushAccessLevel"`
}

type BranchProtectionState struct {
	// Name of the branch.
	Branch pulumi.StringPtrInput
	// Bool, defaults to false. Can be set to true to require code owner approval before merging.
	CodeOwnerApprovalRequired pulumi.BoolPtrInput
	// One of five levels of access to the project.
	MergeAccessLevel pulumi.StringPtrInput
	// The id of the project.
	Project pulumi.StringPtrInput
	// One of five levels of access to the project.
	PushAccessLevel pulumi.StringPtrInput
}

func (BranchProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionState)(nil)).Elem()
}

type branchProtectionArgs struct {
	// Name of the branch.
	Branch string `pulumi:"branch"`
	// Bool, defaults to false. Can be set to true to require code owner approval before merging.
	CodeOwnerApprovalRequired *bool `pulumi:"codeOwnerApprovalRequired"`
	// One of five levels of access to the project.
	MergeAccessLevel string `pulumi:"mergeAccessLevel"`
	// The id of the project.
	Project string `pulumi:"project"`
	// One of five levels of access to the project.
	PushAccessLevel string `pulumi:"pushAccessLevel"`
}

// The set of arguments for constructing a BranchProtection resource.
type BranchProtectionArgs struct {
	// Name of the branch.
	Branch pulumi.StringInput
	// Bool, defaults to false. Can be set to true to require code owner approval before merging.
	CodeOwnerApprovalRequired pulumi.BoolPtrInput
	// One of five levels of access to the project.
	MergeAccessLevel pulumi.StringInput
	// The id of the project.
	Project pulumi.StringInput
	// One of five levels of access to the project.
	PushAccessLevel pulumi.StringInput
}

func (BranchProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionArgs)(nil)).Elem()
}

type BranchProtectionInput interface {
	pulumi.Input

	ToBranchProtectionOutput() BranchProtectionOutput
	ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput
}

func (BranchProtection) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchProtection)(nil)).Elem()
}

func (i BranchProtection) ToBranchProtectionOutput() BranchProtectionOutput {
	return i.ToBranchProtectionOutputWithContext(context.Background())
}

func (i BranchProtection) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionOutput)
}

type BranchProtectionOutput struct {
	*pulumi.OutputState
}

func (BranchProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchProtectionOutput)(nil)).Elem()
}

func (o BranchProtectionOutput) ToBranchProtectionOutput() BranchProtectionOutput {
	return o
}

func (o BranchProtectionOutput) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BranchProtectionOutput{})
}
