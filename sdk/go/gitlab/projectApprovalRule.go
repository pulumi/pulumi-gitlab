// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # gitlab\_project\_approval\_rule
//
// This resource allows you to create and manage multiple approval rules for your GitLab
// projects. For further information on approval rules, consult the [gitlab
// documentation](https://docs.gitlab.com/ee/api/merge_request_approvals.html#project-level-mr-approvals).
//
// > This feature requires GitLab Premium.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewProjectApprovalRule(ctx, "example_one", &gitlab.ProjectApprovalRuleArgs{
// 			ApprovalsRequired: pulumi.Int(3),
// 			GroupIds: pulumi.IntArray{
// 				pulumi.Int(51),
// 			},
// 			Project: pulumi.String("5"),
// 			UserIds: pulumi.IntArray{
// 				pulumi.Int(50),
// 				pulumi.Int(500),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Protected Branch IDs
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBranchProtection, err := gitlab.NewBranchProtection(ctx, "exampleBranchProtection", &gitlab.BranchProtectionArgs{
// 			Project:          pulumi.String("5"),
// 			Branch:           pulumi.String("release/*"),
// 			PushAccessLevel:  pulumi.String("maintainer"),
// 			MergeAccessLevel: pulumi.String("developer"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewProjectApprovalRule(ctx, "exampleProjectApprovalRule", &gitlab.ProjectApprovalRuleArgs{
// 			Project:           pulumi.String("5"),
// 			ApprovalsRequired: pulumi.Int(3),
// 			UserIds: pulumi.IntArray{
// 				pulumi.Int(50),
// 				pulumi.Int(500),
// 			},
// 			GroupIds: pulumi.IntArray{
// 				pulumi.Int(51),
// 			},
// 			ProtectedBranchIds: pulumi.IntArray{
// 				exampleBranchProtection.BranchProtectionId,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// GitLab project approval rules can be imported using an id consisting of `project-id:rule-id`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/projectApprovalRule:ProjectApprovalRule example "12345:6"
// ```
type ProjectApprovalRule struct {
	pulumi.CustomResourceState

	// The number of approvals required for this rule.
	ApprovalsRequired pulumi.IntOutput `pulumi:"approvalsRequired"`
	// A list of group IDs whose members can approve of the merge request.
	GroupIds pulumi.IntArrayOutput `pulumi:"groupIds"`
	// The name of the approval rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the approval rules.
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds pulumi.IntArrayOutput `pulumi:"protectedBranchIds"`
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
	// The number of approvals required for this rule.
	ApprovalsRequired *int `pulumi:"approvalsRequired"`
	// A list of group IDs whose members can approve of the merge request.
	GroupIds []int `pulumi:"groupIds"`
	// The name of the approval rule.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the approval rules.
	Project *string `pulumi:"project"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds []int `pulumi:"protectedBranchIds"`
	// A list of specific User IDs to add to the list of approvers.
	UserIds []int `pulumi:"userIds"`
}

type ProjectApprovalRuleState struct {
	// The number of approvals required for this rule.
	ApprovalsRequired pulumi.IntPtrInput
	// A list of group IDs whose members can approve of the merge request.
	GroupIds pulumi.IntArrayInput
	// The name of the approval rule.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the approval rules.
	Project pulumi.StringPtrInput
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds pulumi.IntArrayInput
	// A list of specific User IDs to add to the list of approvers.
	UserIds pulumi.IntArrayInput
}

func (ProjectApprovalRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectApprovalRuleState)(nil)).Elem()
}

type projectApprovalRuleArgs struct {
	// The number of approvals required for this rule.
	ApprovalsRequired int `pulumi:"approvalsRequired"`
	// A list of group IDs whose members can approve of the merge request.
	GroupIds []int `pulumi:"groupIds"`
	// The name of the approval rule.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the approval rules.
	Project string `pulumi:"project"`
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds []int `pulumi:"protectedBranchIds"`
	// A list of specific User IDs to add to the list of approvers.
	UserIds []int `pulumi:"userIds"`
}

// The set of arguments for constructing a ProjectApprovalRule resource.
type ProjectApprovalRuleArgs struct {
	// The number of approvals required for this rule.
	ApprovalsRequired pulumi.IntInput
	// A list of group IDs whose members can approve of the merge request.
	GroupIds pulumi.IntArrayInput
	// The name of the approval rule.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the approval rules.
	Project pulumi.StringInput
	// A list of protected branch IDs (not branch names) for which the rule applies.
	ProtectedBranchIds pulumi.IntArrayInput
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
	return reflect.TypeOf((*ProjectApprovalRule)(nil))
}

func (i *ProjectApprovalRule) ToProjectApprovalRuleOutput() ProjectApprovalRuleOutput {
	return i.ToProjectApprovalRuleOutputWithContext(context.Background())
}

func (i *ProjectApprovalRule) ToProjectApprovalRuleOutputWithContext(ctx context.Context) ProjectApprovalRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApprovalRuleOutput)
}

func (i *ProjectApprovalRule) ToProjectApprovalRulePtrOutput() ProjectApprovalRulePtrOutput {
	return i.ToProjectApprovalRulePtrOutputWithContext(context.Background())
}

func (i *ProjectApprovalRule) ToProjectApprovalRulePtrOutputWithContext(ctx context.Context) ProjectApprovalRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApprovalRulePtrOutput)
}

type ProjectApprovalRulePtrInput interface {
	pulumi.Input

	ToProjectApprovalRulePtrOutput() ProjectApprovalRulePtrOutput
	ToProjectApprovalRulePtrOutputWithContext(ctx context.Context) ProjectApprovalRulePtrOutput
}

type projectApprovalRulePtrType ProjectApprovalRuleArgs

func (*projectApprovalRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectApprovalRule)(nil))
}

func (i *projectApprovalRulePtrType) ToProjectApprovalRulePtrOutput() ProjectApprovalRulePtrOutput {
	return i.ToProjectApprovalRulePtrOutputWithContext(context.Background())
}

func (i *projectApprovalRulePtrType) ToProjectApprovalRulePtrOutputWithContext(ctx context.Context) ProjectApprovalRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApprovalRulePtrOutput)
}

// ProjectApprovalRuleArrayInput is an input type that accepts ProjectApprovalRuleArray and ProjectApprovalRuleArrayOutput values.
// You can construct a concrete instance of `ProjectApprovalRuleArrayInput` via:
//
//          ProjectApprovalRuleArray{ ProjectApprovalRuleArgs{...} }
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
//          ProjectApprovalRuleMap{ "key": ProjectApprovalRuleArgs{...} }
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
	return reflect.TypeOf((*ProjectApprovalRule)(nil))
}

func (o ProjectApprovalRuleOutput) ToProjectApprovalRuleOutput() ProjectApprovalRuleOutput {
	return o
}

func (o ProjectApprovalRuleOutput) ToProjectApprovalRuleOutputWithContext(ctx context.Context) ProjectApprovalRuleOutput {
	return o
}

func (o ProjectApprovalRuleOutput) ToProjectApprovalRulePtrOutput() ProjectApprovalRulePtrOutput {
	return o.ToProjectApprovalRulePtrOutputWithContext(context.Background())
}

func (o ProjectApprovalRuleOutput) ToProjectApprovalRulePtrOutputWithContext(ctx context.Context) ProjectApprovalRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectApprovalRule) *ProjectApprovalRule {
		return &v
	}).(ProjectApprovalRulePtrOutput)
}

type ProjectApprovalRulePtrOutput struct{ *pulumi.OutputState }

func (ProjectApprovalRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectApprovalRule)(nil))
}

func (o ProjectApprovalRulePtrOutput) ToProjectApprovalRulePtrOutput() ProjectApprovalRulePtrOutput {
	return o
}

func (o ProjectApprovalRulePtrOutput) ToProjectApprovalRulePtrOutputWithContext(ctx context.Context) ProjectApprovalRulePtrOutput {
	return o
}

func (o ProjectApprovalRulePtrOutput) Elem() ProjectApprovalRuleOutput {
	return o.ApplyT(func(v *ProjectApprovalRule) ProjectApprovalRule {
		if v != nil {
			return *v
		}
		var ret ProjectApprovalRule
		return ret
	}).(ProjectApprovalRuleOutput)
}

type ProjectApprovalRuleArrayOutput struct{ *pulumi.OutputState }

func (ProjectApprovalRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectApprovalRule)(nil))
}

func (o ProjectApprovalRuleArrayOutput) ToProjectApprovalRuleArrayOutput() ProjectApprovalRuleArrayOutput {
	return o
}

func (o ProjectApprovalRuleArrayOutput) ToProjectApprovalRuleArrayOutputWithContext(ctx context.Context) ProjectApprovalRuleArrayOutput {
	return o
}

func (o ProjectApprovalRuleArrayOutput) Index(i pulumi.IntInput) ProjectApprovalRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectApprovalRule {
		return vs[0].([]ProjectApprovalRule)[vs[1].(int)]
	}).(ProjectApprovalRuleOutput)
}

type ProjectApprovalRuleMapOutput struct{ *pulumi.OutputState }

func (ProjectApprovalRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProjectApprovalRule)(nil))
}

func (o ProjectApprovalRuleMapOutput) ToProjectApprovalRuleMapOutput() ProjectApprovalRuleMapOutput {
	return o
}

func (o ProjectApprovalRuleMapOutput) ToProjectApprovalRuleMapOutputWithContext(ctx context.Context) ProjectApprovalRuleMapOutput {
	return o
}

func (o ProjectApprovalRuleMapOutput) MapIndex(k pulumi.StringInput) ProjectApprovalRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProjectApprovalRule {
		return vs[0].(map[string]ProjectApprovalRule)[vs[1].(string)]
	}).(ProjectApprovalRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRuleInput)(nil)).Elem(), &ProjectApprovalRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRulePtrInput)(nil)).Elem(), &ProjectApprovalRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRuleArrayInput)(nil)).Elem(), ProjectApprovalRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApprovalRuleMapInput)(nil)).Elem(), ProjectApprovalRuleMap{})
	pulumi.RegisterOutputType(ProjectApprovalRuleOutput{})
	pulumi.RegisterOutputType(ProjectApprovalRulePtrOutput{})
	pulumi.RegisterOutputType(ProjectApprovalRuleArrayOutput{})
	pulumi.RegisterOutputType(ProjectApprovalRuleMapOutput{})
}
