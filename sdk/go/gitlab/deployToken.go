// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_deploy\_token
//
// This resource allows you to create and manage deploy token for your GitLab projects and groups.
//
// ## Example Usage
// ### Project
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewDeployToken(ctx, "example", &gitlab.DeployTokenArgs{
// 			ExpiresAt: pulumi.String("2020-03-14T00:00:00.000Z"),
// 			Project:   pulumi.String("example/deploying"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("read_repository"),
// 				pulumi.String("read_registry"),
// 			},
// 			Username: pulumi.String("example-username"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewDeployToken(ctx, "example", &gitlab.DeployTokenArgs{
// 			Group: pulumi.String("example/deploying"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("read_repository"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DeployToken struct {
	pulumi.CustomResourceState

	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	// Either `project` or `group` must be set.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// A name to describe the deploy token with.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	// Either `project` or `group` must be set.
	Project pulumi.StringPtrOutput `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The secret token. This is only populated when creating a new deploy token.
	Token pulumi.StringOutput `pulumi:"token"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewDeployToken registers a new resource with the given unique name, arguments, and options.
func NewDeployToken(ctx *pulumi.Context,
	name string, args *DeployTokenArgs, opts ...pulumi.ResourceOption) (*DeployToken, error) {
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil {
		args = &DeployTokenArgs{}
	}
	var resource DeployToken
	err := ctx.RegisterResource("gitlab:index/deployToken:DeployToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployToken gets an existing DeployToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployTokenState, opts ...pulumi.ResourceOption) (*DeployToken, error) {
	var resource DeployToken
	err := ctx.ReadResource("gitlab:index/deployToken:DeployToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployToken resources.
type deployTokenState struct {
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	// Either `project` or `group` must be set.
	Group *string `pulumi:"group"`
	// A name to describe the deploy token with.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	// Either `project` or `group` must be set.
	Project *string `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`.
	Scopes []string `pulumi:"scopes"`
	// The secret token. This is only populated when creating a new deploy token.
	Token *string `pulumi:"token"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username *string `pulumi:"username"`
}

type DeployTokenState struct {
	ExpiresAt pulumi.StringPtrInput
	// The name or id of the group to add the deploy token to.
	// Either `project` or `group` must be set.
	Group pulumi.StringPtrInput
	// A name to describe the deploy token with.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the deploy token to.
	// Either `project` or `group` must be set.
	Project pulumi.StringPtrInput
	// Valid values: `readRepository`, `readRegistry`.
	Scopes pulumi.StringArrayInput
	// The secret token. This is only populated when creating a new deploy token.
	Token pulumi.StringPtrInput
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringPtrInput
}

func (DeployTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployTokenState)(nil)).Elem()
}

type deployTokenArgs struct {
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	// Either `project` or `group` must be set.
	Group *string `pulumi:"group"`
	// A name to describe the deploy token with.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	// Either `project` or `group` must be set.
	Project *string `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`.
	Scopes []string `pulumi:"scopes"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a DeployToken resource.
type DeployTokenArgs struct {
	ExpiresAt pulumi.StringPtrInput
	// The name or id of the group to add the deploy token to.
	// Either `project` or `group` must be set.
	Group pulumi.StringPtrInput
	// A name to describe the deploy token with.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the deploy token to.
	// Either `project` or `group` must be set.
	Project pulumi.StringPtrInput
	// Valid values: `readRepository`, `readRegistry`.
	Scopes pulumi.StringArrayInput
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringPtrInput
}

func (DeployTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployTokenArgs)(nil)).Elem()
}

type DeployTokenInput interface {
	pulumi.Input

	ToDeployTokenOutput() DeployTokenOutput
	ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput
}

func (DeployToken) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployToken)(nil)).Elem()
}

func (i DeployToken) ToDeployTokenOutput() DeployTokenOutput {
	return i.ToDeployTokenOutputWithContext(context.Background())
}

func (i DeployToken) ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenOutput)
}

type DeployTokenOutput struct {
	*pulumi.OutputState
}

func (DeployTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployTokenOutput)(nil)).Elem()
}

func (o DeployTokenOutput) ToDeployTokenOutput() DeployTokenOutput {
	return o
}

func (o DeployTokenOutput) ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeployTokenOutput{})
}
