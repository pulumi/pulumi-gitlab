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

// ## Example Usage
//
// ## Import
//
// Starting in Terraform v1.5.0 you can use an import block to import `gitlab_project_approval_rule`. For example:
//
// terraform
//
// import {
//
//	to = gitlab_project_approval_rule.example
//
//	id = "see CLI command below for ID"
//
// }
//
// Import using the CLI is supported using the following syntax:
//
// GitLab project approval rules can be imported using a key composed of `<project-id>:<rule-id>`, e.g.
//
// ```sh
// $ pulumi import gitlab:index/projectApprovalRule:ProjectApprovalRule example "12345:6"
// ```
type ProjectApprovalRule struct {
	pulumi.CustomResourceState

	// Whether the rule is applied to all protected branches. If set to 'true', the value of `protectedBranchIds` is ignored. Default is 'false'.
	AppliesToAllProtectedBranches pulumi.BoolPtrOutput `pulumi:"appliesToAllProtectedBranches"`
	// The number of approvals required for this rule.
	ApprovalsRequired pulumi.IntOutput `pulumi:"approvalsRequired"`
	// When this flag is set, the default `anyApprover` rule will not be imported if present.
	DisableImportingDefaultAnyApproverRuleOnCreate pulumi.BoolPtrOutput `pulumi:"disableImportingDefaultAnyApproverRuleOnCreate"`
	// A list of group IDs whose members can approve of the merge request.
	GroupIds pulumi.IntArrayOutput `pulumi:"groupIds"`
	// The name of the approval rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the approval rules.
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds pulumi.IntArrayOutput `pulumi:"protectedBranchIds"`
	// Report type is required when the ruleType is `reportApprover`. Valid values are `codeCoverage`.
	ReportType pulumi.StringOutput `pulumi:"reportType"`
	// String, defaults to 'regular'. The type of rule. `anyApprover` is a pre-configured default rule with `approvalsRequired` at `0`. Valid values are `regular`, `anyApprover`, `reportApprover`.
	RuleType pulumi.StringOutput `pulumi:"ruleType"`
	// A list of specific User IDs to add to the list of approvers.
	UserIds pulumi.IntArrayOutput `pulumi:"userIds"`
}

