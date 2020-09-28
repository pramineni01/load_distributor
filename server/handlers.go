package server

import (
	"fmt"
	"math"
	"net/http"
	"sync"
)

const (
	shareForA float64 = 0.1
	shareForB float64 = 0.2
	shareForC float64 = 0.7
)

type HttpHandlers struct {
	buckets   sync.Map
	totalReqs float64
}

// AddReqFunc adds a request and distributes among buckets
func (h *HttpHandlers) AddReqFunc(w http.ResponseWriter, _ *http.Request) {
	h.totalReqs++

	h.buckets.Store("A", math.Round(h.totalReqs*shareForA))
	h.buckets.Store("B", math.Round(h.totalReqs*shareForB))
	h.buckets.Store("C", math.Round(h.totalReqs*shareForC))

	w.WriteHeader(200)
	w.Write([]byte("Accepted request"))
}

// GetStatsFunc returns the distribution
func (h *HttpHandlers) GetStatsFunc(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	h.buckets.Range(func(k, v interface{}) bool {
		w.Write([]byte(fmt.Sprintf("%s : %d\n", k, int(v.(float64)))))
		return true
	})
}
