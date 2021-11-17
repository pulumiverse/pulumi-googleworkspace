package examples

import (
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

func checkGoogleworkspaceCredentials(t *testing.T) {
	_, exists := os.LookupEnv("GOOGLEWORKSPACE_CREDENTIALS")
	if !exists {
		t.Skipf("Skipping test due to missing GOOGLEWORKSPACE_CREDENTIALS environment variable")
	}
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	checkGoogleworkspaceCredentials(t)
	return integration.ProgramTestOptions{
		RunUpdateTest:        false,
		ExpectRefreshChanges: true,
	}
}
