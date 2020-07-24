package main

import (
	"log"
	"net/http"
	"time"
	"flag"
	"os"

	"github.com/myzhan/boomer"
)

var targetUrl = ""

func getIndex() {
	start := time.Now()

	url := string(targetUrl)
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	elapsed := time.Since(start)

	if resp.Status == "200 OK" {
		boomer.RecordSuccess("Get", "/", elapsed.Nanoseconds()/int64(time.Millisecond), resp.ContentLength)
	} else if resp.Status == "201 Created" {
		boomer.RecordSuccess("Get", "/", elapsed.Nanoseconds()/int64(time.Millisecond), resp.ContentLength)
	} else {
		boomer.RecordFailure("Get", "/", elapsed.Nanoseconds()/int64(time.Millisecond), resp.Status)
	}
}

func main() {

	flag.Parse()

	argsWithPort := os.Args

	targetUrl = argsWithPort[3]

	log.Println("Boomer target url is", targetUrl)

	task := &boomer.Task{
		Name:   "getIndex",
		Fn:     getIndex,
	}

	boomer.Run(task)
}