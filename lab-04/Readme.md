# Lab 4 — Docker Containers

---

## Objective

To explore Docker commands, create a custom Docker image using a Dockerfile, and deploy a Node.js application inside a container.

---

# Installation (Docker)

## Install Docker (Ubuntu / WSL)

```bash
sudo apt update
sudo apt install -y docker.io
```

---

## Start Docker

```bash
sudo systemctl start docker
sudo systemctl enable docker
```

---

## Verify Installation

```bash
docker --version
```

---

## Run without sudo (optional)

```bash
sudo usermod -aG docker $USER
newgrp docker
```

---

# Basic Docker Commands (VERY IMPORTANT)

---

## Pull Image

```bash
docker pull ubuntu
```

## Run Container

```bash
docker run -it ubuntu
```

## List Running Containers

```bash
docker ps
```

## List All Containers

```bash
docker ps -a
```

## List Images

```bash
docker images
```

## Stop Container

```bash
docker stop <container_id>
```

## Remove Container

```bash
docker rm <container_id>
```

## Remove Image

```bash
docker rmi <image_name>
```

---

# Task 2 — Create Docker Image (Simple)

---

## Dockerfile (Dockerfile-simple)

```dockerfile
FROM ubuntu:22.04

LABEL maintainer="Sarthak Yadav"

RUN apt update && apt install -y curl

CMD ["echo", "Hello from Sarthak Yadav Container"]
```

---

## Build Image

```bash
docker build -t myapp .
```

---

## Run Container

```bash
docker run myapp
```

---

## Expected Output

```
Hello from Sarthak Yadav Container
```

---

# Task 3 — Node.js App in Docker

---

## app.js

```js
const http = require('http');

http.createServer((req, res) => {
    res.end("Hello from Docker Node App");
}).listen(3000);

console.log("Server running on port 3000");
```

---

## Dockerfile (Dockerfile-node)

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
docker build -t node-app .
```

---

## Run Container (Port Mapping IMPORTANT)

```bash
docker run -p 3000:3000 node-app
```

---

## Output

Open browser:

```
http://localhost:3000
```

Expected:

```
Hello from Docker Node App
```

---
