// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getProjects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.
//
// > This data source supports all available filters exposed by the [client-go](https://gitlab.com/gitlab-org/api/client-go) package, which might not expose all available filters exposed by the GitLab APIs.
//
// > The owner sub-attributes are only populated if the GitLab token used has an administrator scope.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/projects/#list-all-projects)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// List projects within a group tree
//			mygroup, err := gitlab.LookupGroup(ctx, &gitlab.LookupGroupArgs{
//				FullPath: pulumi.StringRef("mygroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.GetProjects(ctx, &gitlab.GetProjectsArgs{
//				GroupId:          pulumi.IntRef(mygroup.Id),
//				OrderBy:          pulumi.StringRef("name"),
//				IncludeSubgroups: pulumi.BoolRef(true),
//				WithShared:       pulumi.BoolRef(false),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// List projects using the search syntax
//			_, err = gitlab.GetProjects(ctx, &gitlab.GetProjectsArgs{
//				Search:     pulumi.StringRef("postgresql"),
//				Visibility: pulumi.StringRef("private"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("gitlab:index/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// Limit by archived status.
	Archived *bool `pulumi:"archived"`
	// The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
	GroupId *int `pulumi:"groupId"`
	// Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
	IncludeSubgroups *bool `pulumi:"includeSubgroups"`
	// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
	MaxQueryablePages *int `pulumi:"maxQueryablePages"`
	// Limit by projects that the current user is a member of.
	Membership *bool `pulumi:"membership"`
	// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `groupId`.
	MinAccessLevel *int `pulumi:"minAccessLevel"`
	// Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
	OrderBy *string `pulumi:"orderBy"`
	// Limit by projects owned by the current user.
	Owned *bool `pulumi:"owned"`
	// The first page to begin the query on.
	Page *int `pulumi:"page"`
	// The number of results to return per page.
	PerPage *int `pulumi:"perPage"`
	// Return list of authorized projects matching the search criteria.
	Search *string `pulumi:"search"`
	// Return only the ID, URL, name, and path of each project.
	Simple *bool `pulumi:"simple"`
	// Return projects sorted in `asc` or `desc` order. Default is `desc`.
	Sort *string `pulumi:"sort"`
	// Limit by projects starred by the current user.
	Starred *bool `pulumi:"starred"`
	// Include project statistics. Cannot be used with `groupId`.
	Statistics *bool `pulumi:"statistics"`
	// Limit by projects that have all of the given topics.
	Topics []string `pulumi:"topics"`
	// Limit by visibility `public`, `internal`, or `private`.
	Visibility *string `pulumi:"visibility"`
	// Include custom attributes in response *(admins only)*.
	WithCustomAttributes *bool `pulumi:"withCustomAttributes"`
	// Limit by projects with issues feature enabled. Default is `false`.
	WithIssuesEnabled *bool `pulumi:"withIssuesEnabled"`
	// Limit by projects with merge requests feature enabled. Default is `false`.
	WithMergeRequestsEnabled *bool `pulumi:"withMergeRequestsEnabled"`
	// Limit by projects which use the given programming language. Cannot be used with `groupId`.
	WithProgrammingLanguage *string `pulumi:"withProgrammingLanguage"`
	// Include projects shared to this group. Default is `true`. Needs `groupId`.
	WithShared *bool `pulumi:"withShared"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// Limit by archived status.
	Archived *bool `pulumi:"archived"`
	// The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
	GroupId *int `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
	IncludeSubgroups *bool `pulumi:"includeSubgroups"`
	// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
	MaxQueryablePages *int `pulumi:"maxQueryablePages"`
	// Limit by projects that the current user is a member of.
	Membership *bool `pulumi:"membership"`
	// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `groupId`.
	MinAccessLevel *int `pulumi:"minAccessLevel"`
	// Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
	OrderBy *string `pulumi:"orderBy"`
	// Limit by projects owned by the current user.
	Owned *bool `pulumi:"owned"`
	// The first page to begin the query on.
	Page *int `pulumi:"page"`
	// The number of results to return per page.
	PerPage *int `pulumi:"perPage"`
	// A list containing the projects matching the supplied arguments
	Projects []GetProjectsProject `pulumi:"projects"`
	// Return list of authorized projects matching the search criteria.
	Search *string `pulumi:"search"`
	// Return only the ID, URL, name, and path of each project.
	Simple *bool `pulumi:"simple"`
	// Return projects sorted in `asc` or `desc` order. Default is `desc`.
	Sort *string `pulumi:"sort"`
	// Limit by projects starred by the current user.
	Starred *bool `pulumi:"starred"`
	// Include project statistics. Cannot be used with `groupId`.
	Statistics *bool `pulumi:"statistics"`
	// Limit by projects that have all of the given topics.
	Topics []string `pulumi:"topics"`
	// Limit by visibility `public`, `internal`, or `private`.
	Visibility *string `pulumi:"visibility"`
	// Include custom attributes in response *(admins only)*.
	WithCustomAttributes *bool `pulumi:"withCustomAttributes"`
	// Limit by projects with issues feature enabled. Default is `false`.
	WithIssuesEnabled *bool `pulumi:"withIssuesEnabled"`
	// Limit by projects with merge requests feature enabled. Default is `false`.
	WithMergeRequestsEnabled *bool `pulumi:"withMergeRequestsEnabled"`
	// Limit by projects which use the given programming language. Cannot be used with `groupId`.
	WithProgrammingLanguage *string `pulumi:"withProgrammingLanguage"`
	// Include projects shared to this group. Default is `true`. Needs `groupId`.
	WithShared *bool `pulumi:"withShared"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectsResultOutput, error) {
			args := v.(GetProjectsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gitlab:index/getProjects:getProjects", args, GetProjectsResultOutput{}, options).(GetProjectsResultOutput), nil
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// Limit by archived status.
	Archived pulumi.BoolPtrInput `pulumi:"archived"`
	// The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
	GroupId pulumi.IntPtrInput `pulumi:"groupId"`
	// Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
	IncludeSubgroups pulumi.BoolPtrInput `pulumi:"includeSubgroups"`
	// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
	MaxQueryablePages pulumi.IntPtrInput `pulumi:"maxQueryablePages"`
	// Limit by projects that the current user is a member of.
	Membership pulumi.BoolPtrInput `pulumi:"membership"`
	// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `groupId`.
	MinAccessLevel pulumi.IntPtrInput `pulumi:"minAccessLevel"`
	// Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Limit by projects owned by the current user.
	Owned pulumi.BoolPtrInput `pulumi:"owned"`
	// The first page to begin the query on.
	Page pulumi.IntPtrInput `pulumi:"page"`
	// The number of results to return per page.
	PerPage pulumi.IntPtrInput `pulumi:"perPage"`
	// Return list of authorized projects matching the search criteria.
	Search pulumi.StringPtrInput `pulumi:"search"`
	// Return only the ID, URL, name, and path of each project.
	Simple pulumi.BoolPtrInput `pulumi:"simple"`
	// Return projects sorted in `asc` or `desc` order. Default is `desc`.
	Sort pulumi.StringPtrInput `pulumi:"sort"`
	// Limit by projects starred by the current user.
	Starred pulumi.BoolPtrInput `pulumi:"starred"`
	// Include project statistics. Cannot be used with `groupId`.
	Statistics pulumi.BoolPtrInput `pulumi:"statistics"`
	// Limit by projects that have all of the given topics.
	Topics pulumi.StringArrayInput `pulumi:"topics"`
	// Limit by visibility `public`, `internal`, or `private`.
	Visibility pulumi.StringPtrInput `pulumi:"visibility"`
	// Include custom attributes in response *(admins only)*.
	WithCustomAttributes pulumi.BoolPtrInput `pulumi:"withCustomAttributes"`
	// Limit by projects with issues feature enabled. Default is `false`.
	WithIssuesEnabled pulumi.BoolPtrInput `pulumi:"withIssuesEnabled"`
	// Limit by projects with merge requests feature enabled. Default is `false`.
	WithMergeRequestsEnabled pulumi.BoolPtrInput `pulumi:"withMergeRequestsEnabled"`
	// Limit by projects which use the given programming language. Cannot be used with `groupId`.
	WithProgrammingLanguage pulumi.StringPtrInput `pulumi:"withProgrammingLanguage"`
	// Include projects shared to this group. Default is `true`. Needs `groupId`.
	WithShared pulumi.BoolPtrInput `pulumi:"withShared"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

