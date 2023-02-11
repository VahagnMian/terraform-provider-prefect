data "prefect_work_queues" "work_queues" {

}

output "work_queues" {
  value = data.prefect_work_queues.work_queues
}