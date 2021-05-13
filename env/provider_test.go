package env

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"env": func() (*schema.Provider, error) { return Provider(), nil },
}

var testValue = []string{"a", "b", "c"}
var testPaths = strings.Join(testValue[:], string(os.PathListSeparator))

func testPreCheck(t *testing.T) {
	os.Unsetenv("TERRAFORM_PROVIDER_ENV_THAT_WE_DO_NOT_SET")
	os.Setenv("TERRAFORM_PROVIDER_ENV", testPaths)
}
