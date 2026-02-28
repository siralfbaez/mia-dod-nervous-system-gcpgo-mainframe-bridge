resource "google_pubsub_topic" "signals" {
  name = "mainframe-signals-topic"
}

resource "google_pubsub_subscription" "worker_sub" {
  name  = "worker-subscription"
  topic = google_pubsub_topic.signals.name

  ack_deadline_seconds = 20
}

resource "google_pubsub_topic" "mainframe_signals" {
  name = "mainframe-signals"
}

resource "google_pubsub_subscription" "worker_subscription" {
  name  = "worker-sub"
  topic = google_pubsub_topic.mainframe_signals.name
}

output "topic_name" {
  value = google_pubsub_topic.mainframe_signals.name
}