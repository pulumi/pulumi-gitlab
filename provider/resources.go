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
	"context"
	"fmt"
	"path/filepath"
	"unicode"

	// The linter requires unnamed imports to have a doc comment
	_ "embed"

	gitlabShim "gitlab.com/gitlab-org/terraform-provider-gitlab/shim"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-gitlab/provider/v6/pkg/version"
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

// gitLabResource manufactures a standard resource token given a module and resource name.
// It automatically uses the GitLab package and names the file by simply lower casing the resource's
// first character.
func gitLabResource(res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return gitLabType(gitLabMod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-gitlab/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the GitLab package.
func Provider() tfbridge.ProviderInfo {
	p := pfbridge.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(gitlabShim.SDKProvider()),
		gitlabShim.PFProvider(),
	)
	prov := tfbridge.ProviderInfo{
		P:                p,
		Name:             "gitlab",
		Description:      "A Pulumi package for creating and managing GitLab resources.",
		Keywords:         []string{"pulumi", "gitlab"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		GitHubOrg:        "gitlabhq",
		Repository:       "https://github.com/pulumi/pulumi-gitlab",
		UpstreamRepoPath: "./upstream",
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		Version:          version.Version,

		Config: map[string]*tfbridge.SchemaInfo{
			"cacert_file": {},
			"insecure":    {},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Both sshkey and gpgkey have non-standard capitalization.
			"gitlab_user_sshkey": {Tok: gitLabResource("UserSshKey")},
			"gitlab_user_gpgkey": {Tok: gitLabResource("UserGpgKey")},
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
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				Requires: map[string]string{
					"pulumi": ">=3.0.0,<4.0.0",
				}}
			i.PyProject.Enabled = true
			return i
		})(),

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

	prov.MustComputeTokens(tks.SingleModule("gitlab_", gitLabMod, tks.MakeStandard(gitLabPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
