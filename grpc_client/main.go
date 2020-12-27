/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"time"

	pb "github.com/jbaojunior/grpc-test/grpctest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultQuestion = "What is the server?"
)

func main() {

	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "5551"
	}

	server, ok := os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		server = "127.0.0.1"
	}
	address := server + ":" + port

	// Verify if TLS is enable and set up a connection to the server
	var conn *grpc.ClientConn
	var err error
	_, ok = os.LookupEnv("SERVER_TLS_ENABLE")
	if !ok {
		conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	}

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcTestClient(conn)

	// Contact the server and print out its response.
	question := defaultQuestion
	/*if len(os.Args) > 1 {
		name = os.Args[1]
	}*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Msg(ctx, &pb.MsgRequest{Server: question})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("%s", r.GetMessage())
}
