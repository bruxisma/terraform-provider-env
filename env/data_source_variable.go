package env

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVariable() *schema.Resource {
	return &schema.Resource{
		Description: "An environment variable from the local machine",
		ReadContext: dataSourceVariableRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "For internal use only",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"exists": &schema.Schema{
				Description: "Whether the environment variable exists",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"value": &schema.Schema{
				Description: "The result of finding the environment variable",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"variable": &schema.Schema{
				Description: "The environment variable to lookup",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceVariableRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	value, exists := os.LookupEnv(data.Get("variable").(string))
	hash := sha256.Sum256([]byte(value))
	data.SetId(hex.EncodeToString(hash[:]))
	data.Set("exists", exists)
	data.Set("value", value)
	return nil
}
