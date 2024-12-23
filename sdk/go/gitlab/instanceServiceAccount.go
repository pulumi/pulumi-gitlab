// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `InstanceServiceAccount` resource allows creating a GitLab instance service account.
//
// > In order for a user to create a user account, they must have admin privileges at the instance level. This makes this feature unavailable on `gitlab.com`
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/user_service_accounts.html)
type InstanceServiceAccount struct {
	pulumi.CustomResourceState

	// The name of the user. If not specified, the default Service account user name is used.
	Name pulumi.StringOutput `pulumi:"name"`
	// The service account id.
	ServiceAccountId pulumi.StringOutput                     `pulumi:"serviceAccountId"`
	Timeouts         InstanceServiceAccountTimeoutsPtrOutput `pulumi:"timeouts"`
	// The username of the user. If not specified, it’s automatically generated.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewInstanceServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewInstanceServiceAccount(ctx *pulumi.Context,
	name string, args *InstanceServiceAccountArgs, opts ...pulumi.ResourceOption) (*InstanceServiceAccount, error) {
	if args == nil {
		args = &InstanceServiceAccountArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceServiceAccount
	err := ctx.RegisterResource("gitlab:index/instanceServiceAccount:InstanceServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceServiceAccount gets an existing InstanceServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceServiceAccountState, opts ...pulumi.ResourceOption) (*InstanceServiceAccount, error) {
	var resource InstanceServiceAccount
	err := ctx.ReadResource("gitlab:index/instanceServiceAccount:InstanceServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceServiceAccount resources.
type instanceServiceAccountState struct {
	// The name of the user. If not specified, the default Service account user name is used.
	Name *string `pulumi:"name"`
	// The service account id.
	ServiceAccountId *string                         `pulumi:"serviceAccountId"`
	Timeouts         *InstanceServiceAccountTimeouts `pulumi:"timeouts"`
	// The username of the user. If not specified, it’s automatically generated.
	Username *string `pulumi:"username"`
}

type InstanceServiceAccountState struct {
	// The name of the user. If not specified, the default Service account user name is used.
	Name pulumi.StringPtrInput
	// The service account id.
	ServiceAccountId pulumi.StringPtrInput
	Timeouts         InstanceServiceAccountTimeoutsPtrInput
	// The username of the user. If not specified, it’s automatically generated.
	Username pulumi.StringPtrInput
}

func (InstanceServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceServiceAccountState)(nil)).Elem()
}

type instanceServiceAccountArgs struct {
	// The name of the user. If not specified, the default Service account user name is used.
	Name     *string                         `pulumi:"name"`
	Timeouts *InstanceServiceAccountTimeouts `pulumi:"timeouts"`
	// The username of the user. If not specified, it’s automatically generated.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a InstanceServiceAccount resource.
type InstanceServiceAccountArgs struct {
	// The name of the user. If not specified, the default Service account user name is used.
	Name     pulumi.StringPtrInput
	Timeouts InstanceServiceAccountTimeoutsPtrInput
	// The username of the user. If not specified, it’s automatically generated.
	Username pulumi.StringPtrInput
}

func (InstanceServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceServiceAccountArgs)(nil)).Elem()
}

type InstanceServiceAccountInput interface {
	pulumi.Input

	ToInstanceServiceAccountOutput() InstanceServiceAccountOutput
	ToInstanceServiceAccountOutputWithContext(ctx context.Context) InstanceServiceAccountOutput
}

func (*InstanceServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceServiceAccount)(nil)).Elem()
}

func (i *InstanceServiceAccount) ToInstanceServiceAccountOutput() InstanceServiceAccountOutput {
	return i.ToInstanceServiceAccountOutputWithContext(context.Background())
}

func (i *InstanceServiceAccount) ToInstanceServiceAccountOutputWithContext(ctx context.Context) InstanceServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceServiceAccountOutput)
}

// InstanceServiceAccountArrayInput is an input type that accepts InstanceServiceAccountArray and InstanceServiceAccountArrayOutput values.
// You can construct a concrete instance of `InstanceServiceAccountArrayInput` via:
//
//	InstanceServiceAccountArray{ InstanceServiceAccountArgs{...} }
type InstanceServiceAccountArrayInput interface {
	pulumi.Input

	ToInstanceServiceAccountArrayOutput() InstanceServiceAccountArrayOutput
	ToInstanceServiceAccountArrayOutputWithContext(context.Context) InstanceServiceAccountArrayOutput
}

type InstanceServiceAccountArray []InstanceServiceAccountInput

