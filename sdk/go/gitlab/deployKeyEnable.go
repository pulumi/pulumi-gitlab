// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `DeployKeyEnable` resource allows to enable an already existing deploy key (see `DeployKey resource`) for a specific project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_keys.html#enable-a-deploy-key)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		parentProject, err := gitlab.NewProject(ctx, "parentProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooProject, err := gitlab.NewProject(ctx, "fooProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		parentDeployKey, err := gitlab.NewDeployKey(ctx, "parentDeployKey", &gitlab.DeployKeyArgs{
// 			Project: parentProject.ID(),
// 			Title:   pulumi.String("Example deploy key"),
// 			Key:     pulumi.String("ssh-rsa AAAA..."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewDeployKeyEnable(ctx, "fooDeployKeyEnable", &gitlab.DeployKeyEnableArgs{
// 			Project: fooProject.ID(),
// 			KeyId:   parentDeployKey.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
// ```
type DeployKeyEnable struct {
	pulumi.CustomResourceState

	// Can deploy key push to the project’s repository.
	CanPush pulumi.BoolOutput `pulumi:"canPush"`
	// Deploy key.
	Key pulumi.StringOutput `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringOutput `pulumi:"project"`
	// Deploy key's title.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewDeployKeyEnable registers a new resource with the given unique name, arguments, and options.
func NewDeployKeyEnable(ctx *pulumi.Context,
	name string, args *DeployKeyEnableArgs, opts ...pulumi.ResourceOption) (*DeployKeyEnable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource DeployKeyEnable
	err := ctx.RegisterResource("gitlab:index/deployKeyEnable:DeployKeyEnable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployKeyEnable gets an existing DeployKeyEnable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployKeyEnable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployKeyEnableState, opts ...pulumi.ResourceOption) (*DeployKeyEnable, error) {
	var resource DeployKeyEnable
	err := ctx.ReadResource("gitlab:index/deployKeyEnable:DeployKeyEnable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployKeyEnable resources.
type deployKeyEnableState struct {
	// Can deploy key push to the project’s repository.
	CanPush *bool `pulumi:"canPush"`
	// Deploy key.
	Key *string `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId *string `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project *string `pulumi:"project"`
	// Deploy key's title.
	Title *string `pulumi:"title"`
}

type DeployKeyEnableState struct {
	// Can deploy key push to the project’s repository.
	CanPush pulumi.BoolPtrInput
	// Deploy key.
	Key pulumi.StringPtrInput
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringPtrInput
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringPtrInput
	// Deploy key's title.
	Title pulumi.StringPtrInput
}

func (DeployKeyEnableState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployKeyEnableState)(nil)).Elem()
}

type deployKeyEnableArgs struct {
	// Can deploy key push to the project’s repository.
	CanPush *bool `pulumi:"canPush"`
	// Deploy key.
	Key *string `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId string `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project string `pulumi:"project"`
	// Deploy key's title.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a DeployKeyEnable resource.
type DeployKeyEnableArgs struct {
	// Can deploy key push to the project’s repository.
	CanPush pulumi.BoolPtrInput
	// Deploy key.
	Key pulumi.StringPtrInput
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringInput
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringInput
	// Deploy key's title.
	Title pulumi.StringPtrInput
}

func (DeployKeyEnableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployKeyEnableArgs)(nil)).Elem()
}

type DeployKeyEnableInput interface {
	pulumi.Input

	ToDeployKeyEnableOutput() DeployKeyEnableOutput
	ToDeployKeyEnableOutputWithContext(ctx context.Context) DeployKeyEnableOutput
}

func (*DeployKeyEnable) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployKeyEnable)(nil)).Elem()
}

func (i *DeployKeyEnable) ToDeployKeyEnableOutput() DeployKeyEnableOutput {
	return i.ToDeployKeyEnableOutputWithContext(context.Background())
}

func (i *DeployKeyEnable) ToDeployKeyEnableOutputWithContext(ctx context.Context) DeployKeyEnableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnableOutput)
}

// DeployKeyEnableArrayInput is an input type that accepts DeployKeyEnableArray and DeployKeyEnableArrayOutput values.
// You can construct a concrete instance of `DeployKeyEnableArrayInput` via:
//
//          DeployKeyEnableArray{ DeployKeyEnableArgs{...} }
type DeployKeyEnableArrayInput interface {
	pulumi.Input

	ToDeployKeyEnableArrayOutput() DeployKeyEnableArrayOutput
	ToDeployKeyEnableArrayOutputWithContext(context.Context) DeployKeyEnableArrayOutput
}

type DeployKeyEnableArray []DeployKeyEnableInput

func (DeployKeyEnableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployKeyEnable)(nil)).Elem()
}

func (i DeployKeyEnableArray) ToDeployKeyEnableArrayOutput() DeployKeyEnableArrayOutput {
	return i.ToDeployKeyEnableArrayOutputWithContext(context.Background())
}

func (i DeployKeyEnableArray) ToDeployKeyEnableArrayOutputWithContext(ctx context.Context) DeployKeyEnableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnableArrayOutput)
}

// DeployKeyEnableMapInput is an input type that accepts DeployKeyEnableMap and DeployKeyEnableMapOutput values.
// You can construct a concrete instance of `DeployKeyEnableMapInput` via:
//
//          DeployKeyEnableMap{ "key": DeployKeyEnableArgs{...} }
type DeployKeyEnableMapInput interface {
	pulumi.Input

	ToDeployKeyEnableMapOutput() DeployKeyEnableMapOutput
	ToDeployKeyEnableMapOutputWithContext(context.Context) DeployKeyEnableMapOutput
}

type DeployKeyEnableMap map[string]DeployKeyEnableInput

func (DeployKeyEnableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployKeyEnable)(nil)).Elem()
}

func (i DeployKeyEnableMap) ToDeployKeyEnableMapOutput() DeployKeyEnableMapOutput {
	return i.ToDeployKeyEnableMapOutputWithContext(context.Background())
}

func (i DeployKeyEnableMap) ToDeployKeyEnableMapOutputWithContext(ctx context.Context) DeployKeyEnableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnableMapOutput)
}

type DeployKeyEnableOutput struct{ *pulumi.OutputState }

func (DeployKeyEnableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployKeyEnable)(nil)).Elem()
}

func (o DeployKeyEnableOutput) ToDeployKeyEnableOutput() DeployKeyEnableOutput {
	return o
}

func (o DeployKeyEnableOutput) ToDeployKeyEnableOutputWithContext(ctx context.Context) DeployKeyEnableOutput {
	return o
}

type DeployKeyEnableArrayOutput struct{ *pulumi.OutputState }

func (DeployKeyEnableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployKeyEnable)(nil)).Elem()
}

func (o DeployKeyEnableArrayOutput) ToDeployKeyEnableArrayOutput() DeployKeyEnableArrayOutput {
	return o
}

func (o DeployKeyEnableArrayOutput) ToDeployKeyEnableArrayOutputWithContext(ctx context.Context) DeployKeyEnableArrayOutput {
	return o
}

func (o DeployKeyEnableArrayOutput) Index(i pulumi.IntInput) DeployKeyEnableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeployKeyEnable {
		return vs[0].([]*DeployKeyEnable)[vs[1].(int)]
	}).(DeployKeyEnableOutput)
}

type DeployKeyEnableMapOutput struct{ *pulumi.OutputState }

func (DeployKeyEnableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployKeyEnable)(nil)).Elem()
}

func (o DeployKeyEnableMapOutput) ToDeployKeyEnableMapOutput() DeployKeyEnableMapOutput {
	return o
}

func (o DeployKeyEnableMapOutput) ToDeployKeyEnableMapOutputWithContext(ctx context.Context) DeployKeyEnableMapOutput {
	return o
}

func (o DeployKeyEnableMapOutput) MapIndex(k pulumi.StringInput) DeployKeyEnableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeployKeyEnable {
		return vs[0].(map[string]*DeployKeyEnable)[vs[1].(string)]
	}).(DeployKeyEnableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployKeyEnableInput)(nil)).Elem(), &DeployKeyEnable{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployKeyEnableArrayInput)(nil)).Elem(), DeployKeyEnableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployKeyEnableMapInput)(nil)).Elem(), DeployKeyEnableMap{})
	pulumi.RegisterOutputType(DeployKeyEnableOutput{})
	pulumi.RegisterOutputType(DeployKeyEnableArrayOutput{})
	pulumi.RegisterOutputType(DeployKeyEnableMapOutput{})
}
