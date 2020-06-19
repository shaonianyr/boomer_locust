package main

import (
	"log"
	"net/http"
	"time"

	"github.com/myzhan/boomer"
)

func getDemo() {
	start := time.Now()
	resp, err := http.Get("https://www.baidu.com/")

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


func main() {

	task := &boomer.Task{
		Name: "baidu",
		// The weight is used to distribute goroutines over multiple tasks.
		// Single task not set.
		// Weight: 20,
		Fn: getDemo,
	}

	boomer.Run(task)
}