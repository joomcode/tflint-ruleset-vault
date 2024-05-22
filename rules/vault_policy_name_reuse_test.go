package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_VaultPolicyNameReuseRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no issue found",
			Content: `
resource "vault_policy" "foo" {
    name = "foo"
}

resource "vault_policy" "bar" {
    name = "bar"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "issue found",
			Content: `
resource "vault_policy" "foo" {
    name = "foo"
}

resource "vault_policy" "foo_dup" {
    name = "foo"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewVaultPolicyNameReuseRule(),
					Message: `vault_policy with name "foo" redeclared, previously used at resource.tf:3,12-17`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 12},
						End:      hcl.Pos{Line: 7, Column: 17},
					},
				},
			},
		},
	}

	rule := NewVaultPolicyNameReuseRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
