data "prefect_work_queues" "work_queues" {

}

resource "prefect_work_queue" "wq" {
    name        = "workqueue2"
    description = "Description2"
    concurrency_limit = 50
}

resource "prefect_work_queue" "wq1" {
    name        = "workqueue3"
    description = "Description3"
    concurrency_limit = 50
}

output "work_queues" {
    value = data.prefect_work_queues.work_queues
}
