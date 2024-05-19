// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gitlab/sdk/v7/go/gitlab/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gitlab:index/application:Application":
		r = &Application{}
	case "gitlab:index/applicationSettings:ApplicationSettings":
		r = &ApplicationSettings{}
	case "gitlab:index/branch:Branch":
		r = &Branch{}
	case "gitlab:index/branchProtection:BranchProtection":
		r = &BranchProtection{}
	case "gitlab:index/clusterAgent:ClusterAgent":
		r = &ClusterAgent{}
	case "gitlab:index/clusterAgentToken:ClusterAgentToken":
		r = &ClusterAgentToken{}
	case "gitlab:index/complianceFramework:ComplianceFramework":
		r = &ComplianceFramework{}
	case "gitlab:index/deployKey:DeployKey":
		r = &DeployKey{}
	case "gitlab:index/deployKeyEnable:DeployKeyEnable":
		r = &DeployKeyEnable{}
	case "gitlab:index/deployToken:DeployToken":
		r = &DeployToken{}
	case "gitlab:index/globalLevelNotifications:GlobalLevelNotifications":
		r = &GlobalLevelNotifications{}
	case "gitlab:index/group:Group":
		r = &Group{}
	case "gitlab:index/groupAccessToken:GroupAccessToken":
		r = &GroupAccessToken{}
	case "gitlab:index/groupBadge:GroupBadge":
		r = &GroupBadge{}
	case "gitlab:index/groupCluster:GroupCluster":
		r = &GroupCluster{}
	case "gitlab:index/groupCustomAttribute:GroupCustomAttribute":
		r = &GroupCustomAttribute{}
	case "gitlab:index/groupEpicBoard:GroupEpicBoard":
		r = &GroupEpicBoard{}
	case "gitlab:index/groupHook:GroupHook":
		r = &GroupHook{}
	case "gitlab:index/groupIssueBoard:GroupIssueBoard":
		r = &GroupIssueBoard{}
	case "gitlab:index/groupLabel:GroupLabel":
		r = &GroupLabel{}
	case "gitlab:index/groupLdapLink:GroupLdapLink":
		r = &GroupLdapLink{}
	case "gitlab:index/groupMembership:GroupMembership":
		r = &GroupMembership{}
	case "gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate":
		r = &GroupProjectFileTemplate{}
	case "gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment":
		r = &GroupProtectedEnvironment{}
	case "gitlab:index/groupSamlLink:GroupSamlLink":
		r = &GroupSamlLink{}
	case "gitlab:index/groupShareGroup:GroupShareGroup":
		r = &GroupShareGroup{}
	case "gitlab:index/groupVariable:GroupVariable":
		r = &GroupVariable{}
	case "gitlab:index/instanceCluster:InstanceCluster":
		r = &InstanceCluster{}
	case "gitlab:index/instanceVariable:InstanceVariable":
		r = &InstanceVariable{}
	case "gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker":
		r = &IntegrationCustomIssueTracker{}
	case "gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush":
		r = &IntegrationEmailsOnPush{}
	case "gitlab:index/integrationExternalWiki:IntegrationExternalWiki":
		r = &IntegrationExternalWiki{}
	case "gitlab:index/integrationGithub:IntegrationGithub":
		r = &IntegrationGithub{}
	case "gitlab:index/integrationJira:IntegrationJira":
		r = &IntegrationJira{}
	case "gitlab:index/integrationMattermost:IntegrationMattermost":
		r = &IntegrationMattermost{}
	case "gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams":
		r = &IntegrationMicrosoftTeams{}
	case "gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail":
		r = &IntegrationPipelinesEmail{}
	case "gitlab:index/integrationSlack:IntegrationSlack":
		r = &IntegrationSlack{}
	case "gitlab:index/label:Label":
		r = &Label{}
	case "gitlab:index/pagesDomain:PagesDomain":
		r = &PagesDomain{}
	case "gitlab:index/personalAccessToken:PersonalAccessToken":
		r = &PersonalAccessToken{}
	case "gitlab:index/pipelineSchedule:PipelineSchedule":
		r = &PipelineSchedule{}
	case "gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable":
		r = &PipelineScheduleVariable{}
	case "gitlab:index/pipelineTrigger:PipelineTrigger":
		r = &PipelineTrigger{}
	case "gitlab:index/project:Project":
		r = &Project{}
	case "gitlab:index/projectAccessToken:ProjectAccessToken":
		r = &ProjectAccessToken{}
	case "gitlab:index/projectApprovalRule:ProjectApprovalRule":
		r = &ProjectApprovalRule{}
	case "gitlab:index/projectBadge:ProjectBadge":
		r = &ProjectBadge{}
	case "gitlab:index/projectCluster:ProjectCluster":
		r = &ProjectCluster{}
	case "gitlab:index/projectComplianceFramework:ProjectComplianceFramework":
		r = &ProjectComplianceFramework{}
	case "gitlab:index/projectCustomAttribute:ProjectCustomAttribute":
		r = &ProjectCustomAttribute{}
	case "gitlab:index/projectEnvironment:ProjectEnvironment":
		r = &ProjectEnvironment{}
	case "gitlab:index/projectFreezePeriod:ProjectFreezePeriod":
		r = &ProjectFreezePeriod{}
	case "gitlab:index/projectHook:ProjectHook":
		r = &ProjectHook{}
	case "gitlab:index/projectIssue:ProjectIssue":
		r = &ProjectIssue{}
	case "gitlab:index/projectIssueBoard:ProjectIssueBoard":
		r = &ProjectIssueBoard{}
	case "gitlab:index/projectJobTokenScope:ProjectJobTokenScope":
		r = &ProjectJobTokenScope{}
	case "gitlab:index/projectLabel:ProjectLabel":
		r = &ProjectLabel{}
	case "gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals":
		r = &ProjectLevelMrApprovals{}
	case "gitlab:index/projectLevelNotifications:ProjectLevelNotifications":
		r = &ProjectLevelNotifications{}
	case "gitlab:index/projectMembership:ProjectMembership":
		r = &ProjectMembership{}
	case "gitlab:index/projectMilestone:ProjectMilestone":
		r = &ProjectMilestone{}
	case "gitlab:index/projectMirror:ProjectMirror":
		r = &ProjectMirror{}
	case "gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment":
		r = &ProjectProtectedEnvironment{}
	case "gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement":
		r = &ProjectRunnerEnablement{}
	case "gitlab:index/projectShareGroup:ProjectShareGroup":
		r = &ProjectShareGroup{}
	case "gitlab:index/projectTag:ProjectTag":
		r = &ProjectTag{}
	case "gitlab:index/projectVariable:ProjectVariable":
		r = &ProjectVariable{}
	case "gitlab:index/releaseLink:ReleaseLink":
		r = &ReleaseLink{}
	case "gitlab:index/repositoryFile:RepositoryFile":
		r = &RepositoryFile{}
	case "gitlab:index/runner:Runner":
		r = &Runner{}
	case "gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker":
		r = &ServiceCustomIssueTracker{}
	case "gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush":
		r = &ServiceEmailsOnPush{}
	case "gitlab:index/serviceExternalWiki:ServiceExternalWiki":
		r = &ServiceExternalWiki{}
	case "gitlab:index/serviceGithub:ServiceGithub":
		r = &ServiceGithub{}
	case "gitlab:index/serviceJira:ServiceJira":
		r = &ServiceJira{}
	case "gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams":
		r = &ServiceMicrosoftTeams{}
	case "gitlab:index/servicePipelinesEmail:ServicePipelinesEmail":
		r = &ServicePipelinesEmail{}
	case "gitlab:index/serviceSlack:ServiceSlack":
		r = &ServiceSlack{}
	case "gitlab:index/systemHook:SystemHook":
		r = &SystemHook{}
	case "gitlab:index/tagProtection:TagProtection":
		r = &TagProtection{}
	case "gitlab:index/topic:Topic":
		r = &Topic{}
	case "gitlab:index/user:User":
		r = &User{}
	case "gitlab:index/userCustomAttribute:UserCustomAttribute":
		r = &UserCustomAttribute{}
	case "gitlab:index/userGpgKey:UserGpgKey":
		r = &UserGpgKey{}
	case "gitlab:index/userRunner:UserRunner":
		r = &UserRunner{}
	case "gitlab:index/userSshKey:UserSshKey":
		r = &UserSshKey{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:gitlab" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/applicationSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/branch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/branchProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/clusterAgent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/clusterAgentToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/complianceFramework",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/deployKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/deployKeyEnable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/deployToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/globalLevelNotifications",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupAccessToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupBadge",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupCustomAttribute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupEpicBoard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupHook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupIssueBoard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupLabel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupLdapLink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupProjectFileTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupProtectedEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupSamlLink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupShareGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/groupVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/instanceCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/instanceVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationCustomIssueTracker",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationEmailsOnPush",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationExternalWiki",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationGithub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationJira",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationMattermost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationMicrosoftTeams",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationPipelinesEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/integrationSlack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/label",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/pagesDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/personalAccessToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/pipelineSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/pipelineScheduleVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/pipelineTrigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectAccessToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectApprovalRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectBadge",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectComplianceFramework",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectCustomAttribute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectFreezePeriod",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectHook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectIssue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectIssueBoard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectJobTokenScope",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectLabel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectLevelMrApprovals",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectLevelNotifications",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectMilestone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectMirror",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectProtectedEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectRunnerEnablement",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectShareGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectTag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/projectVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/releaseLink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/repositoryFile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/runner",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceCustomIssueTracker",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceEmailsOnPush",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceExternalWiki",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceGithub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceJira",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceMicrosoftTeams",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/servicePipelinesEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/serviceSlack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/systemHook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/tagProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/topic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/userCustomAttribute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/userGpgKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/userRunner",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gitlab",
		"index/userSshKey",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"gitlab",
		&pkg{version},
	)
}
