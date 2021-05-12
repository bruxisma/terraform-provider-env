package env

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testDataSourceVariable = `
data "env_variable" "path" {
  key = "PATH"
}
`

func TestDataSourceVariable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceVariable,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("data.env_variable.path", "variable")),
			},
		},
	})
}
