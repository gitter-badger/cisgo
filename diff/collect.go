package diff

import (
	"log"
)

type Result struct{}

type Collector interface {
	Diff(device *Device) ([]*Result, error)
}

func Collect(collector Collector, device *Device, results chan<- *Result) {
	diffResults, err := collector.Diff(device)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range diffResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%v:\n", result)
	}
}
