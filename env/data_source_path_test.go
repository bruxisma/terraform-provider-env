package env

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testDataSourcePathExists = `
data "env_path" "path" {
	variable = "TERRAFORM_PROVIDER_ENV"
}
`

const testDataSourcePathMissing = `
data "env_path" "path" {
	variable = "TERRAFORM_PROVIDER_ENV_THAT_WE_DO_NOT_SET"
}
`

func TestDataSourcePath(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		IsUnitTest:        true,
		PreCheck:          func() { testPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testDataSourcePathExists,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_path.path", "exists", "true"),
					resource.TestCheckResourceAttr("data.env_path.path", "paths.0", testValue[0]),
					resource.TestCheckResourceAttr("data.env_path.path", "paths.1", testValue[1]),
					resource.TestCheckResourceAttr("data.env_path.path", "paths.2", testValue[2]),
					resource.TestCheckResourceAttrSet("data.env_path.path", "id")),
			},
			{
				Config: testDataSourcePathMissing,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_path.path", "exists", "false"),
					resource.TestCheckNoResourceAttr("data.env_path.path", "paths"),
					resource.TestCheckResourceAttrSet("data.env_path.path", "id")),
			},
		},
	})
}
