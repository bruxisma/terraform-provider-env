provider "env" { }

data "env_path" "path" {
  variable = "PATH"
}

data "env_variable" "home" {
  variable = "HOME"
}
