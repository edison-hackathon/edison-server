package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tarantool/go-tarantool"
)

func main() {
	server := "100.100.157.192:3302"
	opts := tarantool.Opts{
		Timeout:       500 * time.Millisecond,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
	}
	tarantoolConn, err := tarantool.Connect(server, opts)
	if err != nil {
		log.Fatalf("Failed to connect: %s", err.Error())
	}

	h := handlers{tarantoolConn: tarantoolConn}
	http.HandleFunc("/api/v1/devices", h.getDevices)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
