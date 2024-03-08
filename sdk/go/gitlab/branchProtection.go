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

// ## Example Usage
//
// ## Import
//
// Gitlab protected branches can be imported with a key composed of `<project_id>:<branch>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/branchProtection:BranchProtection BranchProtect "12345:main"
// ```
type BranchProtection struct {
	pulumi.CustomResourceState

	// Can be set to true to allow users with push access to force push.
	AllowForcePush pulumi.BoolOutput `pulumi:"allowForcePush"`
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	AllowedToMerges BranchProtectionAllowedToMergeArrayOutput `pulumi:"allowedToMerges"`
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	AllowedToPushes BranchProtectionAllowedToPushArrayOutput `pulumi:"allowedToPushes"`
	// Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
	AllowedToUnprotects BranchProtectionAllowedToUnprotectArrayOutput `pulumi:"allowedToUnprotects"`
	// Name of the branch.
	Branch pulumi.StringOutput `pulumi:"branch"`
	// The ID of the branch protection (not the branch name).
	BranchProtectionId pulumi.IntOutput `pulumi:"branchProtectionId"`
	// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
	CodeOwnerApprovalRequired pulumi.BoolOutput `pulumi:"codeOwnerApprovalRequired"`
	// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
	MergeAccessLevel pulumi.StringOutput `pulumi:"mergeAccessLevel"`
	// The id of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
	PushAccessLevel pulumi.StringOutput `pulumi:"pushAccessLevel"`
	// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
	UnprotectAccessLevel pulumi.StringOutput `pulumi:"unprotectAccessLevel"`
}

// NewBranchProtection registers a new resource with the given unique name, arguments, and options.
func NewBranchProtection(ctx *pulumi.Context,
	name string, args *BranchProtectionArgs, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Can be set to true to allow users with push access to force push.
	AllowForcePush *bool `pulumi:"allowForcePush"`
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	AllowedToMerges []BranchProtectionAllowedToMerge `pulumi:"allowedToMerges"`
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	AllowedToPushes []BranchProtectionAllowedToPush `pulumi:"allowedToPushes"`
	// Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
	AllowedToUnprotects []BranchProtectionAllowedToUnprotect `pulumi:"allowedToUnprotects"`
	// Name of the branch.
	Branch *string `pulumi:"branch"`
	// The ID of the branch protection (not the branch name).
	BranchProtectionId *int `pulumi:"branchProtectionId"`
	// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
	CodeOwnerApprovalRequired *bool `pulumi:"codeOwnerApprovalRequired"`
	// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
	MergeAccessLevel *string `pulumi:"mergeAccessLevel"`
	// The id of the project.
	Project *string `pulumi:"project"`
	// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
	PushAccessLevel *string `pulumi:"pushAccessLevel"`
	// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
	UnprotectAccessLevel *string `pulumi:"unprotectAccessLevel"`
}

type BranchProtectionState struct {
	// Can be set to true to allow users with push access to force push.
	AllowForcePush pulumi.BoolPtrInput
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	AllowedToMerges BranchProtectionAllowedToMergeArrayInput
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	AllowedToPushes BranchProtectionAllowedToPushArrayInput
	// Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
	AllowedToUnprotects BranchProtectionAllowedToUnprotectArrayInput
	// Name of the branch.
	Branch pulumi.StringPtrInput
	// The ID of the branch protection (not the branch name).
	BranchProtectionId pulumi.IntPtrInput
	// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
	CodeOwnerApprovalRequired pulumi.BoolPtrInput
	// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
	MergeAccessLevel pulumi.StringPtrInput
	// The id of the project.
	Project pulumi.StringPtrInput
	// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
	PushAccessLevel pulumi.StringPtrInput
	// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
	UnprotectAccessLevel pulumi.StringPtrInput
}

func (BranchProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionState)(nil)).Elem()
}

