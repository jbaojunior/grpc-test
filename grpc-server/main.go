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

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/jbaojunior/grpc-test/grpctest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// server is used to implement GrpcTestServer.
type server struct {
	pb.UnimplementedGrpcTestServer
}

// SayHello implements helloworld.GrpcTestServer
func (s *server) Msg(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	p, _ := peer.FromContext(ctx)
	log.Printf("Received: %v\t ClientAddress: %v", in.GetServer(), p.Addr.String())

	host, _ := os.Hostname()
	message := fmt.Sprintf("Server: %v", host)
	return &pb.MsgReply{Message: message}, nil
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5551"
	}

	port = ":" + port

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Server running on port %v...", port)
	pb.RegisterGrpcTestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
