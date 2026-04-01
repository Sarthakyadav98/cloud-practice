# Lab 9 — OpenFaaS (Serverless) + CI/CD

---

## Objective

* Install OpenFaaS on Minikube
* Create and deploy a serverless function
* Monitor using Prometheus
* Automate deployment using Terraform + GitHub Actions

---

# Task 1 — OpenFaaS Setup & Function Deployment

---

## Step 1: Start Minikube

```bash
minikube start --driver=docker
```

---

## Step 2: Install CLI Tools

```bash
# Install faas-cli
curl -sSL https://cli.openfaas.com | sudo sh

# Install helm
sudo apt install -y helm
```

---

## Step 3: Add OpenFaaS Helm Repo

```bash
helm repo add openfaas https://openfaas.github.io/faas-netes/
helm repo update
```

---

## Step 4: Create Namespaces

```bash
kubectl create namespace openfaas
kubectl create namespace openfaas-fn
```

---

## Step 5: Install OpenFaaS

```bash
helm upgrade openfaas openfaas/openfaas \
  --install \
  --namespace openfaas \
  --set functionNamespace=openfaas-fn
```

Check pods:

```bash
kubectl get pods -n openfaas
```

---

## Step 6: Port Forward

```bash
kubectl port-forward -n openfaas svc/gateway 8080:8080
```

Open:

```
http://localhost:8080
```

---

## Step 7: Login

Get password:

```bash
kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode
```

Login:

```bash
faas-cli login --username admin --password <password>
```

---

## Step 8: Create Function

Pull templates:

```bash
faas-cli template store pull node
faas-cli template pull
```

Create function:

```bash
faas-cli new hello-fn --lang node18
```

Replace handler.js with:

```js
module.exports = async (event, context) => {
    return "Hello from OpenFaaS Function";
}
```

---

## Step 9: Deploy Function

Set Docker environment to Minikube:

```bash
eval $(minikube docker-env)
```

Deploy:

```bash
faas-cli up -f hello-fn.yml
```

---

## Step 10: Invoke Function

```bash
faas-cli invoke hello-fn
```

---

## Expected Output

```
Hello from OpenFaaS Function
```

Optional: Test in browser:

```
http://localhost:8080/function/hello-fn
```

---

# Monitoring (Prometheus)

Open Prometheus UI and run:

```
gateway_function_invocation_total
```

This shows the number of times the function has been invoked.

---

# Task 2 — CI/CD using Terraform + GitHub Actions

---

## Step 1: Terraform Code

Terraform will deploy and destroy function using faas-cli.

---

## Step 2: GitHub Workflow

Create folder:

```bash
mkdir -p .github/workflows
```

---

## Step 3: Push to GitHub

```bash
git add .
git commit -m "OpenFaaS CI/CD"
git push
```

---

## Expected Output

* Function deployed using Terraform
* Function destroyed automatically
* CI/CD pipeline runs successfully

---

# Quick Fix

```bash
kubectl get pods -n openfaas
kubectl logs <pod-name> -n openfaas
```

---

# Quick Flow

```
Install → Deploy Function → Invoke → Monitor → CI/CD
```

---

# Final Result

* OpenFaaS installed
* Function deployed and executed
* Prometheus monitoring working
* CI/CD pipeline automated

---
