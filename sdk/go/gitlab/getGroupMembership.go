// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupMembership` data source allows to list and filter all members of a group specified by either its id or full path.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
func LookupGroupMembership(ctx *pulumi.Context, args *LookupGroupMembershipArgs, opts ...pulumi.InvokeOption) (*LookupGroupMembershipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupMembershipResult
	err := ctx.Invoke("gitlab:index/getGroupMembership:getGroupMembership", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupMembership.
type LookupGroupMembershipArgs struct {
	// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The full path of the group.
	FullPath *string `pulumi:"fullPath"`
	// The ID of the group.
	GroupId *int `pulumi:"groupId"`
	// Return all project members including members through ancestor groups.
	Inherited *bool `pulumi:"inherited"`
}

// A collection of values returned by getGroupMembership.
type LookupGroupMembershipResult struct {
	// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel string `pulumi:"accessLevel"`
	// The full path of the group.
	FullPath string `pulumi:"fullPath"`
	// The ID of the group.
	GroupId int `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Return all project members including members through ancestor groups.
	Inherited *bool `pulumi:"inherited"`
	// The list of group members.
	Members []GetGroupMembershipMember `pulumi:"members"`
}

func LookupGroupMembershipOutput(ctx *pulumi.Context, args LookupGroupMembershipOutputArgs, opts ...pulumi.InvokeOption) LookupGroupMembershipResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupMembershipResultOutput, error) {
			args := v.(LookupGroupMembershipArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gitlab:index/getGroupMembership:getGroupMembership", args, LookupGroupMembershipResultOutput{}, options).(LookupGroupMembershipResultOutput), nil
		}).(LookupGroupMembershipResultOutput)
}

// A collection of arguments for invoking getGroupMembership.
type LookupGroupMembershipOutputArgs struct {
	// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
	AccessLevel pulumi.StringPtrInput `pulumi:"accessLevel"`
	// The full path of the group.
	FullPath pulumi.StringPtrInput `pulumi:"fullPath"`
	// The ID of the group.
	GroupId pulumi.IntPtrInput `pulumi:"groupId"`
	// Return all project members including members through ancestor groups.
	Inherited pulumi.BoolPtrInput `pulumi:"inherited"`
}

func (LookupGroupMembershipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupMembershipArgs)(nil)).Elem()
}

// A collection of values returned by getGroupMembership.
type LookupGroupMembershipResultOutput struct{ *pulumi.OutputState }

func (LookupGroupMembershipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupMembershipResult)(nil)).Elem()
}

func (o LookupGroupMembershipResultOutput) ToLookupGroupMembershipResultOutput() LookupGroupMembershipResultOutput {
	return o
}

func (o LookupGroupMembershipResultOutput) ToLookupGroupMembershipResultOutputWithContext(ctx context.Context) LookupGroupMembershipResultOutput {
	return o
}

// Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
func (o LookupGroupMembershipResultOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupMembershipResult) string { return v.AccessLevel }).(pulumi.StringOutput)
}

// The full path of the group.
func (o LookupGroupMembershipResultOutput) FullPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupMembershipResult) string { return v.FullPath }).(pulumi.StringOutput)
}

// The ID of the group.
func (o LookupGroupMembershipResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupMembershipResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupMembershipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupMembershipResult) string { return v.Id }).(pulumi.StringOutput)
}

// Return all project members including members through ancestor groups.
func (o LookupGroupMembershipResultOutput) Inherited() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGroupMembershipResult) *bool { return v.Inherited }).(pulumi.BoolPtrOutput)
}

// The list of group members.
func (o LookupGroupMembershipResultOutput) Members() GetGroupMembershipMemberArrayOutput {
	return o.ApplyT(func(v LookupGroupMembershipResult) []GetGroupMembershipMember { return v.Members }).(GetGroupMembershipMemberArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupMembershipResultOutput{})
}
