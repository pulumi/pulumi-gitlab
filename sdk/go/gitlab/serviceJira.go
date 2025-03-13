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

// The `ServiceJira` resource allows to manage the lifecycle of a project integration with Jira.
//
// > This resource is deprecated. use `IntegrationJira`instead!
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/integrations/#jira)
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
//			_, err = gitlab.NewServiceJira(ctx, "jira", &gitlab.ServiceJiraArgs{
//				Project:  awesomeProject.ID(),
//				Url:      pulumi.String("https://jira.example.com"),
//				Username: pulumi.String("user"),
//				Password: pulumi.String("mypass"),
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
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_service_jira`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_service_jira.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// You can import a gitlab_service_jira state using the project ID, e.g.
//
// ```sh
// $ pulumi import gitlab:index/serviceJira:ServiceJira jira 1
// ```
type ServiceJira struct {
	pulumi.CustomResourceState

	// Whether the integration is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	ApiUrl pulumi.StringOutput `pulumi:"apiUrl"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled pulumi.BoolOutput `pulumi:"commentOnEventEnabled"`
	// Enable notifications for commit events
	CommitEvents pulumi.BoolOutput `pulumi:"commitEvents"`
	// Create time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Enable viewing Jira issues in GitLab.
	IssuesEnabled pulumi.BoolPtrOutput `pulumi:"issuesEnabled"`
	// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
	JiraAuthType pulumi.IntPtrOutput `pulumi:"jiraAuthType"`
	// Prefix to match Jira issue keys.
	JiraIssuePrefix pulumi.StringPtrOutput `pulumi:"jiraIssuePrefix"`
	// Regular expression to match Jira issue keys.
	JiraIssueRegex               pulumi.StringPtrOutput `pulumi:"jiraIssueRegex"`
	JiraIssueTransitionAutomatic pulumi.BoolPtrOutput   `pulumi:"jiraIssueTransitionAutomatic"`
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId pulumi.StringPtrOutput `pulumi:"jiraIssueTransitionId"`
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolOutput `pulumi:"mergeRequestsEvents"`
	// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
	Password pulumi.StringOutput `pulumi:"password"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
	ProjectKeys pulumi.StringArrayOutput `pulumi:"projectKeys"`
	// Title.
	Title pulumi.StringOutput `pulumi:"title"`
	// Update time.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url pulumi.StringOutput `pulumi:"url"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings pulumi.BoolPtrOutput `pulumi:"useInheritedSettings"`
	// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewServiceJira registers a new resource with the given unique name, arguments, and options.
func NewServiceJira(ctx *pulumi.Context,
	name string, args *ServiceJiraArgs, opts ...pulumi.ResourceOption) (*ServiceJira, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceJira
	err := ctx.RegisterResource("gitlab:index/serviceJira:ServiceJira", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceJira gets an existing ServiceJira resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceJira(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceJiraState, opts ...pulumi.ResourceOption) (*ServiceJira, error) {
	var resource ServiceJira
	err := ctx.ReadResource("gitlab:index/serviceJira:ServiceJira", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceJira resources.
type serviceJiraState struct {
	// Whether the integration is active.
	Active *bool `pulumi:"active"`
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	ApiUrl *string `pulumi:"apiUrl"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled *bool `pulumi:"commentOnEventEnabled"`
	// Enable notifications for commit events
	CommitEvents *bool `pulumi:"commitEvents"`
	// Create time.
	CreatedAt *string `pulumi:"createdAt"`
	// Enable viewing Jira issues in GitLab.
	IssuesEnabled *bool `pulumi:"issuesEnabled"`
	// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
	JiraAuthType *int `pulumi:"jiraAuthType"`
	// Prefix to match Jira issue keys.
	JiraIssuePrefix *string `pulumi:"jiraIssuePrefix"`
	// Regular expression to match Jira issue keys.
	JiraIssueRegex               *string `pulumi:"jiraIssueRegex"`
	JiraIssueTransitionAutomatic *bool   `pulumi:"jiraIssueTransitionAutomatic"`
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId *string `pulumi:"jiraIssueTransitionId"`
	// Enable notifications for merge request events
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
	Password *string `pulumi:"password"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey *string `pulumi:"projectKey"`
	// Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
	ProjectKeys []string `pulumi:"projectKeys"`
	// Title.
	Title *string `pulumi:"title"`
	// Update time.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url *string `pulumi:"url"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings *bool `pulumi:"useInheritedSettings"`
	// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
	Username *string `pulumi:"username"`
}

type ServiceJiraState struct {
	// Whether the integration is active.
	Active pulumi.BoolPtrInput
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	ApiUrl pulumi.StringPtrInput
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled pulumi.BoolPtrInput
	// Enable notifications for commit events
	CommitEvents pulumi.BoolPtrInput
	// Create time.
	CreatedAt pulumi.StringPtrInput
	// Enable viewing Jira issues in GitLab.
	IssuesEnabled pulumi.BoolPtrInput
	// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
	JiraAuthType pulumi.IntPtrInput
	// Prefix to match Jira issue keys.
	JiraIssuePrefix pulumi.StringPtrInput
	// Regular expression to match Jira issue keys.
	JiraIssueRegex               pulumi.StringPtrInput
	JiraIssueTransitionAutomatic pulumi.BoolPtrInput
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId pulumi.StringPtrInput
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrInput
	// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
	Password pulumi.StringPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey pulumi.StringPtrInput
	// Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
	ProjectKeys pulumi.StringArrayInput
	// Title.
	Title pulumi.StringPtrInput
	// Update time.
	UpdatedAt pulumi.StringPtrInput
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url pulumi.StringPtrInput
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings pulumi.BoolPtrInput
	// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
	Username pulumi.StringPtrInput
}

func (ServiceJiraState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceJiraState)(nil)).Elem()
}

type serviceJiraArgs struct {
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	ApiUrl *string `pulumi:"apiUrl"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled *bool `pulumi:"commentOnEventEnabled"`
	// Enable notifications for commit events
	CommitEvents *bool `pulumi:"commitEvents"`
	// Enable viewing Jira issues in GitLab.
	IssuesEnabled *bool `pulumi:"issuesEnabled"`
	// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
	JiraAuthType *int `pulumi:"jiraAuthType"`
	// Prefix to match Jira issue keys.
	JiraIssuePrefix *string `pulumi:"jiraIssuePrefix"`
	// Regular expression to match Jira issue keys.
	JiraIssueRegex               *string `pulumi:"jiraIssueRegex"`
	JiraIssueTransitionAutomatic *bool   `pulumi:"jiraIssueTransitionAutomatic"`
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId *string `pulumi:"jiraIssueTransitionId"`
	// Enable notifications for merge request events
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
	Password string `pulumi:"password"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey *string `pulumi:"projectKey"`
	// Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
	ProjectKeys []string `pulumi:"projectKeys"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url string `pulumi:"url"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings *bool `pulumi:"useInheritedSettings"`
	// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceJira resource.
type ServiceJiraArgs struct {
	// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
	ApiUrl pulumi.StringPtrInput
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled pulumi.BoolPtrInput
	// Enable notifications for commit events
	CommitEvents pulumi.BoolPtrInput
	// Enable viewing Jira issues in GitLab.
	IssuesEnabled pulumi.BoolPtrInput
	// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
	JiraAuthType pulumi.IntPtrInput
	// Prefix to match Jira issue keys.
	JiraIssuePrefix pulumi.StringPtrInput
	// Regular expression to match Jira issue keys.
	JiraIssueRegex               pulumi.StringPtrInput
	JiraIssueTransitionAutomatic pulumi.BoolPtrInput
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId pulumi.StringPtrInput
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrInput
	// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
	Password pulumi.StringInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey pulumi.StringPtrInput
	// Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
	ProjectKeys pulumi.StringArrayInput
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url pulumi.StringInput
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings pulumi.BoolPtrInput
	// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
	Username pulumi.StringPtrInput
}

func (ServiceJiraArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceJiraArgs)(nil)).Elem()
}

type ServiceJiraInput interface {
	pulumi.Input

	ToServiceJiraOutput() ServiceJiraOutput
	ToServiceJiraOutputWithContext(ctx context.Context) ServiceJiraOutput
}

func (*ServiceJira) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceJira)(nil)).Elem()
}

