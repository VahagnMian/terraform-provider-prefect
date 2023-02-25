data "prefect_work_queues" "work_queues" {}

resource "prefect_work_queue" "wq" {
    name        = "work_queue1"
    description = "description1"
    concurrency_limit = 50
}

output "work_queues" {
    value = data.prefect_work_queues.work_queues
}
