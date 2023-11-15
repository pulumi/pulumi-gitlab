// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v6/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ServiceEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.
//
// > This resource is deprecated. Please use `IntegrationEmailsOnPush` instead!
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#emails-on-push)
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
//			_, err = gitlab.NewServiceEmailsOnPush(ctx, "emails", &gitlab.ServiceEmailsOnPushArgs{
//				Project:    awesomeProject.ID(),
//				Recipients: pulumi.String("myrecipient@example.com myotherrecipient@example.com"),
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
// You can import a gitlab_service_emails_on_push state using the project ID, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush emails 1
//
// ```
type ServiceEmailsOnPush struct {
	pulumi.CustomResourceState

	// Whether the integration is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
	BranchesToBeNotified pulumi.StringPtrOutput `pulumi:"branchesToBeNotified"`
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Disable code diffs.
	DisableDiffs pulumi.BoolPtrOutput `pulumi:"disableDiffs"`
	// ID or full-path of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// Enable notifications for push events.
	PushEvents pulumi.BoolPtrOutput `pulumi:"pushEvents"`
	// Emails separated by whitespace.
	Recipients pulumi.StringOutput `pulumi:"recipients"`
	// Send from committer.
	SendFromCommitterEmail pulumi.BoolPtrOutput `pulumi:"sendFromCommitterEmail"`
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolPtrOutput `pulumi:"tagPushEvents"`
	// Title of the integration.
	Title pulumi.StringOutput `pulumi:"title"`
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewServiceEmailsOnPush registers a new resource with the given unique name, arguments, and options.
func NewServiceEmailsOnPush(ctx *pulumi.Context,
	name string, args *ServiceEmailsOnPushArgs, opts ...pulumi.ResourceOption) (*ServiceEmailsOnPush, error) {
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
	var resource ServiceEmailsOnPush
	err := ctx.RegisterResource("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEmailsOnPush gets an existing ServiceEmailsOnPush resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEmailsOnPush(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEmailsOnPushState, opts ...pulumi.ResourceOption) (*ServiceEmailsOnPush, error) {
	var resource ServiceEmailsOnPush
	err := ctx.ReadResource("gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEmailsOnPush resources.
type serviceEmailsOnPushState struct {
	// Whether the integration is active.
	Active *bool `pulumi:"active"`
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Disable code diffs.
	DisableDiffs *bool `pulumi:"disableDiffs"`
	// ID or full-path of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// Enable notifications for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Emails separated by whitespace.
	Recipients *string `pulumi:"recipients"`
	// Send from committer.
	SendFromCommitterEmail *bool `pulumi:"sendFromCommitterEmail"`
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug *string `pulumi:"slug"`
	// Enable notifications for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
	// Title of the integration.
	Title *string `pulumi:"title"`
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ServiceEmailsOnPushState struct {
	// Whether the integration is active.
	Active pulumi.BoolPtrInput
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
	BranchesToBeNotified pulumi.StringPtrInput
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt pulumi.StringPtrInput
	// Disable code diffs.
	DisableDiffs pulumi.BoolPtrInput
	// ID or full-path of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// Enable notifications for push events.
	PushEvents pulumi.BoolPtrInput
	// Emails separated by whitespace.
	Recipients pulumi.StringPtrInput
	// Send from committer.
	SendFromCommitterEmail pulumi.BoolPtrInput
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringPtrInput
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolPtrInput
	// Title of the integration.
	Title pulumi.StringPtrInput
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (ServiceEmailsOnPushState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEmailsOnPushState)(nil)).Elem()
}

type serviceEmailsOnPushArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Disable code diffs.
	DisableDiffs *bool `pulumi:"disableDiffs"`
	// ID or full-path of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// Enable notifications for push events.
	PushEvents *bool `pulumi:"pushEvents"`
	// Emails separated by whitespace.
	Recipients string `pulumi:"recipients"`
	// Send from committer.
	SendFromCommitterEmail *bool `pulumi:"sendFromCommitterEmail"`
	// Enable notifications for tag push events.
	TagPushEvents *bool `pulumi:"tagPushEvents"`
}

// The set of arguments for constructing a ServiceEmailsOnPush resource.
type ServiceEmailsOnPushArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
	BranchesToBeNotified pulumi.StringPtrInput
	// Disable code diffs.
	DisableDiffs pulumi.BoolPtrInput
	// ID or full-path of the project you want to activate integration on.
	Project pulumi.StringInput
	// Enable notifications for push events.
	PushEvents pulumi.BoolPtrInput
	// Emails separated by whitespace.
	Recipients pulumi.StringInput
	// Send from committer.
	SendFromCommitterEmail pulumi.BoolPtrInput
	// Enable notifications for tag push events.
	TagPushEvents pulumi.BoolPtrInput
}

func (ServiceEmailsOnPushArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEmailsOnPushArgs)(nil)).Elem()
}

