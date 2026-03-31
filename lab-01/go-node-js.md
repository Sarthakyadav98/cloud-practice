## Lab 1 - Go & Node.js

### Node Server

```
const http = require('http');
http.createServer((req, res) => {
    res.end("Hello Node");
}).listen(3000);
```

Run:
node node-server.js



```go
package main
import (
    "fmt"
    "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Go")
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Run:
go run go-server.go