// NewProjectApprovalRule registers a new resource with the given unique name, arguments, and options.
func NewProjectApprovalRule(ctx *pulumi.Context,
	name string, args *ProjectApprovalRuleArgs, opts ...pulumi.ResourceOption) (*ProjectApprovalRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApprovalsRequired == nil {
		return nil, errors.New("invalid value for required argument 'ApprovalsRequired'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectApprovalRule
	err := ctx.RegisterResource("gitlab:index/projectApprovalRule:ProjectApprovalRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectApprovalRule gets an existing ProjectApprovalRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectApprovalRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectApprovalRuleState, opts ...pulumi.ResourceOption) (*ProjectApprovalRule, error) {
	var resource ProjectApprovalRule
	err := ctx.ReadResource("gitlab:index/projectApprovalRule:ProjectApprovalRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectApprovalRule resources.
type projectApprovalRuleState struct {
	// Whether the rule is applied to all protected branches. If set to 'true', the value of `protectedBranchIds` is ignored. Default is 'false'.
	AppliesToAllProtectedBranches *bool `pulumi:"appliesToAllProtectedBranches"`
	// The number of approvals required for this rule.
	ApprovalsRequired *int `pulumi:"approvalsRequired"`
	// When this flag is set, the default `anyApprover` rule will not be imported if present.
	DisableImportingDefaultAnyApproverRuleOnCreate *bool `pulumi:"disableImportingDefaultAnyApproverRuleOnCreate"`
	// A list of group IDs whose members can approve of the merge request.
	GroupIds []int `pulumi:"groupIds"`
	// The name of the approval rule.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the approval rules.
	Project *string `pulumi:"project"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds []int `pulumi:"protectedBranchIds"`
	// Report type is required when the ruleType is `reportApprover`. Valid values are `codeCoverage`.
	ReportType *string `pulumi:"reportType"`
	// String, defaults to 'regular'. The type of rule. `anyApprover` is a pre-configured default rule with `approvalsRequired` at `0`. Valid values are `regular`, `anyApprover`, `reportApprover`.
	RuleType *string `pulumi:"ruleType"`
	// A list of specific User IDs to add to the list of approvers.
	UserIds []int `pulumi:"userIds"`
}

type ProjectApprovalRuleState struct {
	// Whether the rule is applied to all protected branches. If set to 'true', the value of `protectedBranchIds` is ignored. Default is 'false'.
	AppliesToAllProtectedBranches pulumi.BoolPtrInput
	// The number of approvals required for this rule.
	ApprovalsRequired pulumi.IntPtrInput
	// When this flag is set, the default `anyApprover` rule will not be imported if present.
	DisableImportingDefaultAnyApproverRuleOnCreate pulumi.BoolPtrInput
	// A list of group IDs whose members can approve of the merge request.
	GroupIds pulumi.IntArrayInput
	// The name of the approval rule.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the approval rules.
	Project pulumi.StringPtrInput
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds pulumi.IntArrayInput
	// Report type is required when the ruleType is `reportApprover`. Valid values are `codeCoverage`.
	ReportType pulumi.StringPtrInput
	// String, defaults to 'regular'. The type of rule. `anyApprover` is a pre-configured default rule with `approvalsRequired` at `0`. Valid values are `regular`, `anyApprover`, `reportApprover`.
	RuleType pulumi.StringPtrInput
	// A list of specific User IDs to add to the list of approvers.
	UserIds pulumi.IntArrayInput
}

func (ProjectApprovalRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectApprovalRuleState)(nil)).Elem()
}

type projectApprovalRuleArgs struct {
	// Whether the rule is applied to all protected branches. If set to 'true', the value of `protectedBranchIds` is ignored. Default is 'false'.
	AppliesToAllProtectedBranches *bool `pulumi:"appliesToAllProtectedBranches"`
	// The number of approvals required for this rule.
	ApprovalsRequired int `pulumi:"approvalsRequired"`
	// When this flag is set, the default `anyApprover` rule will not be imported if present.
	DisableImportingDefaultAnyApproverRuleOnCreate *bool `pulumi:"disableImportingDefaultAnyApproverRuleOnCreate"`
	// A list of group IDs whose members can approve of the merge request.
	GroupIds []int `pulumi:"groupIds"`
	// The name of the approval rule.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the approval rules.
	Project string `pulumi:"project"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds []int `pulumi:"protectedBranchIds"`
	// Report type is required when the ruleType is `reportApprover`. Valid values are `codeCoverage`.
	ReportType *string `pulumi:"reportType"`
	// String, defaults to 'regular'. The type of rule. `anyApprover` is a pre-configured default rule with `approvalsRequired` at `0`. Valid values are `regular`, `anyApprover`, `reportApprover`.
	RuleType *string `pulumi:"ruleType"`
	// A list of specific User IDs to add to the list of approvers.
	UserIds []int `pulumi:"userIds"`
}

// The set of arguments for constructing a ProjectApprovalRule resource.
type ProjectApprovalRuleArgs struct {
	// Whether the rule is applied to all protected branches. If set to 'true', the value of `protectedBranchIds` is ignored. Default is 'false'.
	AppliesToAllProtectedBranches pulumi.BoolPtrInput
	// The number of approvals required for this rule.
	ApprovalsRequired pulumi.IntInput
	// When this flag is set, the default `anyApprover` rule will not be imported if present.
	DisableImportingDefaultAnyApproverRuleOnCreate pulumi.BoolPtrInput
	// A list of group IDs whose members can approve of the merge request.
	GroupIds pulumi.IntArrayInput
	// The name of the approval rule.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the approval rules.
	Project pulumi.StringInput
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds pulumi.IntArrayInput
	// Report type is required when the ruleType is `reportApprover`. Valid values are `codeCoverage`.
	ReportType pulumi.StringPtrInput
	// String, defaults to 'regular'. The type of rule. `anyApprover` is a pre-configured default rule with `approvalsRequired` at `0`. Valid values are `regular`, `anyApprover`, `reportApprover`.
	RuleType pulumi.StringPtrInput
	// A list of specific User IDs to add to the list of approvers.
	UserIds pulumi.IntArrayInput
}

func (ProjectApprovalRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectApprovalRuleArgs)(nil)).Elem()
}

type ProjectApprovalRuleInput interface {
	pulumi.Input

	ToProjectApprovalRuleOutput() ProjectApprovalRuleOutput
	ToProjectApprovalRuleOutputWithContext(ctx context.Context) ProjectApprovalRuleOutput
}

func (*ProjectApprovalRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectApprovalRule)(nil)).Elem()
}

func (i *ProjectApprovalRule) ToProjectApprovalRuleOutput() ProjectApprovalRuleOutput {
	return i.ToProjectApprovalRuleOutputWithContext(context.Background())
}

func (i *ProjectApprovalRule) ToProjectApprovalRuleOutputWithContext(ctx context.Context) ProjectApprovalRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApprovalRuleOutput)
}

