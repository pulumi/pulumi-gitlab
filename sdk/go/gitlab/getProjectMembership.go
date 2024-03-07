// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectMembership` data source allows to list and filter all members of a project specified by either its id or full path.
//
// > **Note** exactly one of projectId or fullPath must be provided.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/members.html#list-all-members-of-a-group-or-project)
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
//			_, err := gitlab.LookupProjectMembership(ctx, &gitlab.LookupProjectMembershipArgs{
//				Inherited: pulumi.BoolRef(true),
//				ProjectId: pulumi.IntRef(123),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupProjectMembership(ctx *pulumi.Context, args *LookupProjectMembershipArgs, opts ...pulumi.InvokeOption) (*LookupProjectMembershipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectMembershipResult
	err := ctx.Invoke("gitlab:index/getProjectMembership:getProjectMembership", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectMembership.
type LookupProjectMembershipArgs struct {
	// The full path of the project.
	FullPath *string `pulumi:"fullPath"`
	// Return all project members including members through ancestor groups
	Inherited *bool `pulumi:"inherited"`
	// The ID of the project.
	ProjectId *int `pulumi:"projectId"`
	// A query string to search for members
	Query *string `pulumi:"query"`
}

// A collection of values returned by getProjectMembership.
type LookupProjectMembershipResult struct {
	// The full path of the project.
	FullPath string `pulumi:"fullPath"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Return all project members including members through ancestor groups
	Inherited *bool `pulumi:"inherited"`
	// The list of project members.
	Members []GetProjectMembershipMember `pulumi:"members"`
	// The ID of the project.
	ProjectId int `pulumi:"projectId"`
	// A query string to search for members
	Query *string `pulumi:"query"`
}

func LookupProjectMembershipOutput(ctx *pulumi.Context, args LookupProjectMembershipOutputArgs, opts ...pulumi.InvokeOption) LookupProjectMembershipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectMembershipResult, error) {
			args := v.(LookupProjectMembershipArgs)
			r, err := LookupProjectMembership(ctx, &args, opts...)
			var s LookupProjectMembershipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectMembershipResultOutput)
}

// A collection of arguments for invoking getProjectMembership.
type LookupProjectMembershipOutputArgs struct {
	// The full path of the project.
	FullPath pulumi.StringPtrInput `pulumi:"fullPath"`
	// Return all project members including members through ancestor groups
	Inherited pulumi.BoolPtrInput `pulumi:"inherited"`
	// The ID of the project.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// A query string to search for members
	Query pulumi.StringPtrInput `pulumi:"query"`
}

func (LookupProjectMembershipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectMembershipArgs)(nil)).Elem()
}

// A collection of values returned by getProjectMembership.
type LookupProjectMembershipResultOutput struct{ *pulumi.OutputState }

func (LookupProjectMembershipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectMembershipResult)(nil)).Elem()
}

func (o LookupProjectMembershipResultOutput) ToLookupProjectMembershipResultOutput() LookupProjectMembershipResultOutput {
	return o
}

func (o LookupProjectMembershipResultOutput) ToLookupProjectMembershipResultOutputWithContext(ctx context.Context) LookupProjectMembershipResultOutput {
	return o
}

// The full path of the project.
func (o LookupProjectMembershipResultOutput) FullPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMembershipResult) string { return v.FullPath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectMembershipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectMembershipResult) string { return v.Id }).(pulumi.StringOutput)
}

// Return all project members including members through ancestor groups
func (o LookupProjectMembershipResultOutput) Inherited() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProjectMembershipResult) *bool { return v.Inherited }).(pulumi.BoolPtrOutput)
}

// The list of project members.
func (o LookupProjectMembershipResultOutput) Members() GetProjectMembershipMemberArrayOutput {
	return o.ApplyT(func(v LookupProjectMembershipResult) []GetProjectMembershipMember { return v.Members }).(GetProjectMembershipMemberArrayOutput)
}

// The ID of the project.
func (o LookupProjectMembershipResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectMembershipResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// A query string to search for members
func (o LookupProjectMembershipResultOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectMembershipResult) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectMembershipResultOutput{})
}
