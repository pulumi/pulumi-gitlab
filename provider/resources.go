// Copyright 2016-2018, Pulumi Corporation.(
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"path/filepath"
	"unicode"

	gitlabShim "github.com/gitlabhq/terraform-provider-gitlab/shim"
	"github.com/pulumi/pulumi-gitlab/provider/v4/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the GitLab token components used below.
const (
	// packages:
	gitLabPkg = "gitlab"
	// modules:
	gitLabMod = "index" // the root index.
)

// gitLabMember manufactures a type token for the GitLab package and the given module and type.
func gitLabMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(gitLabPkg + ":" + mod + ":" + mem)
}

// gitLabType manufactures a type token for the GitLab package and the given module and type.
func gitLabType(mod string, typ string) tokens.Type {
	return tokens.Type(gitLabMember(mod, typ))
}

// gitLabDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the GitLab package and names the file by simply lower casing the data
// source's first character.
func gitLabDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return gitLabMember(mod+"/"+fn, res)
}

// gitLabResource manufactures a standard resource token given a module and resource name.
// It automatically uses the GitLab package and names the file by simply lower casing the resource's
// first character.
func gitLabResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return gitLabType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the GitLab package.
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(gitlabShim.NewProvider())
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "gitlab",
		Description: "A Pulumi package for creating and managing GitLab resources.",
		Keywords:    []string{"pulumi", "gitlab"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		GitHubOrg:   "gitlabhq",
		Repository:  "https://github.com/pulumi/pulumi-gitlab",
		Config: map[string]*tfbridge.SchemaInfo{
			"cacert_file": {},
			"insecure":    {},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"gitlab_branch_protection":             {Tok: gitLabResource(gitLabMod, "BranchProtection")},
			"gitlab_tag_protection":                {Tok: gitLabResource(gitLabMod, "TagProtection")},
			"gitlab_group":                         {Tok: gitLabResource(gitLabMod, "Group")},
			"gitlab_group_label":                   {Tok: gitLabResource(gitLabMod, "GroupLabel")},
			"gitlab_group_cluster":                 {Tok: gitLabResource(gitLabMod, "GroupCluster")},
			"gitlab_project":                       {Tok: gitLabResource(gitLabMod, "Project")},
			"gitlab_label":                         {Tok: gitLabResource(gitLabMod, "Label")},
			"gitlab_pipeline_schedule":             {Tok: gitLabResource(gitLabMod, "PipelineSchedule")},
			"gitlab_pipeline_trigger":              {Tok: gitLabResource(gitLabMod, "PipelineTrigger")},
			"gitlab_pipeline_schedule_variable":    {Tok: gitLabResource(gitLabMod, "PipelineScheduleVariable")},
			"gitlab_project_hook":                  {Tok: gitLabResource(gitLabMod, "ProjectHook")},
			"gitlab_deploy_key":                    {Tok: gitLabResource(gitLabMod, "DeployKey")},
			"gitlab_deploy_key_enable":             {Tok: gitLabResource(gitLabMod, "DeployKeyEnable")},
			"gitlab_user":                          {Tok: gitLabResource(gitLabMod, "User")},
			"gitlab_project_membership":            {Tok: gitLabResource(gitLabMod, "ProjectMembership")},
			"gitlab_group_membership":              {Tok: gitLabResource(gitLabMod, "GroupMembership")},
			"gitlab_project_variable":              {Tok: gitLabResource(gitLabMod, "ProjectVariable")},
			"gitlab_project_share_group":           {Tok: gitLabResource(gitLabMod, "ProjectShareGroup")},
			"gitlab_group_variable":                {Tok: gitLabResource(gitLabMod, "GroupVariable")},
			"gitlab_project_cluster":               {Tok: gitLabResource(gitLabMod, "ProjectCluster")},
			"gitlab_service_slack":                 {Tok: gitLabResource(gitLabMod, "ServiceSlack")},
			"gitlab_service_jira":                  {Tok: gitLabResource(gitLabMod, "ServiceJira")},
			"gitlab_service_github":                {Tok: gitLabResource(gitLabMod, "ServiceGithub")},
			"gitlab_group_ldap_link":               {Tok: gitLabResource(gitLabMod, "GroupLdapLink")},
			"gitlab_deploy_token":                  {Tok: gitLabResource(gitLabMod, "DeployToken")},
			"gitlab_instance_cluster":              {Tok: gitLabResource(gitLabMod, "InstanceCluster")},
			"gitlab_project_level_mr_approvals":    {Tok: gitLabResource(gitLabMod, "ProjectLevelMrApprovals")},
			"gitlab_project_mirror":                {Tok: gitLabResource(gitLabMod, "ProjectMirror")},
			"gitlab_service_pipelines_email":       {Tok: gitLabResource(gitLabMod, "ServicePipelinesEmail")},
			"gitlab_project_approval_rule":         {Tok: gitLabResource(gitLabMod, "ProjectApprovalRule")},
			"gitlab_instance_variable":             {Tok: gitLabResource(gitLabMod, "InstanceVariable")},
			"gitlab_group_share_group":             {Tok: gitLabResource(gitLabMod, "GroupShareGroup")},
			"gitlab_project_freeze_period":         {Tok: gitLabResource(gitLabMod, "ProjectFreezePeriod")},
			"gitlab_project_badge":                 {Tok: gitLabResource(gitLabMod, "ProjectBadge")},
			"gitlab_group_badge":                   {Tok: gitLabResource(gitLabMod, "GroupBadge")},
			"gitlab_group_custom_attribute":        {Tok: gitLabResource(gitLabMod, "GroupCustomAttribute")},
			"gitlab_managed_license":               {Tok: gitLabResource(gitLabMod, "ManagedLicense")},
			"gitlab_project_access_token":          {Tok: gitLabResource(gitLabMod, "ProjectAccessToken")},
			"gitlab_project_custom_attribute":      {Tok: gitLabResource(gitLabMod, "ProjectCustomAttribute")},
			"gitlab_repository_file":               {Tok: gitLabResource(gitLabMod, "RepositoryFile")},
			"gitlab_service_microsoft_teams":       {Tok: gitLabResource(gitLabMod, "ServiceMicrosoftTeams")},
			"gitlab_user_custom_attribute":         {Tok: gitLabResource(gitLabMod, "UserCustomAttribute")},
			"gitlab_branch":                        {Tok: gitLabResource(gitLabMod, "Branch")},
			"gitlab_group_access_token":            {Tok: gitLabResource(gitLabMod, "GroupAccessToken")},
			"gitlab_topic":                         {Tok: gitLabResource(gitLabMod, "Topic")},
			"gitlab_user_sshkey":                   {Tok: gitLabResource(gitLabMod, "UserSshKey")},
			"gitlab_project_issue":                 {Tok: gitLabResource(gitLabMod, "ProjectIssue")},
			"gitlab_project_tag":                   {Tok: gitLabResource(gitLabMod, "ProjectTag")},
			"gitlab_project_environment":           {Tok: gitLabResource(gitLabMod, "ProjectEnvironment")},
			"gitlab_project_protected_environment": {Tok: gitLabResource(gitLabMod, "ProjectProtectedEnvironment")},
			"gitlab_system_hook":                   {Tok: gitLabResource(gitLabMod, "SystemHook")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"gitlab_group": {Tok: gitLabDataSource(gitLabMod, "getGroup")},
			"gitlab_project": {
				Tok: gitLabDataSource(gitLabMod, "getProject"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"push_rules": {
						Name:        "pushRules",
						MaxItemsOne: tfbridge.True(),
					},
				},
			},
			"gitlab_user":  {Tok: gitLabDataSource(gitLabMod, "getUser")},
			"gitlab_users": {Tok: gitLabDataSource(gitLabMod, "getUsers")},
			"gitlab_projects": {
				Tok: gitLabDataSource(gitLabMod, "getProjects"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"projects": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"permissions": {
									Name:        "permissions",
									MaxItemsOne: tfbridge.True(),
								},
								"namespace": {
									Name:        "namespace",
									MaxItemsOne: tfbridge.True(),
								},
								"forked_from_project": {
									Name:        "forkedFromProject",
									MaxItemsOne: tfbridge.True(),
								},
								"owner": {
									Name:        "owner",
									MaxItemsOne: tfbridge.True(),
								},
							},
						},
					},
				},
			},
			"gitlab_group_membership":           {Tok: gitLabDataSource(gitLabMod, "getGroupMembership")},
			"gitlab_project_protected_branch":   {Tok: gitLabDataSource(gitLabMod, "getProjectProtectedBranch")},
			"gitlab_project_protected_branches": {Tok: gitLabDataSource(gitLabMod, "getProjectProtectedBranches")},
			"gitlab_branch":                     {Tok: gitLabDataSource(gitLabMod, "getBranch")},
			"gitlab_project_issue":              {Tok: gitLabDataSource(gitLabMod, "getProjectIssue")},
			"gitlab_project_issues":             {Tok: gitLabDataSource(gitLabMod, "getProjectIssues")},
			"gitlab_project_tag":                {Tok: gitLabDataSource(gitLabMod, "getProjectTag")},
			"gitlab_instance_deploy_keys":       {Tok: gitLabDataSource(gitLabMod, "getInstanceDeployKeys")},
			"gitlab_project_tags":               {Tok: gitLabDataSource(gitLabMod, "getProjectTags")},
			"gitlab_repository_file":            {Tok: gitLabDataSource(gitLabMod, "getRepositoryFile")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^3.0.0",
				"builtin-modules":   "3.0.0",
				"read-package-tree": "^5.2.1",
				"resolve":           "^1.7.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", gitLabPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				gitLabPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				gitLabPkg: "GitLab",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
