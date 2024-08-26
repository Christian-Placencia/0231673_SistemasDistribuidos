package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

var commitLog = &Log{}

// Add new record to the log.
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read a record from the log by offset.
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, http.ErrNoLocation
	}
	return c.records[offset], nil
}

// Endpoint to add a new record to the log.
func handleProduce(w http.ResponseWriter, r *http.Request) {
	var record Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := commitLog.Append(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]uint64{"offset": offset}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Endpoint to read a record from the log by offset.
func handleConsume(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Offset uint64 `json:"offset"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := commitLog.Read(req.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Confirgure the HTTP server with two endpoints:
func main() {
	http.HandleFunc("/produce", handleProduce)
	http.HandleFunc("/consume", handleConsume)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
