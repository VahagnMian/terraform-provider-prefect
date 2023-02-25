# This is example configuration of provider
terraform {
  required_providers {
    prefect = {
      source = "VahagnMian/prefect"
      version = "1.0.0"
    }
  }
}

provider "prefect" {
  #url = "http[s]://<url>"
  url = "http://localhost:4200"
}