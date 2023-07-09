package main

import (
	"math/rand"
	"runtime"
	"time"
)

const (
	ServerAdr      = "http://localhost:8080/update/"
	pollInterval   = 2
	reportInterval = 10
)

type counter int64
type gauge float64
type CounterMetric struct {
	name string
	Val  counter
}
type GaugeMetric struct {
	name string
	Val  gauge
}

var ListCounter map[int]CounterMetric
var ListGauge map[int]GaugeMetric

func main() {

	var m runtime.MemStats
	PollCount := counter(0) //Счетчик циклов опроса
	RandomValue := gauge(0)
	ListGauge = make(map[int]GaugeMetric)
	ListCounter = make(map[int]CounterMetric)

	client := NewClient(ServerAdr)

	for PollCount < 1000 {
		runtime.ReadMemStats(&m)
		RandomValue = gauge(rand.Float64())
		ListGauge[1] = GaugeMetric{"Alloc", gauge(m.Alloc)}
		ListGauge[2] = GaugeMetric{"BuckHashSys", gauge(m.BuckHashSys)}
		ListGauge[3] = GaugeMetric{"Frees", gauge(m.Frees)}
		ListGauge[4] = GaugeMetric{"GCCPUFraction", gauge(m.GCCPUFraction)}
		ListGauge[5] = GaugeMetric{"GCSys", gauge(m.GCSys)}
		ListGauge[6] = GaugeMetric{"HeapAlloc", gauge(m.HeapAlloc)}
		ListGauge[7] = GaugeMetric{"HeapIdle", gauge(m.HeapIdle)}
		ListGauge[8] = GaugeMetric{"HeapObjects", gauge(m.HeapObjects)}
		ListGauge[9] = GaugeMetric{"HeapReleased", gauge(m.HeapReleased)}
		ListGauge[10] = GaugeMetric{"HeapSys", gauge(m.HeapSys)}
		ListGauge[11] = GaugeMetric{"LastGC", gauge(m.LastGC)}
		ListGauge[12] = GaugeMetric{"Lookups", gauge(m.Lookups)}
		ListGauge[13] = GaugeMetric{"MCacheInuse", gauge(m.MCacheInuse)}
		ListGauge[14] = GaugeMetric{"MCacheSys", gauge(m.MCacheSys)}
		ListGauge[15] = GaugeMetric{"Mallocs", gauge(m.Mallocs)}
		ListGauge[16] = GaugeMetric{"NextGC", gauge(m.NextGC)}
		ListGauge[17] = GaugeMetric{"NumForcedGC", gauge(m.NumForcedGC)}
		ListGauge[18] = GaugeMetric{"NumGC", gauge(m.NumGC)}
		ListGauge[19] = GaugeMetric{"OtherSys", gauge(m.OtherSys)}
		ListGauge[20] = GaugeMetric{"PauseTotalNs", gauge(m.PauseTotalNs)}
		ListGauge[21] = GaugeMetric{"StackInuse", gauge(m.StackInuse)}
		ListGauge[22] = GaugeMetric{"StackSys", gauge(m.StackSys)}
		ListGauge[23] = GaugeMetric{"Sys", gauge(m.Sys)}
		ListGauge[24] = GaugeMetric{"TotalAlloc", gauge(m.TotalAlloc)}
		ListGauge[25] = GaugeMetric{"RandomValue", gauge(RandomValue)}

		ListCounter[1] = CounterMetric{"PollCount", counter(PollCount)}

		if PollCount%(reportInterval/pollInterval) == 0 {
			//			_ = PostCounter(ListGauge, ListCounter)
			_ = client.PostCounter(ListGauge, ListCounter)
		}

		time.Sleep(pollInterval * time.Second)
		/*		fmt.Printf("%v\n", m.Alloc)
				fmt.Printf("%v\n", m.BuckHashSys)
				fmt.Printf("%v\n", m.Frees)
				fmt.Printf("%v\n", m.GCCPUFraction)
				fmt.Printf("%v\n", m.GCSys)
				fmt.Printf("%v\n", m.HeapAlloc)
				fmt.Printf("%v\n", m.HeapIdle)
				fmt.Printf("%v\n", m.HeapObjects)
				fmt.Printf("%v\n", m.HeapReleased)
				fmt.Printf("%v\n", m.HeapSys)
				fmt.Printf("%v\n", m.LastGC)
				fmt.Printf("%v\n", m.Lookups)
				fmt.Printf("%v\n", m.MCacheInuse)
				fmt.Printf("%v\n", m.MCacheSys)
				fmt.Printf("%v\n", m.Mallocs)
				fmt.Printf("%v\n", m.NextGC)
				fmt.Printf("%v\n", m.NumForcedGC)
				fmt.Printf("%v\n", m.NumGC)
				fmt.Printf("%v\n", m.OtherSys)
				fmt.Printf("%v\n", m.PauseTotalNs)
				fmt.Printf("%v\n", m.StackInuse)
				fmt.Printf("%v\n", m.StackSys)
				fmt.Printf("%v\n", m.Sys)
				fmt.Printf("%v\n", m.TotalAlloc)*/
		//		fmt.Printf("%v\n", PollCount)
		//		fmt.Printf("%v\n", RandomValue)

		PollCount++
	}

}
