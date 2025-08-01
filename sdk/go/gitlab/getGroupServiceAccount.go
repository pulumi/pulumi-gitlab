// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupServiceAccount` data source retrieves information about a GitLab service account in a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/group_service_accounts/#list-service-account-users)
func LookupGroupServiceAccount(ctx *pulumi.Context, args *LookupGroupServiceAccountArgs, opts ...pulumi.InvokeOption) (*LookupGroupServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupServiceAccountResult
	err := ctx.Invoke("gitlab:index/getGroupServiceAccount:getGroupServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupServiceAccount.
type LookupGroupServiceAccountArgs struct {
	// The ID or URL-encoded path of the target group. Must be a top-level group.
	Group string `pulumi:"group"`
	// The service account id.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// A collection of values returned by getGroupServiceAccount.
type LookupGroupServiceAccountResult struct {
	// The ID or URL-encoded path of the target group. Must be a top-level group.
	Group string `pulumi:"group"`
	Id    string `pulumi:"id"`
	// The name of the user. If not specified, the default Service account user name is used.
	Name string `pulumi:"name"`
	// The service account id.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// The username of the user. If not specified, it's automatically generated.
	Username string `pulumi:"username"`
}

func LookupGroupServiceAccountOutput(ctx *pulumi.Context, args LookupGroupServiceAccountOutputArgs, opts ...pulumi.InvokeOption) LookupGroupServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupServiceAccountResultOutput, error) {
			args := v.(LookupGroupServiceAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gitlab:index/getGroupServiceAccount:getGroupServiceAccount", args, LookupGroupServiceAccountResultOutput{}, options).(LookupGroupServiceAccountResultOutput), nil
		}).(LookupGroupServiceAccountResultOutput)
}

// A collection of arguments for invoking getGroupServiceAccount.
type LookupGroupServiceAccountOutputArgs struct {
	// The ID or URL-encoded path of the target group. Must be a top-level group.
	Group pulumi.StringInput `pulumi:"group"`
	// The service account id.
	ServiceAccountId pulumi.StringInput `pulumi:"serviceAccountId"`
}

func (LookupGroupServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getGroupServiceAccount.
type LookupGroupServiceAccountResultOutput struct{ *pulumi.OutputState }

func (LookupGroupServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupServiceAccountResult)(nil)).Elem()
}

func (o LookupGroupServiceAccountResultOutput) ToLookupGroupServiceAccountResultOutput() LookupGroupServiceAccountResultOutput {
	return o
}

func (o LookupGroupServiceAccountResultOutput) ToLookupGroupServiceAccountResultOutputWithContext(ctx context.Context) LookupGroupServiceAccountResultOutput {
	return o
}

// The ID or URL-encoded path of the target group. Must be a top-level group.
func (o LookupGroupServiceAccountResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupServiceAccountResult) string { return v.Group }).(pulumi.StringOutput)
}

func (o LookupGroupServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the user. If not specified, the default Service account user name is used.
func (o LookupGroupServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service account id.
func (o LookupGroupServiceAccountResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupServiceAccountResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// The username of the user. If not specified, it's automatically generated.
func (o LookupGroupServiceAccountResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupServiceAccountResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupServiceAccountResultOutput{})
}
