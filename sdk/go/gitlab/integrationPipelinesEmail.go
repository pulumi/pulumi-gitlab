// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `IntegrationPipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#pipeline-emails)
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
//			awesomeProject, err := gitlab.NewProject(ctx, "awesomeProject", &gitlab.ProjectArgs{
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewIntegrationPipelinesEmail(ctx, "email", &gitlab.IntegrationPipelinesEmailArgs{
//				Project: awesomeProject.ID(),
//				Recipients: pulumi.StringArray{
//					pulumi.String("gitlab@user.create"),
//				},
//				NotifyOnlyBrokenPipelines: pulumi.Bool(true),
//				BranchesToBeNotified:      pulumi.String("all"),
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
// You can import a gitlab_integration_pipelines_email state using the project ID, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail email 1
//
// ```
type IntegrationPipelinesEmail struct {
	pulumi.CustomResourceState

	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrOutput `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrOutput `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayOutput `pulumi:"recipients"`
}

// NewIntegrationPipelinesEmail registers a new resource with the given unique name, arguments, and options.
func NewIntegrationPipelinesEmail(ctx *pulumi.Context,
	name string, args *IntegrationPipelinesEmailArgs, opts ...pulumi.ResourceOption) (*IntegrationPipelinesEmail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	var resource IntegrationPipelinesEmail
	err := ctx.RegisterResource("gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationPipelinesEmail gets an existing IntegrationPipelinesEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationPipelinesEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationPipelinesEmailState, opts ...pulumi.ResourceOption) (*IntegrationPipelinesEmail, error) {
	var resource IntegrationPipelinesEmail
	err := ctx.ReadResource("gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationPipelinesEmail resources.
type integrationPipelinesEmailState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients []string `pulumi:"recipients"`
}

type IntegrationPipelinesEmailState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrInput
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayInput
}

func (IntegrationPipelinesEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationPipelinesEmailState)(nil)).Elem()
}

type integrationPipelinesEmailArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients []string `pulumi:"recipients"`
}

// The set of arguments for constructing a IntegrationPipelinesEmail resource.
type IntegrationPipelinesEmailArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrInput
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayInput
}

func (IntegrationPipelinesEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationPipelinesEmailArgs)(nil)).Elem()
}

type IntegrationPipelinesEmailInput interface {
	pulumi.Input

	ToIntegrationPipelinesEmailOutput() IntegrationPipelinesEmailOutput
	ToIntegrationPipelinesEmailOutputWithContext(ctx context.Context) IntegrationPipelinesEmailOutput
}

func (*IntegrationPipelinesEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationPipelinesEmail)(nil)).Elem()
}

func (i *IntegrationPipelinesEmail) ToIntegrationPipelinesEmailOutput() IntegrationPipelinesEmailOutput {
	return i.ToIntegrationPipelinesEmailOutputWithContext(context.Background())
}

func (i *IntegrationPipelinesEmail) ToIntegrationPipelinesEmailOutputWithContext(ctx context.Context) IntegrationPipelinesEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationPipelinesEmailOutput)
}

// IntegrationPipelinesEmailArrayInput is an input type that accepts IntegrationPipelinesEmailArray and IntegrationPipelinesEmailArrayOutput values.
// You can construct a concrete instance of `IntegrationPipelinesEmailArrayInput` via:
//
//	IntegrationPipelinesEmailArray{ IntegrationPipelinesEmailArgs{...} }
type IntegrationPipelinesEmailArrayInput interface {
	pulumi.Input

	ToIntegrationPipelinesEmailArrayOutput() IntegrationPipelinesEmailArrayOutput
	ToIntegrationPipelinesEmailArrayOutputWithContext(context.Context) IntegrationPipelinesEmailArrayOutput
}

type IntegrationPipelinesEmailArray []IntegrationPipelinesEmailInput

func (IntegrationPipelinesEmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationPipelinesEmail)(nil)).Elem()
}

func (i IntegrationPipelinesEmailArray) ToIntegrationPipelinesEmailArrayOutput() IntegrationPipelinesEmailArrayOutput {
	return i.ToIntegrationPipelinesEmailArrayOutputWithContext(context.Background())
}

func (i IntegrationPipelinesEmailArray) ToIntegrationPipelinesEmailArrayOutputWithContext(ctx context.Context) IntegrationPipelinesEmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationPipelinesEmailArrayOutput)
}

