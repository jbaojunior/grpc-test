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

&nbsp;&nbsp;&nbsp;&nbsp;SERVER_TLS_ENABLE: Active TLS on client. If you pretend use Ningx ingress to test using TLS you need use this parameter.

**Execute:**
```
docker run -it --rm --name grpc-client jbaojunior/grpc-test grcp-client
``` 

If want do a looping to test some kind of LB:
```
docker run -it --rm --name grpc-client -e SERVER_ADDRESS=${SERVER_ADDRESS} --entrypoint /bin/bash jbaojunior/grpc-test -c 'while true; do /usr/local/bin/grpc-client; sleep 0.5; done'
``` 

### Build docker
```
docker build . -t grpc-test
```

### Deploy Kubernetes
To do tests was created a certificate and a key using the domain casa.com. This address is just to test and does not reflect anything.

Start creating a namespaces:
```
kubectl create grpc
```

Deploy the resource to namespace:
```
kubectl -n grpc -f k8s_resources
```

To test on Kubernetes using ingress active the env SERVER_TLS_ENABLE:
```
SERVER_ADDRESS=grpc-test.casa.com SERVER_PORT=443 SERVER_TLS_ENABLE=true grpc_client
```

#### Update proto
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/grpc-test.proto
```

*Using https://grpc.io/docs/languages/go/quickstart and https://github.com/grpc/grpc-go as reference*
