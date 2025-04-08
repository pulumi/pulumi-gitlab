# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .application import *
from .application_settings import *
from .branch import *
from .branch_protection import *
from .cluster_agent import *
from .cluster_agent_token import *
from .compliance_framework import *
from .deploy_key import *
from .deploy_key_enable import *
from .deploy_token import *
from .get_application import *
from .get_branch import *
from .get_cluster_agent import *
from .get_cluster_agents import *
from .get_compliance_framework import *
from .get_current_user import *
from .get_group import *
from .get_group_access_tokens import *
from .get_group_billable_member_memberships import *
from .get_group_hook import *
from .get_group_hooks import *
from .get_group_ids import *
from .get_group_membership import *
from .get_group_provisioned_users import *
from .get_group_service_account import *
from .get_group_subgroups import *
from .get_group_variable import *
from .get_group_variables import *
from .get_groups import *
from .get_instance_deploy_keys import *
from .get_instance_service_account import *
from .get_instance_variable import *
from .get_instance_variables import *
from .get_metadata import *
from .get_pipeline_schedule import *
from .get_pipeline_schedules import *
from .get_project import *
from .get_project_branches import *
from .get_project_environments import *
from .get_project_hook import *
from .get_project_hooks import *
from .get_project_ids import *
from .get_project_issue import *
from .get_project_issues import *
from .get_project_membership import *
from .get_project_merge_request import *
from .get_project_milestone import *
from .get_project_milestones import *
from .get_project_mirror_public_key import *
from .get_project_protected_branch import *
from .get_project_protected_branches import *
from .get_project_protected_tag import *
from .get_project_protected_tags import *
from .get_project_tag import *
from .get_project_tags import *
from .get_project_variable import *
from .get_project_variables import *
from .get_projects import *
from .get_release import *
from .get_release_link import *
from .get_release_links import *
from .get_repository_file import *
from .get_repository_tree import *
from .get_runners import *
from .get_user import *
from .get_user_sshkeys import *
from .get_users import *
from .global_level_notifications import *
from .group import *
from .group_access_token import *
from .group_badge import *
from .group_cluster import *
from .group_custom_attribute import *
from .group_epic_board import *
from .group_hook import *
from .group_issue_board import *
from .group_label import *
from .group_ldap_link import *
from .group_membership import *
from .group_project_file_template import *
from .group_protected_environment import *
from .group_saml_link import *
from .group_security_policy_attachment import *
from .group_service_account import *
from .group_service_account_access_token import *
from .group_share_group import *
from .group_variable import *
from .instance_cluster import *
from .instance_service_account import *
from .instance_variable import *
from .integration_custom_issue_tracker import *
from .integration_emails_on_push import *
from .integration_external_wiki import *
from .integration_github import *
from .integration_harbor import *
from .integration_jenkins import *
from .integration_jira import *
from .integration_mattermost import *
from .integration_microsoft_teams import *
from .integration_pipelines_email import *
from .integration_slack import *
from .integration_telegram import *
from .label import *
from .member_role import *
from .pages_domain import *
from .personal_access_token import *
from .pipeline_schedule import *
from .pipeline_schedule_variable import *
from .pipeline_trigger import *
from .project import *
from .project_access_token import *
from .project_approval_rule import *
from .project_badge import *
from .project_cluster import *
from .project_compliance_framework import *
from .project_compliance_frameworks import *
from .project_custom_attribute import *
from .project_environment import *
from .project_freeze_period import *
from .project_hook import *
from .project_issue import *
from .project_issue_board import *
from .project_job_token_scope import *
from .project_job_token_scopes import *
from .project_label import *
from .project_level_mr_approvals import *
from .project_level_notifications import *
from .project_membership import *
from .project_milestone import *
from .project_mirror import *
from .project_protected_environment import *
from .project_push_rules import *
from .project_runner_enablement import *
from .project_security_policy_attachment import *
from .project_share_group import *
from .project_tag import *
from .project_target_branch_rule import *
from .project_variable import *
from .project_wiki_page import *
from .provider import *
from .release import *
from .release_link import *
from .repository_file import *
from .runner import *
from .service_custom_issue_tracker import *
from .service_emails_on_push import *
from .service_external_wiki import *
from .service_github import *
from .service_jira import *
from .service_microsoft_teams import *
from .service_pipelines_email import *
from .service_slack import *
from .system_hook import *
from .tag_protection import *
from .topic import *
from .user import *
from .user_custom_attribute import *
from .user_gpg_key import *
from .user_identity import *
from .user_impersonation_token import *
from .user_runner import *
from .user_ssh_key import *
from .value_stream_analytics import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_gitlab.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_gitlab.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "gitlab",
  "mod": "index/application",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/application:Application": "Application"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/applicationSettings",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/applicationSettings:ApplicationSettings": "ApplicationSettings"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/branch",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/branch:Branch": "Branch"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/branchProtection",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/branchProtection:BranchProtection": "BranchProtection"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/clusterAgent",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/clusterAgent:ClusterAgent": "ClusterAgent"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/clusterAgentToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/clusterAgentToken:ClusterAgentToken": "ClusterAgentToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/complianceFramework",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/complianceFramework:ComplianceFramework": "ComplianceFramework"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/deployKey",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/deployKey:DeployKey": "DeployKey"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/deployKeyEnable",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/deployKeyEnable:DeployKeyEnable": "DeployKeyEnable"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/deployToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/deployToken:DeployToken": "DeployToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/globalLevelNotifications",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/globalLevelNotifications:GlobalLevelNotifications": "GlobalLevelNotifications"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/group",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/group:Group": "Group"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupAccessToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupAccessToken:GroupAccessToken": "GroupAccessToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupBadge",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupBadge:GroupBadge": "GroupBadge"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupCluster",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupCluster:GroupCluster": "GroupCluster"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupCustomAttribute",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupCustomAttribute:GroupCustomAttribute": "GroupCustomAttribute"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupEpicBoard",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupEpicBoard:GroupEpicBoard": "GroupEpicBoard"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupHook",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupHook:GroupHook": "GroupHook"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupIssueBoard",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupIssueBoard:GroupIssueBoard": "GroupIssueBoard"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupLabel",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupLabel:GroupLabel": "GroupLabel"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupLdapLink",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupLdapLink:GroupLdapLink": "GroupLdapLink"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupMembership",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupMembership:GroupMembership": "GroupMembership"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupProjectFileTemplate",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate": "GroupProjectFileTemplate"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupProtectedEnvironment",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupProtectedEnvironment:GroupProtectedEnvironment": "GroupProtectedEnvironment"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupSamlLink",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupSamlLink:GroupSamlLink": "GroupSamlLink"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupSecurityPolicyAttachment",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupSecurityPolicyAttachment:GroupSecurityPolicyAttachment": "GroupSecurityPolicyAttachment"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupServiceAccount",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupServiceAccount:GroupServiceAccount": "GroupServiceAccount"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupServiceAccountAccessToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupServiceAccountAccessToken:GroupServiceAccountAccessToken": "GroupServiceAccountAccessToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupShareGroup",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupShareGroup:GroupShareGroup": "GroupShareGroup"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/groupVariable",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/groupVariable:GroupVariable": "GroupVariable"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/instanceCluster",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/instanceCluster:InstanceCluster": "InstanceCluster"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/instanceServiceAccount",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/instanceServiceAccount:InstanceServiceAccount": "InstanceServiceAccount"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/instanceVariable",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/instanceVariable:InstanceVariable": "InstanceVariable"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationCustomIssueTracker",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationCustomIssueTracker:IntegrationCustomIssueTracker": "IntegrationCustomIssueTracker"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationEmailsOnPush",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationEmailsOnPush:IntegrationEmailsOnPush": "IntegrationEmailsOnPush"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationExternalWiki",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationExternalWiki:IntegrationExternalWiki": "IntegrationExternalWiki"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationGithub",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationGithub:IntegrationGithub": "IntegrationGithub"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationHarbor",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationHarbor:IntegrationHarbor": "IntegrationHarbor"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationJenkins",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationJenkins:IntegrationJenkins": "IntegrationJenkins"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationJira",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationJira:IntegrationJira": "IntegrationJira"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationMattermost",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationMattermost:IntegrationMattermost": "IntegrationMattermost"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationMicrosoftTeams",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationMicrosoftTeams:IntegrationMicrosoftTeams": "IntegrationMicrosoftTeams"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationPipelinesEmail",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationPipelinesEmail:IntegrationPipelinesEmail": "IntegrationPipelinesEmail"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationSlack",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationSlack:IntegrationSlack": "IntegrationSlack"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/integrationTelegram",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/integrationTelegram:IntegrationTelegram": "IntegrationTelegram"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/label",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/label:Label": "Label"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/memberRole",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/memberRole:MemberRole": "MemberRole"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/pagesDomain",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/pagesDomain:PagesDomain": "PagesDomain"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/personalAccessToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/personalAccessToken:PersonalAccessToken": "PersonalAccessToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/pipelineSchedule",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/pipelineSchedule:PipelineSchedule": "PipelineSchedule"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/pipelineScheduleVariable",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable": "PipelineScheduleVariable"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/pipelineTrigger",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/pipelineTrigger:PipelineTrigger": "PipelineTrigger"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/project",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/project:Project": "Project"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectAccessToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectAccessToken:ProjectAccessToken": "ProjectAccessToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectApprovalRule",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectApprovalRule:ProjectApprovalRule": "ProjectApprovalRule"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectBadge",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectBadge:ProjectBadge": "ProjectBadge"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectCluster",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectCluster:ProjectCluster": "ProjectCluster"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectComplianceFramework",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectComplianceFramework:ProjectComplianceFramework": "ProjectComplianceFramework"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectComplianceFrameworks",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectComplianceFrameworks:ProjectComplianceFrameworks": "ProjectComplianceFrameworks"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectCustomAttribute",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectCustomAttribute:ProjectCustomAttribute": "ProjectCustomAttribute"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectEnvironment",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectEnvironment:ProjectEnvironment": "ProjectEnvironment"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectFreezePeriod",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectFreezePeriod:ProjectFreezePeriod": "ProjectFreezePeriod"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectHook",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectHook:ProjectHook": "ProjectHook"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectIssue",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectIssue:ProjectIssue": "ProjectIssue"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectIssueBoard",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectIssueBoard:ProjectIssueBoard": "ProjectIssueBoard"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectJobTokenScope",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectJobTokenScope:ProjectJobTokenScope": "ProjectJobTokenScope"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectJobTokenScopes",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectJobTokenScopes:ProjectJobTokenScopes": "ProjectJobTokenScopes"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectLabel",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectLabel:ProjectLabel": "ProjectLabel"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectLevelMrApprovals",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals": "ProjectLevelMrApprovals"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectLevelNotifications",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectLevelNotifications:ProjectLevelNotifications": "ProjectLevelNotifications"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectMembership",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectMembership:ProjectMembership": "ProjectMembership"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectMilestone",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectMilestone:ProjectMilestone": "ProjectMilestone"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectMirror",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectMirror:ProjectMirror": "ProjectMirror"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectProtectedEnvironment",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment": "ProjectProtectedEnvironment"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectPushRules",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectPushRules:ProjectPushRules": "ProjectPushRules"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectRunnerEnablement",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement": "ProjectRunnerEnablement"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectSecurityPolicyAttachment",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectSecurityPolicyAttachment:ProjectSecurityPolicyAttachment": "ProjectSecurityPolicyAttachment"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectShareGroup",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectShareGroup:ProjectShareGroup": "ProjectShareGroup"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectTag",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectTag:ProjectTag": "ProjectTag"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectTargetBranchRule",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectTargetBranchRule:ProjectTargetBranchRule": "ProjectTargetBranchRule"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectVariable",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectVariable:ProjectVariable": "ProjectVariable"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/projectWikiPage",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/projectWikiPage:ProjectWikiPage": "ProjectWikiPage"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/release",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/release:Release": "Release"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/releaseLink",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/releaseLink:ReleaseLink": "ReleaseLink"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/repositoryFile",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/repositoryFile:RepositoryFile": "RepositoryFile"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/runner",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/runner:Runner": "Runner"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceCustomIssueTracker",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceCustomIssueTracker:ServiceCustomIssueTracker": "ServiceCustomIssueTracker"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceEmailsOnPush",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceEmailsOnPush:ServiceEmailsOnPush": "ServiceEmailsOnPush"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceExternalWiki",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceExternalWiki:ServiceExternalWiki": "ServiceExternalWiki"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceGithub",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceGithub:ServiceGithub": "ServiceGithub"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceJira",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceJira:ServiceJira": "ServiceJira"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceMicrosoftTeams",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams": "ServiceMicrosoftTeams"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/servicePipelinesEmail",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/servicePipelinesEmail:ServicePipelinesEmail": "ServicePipelinesEmail"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/serviceSlack",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/serviceSlack:ServiceSlack": "ServiceSlack"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/systemHook",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/systemHook:SystemHook": "SystemHook"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/tagProtection",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/tagProtection:TagProtection": "TagProtection"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/topic",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/topic:Topic": "Topic"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/user",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/user:User": "User"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/userCustomAttribute",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/userCustomAttribute:UserCustomAttribute": "UserCustomAttribute"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/userGpgKey",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/userGpgKey:UserGpgKey": "UserGpgKey"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/userIdentity",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/userIdentity:UserIdentity": "UserIdentity"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/userImpersonationToken",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/userImpersonationToken:UserImpersonationToken": "UserImpersonationToken"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/userRunner",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/userRunner:UserRunner": "UserRunner"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/userSshKey",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/userSshKey:UserSshKey": "UserSshKey"
  }
 },
 {
  "pkg": "gitlab",
  "mod": "index/valueStreamAnalytics",
  "fqn": "pulumi_gitlab",
  "classes": {
   "gitlab:index/valueStreamAnalytics:ValueStreamAnalytics": "ValueStreamAnalytics"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "gitlab",
  "token": "pulumi:providers:gitlab",
  "fqn": "pulumi_gitlab",
  "class": "Provider"
 }
]
"""
)
