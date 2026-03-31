const http = require('http');

// Create server
const server = http.createServer((req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.write("Hello from Node.js Server");
    res.end();
});

// Listen on port
server.listen(3000, () => {
    console.log("Server running on port 3000");
});