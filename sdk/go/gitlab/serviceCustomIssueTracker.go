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

// The `ServiceCustomIssueTracker` resource allows to manage the lifecycle of a project integration with Custom Issue Tracker.
//
// > This resource is deprecated. use `IntegrationCustomIssueTracker`instead!
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#custom-issue-tracker)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = gitlab.NewServiceCustomIssueTracker(ctx, "tracker", &gitlab.ServiceCustomIssueTrackerArgs{
//				Project:    awesomeProject.ID(),
//				ProjectUrl: pulumi.String("https://customtracker.com/issues"),
//				IssuesUrl:  pulumi.String("https://customtracker.com/TEST-:id"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// You can import a gitlab_service_custom_issue_tracker state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker tracker 1
// ```
type ServiceCustomIssueTracker struct {
	pulumi.CustomResourceState

	// Whether the integration is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The URL to view an issue in the external issue tracker. Must contain :id.
	IssuesUrl pulumi.StringOutput `pulumi:"issuesUrl"`
	// The ID or full path of the project for the custom issue tracker.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL to the project in the external issue tracker.
	ProjectUrl pulumi.StringOutput `pulumi:"projectUrl"`
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewServiceCustomIssueTracker registers a new resource with the given unique name, arguments, and options.
func NewServiceCustomIssueTracker(ctx *pulumi.Context,
	name string, args *ServiceCustomIssueTrackerArgs, opts ...pulumi.ResourceOption) (*ServiceCustomIssueTracker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IssuesUrl == nil {
		return nil, errors.New("invalid value for required argument 'IssuesUrl'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ProjectUrl == nil {
		return nil, errors.New("invalid value for required argument 'ProjectUrl'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceCustomIssueTracker
	err := ctx.RegisterResource("gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceCustomIssueTracker gets an existing ServiceCustomIssueTracker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceCustomIssueTracker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceCustomIssueTrackerState, opts ...pulumi.ResourceOption) (*ServiceCustomIssueTracker, error) {
	var resource ServiceCustomIssueTracker
	err := ctx.ReadResource("gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceCustomIssueTracker resources.
type serviceCustomIssueTrackerState struct {
	// Whether the integration is active.
	Active *bool `pulumi:"active"`
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// The URL to view an issue in the external issue tracker. Must contain :id.
	IssuesUrl *string `pulumi:"issuesUrl"`
	// The ID or full path of the project for the custom issue tracker.
	Project *string `pulumi:"project"`
	// The URL to the project in the external issue tracker.
	ProjectUrl *string `pulumi:"projectUrl"`
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug *string `pulumi:"slug"`
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ServiceCustomIssueTrackerState struct {
	// Whether the integration is active.
	Active pulumi.BoolPtrInput
	// The ISO8601 date/time that this integration was activated at in UTC.
	CreatedAt pulumi.StringPtrInput
	// The URL to view an issue in the external issue tracker. Must contain :id.
	IssuesUrl pulumi.StringPtrInput
	// The ID or full path of the project for the custom issue tracker.
	Project pulumi.StringPtrInput
	// The URL to the project in the external issue tracker.
	ProjectUrl pulumi.StringPtrInput
	// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
	Slug pulumi.StringPtrInput
	// The ISO8601 date/time that this integration was last updated at in UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (ServiceCustomIssueTrackerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceCustomIssueTrackerState)(nil)).Elem()
}

type serviceCustomIssueTrackerArgs struct {
	// The URL to view an issue in the external issue tracker. Must contain :id.
	IssuesUrl string `pulumi:"issuesUrl"`
	// The ID or full path of the project for the custom issue tracker.
	Project string `pulumi:"project"`
	// The URL to the project in the external issue tracker.
	ProjectUrl string `pulumi:"projectUrl"`
}

// The set of arguments for constructing a ServiceCustomIssueTracker resource.
type ServiceCustomIssueTrackerArgs struct {
	// The URL to view an issue in the external issue tracker. Must contain :id.
	IssuesUrl pulumi.StringInput
	// The ID or full path of the project for the custom issue tracker.
	Project pulumi.StringInput
	// The URL to the project in the external issue tracker.
	ProjectUrl pulumi.StringInput
}

func (ServiceCustomIssueTrackerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceCustomIssueTrackerArgs)(nil)).Elem()
}

type ServiceCustomIssueTrackerInput interface {
	pulumi.Input

	ToServiceCustomIssueTrackerOutput() ServiceCustomIssueTrackerOutput
	ToServiceCustomIssueTrackerOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerOutput
}

func (*ServiceCustomIssueTracker) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCustomIssueTracker)(nil)).Elem()
}

func (i *ServiceCustomIssueTracker) ToServiceCustomIssueTrackerOutput() ServiceCustomIssueTrackerOutput {
	return i.ToServiceCustomIssueTrackerOutputWithContext(context.Background())
}

func (i *ServiceCustomIssueTracker) ToServiceCustomIssueTrackerOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCustomIssueTrackerOutput)
}

// ServiceCustomIssueTrackerArrayInput is an input type that accepts ServiceCustomIssueTrackerArray and ServiceCustomIssueTrackerArrayOutput values.
// You can construct a concrete instance of `ServiceCustomIssueTrackerArrayInput` via:
//
//	ServiceCustomIssueTrackerArray{ ServiceCustomIssueTrackerArgs{...} }
type ServiceCustomIssueTrackerArrayInput interface {
	pulumi.Input

	ToServiceCustomIssueTrackerArrayOutput() ServiceCustomIssueTrackerArrayOutput
	ToServiceCustomIssueTrackerArrayOutputWithContext(context.Context) ServiceCustomIssueTrackerArrayOutput
}

type ServiceCustomIssueTrackerArray []ServiceCustomIssueTrackerInput

func (ServiceCustomIssueTrackerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceCustomIssueTracker)(nil)).Elem()
}