type ServiceEmailsOnPushInput interface {
	pulumi.Input

	ToServiceEmailsOnPushOutput() ServiceEmailsOnPushOutput
	ToServiceEmailsOnPushOutputWithContext(ctx context.Context) ServiceEmailsOnPushOutput
}

func (*ServiceEmailsOnPush) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEmailsOnPush)(nil)).Elem()
}

func (i *ServiceEmailsOnPush) ToServiceEmailsOnPushOutput() ServiceEmailsOnPushOutput {
	return i.ToServiceEmailsOnPushOutputWithContext(context.Background())
}

func (i *ServiceEmailsOnPush) ToServiceEmailsOnPushOutputWithContext(ctx context.Context) ServiceEmailsOnPushOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEmailsOnPushOutput)
}

// ServiceEmailsOnPushArrayInput is an input type that accepts ServiceEmailsOnPushArray and ServiceEmailsOnPushArrayOutput values.
// You can construct a concrete instance of `ServiceEmailsOnPushArrayInput` via:
//
//	ServiceEmailsOnPushArray{ ServiceEmailsOnPushArgs{...} }
type ServiceEmailsOnPushArrayInput interface {
	pulumi.Input

	ToServiceEmailsOnPushArrayOutput() ServiceEmailsOnPushArrayOutput
	ToServiceEmailsOnPushArrayOutputWithContext(context.Context) ServiceEmailsOnPushArrayOutput
}

type ServiceEmailsOnPushArray []ServiceEmailsOnPushInput

func (ServiceEmailsOnPushArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEmailsOnPush)(nil)).Elem()
}

func (i ServiceEmailsOnPushArray) ToServiceEmailsOnPushArrayOutput() ServiceEmailsOnPushArrayOutput {
	return i.ToServiceEmailsOnPushArrayOutputWithContext(context.Background())
}

func (i ServiceEmailsOnPushArray) ToServiceEmailsOnPushArrayOutputWithContext(ctx context.Context) ServiceEmailsOnPushArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEmailsOnPushArrayOutput)
}

// ServiceEmailsOnPushMapInput is an input type that accepts ServiceEmailsOnPushMap and ServiceEmailsOnPushMapOutput values.
// You can construct a concrete instance of `ServiceEmailsOnPushMapInput` via:
//
//	ServiceEmailsOnPushMap{ "key": ServiceEmailsOnPushArgs{...} }
type ServiceEmailsOnPushMapInput interface {
	pulumi.Input

	ToServiceEmailsOnPushMapOutput() ServiceEmailsOnPushMapOutput
	ToServiceEmailsOnPushMapOutputWithContext(context.Context) ServiceEmailsOnPushMapOutput
}

type ServiceEmailsOnPushMap map[string]ServiceEmailsOnPushInput

func (ServiceEmailsOnPushMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEmailsOnPush)(nil)).Elem()
}

func (i ServiceEmailsOnPushMap) ToServiceEmailsOnPushMapOutput() ServiceEmailsOnPushMapOutput {
	return i.ToServiceEmailsOnPushMapOutputWithContext(context.Background())
}

func (i ServiceEmailsOnPushMap) ToServiceEmailsOnPushMapOutputWithContext(ctx context.Context) ServiceEmailsOnPushMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEmailsOnPushMapOutput)
}

type ServiceEmailsOnPushOutput struct{ *pulumi.OutputState }

func (ServiceEmailsOnPushOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEmailsOnPush)(nil)).Elem()
}

