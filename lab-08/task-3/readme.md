🚀 Task 3 — CI/CD Pipeline (GitHub Actions + Terraform)
🧠 What You’re Building

When you push code to GitHub:

👉 GitHub Actions will:

Install Terraform
Run terraform init
Run terraform plan
Run terraform apply
Deploy your Kubernetes resource (e.g., MongoDB pod/service)
🏗️ Step 1 — Prepare Terraform Code

Make sure your Terraform actually deploys MongoDB (not just nginx).

Example: main.tf
provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_pod" "mongodb" {
  metadata {
    name = "mongodb-pod"
    labels = {
      app = "mongodb"
    }
  }

  spec {
    container {
      name  = "mongodb"
      image = "mongo:6.0"

      port {
        container_port = 27017
      }
    }
  }
}

resource "kubernetes_service" "mongodb_service" {
  metadata {
    name = "mongodb-service"
  }

  spec {
    selector = {
      app = "mongodb"
    }

    port {
      port        = 27017
      target_port = 27017
    }

    type = "NodePort"
  }
}
🔐 Step 2 — Setup GitHub Secrets

Go to your GitHub repo:

👉 Settings → Secrets → Actions

Add:

Name	Value
KUBE_CONFIG	Your kubeconfig file content
Get kubeconfig:
cat ~/.kube/config

Copy full output → paste into secret

⚙️ Step 3 — Create GitHub Actions Workflow

Create file:

.github/workflows/terraform.yml
✅ Full Working Workflow
name: Terraform CI/CD

on:
  push:
    branches:
      - main

jobs:
  terraform:
    name: Deploy using Terraform
    runs-on: ubuntu-latest

    steps:
    
    - name: Checkout Code
      uses: actions/checkout@v3

    - name: Setup Terraform
      uses: hashicorp/setup-terraform@v2

    - name: Setup Kubeconfig
      run: |
        mkdir -p ~/.kube
        echo "${{ secrets.KUBE_CONFIG }}" > ~/.kube/config

    - name: Terraform Init
      run: terraform init

    - name: Terraform Plan
      run: terraform plan

    - name: Terraform Apply
      run: terraform apply -auto-approve
🔄 Step 4 — Push Code
git add .
git commit -m "Add CI/CD pipeline"
git push origin main
👀 Step 5 — Verify Pipeline

Go to your GitHub repo:

👉 Actions tab

You should see:
✔ Workflow triggered
✔ Terraform executed
✔ Deployment success

🧪 Step 6 — Verify Kubernetes

Run locally:

kubectl get pods
kubectl get svc

You should see:

mongodb-pod
mongodb-service
🔗 Step 7 — Test MongoDB (Optional)

Port forward:

kubectl port-forward pod/mongodb-pod 27017:27017

Then connect:

mongosh