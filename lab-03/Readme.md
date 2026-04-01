# Lab 3 — DynamoDB using Node.js & Go

---



# Installation

## Install AWS CLI

```
sudo apt install awscli -y
```

---

## Configure AWS

```
aws configure
```

Enter:

* Access Key
* Secret Key
* Region (e.g., ap-south-1)

---

## Install SDKs

### Node.js

```bash id="4xy59k"
npm init -y
npm install aws-sdk
```

### Go

```bash id="nsbqnb"
go get github.com/aws/aws-sdk-go
```

---

# Operations

* Create table
* Insert / Update item
* Delete table

---

# Run

### Node.js

```bash id="pm8l3u"
node node-dynamodb.js
```

### Go

```bash id="3r1q9h"
go run go-dynamodb.go
```

---

# Expected Output

* Table created
* Item inserted / updated
* Table deleted

---

# Quick Flow

```text id="p6k9o7"
Create Table → Insert/Update → Delete Table
```

---
