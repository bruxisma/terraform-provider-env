package env

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"env_variable": dataSourceVariable(),
			"env_path":     dataSourcePath(),
		},
	}
}
