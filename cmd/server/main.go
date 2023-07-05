package main

import (
	"net/http"
	"strconv"
	"strings"
)

const LenArr = 10000

type tMetric struct {
	gauge   map[string]float64
	counter map[string]int64
}

var MemMetric tMetric

func gaugePage(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		//		res.WriteHeader(http.StatusOK)
		st := req.URL.Path
		// Проверка корректности
		if len(st) > 100 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		aSt := strings.Split(st, "/")
		if len(aSt) != 4 || len(aSt[2]) == 0 || aSt[1] != "gauge" {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		metricVal, err := strconv.ParseFloat(aSt[3], 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// контроль длинны карты
		if _, ok := MemMetric.gauge[aSt[2]]; ok == false && len(MemMetric.gauge) > LenArr {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// Добавляю новую метрику

		MemMetric.gauge[aSt[2]] = metricVal

		//		body += fmt.Sprintf("Длинна: %s\r\n", len(aSt))
		//		body := fmt.Sprintf("Длинна: %s\r\n", len(MemMetric.gauge))
		//		res.Write([]byte(body))
		//		io.WriteString(res, req.URL.Path)

		return

	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func counterPage(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		//		res.WriteHeader(http.StatusOK)
		st := req.URL.Path
		// Проверка корректности
		if len(st) > 100 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		aSt := strings.Split(st, "/")
		if len(aSt) != 4 || len(aSt[2]) == 0 || aSt[1] != "counter" {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		metricVal, err := strconv.ParseInt(aSt[3], 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// контроль длинны карты
		if _, ok := MemMetric.counter[aSt[2]]; ok == false && len(MemMetric.counter) > LenArr {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// Добавляю новую метрику

		MemMetric.counter[aSt[2]] += metricVal

		//		body += fmt.Sprintf("Длинна: %s\r\n", len(aSt))
		//		body := fmt.Sprintf("Длинна: %s\r\n", len(MemMetric.counter))
		//		res.Write([]byte(body))
		//		io.WriteString(res, req.URL.Path)

		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(`/gauge/`, gaugePage)
	mux.HandleFunc(`/counter/`, counterPage)
	MemMetric.gauge = make(map[string]float64, LenArr)
	MemMetric.counter = make(map[string]int64, LenArr)
	//gaugeMetric.val = m

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
