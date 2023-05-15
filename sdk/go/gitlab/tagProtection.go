// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `TagProtection` resource allows to manage the lifecycle of a tag protection.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_tags.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v5/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewTagProtection(ctx, "tagProtect", &gitlab.TagProtectionArgs{
//				CreateAccessLevel: pulumi.String("developer"),
//				Project:           pulumi.String("12345"),
//				Tag:               pulumi.String("TagProtected"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Tag protections can be imported using an id made up of `project_id:tag_name`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/tagProtection:TagProtection example 123456789:v1.0.0
//
// ```
type TagProtection struct {
	pulumi.CustomResourceState

	// Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
	CreateAccessLevel pulumi.StringOutput `pulumi:"createAccessLevel"`
	// The id of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// Name of the tag or wildcard.
	Tag pulumi.StringOutput `pulumi:"tag"`
}

// NewTagProtection registers a new resource with the given unique name, arguments, and options.
func NewTagProtection(ctx *pulumi.Context,
	name string, args *TagProtectionArgs, opts ...pulumi.ResourceOption) (*TagProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreateAccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'CreateAccessLevel'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Tag == nil {
		return nil, errors.New("invalid value for required argument 'Tag'")
	}
	var resource TagProtection
	err := ctx.RegisterResource("gitlab:index/tagProtection:TagProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagProtection gets an existing TagProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagProtectionState, opts ...pulumi.ResourceOption) (*TagProtection, error) {
	var resource TagProtection
	err := ctx.ReadResource("gitlab:index/tagProtection:TagProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagProtection resources.
type tagProtectionState struct {
	// Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
	CreateAccessLevel *string `pulumi:"createAccessLevel"`
	// The id of the project.
	Project *string `pulumi:"project"`
	// Name of the tag or wildcard.
	Tag *string `pulumi:"tag"`
}

type TagProtectionState struct {
	// Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
	CreateAccessLevel pulumi.StringPtrInput
	// The id of the project.
	Project pulumi.StringPtrInput
	// Name of the tag or wildcard.
	Tag pulumi.StringPtrInput
}

func (TagProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagProtectionState)(nil)).Elem()
}

type tagProtectionArgs struct {
	// Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
	CreateAccessLevel string `pulumi:"createAccessLevel"`
	// The id of the project.
	Project string `pulumi:"project"`
	// Name of the tag or wildcard.
	Tag string `pulumi:"tag"`
}

// The set of arguments for constructing a TagProtection resource.
type TagProtectionArgs struct {
	// Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
	CreateAccessLevel pulumi.StringInput
	// The id of the project.
	Project pulumi.StringInput
	// Name of the tag or wildcard.
	Tag pulumi.StringInput
}

func (TagProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagProtectionArgs)(nil)).Elem()
}

type TagProtectionInput interface {
	pulumi.Input

	ToTagProtectionOutput() TagProtectionOutput
	ToTagProtectionOutputWithContext(ctx context.Context) TagProtectionOutput
}

func (*TagProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**TagProtection)(nil)).Elem()
}

func (i *TagProtection) ToTagProtectionOutput() TagProtectionOutput {
	return i.ToTagProtectionOutputWithContext(context.Background())
}

func (i *TagProtection) ToTagProtectionOutputWithContext(ctx context.Context) TagProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagProtectionOutput)
}

// TagProtectionArrayInput is an input type that accepts TagProtectionArray and TagProtectionArrayOutput values.
// You can construct a concrete instance of `TagProtectionArrayInput` via:
//
//	TagProtectionArray{ TagProtectionArgs{...} }
type TagProtectionArrayInput interface {
	pulumi.Input

	ToTagProtectionArrayOutput() TagProtectionArrayOutput
	ToTagProtectionArrayOutputWithContext(context.Context) TagProtectionArrayOutput
}

type TagProtectionArray []TagProtectionInput

func (TagProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagProtection)(nil)).Elem()
}

func (i TagProtectionArray) ToTagProtectionArrayOutput() TagProtectionArrayOutput {
	return i.ToTagProtectionArrayOutputWithContext(context.Background())
}

func (i TagProtectionArray) ToTagProtectionArrayOutputWithContext(ctx context.Context) TagProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagProtectionArrayOutput)
}

// TagProtectionMapInput is an input type that accepts TagProtectionMap and TagProtectionMapOutput values.
// You can construct a concrete instance of `TagProtectionMapInput` via:
//
//	TagProtectionMap{ "key": TagProtectionArgs{...} }
type TagProtectionMapInput interface {
	pulumi.Input

	ToTagProtectionMapOutput() TagProtectionMapOutput
	ToTagProtectionMapOutputWithContext(context.Context) TagProtectionMapOutput
}

type TagProtectionMap map[string]TagProtectionInput

func (TagProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagProtection)(nil)).Elem()
}

func (i TagProtectionMap) ToTagProtectionMapOutput() TagProtectionMapOutput {
	return i.ToTagProtectionMapOutputWithContext(context.Background())
}

func (i TagProtectionMap) ToTagProtectionMapOutputWithContext(ctx context.Context) TagProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagProtectionMapOutput)
}

type TagProtectionOutput struct{ *pulumi.OutputState }

func (TagProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagProtection)(nil)).Elem()
}

func (o TagProtectionOutput) ToTagProtectionOutput() TagProtectionOutput {
	return o
}

func (o TagProtectionOutput) ToTagProtectionOutputWithContext(ctx context.Context) TagProtectionOutput {
	return o
}

// Access levels which are allowed to create. Valid values are: `no one`, `developer`, `maintainer`.
func (o TagProtectionOutput) CreateAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *TagProtection) pulumi.StringOutput { return v.CreateAccessLevel }).(pulumi.StringOutput)
}

// The id of the project.
func (o TagProtectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TagProtection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Name of the tag or wildcard.
func (o TagProtectionOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagProtection) pulumi.StringOutput { return v.Tag }).(pulumi.StringOutput)
}

type TagProtectionArrayOutput struct{ *pulumi.OutputState }

func (TagProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagProtection)(nil)).Elem()
}

func (o TagProtectionArrayOutput) ToTagProtectionArrayOutput() TagProtectionArrayOutput {
	return o
}

func (o TagProtectionArrayOutput) ToTagProtectionArrayOutputWithContext(ctx context.Context) TagProtectionArrayOutput {
	return o
}

func (o TagProtectionArrayOutput) Index(i pulumi.IntInput) TagProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagProtection {
		return vs[0].([]*TagProtection)[vs[1].(int)]
	}).(TagProtectionOutput)
}

type TagProtectionMapOutput struct{ *pulumi.OutputState }

func (TagProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagProtection)(nil)).Elem()
}

func (o TagProtectionMapOutput) ToTagProtectionMapOutput() TagProtectionMapOutput {
	return o
}

func (o TagProtectionMapOutput) ToTagProtectionMapOutputWithContext(ctx context.Context) TagProtectionMapOutput {
	return o
}

func (o TagProtectionMapOutput) MapIndex(k pulumi.StringInput) TagProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagProtection {
		return vs[0].(map[string]*TagProtection)[vs[1].(string)]
	}).(TagProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagProtectionInput)(nil)).Elem(), &TagProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagProtectionArrayInput)(nil)).Elem(), TagProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagProtectionMapInput)(nil)).Elem(), TagProtectionMap{})
	pulumi.RegisterOutputType(TagProtectionOutput{})
	pulumi.RegisterOutputType(TagProtectionArrayOutput{})
	pulumi.RegisterOutputType(TagProtectionMapOutput{})
}
