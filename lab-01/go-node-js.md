# Lab 1 — Go & Node.js Basics

## Objective

To understand basic syntax of Go and Node.js and create simple web services using both.

---

# Installation

## Install Node.js

```bash
sudo apt update
sudo apt install -y nodejs npm
```

## Verify

```bash
node -v
npm -v
```

---

## Install Go (Golang)

```bash
sudo apt install -y golang-go
```

## Verify

```bash
go version
```

---

# Node.js Web Server

## Code (node-server.js)

```js
const http = require('http');

http.createServer((req, res) => {
    res.write("Hello from Node.js Server");
    res.end();
}).listen(3000);

console.log("Server running on port 3000");
```

---

## Run

```bash
node node-server.js
```

## Output

Open in browser:

```
http://localhost:3000
```

Expected Output:

```
Hello from Node.js Server
```

---

# Go Web Server

## Code (go-server.go)

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go Server")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
```

---

## Run

```bash
go run go-server.go
```

## Output

Open in browser:

```
http://localhost:8080
```

Expected Output:

```
Hello from Go Server
```

---

# Expected Result

* Node.js server runs on port 3000
* Go server runs on port 8080
* Both return simple text response in browser

---