func (i *ServiceJira) ToServiceJiraOutput() ServiceJiraOutput {
	return i.ToServiceJiraOutputWithContext(context.Background())
}

func (i *ServiceJira) ToServiceJiraOutputWithContext(ctx context.Context) ServiceJiraOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceJiraOutput)
}

// ServiceJiraArrayInput is an input type that accepts ServiceJiraArray and ServiceJiraArrayOutput values.
// You can construct a concrete instance of `ServiceJiraArrayInput` via:
//
//	ServiceJiraArray{ ServiceJiraArgs{...} }
type ServiceJiraArrayInput interface {
	pulumi.Input

	ToServiceJiraArrayOutput() ServiceJiraArrayOutput
	ToServiceJiraArrayOutputWithContext(context.Context) ServiceJiraArrayOutput
}

type ServiceJiraArray []ServiceJiraInput

func (ServiceJiraArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceJira)(nil)).Elem()
}

func (i ServiceJiraArray) ToServiceJiraArrayOutput() ServiceJiraArrayOutput {
	return i.ToServiceJiraArrayOutputWithContext(context.Background())
}

func (i ServiceJiraArray) ToServiceJiraArrayOutputWithContext(ctx context.Context) ServiceJiraArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceJiraArrayOutput)
}

// ServiceJiraMapInput is an input type that accepts ServiceJiraMap and ServiceJiraMapOutput values.
// You can construct a concrete instance of `ServiceJiraMapInput` via:
//
//	ServiceJiraMap{ "key": ServiceJiraArgs{...} }
type ServiceJiraMapInput interface {
	pulumi.Input

	ToServiceJiraMapOutput() ServiceJiraMapOutput
	ToServiceJiraMapOutputWithContext(context.Context) ServiceJiraMapOutput
}

