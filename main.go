package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/joomcode/tflint-ruleset-vault/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "vault",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewVaultPolicyNameReuseRule(),
			},
		},
	})
}
