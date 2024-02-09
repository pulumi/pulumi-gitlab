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

// The `IntegrationEmailsOnPush` resource allows to manage the lifecycle of a project integration with Emails on Push Service.
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
//			_, err = gitlab.NewIntegrationEmailsOnPush(ctx, "emails", &gitlab.IntegrationEmailsOnPushArgs{
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
// You can import a gitlab_integration_emails_on_push state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush emails 1
// ```
type IntegrationEmailsOnPush struct {
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

// NewIntegrationEmailsOnPush registers a new resource with the given unique name, arguments, and options.
func NewIntegrationEmailsOnPush(ctx *pulumi.Context,
	name string, args *IntegrationEmailsOnPushArgs, opts ...pulumi.ResourceOption) (*IntegrationEmailsOnPush, error) {
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
	var resource IntegrationEmailsOnPush
	err := ctx.RegisterResource("gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationEmailsOnPush gets an existing IntegrationEmailsOnPush resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationEmailsOnPush(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationEmailsOnPushState, opts ...pulumi.ResourceOption) (*IntegrationEmailsOnPush, error) {
	var resource IntegrationEmailsOnPush
	err := ctx.ReadResource("gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationEmailsOnPush resources.
type integrationEmailsOnPushState struct {
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

type IntegrationEmailsOnPushState struct {
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

func (IntegrationEmailsOnPushState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationEmailsOnPushState)(nil)).Elem()
}

type integrationEmailsOnPushArgs struct {
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

// The set of arguments for constructing a IntegrationEmailsOnPush resource.
type IntegrationEmailsOnPushArgs struct {
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

func (IntegrationEmailsOnPushArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationEmailsOnPushArgs)(nil)).Elem()
}

type IntegrationEmailsOnPushInput interface {
	pulumi.Input

	ToIntegrationEmailsOnPushOutput() IntegrationEmailsOnPushOutput
	ToIntegrationEmailsOnPushOutputWithContext(ctx context.Context) IntegrationEmailsOnPushOutput
}

func (*IntegrationEmailsOnPush) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationEmailsOnPush)(nil)).Elem()
}

func (i *IntegrationEmailsOnPush) ToIntegrationEmailsOnPushOutput() IntegrationEmailsOnPushOutput {
	return i.ToIntegrationEmailsOnPushOutputWithContext(context.Background())
}

func (i *IntegrationEmailsOnPush) ToIntegrationEmailsOnPushOutputWithContext(ctx context.Context) IntegrationEmailsOnPushOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationEmailsOnPushOutput)
}

// IntegrationEmailsOnPushArrayInput is an input type that accepts IntegrationEmailsOnPushArray and IntegrationEmailsOnPushArrayOutput values.
// You can construct a concrete instance of `IntegrationEmailsOnPushArrayInput` via:
//
//	IntegrationEmailsOnPushArray{ IntegrationEmailsOnPushArgs{...} }
type IntegrationEmailsOnPushArrayInput interface {
	pulumi.Input

	ToIntegrationEmailsOnPushArrayOutput() IntegrationEmailsOnPushArrayOutput
	ToIntegrationEmailsOnPushArrayOutputWithContext(context.Context) IntegrationEmailsOnPushArrayOutput
}

type IntegrationEmailsOnPushArray []IntegrationEmailsOnPushInput

func (IntegrationEmailsOnPushArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationEmailsOnPush)(nil)).Elem()
}

func (i IntegrationEmailsOnPushArray) ToIntegrationEmailsOnPushArrayOutput() IntegrationEmailsOnPushArrayOutput {
	return i.ToIntegrationEmailsOnPushArrayOutputWithContext(context.Background())
}

func (i IntegrationEmailsOnPushArray) ToIntegrationEmailsOnPushArrayOutputWithContext(ctx context.Context) IntegrationEmailsOnPushArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationEmailsOnPushArrayOutput)
}

// IntegrationEmailsOnPushMapInput is an input type that accepts IntegrationEmailsOnPushMap and IntegrationEmailsOnPushMapOutput values.
// You can construct a concrete instance of `IntegrationEmailsOnPushMapInput` via:
//
//	IntegrationEmailsOnPushMap{ "key": IntegrationEmailsOnPushArgs{...} }
type IntegrationEmailsOnPushMapInput interface {
	pulumi.Input

	ToIntegrationEmailsOnPushMapOutput() IntegrationEmailsOnPushMapOutput
	ToIntegrationEmailsOnPushMapOutputWithContext(context.Context) IntegrationEmailsOnPushMapOutput
}

type IntegrationEmailsOnPushMap map[string]IntegrationEmailsOnPushInput

func (IntegrationEmailsOnPushMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationEmailsOnPush)(nil)).Elem()
}

func (i IntegrationEmailsOnPushMap) ToIntegrationEmailsOnPushMapOutput() IntegrationEmailsOnPushMapOutput {
	return i.ToIntegrationEmailsOnPushMapOutputWithContext(context.Background())
}

