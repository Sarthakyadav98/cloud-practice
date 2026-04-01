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