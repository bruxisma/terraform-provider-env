package env

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testDataSourceVariable = `
data "env_variable" "path" {
  variable = "TERRAFORM_PROVIDER_ENV"
}
`

func TestDataSourceVariable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		IsUnitTest:        true,
		PreCheck:          func() { testPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testDataSourceVariable,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_variable.path", "exists", "true"),
					resource.TestCheckResourceAttr("data.env_variable.path", "value", testPaths),
					resource.TestCheckResourceAttrSet("data.env_variable.path", "id")),
			},
		},
	})
}
