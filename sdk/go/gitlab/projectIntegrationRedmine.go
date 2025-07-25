// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ProjectIntegrationRedmine` resource manages the lifecycle of a project integration with Redmine.
//
// > Using Redmine requires that GitLab internal issue tracking is disabled for the project.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/api/project_integrations/#redmine)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
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
//			_, err = gitlab.NewProjectIntegrationRedmine(ctx, "redmine", &gitlab.ProjectIntegrationRedmineArgs{
//				Project:     awesomeProject.ID(),
//				NewIssueUrl: pulumi.String("https://redmine.example.com/issue"),
//				ProjectUrl:  pulumi.String("https://redmine.example.com/project"),
//				IssuesUrl:   pulumi.String("https://redmine.example.com/issue/:id"),
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
// Starting in Terraform v1.5.0, you can use an import block to import `gitlab_project_integration_redmine`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_integration_redmine.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Importing using the CLI is supported with the following syntax:
//
// ```sh
// $ pulumi import gitlab:index/projectIntegrationRedmine:ProjectIntegrationRedmine You can import a gitlab_project_integration_redmine state using `<resource> <project_id>`:
// ```
//
// ```sh
// $ pulumi import gitlab:index/projectIntegrationRedmine:ProjectIntegrationRedmine redmine 1
// ```
type ProjectIntegrationRedmine struct {
	pulumi.CustomResourceState

	// The URL to the Redmine project issue to link to this GitLab project.
	IssuesUrl pulumi.StringOutput `pulumi:"issuesUrl"`
	// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
	NewIssueUrl pulumi.StringOutput `pulumi:"newIssueUrl"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL to the Redmine project to link to this GitLab project.
	ProjectUrl pulumi.StringOutput `pulumi:"projectUrl"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings pulumi.BoolOutput `pulumi:"useInheritedSettings"`
}

