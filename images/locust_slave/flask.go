package main

import (
	"log"
	"net/http"
	"time"

	"github.com/myzhan/boomer"
)

func getIndex() {
	start := time.Now()
	resp, err := http.Get("http://flask-demo:5000/")

	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	if resp.Status == "200 OK" {
		boomer.RecordSuccess("Get", "/", elapsed.Nanoseconds()/int64(time.Millisecond), resp.ContentLength)
	} else {
		boomer.RecordFailure("Get", "/", elapsed.Nanoseconds()/int64(time.Millisecond), resp.Status)
	}
}

func getText() {
	start := time.Now()
	resp, err := http.Get("http://flask-demo:5000/text")

	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	if resp.Status == "200 OK" {
		boomer.RecordSuccess("Get", "/text", elapsed.Nanoseconds()/int64(time.Millisecond), resp.ContentLength)
	} else {
		boomer.RecordFailure("Get", "/text", elapsed.Nanoseconds()/int64(time.Millisecond), resp.Status)
	}
}

func main() {

	task1 := &boomer.Task{
		Name:   "getIndex",
		Weight: 10,
		Fn:     getIndex,
	}

	task2 := &boomer.Task{
		Name:   "getText",
		Weight: 30,
		Fn:     getText,
	}

	boomer.Run(task1, task2)
}