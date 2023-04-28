// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getUserSshkeys` data source allows a list of SSH keys to be retrieved by either the user ID or username.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/users.html#list-ssh-keys-for-user)
func GetUserSshkeys(ctx *pulumi.Context, args *GetUserSshkeysArgs, opts ...pulumi.InvokeOption) (*GetUserSshkeysResult, error) {
	var rv GetUserSshkeysResult
	err := ctx.Invoke("gitlab:index/getUserSshkeys:getUserSshkeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserSshkeys.
type GetUserSshkeysArgs struct {
	// ID of the user to get the SSH keys for.
	UserId *int `pulumi:"userId"`
	// Username of the user to get the SSH keys for.
	Username *string `pulumi:"username"`
}

// A collection of values returned by getUserSshkeys.
type GetUserSshkeysResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The user's keys.
	Keys []GetUserSshkeysKey `pulumi:"keys"`
	// ID of the user to get the SSH keys for.
	UserId int `pulumi:"userId"`
	// Username of the user to get the SSH keys for.
	Username string `pulumi:"username"`
}

func GetUserSshkeysOutput(ctx *pulumi.Context, args GetUserSshkeysOutputArgs, opts ...pulumi.InvokeOption) GetUserSshkeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserSshkeysResult, error) {
			args := v.(GetUserSshkeysArgs)
			r, err := GetUserSshkeys(ctx, &args, opts...)
			var s GetUserSshkeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserSshkeysResultOutput)
}

// A collection of arguments for invoking getUserSshkeys.
type GetUserSshkeysOutputArgs struct {
	// ID of the user to get the SSH keys for.
	UserId pulumi.IntPtrInput `pulumi:"userId"`
	// Username of the user to get the SSH keys for.
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (GetUserSshkeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSshkeysArgs)(nil)).Elem()
}

// A collection of values returned by getUserSshkeys.
type GetUserSshkeysResultOutput struct{ *pulumi.OutputState }

func (GetUserSshkeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSshkeysResult)(nil)).Elem()
}

func (o GetUserSshkeysResultOutput) ToGetUserSshkeysResultOutput() GetUserSshkeysResultOutput {
	return o
}

func (o GetUserSshkeysResultOutput) ToGetUserSshkeysResultOutputWithContext(ctx context.Context) GetUserSshkeysResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserSshkeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshkeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// The user's keys.
func (o GetUserSshkeysResultOutput) Keys() GetUserSshkeysKeyArrayOutput {
	return o.ApplyT(func(v GetUserSshkeysResult) []GetUserSshkeysKey { return v.Keys }).(GetUserSshkeysKeyArrayOutput)
}

// ID of the user to get the SSH keys for.
func (o GetUserSshkeysResultOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserSshkeysResult) int { return v.UserId }).(pulumi.IntOutput)
}

// Username of the user to get the SSH keys for.
func (o GetUserSshkeysResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshkeysResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserSshkeysResultOutput{})
}