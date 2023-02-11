terraform {
  required_providers {
    prefect = {
      version = "0.2.0"
      source  = "miandevops/edu/prefect"
    }
  }
}

provider "prefect" {}

module "prefect" {
  source = "./prefect"
}

output "work_queues" {
  value = module.prefect.work_queues
}
