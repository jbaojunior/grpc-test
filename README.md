## GRPC Test

GRPC Test return server's hostname

### GRPC Server
**Env:**<br>
&nbsp;&nbsp;&nbsp;&nbsp;PORT: Port to server do binding. Default is 5551.

**Execute:**
```
docker run -it --rm --name grpc-server -p 5551:5551 jbaojunior/grcp-test
```

### GRPC Client
**Env:**<br>
&nbsp;&nbsp;&nbsp;&nbsp;SERVER_ADDRESS: Server Address. Default is 127.0.0.1<br>
&nbsp;&nbsp;&nbsp;&nbsp;SERVER_PORT: Server port. Default is 5551

**Execute:**
```
docker run -it --rm --name grpc-client jbaojunior/grpc-test grcp-client
``` 

If want do a looping to test some kind of LB:
```
docker run -it --rm --name grpc-client -e SERVER_ADDRESS=${SERVER_ADDRESS} --entrypoint /bin/bash jbaojunior/grpc-test -c 'while true; do /usr/local/bin/grpc-client; sleep 0.5; done'
``` 

*Using https://grpc.io/docs/languages/go/quickstart and https://github.com/grpc/grpc-go as reference*