// NewProjectIntegrationRedmine registers a new resource with the given unique name, arguments, and options.
func NewProjectIntegrationRedmine(ctx *pulumi.Context,
	name string, args *ProjectIntegrationRedmineArgs, opts ...pulumi.ResourceOption) (*ProjectIntegrationRedmine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IssuesUrl == nil {
		return nil, errors.New("invalid value for required argument 'IssuesUrl'")
	}
	if args.NewIssueUrl == nil {
		return nil, errors.New("invalid value for required argument 'NewIssueUrl'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ProjectUrl == nil {
		return nil, errors.New("invalid value for required argument 'ProjectUrl'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectIntegrationRedmine
	err := ctx.RegisterResource("gitlab:index/projectIntegrationRedmine:ProjectIntegrationRedmine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectIntegrationRedmine gets an existing ProjectIntegrationRedmine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectIntegrationRedmine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectIntegrationRedmineState, opts ...pulumi.ResourceOption) (*ProjectIntegrationRedmine, error) {
	var resource ProjectIntegrationRedmine
	err := ctx.ReadResource("gitlab:index/projectIntegrationRedmine:ProjectIntegrationRedmine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectIntegrationRedmine resources.
type projectIntegrationRedmineState struct {
	// The URL to the Redmine project issue to link to this GitLab project.
	IssuesUrl *string `pulumi:"issuesUrl"`
	// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
	NewIssueUrl *string `pulumi:"newIssueUrl"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// The URL to the Redmine project to link to this GitLab project.
	ProjectUrl *string `pulumi:"projectUrl"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings *bool `pulumi:"useInheritedSettings"`
}

type ProjectIntegrationRedmineState struct {
	// The URL to the Redmine project issue to link to this GitLab project.
	IssuesUrl pulumi.StringPtrInput
	// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
	NewIssueUrl pulumi.StringPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// The URL to the Redmine project to link to this GitLab project.
	ProjectUrl pulumi.StringPtrInput
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings pulumi.BoolPtrInput
}

func (ProjectIntegrationRedmineState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIntegrationRedmineState)(nil)).Elem()
}

type projectIntegrationRedmineArgs struct {
	// The URL to the Redmine project issue to link to this GitLab project.
	IssuesUrl string `pulumi:"issuesUrl"`
	// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
	NewIssueUrl string `pulumi:"newIssueUrl"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// The URL to the Redmine project to link to this GitLab project.
	ProjectUrl string `pulumi:"projectUrl"`
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings *bool `pulumi:"useInheritedSettings"`
}

// The set of arguments for constructing a ProjectIntegrationRedmine resource.
type ProjectIntegrationRedmineArgs struct {
	// The URL to the Redmine project issue to link to this GitLab project.
	IssuesUrl pulumi.StringInput
	// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
	NewIssueUrl pulumi.StringInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// The URL to the Redmine project to link to this GitLab project.
	ProjectUrl pulumi.StringInput
	// Indicates whether or not to inherit default settings. Defaults to false.
	UseInheritedSettings pulumi.BoolPtrInput
}

func (ProjectIntegrationRedmineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectIntegrationRedmineArgs)(nil)).Elem()
}

type ProjectIntegrationRedmineInput interface {
	pulumi.Input

	ToProjectIntegrationRedmineOutput() ProjectIntegrationRedmineOutput
	ToProjectIntegrationRedmineOutputWithContext(ctx context.Context) ProjectIntegrationRedmineOutput
}

func (*ProjectIntegrationRedmine) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIntegrationRedmine)(nil)).Elem()
}

func (i *ProjectIntegrationRedmine) ToProjectIntegrationRedmineOutput() ProjectIntegrationRedmineOutput {
	return i.ToProjectIntegrationRedmineOutputWithContext(context.Background())
}

func (i *ProjectIntegrationRedmine) ToProjectIntegrationRedmineOutputWithContext(ctx context.Context) ProjectIntegrationRedmineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIntegrationRedmineOutput)
}

// ProjectIntegrationRedmineArrayInput is an input type that accepts ProjectIntegrationRedmineArray and ProjectIntegrationRedmineArrayOutput values.
// You can construct a concrete instance of `ProjectIntegrationRedmineArrayInput` via:
//
//	ProjectIntegrationRedmineArray{ ProjectIntegrationRedmineArgs{...} }
type ProjectIntegrationRedmineArrayInput interface {
	pulumi.Input

	ToProjectIntegrationRedmineArrayOutput() ProjectIntegrationRedmineArrayOutput
	ToProjectIntegrationRedmineArrayOutputWithContext(context.Context) ProjectIntegrationRedmineArrayOutput
}

type ProjectIntegrationRedmineArray []ProjectIntegrationRedmineInput

func (ProjectIntegrationRedmineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIntegrationRedmine)(nil)).Elem()
}

func (i ProjectIntegrationRedmineArray) ToProjectIntegrationRedmineArrayOutput() ProjectIntegrationRedmineArrayOutput {
	return i.ToProjectIntegrationRedmineArrayOutputWithContext(context.Background())
}

func (i ProjectIntegrationRedmineArray) ToProjectIntegrationRedmineArrayOutputWithContext(ctx context.Context) ProjectIntegrationRedmineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIntegrationRedmineArrayOutput)
}

// ProjectIntegrationRedmineMapInput is an input type that accepts ProjectIntegrationRedmineMap and ProjectIntegrationRedmineMapOutput values.
// You can construct a concrete instance of `ProjectIntegrationRedmineMapInput` via:
//
//	ProjectIntegrationRedmineMap{ "key": ProjectIntegrationRedmineArgs{...} }
type ProjectIntegrationRedmineMapInput interface {
	pulumi.Input

	ToProjectIntegrationRedmineMapOutput() ProjectIntegrationRedmineMapOutput
	ToProjectIntegrationRedmineMapOutputWithContext(context.Context) ProjectIntegrationRedmineMapOutput
}

type ProjectIntegrationRedmineMap map[string]ProjectIntegrationRedmineInput

func (ProjectIntegrationRedmineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIntegrationRedmine)(nil)).Elem()
}

func (i ProjectIntegrationRedmineMap) ToProjectIntegrationRedmineMapOutput() ProjectIntegrationRedmineMapOutput {
	return i.ToProjectIntegrationRedmineMapOutputWithContext(context.Background())
}

func (i ProjectIntegrationRedmineMap) ToProjectIntegrationRedmineMapOutputWithContext(ctx context.Context) ProjectIntegrationRedmineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectIntegrationRedmineMapOutput)
}

type ProjectIntegrationRedmineOutput struct{ *pulumi.OutputState }

func (ProjectIntegrationRedmineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectIntegrationRedmine)(nil)).Elem()
}

func (o ProjectIntegrationRedmineOutput) ToProjectIntegrationRedmineOutput() ProjectIntegrationRedmineOutput {
	return o
}

func (o ProjectIntegrationRedmineOutput) ToProjectIntegrationRedmineOutputWithContext(ctx context.Context) ProjectIntegrationRedmineOutput {
	return o
}

// The URL to the Redmine project issue to link to this GitLab project.
func (o ProjectIntegrationRedmineOutput) IssuesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationRedmine) pulumi.StringOutput { return v.IssuesUrl }).(pulumi.StringOutput)
}

// The URL to use to create a new issue in the Redmine project linked to this GitLab project.
func (o ProjectIntegrationRedmineOutput) NewIssueUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationRedmine) pulumi.StringOutput { return v.NewIssueUrl }).(pulumi.StringOutput)
}

// ID of the project you want to activate integration on.
func (o ProjectIntegrationRedmineOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationRedmine) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL to the Redmine project to link to this GitLab project.
func (o ProjectIntegrationRedmineOutput) ProjectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectIntegrationRedmine) pulumi.StringOutput { return v.ProjectUrl }).(pulumi.StringOutput)
}

// Indicates whether or not to inherit default settings. Defaults to false.
func (o ProjectIntegrationRedmineOutput) UseInheritedSettings() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProjectIntegrationRedmine) pulumi.BoolOutput { return v.UseInheritedSettings }).(pulumi.BoolOutput)
}

type ProjectIntegrationRedmineArrayOutput struct{ *pulumi.OutputState }

func (ProjectIntegrationRedmineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectIntegrationRedmine)(nil)).Elem()
}

func (o ProjectIntegrationRedmineArrayOutput) ToProjectIntegrationRedmineArrayOutput() ProjectIntegrationRedmineArrayOutput {
	return o
}

func (o ProjectIntegrationRedmineArrayOutput) ToProjectIntegrationRedmineArrayOutputWithContext(ctx context.Context) ProjectIntegrationRedmineArrayOutput {
	return o
}

func (o ProjectIntegrationRedmineArrayOutput) Index(i pulumi.IntInput) ProjectIntegrationRedmineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectIntegrationRedmine {
		return vs[0].([]*ProjectIntegrationRedmine)[vs[1].(int)]
	}).(ProjectIntegrationRedmineOutput)
}

type ProjectIntegrationRedmineMapOutput struct{ *pulumi.OutputState }

func (ProjectIntegrationRedmineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectIntegrationRedmine)(nil)).Elem()
}

func (o ProjectIntegrationRedmineMapOutput) ToProjectIntegrationRedmineMapOutput() ProjectIntegrationRedmineMapOutput {
	return o
}

func (o ProjectIntegrationRedmineMapOutput) ToProjectIntegrationRedmineMapOutputWithContext(ctx context.Context) ProjectIntegrationRedmineMapOutput {
	return o
}

func (o ProjectIntegrationRedmineMapOutput) MapIndex(k pulumi.StringInput) ProjectIntegrationRedmineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectIntegrationRedmine {
		return vs[0].(map[string]*ProjectIntegrationRedmine)[vs[1].(string)]
	}).(ProjectIntegrationRedmineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIntegrationRedmineInput)(nil)).Elem(), &ProjectIntegrationRedmine{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIntegrationRedmineArrayInput)(nil)).Elem(), ProjectIntegrationRedmineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectIntegrationRedmineMapInput)(nil)).Elem(), ProjectIntegrationRedmineMap{})
	pulumi.RegisterOutputType(ProjectIntegrationRedmineOutput{})
	pulumi.RegisterOutputType(ProjectIntegrationRedmineArrayOutput{})
	pulumi.RegisterOutputType(ProjectIntegrationRedmineMapOutput{})
}
