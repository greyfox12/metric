package main

import (
	"fmt"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) PostCounter(ga map[int]GaugeMetric, co map[int]CounterMetric) int {
	//	fmt.Printf("Time: %v\n", time.Now().Unix())
	fmt.Printf("URL: %v\n", c.url)

	for _, val := range ga {
		//   resp, err := http.Post(url string, contentType string, body io.Reader)
		//		fmt.Printf("Name: %v  Val: %v\n", val.name, val.Val)
		s := fmt.Sprintf("%s/update/gauge/%s/%v", c.url, val.name, val.Val)
		//		fmt.Printf("%s\n", s)
		resp, err := http.Post(s, "Content-Type: text/plain", nil)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

	}

	for _, val := range co {
		//   resp, err := http.Post(url string, contentType string, body io.Reader)
		//		fmt.Printf("Name: %v  Val: %v\n", val.name, val.Val)
		s := fmt.Sprintf("%s/update/counter/%s/%v", c.url, val.name, val.Val)
		//		fmt.Printf("%s\n", s)
		resp, err := http.Post(s, "Content-Type: text/plain", nil)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

	}
	//	fmt.Printf("\n")
	return 0
}
