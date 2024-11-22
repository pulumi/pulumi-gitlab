// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ServicePipelinesEmail` resource allows to manage the lifecycle of a project integration with Pipeline Emails Service.
//
// > This resource is deprecated. use `IntegrationPipelinesEmail`instead!
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
//	"github.com/pulumi/pulumi-gitlab/sdk/v8/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			awesomeProject, err := gitlab.NewProject(ctx, "awesome_project", &gitlab.ProjectArgs{
//				Name:            pulumi.String("awesome_project"),
//				Description:     pulumi.String("My awesome project."),
//				VisibilityLevel: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gitlab.NewServicePipelinesEmail(ctx, "email", &gitlab.ServicePipelinesEmailArgs{
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_pipelines_email`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_service_pipelines_email.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// You can import a gitlab_service_pipelines_email state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/servicePipelinesEmail:ServicePipelinesEmail email 1
// ```
type ServicePipelinesEmail struct {
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

// NewServicePipelinesEmail registers a new resource with the given unique name, arguments, and options.
func NewServicePipelinesEmail(ctx *pulumi.Context,
	name string, args *ServicePipelinesEmailArgs, opts ...pulumi.ResourceOption) (*ServicePipelinesEmail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicePipelinesEmail
	err := ctx.RegisterResource("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePipelinesEmail gets an existing ServicePipelinesEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePipelinesEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePipelinesEmailState, opts ...pulumi.ResourceOption) (*ServicePipelinesEmail, error) {
	var resource ServicePipelinesEmail
	err := ctx.ReadResource("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePipelinesEmail resources.
type servicePipelinesEmailState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients []string `pulumi:"recipients"`
}

type ServicePipelinesEmailState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrInput
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayInput
}

func (ServicePipelinesEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePipelinesEmailState)(nil)).Elem()
}

type servicePipelinesEmailArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients []string `pulumi:"recipients"`
}

// The set of arguments for constructing a ServicePipelinesEmail resource.
type ServicePipelinesEmailArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrInput
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayInput
}

func (ServicePipelinesEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePipelinesEmailArgs)(nil)).Elem()
}

type ServicePipelinesEmailInput interface {
	pulumi.Input

	ToServicePipelinesEmailOutput() ServicePipelinesEmailOutput
	ToServicePipelinesEmailOutputWithContext(ctx context.Context) ServicePipelinesEmailOutput
}

func (*ServicePipelinesEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePipelinesEmail)(nil)).Elem()
}

func (i *ServicePipelinesEmail) ToServicePipelinesEmailOutput() ServicePipelinesEmailOutput {
	return i.ToServicePipelinesEmailOutputWithContext(context.Background())
}

func (i *ServicePipelinesEmail) ToServicePipelinesEmailOutputWithContext(ctx context.Context) ServicePipelinesEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailOutput)
}

// ServicePipelinesEmailArrayInput is an input type that accepts ServicePipelinesEmailArray and ServicePipelinesEmailArrayOutput values.
// You can construct a concrete instance of `ServicePipelinesEmailArrayInput` via:
//
//	ServicePipelinesEmailArray{ ServicePipelinesEmailArgs{...} }
type ServicePipelinesEmailArrayInput interface {
	pulumi.Input

	ToServicePipelinesEmailArrayOutput() ServicePipelinesEmailArrayOutput
	ToServicePipelinesEmailArrayOutputWithContext(context.Context) ServicePipelinesEmailArrayOutput
}

type ServicePipelinesEmailArray []ServicePipelinesEmailInput

func (ServicePipelinesEmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePipelinesEmail)(nil)).Elem()
}

func (i ServicePipelinesEmailArray) ToServicePipelinesEmailArrayOutput() ServicePipelinesEmailArrayOutput {
	return i.ToServicePipelinesEmailArrayOutputWithContext(context.Background())
}

func (i ServicePipelinesEmailArray) ToServicePipelinesEmailArrayOutputWithContext(ctx context.Context) ServicePipelinesEmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailArrayOutput)
}