// ProjectApprovalRuleArrayInput is an input type that accepts ProjectApprovalRuleArray and ProjectApprovalRuleArrayOutput values.
// You can construct a concrete instance of `ProjectApprovalRuleArrayInput` via:
//
//	ProjectApprovalRuleArray{ ProjectApprovalRuleArgs{...} }
type ProjectApprovalRuleArrayInput interface {
	pulumi.Input

	ToProjectApprovalRuleArrayOutput() ProjectApprovalRuleArrayOutput
	ToProjectApprovalRuleArrayOutputWithContext(context.Context) ProjectApprovalRuleArrayOutput
}

type ProjectApprovalRuleArray []ProjectApprovalRuleInput

func (ProjectApprovalRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectApprovalRule)(nil)).Elem()
}

func (i ProjectApprovalRuleArray) ToProjectApprovalRuleArrayOutput() ProjectApprovalRuleArrayOutput {
	return i.ToProjectApprovalRuleArrayOutputWithContext(context.Background())
}

func (i ProjectApprovalRuleArray) ToProjectApprovalRuleArrayOutputWithContext(ctx context.Context) ProjectApprovalRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApprovalRuleArrayOutput)
}

// ProjectApprovalRuleMapInput is an input type that accepts ProjectApprovalRuleMap and ProjectApprovalRuleMapOutput values.
// You can construct a concrete instance of `ProjectApprovalRuleMapInput` via:
//
//	ProjectApprovalRuleMap{ "key": ProjectApprovalRuleArgs{...} }
type ProjectApprovalRuleMapInput interface {
	pulumi.Input

	ToProjectApprovalRuleMapOutput() ProjectApprovalRuleMapOutput
	ToProjectApprovalRuleMapOutputWithContext(context.Context) ProjectApprovalRuleMapOutput
}

type ProjectApprovalRuleMap map[string]ProjectApprovalRuleInput

func (ProjectApprovalRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectApprovalRule)(nil)).Elem()
}

func (i ProjectApprovalRuleMap) ToProjectApprovalRuleMapOutput() ProjectApprovalRuleMapOutput {
	return i.ToProjectApprovalRuleMapOutputWithContext(context.Background())
}

func (i ProjectApprovalRuleMap) ToProjectApprovalRuleMapOutputWithContext(ctx context.Context) ProjectApprovalRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApprovalRuleMapOutput)
}

type ProjectApprovalRuleOutput struct{ *pulumi.OutputState }

func (ProjectApprovalRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectApprovalRule)(nil)).Elem()
}

func (o ProjectApprovalRuleOutput) ToProjectApprovalRuleOutput() ProjectApprovalRuleOutput {
	return o
}

func (o ProjectApprovalRuleOutput) ToProjectApprovalRuleOutputWithContext(ctx context.Context) ProjectApprovalRuleOutput {
	return o
}

