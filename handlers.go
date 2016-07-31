package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tarantool/go-tarantool"
)

type handlers struct {
	tarantoolConn *tarantool.Connection
}

func (h handlers) getDevices(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.ParseUint(
		r.URL.Query().Get("offset"),
		10,
		32,
	)

	limit, _ := strconv.ParseUint(
		r.URL.Query().Get("limit"),
		10,
		32,
	)
	if limit == 0 {
		limit = 10
	}

	indexNo := uint32(0)

	resp, err := h.tarantoolConn.Select(
		"devices",
		indexNo,
		uint32(offset),
		uint32(limit),
		tarantool.IterEq,
		[]interface{}{},
	)

	if err != nil {
		log.Fatalf("Failed to select: %s", err.Error())
	}

	devices, err := parseDevices(resp.Tuples())
	if err != nil {
		log.Fatalf("Cannot parse measurements: %s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	json.NewEncoder(w).Encode(devices)
}
