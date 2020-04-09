// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to create and manage push rules for your GitLab projects.
// For further information on push rules, consult the [gitlab
// documentation](https://docs.gitlab.com/ce/push_rules/push_rules.html#push-rules).
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project_push_rules.html.markdown.
type ProjectPushRules struct {
	pulumi.CustomResourceState

	// All commit author emails must match this regex, e.g. "@my-company.com$"
	AuthorEmailRegex pulumi.StringPtrOutput `pulumi:"authorEmailRegex"`
	// All branch names must match this regex, e.g. "(feature|hotfix)\/*"
	BranchNameRegex pulumi.StringPtrOutput `pulumi:"branchNameRegex"`
	// All commit messages must match this regex, e.g. "Fixed \d+\..*"
	CommitMessageRegex pulumi.StringPtrOutput `pulumi:"commitMessageRegex"`
	// Deny deleting a tag
	DenyDeleteTag pulumi.BoolPtrOutput `pulumi:"denyDeleteTag"`
	// All commited filenames must not match this regex, e.g. "(jar|exe)$"
	FileNameRegex pulumi.StringPtrOutput `pulumi:"fileNameRegex"`
	// Maximum file size (MB)
	MaxFileSize pulumi.IntPtrOutput `pulumi:"maxFileSize"`
	// Restrict commits by author (email) to existing GitLab users
	MemberCheck pulumi.BoolPtrOutput `pulumi:"memberCheck"`
	// GitLab will reject any files that are likely to contain secrets
	PreventSecrets pulumi.BoolPtrOutput `pulumi:"preventSecrets"`
	// The name or id of the project to add the push rules to.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectPushRules registers a new resource with the given unique name, arguments, and options.
func NewProjectPushRules(ctx *pulumi.Context,
	name string, args *ProjectPushRulesArgs, opts ...pulumi.ResourceOption) (*ProjectPushRules, error) {
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil {
		args = &ProjectPushRulesArgs{}
	}
	var resource ProjectPushRules
	err := ctx.RegisterResource("gitlab:index/projectPushRules:ProjectPushRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectPushRules gets an existing ProjectPushRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectPushRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectPushRulesState, opts ...pulumi.ResourceOption) (*ProjectPushRules, error) {
	var resource ProjectPushRules
	err := ctx.ReadResource("gitlab:index/projectPushRules:ProjectPushRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectPushRules resources.
type projectPushRulesState struct {
	// All commit author emails must match this regex, e.g. "@my-company.com$"
	AuthorEmailRegex *string `pulumi:"authorEmailRegex"`
	// All branch names must match this regex, e.g. "(feature|hotfix)\/*"
	BranchNameRegex *string `pulumi:"branchNameRegex"`
	// All commit messages must match this regex, e.g. "Fixed \d+\..*"
	CommitMessageRegex *string `pulumi:"commitMessageRegex"`
	// Deny deleting a tag
	DenyDeleteTag *bool `pulumi:"denyDeleteTag"`
	// All commited filenames must not match this regex, e.g. "(jar|exe)$"
	FileNameRegex *string `pulumi:"fileNameRegex"`
	// Maximum file size (MB)
	MaxFileSize *int `pulumi:"maxFileSize"`
	// Restrict commits by author (email) to existing GitLab users
	MemberCheck *bool `pulumi:"memberCheck"`
	// GitLab will reject any files that are likely to contain secrets
	PreventSecrets *bool `pulumi:"preventSecrets"`
	// The name or id of the project to add the push rules to.
	Project *string `pulumi:"project"`
}

type ProjectPushRulesState struct {
	// All commit author emails must match this regex, e.g. "@my-company.com$"
	AuthorEmailRegex pulumi.StringPtrInput
	// All branch names must match this regex, e.g. "(feature|hotfix)\/*"
	BranchNameRegex pulumi.StringPtrInput
	// All commit messages must match this regex, e.g. "Fixed \d+\..*"
	CommitMessageRegex pulumi.StringPtrInput
	// Deny deleting a tag
	DenyDeleteTag pulumi.BoolPtrInput
	// All commited filenames must not match this regex, e.g. "(jar|exe)$"
	FileNameRegex pulumi.StringPtrInput
	// Maximum file size (MB)
	MaxFileSize pulumi.IntPtrInput
	// Restrict commits by author (email) to existing GitLab users
	MemberCheck pulumi.BoolPtrInput
	// GitLab will reject any files that are likely to contain secrets
	PreventSecrets pulumi.BoolPtrInput
	// The name or id of the project to add the push rules to.
	Project pulumi.StringPtrInput
}

func (ProjectPushRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPushRulesState)(nil)).Elem()
}

type projectPushRulesArgs struct {
	// All commit author emails must match this regex, e.g. "@my-company.com$"
	AuthorEmailRegex *string `pulumi:"authorEmailRegex"`
	// All branch names must match this regex, e.g. "(feature|hotfix)\/*"
	BranchNameRegex *string `pulumi:"branchNameRegex"`
	// All commit messages must match this regex, e.g. "Fixed \d+\..*"
	CommitMessageRegex *string `pulumi:"commitMessageRegex"`
	// Deny deleting a tag
	DenyDeleteTag *bool `pulumi:"denyDeleteTag"`
	// All commited filenames must not match this regex, e.g. "(jar|exe)$"
	FileNameRegex *string `pulumi:"fileNameRegex"`
	// Maximum file size (MB)
	MaxFileSize *int `pulumi:"maxFileSize"`
	// Restrict commits by author (email) to existing GitLab users
	MemberCheck *bool `pulumi:"memberCheck"`
	// GitLab will reject any files that are likely to contain secrets
	PreventSecrets *bool `pulumi:"preventSecrets"`
	// The name or id of the project to add the push rules to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectPushRules resource.
type ProjectPushRulesArgs struct {
	// All commit author emails must match this regex, e.g. "@my-company.com$"
	AuthorEmailRegex pulumi.StringPtrInput
	// All branch names must match this regex, e.g. "(feature|hotfix)\/*"
	BranchNameRegex pulumi.StringPtrInput
	// All commit messages must match this regex, e.g. "Fixed \d+\..*"
	CommitMessageRegex pulumi.StringPtrInput
	// Deny deleting a tag
	DenyDeleteTag pulumi.BoolPtrInput
	// All commited filenames must not match this regex, e.g. "(jar|exe)$"
	FileNameRegex pulumi.StringPtrInput
	// Maximum file size (MB)
	MaxFileSize pulumi.IntPtrInput
	// Restrict commits by author (email) to existing GitLab users
	MemberCheck pulumi.BoolPtrInput
	// GitLab will reject any files that are likely to contain secrets
	PreventSecrets pulumi.BoolPtrInput
	// The name or id of the project to add the push rules to.
	Project pulumi.StringInput
}

func (ProjectPushRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectPushRulesArgs)(nil)).Elem()
}
