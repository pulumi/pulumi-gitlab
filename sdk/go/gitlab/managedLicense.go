// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ManagedLicense` resource allows to manage the lifecycle of a managed license.
//
// > This resource requires a GitLab Enterprise instance.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/managed_licenses.html)
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
// 		foo, err := gitlab.NewProject(ctx, "foo", &gitlab.ProjectArgs{
// 			Description:     pulumi.String("Lorem Ipsum"),
// 			VisibilityLevel: pulumi.String("public"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewManagedLicense(ctx, "mit", &gitlab.ManagedLicenseArgs{
// 			Project:        foo.ID(),
// 			ApprovalStatus: pulumi.String("allowed"),
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
// # You can import this resource with an id made up of `{project-id}:{license-id}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/managedLicense:ManagedLicense foo 1:2
// ```
type ManagedLicense struct {
	pulumi.CustomResourceState

	// The approval status of the license. Valid values are: `approved`, `blacklisted`, `allowed`, `denied`. "approved" and
	// "blacklisted" have been deprecated in favor of "allowed" and "denied"; use "allowed" and "denied" for GitLab versions
	// 15.0 and higher. Prior to version 15.0 and after 14.6, the values are equivalent.
	ApprovalStatus pulumi.StringOutput `pulumi:"approvalStatus"`
	// The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license')
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project under which the managed license will be created.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewManagedLicense registers a new resource with the given unique name, arguments, and options.
func NewManagedLicense(ctx *pulumi.Context,
	name string, args *ManagedLicenseArgs, opts ...pulumi.ResourceOption) (*ManagedLicense, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApprovalStatus == nil {
		return nil, errors.New("invalid value for required argument 'ApprovalStatus'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource ManagedLicense
	err := ctx.RegisterResource("gitlab:index/managedLicense:ManagedLicense", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedLicense gets an existing ManagedLicense resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedLicenseState, opts ...pulumi.ResourceOption) (*ManagedLicense, error) {
	var resource ManagedLicense
	err := ctx.ReadResource("gitlab:index/managedLicense:ManagedLicense", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedLicense resources.
type managedLicenseState struct {
	// The approval status of the license. Valid values are: `approved`, `blacklisted`, `allowed`, `denied`. "approved" and
	// "blacklisted" have been deprecated in favor of "allowed" and "denied"; use "allowed" and "denied" for GitLab versions
	// 15.0 and higher. Prior to version 15.0 and after 14.6, the values are equivalent.
	ApprovalStatus *string `pulumi:"approvalStatus"`
	// The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license')
	Name *string `pulumi:"name"`
	// The ID of the project under which the managed license will be created.
	Project *string `pulumi:"project"`
}

type ManagedLicenseState struct {
	// The approval status of the license. Valid values are: `approved`, `blacklisted`, `allowed`, `denied`. "approved" and
	// "blacklisted" have been deprecated in favor of "allowed" and "denied"; use "allowed" and "denied" for GitLab versions
	// 15.0 and higher. Prior to version 15.0 and after 14.6, the values are equivalent.
	ApprovalStatus pulumi.StringPtrInput
	// The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license')
	Name pulumi.StringPtrInput
	// The ID of the project under which the managed license will be created.
	Project pulumi.StringPtrInput
}

func (ManagedLicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedLicenseState)(nil)).Elem()
}

type managedLicenseArgs struct {
	// The approval status of the license. Valid values are: `approved`, `blacklisted`, `allowed`, `denied`. "approved" and
	// "blacklisted" have been deprecated in favor of "allowed" and "denied"; use "allowed" and "denied" for GitLab versions
	// 15.0 and higher. Prior to version 15.0 and after 14.6, the values are equivalent.
	ApprovalStatus string `pulumi:"approvalStatus"`
	// The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license')
	Name *string `pulumi:"name"`
	// The ID of the project under which the managed license will be created.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ManagedLicense resource.
type ManagedLicenseArgs struct {
	// The approval status of the license. Valid values are: `approved`, `blacklisted`, `allowed`, `denied`. "approved" and
	// "blacklisted" have been deprecated in favor of "allowed" and "denied"; use "allowed" and "denied" for GitLab versions
	// 15.0 and higher. Prior to version 15.0 and after 14.6, the values are equivalent.
	ApprovalStatus pulumi.StringInput
	// The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license')
	Name pulumi.StringPtrInput
	// The ID of the project under which the managed license will be created.
	Project pulumi.StringInput
}

func (ManagedLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedLicenseArgs)(nil)).Elem()
}

type ManagedLicenseInput interface {
	pulumi.Input

	ToManagedLicenseOutput() ManagedLicenseOutput
	ToManagedLicenseOutputWithContext(ctx context.Context) ManagedLicenseOutput
}

func (*ManagedLicense) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedLicense)(nil)).Elem()
}

