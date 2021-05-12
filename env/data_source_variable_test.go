package env

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testDataSourceVariable = `
data "env_variable" "path" {
  variable = "PATH"
}
`

func TestDataSourceVariable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"env": func() (*schema.Provider, error) { return Provider(), nil },
		},
		Steps: []resource.TestStep{
			{
				Config: testDataSourceVariable,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("data.env_variable.path", "variable")),
			},
		},
	})
}
