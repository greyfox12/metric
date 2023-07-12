package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const LenArr = 10000

type tMetric struct {
	gauge   map[string]float64
	counter map[string]int64
}

var MemMetric tMetric

func main() {

	IPAddress := flag.String("a", "localhost:8080", "Endpoint server IP address host:port")
	flag.Parse()

	r := chi.NewRouter()

	// определяем хендлер, который выводит определённую машину
	r.Route("/", func(r chi.Router) {
		r.Get("/", ListMetricPage)
		r.Get("/value/gauge/{metricName}", OneMetricPage)
		r.Get("/value/counter/{metricName}", OneMetricPage)
		r.Route("/update", func(r chi.Router) {
			r.Post("/gauge/{metricName}/{metricVal}", GaugePage)
			r.Post("/counter/{metricName}/{metricVal}", CounterPage)
			r.Post("/*", ErrorPage)
		})
	})

	log.Fatal(http.ListenAndServe(*IPAddress, r))
	/*	mux := http.NewServeMux()

		mux.HandleFunc(`/update/gauge/`, GaugePage)
		mux.HandleFunc(`/update/counter/`, CounterPage)
		mux.HandleFunc(`/`, ErrorPage)
		MemMetric.gauge = make(map[string]float64, LenArr)
		MemMetric.counter = make(map[string]int64, LenArr)

		err := http.ListenAndServe(`:8080`, mux)
		if err != nil {
			panic(err)
		} */
}
