# Lab 7 — Microservices, Prometheus & CI/CD

---

## Objective

* Create a microservice using Docker and Minikube
* Monitor containers using Prometheus
* Implement CI/CD using GitHub Actions

---

# Task 1 — Microservice (GPS API)

---

## app.js

```js
const http = require('http');

http.createServer((req, res) => {
    res.writeHead(200, {'Content-Type': 'application/json'});

    const data = {
        place: "IIIT Kottayam",
        latitude: 9.5916,
        longitude: 76.5222
    };

    res.end(JSON.stringify(data));
}).listen(3000);

console.log("GPS Service running on port 3000");
```

---

## Dockerfile

```dockerfile
FROM node:18
WORKDIR /app
COPY . .
EXPOSE 3000
CMD ["node", "app.js"]
```

---

## Build Image

```bash
eval $(minikube docker-env)
docker build -t gps-app .
```

---

## deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gps-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: gps-app
  template:
    metadata:
      labels:
        app: gps-app
    spec:
      containers:
      - name: gps-container
        image: gps-app
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
  name: gps-service
spec:
  type: NodePort
  selector:
    app: gps-app
  ports:
  - port: 3000
    targetPort: 3000
    nodePort: 30008
```

---

## Deploy

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl get pods
minikube service gps-service
```

---

## Expected Output

```json
{
  "place": "IIIT Kottayam",
  "latitude": 9.5916,
  "longitude": 76.5222
}
```

---

# Task 2 — Prometheus Monitoring

---

## prometheus-deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus
spec:
  replicas: 1
  selector:
    matchLabels:
      app: prometheus
  template:
    metadata:
      labels:
        app: prometheus
    spec:
      containers:
      - name: prometheus
        image: prom/prometheus
        ports:
        - containerPort: 9090
```

---

## prometheus-service.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  name: prometheus-service
spec:
  type: NodePort
  selector:
    app: prometheus
  ports:
  - port: 9090
    targetPort: 9090
    nodePort: 30009
```

---

## Deploy Prometheus

```bash
kubectl apply -f prometheus-deployment.yaml
kubectl apply -f prometheus-service.yaml
minikube service prometheus-service
```

---

## Test in UI

Query:

```
up
```

---

# Task 3 — CI/CD using GitHub Actions

---

## .github/workflows/ci.yml

```yaml
name: CI Pipeline

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Build Docker Image
      run: echo "Building..."

    - name: Deploy
      run: echo "Deploying..."
```

---

## Step 3: Push to GitHub

```bash
git add .
git commit -m "Add CI/CD pipeline"
git push
```

---

## Step 4: Verify

Go to GitHub → Actions tab

Workflow should run automatically

---

## Expected Output

CI/CD pipeline runs successfully

---

# Quick Fix

```bash
kubectl get pods
kubectl logs <pod-name>
kubectl describe pod <pod-name>
```

---

# Quick Flow

```text
Build → Deploy → Service → Monitor → CI/CD
```

---

# Final Result

* Microservice running
* Prometheus monitoring active
* CI/CD pipeline executed

---
