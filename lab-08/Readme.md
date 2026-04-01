# Lab 8 — Terraform, MongoDB & CI/CD

---

## Objective

* Create Kubernetes resources using Terraform
* Perform MongoDB operations
* Automate deployment using CI/CD (GitHub Actions + Terraform)

---

# Task 1 — Terraform (Kubernetes Resource)

---

## Step 1: Install Terraform

```bash
sudo apt update
sudo apt install -y gnupg software-properties-common
```

```bash
wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | \
sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
```

```bash
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list
```

```bash
sudo apt update
sudo apt install terraform -y
```

---

## Step 2: Verify

```bash
terraform -version
```

---

## Step 3: Start Minikube

```bash
minikube start --driver=docker
```

---

## Step 4: Initialize Terraform

```bash
terraform init
```

---

## Step 5: Plan

```bash
terraform plan
```

---

## Step 6: Apply

```bash
terraform apply
```

Type:

```text
yes
```

---

## Step 7: Verify

```bash
kubectl get pods
```

---

## Step 8: Destroy

```bash
terraform destroy
```

Port forwarding:

```bash
kubectl port-forward pod/nginx-pod-2 8080:80
```

---

# Task 2 — MongoDB Operations

---

## Install MongoDB

```bash
sudo apt update
sudo apt install -y gnupg curl

curl -fsSL https://pgp.mongodb.com/server-7.0.asc | \
sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor

echo "deb [ arch=amd64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | \
sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list

sudo apt update
sudo apt install -y mongodb-org
```

---

## Start MongoDB

```bash
sudo systemctl start mongodb
```

---

## Open Shell

```bash
mongosh
```

---

## Create Database

```js
use btechDB
```

---

## Insert Data

```js
db.success.insertOne({
  name: "Sarthak",
  achievement: "Cloud Lab Completed",
  year: 2026
})
```

---

## View Data

```js
db.success.find()
```

---

## Delete Data

```js
db.success.deleteOne({ name: "Sarthak" })
```

---

# Task 3 — CI/CD using GitHub Actions

---

## Step 1: Create Workflow Folder

```bash
mkdir -p .github/workflows
```

---

## Step 2: Create Workflow File

File: `.github/workflows/terraform.yml`

---

## Step 3: Push Code

```bash
git add .
git commit -m "Add Terraform CI/CD"
git push
```

---

## Expected Output

* Terraform creates Kubernetes resource
* MongoDB operations executed
* GitHub Actions pipeline runs successfully

---

# Common Errors

```bash
terraform init
kubectl get pods
```

---

# Quick Flow

```text
Terraform → MongoDB → CI/CD
```

---

# Final Result

* Kubernetes resource created using Terraform
* MongoDB operations successful
* CI/CD pipeline triggered

---
