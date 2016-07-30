package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tarantool/go-tarantool"
)

func main() {
	server := "100.100.147.43:3301"
	opts := tarantool.Opts{
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
		//User:          "test",
		//Pass:          "test",
	}
	tarantoolConn, err := tarantool.Connect(server, opts)
	if err != nil {
		log.Fatalf("Failed to connect: %s", err.Error())
	}

	h := handlers{tarantoolConn: tarantoolConn}
	http.HandleFunc("/api/v1/measurements", h.getMeasurements)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
