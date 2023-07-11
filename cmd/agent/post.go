package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) PostCounter(ga map[int]GaugeMetric, co map[int]CounterMetric) int {
	fmt.Printf("Time: %v\n", time.Now().Unix())
	//	fmt.Printf("URL: %v\n", c.url)
	//	fmt.Printf("Len GA= %v\n", len(ga))

	for _, val := range ga {
		s := fmt.Sprintf("%s/update/counter/%s/%v", c.url, val.name, val.Val)
		resp, err := http.Post(s, "Content-Type: text/plain", nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return -1
			//			panic(err)
		}

		defer resp.Body.Close()
		_, _ = ioutil.ReadAll(resp.Body)
	}

	for _, val := range co {
		s := fmt.Sprintf("%s/update/counter/%s/%v", c.url, val.name, val.Val)
		resp, err := http.Post(s, "Content-Type: text/plain", nil)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return -1
			//			panic(err)
		}
		defer resp.Body.Close()
		_, _ = ioutil.ReadAll(resp.Body)

	}
	return 0
}
