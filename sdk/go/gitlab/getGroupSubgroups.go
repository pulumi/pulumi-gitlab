// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getGroupSubgroups` data source allows to get subgroups of a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-a-groups-subgroups)
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
//			subgroups, err := gitlab.GetGroupSubgroups(ctx, &gitlab.GetGroupSubgroupsArgs{
//				GroupId: 123456,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("subgroups", subgroups)
//			return nil
//		})
//	}
//
// ```
func GetGroupSubgroups(ctx *pulumi.Context, args *GetGroupSubgroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupSubgroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupSubgroupsResult
	err := ctx.Invoke("gitlab:index/getGroupSubgroups:getGroupSubgroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupSubgroups.
type GetGroupSubgroupsArgs struct {
	// Show all the groups you have access to.
	AllAvailable *bool `pulumi:"allAvailable"`
	// The ID of the group.
	GroupId int `pulumi:"groupId"`
	// Limit to groups where current user has at least this access level.
	MinAccessLevel *string `pulumi:"minAccessLevel"`
	// Order groups by name, path or id.
	OrderBy *string `pulumi:"orderBy"`
	// Limit to groups explicitly owned by the current user.
	Owned *bool `pulumi:"owned"`
	// Return the list of authorized groups matching the search criteria.
	Search *string `pulumi:"search"`
	// Skip the group IDs passed.
	SkipGroups []int `pulumi:"skipGroups"`
	// Order groups in asc or desc order.
	Sort *string `pulumi:"sort"`
	// Include group statistics (administrators only).
	Statistics *bool `pulumi:"statistics"`
	// Include custom attributes in response (administrators only).
	WithCustomAttributes *bool `pulumi:"withCustomAttributes"`
}

// A collection of values returned by getGroupSubgroups.
type GetGroupSubgroupsResult struct {
	// Show all the groups you have access to.
	AllAvailable bool `pulumi:"allAvailable"`
	// The ID of the group.
	GroupId int `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Limit to groups where current user has at least this access level.
	MinAccessLevel string `pulumi:"minAccessLevel"`
	// Order groups by name, path or id.
	OrderBy string `pulumi:"orderBy"`
	// Limit to groups explicitly owned by the current user.
	Owned bool `pulumi:"owned"`
	// Return the list of authorized groups matching the search criteria.
	Search string `pulumi:"search"`
	// Skip the group IDs passed.
	SkipGroups []int `pulumi:"skipGroups"`
	// Order groups in asc or desc order.
	Sort string `pulumi:"sort"`
	// Include group statistics (administrators only).
	Statistics bool `pulumi:"statistics"`
	// Subgroups of the parent group.
	Subgroups []GetGroupSubgroupsSubgroup `pulumi:"subgroups"`
	// Include custom attributes in response (administrators only).
	WithCustomAttributes bool `pulumi:"withCustomAttributes"`
}

func GetGroupSubgroupsOutput(ctx *pulumi.Context, args GetGroupSubgroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupSubgroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupSubgroupsResult, error) {
			args := v.(GetGroupSubgroupsArgs)
			r, err := GetGroupSubgroups(ctx, &args, opts...)
			var s GetGroupSubgroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupSubgroupsResultOutput)
}

// A collection of arguments for invoking getGroupSubgroups.
type GetGroupSubgroupsOutputArgs struct {
	// Show all the groups you have access to.
	AllAvailable pulumi.BoolPtrInput `pulumi:"allAvailable"`
	// The ID of the group.
	GroupId pulumi.IntInput `pulumi:"groupId"`
	// Limit to groups where current user has at least this access level.
	MinAccessLevel pulumi.StringPtrInput `pulumi:"minAccessLevel"`
	// Order groups by name, path or id.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Limit to groups explicitly owned by the current user.
	Owned pulumi.BoolPtrInput `pulumi:"owned"`
	// Return the list of authorized groups matching the search criteria.
	Search pulumi.StringPtrInput `pulumi:"search"`
	// Skip the group IDs passed.
	SkipGroups pulumi.IntArrayInput `pulumi:"skipGroups"`
	// Order groups in asc or desc order.
	Sort pulumi.StringPtrInput `pulumi:"sort"`
	// Include group statistics (administrators only).
	Statistics pulumi.BoolPtrInput `pulumi:"statistics"`
	// Include custom attributes in response (administrators only).
	WithCustomAttributes pulumi.BoolPtrInput `pulumi:"withCustomAttributes"`
}

func (GetGroupSubgroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupSubgroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroupSubgroups.
type GetGroupSubgroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupSubgroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupSubgroupsResult)(nil)).Elem()
}

func (o GetGroupSubgroupsResultOutput) ToGetGroupSubgroupsResultOutput() GetGroupSubgroupsResultOutput {
	return o
}

func (o GetGroupSubgroupsResultOutput) ToGetGroupSubgroupsResultOutputWithContext(ctx context.Context) GetGroupSubgroupsResultOutput {
	return o
}

// Show all the groups you have access to.
func (o GetGroupSubgroupsResultOutput) AllAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) bool { return v.AllAvailable }).(pulumi.BoolOutput)
}

// The ID of the group.
func (o GetGroupSubgroupsResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupSubgroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Limit to groups where current user has at least this access level.
func (o GetGroupSubgroupsResultOutput) MinAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) string { return v.MinAccessLevel }).(pulumi.StringOutput)
}

// Order groups by name, path or id.
func (o GetGroupSubgroupsResultOutput) OrderBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) string { return v.OrderBy }).(pulumi.StringOutput)
}

// Limit to groups explicitly owned by the current user.
func (o GetGroupSubgroupsResultOutput) Owned() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) bool { return v.Owned }).(pulumi.BoolOutput)
}

// Return the list of authorized groups matching the search criteria.
func (o GetGroupSubgroupsResultOutput) Search() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) string { return v.Search }).(pulumi.StringOutput)
}

// Skip the group IDs passed.
func (o GetGroupSubgroupsResultOutput) SkipGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) []int { return v.SkipGroups }).(pulumi.IntArrayOutput)
}

// Order groups in asc or desc order.
func (o GetGroupSubgroupsResultOutput) Sort() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) string { return v.Sort }).(pulumi.StringOutput)
}

// Include group statistics (administrators only).
func (o GetGroupSubgroupsResultOutput) Statistics() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) bool { return v.Statistics }).(pulumi.BoolOutput)
}

// Subgroups of the parent group.
func (o GetGroupSubgroupsResultOutput) Subgroups() GetGroupSubgroupsSubgroupArrayOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) []GetGroupSubgroupsSubgroup { return v.Subgroups }).(GetGroupSubgroupsSubgroupArrayOutput)
}

// Include custom attributes in response (administrators only).
func (o GetGroupSubgroupsResultOutput) WithCustomAttributes() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupSubgroupsResult) bool { return v.WithCustomAttributes }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupSubgroupsResultOutput{})
}
