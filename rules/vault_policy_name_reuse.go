package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// VaultPolicyNameReuseRule checks whether module contains a vault_policy resources
// with the same policy name, which might cause a perpetual diff.
type VaultPolicyNameReuseRule struct {
	tflint.DefaultRule
}

// NewVaultPolicyNameReuseRule returns a new rule
func NewVaultPolicyNameReuseRule() *VaultPolicyNameReuseRule {
	return &VaultPolicyNameReuseRule{}
}

// Name returns the rule name
func (r *VaultPolicyNameReuseRule) Name() string {
	return "vault_policy_reuse"
}

// Enabled returns whether the rule is enabled by default
func (r *VaultPolicyNameReuseRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *VaultPolicyNameReuseRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *VaultPolicyNameReuseRule) Link() string {
	return ""
}

// Check checks whether module contains a vault_policy resources
// with the same policy name, which might cause a perpetual diff.
func (r *VaultPolicyNameReuseRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("vault_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	logger.Debug("Got %d policies", len(resources.Blocks))

	seenPolicies := make(map[string]hcl.Range)

	for _, resource := range resources.Blocks {
		attribute, ok := resource.Body.Attributes["name"]
		if !ok {
			// should not happen, but ignore, it's not our concern
			logger.Warn("found vault_policy resource without name attribute at %v", resource.DefRange)
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(policyName string) error {
			if range_, ok := seenPolicies[policyName]; ok {
				return runner.EmitIssue(
					r,
					fmt.Sprintf(
						`vault_policy with name "%s" redeclared, previously used at %v`,
						policyName, range_,
					),
					attribute.Expr.Range(),
				)
			}

			seenPolicies[policyName] = attribute.Expr.Range()
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