// IntegrationPipelinesEmailMapInput is an input type that accepts IntegrationPipelinesEmailMap and IntegrationPipelinesEmailMapOutput values.
// You can construct a concrete instance of `IntegrationPipelinesEmailMapInput` via:
//
//	IntegrationPipelinesEmailMap{ "key": IntegrationPipelinesEmailArgs{...} }
type IntegrationPipelinesEmailMapInput interface {
	pulumi.Input

	ToIntegrationPipelinesEmailMapOutput() IntegrationPipelinesEmailMapOutput
	ToIntegrationPipelinesEmailMapOutputWithContext(context.Context) IntegrationPipelinesEmailMapOutput
}

type IntegrationPipelinesEmailMap map[string]IntegrationPipelinesEmailInput

func (IntegrationPipelinesEmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationPipelinesEmail)(nil)).Elem()
}

func (i IntegrationPipelinesEmailMap) ToIntegrationPipelinesEmailMapOutput() IntegrationPipelinesEmailMapOutput {
	return i.ToIntegrationPipelinesEmailMapOutputWithContext(context.Background())
}

func (i IntegrationPipelinesEmailMap) ToIntegrationPipelinesEmailMapOutputWithContext(ctx context.Context) IntegrationPipelinesEmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationPipelinesEmailMapOutput)
}

type IntegrationPipelinesEmailOutput struct{ *pulumi.OutputState }

func (IntegrationPipelinesEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationPipelinesEmail)(nil)).Elem()
}

func (o IntegrationPipelinesEmailOutput) ToIntegrationPipelinesEmailOutput() IntegrationPipelinesEmailOutput {
	return o
}

func (o IntegrationPipelinesEmailOutput) ToIntegrationPipelinesEmailOutputWithContext(ctx context.Context) IntegrationPipelinesEmailOutput {
	return o
}

// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
func (o IntegrationPipelinesEmailOutput) BranchesToBeNotified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationPipelinesEmail) pulumi.StringPtrOutput { return v.BranchesToBeNotified }).(pulumi.StringPtrOutput)
}

// Notify only broken pipelines. Default is true.
func (o IntegrationPipelinesEmailOutput) NotifyOnlyBrokenPipelines() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationPipelinesEmail) pulumi.BoolPtrOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolPtrOutput)
}

// ID of the project you want to activate integration on.
func (o IntegrationPipelinesEmailOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationPipelinesEmail) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// ) email addresses where notifications are sent.
func (o IntegrationPipelinesEmailOutput) Recipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IntegrationPipelinesEmail) pulumi.StringArrayOutput { return v.Recipients }).(pulumi.StringArrayOutput)
}

type IntegrationPipelinesEmailArrayOutput struct{ *pulumi.OutputState }

func (IntegrationPipelinesEmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationPipelinesEmail)(nil)).Elem()
}

func (o IntegrationPipelinesEmailArrayOutput) ToIntegrationPipelinesEmailArrayOutput() IntegrationPipelinesEmailArrayOutput {
	return o
}

func (o IntegrationPipelinesEmailArrayOutput) ToIntegrationPipelinesEmailArrayOutputWithContext(ctx context.Context) IntegrationPipelinesEmailArrayOutput {
	return o
}

func (o IntegrationPipelinesEmailArrayOutput) Index(i pulumi.IntInput) IntegrationPipelinesEmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationPipelinesEmail {
		return vs[0].([]*IntegrationPipelinesEmail)[vs[1].(int)]
	}).(IntegrationPipelinesEmailOutput)
}

type IntegrationPipelinesEmailMapOutput struct{ *pulumi.OutputState }

func (IntegrationPipelinesEmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationPipelinesEmail)(nil)).Elem()
}

func (o IntegrationPipelinesEmailMapOutput) ToIntegrationPipelinesEmailMapOutput() IntegrationPipelinesEmailMapOutput {
	return o
}

func (o IntegrationPipelinesEmailMapOutput) ToIntegrationPipelinesEmailMapOutputWithContext(ctx context.Context) IntegrationPipelinesEmailMapOutput {
	return o
}

func (o IntegrationPipelinesEmailMapOutput) MapIndex(k pulumi.StringInput) IntegrationPipelinesEmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationPipelinesEmail {
		return vs[0].(map[string]*IntegrationPipelinesEmail)[vs[1].(string)]
	}).(IntegrationPipelinesEmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationPipelinesEmailInput)(nil)).Elem(), &IntegrationPipelinesEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationPipelinesEmailArrayInput)(nil)).Elem(), IntegrationPipelinesEmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationPipelinesEmailMapInput)(nil)).Elem(), IntegrationPipelinesEmailMap{})
	pulumi.RegisterOutputType(IntegrationPipelinesEmailOutput{})
	pulumi.RegisterOutputType(IntegrationPipelinesEmailArrayOutput{})
	pulumi.RegisterOutputType(IntegrationPipelinesEmailMapOutput{})
}