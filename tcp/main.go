package main

import (
	"flag"
	"network-learn/tcp/client"
	"network-learn/tcp/learn-server"
)

var mode = flag.String("mode", "", "启动服务端还是客户端")

func main()  {
	flag.Parse()
	//服务端
	if *mode == "server" {
		server := learn_server.NewServer()
		server.Start()
	}

	//客户端
	if *mode == "client" {
		client.Run()
	}
}
