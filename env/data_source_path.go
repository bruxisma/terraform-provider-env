package env

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePath() *schema.Resource {
	return &schema.Resource{
		Description: "An environment variable for a list of paths from the local machine",
		ReadContext: dataSourcePathRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "For internal use only",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"variable": &schema.Schema{
				Description: "The environment variable to lookup",
				Type:        schema.TypeString,
				Required:    true,
			},
			"exists": &schema.Schema{
				Description: "Whether the environment variable exists",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"paths": &schema.Schema{
				Description: "The environment variable as a list of paths",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourcePathRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	value, exists := os.LookupEnv(data.Get("variable").(string))
	paths := strings.Split(value, string(os.PathListSeparator))
	hash := sha256.Sum256([]byte(value))
	id := hex.EncodeToString(hash[:])
	if !exists {
		paths = []string{}
	}
	data.SetId(id)
	data.Set("exists", exists)
	data.Set("paths", paths)
	return nil
}