// Limit by archived status.
func (o GetProjectsResultOutput) Archived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Archived }).(pulumi.BoolPtrOutput)
}

// The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `minAccessLevel`, `withProgrammingLanguage` or `statistics`.
func (o GetProjectsResultOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *int { return v.GroupId }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Include projects in subgroups of this group. Default is `false`. Needs `groupId`.
func (o GetProjectsResultOutput) IncludeSubgroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.IncludeSubgroups }).(pulumi.BoolPtrOutput)
}

// The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
func (o GetProjectsResultOutput) MaxQueryablePages() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *int { return v.MaxQueryablePages }).(pulumi.IntPtrOutput)
}

// Limit by projects that the current user is a member of.
func (o GetProjectsResultOutput) Membership() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Membership }).(pulumi.BoolPtrOutput)
}

// Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/api/members/) for values. Cannot be used with `groupId`.
func (o GetProjectsResultOutput) MinAccessLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *int { return v.MinAccessLevel }).(pulumi.IntPtrOutput)
}

// Return projects ordered ordered by: `id`, `name`, `path`, `createdAt`, `updatedAt`, `lastActivityAt`, `similarity`, `repositorySize`, `storageSize`, `packagesSize`, `wikiSize`. Some values or only available in certain circumstances. See [upstream docs](https://docs.gitlab.com/api/projects/#list-all-projects) for details.
func (o GetProjectsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

// Limit by projects owned by the current user.
func (o GetProjectsResultOutput) Owned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Owned }).(pulumi.BoolPtrOutput)
}

