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
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-gitlab/gitlab"
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
	p := gitlab.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "gitlab",
		Description: "A Pulumi package for creating and managing GitLab resources.",
		Keywords:    []string{"pulumi", "gitlab"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-gitlab",
		Config: map[string]*tfbridge.SchemaInfo{
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITLAB_TOKEN"},
				},
			},
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITLAB_BASE_URL"},
				},
			},
			"cacert_file": {},
			"insecure":    {},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"gitlab_branch_protection":          {Tok: gitLabResource(gitLabMod, "BranchProtection")},
			"gitlab_tag_protection":             {Tok: gitLabResource(gitLabMod, "TagProtection")},
			"gitlab_group":                      {Tok: gitLabResource(gitLabMod, "Group")},
			"gitlab_group_label":                {Tok: gitLabResource(gitLabMod, "GroupLabel")},
			"gitlab_group_cluster":              {Tok: gitLabResource(gitLabMod, "GroupCluster")},
			"gitlab_project":                    {Tok: gitLabResource(gitLabMod, "Project")},
			"gitlab_project_push_rules":         {Tok: gitLabResource(gitLabMod, "ProjectPushRules")},
			"gitlab_label":                      {Tok: gitLabResource(gitLabMod, "Label")},
			"gitlab_pipeline_schedule":          {Tok: gitLabResource(gitLabMod, "PipelineSchedule")},
			"gitlab_pipeline_trigger":           {Tok: gitLabResource(gitLabMod, "PipelineTrigger")},
			"gitlab_pipeline_schedule_variable": {Tok: gitLabResource(gitLabMod, "PipelineScheduleVariable")},
			"gitlab_project_hook":               {Tok: gitLabResource(gitLabMod, "ProjectHook")},
			"gitlab_deploy_key":                 {Tok: gitLabResource(gitLabMod, "DeployKey")},
			"gitlab_deploy_key_enable":          {Tok: gitLabResource(gitLabMod, "DeployKeyEnable")},
			"gitlab_user":                       {Tok: gitLabResource(gitLabMod, "User")},
			"gitlab_project_membership":         {Tok: gitLabResource(gitLabMod, "ProjectMembership")},
			"gitlab_group_membership":           {Tok: gitLabResource(gitLabMod, "GroupMembership")},
			"gitlab_project_variable":           {Tok: gitLabResource(gitLabMod, "ProjectVariable")},
			"gitlab_project_share_group":        {Tok: gitLabResource(gitLabMod, "ProjectShareGroup")},
			"gitlab_group_variable":             {Tok: gitLabResource(gitLabMod, "GroupVariable")},
			"gitlab_project_cluster":            {Tok: gitLabResource(gitLabMod, "ProjectCluster")},
			"gitlab_service_slack":              {Tok: gitLabResource(gitLabMod, "ServiceSlack")},
			"gitlab_service_jira":               {Tok: gitLabResource(gitLabMod, "ServiceJira")},
			"gitlab_group_ldap_link":            {Tok: gitLabResource(gitLabMod, "GroupLdapLink")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"gitlab_group":   {Tok: gitLabDataSource(gitLabMod, "getGroup")},
			"gitlab_project": {Tok: gitLabDataSource(gitLabMod, "getProject")},
			"gitlab_user":    {Tok: gitLabDataSource(gitLabMod, "getUser")},
			"gitlab_users":   {Tok: gitLabDataSource(gitLabMod, "getUsers")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^2.0.0",
				"builtin-modules":   "3.0.0",
				"read-package-tree": "^5.2.1",
				"resolve":           "^1.7.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
			Overlay: &tfbridge.OverlayInfo{
				DestFiles: []string{},
				Modules:   map[string]*tfbridge.OverlayInfo{},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				gitLabPkg: "GitLab",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameField = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameField]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameField]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameField] = tfbridge.AutoName(nameField, 255)
				}
			}
		}
	}

	return prov
}