type branchProtectionArgs struct {
	// Can be set to true to allow users with push access to force push.
	AllowForcePush *bool `pulumi:"allowForcePush"`
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	AllowedToMerges []BranchProtectionAllowedToMerge `pulumi:"allowedToMerges"`
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	AllowedToPushes []BranchProtectionAllowedToPush `pulumi:"allowedToPushes"`
	// Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
	AllowedToUnprotects []BranchProtectionAllowedToUnprotect `pulumi:"allowedToUnprotects"`
	// Name of the branch.
	Branch string `pulumi:"branch"`
	// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
	CodeOwnerApprovalRequired *bool `pulumi:"codeOwnerApprovalRequired"`
	// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
	MergeAccessLevel *string `pulumi:"mergeAccessLevel"`
	// The id of the project.
	Project string `pulumi:"project"`
	// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
	PushAccessLevel *string `pulumi:"pushAccessLevel"`
	// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
	UnprotectAccessLevel *string `pulumi:"unprotectAccessLevel"`
}

// The set of arguments for constructing a BranchProtection resource.
type BranchProtectionArgs struct {
	// Can be set to true to allow users with push access to force push.
	AllowForcePush pulumi.BoolPtrInput
	// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
	AllowedToMerges BranchProtectionAllowedToMergeArrayInput
	// Array of access levels and user(s)/group(s) allowed to push to protected branch.
	AllowedToPushes BranchProtectionAllowedToPushArrayInput
	// Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
	AllowedToUnprotects BranchProtectionAllowedToUnprotectArrayInput
	// Name of the branch.
	Branch pulumi.StringInput
	// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
	CodeOwnerApprovalRequired pulumi.BoolPtrInput
	// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
	MergeAccessLevel pulumi.StringPtrInput
	// The id of the project.
	Project pulumi.StringInput
	// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
	PushAccessLevel pulumi.StringPtrInput
	// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
	UnprotectAccessLevel pulumi.StringPtrInput
}

func (BranchProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionArgs)(nil)).Elem()
}

type BranchProtectionInput interface {
	pulumi.Input

	ToBranchProtectionOutput() BranchProtectionOutput
	ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput
}

func (*BranchProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtection)(nil)).Elem()
}

func (i *BranchProtection) ToBranchProtectionOutput() BranchProtectionOutput {
	return i.ToBranchProtectionOutputWithContext(context.Background())
}

func (i *BranchProtection) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionOutput)
}

// BranchProtectionArrayInput is an input type that accepts BranchProtectionArray and BranchProtectionArrayOutput values.
// You can construct a concrete instance of `BranchProtectionArrayInput` via:
//
//	BranchProtectionArray{ BranchProtectionArgs{...} }
type BranchProtectionArrayInput interface {
	pulumi.Input

	ToBranchProtectionArrayOutput() BranchProtectionArrayOutput
	ToBranchProtectionArrayOutputWithContext(context.Context) BranchProtectionArrayOutput
}

type BranchProtectionArray []BranchProtectionInput

func (BranchProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtection)(nil)).Elem()
}

func (i BranchProtectionArray) ToBranchProtectionArrayOutput() BranchProtectionArrayOutput {
	return i.ToBranchProtectionArrayOutputWithContext(context.Background())
}

func (i BranchProtectionArray) ToBranchProtectionArrayOutputWithContext(ctx context.Context) BranchProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionArrayOutput)
}

// BranchProtectionMapInput is an input type that accepts BranchProtectionMap and BranchProtectionMapOutput values.
// You can construct a concrete instance of `BranchProtectionMapInput` via:
//
//	BranchProtectionMap{ "key": BranchProtectionArgs{...} }
type BranchProtectionMapInput interface {
	pulumi.Input

	ToBranchProtectionMapOutput() BranchProtectionMapOutput
	ToBranchProtectionMapOutputWithContext(context.Context) BranchProtectionMapOutput
}

type BranchProtectionMap map[string]BranchProtectionInput

func (BranchProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtection)(nil)).Elem()
}

func (i BranchProtectionMap) ToBranchProtectionMapOutput() BranchProtectionMapOutput {
	return i.ToBranchProtectionMapOutputWithContext(context.Background())
}

func (i BranchProtectionMap) ToBranchProtectionMapOutputWithContext(ctx context.Context) BranchProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionMapOutput)
}

type BranchProtectionOutput struct{ *pulumi.OutputState }

