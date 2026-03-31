# Lab 5 — Kubernetes (Minikube)

---

## Objective

Deploy and expose a containerized application using Minikube and Kubernetes.

---

# Installation (One-time setup)

```bash
sudo apt update
sudo apt install -y docker.io
sudo apt install -y kubectl
```

---

## Install Minikube

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

---

# Start Cluster

```bash
minikube start --driver=docker
```

---

# Deployment

```bash
kubectl apply -f deployment.yaml
```

---

# Service

```bash
kubectl apply -f service.yaml
```

---

# Check Pods

```bash
kubectl get pods
```

---

# Access Service

```bash
minikube service my-service
```

---

# Expected Output

* Pods status: **Running**
* Service opens in browser

---

# Quick Fix

```bash
kubectl get pods
kubectl describe pod <pod-name>
kubectl logs <pod-name>
```

---

# Quick Flow

```text
minikube start → deploy → service → get pods → access
```

---
