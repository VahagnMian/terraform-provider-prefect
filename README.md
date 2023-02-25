# Terraform Provider for Prefect

## About 
Prefect is a data orchestration and observation tool that helps you manage and monitor your data workflows.

### Supported Resources
 - Work Queue

Provider Initialization

```
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
```

To use this provider, you need to specify the version and source of the provider in the required_providers block. The prefect block specifies the URL of your Prefect instance.

## Work Queue Creation

```
resource "prefect_work_queue" "wq" {
    name        = "work_queue1"
    description = "description1"
    concurrency_limit = 50
}
```
You can use the prefect_work_queue resource to create a new work queue. In this example, we create a new work queue with the name work_queue1, a description of description1, and a concurrency limit of 50.

## Getting All Work Queues
```
data "prefect_work_queues" "work_queues" {}

output "work_queues" {
    value = data.prefect_work_queues.work_queues
}
```

The prefect_work_queues data source allows you to retrieve all of the work queues that exist on your Prefect instance. In this example, we retrieve all work queues and output them.

Hope you find this provider helpful!

Note: Prefect-client-go was in my different repository, but development of this client will be continued in this repository

Old repository
```
https://github.com/VahagnMian/prefect-client-go
```