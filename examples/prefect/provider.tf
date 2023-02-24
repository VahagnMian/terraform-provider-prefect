terraform {
  required_providers {
    prefect = {
      version = "0.2"
      source  = "miandevops/edu/prefect"
    }
  }
}

provider "prefect" {
  //url = "localhost:4300"
  url = "https://prefect.dev.ben-energy.com"
}