// The first page to begin the query on.
func (o GetProjectsResultOutput) Page() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *int { return v.Page }).(pulumi.IntPtrOutput)
}

// The number of results to return per page.
func (o GetProjectsResultOutput) PerPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *int { return v.PerPage }).(pulumi.IntPtrOutput)
}

// A list containing the projects matching the supplied arguments
func (o GetProjectsResultOutput) Projects() GetProjectsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsProject { return v.Projects }).(GetProjectsProjectArrayOutput)
}

// Return list of authorized projects matching the search criteria.
func (o GetProjectsResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

// Return only the ID, URL, name, and path of each project.
func (o GetProjectsResultOutput) Simple() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Simple }).(pulumi.BoolPtrOutput)
}

// Return projects sorted in `asc` or `desc` order. Default is `desc`.
func (o GetProjectsResultOutput) Sort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Sort }).(pulumi.StringPtrOutput)
}

// Limit by projects starred by the current user.
func (o GetProjectsResultOutput) Starred() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Starred }).(pulumi.BoolPtrOutput)
}

// Include project statistics. Cannot be used with `groupId`.
func (o GetProjectsResultOutput) Statistics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Statistics }).(pulumi.BoolPtrOutput)
}

// Limit by projects that have all of the given topics.
func (o GetProjectsResultOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []string { return v.Topics }).(pulumi.StringArrayOutput)
}

// Limit by visibility `public`, `internal`, or `private`.
func (o GetProjectsResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

// Include custom attributes in response *(admins only)*.
func (o GetProjectsResultOutput) WithCustomAttributes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.WithCustomAttributes }).(pulumi.BoolPtrOutput)
}

// Limit by projects with issues feature enabled. Default is `false`.
func (o GetProjectsResultOutput) WithIssuesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.WithIssuesEnabled }).(pulumi.BoolPtrOutput)
}

// Limit by projects with merge requests feature enabled. Default is `false`.
func (o GetProjectsResultOutput) WithMergeRequestsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.WithMergeRequestsEnabled }).(pulumi.BoolPtrOutput)
}

// Limit by projects which use the given programming language. Cannot be used with `groupId`.
func (o GetProjectsResultOutput) WithProgrammingLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.WithProgrammingLanguage }).(pulumi.StringPtrOutput)
}

// Include projects shared to this group. Default is `true`. Needs `groupId`.
func (o GetProjectsResultOutput) WithShared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.WithShared }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
