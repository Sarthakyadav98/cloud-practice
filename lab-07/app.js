const http = require('http');
http.createServer((req,res)=>{
    res.end(JSON.stringify({
        lat: 9.59,
        lon: 76.52
    }));
}).listen(3000);