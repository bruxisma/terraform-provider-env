---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env Provider"
subcategory: ""
description: |-
  
---

# env Provider



## Example Usage

```terraform
provider "env" { }

data "env_path" "path" {
  variable = "PATH"
}

data "env_variable" "home" {
  variable = "HOME"
}
```

<!-- schema generated by tfplugindocs -->
## Schema
