// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Group` data source allows details of a group to be retrieved by its id or full path.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#details-of-a-group)
//
// ## Example Usage
//
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
//			_, err := gitlab.LookupGroup(ctx, &gitlab.LookupGroupArgs{
//				FullPath: pulumi.StringRef("foo/bar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("gitlab:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The full path of the group.
	FullPath *string `pulumi:"fullPath"`
	// The ID of the group.
	GroupId *int `pulumi:"groupId"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// Whether developers and maintainers can push to the applicable default branch.
	DefaultBranchProtection int `pulumi:"defaultBranchProtection"`
	// The description of the group.
	Description string `pulumi:"description"`
	// Can be set by administrators only. Additional CI/CD minutes for this group.
	ExtraSharedRunnersMinutesLimit int `pulumi:"extraSharedRunnersMinutesLimit"`
	// The full name of the group.
	FullName string `pulumi:"fullName"`
	// The full path of the group.
	FullPath string `pulumi:"fullPath"`
	// The ID of the group.
	GroupId int `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Boolean, is LFS enabled for projects in this group.
	LfsEnabled bool `pulumi:"lfsEnabled"`
	// Users cannot be added to projects in this group.
	MembershipLock bool `pulumi:"membershipLock"`
	// The name of this group.
	Name string `pulumi:"name"`
	// Integer, ID of the parent group.
	ParentId int `pulumi:"parentId"`
	// The path of the group.
	Path string `pulumi:"path"`
	// When enabled, users can not fork projects from this group to external namespaces.
	PreventForkingOutsideGroup bool `pulumi:"preventForkingOutsideGroup"`
	// Boolean, is request for access enabled to the group.
	RequestAccessEnabled bool `pulumi:"requestAccessEnabled"`
	// The group level registration token to use during runner setup.
	RunnersToken string `pulumi:"runnersToken"`
	// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
	SharedRunnersMinutesLimit int `pulumi:"sharedRunnersMinutesLimit"`
	// Visibility level of the group. Possible values are `private`, `internal`, `public`.
	VisibilityLevel string `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl string `pulumi:"webUrl"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// The full path of the group.
	FullPath pulumi.StringPtrInput `pulumi:"fullPath"`
	// The ID of the group.
	GroupId pulumi.IntPtrInput `pulumi:"groupId"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// Whether developers and maintainers can push to the applicable default branch.
func (o LookupGroupResultOutput) DefaultBranchProtection() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.DefaultBranchProtection }).(pulumi.IntOutput)
}

// The description of the group.
func (o LookupGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// Can be set by administrators only. Additional CI/CD minutes for this group.
func (o LookupGroupResultOutput) ExtraSharedRunnersMinutesLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.ExtraSharedRunnersMinutesLimit }).(pulumi.IntOutput)
}

// The full name of the group.
func (o LookupGroupResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The full path of the group.
func (o LookupGroupResultOutput) FullPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.FullPath }).(pulumi.StringOutput)
}

// The ID of the group.
func (o LookupGroupResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Boolean, is LFS enabled for projects in this group.
func (o LookupGroupResultOutput) LfsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.LfsEnabled }).(pulumi.BoolOutput)
}

// Users cannot be added to projects in this group.
func (o LookupGroupResultOutput) MembershipLock() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.MembershipLock }).(pulumi.BoolOutput)
}

// The name of this group.
func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Integer, ID of the parent group.
func (o LookupGroupResultOutput) ParentId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.ParentId }).(pulumi.IntOutput)
}

// The path of the group.
func (o LookupGroupResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Path }).(pulumi.StringOutput)
}

// When enabled, users can not fork projects from this group to external namespaces.
func (o LookupGroupResultOutput) PreventForkingOutsideGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.PreventForkingOutsideGroup }).(pulumi.BoolOutput)
}

// Boolean, is request for access enabled to the group.
func (o LookupGroupResultOutput) RequestAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.RequestAccessEnabled }).(pulumi.BoolOutput)
}

// The group level registration token to use during runner setup.
func (o LookupGroupResultOutput) RunnersToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.RunnersToken }).(pulumi.StringOutput)
}

// Can be set by administrators only. Maximum number of monthly CI/CD minutes for this group. Can be nil (default; inherit system default), 0 (unlimited), or > 0.
func (o LookupGroupResultOutput) SharedRunnersMinutesLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.SharedRunnersMinutesLimit }).(pulumi.IntOutput)
}

// Visibility level of the group. Possible values are `private`, `internal`, `public`.
func (o LookupGroupResultOutput) VisibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.VisibilityLevel }).(pulumi.StringOutput)
}

// Web URL of the group.
func (o LookupGroupResultOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.WebUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
