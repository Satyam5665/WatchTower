 ![image](https://github.com/user-attachments/assets/491c622c-a69d-438e-a173-365482c0fcdb)

# 🚨 WatchTower

**Real-Time Anomaly Detection & Alerting System** built with Go, Apache Kafka, Redis, Docker.

> WatchTower keeps a vigilant eye on your systems — ingesting telemetry in real-time, detecting anomalies instantly, and triggering alerts before things go wrong.

---

## ⚙️ What It Does

* 📡 Ingests telemetry data from multiple sources via Apache Kafka
* 🤖 Processes data with custom anomaly detection logic in Go
* 🔔 Publishes real-time alerts through Redis Pub/Sub
* 🧾 Logs alerts in Redis Streams for durable tracking
* 🐳 Fully containerized with Docker and orchestrated using Docker Compose

---

## 🧠 Why WatchTower?

* **Real-Time**: Built for speed and low latency
* **Scalable**: Decoupled architecture using Kafka and Redis
* **Extensible**: Add ML models or new alert channels easily

---
Future Work : To integerate Prometheus for Data visualization.