func (i IntegrationEmailsOnPushMap) ToIntegrationEmailsOnPushMapOutputWithContext(ctx context.Context) IntegrationEmailsOnPushMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationEmailsOnPushMapOutput)
}

type IntegrationEmailsOnPushOutput struct{ *pulumi.OutputState }

func (IntegrationEmailsOnPushOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationEmailsOnPush)(nil)).Elem()
}

func (o IntegrationEmailsOnPushOutput) ToIntegrationEmailsOnPushOutput() IntegrationEmailsOnPushOutput {
	return o
}

func (o IntegrationEmailsOnPushOutput) ToIntegrationEmailsOnPushOutputWithContext(ctx context.Context) IntegrationEmailsOnPushOutput {
	return o
}

// Whether the integration is active.
func (o IntegrationEmailsOnPushOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `defaultAndProtected`. Notifications are always fired for tag pushes.
func (o IntegrationEmailsOnPushOutput) BranchesToBeNotified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringPtrOutput { return v.BranchesToBeNotified }).(pulumi.StringPtrOutput)
}

// The ISO8601 date/time that this integration was activated at in UTC.
func (o IntegrationEmailsOnPushOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Disable code diffs.
func (o IntegrationEmailsOnPushOutput) DisableDiffs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.BoolPtrOutput { return v.DisableDiffs }).(pulumi.BoolPtrOutput)
}

// ID or full-path of the project you want to activate integration on.
func (o IntegrationEmailsOnPushOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Enable notifications for push events.
func (o IntegrationEmailsOnPushOutput) PushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.BoolPtrOutput { return v.PushEvents }).(pulumi.BoolPtrOutput)
}

// Emails separated by whitespace.
func (o IntegrationEmailsOnPushOutput) Recipients() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringOutput { return v.Recipients }).(pulumi.StringOutput)
}

// Send from committer.
func (o IntegrationEmailsOnPushOutput) SendFromCommitterEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.BoolPtrOutput { return v.SendFromCommitterEmail }).(pulumi.BoolPtrOutput)
}

// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
func (o IntegrationEmailsOnPushOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Enable notifications for tag push events.
func (o IntegrationEmailsOnPushOutput) TagPushEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.BoolPtrOutput { return v.TagPushEvents }).(pulumi.BoolPtrOutput)
}

// Title of the integration.
func (o IntegrationEmailsOnPushOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The ISO8601 date/time that this integration was last updated at in UTC.
func (o IntegrationEmailsOnPushOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationEmailsOnPush) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IntegrationEmailsOnPushArrayOutput struct{ *pulumi.OutputState }

func (IntegrationEmailsOnPushArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationEmailsOnPush)(nil)).Elem()
}

func (o IntegrationEmailsOnPushArrayOutput) ToIntegrationEmailsOnPushArrayOutput() IntegrationEmailsOnPushArrayOutput {
	return o
}

func (o IntegrationEmailsOnPushArrayOutput) ToIntegrationEmailsOnPushArrayOutputWithContext(ctx context.Context) IntegrationEmailsOnPushArrayOutput {
	return o
}

func (o IntegrationEmailsOnPushArrayOutput) Index(i pulumi.IntInput) IntegrationEmailsOnPushOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationEmailsOnPush {
		return vs[0].([]*IntegrationEmailsOnPush)[vs[1].(int)]
	}).(IntegrationEmailsOnPushOutput)
}

type IntegrationEmailsOnPushMapOutput struct{ *pulumi.OutputState }

func (IntegrationEmailsOnPushMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationEmailsOnPush)(nil)).Elem()
}

func (o IntegrationEmailsOnPushMapOutput) ToIntegrationEmailsOnPushMapOutput() IntegrationEmailsOnPushMapOutput {
	return o
}

func (o IntegrationEmailsOnPushMapOutput) ToIntegrationEmailsOnPushMapOutputWithContext(ctx context.Context) IntegrationEmailsOnPushMapOutput {
	return o
}

func (o IntegrationEmailsOnPushMapOutput) MapIndex(k pulumi.StringInput) IntegrationEmailsOnPushOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationEmailsOnPush {
		return vs[0].(map[string]*IntegrationEmailsOnPush)[vs[1].(string)]
	}).(IntegrationEmailsOnPushOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationEmailsOnPushInput)(nil)).Elem(), &IntegrationEmailsOnPush{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationEmailsOnPushArrayInput)(nil)).Elem(), IntegrationEmailsOnPushArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationEmailsOnPushMapInput)(nil)).Elem(), IntegrationEmailsOnPushMap{})
	pulumi.RegisterOutputType(IntegrationEmailsOnPushOutput{})
	pulumi.RegisterOutputType(IntegrationEmailsOnPushArrayOutput{})
	pulumi.RegisterOutputType(IntegrationEmailsOnPushMapOutput{})
}