func (o ServiceEmailsOnPushOutput) ToServiceEmailsOnPushOutput() ServiceEmailsOnPushOutput {
	return o
}

func (o ServiceEmailsOnPushOutput) ToServiceEmailsOnPushOutputWithContext(ctx context.Context) ServiceEmailsOnPushOutput {
	return o
}

// Whether the integration is active.
func (o ServiceEmailsOnPushOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
func (o ServiceEmailsOnPushOutput) BranchesToBeNotified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringPtrOutput { return v.BranchesToBeNotified }).(pulumi.StringPtrOutput)
}

// The ISO8601 date/time that this integration was activated at in UTC.
func (o ServiceEmailsOnPushOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Disable code diffs.
func (o ServiceEmailsOnPushOutput) DisableDiffs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.BoolPtrOutput { return v.DisableDiffs }).(pulumi.BoolPtrOutput)
}

// ID or full-path of the project you want to activate integration on.
func (o ServiceEmailsOnPushOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Enable notifications for push events.
func (o ServiceEmailsOnPushOutput) PushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.BoolPtrOutput { return v.PushEvents }).(pulumi.BoolPtrOutput)
}

// Emails separated by whitespace.
func (o ServiceEmailsOnPushOutput) Recipients() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringOutput { return v.Recipients }).(pulumi.StringOutput)
}

// Send from committer.
func (o ServiceEmailsOnPushOutput) SendFromCommitterEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.BoolPtrOutput { return v.SendFromCommitterEmail }).(pulumi.BoolPtrOutput)
}

// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
func (o ServiceEmailsOnPushOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Enable notifications for tag push events.
func (o ServiceEmailsOnPushOutput) TagPushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.BoolPtrOutput { return v.TagPushEvents }).(pulumi.BoolPtrOutput)
}

// Title of the integration.
func (o ServiceEmailsOnPushOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The ISO8601 date/time that this integration was last updated at in UTC.
func (o ServiceEmailsOnPushOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEmailsOnPush) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ServiceEmailsOnPushArrayOutput struct{ *pulumi.OutputState }

func (ServiceEmailsOnPushArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceEmailsOnPush)(nil)).Elem()
}

func (o ServiceEmailsOnPushArrayOutput) ToServiceEmailsOnPushArrayOutput() ServiceEmailsOnPushArrayOutput {
	return o
}

func (o ServiceEmailsOnPushArrayOutput) ToServiceEmailsOnPushArrayOutputWithContext(ctx context.Context) ServiceEmailsOnPushArrayOutput {
	return o
}

func (o ServiceEmailsOnPushArrayOutput) Index(i pulumi.IntInput) ServiceEmailsOnPushOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceEmailsOnPush {
		return vs[0].([]*ServiceEmailsOnPush)[vs[1].(int)]
	}).(ServiceEmailsOnPushOutput)
}

type ServiceEmailsOnPushMapOutput struct{ *pulumi.OutputState }

func (ServiceEmailsOnPushMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceEmailsOnPush)(nil)).Elem()
}

func (o ServiceEmailsOnPushMapOutput) ToServiceEmailsOnPushMapOutput() ServiceEmailsOnPushMapOutput {
	return o
}

func (o ServiceEmailsOnPushMapOutput) ToServiceEmailsOnPushMapOutputWithContext(ctx context.Context) ServiceEmailsOnPushMapOutput {
	return o
}

func (o ServiceEmailsOnPushMapOutput) MapIndex(k pulumi.StringInput) ServiceEmailsOnPushOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceEmailsOnPush {
		return vs[0].(map[string]*ServiceEmailsOnPush)[vs[1].(string)]
	}).(ServiceEmailsOnPushOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEmailsOnPushInput)(nil)).Elem(), &ServiceEmailsOnPush{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEmailsOnPushArrayInput)(nil)).Elem(), ServiceEmailsOnPushArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEmailsOnPushMapInput)(nil)).Elem(), ServiceEmailsOnPushMap{})
	pulumi.RegisterOutputType(ServiceEmailsOnPushOutput{})
	pulumi.RegisterOutputType(ServiceEmailsOnPushArrayOutput{})
	pulumi.RegisterOutputType(ServiceEmailsOnPushMapOutput{})
}
