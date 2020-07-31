package main

import (
	"flag"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"time"

	"github.com/myzhan/boomer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var verbose = false
var targetUrl string
var conn *grpc.ClientConn

const defaultName = "world"

func NewConn() *grpc.ClientConn {
	conn, err := grpc.Dial(targetUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func worker() {
	startTime := time.Now()

	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := defaultName

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	
	elapsed := time.Since(startTime)

	if err != nil {
		if verbose {
			log.Fatalf("could not greet: %v", err)
		}
		boomer.RecordFailure("grpc", "helloworld.proto", 0.0, err.Error())
	} else {
		if verbose {
			log.Println("Resp Time:",elapsed.Nanoseconds()/int64(time.Millisecond), "ms")
			log.Println("message:", r.GetMessage())
		}
		boomer.RecordSuccess("grpc", "helloworld.proto",
			elapsed.Nanoseconds()/int64(time.Millisecond), 10)
	}
}

func main() {

	defer func() {
		_ = conn.Close()
	}()

	flag.StringVar(&targetUrl, "url", "", "url or app:port")
	flag.Parse()

	if targetUrl == "" {
		log.Println("Boomer target url is null")
		return
	}

	log.Println("Boomer target url is", string(targetUrl))

	conn = NewConn()

	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn:     worker,
	}

	boomer.Run(task)
}