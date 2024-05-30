// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getGroups` data source allows details of multiple groups to be retrieved given some optional filter criteria.
//
// > Some attributes might not be returned depending on if you're an admin or not.
//
// > Some available options require administrator privileges.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-groups)
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
//			_, err := gitlab.GetGroups(ctx, &gitlab.GetGroupsArgs{
//				Sort:    pulumi.StringRef("desc"),
//				OrderBy: pulumi.StringRef("name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.GetGroups(ctx, &gitlab.GetGroupsArgs{
//				Search: pulumi.StringRef("GitLab"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupsResult
	err := ctx.Invoke("gitlab:index/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
	OrderBy *string `pulumi:"orderBy"`
	// Search groups by name or path.
	Search *string `pulumi:"search"`
	// Sort groups' list in asc or desc order. (Requires administrator privileges)
	Sort *string `pulumi:"sort"`
	// Limit to top level groups, excluding all subgroups.
	TopLevelOnly *bool `pulumi:"topLevelOnly"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	// The list of groups.
	Groups []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
	OrderBy *string `pulumi:"orderBy"`
	// Search groups by name or path.
	Search *string `pulumi:"search"`
	// Sort groups' list in asc or desc order. (Requires administrator privileges)
	Sort *string `pulumi:"sort"`
	// Limit to top level groups, excluding all subgroups.
	TopLevelOnly *bool `pulumi:"topLevelOnly"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			var s GetGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Search groups by name or path.
	Search pulumi.StringPtrInput `pulumi:"search"`
	// Sort groups' list in asc or desc order. (Requires administrator privileges)
	Sort pulumi.StringPtrInput `pulumi:"sort"`
	// Limit to top level groups, excluding all subgroups.
	TopLevelOnly pulumi.BoolPtrInput `pulumi:"topLevelOnly"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

// The list of groups.
func (o GetGroupsResultOutput) Groups() GetGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroup { return v.Groups }).(GetGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Order the groups' list by `id`, `name`, `path`, or `similarity`. (Requires administrator privileges)
func (o GetGroupsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

// Search groups by name or path.
func (o GetGroupsResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

// Sort groups' list in asc or desc order. (Requires administrator privileges)
func (o GetGroupsResultOutput) Sort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Sort }).(pulumi.StringPtrOutput)
}

// Limit to top level groups, excluding all subgroups.
func (o GetGroupsResultOutput) TopLevelOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *bool { return v.TopLevelOnly }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}
