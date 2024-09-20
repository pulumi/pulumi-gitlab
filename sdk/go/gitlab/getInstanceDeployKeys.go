// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getInstanceDeployKeys` data source allows to retrieve a list of deploy keys for a GitLab instance.
//
// > This data source requires administration privileges.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#list-all-deploy-keys)
func GetInstanceDeployKeys(ctx *pulumi.Context, args *GetInstanceDeployKeysArgs, opts ...pulumi.InvokeOption) (*GetInstanceDeployKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceDeployKeysResult
	err := ctx.Invoke("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceDeployKeys.
type GetInstanceDeployKeysArgs struct {
	// Only return deploy keys that are public.
	Public *bool `pulumi:"public"`
}

// A collection of values returned by getInstanceDeployKeys.
type GetInstanceDeployKeysResult struct {
	// The list of all deploy keys across all projects of the GitLab instance.
	DeployKeys []GetInstanceDeployKeysDeployKey `pulumi:"deployKeys"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Only return deploy keys that are public.
	Public *bool `pulumi:"public"`
}

func GetInstanceDeployKeysOutput(ctx *pulumi.Context, args GetInstanceDeployKeysOutputArgs, opts ...pulumi.InvokeOption) GetInstanceDeployKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceDeployKeysResultOutput, error) {
			args := v.(GetInstanceDeployKeysArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetInstanceDeployKeysResult
			secret, err := ctx.InvokePackageRaw("gitlab:index/getInstanceDeployKeys:getInstanceDeployKeys", args, &rv, "", opts...)
			if err != nil {
				return GetInstanceDeployKeysResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetInstanceDeployKeysResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetInstanceDeployKeysResultOutput), nil
			}
			return output, nil
		}).(GetInstanceDeployKeysResultOutput)
}

// A collection of arguments for invoking getInstanceDeployKeys.
type GetInstanceDeployKeysOutputArgs struct {
	// Only return deploy keys that are public.
	Public pulumi.BoolPtrInput `pulumi:"public"`
}

func (GetInstanceDeployKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceDeployKeysArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceDeployKeys.
type GetInstanceDeployKeysResultOutput struct{ *pulumi.OutputState }

func (GetInstanceDeployKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceDeployKeysResult)(nil)).Elem()
}

func (o GetInstanceDeployKeysResultOutput) ToGetInstanceDeployKeysResultOutput() GetInstanceDeployKeysResultOutput {
	return o
}

func (o GetInstanceDeployKeysResultOutput) ToGetInstanceDeployKeysResultOutputWithContext(ctx context.Context) GetInstanceDeployKeysResultOutput {
	return o
}

// The list of all deploy keys across all projects of the GitLab instance.
func (o GetInstanceDeployKeysResultOutput) DeployKeys() GetInstanceDeployKeysDeployKeyArrayOutput {
	return o.ApplyT(func(v GetInstanceDeployKeysResult) []GetInstanceDeployKeysDeployKey { return v.DeployKeys }).(GetInstanceDeployKeysDeployKeyArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceDeployKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceDeployKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Only return deploy keys that are public.
func (o GetInstanceDeployKeysResultOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstanceDeployKeysResult) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceDeployKeysResultOutput{})
}
