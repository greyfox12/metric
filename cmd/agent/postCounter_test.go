package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"
)

//type Client struct {
//	url string
//}

//func NewClient(url string) Client {
//	return Client{url}
//}

func TestClientPostCounter(t *testing.T) {
	ListGauge = make(map[int]GaugeMetric)
	ListCounter = make(map[int]CounterMetric)

	ListGauge[1] = GaugeMetric{"Alloc", gauge(5.5)}
	ListGauge[2] = GaugeMetric{"BuckHashSys", gauge(6)}

	ListCounter[1] = CounterMetric{"PollCount", counter(100)}

	expected := "dummy data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	err := c.PostCounter(ListGauge, ListCounter)
	if err != 0 {
		t.Errorf("expected err to be nil got %v", err)
	}

}
