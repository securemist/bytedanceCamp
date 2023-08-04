package main

import (
	"bytedanceCamp/util/initialize"
	"google.golang.org/grpc"
)

func main() {
	initialize.Initialize()

	// 注册grpc服务
	server := grpc.NewServer()
	
}