func (BranchProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtection)(nil)).Elem()
}

func (o BranchProtectionOutput) ToBranchProtectionOutput() BranchProtectionOutput {
	return o
}

func (o BranchProtectionOutput) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return o
}

// Can be set to true to allow users with push access to force push.
func (o BranchProtectionOutput) AllowForcePush() pulumi.BoolOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolOutput { return v.AllowForcePush }).(pulumi.BoolOutput)
}

// Array of access levels and user(s)/group(s) allowed to merge to protected branch.
func (o BranchProtectionOutput) AllowedToMerges() BranchProtectionAllowedToMergeArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionAllowedToMergeArrayOutput { return v.AllowedToMerges }).(BranchProtectionAllowedToMergeArrayOutput)
}

// Array of access levels and user(s)/group(s) allowed to push to protected branch.
func (o BranchProtectionOutput) AllowedToPushes() BranchProtectionAllowedToPushArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionAllowedToPushArrayOutput { return v.AllowedToPushes }).(BranchProtectionAllowedToPushArrayOutput)
}

// Array of access levels and user(s)/group(s) allowed to unprotect push to protected branch.
func (o BranchProtectionOutput) AllowedToUnprotects() BranchProtectionAllowedToUnprotectArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionAllowedToUnprotectArrayOutput { return v.AllowedToUnprotects }).(BranchProtectionAllowedToUnprotectArrayOutput)
}

// Name of the branch.
func (o BranchProtectionOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.Branch }).(pulumi.StringOutput)
}

// The ID of the branch protection (not the branch name).
func (o BranchProtectionOutput) BranchProtectionId() pulumi.IntOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.IntOutput { return v.BranchProtectionId }).(pulumi.IntOutput)
}

// Can be set to true to require code owner approval before merging. Only available for Premium and Ultimate instances.
func (o BranchProtectionOutput) CodeOwnerApprovalRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolOutput { return v.CodeOwnerApprovalRequired }).(pulumi.BoolOutput)
}

// Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
func (o BranchProtectionOutput) MergeAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.MergeAccessLevel }).(pulumi.StringOutput)
}

// The id of the project.
func (o BranchProtectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
func (o BranchProtectionOutput) PushAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.PushAccessLevel }).(pulumi.StringOutput)
}

// Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`, `admin`.
func (o BranchProtectionOutput) UnprotectAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.UnprotectAccessLevel }).(pulumi.StringOutput)
}

type BranchProtectionArrayOutput struct{ *pulumi.OutputState }

func (BranchProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtection)(nil)).Elem()
}

func (o BranchProtectionArrayOutput) ToBranchProtectionArrayOutput() BranchProtectionArrayOutput {
	return o
}

func (o BranchProtectionArrayOutput) ToBranchProtectionArrayOutputWithContext(ctx context.Context) BranchProtectionArrayOutput {
	return o
}

func (o BranchProtectionArrayOutput) Index(i pulumi.IntInput) BranchProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchProtection {
		return vs[0].([]*BranchProtection)[vs[1].(int)]
	}).(BranchProtectionOutput)
}

type BranchProtectionMapOutput struct{ *pulumi.OutputState }

func (BranchProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtection)(nil)).Elem()
}

func (o BranchProtectionMapOutput) ToBranchProtectionMapOutput() BranchProtectionMapOutput {
	return o
}

func (o BranchProtectionMapOutput) ToBranchProtectionMapOutputWithContext(ctx context.Context) BranchProtectionMapOutput {
	return o
}

func (o BranchProtectionMapOutput) MapIndex(k pulumi.StringInput) BranchProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchProtection {
		return vs[0].(map[string]*BranchProtection)[vs[1].(string)]
	}).(BranchProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionInput)(nil)).Elem(), &BranchProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionArrayInput)(nil)).Elem(), BranchProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionMapInput)(nil)).Elem(), BranchProtectionMap{})
	pulumi.RegisterOutputType(BranchProtectionOutput{})
	pulumi.RegisterOutputType(BranchProtectionArrayOutput{})
	pulumi.RegisterOutputType(BranchProtectionMapOutput{})
}