type ServiceJiraMap map[string]ServiceJiraInput

func (ServiceJiraMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceJira)(nil)).Elem()
}

func (i ServiceJiraMap) ToServiceJiraMapOutput() ServiceJiraMapOutput {
	return i.ToServiceJiraMapOutputWithContext(context.Background())
}

func (i ServiceJiraMap) ToServiceJiraMapOutputWithContext(ctx context.Context) ServiceJiraMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceJiraMapOutput)
}

type ServiceJiraOutput struct{ *pulumi.OutputState }

func (ServiceJiraOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceJira)(nil)).Elem()
}

func (o ServiceJiraOutput) ToServiceJiraOutput() ServiceJiraOutput {
	return o
}

func (o ServiceJiraOutput) ToServiceJiraOutputWithContext(ctx context.Context) ServiceJiraOutput {
	return o
}

// Whether the integration is active.
func (o ServiceJiraOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// The base URL to the Jira instance API. Web URL value is used if not set. For example, https://jira-api.example.com.
func (o ServiceJiraOutput) ApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.ApiUrl }).(pulumi.StringOutput)
}

// Enable comments inside Jira issues on each GitLab event (commit / merge request)
func (o ServiceJiraOutput) CommentOnEventEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolOutput { return v.CommentOnEventEnabled }).(pulumi.BoolOutput)
}

// Enable notifications for commit events
func (o ServiceJiraOutput) CommitEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolOutput { return v.CommitEvents }).(pulumi.BoolOutput)
}

// Create time.
func (o ServiceJiraOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Enable viewing Jira issues in GitLab.
func (o ServiceJiraOutput) IssuesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolPtrOutput { return v.IssuesEnabled }).(pulumi.BoolPtrOutput)
}

// The authentication method to be used with Jira. 0 means Basic Authentication. 1 means Jira personal access token. Defaults to 0.
func (o ServiceJiraOutput) JiraAuthType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.IntPtrOutput { return v.JiraAuthType }).(pulumi.IntPtrOutput)
}

