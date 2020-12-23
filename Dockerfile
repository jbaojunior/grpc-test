FROM golang:1.15-buster as build

RUN mkdir /tmp/binaries

COPY . .

WORKDIR /go/grpc-client
RUN go build -o grpc-client . && chmod +x grpc-client && cp grpc-client /tmp/binaries/

WORKDIR /go/grpc-server
RUN go build -o grpc-server . && chmod +x grpc-server && cp grpc-server /tmp/binaries/

FROM ubuntu
COPY --from=build /tmp/binaries /tmp
RUN mv /tmp/grpc-* /usr/local/bin/

ENTRYPOINT ["/bin/bash", "-c"]
CMD ["/usr/local/bin/grpc-server"]