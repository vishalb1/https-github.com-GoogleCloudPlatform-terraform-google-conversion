// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/caiasset"
	"github.com/google/go-cmp/cmp"
)

// TestCLI tests the "convert" and "validate" subcommand against a generated .tfplan file.
func TestCLI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
		return
	}

	// Test cases for each type of resource is defined here.
	cases := []struct {
		name                 string
		compareConvertOutput compareConvertOutputFunc
	}{
		{name: "example_project_iam_binding", compareConvertOutput: compareMergedIamBindingOutput},
		{name: "example_project_iam_member", compareConvertOutput: compareMergedIamMemberOutput},
	}

	for i := range cases {
		// Allocate a variable to make sure test can run in parallel.
		c := cases[i]

		// Add default convert comparison func if not set
		if c.compareConvertOutput == nil {
			c.compareConvertOutput = compareUnmergedConvertOutput
		}

		// Test both offline and online mode.
		for _, offline := range []bool{true, false} {
			offline := offline
			t.Run(fmt.Sprintf("v=0.12/tf=%s/offline=%t", c.name, offline), func(t *testing.T) {
				t.Parallel()
				// Create a temporary directory for running terraform.
				dir, err := os.MkdirTemp(tmpDir, "terraform")
				if err != nil {
					log.Fatal(err)
				}
				defer os.RemoveAll(dir)

				// Generate the <name>.tf and <name>_assets.json files into the temporary directory.
				generateTestFiles(t, "../testdata/templates", dir, c.name+".tf")
				generateTestFiles(t, "../testdata/templates", dir, c.name+".json")

				// Uses glob matching to match generateTestFiles internals.
				tfstateMatches, err := filepath.Glob(filepath.Join("../testdata/templates", c.name+".tfstate"))
				if err != nil {
					t.Fatalf("malformed glob: %v", err)
				}
				if tfstateMatches != nil {
					generateTestFiles(t, "../testdata/templates", dir, c.name+".tfstate")
					err = os.Rename(
						filepath.Join(dir, c.name+".tfstate"),
						filepath.Join(dir, "terraform.tfstate"),
					)
					if err != nil {
						t.Fatalf("renaming tfstate: %v", err)
					}
				}

				terraformWorkflow(t, dir, c.name)

				t.Run("cmd=convert", func(t *testing.T) {
					testConvertCommand(t, dir, c.name, c.name, offline, true, c.compareConvertOutput)
				})
			})
		}
	}
}

type compareConvertOutputFunc func(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool)

func compareUnmergedConvertOutput(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool) {
	expectedAssets := normalizeAssets(t, expected, offline)
	actualAssets := normalizeAssets(t, actual, offline)
	if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
		t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
	}
}

// For merged IAM members, only consider whether the expected members are present.
func compareMergedIamMemberOutput(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool) {
	var normalizedActual []caiasset.Asset
	for i := range expected {
		expectedAsset := expected[i]
		actualAsset := actual[i]

		// Copy actualAsset
		normalizedActualAsset := actualAsset

		expectedBindings := map[string]map[string]struct{}{}
		for _, binding := range expectedAsset.IAMPolicy.Bindings {
			expectedBindings[binding.Role] = map[string]struct{}{}
			for _, member := range binding.Members {
				expectedBindings[binding.Role][member] = struct{}{}
			}
		}

		iamPolicy := caiasset.IAMPolicy{}
		for _, binding := range actualAsset.IAMPolicy.Bindings {
			if expectedMembers, exists := expectedBindings[binding.Role]; exists {
				iamBinding := caiasset.IAMBinding{
					Role: binding.Role,
				}
				for _, member := range binding.Members {
					if _, exists := expectedMembers[member]; exists {
						iamBinding.Members = append(iamBinding.Members, member)
					}
				}
				iamPolicy.Bindings = append(iamPolicy.Bindings, iamBinding)
			}
		}
		normalizedActualAsset.IAMPolicy = &iamPolicy
		normalizedActual = append(normalizedActual, normalizedActualAsset)
	}

	expectedAssets := normalizeAssets(t, expected, offline)
	actualAssets := normalizeAssets(t, normalizedActual, offline)
	if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
		t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
	}
}

// For merged IAM bindings, only consider whether the expected bindings are as expected.
func compareMergedIamBindingOutput(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool) {
	var normalizedActual []caiasset.Asset
	for i := range expected {
		expectedAsset := expected[i]
		actualAsset := actual[i]

		// Copy actualAsset
		normalizedActualAsset := actualAsset

		expectedBindings := map[string]struct{}{}
		for _, binding := range expectedAsset.IAMPolicy.Bindings {
			expectedBindings[binding.Role] = struct{}{}
		}

		iamPolicy := caiasset.IAMPolicy{}
		for _, binding := range actualAsset.IAMPolicy.Bindings {
			if _, exists := expectedBindings[binding.Role]; exists {
				iamPolicy.Bindings = append(iamPolicy.Bindings, binding)
			}
		}
		normalizedActualAsset.IAMPolicy = &iamPolicy
		normalizedActual = append(normalizedActual, normalizedActualAsset)
	}

	expectedAssets := normalizeAssets(t, expected, offline)
	actualAssets := normalizeAssets(t, normalizedActual, offline)
	if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
		t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
	}
}
