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

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing GITLAB_TOKEN environment variable")
	}

	baseUrl := os.Getenv("GITLAB_BASE_URL")
	if baseUrl == "" {
		t.Skipf("Skipping test due to missing GITLAB_BASE_URL environment variable")
	}

	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	var base = integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		Config: map[string]string{},
	}

	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/gitlab",
		},
	})

	tests := []integration.ProgramTestOptions{
		baseJS.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "project")}),
	}

	for _, ex := range tests {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, &example)
		})
	}
}
