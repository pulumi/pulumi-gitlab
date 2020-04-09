// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to manage Jira integration.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/service_jira.html.markdown.
type ServiceJira struct {
	pulumi.CustomResourceState

	Active pulumi.BoolOutput `pulumi:"active"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled pulumi.BoolOutput `pulumi:"commentOnEventEnabled"`
	// Enable notifications for commit events
	CommitEvents pulumi.BoolOutput   `pulumi:"commitEvents"`
	CreatedAt    pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId pulumi.StringPtrOutput `pulumi:"jiraIssueTransitionId"`
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolOutput `pulumi:"mergeRequestsEvents"`
	// The password of the user created to be used with GitLab/JIRA.
	Password pulumi.StringOutput `pulumi:"password"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	Title      pulumi.StringOutput    `pulumi:"title"`
	UpdatedAt  pulumi.StringOutput    `pulumi:"updatedAt"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url pulumi.StringOutput `pulumi:"url"`
	// The username of the user created to be used with GitLab/JIRA.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewServiceJira registers a new resource with the given unique name, arguments, and options.
func NewServiceJira(ctx *pulumi.Context,
	name string, args *ServiceJiraArgs, opts ...pulumi.ResourceOption) (*ServiceJira, error) {
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &ServiceJiraArgs{}
	}
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
	Active *bool `pulumi:"active"`
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled *bool `pulumi:"commentOnEventEnabled"`
	// Enable notifications for commit events
	CommitEvents *bool   `pulumi:"commitEvents"`
	CreatedAt    *string `pulumi:"createdAt"`
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId *string `pulumi:"jiraIssueTransitionId"`
	// Enable notifications for merge request events
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// The password of the user created to be used with GitLab/JIRA.
	Password *string `pulumi:"password"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey *string `pulumi:"projectKey"`
	Title      *string `pulumi:"title"`
	UpdatedAt  *string `pulumi:"updatedAt"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url *string `pulumi:"url"`
	// The username of the user created to be used with GitLab/JIRA.
	Username *string `pulumi:"username"`
}

type ServiceJiraState struct {
	Active pulumi.BoolPtrInput
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled pulumi.BoolPtrInput
	// Enable notifications for commit events
	CommitEvents pulumi.BoolPtrInput
	CreatedAt    pulumi.StringPtrInput
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId pulumi.StringPtrInput
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrInput
	// The password of the user created to be used with GitLab/JIRA.
	Password pulumi.StringPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey pulumi.StringPtrInput
	Title      pulumi.StringPtrInput
	UpdatedAt  pulumi.StringPtrInput
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url pulumi.StringPtrInput
	// The username of the user created to be used with GitLab/JIRA.
	Username pulumi.StringPtrInput
}

func (ServiceJiraState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceJiraState)(nil)).Elem()
}

type serviceJiraArgs struct {
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled *bool `pulumi:"commentOnEventEnabled"`
	// Enable notifications for commit events
	CommitEvents *bool `pulumi:"commitEvents"`
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId *string `pulumi:"jiraIssueTransitionId"`
	// Enable notifications for merge request events
	MergeRequestsEvents *bool `pulumi:"mergeRequestsEvents"`
	// The password of the user created to be used with GitLab/JIRA.
	Password string `pulumi:"password"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey *string `pulumi:"projectKey"`
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url string `pulumi:"url"`
	// The username of the user created to be used with GitLab/JIRA.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ServiceJira resource.
type ServiceJiraArgs struct {
	// Enable comments inside Jira issues on each GitLab event (commit / merge request)
	CommentOnEventEnabled pulumi.BoolPtrInput
	// Enable notifications for commit events
	CommitEvents pulumi.BoolPtrInput
	// The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.
	JiraIssueTransitionId pulumi.StringPtrInput
	// Enable notifications for merge request events
	MergeRequestsEvents pulumi.BoolPtrInput
	// The password of the user created to be used with GitLab/JIRA.
	Password pulumi.StringInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// The short identifier for your JIRA project, all uppercase, e.g., PROJ.
	ProjectKey pulumi.StringPtrInput
	// The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.
	Url pulumi.StringInput
	// The username of the user created to be used with GitLab/JIRA.
	Username pulumi.StringInput
}

func (ServiceJiraArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceJiraArgs)(nil)).Elem()
}
