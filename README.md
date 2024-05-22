# TFLint Vault Ruleset

[![Build Status](https://github.com/joomcode/tflint-ruleset-vault/workflows/build/badge.svg?branch=main)](https://github.com/joomcode/tflint-ruleset-vault/actions)

This a small ruleset for TFLint that checks for some mistakes in HashiCorp Vault configurations.

## Requirements

- TFLint v0.42+
- Go v1.22

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "vault" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/joomcode/tflint-ruleset-vault"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  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
  wYeMFkqSsluCjDEL47bGvq0=
  =hItF
  -----END PGP PUBLIC KEY BLOCK-----
  KEY
}
```

## Rules

| Name                    | Description                                                  | Severity | Enabled | Link |
| ----------------------- | ------------------------------------------------------------ | -------- | ------- | ---- |
| vault_policy_name_reuse | Checks if multiple vault_policy resources have the same name | ERROR    | ✔      |      |

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "vault" {
  enabled = true
}
EOS
$ tflint
```