// Prefix to match Jira issue keys.
func (o ServiceJiraOutput) JiraIssuePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringPtrOutput { return v.JiraIssuePrefix }).(pulumi.StringPtrOutput)
}

// Regular expression to match Jira issue keys.
func (o ServiceJiraOutput) JiraIssueRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringPtrOutput { return v.JiraIssueRegex }).(pulumi.StringPtrOutput)
}

func (o ServiceJiraOutput) JiraIssueTransitionAutomatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolPtrOutput { return v.JiraIssueTransitionAutomatic }).(pulumi.BoolPtrOutput)
}

// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
func (o ServiceJiraOutput) JiraIssueTransitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringPtrOutput { return v.JiraIssueTransitionId }).(pulumi.StringPtrOutput)
}

// Enable notifications for merge request events
func (o ServiceJiraOutput) MergeRequestsEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolOutput { return v.MergeRequestsEvents }).(pulumi.BoolOutput)
}

// The Jira API token, password, or personal access token to be used with Jira. When your authentication method is basic (jira*auth*type is 0), use an API token for Jira Cloud or a password for Jira Data Center or Jira Server. When your authentication method is a Jira personal access token (jira*auth*type is 1), use the personal access token.
func (o ServiceJiraOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// ID of the project you want to activate integration on.
func (o ServiceJiraOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
func (o ServiceJiraOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Keys of Jira projects. When issuesEnabled is true, this setting specifies which Jira projects to view issues from in GitLab.
func (o ServiceJiraOutput) ProjectKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringArrayOutput { return v.ProjectKeys }).(pulumi.StringArrayOutput)
}

// Title.
func (o ServiceJiraOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// Update time.
func (o ServiceJiraOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
func (o ServiceJiraOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Indicates whether or not to inherit default settings. Defaults to false.
func (o ServiceJiraOutput) UseInheritedSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.BoolPtrOutput { return v.UseInheritedSettings }).(pulumi.BoolPtrOutput)
}

// The email or username to be used with Jira. For Jira Cloud use an email, for Jira Data Center and Jira Server use a username. Required when using Basic authentication (jira*auth*type is 0).
func (o ServiceJiraOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceJira) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type ServiceJiraArrayOutput struct{ *pulumi.OutputState }

func (ServiceJiraArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceJira)(nil)).Elem()
}

func (o ServiceJiraArrayOutput) ToServiceJiraArrayOutput() ServiceJiraArrayOutput {
	return o
}

func (o ServiceJiraArrayOutput) ToServiceJiraArrayOutputWithContext(ctx context.Context) ServiceJiraArrayOutput {
	return o
}

func (o ServiceJiraArrayOutput) Index(i pulumi.IntInput) ServiceJiraOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceJira {
		return vs[0].([]*ServiceJira)[vs[1].(int)]
	}).(ServiceJiraOutput)
}

type ServiceJiraMapOutput struct{ *pulumi.OutputState }

func (ServiceJiraMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceJira)(nil)).Elem()
}

func (o ServiceJiraMapOutput) ToServiceJiraMapOutput() ServiceJiraMapOutput {
	return o
}

func (o ServiceJiraMapOutput) ToServiceJiraMapOutputWithContext(ctx context.Context) ServiceJiraMapOutput {
	return o
}

func (o ServiceJiraMapOutput) MapIndex(k pulumi.StringInput) ServiceJiraOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceJira {
		return vs[0].(map[string]*ServiceJira)[vs[1].(string)]
	}).(ServiceJiraOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceJiraInput)(nil)).Elem(), &ServiceJira{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceJiraArrayInput)(nil)).Elem(), ServiceJiraArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceJiraMapInput)(nil)).Elem(), ServiceJiraMap{})
	pulumi.RegisterOutputType(ServiceJiraOutput{})
	pulumi.RegisterOutputType(ServiceJiraArrayOutput{})
	pulumi.RegisterOutputType(ServiceJiraMapOutput{})
}