func (InstanceServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceServiceAccount)(nil)).Elem()
}

func (i InstanceServiceAccountArray) ToInstanceServiceAccountArrayOutput() InstanceServiceAccountArrayOutput {
	return i.ToInstanceServiceAccountArrayOutputWithContext(context.Background())
}

func (i InstanceServiceAccountArray) ToInstanceServiceAccountArrayOutputWithContext(ctx context.Context) InstanceServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceServiceAccountArrayOutput)
}

// InstanceServiceAccountMapInput is an input type that accepts InstanceServiceAccountMap and InstanceServiceAccountMapOutput values.
// You can construct a concrete instance of `InstanceServiceAccountMapInput` via:
//
//	InstanceServiceAccountMap{ "key": InstanceServiceAccountArgs{...} }
type InstanceServiceAccountMapInput interface {
	pulumi.Input

	ToInstanceServiceAccountMapOutput() InstanceServiceAccountMapOutput
	ToInstanceServiceAccountMapOutputWithContext(context.Context) InstanceServiceAccountMapOutput
}

type InstanceServiceAccountMap map[string]InstanceServiceAccountInput

func (InstanceServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceServiceAccount)(nil)).Elem()
}

func (i InstanceServiceAccountMap) ToInstanceServiceAccountMapOutput() InstanceServiceAccountMapOutput {
	return i.ToInstanceServiceAccountMapOutputWithContext(context.Background())
}

func (i InstanceServiceAccountMap) ToInstanceServiceAccountMapOutputWithContext(ctx context.Context) InstanceServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceServiceAccountMapOutput)
}

type InstanceServiceAccountOutput struct{ *pulumi.OutputState }

func (InstanceServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceServiceAccount)(nil)).Elem()
}

func (o InstanceServiceAccountOutput) ToInstanceServiceAccountOutput() InstanceServiceAccountOutput {
	return o
}

func (o InstanceServiceAccountOutput) ToInstanceServiceAccountOutputWithContext(ctx context.Context) InstanceServiceAccountOutput {
	return o
}

// The name of the user. If not specified, the default Service account user name is used.
func (o InstanceServiceAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceServiceAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The service account id.
func (o InstanceServiceAccountOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceServiceAccount) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

func (o InstanceServiceAccountOutput) Timeouts() InstanceServiceAccountTimeoutsPtrOutput {
	return o.ApplyT(func(v *InstanceServiceAccount) InstanceServiceAccountTimeoutsPtrOutput { return v.Timeouts }).(InstanceServiceAccountTimeoutsPtrOutput)
}

// The username of the user. If not specified, it’s automatically generated.
func (o InstanceServiceAccountOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceServiceAccount) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type InstanceServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (InstanceServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceServiceAccount)(nil)).Elem()
}

func (o InstanceServiceAccountArrayOutput) ToInstanceServiceAccountArrayOutput() InstanceServiceAccountArrayOutput {
	return o
}

func (o InstanceServiceAccountArrayOutput) ToInstanceServiceAccountArrayOutputWithContext(ctx context.Context) InstanceServiceAccountArrayOutput {
	return o
}

func (o InstanceServiceAccountArrayOutput) Index(i pulumi.IntInput) InstanceServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceServiceAccount {
		return vs[0].([]*InstanceServiceAccount)[vs[1].(int)]
	}).(InstanceServiceAccountOutput)
}

type InstanceServiceAccountMapOutput struct{ *pulumi.OutputState }

func (InstanceServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceServiceAccount)(nil)).Elem()
}

func (o InstanceServiceAccountMapOutput) ToInstanceServiceAccountMapOutput() InstanceServiceAccountMapOutput {
	return o
}

func (o InstanceServiceAccountMapOutput) ToInstanceServiceAccountMapOutputWithContext(ctx context.Context) InstanceServiceAccountMapOutput {
	return o
}

func (o InstanceServiceAccountMapOutput) MapIndex(k pulumi.StringInput) InstanceServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceServiceAccount {
		return vs[0].(map[string]*InstanceServiceAccount)[vs[1].(string)]
	}).(InstanceServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceServiceAccountInput)(nil)).Elem(), &InstanceServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceServiceAccountArrayInput)(nil)).Elem(), InstanceServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceServiceAccountMapInput)(nil)).Elem(), InstanceServiceAccountMap{})
	pulumi.RegisterOutputType(InstanceServiceAccountOutput{})
	pulumi.RegisterOutputType(InstanceServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(InstanceServiceAccountMapOutput{})
}
