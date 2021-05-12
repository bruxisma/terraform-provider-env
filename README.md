# Terraform Provider for Environment Variables

[![Tests][test-badge]][test-link] [![codecov][coverage-badge]][coverage-link]

This terraform provider allows you to read environment variables (either as
strings or as a list of paths) from your local machine.

Do not use this in multi-person (or even multi-system) deployment mechanisms.
It is meant entirely for debugging/configuration purposes.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.15.x
-	[Go](https://golang.org/doc/install) 1.16.x (to build the provider plugin)

[coverage-badge]: https://codecov.io/gh/slurps-mad-rips/terraform-provider-env/branch/main/graph/badge.svg?token=DWOcB8YHRu
[coverage-link]: https://codecov.io/gh/slurps-mad-rips/terraform-provider-env
[test-badge]: https://github.com/slurps-mad-rips/terraform-provider-env/actions/workflows/test.yml/badge.svg
[test-link]: https://github.com/slurps-mad-rips/terraform-provider-env/actions/workflows/test.yml
