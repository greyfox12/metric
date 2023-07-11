package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GaugePage(res http.ResponseWriter, req *http.Request) {
	//	fmt.Printf("req.Method1=%v\n", req.Method)
	if MemMetric.gauge == nil {
		MemMetric.gauge = make(map[string]float64, LenArr)
	}

	if req.Method == http.MethodPost {
		aSt := strings.Split(req.URL.Path, "/")

		metricName := aSt[3]
		metricVal := aSt[4]
		//		fmt.Printf("request.URL.Path=%v\n", req.URL.Path)
		//		fmt.Printf("metricVal=%v\n", metricVal)
		if metricName == "" || metricVal == "" {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		// Проверка корректности
		if len(metricName) > 100 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		metricCn, err := strconv.ParseFloat(metricVal, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// контроль длинны карты
		if _, ok := MemMetric.gauge[metricName]; !ok && len(MemMetric.gauge) > LenArr {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// Добавляю новую метрику

		MemMetric.gauge[metricName] = metricCn

		return

	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func CounterPage(res http.ResponseWriter, req *http.Request) {
	//	fmt.Printf("req.Method2=%v\n", req.Method)
	if MemMetric.counter == nil {
		MemMetric.counter = make(map[string]int64, LenArr)
	}

	if req.Method == http.MethodPost {
		aSt := strings.Split(req.URL.Path, "/")

		metricName := aSt[3]
		metricVal := aSt[4]
		//		fmt.Printf("request.URL.Path=%v\n", req.URL.Path)
		//		fmt.Printf("metricVal=%v\n", metricVal)
		if metricName == "" || metricVal == "" {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		// Проверка корректности
		if len(metricName) > 100 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		metricCn, err := strconv.ParseInt(metricVal, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// контроль длинны карты
		if _, ok := MemMetric.counter[metricName]; !ok && len(MemMetric.counter) > LenArr {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// Добавляю новую метрику

		MemMetric.counter[metricName] += metricCn

		//		body += fmt.Sprintf("Длинна: %s\r\n", len(aSt))
		//		body := fmt.Sprintf("Длинна: %s\r\n", len(MemMetric.counter))
		//		res.Write([]byte(body))
		//		io.WriteString(res, req.URL.Path)

		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func ErrorPage(res http.ResponseWriter, req *http.Request) {
	//	fmt.Printf("req.Method3=%v\n", req.Method)
	if req.Method == http.MethodPost {

		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusMethodNotAllowed)
}

func ListMetricPage(res http.ResponseWriter, req *http.Request) {

	var body []string
	//	fmt.Printf("req.Method=%v\n", req.Method)

	for key, val := range MemMetric.gauge {
		body = append(body, fmt.Sprintf("%s = %v", key, val))
	}

	for key, val := range MemMetric.counter {
		body = append(body, fmt.Sprintf("%s = %v", key, val))
	}
	io.WriteString(res, strings.Join(body, "\n"))
}

func OneMetricPage(res http.ResponseWriter, req *http.Request) {
	var Val string
	aSt := strings.Split(req.URL.Path, "/")

	metricName := aSt[3]
	if aSt[2] == "gauge" {
		Val = fmt.Sprintf("%v", MemMetric.gauge[metricName])
	} else {
		Val = fmt.Sprintf("%v", MemMetric.counter[metricName])
	}

	//	fmt.Printf("metricName=%v\n", metricName)

	io.WriteString(res, fmt.Sprintf("%s = %v", metricName, Val))
}
