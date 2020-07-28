package main

import (
	"log"
	"net/http"
	"time"
	"flag"

	"github.com/myzhan/boomer"
)

var targetUrl string

func getIndex() {
	start := time.Now()

	resp, err := http.Get(string(targetUrl))

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

	flag.StringVar(&targetUrl, "url", "", "url or app:port")
	flag.Parse()

	if targetUrl == "" {
		log.Println("Boomer target url is null")
		return
	}

	log.Println("Boomer target url is", string(targetUrl))

	task := &boomer.Task{
		Name:   "getIndex",
		Fn:     getIndex,
	}

	boomer.Run(task)
}