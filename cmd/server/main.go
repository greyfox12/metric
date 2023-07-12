package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
)

const LenArr = 10000
const DefServerAdr = "localhost:8080"

type tMetric struct {
	gauge   map[string]float64
	counter map[string]int64
}

var MemMetric tMetric

type Config struct {
	address string `env:"ADDRESS"`
}

var cfg Config

func main() {

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.address == "" {
		cfg.address = DefServerAdr
	}
	fmt.Printf("ServerAdr = %v\n", cfg.address)

	IPAddress := flag.String("a", cfg.address, "Endpoint server IP address host:port")
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
