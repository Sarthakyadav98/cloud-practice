# 🧪 Lab 8 — Terraform, MongoDB & CI/CD

---
## 🎯 Objective

* Create Kubernetes resources using Terraform
* Perform MongoDB operations
* Automate deployment using CI/CD (GitHub Actions + Terraform)

---

# ⚙️ Task 1 — Terraform (Kubernetes Resource)

---

## 🔹 Step 1: Install Terraform

```bash id="l7f4b0"
sudo apt update
sudo apt install -y gnupg software-properties-common
```

```bash id="c1xq0r"
wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | \
sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
```

```bash id="jz8b8m"
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list
```

```bash id="3s2x9f"
sudo apt update
sudo apt install terraform -y
```

---

## 🔹 Step 2: Verify

```bash id="o5g5dp"
terraform -version
```

---

## 🔹 Step 3: Start Minikube

```bash id="g9b1zi"
minikube start --driver=docker
```

---

## 🔹 Step 4: Initialize Terraform

```bash id="6b6p9o"
terraform init
```

---

## 🔹 Step 5: Plan

```bash id="q91gqe"
terraform plan
```

---

## 🔹 Step 6: Apply

```bash id="xejp9d"
terraform apply
```

Type:

```text id="z6u2j2"
yes
```

---

## 🔹 Step 7: Verify

```bash id="2wcc5i"
kubectl get pods
```

---

## 🔹 Step 8: Destroy

```bash id="c4m7c7"
terraform destroy
```

Port forwarding - 
kubectl port-forward pod/nginx-pod-2 8080:80


---

# ⚙️ Task 2 — MongoDB Operations

---

## 🔹 Install MongoDB

```bash id="fzzvwn"
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

## 🔹 Start MongoDB

```bash id="rph0tz"
sudo systemctl start mongodb
```

---

## 🔹 Open Shell

```bash id="zqlq3k"
mongosh
```

---

## 🔹 Create Database

```js id="0c6c5u"
use btechDB
```

---

## 🔹 Insert Data

```js id="a7pdjz"
db.success.insertOne({
  name: "Sarthak",
  achievement: "Cloud Lab Completed",
  year: 2026
})
```

---

## 🔹 View Data

```js id="lql6b9"
db.success.find()
```

---

## 🔹 Delete Data

```js id="fj9q2c"
db.success.deleteOne({ name: "Sarthak" })
```

---

# ⚙️ Task 3 — CI/CD using GitHub Actions

---

## 🔹 Step 1: Create Workflow Folder

```bash id="s1y3zy"
mkdir -p .github/workflows
```

---

## 🔹 Step 2: Create Workflow File

File: `.github/workflows/terraform.yml`

---

## 🔹 Step 3: Push Code

```bash id="y2eq9v"
git add .
git commit -m "Add Terraform CI/CD"
git push
```

---

## 📤 Expected Output

* Terraform creates Kubernetes resource
* MongoDB operations executed
* GitHub Actions pipeline runs successfully

---

# ⚠️ Common Errors

```bash id="0v4d2k"
terraform init
kubectl get pods
```

---

# 💡 Quick Flow

```text id="a1jz6o"
Terraform → MongoDB → CI/CD
```

---

# ✅ Final Result

* Kubernetes resource created using Terraform
* MongoDB operations successful
* CI/CD pipeline triggered

---
