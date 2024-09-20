// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Branch` data source allows details of a repository branch to be retrieved by its name and project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/branches.html#get-single-repository-branch)
func LookupBranch(ctx *pulumi.Context, args *LookupBranchArgs, opts ...pulumi.InvokeOption) (*LookupBranchResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBranchResult
	err := ctx.Invoke("gitlab:index/getBranch:getBranch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBranch.
type LookupBranchArgs struct {
	// The name of the branch.
	Name string `pulumi:"name"`
	// The full path or id of the project.
	Project string `pulumi:"project"`
}

// A collection of values returned by getBranch.
type LookupBranchResult struct {
	// Bool, true if you can push to the branch.
	CanPush bool `pulumi:"canPush"`
	// The commit associated with the branch ref.
	Commits []GetBranchCommit `pulumi:"commits"`
	// Bool, true if branch is the default branch for the project.
	Default bool `pulumi:"default"`
	// Bool, true if developer level access allows to merge branch.
	DeveloperCanMerge bool `pulumi:"developerCanMerge"`
	// Bool, true if developer level access allows git push.
	DeveloperCanPush bool `pulumi:"developerCanPush"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Bool, true if the branch has been merged into it's parent.
	Merged bool `pulumi:"merged"`
	// The name of the branch.
	Name string `pulumi:"name"`
	// The full path or id of the project.
	Project string `pulumi:"project"`
	// Bool, true if branch has branch protection.
	Protected bool `pulumi:"protected"`
	// The url of the created branch (https.)
	WebUrl string `pulumi:"webUrl"`
}

func LookupBranchOutput(ctx *pulumi.Context, args LookupBranchOutputArgs, opts ...pulumi.InvokeOption) LookupBranchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBranchResultOutput, error) {
			args := v.(LookupBranchArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupBranchResult
			secret, err := ctx.InvokePackageRaw("gitlab:index/getBranch:getBranch", args, &rv, "", opts...)
			if err != nil {
				return LookupBranchResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupBranchResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupBranchResultOutput), nil
			}
			return output, nil
		}).(LookupBranchResultOutput)
}

// A collection of arguments for invoking getBranch.
type LookupBranchOutputArgs struct {
	// The name of the branch.
	Name pulumi.StringInput `pulumi:"name"`
	// The full path or id of the project.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupBranchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBranchArgs)(nil)).Elem()
}

// A collection of values returned by getBranch.
type LookupBranchResultOutput struct{ *pulumi.OutputState }

func (LookupBranchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBranchResult)(nil)).Elem()
}

func (o LookupBranchResultOutput) ToLookupBranchResultOutput() LookupBranchResultOutput {
	return o
}

func (o LookupBranchResultOutput) ToLookupBranchResultOutputWithContext(ctx context.Context) LookupBranchResultOutput {
	return o
}

// Bool, true if you can push to the branch.
func (o LookupBranchResultOutput) CanPush() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBranchResult) bool { return v.CanPush }).(pulumi.BoolOutput)
}

// The commit associated with the branch ref.
func (o LookupBranchResultOutput) Commits() GetBranchCommitArrayOutput {
	return o.ApplyT(func(v LookupBranchResult) []GetBranchCommit { return v.Commits }).(GetBranchCommitArrayOutput)
}

// Bool, true if branch is the default branch for the project.
func (o LookupBranchResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBranchResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// Bool, true if developer level access allows to merge branch.
func (o LookupBranchResultOutput) DeveloperCanMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBranchResult) bool { return v.DeveloperCanMerge }).(pulumi.BoolOutput)
}

// Bool, true if developer level access allows git push.
func (o LookupBranchResultOutput) DeveloperCanPush() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBranchResult) bool { return v.DeveloperCanPush }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBranchResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Id }).(pulumi.StringOutput)
}

// Bool, true if the branch has been merged into it's parent.
func (o LookupBranchResultOutput) Merged() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBranchResult) bool { return v.Merged }).(pulumi.BoolOutput)
}

// The name of the branch.
func (o LookupBranchResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Name }).(pulumi.StringOutput)
}

// The full path or id of the project.
func (o LookupBranchResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Project }).(pulumi.StringOutput)
}

// Bool, true if branch has branch protection.
func (o LookupBranchResultOutput) Protected() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBranchResult) bool { return v.Protected }).(pulumi.BoolOutput)
}

// The url of the created branch (https.)
func (o LookupBranchResultOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.WebUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBranchResultOutput{})
}