func (i ServiceCustomIssueTrackerArray) ToServiceCustomIssueTrackerArrayOutput() ServiceCustomIssueTrackerArrayOutput {
	return i.ToServiceCustomIssueTrackerArrayOutputWithContext(context.Background())
}

func (i ServiceCustomIssueTrackerArray) ToServiceCustomIssueTrackerArrayOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCustomIssueTrackerArrayOutput)
}

// ServiceCustomIssueTrackerMapInput is an input type that accepts ServiceCustomIssueTrackerMap and ServiceCustomIssueTrackerMapOutput values.
// You can construct a concrete instance of `ServiceCustomIssueTrackerMapInput` via:
//
//	ServiceCustomIssueTrackerMap{ "key": ServiceCustomIssueTrackerArgs{...} }
type ServiceCustomIssueTrackerMapInput interface {
	pulumi.Input

	ToServiceCustomIssueTrackerMapOutput() ServiceCustomIssueTrackerMapOutput
	ToServiceCustomIssueTrackerMapOutputWithContext(context.Context) ServiceCustomIssueTrackerMapOutput
}

type ServiceCustomIssueTrackerMap map[string]ServiceCustomIssueTrackerInput

func (ServiceCustomIssueTrackerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceCustomIssueTracker)(nil)).Elem()
}

func (i ServiceCustomIssueTrackerMap) ToServiceCustomIssueTrackerMapOutput() ServiceCustomIssueTrackerMapOutput {
	return i.ToServiceCustomIssueTrackerMapOutputWithContext(context.Background())
}

func (i ServiceCustomIssueTrackerMap) ToServiceCustomIssueTrackerMapOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCustomIssueTrackerMapOutput)
}

type ServiceCustomIssueTrackerOutput struct{ *pulumi.OutputState }

func (ServiceCustomIssueTrackerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCustomIssueTracker)(nil)).Elem()
}

func (o ServiceCustomIssueTrackerOutput) ToServiceCustomIssueTrackerOutput() ServiceCustomIssueTrackerOutput {
	return o
}

func (o ServiceCustomIssueTrackerOutput) ToServiceCustomIssueTrackerOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerOutput {
	return o
}

// Whether the integration is active.
func (o ServiceCustomIssueTrackerOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// The ISO8601 date/time that this integration was activated at in UTC.
func (o ServiceCustomIssueTrackerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The URL to view an issue in the external issue tracker. Must contain :id.
func (o ServiceCustomIssueTrackerOutput) IssuesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.StringOutput { return v.IssuesUrl }).(pulumi.StringOutput)
}

// The ID or full path of the project for the custom issue tracker.
func (o ServiceCustomIssueTrackerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL to the project in the external issue tracker.
func (o ServiceCustomIssueTrackerOutput) ProjectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.StringOutput { return v.ProjectUrl }).(pulumi.StringOutput)
}

// The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
func (o ServiceCustomIssueTrackerOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// The ISO8601 date/time that this integration was last updated at in UTC.
func (o ServiceCustomIssueTrackerOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceCustomIssueTracker) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ServiceCustomIssueTrackerArrayOutput struct{ *pulumi.OutputState }

func (ServiceCustomIssueTrackerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceCustomIssueTracker)(nil)).Elem()
}

func (o ServiceCustomIssueTrackerArrayOutput) ToServiceCustomIssueTrackerArrayOutput() ServiceCustomIssueTrackerArrayOutput {
	return o
}

func (o ServiceCustomIssueTrackerArrayOutput) ToServiceCustomIssueTrackerArrayOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerArrayOutput {
	return o
}

func (o ServiceCustomIssueTrackerArrayOutput) Index(i pulumi.IntInput) ServiceCustomIssueTrackerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceCustomIssueTracker {
		return vs[0].([]*ServiceCustomIssueTracker)[vs[1].(int)]
	}).(ServiceCustomIssueTrackerOutput)
}

type ServiceCustomIssueTrackerMapOutput struct{ *pulumi.OutputState }

func (ServiceCustomIssueTrackerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceCustomIssueTracker)(nil)).Elem()
}

func (o ServiceCustomIssueTrackerMapOutput) ToServiceCustomIssueTrackerMapOutput() ServiceCustomIssueTrackerMapOutput {
	return o
}

func (o ServiceCustomIssueTrackerMapOutput) ToServiceCustomIssueTrackerMapOutputWithContext(ctx context.Context) ServiceCustomIssueTrackerMapOutput {
	return o
}

func (o ServiceCustomIssueTrackerMapOutput) MapIndex(k pulumi.StringInput) ServiceCustomIssueTrackerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceCustomIssueTracker {
		return vs[0].(map[string]*ServiceCustomIssueTracker)[vs[1].(string)]
	}).(ServiceCustomIssueTrackerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCustomIssueTrackerInput)(nil)).Elem(), &ServiceCustomIssueTracker{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCustomIssueTrackerArrayInput)(nil)).Elem(), ServiceCustomIssueTrackerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCustomIssueTrackerMapInput)(nil)).Elem(), ServiceCustomIssueTrackerMap{})
	pulumi.RegisterOutputType(ServiceCustomIssueTrackerOutput{})
	pulumi.RegisterOutputType(ServiceCustomIssueTrackerArrayOutput{})
	pulumi.RegisterOutputType(ServiceCustomIssueTrackerMapOutput{})
}