func (i *ManagedLicense) ToManagedLicenseOutput() ManagedLicenseOutput {
	return i.ToManagedLicenseOutputWithContext(context.Background())
}

func (i *ManagedLicense) ToManagedLicenseOutputWithContext(ctx context.Context) ManagedLicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedLicenseOutput)
}

// ManagedLicenseArrayInput is an input type that accepts ManagedLicenseArray and ManagedLicenseArrayOutput values.
// You can construct a concrete instance of `ManagedLicenseArrayInput` via:
//
//          ManagedLicenseArray{ ManagedLicenseArgs{...} }
type ManagedLicenseArrayInput interface {
	pulumi.Input

	ToManagedLicenseArrayOutput() ManagedLicenseArrayOutput
	ToManagedLicenseArrayOutputWithContext(context.Context) ManagedLicenseArrayOutput
}

type ManagedLicenseArray []ManagedLicenseInput

func (ManagedLicenseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedLicense)(nil)).Elem()
}

func (i ManagedLicenseArray) ToManagedLicenseArrayOutput() ManagedLicenseArrayOutput {
	return i.ToManagedLicenseArrayOutputWithContext(context.Background())
}

func (i ManagedLicenseArray) ToManagedLicenseArrayOutputWithContext(ctx context.Context) ManagedLicenseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedLicenseArrayOutput)
}

// ManagedLicenseMapInput is an input type that accepts ManagedLicenseMap and ManagedLicenseMapOutput values.
// You can construct a concrete instance of `ManagedLicenseMapInput` via:
//
//          ManagedLicenseMap{ "key": ManagedLicenseArgs{...} }
type ManagedLicenseMapInput interface {
	pulumi.Input

	ToManagedLicenseMapOutput() ManagedLicenseMapOutput
	ToManagedLicenseMapOutputWithContext(context.Context) ManagedLicenseMapOutput
}

type ManagedLicenseMap map[string]ManagedLicenseInput

func (ManagedLicenseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedLicense)(nil)).Elem()
}

func (i ManagedLicenseMap) ToManagedLicenseMapOutput() ManagedLicenseMapOutput {
	return i.ToManagedLicenseMapOutputWithContext(context.Background())
}

func (i ManagedLicenseMap) ToManagedLicenseMapOutputWithContext(ctx context.Context) ManagedLicenseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedLicenseMapOutput)
}

type ManagedLicenseOutput struct{ *pulumi.OutputState }

func (ManagedLicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedLicense)(nil)).Elem()
}

func (o ManagedLicenseOutput) ToManagedLicenseOutput() ManagedLicenseOutput {
	return o
}

func (o ManagedLicenseOutput) ToManagedLicenseOutputWithContext(ctx context.Context) ManagedLicenseOutput {
	return o
}

type ManagedLicenseArrayOutput struct{ *pulumi.OutputState }

func (ManagedLicenseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedLicense)(nil)).Elem()
}

func (o ManagedLicenseArrayOutput) ToManagedLicenseArrayOutput() ManagedLicenseArrayOutput {
	return o
}

func (o ManagedLicenseArrayOutput) ToManagedLicenseArrayOutputWithContext(ctx context.Context) ManagedLicenseArrayOutput {
	return o
}

func (o ManagedLicenseArrayOutput) Index(i pulumi.IntInput) ManagedLicenseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedLicense {
		return vs[0].([]*ManagedLicense)[vs[1].(int)]
	}).(ManagedLicenseOutput)
}

type ManagedLicenseMapOutput struct{ *pulumi.OutputState }

func (ManagedLicenseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedLicense)(nil)).Elem()
}

func (o ManagedLicenseMapOutput) ToManagedLicenseMapOutput() ManagedLicenseMapOutput {
	return o
}

func (o ManagedLicenseMapOutput) ToManagedLicenseMapOutputWithContext(ctx context.Context) ManagedLicenseMapOutput {
	return o
}

func (o ManagedLicenseMapOutput) MapIndex(k pulumi.StringInput) ManagedLicenseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedLicense {
		return vs[0].(map[string]*ManagedLicense)[vs[1].(string)]
	}).(ManagedLicenseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedLicenseInput)(nil)).Elem(), &ManagedLicense{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedLicenseArrayInput)(nil)).Elem(), ManagedLicenseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedLicenseMapInput)(nil)).Elem(), ManagedLicenseMap{})
	pulumi.RegisterOutputType(ManagedLicenseOutput{})
	pulumi.RegisterOutputType(ManagedLicenseArrayOutput{})
	pulumi.RegisterOutputType(ManagedLicenseMapOutput{})
}