// ServicePipelinesEmailMapInput is an input type that accepts ServicePipelinesEmailMap and ServicePipelinesEmailMapOutput values.
// You can construct a concrete instance of `ServicePipelinesEmailMapInput` via:
//
//	ServicePipelinesEmailMap{ "key": ServicePipelinesEmailArgs{...} }
type ServicePipelinesEmailMapInput interface {
	pulumi.Input

	ToServicePipelinesEmailMapOutput() ServicePipelinesEmailMapOutput
	ToServicePipelinesEmailMapOutputWithContext(context.Context) ServicePipelinesEmailMapOutput
}

type ServicePipelinesEmailMap map[string]ServicePipelinesEmailInput

func (ServicePipelinesEmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePipelinesEmail)(nil)).Elem()
}

func (i ServicePipelinesEmailMap) ToServicePipelinesEmailMapOutput() ServicePipelinesEmailMapOutput {
	return i.ToServicePipelinesEmailMapOutputWithContext(context.Background())
}

func (i ServicePipelinesEmailMap) ToServicePipelinesEmailMapOutputWithContext(ctx context.Context) ServicePipelinesEmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailMapOutput)
}

type ServicePipelinesEmailOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePipelinesEmail)(nil)).Elem()
}

func (o ServicePipelinesEmailOutput) ToServicePipelinesEmailOutput() ServicePipelinesEmailOutput {
	return o
}

func (o ServicePipelinesEmailOutput) ToServicePipelinesEmailOutputWithContext(ctx context.Context) ServicePipelinesEmailOutput {
	return o
}

// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
func (o ServicePipelinesEmailOutput) BranchesToBeNotified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePipelinesEmail) pulumi.StringPtrOutput { return v.BranchesToBeNotified }).(pulumi.StringPtrOutput)
}

// Notify only broken pipelines. Default is true.
func (o ServicePipelinesEmailOutput) NotifyOnlyBrokenPipelines() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServicePipelinesEmail) pulumi.BoolPtrOutput { return v.NotifyOnlyBrokenPipelines }).(pulumi.BoolPtrOutput)
}

// ID of the project you want to activate integration on.
func (o ServicePipelinesEmailOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePipelinesEmail) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// ) email addresses where notifications are sent.
func (o ServicePipelinesEmailOutput) Recipients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServicePipelinesEmail) pulumi.StringArrayOutput { return v.Recipients }).(pulumi.StringArrayOutput)
}

type ServicePipelinesEmailArrayOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePipelinesEmail)(nil)).Elem()
}

func (o ServicePipelinesEmailArrayOutput) ToServicePipelinesEmailArrayOutput() ServicePipelinesEmailArrayOutput {
	return o
}

func (o ServicePipelinesEmailArrayOutput) ToServicePipelinesEmailArrayOutputWithContext(ctx context.Context) ServicePipelinesEmailArrayOutput {
	return o
}

func (o ServicePipelinesEmailArrayOutput) Index(i pulumi.IntInput) ServicePipelinesEmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePipelinesEmail {
		return vs[0].([]*ServicePipelinesEmail)[vs[1].(int)]
	}).(ServicePipelinesEmailOutput)
}

type ServicePipelinesEmailMapOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePipelinesEmail)(nil)).Elem()
}

func (o ServicePipelinesEmailMapOutput) ToServicePipelinesEmailMapOutput() ServicePipelinesEmailMapOutput {
	return o
}

func (o ServicePipelinesEmailMapOutput) ToServicePipelinesEmailMapOutputWithContext(ctx context.Context) ServicePipelinesEmailMapOutput {
	return o
}

func (o ServicePipelinesEmailMapOutput) MapIndex(k pulumi.StringInput) ServicePipelinesEmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePipelinesEmail {
		return vs[0].(map[string]*ServicePipelinesEmail)[vs[1].(string)]
	}).(ServicePipelinesEmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailInput)(nil)).Elem(), &ServicePipelinesEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailArrayInput)(nil)).Elem(), ServicePipelinesEmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailMapInput)(nil)).Elem(), ServicePipelinesEmailMap{})
	pulumi.RegisterOutputType(ServicePipelinesEmailOutput{})
	pulumi.RegisterOutputType(ServicePipelinesEmailArrayOutput{})
	pulumi.RegisterOutputType(ServicePipelinesEmailMapOutput{})
}