// Whether the rule is applied to all protected branches. If set to 'true', the value of `protectedBranchIds` is ignored. Default is 'false'.
func (o ProjectApprovalRuleOutput) AppliesToAllProtectedBranches() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.BoolPtrOutput { return v.AppliesToAllProtectedBranches }).(pulumi.BoolPtrOutput)
}

// The number of approvals required for this rule.
func (o ProjectApprovalRuleOutput) ApprovalsRequired() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.IntOutput { return v.ApprovalsRequired }).(pulumi.IntOutput)
}

// When this flag is set, the default `anyApprover` rule will not be imported if present.
func (o ProjectApprovalRuleOutput) DisableImportingDefaultAnyApproverRuleOnCreate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.BoolPtrOutput {
		return v.DisableImportingDefaultAnyApproverRuleOnCreate
	}).(pulumi.BoolPtrOutput)
}

// A list of group IDs whose members can approve of the merge request.
func (o ProjectApprovalRuleOutput) GroupIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.IntArrayOutput { return v.GroupIds }).(pulumi.IntArrayOutput)
}

// The name of the approval rule.
func (o ProjectApprovalRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name or id of the project to add the approval rules.
func (o ProjectApprovalRuleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A list of protected branch IDs (not branch names) for which the rule applies.
func (o ProjectApprovalRuleOutput) ProtectedBranchIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.IntArrayOutput { return v.ProtectedBranchIds }).(pulumi.IntArrayOutput)
}

// Report type is required when the ruleType is `reportApprover`. Valid values are `codeCoverage`.
func (o ProjectApprovalRuleOutput) ReportType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.StringOutput { return v.ReportType }).(pulumi.StringOutput)
}

// String, defaults to 'regular'. The type of rule. `anyApprover` is a pre-configured default rule with `approvalsRequired` at `0`. Valid values are `regular`, `anyApprover`, `reportApprover`.
func (o ProjectApprovalRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

// A list of specific User IDs to add to the list of approvers.
func (o ProjectApprovalRuleOutput) UserIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) pulumi.IntArrayOutput { return v.UserIds }).(pulumi.IntArrayOutput)
}

type ProjectApprovalRuleArrayOutput struct{ *pulumi.OutputState }

func (ProjectApprovalRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectApprovalRule)(nil)).Elem()
}

func (o ProjectApprovalRuleArrayOutput) ToProjectApprovalRuleArrayOutput() ProjectApprovalRuleArrayOutput {
	return o
}

func (o ProjectApprovalRuleArrayOutput) ToProjectApprovalRuleArrayOutputWithContext(ctx context.Context) ProjectApprovalRuleArrayOutput {
	return o
}

func (o ProjectApprovalRuleArrayOutput) Index(i pulumi.IntInput) ProjectApprovalRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectApprovalRule {
		return vs[0].([]*ProjectApprovalRule)[vs[1].(int)]
	}).(ProjectApprovalRuleOutput)
}

type ProjectApprovalRuleMapOutput struct{ *pulumi.OutputState }

func (ProjectApprovalRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectApprovalRule)(nil)).Elem()
}

func (o ProjectApprovalRuleMapOutput) ToProjectApprovalRuleMapOutput() ProjectApprovalRuleMapOutput {
	return o
}

func (o ProjectApprovalRuleMapOutput) ToProjectApprovalRuleMapOutputWithContext(ctx context.Context) ProjectApprovalRuleMapOutput {
	return o
}

func (o ProjectApprovalRuleMapOutput) MapIndex(k pulumi.StringInput) ProjectApprovalRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectApprovalRule {
		return vs[0].(map[string]*ProjectApprovalRule)[vs[1].(string)]
	}).(ProjectApprovalRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRuleInput)(nil)).Elem(), &ProjectApprovalRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRuleArrayInput)(nil)).Elem(), ProjectApprovalRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRuleMapInput)(nil)).Elem(), ProjectApprovalRuleMap{})
	pulumi.RegisterOutputType(ProjectApprovalRuleOutput{})
	pulumi.RegisterOutputType(ProjectApprovalRuleArrayOutput{})
	pulumi.RegisterOutputType(ProjectApprovalRuleMapOutput{})
}
