terraform {
  required_providers {
    prefect = {
      version = "0.2"
      source  = "miandevops/edu/prefect"
    }
  }
}

provider "prefect" {

  url = "http[s]://<your-prefect>"
}