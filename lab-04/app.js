const http = require('http');
http.createServer((req,res)=>{
    res.end("Hello Docker Node");
}).listen(3000);