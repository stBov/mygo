package myrpc

import (
	"net/rpc"
	"github.com/chai2010/pbgo/examples/hello.pb"
	"net"
	"log"
	"fmt"
	"io/ioutil"
	"mime"
)

type HelloService struct{}

func (p *HelloService) Hello(request *hello_pb.String, reply *hello_pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
func (p *HelloService) Echo(request *hello_pb.Message, reply *hello_pb.Message) error {
	*reply = *request
	return nil
}

func (p *HelloService) Static(request *hello_pb.String, reply *hello_pb.StaticFile) error {
	data, err := ioutil.ReadFile("./testdata/" + request.Value)
	if err != nil {
		return err
	}

	reply.ContentType = mime.TypeByExtension(request.Value)
	reply.ContentBody = data
	return nil
}

func startRpcServer() {
	hello_pb.RegisterHelloService(rpc.DefaultServer, new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}

func tryRpcClient() {
	client, err := hello_pb.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := client.Hello(&hello_pb.String{Value: "gopher"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}

func Rpcgo()  {
	go startRpcServer()
	tryRpcClient()
}