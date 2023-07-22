// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getInstanceVariables` data source allows to retrieve all instance-level CI/CD variables.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)
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
//			_, err := gitlab.GetInstanceVariables(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceVariables(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetInstanceVariablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceVariablesResult
	err := ctx.Invoke("gitlab:index/getInstanceVariables:getInstanceVariables", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getInstanceVariables.
type GetInstanceVariablesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of variables returned by the search
	Variables []GetInstanceVariablesVariable `pulumi:"variables"`
}
