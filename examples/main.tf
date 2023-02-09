terraform {
  required_providers {
    prefect = {
      version = "0.2"
      source  = "miandevops/edu/prefect"
    }
  }
}

provider "prefect" {}

module "prefect" {
  source = "./prefect"
}
