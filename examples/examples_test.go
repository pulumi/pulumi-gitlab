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
	"fmt"
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	token := os.Getenv("PULUMI_GITLAB_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing PULUMI_GITLAB_TOKEN environment variable")
	}
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		Env:                  []string{fmt.Sprintf("GITLAB_TOKEN=%s", token)},
	}
}
