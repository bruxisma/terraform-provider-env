package env

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testDataSourceVariableExists = `
data "env_variable" "path" {
  variable = "TERRAFORM_PROVIDER_ENV"
}
`

const testDataSourceVariableMissing = `
data "env_variable" "path" {
	variable = "TERRAFORM_PROVIDER_ENV_THAT_WE_DO_NOT_SET"
}
`

func TestDataSourceVariable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		IsUnitTest:        true,
		PreCheck:          func() { testPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testDataSourceVariableExists,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_variable.path", "exists", "true"),
					resource.TestCheckResourceAttr("data.env_variable.path", "value", testPaths),
					resource.TestCheckResourceAttrSet("data.env_variable.path", "id")),
			},
			{
				Config: testDataSourceVariableMissing,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_variable.path", "exists", "false"),
					resource.TestCheckResourceAttr("data.env_variable.path", "value", ""),
					resource.TestCheckResourceAttrSet("data.env_variable.path", "id")),
			},
		},
	})
}
