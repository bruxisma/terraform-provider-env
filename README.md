# Terraform Provider for Environment Variables

This terraform provider allows you to read environment variables (either as
strings or as a list of paths) from your local machine.

Do not use this in multi-person (or even multi-system) deployment mechanisms.
It is meant entirely for debugging/configuration purposes.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.15.x
-	[Go](https://golang.org/doc/install) 1.16.x (to build the provider plugin)

