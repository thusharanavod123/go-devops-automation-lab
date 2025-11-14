



# 🚀 Go DevOps Observability Starter

A beginner-friendly DevOps project built using **Golang**, designed to teach how real-world observability works.
This project simulates **CPU, Memory, and Disk metrics**, exposes them via an HTTP endpoint, and prepares you for full integration with **Prometheus & Grafana**.

---

## 🎯 **What This Project Does (Casually Explained)**

This is a simple Go application that:

* Pretends to be a real server
* Generates **fake CPU, memory, and disk usage metrics**
* Exposes them on:

  ```
  http://localhost:8080/metrics
  ```
* Prints new system states continuously (like DevOps monitoring dashboards)
* Prepares you to integrate with:

    * Prometheus (metrics scraping)
    * Grafana (visual dashboards)
    * Docker + Kubernetes later

Think of it as your **first tiny “observability microservice”**.

---

## 🧰 **Tech Stack**

| Component                                | Purpose                                       |
| ---------------------------------------- | --------------------------------------------- |
| **Go (Golang)**                          | Build lightweight, fast monitoring services   |
| **HTTP Server**                          | Expose metrics endpoint                       |
| **JSON Output**                          | Standard observability format                 |
| **Simulated Metrics**                    | Helps understand real DevOps monitoring flows |
| **Future Support: Prometheus + Grafana** | Real monitoring dashboards                    |

---

## 📁 **Project Structure**

```
go-observability-starter/
│── main.go
│── go.mod
│── README.md
```

---

## ⚙️ **How to Run**

### 1️⃣ Clone the repository

```bash
git clone https://github.com/yourusername/go-observability-starter.git
cd go-observability-starter
```

### 2️⃣ Run the Go program

```bash
go run main.go
```

### 3️⃣ Visit metrics in browser

Open:

```
http://localhost:8080/metrics
```

### 4️⃣ Example Output

```
{
  "cpu": 74.88,
  "memory": 1286,
  "disk": 75.65
}
```

---

## 🔮 **Next Stage of This Project (Roadmap)**

This project will evolve into a fully functional DevOps monitoring environment.

### ✔ Phase 1: (Done)

* Go metrics generator
* Local HTTP server

### 🚀 Phase 2: (Next Step)

**Integrate Prometheus**

* Prometheus scrapes `/metrics`

### 📊 Phase 3:

**Add Grafana Dashboard**

* Visual CPU/Mem/Disk charts

### 🐳 Phase 4:

**Docker + Docker Compose**

* Prometheus
* Grafana
* Your Go app
  running together

### ☁️ Phase 5:

**Deploy to AWS EC2**

* Systemd service
* Reverse proxy
* Production monitoring

### 🧠 Phase 6:

**Add Real System Metrics**
Replace simulated values with:

* Actual CPU usage
* Actual RAM
* Actual Disk usage
  (using gopsutil)

---

## 🧩 **Why This Project Matters**

This small project teaches the foundations of:

* DevOps monitoring
* Observability
* Metrics pipelines
* Go microservices
* Cloud deployment patterns

These concepts are used daily in companies like **Uber, Netflix, Meta, Tesla, Google**.

---

## 👐 **Contributing**

Contributions are welcome!

Ways to help:

* Improve the Go code
* Add real system metric collectors
* Add Prometheus exporter format
* Create dashboards
* Write documentation

Don’t forget to ⭐ star the repo if you find it helpful!

---

