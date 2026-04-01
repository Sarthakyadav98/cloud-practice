# Lab 5 — Kubernetes (Minikube)

---

## Objective

Deploy and expose a containerized application using Minikube and Kubernetes, and implement a simple ML-as-a-Service.

---

# Installation (One-time setup)

```bash
sudo apt update
sudo apt install -y docker.io
sudo apt install -y kubectl
```

OR

```bash
chmod +x install.sh
sudo ./install.sh
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

# Task 2 — Deploy Application

---

## deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  replicas: 2
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: node-app
        image: node-app
        imagePullPolicy: Never
        ports:
        - containerPort: 3000
```

---

## service.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  selector:
    app: my-app
  ports:
  - port: 3000
    targetPort: 3000
    nodePort: 30007
```

---

## Build Image (IMPORTANT)

```bash
eval $(minikube docker-env)
docker build -t node-app .
```

---

## Deploy

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

---

## Check Pods

```bash
kubectl get pods
```

---

## Access Service

```bash
minikube service my-service
```

---

# Expected Output

* Pods status: **Running**
* Service opens in browser

---

# Task 3 — ML as a Service

---

## app.py (Simple ML Logic)

```python
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/predict')
def predict():
    x = int(request.args.get('x', 0))
    result = "Even" if x % 2 == 0 else "Odd"
    return jsonify({"input": x, "prediction": result})

app.run(host='0.0.0.0', port=5000)
```

---

## Dockerfile

```dockerfile
FROM python:3.9

WORKDIR /app
COPY . .

RUN pip install flask

EXPOSE 5000

CMD ["python", "app.py"]
```

---

## Build Image

```bash
eval $(minikube docker-env)
docker build -t ml-app .
```

---

## deployment.yaml (ML)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ml-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ml-app
  template:
    metadata:
      labels:
        app: ml-app
    spec:
      containers:
      - name: ml-container
        image: ml-app
        imagePullPolicy: Never
        ports:
        - containerPort: 5000
```

---

## service.yaml (ML)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: ml-service
spec:
  type: NodePort
  selector:
    app: ml-app
  ports:
  - port: 5000
    targetPort: 5000
    nodePort: 30008
```

---

## Deploy ML Service

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

---

## Access ML API

```bash
minikube service ml-service
```

OR

```bash
curl http://<node-ip>:30008/predict?x=5
```

---

## Expected Output

```json
{
  "input": 5,
  "prediction": "Odd"
}
```

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
minikube start → build image → deploy → service → access
```

---
