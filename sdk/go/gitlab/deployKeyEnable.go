// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_deploy\_key\_enable
//
// This resource allows you to enable pre-existing deploy keys for your GitLab projects.
//
// **the GITLAB KEY_ID for the deploy key must be known**
//
// ## Example Usage
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
// 		parentProject, err := gitlab.NewProject(ctx, "parentProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooProject, err := gitlab.NewProject(ctx, "fooProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		parentDeployKey, err := gitlab.NewDeployKey(ctx, "parentDeployKey", &gitlab.DeployKeyArgs{
// 			Key:     pulumi.String("ssh-rsa AAAA..."),
// 			Project: parentProject.ID(),
// 			Title:   pulumi.String("Example deploy key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewDeployKeyEnable(ctx, "fooDeployKeyEnable", &gitlab.DeployKeyEnableArgs{
// 			KeyId:   parentDeployKey.ID(),
// 			Project: fooProject.ID(),
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
// GitLab enabled deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/deployKeyEnable:DeployKeyEnable example 12345:67890
// ```
type DeployKeyEnable struct {
	pulumi.CustomResourceState

	CanPush pulumi.BoolOutput   `pulumi:"canPush"`
	Key     pulumi.StringOutput `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringOutput `pulumi:"project"`
	Title   pulumi.StringOutput `pulumi:"title"`
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
	CanPush *bool   `pulumi:"canPush"`
	Key     *string `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId *string `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project *string `pulumi:"project"`
	Title   *string `pulumi:"title"`
}

type DeployKeyEnableState struct {
	CanPush pulumi.BoolPtrInput
	Key     pulumi.StringPtrInput
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringPtrInput
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringPtrInput
	Title   pulumi.StringPtrInput
}

func (DeployKeyEnableState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployKeyEnableState)(nil)).Elem()
}

type deployKeyEnableArgs struct {
	CanPush *bool   `pulumi:"canPush"`
	Key     *string `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId string `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project string  `pulumi:"project"`
	Title   *string `pulumi:"title"`
}

// The set of arguments for constructing a DeployKeyEnable resource.
type DeployKeyEnableArgs struct {
	CanPush pulumi.BoolPtrInput
	Key     pulumi.StringPtrInput
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringInput
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringInput
	Title   pulumi.StringPtrInput
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
	return reflect.TypeOf((*DeployKeyEnable)(nil))
}

func (i *DeployKeyEnable) ToDeployKeyEnableOutput() DeployKeyEnableOutput {
	return i.ToDeployKeyEnableOutputWithContext(context.Background())
}

func (i *DeployKeyEnable) ToDeployKeyEnableOutputWithContext(ctx context.Context) DeployKeyEnableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnableOutput)
}

func (i *DeployKeyEnable) ToDeployKeyEnablePtrOutput() DeployKeyEnablePtrOutput {
	return i.ToDeployKeyEnablePtrOutputWithContext(context.Background())
}

func (i *DeployKeyEnable) ToDeployKeyEnablePtrOutputWithContext(ctx context.Context) DeployKeyEnablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnablePtrOutput)
}

type DeployKeyEnablePtrInput interface {
	pulumi.Input

	ToDeployKeyEnablePtrOutput() DeployKeyEnablePtrOutput
	ToDeployKeyEnablePtrOutputWithContext(ctx context.Context) DeployKeyEnablePtrOutput
}

type deployKeyEnablePtrType DeployKeyEnableArgs

func (*deployKeyEnablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployKeyEnable)(nil))
}

func (i *deployKeyEnablePtrType) ToDeployKeyEnablePtrOutput() DeployKeyEnablePtrOutput {
	return i.ToDeployKeyEnablePtrOutputWithContext(context.Background())
}

func (i *deployKeyEnablePtrType) ToDeployKeyEnablePtrOutputWithContext(ctx context.Context) DeployKeyEnablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnablePtrOutput)
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
	return reflect.TypeOf(([]*DeployKeyEnable)(nil))
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
	return reflect.TypeOf((map[string]*DeployKeyEnable)(nil))
}

func (i DeployKeyEnableMap) ToDeployKeyEnableMapOutput() DeployKeyEnableMapOutput {
	return i.ToDeployKeyEnableMapOutputWithContext(context.Background())
}

func (i DeployKeyEnableMap) ToDeployKeyEnableMapOutputWithContext(ctx context.Context) DeployKeyEnableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployKeyEnableMapOutput)
}

type DeployKeyEnableOutput struct {
	*pulumi.OutputState
}

func (DeployKeyEnableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployKeyEnable)(nil))
}

func (o DeployKeyEnableOutput) ToDeployKeyEnableOutput() DeployKeyEnableOutput {
	return o
}

func (o DeployKeyEnableOutput) ToDeployKeyEnableOutputWithContext(ctx context.Context) DeployKeyEnableOutput {
	return o
}

func (o DeployKeyEnableOutput) ToDeployKeyEnablePtrOutput() DeployKeyEnablePtrOutput {
	return o.ToDeployKeyEnablePtrOutputWithContext(context.Background())
}

func (o DeployKeyEnableOutput) ToDeployKeyEnablePtrOutputWithContext(ctx context.Context) DeployKeyEnablePtrOutput {
	return o.ApplyT(func(v DeployKeyEnable) *DeployKeyEnable {
		return &v
	}).(DeployKeyEnablePtrOutput)
}

type DeployKeyEnablePtrOutput struct {
	*pulumi.OutputState
}

func (DeployKeyEnablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployKeyEnable)(nil))
}

func (o DeployKeyEnablePtrOutput) ToDeployKeyEnablePtrOutput() DeployKeyEnablePtrOutput {
	return o
}

func (o DeployKeyEnablePtrOutput) ToDeployKeyEnablePtrOutputWithContext(ctx context.Context) DeployKeyEnablePtrOutput {
	return o
}

type DeployKeyEnableArrayOutput struct{ *pulumi.OutputState }

func (DeployKeyEnableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeployKeyEnable)(nil))
}

func (o DeployKeyEnableArrayOutput) ToDeployKeyEnableArrayOutput() DeployKeyEnableArrayOutput {
	return o
}

func (o DeployKeyEnableArrayOutput) ToDeployKeyEnableArrayOutputWithContext(ctx context.Context) DeployKeyEnableArrayOutput {
	return o
}

func (o DeployKeyEnableArrayOutput) Index(i pulumi.IntInput) DeployKeyEnableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeployKeyEnable {
		return vs[0].([]DeployKeyEnable)[vs[1].(int)]
	}).(DeployKeyEnableOutput)
}

type DeployKeyEnableMapOutput struct{ *pulumi.OutputState }

func (DeployKeyEnableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeployKeyEnable)(nil))
}

func (o DeployKeyEnableMapOutput) ToDeployKeyEnableMapOutput() DeployKeyEnableMapOutput {
	return o
}

func (o DeployKeyEnableMapOutput) ToDeployKeyEnableMapOutputWithContext(ctx context.Context) DeployKeyEnableMapOutput {
	return o
}

func (o DeployKeyEnableMapOutput) MapIndex(k pulumi.StringInput) DeployKeyEnableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeployKeyEnable {
		return vs[0].(map[string]DeployKeyEnable)[vs[1].(string)]
	}).(DeployKeyEnableOutput)
}

func init() {
	pulumi.RegisterOutputType(DeployKeyEnableOutput{})
	pulumi.RegisterOutputType(DeployKeyEnablePtrOutput{})
	pulumi.RegisterOutputType(DeployKeyEnableArrayOutput{})
	pulumi.RegisterOutputType(DeployKeyEnableMapOutput{})
}
